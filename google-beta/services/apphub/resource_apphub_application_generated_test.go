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

package apphub_test

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

func TestAccApphubApplication_apphubApplicationBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"location":      "us-east1",
		"scope_type":    "REGIONAL",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckApphubApplicationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApphubApplication_apphubApplicationBasicExample(context),
			},
			{
				ResourceName:            "google_apphub_application.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"application_id", "location"},
			},
		},
	})
}

func testAccApphubApplication_apphubApplicationBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_apphub_application" "example" {
  location = "%{location}"
  application_id = "tf-test-example-application%{random_suffix}"
  scope {
    type = "%{scope_type}"
  }
}
`, context)
}

func TestAccApphubApplication_apphubApplicationGlobalBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"location":      "global",
		"scope_type":    "GLOBAL",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckApphubApplicationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApphubApplication_apphubApplicationGlobalBasicExample(context),
			},
			{
				ResourceName:            "google_apphub_application.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"application_id", "location"},
			},
		},
	})
}

func testAccApphubApplication_apphubApplicationGlobalBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_apphub_application" "example" {
  location = "%{location}"
  application_id = "tf-test-example-application%{random_suffix}"
  scope {
    type = "%{scope_type}"
  }
}
`, context)
}

func TestAccApphubApplication_apphubApplicationFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckApphubApplicationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApphubApplication_apphubApplicationFullExample(context),
			},
			{
				ResourceName:            "google_apphub_application.example2",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"application_id", "location"},
			},
		},
	})
}

func testAccApphubApplication_apphubApplicationFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_apphub_application" "example2" {
  location = "us-east1"
  application_id = "tf-test-example-application%{random_suffix}"
  display_name = "Application Full%{random_suffix}"
  scope {
    type = "REGIONAL"
  }
  description = "Application for testing%{random_suffix}"
  attributes {
    environment {
      type = "STAGING"
		}
		criticality {  
      type = "MISSION_CRITICAL"
		}
		business_owners {
		  display_name =  "Alice%{random_suffix}"
		  email        =  "alice@google.com%{random_suffix}"
		}
		developer_owners {
		  display_name =  "Bob%{random_suffix}"
		  email        =  "bob@google.com%{random_suffix}"
		}
		operator_owners {
		  display_name =  "Charlie%{random_suffix}"
		  email        =  "charlie@google.com%{random_suffix}"
		}
  }
}
`, context)
}

func testAccCheckApphubApplicationDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_apphub_application" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ApphubBasePath}}projects/{{project}}/locations/{{location}}/applications/{{application_id}}")
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
				return fmt.Errorf("ApphubApplication still exists at %s", url)
			}
		}

		return nil
	}
}
