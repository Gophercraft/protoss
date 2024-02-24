package protoss

import (
	"fmt"

	"google.golang.org/protobuf/types/descriptorpb"
)

func (fd *filedumper) dump_extension(extendee string) (err error) {
	var resolved_extendee string
	resolved_extendee, err = fd.resolve_type_name_in_scope(extendee)
	if err != nil {
		return
	}

	fd.printf("extend %s {\n", resolved_extendee)
	fd.line_indentation++

	for _, extension := range fd.file_descriptor.Extension {
		if extension.GetExtendee() == extendee {

			label := extension.GetLabel()
			label_prefix := ""
			switch label {
			case descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL:
				label_prefix = "optional"
				// all fields are optional in proto3
				if fd.syntax == "proto3" {
					label_prefix = ""
				}
			case descriptorpb.FieldDescriptorProto_LABEL_REPEATED:
				label_prefix = "repeated"
			case descriptorpb.FieldDescriptorProto_LABEL_REQUIRED:
				if fd.syntax == "proto3" {
					return fmt.Errorf("protoss: '%s': a field cannot be required in proto3", fd.file_descriptor.GetName())
				}
				label_prefix = "required"
			default:
				panic(label.String())
			}

			if label_prefix != "" {
				label_prefix += " "
			}

			type_name := extension.GetTypeName()
			if type_name == "" {
				switch extension.GetType() {
				case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
					type_name = "double"
				case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
					type_name = "float"
				case descriptorpb.FieldDescriptorProto_TYPE_INT64:
					type_name = "int64"
				case descriptorpb.FieldDescriptorProto_TYPE_UINT64:
					type_name = "uint64"
				case descriptorpb.FieldDescriptorProto_TYPE_INT32:
					type_name = "int32"
				case descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
					type_name = "fixed64"
				case descriptorpb.FieldDescriptorProto_TYPE_FIXED32:
					type_name = "fixed32"
				case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
					type_name = "bool"
				case descriptorpb.FieldDescriptorProto_TYPE_STRING:
					type_name = "string"
				case descriptorpb.FieldDescriptorProto_TYPE_GROUP:
					type_name = "group"
				case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
					type_name = "message"
				case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
					type_name = "bytes"
				case descriptorpb.FieldDescriptorProto_TYPE_UINT32:
					type_name = "uint32"
				case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
					type_name = "enum"
				case descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
					type_name = "sfixed32"
				case descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
					type_name = "sfixed64"
				case descriptorpb.FieldDescriptorProto_TYPE_SINT32:
					type_name = "sint32"
				case descriptorpb.FieldDescriptorProto_TYPE_SINT64:
					type_name = "sint64"
				default:
					panic(extension.GetType())
				}
			} else {
				type_name, err = fd.resolve_type_name_in_scope(type_name)
				if err != nil {
					return
				}
			}

			number := extension.GetNumber()

			fd.indent()
			fd.printf("%s%s %s = %d", label_prefix, type_name, extension.GetName(), number)
			fd.printf(";\n")
		}
	}

	fd.line_indentation--
	fd.printf("}\n")

	return
}
