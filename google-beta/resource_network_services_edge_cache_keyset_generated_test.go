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

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccNetworkServicesEdgeCacheKeyset_networkServicesEdgeCacheKeysetBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesEdgeCacheKeysetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesEdgeCacheKeyset_networkServicesEdgeCacheKeysetBasicExample(context),
			},
			{
				ResourceName:            "google_network_services_edge_cache_keyset.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesEdgeCacheKeyset_networkServicesEdgeCacheKeysetBasicExample(context map[string]interface{}) string {
	return Nprintf(`

resource "google_network_services_edge_cache_keyset" "default" {
  name                 = "tf-test-my-keyset%{random_suffix}"
  description          = "The default keyset"
  public_key {
    id = "my-public-key"
    value = "FHsTyFHNmvNpw4o7-rp-M1yqMyBF8vXSBRkZtkQ0RKY"
  }
  public_key {
    id = "my-public-key-2"
    value = "hzd03llxB1u5FOLKFkZ6_wCJqC7jtN0bg7xlBqS6WVM"
  }
}
`, context)
}

func TestAccNetworkServicesEdgeCacheKeyset_networkServicesEdgeCacheKeysetDualTokenExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesEdgeCacheKeysetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesEdgeCacheKeyset_networkServicesEdgeCacheKeysetDualTokenExample(context),
			},
			{
				ResourceName:            "google_network_services_edge_cache_keyset.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesEdgeCacheKeyset_networkServicesEdgeCacheKeysetDualTokenExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "tf-test-secret-name%{random_suffix}"

  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret-version-basic" {
  secret = google_secret_manager_secret.secret-basic.id

  secret_data = "secret-data"
}

resource "google_network_services_edge_cache_keyset" "default" {
  name        = "tf-test-my-keyset%{random_suffix}"
  description = "The default keyset"
  public_key {
    id      = "my-public-key"
    managed = true
  }
  validation_shared_keys {
    secret_version = google_secret_manager_secret_version.secret-version-basic.id
  }
}
`, context)
}

func testAccCheckNetworkServicesEdgeCacheKeysetDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_services_edge_cache_keyset" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("NetworkServicesEdgeCacheKeyset still exists at %s", url)
			}
		}

		return nil
	}
}
