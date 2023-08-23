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

package identityplatform_test

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccIdentityPlatformConfig_identityPlatformConfigBasicExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"org_id":           envvar.GetTestOrgFromEnv(t),
		"billing_acct":     envvar.GetTestBillingAccountFromEnv(t),
		"quota_start_time": time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
		"random_suffix":    acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIdentityPlatformConfig_identityPlatformConfigBasicExample(context),
			},
			{
				ResourceName:      "google_identity_platform_config.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIdentityPlatformConfig_identityPlatformConfigBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "default" {
  project_id = "tf-test-my-project%{random_suffix}"
  name       = "tf-test-my-project%{random_suffix}"
  org_id     = "%{org_id}"
  billing_account =  "%{billing_acct}"
  labels = {
    firebase = "enabled"
  }
}

resource "google_project_service" "identitytoolkit" {
  project = google_project.default.project_id
  service = "identitytoolkit.googleapis.com"
}


resource "google_identity_platform_config" "default" {
  project = google_project.default.project_id
  autodelete_anonymous_users = true
  blocking_functions {
    triggers {
      event_type = "beforeSignIn"
      function_uri = "https://us-east1-tf-test-my-project%{random_suffix}.cloudfunctions.net/before-sign-in"
    }
    forward_inbound_credentials {
      refresh_token = true
      access_token = true
      id_token = true
    }
  }
  quota {
    sign_up_quota_config {
      quota = 1000
      start_time = "%{quota_start_time}"
      quota_duration = "7200s"
    }
  }
  authorized_domains = [
    "localhost",
    "tf-test-my-project%{random_suffix}.firebaseapp.com",
    "tf-test-my-project%{random_suffix}.web.app",
  ]
}
`, context)
}

func TestAccIdentityPlatformConfig_identityPlatformConfigMinimalExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"billing_acct":  envvar.GetTestBillingAccountFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIdentityPlatformConfig_identityPlatformConfigMinimalExample(context),
			},
			{
				ResourceName:      "google_identity_platform_config.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIdentityPlatformConfig_identityPlatformConfigMinimalExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "default" {
  project_id = "tf-test-my-project-1%{random_suffix}"
  name       = "tf-test-my-project-1%{random_suffix}"
  org_id     = "%{org_id}"
  billing_account =  "%{billing_acct}"
  labels = {
    firebase = "enabled"
  }
}

resource "google_project_service" "identitytoolkit" {
  project = google_project.default.project_id
  service = "identitytoolkit.googleapis.com"
}


resource "google_identity_platform_config" "default" {
  project = google_project.default.project_id
}
`, context)
}
