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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSecretManagerSecretVersion_secretVersionBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecretManagerSecretVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerSecretVersion_secretVersionBasicExample(context),
			},
			{
				ResourceName:            "google_secret_manager_secret_version.secret-version-basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"secret"},
			},
		},
	})
}

func testAccSecretManagerSecretVersion_secretVersionBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "tf-test-secret-version%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    automatic = true
  }
}


resource "google_secret_manager_secret_version" "secret-version-basic" {
  secret = google_secret_manager_secret.secret-basic.id

  secret_data = "tf-test-secret-data%{random_suffix}"
}
`, context)
}

func testAccCheckSecretManagerSecretVersionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_secret_manager_secret_version" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{SecretManagerBasePath}}{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("SecretManagerSecretVersion still exists at %s", url)
			}
		}

		return nil
	}
}
