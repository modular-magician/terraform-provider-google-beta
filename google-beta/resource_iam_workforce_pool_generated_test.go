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

func TestAccIAMWorkforcePoolWorkforcePool_iamWorkforcePoolBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        acctest.GetTestOrgFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAMWorkforcePoolWorkforcePoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMWorkforcePoolWorkforcePool_iamWorkforcePoolBasicExample(context),
			},
			{
				ResourceName:            "google_iam_workforce_pool.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "workforce_pool_id"},
			},
		},
	})
}

func testAccIAMWorkforcePoolWorkforcePool_iamWorkforcePoolBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workforce_pool" "example" {
  workforce_pool_id = "tf-test-example-pool%{random_suffix}"
  parent            = "organizations/%{org_id}"
  location          = "global"
}
`, context)
}

func TestAccIAMWorkforcePoolWorkforcePool_iamWorkforcePoolFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        acctest.GetTestOrgFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAMWorkforcePoolWorkforcePoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMWorkforcePoolWorkforcePool_iamWorkforcePoolFullExample(context),
			},
			{
				ResourceName:            "google_iam_workforce_pool.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "workforce_pool_id"},
			},
		},
	})
}

func testAccIAMWorkforcePoolWorkforcePool_iamWorkforcePoolFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workforce_pool" "example" {
  workforce_pool_id = "tf-test-example-pool%{random_suffix}"
  parent            = "organizations/%{org_id}"
  location          = "global"
  display_name      = "Display name"
  description       = "A sample workforce pool."
  disabled          = false
  session_duration  = "7200s"
}
`, context)
}

func testAccCheckIAMWorkforcePoolWorkforcePoolDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_iam_workforce_pool" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{IAMWorkforcePoolBasePath}}locations/{{location}}/workforcePools/{{workforce_pool_id}}")
			if err != nil {
				return err
			}

			res, err := transport_tpg.SendRequest(config, "GET", "", url, config.UserAgent, nil)
			if err != nil {
				return nil
			}

			if v := res["state"]; v == "DELETED" {
				return nil
			}

			return fmt.Errorf("IAMWorkforcePool still exists at %s", url)
		}

		return nil
	}
}
