// Code generated by protoc-gen-go.
// source: wire.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	wire.proto

It has these top-level messages:
	Login
	Contact
	Text
	File
	Av
	Haber
*/
package main

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

// Identifies which field is filled in
type Haber_Which int32

const (
	Haber_LOGIN    Haber_Which = 0
	Haber_ROSTER   Haber_Which = 1
	Haber_PRESENCE Haber_Which = 2
	Haber_TEXT     Haber_Which = 3
	Haber_FILE     Haber_Which = 4
	Haber_AV       Haber_Which = 5
)

var Haber_Which_name = map[int32]string{
	0: "LOGIN",
	1: "ROSTER",
	2: "PRESENCE",
	3: "TEXT",
	4: "FILE",
	5: "AV",
}
var Haber_Which_value = map[string]int32{
	"LOGIN":    0,
	"ROSTER":   1,
	"PRESENCE": 2,
	"TEXT":     3,
	"FILE":     4,
	"AV":       5,
}

func (x Haber_Which) String() string {
	return proto.EnumName(Haber_Which_name, int32(x))
}
func (Haber_Which) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type Login struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *Login) Reset()                    { *m = Login{} }
func (m *Login) String() string            { return proto.CompactTextString(m) }
func (*Login) ProtoMessage()               {}
func (*Login) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Login) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Contact struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Online bool   `protobuf:"varint,2,opt,name=online" json:"online,omitempty"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Contact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Contact) GetOnline() bool {
	if m != nil {
		return m.Online
	}
	return false
}

type Text struct {
	Body string `protobuf:"bytes,1,opt,name=body" json:"body,omitempty"`
}

func (m *Text) Reset()                    { *m = Text{} }
func (m *Text) String() string            { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()               {}
func (*Text) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Text) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type File struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *File) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Av struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Av) Reset()                    { *m = Av{} }
func (m *Av) String() string            { return proto.CompactTextString(m) }
func (*Av) ProtoMessage()               {}
func (*Av) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Av) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Haber struct {
	Version   uint32      `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	SessionId string      `protobuf:"bytes,2,opt,name=sessionId" json:"sessionId,omitempty"`
	From      string      `protobuf:"bytes,3,opt,name=from" json:"from,omitempty"`
	To        string      `protobuf:"bytes,4,opt,name=to" json:"to,omitempty"`
	Which     Haber_Which `protobuf:"varint,5,opt,name=which,enum=Haber_Which" json:"which,omitempty"`
	// One of the following will be filled in
	Login    *Login     `protobuf:"bytes,101,opt,name=login" json:"login,omitempty"`
	Roster   []*Contact `protobuf:"bytes,102,rep,name=roster" json:"roster,omitempty"`
	Presence *Contact   `protobuf:"bytes,103,opt,name=presence" json:"presence,omitempty"`
	Text     *Text      `protobuf:"bytes,104,opt,name=text" json:"text,omitempty"`
	Av       *Av        `protobuf:"bytes,105,opt,name=av" json:"av,omitempty"`
	File     *File      `protobuf:"bytes,106,opt,name=file" json:"file,omitempty"`
}

func (m *Haber) Reset()                    { *m = Haber{} }
func (m *Haber) String() string            { return proto.CompactTextString(m) }
func (*Haber) ProtoMessage()               {}
func (*Haber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Haber) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Haber) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Haber) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Haber) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Haber) GetWhich() Haber_Which {
	if m != nil {
		return m.Which
	}
	return Haber_LOGIN
}

func (m *Haber) GetLogin() *Login {
	if m != nil {
		return m.Login
	}
	return nil
}

func (m *Haber) GetRoster() []*Contact {
	if m != nil {
		return m.Roster
	}
	return nil
}

func (m *Haber) GetPresence() *Contact {
	if m != nil {
		return m.Presence
	}
	return nil
}

func (m *Haber) GetText() *Text {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *Haber) GetAv() *Av {
	if m != nil {
		return m.Av
	}
	return nil
}

func (m *Haber) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

func init() {
	proto.RegisterType((*Login)(nil), "Login")
	proto.RegisterType((*Contact)(nil), "Contact")
	proto.RegisterType((*Text)(nil), "Text")
	proto.RegisterType((*File)(nil), "File")
	proto.RegisterType((*Av)(nil), "Av")
	proto.RegisterType((*Haber)(nil), "Haber")
	proto.RegisterEnum("Haber_Which", Haber_Which_name, Haber_Which_value)
}

func init() { proto.RegisterFile("wire.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xf1, 0x9f, 0x75, 0x9d, 0x69, 0xa8, 0xac, 0x41, 0x42, 0x4b, 0xd5, 0x83, 0x65, 0x38,
	0xf8, 0xe4, 0x43, 0x10, 0x1f, 0x20, 0x54, 0x2e, 0x8d, 0x14, 0xb5, 0x68, 0x1b, 0x01, 0xe2, 0xe6,
	0x24, 0x93, 0x66, 0x51, 0xe2, 0xad, 0xd6, 0x8b, 0x5b, 0xae, 0x7c, 0x72, 0xb4, 0x13, 0xb7, 0xe5,
	0xd0, 0xdb, 0x9b, 0xf7, 0x7b, 0xbb, 0x1e, 0x3f, 0x2d, 0xc0, 0xbd, 0xb6, 0x54, 0xdd, 0x59, 0xe3,
	0x4c, 0xf1, 0x1e, 0xc4, 0xdc, 0xdc, 0xea, 0x16, 0x4f, 0x21, 0xfd, 0xdd, 0x91, 0x6d, 0x9b, 0x3d,
	0xc9, 0x20, 0x0f, 0xca, 0x91, 0x7a, 0x9a, 0x8b, 0x4f, 0x70, 0x74, 0x6e, 0x5a, 0xd7, 0xac, 0x1c,
	0x22, 0xc4, 0xff, 0x45, 0x58, 0xe3, 0x5b, 0x48, 0x4c, 0xbb, 0xd3, 0x2d, 0xc9, 0x30, 0x0f, 0xca,
	0x54, 0x0d, 0x53, 0x71, 0x0a, 0xf1, 0x82, 0x1e, 0xf8, 0xcc, 0xd2, 0xac, 0xff, 0x3c, 0x9e, 0xf1,
	0xda, 0xb3, 0x0b, 0xbd, 0x23, 0xcf, 0xd6, 0x8d, 0x6b, 0x98, 0x8d, 0x15, 0xeb, 0x42, 0x42, 0x38,
	0xed, 0x5f, 0x24, 0x7f, 0x23, 0x10, 0x97, 0xcd, 0x92, 0x2c, 0x4a, 0x38, 0xea, 0xc9, 0x76, 0xda,
	0xb4, 0x1c, 0x78, 0xad, 0x1e, 0x47, 0x3c, 0x83, 0x51, 0x47, 0x9d, 0x97, 0xb3, 0x35, 0x2f, 0x34,
	0x52, 0xcf, 0x86, 0xbf, 0x75, 0x63, 0xcd, 0x5e, 0x46, 0x87, 0x5d, 0xbc, 0xc6, 0x13, 0x08, 0x9d,
	0x91, 0x31, 0x3b, 0xa1, 0x33, 0x58, 0x80, 0xb8, 0xdf, 0xea, 0xd5, 0x56, 0x8a, 0x3c, 0x28, 0x4f,
	0x26, 0xe3, 0x8a, 0x3f, 0x59, 0x7d, 0xf7, 0x9e, 0x3a, 0x20, 0x3c, 0x03, 0xb1, 0xf3, 0xbd, 0x49,
	0xca, 0x83, 0xf2, 0x78, 0x92, 0x54, 0xdc, 0xa2, 0x3a, 0x98, 0x98, 0x43, 0x62, 0x4d, 0xe7, 0xc8,
	0xca, 0x4d, 0x1e, 0x95, 0xc7, 0x93, 0xb4, 0x1a, 0xfa, 0x53, 0x83, 0x8f, 0x1f, 0x20, 0xbd, 0xb3,
	0xd4, 0x51, 0xbb, 0x22, 0x79, 0xcb, 0x57, 0x3c, 0x67, 0x9e, 0x08, 0xbe, 0x83, 0xd8, 0xd1, 0x83,
	0x93, 0x5b, 0x4e, 0x88, 0xca, 0xd7, 0xa9, 0xd8, 0xc2, 0x37, 0x10, 0x36, 0xbd, 0xd4, 0x0c, 0xa2,
	0x6a, 0xda, 0xab, 0xb0, 0xe9, 0x7d, 0x7e, 0xa3, 0x77, 0x24, 0x7f, 0x0d, 0x79, 0x5f, 0xb1, 0x62,
	0xab, 0xb8, 0x04, 0xc1, 0x3f, 0x80, 0x23, 0x10, 0xf3, 0xeb, 0x2f, 0xb3, 0xab, 0xec, 0x15, 0x02,
	0x24, 0xea, 0xfa, 0x66, 0x51, 0xab, 0x2c, 0xc0, 0x31, 0xa4, 0x5f, 0x55, 0x7d, 0x53, 0x5f, 0x9d,
	0xd7, 0x59, 0x88, 0x29, 0xc4, 0x8b, 0xfa, 0xc7, 0x22, 0x8b, 0xbc, 0xba, 0x98, 0xcd, 0xeb, 0x2c,
	0xc6, 0x04, 0xc2, 0xe9, 0xb7, 0x4c, 0x7c, 0x4e, 0x7e, 0xc6, 0xfb, 0x46, 0xb7, 0xcb, 0x84, 0x5f,
	0xd0, 0xc7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x76, 0xf4, 0xa3, 0x4f, 0x02, 0x00, 0x00,
}
