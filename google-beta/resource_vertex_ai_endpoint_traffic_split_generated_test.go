// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"fmt"
	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	vertexai "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai/beta"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

func TestAccVertexAiEndpointTrafficSplit_BasicEndpointTrafficSplit(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  getTestProjectFromEnv(),
		"region":        getTestRegionFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVertexAiEndpointTrafficSplitDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAiEndpointTrafficSplit_BasicEndpointTrafficSplit(context),
			},
			{
				ResourceName:            "google_vertex_ai_endpoint_traffic_split.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func testAccVertexAiEndpointTrafficSplit_BasicEndpointTrafficSplit(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vertex_ai_endpoint_traffic_split" "primary" {
  endpoint = google_vertex_ai_endpoint.minimal.name
  location = "%{region}"

  traffic_split {
    deployed_model_id  = google_vertex_ai_model_deployment.minimal.deployed_model_id
    traffic_percentage = 100
  }

  project = "%{project_name}"
}

resource "google_vertex_ai_model_deployment" "minimal" {
  dedicated_resources {
    machine_spec {
      machine_type = "n1-standard-2"
    }

    min_replica_count = 1
  }

  endpoint = "projects/%{project_name}/locations/%{region}/endpoints/${google_vertex_ai_endpoint.minimal.name}"
  model    = "projects/%{project_name}/locations/%{region}/models/${google_vertex_ai_model.basic.name}"
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
  location     = "%{region}"
  artifact_uri = "gs://cloud-samples-data/vertex-ai/google-cloud-aiplatform-ci-artifacts/models/iris_xgboost/"
  description  = "A sample model"

  labels = {
    label-one = "value-one"
  }

  project             = "%{project_name}"
  version_aliases     = ["default", "v1", "v2"]
  version_description = "A sample model version"
}

resource "google_vertex_ai_endpoint" "minimal" {
  display_name = "sample-endpoint"
  location     = "%{region}"
  labels       = {}
  project      = "%{project_name}"
}


`, context)
}

func testAccCheckVertexAiEndpointTrafficSplitDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "rs.google_vertex_ai_endpoint_traffic_split" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			billingProject := ""
			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			obj := &vertexai.EndpointTrafficSplit{
				Endpoint: dcl.String(rs.Primary.Attributes["endpoint"]),
				Location: dcl.String(rs.Primary.Attributes["location"]),
				Project:  dcl.StringOrNil(rs.Primary.Attributes["project"]),
				Etag:     dcl.StringOrNil(rs.Primary.Attributes["etag"]),
			}

			client := NewDCLVertexAiClient(config, config.userAgent, billingProject, 0)
			_, err := client.GetEndpointTrafficSplit(context.Background(), obj)
			if err == nil {
				return fmt.Errorf("google_vertex_ai_endpoint_traffic_split still exists %v", obj)
			}
		}
		return nil
	}
}
