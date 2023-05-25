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

func TestAccBigqueryReservationCapacityCommitment_bigqueryReservationCapacityCommitmentBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckBigqueryReservationCapacityCommitmentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryReservationCapacityCommitment_bigqueryReservationCapacityCommitmentBasicExample(context),
			},
			{
				ResourceName:            "google_bigquery_capacity_commitment.commitment",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"capacity_commitment_id", "location", "enforce_single_admin_project_per_org"},
			},
		},
	})
}

func testAccBigqueryReservationCapacityCommitment_bigqueryReservationCapacityCommitmentBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_capacity_commitment" "commitment" {
	capacity_commitment_id = "capacity-tf-test%{random_suffix}"

	location   = "us-west2"
	slot_count = 100
	plan       = "FLEX_FLAT_RATE"
	edition    = "ENTERPRISE"
}

resource "time_sleep" "wait_61_seconds" {
	depends_on = [google_bigquery_capacity_commitment.commitment]
    
	# Only needed for CI tests to be able to tear down the commitment once it's expired
    create_duration = "61s"
}
`, context)
}

func testAccCheckBigqueryReservationCapacityCommitmentDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_bigquery_capacity_commitment" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/capacityCommitments/{{capacity_commitment_id}}")
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
				return fmt.Errorf("BigqueryReservationCapacityCommitment still exists at %s", url)
			}
		}

		return nil
	}
}
