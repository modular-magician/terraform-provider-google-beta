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
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccVertexAIIndex_vertexAiIndexExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckVertexAIIndexDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIIndex_vertexAiIndexExample(context),
			},
			{
				ResourceName:            "google_vertex_ai_index.index",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "region", "metadata.0.contents_delta_uri", "metadata.0.is_complete_overwrite"},
			},
		},
	})
}

func testAccVertexAIIndex_vertexAiIndexExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
  name     = "tf-test-vertex-ai-index-test%{random_suffix}"
  location = "us-central1"
  uniform_bucket_level_access = true
}

# The sample data comes from the following link:
# https://cloud.google.com/vertex-ai/docs/matching-engine/filtering#specify-namespaces-tokens
resource "google_storage_bucket_object" "data" {
  name   = "contents/data.json"
  bucket = google_storage_bucket.bucket.name
  content = <<EOF
{"id": "42", "embedding": [0.5, 1.0], "restricts": [{"namespace": "class", "allow": ["cat", "pet"]},{"namespace": "category", "allow": ["feline"]}]}
{"id": "43", "embedding": [0.6, 1.0], "restricts": [{"namespace": "class", "allow": ["dog", "pet"]},{"namespace": "category", "allow": ["canine"]}]}
EOF
}

resource "google_vertex_ai_index" "index" {
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  display_name = "tf-test-test-index%{random_suffix}"
  description = "index for test"
  metadata {
    contents_delta_uri = "gs://${google_storage_bucket.bucket.name}/contents"
    config {
      dimensions = 2
      approximate_neighbors_count = 150
      distance_measure_type = "DOT_PRODUCT_DISTANCE"
      algorithm_config {
        tree_ah_config {
          leaf_node_embedding_count = 500
          leaf_nodes_to_search_percent = 7
        }
      }
    }
  }
  index_update_method = "BATCH_UPDATE"
}
`, context)
}

func TestAccVertexAIIndex_vertexAiIndexStreamingExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckVertexAIIndexDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIIndex_vertexAiIndexStreamingExample(context),
			},
			{
				ResourceName:            "google_vertex_ai_index.index",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "region", "metadata.0.contents_delta_uri", "metadata.0.is_complete_overwrite"},
			},
		},
	})
}

func testAccVertexAIIndex_vertexAiIndexStreamingExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
  name     = "tf-test-vertex-ai-index-test%{random_suffix}"
  location = "us-central1"
  uniform_bucket_level_access = true
}

# The sample data comes from the following link:
# https://cloud.google.com/vertex-ai/docs/matching-engine/filtering#specify-namespaces-tokens
resource "google_storage_bucket_object" "data" {
  name   = "contents/data.json"
  bucket = google_storage_bucket.bucket.name
  content = <<EOF
{"id": "42", "embedding": [0.5, 1.0], "restricts": [{"namespace": "class", "allow": ["cat", "pet"]},{"namespace": "category", "allow": ["feline"]}]}
{"id": "43", "embedding": [0.6, 1.0], "restricts": [{"namespace": "class", "allow": ["dog", "pet"]},{"namespace": "category", "allow": ["canine"]}]}
EOF
}

resource "google_vertex_ai_index" "index" {
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  display_name = "tf-test-test-index%{random_suffix}"
  description = "index for test"
  metadata {
    contents_delta_uri = "gs://${google_storage_bucket.bucket.name}/contents"
    config {
      dimensions = 2
      distance_measure_type = "COSINE_DISTANCE"
      feature_norm_type = "UNIT_L2_NORM"
      algorithm_config {
        brute_force_config {}
      }
    }
  }
  index_update_method = "STREAM_UPDATE"
}
`, context)
}

func testAccCheckVertexAIIndexDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_vertex_ai_index" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/indexes/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("VertexAIIndex still exists at %s", url)
			}
		}

		return nil
	}
}
