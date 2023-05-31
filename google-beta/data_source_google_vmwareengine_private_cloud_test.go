// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceVmwareEnginePrivateCloud_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"region":        acctest.GetTestRegionFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckVmwareenginePrivateCloudDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceVmwareEnginePrivateCloudConfig(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceStateWithIgnores("data.google_vmwareengine_private_cloud.ds", "google_vmwareengine_private_cloud.pc", map[string]struct{}{}),
				),
			},
		},
	})
}

func testAccDataSourceVmwareEnginePrivateCloudConfig(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vmwareengine_network" "default-nw" {
   name              = "%{region}-default"
   provider 		 = google-beta
   location          = "%{region}"
   type              = "LEGACY"
}

resource "google_vmwareengine_private_cloud" "pc" {
  location = "%{region}-a"
  name = "tf-test-sample-pc%{random_suffix}"
  provider = google-beta
  description = ""
  network_config {
    management_cidr = "192.168.30.0/24"
    vmware_engine_network = google_vmwareengine_network.default-nw.id
  }

  management_cluster {
    cluster_id = "tf-test-sample-mgmt-cluster%{random_suffix}"
    node_type_configs {
      node_type_id = "standard-72"
      node_count = 3
    }
  }
}

data "google_vmwareengine_private_cloud" ds {
	location = "%{region}-a"
	provider = google-beta
  	name = "tf-test-sample-pc%{random_suffix}"
	depends_on = [
   	 google_vmwareengine_private_cloud.pc,
  	]
}
`, context)
}
