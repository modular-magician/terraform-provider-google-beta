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

func TestAccBiglakeDatabase_bigqueryBiglakeDatabaseExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBiglakeDatabaseDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBiglakeDatabase_bigqueryBiglakeDatabaseExample(context),
			},
			{
				ResourceName:            "google_biglake_database.database",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name", "catalog_id"},
			},
		},
	})
}

func testAccBiglakeDatabase_bigqueryBiglakeDatabaseExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_biglake_catalog" "catalog" {
    name = "tf_test_my_catalog%{random_suffix}"
    # Hard code to avoid invalid random id suffix
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

resource "google_biglake_database" "database"  {
    name = "tf_test_my_database%{random_suffix}"
    catalog_id = google_biglake_catalog.catalog.name
    location = google_biglake_catalog.catalog.location
    type = "HIVE"
    hive_options {
        location_uri = "gs://${google_storage_bucket.bucket.name}/${google_storage_bucket_object.metadata_folder.name}"
        parameters = {
          "tool" = "hammer"
        }
    }
}
`, context)
}

func testAccCheckBiglakeDatabaseDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_biglake_database" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{BiglakeBasePath}}projects/{{project}}/locations/{{location}}/catalogs/{{catalog_id}}/databases/{{name}}")
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
				return fmt.Errorf("BiglakeDatabase still exists at %s", url)
			}
		}

		return nil
	}
}
