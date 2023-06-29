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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccRuntimeConfigConfigIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRuntimeConfigConfigIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_runtimeconfig_config_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/configs/%s roles/viewer", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-my-config%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccRuntimeConfigConfigIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_runtimeconfig_config_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/configs/%s roles/viewer", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-my-config%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccRuntimeConfigConfigIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccRuntimeConfigConfigIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_runtimeconfig_config_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/configs/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-my-config%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccRuntimeConfigConfigIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRuntimeConfigConfigIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_runtimeconfig_config_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_runtimeconfig_config_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/configs/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-my-config%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccRuntimeConfigConfigIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_runtimeconfig_config_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/configs/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-my-config%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccRuntimeConfigConfigIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_runtimeconfig_config" "config" {
  provider    = google-beta
  name        = "tf-test-my-config%{random_suffix}"
  description = "Runtime configuration values for my service"
}

resource "google_runtimeconfig_config_iam_member" "foo" {
  provider = google-beta
  project = google_runtimeconfig_config.config.project
  config = google_runtimeconfig_config.config.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccRuntimeConfigConfigIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_runtimeconfig_config" "config" {
  provider    = google-beta
  name        = "tf-test-my-config%{random_suffix}"
  description = "Runtime configuration values for my service"
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_runtimeconfig_config_iam_policy" "foo" {
  provider = google-beta
  project = google_runtimeconfig_config.config.project
  config = google_runtimeconfig_config.config.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_runtimeconfig_config_iam_policy" "foo" {
  provider = google-beta
  project = google_runtimeconfig_config.config.project
  config = google_runtimeconfig_config.config.name
  depends_on = [
    google_runtimeconfig_config_iam_policy.foo
  ]
}
`, context)
}

func testAccRuntimeConfigConfigIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_runtimeconfig_config" "config" {
  provider    = google-beta
  name        = "tf-test-my-config%{random_suffix}"
  description = "Runtime configuration values for my service"
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_runtimeconfig_config_iam_policy" "foo" {
  provider = google-beta
  project = google_runtimeconfig_config.config.project
  config = google_runtimeconfig_config.config.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccRuntimeConfigConfigIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_runtimeconfig_config" "config" {
  provider    = google-beta
  name        = "tf-test-my-config%{random_suffix}"
  description = "Runtime configuration values for my service"
}

resource "google_runtimeconfig_config_iam_binding" "foo" {
  provider = google-beta
  project = google_runtimeconfig_config.config.project
  config = google_runtimeconfig_config.config.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccRuntimeConfigConfigIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_runtimeconfig_config" "config" {
  provider    = google-beta
  name        = "tf-test-my-config%{random_suffix}"
  description = "Runtime configuration values for my service"
}

resource "google_runtimeconfig_config_iam_binding" "foo" {
  provider = google-beta
  project = google_runtimeconfig_config.config.project
  config = google_runtimeconfig_config.config.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
