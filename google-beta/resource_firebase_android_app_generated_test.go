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

func TestAccFirebaseAndroidApp_firebaseAndroidAppBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"project_id":    envvar.GetTestProjectFromEnv(),
		"package_name":  "android.package.app" + RandString(t, 4),
		"display_name":  "tf-test Display Name Basic",
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckFirebaseAndroidAppDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAndroidApp_firebaseAndroidAppBasicExample(context),
			},
			{
				ResourceName:            "google_firebase_android_app.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project", "deletion_policy"},
			},
		},
	})
}

func testAccFirebaseAndroidApp_firebaseAndroidAppBasicExample(context map[string]interface{}) string {
	return tpgresource.Nprintf(`
resource "google_firebase_android_app" "basic" {
  provider = google-beta
  project = "%{project_id}"
  display_name = "%{display_name}"
  package_name = "%{package_name}"
  sha1_hashes = ["2145bdf698b8715039bd0e83f2069bed435ac21c"]
  sha256_hashes = ["2145bdf698b8715039bd0e83f2069bed435ac21ca1b2c3d4e5f6123456789abc"]
}
`, context)
}

func testAccCheckFirebaseAndroidAppDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_firebase_android_app" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{FirebaseBasePath}}projects/{{project}}/androidApps/{{app_id}}")
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
				return fmt.Errorf("FirebaseAndroidApp still exists at %s", url)
			}
		}

		return nil
	}
}
