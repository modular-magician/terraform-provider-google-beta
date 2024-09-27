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

package spanner_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccSpannerInstance_spannerInstanceBasicExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSpannerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSpannerInstance_spannerInstanceBasicExample(context),
			},
			{
				ResourceName:            "google_spanner_instance.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"config", "labels", "tags", "terraform_labels"},
			},
		},
	})
}

func testAccSpannerInstance_spannerInstanceBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_spanner_instance" "example" {
  config       = "regional-us-central1"
  display_name = "Test Spanner Instance"
  num_nodes    = 2
  edition      = "STANDARD"
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
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSpannerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSpannerInstance_spannerInstanceProcessingUnitsExample(context),
			},
			{
				ResourceName:            "google_spanner_instance.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"config", "labels", "tags", "terraform_labels"},
			},
		},
	})
}

func testAccSpannerInstance_spannerInstanceProcessingUnitsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
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

func TestAccSpannerInstance_spannerInstanceWithAutoscalingExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSpannerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSpannerInstance_spannerInstanceWithAutoscalingExample(context),
			},
			{
				ResourceName:            "google_spanner_instance.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"config", "labels", "tags", "terraform_labels"},
			},
		},
	})
}

func testAccSpannerInstance_spannerInstanceWithAutoscalingExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_spanner_instance" "example" {
  config       = "regional-us-central1"
  display_name = "Test Spanner Instance"
  autoscaling_config {
    autoscaling_limits {
      // Define the minimum and maximum compute capacity allocated to the instance
      // Either use nodes or processing units to specify the limits,
      // but should use the same unit to set both the min_limit and max_limit.
      max_processing_units            = 3000 // OR max_nodes  = 3
      min_processing_units            = 2000 // OR min_nodes = 2
    }
    autoscaling_targets {
      high_priority_cpu_utilization_percent = 75
      storage_utilization_percent           = 90
    }
  }
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
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSpannerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSpannerInstance_spannerInstanceMultiRegionalExample(context),
			},
			{
				ResourceName:            "google_spanner_instance.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"config", "labels", "tags", "terraform_labels"},
			},
		},
	})
}

func testAccSpannerInstance_spannerInstanceMultiRegionalExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
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

			config := acctest.GoogleProviderConfig(t)

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
