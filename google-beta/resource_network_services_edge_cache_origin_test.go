package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkServicesEdgeCacheOrigin_updateAndImport(t *testing.T) {
	t.Parallel()

	name := "tf-test-origin-" + randString(t, 10)

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigqueryReservationReservationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocMetastoreService_update_0(name),
			},
			{
				ResourceName:      "google_network_services_edge_cache_origin.instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataprocMetastoreService_update_1(name),
			},
			{
				ResourceName:      "google_network_services_edge_cache_origin.instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataprocMetastoreService_update_0(name string) string {
	return fmt.Sprintf(`
	resource "google_network_services_edge_cache_origin" "instance" {
		name                 = "%s"
		origin_address       = "gs://media-edge-default"
		description          = "The default bucket for media edge test"
		max_attempts         = 2
		labels = {
			a = "b"
		}

		timeout {
			connect_timeout = "10s"
		}
	}
`, name)
}

func testAccDataprocMetastoreService_update_1(name string) string {
	return fmt.Sprintf(`
	resource "google_network_services_edge_cache_origin" "instance" {
		name                 = "%s"
		origin_address       = "gs://media-edge-fallback"
		description          = "The default bucket for media edge test"
		max_attempts         = 3

		timeout {
			connect_timeout = "9s"
		}
	}
`, name)
}
