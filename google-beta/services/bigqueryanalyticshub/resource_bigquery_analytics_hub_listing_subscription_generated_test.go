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

package bigqueryanalyticshub_test

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

func TestAccBigqueryAnalyticsHubListingSubscription_bigqueryAnalyticshubListingSubscriptionBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBigqueryAnalyticsHubListingSubscriptionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryAnalyticsHubListingSubscription_bigqueryAnalyticshubListingSubscriptionBasicExample(context),
			},
			{
				ResourceName:            "google_bigquery_analytics_hub_listing_subscription.listing_subscription",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"data_exchange_id", "listing_id", "location"},
			},
		},
	})
}

func testAccBigqueryAnalyticsHubListingSubscription_bigqueryAnalyticshubListingSubscriptionBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_analytics_hub_data_exchange" "listing_subscription" {
  location         = "US"
  data_exchange_id = "tf_test_my_data_exchange%{random_suffix}"
  display_name     = "tf_test_my_data_exchange%{random_suffix}"
  description      = "example data exchange%{random_suffix}"
}

resource "google_bigquery_analytics_hub_listing" "listing_subscription" {
  location         = "US"
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.listing_subscription.data_exchange_id
  listing_id       = "tf_test_my_listing%{random_suffix}"
  display_name     = "tf_test_my_listing%{random_suffix}"
  description      = "example data exchange%{random_suffix}"

  bigquery_dataset {
    dataset = google_bigquery_dataset.listing_subscription.id
  }
}

resource "google_bigquery_analytics_hub_listing_subscription" "listing_subscription" {
  location = "US"
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.listing_subscription.data_exchange_id
  listing_id       = google_bigquery_analytics_hub_listing.listing_subscription.listing_id
  destination_dataset {
    description = "A test subscription"
    friendly_name = "👋"
    labels = {
      testing = "123"
    }
    location = "US"
    dataset_reference {
      dataset_id = "tf_test_destination_dataset%{random_suffix}"
      project_id = google_bigquery_dataset.listing_subscription.project
    }
  }
}

resource "google_bigquery_dataset" "listing_subscription" {
  dataset_id                  = "tf_test_my_listing%{random_suffix}"
  friendly_name               = "tf_test_my_listing%{random_suffix}"
  description                 = "example data exchange%{random_suffix}"
  location                    = "US"
}
`, context)
}

func testAccCheckBigqueryAnalyticsHubListingSubscriptionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_bigquery_analytics_hub_listing_subscription" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{BigqueryAnalyticsHubBasePath}}{{name}}")
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
				return fmt.Errorf("BigqueryAnalyticsHubListingSubscription still exists at %s", url)
			}
		}

		return nil
	}
}
