# Order Cancellations, Refunds and Returns

This guide outlines how to handle issues with orders
around cancellations, refunds and returns.

## Merchant Initiated

If an order is unfulfillable,
the Zentail seller may need to cancel the order.
This information can be obtained similarly to the shipment notifications 
by polling the 
[`GET /v1/salesOrder`](https://developer.zentail.com/docs/sales-channel-integration/b3A6MjczMjk2MjM-retrieve-detailed-sales-order-information-for-a-list-of-orders)
endpoint.
In this case, filter for the status `CANCELLED`.

We refer to this process as a `Merchant Initiated Cancellation`.

You will want to check the products array to see the `cancelQuantity` 
on each line item to confirm what part of the order is actually cancelled.

## Customer Initiated

If, for some reason, an order needs to be cancelled by the customer
use the
[`POST /salesOrder/cancel/{orderNumber}`](https://developer.zentail.com/docs/sales-channel-integration/b3A6MjczMjk2Nzg-cancel-sales-order-in-zentail)
API to communicate that to Zentail.

We refer to this process as a `Customer Initiated Cancellation`.

Please refer to the specification for details on calling that endpoint.

## Returns / Refunds

Currently the Zentail APIs do not support returns, refunds or replacements.
This means sellers will need to be able to manage these directly on your sales channel.
