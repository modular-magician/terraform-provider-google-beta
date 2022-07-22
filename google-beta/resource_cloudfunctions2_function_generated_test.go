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
)

func TestAccCloudfunctions2function_cloudfunctions2BasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":             getTestProjectFromEnv(),
		"zip_path":            "./test-fixtures/cloudfunctions2/function-source.zip",
		"primary_resource_id": "function",
		"location":            "us-central1",
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckCloudfunctions2functionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudfunctions2function_cloudfunctions2BasicExample(context),
			},
			{
				ResourceName:            "google_cloudfunctions2_function.function",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "build_config.0.source.0.storage_source.0.object", "build_config.0.source.0.storage_source.0.bucket"},
			},
		},
	})
}

func testAccCloudfunctions2function_cloudfunctions2BasicExample(context map[string]interface{}) string {
	return Nprintf(`
# [START functions_v2_basic]
locals {
  project = "%{project}" # Google Cloud Platform Project ID
}

provider "google-beta" {
   project = local.project
}

resource "google_storage_bucket" "bucket" {
  provider = google-beta
  name     = "${local.project}-tf-test-gcf-source%{random_suffix}"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  provider = google-beta
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}
 
resource "google_cloudfunctions2_function" "function" {
  provider = google-beta
  name = "tf-test-test-function%{random_suffix}"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime = "nodejs16"
    entry_point = "helloHttp"  # Set the entry point 
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }
 
  service_config {
    max_instance_count  = 1
    available_memory    = "256M"
    timeout_seconds     = 60
  }
}

output "function_uri" { 
  value = google_cloudfunctions2_function.function.service_config[0].uri
}
# [END functions_v2_basic]
`, context)
}

func TestAccCloudfunctions2function_cloudfunctions2FullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":             getTestProjectFromEnv(),
		"zip_path":            "./test-fixtures/cloudfunctions2/function-source-pubsub.zip",
		"primary_resource_id": "terraform-test",
		"location":            "us-central1",
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckCloudfunctions2functionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudfunctions2function_cloudfunctions2FullExample(context),
			},
			{
				ResourceName:            "google_cloudfunctions2_function.function",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "build_config.0.source.0.storage_source.0.object", "build_config.0.source.0.storage_source.0.bucket"},
			},
		},
	})
}

func testAccCloudfunctions2function_cloudfunctions2FullExample(context map[string]interface{}) string {
	return Nprintf(`
# [START functions_v2_full]
locals {
  project = "%{project}" # Google Cloud Platform Project ID
}

provider "google-beta" {
   project = local.project
}

resource "google_service_account" "account" {
  provider = google-beta
  account_id = "tf-test-test-sa%{random_suffix}"
  display_name = "Test Service Account"
}

resource "google_pubsub_topic" "topic" {
  provider = google-beta
  name = "tf-test-functions2-topic%{random_suffix}"
}

resource "google_storage_bucket" "bucket" {
  provider = google-beta
  name     = "${local.project}-tf-test-gcf-source%{random_suffix}"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  provider = google-beta
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}
 
resource "google_cloudfunctions2_function" "function" {
  provider = google-beta
  name = "tf-test-test-function%{random_suffix}"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime = "nodejs16"
    entry_point = "helloPubSub"  # Set the entry point 
    environment_variables = {
        BUILD_CONFIG_TEST = "build_test"
    }
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }
 
  service_config {
    max_instance_count  = 3
    min_instance_count = 1
    available_memory    = "256M"
    timeout_seconds     = 60
    environment_variables = {
        SERVICE_CONFIG_TEST = "config_test"
    }
    ingress_settings = "ALLOW_INTERNAL_ONLY"
    all_traffic_on_latest_revision = true
    service_account_email = google_service_account.account.email
  }

  event_trigger {
    trigger_region = "us-central1"
    event_type = "google.cloud.pubsub.topic.v1.messagePublished"
    pubsub_topic = google_pubsub_topic.topic.id
    retry_policy = "RETRY_POLICY_RETRY"
  }
}
# [END functions_v2_full]
`, context)
}

func testAccCheckCloudfunctions2functionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloudfunctions2_function" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{Cloudfunctions2BasePath}}projects/{{project}}/locations/{{location}}/functions/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("Cloudfunctions2function still exists at %s", url)
			}
		}

		return nil
	}
}
