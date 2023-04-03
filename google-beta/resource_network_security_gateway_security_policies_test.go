package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkSecurityGatewaySecurityPolicies_update(t *testing.T) {
	t.Parallel()

	gatewaySecurityPolicyName := fmt.Sprintf("tf-test-gateway-sp-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityGatewaySecurityPoliciesDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityGatewaySecurityPolicies_basic(gatewaySecurityPolicyName),
			},
			{
				ResourceName:      "google_network_security_gateway_security_policies.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkSecurityGatewaySecurityPolicies_update(gatewaySecurityPolicyName),
			},
			{
				ResourceName:      "google_network_security_gateway_security_policies.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkSecurityGatewaySecurityPolicies_basic(gatewaySecurityPolicyName string) string {
	return fmt.Sprintf(`
resource "google_network_security_gateway_security_policies" "foobar" {
  name        = "%s"
  location    = "us-central1"
  description = "my description"
}
`, gatewaySecurityPolicyName)
}

func testAccNetworkSecurityGatewaySecurityPolicies_update(gatewaySecurityPolicyName string) string {
	return fmt.Sprintf(`
resource "google_network_security_gateway_security_policies" "foobar" {
  name        = "%s"
  location    = "us-central1"
  description = "update description"
}
`, gatewaySecurityPolicyName)
}
