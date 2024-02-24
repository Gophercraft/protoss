package protoss

import "google.golang.org/protobuf/types/descriptorpb"

func (fd *filedumper) dump_enum_definition(enum_descriptor *descriptorpb.EnumDescriptorProto) (err error) {
	fd.indent()
	fd.printf("enum %s {\n", enum_descriptor.GetName())

	fd.line_indentation++

	if enum_descriptor.GetOptions() != nil {
		any_options := false

		options := enum_descriptor.GetOptions()
		if options.AllowAlias != nil {
			fd.indent()
			fd.printf("option allow_alias = %t;\n", options.GetAllowAlias())
			any_options = true
		}

		if options.Deprecated != nil {
			fd.indent()
			fd.printf("option deprecated = %t;\n", options.GetDeprecated())
			any_options = true
		}

		if any_options {
			fd.printf("\n")
		}
	}

	for _, enum_value := range enum_descriptor.Value {
		fd.indent()
		fd.printf("%s = %d", enum_value.GetName(), enum_value.GetNumber())
		// if enum_value.Options != nil {
		// 	if enum_value.
		// }
		fd.printf(";\n")
	}
	fd.line_indentation--

	fd.indent()
	fd.printf("}\n")

	return
}
