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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccCloudbuildv2ConnectionIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/cloudbuild.connectionViewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudbuildv2ConnectionIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccCloudbuildv2ConnectionIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccCloudbuildv2ConnectionIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/cloudbuild.connectionViewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccCloudbuildv2ConnectionIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccCloudbuildv2ConnectionIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/cloudbuild.connectionViewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudbuildv2ConnectionIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_cloudbuildv2_connection_iam_policy.foo", "policy_data"),
			},
			{
				Config: testAccCloudbuildv2ConnectionIamPolicy_emptyBinding(context),
			},
		},
	})
}

func testAccCloudbuildv2ConnectionIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "my-connection" {
  provider = google-beta
  location = "us-central1"
  name = "tf-test-tf-test-connection%{random_suffix}"

  github_config {
    app_installation_id = 0

    authorizer_credential {
      oauth_token_secret_version = "projects/gcb-terraform-creds/secrets/github-pat/versions/1"
    }
  }
}

resource "google_cloudbuildv2_connection_iam_member" "foo" {
  provider = google-beta
  project = google_cloudbuildv2_connection.my-connection.project
  location = google_cloudbuildv2_connection.my-connection.location
  name = google_cloudbuildv2_connection.my-connection.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccCloudbuildv2ConnectionIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "my-connection" {
  provider = google-beta
  location = "us-central1"
  name = "tf-test-tf-test-connection%{random_suffix}"

  github_config {
    app_installation_id = 0

    authorizer_credential {
      oauth_token_secret_version = "projects/gcb-terraform-creds/secrets/github-pat/versions/1"
    }
  }
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_cloudbuildv2_connection_iam_policy" "foo" {
  provider = google-beta
  project = google_cloudbuildv2_connection.my-connection.project
  location = google_cloudbuildv2_connection.my-connection.location
  name = google_cloudbuildv2_connection.my-connection.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_cloudbuildv2_connection_iam_policy" "foo" {
  provider = google-beta
  project = google_cloudbuildv2_connection.my-connection.project
  location = google_cloudbuildv2_connection.my-connection.location
  name = google_cloudbuildv2_connection.my-connection.name
  depends_on = [
    google_cloudbuildv2_connection_iam_policy.foo
  ]
}
`, context)
}

func testAccCloudbuildv2ConnectionIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "my-connection" {
  provider = google-beta
  location = "us-central1"
  name = "tf-test-tf-test-connection%{random_suffix}"

  github_config {
    app_installation_id = 0

    authorizer_credential {
      oauth_token_secret_version = "projects/gcb-terraform-creds/secrets/github-pat/versions/1"
    }
  }
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_cloudbuildv2_connection_iam_policy" "foo" {
  provider = google-beta
  project = google_cloudbuildv2_connection.my-connection.project
  location = google_cloudbuildv2_connection.my-connection.location
  name = google_cloudbuildv2_connection.my-connection.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccCloudbuildv2ConnectionIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "my-connection" {
  provider = google-beta
  location = "us-central1"
  name = "tf-test-tf-test-connection%{random_suffix}"

  github_config {
    app_installation_id = 0

    authorizer_credential {
      oauth_token_secret_version = "projects/gcb-terraform-creds/secrets/github-pat/versions/1"
    }
  }
}

resource "google_cloudbuildv2_connection_iam_binding" "foo" {
  provider = google-beta
  project = google_cloudbuildv2_connection.my-connection.project
  location = google_cloudbuildv2_connection.my-connection.location
  name = google_cloudbuildv2_connection.my-connection.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccCloudbuildv2ConnectionIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloudbuildv2_connection" "my-connection" {
  provider = google-beta
  location = "us-central1"
  name = "tf-test-tf-test-connection%{random_suffix}"

  github_config {
    app_installation_id = 0

    authorizer_credential {
      oauth_token_secret_version = "projects/gcb-terraform-creds/secrets/github-pat/versions/1"
    }
  }
}

resource "google_cloudbuildv2_connection_iam_binding" "foo" {
  provider = google-beta
  project = google_cloudbuildv2_connection.my-connection.project
  location = google_cloudbuildv2_connection.my-connection.location
  name = google_cloudbuildv2_connection.my-connection.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
