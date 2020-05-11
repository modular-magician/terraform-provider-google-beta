package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccCloudRunService_cloudRunVPCServiceUpdate(t *testing.T) {
	t.Parallel()

	project := getTestProjectFromEnv()
	name := "tfvpctest-cloudrun-" + randString(t, 6)
	networkName := fmt.Sprintf("tf-test-net-%d", randInt(t))
	vpcConnectorName := fmt.Sprintf("tf-test-conn-%s", randString(t, 5))

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunVPCServiceUpdate(name, project, networkName, vpcConnectorName, "10.10.0.0/28", "10"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_cloudRunVPCServiceUpdate(name, project, networkName, vpcConnectorName, "10.10.0.0/28", "50"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunVPCServiceUpdate(name, project, networkName, vpcConnectorName, vpcIp, concurrency string) string {
	return fmt.Sprintf(`
	resource "google_compute_network" "vpc" {
		name = "%s"
		auto_create_subnetworks = false
	}
	
	resource "google_vpc_access_connector" "%s" {
		name          = "%s"
		region        = "us-central1"
		ip_cidr_range = "%s"
		network       = google_compute_network.vpc.name
	}

	resource "google_cloud_run_service" "default" {
		name     = "%s"
		location = "us-central1"
    provider = google-beta

		metadata {
			namespace = "%s"
		}
		
		vpc_connector	=	google_vpc_access_connector.%s.self_link

		template {
			spec {
				containers {
					image = "gcr.io/cloudrun/hello"
					args  = ["arrgs"]
				}
			container_concurrency = %s
			}
		}

		traffic {
			percent         = 100
			latest_revision = true
		}
	}
`, networkName, vpcConnectorName, vpcConnectorName, vpcIp, name, project, vpcConnectorName, concurrency)
}
