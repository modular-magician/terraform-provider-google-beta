// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package securitycenterv2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"google.golang.org/api/googleapi"
)

func TestAccSecurityCenterOrganizationMuteConfig_createUpdateDelete(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
		"location":      "global",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck: func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecurityCenterOrganizationMuteConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityCenterOrganizationMuteConfig(context),
			},
			{
				ResourceName:            "google_scc_v2_organization_mute_config.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"organizationmute_ConfigId", "parent"},
			},
			{
				Config: testAccSecurityCenterOrganizationMuteConfig_update(context),
			},
			{
				ResourceName:            "google_scc_v2_organization_mute_config.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"organizationmute_ConfigId", "parent"},
			},
			{
				Config: testAccSecurityCenterOrganizationMuteConfig_delete(context),
			},
		},
	})
}

func testAccSecurityCenterOrganizationMuteConfig(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_scc_v2_organization_mute_config" "primary" {
  mute_config_id = "tf-test-my-config-%{random_suffix}"
  parent         = "organizations/%{org_id}"
  filter         = "category: \"OS_VULNERABILITY\""
  location       = "%{location}"
  description    = "My Mute Config"
  type           = "STATIC"
}
`, context)
}

func testAccSecurityCenterOrganizationMuteConfig_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_scc_v2_organization_mute_config" "primary" {
  mute_config_id = "tf-test-my-config-%{random_suffix}"
  parent         = "organizations/%{org_id}"
  filter         = "category: \"WEB_SECURITY\""
  location       = "%{location}"
  description    = "My Mute Config Updated"
  type           = "STATIC"
}
`, context)
}

func testAccSecurityCenterOrganizationMuteConfig_delete(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_scc_v2_organization_mute_config" "primary" {
  mute_config_id = "tf-test-my-config-%{random_suffix}"
  parent         = "organizations/%{org_id}"
  filter         = "category: \"OS_VULNERABILITY\""
  location       = "%{location}"
  description    = "My Mute Config"
  type           = "STATIC"
}
`, context)
}

func testAccCheckSecurityCenterOrganizationMuteConfigDestroy(s *terraform.State) error {
		config := acctest.Provider.Meta().(*google.Config)
		client := config.SecurityCenterClientV2
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "google_scc_v2_organization_mute_config" {
				continue
			}

			_, err := client.Organizations.MuteConfigs.Get(rs.Primary.ID).Do()
			if err == nil {
				return fmt.Errorf("Organization Mute Config %s still exists", rs.Primary.ID)
			}

			if !isGoogleAPIErrorWithCode(err, 404) {
				return fmt.Errorf("error fetching Organization Mute Config: %s", err)
			}
		}

		return nil
}

func isGoogleAPIErrorWithCode(err error, code int) bool {
	apiErr, ok := err.(*googleapi.Error)
	return ok && apiErr.Code == code
}
