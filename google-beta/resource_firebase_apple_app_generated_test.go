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

package google

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

func TestAccFirebaseAppleApp_firebaseAppleAppBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"project_id":    envvar.GetTestProjectFromEnv(),
		"display_name":  "tf-test Display Name Basic",
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckFirebaseAppleAppDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAppleApp_firebaseAppleAppBasicExample(context),
			},
			{
				ResourceName:      "google_firebase_apple_app.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccFirebaseAppleApp_firebaseAppleAppBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_apple_app" "default" {
  provider = google-beta
  project = "%{project_id}"
  display_name = "%{display_name}"
  bundle_id = "apple.app.12345%{random_suffix}"
}
`, context)
}

func TestAccFirebaseAppleApp_firebaseAppleAppFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"project_id":    envvar.GetTestProjectFromEnv(),
		"app_store_id":  12345,
		"team_id":       9987654321,
		"display_name":  "tf-test Display Name Full",
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckFirebaseAppleAppDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAppleApp_firebaseAppleAppFullExample(context),
			},
			{
				ResourceName:            "google_firebase_apple_app.full",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project", "deletion_policy"},
			},
		},
	})
}

func testAccFirebaseAppleApp_firebaseAppleAppFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_apple_app" "full" {
  provider = google-beta
  project = "%{project_id}"
  display_name = "%{display_name}"
  bundle_id = "apple.app.12345%{random_suffix}"
  app_store_id = "%{app_store_id}"
  team_id = "%{team_id}"
}
`, context)
}

func testAccCheckFirebaseAppleAppDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_firebase_apple_app" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{FirebaseBasePath}}projects/{{project}}/iosApps/{{app_id}}")
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
				return fmt.Errorf("FirebaseAppleApp still exists at %s", url)
			}
		}

		return nil
	}
}
