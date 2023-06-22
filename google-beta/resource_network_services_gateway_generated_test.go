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

func TestAccNetworkServicesGateway_networkServicesGatewaySecureWebProxyExample(t *testing.T) {
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
				Config: testAccNetworkServicesGateway_networkServicesGatewaySecureWebProxyExample(context),
			},
			{
				ResourceName:            "google_network_services_gateway.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "delete_swg_autogen_router_on_destroy"},
			},
		},
	})
}

func testAccNetworkServicesGateway_networkServicesGatewaySecureWebProxyExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_certificate_manager_certificate" "default" {
  provider    = google-beta
  name        = "tf-test-my-certificate%{random_suffix}"
  location    = "us-central1"
  self_managed {
    pem_certificate = file("test-fixtures/certificatemanager/cert.pem")
    pem_private_key = file("test-fixtures/certificatemanager/private-key.pem")
  }
}

resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-my-network%{random_suffix}"
  routing_mode            = "REGIONAL"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-my-subnetwork-name%{random_suffix}"
  purpose       = "PRIVATE"
  ip_cidr_range = "10.128.0.0/20"
  region        = "us-central1"
  network       = google_compute_network.default.id
  role          = "ACTIVE"
}

resource "google_compute_subnetwork" "proxyonlysubnet" {
  provider      = google-beta
  name          = "tf-test-my-proxy-only-subnetwork%{random_suffix}"
  purpose       = "REGIONAL_MANAGED_PROXY"
  ip_cidr_range = "192.168.0.0/23"
  region        = "us-central1"
  network       = google_compute_network.default.id
  role          = "ACTIVE"
}

resource "google_network_security_gateway_security_policy" "default" {
  provider    = google-beta
  name        = "tf-test-my-policy-name%{random_suffix}"
  location    = "us-central1"
}

resource "google_network_security_gateway_security_policy_rule" "default" {
  provider                = google-beta
  name                    = "tf-test-my-policyrule-name%{random_suffix}"
  location                = "us-central1"
  gateway_security_policy = google_network_security_gateway_security_policy.default.name
  enabled                 = true  
  priority                = 1
  session_matcher         = "host() == 'example.com'"
  basic_profile           = "ALLOW"
}

resource "google_network_services_gateway" "default" {
  provider                             = google-beta
  name                                 = "tf-test-my-gateway1%{random_suffix}"
  location                             = "us-central1"
  addresses                            = ["10.128.0.99"]
  type                                 = "SECURE_WEB_GATEWAY"
  ports                                = [443]
  scope                                = "tf-test-my-default-scope1%{random_suffix}"
  certificate_urls                     = [google_certificate_manager_certificate.default.id]
  gateway_security_policy              = google_network_security_gateway_security_policy.default.id
  network                              = google_compute_network.default.id
  subnetwork                           = google_compute_subnetwork.default.id
  delete_swg_autogen_router_on_destroy = true
  depends_on                           = [google_compute_subnetwork.proxyonlysubnet]
}
`, context)
}

func TestAccNetworkServicesGateway_networkServicesGatewayMultipleSwpSameNetworkExample(t *testing.T) {
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
				Config: testAccNetworkServicesGateway_networkServicesGatewayMultipleSwpSameNetworkExample(context),
			},
			{
				ResourceName:            "google_network_services_gateway.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "delete_swg_autogen_router_on_destroy"},
			},
		},
	})
}

func testAccNetworkServicesGateway_networkServicesGatewayMultipleSwpSameNetworkExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_certificate_manager_certificate" "default" {
  provider    = google-beta
  name        = "tf-test-my-certificate%{random_suffix}"
  location    = "us-south1"
  self_managed {
    pem_certificate = file("test-fixtures/certificatemanager/cert.pem")
    pem_private_key = file("test-fixtures/certificatemanager/private-key.pem")
  }
}

resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-my-network%{random_suffix}"
  routing_mode            = "REGIONAL"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-my-subnetwork-name%{random_suffix}"
  purpose       = "PRIVATE"
  ip_cidr_range = "10.128.0.0/20"
  region        = "us-south1"
  network       = google_compute_network.default.id
  role          = "ACTIVE"
}

resource "google_compute_subnetwork" "proxyonlysubnet" {
  provider      = google-beta
  name          = "tf-test-my-proxy-only-subnetwork%{random_suffix}"
  purpose       = "REGIONAL_MANAGED_PROXY"
  ip_cidr_range = "192.168.0.0/23"
  region        = "us-south1"
  network       = google_compute_network.default.id
  role          = "ACTIVE"
}

resource "google_network_security_gateway_security_policy" "default" {
  provider    = google-beta
  name        = "tf-test-my-policy-name%{random_suffix}"
  location    = "us-south1"
}

resource "google_network_security_gateway_security_policy_rule" "default" {
  provider                = google-beta
  name                    = "tf-test-my-policyrule-name%{random_suffix}"
  location                = "us-south1"
  gateway_security_policy = google_network_security_gateway_security_policy.default.name
  enabled                 = true  
  priority                = 1
  session_matcher         = "host() == 'example.com'"
  basic_profile           = "ALLOW"
}

resource "google_network_services_gateway" "default" {
  provider                             = google-beta
  name                                 = "tf-test-my-gateway1%{random_suffix}"
  location                             = "us-south1"
  addresses                            = ["10.128.0.99"]
  type                                 = "SECURE_WEB_GATEWAY"
  ports                                = [443]
  scope                                = "tf-test-my-default-scope1%{random_suffix}"
  certificate_urls                     = [google_certificate_manager_certificate.default.id]
  gateway_security_policy              = google_network_security_gateway_security_policy.default.id
  network                              = google_compute_network.default.id
  subnetwork                           = google_compute_subnetwork.default.id
  delete_swg_autogen_router_on_destroy = true
  depends_on                           = [google_compute_subnetwork.proxyonlysubnet]
}

resource "google_network_services_gateway" "gateway2" {
  provider                             = google-beta
  name                                 = "tf-test-my-gateway2%{random_suffix}"
  location                             = "us-south1"
  addresses                            = ["10.128.0.98"]
  type                                 = "SECURE_WEB_GATEWAY"
  ports                                = [443]
  scope                                = "tf-test-my-default-scope2%{random_suffix}"
  certificate_urls                     = [google_certificate_manager_certificate.default.id]
  gateway_security_policy              = google_network_security_gateway_security_policy.default.id
  network                              = google_compute_network.default.id
  subnetwork                           = google_compute_subnetwork.default.id
  delete_swg_autogen_router_on_destroy = true
  depends_on                           = [google_compute_subnetwork.proxyonlysubnet]
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
