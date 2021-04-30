// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package broker is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Message
*/
package broker

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Header map[string]string `protobuf:"bytes,1,rep,name=Header" json:"Header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body   []byte            `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "broker.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x2a, 0xca, 0xcf, 0x4e,
	0x2d, 0x52, 0xea, 0x65, 0xe4, 0x62, 0xf7, 0x85, 0xc8, 0x08, 0x19, 0x73, 0xb1, 0x79, 0xa4, 0x26,
	0xa6, 0xa4, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x49, 0xeb, 0x41, 0x14, 0xe9, 0x41,
	0x15, 0xe8, 0x41, 0x64, 0x5d, 0xf3, 0x4a, 0x8a, 0x2a, 0x83, 0xa0, 0x4a, 0x85, 0x84, 0xb8, 0x58,
	0x9c, 0xf2, 0x53, 0x2a, 0x25, 0x98, 0x14, 0x18, 0x35, 0x78, 0x82, 0xc0, 0x6c, 0x29, 0x4b, 0x2e,
	0x6e, 0x24, 0xa5, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0x2a, 0x58, 0x17, 0x67, 0x10,
	0x84, 0x63, 0xc5, 0x64, 0xc1, 0x98, 0xc4, 0x06, 0x76, 0x9e, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x74, 0xa1, 0xc2, 0xc8, 0xaf, 0x00, 0x00, 0x00,
}