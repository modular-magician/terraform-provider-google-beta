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

package firestore_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccFirestoreIndex_firestoreIndexBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestFirestoreProjectFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckFirestoreIndexDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirestoreIndex_firestoreIndexBasicExample(context),
			},
			{
				ResourceName:            "google_firestore_index.my-index",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"database", "collection"},
			},
		},
	})
}

func testAccFirestoreIndex_firestoreIndexBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firestore_index" "my-index" {
  project = "%{project_id}"

  collection = "chatrooms"

  fields {
    field_path = "name"
    order      = "ASCENDING"
  }

  fields {
    field_path = "description"
    order      = "DESCENDING"
  }

}
`, context)
}

func TestAccFirestoreIndex_firestoreIndexDatastoreModeExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestFirestoreProjectFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckFirestoreIndexDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirestoreIndex_firestoreIndexDatastoreModeExample(context),
			},
			{
				ResourceName:            "google_firestore_index.my-datastore-mode-index",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"database", "collection"},
			},
		},
	})
}

func testAccFirestoreIndex_firestoreIndexDatastoreModeExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firestore_index" "my-datastore-mode-index" {
  project = "%{project_id}"

  collection = "chatrooms"

  query_scope = "COLLECTION_RECURSIVE"
  api_scope = "DATASTORE_MODE_API"

  fields {
    field_path = "name"
    order      = "ASCENDING"
  }

  fields {
    field_path = "description"
    order      = "DESCENDING"
  }

}
`, context)
}

func testAccCheckFirestoreIndexDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_firestore_index" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{FirestoreBasePath}}{{name}}")
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
				return fmt.Errorf("FirestoreIndex still exists at %s", url)
			}
		}

		return nil
	}
}
