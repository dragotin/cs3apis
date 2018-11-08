// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cernbox/publicshare/v1/resources.proto

package publicsharev1pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Permission int32

const (
	Permission_PERMISSION_INVALID   Permission = 0
	Permission_PERMISSION_VIEW_ONLY Permission = 1
	Permission_PERMISSION_MODIFY    Permission = 2
	Permission_PERMISSION_DROP_ONLY Permission = 3
)

var Permission_name = map[int32]string{
	0: "PERMISSION_INVALID",
	1: "PERMISSION_VIEW_ONLY",
	2: "PERMISSION_MODIFY",
	3: "PERMISSION_DROP_ONLY",
}

var Permission_value = map[string]int32{
	"PERMISSION_INVALID":   0,
	"PERMISSION_VIEW_ONLY": 1,
	"PERMISSION_MODIFY":    2,
	"PERMISSION_DROP_ONLY": 3,
}

func (x Permission) String() string {
	return proto.EnumName(Permission_name, int32(x))
}

func (Permission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4fd4c7b580ef6f38, []int{0}
}

type FileType int32

const (
	FileType_FILE_TYPE_INVALID FileType = 0
	FileType_FILE_TYPE_FILE    FileType = 1
	FileType_FILE_TYPE_FOLDER  FileType = 2
)

var FileType_name = map[int32]string{
	0: "FILE_TYPE_INVALID",
	1: "FILE_TYPE_FILE",
	2: "FILE_TYPE_FOLDER",
}

var FileType_value = map[string]int32{
	"FILE_TYPE_INVALID": 0,
	"FILE_TYPE_FILE":    1,
	"FILE_TYPE_FOLDER":  2,
}

func (x FileType) String() string {
	return proto.EnumName(FileType_name, int32(x))
}

func (FileType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4fd4c7b580ef6f38, []int{1}
}

type UpdatePolicy struct {
	SetPassword          bool     `protobuf:"varint,1,opt,name=set_password,json=setPassword,proto3" json:"set_password,omitempty"`
	SetExpiration        bool     `protobuf:"varint,2,opt,name=set_expiration,json=setExpiration,proto3" json:"set_expiration,omitempty"`
	SetPermission        bool     `protobuf:"varint,3,opt,name=set_permission,json=setPermission,proto3" json:"set_permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePolicy) Reset()         { *m = UpdatePolicy{} }
func (m *UpdatePolicy) String() string { return proto.CompactTextString(m) }
func (*UpdatePolicy) ProtoMessage()    {}
func (*UpdatePolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd4c7b580ef6f38, []int{0}
}

func (m *UpdatePolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePolicy.Unmarshal(m, b)
}
func (m *UpdatePolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePolicy.Marshal(b, m, deterministic)
}
func (m *UpdatePolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePolicy.Merge(m, src)
}
func (m *UpdatePolicy) XXX_Size() int {
	return xxx_messageInfo_UpdatePolicy.Size(m)
}
func (m *UpdatePolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePolicy.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePolicy proto.InternalMessageInfo

func (m *UpdatePolicy) GetSetPassword() bool {
	if m != nil {
		return m.SetPassword
	}
	return false
}

func (m *UpdatePolicy) GetSetExpiration() bool {
	if m != nil {
		return m.SetExpiration
	}
	return false
}

func (m *UpdatePolicy) GetSetPermission() bool {
	if m != nil {
		return m.SetPermission
	}
	return false
}

type PublicShare struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string     `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Filename             string     `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	FileType             FileType   `protobuf:"varint,4,opt,name=file_type,json=fileType,proto3,enum=cernbox.publicsharev1.FileType" json:"file_type,omitempty"`
	Expiration           uint64     `protobuf:"varint,5,opt,name=expiration,proto3" json:"expiration,omitempty"`
	PasswordProtected    bool       `protobuf:"varint,6,opt,name=password_protected,json=passwordProtected,proto3" json:"password_protected,omitempty"`
	Permission           Permission `protobuf:"varint,7,opt,name=permission,proto3,enum=cernbox.publicsharev1.Permission" json:"permission,omitempty"`
	DisplayName          string     `protobuf:"bytes,8,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Owner                string     `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PublicShare) Reset()         { *m = PublicShare{} }
func (m *PublicShare) String() string { return proto.CompactTextString(m) }
func (*PublicShare) ProtoMessage()    {}
func (*PublicShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd4c7b580ef6f38, []int{1}
}

func (m *PublicShare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicShare.Unmarshal(m, b)
}
func (m *PublicShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicShare.Marshal(b, m, deterministic)
}
func (m *PublicShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicShare.Merge(m, src)
}
func (m *PublicShare) XXX_Size() int {
	return xxx_messageInfo_PublicShare.Size(m)
}
func (m *PublicShare) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicShare.DiscardUnknown(m)
}

var xxx_messageInfo_PublicShare proto.InternalMessageInfo

func (m *PublicShare) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PublicShare) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PublicShare) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *PublicShare) GetFileType() FileType {
	if m != nil {
		return m.FileType
	}
	return FileType_FILE_TYPE_INVALID
}

func (m *PublicShare) GetExpiration() uint64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

func (m *PublicShare) GetPasswordProtected() bool {
	if m != nil {
		return m.PasswordProtected
	}
	return false
}

func (m *PublicShare) GetPermission() Permission {
	if m != nil {
		return m.Permission
	}
	return Permission_PERMISSION_INVALID
}

func (m *PublicShare) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *PublicShare) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func init() {
	proto.RegisterEnum("cernbox.publicsharev1.Permission", Permission_name, Permission_value)
	proto.RegisterEnum("cernbox.publicsharev1.FileType", FileType_name, FileType_value)
	proto.RegisterType((*UpdatePolicy)(nil), "cernbox.publicsharev1.UpdatePolicy")
	proto.RegisterType((*PublicShare)(nil), "cernbox.publicsharev1.PublicShare")
}

func init() {
	proto.RegisterFile("cernbox/publicshare/v1/resources.proto", fileDescriptor_4fd4c7b580ef6f38)
}

var fileDescriptor_4fd4c7b580ef6f38 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xd1, 0x8e, 0xd2, 0x40,
	0x14, 0xb5, 0xdd, 0x5d, 0x84, 0xcb, 0x8a, 0xdd, 0x09, 0x98, 0xea, 0x83, 0xb2, 0x9b, 0x68, 0xc8,
	0x26, 0x42, 0xd0, 0x57, 0x5f, 0x28, 0x94, 0xa4, 0x91, 0xa5, 0xcd, 0xb0, 0xa2, 0x18, 0x93, 0xa6,
	0x94, 0x6b, 0x6c, 0x84, 0x4e, 0x33, 0xd3, 0x5d, 0xe0, 0xc9, 0x7f, 0xf1, 0xd1, 0x4f, 0xf1, 0x0b,
	0xfc, 0x1c, 0x33, 0x2d, 0x85, 0x59, 0xb3, 0xbe, 0xcd, 0x3d, 0xf7, 0x9c, 0xb9, 0xf7, 0x9c, 0x4e,
	0xe1, 0x55, 0x88, 0x3c, 0x9e, 0xb3, 0x4d, 0x27, 0xb9, 0x99, 0x2f, 0xa3, 0x50, 0x7c, 0x0b, 0x38,
	0x76, 0x6e, 0xbb, 0x1d, 0x8e, 0x82, 0xdd, 0xf0, 0x10, 0x45, 0x3b, 0xe1, 0x2c, 0x65, 0xa4, 0xb1,
	0xe3, 0xb5, 0x15, 0xde, 0x6d, 0xf7, 0xe2, 0x07, 0x9c, 0x7e, 0x48, 0x16, 0x41, 0x8a, 0x1e, 0x5b,
	0x46, 0xe1, 0x96, 0x9c, 0xc3, 0xa9, 0xc0, 0xd4, 0x4f, 0x02, 0x21, 0xd6, 0x8c, 0x2f, 0x4c, 0xad,
	0xa9, 0xb5, 0xca, 0xb4, 0x2a, 0x30, 0xf5, 0x76, 0x10, 0x79, 0x09, 0x35, 0x49, 0xc1, 0x4d, 0x12,
	0xf1, 0x20, 0x8d, 0x58, 0x6c, 0xea, 0x19, 0xe9, 0x91, 0xc0, 0xd4, 0xde, 0x83, 0x05, 0x2d, 0x41,
	0xbe, 0x8a, 0x84, 0x90, 0xb4, 0xa3, 0x3d, 0xcd, 0xdb, 0x83, 0x17, 0x7f, 0x74, 0xa8, 0x7a, 0xd9,
	0x4a, 0x13, 0xb9, 0x12, 0xa9, 0x81, 0x1e, 0xe5, 0x63, 0x2b, 0x54, 0x8f, 0x16, 0xa4, 0x0e, 0x27,
	0x29, 0xfb, 0x8e, 0xf9, 0x90, 0x0a, 0xcd, 0x0b, 0xf2, 0x0c, 0xca, 0x5f, 0xa3, 0x25, 0xc6, 0xc1,
	0x0a, 0xb3, 0x6b, 0x2b, 0x74, 0x5f, 0x93, 0x77, 0x50, 0x91, 0x67, 0x3f, 0xdd, 0x26, 0x68, 0x1e,
	0x37, 0xb5, 0x56, 0xed, 0xcd, 0x8b, 0xf6, 0xbd, 0xee, 0xdb, 0xc3, 0x68, 0x89, 0xd7, 0xdb, 0x04,
	0x73, 0xb5, 0x3c, 0x91, 0xe7, 0x00, 0x8a, 0xb3, 0x93, 0xa6, 0xd6, 0x3a, 0xa6, 0x0a, 0x42, 0x5e,
	0x03, 0x29, 0xc2, 0xf1, 0x65, 0xb2, 0x18, 0xa6, 0xb8, 0x30, 0x4b, 0x99, 0xb5, 0xb3, 0xa2, 0xe3,
	0x15, 0x0d, 0xd2, 0x03, 0x50, 0x12, 0x78, 0x98, 0x6d, 0x73, 0xfe, 0x9f, 0x6d, 0x0e, 0xa9, 0x50,
	0x45, 0x24, 0x3f, 0xc9, 0x22, 0x12, 0xc9, 0x32, 0xd8, 0xfa, 0x99, 0xdf, 0x72, 0xe6, 0xb7, 0xba,
	0xc3, 0xc6, 0xd2, 0x72, 0x1d, 0x4e, 0xd8, 0x3a, 0x46, 0x6e, 0x56, 0xf2, 0x90, 0xb2, 0xe2, 0x92,
	0x01, 0x1c, 0xae, 0x24, 0x4f, 0x80, 0x78, 0x36, 0xbd, 0x72, 0x26, 0x13, 0xc7, 0x1d, 0xfb, 0xce,
	0x78, 0xda, 0x1b, 0x39, 0x03, 0xe3, 0x01, 0x31, 0xa1, 0xae, 0xe0, 0x53, 0xc7, 0xfe, 0xe8, 0xbb,
	0xe3, 0xd1, 0xcc, 0xd0, 0x48, 0x03, 0xce, 0x94, 0xce, 0x95, 0x3b, 0x70, 0x86, 0x33, 0x43, 0xff,
	0x47, 0x30, 0xa0, 0xae, 0x97, 0x0b, 0x8e, 0x2e, 0xdf, 0x43, 0xb9, 0x48, 0x54, 0x8a, 0x87, 0xce,
	0xc8, 0xf6, 0xaf, 0x67, 0x9e, 0xad, 0x4c, 0x23, 0x50, 0x3b, 0xc0, 0xf2, 0x64, 0x68, 0xa4, 0x0e,
	0x86, 0x82, 0xb9, 0xa3, 0x81, 0x4d, 0x0d, 0xdd, 0x5a, 0xc3, 0xd3, 0x90, 0xad, 0xee, 0x8f, 0xca,
	0xaa, 0xd1, 0xe2, 0x79, 0xcb, 0xa8, 0x99, 0xa7, 0x7d, 0x7e, 0x7c, 0x87, 0x90, 0xcc, 0x7f, 0xea,
	0xa5, 0xbe, 0xe5, 0x7e, 0xea, 0x59, 0xbf, 0xf4, 0x46, 0xdf, 0xa6, 0x63, 0x8b, 0x6d, 0xda, 0xca,
	0x3b, 0x9b, 0x76, 0x7f, 0xeb, 0x8d, 0x3e, 0xf2, 0xd8, 0x62, 0x9b, 0x2f, 0x77, 0xf0, 0x79, 0x29,
	0xfb, 0x61, 0xde, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x36, 0x1d, 0x34, 0xb4, 0x5a, 0x03, 0x00,
	0x00,
}
