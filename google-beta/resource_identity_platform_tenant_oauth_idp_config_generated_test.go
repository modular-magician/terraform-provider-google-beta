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

func TestAccIdentityPlatformTenantOauthIdpConfig_identityPlatformTenantOauthIdpConfigBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"name":          "oidc.oauth-idp-config-" + RandString(t, 10),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckIdentityPlatformTenantOauthIdpConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIdentityPlatformTenantOauthIdpConfig_identityPlatformTenantOauthIdpConfigBasicExample(context),
			},
			{
				ResourceName:            "google_identity_platform_tenant_oauth_idp_config.tenant_oauth_idp_config",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"tenant"},
			},
		},
	})
}

func testAccIdentityPlatformTenantOauthIdpConfig_identityPlatformTenantOauthIdpConfigBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_identity_platform_tenant" "tenant" {
  display_name  = "tenant"
}

resource "google_identity_platform_tenant_oauth_idp_config" "tenant_oauth_idp_config" {
  name          = "%{name}"
  tenant        = google_identity_platform_tenant.tenant.name
  display_name  = "Display Name"
  client_id     = "client-id"
  issuer        = "issuer"
  enabled       = true
  client_secret = "secret"
}
`, context)
}

func testAccCheckIdentityPlatformTenantOauthIdpConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_identity_platform_tenant_oauth_idp_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("IdentityPlatformTenantOauthIdpConfig still exists at %s", url)
			}
		}

		return nil
	}
}
