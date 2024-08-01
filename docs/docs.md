# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/listing/listing.proto](#api_listing_listing-proto)
    - [Attribute](#listing_api-Attribute)
    - [Attribute.MultiText](#listing_api-Attribute-MultiText)
    - [Attribute.NumericWithUnits](#listing_api-Attribute-NumericWithUnits)
    - [Listing](#listing_api-Listing)
    - [Listing.ProductData](#listing_api-Listing-ProductData)
    - [Variant](#listing_api-Variant)
    - [Variant.Attributes](#listing_api-Variant-Attributes)
    - [Variant.Inventory](#listing_api-Variant-Inventory)
    - [Variant.SettingsEntry](#listing_api-Variant-SettingsEntry)
  
- [api/listing/service.proto](#api_listing_service-proto)
    - [Error](#listing_api-Error)
    - [GetBySKURequest](#listing_api-GetBySKURequest)
    - [GetRequest](#listing_api-GetRequest)
    - [GetVariantRequest](#listing_api-GetVariantRequest)
    - [ListInventorySinceRequest](#listing_api-ListInventorySinceRequest)
    - [ListListingsResponse](#listing_api-ListListingsResponse)
    - [ListSinceRequest](#listing_api-ListSinceRequest)
    - [ListVariantsResponse](#listing_api-ListVariantsResponse)
    - [ReplaceErrorsRequest](#listing_api-ReplaceErrorsRequest)
    - [ReplaceErrorsResponse](#listing_api-ReplaceErrorsResponse)
    - [ReplaceSubmissionsRequest](#listing_api-ReplaceSubmissionsRequest)
    - [ReplaceSubmissionsResponse](#listing_api-ReplaceSubmissionsResponse)
    - [SetInventorySubmissionDetailsRequest](#listing_api-SetInventorySubmissionDetailsRequest)
    - [SetInventorySubmissionDetailsResponse](#listing_api-SetInventorySubmissionDetailsResponse)
    - [UpdateChannelListingIDRequest](#listing_api-UpdateChannelListingIDRequest)
    - [UpdateChannelListingIDResponse](#listing_api-UpdateChannelListingIDResponse)
    - [UpdateStatusRequest](#listing_api-UpdateStatusRequest)
    - [UpdateStatusResponse](#listing_api-UpdateStatusResponse)
  
    - [Error.Severity](#listing_api-Error-Severity)
    - [Error.Type](#listing_api-Error-Type)
    - [StandardStatus](#listing_api-StandardStatus)
  
    - [ListingService](#listing_api-ListingService)
  
- [api/listing/settings.proto](#api_listing_settings-proto)
    - [Settings](#listing_api-Settings)
  
- [api/listing/submission.proto](#api_listing_submission-proto)
    - [Submission](#listing_api-Submission)
    - [Submission.MetadataEntry](#listing_api-Submission-MetadataEntry)
  
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
LATER: comment this more


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| product_data | [Listing.ProductData](#listing_api-Listing-ProductData) |  |  |
| variants | [Variant](#listing_api-Variant) | repeated |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="listing_api-Listing-ProductData"></a>

### Listing.ProductData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  |  |
| category_id | [string](#string) |  |  |
| product_type_id | [string](#string) |  |  |
| pivot_attributes | [string](#string) | repeated |  |
| attributes | [Attribute](#listing_api-Attribute) | repeated |  |






<a name="listing_api-Variant"></a>

### Variant
Variant is a representation of a specific SKU in a listing
LATER: comment this more


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku | [string](#string) |  |  |
| channel_id | [string](#string) |  |  |
| settings | [Variant.SettingsEntry](#listing_api-Variant-SettingsEntry) | repeated | Any Integration-level Settings will be merged with Variant-level Settings to produce these Settings. |
| inventory | [Variant.Inventory](#listing_api-Variant-Inventory) |  | LATER: do we need desired_status / intended action? |
| identifiers | [Variant.Attributes](#listing_api-Variant-Attributes) |  |  |
| product_data | [Variant.Attributes](#listing_api-Variant-Attributes) |  |  |
| pricing | [Variant.Attributes](#listing_api-Variant-Attributes) |  |  |
| logistics | [Variant.Attributes](#listing_api-Variant-Attributes) |  |  |






<a name="listing_api-Variant-Attributes"></a>

### Variant.Attributes
Attributes is a container for a list of attributes of a similar
classification


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  |  |
| attributes | [Attribute](#listing_api-Attribute) | repeated |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="listing_api-Variant-Inventory"></a>

### Variant.Inventory
Inventory contains information about the availability of this variant


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  |  |
| merchant_fulfillable_quantity | [int64](#int64) |  | total_quantity should be the sum of merchant_fulfillable_inventory and storefront_fulfillable_quantity -- removing for now since it isn&#39;t implemented and will always be 0 -- int64 total_quantity = 2; |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | -- removing for now since isn&#39;t implemented and will always be 0 -- int64 storefront_fulfillable_quantity = 4; |
| updated_externally_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="listing_api-Variant-SettingsEntry"></a>

### Variant.SettingsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 

 

 

 



<a name="api_listing_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/listing/service.proto



<a name="listing_api-Error"></a>

### Error



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attribute_ids | [string](#string) | repeated | One or more attributes that this error applies to |
| severity | [Error.Severity](#listing_api-Error-Severity) |  |  |
| type | [Error.Type](#listing_api-Error-Type) |  |  |
| message | [string](#string) |  | Message is required if Type is TYPE_OTHER |






<a name="listing_api-GetBySKURequest"></a>

### GetBySKURequest
GetBySKURequest is the request object for the GetBySKU method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku | [string](#string) |  |  |






<a name="listing_api-GetRequest"></a>

### GetRequest
GetRequest is the request object for the Get method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listing_id | [string](#string) |  |  |






<a name="listing_api-GetVariantRequest"></a>

### GetVariantRequest
GetVariantRequest is the request object for the GetVariant method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku | [string](#string) |  |  |






<a name="listing_api-ListInventorySinceRequest"></a>

### ListInventorySinceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| since | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| external_changes_only | [bool](#bool) |  | Only return variants with inventory changes that have happened outside of sales on this channel. |
| cursor | [string](#string) |  | TODO: how to get these in the query params |






<a name="listing_api-ListListingsResponse"></a>

### ListListingsResponse
ListListingsResponse is the response object containing a list of Listings


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listings | [Listing](#listing_api-Listing) | repeated |  |
| next_page_cursor | [string](#string) |  | The cursor token for the next page of results. If empty, there are no more results. |






<a name="listing_api-ListSinceRequest"></a>

### ListSinceRequest
ListSinceRequestRequest is the request object for the ListSinceRequest method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| since | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| cursor | [string](#string) |  | TODO: how to get these in the query params |






<a name="listing_api-ListVariantsResponse"></a>

### ListVariantsResponse
ListVariantsResponse is the response object containing a list of Variants


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variants | [Variant](#listing_api-Variant) | repeated |  |
| next_page_cursor | [string](#string) |  | The cursor token for the next page of results. If empty, there are no more results. |






<a name="listing_api-ReplaceErrorsRequest"></a>

### ReplaceErrorsRequest
ReplaceErrorsRequest provides all the channel-generated errors for a SKU


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku | [string](#string) |  |  |
| errors | [Error](#listing_api-Error) | repeated |  |






<a name="listing_api-ReplaceErrorsResponse"></a>

### ReplaceErrorsResponse
ReplaceErrorsResponse is empty






<a name="listing_api-ReplaceSubmissionsRequest"></a>

### ReplaceSubmissionsRequest
ReplaceSubmissionsRequest is used to replace all submissions for the given
type and sku


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| submissions | [Submission](#listing_api-Submission) | repeated |  |






<a name="listing_api-ReplaceSubmissionsResponse"></a>

### ReplaceSubmissionsResponse
ReplaceSubmissionsResponse contains the submissions created


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| submissions | [Submission](#listing_api-Submission) | repeated |  |






<a name="listing_api-SetInventorySubmissionDetailsRequest"></a>

### SetInventorySubmissionDetailsRequest
SetInventorySubmissionDetailsRequest is used to create inventory submission
details


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| submission_id | [int64](#int64) |  |  |
| inventory_level_sent | [int64](#int64) |  |  |
| successful | [bool](#bool) |  |  |






<a name="listing_api-SetInventorySubmissionDetailsResponse"></a>

### SetInventorySubmissionDetailsResponse
SetInventorySubmissionDetailsResponse is currently an empty response






<a name="listing_api-UpdateChannelListingIDRequest"></a>

### UpdateChannelListingIDRequest
UpdateChannelListingIDRequest provides the status of a SKU


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku | [string](#string) |  |  |
| channel_id | [string](#string) |  |  |






<a name="listing_api-UpdateChannelListingIDResponse"></a>

### UpdateChannelListingIDResponse
UpdateChannelListingIDResponse is empty






<a name="listing_api-UpdateStatusRequest"></a>

### UpdateStatusRequest
UpdateStatusRequest provides the status of a SKU


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku | [string](#string) |  |  |
| status | [StandardStatus](#listing_api-StandardStatus) |  |  |
| channel_status | [string](#string) |  |  |
| channel_id | [string](#string) |  | The channel&#39;s ID for this particular SKU. Should be more specific than Listing. |
| listing_url | [string](#string) |  |  |






<a name="listing_api-UpdateStatusResponse"></a>

### UpdateStatusResponse
UpdateStatusResponse is empty





 


<a name="listing_api-Error-Severity"></a>

### Error.Severity


| Name | Number | Description |
| ---- | ------ | ----------- |
| SEVERITY_UNSPECIFIED | 0 |  |
| SEVERITY_NOTICE | 1 |  |
| SEVERITY_WARNING | 2 |  |
| SEVERITY_ERROR | 3 |  |



<a name="listing_api-Error-Type"></a>

### Error.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| TYPE_MISSING_VALUE | 1 |  |
| TYPE_INVALID_VALUE | 2 |  |
| TYPE_RESTRICTED_VALUE | 3 |  |
| TYPE_OTHER | 4 |  |



<a name="listing_api-StandardStatus"></a>

### StandardStatus
StandardStatus is the standardized status of a listing as defined by Zentail

| Name | Number | Description |
| ---- | ------ | ----------- |
| STANDARD_STATUS_UNSPECIFIED | 0 |  |
| STANDARD_STATUS_PUBLISHED | 1 |  |
| STANDARD_STATUS_UNPUBLISHED | 2 |  |
| STANDARD_STATUS_SUPPRESSED | 3 |  |
| STANDARD_STATUS_RETIRED | 4 |  |


 

 


<a name="listing_api-ListingService"></a>

### ListingService
ListingService provides a service for listing related operations
This includes, retrieving listings as well as sending back statuses to
Zentail.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#listing_api-GetRequest) | [Listing](#listing_api-Listing) | Get retrieves a single listing by its ID |
| GetBySKU | [GetBySKURequest](#listing_api-GetBySKURequest) | [Listing](#listing_api-Listing) | GetBySKU retrieves a single listing which contains a variant with the given SKU |
| GetVariant | [GetVariantRequest](#listing_api-GetVariantRequest) | [Variant](#listing_api-Variant) | GetVariant retrieves a single variant by its SKU |
| ListNewListings | [ListSinceRequest](#listing_api-ListSinceRequest) | [ListListingsResponse](#listing_api-ListListingsResponse) | ListNewListings will list any listing created or updated since the given timestamp where:

1. Product data is enabled for at least one Variant in the Listing

2. No variants have a channel ID |
| ListUpdatedListings | [ListSinceRequest](#listing_api-ListSinceRequest) | [ListListingsResponse](#listing_api-ListListingsResponse) | ListUpdateListings will return any listing that:

1. Has at least one Variant with a channel ID

2. Has a Product Data change since the last timestamp (including Variants)

3. Product Data is enabled for the Listing |
| ListVariantsWithUpdatedInventory | [ListInventorySinceRequest](#listing_api-ListInventorySinceRequest) | [ListVariantsResponse](#listing_api-ListVariantsResponse) | ListVariantsWithUpdatedInventory will return any variant that:

1. Has an inventory change since the last timestamp

2. Inventory Data is enabled for the Variant |
| ListVariantsWithUpdatedPricing | [ListSinceRequest](#listing_api-ListSinceRequest) | [ListVariantsResponse](#listing_api-ListVariantsResponse) | ListVariantsWithUpdatedPricing will return any variant that:

1. Has a pricing change since the last timestamp

2. Pricing Data is enabled for the Variant |
| UpdateStatus | [UpdateStatusRequest](#listing_api-UpdateStatusRequest) | [UpdateStatusResponse](#listing_api-UpdateStatusResponse) | UpdateStatus updates the status of a listing |
| UpdateChannelListingID | [UpdateChannelListingIDRequest](#listing_api-UpdateChannelListingIDRequest) | [UpdateChannelListingIDResponse](#listing_api-UpdateChannelListingIDResponse) | UpdateChannelListingID updates the channel ID for the listing |
| ReplaceErrors | [ReplaceErrorsRequest](#listing_api-ReplaceErrorsRequest) | [ReplaceErrorsResponse](#listing_api-ReplaceErrorsResponse) | ReplaceErrors replaces the errors for a variant |
| ReplaceSubmissions | [ReplaceSubmissionsRequest](#listing_api-ReplaceSubmissionsRequest) | [ReplaceSubmissionsResponse](#listing_api-ReplaceSubmissionsResponse) | ReplaceSubmissions replaces the submissions for a variant |
| SetInventorySubmissionDetails | [SetInventorySubmissionDetailsRequest](#listing_api-SetInventorySubmissionDetailsRequest) | [SetInventorySubmissionDetailsResponse](#listing_api-SetInventorySubmissionDetailsResponse) | SetInventorySubmissionDetails is used to set the inventory details for a given submission |

 



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





 

 

 

 



<a name="api_listing_submission-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/listing/submission.proto



<a name="listing_api-Submission"></a>

### Submission
Submission represents specific data sent to a channel for a particular SKU


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| submission_id | [int64](#int64) |  | submission_id is returned by the API, setting this when calling ReplaceSubmissions will have no effect |
| sku | [string](#string) |  |  |
| content_sent | [string](#string) |  |  |
| content_received | [string](#string) |  |  |
| metadata | [Submission.MetadataEntry](#listing_api-Submission-MetadataEntry) | repeated |  |
| type | [string](#string) |  | type is a user-specified type which is used to match like submissions when doing a replace operation |
| successful | [bool](#bool) |  |  |
| prepared_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| submitted_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| acknowledged_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="listing_api-Submission-MetadataEntry"></a>

### Submission.MetadataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 

 

 

 



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

