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

package dataform_test

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

func TestAccDataformRepository_dataformRepositoryWithCloudsourceRepoExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckDataformRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataformRepository_dataformRepositoryWithCloudsourceRepoExample(context),
			},
			{
				ResourceName:            "google_dataform_repository.dataform_repository",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "region", "terraform_labels"},
			},
		},
	})
}

func testAccDataformRepository_dataformRepositoryWithCloudsourceRepoExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
  provider = google-beta
}

resource "google_sourcerepo_repository" "git_repository" {
  provider = google-beta
  name = "my/repository%{random_suffix}"
}

resource "google_secret_manager_secret" "secret" {
  provider = google-beta
  secret_id = "tf-test-my-secret%{random_suffix}"

  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "secret_version" {
  provider = google-beta
  secret = google_secret_manager_secret.secret.id

  secret_data = "tf-test-secret-data%{random_suffix}"
}

resource "google_kms_key_ring" "keyring" {
  provider = google-beta
  
  name     = "tf-test-example-key-ring%{random_suffix}"
  location = "us-central1"
}

resource "google_kms_crypto_key" "example_key" {
  provider = google-beta
  
  name            = "tf-test-example-crypto-key-name%{random_suffix}"
  key_ring        = google_kms_key_ring.keyring.id
}

resource "google_kms_crypto_key_iam_binding" "crypto_key_binding" {
  provider = google-beta

  crypto_key_id = google_kms_crypto_key.example_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"

  members = [
    "serviceAccount:service-${data.google_project.project.number}@gcp-sa-dataform.iam.gserviceaccount.com",
  ]
}

resource "google_dataform_repository" "dataform_repository" {
  provider = google-beta
  name = "tf_test_dataform_repository%{random_suffix}"
  display_name = "tf_test_dataform_repository%{random_suffix}"
  npmrc_environment_variables_secret_version = google_secret_manager_secret_version.secret_version.id
  kms_key_name = google_kms_crypto_key.example_key.id

  labels = {
    label_foo1 = "label-bar1"
  }

  git_remote_settings {
      url = google_sourcerepo_repository.git_repository.url
      default_branch = "main"
      authentication_token_secret_version = google_secret_manager_secret_version.secret_version.id
  }

  workspace_compilation_overrides {
    default_database = "database"
    schema_suffix = "_suffix"
    table_prefix = "prefix_"
  }

  depends_on = [
    google_kms_crypto_key_iam_binding.crypto_key_binding
  ]
}
`, context)
}

func TestAccDataformRepository_dataformRepositoryWithCloudsourceRepoAndSshExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckDataformRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataformRepository_dataformRepositoryWithCloudsourceRepoAndSshExample(context),
			},
			{
				ResourceName:            "google_dataform_repository.dataform_repository",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "region", "terraform_labels"},
			},
		},
	})
}

func testAccDataformRepository_dataformRepositoryWithCloudsourceRepoAndSshExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sourcerepo_repository" "git_repository" {
  provider = google-beta
  name = "my/repository%{random_suffix}"
}

resource "google_secret_manager_secret" "secret" {
  provider = google-beta
  secret_id = "tf-test-my-secret%{random_suffix}"

  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "secret_version" {
  provider = google-beta
  secret = google_secret_manager_secret.secret.id

  secret_data = "tf-test-secret-data%{random_suffix}"
}

resource "google_dataform_repository" "dataform_repository" {
  provider = google-beta
  name = "tf_test_dataform_repository%{random_suffix}"

  git_remote_settings {
      url = google_sourcerepo_repository.git_repository.url
      default_branch = "main"
      ssh_authentication_config {
        user_private_key_secret_version = google_secret_manager_secret_version.secret_version.id
        host_public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAklOUpkDHrfHY17SbrmTIpNLTGK9Tjom/BWDSU"
      }
  }

  workspace_compilation_overrides {
    default_database = "database"
    schema_suffix = "_suffix"
    table_prefix = "prefix_"
  }

  service_account = "1234567890-compute@developer.gserviceaccount.com"
}
`, context)
}

func testAccCheckDataformRepositoryDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dataform_repository" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DataformBasePath}}projects/{{project}}/locations/{{region}}/repositories/{{name}}")
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
				return fmt.Errorf("DataformRepository still exists at %s", url)
			}
		}

		return nil
	}
}
