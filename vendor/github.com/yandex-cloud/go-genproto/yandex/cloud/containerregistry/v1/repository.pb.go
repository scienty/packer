// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/containerregistry/v1/repository.proto

package containerregistry

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A Repository resource. For more information, see [Repository](/docs/cloud/container-registry/repository).
type Repository struct {
	// Name of the repository.
	// The name is unique within the registry.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Repository) Reset()         { *m = Repository{} }
func (m *Repository) String() string { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()    {}
func (*Repository) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a476f200b225be9, []int{0}
}

func (m *Repository) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repository.Unmarshal(m, b)
}
func (m *Repository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repository.Marshal(b, m, deterministic)
}
func (m *Repository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repository.Merge(m, src)
}
func (m *Repository) XXX_Size() int {
	return xxx_messageInfo_Repository.Size(m)
}
func (m *Repository) XXX_DiscardUnknown() {
	xxx_messageInfo_Repository.DiscardUnknown(m)
}

var xxx_messageInfo_Repository proto.InternalMessageInfo

func (m *Repository) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Repository)(nil), "yandex.cloud.containerregistry.v1.Repository")
}

func init() {
	proto.RegisterFile("yandex/cloud/containerregistry/v1/repository.proto", fileDescriptor_3a476f200b225be9)
}

var fileDescriptor_3a476f200b225be9 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xaa, 0x4c, 0xcc, 0x4b,
	0x49, 0xad, 0xd0, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcc, 0xcc,
	0x4b, 0x2d, 0x2a, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x29, 0xaa, 0xd4, 0x2f, 0x33, 0xd4, 0x2f, 0x4a,
	0x2d, 0xc8, 0x2f, 0xce, 0x2c, 0xc9, 0x2f, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52,
	0x84, 0xe8, 0xd1, 0x03, 0xeb, 0xd1, 0xc3, 0xd0, 0xa3, 0x57, 0x66, 0xa8, 0xa4, 0xc0, 0xc5, 0x15,
	0x04, 0xd7, 0x26, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x19, 0x04, 0x66, 0x3b, 0x45, 0x46, 0x85, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0xea, 0x43, 0x4c, 0xd4, 0x85, 0xb8, 0x22, 0x3d, 0x5f, 0x37, 0x3d, 0x35, 0x0f, 0x6c, 0x97,
	0x3e, 0x41, 0xe7, 0x59, 0x63, 0x08, 0x26, 0xb1, 0x81, 0xb5, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x1d, 0x25, 0xf5, 0x56, 0xdc, 0x00, 0x00, 0x00,
}
