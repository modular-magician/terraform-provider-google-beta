// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

package iam2_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccIAM2DenyPolicy_iamDenyPolicyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAM2DenyPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAM2DenyPolicy_iamDenyPolicyBasicExample(context),
			},
			{
				ResourceName:            "google_iam_deny_policy.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "parent"},
			},
		},
	})
}

func testAccIAM2DenyPolicy_iamDenyPolicyBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "project" {
  project_id      = "tf-test-my-project%{random_suffix}"
  name            = "tf-test-my-project%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
  deletion_policy = "NONE"
}

resource "google_iam_deny_policy" "example" {
  parent   = urlencode("cloudresourcemanager.googleapis.com/projects/${google_project.project.project_id}")
  name     = "tf-test-my-deny-policy%{random_suffix}"
  display_name = "A deny rule"
  rules {
    description = "First rule"
    deny_rule {
      denied_principals = ["principalSet://goog/public:all"]
      denial_condition {
        title = "Some expr"
        expression = "!resource.matchTag('12345678/env', 'test')"
      }
      denied_permissions = ["cloudresourcemanager.googleapis.com/projects.update"]
    }
  }
  rules {
    description = "Second rule"
    deny_rule {
      denied_principals = ["principalSet://goog/public:all"]
      denial_condition {
        title = "Some expr"
        expression = "!resource.matchTag('12345678/env', 'test')"
      }
      denied_permissions = ["cloudresourcemanager.googleapis.com/projects.update"]
      exception_principals = ["principal://iam.googleapis.com/projects/-/serviceAccounts/${google_service_account.test-account.email}"]
    }
  }
}

resource "google_service_account" "test-account" {
  account_id   = "tf-test-svc-acc%{random_suffix}"
  display_name = "Test Service Account"
  project      = google_project.project.project_id
}
`, context)
}

func testAccCheckIAM2DenyPolicyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_iam_deny_policy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{IAM2BasePath}}policies/{{parent}}/denypolicies/{{name}}")
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
				return fmt.Errorf("IAM2DenyPolicy still exists at %s", url)
			}
		}

		return nil
	}
}
