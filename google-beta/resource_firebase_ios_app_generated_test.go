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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFirebaseIosApp_firebaseIosAppBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseIosApp_firebaseIosAppBasicExample(context),
			},
		},
	})
}

func testAccFirebaseIosApp_firebaseIosAppBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "default" {
	provider = google-beta

	project_id = "tf-test%{random_suffix}"
	name       = "tf-test%{random_suffix}"
	org_id     = "%{org_id}"
}

resource "google_firebase_project" "default" {
	provider = google-beta
	project  = google_project.default.project_id
}

resource "google_firebase_ios_app" "basic" {
	provider = google-beta
	project = google_project.default.project_id
	display_name = "Display Name Basic%{random_suffix}"

	depends_on = [google_firebase_project.default]
}
`, context)
}
