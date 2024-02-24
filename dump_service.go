package protoss

import (
	"github.com/Gophercraft/protoss/extensions/bgs/protocol"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (fd *filedumper) dump_service_definition(service_descriptor *descriptorpb.ServiceDescriptorProto) (err error) {
	fd.printf("service %s {\n", service_descriptor.GetName())
	fd.line_indentation++

	for _, rpc := range service_descriptor.Method {
		// decompile basic rpc type information
		method_name := rpc.GetName()
		argument_type := rpc.GetInputType()
		argument_type, err = fd.resolve_type_name_in_scope(argument_type)
		if err != nil {
			return
		}
		return_type := rpc.GetOutputType()
		return_type, err = fd.resolve_type_name_in_scope(return_type)
		if err != nil {
			return
		}

		stream_client_prefix := ""
		if rpc.GetClientStreaming() {
			stream_client_prefix = "stream "
		}
		stream_server_prefix := ""
		if rpc.GetServerStreaming() {
			stream_server_prefix = "stream "
		}

		fd.indent()
		fd.printf("rpc %s(%s%s) returns (%s%s) {",
			method_name,
			stream_client_prefix,
			argument_type,
			stream_server_prefix,
			return_type,
		)

		// decompile BGS options
		if proto.HasExtension(rpc.GetOptions(), protocol.E_MethodOptions) {
			method_options := proto.GetExtension(rpc.GetOptions(), protocol.E_MethodOptions).(*protocol.BGSMethodOptions)
			fd.printf("\n")
			fd.line_indentation++

			fd.indent()
			fd.printf("option (method_options).id = %d;\n", method_options.GetId())

			fd.line_indentation--
			fd.indent()
		}

		// write options
		fd.printf("}\n")
	}
	fd.line_indentation--

	// for _, field := range message_descriptor.Field {
	// 	type_name := field.GetTypeName()
	// 	// if type_name == "" {
	// 	// 	// no type name
	// 	// 	type_name = field.Type.
	// 	// }

	// 	// if field.OneofIndex != nil {
	// 	// 	oneof := message_def.OneofDecl[*field.OneofIndex]
	// 	// 	oneof.Options.
	// 	// }
	// 	number := field.GetNumber()

	// 	fd.printf("\t%s %s = %d;", type_name, field.GetName(), number)
	// }
	fd.printf("}\n")
	return
}
