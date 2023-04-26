// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spn/monitoringc/launch_id_from_channel_id.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type LaunchIDFromChannelID struct {
	ChannelID string `protobuf:"bytes,1,opt,name=channelID,proto3" json:"channelID,omitempty"`
	LaunchID  uint64 `protobuf:"varint,2,opt,name=launchID,proto3" json:"launchID,omitempty"`
}

func (m *LaunchIDFromChannelID) Reset()         { *m = LaunchIDFromChannelID{} }
func (m *LaunchIDFromChannelID) String() string { return proto.CompactTextString(m) }
func (*LaunchIDFromChannelID) ProtoMessage()    {}
func (*LaunchIDFromChannelID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a03e8149f64ce63c, []int{0}
}
func (m *LaunchIDFromChannelID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LaunchIDFromChannelID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LaunchIDFromChannelID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LaunchIDFromChannelID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchIDFromChannelID.Merge(m, src)
}
func (m *LaunchIDFromChannelID) XXX_Size() int {
	return m.Size()
}
func (m *LaunchIDFromChannelID) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchIDFromChannelID.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchIDFromChannelID proto.InternalMessageInfo

func (m *LaunchIDFromChannelID) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *LaunchIDFromChannelID) GetLaunchID() uint64 {
	if m != nil {
		return m.LaunchID
	}
	return 0
}

func init() {
	proto.RegisterType((*LaunchIDFromChannelID)(nil), "spn.monitoringc.LaunchIDFromChannelID")
}

func init() {
	proto.RegisterFile("spn/monitoringc/launch_id_from_channel_id.proto", fileDescriptor_a03e8149f64ce63c)
}

var fileDescriptor_a03e8149f64ce63c = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2f, 0x2e, 0xc8, 0xd3,
	0xcf, 0xcd, 0xcf, 0xcb, 0x2c, 0xc9, 0x2f, 0xca, 0xcc, 0x4b, 0x4f, 0xd6, 0xcf, 0x49, 0x2c, 0xcd,
	0x4b, 0xce, 0x88, 0xcf, 0x4c, 0x89, 0x4f, 0x2b, 0xca, 0xcf, 0x8d, 0x4f, 0xce, 0x48, 0xcc, 0xcb,
	0x4b, 0xcd, 0x89, 0xcf, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2f, 0x2e, 0xc8,
	0xd3, 0x43, 0xd2, 0xa0, 0x14, 0xc8, 0x25, 0xea, 0x03, 0xd6, 0xe3, 0xe9, 0xe2, 0x56, 0x94, 0x9f,
	0xeb, 0x0c, 0xd1, 0xe0, 0xe9, 0x22, 0x24, 0xc3, 0xc5, 0x99, 0x0c, 0xe3, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0x21, 0x04, 0x84, 0xa4, 0xb8, 0x38, 0x72, 0xa0, 0xda, 0x24, 0x98, 0x14, 0x18,
	0x35, 0x58, 0x82, 0xe0, 0x7c, 0x27, 0xf7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63,
	0x88, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x49, 0xcd,
	0x4b, 0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0x01, 0x7b, 0xa2, 0x02, 0xc5, 0x1b, 0x25, 0x95, 0x05,
	0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x37, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x30, 0xd3, 0x09,
	0xf0, 0xe6, 0x00, 0x00, 0x00,
}

func (m *LaunchIDFromChannelID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LaunchIDFromChannelID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LaunchIDFromChannelID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LaunchID != 0 {
		i = encodeVarintLaunchIdFromChannelId(dAtA, i, uint64(m.LaunchID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintLaunchIdFromChannelId(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLaunchIdFromChannelId(dAtA []byte, offset int, v uint64) int {
	offset -= sovLaunchIdFromChannelId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LaunchIDFromChannelID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovLaunchIdFromChannelId(uint64(l))
	}
	if m.LaunchID != 0 {
		n += 1 + sovLaunchIdFromChannelId(uint64(m.LaunchID))
	}
	return n
}

func sovLaunchIdFromChannelId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLaunchIdFromChannelId(x uint64) (n int) {
	return sovLaunchIdFromChannelId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LaunchIDFromChannelID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLaunchIdFromChannelId
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
			return fmt.Errorf("proto: LaunchIDFromChannelID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LaunchIDFromChannelID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLaunchIdFromChannelId
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
				return ErrInvalidLengthLaunchIdFromChannelId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLaunchIdFromChannelId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchID", wireType)
			}
			m.LaunchID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLaunchIdFromChannelId
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LaunchID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLaunchIdFromChannelId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLaunchIdFromChannelId
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
func skipLaunchIdFromChannelId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLaunchIdFromChannelId
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
					return 0, ErrIntOverflowLaunchIdFromChannelId
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
					return 0, ErrIntOverflowLaunchIdFromChannelId
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
				return 0, ErrInvalidLengthLaunchIdFromChannelId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLaunchIdFromChannelId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLaunchIdFromChannelId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLaunchIdFromChannelId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLaunchIdFromChannelId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLaunchIdFromChannelId = fmt.Errorf("proto: unexpected end of group")
)
