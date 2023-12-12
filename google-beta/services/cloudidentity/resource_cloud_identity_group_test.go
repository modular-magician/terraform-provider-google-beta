// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package cloudidentity_test

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

// Intended to fix https://github.com/hashicorp/terraform-provider-google/issues/10001
// These are all of the tests that use a cloud_identity_group, except for
// testAccAccessContextManagerGcpUserAccessBinding_basicTest. The theory is that they sometimes
// fail with a 409 because of concurrent roster mutations, so running them serially should prevent
// the error.
func TestAccCloudIdentityGroup(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"basic":                           testAccCloudIdentityGroup_cloudIdentityGroupsBasicExampleTest,
		"update":                          testAccCloudIdentityGroup_updateTest,
		"membership_basic":                testAccCloudIdentityGroupMembership_cloudIdentityGroupMembershipExampleTest,
		"membership_update":               testAccCloudIdentityGroupMembership_updateTest,
		"membership_import":               testAccCloudIdentityGroupMembership_importTest,
		"membership_dne":                  testAccCloudIdentityGroupMembership_membershipDoesNotExistTest,
		"membership_user":                 testAccCloudIdentityGroupMembership_cloudIdentityGroupMembershipUserExampleTest,
		"membership_with_member_key":      testAccCloudIdentityGroupMembership_cloudIdentityGroupMembershipWithMemberKeyTest,
		"membership_user_with_member_key": testAccCloudIdentityGroupMembership_cloudIdentityGroupMembershipUserWithMemberKeyTest,
		"data_source_basic":               testAccDataSourceCloudIdentityGroups_basicTest,
		"data_source_membership_basic":    testAccDataSourceCloudIdentityGroupMemberships_basicTest,
		"data_source_group_lookup":        testAccDataSourceCloudIdentityGroupLookup_basicTest,
	}

	for name, tc := range testCases {
		// shadow the tc variable into scope so that when
		// the loop continues, if t.Run hasn't executed tc(t)
		// yet, we don't have a race condition
		// see https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		tc := tc
		t.Run(name, func(t *testing.T) {
			tc(t)
		})
	}
}

func testAccCloudIdentityGroup_updateTest(t *testing.T) {
	context := map[string]interface{}{
		"org_domain":    envvar.GetTestOrgDomainFromEnv(t),
		"cust_id":       envvar.GetTestCustIdFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudIdentityGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudIdentityGroup_cloudIdentityGroupsBasicExample(context),
			},
			{
				Config: testAccCloudIdentityGroup_update(context),
			},
		},
	})
}

func testAccCloudIdentityGroup_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_identity_group" "cloud_identity_group_basic" {
  display_name = "tf-test-my-identity-group%{random_suffix}-update"
  description  = "my-description"

  parent = "customers/%{cust_id}"

  group_key {
    id = "tf-test-my-identity-group%{random_suffix}@%{org_domain}"
  }

  labels = {
    "cloudidentity.googleapis.com/groups.discussion_forum" = ""
    "cloudidentity.googleapis.com/groups.security" = ""
  }
}
`, context)
}

func testAccCloudIdentityGroup_cloudIdentityGroupsBasicExampleTest(t *testing.T) {
	context := map[string]interface{}{
		"org_domain":    envvar.GetTestOrgDomainFromEnv(t),
		"cust_id":       envvar.GetTestCustIdFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudIdentityGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudIdentityGroup_cloudIdentityGroupsBasicExample(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("google_cloud_identity_group.cloud_identity_group_basic",
						"additional_group_keys.#"),
				),
			},
			{
				ResourceName:            "google_cloud_identity_group.cloud_identity_group_basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"initial_group_config"},
			},
		},
	})
}

func testAccCloudIdentityGroup_cloudIdentityGroupsBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_identity_group" "cloud_identity_group_basic" {
  display_name         = "tf-test-my-identity-group%{random_suffix}"
  initial_group_config = "WITH_INITIAL_OWNER"

  parent = "customers/%{cust_id}"

  group_key {
    id = "tf-test-my-identity-group%{random_suffix}@%{org_domain}"
  }

  labels = {
    "cloudidentity.googleapis.com/groups.discussion_forum" = ""
  }
}
`, context)
}

func testAccCheckCloudIdentityGroupDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloud_identity_group" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{CloudIdentityBasePath}}{{name}}")
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
				return fmt.Errorf("CloudIdentityGroup still exists at %s", url)
			}
		}

		return nil
	}
}

func TestAccCloudIdentityGroup_cloudIdentityGroupsDynamicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_domain":    envvar.GetTestOrgDomainFromEnv(t),
		"cust_id":       envvar.GetTestCustIdFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudIdentityGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config:             testAccCloudIdentityGroup_cloudIdentityGroupsDynamicExample(context),
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName: "google_cloud_identity_group.cloud_identity_groups_dynamic",
				ImportState:  true,
			},
			{
				Config: testAccCloudIdentityGroup_cloudIdentityGroupsDynamicUpdate(context),
			},
			{
				ResourceName:      "google_cloud_identity_group.cloud_identity_groups_dynamic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudIdentityGroup_cloudIdentityGroupsDynamicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_identity_group" "cloud_identity_groups_dynamic" {
  display_name         = "tf-test-my-identity-dynamic-group%{random_suffix}"

  parent = "customers/%{cust_id}"

  group_key {
    id = "tf-test-my-identity-dynamic-group%{random_suffix}@%{org_domain}"
  }

  labels = {
    "cloudidentity.googleapis.com/groups.discussion_forum" = ""
  }

  dynamic_group_metadata {
    queries {
      resource_type = "USER"
      query         = "user.addresses.exists(ad, ad.locality=='Seattle')"
    }
  }
}
`, context)
}

func testAccCloudIdentityGroup_cloudIdentityGroupsDynamicUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_identity_group" "cloud_identity_groups_dynamic" {
  display_name         = "tf-test-my-identity-dynamic-group%{random_suffix}"

  parent = "customers/%{cust_id}"

  group_key {
    id = "tf-test-my-identity-dynamic-group%{random_suffix}@%{org_domain}"
  }

  labels = {
    "cloudidentity.googleapis.com/groups.discussion_forum" = ""
  }

  dynamic_group_metadata {
    queries {
      resource_type = "USER"
      query         = "user.addresses.exists(ad, ad.locality=='Seattle')"
    }
  }
}
`, context)
}
