// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: api/listing/taxonomy.proto

package listing

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Type specifies what form the related attribute should be rendered into
type AttributeSpec_Type int32

const (
	AttributeSpec_TYPE_UNSPECIFIED        AttributeSpec_Type = 0
	AttributeSpec_TYPE_TEXT               AttributeSpec_Type = 1
	AttributeSpec_TYPE_SELECT             AttributeSpec_Type = 2
	AttributeSpec_TYPE_NUMERIC            AttributeSpec_Type = 3
	AttributeSpec_TYPE_NUMERIC_WITH_UNITS AttributeSpec_Type = 4
	AttributeSpec_TYPE_MULTI_TEXT         AttributeSpec_Type = 5
	AttributeSpec_TYPE_MULTI_SELECT       AttributeSpec_Type = 6
	AttributeSpec_TYPE_MONEY              AttributeSpec_Type = 7
)

// Enum value maps for AttributeSpec_Type.
var (
	AttributeSpec_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_TEXT",
		2: "TYPE_SELECT",
		3: "TYPE_NUMERIC",
		4: "TYPE_NUMERIC_WITH_UNITS",
		5: "TYPE_MULTI_TEXT",
		6: "TYPE_MULTI_SELECT",
		7: "TYPE_MONEY",
	}
	AttributeSpec_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":        0,
		"TYPE_TEXT":               1,
		"TYPE_SELECT":             2,
		"TYPE_NUMERIC":            3,
		"TYPE_NUMERIC_WITH_UNITS": 4,
		"TYPE_MULTI_TEXT":         5,
		"TYPE_MULTI_SELECT":       6,
		"TYPE_MONEY":              7,
	}
)

func (x AttributeSpec_Type) Enum() *AttributeSpec_Type {
	p := new(AttributeSpec_Type)
	*p = x
	return p
}

func (x AttributeSpec_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttributeSpec_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_api_listing_taxonomy_proto_enumTypes[0].Descriptor()
}

func (AttributeSpec_Type) Type() protoreflect.EnumType {
	return &file_api_listing_taxonomy_proto_enumTypes[0]
}

func (x AttributeSpec_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttributeSpec_Type.Descriptor instead.
func (AttributeSpec_Type) EnumDescriptor() ([]byte, []int) {
	return file_api_listing_taxonomy_proto_rawDescGZIP(), []int{0, 0}
}

// Level specifies whether this AttributeSpec should be rendered
// on the Listing-level or the Variant-level
type AttributeSpec_Level int32

const (
	AttributeSpec_LEVEL_UNSPECIFIED AttributeSpec_Level = 0
	AttributeSpec_LEVEL_LISTING     AttributeSpec_Level = 1
	AttributeSpec_LEVEL_VARIANT     AttributeSpec_Level = 2
)

// Enum value maps for AttributeSpec_Level.
var (
	AttributeSpec_Level_name = map[int32]string{
		0: "LEVEL_UNSPECIFIED",
		1: "LEVEL_LISTING",
		2: "LEVEL_VARIANT",
	}
	AttributeSpec_Level_value = map[string]int32{
		"LEVEL_UNSPECIFIED": 0,
		"LEVEL_LISTING":     1,
		"LEVEL_VARIANT":     2,
	}
)

func (x AttributeSpec_Level) Enum() *AttributeSpec_Level {
	p := new(AttributeSpec_Level)
	*p = x
	return p
}

func (x AttributeSpec_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttributeSpec_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_api_listing_taxonomy_proto_enumTypes[1].Descriptor()
}

func (AttributeSpec_Level) Type() protoreflect.EnumType {
	return &file_api_listing_taxonomy_proto_enumTypes[1]
}

func (x AttributeSpec_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttributeSpec_Level.Descriptor instead.
func (AttributeSpec_Level) EnumDescriptor() ([]byte, []int) {
	return file_api_listing_taxonomy_proto_rawDescGZIP(), []int{0, 1}
}

// Classification specifies whether this AttributeSpec relates to
// Product, Inventory, or Pricing data
type AttributeSpec_Classification int32

const (
	AttributeSpec_CLASSIFICATION_UNSPECIFIED AttributeSpec_Classification = 0
	AttributeSpec_CLASSIFICATION_PRODUCT     AttributeSpec_Classification = 1
	AttributeSpec_CLASSIFICATION_INVENTORY   AttributeSpec_Classification = 2
	AttributeSpec_CLASSIFICATION_PRICING     AttributeSpec_Classification = 3
	AttributeSpec_CLASSIFICATION_LOGISTICS   AttributeSpec_Classification = 4
	AttributeSpec_CLASSIFICATION_IDENTIFIER  AttributeSpec_Classification = 5
)

// Enum value maps for AttributeSpec_Classification.
var (
	AttributeSpec_Classification_name = map[int32]string{
		0: "CLASSIFICATION_UNSPECIFIED",
		1: "CLASSIFICATION_PRODUCT",
		2: "CLASSIFICATION_INVENTORY",
		3: "CLASSIFICATION_PRICING",
		4: "CLASSIFICATION_LOGISTICS",
		5: "CLASSIFICATION_IDENTIFIER",
	}
	AttributeSpec_Classification_value = map[string]int32{
		"CLASSIFICATION_UNSPECIFIED": 0,
		"CLASSIFICATION_PRODUCT":     1,
		"CLASSIFICATION_INVENTORY":   2,
		"CLASSIFICATION_PRICING":     3,
		"CLASSIFICATION_LOGISTICS":   4,
		"CLASSIFICATION_IDENTIFIER":  5,
	}
)

func (x AttributeSpec_Classification) Enum() *AttributeSpec_Classification {
	p := new(AttributeSpec_Classification)
	*p = x
	return p
}

func (x AttributeSpec_Classification) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttributeSpec_Classification) Descriptor() protoreflect.EnumDescriptor {
	return file_api_listing_taxonomy_proto_enumTypes[2].Descriptor()
}

func (AttributeSpec_Classification) Type() protoreflect.EnumType {
	return &file_api_listing_taxonomy_proto_enumTypes[2]
}

func (x AttributeSpec_Classification) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttributeSpec_Classification.Descriptor instead.
func (AttributeSpec_Classification) EnumDescriptor() ([]byte, []int) {
	return file_api_listing_taxonomy_proto_rawDescGZIP(), []int{0, 2}
}

// Usage indicates if the spec is required or just recommended
type AttributeSpec_Usage int32

const (
	AttributeSpec_USAGE_UNSPECIFIED AttributeSpec_Usage = 0
	AttributeSpec_USAGE_AVAILABLE   AttributeSpec_Usage = 1
	AttributeSpec_USAGE_REQUIRED    AttributeSpec_Usage = 2
	AttributeSpec_USAGE_RECOMMENDED AttributeSpec_Usage = 3
)

// Enum value maps for AttributeSpec_Usage.
var (
	AttributeSpec_Usage_name = map[int32]string{
		0: "USAGE_UNSPECIFIED",
		1: "USAGE_AVAILABLE",
		2: "USAGE_REQUIRED",
		3: "USAGE_RECOMMENDED",
	}
	AttributeSpec_Usage_value = map[string]int32{
		"USAGE_UNSPECIFIED": 0,
		"USAGE_AVAILABLE":   1,
		"USAGE_REQUIRED":    2,
		"USAGE_RECOMMENDED": 3,
	}
)

func (x AttributeSpec_Usage) Enum() *AttributeSpec_Usage {
	p := new(AttributeSpec_Usage)
	*p = x
	return p
}

func (x AttributeSpec_Usage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttributeSpec_Usage) Descriptor() protoreflect.EnumDescriptor {
	return file_api_listing_taxonomy_proto_enumTypes[3].Descriptor()
}

func (AttributeSpec_Usage) Type() protoreflect.EnumType {
	return &file_api_listing_taxonomy_proto_enumTypes[3]
}

func (x AttributeSpec_Usage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttributeSpec_Usage.Descriptor instead.
func (AttributeSpec_Usage) EnumDescriptor() ([]byte, []int) {
	return file_api_listing_taxonomy_proto_rawDescGZIP(), []int{0, 3}
}

// AttributeSpec is the specification of an attribute within a Product Type on
// your channel.
type AttributeSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is the ID that will be used when retrieving listing data or submitting
	// error to Zentail. It should reflect whatever meaningful identifier can be
	// used to ingest this attribute into your system. In the past we have seen it
	// be a numeric identifier or an xpath/jsonpath string
	AttributeSpecId string `protobuf:"bytes,1,opt,name=attribute_spec_id,json=attributeSpecId,proto3" json:"attribute_spec_id,omitempty"`
	// Human-readable name for the attribute. Used for display purposes only.
	DisplayName string             `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Type        AttributeSpec_Type `protobuf:"varint,3,opt,name=type,proto3,enum=listing_api.AttributeSpec_Type" json:"type,omitempty"`
	// If the type is SELECT or MULTI_SELECT, provide the valid values here
	ValidValues []string `protobuf:"bytes,4,rep,name=valid_values,json=validValues,proto3" json:"valid_values,omitempty"`
	// If the type is NUMERIC_WITH_UNITS, provide the valid units here
	ValidUnits []string `protobuf:"bytes,5,rep,name=valid_units,json=validUnits,proto3" json:"valid_units,omitempty"`
	// Optional, used to specify the unit of the numeric value when the type is
	// TYPE_NUMERIC This will allow Zentail to convert values with units in
	// Zentail to a single numeric value if your channel does not support units
	Unit           string                       `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	Level          AttributeSpec_Level          `protobuf:"varint,7,opt,name=level,proto3,enum=listing_api.AttributeSpec_Level" json:"level,omitempty"`
	Classification AttributeSpec_Classification `protobuf:"varint,8,opt,name=classification,proto3,enum=listing_api.AttributeSpec_Classification" json:"classification,omitempty"`
	Usage          AttributeSpec_Usage          `protobuf:"varint,9,opt,name=usage,proto3,enum=listing_api.AttributeSpec_Usage" json:"usage,omitempty"`
	UpdatedAt      *timestamppb.Timestamp       `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt      *timestamppb.Timestamp       `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *AttributeSpec) Reset() {
	*x = AttributeSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_listing_taxonomy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeSpec) ProtoMessage() {}

func (x *AttributeSpec) ProtoReflect() protoreflect.Message {
	mi := &file_api_listing_taxonomy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeSpec.ProtoReflect.Descriptor instead.
func (*AttributeSpec) Descriptor() ([]byte, []int) {
	return file_api_listing_taxonomy_proto_rawDescGZIP(), []int{0}
}

func (x *AttributeSpec) GetAttributeSpecId() string {
	if x != nil {
		return x.AttributeSpecId
	}
	return ""
}

func (x *AttributeSpec) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AttributeSpec) GetType() AttributeSpec_Type {
	if x != nil {
		return x.Type
	}
	return AttributeSpec_TYPE_UNSPECIFIED
}

func (x *AttributeSpec) GetValidValues() []string {
	if x != nil {
		return x.ValidValues
	}
	return nil
}

func (x *AttributeSpec) GetValidUnits() []string {
	if x != nil {
		return x.ValidUnits
	}
	return nil
}

func (x *AttributeSpec) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *AttributeSpec) GetLevel() AttributeSpec_Level {
	if x != nil {
		return x.Level
	}
	return AttributeSpec_LEVEL_UNSPECIFIED
}

func (x *AttributeSpec) GetClassification() AttributeSpec_Classification {
	if x != nil {
		return x.Classification
	}
	return AttributeSpec_CLASSIFICATION_UNSPECIFIED
}

func (x *AttributeSpec) GetUsage() AttributeSpec_Usage {
	if x != nil {
		return x.Usage
	}
	return AttributeSpec_USAGE_UNSPECIFIED
}

func (x *AttributeSpec) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *AttributeSpec) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// ProductType are detailed classifications dictated by the channel.
// They should be at least as specific as the Category.
// Generally they should be more specific than the Category.
// Product Types are used to specify which attributes are required and
// recommended for a given Listing.
// If this is not a concept in your system, you should create a product type for
// each category.
type ProductType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductTypeId string                 `protobuf:"bytes,1,opt,name=product_type_id,json=productTypeId,proto3" json:"product_type_id,omitempty"`
	DisplayName   string                 `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ProductType) Reset() {
	*x = ProductType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_listing_taxonomy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductType) ProtoMessage() {}

func (x *ProductType) ProtoReflect() protoreflect.Message {
	mi := &file_api_listing_taxonomy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductType.ProtoReflect.Descriptor instead.
func (*ProductType) Descriptor() ([]byte, []int) {
	return file_api_listing_taxonomy_proto_rawDescGZIP(), []int{1}
}

func (x *ProductType) GetProductTypeId() string {
	if x != nil {
		return x.ProductTypeId
	}
	return ""
}

func (x *ProductType) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ProductType) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ProductType) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_api_listing_taxonomy_proto protoreflect.FileDescriptor

var file_api_listing_taxonomy_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x61,
	0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x08, 0x0a, 0x0d, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2a, 0x0a, 0x11,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x75, 0x6e, 0x69,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55,
	0x6e, 0x69, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x51, 0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xa7, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x10, 0x02,
	0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x49, 0x43,
	0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x45, 0x52,
	0x49, 0x43, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x53, 0x10, 0x04, 0x12,
	0x13, 0x0a, 0x0f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x5f, 0x54, 0x45,
	0x58, 0x54, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x55, 0x4c,
	0x54, 0x49, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x45, 0x59, 0x10, 0x07, 0x22, 0x44, 0x0a, 0x05, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x56, 0x41, 0x52, 0x49, 0x41, 0x4e, 0x54, 0x10,
	0x02, 0x22, 0xc3, 0x01, 0x0a, 0x0e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x46, 0x49,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x46, 0x49,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x10, 0x01,
	0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x45, 0x4e, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x02, 0x12, 0x1a,
	0x0a, 0x16, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x50, 0x52, 0x49, 0x43, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4c,
	0x41, 0x53, 0x53, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x4f, 0x47,
	0x49, 0x53, 0x54, 0x49, 0x43, 0x53, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4c, 0x41, 0x53,
	0x53, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54,
	0x49, 0x46, 0x49, 0x45, 0x52, 0x10, 0x05, 0x22, 0x5e, 0x0a, 0x05, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x53, 0x41, 0x47, 0x45,
	0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x15, 0x0a, 0x11, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x4d, 0x4d,
	0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x03, 0x22, 0xce, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64, 0x74, 0x72, 0x61, 0x64, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_listing_taxonomy_proto_rawDescOnce sync.Once
	file_api_listing_taxonomy_proto_rawDescData = file_api_listing_taxonomy_proto_rawDesc
)

func file_api_listing_taxonomy_proto_rawDescGZIP() []byte {
	file_api_listing_taxonomy_proto_rawDescOnce.Do(func() {
		file_api_listing_taxonomy_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_listing_taxonomy_proto_rawDescData)
	})
	return file_api_listing_taxonomy_proto_rawDescData
}

var file_api_listing_taxonomy_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_api_listing_taxonomy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_listing_taxonomy_proto_goTypes = []interface{}{
	(AttributeSpec_Type)(0),           // 0: listing_api.AttributeSpec.Type
	(AttributeSpec_Level)(0),          // 1: listing_api.AttributeSpec.Level
	(AttributeSpec_Classification)(0), // 2: listing_api.AttributeSpec.Classification
	(AttributeSpec_Usage)(0),          // 3: listing_api.AttributeSpec.Usage
	(*AttributeSpec)(nil),             // 4: listing_api.AttributeSpec
	(*ProductType)(nil),               // 5: listing_api.ProductType
	(*timestamppb.Timestamp)(nil),     // 6: google.protobuf.Timestamp
}
var file_api_listing_taxonomy_proto_depIdxs = []int32{
	0, // 0: listing_api.AttributeSpec.type:type_name -> listing_api.AttributeSpec.Type
	1, // 1: listing_api.AttributeSpec.level:type_name -> listing_api.AttributeSpec.Level
	2, // 2: listing_api.AttributeSpec.classification:type_name -> listing_api.AttributeSpec.Classification
	3, // 3: listing_api.AttributeSpec.usage:type_name -> listing_api.AttributeSpec.Usage
	6, // 4: listing_api.AttributeSpec.updated_at:type_name -> google.protobuf.Timestamp
	6, // 5: listing_api.AttributeSpec.created_at:type_name -> google.protobuf.Timestamp
	6, // 6: listing_api.ProductType.updated_at:type_name -> google.protobuf.Timestamp
	6, // 7: listing_api.ProductType.created_at:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_api_listing_taxonomy_proto_init() }
func file_api_listing_taxonomy_proto_init() {
	if File_api_listing_taxonomy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_listing_taxonomy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_listing_taxonomy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_listing_taxonomy_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_listing_taxonomy_proto_goTypes,
		DependencyIndexes: file_api_listing_taxonomy_proto_depIdxs,
		EnumInfos:         file_api_listing_taxonomy_proto_enumTypes,
		MessageInfos:      file_api_listing_taxonomy_proto_msgTypes,
	}.Build()
	File_api_listing_taxonomy_proto = out.File
	file_api_listing_taxonomy_proto_rawDesc = nil
	file_api_listing_taxonomy_proto_goTypes = nil
	file_api_listing_taxonomy_proto_depIdxs = nil
}
