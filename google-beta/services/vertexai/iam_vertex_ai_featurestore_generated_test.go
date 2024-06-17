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

package vertexai_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccVertexAIFeaturestoreIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(t, 10),
		"role":            "roles/viewer",
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"org_id":          envvar.GetTestOrgFromEnv(t),

		"kms_key_name": acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIFeaturestoreIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_featurestore_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/featurestores/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("terraform%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccVertexAIFeaturestoreIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_featurestore_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/featurestores/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("terraform%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVertexAIFeaturestoreIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(t, 10),
		"role":            "roles/viewer",
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"org_id":          envvar.GetTestOrgFromEnv(t),

		"kms_key_name": acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccVertexAIFeaturestoreIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_featurestore_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/featurestores/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("terraform%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVertexAIFeaturestoreIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(t, 10),
		"role":            "roles/viewer",
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"org_id":          envvar.GetTestOrgFromEnv(t),

		"kms_key_name": acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIFeaturestoreIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_vertex_ai_featurestore_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_vertex_ai_featurestore_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("terraform%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccVertexAIFeaturestoreIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_vertex_ai_featurestore_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("terraform%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccVertexAIFeaturestoreIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
  force_destroy = true
}

resource "google_vertex_ai_featurestore_iam_member" "foo" {
  project = google_vertex_ai_featurestore.featurestore.project
  region = google_vertex_ai_featurestore.featurestore.region
  featurestore = google_vertex_ai_featurestore.featurestore.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccVertexAIFeaturestoreIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
  force_destroy = true
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_vertex_ai_featurestore_iam_policy" "foo" {
  project = google_vertex_ai_featurestore.featurestore.project
  region = google_vertex_ai_featurestore.featurestore.region
  featurestore = google_vertex_ai_featurestore.featurestore.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_vertex_ai_featurestore_iam_policy" "foo" {
  project = google_vertex_ai_featurestore.featurestore.project
  region = google_vertex_ai_featurestore.featurestore.region
  featurestore = google_vertex_ai_featurestore.featurestore.name
  depends_on = [
    google_vertex_ai_featurestore_iam_policy.foo
  ]
}
`, context)
}

func testAccVertexAIFeaturestoreIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
  force_destroy = true
}

data "google_iam_policy" "foo" {
}

resource "google_vertex_ai_featurestore_iam_policy" "foo" {
  project = google_vertex_ai_featurestore.featurestore.project
  region = google_vertex_ai_featurestore.featurestore.region
  featurestore = google_vertex_ai_featurestore.featurestore.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccVertexAIFeaturestoreIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
  force_destroy = true
}

resource "google_vertex_ai_featurestore_iam_binding" "foo" {
  project = google_vertex_ai_featurestore.featurestore.project
  region = google_vertex_ai_featurestore.featurestore.region
  featurestore = google_vertex_ai_featurestore.featurestore.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccVertexAIFeaturestoreIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
  force_destroy = true
}

resource "google_vertex_ai_featurestore_iam_binding" "foo" {
  project = google_vertex_ai_featurestore.featurestore.project
  region = google_vertex_ai_featurestore.featurestore.region
  featurestore = google_vertex_ai_featurestore.featurestore.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
