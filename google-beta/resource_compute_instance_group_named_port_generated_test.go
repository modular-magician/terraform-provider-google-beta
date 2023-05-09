// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccComputeInstanceGroupNamedPort_instanceGroupNamedPortGkeExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceGroupNamedPortDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceGroupNamedPort_instanceGroupNamedPortGkeExample(context),
			},
			{
				ResourceName:            "google_compute_instance_group_named_port.my_port",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"group", "zone"},
			},
		},
	})
}

func testAccComputeInstanceGroupNamedPort_instanceGroupNamedPortGkeExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance_group_named_port" "my_port" {
  group = google_container_cluster.my_cluster.node_pool[0].instance_group_urls[0]
  zone = "us-central1-a"

  name = "http"
  port = 8080
}

resource "google_compute_instance_group_named_port" "my_ports" {
  group = google_container_cluster.my_cluster.node_pool[0].instance_group_urls[0]
  zone = "us-central1-a"

  name = "https"
  port = 4443
}

resource "google_compute_network" "container_network" {
  name                    = "tf-test-container-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "container_subnetwork" {
  name                     = "tf-test-container-subnetwork%{random_suffix}"
  region                   = "us-central1"
  network                  = google_compute_network.container_network.name
  ip_cidr_range            = "10.0.36.0/24"
}

resource "google_container_cluster" "my_cluster" {
  name               = "tf-test-my-cluster%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1

  network    = google_compute_network.container_network.name
  subnetwork = google_compute_subnetwork.container_subnetwork.name

  ip_allocation_policy {
    cluster_ipv4_cidr_block  = "/19"
    services_ipv4_cidr_block = "/22"
  }
}
`, context)
}

func testAccCheckComputeInstanceGroupNamedPortDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_instance_group_named_port" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeInstanceGroupNamedPort still exists at %s", url)
			}
		}

		return nil
	}
}
