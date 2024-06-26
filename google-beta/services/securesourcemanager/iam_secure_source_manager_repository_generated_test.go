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

package securesourcemanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccSecureSourceManagerRepositoryIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/securesourcemanager.repoAdmin",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecureSourceManagerRepositoryIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_secure_source_manager_repository_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/repositories/%s roles/securesourcemanager.repoAdmin", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccSecureSourceManagerRepositoryIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_secure_source_manager_repository_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/repositories/%s roles/securesourcemanager.repoAdmin", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSecureSourceManagerRepositoryIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/securesourcemanager.repoAdmin",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccSecureSourceManagerRepositoryIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_secure_source_manager_repository_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/repositories/%s roles/securesourcemanager.repoAdmin user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSecureSourceManagerRepositoryIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/securesourcemanager.repoAdmin",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecureSourceManagerRepositoryIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_secure_source_manager_repository_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_secure_source_manager_repository_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/repositories/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSecureSourceManagerRepositoryIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_secure_source_manager_repository_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/repositories/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSecureSourceManagerRepositoryIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secure_source_manager_instance" "instance" {
    location = "us-central1"
    instance_id = "tf-test-my-instance%{random_suffix}"
}

resource "google_secure_source_manager_repository" "default" {
    location = "us-central1"
    repository_id = "tf-test-my-repository%{random_suffix}"
    instance = google_secure_source_manager_instance.instance.name
}

resource "google_secure_source_manager_repository_iam_member" "foo" {
  project = google_secure_source_manager_repository.default.project
  location = google_secure_source_manager_repository.default.location
  repository_id = google_secure_source_manager_repository.default.repository_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccSecureSourceManagerRepositoryIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secure_source_manager_instance" "instance" {
    location = "us-central1"
    instance_id = "tf-test-my-instance%{random_suffix}"
}

resource "google_secure_source_manager_repository" "default" {
    location = "us-central1"
    repository_id = "tf-test-my-repository%{random_suffix}"
    instance = google_secure_source_manager_instance.instance.name
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_secure_source_manager_repository_iam_policy" "foo" {
  project = google_secure_source_manager_repository.default.project
  location = google_secure_source_manager_repository.default.location
  repository_id = google_secure_source_manager_repository.default.repository_id
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_secure_source_manager_repository_iam_policy" "foo" {
  project = google_secure_source_manager_repository.default.project
  location = google_secure_source_manager_repository.default.location
  repository_id = google_secure_source_manager_repository.default.repository_id
  depends_on = [
    google_secure_source_manager_repository_iam_policy.foo
  ]
}
`, context)
}

func testAccSecureSourceManagerRepositoryIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secure_source_manager_instance" "instance" {
    location = "us-central1"
    instance_id = "tf-test-my-instance%{random_suffix}"
}

resource "google_secure_source_manager_repository" "default" {
    location = "us-central1"
    repository_id = "tf-test-my-repository%{random_suffix}"
    instance = google_secure_source_manager_instance.instance.name
}

data "google_iam_policy" "foo" {
}

resource "google_secure_source_manager_repository_iam_policy" "foo" {
  project = google_secure_source_manager_repository.default.project
  location = google_secure_source_manager_repository.default.location
  repository_id = google_secure_source_manager_repository.default.repository_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccSecureSourceManagerRepositoryIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secure_source_manager_instance" "instance" {
    location = "us-central1"
    instance_id = "tf-test-my-instance%{random_suffix}"
}

resource "google_secure_source_manager_repository" "default" {
    location = "us-central1"
    repository_id = "tf-test-my-repository%{random_suffix}"
    instance = google_secure_source_manager_instance.instance.name
}

resource "google_secure_source_manager_repository_iam_binding" "foo" {
  project = google_secure_source_manager_repository.default.project
  location = google_secure_source_manager_repository.default.location
  repository_id = google_secure_source_manager_repository.default.repository_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccSecureSourceManagerRepositoryIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secure_source_manager_instance" "instance" {
    location = "us-central1"
    instance_id = "tf-test-my-instance%{random_suffix}"
}

resource "google_secure_source_manager_repository" "default" {
    location = "us-central1"
    repository_id = "tf-test-my-repository%{random_suffix}"
    instance = google_secure_source_manager_instance.instance.name
}

resource "google_secure_source_manager_repository_iam_binding" "foo" {
  project = google_secure_source_manager_repository.default.project
  location = google_secure_source_manager_repository.default.location
  repository_id = google_secure_source_manager_repository.default.repository_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
