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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFirebaseProjectLocation_firebaseProjectLocationBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        GetTestOrgFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseProjectLocation_firebaseProjectLocationBasicExample(context),
			},
			{
				ResourceName:      "google_firebase_project_location.basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccFirebaseProjectLocation_firebaseProjectLocationBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "default" {
  provider = google-beta

  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"

  labels = {
    "firebase" = "enabled"
  }
}

resource "google_firebase_project" "default" {
  provider = google-beta
  project  = google_project.default.project_id
}

resource "google_firebase_project_location" "basic" {
	provider = google-beta
	project = google_firebase_project.default.project

	location_id = "us-central"
}
`, context)
}
