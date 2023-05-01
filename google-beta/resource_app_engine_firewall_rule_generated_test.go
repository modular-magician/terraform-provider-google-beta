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
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccAppEngineFirewallRule_appEngineFirewallRuleBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        acctest.GetTestOrgFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckAppEngineFirewallRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAppEngineFirewallRule_appEngineFirewallRuleBasicExample(context),
			},
			{
				ResourceName:      "google_app_engine_firewall_rule.rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAppEngineFirewallRule_appEngineFirewallRuleBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "my_project" {
  name       = "tf-test-project"
  project_id = "tf-test-ae-project%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "google_app_engine_application" "app" {
  project     = google_project.my_project.project_id
  location_id = "us-central"
}

resource "google_app_engine_firewall_rule" "rule" {
  project      = google_app_engine_application.app.project
  priority     = 1000
  action       = "ALLOW"
  source_range = "*"
}
`, context)
}

func testAccCheckAppEngineFirewallRuleDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_app_engine_firewall_rule" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := acctest.ReplaceVarsForTest(config, rs, "{{AppEngineBasePath}}apps/{{project}}/firewall/ingressRules/{{priority}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("AppEngineFirewallRule still exists at %s", url)
			}
		}

		return nil
	}
}
