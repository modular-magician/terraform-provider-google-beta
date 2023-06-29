// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccComputeGlobalNetworkEndpoint_networkEndpointsBasic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"default_port":  90,
		"modified_port": 100,
	}
	negId := fmt.Sprintf("projects/%s/global/networkEndpointGroups/neg-%s",
		envvar.GetTestProjectFromEnv(), context["random_suffix"])

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Create one endpoint
				Config: testAccComputeGlobalNetworkEndpoint_networkEndpointsBasic(context),
			},
			{
				ResourceName:      "google_compute_global_network_endpoint.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Force-recreate old endpoint
				Config: testAccComputeGlobalNetworkEndpoint_networkEndpointsModified(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeNetworkEndpointWithPortsDestroyed(t, negId, "90"),
				),
			},
			{
				ResourceName:      "google_compute_global_network_endpoint.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// delete all endpoints
				Config: testAccComputeGlobalNetworkEndpoint_noNetworkEndpoints(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeNetworkEndpointWithPortsDestroyed(t, negId, "100"),
				),
			},
		},
	})
}

func testAccComputeGlobalNetworkEndpoint_networkEndpointsBasic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_global_network_endpoint" "default" {
  global_network_endpoint_group = google_compute_global_network_endpoint_group.neg.id

  ip_address = "8.8.8.8"
  port       = google_compute_global_network_endpoint_group.neg.default_port
}
`, context) + testAccComputeGlobalNetworkEndpoint_noNetworkEndpoints(context)
}

func testAccComputeGlobalNetworkEndpoint_networkEndpointsModified(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_global_network_endpoint" "default" {
  global_network_endpoint_group = google_compute_global_network_endpoint_group.neg.name

  ip_address = "8.8.8.8"
  port = "%{modified_port}"
}
`, context) + testAccComputeGlobalNetworkEndpoint_noNetworkEndpoints(context)
}

func testAccComputeGlobalNetworkEndpoint_noNetworkEndpoints(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_global_network_endpoint_group" "neg" {
  name                  = "neg-%{random_suffix}"
  default_port          = "%{default_port}"
  network_endpoint_type = "INTERNET_IP_PORT"
}
`, context)
}
