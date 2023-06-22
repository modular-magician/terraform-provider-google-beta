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

package google

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

func TestAccPubsubLiteTopic_pubsubLiteTopicBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPubsubLiteTopicDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubLiteTopic_pubsubLiteTopicBasicExample(context),
			},
			{
				ResourceName:            "google_pubsub_lite_topic.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "zone", "name"},
			},
		},
	})
}

func testAccPubsubLiteTopic_pubsubLiteTopicBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_lite_reservation" "example" {
  name = "tf-test-example-reservation%{random_suffix}"
  project = data.google_project.project.number
  throughput_capacity = 2
}

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

  reservation_config {
    throughput_reservation = google_pubsub_lite_reservation.example.name
  }
}

data "google_project" "project" {
}
`, context)
}

func testAccCheckPubsubLiteTopicDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_pubsub_lite_topic" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{PubsubLiteBasePath}}projects/{{project}}/locations/{{zone}}/topics/{{name}}")
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
				return fmt.Errorf("PubsubLiteTopic still exists at %s", url)
			}
		}

		return nil
	}
}
