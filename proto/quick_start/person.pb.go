// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go-study/proto/quick_start/person.proto

package proto_quick_start

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_93699bb8dced1bcb, []int{0, 0}
}

type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated          *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_93699bb8dced1bcb, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=proto.quick_start.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_93699bb8dced1bcb, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

func init() {
	proto.RegisterEnum("proto.quick_start.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "proto.quick_start.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "proto.quick_start.Person.PhoneNumber")
}

func init() {
	proto.RegisterFile("go-study/proto/quick_start/person.proto", fileDescriptor_93699bb8dced1bcb)
}

var fileDescriptor_93699bb8dced1bcb = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0xff, 0xa4, 0x69, 0xf8, 0x7b, 0x23, 0xa5, 0x0e, 0x22, 0xa1, 0x1b, 0x43, 0x45, 0x0c,
	0x88, 0x13, 0xa8, 0x0b, 0x57, 0xdd, 0x08, 0x05, 0x45, 0x6b, 0xca, 0x50, 0x71, 0x67, 0x99, 0x98,
	0x31, 0x06, 0x93, 0xcc, 0x98, 0x99, 0x2c, 0xf2, 0x90, 0xbe, 0x93, 0x64, 0x26, 0x95, 0x82, 0x0b,
	0x57, 0x39, 0x37, 0xf7, 0x9c, 0x33, 0xdf, 0x85, 0xf3, 0x8c, 0x5f, 0x4a, 0xd5, 0xa4, 0x6d, 0x24,
	0x6a, 0xae, 0x78, 0xf4, 0xd9, 0xe4, 0xaf, 0x1f, 0x5b, 0xa9, 0x68, 0xad, 0x22, 0xc1, 0x6a, 0xc9,
	0x2b, 0xac, 0x17, 0xe8, 0x50, 0x7f, 0xf0, 0xde, 0x7e, 0x7a, 0x92, 0x71, 0x9e, 0x15, 0xcc, 0x24,
	0x93, 0xe6, 0x2d, 0x52, 0x79, 0xc9, 0xa4, 0xa2, 0xa5, 0x30, 0x99, 0xd9, 0x97, 0x0d, 0xee, 0x5a,
	0x97, 0x20, 0x04, 0x4e, 0x45, 0x4b, 0xe6, 0x5b, 0x81, 0x15, 0x8e, 0x88, 0xd6, 0x68, 0x0c, 0x76,
	0x9e, 0xfa, 0x76, 0x60, 0x85, 0x43, 0x62, 0xe7, 0x29, 0x3a, 0x82, 0x21, 0x2b, 0x69, 0x5e, 0xf8,
	0x03, 0x6d, 0x32, 0x03, 0x5a, 0x80, 0x2b, 0xde, 0x79, 0xc5, 0xa4, 0xef, 0x04, 0x83, 0xd0, 0x9b,
	0x9f, 0xe1, 0x5f, 0x24, 0xd8, 0x3c, 0x82, 0xd7, 0x9d, 0xef, 0xb1, 0x29, 0x13, 0x56, 0x93, 0x3e,
	0x84, 0x16, 0x70, 0x50, 0x50, 0xa9, 0xb6, 0x8d, 0x48, 0xa9, 0x62, 0xa9, 0x3f, 0x0c, 0xac, 0xd0,
	0x9b, 0x4f, 0xb1, 0x61, 0xc7, 0x3b, 0x76, 0xbc, 0xd9, 0xb1, 0x13, 0xaf, 0xf3, 0x3f, 0x19, 0xfb,
	0xf4, 0x05, 0xbc, 0xbd, 0x56, 0x74, 0x0c, 0x6e, 0xa5, 0x55, 0x7f, 0x48, 0x3f, 0xa1, 0x6b, 0x70,
	0x54, 0x2b, 0x98, 0x3e, 0x66, 0x3c, 0x3f, 0xfd, 0x03, 0x71, 0xd3, 0x0a, 0x46, 0x74, 0x60, 0x76,
	0x01, 0xa3, 0x9f, 0x5f, 0x08, 0xc0, 0x5d, 0xc5, 0x37, 0x77, 0x0f, 0xcb, 0xc9, 0x3f, 0xf4, 0x1f,
	0x9c, 0xdb, 0x78, 0xb5, 0x9c, 0x58, 0x9d, 0x7a, 0x8e, 0xc9, 0xfd, 0xc4, 0x4e, 0x5c, 0x5d, 0x7b,
	0xf5, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x25, 0x6e, 0x9f, 0x4c, 0xb5, 0x01, 0x00, 0x00,
}
