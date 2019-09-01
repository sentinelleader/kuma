// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mesh/v1alpha1/mesh.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Mesh defines configuration of a single mesh.
type Mesh struct {
	// mTLS settings.
	// +optional
	Mtls *Mesh_Mtls `protobuf:"bytes,1,opt,name=mtls,proto3" json:"mtls,omitempty"`
	// Tracing settings.
	// +optional
	Tracing              *Tracing `protobuf:"bytes,2,opt,name=tracing,proto3" json:"tracing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mesh) Reset()         { *m = Mesh{} }
func (m *Mesh) String() string { return proto.CompactTextString(m) }
func (*Mesh) ProtoMessage()    {}
func (*Mesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b3cd8c92bbf6a, []int{0}
}
func (m *Mesh) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mesh.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mesh.Merge(m, src)
}
func (m *Mesh) XXX_Size() int {
	return m.Size()
}
func (m *Mesh) XXX_DiscardUnknown() {
	xxx_messageInfo_Mesh.DiscardUnknown(m)
}

var xxx_messageInfo_Mesh proto.InternalMessageInfo

func (m *Mesh) GetMtls() *Mesh_Mtls {
	if m != nil {
		return m.Mtls
	}
	return nil
}

func (m *Mesh) GetTracing() *Tracing {
	if m != nil {
		return m.Tracing
	}
	return nil
}

// mTLS settings of a Mesh.
type Mesh_Mtls struct {
	// Certificate Authority of a Mesh.
	// +optional
	Ca                   *CertificateAuthority `protobuf:"bytes,1,opt,name=ca,proto3" json:"ca,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Mesh_Mtls) Reset()         { *m = Mesh_Mtls{} }
func (m *Mesh_Mtls) String() string { return proto.CompactTextString(m) }
func (*Mesh_Mtls) ProtoMessage()    {}
func (*Mesh_Mtls) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b3cd8c92bbf6a, []int{0, 0}
}
func (m *Mesh_Mtls) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mesh_Mtls) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mesh_Mtls.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mesh_Mtls) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mesh_Mtls.Merge(m, src)
}
func (m *Mesh_Mtls) XXX_Size() int {
	return m.Size()
}
func (m *Mesh_Mtls) XXX_DiscardUnknown() {
	xxx_messageInfo_Mesh_Mtls.DiscardUnknown(m)
}

var xxx_messageInfo_Mesh_Mtls proto.InternalMessageInfo

func (m *Mesh_Mtls) GetCa() *CertificateAuthority {
	if m != nil {
		return m.Ca
	}
	return nil
}

// CertificateAuthority defines configuration of a CA.
type CertificateAuthority struct {
	// Types that are valid to be assigned to Type:
	//	*CertificateAuthority_Builtin_
	Type                 isCertificateAuthority_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *CertificateAuthority) Reset()         { *m = CertificateAuthority{} }
func (m *CertificateAuthority) String() string { return proto.CompactTextString(m) }
func (*CertificateAuthority) ProtoMessage()    {}
func (*CertificateAuthority) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b3cd8c92bbf6a, []int{1}
}
func (m *CertificateAuthority) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CertificateAuthority) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CertificateAuthority.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CertificateAuthority) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateAuthority.Merge(m, src)
}
func (m *CertificateAuthority) XXX_Size() int {
	return m.Size()
}
func (m *CertificateAuthority) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateAuthority.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateAuthority proto.InternalMessageInfo

type isCertificateAuthority_Type interface {
	isCertificateAuthority_Type()
	MarshalTo([]byte) (int, error)
	Size() int
}

type CertificateAuthority_Builtin_ struct {
	Builtin *CertificateAuthority_Builtin `protobuf:"bytes,1,opt,name=builtin,proto3,oneof"`
}

func (*CertificateAuthority_Builtin_) isCertificateAuthority_Type() {}

func (m *CertificateAuthority) GetType() isCertificateAuthority_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *CertificateAuthority) GetBuiltin() *CertificateAuthority_Builtin {
	if x, ok := m.GetType().(*CertificateAuthority_Builtin_); ok {
		return x.Builtin
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CertificateAuthority) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CertificateAuthority_OneofMarshaler, _CertificateAuthority_OneofUnmarshaler, _CertificateAuthority_OneofSizer, []interface{}{
		(*CertificateAuthority_Builtin_)(nil),
	}
}

func _CertificateAuthority_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CertificateAuthority)
	// type
	switch x := m.Type.(type) {
	case *CertificateAuthority_Builtin_:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Builtin); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CertificateAuthority.Type has unexpected type %T", x)
	}
	return nil
}

func _CertificateAuthority_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CertificateAuthority)
	switch tag {
	case 1: // type.builtin
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CertificateAuthority_Builtin)
		err := b.DecodeMessage(msg)
		m.Type = &CertificateAuthority_Builtin_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CertificateAuthority_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CertificateAuthority)
	// type
	switch x := m.Type.(type) {
	case *CertificateAuthority_Builtin_:
		s := proto.Size(x.Builtin)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Builtin defines configuration of the builtin CA.
type CertificateAuthority_Builtin struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateAuthority_Builtin) Reset()         { *m = CertificateAuthority_Builtin{} }
func (m *CertificateAuthority_Builtin) String() string { return proto.CompactTextString(m) }
func (*CertificateAuthority_Builtin) ProtoMessage()    {}
func (*CertificateAuthority_Builtin) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b3cd8c92bbf6a, []int{1, 0}
}
func (m *CertificateAuthority_Builtin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CertificateAuthority_Builtin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CertificateAuthority_Builtin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CertificateAuthority_Builtin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateAuthority_Builtin.Merge(m, src)
}
func (m *CertificateAuthority_Builtin) XXX_Size() int {
	return m.Size()
}
func (m *CertificateAuthority_Builtin) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateAuthority_Builtin.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateAuthority_Builtin proto.InternalMessageInfo

// Tracing defines tracing configuration of the mesh.
type Tracing struct {
	// Types that are valid to be assigned to Type:
	//	*Tracing_Zipkin_
	Type                 isTracing_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Tracing) Reset()         { *m = Tracing{} }
func (m *Tracing) String() string { return proto.CompactTextString(m) }
func (*Tracing) ProtoMessage()    {}
func (*Tracing) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b3cd8c92bbf6a, []int{2}
}
func (m *Tracing) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tracing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tracing.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tracing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing.Merge(m, src)
}
func (m *Tracing) XXX_Size() int {
	return m.Size()
}
func (m *Tracing) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing proto.InternalMessageInfo

type isTracing_Type interface {
	isTracing_Type()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Tracing_Zipkin_ struct {
	Zipkin *Tracing_Zipkin `protobuf:"bytes,1,opt,name=zipkin,proto3,oneof"`
}

func (*Tracing_Zipkin_) isTracing_Type() {}

func (m *Tracing) GetType() isTracing_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Tracing) GetZipkin() *Tracing_Zipkin {
	if x, ok := m.GetType().(*Tracing_Zipkin_); ok {
		return x.Zipkin
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Tracing) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Tracing_OneofMarshaler, _Tracing_OneofUnmarshaler, _Tracing_OneofSizer, []interface{}{
		(*Tracing_Zipkin_)(nil),
	}
}

func _Tracing_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Tracing)
	// type
	switch x := m.Type.(type) {
	case *Tracing_Zipkin_:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Zipkin); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Tracing.Type has unexpected type %T", x)
	}
	return nil
}

func _Tracing_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Tracing)
	switch tag {
	case 1: // type.zipkin
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Tracing_Zipkin)
		err := b.DecodeMessage(msg)
		m.Type = &Tracing_Zipkin_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Tracing_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Tracing)
	// type
	switch x := m.Type.(type) {
	case *Tracing_Zipkin_:
		s := proto.Size(x.Zipkin)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Zipkin defined configuration of Zipkin tracer.
type Tracing_Zipkin struct {
	// Address of Zipkin collector.
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tracing_Zipkin) Reset()         { *m = Tracing_Zipkin{} }
func (m *Tracing_Zipkin) String() string { return proto.CompactTextString(m) }
func (*Tracing_Zipkin) ProtoMessage()    {}
func (*Tracing_Zipkin) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b3cd8c92bbf6a, []int{2, 0}
}
func (m *Tracing_Zipkin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tracing_Zipkin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tracing_Zipkin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tracing_Zipkin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing_Zipkin.Merge(m, src)
}
func (m *Tracing_Zipkin) XXX_Size() int {
	return m.Size()
}
func (m *Tracing_Zipkin) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing_Zipkin.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing_Zipkin proto.InternalMessageInfo

func (m *Tracing_Zipkin) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Mesh)(nil), "konvoy.mesh.v1alpha1.Mesh")
	proto.RegisterType((*Mesh_Mtls)(nil), "konvoy.mesh.v1alpha1.Mesh.Mtls")
	proto.RegisterType((*CertificateAuthority)(nil), "konvoy.mesh.v1alpha1.CertificateAuthority")
	proto.RegisterType((*CertificateAuthority_Builtin)(nil), "konvoy.mesh.v1alpha1.CertificateAuthority.Builtin")
	proto.RegisterType((*Tracing)(nil), "konvoy.mesh.v1alpha1.Tracing")
	proto.RegisterType((*Tracing_Zipkin)(nil), "konvoy.mesh.v1alpha1.Tracing.Zipkin")
}

func init() { proto.RegisterFile("mesh/v1alpha1/mesh.proto", fileDescriptor_ae9b3cd8c92bbf6a) }

var fileDescriptor_ae9b3cd8c92bbf6a = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0x4d, 0x2d, 0xce,
	0xd0, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4, 0x07, 0xf1, 0xf4, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0x44, 0xb2, 0xf3, 0xf3, 0xca, 0xf2, 0x2b, 0xf5, 0xc0, 0x42, 0x30, 0x05, 0x4a,
	0x3b, 0x18, 0xb9, 0x58, 0x7c, 0x53, 0x8b, 0x33, 0x84, 0x8c, 0xb9, 0x58, 0x72, 0x4b, 0x72, 0x8a,
	0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xe4, 0xf5, 0xb0, 0xa9, 0xd6, 0x03, 0xa9, 0xd4, 0xf3,
	0x2d, 0xc9, 0x29, 0x0e, 0x02, 0x2b, 0x16, 0x32, 0xe7, 0x62, 0x2f, 0x29, 0x4a, 0x4c, 0xce, 0xcc,
	0x4b, 0x97, 0x60, 0x02, 0xeb, 0x93, 0xc5, 0xae, 0x2f, 0x04, 0xa2, 0x28, 0x08, 0xa6, 0x5a, 0xca,
	0x89, 0x8b, 0x05, 0x64, 0x8c, 0x90, 0x15, 0x17, 0x53, 0x72, 0x22, 0xd4, 0x4e, 0x2d, 0xec, 0x7a,
	0x9d, 0x53, 0x8b, 0x4a, 0x32, 0xd3, 0x32, 0x93, 0x13, 0x4b, 0x52, 0x1d, 0x4b, 0x4b, 0x32, 0xf2,
	0x8b, 0x32, 0x4b, 0x2a, 0x83, 0x98, 0x92, 0x13, 0x95, 0x2a, 0xb9, 0x44, 0xb0, 0xc9, 0x09, 0xf9,
	0x71, 0xb1, 0x27, 0x95, 0x66, 0xe6, 0x94, 0x64, 0xe6, 0x41, 0x0d, 0x36, 0x22, 0xde, 0x60, 0x3d,
	0x27, 0x88, 0x4e, 0x0f, 0x86, 0x20, 0x98, 0x21, 0x52, 0x9c, 0x5c, 0xec, 0x50, 0x51, 0x27, 0x36,
	0x2e, 0x96, 0x92, 0xca, 0x82, 0x54, 0xa5, 0x52, 0x2e, 0x76, 0xa8, 0x97, 0x84, 0xec, 0xb8, 0xd8,
	0xaa, 0x32, 0x0b, 0xb2, 0xe1, 0x96, 0xa9, 0xe0, 0x0d, 0x01, 0xbd, 0x28, 0xb0, 0x5a, 0x0f, 0x86,
	0x20, 0xa8, 0x2e, 0x29, 0x25, 0x2e, 0x36, 0x88, 0x98, 0x90, 0x04, 0x17, 0x7b, 0x62, 0x4a, 0x4a,
	0x51, 0x6a, 0x31, 0x24, 0x12, 0x38, 0x83, 0x60, 0x5c, 0x98, 0xb5, 0x4e, 0x62, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0x63, 0x14, 0x07, 0xcc, 0xf0, 0x24, 0x36,
	0x70, 0x0c, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x98, 0xfd, 0x28, 0xfd, 0x01, 0x00,
	0x00,
}

func (m *Mesh) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mesh) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Mtls != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMesh(dAtA, i, uint64(m.Mtls.Size()))
		n1, err := m.Mtls.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Tracing != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMesh(dAtA, i, uint64(m.Tracing.Size()))
		n2, err := m.Tracing.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Mesh_Mtls) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mesh_Mtls) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Ca != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMesh(dAtA, i, uint64(m.Ca.Size()))
		n3, err := m.Ca.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CertificateAuthority) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CertificateAuthority) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != nil {
		nn4, err := m.Type.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn4
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CertificateAuthority_Builtin_) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Builtin != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMesh(dAtA, i, uint64(m.Builtin.Size()))
		n5, err := m.Builtin.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func (m *CertificateAuthority_Builtin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CertificateAuthority_Builtin) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Tracing) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tracing) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != nil {
		nn6, err := m.Type.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn6
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Tracing_Zipkin_) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Zipkin != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMesh(dAtA, i, uint64(m.Zipkin.Size()))
		n7, err := m.Zipkin.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}
func (m *Tracing_Zipkin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tracing_Zipkin) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMesh(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMesh(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Mesh) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Mtls != nil {
		l = m.Mtls.Size()
		n += 1 + l + sovMesh(uint64(l))
	}
	if m.Tracing != nil {
		l = m.Tracing.Size()
		n += 1 + l + sovMesh(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Mesh_Mtls) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ca != nil {
		l = m.Ca.Size()
		n += 1 + l + sovMesh(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CertificateAuthority) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != nil {
		n += m.Type.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CertificateAuthority_Builtin_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Builtin != nil {
		l = m.Builtin.Size()
		n += 1 + l + sovMesh(uint64(l))
	}
	return n
}
func (m *CertificateAuthority_Builtin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Tracing) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != nil {
		n += m.Type.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Tracing_Zipkin_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Zipkin != nil {
		l = m.Zipkin.Size()
		n += 1 + l + sovMesh(uint64(l))
	}
	return n
}
func (m *Tracing_Zipkin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMesh(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMesh(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMesh(x uint64) (n int) {
	return sovMesh(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Mesh) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMesh
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Mesh: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mesh: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mtls", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMesh
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMesh
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Mtls == nil {
				m.Mtls = &Mesh_Mtls{}
			}
			if err := m.Mtls.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tracing", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMesh
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMesh
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tracing == nil {
				m.Tracing = &Tracing{}
			}
			if err := m.Tracing.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMesh(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMesh
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMesh
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Mesh_Mtls) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMesh
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Mtls: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mtls: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ca", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMesh
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMesh
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ca == nil {
				m.Ca = &CertificateAuthority{}
			}
			if err := m.Ca.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMesh(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMesh
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMesh
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CertificateAuthority) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMesh
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CertificateAuthority: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertificateAuthority: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Builtin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMesh
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMesh
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &CertificateAuthority_Builtin{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Type = &CertificateAuthority_Builtin_{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMesh(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMesh
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMesh
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CertificateAuthority_Builtin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMesh
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Builtin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Builtin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMesh(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMesh
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMesh
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Tracing) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMesh
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Tracing: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tracing: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Zipkin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMesh
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMesh
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Tracing_Zipkin{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Type = &Tracing_Zipkin_{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMesh(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMesh
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMesh
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Tracing_Zipkin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMesh
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Zipkin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Zipkin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMesh
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMesh
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMesh(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMesh
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMesh
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMesh(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMesh
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMesh
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthMesh
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMesh
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMesh(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthMesh
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMesh = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMesh   = fmt.Errorf("proto: integer overflow")
)
