# Inventory

Zentail's API aims to make inventory ingestion as simple as possible.
Having the most up to date and correct inventory levels in your system
is essential to avoiding overselling.

With that in mind, you should aim to perform the below actions as frequently
as possible.

## Polling for Inventory Changes

To retrieve a list of all SKUs (variants) with inventory changes,
you can make a request to:

```
GET /storefront/v2/listing/variant/inventory/updated
```

with the following parameters:

| Parameter | Type | Description | Example |
| --------- | ---- | ----------- | ------- |
| `since`   | `string($date-time)` | A timestamp value which indicates the earliest changes that should be reflected | `2024-10-22T17:10:37.067Z` |
| `externalChangesOnly` | `bool` | If true, only return variants with inventory changes that have happened outside of sales on this channel. Defaults to `false` | `false` |
| `cursor`  | `string` | Use this if a `nextPageCursor` is returned to retrieve the next page of results |  | 

This API will return all variants that have had an inventory change since the provided time.

> [!NOTE]
> `externalChangesOnly` should only be used if you are upkeeping your own
> internal inventory account for orders that happen directly on your channel and you want
> to receive less updates and try to avoid race conditions with changes in Zentail.
> Setting the flag will mean that if inventory has only changed on your sales channel
> those results will be omitted from the response.

## Handling the Response

The response to the above request will include all of the information about variants with
inventory changes. For the sake of ingesting inventory, we can focus in on the `inventory` object.

```json
{
  "variants": [
    {
      "sku": "string",
      ...
      "inventory": {
        "enabled": true,
        "totalQuantity": 100,
        "merchantFulfillableQuantity": 100,
        "storefrontFulfillableQuantity": 0,
        "updatedAt": "2024-10-22T17:10:37.067Z",
        "updatedExternallyAt": "2024-10-22T17:10:37.067Z"
      },
      ...
    }
  ]
}
```

For most use cases, the `totalQuantity` is all we really need to leverage.
That represents the currently available quantity for sale after accounting
for pending orders on all of the sellers channels.

`updatedAt` and `updatedExternallyAt` can also be useful for understanding exactly
when the changes to the inventory levels occurred.

### Recommended Implementation Strategy

You can use the above API to routinely request newly updated variants.
We recommend that you maintain an internal timestamp for each seller to 
keep track of the most recently successful request.

You can then run, for each seller, a process that requests updates since
that time, process them and update the timestamp to when you began that process.

## Next Steps

The process for [ingesting Pricing](./pricing.md) 
changes is very similar, so we recommend tackling that next.
