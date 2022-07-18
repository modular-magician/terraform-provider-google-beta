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

func TestAccVertexAiModelDeployment_BasicModelDeployment(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  getTestProjectFromEnv(),
		"region":        getTestRegionFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVertexAiModelDeploymentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAiModelDeployment_BasicModelDeployment(context),
			},
			{
				ResourceName:            "google_vertex_ai_model_deployment.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"endpoint", "model"},
			},
		},
	})
}

func testAccVertexAiModelDeployment_BasicModelDeployment(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vertex_ai_model_deployment" "primary" {
  dedicated_resources {
    machine_spec {
      machine_type = "n1-standard-2"
    }

    min_replica_count = 1
    max_replica_count = 1
  }

  endpoint = "projects/%{project_name}/locations/%{region}/endpoints/${google_vertex_ai_endpoint.minimal.name}"
  model    = "projects/%{project_name}/locations/%{region}/models/${google_vertex_ai_model.basic.name}"
  location = "%{region}"
  project  = "%{project_name}"
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

func testAccCheckVertexAiModelDeploymentDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "rs.google_vertex_ai_model_deployment" {
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

			obj := &vertexai.ModelDeployment{
				Endpoint:        dcl.String(rs.Primary.Attributes["endpoint"]),
				Model:           dcl.String(rs.Primary.Attributes["model"]),
				Location:        dcl.StringOrNil(rs.Primary.Attributes["location"]),
				Project:         dcl.StringOrNil(rs.Primary.Attributes["project"]),
				DeployedModelId: dcl.StringOrNil(rs.Primary.Attributes["deployed_model_id"]),
			}

			client := NewDCLVertexAiClient(config, config.userAgent, billingProject, 0)
			_, err := client.GetModelDeployment(context.Background(), obj)
			if err == nil {
				return fmt.Errorf("google_vertex_ai_model_deployment still exists %v", obj)
			}
		}
		return nil
	}
}
