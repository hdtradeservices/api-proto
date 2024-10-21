# Taxonomy Ingestion

Taxonomy ingestion is a critical part of how Zentail delivers a simple
pre-transcribed listing via the API.

This guide will outline how to send us the right structured metadata
so we can construct high-quality listings that are ready-to-ingest by your system.

## Terminology

This section will outline some key terms that we use in the API and what their 
role in the integration is.

| Term                       | Meaning |
| -------------------------- | ------- |
| **Attribute**    | An Attribute is the specification (name, type, product type etc...) of an Attribute within a ProductType on your marketplace. |
| **Product Type** | ProductType are detailed classifications dictated by the marketplace. Product Types are used to specify which attributes are available, required and recommended for Listing that are classified into them. |

## Product Types

The first step in building a Taxonomy is to populate your ProductTypes.
To create ProductTypes, use the `CreateProductType` action:

```
POST /v2/storefront/taxonomy/product_type
```

with a request like:

```json
{
  "productType": {
    "productTypeId": "<a unique identifier for this product type>",
    "displayName": "< a user-friendly name for this product type>"
  }
}
```

## Attribute

Attributes can now be created since we have constructed
the ProductTypes to associate them with.
To create them, use the CreateAttribute action:

```
POST /v2/storefront/taxonomy/attribute/{productTypeId}
```

with a body like:

```jsonc
{
  "attribute": {
    "attributeId": "<a unique identifier for this attribute, must only be unique within the productType ID>",
    "displayName": "<a user friendly display name>",
    "type": "<one of: TYPE_TEXT, TYPE_SELECT, TYPE_NUMERIC, TYPE_NUMERIC_WITH_UNITS, TYPE_MULTI_TEXT, TYPE_MULTI_SELECT, TYPE_MONEY>",
    // only include validValues if the type is SELECT or MULTI_SELECT
    "validValues": [
      "string"
    ],
    // only include validUnits if the type is NUMERIC_WITH_UNITS
    "validUnits": [
      "string"
    ],
    // unit is used only by NUMERIC types. If this numeric value has some 
    // intrinsic unit associated with it, provide it here and Zentail will
    // ensure values in Zentail with differing units are converted properly
    // before including them on the listing.
    "unit": "string",
    "can_be_pivot": true,
    "level": "<one of LEVEL_LISTING, LEVEL_VARIANT>",
    "usage": "<one of USAGE_AVAILABLE, USAGE_RECOMMENDED, USAGE_REQUIRED>",
    "classification": "<one of CLASSIFICATION_PRODUCT, CLASSIFICATION_INVENTORY, CLASSIFICATION_PRICING, CLASSIFICATION_LOGISTICS, CLASSIFICATION_IDENTIFIER>",
  }
}
```

### Listing Level

The `level` property shouldo only used to enforce that all variants must always have the same value
for this attribute spec.
In general, it is safe to use `LEVEL_VARIANT` for all Attributes,
but `LEVEL_LISTING` is available to allow additional constraints on the valid configuration
of each Attribute.

> [!NOTE]
> Attributes are always associated with a Product Type,
> if a certain Attribute is available in multiple Product Types,
> you should invoke the `CreateAttribute` action for each `productTypeID`.

