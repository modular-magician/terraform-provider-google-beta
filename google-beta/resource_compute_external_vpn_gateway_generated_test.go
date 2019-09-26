// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccComputeExternalVpnGateway_externalVpnGatewayExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeExternalVpnGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeExternalVpnGateway_externalVpnGatewayExample(context),
			},
		},
	})
}

func testAccComputeExternalVpnGateway_externalVpnGatewayExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_ha_vpn_gateway" "ha_gateway" {
  provider = "google-beta"
  region   = "us-central1"
  name     = "ha-vpn%{random_suffix}"
  network  = "${google_compute_network.network.self_link}"
}

resource "google_compute_external_vpn_gateway" "external_gateway" {
  provider        = "google-beta"
  name            = "external-gateway%{random_suffix}"
  redundancy_type = "SINGLE_IP_INTERNALLY_REDUNDANT"
  description     = "An externally managed VPN gateway"
  interface {
    id = 0
    ip_address = "8.8.8.8"
  }
}

resource "google_compute_network" "network" {
  provider                = "google-beta"
  name                    = "network%{random_suffix}"
  routing_mode            = "GLOBAL"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "network_subnet1" {
  provider = "google-beta"
  name          = "ha-vpn-subnet-1"
  ip_cidr_range = "10.0.1.0/24"
  region        = "us-central1"
  network       = "${google_compute_network.network.self_link}"
}

resource "google_compute_subnetwork" "network_subnet2" {
  provider = "google-beta"
  name          = "ha-vpn-subnet-2"
  ip_cidr_range = "10.0.2.0/24"
  region        = "us-west1"
  network       = "${google_compute_network.network.self_link}"
}

resource "google_compute_router" "router1" {
  provider = "google-beta"
  name    = "ha-vpn-router1"
  network = "${google_compute_network.network.name}"
  bgp {
    asn = 64514
  }
}

resource "google_compute_vpn_tunnel" "tunnel1" {
  provider         = "google-beta"
  name             = "ha-vpn-tunnel1"
  region           = "us-central1"
  vpn_gateway      = "${google_compute_ha_vpn_gateway.ha_gateway.self_link}"
  peer_external_gateway = "${google_compute_external_vpn_gateway.external_gateway.self_link}"
  peer_external_gateway_interface = 0
  shared_secret    = "a secret message"
  router           = "${google_compute_router.router1.self_link}"
  vpn_gateway_interface = 0
}

resource "google_compute_vpn_tunnel" "tunnel2" {
  provider         = "google-beta"
  name             = "ha-vpn-tunnel2"
  region           = "us-central1"
  vpn_gateway      = "${google_compute_ha_vpn_gateway.ha_gateway.self_link}"
  peer_external_gateway = "${google_compute_external_vpn_gateway.external_gateway.self_link}"
  peer_external_gateway_interface = 0
  shared_secret    = "a secret message"
  router           = " ${google_compute_router.router1.self_link}"
  vpn_gateway_interface = 1
}

resource "google_compute_router_interface" "router1_interface1" {
  provider = "google-beta"
  name       = "router1-interface1"
  router     = "${google_compute_router.router1.name}"
  region     = "us-central1"
  ip_range   = "169.254.0.1/30"
  vpn_tunnel = "${google_compute_vpn_tunnel.tunnel1.name}"
}

resource "google_compute_router_peer" "router1_peer1" {
  provider = "google-beta"
  name                      = "router1-peer1"
  router                    = "${google_compute_router.router1.name}"
  region                    = "us-central1"
  peer_ip_address           = "169.254.0.2"
  peer_asn                  = 64515
  advertised_route_priority = 100
  interface                 = "${google_compute_router_interface.router1_interface1.name}"
}

resource "google_compute_router_interface" "router1_interface2" {
  provider = "google-beta"
  name       = "router1-interface2"
  router     = "${google_compute_router.router1.name}"
  region     = "us-central1"
  ip_range   = "169.254.1.1/30"
  vpn_tunnel = "${google_compute_vpn_tunnel.tunnel2.name}"
}

resource "google_compute_router_peer" "router1_peer2" {
  provider = "google-beta"
  name                      = "router1-peer2"
  router                    = "${google_compute_router.router1.name}"
  region                    = "us-central1"
  peer_ip_address           = "169.254.1.2"
  peer_asn                  = 64515
  advertised_route_priority = 100
  interface                 = "${google_compute_router_interface.router1_interface2.name}"
}
`, context)
}

func testAccCheckComputeExternalVpnGatewayDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_external_vpn_gateway" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/externalVpnGateways/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeExternalVpnGateway still exists at %s", url)
		}
	}

	return nil
}
