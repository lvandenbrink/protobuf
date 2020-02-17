package graphql_test

import (
	bytes "bytes"
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

// A delicious dessert dish on the menu
type TestDessert struct {
	// The name of the dish
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// How sweet is the dish, an integer between 0 and 10
	Sweetness            int32    `protobuf:"varint,2,opt,name=sweetness,proto3" json:"sweetness,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestDessert) Reset()         { *m = TestDessert{} }
func (m *TestDessert) String() string { return proto.CompactTextString(m) }
func (*TestDessert) ProtoMessage()    {}
func (*TestDessert) Descriptor() ([]byte, []int) {
	return fileDescriptor_592790b6763b2c67, []int{0}
}
func (m *TestDessert) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestDessert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestDessert.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestDessert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestDessert.Merge(m, src)
}
func (m *TestDessert) XXX_Size() int {
	return m.Size()
}
func (m *TestDessert) XXX_DiscardUnknown() {
	xxx_messageInfo_TestDessert.DiscardUnknown(m)
}

var xxx_messageInfo_TestDessert proto.InternalMessageInfo

func (m *TestDessert) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestDessert) GetSweetness() int32 {
	if m != nil {
		return m.Sweetness
	}
	return 0
}

func init() {
	proto.RegisterType((*TestDessert)(nil), "flavortown.dessert.TestDessert")
}

func init() {
	proto.RegisterFile("dessert.proto", fileDescriptor_592790b6763b2c67)
}

var fileDescriptor_592790b6763b2c67 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x49, 0x2d, 0x2e,
	0x4e, 0x2d, 0x2a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a, 0xcb, 0x49, 0x2c, 0xcb,
	0x2f, 0x2a, 0xc9, 0x2f, 0xcf, 0xd3, 0x83, 0xca, 0x48, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0x95, 0x26, 0x95, 0xa6, 0x81,
	0x79, 0x60, 0x0e, 0x98, 0x05, 0x31, 0x42, 0xca, 0x00, 0x49, 0x79, 0x7e, 0x41, 0x71, 0x6a, 0x2a,
	0x42, 0x3d, 0x98, 0x0b, 0xd1, 0x00, 0x66, 0x42, 0x74, 0x28, 0x59, 0x73, 0xb1, 0xbb, 0x40, 0xec,
	0x12, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02,
	0xb3, 0x85, 0x64, 0xb8, 0x38, 0x8b, 0xcb, 0x53, 0x53, 0x4b, 0xf2, 0x52, 0x8b, 0x8b, 0x25, 0x98,
	0x14, 0x18, 0x35, 0x58, 0x83, 0x10, 0x02, 0x4e, 0x32, 0x3f, 0x1e, 0xca, 0x31, 0xae, 0x78, 0x24,
	0xc7, 0xb8, 0xe3, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78,
	0x24, 0xc7, 0x78, 0x60, 0x91, 0x3c, 0x63, 0x12, 0x1b, 0xd8, 0x06, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xdd, 0x92, 0xb5, 0xa7, 0xe7, 0x00, 0x00, 0x00,
}

func (this *TestDessert) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TestDessert)
	if !ok {
		that2, ok := that.(TestDessert)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Sweetness != that1.Sweetness {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

type TestDessertGetter interface {
	GetTestDessert() *TestDessert
}

func (m *TestDessert) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestDessert) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTestDessert(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Sweetness != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTestDessert(dAtA, i, uint64(m.Sweetness))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTestDessert(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedTestDessert(r randyTestDessert, easy bool) *TestDessert {
	this := &TestDessert{}
	this.Name = string(randStringTestDessert(r))
	this.Sweetness = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Sweetness *= -1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTestDessert(r, 3)
	}
	return this
}

type randyTestDessert interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTestDessert(r randyTestDessert) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTestDessert(r randyTestDessert) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneTestDessert(r)
	}
	return string(tmps)
}
func randUnrecognizedTestDessert(r randyTestDessert, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTestDessert(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTestDessert(dAtA []byte, r randyTestDessert, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTestDessert(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateTestDessert(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateTestDessert(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTestDessert(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTestDessert(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTestDessert(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTestDessert(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *TestDessert) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTestDessert(uint64(l))
	}
	if m.Sweetness != 0 {
		n += 1 + sovTestDessert(uint64(m.Sweetness))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTestDessert(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTestDessert(x uint64) (n int) {
	return sovTestDessert(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TestDessert) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestDessert
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
			return fmt.Errorf("proto: TestDessert: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestDessert: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestDessert
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
				return ErrInvalidLengthTestDessert
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTestDessert
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sweetness", wireType)
			}
			m.Sweetness = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestDessert
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sweetness |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTestDessert(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTestDessert
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTestDessert
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
func skipTestDessert(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTestDessert
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
					return 0, ErrIntOverflowTestDessert
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
					return 0, ErrIntOverflowTestDessert
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
				return 0, ErrInvalidLengthTestDessert
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTestDessert
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTestDessert
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
				next, err := skipTestDessert(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTestDessert
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
	ErrInvalidLengthTestDessert = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTestDessert   = fmt.Errorf("proto: integer overflow")
)
