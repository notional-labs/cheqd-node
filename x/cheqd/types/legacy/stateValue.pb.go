// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cheqd/did/v1/stateValue.proto

package legacy

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type StateValue struct {
	Data     *types.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Metadata *Metadata  `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *StateValue) Reset()         { *m = StateValue{} }
func (m *StateValue) String() string { return proto.CompactTextString(m) }
func (*StateValue) ProtoMessage()    {}
func (*StateValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_fadc6edf051fd79c, []int{0}
}
func (m *StateValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StateValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StateValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StateValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateValue.Merge(m, src)
}
func (m *StateValue) XXX_Size() int {
	return m.Size()
}
func (m *StateValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StateValue.DiscardUnknown(m)
}

var xxx_messageInfo_StateValue proto.InternalMessageInfo

func (m *StateValue) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *StateValue) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// metadata
type Metadata struct {
	Created     string   `protobuf:"bytes,1,opt,name=created,proto3" json:"created,omitempty"`
	Updated     string   `protobuf:"bytes,2,opt,name=updated,proto3" json:"updated,omitempty"`
	Deactivated bool     `protobuf:"varint,3,opt,name=deactivated,proto3" json:"deactivated,omitempty"`
	VersionId   string   `protobuf:"bytes,4,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	Resources   []string `protobuf:"bytes,5,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_fadc6edf051fd79c, []int{1}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Metadata) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

func (m *Metadata) GetDeactivated() bool {
	if m != nil {
		return m.Deactivated
	}
	return false
}

func (m *Metadata) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *Metadata) GetResources() []string {
	if m != nil {
		return m.Resources
	}
	return nil
}

func init() {
	proto.RegisterType((*StateValue)(nil), "cheqdid.cheqdnode.cheqd.v1.StateValue")
	proto.RegisterType((*Metadata)(nil), "cheqdid.cheqdnode.cheqd.v1.Metadata")
}

func init() { proto.RegisterFile("cheqd/did/v1/stateValue.proto", fileDescriptor_fadc6edf051fd79c) }

var fileDescriptor_fadc6edf051fd79c = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbd, 0x4e, 0xfb, 0x30,
	0x14, 0xc5, 0xeb, 0xb6, 0xff, 0x3f, 0xad, 0xbb, 0x59, 0x0c, 0xa1, 0xa2, 0x56, 0x54, 0x31, 0x74,
	0x00, 0x5b, 0x85, 0x17, 0x00, 0x06, 0x24, 0x06, 0x96, 0x20, 0x31, 0xb0, 0x20, 0x37, 0xbe, 0xb4,
	0x96, 0xda, 0xb8, 0x24, 0x4e, 0xd4, 0xbe, 0x05, 0x2f, 0xc0, 0xfb, 0x30, 0x76, 0x64, 0x44, 0xc9,
	0x8b, 0xa0, 0xdc, 0x7c, 0x88, 0x85, 0xc5, 0xf6, 0x3d, 0xbf, 0x73, 0xfc, 0x75, 0xe9, 0x24, 0x5c,
	0xc1, 0x9b, 0x96, 0xda, 0x68, 0x99, 0xcd, 0x65, 0xe2, 0x94, 0x83, 0x27, 0xb5, 0x4e, 0x41, 0x6c,
	0x63, 0xeb, 0x2c, 0x1b, 0x23, 0x36, 0x5a, 0xe0, 0x1c, 0x59, 0x0d, 0xd5, 0x4a, 0x64, 0xf3, 0xf1,
	0xc9, 0xd2, 0xda, 0xe5, 0x1a, 0x24, 0x3a, 0x17, 0xe9, 0xab, 0x54, 0xd1, 0xbe, 0x8a, 0x4d, 0x77,
	0x94, 0x3e, 0xb6, 0x5b, 0xb1, 0x19, 0xed, 0x6b, 0xe5, 0x94, 0x47, 0x7c, 0x32, 0x1b, 0x5d, 0x1e,
	0x8b, 0x2a, 0x27, 0x9a, 0x9c, 0xb8, 0x89, 0xf6, 0x01, 0x3a, 0xd8, 0x35, 0x1d, 0x6c, 0xc0, 0x29,
	0x74, 0x77, 0xd1, 0x7d, 0x26, 0xfe, 0xbe, 0x81, 0x78, 0xa8, 0xbd, 0x41, 0x9b, 0x9a, 0x7e, 0x10,
	0x3a, 0x68, 0x64, 0xe6, 0xd1, 0xa3, 0x30, 0x06, 0xe5, 0x40, 0xe3, 0xd9, 0xc3, 0xa0, 0x29, 0x4b,
	0x92, 0x6e, 0x35, 0x92, 0x6e, 0x45, 0xea, 0x92, 0xf9, 0x74, 0xa4, 0x41, 0x85, 0xce, 0x64, 0x48,
	0x7b, 0x3e, 0x99, 0x0d, 0x82, 0xdf, 0x12, 0x9b, 0x50, 0x9a, 0x41, 0x9c, 0x18, 0x1b, 0xbd, 0x18,
	0xed, 0xf5, 0x31, 0x3e, 0xac, 0x95, 0x7b, 0xcd, 0x4e, 0xe9, 0x30, 0x86, 0xc4, 0xa6, 0x71, 0x08,
	0x89, 0xf7, 0xcf, 0xef, 0x95, 0xb4, 0x15, 0x6e, 0xef, 0x3e, 0x73, 0x4e, 0x0e, 0x39, 0x27, 0xdf,
	0x39, 0x27, 0xef, 0x05, 0xef, 0x1c, 0x0a, 0xde, 0xf9, 0x2a, 0x78, 0xe7, 0xf9, 0x7c, 0x69, 0xdc,
	0x2a, 0x5d, 0x88, 0xd0, 0x6e, 0x64, 0xd5, 0x14, 0x1c, 0x2f, 0xca, 0x27, 0xcb, 0x5d, 0x2d, 0xb9,
	0xfd, 0x16, 0x12, 0x99, 0xcd, 0x17, 0xff, 0xf1, 0xf7, 0xae, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xb1, 0x56, 0x86, 0x70, 0xc0, 0x01, 0x00, 0x00,
}

func (m *StateValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StateValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStateValue(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStateValue(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Resources) > 0 {
		for iNdEx := len(m.Resources) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Resources[iNdEx])
			copy(dAtA[i:], m.Resources[iNdEx])
			i = encodeVarintStateValue(dAtA, i, uint64(len(m.Resources[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.VersionId) > 0 {
		i -= len(m.VersionId)
		copy(dAtA[i:], m.VersionId)
		i = encodeVarintStateValue(dAtA, i, uint64(len(m.VersionId)))
		i--
		dAtA[i] = 0x22
	}
	if m.Deactivated {
		i--
		if m.Deactivated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Updated) > 0 {
		i -= len(m.Updated)
		copy(dAtA[i:], m.Updated)
		i = encodeVarintStateValue(dAtA, i, uint64(len(m.Updated)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Created) > 0 {
		i -= len(m.Created)
		copy(dAtA[i:], m.Created)
		i = encodeVarintStateValue(dAtA, i, uint64(len(m.Created)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStateValue(dAtA []byte, offset int, v uint64) int {
	offset -= sovStateValue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StateValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovStateValue(uint64(l))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovStateValue(uint64(l))
	}
	return n
}

func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Created)
	if l > 0 {
		n += 1 + l + sovStateValue(uint64(l))
	}
	l = len(m.Updated)
	if l > 0 {
		n += 1 + l + sovStateValue(uint64(l))
	}
	if m.Deactivated {
		n += 2
	}
	l = len(m.VersionId)
	if l > 0 {
		n += 1 + l + sovStateValue(uint64(l))
	}
	if len(m.Resources) > 0 {
		for _, s := range m.Resources {
			l = len(s)
			n += 1 + l + sovStateValue(uint64(l))
		}
	}
	return n
}

func sovStateValue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStateValue(x uint64) (n int) {
	return sovStateValue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StateValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStateValue
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
			return fmt.Errorf("proto: StateValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateValue
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
				return ErrInvalidLengthStateValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateValue
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
				return ErrInvalidLengthStateValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStateValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStateValue
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStateValue
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateValue
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
				return ErrInvalidLengthStateValue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Created = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updated", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateValue
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
				return ErrInvalidLengthStateValue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Updated = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deactivated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateValue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deactivated = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VersionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateValue
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
				return ErrInvalidLengthStateValue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VersionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateValue
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
				return ErrInvalidLengthStateValue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resources = append(m.Resources, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStateValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStateValue
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStateValue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStateValue
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
					return 0, ErrIntOverflowStateValue
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStateValue
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
				return 0, ErrInvalidLengthStateValue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStateValue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStateValue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStateValue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStateValue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStateValue = fmt.Errorf("proto: unexpected end of group")
)
