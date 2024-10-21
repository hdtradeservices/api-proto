# Sales Order Injection

For this step, we will implement the process for sending orders that are placed by your customers.
In Zentail terminology, these are referred to as Sales Orders.

In the sales channel system,
you should identify which seller should fulfill the order,
and use their corresponding Zentail API Token to make all calls below.

For this, we will primarily invoke the
[`POST /v1/salesOrder`](https://developer.zentail.com/docs/sales-channel-integration/b3A6MjczMjk2MjQ-create-new-sales-orders-in-zentail)
endpoint.

The documentation there should provide all of the necessary info to make a successful call, but there are some special callouts below:

## Integration ID
Since we are using a registered application, `integrationID` will not be required (in fact, it'll be ignored).

## Order Status / Payment

If you are sending an order to Zentail for fulfillment,
you should use either the `PENDING` or `PENDING_PAYMENT` status.

For orders that have not yet cleared payment verification,
you can send the `PENDING_PAYMENT` status.
In that status, no fulfillment action will be taken,
but Zentail will reserve the inventory to fulfill this order for a short period of time.

If you used the `PENDING_PAYMENT` status,
you will need to call the
[`POST /v1/salesOrder/markPaid/{orderNumber}`](https://developer.zentail.com/docs/sales-channel-integration/b3A6MjczMjk2Nzk-mark-order-paid-in-zentail)
endpoint for that order.
Doing so will mark the order `PENDING` and from there,
we will advance the order to the fulfillment process.

## Commission

If it is relevant (and possible),
you can provide the channel commission charged for this order in the resellerCommission field.
This can be very helpful for your sellers
so they can have a better picture of their overall profitability on your channel.
