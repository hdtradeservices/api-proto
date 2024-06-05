# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/listing/listing.proto](#api_listing_listing-proto)
    - [Attribute](#listing_api-Attribute)
    - [Attribute.MultiText](#listing_api-Attribute-MultiText)
    - [Attribute.NumericWithUnits](#listing_api-Attribute-NumericWithUnits)
    - [Listing](#listing_api-Listing)
    - [Listing.Variant](#listing_api-Listing-Variant)
  
    - [Attribute.Source](#listing_api-Attribute-Source)
  
- [api/listing/service.proto](#api_listing_service-proto)
    - [GetRequest](#listing_api-GetRequest)
    - [ListRequest](#listing_api-ListRequest)
    - [ListResponse](#listing_api-ListResponse)
  
    - [ListingService](#listing_api-ListingService)
  
- [api/listing/settings.proto](#api_listing_settings-proto)
    - [Settings](#listing_api-Settings)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_listing_listing-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/listing/listing.proto



<a name="listing_api-Attribute"></a>

### Attribute
Attribute has data that describes a Listing or a Listing&#39;s Variants


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| text_value | [string](#string) |  |  |
| numeric_value | [double](#double) |  |  |
| numeric_with_units_value | [Attribute.NumericWithUnits](#listing_api-Attribute-NumericWithUnits) |  |  |
| multi_text_value | [Attribute.MultiText](#listing_api-Attribute-MultiText) |  |  |
| source | [Attribute.Source](#listing_api-Attribute-Source) |  |  |






<a name="listing_api-Attribute-MultiText"></a>

### Attribute.MultiText
MultiText supports a value that has multiple sub-values
(e.g. bullet points)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| values | [string](#string) | repeated |  |






<a name="listing_api-Attribute-NumericWithUnits"></a>

### Attribute.NumericWithUnits
NumericWithUnits supports values like 1 lb


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| numeric | [double](#double) |  |  |
| units | [string](#string) |  |  |






<a name="listing_api-Listing"></a>

### Listing
Listing is a representation of a product to be sold on a Channel


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| category_id | [string](#string) |  |  |
| product_type_id | [string](#string) |  |  |
| attributes | [Attribute](#listing_api-Attribute) | repeated |  |
| variants | [Listing.Variant](#listing_api-Listing-Variant) | repeated | the channel These should be in the order that are intended to be displayed one |
| pivot_attributes | [string](#string) | repeated | These are attribute IDs that should have a corresponding Attribute in each Variant&#39;s list of Attributes |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="listing_api-Listing-Variant"></a>

### Listing.Variant



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku | [string](#string) |  |  |
| attributes | [Attribute](#listing_api-Attribute) | repeated |  |
| total_inventory | [int64](#int64) |  | total_inventory should be the sum of merchant_fulfillable_inventory and storefront_fulfillable_inventory |
| merchant_fulfillable_inventory | [int64](#int64) |  |  |
| storefront_fulfillable_inventory | [int64](#int64) |  |  |
| settings | [Settings](#listing_api-Settings) |  | Any Integration-level Settings will be merged with Variant-level Settings to produce these Settings. |





 


<a name="listing_api-Attribute-Source"></a>

### Attribute.Source


| Name | Number | Description |
| ---- | ------ | ----------- |
| SOURCE_UNSPECIFIED | 0 |  |
| SOURCE_MAPPED_FROM_CODEC | 1 |  |
| SOURCE_ATTRIBUTE_OVERRIDE | 2 |  |
| SOURCE_ATTRIBUTE_DEFAULT_VALUE | 3 |  |


 

 

 



<a name="api_listing_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/listing/service.proto



<a name="listing_api-GetRequest"></a>

### GetRequest
GetRequest is the request object for the Get method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listing_id | [string](#string) |  |  |






<a name="listing_api-ListRequest"></a>

### ListRequest
ListRequest is the request object for the List method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| last_update_ts | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| next_token | [string](#string) |  |  |
| page_length | [int32](#int32) |  |  |
| skus | [string](#string) | repeated | protolint:disable:next REPEATED_FIELD_NAMES_PLURALIZED |






<a name="listing_api-ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listings | [Listing](#listing_api-Listing) | repeated |  |
| next_token | [string](#string) |  | next_token is a token that can be provided to List request to get the next page of results. Next tokens are only valid for 5 minutes after being returned. If no `nextToken` is provided, there are no more results to return. |





 

 

 


<a name="listing_api-ListingService"></a>

### ListingService
ListingService provides a service for listing related operations
This includes, retrieving listings as well as sending back statuses to
Zentail.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#listing_api-GetRequest) | [Listing](#listing_api-Listing) | Get retrieves a single listing by its ID |
| List | [ListRequest](#listing_api-ListRequest) | [ListResponse](#listing_api-ListResponse) | List retrieves a list of listings based on the provided query parameters |

 



<a name="api_listing_settings-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/listing/settings.proto



<a name="listing_api-Settings"></a>

### Settings
Settings represent Listing Settings at any level (Variant, Integration, etc).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| string_value | [string](#string) |  |  |
| bool_value | [bool](#bool) |  |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

