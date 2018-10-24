// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccComputeRegionDisk_RegionDiskBasicExample(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeRegionDiskDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionDisk_RegionDiskBasicExample(acctest.RandString(10)),
			},
			{
				ResourceName:      "google_compute_region_disk.regiondisk",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionDisk_RegionDiskBasicExample(val string) string {
	return fmt.Sprintf(`
resource "google_compute_region_disk" "regiondisk" {
  name = "my-region-disk-%s"
  snapshot = "${google_compute_snapshot.snapdisk.self_link}"
  type = "pd-ssd"
  region = "us-central1"

  replica_zones = ["us-central1-a", "us-central1-f"]
}

resource "google_compute_disk" "disk" {
  name = "my-disk-%s"
  image = "debian-cloud/debian-9"
  size = 50
  type = "pd-ssd"
  zone = "us-central1-a"
}

resource "google_compute_snapshot" "snapdisk" {
  name = "my-snapshot-%s"
  source_disk = "${google_compute_disk.disk.name}"
  zone = "us-central1-a"
}
`, val, val, val,
	)
}
