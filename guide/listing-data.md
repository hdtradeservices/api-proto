# Listing Data

The process for ingesting Listing Data is very similar to 
[pricing](./pricing.md) and [inventory](./inventory.md).

The difference here is that there are a few different types
of events that can happen for listings:

1. Listings that Zentail believes should be created on your channel
1. Listings that have had product data changes

## "New" Listings

To find out about listings Zentail thinks should be created on your channel
you can use this request:

```
GET /storefront/v2/listing/new
```

with the following parameters:

| Parameter | Type | Description | Example |
| --------- | ---- | ----------- | ------- |
| `since`   | `string ($date-time)` | All listings updated or created after this time will be returned | `2024-10-22T18:01:15.094Z` |
| `cursor`  | `string` | Use this if a nextPageCursor is returned to retrieve the next page of results | |

Zentail thinks all listings are "new" until you send us back a channel ID
(see [listing status](./listing-status.md)).

It is still advisable to check the data in the listing against
existing sku data in your system to ensure that duplicate listings
are not created.

> [!IMPORTANT] 
> Zentail *HIGHLY* recommends you enforce SKUs to be unique on your channel.
> This is congruent with how things work in Zentail and will massively simplify
> the integration and ensure listings are unique and accurate.

## Updated Listings

For updated listings, use this endpoint:

```
GET /storefront/v2/listing/product_data/updated
```

with the following parameters:

| Parameter | Type | Description | Example |
| --------- | ---- | ----------- | ------- |
| `since`   | `string ($date-time)` | All listings with listing data updated after this time will be returned | `2024-10-22T18:01:15.094Z` |
| `cursor`  | `string` | Use this if a nextPageCursor is returned to retrieve the next page of results | |

## Ingesting Listings

The listing data returned will include all of the information needed
to construct a listing on your channel.
The attributes will be provided using the `channel_attribute_id` you
provided during [taxonomy ingestion](./taxonomy-ingestion.md) and values
in the corresponding type.

The response will now include some extra contextual information
not included in the `Variant` responses:

```jsonc
{
  "listings": [
    {
      "id": "a unique zentail ID for this listing",
      "productData": {
        "enabled": true,
        "productTypeId": "string",
        "pivotAttributes": [
           "size", "color"
        ],
        "attributes": [
          ...
        ]
      },
      "variants": [
        {
          "sku": "string",
          "channelId": "string",
          "inventory": {
             ...
          },
          "identifiers": {
            "enabled": true,
            "attributes": [
              ...
            ],
            "updatedAt": "2024-10-22T18:23:49.457Z"
          },
          "productData": {
            "enabled": true,
            "attributes": [
              ...
            ],
            "updatedAt": "2024-10-22T18:23:49.457Z"
          },
          "pricing": {
            ...
          },
          "logistics": {
            "enabled": true,
            "attributes": [
              ...
            ],
            "updatedAt": "2024-10-22T18:23:49.457Z"
          }
        }
      ],
      "createdAt": "2024-10-22T18:23:49.457Z",
      "updatedAt": "2024-10-22T18:23:49.457Z"
    }
  ],
  "nextPageCursor": "string"
}
```

There is a top-level `productData` object which will contain
any attributes that were marked as `Listing`-level attributes
during taxonomy ingestion.

The `pivotAttributes` property dictates which Attributes the listing should
vary on. For example, if its `["size", "color"]` then the listing should
provide users the ability to select a size and color based on the variants in the listing.

> [!NOTE] Zentail sellers have the ability to restrict sending
> certain types of data to sales channels.
> This is the reason for the `enabled` flag present in each block.
> If the seller has disabled that particular classification,
> `enabled` with be false and no `attributes` will be provided.

In the variants, you can see `inventory` and `pricing` which have been
covered in the previous steps.
We also have `identifiers`, `logistics` and `productData`, these
will provide attributes with the corresponding classifications based on
what you provided us during taxonomy ingestion.

> [!IMPORTANT]
> Inventory and Price changes _will not_ cause listings to be returned by this endpoint.
> They are provided here for completeness and convenience but to receive updates when
> they change, see the [inventory](./inventory.md) and [pricing](./pricing.md) flows.

In all locations with `attributes` arrays, the attributes will have an `id`
which corresponds to a `channel_attribute_id` provided during taxonomy ingestion.

They will also have a value that is based on the type of the attribute.

Refer to the spec for the value types and structures.

## Recommended Implementation Strategy

You can use the above APIs to routinely request newly created and updated listings.
We recommend that you maintain an internal timestamp for each seller to 
keep track of the most recently successful request.

You can then run, for each seller, a process that requests updates since
that time, process them and update the timestamp to when you began that process.

When ingesting the listing, it likely makes sense to ingest all of the provided data
to ensure it is as up to date as possible.

## Next Steps

Updating Zentail with the status and any ingestion errors is critical
to ensuring sellers can quickly identify and resolve any data quality
issues as quickly as possible.

See the [next guide](./listing-status.md) to start sending statuses back to Zentail.
