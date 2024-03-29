package protoss

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/descriptorpb"
)

type DumpOptions struct {
	// The directory to output decompiled definitions to
	Output string
	// The Go package path to prepend to Protocol Buffers packages
	// when generating "go_package" attribute
	GoPrefix string
	// If true, output additional .json files to assist in ensuring accurate .proto decompilation
	JSON bool
	// If true, don't output descriptors that are in the google/ path
	IgnoreGoogle bool
}

type filedumper struct {
	options          *DumpOptions
	output_file      *os.File
	file_descriptor  *descriptorpb.FileDescriptorProto
	syntax           string
	import_scope     string
	package_scope    string
	line_indentation int
}

func DumpBinary(binary_path string, options *DumpOptions) (err error) {
	var pb_reader *Reader
	filedumper := new(filedumper)
	filedumper.options = options
	pb_reader, err = OpenBinary(binary_path)
	if err != nil {
		return
	}

	defer pb_reader.Close()

	for {
		filedumper.file_descriptor, err = pb_reader.ReadDescriptor()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return
		}

		err = filedumper.dump()
		if err != nil {
			return
		}
	}

	return
}

func (fd *filedumper) printf(fomt string, args ...any) {
	fmt.Fprintf(fd.output_file, fomt, args...)
}

func (fd *filedumper) indent() {
	for i := 0; i < fd.line_indentation; i++ {
		fmt.Fprintf(fd.output_file, "  ")
	}
}

// returns the directory containing a proto definition (forward slash only)
func directory_path(fd_name string) string {
	elements := strings.Split(fd_name, "/")
	if len(elements) == 1 {
		return ""
	}

	joined := strings.Join(elements[:len(elements)-1], "/")
	if !strings.HasSuffix(joined, "/") {
		joined += "/"
	}

	return joined
}

func (fd *filedumper) dump() (err error) {
	fd_name := fd.file_descriptor.GetName()
	if fd.options.IgnoreGoogle {
		if strings.HasPrefix(fd_name, "google/") {
			return nil
		}
	}

	if fd_name == "" {
		err = fmt.Errorf("protoss: dump(): file descriptor has no name")
		return
	}

	if !(strings.HasSuffix(fd_name, ".proto")) {
		err = fmt.Errorf("protoss: dump(): incorrect file descriptor name: %s ", fd_name)
		return
	}

	fd.import_scope = directory_path(fd_name)
	fd.package_scope = fd.file_descriptor.GetPackage()

	output_path := filepath.Join(fd.options.Output, fd_name)

	output_path_directory := filepath.Dir(output_path)

	fmt.Println("mkdir", output_path_directory)

	if err = os.MkdirAll(output_path_directory, 0700); err != nil {
		return
	}

	if fd.options.JSON {
		js := protojson.Format(fd.file_descriptor)

		if err = os.WriteFile(output_path+".json", []byte(js), 0700); err != nil {
			return
		}
	}

	fd.output_file, err = os.OpenFile(output_path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0700)
	if err != nil {
		return
	}

	fd.syntax = "proto2"

	if fd.file_descriptor.Syntax != nil {
		fd.syntax = fd.file_descriptor.GetSyntax()
	}

	fd.printf("// DO NOT EDIT: this file was auto-generated by Gophercraft/protoss\n\n")

	fd.printf("syntax = \"%s\";\n\n", fd.syntax)

	if fd.file_descriptor.Package != nil {
		fd.printf("package %s;\n\n", fd.file_descriptor.GetPackage())

		go_prefix := fd.options.GoPrefix
		if !strings.HasSuffix(go_prefix, "/") {
			go_prefix += "/"
		}

		fd.printf("option go_package = \"%s%s\";\n\n", go_prefix, strings.ReplaceAll(fd.file_descriptor.GetPackage(), ".", "/"))
	}

	for i, dep := range fd.file_descriptor.Dependency {
		index := int32(i)

		prefix := ""

		for _, idx := range fd.file_descriptor.WeakDependency {
			if index == idx {
				prefix = " weak"
				break
			}
		}

		for _, idx := range fd.file_descriptor.PublicDependency {
			if index == idx {
				prefix = " public"
				break
			}
		}

		fmt.Println(dep, "=>", fd.import_scope)

		// if strings.HasPrefix(dep, fd.import_scope) {
		// 	dep = dep[len(fd.import_scope):]
		// }

		fd.printf("import%s \"%s\";\n", prefix, dep)
	}

	if len(fd.file_descriptor.Dependency) > 0 {
		fd.printf("\n")
	}

	fd.printf("\n// Enums\n\n")

	for _, enum_type := range fd.file_descriptor.EnumType {
		if err = fd.dump_enum_definition(enum_type); err != nil {
			return
		}
	}

	fd.printf("\n// Messages\n\n")

	for i, message := range fd.file_descriptor.MessageType {
		if i > 0 {
			fd.printf("\n")
		}

		if err = fd.dump_message_definition(message); err != nil {
			return
		}
	}

	fd.printf("\n// Services\n\n")

	for i, service := range fd.file_descriptor.Service {
		if i > 0 {
			fd.printf("\n")
		}

		if err = fd.dump_service_definition(service); err != nil {
			return
		}
	}

	fd.printf("// Extensions \n\n")

	// look for extendees

	var extendees []string
	for _, extension := range fd.file_descriptor.Extension {
		if !slices.Contains(extendees, extension.GetExtendee()) {
			extendees = append(extendees, extension.GetExtendee())
		}
	}

	for i, extendee := range extendees {
		if i > 0 {
			fd.printf("\n")
		}

		if err = fd.dump_extension(extendee); err != nil {
			return
		}
	}

	return fd.output_file.Close()
}
