// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cheqd/did/v2/fee.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// FeeParams defines the parameters for the `did` module fixed fee.
type FeeParams struct {
	// Tx types define the fixed fee each for the `did` module.
	TxTypes    map[string]types.Coin                  `protobuf:"bytes,1,rep,name=tx_types,json=txTypes,proto3" json:"tx_types" yaml:"tx_types" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BurnFactor github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=burn_factor,json=burnFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"burn_factor"`
}

func (m *FeeParams) Reset()         { *m = FeeParams{} }
func (m *FeeParams) String() string { return proto.CompactTextString(m) }
func (*FeeParams) ProtoMessage()    {}
func (*FeeParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0cfbae270deaac7, []int{0}
}
func (m *FeeParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeParams.Merge(m, src)
}
func (m *FeeParams) XXX_Size() int {
	return m.Size()
}
func (m *FeeParams) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeParams.DiscardUnknown(m)
}

var xxx_messageInfo_FeeParams proto.InternalMessageInfo

func (m *FeeParams) GetTxTypes() map[string]types.Coin {
	if m != nil {
		return m.TxTypes
	}
	return nil
}

func init() {
	proto.RegisterType((*FeeParams)(nil), "cheqd.did.v2.FeeParams")
	proto.RegisterMapType((map[string]types.Coin)(nil), "cheqd.did.v2.FeeParams.TxTypesEntry")
}

func init() { proto.RegisterFile("cheqd/did/v2/fee.proto", fileDescriptor_b0cfbae270deaac7) }

var fileDescriptor_b0cfbae270deaac7 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x51, 0xcd, 0x4e, 0x32, 0x31,
	0x14, 0x9d, 0x42, 0xbe, 0x1f, 0x0a, 0x89, 0x66, 0x62, 0x14, 0x58, 0x14, 0x42, 0x8c, 0xc1, 0x05,
	0x6d, 0x18, 0x37, 0x86, 0xb8, 0x42, 0xc4, 0xad, 0x99, 0x60, 0x62, 0x4c, 0x0c, 0x99, 0x9f, 0x02,
	0x13, 0x98, 0x29, 0xce, 0x94, 0x09, 0xf3, 0x16, 0xee, 0x7d, 0x01, 0x1f, 0xc0, 0x87, 0x60, 0x49,
	0x5c, 0x19, 0x17, 0xc4, 0x0c, 0x6f, 0xe0, 0x13, 0x98, 0x69, 0xab, 0x61, 0xd3, 0xde, 0xde, 0x73,
	0xee, 0xb9, 0x27, 0xa7, 0xf0, 0xd0, 0x99, 0xd0, 0x47, 0x97, 0xb8, 0x9e, 0x4b, 0x62, 0x83, 0x8c,
	0x28, 0xc5, 0xf3, 0x90, 0x71, 0xa6, 0x97, 0x44, 0x1f, 0xbb, 0x9e, 0x8b, 0x63, 0xa3, 0x7a, 0x30,
	0x66, 0x63, 0x26, 0x00, 0x92, 0x55, 0x92, 0x53, 0x45, 0x0e, 0x8b, 0x7c, 0x16, 0x11, 0xdb, 0x8a,
	0x28, 0x89, 0xdb, 0x36, 0xe5, 0x56, 0x9b, 0x38, 0xcc, 0x0b, 0x14, 0x5e, 0x91, 0xf8, 0x50, 0x0e,
	0xca, 0x87, 0x84, 0x1a, 0xcf, 0x39, 0x58, 0xe8, 0x53, 0x7a, 0x63, 0x85, 0x96, 0x1f, 0xe9, 0x77,
	0xf0, 0x3f, 0x5f, 0x0e, 0x79, 0x32, 0xa7, 0x51, 0x19, 0xd4, 0xf3, 0xcd, 0xa2, 0x71, 0x8c, 0x77,
	0xf7, 0xe3, 0x5f, 0x2a, 0x1e, 0x2c, 0x07, 0x19, 0xed, 0x2a, 0xe0, 0x61, 0xd2, 0x3d, 0x5a, 0x6d,
	0x6a, 0xda, 0xd7, 0xa6, 0xb6, 0x97, 0x58, 0xfe, 0xac, 0xd3, 0xf8, 0xd1, 0x68, 0x98, 0xff, 0xb8,
	0xa4, 0xe9, 0x0f, 0xb0, 0x68, 0x2f, 0xc2, 0x60, 0x38, 0xb2, 0x1c, 0xce, 0xc2, 0x72, 0xae, 0x0e,
	0x9a, 0x85, 0xee, 0x45, 0x36, 0xf6, 0xb1, 0xa9, 0x9d, 0x8c, 0x3d, 0x3e, 0x59, 0xd8, 0xd8, 0x61,
	0xbe, 0x72, 0xa7, 0xae, 0x56, 0xe4, 0x4e, 0x89, 0x90, 0xc2, 0x3d, 0xea, 0xbc, 0xbd, 0xb6, 0xa0,
	0x32, 0xdf, 0xa3, 0x8e, 0x09, 0x33, 0xc1, 0xbe, 0xd0, 0xab, 0xde, 0xc2, 0xd2, 0xae, 0x21, 0x7d,
	0x1f, 0xe6, 0xa7, 0x34, 0x29, 0x83, 0x6c, 0x8d, 0x99, 0x95, 0x3a, 0x81, 0x7f, 0x62, 0x6b, 0xb6,
	0xa0, 0x62, 0x75, 0xd1, 0xa8, 0x60, 0xa5, 0x94, 0x65, 0x86, 0x55, 0x66, 0xf8, 0x92, 0x79, 0x81,
	0x29, 0x79, 0x9d, 0xdc, 0x39, 0xe8, 0x5e, 0xbf, 0xa4, 0x08, 0xac, 0x52, 0x04, 0xd6, 0x29, 0x02,
	0x9f, 0x29, 0x02, 0x4f, 0x5b, 0xa4, 0xad, 0xb7, 0x48, 0x7b, 0xdf, 0x22, 0xed, 0xfe, 0x74, 0xd7,
	0xb6, 0xf8, 0x3d, 0x71, 0xb6, 0x02, 0xe6, 0x52, 0xb2, 0x54, 0x2d, 0xe1, 0xde, 0xfe, 0x2b, 0xd2,
	0x3e, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x63, 0x5e, 0x4e, 0xb6, 0xe6, 0x01, 0x00, 0x00,
}

func (this *FeeParams) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FeeParams)
	if !ok {
		that2, ok := that.(FeeParams)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.TxTypes) != len(that1.TxTypes) {
		return false
	}
	for i := range this.TxTypes {
		a := this.TxTypes[i]
		b := that1.TxTypes[i]
		if !(&a).Equal(&b) {
			return false
		}
	}
	if !this.BurnFactor.Equal(that1.BurnFactor) {
		return false
	}
	return true
}
func (m *FeeParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BurnFactor.Size()
		i -= size
		if _, err := m.BurnFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.TxTypes) > 0 {
		for k := range m.TxTypes {
			v := m.TxTypes[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintFee(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintFee(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FeeParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TxTypes) > 0 {
		for k, v := range m.TxTypes {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovFee(uint64(len(k))) + 1 + l + sovFee(uint64(l))
			n += mapEntrySize + 1 + sovFee(uint64(mapEntrySize))
		}
	}
	l = m.BurnFactor.Size()
	n += 1 + l + sovFee(uint64(l))
	return n
}

func sovFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFee(x uint64) (n int) {
	return sovFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FeeParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
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
			return fmt.Errorf("proto: FeeParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxTypes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TxTypes == nil {
				m.TxTypes = make(map[string]types.Coin)
			}
			var mapkey string
			mapvalue := &types.Coin{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFee
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFee
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthFee
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthFee
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFee
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthFee
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthFee
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &types.Coin{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipFee(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthFee
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.TxTypes[mapkey] = *mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BurnFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
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
func skipFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFee
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
					return 0, ErrIntOverflowFee
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
					return 0, ErrIntOverflowFee
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
				return 0, ErrInvalidLengthFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFee = fmt.Errorf("proto: unexpected end of group")
)
