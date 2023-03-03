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
)

func TestAccOrgPolicyCustomConstraint_orgPolicyCustomConstraintBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        GetTestOrgFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProvidersOiCS,
		CheckDestroy: testAccCheckOrgPolicyCustomConstraintDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccOrgPolicyCustomConstraint_orgPolicyCustomConstraintBasicExample(context),
			},
			{
				ResourceName:            "google_org_policy_custom_constraint.constraint",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccOrgPolicyCustomConstraint_orgPolicyCustomConstraintBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_org_policy_custom_constraint" "constraint" {
  provider = google-beta

  name         = "custom.disableGkeAutoUpgrade"
  parent       = "organizations/%{org_id}"

  action_type    = "ALLOW"
  condition      = "resource.management.autoUpgrade == false"
  method_types   = ["CREATE", "UPDATE"]
  resource_types = ["container.googleapis.com/NodePool"]
}
`, context)
}

func TestAccOrgPolicyCustomConstraint_orgPolicyCustomConstraintFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        GetTestOrgTargetFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProvidersOiCS,
		CheckDestroy: testAccCheckOrgPolicyCustomConstraintDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccOrgPolicyCustomConstraint_orgPolicyCustomConstraintFullExample(context),
			},
			{
				ResourceName:            "google_org_policy_custom_constraint.constraint",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccOrgPolicyCustomConstraint_orgPolicyCustomConstraintFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_org_policy_custom_constraint" "constraint" {
  provider = google-beta

  name         = "custom.disableGkeAutoUpgrade"
  parent       = "organizations/%{org_id}"
  display_name = "Disable GKE auto upgrade"
  description  = "Only allow GKE NodePool resource to be created or updated if AutoUpgrade is not enabled where this custom constraint is enforced."

  action_type    = "ALLOW"
  condition      = "resource.management.autoUpgrade == false"
  method_types   = ["CREATE", "UPDATE"]
  resource_types = ["container.googleapis.com/NodePool"]
}

resource "google_org_policy_policy" "bool" {
  provider = google-beta

  name   = "organizations/%{org_id}/policies/${google_org_policy_custom_constraint.constraint.name}"
  parent = "organizations/%{org_id}"

  spec {
    rules {
      enforce = "TRUE"
    }
  }
}
`, context)
}

func testAccCheckOrgPolicyCustomConstraintDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_org_policy_custom_constraint" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{OrgPolicyBasePath}}{{parent}}/customConstraints/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("OrgPolicyCustomConstraint still exists at %s", url)
			}
		}

		return nil
	}
}
