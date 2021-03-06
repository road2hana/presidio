// Code generated by protoc-gen-go. DO NOT EDIT.
// source: anonymize.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AnonymizeApiRequest represents the request to the API HTTP service
type AnonymizeApiRequest struct {
	// The text to anonymize
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// The analyze template id - anonymization is done according to analyzing results.
	// One of analyzeTemplateId or analyzeTemplate have to be configured.
	AnalyzeTemplateId string `protobuf:"bytes,2,opt,name=analyzeTemplateId,proto3" json:"analyzeTemplateId,omitempty"`
	// The anonymize template id - represents the anonymize configuration, which fields to anonymize and how.
	AnonymizeTemplateId string `protobuf:"bytes,3,opt,name=anonymizeTemplateId,proto3" json:"anonymizeTemplateId,omitempty"`
	// Optional parameter for running the analyzer without creating a template.
	AnalyzeTemplate *AnalyzeTemplate `protobuf:"bytes,4,opt,name=analyzeTemplate,proto3" json:"analyzeTemplate,omitempty"`
	// Optional parameter for running the anonymizer without creating a template.
	AnonymizeTemplate *AnonymizeTemplate `protobuf:"bytes,5,opt,name=anonymizeTemplate,proto3" json:"anonymizeTemplate,omitempty"`
	// Optional parameter for anonymizing text for a given context
	AnonymizeTextContext *AnonymizeTextContext `protobuf:"bytes,6,opt,name=anonymizeTextContext,proto3" json:"anonymizeTextContext,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AnonymizeApiRequest) Reset()         { *m = AnonymizeApiRequest{} }
func (m *AnonymizeApiRequest) String() string { return proto.CompactTextString(m) }
func (*AnonymizeApiRequest) ProtoMessage()    {}
func (*AnonymizeApiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_anonymize_5f4527c3d9864e50, []int{0}
}
func (m *AnonymizeApiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnonymizeApiRequest.Unmarshal(m, b)
}
func (m *AnonymizeApiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnonymizeApiRequest.Marshal(b, m, deterministic)
}
func (dst *AnonymizeApiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnonymizeApiRequest.Merge(dst, src)
}
func (m *AnonymizeApiRequest) XXX_Size() int {
	return xxx_messageInfo_AnonymizeApiRequest.Size(m)
}
func (m *AnonymizeApiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnonymizeApiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnonymizeApiRequest proto.InternalMessageInfo

func (m *AnonymizeApiRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *AnonymizeApiRequest) GetAnalyzeTemplateId() string {
	if m != nil {
		return m.AnalyzeTemplateId
	}
	return ""
}

func (m *AnonymizeApiRequest) GetAnonymizeTemplateId() string {
	if m != nil {
		return m.AnonymizeTemplateId
	}
	return ""
}

func (m *AnonymizeApiRequest) GetAnalyzeTemplate() *AnalyzeTemplate {
	if m != nil {
		return m.AnalyzeTemplate
	}
	return nil
}

func (m *AnonymizeApiRequest) GetAnonymizeTemplate() *AnonymizeTemplate {
	if m != nil {
		return m.AnonymizeTemplate
	}
	return nil
}

func (m *AnonymizeApiRequest) GetAnonymizeTextContext() *AnonymizeTextContext {
	if m != nil {
		return m.AnonymizeTextContext
	}
	return nil
}

// Anonymize text metadata
type AnonymizeTextContext struct {
	// Text Create Date
	CreateDate           string   `protobuf:"bytes,1,opt,name=CreateDate,proto3" json:"CreateDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnonymizeTextContext) Reset()         { *m = AnonymizeTextContext{} }
func (m *AnonymizeTextContext) String() string { return proto.CompactTextString(m) }
func (*AnonymizeTextContext) ProtoMessage()    {}
func (*AnonymizeTextContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_anonymize_5f4527c3d9864e50, []int{1}
}
func (m *AnonymizeTextContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnonymizeTextContext.Unmarshal(m, b)
}
func (m *AnonymizeTextContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnonymizeTextContext.Marshal(b, m, deterministic)
}
func (dst *AnonymizeTextContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnonymizeTextContext.Merge(dst, src)
}
func (m *AnonymizeTextContext) XXX_Size() int {
	return xxx_messageInfo_AnonymizeTextContext.Size(m)
}
func (m *AnonymizeTextContext) XXX_DiscardUnknown() {
	xxx_messageInfo_AnonymizeTextContext.DiscardUnknown(m)
}

var xxx_messageInfo_AnonymizeTextContext proto.InternalMessageInfo

func (m *AnonymizeTextContext) GetCreateDate() string {
	if m != nil {
		return m.CreateDate
	}
	return ""
}

// AnonymizeRequest represents the request to the anonymize service via GRPC
type AnonymizeRequest struct {
	// The text to anonymize
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// The anonymize template represent the anonymize configuration, which fields to anonymize and how
	Template *AnonymizeTemplate `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	// The analyze result containing the field type and location of the sensetive data to be anonymized.
	AnalyzeResults []*AnalyzeResult `protobuf:"bytes,3,rep,name=analyzeResults,proto3" json:"analyzeResults,omitempty"`
	// The context of the anonymize text
	AnonymizeTextContext *AnonymizeTextContext `protobuf:"bytes,4,opt,name=anonymizeTextContext,proto3" json:"anonymizeTextContext,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AnonymizeRequest) Reset()         { *m = AnonymizeRequest{} }
func (m *AnonymizeRequest) String() string { return proto.CompactTextString(m) }
func (*AnonymizeRequest) ProtoMessage()    {}
func (*AnonymizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_anonymize_5f4527c3d9864e50, []int{2}
}
func (m *AnonymizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnonymizeRequest.Unmarshal(m, b)
}
func (m *AnonymizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnonymizeRequest.Marshal(b, m, deterministic)
}
func (dst *AnonymizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnonymizeRequest.Merge(dst, src)
}
func (m *AnonymizeRequest) XXX_Size() int {
	return xxx_messageInfo_AnonymizeRequest.Size(m)
}
func (m *AnonymizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnonymizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnonymizeRequest proto.InternalMessageInfo

func (m *AnonymizeRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *AnonymizeRequest) GetTemplate() *AnonymizeTemplate {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *AnonymizeRequest) GetAnalyzeResults() []*AnalyzeResult {
	if m != nil {
		return m.AnalyzeResults
	}
	return nil
}

func (m *AnonymizeRequest) GetAnonymizeTextContext() *AnonymizeTextContext {
	if m != nil {
		return m.AnonymizeTextContext
	}
	return nil
}

// AnonymizeResponse represents the anonymize service response
type AnonymizeResponse struct {
	// The text with the senstive fields anonymized
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnonymizeResponse) Reset()         { *m = AnonymizeResponse{} }
func (m *AnonymizeResponse) String() string { return proto.CompactTextString(m) }
func (*AnonymizeResponse) ProtoMessage()    {}
func (*AnonymizeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_anonymize_5f4527c3d9864e50, []int{3}
}
func (m *AnonymizeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnonymizeResponse.Unmarshal(m, b)
}
func (m *AnonymizeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnonymizeResponse.Marshal(b, m, deterministic)
}
func (dst *AnonymizeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnonymizeResponse.Merge(dst, src)
}
func (m *AnonymizeResponse) XXX_Size() int {
	return xxx_messageInfo_AnonymizeResponse.Size(m)
}
func (m *AnonymizeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AnonymizeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AnonymizeResponse proto.InternalMessageInfo

func (m *AnonymizeResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*AnonymizeApiRequest)(nil), "types.AnonymizeApiRequest")
	proto.RegisterType((*AnonymizeTextContext)(nil), "types.AnonymizeTextContext")
	proto.RegisterType((*AnonymizeRequest)(nil), "types.AnonymizeRequest")
	proto.RegisterType((*AnonymizeResponse)(nil), "types.AnonymizeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AnonymizeServiceClient is the client API for AnonymizeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnonymizeServiceClient interface {
	// Apply method will execute on the given request and return the anonymize response with the sensitive text anonymized
	Apply(ctx context.Context, in *AnonymizeRequest, opts ...grpc.CallOption) (*AnonymizeResponse, error)
}

type anonymizeServiceClient struct {
	cc *grpc.ClientConn
}

func NewAnonymizeServiceClient(cc *grpc.ClientConn) AnonymizeServiceClient {
	return &anonymizeServiceClient{cc}
}

func (c *anonymizeServiceClient) Apply(ctx context.Context, in *AnonymizeRequest, opts ...grpc.CallOption) (*AnonymizeResponse, error) {
	out := new(AnonymizeResponse)
	err := c.cc.Invoke(ctx, "/types.AnonymizeService/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnonymizeServiceServer is the server API for AnonymizeService service.
type AnonymizeServiceServer interface {
	// Apply method will execute on the given request and return the anonymize response with the sensitive text anonymized
	Apply(context.Context, *AnonymizeRequest) (*AnonymizeResponse, error)
}

func RegisterAnonymizeServiceServer(s *grpc.Server, srv AnonymizeServiceServer) {
	s.RegisterService(&_AnonymizeService_serviceDesc, srv)
}

func _AnonymizeService_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnonymizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnonymizeServiceServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.AnonymizeService/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnonymizeServiceServer).Apply(ctx, req.(*AnonymizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AnonymizeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.AnonymizeService",
	HandlerType: (*AnonymizeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Apply",
			Handler:    _AnonymizeService_Apply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "anonymize.proto",
}

func init() { proto.RegisterFile("anonymize.proto", fileDescriptor_anonymize_5f4527c3d9864e50) }

var fileDescriptor_anonymize_5f4527c3d9864e50 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcb, 0x4a, 0xc3, 0x40,
	0x14, 0x35, 0x7d, 0xa1, 0xb7, 0x52, 0xed, 0xb4, 0xe8, 0x50, 0x41, 0x4a, 0x36, 0x76, 0x21, 0x45,
	0xa2, 0xb8, 0x72, 0x61, 0xa8, 0x08, 0xae, 0x94, 0xd1, 0x1f, 0x88, 0xf5, 0x2e, 0x02, 0xc9, 0xcc,
	0x98, 0x99, 0x4a, 0xd3, 0xef, 0x75, 0xe7, 0x4f, 0x48, 0x26, 0x69, 0xc8, 0xcb, 0x2e, 0xdc, 0x85,
	0x7b, 0x1e, 0x73, 0xcf, 0xb9, 0x04, 0x8e, 0x3c, 0x2e, 0x78, 0x1c, 0xfa, 0x1b, 0x9c, 0xcb, 0x48,
	0x68, 0x41, 0xba, 0x3a, 0x96, 0xa8, 0x26, 0x87, 0x4b, 0x11, 0x86, 0x82, 0xa7, 0xc3, 0xc9, 0x40,
	0x63, 0x28, 0x03, 0x4f, 0x67, 0x24, 0xfb, 0xbb, 0x05, 0x23, 0x77, 0x2b, 0x74, 0xa5, 0xcf, 0xf0,
	0x73, 0x85, 0x4a, 0x13, 0x02, 0x1d, 0x8d, 0x6b, 0x4d, 0xad, 0xa9, 0x35, 0x3b, 0x60, 0xe6, 0x9b,
	0x5c, 0xc2, 0xd0, 0xe3, 0x5e, 0x10, 0x6f, 0xf0, 0x2d, 0x33, 0x79, 0xfa, 0xa0, 0x2d, 0x43, 0xa8,
	0x03, 0xe4, 0x0a, 0x46, 0xf9, 0x46, 0x05, 0x7e, 0xdb, 0xf0, 0x9b, 0x20, 0x72, 0x9f, 0x64, 0x28,
	0xd9, 0xd0, 0xce, 0xd4, 0x9a, 0xf5, 0x9d, 0x93, 0xb9, 0x89, 0x32, 0x77, 0xcb, 0x28, 0xab, 0xd2,
	0xc9, 0x63, 0xb2, 0x61, 0xc5, 0x98, 0x76, 0x8d, 0x07, 0xcd, 0x3d, 0x2a, 0x38, 0xab, 0x4b, 0xc8,
	0x33, 0x8c, 0x0b, 0xc3, 0xb5, 0x5e, 0x08, 0x6e, 0xda, 0xe8, 0x19, 0xab, 0xb3, 0xba, 0x55, 0x4e,
	0x61, 0x8d, 0x42, 0xfb, 0x16, 0xc6, 0x4d, 0x6c, 0x72, 0x0e, 0xb0, 0x88, 0xd0, 0xd3, 0xf8, 0x90,
	0x6c, 0x9a, 0x96, 0x5d, 0x98, 0xd8, 0x3f, 0x16, 0x1c, 0xe7, 0xc2, 0x5d, 0xb7, 0xb9, 0x81, 0xfd,
	0xed, 0x65, 0xcd, 0x49, 0x76, 0x05, 0xce, 0x99, 0xe4, 0x0e, 0x06, 0x59, 0x85, 0x0c, 0xd5, 0x2a,
	0xd0, 0x8a, 0xb6, 0xa7, 0xed, 0x59, 0xdf, 0x19, 0x97, 0x0b, 0x4f, 0x41, 0x56, 0xe1, 0xfe, 0xd9,
	0x52, 0xe7, 0xbf, 0x2d, 0x5d, 0xc0, 0xb0, 0x10, 0x56, 0x49, 0xc1, 0x15, 0x36, 0xa5, 0x75, 0x5e,
	0x0a, 0xad, 0xbc, 0x62, 0xf4, 0xe5, 0x2f, 0x93, 0x2c, 0x5d, 0x57, 0xca, 0x20, 0x26, 0xa7, 0xd5,
	0x87, 0xb3, 0xde, 0x26, 0xb4, 0x0e, 0xa4, 0x6f, 0xd8, 0x7b, 0xef, 0x3d, 0xf3, 0x3b, 0x5c, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xa5, 0x5c, 0x68, 0x46, 0x03, 0x00, 0x00,
}
