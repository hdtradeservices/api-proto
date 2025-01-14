# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/listing/listing.proto](#api_listing_listing-proto)
    - [Attribute](#listing_api-Attribute)
    - [Attribute.Money](#listing_api-Attribute-Money)
    - [Attribute.MultiText](#listing_api-Attribute-MultiText)
    - [Attribute.NumericWithUnits](#listing_api-Attribute-NumericWithUnits)
    - [Listing](#listing_api-Listing)
    - [Listing.ProductData](#listing_api-Listing-ProductData)
    - [Variant](#listing_api-Variant)
    - [Variant.Attributes](#listing_api-Variant-Attributes)
    - [Variant.Inventory](#listing_api-Variant-Inventory)
    - [Variant.Pricing](#listing_api-Variant-Pricing)
    - [Variant.Pricing.ScheduledSale](#listing_api-Variant-Pricing-ScheduledSale)
    - [Variant.SettingsEntry](#listing_api-Variant-SettingsEntry)
  
- [api/listing/service.proto](#api_listing_service-proto)
    - [CreateSubmissionsRequest](#listing_api-CreateSubmissionsRequest)
    - [CreateSubmissionsResponse](#listing_api-CreateSubmissionsResponse)
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
    - [SetInventorySubmissionDetailsRequest](#listing_api-SetInventorySubmissionDetailsRequest)
    - [SetInventorySubmissionDetailsResponse](#listing_api-SetInventorySubmissionDetailsResponse)
    - [UpdateChannelListingIDRequest](#listing_api-UpdateChannelListingIDRequest)
    - [UpdateChannelListingIDResponse](#listing_api-UpdateChannelListingIDResponse)
    - [UpdateStatusRequest](#listing_api-UpdateStatusRequest)
    - [UpdateStatusResponse](#listing_api-UpdateStatusResponse)
    - [UpdateSubmissionRequest](#listing_api-UpdateSubmissionRequest)
    - [UpdateSubmissionRequest.MetadataEntry](#listing_api-UpdateSubmissionRequest-MetadataEntry)
  
    - [Error.Severity](#listing_api-Error-Severity)
    - [Error.Type](#listing_api-Error-Type)
    - [StandardStatus](#listing_api-StandardStatus)
    - [UpdateSubmissionRequest.Status](#listing_api-UpdateSubmissionRequest-Status)
  
    - [ListingService](#listing_api-ListingService)
  
- [api/listing/settings.proto](#api_listing_settings-proto)
    - [Settings](#listing_api-Settings)
  
- [api/listing/submission.proto](#api_listing_submission-proto)
    - [Submission](#listing_api-Submission)
    - [Submission.MetadataEntry](#listing_api-Submission-MetadataEntry)
  
- [api/listing/taxonomy.proto](#api_listing_taxonomy-proto)
    - [AttributeSpec](#listing_api-AttributeSpec)
    - [ProductType](#listing_api-ProductType)
  
    - [AttributeSpec.Classification](#listing_api-AttributeSpec-Classification)
    - [AttributeSpec.Level](#listing_api-AttributeSpec-Level)
    - [AttributeSpec.Type](#listing_api-AttributeSpec-Type)
    - [AttributeSpec.Usage](#listing_api-AttributeSpec-Usage)
  
- [api/listing/taxonomy_service.proto](#api_listing_taxonomy_service-proto)
    - [AttributeSpecRequest](#listing_api-AttributeSpecRequest)
    - [CreateAttributeSpecRequest](#listing_api-CreateAttributeSpecRequest)
    - [CreateProductTypeRequest](#listing_api-CreateProductTypeRequest)
    - [CreateProductTypeResponse](#listing_api-CreateProductTypeResponse)
    - [DeleteAttributeSpecRequest](#listing_api-DeleteAttributeSpecRequest)
    - [DeleteAttributeSpecResponse](#listing_api-DeleteAttributeSpecResponse)
    - [DeleteProductTypeRequest](#listing_api-DeleteProductTypeRequest)
    - [DeleteProductTypeResponse](#listing_api-DeleteProductTypeResponse)
    - [ListProductTypesRequest](#listing_api-ListProductTypesRequest)
    - [ListProductTypesResponse](#listing_api-ListProductTypesResponse)
    - [UpdateAttributeSpecRequest](#listing_api-UpdateAttributeSpecRequest)
  
    - [TaxonomyService](#listing_api-TaxonomyService)
  
- [api/orders/orders_service.proto](#api_orders_orders_service-proto)
    - [CancelItemsRequest](#orders_api-CancelItemsRequest)
    - [CancelItemsRequest.CancelQuantitiesEntry](#orders_api-CancelItemsRequest-CancelQuantitiesEntry)
    - [CancelItemsResponse](#orders_api-CancelItemsResponse)
  
    - [OrdersService](#orders_api-OrdersService)
  
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
| money_value | [Attribute.Money](#listing_api-Attribute-Money) |  |  |






<a name="listing_api-Attribute-Money"></a>

### Attribute.Money
Money supports values like with amount and currency_code


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount | [string](#string) |  |  |
| currency_code | [string](#string) |  |  |






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
| pricing | [Variant.Pricing](#listing_api-Variant-Pricing) |  |  |
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
| total_quantity | [int64](#int64) |  | total_quantity should be the sum of merchant_fulfillable_inventory and storefront_fulfillable_quantity |
| merchant_fulfillable_quantity | [int64](#int64) |  |  |
| storefront_fulfillable_quantity | [int64](#int64) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_externally_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="listing_api-Variant-Pricing"></a>

### Variant.Pricing
Pricing is a container for a list of pricing attributes as well as
structured information about scheduled sales


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enabled | [bool](#bool) |  |  |
| attributes | [Attribute](#listing_api-Attribute) | repeated |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| sales | [Variant.Pricing.ScheduledSale](#listing_api-Variant-Pricing-ScheduledSale) | repeated |  |






<a name="listing_api-Variant-Pricing-ScheduledSale"></a>

### Variant.Pricing.ScheduledSale



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| strike_through_price | [double](#double) |  |  |
| sale_price | [double](#double) |  |  |
| start_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| end_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="listing_api-Variant-SettingsEntry"></a>

### Variant.SettingsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 

 

 

 



<a name="api_listing_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/listing/service.proto



<a name="listing_api-CreateSubmissionsRequest"></a>

### CreateSubmissionsRequest
CreateSubmissionsRequest is used to replace all submissions for the given
type and sku


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| submissions | [Submission](#listing_api-Submission) | repeated |  |






<a name="listing_api-CreateSubmissionsResponse"></a>

### CreateSubmissionsResponse
CreateSubmissionsResponse contains the submissions created


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| submissions | [Submission](#listing_api-Submission) | repeated |  |






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
| submission_type | [string](#string) |  |  |
| errors | [Error](#listing_api-Error) | repeated |  |






<a name="listing_api-ReplaceErrorsResponse"></a>

### ReplaceErrorsResponse
ReplaceErrorsResponse is empty






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
| buyable | [bool](#bool) |  |  |






<a name="listing_api-UpdateStatusResponse"></a>

### UpdateStatusResponse
UpdateStatusResponse is empty






<a name="listing_api-UpdateSubmissionRequest"></a>

### UpdateSubmissionRequest
UpdateSubmissionRequest is used to move a submission to another status


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sku | [string](#string) |  |  |
| type | [string](#string) |  | the submission type used when creating the submission |
| content_received | [string](#string) |  |  |
| metadata | [UpdateSubmissionRequest.MetadataEntry](#listing_api-UpdateSubmissionRequest-MetadataEntry) | repeated |  |
| changed_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| status | [UpdateSubmissionRequest.Status](#listing_api-UpdateSubmissionRequest-Status) |  |  |






<a name="listing_api-UpdateSubmissionRequest-MetadataEntry"></a>

### UpdateSubmissionRequest.MetadataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 


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



<a name="listing_api-UpdateSubmissionRequest-Status"></a>

### UpdateSubmissionRequest.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 |  |
| STATUS_SUBMITTED | 1 |  |
| STATUS_ACCEPTED | 2 |  |
| STATUS_REJECTED | 3 |  |


 

 


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
| CreateSubmissions | [CreateSubmissionsRequest](#listing_api-CreateSubmissionsRequest) | [CreateSubmissionsResponse](#listing_api-CreateSubmissionsResponse) | CreateSubmissions replaces the submissions for a variant This can be used at any stage, but is most recommended for the intitial creation of a new submission. For updating submissions after creation see UpdateSubmission |
| UpdateSubmission | [UpdateSubmissionRequest](#listing_api-UpdateSubmissionRequest) | [Submission](#listing_api-Submission) | UpdateSubmission is used to move a submission to another status this can be done when its actually submitted for ingestion and/or when the ingestion is complete. |
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





 

 

 

 



<a name="api_listing_taxonomy-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/listing/taxonomy.proto



<a name="listing_api-AttributeSpec"></a>

### AttributeSpec
AttributeSpec is the specification of an attribute within a Product Type on
your channel.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attribute_spec_id | [string](#string) |  | This is the ID that will be used when retrieving listing data or submitting error to Zentail. It should reflect whatever meaningful identifier can be used to ingest this attribute into your system. In the past we have seen it be a numeric identifier or an xpath/jsonpath string |
| display_name | [string](#string) |  | Human-readable name for the attribute. Used for display purposes only. |
| type | [AttributeSpec.Type](#listing_api-AttributeSpec-Type) |  |  |
| valid_values | [string](#string) | repeated | If the type is SELECT or MULTI_SELECT, provide the valid values here |
| valid_units | [string](#string) | repeated | If the type is NUMERIC_WITH_UNITS, provide the valid units here |
| unit | [string](#string) |  | Optional, used to specify the unit of the numeric value when the type is TYPE_NUMERIC This will allow Zentail to convert values with units in Zentail to a single numeric value if your channel does not support units |
| level | [AttributeSpec.Level](#listing_api-AttributeSpec-Level) |  |  |
| classification | [AttributeSpec.Classification](#listing_api-AttributeSpec-Classification) |  |  |
| usage | [AttributeSpec.Usage](#listing_api-AttributeSpec-Usage) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="listing_api-ProductType"></a>

### ProductType
ProductType are detailed classifications dictated by the channel.
They should be at least as specific as the Category.
Generally they should be more specific than the Category.
Product Types are used to specify which attributes are required and
recommended for a given Listing.
If this is not a concept in your system, you should create a product type for
each category.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_type_id | [string](#string) |  |  |
| display_name | [string](#string) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 


<a name="listing_api-AttributeSpec-Classification"></a>

### AttributeSpec.Classification
Classification specifies whether this AttributeSpec relates to
Product, Inventory, or Pricing data

| Name | Number | Description |
| ---- | ------ | ----------- |
| CLASSIFICATION_UNSPECIFIED | 0 |  |
| CLASSIFICATION_PRODUCT | 1 |  |
| CLASSIFICATION_INVENTORY | 2 |  |
| CLASSIFICATION_PRICING | 3 |  |
| CLASSIFICATION_LOGISTICS | 4 |  |
| CLASSIFICATION_IDENTIFIER | 5 |  |



<a name="listing_api-AttributeSpec-Level"></a>

### AttributeSpec.Level
Level specifies whether this AttributeSpec should be rendered
on the Listing-level or the Variant-level

| Name | Number | Description |
| ---- | ------ | ----------- |
| LEVEL_UNSPECIFIED | 0 |  |
| LEVEL_LISTING | 1 |  |
| LEVEL_VARIANT | 2 |  |



<a name="listing_api-AttributeSpec-Type"></a>

### AttributeSpec.Type
Type specifies what form the related attribute should be rendered into

| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| TYPE_TEXT | 1 |  |
| TYPE_SELECT | 2 |  |
| TYPE_NUMERIC | 3 |  |
| TYPE_NUMERIC_WITH_UNITS | 4 |  |
| TYPE_MULTI_TEXT | 5 |  |
| TYPE_MULTI_SELECT | 6 |  |
| TYPE_MONEY | 7 |  |



<a name="listing_api-AttributeSpec-Usage"></a>

### AttributeSpec.Usage
Usage indicates if the spec is required or just recommended

| Name | Number | Description |
| ---- | ------ | ----------- |
| USAGE_UNSPECIFIED | 0 |  |
| USAGE_AVAILABLE | 1 |  |
| USAGE_REQUIRED | 2 |  |
| USAGE_RECOMMENDED | 3 |  |


 

 

 



<a name="api_listing_taxonomy_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/listing/taxonomy_service.proto



<a name="listing_api-AttributeSpecRequest"></a>

### AttributeSpecRequest
AttributeSpecRequest is used to retrieve a single Attribute Spec


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_type_id | [string](#string) |  |  |
| attribute_spec_id | [string](#string) |  |  |






<a name="listing_api-CreateAttributeSpecRequest"></a>

### CreateAttributeSpecRequest
CreateAttributeSpecRequest is used to create a new AttributeSpec


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_type_id | [string](#string) |  |  |
| attribute_spec | [AttributeSpec](#listing_api-AttributeSpec) |  |  |






<a name="listing_api-CreateProductTypeRequest"></a>

### CreateProductTypeRequest
CreateProductTypeRequest is used to create a new ProductType


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_type | [ProductType](#listing_api-ProductType) |  |  |






<a name="listing_api-CreateProductTypeResponse"></a>

### CreateProductTypeResponse
CreateProductTypeResponse returns the newly created ProductType


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_type | [ProductType](#listing_api-ProductType) |  |  |






<a name="listing_api-DeleteAttributeSpecRequest"></a>

### DeleteAttributeSpecRequest
DeleteAttributeSpecRequest is used to Delete an AttributeSpec


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_type_id | [string](#string) |  |  |
| attribute_spec_id | [string](#string) |  |  |






<a name="listing_api-DeleteAttributeSpecResponse"></a>

### DeleteAttributeSpecResponse
DeleteAttributeSpecResponse is returned after deleting an AttributeSpec






<a name="listing_api-DeleteProductTypeRequest"></a>

### DeleteProductTypeRequest
DeleteProductTypeRequest is used to delete an existing ProductType


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_type_id | [string](#string) |  |  |






<a name="listing_api-DeleteProductTypeResponse"></a>

### DeleteProductTypeResponse
DeleteProductTypeResponse is returned by the DeleteProduct type action






<a name="listing_api-ListProductTypesRequest"></a>

### ListProductTypesRequest
ListProductTypesRequest is used to send the ListProductTypes request






<a name="listing_api-ListProductTypesResponse"></a>

### ListProductTypesResponse
ListProductTypesResponse contains all of the product types returned


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_types | [ProductType](#listing_api-ProductType) | repeated |  |






<a name="listing_api-UpdateAttributeSpecRequest"></a>

### UpdateAttributeSpecRequest
UpdateAttributeSpecRequest is used to update an existing AttributeSpec


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_type_id | [string](#string) |  |  |
| attribute_spec | [AttributeSpec](#listing_api-AttributeSpec) |  |  |





 

 

 


<a name="listing_api-TaxonomyService"></a>

### TaxonomyService
TaxonomyService provides a service for managing the Taxonomy data
for a Listing integration.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListProductTypes | [ListProductTypesRequest](#listing_api-ListProductTypesRequest) | [ListProductTypesResponse](#listing_api-ListProductTypesResponse) | ListProductTypes will return all product types |
| CreateProductType | [CreateProductTypeRequest](#listing_api-CreateProductTypeRequest) | [CreateProductTypeResponse](#listing_api-CreateProductTypeResponse) | CreateProductType will create a new ProductType |
| DeleteProductType | [DeleteProductTypeRequest](#listing_api-DeleteProductTypeRequest) | [DeleteProductTypeResponse](#listing_api-DeleteProductTypeResponse) | DeleteProductType will delete the ProductType with the given ID |
| AttributeSpec | [AttributeSpecRequest](#listing_api-AttributeSpecRequest) | [AttributeSpec](#listing_api-AttributeSpec) | AttributeSpec can be used to retrieve a single AttributeSpec |
| CreateAttributeSpec | [CreateAttributeSpecRequest](#listing_api-CreateAttributeSpecRequest) | [AttributeSpec](#listing_api-AttributeSpec) | CreateAttributeSpec will create a new AttributeSpec |
| UpdateAttributeSpec | [UpdateAttributeSpecRequest](#listing_api-UpdateAttributeSpecRequest) | [AttributeSpec](#listing_api-AttributeSpec) | UpdateAttributeSpec will update the AttributeSpec with the given product_type_id and attribute_spec_id |
| DeleteAttributeSpec | [DeleteAttributeSpecRequest](#listing_api-DeleteAttributeSpecRequest) | [DeleteAttributeSpecResponse](#listing_api-DeleteAttributeSpecResponse) | DeleteAttributeSpec will delete the AttributeSpec with the given product_type_id and attribute_spec_id |

 



<a name="api_orders_orders_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/orders/orders_service.proto



<a name="orders_api-CancelItemsRequest"></a>

### CancelItemsRequest
CancelItemsRequest is the request to cancel items in an order.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| channel_order_id | [string](#string) |  | The order ID used by your channel when creating the order |
| cancel_quantities | [CancelItemsRequest.CancelQuantitiesEntry](#orders_api-CancelItemsRequest-CancelQuantitiesEntry) | repeated | A map of line item IDs to quantities to cancel |






<a name="orders_api-CancelItemsRequest-CancelQuantitiesEntry"></a>

### CancelItemsRequest.CancelQuantitiesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [int64](#int64) |  |  |






<a name="orders_api-CancelItemsResponse"></a>

### CancelItemsResponse
CancelItemsResponse is the response to a CancelItems request.





 

 

 


<a name="orders_api-OrdersService"></a>

### OrdersService
OrdersService provides a service for managing sales order data
for a Listing integration.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CancelItems | [CancelItemsRequest](#orders_api-CancelItemsRequest) | [CancelItemsResponse](#orders_api-CancelItemsResponse) | CancelItems cancels items in an order. |

 



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

