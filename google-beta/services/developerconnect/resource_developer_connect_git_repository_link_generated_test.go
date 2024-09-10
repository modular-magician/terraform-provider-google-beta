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

package developerconnect_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccDeveloperConnectGitRepositoryLink_developerConnectGitRepositoryLinkGithubExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckDeveloperConnectGitRepositoryLinkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDeveloperConnectGitRepositoryLink_developerConnectGitRepositoryLinkGithubExample(context),
			},
			{
				ResourceName:            "google_developer_connect_git_repository_link.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "git_repository_link_id", "labels", "location", "parent_connection", "terraform_labels"},
			},
		},
	})
}

func testAccDeveloperConnectGitRepositoryLink_developerConnectGitRepositoryLinkGithubExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_developer_connect_git_repository_link" "primary" {
  provider = google-beta
  git_repository_link_id              = "tf-test-my-repository%{random_suffix}"
  parent_connection = google_developer_connect_connection.github_conn.connection_id
  clone_uri        = "https://github.com/gcb-developerconnect-robot/tf-demo.git"
  location          = "us-central1"
  annotations       = {}
}

resource "google_developer_connect_connection" "github_conn" {
  
  provider = google-beta
  location = "us-central1"
  connection_id     = "tf-test-my-connection%{random_suffix}"
  disabled = false

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

func testAccCheckDeveloperConnectGitRepositoryLinkDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_developer_connect_git_repository_link" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DeveloperConnectBasePath}}projects/{{project}}/locations/{{location}}/connections/{{parent_connection}}/gitRepositoryLinks/{{git_repository_link_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("DeveloperConnectGitRepositoryLink still exists at %s", url)
			}
		}

		return nil
	}
}
