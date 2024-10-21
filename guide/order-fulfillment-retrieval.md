# Retrieving Order Fulfillments

Once you have sent orders into Zentail for fulfillment,
the next step will be to wait until it's shipped and then retrieve the shipment details.

Currently, there are 2 ways to do this which are each explained in detail below:

* Polling the
[`GET /v1/salesOrder`](https://developer.zentail.com/docs/sales-channel-integration/b3A6MjczMjk2MjM-retrieve-detailed-sales-order-information-for-a-list-of-orders)
endpoint,
with the lastUpdatedTs and status filter (recommended)
* Polling the
[`GET /v1/salesOrder/{orderId}`](https://developer.zentail.com/docs/sales-channel-integration/b3A6MjczMjk2MjU-get-detailed-sales-order-information)
endpoint.

In either case, you will be looking for a few key things:

* The order status is `SHIPPED` (or `PARTIALLY_SHIPPED`, more on that below).
* The Package objects populated under the packages key on the order.
* Each package contains basic package level information like carrier,
tracking and shipment timestamp.
There is also the products key which can be used to see which SKUs (and how many) are in each package.

## Polling with lastUpdatedTs and status (recommended)

For this approach,
we will make a request to list all orders in a given status
that have been updated on or after the given time.
This is particularly useful since it allows us
to target a specific set of orders
and not need to rely on checking orders
we don't care about 
or checking each order individually.

We recommend keeping a timestamp internal to each token
indicating the last timestamp you received a response
and querying our API with a time that slightly overlaps that time
(maybe one minute before it).
Then when you successfully proccess the results of that call,
advance your internal timestamp to the time you made the initial request.

You can also add the status filter if you are particularly interested in orders that have shipped.
For this, provide the value `SHIPPED` to the status param in each request.

> [!NOTE]
> Since this order data is so critical,
> we recommend polling at least once every 15 minutes and as fast as once every 2 minutes.

## Polling individual orders

You may also use the individual order
[`GET /v1/salesOrder/{orderId}`](https://developer.zentail.com/docs/sales-channel-integration/b3A6MjczMjk2MjU-get-detailed-sales-order-information)
endpoint to check on each unshipped order one at a time.
Procedurally, this is simpler
but if you are sending a large volume of orders for a particular seller,
this can be a bit unweildy and throttling can start to become an issue.

## A Note on Partial Shipments

Zentail supports order in a partially shipped state.
This indicates that some but not all of the units in the order have been shipped.

You can still process orders in this state,
but note the contents of the packages array may change as the remaining items are shipped.

You can identify a partially shipped order by the status `PARTIALLY_SHIPPED`.
Once all of the items are accounted for in shipments, the status will advance to `SHIPPED`.
