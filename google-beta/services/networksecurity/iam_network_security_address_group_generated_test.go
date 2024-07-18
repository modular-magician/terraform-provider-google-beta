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

package networksecurity_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccNetworkSecurityProjectAddressGroupIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/compute.networkAdmin",
		"project":       envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityProjectAddressGroupIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccNetworkSecurityProjectAddressGroupIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccNetworkSecurityProjectAddressGroupIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/compute.networkAdmin",
		"project":       envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccNetworkSecurityProjectAddressGroupIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccNetworkSecurityProjectAddressGroupIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/compute.networkAdmin",
		"project":       envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityProjectAddressGroupIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_network_security_address_group_iam_policy.foo", "policy_data"),
			},
			{
				Config: testAccNetworkSecurityProjectAddressGroupIamPolicy_emptyBinding(context),
			},
		},
	})
}

func testAccNetworkSecurityProjectAddressGroupIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_address_group" "default" {
  name        = "tf-test-my-project-address-group%{random_suffix}"
  parent      = "projects/%{project}"
  location    = "us-central1"
  type        = "IPV4"
  capacity    = "100"
  items       = ["208.80.154.224/32"]
}

resource "google_network_security_address_group_iam_member" "foo" {
project = "%{project}"
location = google_network_security_address_group.default.location
name = google_network_security_address_group.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccNetworkSecurityProjectAddressGroupIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_address_group" "default" {
  name        = "tf-test-my-project-address-group%{random_suffix}"
  parent      = "projects/%{project}"
  location    = "us-central1"
  type        = "IPV4"
  capacity    = "100"
  items       = ["208.80.154.224/32"]
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_network_security_address_group_iam_policy" "foo" {
project = "%{project}"
location = google_network_security_address_group.default.location
name = google_network_security_address_group.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_network_security_address_group_iam_policy" "foo" {
project = "%{project}"
location = google_network_security_address_group.default.location
name = google_network_security_address_group.default.name
  depends_on = [
    google_network_security_address_group_iam_policy.foo
  ]
}
`, context)
}

func testAccNetworkSecurityProjectAddressGroupIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_address_group" "default" {
  name        = "tf-test-my-project-address-group%{random_suffix}"
  parent      = "projects/%{project}"
  location    = "us-central1"
  type        = "IPV4"
  capacity    = "100"
  items       = ["208.80.154.224/32"]
}

data "google_iam_policy" "foo" {
}

resource "google_network_security_address_group_iam_policy" "foo" {
project = "%{project}"
location = google_network_security_address_group.default.location
name = google_network_security_address_group.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccNetworkSecurityProjectAddressGroupIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_address_group" "default" {
  name        = "tf-test-my-project-address-group%{random_suffix}"
  parent      = "projects/%{project}"
  location    = "us-central1"
  type        = "IPV4"
  capacity    = "100"
  items       = ["208.80.154.224/32"]
}

resource "google_network_security_address_group_iam_binding" "foo" {
project = "%{project}"
location = google_network_security_address_group.default.location
name = google_network_security_address_group.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccNetworkSecurityProjectAddressGroupIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_address_group" "default" {
  name        = "tf-test-my-project-address-group%{random_suffix}"
  parent      = "projects/%{project}"
  location    = "us-central1"
  type        = "IPV4"
  capacity    = "100"
  items       = ["208.80.154.224/32"]
}

resource "google_network_security_address_group_iam_binding" "foo" {
project = "%{project}"
location = google_network_security_address_group.default.location
name = google_network_security_address_group.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
