package protoss

import (
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/protobuf/types/descriptorpb"
)

func (fd *filedumper) dump_message_definition(message_descriptor *descriptorpb.DescriptorProto) (err error) {
	// original_scope := fd.package_scope

	// fd.package_scope += fmt.Sprintf(".%s", message_descriptor.GetName())

	// Open message definition
	fd.indent()
	fd.printf("message %s {\n", message_descriptor.GetName())

	fd.line_indentation++

	// Encode enums
	for _, enum_type := range message_descriptor.EnumType {
		if err = fd.dump_enum_definition(enum_type); err != nil {
			return
		}
	}

	// Encode nested types
	for _, nested_type := range message_descriptor.NestedType {
		if err = fd.dump_message_definition(nested_type); err != nil {
			return
		}
	}

	was_writing_oneof := false

	for _, field := range message_descriptor.Field {
		is_oneof := field.OneofIndex != nil

		label_prefix := ""

		if is_oneof {
			if !was_writing_oneof {
				fd.indent()
				fd.printf("oneof %s {\n", *message_descriptor.OneofDecl[field.GetOneofIndex()].Name)
				fd.line_indentation++
				was_writing_oneof = true
			}
		} else {
			if was_writing_oneof {
				fd.line_indentation--
				fd.indent()
				fd.printf("}\n")
				was_writing_oneof = false
			}

			label := field.GetLabel()
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
		}

		type_name := field.GetTypeName()
		if type_name == "" {
			switch field.GetType() {
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
				panic(field.GetType())
			}
		} else {
			type_name, err = fd.resolve_type_name_in_scope(type_name)
			if err != nil {
				return
			}
		}
		// if type_name == "" {
		// 	// no type name
		// 	type_name = field.Type.
		// }

		// if field.OneofIndex != nil {
		// 	oneof := message_def.OneofDecl[*field.OneofIndex]
		// 	oneof.Options.
		// }
		number := field.GetNumber()

		fd.indent()
		fd.printf("%s%s %s = %d", label_prefix, type_name, field.GetName(), number)
		var options []string
		if field.Options != nil {
			if field.Options.GetDeprecated() {
				options = append(options, fmt.Sprintf("deprecated = %t", field.Options.GetDeprecated()))
			}
		}

		if field.DefaultValue != nil {
			default_string := field.GetDefaultValue()
			if field.GetType() == descriptorpb.FieldDescriptorProto_TYPE_STRING {
				default_string = strconv.Quote(default_string)
			}
			options = append(options, fmt.Sprintf("default = %s", default_string))
		}

		if len(options) > 0 {
			fd.printf(" [%s]", strings.Join(options, ", "))
		}

		fd.printf(";\n")
	}

	if was_writing_oneof {
		fd.line_indentation--
		fd.indent()
		fd.printf("}\n")
		was_writing_oneof = false
	}

	for _, extension_range := range message_descriptor.GetExtensionRange() {
		upper_range := extension_range.GetEnd()
		upper_string := ""
		if upper_range == 536870912 {
			upper_string = "max"
		} else {
			upper_string = fmt.Sprintf("%d", upper_range)
		}
		fd.indent()
		fd.printf("extensions %d to %s;\n", extension_range.GetStart(), upper_string)
	}

	fd.line_indentation--

	fd.indent()
	fd.printf("}\n")

	// fd.package_scope = original_scope

	return
}
