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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccVertexAIFeatureGroup_vertexAiFeatureGroupExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckVertexAIFeatureGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIFeatureGroup_vertexAiFeatureGroupExample(context),
			},
			{
				ResourceName:            "google_vertex_ai_feature_group.feature_group",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"description", "etag", "labels", "region", "terraform_labels"},
			},
		},
	})
}

func testAccVertexAIFeatureGroup_vertexAiFeatureGroupExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
  resource "google_vertex_ai_feature_group" "feature_group" {
  name = "tf_test_example_feature_group%{random_suffix}"
  description = "A sample feature group"
  region = "us-central1"
  labels = {
      label-one = "value-one"
  }
  big_query {
    big_query_source {
        # The source table must have a column named 'feature_timestamp' of type TIMESTAMP.
        input_uri = "bq://${google_bigquery_table.sample_table.project}.${google_bigquery_table.sample_table.dataset_id}.${google_bigquery_table.sample_table.table_id}"
    }
    entity_id_columns = ["feature_id"]
  }
}

resource "google_bigquery_dataset" "sample_dataset" {
  dataset_id                  = "tf_test_job_load%{random_suffix}_dataset"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "US"
}

resource "google_bigquery_table" "sample_table" {
  deletion_protection = false
  dataset_id = google_bigquery_dataset.sample_dataset.dataset_id
  table_id   = "tf_test_job_load%{random_suffix}_table"

  schema = <<EOF
[
    {
        "name": "feature_id",
        "type": "STRING",
        "mode": "NULLABLE"
    },
    {
        "name": "feature_timestamp",
        "type": "TIMESTAMP",
        "mode": "NULLABLE"
    }
]
EOF
}
`, context)
}

func testAccCheckVertexAIFeatureGroupDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_vertex_ai_feature_group" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featureGroups/{{name}}")
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
				return fmt.Errorf("VertexAIFeatureGroup still exists at %s", url)
			}
		}

		return nil
	}
}
