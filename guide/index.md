# Zentail API Sales Channel Integration Guide

This guide aims to outline all of the steps necessary to
build a complete Sales Channel integration into Zentail.

Our API is designed to abstract away as much of the complexity
around managing listing data, inventory and sales orders from
integration partners and provide a simple to use and understand
interface.

See below for the major steps of building an integration:

1. [Application Registration](./application-registration.md)
1. [Authentication](./authentication.md)
1. [Taxonomy Ingestion](./taxonomy-ingestion.md)
1. [Inventory Retrieval](./inventory.md)
1. [Pricing Data Retrieval](./pricing.md)
1. [Listing Data Retrieval](./listing-data.md)
1. [Listing Status / Errors](./listing-status.md)
1. [Sales Order Injection](./sales-order-injection.md)
1. [Order Fulfillment Retrieval](./order-fulfillment-retrieval.md)
1. [Order Cancellation, Returns and Refunds](./order-cancellation-refund-returns.md)

The steps are ordered in such a way as to provide an increasingly complete integration
that can be tested along the way as the pieces are built out.

> [!NOTE]
> Some of the process of Taxonomy Ingestion requires implementation
> work from Zentail. This means it is be valuable to start sending us the 
> Taxonomy information ASAP to avoid delays later in the process.


