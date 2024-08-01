// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: api/listing/listing.proto

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

// Variant is a representation of a specific SKU in a listing
// LATER: comment this more
type Variant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku       string `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty"`
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// Any Integration-level Settings will be merged
	// with Variant-level Settings to produce these Settings.
	Settings map[string]string `protobuf:"bytes,3,rep,name=settings,proto3" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// LATER: do we need desired_status / intended action?
	Inventory   *Variant_Inventory  `protobuf:"bytes,4,opt,name=inventory,proto3" json:"inventory,omitempty"`
	Identifiers *Variant_Attributes `protobuf:"bytes,8,opt,name=identifiers,proto3" json:"identifiers,omitempty"`
	ProductData *Variant_Attributes `protobuf:"bytes,5,opt,name=product_data,json=productData,proto3" json:"product_data,omitempty"`
	Pricing     *Variant_Attributes `protobuf:"bytes,6,opt,name=pricing,proto3" json:"pricing,omitempty"`
	Logistics   *Variant_Attributes `protobuf:"bytes,7,opt,name=logistics,proto3" json:"logistics,omitempty"`
}

func (x *Variant) Reset() {
	*x = Variant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_listing_listing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Variant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Variant) ProtoMessage() {}

func (x *Variant) ProtoReflect() protoreflect.Message {
	mi := &file_api_listing_listing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Variant.ProtoReflect.Descriptor instead.
func (*Variant) Descriptor() ([]byte, []int) {
	return file_api_listing_listing_proto_rawDescGZIP(), []int{0}
}

func (x *Variant) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *Variant) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *Variant) GetSettings() map[string]string {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *Variant) GetInventory() *Variant_Inventory {
	if x != nil {
		return x.Inventory
	}
	return nil
}

func (x *Variant) GetIdentifiers() *Variant_Attributes {
	if x != nil {
		return x.Identifiers
	}
	return nil
}

func (x *Variant) GetProductData() *Variant_Attributes {
	if x != nil {
		return x.ProductData
	}
	return nil
}

func (x *Variant) GetPricing() *Variant_Attributes {
	if x != nil {
		return x.Pricing
	}
	return nil
}

func (x *Variant) GetLogistics() *Variant_Attributes {
	if x != nil {
		return x.Logistics
	}
	return nil
}

// Listing is a representation of a product to be sold on a Channel
// LATER: comment this more
type Listing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductData *Listing_ProductData   `protobuf:"bytes,2,opt,name=product_data,json=productData,proto3" json:"product_data,omitempty"`
	Variants    []*Variant             `protobuf:"bytes,3,rep,name=variants,proto3" json:"variants,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Listing) Reset() {
	*x = Listing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_listing_listing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Listing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Listing) ProtoMessage() {}

func (x *Listing) ProtoReflect() protoreflect.Message {
	mi := &file_api_listing_listing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Listing.ProtoReflect.Descriptor instead.
func (*Listing) Descriptor() ([]byte, []int) {
	return file_api_listing_listing_proto_rawDescGZIP(), []int{1}
}

func (x *Listing) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Listing) GetProductData() *Listing_ProductData {
	if x != nil {
		return x.ProductData
	}
	return nil
}

func (x *Listing) GetVariants() []*Variant {
	if x != nil {
		return x.Variants
	}
	return nil
}

func (x *Listing) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Listing) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Attribute has data that describes a Listing or a Listing's Variants
type Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Value:
	//	*Attribute_TextValue
	//	*Attribute_NumericValue
	//	*Attribute_NumericWithUnitsValue
	//	*Attribute_MultiTextValue
	Value isAttribute_Value `protobuf_oneof:"value"`
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_listing_listing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

func (x *Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_api_listing_listing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute.ProtoReflect.Descriptor instead.
func (*Attribute) Descriptor() ([]byte, []int) {
	return file_api_listing_listing_proto_rawDescGZIP(), []int{2}
}

func (x *Attribute) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *Attribute) GetValue() isAttribute_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Attribute) GetTextValue() string {
	if x, ok := x.GetValue().(*Attribute_TextValue); ok {
		return x.TextValue
	}
	return ""
}

func (x *Attribute) GetNumericValue() float64 {
	if x, ok := x.GetValue().(*Attribute_NumericValue); ok {
		return x.NumericValue
	}
	return 0
}

func (x *Attribute) GetNumericWithUnitsValue() *Attribute_NumericWithUnits {
	if x, ok := x.GetValue().(*Attribute_NumericWithUnitsValue); ok {
		return x.NumericWithUnitsValue
	}
	return nil
}

func (x *Attribute) GetMultiTextValue() *Attribute_MultiText {
	if x, ok := x.GetValue().(*Attribute_MultiTextValue); ok {
		return x.MultiTextValue
	}
	return nil
}

type isAttribute_Value interface {
	isAttribute_Value()
}

type Attribute_TextValue struct {
	TextValue string `protobuf:"bytes,2,opt,name=text_value,json=textValue,proto3,oneof"`
}

type Attribute_NumericValue struct {
	NumericValue float64 `protobuf:"fixed64,3,opt,name=numeric_value,json=numericValue,proto3,oneof"`
}

type Attribute_NumericWithUnitsValue struct {
	NumericWithUnitsValue *Attribute_NumericWithUnits `protobuf:"bytes,4,opt,name=numeric_with_units_value,json=numericWithUnitsValue,proto3,oneof"`
}

type Attribute_MultiTextValue struct {
	MultiTextValue *Attribute_MultiText `protobuf:"bytes,5,opt,name=multi_text_value,json=multiTextValue,proto3,oneof"`
}

func (*Attribute_TextValue) isAttribute_Value() {}

func (*Attribute_NumericValue) isAttribute_Value() {}

func (*Attribute_NumericWithUnitsValue) isAttribute_Value() {}

func (*Attribute_MultiTextValue) isAttribute_Value() {}

// Inventory contains information about the availability of this variant
type Variant_Inventory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// total_quantity should be the sum of merchant_fulfillable_inventory
	// and storefront_fulfillable_quantity
	// -- removing for now since it isn't implemented and will always be 0 --
	// int64 total_quantity                            = 2;
	MerchantFulfillableQuantity int64 `protobuf:"varint,3,opt,name=merchant_fulfillable_quantity,json=merchantFulfillableQuantity,proto3" json:"merchant_fulfillable_quantity,omitempty"`
	// -- removing for now since isn't implemented and will always be 0 --
	// int64 storefront_fulfillable_quantity           = 4;
	UpdatedAt           *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedExternallyAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_externally_at,json=updatedExternallyAt,proto3" json:"updated_externally_at,omitempty"`
}

func (x *Variant_Inventory) Reset() {
	*x = Variant_Inventory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_listing_listing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Variant_Inventory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Variant_Inventory) ProtoMessage() {}

func (x *Variant_Inventory) ProtoReflect() protoreflect.Message {
	mi := &file_api_listing_listing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Variant_Inventory.ProtoReflect.Descriptor instead.
func (*Variant_Inventory) Descriptor() ([]byte, []int) {
	return file_api_listing_listing_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Variant_Inventory) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Variant_Inventory) GetMerchantFulfillableQuantity() int64 {
	if x != nil {
		return x.MerchantFulfillableQuantity
	}
	return 0
}

func (x *Variant_Inventory) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Variant_Inventory) GetUpdatedExternallyAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedExternallyAt
	}
	return nil
}

// Attributes is a container for a list of attributes of a similar
// classification
type Variant_Attributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled    bool                   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Attributes []*Attribute           `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Variant_Attributes) Reset() {
	*x = Variant_Attributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_listing_listing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Variant_Attributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Variant_Attributes) ProtoMessage() {}

func (x *Variant_Attributes) ProtoReflect() protoreflect.Message {
	mi := &file_api_listing_listing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Variant_Attributes.ProtoReflect.Descriptor instead.
func (*Variant_Attributes) Descriptor() ([]byte, []int) {
	return file_api_listing_listing_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Variant_Attributes) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Variant_Attributes) GetAttributes() []*Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Variant_Attributes) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Listing_ProductData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled         bool         `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	CategoryId      string       `protobuf:"bytes,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	ProductTypeId   string       `protobuf:"bytes,3,opt,name=product_type_id,json=productTypeId,proto3" json:"product_type_id,omitempty"`
	PivotAttributes []string     `protobuf:"bytes,4,rep,name=pivot_attributes,json=pivotAttributes,proto3" json:"pivot_attributes,omitempty"`
	Attributes      []*Attribute `protobuf:"bytes,5,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *Listing_ProductData) Reset() {
	*x = Listing_ProductData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_listing_listing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Listing_ProductData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Listing_ProductData) ProtoMessage() {}

func (x *Listing_ProductData) ProtoReflect() protoreflect.Message {
	mi := &file_api_listing_listing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Listing_ProductData.ProtoReflect.Descriptor instead.
func (*Listing_ProductData) Descriptor() ([]byte, []int) {
	return file_api_listing_listing_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Listing_ProductData) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Listing_ProductData) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *Listing_ProductData) GetProductTypeId() string {
	if x != nil {
		return x.ProductTypeId
	}
	return ""
}

func (x *Listing_ProductData) GetPivotAttributes() []string {
	if x != nil {
		return x.PivotAttributes
	}
	return nil
}

func (x *Listing_ProductData) GetAttributes() []*Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

// NumericWithUnits supports values like 1 lb
type Attribute_NumericWithUnits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Numeric float64 `protobuf:"fixed64,1,opt,name=numeric,proto3" json:"numeric,omitempty"`
	Units   string  `protobuf:"bytes,2,opt,name=units,proto3" json:"units,omitempty"`
}

func (x *Attribute_NumericWithUnits) Reset() {
	*x = Attribute_NumericWithUnits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_listing_listing_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute_NumericWithUnits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute_NumericWithUnits) ProtoMessage() {}

func (x *Attribute_NumericWithUnits) ProtoReflect() protoreflect.Message {
	mi := &file_api_listing_listing_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute_NumericWithUnits.ProtoReflect.Descriptor instead.
func (*Attribute_NumericWithUnits) Descriptor() ([]byte, []int) {
	return file_api_listing_listing_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Attribute_NumericWithUnits) GetNumeric() float64 {
	if x != nil {
		return x.Numeric
	}
	return 0
}

func (x *Attribute_NumericWithUnits) GetUnits() string {
	if x != nil {
		return x.Units
	}
	return ""
}

// MultiText supports a value that has multiple sub-values
// (e.g. bullet points)
type Attribute_MultiText struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Attribute_MultiText) Reset() {
	*x = Attribute_MultiText{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_listing_listing_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute_MultiText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute_MultiText) ProtoMessage() {}

func (x *Attribute_MultiText) ProtoReflect() protoreflect.Message {
	mi := &file_api_listing_listing_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute_MultiText.ProtoReflect.Descriptor instead.
func (*Attribute_MultiText) Descriptor() ([]byte, []int) {
	return file_api_listing_listing_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Attribute_MultiText) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_api_listing_listing_proto protoreflect.FileDescriptor

var file_api_listing_listing_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x07, 0x0a, 0x07, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x2e,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x41, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0b, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x42, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x07, 0x70,
	0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x07, 0x70,
	0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0xf4, 0x01, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x42, 0x0a, 0x1d, 0x6d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x1b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x46, 0x75, 0x6c, 0x66, 0x69,
	0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4e, 0x0a, 0x15, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x6c, 0x79, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x6c, 0x79, 0x41, 0x74, 0x1a, 0x99, 0x01, 0x0a, 0x0a, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x36, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xdc, 0x03, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x43, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x08,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0xd3,
	0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x69, 0x76, 0x6f, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x69, 0x76,
	0x6f, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0a,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x22, 0x87, 0x03, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x74, 0x65, 0x78, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0c, 0x6e, 0x75,
	0x6d, 0x65, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x62, 0x0a, 0x18, 0x6e, 0x75,
	0x6d, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x57, 0x69, 0x74, 0x68,
	0x55, 0x6e, 0x69, 0x74, 0x73, 0x48, 0x00, 0x52, 0x15, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63,
	0x57, 0x69, 0x74, 0x68, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4c,
	0x0a, 0x10, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x54, 0x65, 0x78, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x54, 0x65, 0x78, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x42, 0x0a, 0x10,
	0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e, 0x69, 0x74, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e,
	0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73,
	0x1a, 0x23, 0x0a, 0x09, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x54, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x64, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_listing_listing_proto_rawDescOnce sync.Once
	file_api_listing_listing_proto_rawDescData = file_api_listing_listing_proto_rawDesc
)

func file_api_listing_listing_proto_rawDescGZIP() []byte {
	file_api_listing_listing_proto_rawDescOnce.Do(func() {
		file_api_listing_listing_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_listing_listing_proto_rawDescData)
	})
	return file_api_listing_listing_proto_rawDescData
}

var file_api_listing_listing_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_listing_listing_proto_goTypes = []interface{}{
	(*Variant)(nil),                    // 0: listing_api.Variant
	(*Listing)(nil),                    // 1: listing_api.Listing
	(*Attribute)(nil),                  // 2: listing_api.Attribute
	nil,                                // 3: listing_api.Variant.SettingsEntry
	(*Variant_Inventory)(nil),          // 4: listing_api.Variant.Inventory
	(*Variant_Attributes)(nil),         // 5: listing_api.Variant.Attributes
	(*Listing_ProductData)(nil),        // 6: listing_api.Listing.ProductData
	(*Attribute_NumericWithUnits)(nil), // 7: listing_api.Attribute.NumericWithUnits
	(*Attribute_MultiText)(nil),        // 8: listing_api.Attribute.MultiText
	(*timestamppb.Timestamp)(nil),      // 9: google.protobuf.Timestamp
}
var file_api_listing_listing_proto_depIdxs = []int32{
	3,  // 0: listing_api.Variant.settings:type_name -> listing_api.Variant.SettingsEntry
	4,  // 1: listing_api.Variant.inventory:type_name -> listing_api.Variant.Inventory
	5,  // 2: listing_api.Variant.identifiers:type_name -> listing_api.Variant.Attributes
	5,  // 3: listing_api.Variant.product_data:type_name -> listing_api.Variant.Attributes
	5,  // 4: listing_api.Variant.pricing:type_name -> listing_api.Variant.Attributes
	5,  // 5: listing_api.Variant.logistics:type_name -> listing_api.Variant.Attributes
	6,  // 6: listing_api.Listing.product_data:type_name -> listing_api.Listing.ProductData
	0,  // 7: listing_api.Listing.variants:type_name -> listing_api.Variant
	9,  // 8: listing_api.Listing.created_at:type_name -> google.protobuf.Timestamp
	9,  // 9: listing_api.Listing.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 10: listing_api.Attribute.numeric_with_units_value:type_name -> listing_api.Attribute.NumericWithUnits
	8,  // 11: listing_api.Attribute.multi_text_value:type_name -> listing_api.Attribute.MultiText
	9,  // 12: listing_api.Variant.Inventory.updated_at:type_name -> google.protobuf.Timestamp
	9,  // 13: listing_api.Variant.Inventory.updated_externally_at:type_name -> google.protobuf.Timestamp
	2,  // 14: listing_api.Variant.Attributes.attributes:type_name -> listing_api.Attribute
	9,  // 15: listing_api.Variant.Attributes.updated_at:type_name -> google.protobuf.Timestamp
	2,  // 16: listing_api.Listing.ProductData.attributes:type_name -> listing_api.Attribute
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_api_listing_listing_proto_init() }
func file_api_listing_listing_proto_init() {
	if File_api_listing_listing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_listing_listing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Variant); i {
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
		file_api_listing_listing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Listing); i {
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
		file_api_listing_listing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute); i {
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
		file_api_listing_listing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Variant_Inventory); i {
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
		file_api_listing_listing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Variant_Attributes); i {
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
		file_api_listing_listing_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Listing_ProductData); i {
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
		file_api_listing_listing_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute_NumericWithUnits); i {
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
		file_api_listing_listing_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute_MultiText); i {
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
	file_api_listing_listing_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Attribute_TextValue)(nil),
		(*Attribute_NumericValue)(nil),
		(*Attribute_NumericWithUnitsValue)(nil),
		(*Attribute_MultiTextValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_listing_listing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_listing_listing_proto_goTypes,
		DependencyIndexes: file_api_listing_listing_proto_depIdxs,
		MessageInfos:      file_api_listing_listing_proto_msgTypes,
	}.Build()
	File_api_listing_listing_proto = out.File
	file_api_listing_listing_proto_rawDesc = nil
	file_api_listing_listing_proto_goTypes = nil
	file_api_listing_listing_proto_depIdxs = nil
}
