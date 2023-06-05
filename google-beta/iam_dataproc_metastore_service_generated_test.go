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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccDataprocMetastoreServiceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocMetastoreServiceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataproc_metastore_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/services/%s roles/viewer", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-metastore-srv%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccDataprocMetastoreServiceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_dataproc_metastore_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/services/%s roles/viewer", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-metastore-srv%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataprocMetastoreServiceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccDataprocMetastoreServiceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataproc_metastore_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/services/%s roles/viewer user:admin@hashicorptest.com", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-metastore-srv%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataprocMetastoreServiceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocMetastoreServiceIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataproc_metastore_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/services/%s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-metastore-srv%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataprocMetastoreServiceIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_dataproc_metastore_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/services/%s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-metastore-srv%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataprocMetastoreServiceIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_metastore_service" "default" {
  service_id = "tf-test-metastore-srv%{random_suffix}"
  location   = "us-central1"
  port       = 9080
  tier       = "DEVELOPER"

  maintenance_window {
    hour_of_day = 2
    day_of_week = "SUNDAY"
  }

  hive_metastore_config {
    version = "2.3.6"
  }
}

resource "google_dataproc_metastore_service_iam_member" "foo" {
  project = google_dataproc_metastore_service.default.project
  location = google_dataproc_metastore_service.default.location
  service_id = google_dataproc_metastore_service.default.service_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccDataprocMetastoreServiceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_metastore_service" "default" {
  service_id = "tf-test-metastore-srv%{random_suffix}"
  location   = "us-central1"
  port       = 9080
  tier       = "DEVELOPER"

  maintenance_window {
    hour_of_day = 2
    day_of_week = "SUNDAY"
  }

  hive_metastore_config {
    version = "2.3.6"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_dataproc_metastore_service_iam_policy" "foo" {
  project = google_dataproc_metastore_service.default.project
  location = google_dataproc_metastore_service.default.location
  service_id = google_dataproc_metastore_service.default.service_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataprocMetastoreServiceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_metastore_service" "default" {
  service_id = "tf-test-metastore-srv%{random_suffix}"
  location   = "us-central1"
  port       = 9080
  tier       = "DEVELOPER"

  maintenance_window {
    hour_of_day = 2
    day_of_week = "SUNDAY"
  }

  hive_metastore_config {
    version = "2.3.6"
  }
}

data "google_iam_policy" "foo" {
}

resource "google_dataproc_metastore_service_iam_policy" "foo" {
  project = google_dataproc_metastore_service.default.project
  location = google_dataproc_metastore_service.default.location
  service_id = google_dataproc_metastore_service.default.service_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataprocMetastoreServiceIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_metastore_service" "default" {
  service_id = "tf-test-metastore-srv%{random_suffix}"
  location   = "us-central1"
  port       = 9080
  tier       = "DEVELOPER"

  maintenance_window {
    hour_of_day = 2
    day_of_week = "SUNDAY"
  }

  hive_metastore_config {
    version = "2.3.6"
  }
}

resource "google_dataproc_metastore_service_iam_binding" "foo" {
  project = google_dataproc_metastore_service.default.project
  location = google_dataproc_metastore_service.default.location
  service_id = google_dataproc_metastore_service.default.service_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccDataprocMetastoreServiceIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_metastore_service" "default" {
  service_id = "tf-test-metastore-srv%{random_suffix}"
  location   = "us-central1"
  port       = 9080
  tier       = "DEVELOPER"

  maintenance_window {
    hour_of_day = 2
    day_of_week = "SUNDAY"
  }

  hive_metastore_config {
    version = "2.3.6"
  }
}

resource "google_dataproc_metastore_service_iam_binding" "foo" {
  project = google_dataproc_metastore_service.default.project
  location = google_dataproc_metastore_service.default.location
  service_id = google_dataproc_metastore_service.default.service_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
