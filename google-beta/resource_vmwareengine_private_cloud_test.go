// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccVmwareenginePrivateCloud_vmwareEnginePrivateCloudBasicExample(t *testing.T) {
	t.Parallel()
	context := map[string]interface{}{
		"region":        GetTestRegionFromEnv(),
		"random_suffix": RandString(t, 10),
	}
	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckVmwareenginePrivateCloudDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVmwareenginePrivateCloud_vmwareEnginePrivateCloudBasicExample(context),
			},
			{
				ResourceName:            "google_vmwareengine_private_cloud.vmw-engine-pc",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name"},
			},
		},
	})
}

func testAccVmwareenginePrivateCloud_vmwareEnginePrivateCloudBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vmwareengine_network" "default-nw" {
   provider      	 = google-beta
   name              = "%{region}-default"
   location          = "%{region}"
   type              = "LEGACY"
}
resource "google_vmwareengine_private_cloud" "vmw-engine-pc" {
  provider = google-beta
  location = "%{region}-a"
  name = "tf-test-sample-pc%{random_suffix}"
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
`, context)
}

func TestAccVmwareenginePrivateCloud_vmwareEnginePrivateCloudUpdateAndExpand(t *testing.T) {
	t.Parallel()
	context := map[string]interface{}{
		"region":        acctest.GetTestRegionFromEnv(),
		"random_suffix": RandString(t, 10),
	}
	configTemplate := privateCloudUpdateConfigTemplate(context)
	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckVmwareenginePrivateCloudDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(configTemplate, "description1", 3),
			},
			{
				ResourceName:            "google_vmwareengine_private_cloud.vmw-engine-pc",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name"},
			},
			{
				Config: fmt.Sprintf(configTemplate, "description2", 4),
			},
			{
				ResourceName:            "google_vmwareengine_private_cloud.vmw-engine-pc",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name"},
			},
		},
	})
}

func TestAccVmwareenginePrivateCloud_vmwareEnginePrivateCloudUpdateAndShrink(t *testing.T) {
	t.Parallel()
	context := map[string]interface{}{
		"region":        acctest.GetTestRegionFromEnv(),
		"random_suffix": RandString(t, 10),
	}
	configTemplate := privateCloudUpdateConfigTemplate(context)
	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckVmwareenginePrivateCloudDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(configTemplate, "description1", 4),
			},
			{
				ResourceName:            "google_vmwareengine_private_cloud.vmw-engine-pc",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name"},
			},
			{
				Config: fmt.Sprintf(configTemplate, "description2", 3),
			},
			{
				ResourceName:            "google_vmwareengine_private_cloud.vmw-engine-pc",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name"},
			},
		},
	})
}

func privateCloudUpdateConfigTemplate(context map[string]interface{}) string {

	return Nprintf(`
resource "google_vmwareengine_network" "default-nw" {
   provider      	 = google-beta
   name              = "%{region}-default"
   location          = "%{region}"
   type              = "LEGACY"
}

resource "google_vmwareengine_private_cloud" "vmw-engine-pc" {
  location = "%{region}-a"
  name = "tf-test-sample-pc%{random_suffix}"
  provider = google-beta
  description = "%s"
  network_config {
    management_cidr = "192.168.30.0/24"
    vmware_engine_network = google_vmwareengine_network.default-nw.id
  }
  management_cluster {
    cluster_id = "tf-test-sample-mgmt-cluster-custom-core-count%{random_suffix}"
    node_type_configs {
      node_type_id = "standard-72"
      node_count = %d
      custom_core_count = 32
    }
  }
}
`, context)
}

func testAccCheckVmwareenginePrivateCloudDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_vmwareengine_private_cloud" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}
			config := GoogleProviderConfig(t)
			url, err := replaceVarsForTest(config, rs, "{{VmwareengineBasePath}}projects/{{project}}/locations/{{location}}/privateClouds/{{name}}")
			if err != nil {
				return err
			}
			billingProject := ""
			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}
			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("VmwareenginePrivateCloud still exists at %s", url)
			}
		}
		return nil
	}
}
