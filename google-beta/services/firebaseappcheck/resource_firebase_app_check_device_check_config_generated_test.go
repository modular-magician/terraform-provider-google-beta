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

package firebaseappcheck_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccFirebaseAppCheckDeviceCheckConfig_firebaseAppCheckDeviceCheckConfigFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":       envvar.GetTestProjectFromEnv(),
		"private_key_path": "test-fixtures/private-key-2.p8",
		"team_id":          "9987654321",
		"token_ttl":        "7200s",
		"random_suffix":    acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAppCheckDeviceCheckConfig_firebaseAppCheckDeviceCheckConfigFullExample(context),
			},
			{
				ResourceName:            "google_firebase_app_check_device_check_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"private_key", "app_id"},
			},
		},
	})
}

func testAccFirebaseAppCheckDeviceCheckConfig_firebaseAppCheckDeviceCheckConfigFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_apple_app" "default" {
  provider = google-beta

  project      = "%{project_id}"
  display_name = "Apple app"
  bundle_id    = "bundle.id.devicecheck%{random_suffix}"
  team_id      = "%{team_id}"
}

# It takes a while for App Check to recognize the new app
# If your app already exists, you don't have to wait 30 seconds.
resource "time_sleep" "wait_30s" {
  depends_on      = [google_firebase_apple_app.default]
  create_duration = "30s"
}

resource "google_firebase_app_check_device_check_config" "default" {
  provider = google-beta

  project     = "%{project_id}"
  app_id      = google_firebase_apple_app.default.app_id
  token_ttl   = "%{token_ttl}"
  key_id      = "Key ID%{random_suffix}"
  private_key = file("%{private_key_path}")

  depends_on = [time_sleep.wait_30s]

  lifecycle {
    precondition {
      condition     = google_firebase_apple_app.default.team_id != ""
      error_message = "Provide a Team ID on the Apple App to use App Check"
    }
  }
}
`, context)
}
