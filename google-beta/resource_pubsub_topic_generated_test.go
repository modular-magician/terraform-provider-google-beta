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

func TestAccPubsubTopic_pubsubTopicBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPubsubTopicDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubTopic_pubsubTopicBasicExample(context),
			},
			{
				ResourceName:      "google_pubsub_topic.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPubsubTopic_pubsubTopicBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"

  labels = {
    foo = "bar"
  }
}
`, context)
}

func TestAccPubsubTopic_pubsubTopicGeoRestrictedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPubsubTopicDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubTopic_pubsubTopicGeoRestrictedExample(context),
			},
			{
				ResourceName:      "google_pubsub_topic.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPubsubTopic_pubsubTopicGeoRestrictedExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"

  message_storage_policy {
    allowed_persistence_regions = [
      "europe-west3",
    ]
  }
}
`, context)
}

func TestAccPubsubTopic_pubsubTopicSchemaSettingsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPubsubTopicDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubTopic_pubsubTopicSchemaSettingsExample(context),
			},
			{
				ResourceName:      "google_pubsub_topic.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPubsubTopic_pubsubTopicSchemaSettingsExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_schema" "example" {
  name = "example%{random_suffix}"
  type = "AVRO"
  definition = "{\n  \"type\" : \"record\",\n  \"name\" : \"Avro\",\n  \"fields\" : [\n    {\n      \"name\" : \"StringField\",\n      \"type\" : \"string\"\n    },\n    {\n      \"name\" : \"IntField\",\n      \"type\" : \"int\"\n    }\n  ]\n}\n"
}

resource "google_pubsub_topic" "example" {
  name = "tf-test-example-topic%{random_suffix}"

  depends_on = [google_pubsub_schema.example]
  schema_settings {
    schema = "projects/%{project_name}/schemas/example%{random_suffix}"
    encoding = "JSON"
  }
}
`, context)
}

func testAccCheckPubsubTopicDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_pubsub_topic" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{PubsubBasePath}}projects/{{project}}/topics/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil, pubsubTopicProjectNotReady)
			if err == nil {
				return fmt.Errorf("PubsubTopic still exists at %s", url)
			}
		}

		return nil
	}
}
