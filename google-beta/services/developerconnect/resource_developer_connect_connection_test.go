// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package developerconnect_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccDeveloperConnectConnection_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDeveloperConnectConnection_basic(context),
			},
			{
				ResourceName:            "google_developer_connect_connection.my-connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "connection_id", "labels", "location", "terraform_labels"},
			},
			{
				Config: testAccDeveloperConnectConnection_update(context),
			},
			{
				ResourceName:            "google_developer_connect_connection.my-connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "connection_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccDeveloperConnectConnection_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_developer_connect_connection" "my-connection" {
  provider = google-beta
  location = "us-central1"
  connection_id = "tf-test-tf-test-connection%{random_suffix}"

  github_config {
    github_app = "DEVELOPER_CONNECT"

    authorizer_credential {
      oauth_token_secret_version = "projects/devconnect-terraform-creds/secrets/tf-test-do-not-change-github-oauthtoken-e0b9e7/versions/1"
    }
  }
}
`, context)
}

func testAccDeveloperConnectConnection_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_developer_connect_connection" "my-connection" {
  provider = google-beta
  location = "us-central1"
  connection_id = "tf-test-tf-test-connection%{random_suffix}"

  github_config {
    github_app = "DEVELOPER_CONNECT"
    app_installation_id = 49439208

    authorizer_credential {
      oauth_token_secret_version = "projects/devconnect-terraform-creds/secrets/tf-test-do-not-change-github-oauthtoken-e0b9e7/versions/1"
    }
  }
}
`, context)
}
