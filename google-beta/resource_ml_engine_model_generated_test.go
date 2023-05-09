// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccMLEngineModel_mlModelBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckMLEngineModelDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMLEngineModel_mlModelBasicExample(context),
			},
			{
				ResourceName:      "google_ml_engine_model.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMLEngineModel_mlModelBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_ml_engine_model" "default" {
  name        = "default%{random_suffix}"
  description = "My model"
  regions     = ["us-central1"]
}
`, context)
}

func TestAccMLEngineModel_mlModelFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckMLEngineModelDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMLEngineModel_mlModelFullExample(context),
			},
			{
				ResourceName:      "google_ml_engine_model.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMLEngineModel_mlModelFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_ml_engine_model" "default" {
  name        = "default%{random_suffix}"
  description = "My model"
  regions     = ["us-central1"]
  labels = {
    my_model = "foo"
  }
  online_prediction_logging         = true
  online_prediction_console_logging = true
}
`, context)
}

func testAccCheckMLEngineModelDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_ml_engine_model" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{MLEngineBasePath}}projects/{{project}}/models/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("MLEngineModel still exists at %s", url)
			}
		}

		return nil
	}
}
