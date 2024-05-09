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

package networksecurity_test

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

func TestAccNetworkSecurityUrlLists_networkSecurityUrlListsBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityUrlListsDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityUrlLists_networkSecurityUrlListsBasicExample(context),
			},
			{
				ResourceName:            "google_network_security_url_lists.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name"},
			},
		},
	})
}

func testAccNetworkSecurityUrlLists_networkSecurityUrlListsBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_url_lists" "default" {
  name        = "tf-test-my-url-lists%{random_suffix}"
  location    = "us-central1"
  values = ["www.example.com"]
}
`, context)
}

func TestAccNetworkSecurityUrlLists_networkSecurityUrlListsAdvancedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityUrlListsDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityUrlLists_networkSecurityUrlListsAdvancedExample(context),
			},
			{
				ResourceName:            "google_network_security_url_lists.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name"},
			},
		},
	})
}

func testAccNetworkSecurityUrlLists_networkSecurityUrlListsAdvancedExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_url_lists" "default" {
  name        = "tf-test-my-url-lists%{random_suffix}"
  location    = "us-central1"
  description = "my description"
  values = ["www.example.com", "about.example.com", "github.com/example-org/*"]
}
`, context)
}

func testAccCheckNetworkSecurityUrlListsDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_security_url_lists" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/urlLists/{{name}}")
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
				return fmt.Errorf("NetworkSecurityUrlLists still exists at %s", url)
			}
		}

		return nil
	}
}
