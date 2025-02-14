// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/change_min_on_chain_balance_proposal.proto

package types

import (
	fmt "fmt"
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

type ChangeMinOnChainBalanceProposal struct {
	Title             string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description       string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ChainReferenceID  string `protobuf:"bytes,3,opt,name=chainReferenceID,proto3" json:"chainReferenceID,omitempty"`
	MinOnChainBalance string `protobuf:"bytes,4,opt,name=minOnChainBalance,proto3" json:"minOnChainBalance,omitempty"`
}

func (m *ChangeMinOnChainBalanceProposal) Reset()         { *m = ChangeMinOnChainBalanceProposal{} }
func (m *ChangeMinOnChainBalanceProposal) String() string { return proto.CompactTextString(m) }
func (*ChangeMinOnChainBalanceProposal) ProtoMessage()    {}
func (*ChangeMinOnChainBalanceProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c4541a20a7d2145, []int{0}
}
func (m *ChangeMinOnChainBalanceProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChangeMinOnChainBalanceProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChangeMinOnChainBalanceProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChangeMinOnChainBalanceProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeMinOnChainBalanceProposal.Merge(m, src)
}
func (m *ChangeMinOnChainBalanceProposal) XXX_Size() int {
	return m.Size()
}
func (m *ChangeMinOnChainBalanceProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeMinOnChainBalanceProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeMinOnChainBalanceProposal proto.InternalMessageInfo

func (m *ChangeMinOnChainBalanceProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ChangeMinOnChainBalanceProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ChangeMinOnChainBalanceProposal) GetChainReferenceID() string {
	if m != nil {
		return m.ChainReferenceID
	}
	return ""
}

func (m *ChangeMinOnChainBalanceProposal) GetMinOnChainBalance() string {
	if m != nil {
		return m.MinOnChainBalance
	}
	return ""
}

func init() {
	proto.RegisterType((*ChangeMinOnChainBalanceProposal)(nil), "palomachain.paloma.evm.ChangeMinOnChainBalanceProposal")
}

func init() {
	proto.RegisterFile("evm/change_min_on_chain_balance_proposal.proto", fileDescriptor_9c4541a20a7d2145)
}

var fileDescriptor_9c4541a20a7d2145 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0x2d, 0xcb, 0xd5,
	0x4f, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0x8d, 0xcf, 0xcd, 0xcc, 0x8b, 0xcf, 0xcf, 0x8b, 0x4f, 0xce,
	0x48, 0xcc, 0xcc, 0x8b, 0x4f, 0x4a, 0xcc, 0x49, 0xcc, 0x4b, 0x4e, 0x8d, 0x2f, 0x28, 0xca, 0x2f,
	0xc8, 0x2f, 0x4e, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2b, 0x48, 0xcc, 0xc9,
	0xcf, 0x4d, 0x04, 0xab, 0xd1, 0x83, 0xb0, 0x41, 0x46, 0x28, 0x6d, 0x66, 0xe4, 0x92, 0x77, 0x06,
	0x1b, 0xe3, 0x9b, 0x99, 0xe7, 0x9f, 0xe7, 0x0c, 0x92, 0x77, 0x82, 0x18, 0x11, 0x00, 0x35, 0x41,
	0x48, 0x84, 0x8b, 0xb5, 0x24, 0xb3, 0x24, 0x27, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08,
	0xc2, 0x11, 0x52, 0xe0, 0xe2, 0x4e, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc, 0xcf,
	0x93, 0x60, 0x02, 0xcb, 0x21, 0x0b, 0x09, 0x69, 0x71, 0x09, 0x80, 0xed, 0x0b, 0x4a, 0x4d, 0x4b,
	0x2d, 0x4a, 0xcd, 0x4b, 0x4e, 0xf5, 0x74, 0x91, 0x60, 0x06, 0x2b, 0xc3, 0x10, 0x17, 0xd2, 0xe1,
	0x12, 0xcc, 0x45, 0x77, 0x80, 0x04, 0x0b, 0x58, 0x31, 0xa6, 0x84, 0x93, 0xf3, 0x89, 0x47, 0x72,
	0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7,
	0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x69, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0x23, 0x79, 0x19, 0xca, 0xd6, 0xaf, 0xd0, 0x07, 0x85, 0x5b, 0x49, 0x65, 0x41,
	0x6a, 0x71, 0x12, 0x1b, 0x38, 0x64, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x12, 0x1c,
	0x9f, 0x4b, 0x01, 0x00, 0x00,
}

func (m *ChangeMinOnChainBalanceProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChangeMinOnChainBalanceProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChangeMinOnChainBalanceProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinOnChainBalance) > 0 {
		i -= len(m.MinOnChainBalance)
		copy(dAtA[i:], m.MinOnChainBalance)
		i = encodeVarintChangeMinOnChainBalanceProposal(dAtA, i, uint64(len(m.MinOnChainBalance)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChainReferenceID) > 0 {
		i -= len(m.ChainReferenceID)
		copy(dAtA[i:], m.ChainReferenceID)
		i = encodeVarintChangeMinOnChainBalanceProposal(dAtA, i, uint64(len(m.ChainReferenceID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintChangeMinOnChainBalanceProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintChangeMinOnChainBalanceProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintChangeMinOnChainBalanceProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovChangeMinOnChainBalanceProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChangeMinOnChainBalanceProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovChangeMinOnChainBalanceProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovChangeMinOnChainBalanceProposal(uint64(l))
	}
	l = len(m.ChainReferenceID)
	if l > 0 {
		n += 1 + l + sovChangeMinOnChainBalanceProposal(uint64(l))
	}
	l = len(m.MinOnChainBalance)
	if l > 0 {
		n += 1 + l + sovChangeMinOnChainBalanceProposal(uint64(l))
	}
	return n
}

func sovChangeMinOnChainBalanceProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChangeMinOnChainBalanceProposal(x uint64) (n int) {
	return sovChangeMinOnChainBalanceProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChangeMinOnChainBalanceProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChangeMinOnChainBalanceProposal
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
			return fmt.Errorf("proto: ChangeMinOnChainBalanceProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChangeMinOnChainBalanceProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChangeMinOnChainBalanceProposal
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
				return ErrInvalidLengthChangeMinOnChainBalanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChangeMinOnChainBalanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChangeMinOnChainBalanceProposal
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
				return ErrInvalidLengthChangeMinOnChainBalanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChangeMinOnChainBalanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainReferenceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChangeMinOnChainBalanceProposal
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
				return ErrInvalidLengthChangeMinOnChainBalanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChangeMinOnChainBalanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainReferenceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinOnChainBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChangeMinOnChainBalanceProposal
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
				return ErrInvalidLengthChangeMinOnChainBalanceProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChangeMinOnChainBalanceProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinOnChainBalance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChangeMinOnChainBalanceProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChangeMinOnChainBalanceProposal
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
func skipChangeMinOnChainBalanceProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChangeMinOnChainBalanceProposal
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
					return 0, ErrIntOverflowChangeMinOnChainBalanceProposal
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
					return 0, ErrIntOverflowChangeMinOnChainBalanceProposal
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
				return 0, ErrInvalidLengthChangeMinOnChainBalanceProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChangeMinOnChainBalanceProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChangeMinOnChainBalanceProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChangeMinOnChainBalanceProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChangeMinOnChainBalanceProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChangeMinOnChainBalanceProposal = fmt.Errorf("proto: unexpected end of group")
)
