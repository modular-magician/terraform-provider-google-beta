package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataprocMetastoreService_updateAndImport(t *testing.T) {
	t.Parallel()

	name := "tf-metastore-" + randString(t, 10)
	tier := [2]string{"DEVELOPER", "ENTERPRISE"}
	resourceName := "google_dataproc_metastore_service.my_metastore"

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigqueryReservationReservationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocMetastoreService_updateAndImport(name, tier[0]),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataprocMetastoreService_updateAndImport(name, tier[1]),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataprocMetastoreService_networking(t *testing.T) {
	t.Parallel()

	name := "tf-metastore-" + randString(t, 10)
	resourceName := "google_dataproc_metastore_service.my_metastore"

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigqueryReservationReservationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocMetastoreService_networking(name, "DEVELOPER"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "network_config.0.consumers.0.endpoint_uri"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataprocMetastoreService_updateAndImport(name, tier string) string {
	return fmt.Sprintf(`
resource "google_dataproc_metastore_service" "my_metastore" {
	service_id = "%s"
	location   = "us-central1"
	tier       = "%s"

	hive_metastore_config {
		version = "2.3.6"
	}
}
`, name, tier)
}

func testAccDataprocMetastoreService_networking(name, tier string) string {
	return fmt.Sprintf(`
data "google_compute_subnetwork" "test" {
  name   = "default"
  region = "us-central1"
}

resource "google_dataproc_metastore_service" "my_metastore" {
	service_id = "%[1]s"
	location   = "us-central1"
	tier       = "%[2]s"

	hive_metastore_config {
		version = "2.3.6"
	}

	network_config {
		consumers {
			subnetwork = google_compute_subnetwork.test.id
		}
	}
}
`, name, tier)
}
