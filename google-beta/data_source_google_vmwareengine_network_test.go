// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceVmwareEngineNetwork_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"region":        acctest.GetTestRegionFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckVmwareengineNetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceVmwareEngineNetworkConfig(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceStateWithIgnores("data.google_vmwareengine_network.ds", "google_vmwareengine_network.nw", map[string]struct{}{}),
				),
			},
		},
	})
}

func testAccDataSourceVmwareEngineNetworkConfig(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vmwareengine_network" "nw" {
    name              = "%{region}-default" #Legacy network IDs are in the format: {region-id}-default
	provider = google-beta
    location          = "%{region}"
    type              = "LEGACY"
    description       = "VMwareEngine legacy network sample"
}


data "google_vmwareengine_network" "ds" {
  name = google_vmwareengine_network.nw.name
  provider = google-beta
  location = "%{region}"
  depends_on = [
    google_vmwareengine_network.nw,
  ]
}
`, context)
}
