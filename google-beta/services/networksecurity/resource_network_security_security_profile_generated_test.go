// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

package networksecurity_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccNetworkSecuritySecurityProfile_networkSecuritySecurityProfileBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecuritySecurityProfileDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecuritySecurityProfile_networkSecuritySecurityProfileBasicExample(context),
			},
			{
				ResourceName:            "google_network_security_security_profile.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "parent", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecuritySecurityProfile_networkSecuritySecurityProfileBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_security_profile" "default" {
  name        = "tf-test-my-security-profile%{random_suffix}"
  parent      = "organizations/%{org_id}"
  description = "my description"
  type        = "THREAT_PREVENTION"

  labels = {
    foo = "bar"
  }
}
`, context)
}

func TestAccNetworkSecuritySecurityProfile_networkSecuritySecurityProfileOverridesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecuritySecurityProfileDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecuritySecurityProfile_networkSecuritySecurityProfileOverridesExample(context),
			},
			{
				ResourceName:            "google_network_security_security_profile.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "parent", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecuritySecurityProfile_networkSecuritySecurityProfileOverridesExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_security_profile" "default" {
  name        = "tf-test-my-security-profile%{random_suffix}"
  parent      = "organizations/%{org_id}"
  description = "my description"
  type        = "THREAT_PREVENTION"

  threat_prevention_profile {
    severity_overrides {
      action   = "ALLOW"
      severity = "INFORMATIONAL"
    }

    severity_overrides {
      action   = "DENY"
      severity = "HIGH"
    }

    threat_overrides {
      action    = "ALLOW"
      threat_id = "280647"
    }
  }
}
`, context)
}

func testAccCheckNetworkSecuritySecurityProfileDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_security_security_profile" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/securityProfiles/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("NetworkSecuritySecurityProfile still exists at %s", url)
			}
		}

		return nil
	}
}
