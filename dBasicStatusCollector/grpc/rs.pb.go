// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rs.proto

package robotStatus

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The request message containing the tag for the server to log, not to query.
type RobotStatusRequest struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RobotStatusRequest) Reset()         { *m = RobotStatusRequest{} }
func (m *RobotStatusRequest) String() string { return proto.CompactTextString(m) }
func (*RobotStatusRequest) ProtoMessage()    {}
func (*RobotStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0ec5498159ba86f, []int{0}
}

func (m *RobotStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RobotStatusRequest.Unmarshal(m, b)
}
func (m *RobotStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RobotStatusRequest.Marshal(b, m, deterministic)
}
func (m *RobotStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RobotStatusRequest.Merge(m, src)
}
func (m *RobotStatusRequest) XXX_Size() int {
	return xxx_messageInfo_RobotStatusRequest.Size(m)
}
func (m *RobotStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RobotStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RobotStatusRequest proto.InternalMessageInfo

func (m *RobotStatusRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

// The response message containing the last RobotStatus, datetime, and errorMesage if any.
type RobotStatusReply struct {
	MoveStatus           string   `protobuf:"bytes,1,opt,name=moveStatus,proto3" json:"moveStatus,omitempty"`
	ChargeStatus         bool     `protobuf:"varint,2,opt,name=chargeStatus,proto3" json:"chargeStatus,omitempty"`
	SoftEstopStatus      bool     `protobuf:"varint,3,opt,name=softEstopStatus,proto3" json:"softEstopStatus,omitempty"`
	HardEstopStatus      bool     `protobuf:"varint,4,opt,name=hardEstopStatus,proto3" json:"hardEstopStatus,omitempty"`
	PowerPercent         int64    `protobuf:"varint,5,opt,name=powerPercent,proto3" json:"powerPercent,omitempty"`
	X                    float64  `protobuf:"fixed64,6,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float64  `protobuf:"fixed64,7,opt,name=y,proto3" json:"y,omitempty"`
	Theta                float64  `protobuf:"fixed64,8,opt,name=theta,proto3" json:"theta,omitempty"`
	Datetime             string   `protobuf:"bytes,9,opt,name=datetime,proto3" json:"datetime,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,10,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RobotStatusReply) Reset()         { *m = RobotStatusReply{} }
func (m *RobotStatusReply) String() string { return proto.CompactTextString(m) }
func (*RobotStatusReply) ProtoMessage()    {}
func (*RobotStatusReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0ec5498159ba86f, []int{1}
}

func (m *RobotStatusReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RobotStatusReply.Unmarshal(m, b)
}
func (m *RobotStatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RobotStatusReply.Marshal(b, m, deterministic)
}
func (m *RobotStatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RobotStatusReply.Merge(m, src)
}
func (m *RobotStatusReply) XXX_Size() int {
	return xxx_messageInfo_RobotStatusReply.Size(m)
}
func (m *RobotStatusReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RobotStatusReply.DiscardUnknown(m)
}

var xxx_messageInfo_RobotStatusReply proto.InternalMessageInfo

func (m *RobotStatusReply) GetMoveStatus() string {
	if m != nil {
		return m.MoveStatus
	}
	return ""
}

func (m *RobotStatusReply) GetChargeStatus() bool {
	if m != nil {
		return m.ChargeStatus
	}
	return false
}

func (m *RobotStatusReply) GetSoftEstopStatus() bool {
	if m != nil {
		return m.SoftEstopStatus
	}
	return false
}

func (m *RobotStatusReply) GetHardEstopStatus() bool {
	if m != nil {
		return m.HardEstopStatus
	}
	return false
}

func (m *RobotStatusReply) GetPowerPercent() int64 {
	if m != nil {
		return m.PowerPercent
	}
	return 0
}

func (m *RobotStatusReply) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *RobotStatusReply) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *RobotStatusReply) GetTheta() float64 {
	if m != nil {
		return m.Theta
	}
	return 0
}

func (m *RobotStatusReply) GetDatetime() string {
	if m != nil {
		return m.Datetime
	}
	return ""
}

func (m *RobotStatusReply) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*RobotStatusRequest)(nil), "robotStatus.RobotStatusRequest")
	proto.RegisterType((*RobotStatusReply)(nil), "robotStatus.RobotStatusReply")
}

func init() { proto.RegisterFile("rs.proto", fileDescriptor_e0ec5498159ba86f) }

var fileDescriptor_e0ec5498159ba86f = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xdd, 0xc6, 0xd6, 0x74, 0x2c, 0x5a, 0x06, 0x0f, 0x4b, 0x41, 0x0d, 0x39, 0x48, 0x4e,
	0x39, 0xe8, 0x33, 0x88, 0x27, 0x41, 0xb6, 0x4f, 0xb0, 0x4d, 0xc7, 0xa4, 0xd0, 0xba, 0x71, 0x76,
	0x5a, 0x9b, 0xb7, 0xf3, 0xd1, 0x24, 0x49, 0xd5, 0xa4, 0xe2, 0x6d, 0xff, 0x6f, 0x3e, 0x18, 0x76,
	0x7e, 0x08, 0xd9, 0xa7, 0x25, 0x3b, 0x71, 0x78, 0xce, 0x6e, 0xe1, 0x64, 0x2e, 0x56, 0xb6, 0x3e,
	0xbe, 0x03, 0x34, 0xbf, 0xd1, 0xd0, 0xfb, 0x96, 0xbc, 0xe0, 0x14, 0x02, 0xb1, 0xb9, 0x56, 0x91,
	0x4a, 0xc6, 0xa6, 0x7e, 0xc6, 0x9f, 0x03, 0x98, 0xf6, 0xc4, 0x72, 0x5d, 0xe1, 0x0d, 0xc0, 0xc6,
	0xed, 0xa8, 0x45, 0x07, 0xbb, 0x43, 0x30, 0x86, 0x49, 0x56, 0x58, 0xce, 0xbf, 0x8d, 0x41, 0xa4,
	0x92, 0xd0, 0xf4, 0x18, 0x26, 0x70, 0xe9, 0xdd, 0xab, 0x3c, 0x7a, 0x71, 0xe5, 0x41, 0x0b, 0x1a,
	0xed, 0x18, 0xd7, 0x66, 0x61, 0x79, 0xd9, 0x35, 0x4f, 0x5b, 0xf3, 0x08, 0xd7, 0x7b, 0x4b, 0xf7,
	0x41, 0xfc, 0x42, 0x9c, 0xd1, 0x9b, 0xe8, 0x61, 0xa4, 0x92, 0xc0, 0xf4, 0x18, 0x4e, 0x40, 0xed,
	0xf5, 0x28, 0x52, 0x89, 0x32, 0x6a, 0x5f, 0xa7, 0x4a, 0x9f, 0xb5, 0xa9, 0xc2, 0x2b, 0x18, 0x4a,
	0x41, 0x62, 0x75, 0xd8, 0x90, 0x36, 0xe0, 0x0c, 0xc2, 0xa5, 0x15, 0x92, 0xd5, 0x86, 0xf4, 0xb8,
	0xf9, 0xeb, 0x4f, 0xae, 0x37, 0x12, 0xb3, 0xe3, 0x67, 0xf2, 0xde, 0xe6, 0xa4, 0xa1, 0x99, 0xf7,
	0xd8, 0x7d, 0xd1, 0x3b, 0xf5, 0x9c, 0x78, 0xb7, 0xca, 0x08, 0x0d, 0x5c, 0x3c, 0x91, 0x74, 0x06,
	0x78, 0x9b, 0x76, 0x0a, 0x4a, 0xff, 0xb6, 0x33, 0xbb, 0xfe, 0x5f, 0x28, 0xd7, 0x55, 0x7c, 0xb2,
	0x18, 0x35, 0x45, 0x3f, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0xed, 0xe0, 0x62, 0x72, 0xf4, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RobotStatusServiceClient is the client API for RobotStatusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RobotStatusServiceClient interface {
	// Sends a greeting
	GetRobotStatus(ctx context.Context, in *RobotStatusRequest, opts ...grpc.CallOption) (*RobotStatusReply, error)
}

type robotStatusServiceClient struct {
	cc *grpc.ClientConn
}

func NewRobotStatusServiceClient(cc *grpc.ClientConn) RobotStatusServiceClient {
	return &robotStatusServiceClient{cc}
}

func (c *robotStatusServiceClient) GetRobotStatus(ctx context.Context, in *RobotStatusRequest, opts ...grpc.CallOption) (*RobotStatusReply, error) {
	out := new(RobotStatusReply)
	err := c.cc.Invoke(ctx, "/robotStatus.RobotStatusService/GetRobotStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotStatusServiceServer is the server API for RobotStatusService service.
type RobotStatusServiceServer interface {
	// Sends a greeting
	GetRobotStatus(context.Context, *RobotStatusRequest) (*RobotStatusReply, error)
}

// UnimplementedRobotStatusServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRobotStatusServiceServer struct {
}

func (*UnimplementedRobotStatusServiceServer) GetRobotStatus(ctx context.Context, req *RobotStatusRequest) (*RobotStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRobotStatus not implemented")
}

func RegisterRobotStatusServiceServer(s *grpc.Server, srv RobotStatusServiceServer) {
	s.RegisterService(&_RobotStatusService_serviceDesc, srv)
}

func _RobotStatusService_GetRobotStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RobotStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotStatusServiceServer).GetRobotStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robotStatus.RobotStatusService/GetRobotStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotStatusServiceServer).GetRobotStatus(ctx, req.(*RobotStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RobotStatusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "robotStatus.RobotStatusService",
	HandlerType: (*RobotStatusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRobotStatus",
			Handler:    _RobotStatusService_GetRobotStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rs.proto",
}
