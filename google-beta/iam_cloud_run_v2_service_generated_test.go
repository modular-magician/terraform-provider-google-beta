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

func TestAccCloudRunV2ServiceIamBindingGenerated(t *testing.T) {
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
				Config: testAccCloudRunV2ServiceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_cloud_run_v2_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/services/%s roles/viewer", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudrun-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccCloudRunV2ServiceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_cloud_run_v2_service_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/services/%s roles/viewer", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudrun-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudRunV2ServiceIamMemberGenerated(t *testing.T) {
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
				Config: testAccCloudRunV2ServiceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_cloud_run_v2_service_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/services/%s roles/viewer user:admin@hashicorptest.com", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudrun-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudRunV2ServiceIamPolicyGenerated(t *testing.T) {
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
				Config: testAccCloudRunV2ServiceIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_cloud_run_v2_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/services/%s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudrun-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCloudRunV2ServiceIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_cloud_run_v2_service_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/services/%s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudrun-service%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudRunV2ServiceIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_v2_service" "default" {
  name     = "tf-test-cloudrun-service%{random_suffix}"
  location = "us-central1"
  ingress = "INGRESS_TRAFFIC_ALL"
  
  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
  }
}

resource "google_cloud_run_v2_service_iam_member" "foo" {
  project = google_cloud_run_v2_service.default.project
  location = google_cloud_run_v2_service.default.location
  name = google_cloud_run_v2_service.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccCloudRunV2ServiceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_v2_service" "default" {
  name     = "tf-test-cloudrun-service%{random_suffix}"
  location = "us-central1"
  ingress = "INGRESS_TRAFFIC_ALL"
  
  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_cloud_run_v2_service_iam_policy" "foo" {
  project = google_cloud_run_v2_service.default.project
  location = google_cloud_run_v2_service.default.location
  name = google_cloud_run_v2_service.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccCloudRunV2ServiceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_v2_service" "default" {
  name     = "tf-test-cloudrun-service%{random_suffix}"
  location = "us-central1"
  ingress = "INGRESS_TRAFFIC_ALL"
  
  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
  }
}

data "google_iam_policy" "foo" {
}

resource "google_cloud_run_v2_service_iam_policy" "foo" {
  project = google_cloud_run_v2_service.default.project
  location = google_cloud_run_v2_service.default.location
  name = google_cloud_run_v2_service.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccCloudRunV2ServiceIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_v2_service" "default" {
  name     = "tf-test-cloudrun-service%{random_suffix}"
  location = "us-central1"
  ingress = "INGRESS_TRAFFIC_ALL"
  
  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
  }
}

resource "google_cloud_run_v2_service_iam_binding" "foo" {
  project = google_cloud_run_v2_service.default.project
  location = google_cloud_run_v2_service.default.location
  name = google_cloud_run_v2_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccCloudRunV2ServiceIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_v2_service" "default" {
  name     = "tf-test-cloudrun-service%{random_suffix}"
  location = "us-central1"
  ingress = "INGRESS_TRAFFIC_ALL"
  
  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
  }
}

resource "google_cloud_run_v2_service_iam_binding" "foo" {
  project = google_cloud_run_v2_service.default.project
  location = google_cloud_run_v2_service.default.location
  name = google_cloud_run_v2_service.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
