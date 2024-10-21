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
| **Attribute Spec**         | AttributeSpec is the specification of an Attribute within a Category on your channel. |
| **Category**               | Categories provide a (usually) high-level classification for a Listing. Categories can be arranged in a tree-like structure with parent and child categories. The category will often be a crucial piece of information about how the data is structured when sending to the channel. For example, on Amazon it dictates which path to take down the XSD when constructing feed messages. Attribute Specs are associated directly with Categories. |
| **Product Type**           | ProductType are detailed classifications dictated by the channel. They should be at least as specific as the Category. Generally they should be more specific than the Category. Product Types are used to specify which attributes are required and recommended for a given Listing. They are not relevant to the structure of the data that we are capable of sending. If this is not a concept in your system, you should create a product type for each category. |
| **Quality Attribute Spec** | QualityAttributeSpec provides a connection between a ProductType and an AttributeSpec. If a Quality Attribute Spec exists for an AttributeSpec in a particular ProductType, that Spec is either Recommended or Required in that ProductType (and the Quality Attribute Spec will indicate which it is). |

## Creating Categories

You should construct the Category tree first.
All of the other objects will depend in one way or another on a Category.

Categories can be created using the `CreateCategory` action:

```
POST /v2/storefront/taxonomy/category
```

With an example request like:

```json
{
  "category": {
    "channelCategoryId": "<an identifier for your channel, must be unique>",
    "parentCategoryId": "<the ID of the parent category, can be omitted or set to an empty string to indicate root level>",
    "displayName": "<a user friendly display name>"
  }
}
```

You can check
[the definitions](../src/api/listing/taxonomy_service.proto)
for details on the other CRUD operations available.

## Attribute Specs

Attribute Specifications (Specs) can now be created since we have constructed
a Category tree to associate them with.
To create them, use the CreateAttributeSpec action:

```
POST /v2/storefront/taxonomy/attribute_spec/{categoryId}
```

with a body like:

```jsonc
{
  "attributeSpec": {
    "attributeSpecId": "<a unique identifier for this attribute, must only be unique within the category ID>",
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
    "level": "<one of LEVEL_LISTING, LEVEL_VARIANT>",
    "classification": "<one of CLASSIFICATION_PRODUCT, CLASSIFICATION_INVENTORY, CLASSIFICATION_PRICING, CLASSIFICATION_LOGISTICS, CLASSIFICATION_IDENTIFIER>",
  }
}
```

> [!NOTE]
> Attribute Specs are always associated with a category,
> if a certain Attribute is available in multiple categories,
> you should invoke the `CreateAttributeSpec` action for each category ID

## Product Types

At this point, we have established the bear minimum for generating a listing
from Zentail. The next 2 sections outline tools you can use to provide greater
visibility into what makes a great listing on your channel. `ProductType`s and 
`QualityAttributeSpec` drive Listing Quality workflows in Zentail that help sellers
understand requirements and recommendations for each Listing.

If your system does not differentiate `ProductTypes` from `Categories`, you should create
a single Product Type for each Category.

To create ProductTypes, use the `CreateProductType` action:

```
POST /v2/storefront/taxonomy/product_type
```

with a request like:

```json
{
  "productType": {
    "categoryId": "<the category ID to associate the product type with>",
    "channelProductTypeId": "<your id for this product type>"
  }
}
```

## Quality Attribute Specs

`Quality Attribute Specs` (`QAS`) are used to inform Zentail sellers
the use of an AttributeSpec within a Product Type.
They can indicate whether the Spec is _required_ or highly _recommended_
in a Product Type.

To create them, use the `CreateQualityAttributeSpec` action:

```
POST /v2/storefront/taxonomy/quality_attribute_spec/{productTypeId}
```

with a request like:

```json
{
  "qualityAttributeSpec": {
    "attributeSpecId": "<the attribute spec id to indicate>",
    "usage": "<one of USAGE_REQUIRED, USAGE_RECOMMENDED>"
  }
}
```
