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

package biglake_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccBiglakeTable_bigqueryBiglakeTableExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBiglakeTableDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBiglakeTable_bigqueryBiglakeTableExample(context),
			},
			{
				ResourceName:            "google_biglake_table.table",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name", "database_id", "catalog_id"},
			},
		},
	})
}

func testAccBiglakeTable_bigqueryBiglakeTableExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_biglake_catalog" "catalog" {
    name = "tf_test_my_catalog%{random_suffix}"
    location = "US"
}

resource "google_storage_bucket" "bucket" {
  name                        = "tf_test_my_bucket%{random_suffix}"
  location                    = "US"
  force_destroy               = true
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "metadata_folder" {
  name    = "metadata/"
  content = " "
  bucket  = google_storage_bucket.bucket.name
}


resource "google_storage_bucket_object" "data_folder" {
  name    = "data/"
  content = " "
  bucket  = google_storage_bucket.bucket.name
}

resource "google_biglake_database" "database" {
    name = "tf_test_my_database%{random_suffix}"
    catalog_id = google_biglake_catalog.catalog.name
    location = google_biglake_catalog.catalog.location
    type = "HIVE"
    hive_options {
        location_uri = "gs://${google_storage_bucket.bucket.name}/${google_storage_bucket_object.metadata_folder.name}"
        parameters = {
          "name" = "wrench"
        }
    }
}

resource "google_biglake_table" "table" {
    name = "tf_test_my_table%{random_suffix}"
    catalog_id = google_biglake_catalog.catalog.name
    database_id = google_biglake_database.database.name
    location = google_biglake_catalog.catalog.location
    type = "HIVE"
    hive_options {
      table_type = "MANAGED_TABLE"
      storage_descriptor {
        location_uri = "gs://${google_storage_bucket.bucket.name}/${google_storage_bucket_object.data_folder.name}"
      }
      parameters = {
        "name" = "screwdriver"
      }
    }
}
`, context)
}

func testAccCheckBiglakeTableDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_biglake_table" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{BiglakeBasePath}}projects/{{project}}/locations/{{location}}/catalogs/{{catalog_id}}/databases/{{database_id}}/tables/{{name}}")
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
				return fmt.Errorf("BiglakeTable still exists at %s", url)
			}
		}

		return nil
	}
}
