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

func TestAccNetworkSecurityGatewaySecurityPolicies_networkSecurityGatewaySecurityPoliciesBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityGatewaySecurityPoliciesDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityGatewaySecurityPolicies_networkSecurityGatewaySecurityPoliciesBasicExample(context),
			},
			{
				ResourceName:            "google_network_security_gateway_security_policies.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccNetworkSecurityGatewaySecurityPolicies_networkSecurityGatewaySecurityPoliciesBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_network_security_gateway_security_policies" "default" {
  provider    = google-beta
  name        = "tf-test-my-gateway-security-policy%{random_suffix}"
  location    = "us-central1"
  description = "my description"
}
`, context)
}

func testAccCheckNetworkSecurityGatewaySecurityPoliciesDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_security_gateway_security_policies" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("NetworkSecurityGatewaySecurityPolicies still exists at %s", url)
			}
		}

		return nil
	}
}
