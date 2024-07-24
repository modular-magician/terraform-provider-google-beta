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

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccFirestoreIndex_firestoreIndexBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckFirestoreIndexDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirestoreIndex_firestoreIndexBasicExample(context),
			},
			{
				ResourceName:            "google_firestore_index.my-index",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"collection", "database"},
			},
		},
	})
}

func testAccFirestoreIndex_firestoreIndexBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firestore_database" "database" {
  project     = "%{project_id}"
  name        = "tf-test-database-id%{random_suffix}"
  location_id = "nam5"
  type        = "FIRESTORE_NATIVE"

  delete_protection_state = "DELETE_PROTECTION_DISABLED"
  deletion_policy         = "DELETE"
}

resource "google_firestore_index" "my-index" {
  project     = "%{project_id}"
  database   = google_firestore_database.database.name
  collection = "atestcollection"

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
		"project_id":    envvar.GetTestProjectFromEnv(),
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
				ResourceName:            "google_firestore_index.my-index",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"collection", "database"},
			},
		},
	})
}

func testAccFirestoreIndex_firestoreIndexDatastoreModeExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firestore_database" "database" {
  project     = "%{project_id}"
  name        = "tf-test-database-id-dm%{random_suffix}"
  location_id = "nam5"
  type        = "DATASTORE_MODE"

  delete_protection_state = "DELETE_PROTECTION_DISABLED"
  deletion_policy         = "DELETE"
}

resource "google_firestore_index" "my-index" {
  project     = "%{project_id}"
  database   = google_firestore_database.database.name
  collection = "atestcollection"

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

func TestAccFirestoreIndex_firestoreIndexVectorExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckFirestoreIndexDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirestoreIndex_firestoreIndexVectorExample(context),
			},
			{
				ResourceName:            "google_firestore_index.my-index",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"collection", "database"},
			},
		},
	})
}

func testAccFirestoreIndex_firestoreIndexVectorExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firestore_database" "database" {
  project     = "%{project_id}"
  name        = "tf-test-database-id-vector%{random_suffix}"
  location_id = "nam5"
  type        = "FIRESTORE_NATIVE"

  delete_protection_state = "DELETE_PROTECTION_DISABLED"
  deletion_policy         = "DELETE"
}

resource "google_firestore_index" "my-index" {
  project     = "%{project_id}"
  database   = google_firestore_database.database.name
  collection = "atestcollection"

  fields {
    field_path = "field_name"
    order      = "ASCENDING"
  }

  fields {
    field_path = "__name__"
    order      = "ASCENDING"
  }

  fields {
    field_path = "description"
    vector_config {
      dimension = 128
      flat {}
    }
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
				Config:               config,
				Method:               "GET",
				Project:              billingProject,
				RawURL:               url,
				UserAgent:            config.UserAgent,
				ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.FirestoreIndex409Retry},
			})
			if err == nil {
				return fmt.Errorf("FirestoreIndex still exists at %s", url)
			}
		}

		return nil
	}
}
