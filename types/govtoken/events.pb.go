// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cysicmint/govtoken/v1/events.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// EventMint
type EventMint struct {
	// owner
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// amount
	Amount cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount" yaml:"amount"`
}

func (m *EventMint) Reset()         { *m = EventMint{} }
func (m *EventMint) String() string { return proto.CompactTextString(m) }
func (*EventMint) ProtoMessage()    {}
func (*EventMint) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b02807b0d31645, []int{0}
}
func (m *EventMint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventMint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventMint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventMint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMint.Merge(m, src)
}
func (m *EventMint) XXX_Size() int {
	return m.Size()
}
func (m *EventMint) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMint.DiscardUnknown(m)
}

var xxx_messageInfo_EventMint proto.InternalMessageInfo

func (m *EventMint) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

// EventBurn
type EventBurn struct {
	// burner
	Burner string `protobuf:"bytes,1,opt,name=burner,proto3" json:"burner,omitempty"`
	// amount
	Amount cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount" yaml:"amount"`
}

func (m *EventBurn) Reset()         { *m = EventBurn{} }
func (m *EventBurn) String() string { return proto.CompactTextString(m) }
func (*EventBurn) ProtoMessage()    {}
func (*EventBurn) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b02807b0d31645, []int{1}
}
func (m *EventBurn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBurn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBurn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBurn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBurn.Merge(m, src)
}
func (m *EventBurn) XXX_Size() int {
	return m.Size()
}
func (m *EventBurn) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBurn.DiscardUnknown(m)
}

var xxx_messageInfo_EventBurn proto.InternalMessageInfo

func (m *EventBurn) GetBurner() string {
	if m != nil {
		return m.Burner
	}
	return ""
}

// EventOwnerChanged
type EventOwnerChanged struct {
	// old_owner
	OldOwner string `protobuf:"bytes,1,opt,name=old_owner,json=oldOwner,proto3" json:"old_owner,omitempty"`
	// new_owner
	NewOwner string `protobuf:"bytes,2,opt,name=new_owner,json=newOwner,proto3" json:"new_owner,omitempty"`
}

func (m *EventOwnerChanged) Reset()         { *m = EventOwnerChanged{} }
func (m *EventOwnerChanged) String() string { return proto.CompactTextString(m) }
func (*EventOwnerChanged) ProtoMessage()    {}
func (*EventOwnerChanged) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b02807b0d31645, []int{2}
}
func (m *EventOwnerChanged) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventOwnerChanged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventOwnerChanged.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventOwnerChanged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventOwnerChanged.Merge(m, src)
}
func (m *EventOwnerChanged) XXX_Size() int {
	return m.Size()
}
func (m *EventOwnerChanged) XXX_DiscardUnknown() {
	xxx_messageInfo_EventOwnerChanged.DiscardUnknown(m)
}

var xxx_messageInfo_EventOwnerChanged proto.InternalMessageInfo

func (m *EventOwnerChanged) GetOldOwner() string {
	if m != nil {
		return m.OldOwner
	}
	return ""
}

func (m *EventOwnerChanged) GetNewOwner() string {
	if m != nil {
		return m.NewOwner
	}
	return ""
}

func init() {
	proto.RegisterType((*EventMint)(nil), "cysicmint.govtoken.v1.EventMint")
	proto.RegisterType((*EventBurn)(nil), "cysicmint.govtoken.v1.EventBurn")
	proto.RegisterType((*EventOwnerChanged)(nil), "cysicmint.govtoken.v1.EventOwnerChanged")
}

func init() {
	proto.RegisterFile("cysicmint/govtoken/v1/events.proto", fileDescriptor_f4b02807b0d31645)
}

var fileDescriptor_f4b02807b0d31645 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x93, 0x82, 0xc5, 0x0e, 0xb8, 0x30, 0xb4, 0x52, 0x2a, 0xa4, 0x12, 0x5c, 0xb8, 0x71,
	0x86, 0xa2, 0x2b, 0x97, 0x15, 0x05, 0x17, 0x55, 0xe8, 0xd2, 0x8d, 0xe4, 0x67, 0x48, 0x87, 0x64,
	0xee, 0x2d, 0x99, 0x49, 0x62, 0xdf, 0xc2, 0xc7, 0xea, 0xb2, 0x4b, 0x71, 0x51, 0xa4, 0x7d, 0x03,
	0x9f, 0x40, 0x32, 0x89, 0xad, 0x0b, 0x97, 0xee, 0xce, 0xdc, 0x73, 0x0e, 0x1f, 0x77, 0x2e, 0xf1,
	0xc2, 0x85, 0x12, 0xa1, 0x14, 0xa0, 0x59, 0x8c, 0x85, 0xc6, 0x84, 0x03, 0x2b, 0x46, 0x8c, 0x17,
	0x1c, 0xb4, 0xa2, 0xf3, 0x0c, 0x35, 0x3a, 0xbd, 0x5d, 0x86, 0xfe, 0x64, 0x68, 0x31, 0x1a, 0x9c,
	0xff, 0x5d, 0xdd, 0x45, 0x4c, 0x79, 0xd0, 0x8d, 0x31, 0x46, 0x23, 0x59, 0xa5, 0xea, 0xa9, 0x27,
	0x48, 0xe7, 0xae, 0x42, 0x4c, 0x04, 0x68, 0xa7, 0x4b, 0x0e, 0xb0, 0x04, 0x9e, 0xf5, 0xed, 0x33,
	0xfb, 0xa2, 0x33, 0xad, 0x1f, 0xce, 0x3d, 0x69, 0xfb, 0x12, 0x73, 0xd0, 0xfd, 0x56, 0x35, 0x1e,
	0xd3, 0xe5, 0x7a, 0x68, 0x7d, 0xac, 0x87, 0xbd, 0x10, 0x95, 0x44, 0xa5, 0xa2, 0x84, 0x0a, 0x64,
	0xd2, 0xd7, 0x33, 0xfa, 0x00, 0xfa, 0x6b, 0x3d, 0x3c, 0x5a, 0xf8, 0x32, 0xbd, 0xf1, 0xea, 0x92,
	0x37, 0x6d, 0xda, 0x5e, 0xd2, 0xa0, 0xc6, 0x79, 0x06, 0xce, 0x09, 0x69, 0x07, 0x79, 0xb6, 0x67,
	0x35, 0xaf, 0x7f, 0x83, 0x4d, 0xc8, 0xb1, 0x81, 0x3d, 0x55, 0x2b, 0xdc, 0xce, 0x7c, 0x88, 0x79,
	0xe4, 0x9c, 0x92, 0x0e, 0xa6, 0xd1, 0xcb, 0xef, 0x1d, 0x0f, 0x31, 0x8d, 0x4c, 0xa6, 0x32, 0x81,
	0x97, 0x8d, 0xd9, 0xaa, 0x4d, 0xe0, 0xa5, 0x31, 0xc7, 0x8f, 0xcb, 0x8d, 0x6b, 0xaf, 0x36, 0xae,
	0xfd, 0xb9, 0x71, 0xed, 0xb7, 0xad, 0x6b, 0xad, 0xb6, 0xae, 0xf5, 0xbe, 0x75, 0xad, 0xe7, 0xeb,
	0x58, 0xe8, 0x59, 0x1e, 0xd0, 0x10, 0x25, 0x33, 0x77, 0xb8, 0x4c, 0xfd, 0x40, 0x35, 0x12, 0xb8,
	0x2e, 0x31, 0x4b, 0xd8, 0xeb, 0xfe, 0x30, 0x7a, 0x31, 0xe7, 0x2a, 0x68, 0x9b, 0xdf, 0xbf, 0xfa,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x96, 0x98, 0x37, 0xcf, 0xf6, 0x01, 0x00, 0x00,
}

func (m *EventMint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventMint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventMint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventBurn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBurn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBurn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Burner) > 0 {
		i -= len(m.Burner)
		copy(dAtA[i:], m.Burner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Burner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventOwnerChanged) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventOwnerChanged) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventOwnerChanged) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewOwner) > 0 {
		i -= len(m.NewOwner)
		copy(dAtA[i:], m.NewOwner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NewOwner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OldOwner) > 0 {
		i -= len(m.OldOwner)
		copy(dAtA[i:], m.OldOwner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.OldOwner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventMint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func (m *EventBurn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Burner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func (m *EventOwnerChanged) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OldOwner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NewOwner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventMint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventMint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventMint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventBurn) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventBurn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBurn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Burner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Burner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventOwnerChanged) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventOwnerChanged: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventOwnerChanged: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OldOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
