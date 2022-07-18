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

func TestAccVertexAiModel_BasicModel(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  getTestProjectFromEnv(),
		"region":        getTestRegionFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVertexAiModelDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAiModel_BasicModel(context),
			},
			{
				ResourceName:            "google_vertex_ai_model.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
			{
				Config: testAccVertexAiModel_BasicModelUpdate0(context),
			},
			{
				ResourceName:            "google_vertex_ai_model.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func testAccVertexAiModel_BasicModel(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vertex_ai_model" "primary" {
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


`, context)
}

func testAccVertexAiModel_BasicModelUpdate0(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vertex_ai_model" "primary" {
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

  display_name = "new-sample-model"
  location     = "%{region}"
  artifact_uri = "gs://cloud-samples-data/vertex-ai/google-cloud-aiplatform-ci-artifacts/models/iris_xgboost/"
  description  = "An updated sample model"

  labels = {
    label-two = "value-two"
  }

  project             = "%{project_name}"
  version_aliases     = ["default", "v1", "v2"]
  version_description = "A sample model version"
}


`, context)
}

func testAccCheckVertexAiModelDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "rs.google_vertex_ai_model" {
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

			obj := &vertexai.Model{
				DisplayName:        dcl.String(rs.Primary.Attributes["display_name"]),
				Location:           dcl.String(rs.Primary.Attributes["location"]),
				ArtifactUri:        dcl.String(rs.Primary.Attributes["artifact_uri"]),
				Description:        dcl.String(rs.Primary.Attributes["description"]),
				Project:            dcl.StringOrNil(rs.Primary.Attributes["project"]),
				VersionDescription: dcl.String(rs.Primary.Attributes["version_description"]),
				CreateTime:         dcl.StringOrNil(rs.Primary.Attributes["create_time"]),
				Etag:               dcl.StringOrNil(rs.Primary.Attributes["etag"]),
				Name:               dcl.StringOrNil(rs.Primary.Attributes["name"]),
				TrainingPipeline:   dcl.StringOrNil(rs.Primary.Attributes["training_pipeline"]),
				UpdateTime:         dcl.StringOrNil(rs.Primary.Attributes["update_time"]),
				VersionCreateTime:  dcl.StringOrNil(rs.Primary.Attributes["version_create_time"]),
				VersionId:          dcl.StringOrNil(rs.Primary.Attributes["version_id"]),
				VersionUpdateTime:  dcl.StringOrNil(rs.Primary.Attributes["version_update_time"]),
			}

			client := NewDCLVertexAiClient(config, config.userAgent, billingProject, 0)
			_, err := client.GetModel(context.Background(), obj)
			if err == nil {
				return fmt.Errorf("google_vertex_ai_model still exists %v", obj)
			}
		}
		return nil
	}
}
