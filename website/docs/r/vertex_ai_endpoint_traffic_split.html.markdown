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
page_title: "Google: google_vertex_ai_endpoint_traffic_split"
sidebar_current: "docs-google-vertex-ai-endpoint-traffic-split"
description: |-
  The VertexAI EndpointTrafficSplit resource
---

# google_vertex_ai_endpoint_traffic_split

The VertexAI EndpointTrafficSplit resource

## Example Usage - basic_endpoint_traffic_split
A test of a vertex endpoint with a traffic split
```hcl
resource "google_vertex_ai_endpoint_traffic_split" "primary" {
  endpoint = google_vertex_ai_endpoint.minimal.name
  location = "us-west1"

  traffic_split {
    deployed_model_id  = google_vertex_ai_model_deployment.minimal.deployed_model_id
    traffic_percentage = 100
  }

  project = "my-project-name"
}

resource "google_vertex_ai_model_deployment" "minimal" {
  dedicated_resources {
    machine_spec {
      machine_type = "n1-standard-2"
    }

    min_replica_count = 1
  }

  endpoint = "projects/my-project-name/locations/us-west1/endpoints/${google_vertex_ai_endpoint.minimal.name}"
  model    = "projects/my-project-name/locations/us-west1/models/${google_vertex_ai_model.basic.name}"
}

resource "google_vertex_ai_model" "basic" {
  container_spec {
    image_uri = "us-docker.pkg.dev/vertex-ai/prediction/xgboost-cpu.1-5:latest"
    args      = ["sample", "args"]
    command   = ["sample", "command"]

    env {
      name  = "env_one"
      value = "value_one"
    }

    health_route = "/health"

    ports {
      container_port = 8080
    }

    predict_route = "/predict"
  }

  display_name = "sample-model"
  location     = "us-west1"
  artifact_uri = "gs://cloud-samples-data/vertex-ai/google-cloud-aiplatform-ci-artifacts/models/iris_xgboost/"
  description  = "A sample model"

  labels = {
    label-one = "value-one"
  }

  project             = "my-project-name"
  version_aliases     = ["default", "v1", "v2"]
  version_description = "A sample model version"
}

resource "google_vertex_ai_endpoint" "minimal" {
  display_name = "sample-endpoint"
  location     = "us-west1"
  labels       = {}
  project      = "my-project-name"
}


```

## Argument Reference

The following arguments are supported:

* `endpoint` -
  (Required)
  The endpoint for the resource
  
* `location` -
  (Required)
  The location for the resource
  
* `traffic_split` -
  (Required)
  A map from a DeployedModel's ID to the percentage of this Endpoint's traffic that should be forwarded to that DeployedModel. If a DeployedModel's ID is not listed in this map, then it receives no traffic. The traffic percentage values must add up to 100, or map must be empty if the Endpoint is to not accept any traffic at a moment.
  


The `traffic_split` block supports:
    
* `deployed_model_id` -
  (Required)
  A deployed model's id.
    
* `traffic_percentage` -
  (Required)
  The percentage of this Endpoint's traffic that should be forwarded to the DeployedModel.
    
- - -

* `project` -
  (Optional)
  The project for the resource
  


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}`

* `etag` -
  Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
  
## Timeouts

This resource provides the following
[Timeouts](/docs/configuration/resources.html#timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import

EndpointTrafficSplit can be imported using any of these accepted formats:

```
$ terraform import google_vertex_ai_endpoint_traffic_split.default projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}
$ terraform import google_vertex_ai_endpoint_traffic_split.default {{project}}/{{location}}/{{endpoint}}
$ terraform import google_vertex_ai_endpoint_traffic_split.default {{location}}/{{endpoint}}
```



