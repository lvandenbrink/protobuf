// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: testdata/dessert.proto

package testdata

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/opsee/protobuf/opseeproto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// A delicious dessert dish on the menu
type Dessert struct {
	// The name of the dish
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// How sweet is the dish, an integer between 0 and 10
	Sweetness            int32    `protobuf:"varint,2,opt,name=sweetness,proto3" json:"sweetness,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dessert) Reset()         { *m = Dessert{} }
func (m *Dessert) String() string { return proto.CompactTextString(m) }
func (*Dessert) ProtoMessage()    {}
func (*Dessert) Descriptor() ([]byte, []int) {
	return fileDescriptor_cce85ba81638c849, []int{0}
}
func (m *Dessert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dessert.Unmarshal(m, b)
}
func (m *Dessert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dessert.Marshal(b, m, deterministic)
}
func (m *Dessert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dessert.Merge(m, src)
}
func (m *Dessert) XXX_Size() int {
	return xxx_messageInfo_Dessert.Size(m)
}
func (m *Dessert) XXX_DiscardUnknown() {
	xxx_messageInfo_Dessert.DiscardUnknown(m)
}

var xxx_messageInfo_Dessert proto.InternalMessageInfo

func (m *Dessert) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Dessert) GetSweetness() int32 {
	if m != nil {
		return m.Sweetness
	}
	return 0
}

func init() {
	proto.RegisterType((*Dessert)(nil), "testdata.Dessert")
}

func init() { proto.RegisterFile("testdata/dessert.proto", fileDescriptor_cce85ba81638c849) }

var fileDescriptor_cce85ba81638c849 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x49, 0x2d, 0x2e,
	0x49, 0x49, 0x2c, 0x49, 0xd4, 0x4f, 0x49, 0x2d, 0x2e, 0x4e, 0x2d, 0x2a, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x80, 0x89, 0x4b, 0x19, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0xe7, 0x17, 0x14, 0xa7, 0xa6, 0xea, 0x83, 0x55, 0x24, 0x95, 0xa6, 0x41, 0xb8,
	0x60, 0x1e, 0x84, 0x09, 0xd1, 0xab, 0x64, 0xcd, 0xc5, 0xee, 0x02, 0x31, 0x4c, 0x48, 0x88, 0x8b,
	0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x92, 0xe1,
	0xe2, 0x2c, 0x2e, 0x4f, 0x4d, 0x2d, 0xc9, 0x4b, 0x2d, 0x2e, 0x96, 0x60, 0x52, 0x60, 0xd4, 0x60,
	0x0d, 0x42, 0x08, 0x38, 0xb1, 0x1c, 0x58, 0x24, 0xcf, 0x98, 0xc4, 0x06, 0x36, 0xc9, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x6c, 0x56, 0xa1, 0x2e, 0x9f, 0x00, 0x00, 0x00,
}
