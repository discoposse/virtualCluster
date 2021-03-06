// Code generated by protoc-gen-go.
// source: PublicCloudDTO.proto
// DO NOT EDIT!

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// enumerate non market entities
type NonMarketEntityDTO_NonMarketEntityType int32

const (
	NonMarketEntityDTO_CLOUD_SERVICE NonMarketEntityDTO_NonMarketEntityType = 0
)

var NonMarketEntityDTO_NonMarketEntityType_name = map[int32]string{
	0: "CLOUD_SERVICE",
}
var NonMarketEntityDTO_NonMarketEntityType_value = map[string]int32{
	"CLOUD_SERVICE": 0,
}

func (x NonMarketEntityDTO_NonMarketEntityType) Enum() *NonMarketEntityDTO_NonMarketEntityType {
	p := new(NonMarketEntityDTO_NonMarketEntityType)
	*p = x
	return p
}
func (x NonMarketEntityDTO_NonMarketEntityType) String() string {
	return proto.EnumName(NonMarketEntityDTO_NonMarketEntityType_name, int32(x))
}
func (x *NonMarketEntityDTO_NonMarketEntityType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NonMarketEntityDTO_NonMarketEntityType_value, data, "NonMarketEntityDTO_NonMarketEntityType")
	if err != nil {
		return err
	}
	*x = NonMarketEntityDTO_NonMarketEntityType(value)
	return nil
}
func (NonMarketEntityDTO_NonMarketEntityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor6, []int{0, 0}
}

// The NonMarketEntityDTO message represents an Entity discovered in the target that your probe is
// monitoring and this entity does not participate in the market.
//
// Each entity must have a unique ID to identify it in the Operations Manager market.
// Many targets provide unique IDs for their entities, or you can generate your own.
// To guarantee that it's unique, you can give the ID a prefix that identifies your
// probe and the given target.
//
// Specify entity type by setting an 'EntityType' value to the 'entity' field.
//
// The 'displayName' value appears in the product GUI and in reports to identify the entity.
//
type NonMarketEntityDTO struct {
	EntityType  *NonMarketEntityDTO_NonMarketEntityType `protobuf:"varint,1,req,name=entityType,enum=common_dto.NonMarketEntityDTO_NonMarketEntityType" json:"entityType,omitempty"`
	Id          *string                                 `protobuf:"bytes,2,req,name=id" json:"id,omitempty"`
	DisplayName *string                                 `protobuf:"bytes,3,opt,name=displayName" json:"displayName,omitempty"`
	// Collection of entity type's specific data.  Currently, we only have CloudService.
	//
	// Types that are valid to be assigned to EntityData:
	//	*NonMarketEntityDTO_CloudServiceData_
	EntityData       isNonMarketEntityDTO_EntityData `protobuf_oneof:"entity_data"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *NonMarketEntityDTO) Reset()                    { *m = NonMarketEntityDTO{} }
func (m *NonMarketEntityDTO) String() string            { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO) ProtoMessage()               {}
func (*NonMarketEntityDTO) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

type isNonMarketEntityDTO_EntityData interface {
	isNonMarketEntityDTO_EntityData()
}

type NonMarketEntityDTO_CloudServiceData_ struct {
	CloudServiceData *NonMarketEntityDTO_CloudServiceData `protobuf:"bytes,500,opt,name=cloud_service_data,json=cloudServiceData,oneof"`
}

func (*NonMarketEntityDTO_CloudServiceData_) isNonMarketEntityDTO_EntityData() {}

func (m *NonMarketEntityDTO) GetEntityData() isNonMarketEntityDTO_EntityData {
	if m != nil {
		return m.EntityData
	}
	return nil
}

func (m *NonMarketEntityDTO) GetEntityType() NonMarketEntityDTO_NonMarketEntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return NonMarketEntityDTO_CLOUD_SERVICE
}

func (m *NonMarketEntityDTO) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *NonMarketEntityDTO) GetDisplayName() string {
	if m != nil && m.DisplayName != nil {
		return *m.DisplayName
	}
	return ""
}

func (m *NonMarketEntityDTO) GetCloudServiceData() *NonMarketEntityDTO_CloudServiceData {
	if x, ok := m.GetEntityData().(*NonMarketEntityDTO_CloudServiceData_); ok {
		return x.CloudServiceData
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*NonMarketEntityDTO) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _NonMarketEntityDTO_OneofMarshaler, _NonMarketEntityDTO_OneofUnmarshaler, _NonMarketEntityDTO_OneofSizer, []interface{}{
		(*NonMarketEntityDTO_CloudServiceData_)(nil),
	}
}

func _NonMarketEntityDTO_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*NonMarketEntityDTO)
	// entity_data
	switch x := m.EntityData.(type) {
	case *NonMarketEntityDTO_CloudServiceData_:
		b.EncodeVarint(500<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CloudServiceData); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("NonMarketEntityDTO.EntityData has unexpected type %T", x)
	}
	return nil
}

func _NonMarketEntityDTO_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*NonMarketEntityDTO)
	switch tag {
	case 500: // entity_data.cloud_service_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NonMarketEntityDTO_CloudServiceData)
		err := b.DecodeMessage(msg)
		m.EntityData = &NonMarketEntityDTO_CloudServiceData_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _NonMarketEntityDTO_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*NonMarketEntityDTO)
	// entity_data
	switch x := m.EntityData.(type) {
	case *NonMarketEntityDTO_CloudServiceData_:
		s := proto.Size(x.CloudServiceData)
		n += proto.SizeVarint(500<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NonMarketEntityDTO_CloudServiceData struct {
	CloudProvider    *string  `protobuf:"bytes,1,req,name=cloud_provider,json=cloudProvider" json:"cloud_provider,omitempty"`
	Spend            *float32 `protobuf:"fixed32,2,req,name=spend" json:"spend,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *NonMarketEntityDTO_CloudServiceData) Reset()         { *m = NonMarketEntityDTO_CloudServiceData{} }
func (m *NonMarketEntityDTO_CloudServiceData) String() string { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO_CloudServiceData) ProtoMessage()    {}
func (*NonMarketEntityDTO_CloudServiceData) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{0, 0}
}

func (m *NonMarketEntityDTO_CloudServiceData) GetCloudProvider() string {
	if m != nil && m.CloudProvider != nil {
		return *m.CloudProvider
	}
	return ""
}

func (m *NonMarketEntityDTO_CloudServiceData) GetSpend() float32 {
	if m != nil && m.Spend != nil {
		return *m.Spend
	}
	return 0
}

func init() {
	proto.RegisterType((*NonMarketEntityDTO)(nil), "common_dto.NonMarketEntityDTO")
	proto.RegisterType((*NonMarketEntityDTO_CloudServiceData)(nil), "common_dto.NonMarketEntityDTO.CloudServiceData")
	proto.RegisterEnum("common_dto.NonMarketEntityDTO_NonMarketEntityType", NonMarketEntityDTO_NonMarketEntityType_name, NonMarketEntityDTO_NonMarketEntityType_value)
}

func init() { proto.RegisterFile("PublicCloudDTO.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x6a, 0x32, 0x31,
	0x10, 0xc7, 0x35, 0xf2, 0x1d, 0x1c, 0x71, 0xf1, 0x4b, 0x3d, 0x2c, 0xbd, 0x74, 0x11, 0x0a, 0x7b,
	0x4a, 0xa9, 0x8f, 0xa0, 0x2b, 0xb4, 0xd0, 0xba, 0x12, 0x6d, 0xaf, 0xdb, 0x98, 0xa4, 0x10, 0xba,
	0xd9, 0x09, 0x31, 0x0a, 0xbe, 0x40, 0x9f, 0xb8, 0x0f, 0x50, 0xdc, 0x15, 0x6a, 0x6d, 0xa1, 0xc7,
	0xf9, 0x31, 0xff, 0x64, 0x7e, 0x33, 0x30, 0x5c, 0x6c, 0xd7, 0xa5, 0x91, 0xd3, 0x12, 0xb7, 0x2a,
	0x5b, 0xe5, 0xcc, 0x79, 0x0c, 0x48, 0x41, 0xa2, 0xb5, 0x58, 0x15, 0x2a, 0xe0, 0xe8, 0xbd, 0x03,
	0x74, 0x8e, 0xd5, 0xa3, 0xf0, 0x6f, 0x3a, 0xcc, 0xaa, 0x60, 0xc2, 0x3e, 0x5b, 0xe5, 0x94, 0x03,
	0xe8, 0xba, 0x58, 0xed, 0x9d, 0x8e, 0xdb, 0x09, 0x49, 0xa3, 0xf1, 0x98, 0x7d, 0xe5, 0xd8, 0xcf,
	0xcc, 0x39, 0x3a, 0x24, 0xf9, 0xc9, 0x2b, 0x34, 0x02, 0x62, 0x54, 0x4c, 0x12, 0x92, 0x76, 0x39,
	0x31, 0x8a, 0x26, 0xd0, 0x53, 0x66, 0xe3, 0x4a, 0xb1, 0x9f, 0x0b, 0xab, 0xe3, 0x4e, 0xd2, 0x4e,
	0xbb, 0xfc, 0x14, 0xd1, 0x17, 0xa0, 0xf2, 0x30, 0x7a, 0xb1, 0xd1, 0x7e, 0x67, 0xa4, 0x2e, 0x94,
	0x08, 0x22, 0xfe, 0x38, 0x74, 0xf6, 0xc6, 0x37, 0x7f, 0x8c, 0x53, 0x4b, 0x2f, 0x9b, 0x60, 0x26,
	0x82, 0xb8, 0x6b, 0xf1, 0x81, 0x3c, 0x63, 0x97, 0x39, 0x0c, 0xce, 0xfb, 0xe8, 0x35, 0x44, 0xcd,
	0xaf, 0xce, 0xe3, 0xce, 0x28, 0xed, 0x6b, 0xff, 0x2e, 0xef, 0xd7, 0x74, 0x71, 0x84, 0x74, 0x08,
	0xff, 0x36, 0x4e, 0x57, 0x8d, 0x11, 0xe1, 0x4d, 0x31, 0x4a, 0xe1, 0xe2, 0x97, 0x3d, 0xd0, 0xff,
	0xd0, 0x9f, 0x3e, 0xe4, 0x4f, 0x59, 0xb1, 0x9c, 0xf1, 0xe7, 0xfb, 0xe9, 0x6c, 0xd0, 0x9a, 0xf4,
	0xa1, 0xd7, 0x2c, 0xa7, 0xb6, 0x9a, 0xdc, 0xc2, 0x95, 0x44, 0xcb, 0x76, 0x36, 0x6c, 0xfd, 0x1a,
	0x99, 0x2b, 0x45, 0x78, 0x45, 0x6f, 0x8f, 0x92, 0x4c, 0x05, 0x9c, 0x44, 0xdf, 0xaf, 0xf9, 0x19,
	0x00, 0x00, 0xff, 0xff, 0x34, 0x40, 0x62, 0x43, 0xde, 0x01, 0x00, 0x00,
}
