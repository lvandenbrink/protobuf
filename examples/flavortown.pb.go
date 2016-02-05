// Code generated by protoc-gen-gogo.
// source: flavortown.proto
// DO NOT EDIT!

/*
Package flavortown is a generated protocol buffer package.

It is generated from these files:
	flavortown.proto

It has these top-level messages:
	Menu
	LineItem
	Lunch
*/
package flavortown

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/opsee/protobuf/proto/google/protobuf"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/opsee/protobuf/gogogqlproto"
import flavortown_dessert "github.com/opsee/protobuf/examples/dessert"

import bytes "bytes"

import github_com_graphql_go_graphql "github.com/graphql-go/graphql"
import github_com_opsee_protobuf_gogogqlproto "github.com/opsee/protobuf/gogogqlproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The menu at Guy’s American Kitchen & Bar reflects his signature style of authentic and surprising flavors
type Menu struct {
	// These dishes are crafted with the heart and soul of hometown favorites and infused with Guy’s big, daring flavors
	Items []*LineItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *Menu) Reset()         { *m = Menu{} }
func (m *Menu) String() string { return proto.CompactTextString(m) }
func (*Menu) ProtoMessage()    {}

func (m *Menu) GetItems() []*LineItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// A line item representing a dish and price
type LineItem struct {
	// The menu dish, can either be lunch or dessert
	//
	// Types that are valid to be assigned to Dish:
	//	*LineItem_Lunch
	//	*LineItem_Dessert
	Dish isLineItem_Dish `protobuf_oneof:"dish"`
	// The price of the dish in cents
	PriceCents int32 `protobuf:"varint,2,opt,name=price_cents,proto3" json:"price_cents,omitempty"`
	// A timestamp representing when the dish was added to the menu
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=created_at" json:"created_at,omitempty"`
	// A timestamp representing when the dish was updated
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=updated_at" json:"updated_at,omitempty"`
}

func (m *LineItem) Reset()         { *m = LineItem{} }
func (m *LineItem) String() string { return proto.CompactTextString(m) }
func (*LineItem) ProtoMessage()    {}

type isLineItem_Dish interface {
	isLineItem_Dish()
	Equal(interface{}) bool
}

type LineItem_Lunch struct {
	Lunch *Lunch `protobuf:"bytes,100,opt,name=lunch,oneof"`
}
type LineItem_Dessert struct {
	Dessert *flavortown_dessert.Dessert `protobuf:"bytes,101,opt,name=dessert,oneof"`
}

func (*LineItem_Lunch) isLineItem_Dish()   {}
func (*LineItem_Dessert) isLineItem_Dish() {}

func (m *LineItem) GetDish() isLineItem_Dish {
	if m != nil {
		return m.Dish
	}
	return nil
}

func (m *LineItem) GetLunch() *Lunch {
	if x, ok := m.GetDish().(*LineItem_Lunch); ok {
		return x.Lunch
	}
	return nil
}

func (m *LineItem) GetDessert() *flavortown_dessert.Dessert {
	if x, ok := m.GetDish().(*LineItem_Dessert); ok {
		return x.Dessert
	}
	return nil
}

func (m *LineItem) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *LineItem) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LineItem) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _LineItem_OneofMarshaler, _LineItem_OneofUnmarshaler, []interface{}{
		(*LineItem_Lunch)(nil),
		(*LineItem_Dessert)(nil),
	}
}

func _LineItem_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LineItem)
	// dish
	switch x := m.Dish.(type) {
	case *LineItem_Lunch:
		_ = b.EncodeVarint(100<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Lunch); err != nil {
			return err
		}
	case *LineItem_Dessert:
		_ = b.EncodeVarint(101<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Dessert); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LineItem.Dish has unexpected type %T", x)
	}
	return nil
}

func _LineItem_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LineItem)
	switch tag {
	case 100: // dish.lunch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Lunch)
		err := b.DecodeMessage(msg)
		m.Dish = &LineItem_Lunch{msg}
		return true, err
	case 101: // dish.dessert
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(flavortown_dessert.Dessert)
		err := b.DecodeMessage(msg)
		m.Dish = &LineItem_Dessert{msg}
		return true, err
	default:
		return false, nil
	}
}

// A delicious lunch dish on the menu
type Lunch struct {
	// The name of the dish
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The description of the dish
	Description []byte `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *Lunch) Reset()         { *m = Lunch{} }
func (m *Lunch) String() string { return proto.CompactTextString(m) }
func (*Lunch) ProtoMessage()    {}

func init() {
	proto.RegisterType((*Menu)(nil), "flavortown.Menu")
	proto.RegisterType((*LineItem)(nil), "flavortown.LineItem")
	proto.RegisterType((*Lunch)(nil), "flavortown.Lunch")
}
func (this *Menu) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Menu)
	if !ok {
		that2, ok := that.(Menu)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Items) != len(that1.Items) {
		return false
	}
	for i := range this.Items {
		if !this.Items[i].Equal(that1.Items[i]) {
			return false
		}
	}
	return true
}
func (this *LineItem) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*LineItem)
	if !ok {
		that2, ok := that.(LineItem)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if that1.Dish == nil {
		if this.Dish != nil {
			return false
		}
	} else if this.Dish == nil {
		return false
	} else if !this.Dish.Equal(that1.Dish) {
		return false
	}
	if this.PriceCents != that1.PriceCents {
		return false
	}
	if !this.CreatedAt.Equal(that1.CreatedAt) {
		return false
	}
	if !this.UpdatedAt.Equal(that1.UpdatedAt) {
		return false
	}
	return true
}
func (this *LineItem_Lunch) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*LineItem_Lunch)
	if !ok {
		that2, ok := that.(LineItem_Lunch)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Lunch.Equal(that1.Lunch) {
		return false
	}
	return true
}
func (this *LineItem_Dessert) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*LineItem_Dessert)
	if !ok {
		that2, ok := that.(LineItem_Dessert)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Dessert.Equal(that1.Dessert) {
		return false
	}
	return true
}
func (this *Lunch) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Lunch)
	if !ok {
		that2, ok := that.(Lunch)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if !bytes.Equal(this.Description, that1.Description) {
		return false
	}
	return true
}

var GraphQLMenuType *github_com_graphql_go_graphql.Object
var GraphQLLineItemType *github_com_graphql_go_graphql.Object
var GraphQLLineItemDishUnion *github_com_graphql_go_graphql.Union
var GraphQLLunchType *github_com_graphql_go_graphql.Object

func init() {
	GraphQLMenuType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "menu",
		Description: "The menu at Guy’s American Kitchen & Bar reflects his signature style of authentic and surprising flavors",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"items": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.NewList(GraphQLLineItemType),
					Description: "These dishes are crafted with the heart and soul of hometown favorites and infused with Guy’s big, daring flavors",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						switch obj := p.Source.(type) {
						case *Menu:
							return obj.Items, nil
						}
						return nil, fmt.Errorf("field items not resolved")
					},
				},
			}
		}),
	})
	GraphQLLineItemType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "line_item",
		Description: "A line item representing a dish and price",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"price_cents": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Int,
					Description: "The price of the dish in cents",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						switch obj := p.Source.(type) {
						case *LineItem:
							return obj.PriceCents, nil
						}
						return nil, fmt.Errorf("field price_cents not resolved")
					},
				},
				"created_at": &github_com_graphql_go_graphql.Field{
					Type:        github_com_opsee_protobuf_gogogqlproto.ByteString,
					Description: "A timestamp representing when the dish was added to the menu",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						switch obj := p.Source.(type) {
						case *LineItem:
							return obj.CreatedAt, nil
						}
						return nil, fmt.Errorf("field created_at not resolved")
					},
				},
				"updated_at": &github_com_graphql_go_graphql.Field{
					Type:        github_com_opsee_protobuf_gogogqlproto.ByteString,
					Description: "A timestamp representing when the dish was updated",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						switch obj := p.Source.(type) {
						case *LineItem:
							return obj.UpdatedAt, nil
						}
						return nil, fmt.Errorf("field updated_at not resolved")
					},
				},
				"dish": &github_com_graphql_go_graphql.Field{
					Type:        GraphQLLineItemDishUnion,
					Description: "The menu dish, can either be lunch or dessert",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*LineItem)
						if !ok {
							return nil, fmt.Errorf("field dish not resolved")
						}
						return obj.GetDish(), nil
					},
				},
			}
		}),
	})
	GraphQLLunchType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "lunch",
		Description: "A delicious lunch dish on the menu",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"name": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "The name of the dish",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						switch obj := p.Source.(type) {
						case *Lunch:
							return obj.Name, nil
						case *LineItem_Lunch:
							return obj.Lunch.Name, nil
						}
						return nil, fmt.Errorf("field name not resolved")
					},
				},
				"description": &github_com_graphql_go_graphql.Field{
					Type:        github_com_opsee_protobuf_gogogqlproto.ByteString,
					Description: "The description of the dish",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						switch obj := p.Source.(type) {
						case *Lunch:
							return obj.Description, nil
						case *LineItem_Lunch:
							return obj.Lunch.Description, nil
						}
						return nil, fmt.Errorf("field description not resolved")
					},
				},
			}
		}),
	})
	GraphQLLineItemDishUnion = github_com_graphql_go_graphql.NewUnion(github_com_graphql_go_graphql.UnionConfig{
		Name:        "LineItemDish",
		Description: "The menu dish, can either be lunch or dessert",
		Types: []*github_com_graphql_go_graphql.Object{
			GraphQLLunchType,
			flavortown_dessert.GraphQLDessertType,
		},
		ResolveType: func(value interface{}, info github_com_graphql_go_graphql.ResolveInfo) *github_com_graphql_go_graphql.Object {
			switch value.(type) {
			case *LineItem_Lunch:
				return GraphQLLunchType
			case *LineItem_Dessert:
				return flavortown_dessert.GraphQLDessertType
			}
			return nil
		},
	})
}
func NewPopulatedMenu(r randyFlavortown, easy bool) *Menu {
	this := &Menu{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.Items = make([]*LineItem, v1)
		for i := 0; i < v1; i++ {
			this.Items[i] = NewPopulatedLineItem(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedLineItem(r randyFlavortown, easy bool) *LineItem {
	this := &LineItem{}
	oneofNumber_Dish := []int32{100, 101}[r.Intn(2)]
	switch oneofNumber_Dish {
	case 100:
		this.Dish = NewPopulatedLineItem_Lunch(r, easy)
	case 101:
		this.Dish = NewPopulatedLineItem_Dessert(r, easy)
	}
	this.PriceCents = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.PriceCents *= -1
	}
	if r.Intn(10) != 0 {
		this.CreatedAt = google_protobuf.NewPopulatedTimestamp(r, easy)
	}
	if r.Intn(10) != 0 {
		this.UpdatedAt = google_protobuf.NewPopulatedTimestamp(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedLineItem_Lunch(r randyFlavortown, easy bool) *LineItem_Lunch {
	this := &LineItem_Lunch{}
	this.Lunch = NewPopulatedLunch(r, easy)
	return this
}
func NewPopulatedLineItem_Dessert(r randyFlavortown, easy bool) *LineItem_Dessert {
	this := &LineItem_Dessert{}
	this.Dessert = flavortown_dessert.NewPopulatedDessert(r, easy)
	return this
}
func NewPopulatedLunch(r randyFlavortown, easy bool) *Lunch {
	this := &Lunch{}
	this.Name = randStringFlavortown(r)
	v2 := r.Intn(100)
	this.Description = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.Description[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyFlavortown interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneFlavortown(r randyFlavortown) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringFlavortown(r randyFlavortown) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneFlavortown(r)
	}
	return string(tmps)
}
func randUnrecognizedFlavortown(r randyFlavortown, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldFlavortown(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldFlavortown(data []byte, r randyFlavortown, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateFlavortown(data, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		data = encodeVarintPopulateFlavortown(data, uint64(v4))
	case 1:
		data = encodeVarintPopulateFlavortown(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateFlavortown(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateFlavortown(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateFlavortown(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateFlavortown(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
