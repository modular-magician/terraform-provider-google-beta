// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceVmwareEngineCluster_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"region":        acctest.GetTestRegionFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckVmwareengineClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceVmwareEngineCluster(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceStateWithIgnores("data.google_vmwareengine_cluster.ds", "google_vmwareengine_cluster.cls", map[string]struct{}{}),
				),
			},
		},
	})
}

func testAccDataSourceVmwareEngineCluster(context map[string]interface{}) string {
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

resource "google_vmwareengine_cluster" "cls" {
    name = "tf-test-ext-cluster%{random_suffix}"
	provider = google-beta
    parent =  google_vmwareengine_private_cloud.pc.id
    node_type_configs {
      node_type_id = "standard-72"
      node_count = 3
      custom_core_count = 32
    }
}

data "google_vmwareengine_cluster" ds {
  	name = "tf-test-ext-cluster%{random_suffix}"
	provider = google-beta
	parent = google_vmwareengine_private_cloud.pc.id
	depends_on = [
   	 google_vmwareengine_cluster.cls,
  	]
}
`, context)
}
