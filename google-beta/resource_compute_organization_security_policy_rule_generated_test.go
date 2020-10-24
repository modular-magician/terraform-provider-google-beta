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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccComputeOrganizationSecurityPolicyRule_organizationSecurityPolicyRuleBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckComputeOrganizationSecurityPolicyRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeOrganizationSecurityPolicyRule_organizationSecurityPolicyRuleBasicExample(context),
			},
			{
				ResourceName:            "google_compute_organization_security_policy_rule.policy",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"policy_id"},
			},
		},
	})
}

func testAccComputeOrganizationSecurityPolicyRule_organizationSecurityPolicyRuleBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_organization_security_policy" "policy" {
  provider = google-beta

  display_name = "tf-test%{random_suffix}"
  parent       = "organizations/%{org_id}"
}

resource "google_compute_organization_security_policy_rule" "policy" {
  provider = google-beta

  policy_id = google_compute_organization_security_policy.policy.id
  action = "allow"

  direction = "INGRESS"
  enable_logging = true
  match {
    config {
      src_ip_ranges = ["192.168.0.0/16", "10.0.0.0/8"]
      layer4_config {
        ip_protocol = "tcp"
        ports = ["22"]
      }
      layer4_config {
        ip_protocol = "icmp"
      }
    }
  }
  priority = 100
}
`, context)
}

func testAccCheckComputeOrganizationSecurityPolicyRuleDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_organization_security_policy_rule" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}{{policy_id}}/getRule?priority={{priority}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeOrganizationSecurityPolicyRule still exists at %s", url)
			}
		}

		return nil
	}
}
