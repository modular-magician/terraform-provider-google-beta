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

package vpcaccess_test

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

func TestAccVPCAccessConnector_vpcAccessConnectorExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckVPCAccessConnectorDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCAccessConnector_vpcAccessConnectorExample(context),
			},
			{
				ResourceName:            "google_vpc_access_connector.connector",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "self_link"},
			},
		},
	})
}

func testAccVPCAccessConnector_vpcAccessConnectorExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vpc_access_connector" "connector" {
  name          = "tf-test-vpc-con%{random_suffix}"
  ip_cidr_range = "10.8.0.0/28"
  network       = "default"
}
`, context)
}

func TestAccVPCAccessConnector_vpcAccessConnectorSharedVpcExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckVPCAccessConnectorDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCAccessConnector_vpcAccessConnectorSharedVpcExample(context),
			},
			{
				ResourceName:            "google_vpc_access_connector.connector",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "self_link"},
			},
		},
	})
}

func testAccVPCAccessConnector_vpcAccessConnectorSharedVpcExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vpc_access_connector" "connector" {
  name          = "tf-test-vpc-con%{random_suffix}"
  subnet {
    name = google_compute_subnetwork.custom_test.name
  }
  machine_type = "e2-standard-4"
}

resource "google_compute_subnetwork" "custom_test" {
  name          = "tf-test-vpc-con%{random_suffix}"
  ip_cidr_range = "10.2.0.0/28"
  region        = "us-central1"
  network       = google_compute_network.custom_test.id
}

resource "google_compute_network" "custom_test" {
  name                    = "tf-test-vpc-con%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func testAccCheckVPCAccessConnectorDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_vpc_access_connector" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{VPCAccessBasePath}}projects/{{project}}/locations/{{region}}/connectors/{{name}}")
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
				return fmt.Errorf("VPCAccessConnector still exists at %s", url)
			}
		}

		return nil
	}
}
