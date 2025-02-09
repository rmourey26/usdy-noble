// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aura/blocklist/v1/events.proto

package blocklist

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

// OwnershipTransferStarted is emitted whenever an ownership transfer is started.
type OwnershipTransferStarted struct {
	// previous_owner is the address of the previous owner.
	PreviousOwner string `protobuf:"bytes,1,opt,name=previous_owner,json=previousOwner,proto3" json:"previous_owner,omitempty"`
	// new_owner is the address of the new owner.
	NewOwner string `protobuf:"bytes,2,opt,name=new_owner,json=newOwner,proto3" json:"new_owner,omitempty"`
}

func (m *OwnershipTransferStarted) Reset()         { *m = OwnershipTransferStarted{} }
func (m *OwnershipTransferStarted) String() string { return proto.CompactTextString(m) }
func (*OwnershipTransferStarted) ProtoMessage()    {}
func (*OwnershipTransferStarted) Descriptor() ([]byte, []int) {
	return fileDescriptor_d01e1c8e4c279093, []int{0}
}
func (m *OwnershipTransferStarted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OwnershipTransferStarted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OwnershipTransferStarted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OwnershipTransferStarted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnershipTransferStarted.Merge(m, src)
}
func (m *OwnershipTransferStarted) XXX_Size() int {
	return m.Size()
}
func (m *OwnershipTransferStarted) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnershipTransferStarted.DiscardUnknown(m)
}

var xxx_messageInfo_OwnershipTransferStarted proto.InternalMessageInfo

func (m *OwnershipTransferStarted) GetPreviousOwner() string {
	if m != nil {
		return m.PreviousOwner
	}
	return ""
}

func (m *OwnershipTransferStarted) GetNewOwner() string {
	if m != nil {
		return m.NewOwner
	}
	return ""
}

// OwnershipTransferStarted is emitted whenever an ownership transfer is finalized.
type OwnershipTransferred struct {
	// previous_owner is the address of the previous owner.
	PreviousOwner string `protobuf:"bytes,1,opt,name=previous_owner,json=previousOwner,proto3" json:"previous_owner,omitempty"`
	// new_owner is the address of the new owner.
	NewOwner string `protobuf:"bytes,2,opt,name=new_owner,json=newOwner,proto3" json:"new_owner,omitempty"`
}

func (m *OwnershipTransferred) Reset()         { *m = OwnershipTransferred{} }
func (m *OwnershipTransferred) String() string { return proto.CompactTextString(m) }
func (*OwnershipTransferred) ProtoMessage()    {}
func (*OwnershipTransferred) Descriptor() ([]byte, []int) {
	return fileDescriptor_d01e1c8e4c279093, []int{1}
}
func (m *OwnershipTransferred) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OwnershipTransferred) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OwnershipTransferred.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OwnershipTransferred) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnershipTransferred.Merge(m, src)
}
func (m *OwnershipTransferred) XXX_Size() int {
	return m.Size()
}
func (m *OwnershipTransferred) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnershipTransferred.DiscardUnknown(m)
}

var xxx_messageInfo_OwnershipTransferred proto.InternalMessageInfo

func (m *OwnershipTransferred) GetPreviousOwner() string {
	if m != nil {
		return m.PreviousOwner
	}
	return ""
}

func (m *OwnershipTransferred) GetNewOwner() string {
	if m != nil {
		return m.NewOwner
	}
	return ""
}

// BlockedAddressesAdded is emitted whenever addresses are added to the blocklist.
type BlockedAddressesAdded struct {
	// accounts is the list of addresses that were added to the blocklist.
	Accounts []string `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (m *BlockedAddressesAdded) Reset()         { *m = BlockedAddressesAdded{} }
func (m *BlockedAddressesAdded) String() string { return proto.CompactTextString(m) }
func (*BlockedAddressesAdded) ProtoMessage()    {}
func (*BlockedAddressesAdded) Descriptor() ([]byte, []int) {
	return fileDescriptor_d01e1c8e4c279093, []int{2}
}
func (m *BlockedAddressesAdded) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockedAddressesAdded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockedAddressesAdded.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockedAddressesAdded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockedAddressesAdded.Merge(m, src)
}
func (m *BlockedAddressesAdded) XXX_Size() int {
	return m.Size()
}
func (m *BlockedAddressesAdded) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockedAddressesAdded.DiscardUnknown(m)
}

var xxx_messageInfo_BlockedAddressesAdded proto.InternalMessageInfo

func (m *BlockedAddressesAdded) GetAccounts() []string {
	if m != nil {
		return m.Accounts
	}
	return nil
}

// BlockedAddressesRemoved is emitted whenever addresses are removed from the blocklist.
type BlockedAddressesRemoved struct {
	// accounts is the list of addresses that were removed from the blocklist.
	Accounts []string `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (m *BlockedAddressesRemoved) Reset()         { *m = BlockedAddressesRemoved{} }
func (m *BlockedAddressesRemoved) String() string { return proto.CompactTextString(m) }
func (*BlockedAddressesRemoved) ProtoMessage()    {}
func (*BlockedAddressesRemoved) Descriptor() ([]byte, []int) {
	return fileDescriptor_d01e1c8e4c279093, []int{3}
}
func (m *BlockedAddressesRemoved) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockedAddressesRemoved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockedAddressesRemoved.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockedAddressesRemoved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockedAddressesRemoved.Merge(m, src)
}
func (m *BlockedAddressesRemoved) XXX_Size() int {
	return m.Size()
}
func (m *BlockedAddressesRemoved) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockedAddressesRemoved.DiscardUnknown(m)
}

var xxx_messageInfo_BlockedAddressesRemoved proto.InternalMessageInfo

func (m *BlockedAddressesRemoved) GetAccounts() []string {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func init() {
	proto.RegisterType((*OwnershipTransferStarted)(nil), "aura.blocklist.v1.OwnershipTransferStarted")
	proto.RegisterType((*OwnershipTransferred)(nil), "aura.blocklist.v1.OwnershipTransferred")
	proto.RegisterType((*BlockedAddressesAdded)(nil), "aura.blocklist.v1.BlockedAddressesAdded")
	proto.RegisterType((*BlockedAddressesRemoved)(nil), "aura.blocklist.v1.BlockedAddressesRemoved")
}

func init() { proto.RegisterFile("aura/blocklist/v1/events.proto", fileDescriptor_d01e1c8e4c279093) }

var fileDescriptor_d01e1c8e4c279093 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x17, 0x05, 0xd9, 0x02, 0x0a, 0x16, 0xc5, 0xa2, 0x10, 0x46, 0x41, 0xf0, 0x62, 0xc3,
	0x18, 0x1e, 0x3c, 0x6e, 0x7f, 0x40, 0xd8, 0x3c, 0xed, 0xa0, 0xb4, 0xcd, 0xa7, 0x2b, 0x76, 0xf9,
	0x4a, 0xbe, 0x34, 0x75, 0xff, 0xc2, 0x9f, 0xe5, 0x71, 0x47, 0x8f, 0xd2, 0xfe, 0x11, 0x69, 0x70,
	0x0a, 0x0a, 0x9e, 0x3c, 0x26, 0xef, 0xf3, 0xbc, 0x90, 0xbc, 0x5c, 0x24, 0x95, 0x49, 0x64, 0x5a,
	0x60, 0xf6, 0x54, 0xe4, 0x64, 0xa5, 0x1b, 0x49, 0x70, 0xa0, 0x2d, 0xc5, 0xa5, 0x41, 0x8b, 0xc1,
	0x61, 0x97, 0xc7, 0x5f, 0x79, 0xec, 0x46, 0xd1, 0x1d, 0x0f, 0x6f, 0x6a, 0x0d, 0x86, 0x96, 0x79,
	0x79, 0x6b, 0x12, 0x4d, 0x0f, 0x60, 0xe6, 0x36, 0x31, 0x16, 0x54, 0x70, 0xce, 0x0f, 0x4a, 0x03,
	0x2e, 0xc7, 0x8a, 0xee, 0xb1, 0x83, 0x42, 0x36, 0x64, 0x17, 0x83, 0xd9, 0xfe, 0xf6, 0xd6, 0x9b,
	0xc1, 0x19, 0x1f, 0x68, 0xa8, 0x3f, 0x89, 0x1d, 0x4f, 0xf4, 0x35, 0xd4, 0x3e, 0x8c, 0x16, 0xfc,
	0xe8, 0x57, 0xbf, 0xf9, 0xa7, 0xee, 0x31, 0x3f, 0x9e, 0x76, 0x6f, 0x01, 0x35, 0x51, 0xca, 0x00,
	0x11, 0xd0, 0x44, 0x29, 0x50, 0xc1, 0x29, 0xef, 0x27, 0x59, 0x86, 0x95, 0xb6, 0x14, 0xb2, 0xe1,
	0x6e, 0x27, 0x6d, 0xcf, 0xd1, 0x15, 0x3f, 0xf9, 0x29, 0xcd, 0x60, 0x85, 0xee, 0x6f, 0x6d, 0x3a,
	0x7f, 0x6d, 0x04, 0xdb, 0x34, 0x82, 0xbd, 0x37, 0x82, 0xbd, 0xb4, 0xa2, 0xb7, 0x69, 0x45, 0xef,
	0xad, 0x15, 0xbd, 0xc5, 0xf5, 0x63, 0x6e, 0x97, 0x55, 0x1a, 0x67, 0xb8, 0x92, 0xa8, 0x15, 0xfa,
	0xaf, 0xce, 0xb0, 0x90, 0x15, 0xa9, 0xf5, 0xa5, 0xc6, 0xb4, 0x00, 0xf9, 0x2c, 0xfd, 0x32, 0x76,
	0x5d, 0x02, 0x7d, 0xef, 0x93, 0xee, 0x79, 0x76, 0xfc, 0x11, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x3d,
	0x8f, 0x35, 0xb8, 0x01, 0x00, 0x00,
}

func (m *OwnershipTransferStarted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OwnershipTransferStarted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OwnershipTransferStarted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.PreviousOwner) > 0 {
		i -= len(m.PreviousOwner)
		copy(dAtA[i:], m.PreviousOwner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.PreviousOwner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OwnershipTransferred) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OwnershipTransferred) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OwnershipTransferred) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.PreviousOwner) > 0 {
		i -= len(m.PreviousOwner)
		copy(dAtA[i:], m.PreviousOwner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.PreviousOwner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockedAddressesAdded) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockedAddressesAdded) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockedAddressesAdded) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Accounts) > 0 {
		for iNdEx := len(m.Accounts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Accounts[iNdEx])
			copy(dAtA[i:], m.Accounts[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.Accounts[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BlockedAddressesRemoved) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockedAddressesRemoved) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockedAddressesRemoved) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Accounts) > 0 {
		for iNdEx := len(m.Accounts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Accounts[iNdEx])
			copy(dAtA[i:], m.Accounts[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.Accounts[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
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
func (m *OwnershipTransferStarted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PreviousOwner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NewOwner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *OwnershipTransferred) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PreviousOwner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NewOwner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *BlockedAddressesAdded) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Accounts) > 0 {
		for _, s := range m.Accounts {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *BlockedAddressesRemoved) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Accounts) > 0 {
		for _, s := range m.Accounts {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OwnershipTransferStarted) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: OwnershipTransferStarted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OwnershipTransferStarted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousOwner", wireType)
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
			m.PreviousOwner = string(dAtA[iNdEx:postIndex])
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
func (m *OwnershipTransferred) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: OwnershipTransferred: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OwnershipTransferred: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousOwner", wireType)
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
			m.PreviousOwner = string(dAtA[iNdEx:postIndex])
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
func (m *BlockedAddressesAdded) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: BlockedAddressesAdded: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockedAddressesAdded: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accounts", wireType)
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
			m.Accounts = append(m.Accounts, string(dAtA[iNdEx:postIndex]))
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
func (m *BlockedAddressesRemoved) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: BlockedAddressesRemoved: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockedAddressesRemoved: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accounts", wireType)
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
			m.Accounts = append(m.Accounts, string(dAtA[iNdEx:postIndex]))
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
