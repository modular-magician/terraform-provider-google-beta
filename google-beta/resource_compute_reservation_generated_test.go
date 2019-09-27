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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccComputeReservation_reservationBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeReservationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeReservation_reservationBasicExample(context),
			},
			{
				ResourceName:      "google_compute_reservation.gce_reservation",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeReservation_reservationBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_reservation" "gce_reservation" {
  name = "gce-reservation%{random_suffix}"
  zone = "us-central1-a"

  specific_reservation {
    count = 1
    instance_properties {
      min_cpu_platform = "Intel Cascade Lake"
      machine_type = "n2-standard-2"
    }
  }
}
`, context)
}

func testAccCheckComputeReservationDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_reservation" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/reservations/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeReservation still exists at %s", url)
		}
	}

	return nil
}
