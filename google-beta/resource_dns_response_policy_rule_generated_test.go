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

func TestAccDNSResponsePolicyRule_dnsResponsePolicyRuleBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckDNSResponsePolicyRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSResponsePolicyRule_dnsResponsePolicyRuleBasicExample(context),
			},
			{
				ResourceName:            "google_dns_response_policy_rule.example-response-policy-rule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"response_policy"},
			},
		},
	})
}

func testAccDNSResponsePolicyRule_dnsResponsePolicyRuleBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network" "network-1" {
  provider = google-beta

  name                    = "tf-test-network-1%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_network" "network-2" {
  provider = google-beta
  
  name                    = "tf-test-network-2%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_dns_response_policy" "response-policy" {
  provider = google-beta

  response_policy_name = "tf-test-example-response-policy%{random_suffix}"
  
  networks {
    network_url = google_compute_network.network-1.id
  }
  networks {
    network_url = google_compute_network.network-2.id
  }
}

resource "google_dns_response_policy_rule" "example-response-policy-rule" {
  provider = google-beta

  response_policy = google_dns_response_policy.response-policy.response_policy_name
  rule_name       = "tf-test-example-rule%{random_suffix}"
  dns_name        = "dns.example.com."

  local_data {
    local_datas {
      name    = "dns.example.com."
      type    = "A"
      ttl     = 300
      rrdatas = ["192.0.2.91"]
    }
  }  

}
`, context)
}

func testAccCheckDNSResponsePolicyRuleDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dns_response_policy_rule" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DNSBasePath}}projects/{{project}}/responsePolicies/{{response_policy}}/rules/{{rule_name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("DNSResponsePolicyRule still exists at %s", url)
			}
		}

		return nil
	}
}
