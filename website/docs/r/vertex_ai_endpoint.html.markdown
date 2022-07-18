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
layout: "google"
page_title: "Google: google_vertex_ai_endpoint"
sidebar_current: "docs-google-vertex-ai-endpoint"
description: |-
  The VertexAI Endpoint resource
---

# google_vertex_ai_endpoint

The VertexAI Endpoint resource

## Example Usage - basic_endpoint
A minimal test of a vertex endpoint
```hcl
resource "google_vertex_ai_endpoint" "primary" {
  display_name = "sample-endpoint"
  location     = "us-west1"
  labels       = {}
  project      = "my-project-name"
}


```
## Example Usage - network
```hcl
resource "google_vertex_ai_endpoint" "primary" {
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  location     = "us-west1"
  labels       = {
    label-one = "value-one"
  }
  network      = google_compute_network.vertex_network.id
  project      = "my-project-name"
  depends_on   = [
    google_service_networking_connection.vertex_vpc_connection
  ]
}

resource "google_service_networking_connection" "vertex_vpc_connection" {
  network                 = google_compute_network.vertex_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.vertex_range.name]
}

resource "google_compute_global_address" "vertex_range" {
  name          = "vertex-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 24
  network       = google_compute_network.vertex_network.id
  project       = "my-project-name"
}

resource "google_compute_network" "vertex_network" {
  name       = "vertex-network"
  project    = "my-project-name"
}

```

## Argument Reference

The following arguments are supported:

* `display_name` -
  (Required)
  Required. The display name of the Endpoint. The name can be up to 128 characters long and can be consist of any UTF-8 characters.
  
* `location` -
  (Required)
  The location for the resource
  


- - -

* `description` -
  (Optional)
  The description of the Endpoint.
  
* `encryption_spec` -
  (Optional)
  Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.
  
* `labels` -
  (Optional)
  The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
  
* `network` -
  (Optional)
  The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
  
* `project` -
  (Optional)
  The project for the resource
  


The `encryption_spec` block supports:
    
* `kms_key_name` -
  (Required)
  Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. The key needs to be in the same region as where the compute resource is created.
    
## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/endpoints/{{name}}`

* `create_time` -
  Output only. Timestamp when this Endpoint was created.
  
* `deployed_models` -
  Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively.
  
* `etag` -
  Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
  
* `model_deployment_monitoring_job` -
  Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: `projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}`
  
* `name` -
  Output only. The resource name of the Endpoint.
  
* `update_time` -
  Output only. Timestamp when this Endpoint was last updated.
  
## Timeouts

This resource provides the following
[Timeouts](/docs/configuration/resources.html#timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import

Endpoint can be imported using any of these accepted formats:

```
$ terraform import google_vertex_ai_endpoint.default projects/{{project}}/locations/{{location}}/endpoints/{{name}}
$ terraform import google_vertex_ai_endpoint.default {{project}}/{{location}}/{{name}}
$ terraform import google_vertex_ai_endpoint.default {{location}}/{{name}}
```



