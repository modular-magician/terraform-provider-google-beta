package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataFusion_basic(t *testing.T) {
	t.Parallel()

	instanceName := fmt.Sprintf("tf-test-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusion_basic(instanceName),
			},
			{
				ResourceName:      "google_data_fusion_instance.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataFusion_extended(t *testing.T) {
	t.Parallel()

	instanceName := fmt.Sprintf("tf-test-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusion_extended(instanceName),
			},
			{
				ResourceName:      "google_data_fusion_instance.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataFusion_updateLabels(t *testing.T) {
	t.Parallel()

	instanceName := fmt.Sprintf("tf-test-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusion_basic(instanceName),
			},
			{
				ResourceName:      "google_data_fusion_instance.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataFusion_withLabels(instanceName),
			},
			{
				ResourceName:      "google_data_fusion_instance.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataFusion_basic(instanceName string) string {
	return fmt.Sprintf(`
resource "google_data_fusion_instance" "foobar" {
	name = "%s"
	region = "us-central1"
	type = "BASIC"
}`, instanceName)
}

func testAccDataFusion_extended(instanceName string) string {
	return fmt.Sprintf(`
resource "google_data_fusion_instance" "foobar" {
	name = "%s"
	description = "Some description"
	region = "us-central1"
	type = "ENTERPRISE"
	enable_stackdriver_monitoring = true
	enable_stackdriver_logging = true
	labels = {
		label1 = "value1"
	}
}`, instanceName)
}

func testAccDataFusion_withLabels(instanceName string) string {
	return fmt.Sprintf(`
resource "google_data_fusion_instance" "foobar" {
	name = "%s"
	region = "us-central1"
	type = "BASIC"

	labels = {
		label1 = "value1"
		label2 = "value2"
	}
}`, instanceName)
}
