# Pricing

The pricing ingestion process should work basically the same
as the [inventory ingestion process](./inventory.md).

Instead of the inventory endpoint, you can use:

```
GET /storefront/v2/listing/variant/pricing/updated
```

with the following parameters:

| Parameter | Type | Description | Example |
| --------- | ---- | ----------- | ------- |
| `since`   | `string ($date-time)` | All variants with pricing changes after this timestamp will be returned | `2024-10-22T18:01:15.094Z` |
| `cursor`  | `string` | Use this if a nextPageCursor is returned to retrieve the next page of results | |

## Response Data

The `pricing` object contains a few structured fields that should be considered

```jsonc
{
  "variants": [
     {
       ...
       "pricing": {
         "enabled": true,
         "attributes": [
           {
             "id": "string",
             ... value parameters ...
           }
           ...
        },
        "updatedAt": "2024-10-22T18:01:15.094Z",
        "sales": [
          {
            "strikeThroughPrice": 0,
            "salePrice": 0,
            "startAt": "2024-10-22T18:01:15.094Z",
            "endAt": "2024-10-22T18:01:15.094Z"
          }
        ]
    },
    ...
  ]
}
```

### Attributes

The `attributes` array provides the attributes from your taxonomy
that you have classified as pricing attributes.

The `id` attribute will correspond to the `channel_attribute_id` provided
during taxonomy ingestion. Alongside that will be a value key which depends
on the type of the attribute.
It should be one of (please refer to the spec for the most up to date possibilities):

```
"textValue": "string",
"numericValue": 0,
"numericWithUnitsValue": {
  "numeric": 0,
  "units": "string"
},
"multiTextValue": {
  "values": [
    "string"
  ]
},
"moneyValue": {
  "amount": "string",
  "currencyCode": "string"
}
```

for example, if you created a `purchasable.price` attribute in your taxonomy
with a money type:

```json
{
  "id": "purchasable.price",
  "moneyValue": {
    "amount": 12.34,
    "currencyCode": "USD"
  }
}
```

### Scheduled Sales

Zentail sellers are able to create scheduled sales in Zentail
that represent short-term price reductions.

Currently planned sales will be included in the `sales` array:

```json
[
    {
      "strikeThroughPrice": 12.34,
      "salePrice": 10.00,
      "startAt": "2024-10-22T18:01:15.094Z",
      "endAt": "2024-10-22T18:01:15.094Z"
    }
]
```

If the current time is between the `startAt` and `endAt`, you should reflect
these sales on your channel.

> [!NOTE]
> You should receive a notification in your polling process at the start
> and end of the sale so you do not _need_ to keep track of the end date.

### Recommended Implementation Strategy

You can use the above API to routinely request newly updated variants.
We recommend that you maintain an internal timestamp for each seller to 
keep track of the most recently successful request.

You can then run, for each seller, a process that requests updates since
that time, process them and update the timestamp to when you began that process.

## Next Steps

The process for [ingesting Listing Data](./listing-dat.md) 
is the final step in the complete listing ingestion process.
