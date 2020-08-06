// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAppEngineServiceSplitTraffic_appEngineServiceSplitTrafficExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppEngineServiceSplitTrafficDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAppEngineServiceSplitTraffic_appEngineServiceSplitTrafficExample(context),
			},
			{
				ResourceName:            "google_app_engine_service_split_traffic.liveapp",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"split", "migrate_traffic"},
			},
		},
	})
}

func testAccAppEngineServiceSplitTraffic_appEngineServiceSplitTrafficExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
	name = "tf-test-appengine-static-content%{random_suffix}"
}

resource "google_storage_bucket_object" "object" {
	name   = "hello-world.zip"
	bucket = google_storage_bucket.bucket.name
	source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "liveapp_v1" {
  version_id = "v1"
  service = "liveapp"
  delete_service_on_destroy = true

  runtime = "nodejs10"
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.name}"
    }  
  }
  env_variables = {
    port = "8080"
  }
}

resource "google_app_engine_standard_app_version" "liveapp_v2" {
  version_id = "v2"
  service = "liveapp"
  noop_on_destroy = true

  runtime = "nodejs10"
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.name}"
    }  
  }
  env_variables = {
    port = "8080"
  }
}

resource "google_app_engine_service_split_traffic" "liveapp" {
  service = google_app_engine_standard_app_version.liveapp_v2.service

  migrate_traffic = false
  split {
    shard_by = "IP"
    allocations = {
      (google_app_engine_standard_app_version.liveapp_v1.version_id) = 0.75
      (google_app_engine_standard_app_version.liveapp_v2.version_id) = 0.25
    }
  }
}
`, context)
}

func testAccCheckAppEngineServiceSplitTrafficDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_app_engine_service_split_traffic" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			log.Printf("[DEBUG] Ignoring destroy during test")
		}

		return nil
	}
}
