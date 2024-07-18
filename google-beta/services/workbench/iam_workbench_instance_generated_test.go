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

package workbench_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccWorkbenchInstanceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkbenchInstanceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_workbench_instance_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s roles/viewer", envvar.GetTestProjectFromEnv(), "us-west1-a", fmt.Sprintf("tf-test-workbench-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccWorkbenchInstanceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_workbench_instance_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s roles/viewer", envvar.GetTestProjectFromEnv(), "us-west1-a", fmt.Sprintf("tf-test-workbench-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccWorkbenchInstanceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccWorkbenchInstanceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_workbench_instance_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), "us-west1-a", fmt.Sprintf("tf-test-workbench-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccWorkbenchInstanceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkbenchInstanceIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_workbench_instance_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_workbench_instance_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s", envvar.GetTestProjectFromEnv(), "us-west1-a", fmt.Sprintf("tf-test-workbench-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccWorkbenchInstanceIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_workbench_instance_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s", envvar.GetTestProjectFromEnv(), "us-west1-a", fmt.Sprintf("tf-test-workbench-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccWorkbenchInstanceIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_workbench_instance" "instance" {
  name = "tf-test-workbench-instance%{random_suffix}"
  location = "us-west1-a"
}

resource "google_workbench_instance_iam_member" "foo" {
  project = google_workbench_instance.instance.project
  location = google_workbench_instance.instance.location
  name = google_workbench_instance.instance.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccWorkbenchInstanceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_workbench_instance" "instance" {
  name = "tf-test-workbench-instance%{random_suffix}"
  location = "us-west1-a"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_workbench_instance_iam_policy" "foo" {
  project = google_workbench_instance.instance.project
  location = google_workbench_instance.instance.location
  name = google_workbench_instance.instance.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_workbench_instance_iam_policy" "foo" {
  project = google_workbench_instance.instance.project
  location = google_workbench_instance.instance.location
  name = google_workbench_instance.instance.name
  depends_on = [
    google_workbench_instance_iam_policy.foo
  ]
}
`, context)
}

func testAccWorkbenchInstanceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_workbench_instance" "instance" {
  name = "tf-test-workbench-instance%{random_suffix}"
  location = "us-west1-a"
}

data "google_iam_policy" "foo" {
}

resource "google_workbench_instance_iam_policy" "foo" {
  project = google_workbench_instance.instance.project
  location = google_workbench_instance.instance.location
  name = google_workbench_instance.instance.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccWorkbenchInstanceIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_workbench_instance" "instance" {
  name = "tf-test-workbench-instance%{random_suffix}"
  location = "us-west1-a"
}

resource "google_workbench_instance_iam_binding" "foo" {
  project = google_workbench_instance.instance.project
  location = google_workbench_instance.instance.location
  name = google_workbench_instance.instance.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccWorkbenchInstanceIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_workbench_instance" "instance" {
  name = "tf-test-workbench-instance%{random_suffix}"
  location = "us-west1-a"
}

resource "google_workbench_instance_iam_binding" "foo" {
  project = google_workbench_instance.instance.project
  location = google_workbench_instance.instance.location
  name = google_workbench_instance.instance.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
