// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package messages is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Echo
	Response
*/
package messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import (
	math "math"
	"github.com/AsynkronIT/protoactor-go/actor"
)
//import actor "."

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// this is the message the actor on node 1 will send to the remote actor on node 2
type Echo struct {
	Sender  *actor.PID `protobuf:"bytes,1,opt,name=Sender" json:"Sender,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
}

func (m *Echo) Reset()                    { *m = Echo{} }
func (m *Echo) String() string            { return proto.CompactTextString(m) }
func (*Echo) ProtoMessage()               {}
func (*Echo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Echo) GetSender() *actor.PID {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Echo) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// this is the message the remote actor should reply with
type Response struct {
	SomeValue string `protobuf:"bytes,1,opt,name=SomeValue" json:"SomeValue,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetSomeValue() string {
	if m != nil {
		return m.SomeValue
	}
	return ""
}

func init() {
	proto.RegisterType((*Echo)(nil), "messages.Echo")
	proto.RegisterType((*Response)(nil), "messages.Response")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0x72, 0x8b, 0xa5, 0x78,
	0xc0, 0x02, 0xc5, 0x10, 0x71, 0x25, 0x17, 0x2e, 0x16, 0xd7, 0xe4, 0x8c, 0x7c, 0x21, 0x25, 0x2e,
	0xb6, 0xe0, 0xd4, 0xbc, 0x94, 0xd4, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x2e, 0xbd,
	0xc4, 0xe4, 0x92, 0xfc, 0x22, 0xbd, 0x00, 0x4f, 0x97, 0x20, 0xa8, 0x8c, 0x90, 0x04, 0x17, 0xbb,
	0x2f, 0xc4, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x49, 0x83, 0x8b, 0x23,
	0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x86, 0x8b, 0x33, 0x38, 0x3f, 0x37, 0x35,
	0x2c, 0x31, 0xa7, 0x34, 0x15, 0x6c, 0x18, 0x67, 0x10, 0x42, 0x20, 0x89, 0x0d, 0x6c, 0xad, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x5e, 0x8f, 0x42, 0x9f, 0x00, 0x00, 0x00,
}
