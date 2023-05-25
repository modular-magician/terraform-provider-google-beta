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

func TestAccNetworkServicesGateway_networkServicesGatewayBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesGatewayDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesGateway_networkServicesGatewayBasicExample(context),
			},
			{
				ResourceName:            "google_network_services_gateway.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccNetworkServicesGateway_networkServicesGatewayBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_network_services_gateway" "default" {
  provider = google-beta
  name     = "tf-test-my-gateway%{random_suffix}"
  scope    = "default-scope-basic"
  type     = "OPEN_MESH"
  ports    = [443]
}
`, context)
}

func TestAccNetworkServicesGateway_networkServicesGatewayAdvancedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesGatewayDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesGateway_networkServicesGatewayAdvancedExample(context),
			},
			{
				ResourceName:            "google_network_services_gateway.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccNetworkServicesGateway_networkServicesGatewayAdvancedExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_network_services_gateway" "default" {
  provider    = google-beta
  name        = "tf-test-my-gateway%{random_suffix}"
  labels      = {
    foo = "bar"
  }
  description = "my description"
  type        = "OPEN_MESH"
  ports       = [443]
  scope       = "default-scope-advance"
}
`, context)
}

func testAccCheckNetworkServicesGatewayDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_services_gateway" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkServicesBasePath}}projects/{{project}}/locations/{{location}}/gateways/{{name}}")
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
				return fmt.Errorf("NetworkServicesGateway still exists at %s", url)
			}
		}

		return nil
	}
}
