// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gamecode.proto

package gameproto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type C2GS_CMD int32

const (
	C2GS_NONE      C2GS_CMD = 0
	C2S_LOGIN      C2GS_CMD = 1
	C2S_Test       C2GS_CMD = 10
	C2S_HEART_INFO C2GS_CMD = 254
	C2S_ACK        C2GS_CMD = 255
)

var C2GS_CMD_name = map[int32]string{
	0:   "C2GS_NONE",
	1:   "C2S_LOGIN",
	10:  "C2S_Test",
	254: "C2S_HEART_INFO",
	255: "C2S_ACK",
}

var C2GS_CMD_value = map[string]int32{
	"C2GS_NONE":      0,
	"C2S_LOGIN":      1,
	"C2S_Test":       10,
	"C2S_HEART_INFO": 254,
	"C2S_ACK":        255,
}

func (C2GS_CMD) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d88e444a8e2e7656, []int{0}
}

type GS2C_CMD int32

const (
	GS2C_NONE           GS2C_CMD = 0
	S2C_CONFIRM         GS2C_CMD = 1
	S2C_LOGIN_END       GS2C_CMD = 2
	S2C_LOGIN_CHAR_INFO GS2C_CMD = 3
	S2C_Test            GS2C_CMD = 10
)

var GS2C_CMD_name = map[int32]string{
	0:  "GS2C_NONE",
	1:  "S2C_CONFIRM",
	2:  "S2C_LOGIN_END",
	3:  "S2C_LOGIN_CHAR_INFO",
	10: "S2C_Test",
}

var GS2C_CMD_value = map[string]int32{
	"GS2C_NONE":           0,
	"S2C_CONFIRM":         1,
	"S2C_LOGIN_END":       2,
	"S2C_LOGIN_CHAR_INFO": 3,
	"S2C_Test":            10,
}

func (GS2C_CMD) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d88e444a8e2e7656, []int{1}
}

func init() {
	proto.RegisterEnum("gameproto.C2GS_CMD", C2GS_CMD_name, C2GS_CMD_value)
	proto.RegisterEnum("gameproto.GS2C_CMD", GS2C_CMD_name, GS2C_CMD_value)
}

func init() { proto.RegisterFile("gamecode.proto", fileDescriptor_d88e444a8e2e7656) }

var fileDescriptor_d88e444a8e2e7656 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4f, 0xcc, 0x4d,
	0x4d, 0xce, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0xf1, 0xc1, 0x4c,
	0xad, 0x48, 0x2e, 0x0e, 0x67, 0x23, 0xf7, 0xe0, 0x78, 0x67, 0x5f, 0x17, 0x21, 0x5e, 0x2e, 0x4e,
	0x30, 0xdb, 0xcf, 0xdf, 0xcf, 0x55, 0x80, 0x01, 0xc2, 0x0d, 0x8e, 0xf7, 0xf1, 0x77, 0xf7, 0xf4,
	0x13, 0x60, 0x14, 0xe2, 0x01, 0xa9, 0x0c, 0x8e, 0x0f, 0x49, 0x2d, 0x2e, 0x11, 0xe0, 0x12, 0x12,
	0xe6, 0xe2, 0x03, 0xf1, 0x3c, 0x5c, 0x1d, 0x83, 0x42, 0xe2, 0x3d, 0xfd, 0xdc, 0xfc, 0x05, 0xfe,
	0x81, 0x94, 0xb0, 0x83, 0x04, 0x1d, 0x9d, 0xbd, 0x05, 0xfe, 0x33, 0x6a, 0xa5, 0x70, 0x71, 0xb8,
	0x07, 0x1b, 0x39, 0xc3, 0x8c, 0x06, 0xb3, 0xa1, 0x46, 0xf3, 0x73, 0x71, 0x83, 0x65, 0xfc, 0xfd,
	0xdc, 0x3c, 0x83, 0x7c, 0x05, 0x18, 0x85, 0x04, 0xb9, 0x78, 0x41, 0x02, 0x60, 0xbb, 0xe2, 0x5d,
	0xfd, 0x5c, 0x04, 0x98, 0x84, 0xc4, 0xb9, 0x84, 0x11, 0x42, 0xce, 0x1e, 0x8e, 0x41, 0x10, 0x6b,
	0x98, 0x41, 0x0e, 0x01, 0x49, 0x40, 0x1c, 0xe2, 0x64, 0x72, 0xe1, 0xa1, 0x1c, 0xc3, 0x8d, 0x87,
	0x72, 0x0c, 0x1f, 0x1e, 0xca, 0x31, 0x36, 0x3c, 0x92, 0x63, 0x5c, 0xf1, 0x48, 0x8e, 0xf1, 0xc4,
	0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x7c, 0xf1, 0x48, 0x8e, 0xe1,
	0xc3, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e,
	0x21, 0x89, 0x0d, 0xec, 0x7b, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0x9d, 0x6b, 0xfa,
	0x1a, 0x01, 0x00, 0x00,
}

func (x C2GS_CMD) String() string {
	s, ok := C2GS_CMD_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x GS2C_CMD) String() string {
	s, ok := GS2C_CMD_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
