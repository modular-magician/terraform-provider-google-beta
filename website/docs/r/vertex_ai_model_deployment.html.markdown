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
page_title: "Google: google_vertex_ai_model_deployment"
description: |-
  The VertexAI ModelDeployment resource
---

# google_vertex_ai_model_deployment

The VertexAI ModelDeployment resource

## Example Usage - basic_model_deployment
A basic test of a vertex model deployment
```hcl
resource "google_vertex_ai_model_deployment" "primary" {
  dedicated_resources {
    machine_spec {
      machine_type = "n1-standard-2"
    }

    min_replica_count = 1
    max_replica_count = 1
  }

  endpoint = "projects/my-project-name/locations/us-west1/endpoints/${google_vertex_ai_endpoint.minimal.name}"
  model    = "projects/my-project-name/locations/us-west1/models/${google_vertex_ai_model.basic.name}"
  location = "us-west1"
  project  = "my-project-name"
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

* `dedicated_resources` -
  (Required)
  A description of resources that are dedicated to the DeployedModel, and that need a higher degree of manual configuration.
  
* `endpoint` -
  (Required)
  The name of the endpoint to deploy to
  
* `model` -
  (Required)
  The name of the model to deploy
  


The `dedicated_resources` block supports:
    
* `machine_spec` -
  (Required)
  Required. Immutable. The specification of a single machine used by the prediction.
    
* `max_replica_count` -
  (Optional)
  Immutable. The maximum number of replicas this DeployedModel may be deployed on when the traffic against it increases. If the requested value is too large, the deployment will error, but if deployment succeeds then the ability to scale the model to that many replicas is guaranteed (barring service outages). If traffic against the DeployedModel increases beyond what its replicas at maximum may handle, a portion of the traffic will be dropped. If this value is not provided, will use min_replica_count as the default value. The value of this field impacts the charge against Vertex CPU and GPU quotas. Specifically, you will be charged for max_replica_count * number of cores in the selected machine type) and (max_replica_count * number of GPUs per replica in the selected machine type).
    
* `min_replica_count` -
  (Required)
  Required. Immutable. The minimum number of machine replicas this DeployedModel will be always deployed on. This value must be greater than or equal to 1. If traffic against the DeployedModel increases, it may dynamically be deployed onto more replicas, and as traffic decreases, some of these extra replicas may be freed.
    
The `machine_spec` block supports:
    
* `machine_type` -
  (Required)
  Immutable. The type of the machine. See the [list of machine types supported for prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types) See the [list of machine types supported for custom training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types). For DeployedModel this field is optional, and the default value is `n1-standard-2`. For BatchPredictionJob or as part of WorkerPoolSpec this field is required. TODO(rsurowka): Try to better unify the required vs optional.
    
- - -

* `location` -
  (Optional)
  The location of the endpoint
  
* `project` -
  (Optional)
  The project of the endpoint
  


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}/models/{{model}}`

* `deployed_model_id` -
  The deployed ID of the model in the endpoint
  
## Timeouts

This resource provides the following
[Timeouts](/docs/configuration/resources.html#timeouts) configuration options:

- `create` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import

ModelDeployment can be imported using any of these accepted formats:

```
$ terraform import google_vertex_ai_model_deployment.default projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}/models/{{model}}
$ terraform import google_vertex_ai_model_deployment.default {{project}}/{{location}}/{{endpoint}}/{{model}}
$ terraform import google_vertex_ai_model_deployment.default {{location}}/{{endpoint}}/{{model}}
```



