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

func TestAccVertexAiEndpoint_BasicEndpoint(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  getTestProjectFromEnv(),
		"region":        getTestRegionFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVertexAiEndpointDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAiEndpoint_BasicEndpoint(context),
			},
			{
				ResourceName:            "google_vertex_ai_endpoint.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}
func TestAccVertexAiEndpoint_NetworkHandWritten(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_acct":   getTestBillingAccountFromEnv(t),
		"org_id":         getTestOrgFromEnv(t),
		"project_name":   getTestProjectFromEnv(),
		"project_number": getTestProjectNumberFromEnv(),
		"region":         getTestRegionFromEnv(),
		"random_suffix":  randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVertexAiEndpointDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAiEndpoint_NetworkHandWritten(context),
			},
			{
				ResourceName:            "google_vertex_ai_endpoint.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func testAccVertexAiEndpoint_BasicEndpoint(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vertex_ai_endpoint" "primary" {
  display_name = "sample-endpoint"
  location     = "%{region}"
  labels       = {}
  project      = "%{project_name}"
}


`, context)
}

func testAccVertexAiEndpoint_NetworkHandWritten(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vertex_ai_endpoint" "primary" {
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  location     = "%{region}"
  labels       = {
    label-one = "value-one"
  }
  network      = google_compute_network.vertex_network.id
  project      = "%{project_name}"
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
  name          = "tf-test-address%{random_suffix}"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 24
  network       = google_compute_network.vertex_network.id
  project       = "%{project_name}"
}

resource "google_compute_network" "vertex_network" {
  name       = "tf-test-network%{random_suffix}"
  project    = "%{project_name}"
}

`, context)
}

func testAccCheckVertexAiEndpointDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "rs.google_vertex_ai_endpoint" {
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

			obj := &vertexai.Endpoint{
				DisplayName:                  dcl.String(rs.Primary.Attributes["display_name"]),
				Location:                     dcl.String(rs.Primary.Attributes["location"]),
				Description:                  dcl.String(rs.Primary.Attributes["description"]),
				Network:                      dcl.String(rs.Primary.Attributes["network"]),
				Project:                      dcl.StringOrNil(rs.Primary.Attributes["project"]),
				CreateTime:                   dcl.StringOrNil(rs.Primary.Attributes["create_time"]),
				Etag:                         dcl.StringOrNil(rs.Primary.Attributes["etag"]),
				ModelDeploymentMonitoringJob: dcl.StringOrNil(rs.Primary.Attributes["model_deployment_monitoring_job"]),
				Name:                         dcl.StringOrNil(rs.Primary.Attributes["name"]),
				UpdateTime:                   dcl.StringOrNil(rs.Primary.Attributes["update_time"]),
			}

			client := NewDCLVertexAiClient(config, config.userAgent, billingProject, 0)
			_, err := client.GetEndpoint(context.Background(), obj)
			if err == nil {
				return fmt.Errorf("google_vertex_ai_endpoint still exists %v", obj)
			}
		}
		return nil
	}
}
