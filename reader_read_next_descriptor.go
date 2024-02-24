package protoss

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

var (
	dot_proto_string      = []byte(".proto")
	devel_string          = []byte("devel")
	dot_protodevel_string = append(dot_proto_string, devel_string...)
)

// Finds the data of a new descriptor
func (reader *Reader) read_next_descriptor() (descriptor_bytes []byte, err error) {
	for {
		if reader.offset >= len(reader.binary) {
			err = io.EOF
			return
		}

		// Read subset of binary after current offset
		// we do this to look for ".proto"
		working_slice := reader.binary[reader.offset:]

		// Find offset to .proto extension
		working_extension_offset := bytes.Index(working_slice, dot_proto_string)

		// If could not find, there's no more descriptors in the binary
		if working_extension_offset == -1 {
			err = io.EOF
			return
		}

		// The absolute position of a descriptor's file (where ".proto" string begins in file)
		extension_position_in_bin := reader.offset + working_extension_offset

		// Get the end of the descriptor by adding the len of ".proto"
		end_of_descriptor_name_in_bin := extension_position_in_bin + len(dot_proto_string)

		// See if the extension is actually .protodevel
		if end_of_descriptor_name_in_bin+len(devel_string) <= len(reader.binary) {
			if bytes.Equal(reader.binary[extension_position_in_bin:extension_position_in_bin+len(dot_proto_string)+len(devel_string)], dot_protodevel_string) {
				end_of_descriptor_name_in_bin += len(devel_string)
			}
		}

		// Find position of string size marker byte
		size_marker_offset := bytes.LastIndexByte(reader.binary[:end_of_descriptor_name_in_bin], byte(marker_byte))

		if reader.binary[size_marker_offset-1] == marker_byte == ((end_of_descriptor_name_in_bin - size_marker_offset - 1) != 0) {
			size_marker_offset--
		}

		// couldn't find marker byte
		if size_marker_offset == -1 {
			reader.offset = end_of_descriptor_name_in_bin
			continue
		}

		// read variable length string size after marker byte
		size_uvarint, size_uvarint_offset := binary.Uvarint(reader.binary[size_marker_offset+1:])

		descriptor_name_start := size_marker_offset + 1 + size_uvarint_offset
		descriptor_name_end := descriptor_name_start + int(size_uvarint)
		// descriptor_name := string(reader.binary[descriptor_name_start:descriptor_name_end])

		// fmt.Println(descriptor_name)

		if descriptor_name_end != end_of_descriptor_name_in_bin {
			// this is not a protobuf string, move to after the occurence of ".proto(devel)"
			reader.offset = end_of_descriptor_name_in_bin
			continue
		}

		// see if this string is really the beginning of a descriptor
		if bytes.IndexByte(descriptor_markers, reader.binary[end_of_descriptor_name_in_bin]) == -1 {
			reader.offset = end_of_descriptor_name_in_bin
			continue
		}

		markers := descriptor_markers

		cursor := descriptor_name_end

		// Loop through descriptor message, reading fields and advancing the read cursor
		for cursor < len(reader.binary) && bytes.IndexByte(markers, reader.binary[cursor]) != -1 {
			field_marker := reader.binary[cursor]
			wire_type := (reader.binary[cursor] & 0b111)

			cursor++

			markers = markers[bytes.IndexByte(markers, field_marker):]

			switch wire_type {
			// varint:
			case 0:
				_, decoded := binary.Uvarint(reader.binary[cursor:])
				cursor += decoded
			case 1:
				cursor += 8
			case 2:
				size, decoded := binary.Uvarint(reader.binary[cursor:])
				cursor += decoded + int(size)
			case 5:
				cursor += 4
			default:
				return nil, fmt.Errorf("invalid wire_type %d at %d", wire_type, cursor-1)
			}
		}

		descriptor_start := size_marker_offset
		descriptor_end := cursor

		descriptor_bytes = reader.binary[descriptor_start:descriptor_end]

		reader.offset = descriptor_end

		return
	}
}

func (reader *Reader) ReadDescriptorBytes() ([]byte, error) {
	return reader.read_next_descriptor()
}

func (reader *Reader) ReadDescriptor() (file_descriptor *descriptorpb.FileDescriptorProto, err error) {
	var descriptor_bytes []byte
	descriptor_bytes, err = reader.ReadDescriptorBytes()
	if err != nil {
		return
	}

	file_descriptor = new(descriptorpb.FileDescriptorProto)
	err = proto.Unmarshal(descriptor_bytes, file_descriptor)
	return
}
