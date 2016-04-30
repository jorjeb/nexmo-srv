// Code generated by protoc-gen-go.
// source: github.com/jorjeb/nexmo-srv/proto/sms/sms.proto
// DO NOT EDIT!

/*
Package proto_sms is a generated protocol buffer package.

It is generated from these files:
	github.com/jorjeb/nexmo-srv/proto/sms/sms.proto

It has these top-level messages:
	SendRequest
	MessagePart
	SendResponse
*/
package proto_sms

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type SendRequest struct {
	From       string `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To         string `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Text       string `protobuf:"bytes,3,opt,name=text" json:"text,omitempty"`
	RegionCode string `protobuf:"bytes,4,opt,name=regionCode" json:"regionCode,omitempty"`
}

func (m *SendRequest) Reset()                    { *m = SendRequest{} }
func (m *SendRequest) String() string            { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()               {}
func (*SendRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type MessagePart struct {
	Status           string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	MessageId        string `protobuf:"bytes,2,opt,name=messageId" json:"messageId,omitempty"`
	To               string `protobuf:"bytes,3,opt,name=to" json:"to,omitempty"`
	ClientRef        string `protobuf:"bytes,4,opt,name=clientRef" json:"clientRef,omitempty"`
	RemainingBalance string `protobuf:"bytes,5,opt,name=remainingBalance" json:"remainingBalance,omitempty"`
	MessagePrice     string `protobuf:"bytes,6,opt,name=messagePrice" json:"messagePrice,omitempty"`
	Network          string `protobuf:"bytes,7,opt,name=network" json:"network,omitempty"`
	ErrorText        string `protobuf:"bytes,8,opt,name=errorText" json:"errorText,omitempty"`
}

func (m *MessagePart) Reset()                    { *m = MessagePart{} }
func (m *MessagePart) String() string            { return proto.CompactTextString(m) }
func (*MessagePart) ProtoMessage()               {}
func (*MessagePart) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SendResponse struct {
	MessageCount string         `protobuf:"bytes,1,opt,name=messageCount" json:"messageCount,omitempty"`
	Messages     []*MessagePart `protobuf:"bytes,2,rep,name=messages" json:"messages,omitempty"`
}

func (m *SendResponse) Reset()                    { *m = SendResponse{} }
func (m *SendResponse) String() string            { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()               {}
func (*SendResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SendResponse) GetMessages() []*MessagePart {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*SendRequest)(nil), "proto.sms.SendRequest")
	proto.RegisterType((*MessagePart)(nil), "proto.sms.MessagePart")
	proto.RegisterType((*SendResponse)(nil), "proto.sms.SendResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for SMS service

type SMSClient interface {
	Send(ctx context.Context, in *SendRequest, opts ...client.CallOption) (*SendResponse, error)
}

type sMSClient struct {
	c           client.Client
	serviceName string
}

func NewSMSClient(serviceName string, c client.Client) SMSClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "proto.sms"
	}
	return &sMSClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *sMSClient) Send(ctx context.Context, in *SendRequest, opts ...client.CallOption) (*SendResponse, error) {
	req := c.c.NewRequest(c.serviceName, "SMS.Send", in)
	out := new(SendResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SMS service

type SMSHandler interface {
	Send(context.Context, *SendRequest, *SendResponse) error
}

func RegisterSMSHandler(s server.Server, hdlr SMSHandler) {
	s.Handle(s.NewHandler(&SMS{hdlr}))
}

type SMS struct {
	SMSHandler
}

func (h *SMS) Send(ctx context.Context, in *SendRequest, out *SendResponse) error {
	return h.SMSHandler.Send(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x6b, 0xea, 0x40,
	0x14, 0xc5, 0x9f, 0xc6, 0xe7, 0x9f, 0xab, 0x3c, 0x1e, 0x77, 0xe1, 0x1b, 0x1e, 0xa5, 0x94, 0xac,
	0xa4, 0xd0, 0x04, 0xec, 0xaa, 0xbb, 0x52, 0x57, 0x5d, 0x08, 0x12, 0xfb, 0x05, 0xa2, 0x5e, 0xd3,
	0x58, 0x33, 0xd7, 0xce, 0x4c, 0x5a, 0xbf, 0x7a, 0x77, 0x4d, 0x26, 0xd3, 0x18, 0x71, 0x21, 0xde,
	0xfb, 0xbb, 0x27, 0x9c, 0xc3, 0x19, 0x08, 0x93, 0xd4, 0xbc, 0xe6, 0xab, 0x60, 0xcd, 0x59, 0xb8,
	0x63, 0xb5, 0xa3, 0x55, 0x28, 0xe9, 0x98, 0xf1, 0x9d, 0x56, 0x1f, 0xe1, 0x41, 0xb1, 0xe1, 0x50,
	0x67, 0xba, 0xfc, 0x05, 0x76, 0xc3, 0x81, 0xfd, 0x0b, 0x0a, 0xe0, 0x13, 0x0c, 0x97, 0x24, 0x37,
	0x11, 0xbd, 0xe7, 0xa4, 0x0d, 0x22, 0x74, 0xb6, 0x8a, 0x33, 0xd1, 0xba, 0x69, 0x4d, 0x06, 0x91,
	0x9d, 0xf1, 0x0f, 0xb4, 0x0d, 0x8b, 0xb6, 0x25, 0xc5, 0x54, 0x6a, 0x0c, 0x1d, 0x8d, 0xf0, 0x2a,
	0x4d, 0x39, 0xe3, 0x35, 0x80, 0xa2, 0x24, 0x65, 0x39, 0xe3, 0x0d, 0x89, 0x8e, 0xbd, 0x34, 0x88,
	0xff, 0xd5, 0x82, 0xe1, 0x9c, 0xb4, 0x8e, 0x13, 0x5a, 0xc4, 0xca, 0xe0, 0x18, 0xba, 0xda, 0xc4,
	0x26, 0xd7, 0xce, 0xc9, 0x6d, 0x78, 0x05, 0x83, 0xac, 0x92, 0x3d, 0x6f, 0x9c, 0xe5, 0x09, 0xb8,
	0x24, 0x5e, 0x9d, 0xa4, 0x50, 0xaf, 0xf7, 0x29, 0x49, 0x13, 0xd1, 0xd6, 0x99, 0x9e, 0x00, 0xde,
	0xc2, 0x5f, 0x45, 0x59, 0x9c, 0xca, 0x54, 0x26, 0x4f, 0xf1, 0x3e, 0x96, 0x6b, 0x12, 0xbf, 0xad,
	0xe8, 0x82, 0xa3, 0x0f, 0x23, 0x67, 0xb3, 0x50, 0x69, 0xa1, 0xeb, 0x5a, 0xdd, 0x19, 0x43, 0x01,
	0x3d, 0x49, 0xe6, 0x93, 0xd5, 0x9b, 0xe8, 0xd9, 0xf3, 0xcf, 0x5a, 0xe6, 0x20, 0xa5, 0x58, 0xbd,
	0x94, 0xb5, 0xf4, 0xab, 0x1c, 0x35, 0xf0, 0xb7, 0x30, 0xaa, 0x2a, 0xd6, 0x07, 0x96, 0xba, 0xe9,
	0x35, 0xe3, 0x5c, 0x1a, 0xd7, 0xc0, 0x19, 0xc3, 0x29, 0xf4, 0xdd, 0xae, 0x8b, 0x1a, 0xbc, 0xc9,
	0x70, 0x3a, 0x0e, 0xea, 0x47, 0x0b, 0x1a, 0x4d, 0x46, 0xb5, 0x6e, 0xfa, 0x08, 0xde, 0x72, 0xbe,
	0xc4, 0x07, 0xe8, 0x94, 0x76, 0xd8, 0xfc, 0xa0, 0xf1, 0xc4, 0xff, 0xff, 0x5d, 0xf0, 0x2a, 0x97,
	0xff, 0x6b, 0xd5, 0xb5, 0x97, 0xfb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x67, 0x00, 0x39, 0xd4,
	0x51, 0x02, 0x00, 0x00,
}
