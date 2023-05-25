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

func TestAccSpannerInstance_spannerInstanceBasicExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSpannerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSpannerInstance_spannerInstanceBasicExample(context),
			},
			{
				ResourceName:            "google_spanner_instance.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"config"},
			},
		},
	})
}

func testAccSpannerInstance_spannerInstanceBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_spanner_instance" "example" {
  config       = "regional-us-central1"
  display_name = "Test Spanner Instance"
  num_nodes    = 2
  labels = {
    "foo" = "bar"
  }
}
`, context)
}

func TestAccSpannerInstance_spannerInstanceProcessingUnitsExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSpannerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSpannerInstance_spannerInstanceProcessingUnitsExample(context),
			},
			{
				ResourceName:            "google_spanner_instance.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"config"},
			},
		},
	})
}

func testAccSpannerInstance_spannerInstanceProcessingUnitsExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_spanner_instance" "example" {
  config       = "regional-us-central1"
  display_name = "Test Spanner Instance"
  processing_units    = 200
  labels = {
    "foo" = "bar"
  }
}
`, context)
}

func TestAccSpannerInstance_spannerInstanceMultiRegionalExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSpannerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSpannerInstance_spannerInstanceMultiRegionalExample(context),
			},
			{
				ResourceName:            "google_spanner_instance.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"config"},
			},
		},
	})
}

func testAccSpannerInstance_spannerInstanceMultiRegionalExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_spanner_instance" "example" {
  config       = "nam-eur-asia1"
  display_name = "Multi Regional Instance"
  num_nodes    = 2
  labels = {
    "foo" = "bar"
  }
}
`, context)
}

func testAccCheckSpannerInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_spanner_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{SpannerBasePath}}projects/{{project}}/instances/{{name}}")
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
				return fmt.Errorf("SpannerInstance still exists at %s", url)
			}
		}

		return nil
	}
}
