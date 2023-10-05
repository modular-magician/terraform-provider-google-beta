// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package netapp_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccNetappstoragePool_storagePoolCreateExample_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetappstoragePool_storagePoolCreateExample_full(context),
			},
			{
				ResourceName:            "google_netapp_storage_pool.test_pool",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "storage_pool_id"},
			},
			{
				Config: testAccNetappstoragePool_storagePoolCreateExample_update(context),
			},
			{
				ResourceName:            "google_netapp_storage_pool.test_pool",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "storage_pool_id"},
			},
		},
	})
}

func testAccNetappstoragePool_storagePoolCreateExample_full(context map[string]interface{}) string {
	return acctest.Nprintf(`

resource "google_netapp_storage_pool" "test_pool" {
  storage_pool_id = "tf-test-test-pool"
  location = "us-central1"
  service_level = "PREMIUM"
  capacity_gib = "2048"
  network = "projects/{{ProjectNumber}}/global/networks/cxo-automation-vpc-01"
}
`, context)
}

func testAccNetappstoragePool_storagePoolCreateExample_update(context map[string]interface{}) string {
	return acctest.Nprintf(`

resource "google_netapp_storage_pool" "test_pool" {
  storage_pool_id = "tf-test-test-pool"
  location = "us-central1"
  service_level = "PREMIUM"
  capacity_gib = "4096"
  network = "projects/{{ProjectNumber}}/global/networks/cxo-automation-vpc-01"
}
`, context)
}
