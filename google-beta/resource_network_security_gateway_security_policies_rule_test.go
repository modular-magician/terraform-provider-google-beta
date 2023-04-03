package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkSecurityGatewaySecurityPoliciesRule_update(t *testing.T) {
	t.Parallel()

	gatewaySecurityPolicyName := fmt.Sprintf("tf-test-gateway-sp-%s", RandString(t, 10))
	gatewaySecurityPolicyRuleName := fmt.Sprintf("tf-test-gateway-sp-rule-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityGatewaySecurityPoliciesRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityGatewaySecurityPoliciesRule_basic(gatewaySecurityPolicyName, gatewaySecurityPolicyRuleName),
			},
			{
				ResourceName:      "google_network_security_gateway_security_policies_rule.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkSecurityGatewaySecurityPoliciesRule_update(gatewaySecurityPolicyName, gatewaySecurityPolicyRuleName),
			},
			{
				ResourceName:      "google_network_security_gateway_security_policies_rule.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkSecurityGatewaySecurityPoliciesRule_basic(gatewaySecurityPolicyName, gatewaySecurityPolicyRuleName),
			},
			{
				ResourceName:      "google_network_security_gateway_security_policies_rule.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkSecurityGatewaySecurityPoliciesRule_basic(gatewaySecurityPolicyName, gatewaySecurityPolicyRuleName string) string {
	return fmt.Sprintf(`
resource "google_network_security_gateway_security_policies" "default" {
  name        = "%s"
  location    = "us-central1"
  description = "gateway security policy created to be used as reference by the rule."
}
	
resource "google_network_security_gateway_security_policies_rule" "foobar" {
  name                    = "%s"
  location                = "us-central1"
  gateway_security_policy = google_network_security_gateway_security_policies.default.name
  enabled                 = true  
  description             = "my description"
  priority                = 0
  session_matcher         = "host() == 'example.com'"
  application_matcher     = "request.method == 'POST'"
  basic_profile           = "ALLOW"
}
`, gatewaySecurityPolicyName, gatewaySecurityPolicyRuleName)
}

func testAccNetworkSecurityGatewaySecurityPoliciesRule_update(gatewaySecurityPolicyName, gatewaySecurityPolicyRuleName string) string {
	return fmt.Sprintf(`
resource "google_network_security_gateway_security_policies" "default" {
  name        = "%s"
  location    = "us-central1"
  description = "gateway security policy created to be used as reference by the rule."
}
	
resource "google_network_security_gateway_security_policies_rule" "foobar" {
  name                    = "%s"
  location                = "us-central1"
  gateway_security_policy = google_network_security_gateway_security_policies.default.name
  enabled                 = false  
  description             = "my description updated"
  priority                = 1
  session_matcher         = "host() == 'update.com'"
  application_matcher     = "request.method == 'GET'"
  tls_inspection_enabled  = false
  basic_profile           = "DENY"
}
`, gatewaySecurityPolicyName, gatewaySecurityPolicyRuleName)
}
