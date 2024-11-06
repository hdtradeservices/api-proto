# Notifying Sellers of Listing Statuses and Errors

An important part of empowering sellers to get their listings live
is providing actionable feedback on the status of their listings 
on your channel. To relay this information, Zentail provides a family of APIs
that enable the communication of updates, statuses and errors that are
part of your ingestion process.

## Storefront Status

The most critical, functional piece is relaying the status of the listing
back to Zentail. This allows sellers to understand whether or not the 
listings are live and published on the channel.
It also helps Zentail understand the state of the listing so we know if 
it needs to be created on the channel or updated.

To update status, use the `UpdateStatus` action:

```
POST /v2/storefront/listing/variant/{sku}/status
```

with the following in the JSON body:

| Parameter | Type | Description |
| --------- | ---- | ----------- |
| `status`  | `StandardStatus` | A standardized status that enables Zentail to take particular actions. See the models for the most up to date set of values |
| `channelStatus` | `string` | A free-form status that reflects the status of the listing in your terms |
| `channelId` | `string` | A unique identifier for this listing on your platform |
| `buyable` | `boolean` | Set to true if the listing is available for purchase. This flag indicates to Zentail sellers that the product is available and sales should be expected. |
| `listingURL` | `string` | (optional) A URL the seller can click to view the listing on your platform |

When making this call, it will replace the current status for the given SKU.
For listings with multiple skus, a call should be made for each one since each
variant may have different statuses, ids and view links.

## Errors

If there are any problems ingesting the listing, particularly data quality issues surrounding
specific attributes. Sending the errors into Zentail can enable powerful workflows for sellers 
to quickly fix them, improving their listing quality and getting more listings live on the channel.

Errors are managed on a SKU level, but also a `submissionType` level.
`submissionType` is useful if you have multiple different processes running separately
that can generate errors (i.e. `inventory`, `pricing` and `listing data`).
Calls to the `ReplaceErrors` action will replace all errors for that SKU and `submissionType`.

> [!IMPORTANT]
> Once errors are created and then resolved, you must call `ReplaceErrors` again, providing no errors
> so Zentail knows it can clear the existing errors out.

To invoke the `ReplaceErrors` action, you can call this endpoint:

```
POST /v2/storefront/listing/variant/{sku}/errors
```

with the following in the JSON body:

| Parameter | Type | Description |
| --------- | ---- | ----------- |
| `submissionType` | `string` | The unique submission type to associate these errors with |
| `errors` | `[]Error` | An array of the errors, structured as outlined below |

Error fields:

| Parameter | Type | Description |
| --------- | ---- | ----------- |
| `attributeIds` | `[]string` | A list of attribute that are linked to this error. This is used to drive the users attention to specific attributes they can use to resolve the error. The IDs should correspond to `channel_attribute_ids` provided during [taxonomy ingestion](./taxonomy-ingestion.md) |
| `severity` | `Severity` | An indicator of how urgent it is the user resolve this error. `Error` should be used in the event this error prevents the listing from being ingested. See the models for the complete list of valid values. |
| `type` | `ErrorType` | The error type can be used to create a consistent experience in Zentail. It can indicate the classification of error and provide a standardized error message and UI experience for resolving the error. |
| `message` | `string` | An error message to show the user to help them understand and resolve the issue. Ideally these would be as user-friendly as possible. Only required if `Other` is used as the type. |

## Submissions

It's very comforting for Zentail users to see the progress of their listings on the channel.
To provide more insight, we make the `Submissions` API available.
This allows you to convey to Zentail and its users, that you have recieved an update and
started processing it in your system.
It can be very helpful for debugging purposes to provide meaningful data to this API.

To update submissions, you can use the `ReplaceSubmissions` action via this call:

```
POST /v2/storefront/listing/submission
```

Submissions are again grouped by SKU and `submissionType` (similar to errors) and this action
will replace previous submissions that match the SKU and `submissionType`.

Submission Fields:

| Parameter | Type | Description |
| --------- | ---- | ----------- |
| `submissionID` | `string` | A unique identifier in your system to identify the submission, does not need to be unique per SKU or submissionType |
| `sku`          | `string` | The SKU that was submitted | 
| `contentSent`  | `string` | The content that was ingested |
| `contentReceived` | `string` | The result of the content ingestion |
| `metadata` | `map[string]string` | Any additional metadata that is helpful to reveal to the seller |
| `type` | `string` | The `submissionType` discussed above |
| `successful` | `bool` | Whether or not the ingestion was successful |
| `preparedAt` | `string ($date-time)` | The time the data was prepared for ingestion |
| `submittedAt` | `string ($date-time)` | The time the ingestion started |
| `acknowledgedAt` | `string ($date-time)` | The time the ingestion was complete |
