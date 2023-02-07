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
)

func TestAccServiceDirectoryNamespaceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccServiceDirectoryNamespaceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_service_directory_namespace_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/namespaces/%s roles/viewer", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-example-namespace%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccServiceDirectoryNamespaceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_service_directory_namespace_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/namespaces/%s roles/viewer", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-example-namespace%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccServiceDirectoryNamespaceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccServiceDirectoryNamespaceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_service_directory_namespace_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/namespaces/%s roles/viewer user:admin@hashicorptest.com", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-example-namespace%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccServiceDirectoryNamespaceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccServiceDirectoryNamespaceIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_service_directory_namespace_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-example-namespace%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccServiceDirectoryNamespaceIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_service_directory_namespace_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-example-namespace%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccServiceDirectoryNamespaceIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_service_directory_namespace" "example" {
  provider     = google-beta
  namespace_id = "tf-test-example-namespace%{random_suffix}"
  location     = "us-central1"

  labels = {
    key = "value"
    foo = "bar"
  }
}

resource "google_service_directory_namespace_iam_member" "foo" {
  provider = google-beta
  namespace_id = google_service_directory_namespace.example.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccServiceDirectoryNamespaceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_service_directory_namespace" "example" {
  provider     = google-beta
  namespace_id = "tf-test-example-namespace%{random_suffix}"
  location     = "us-central1"

  labels = {
    key = "value"
    foo = "bar"
  }
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_service_directory_namespace_iam_policy" "foo" {
  provider = google-beta
  namespace_id = google_service_directory_namespace.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccServiceDirectoryNamespaceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_service_directory_namespace" "example" {
  provider     = google-beta
  namespace_id = "tf-test-example-namespace%{random_suffix}"
  location     = "us-central1"

  labels = {
    key = "value"
    foo = "bar"
  }
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_service_directory_namespace_iam_policy" "foo" {
  provider = google-beta
  namespace_id = google_service_directory_namespace.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccServiceDirectoryNamespaceIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_service_directory_namespace" "example" {
  provider     = google-beta
  namespace_id = "tf-test-example-namespace%{random_suffix}"
  location     = "us-central1"

  labels = {
    key = "value"
    foo = "bar"
  }
}

resource "google_service_directory_namespace_iam_binding" "foo" {
  provider = google-beta
  namespace_id = google_service_directory_namespace.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccServiceDirectoryNamespaceIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_service_directory_namespace" "example" {
  provider     = google-beta
  namespace_id = "tf-test-example-namespace%{random_suffix}"
  location     = "us-central1"

  labels = {
    key = "value"
    foo = "bar"
  }
}

resource "google_service_directory_namespace_iam_binding" "foo" {
  provider = google-beta
  namespace_id = google_service_directory_namespace.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
