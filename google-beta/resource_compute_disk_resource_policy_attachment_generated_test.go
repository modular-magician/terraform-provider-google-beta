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

func TestAccComputeDiskResourcePolicyAttachment_diskResourcePolicyAttachmentBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeDiskResourcePolicyAttachmentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeDiskResourcePolicyAttachment_diskResourcePolicyAttachmentBasicExample(context),
			},
			{
				ResourceName:            "google_compute_disk_resource_policy_attachment.attachment",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"disk", "zone"},
			},
		},
	})
}

func testAccComputeDiskResourcePolicyAttachment_diskResourcePolicyAttachmentBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_disk_resource_policy_attachment" "attachment" {
  name = google_compute_resource_policy.policy.name
  disk = google_compute_disk.ssd.name
  zone = "us-central1-a"
}

resource "google_compute_disk" "ssd" {
  name  = "tf-test-my-disk%{random_suffix}"
  image = data.google_compute_image.my_image.self_link
  size  = 50
  type  = "pd-ssd"
  zone  = "us-central1-a"
}

resource "google_compute_resource_policy" "policy" {
  name = "tf-test-my-resource-policy%{random_suffix}"
  region = "us-central1"
  snapshot_schedule_policy {
    schedule {
      daily_schedule {
        days_in_cycle = 1
        start_time = "04:00"
      }
    }
  }
}

data "google_compute_image" "my_image" {
  family  = "debian-11"
  project = "debian-cloud"
}
`, context)
}

func testAccCheckComputeDiskResourcePolicyAttachmentDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_disk_resource_policy_attachment" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/disks/{{disk}}")
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
				return fmt.Errorf("ComputeDiskResourcePolicyAttachment still exists at %s", url)
			}
		}

		return nil
	}
}
