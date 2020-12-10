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
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPubsubLiteSubscription_pubsubLiteSubscriptionBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckPubsubLiteSubscriptionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubLiteSubscription_pubsubLiteSubscriptionBasicExample(context),
			},
			{
				ResourceName:            "google_pubsub_lite_subscription.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"topic", "region", "zone", "name"},
			},
		},
	})
}

func testAccPubsubLiteSubscription_pubsubLiteSubscriptionBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_lite_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"
  project = data.google_project.project.number
  partition_config {
    count = 1
    capacity {
      publish_mib_per_sec = 4
      subscribe_mib_per_sec = 8
    }
  }

  retention_config {
    per_partition_bytes = 32212254720
  }
}

resource "google_pubsub_lite_subscription" "example" {
  name  = "tf-test-example-subscription%{random_suffix}"
  topic = google_pubsub_lite_topic.example.name
  delivery_config {
    delivery_requirement = "DELIVER_AFTER_STORED"
  }
}

data "google_project" "project" {
}
`, context)
}

func testAccCheckPubsubLiteSubscriptionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_pubsub_lite_subscription" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{PubsubLiteBasePath}}projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("PubsubLiteSubscription still exists at %s", url)
			}
		}

		return nil
	}
}
