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
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccNetworkSecurityAddressGroup_networkSecurityAddressGroupsBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityAddressGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityAddressGroup_networkSecurityAddressGroupsBasicExample(context),
			},
			{
				ResourceName:            "google_network_security_address_group.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels", "parent", "name", "location"},
			},
		},
	})
}

func testAccNetworkSecurityAddressGroup_networkSecurityAddressGroupsBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_address_group" "default" {
  name        = "tf-test-my-address-groups%{random_suffix}"
  parent      = "projects/%{project}"
  location    = "us-central1"
  type        = "IPV4"
  capacity    = "100"
  items       = ["208.80.154.224/32"]
}
`, context)
}

func TestAccNetworkSecurityAddressGroup_networkSecurityAddressGroupsOrganizationBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityAddressGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityAddressGroup_networkSecurityAddressGroupsOrganizationBasicExample(context),
			},
			{
				ResourceName:            "google_network_security_address_group.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels", "parent", "name", "location"},
			},
		},
	})
}

func testAccNetworkSecurityAddressGroup_networkSecurityAddressGroupsOrganizationBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_address_group" "default" {
  name        = "tf-test-my-address-groups%{random_suffix}"
  parent      = "organizations/%{org_id}"
  location    = "us-central1"
  type        = "IPV4"
  capacity    = "100"
  items       = ["208.80.154.224/32"]
}
`, context)
}

func TestAccNetworkSecurityAddressGroup_networkSecurityAddressGroupsAdvancedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityAddressGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityAddressGroup_networkSecurityAddressGroupsAdvancedExample(context),
			},
			{
				ResourceName:            "google_network_security_address_group.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels", "parent", "name", "location"},
			},
		},
	})
}

func testAccNetworkSecurityAddressGroup_networkSecurityAddressGroupsAdvancedExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_address_group" "default" {
  name        = "tf-test-my-address-groups%{random_suffix}"
  parent      = "projects/%{project}"
  location    = "us-central1"
  description = "my description"
  type        = "IPV4"
  capacity    = "100"
  items       = ["208.80.154.224/32"]
}
`, context)
}

func testAccCheckNetworkSecurityAddressGroupDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_security_address_group" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/addressGroups/{{name}}")
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
				return fmt.Errorf("NetworkSecurityAddressGroup still exists at %s", url)
			}
		}

		return nil
	}
}
