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

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccDataFusionInstance_dataFusionInstanceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckDataFusionInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusionInstance_dataFusionInstanceBasicExample(context),
			},
		},
	})
}

func testAccDataFusionInstance_dataFusionInstanceBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_fusion_instance" "basic_instance" {
  provider = "google-beta"
  name = "my-instance%<random_suffix>s"
  region = "us-central1"
  type = "BASIC"
}
`, context)
}

func TestAccDataFusionInstance_dataFusionInstanceFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckDataFusionInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusionInstance_dataFusionInstanceFullExample(context),
			},
		},
	})
}

func testAccDataFusionInstance_dataFusionInstanceFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_fusion_instance" "extended_instance" {
  provider = "google-beta"
  name = "my-instance%<random_suffix>s"
  description = "My Data Fusion instance"
  region = "us-central1"
  type = "BASIC"
  enable_stackdriver_logging = true
  enable_stackdriver_monitoring = true
  labels = {
    example_key = "example_value"
  }
  private_instance = true
  network_config {
    network = "default"
    ip_allocation = "10.89.48.0/22"
  }
}
`, context)
}

func testAccCheckDataFusionInstanceDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_data_fusion_instance" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{DataFusionBasePath}}projects/{{project}}/locations/{{region}}/instances/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("DataFusionInstance still exists at %s", url)
		}
	}

	return nil
}
