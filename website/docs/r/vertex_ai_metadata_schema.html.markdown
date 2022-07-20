---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
#
# ----------------------------------------------------------------------------
#
#     This file is managed by Magic Modules (https:#github.com/GoogleCloudPlatform/magic-modules)
#     and is based on the DCL (https:#github.com/GoogleCloudPlatform/declarative-resource-client-library).
#     Changes will need to be made to the DCL or Magic Modules instead of here.
#
#     We are not currently able to accept contributions to this file. If changes
#     are required, please file an issue at https:#github.com/hashicorp/terraform-provider-google/issues/new/choose
#
# ----------------------------------------------------------------------------
subcategory: "VertexAi"
page_title: "Google: google_vertex_ai_metadata_schema"
description: |-
  The VertexAI MetadataSchema resource
---

# google_vertex_ai_metadata_schema

The VertexAI MetadataSchema resource

## Example Usage - basic
```hcl
resource "google_vertex_ai_metadata_schema" "primary" {
  location       = "us-west1"
  metadata_store = google_vertex_ai_metadata_store.minimal.name
  name           = "schema"
  schema         = "title: sample.Schema\ntype: object"
  schema_type    = "ARTIFACT_TYPE"
  schema_version = "1.0.0"
  project        = "my-project-name"
}

resource "google_vertex_ai_metadata_store" "minimal" {
  region  = "us-west1"
  name    = "store"
  project = "my-project-name"
}

```

## Argument Reference

The following arguments are supported:

* `location` -
  (Required)
  The location for the resource
  
* `metadata_store` -
  (Required)
  The metadata store for the resource
  
* `name` -
  (Required)
  Output only. The resource name of the MetadataSchema.
  
* `schema` -
  (Required)
  Required. The raw YAML string representation of the MetadataSchema. The combination of [MetadataSchema.version] and the schema name given by `title` in [MetadataSchema.schema] must be unique within a MetadataStore. The schema is defined as an OpenAPI 3.0.2 [MetadataSchema Object](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#schemaObject)
  
* `schema_type` -
  (Required)
  The type of the MetadataSchema. This is a property that identifies which metadata types will use the MetadataSchema. Possible values: METADATA_SCHEMA_TYPE_UNSPECIFIED, ARTIFACT_TYPE, EXECUTION_TYPE, CONTEXT_TYPE
  
* `schema_version` -
  (Required)
  The version of the MetadataSchema. The version's format must match the following regular expression: `^[0-9]+.+.+$`, which would allow to order/compare different versions. Example: 1.0.0, 1.0.1, etc.
  


- - -

* `project` -
  (Optional)
  The project for the resource
  


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/metadataStores/{{metadata_store}}/metadataSchemas/{{name}}`

* `create_time` -
  Output only. Timestamp when this MetadataSchema was created.
  
## Timeouts

This resource provides the following
[Timeouts](/docs/configuration/resources.html#timeouts) configuration options:

- `create` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import

MetadataSchema can be imported using any of these accepted formats:

```
$ terraform import google_vertex_ai_metadata_schema.default projects/{{project}}/locations/{{location}}/metadataStores/{{metadata_store}}/metadataSchemas/{{name}}
$ terraform import google_vertex_ai_metadata_schema.default {{project}}/{{location}}/{{metadata_store}}/{{name}}
$ terraform import google_vertex_ai_metadata_schema.default {{location}}/{{metadata_store}}/{{name}}
```



