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

func TestAccCloudIdentityGroupMembership_cloudIdentityGroupMembershipExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_domain":    getTestOrgDomainFromEnv(t),
		"cust_id":       getTestCustIdFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckCloudIdentityGroupMembershipDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudIdentityGroupMembership_cloudIdentityGroupMembershipExample(context),
			},
		},
	})
}

func testAccCloudIdentityGroupMembership_cloudIdentityGroupMembershipExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_identity_group" "group" {
  provider = google-beta
  display_name = "tf-test-my-identity-group%{random_suffix}"

  parent = "customers/%{cust_id}"

  group_key {
  	id = "tf-test-my-identity-group%{random_suffix}@%{org_domain}"
  }

  labels = {
    "cloudidentity.googleapis.com/groups.discussion_forum" = ""
  }
}

resource "google_cloud_identity_group" "child-group" {
  provider = google-beta
  display_name = "tf-test-my-identity-group%{random_suffix}-child"

  parent = "customers/%{cust_id}"

  group_key {
  	id = "tf-test-my-identity-group%{random_suffix}-child@%{org_domain}"
  }

  labels = {
    "cloudidentity.googleapis.com/groups.discussion_forum" = ""
  }
}

resource "google_cloud_identity_group_membership" "cloud_identity_group_membership_basic" {
  provider = google-beta
  group    = google_cloud_identity_group.group.id

  member_key {
    id = google_cloud_identity_group.child-group.group_key[0].id
  }

  roles {
  	name = "MEMBER"
  }
}
`, context)
}

func TestAccCloudIdentityGroupMembership_cloudIdentityGroupMembershipUserExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_domain":    getTestOrgDomainFromEnv(t),
		"cust_id":       getTestCustIdFromEnv(t),
		"identity_user": getTestIdentityUserFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckCloudIdentityGroupMembershipDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudIdentityGroupMembership_cloudIdentityGroupMembershipUserExample(context),
			},
		},
	})
}

func testAccCloudIdentityGroupMembership_cloudIdentityGroupMembershipUserExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_identity_group" "group" {
  provider = google-beta
  display_name = "tf-test-my-identity-group%{random_suffix}"

  parent = "customers/%{cust_id}"

  group_key {
    id = "tf-test-my-identity-group%{random_suffix}@%{org_domain}"
  }

  labels = {
    "cloudidentity.googleapis.com/groups.discussion_forum" = ""
  }
}

resource "google_cloud_identity_group_membership" "cloud_identity_group_membership_basic" {
  provider = google-beta
  group    = google_cloud_identity_group.group.id

  member_key {
    id = "%{identity_user}@%{org_domain}"
  }

  roles {
    name = "MEMBER"
  }

  roles {
    name = "MANAGER"
  }
}
`, context)
}

func testAccCheckCloudIdentityGroupMembershipDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloud_identity_group_membership" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{CloudIdentityBasePath}}{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("CloudIdentityGroupMembership still exists at %s", url)
			}
		}

		return nil
	}
}
