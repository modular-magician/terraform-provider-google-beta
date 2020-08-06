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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccComputeHaVpnGateway_haVpnGatewayBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeHaVpnGatewayDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHaVpnGateway_haVpnGatewayBasicExample(context),
			},
		},
	})
}

func testAccComputeHaVpnGateway_haVpnGatewayBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_ha_vpn_gateway" "ha_gateway1" {
  provider = google-beta
  region   = "us-central1"
  name     = "tf-test-ha-vpn-1%{random_suffix}"
  network  = google_compute_network.network1.id
}

resource "google_compute_network" "network1" {
  provider                = google-beta
  name                    = "network1%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccComputeHaVpnGateway_haVpnGatewayGcpToGcpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeHaVpnGatewayDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHaVpnGateway_haVpnGatewayGcpToGcpExample(context),
			},
		},
	})
}

func testAccComputeHaVpnGateway_haVpnGatewayGcpToGcpExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_ha_vpn_gateway" "ha_gateway1" {
  provider = google-beta
  region   = "us-central1"
  name     = "tf-test-ha-vpn-1%{random_suffix}"
  network  = google_compute_network.network1.id
}

resource "google_compute_ha_vpn_gateway" "ha_gateway2" {
  provider = google-beta
  region   = "us-central1"
  name     = "tf-test-ha-vpn-2%{random_suffix}"
  network  = google_compute_network.network2.id
}

resource "google_compute_network" "network1" {
  provider                = google-beta
  name                    = "network1%{random_suffix}"
  routing_mode            = "GLOBAL"
  auto_create_subnetworks = false
}

resource "google_compute_network" "network2" {
  provider                = google-beta
  name                    = "network2%{random_suffix}"
  routing_mode            = "GLOBAL"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "network1_subnet1" {
  provider      = google-beta
  name          = "ha-vpn-subnet-1"
  ip_cidr_range = "10.0.1.0/24"
  region        = "us-central1"
  network       = google_compute_network.network1.id
}

resource "google_compute_subnetwork" "network1_subnet2" {
  provider      = google-beta
  name          = "ha-vpn-subnet-2"
  ip_cidr_range = "10.0.2.0/24"
  region        = "us-west1"
  network       = google_compute_network.network1.id
}

resource "google_compute_subnetwork" "network2_subnet1" {
  provider      = google-beta
  name          = "ha-vpn-subnet-3"
  ip_cidr_range = "192.168.1.0/24"
  region        = "us-central1"
  network       = google_compute_network.network2.id
}

resource "google_compute_subnetwork" "network2_subnet2" {
  provider      = google-beta
  name          = "ha-vpn-subnet-4"
  ip_cidr_range = "192.168.2.0/24"
  region        = "us-east1"
  network       = google_compute_network.network2.id
}

resource "google_compute_router" "router1" {
  provider = google-beta
  name     = "ha-vpn-router1"
  network  = google_compute_network.network1.name
  bgp {
    asn = 64514
  }
}

resource "google_compute_router" "router2" {
  provider = google-beta
  name     = "ha-vpn-router2"
  network  = google_compute_network.network2.name
  bgp {
    asn = 64515
  }
}

resource "google_compute_vpn_tunnel" "tunnel1" {
  provider              = google-beta
  name                  = "ha-vpn-tunnel1"
  region                = "us-central1"
  vpn_gateway           = google_compute_ha_vpn_gateway.ha_gateway1.id
  peer_gcp_gateway      = google_compute_ha_vpn_gateway.ha_gateway2.id
  shared_secret         = "a secret message"
  router                = google_compute_router.router1.id
  vpn_gateway_interface = 0
}

resource "google_compute_vpn_tunnel" "tunnel2" {
  provider              = google-beta
  name                  = "ha-vpn-tunnel2"
  region                = "us-central1"
  vpn_gateway           = google_compute_ha_vpn_gateway.ha_gateway1.id
  peer_gcp_gateway      = google_compute_ha_vpn_gateway.ha_gateway2.id
  shared_secret         = "a secret message"
  router                = google_compute_router.router1.id
  vpn_gateway_interface = 1
}

resource "google_compute_vpn_tunnel" "tunnel3" {
  provider              = google-beta
  name                  = "ha-vpn-tunnel3"
  region                = "us-central1"
  vpn_gateway           = google_compute_ha_vpn_gateway.ha_gateway2.id
  peer_gcp_gateway      = google_compute_ha_vpn_gateway.ha_gateway1.id
  shared_secret         = "a secret message"
  router                = google_compute_router.router2.id
  vpn_gateway_interface = 0
}

resource "google_compute_vpn_tunnel" "tunnel4" {
  provider              = google-beta
  name                  = "ha-vpn-tunnel4"
  region                = "us-central1"
  vpn_gateway           = google_compute_ha_vpn_gateway.ha_gateway2.id
  peer_gcp_gateway      = google_compute_ha_vpn_gateway.ha_gateway1.id
  shared_secret         = "a secret message"
  router                = google_compute_router.router2.id
  vpn_gateway_interface = 1
}

resource "google_compute_router_interface" "router1_interface1" {
  provider   = google-beta
  name       = "router1-interface1"
  router     = google_compute_router.router1.name
  region     = "us-central1"
  ip_range   = "169.254.0.1/30"
  vpn_tunnel = google_compute_vpn_tunnel.tunnel1.name
}

resource "google_compute_router_peer" "router1_peer1" {
  provider                  = google-beta
  name                      = "router1-peer1"
  router                    = google_compute_router.router1.name
  region                    = "us-central1"
  peer_ip_address           = "169.254.0.2"
  peer_asn                  = 64515
  advertised_route_priority = 100
  interface                 = google_compute_router_interface.router1_interface1.name
}

resource "google_compute_router_interface" "router1_interface2" {
  provider   = google-beta
  name       = "router1-interface2"
  router     = google_compute_router.router1.name
  region     = "us-central1"
  ip_range   = "169.254.1.1/30"
  vpn_tunnel = google_compute_vpn_tunnel.tunnel2.name
}

resource "google_compute_router_peer" "router1_peer2" {
  provider                  = google-beta
  name                      = "router1-peer2"
  router                    = google_compute_router.router1.name
  region                    = "us-central1"
  peer_ip_address           = "169.254.1.2"
  peer_asn                  = 64515
  advertised_route_priority = 100
  interface                 = google_compute_router_interface.router1_interface2.name
}

resource "google_compute_router_interface" "router2_interface1" {
  provider   = google-beta
  name       = "router2-interface1"
  router     = google_compute_router.router2.name
  region     = "us-central1"
  ip_range   = "169.254.0.1/30"
  vpn_tunnel = google_compute_vpn_tunnel.tunnel3.name
}

resource "google_compute_router_peer" "router2_peer1" {
  provider                  = google-beta
  name                      = "router2-peer1"
  router                    = google_compute_router.router2.name
  region                    = "us-central1"
  peer_ip_address           = "169.254.0.2"
  peer_asn                  = 64514
  advertised_route_priority = 100
  interface                 = google_compute_router_interface.router2_interface1.name
}

resource "google_compute_router_interface" "router2_interface2" {
  provider   = google-beta
  name       = "router2-interface2"
  router     = google_compute_router.router2.name
  region     = "us-central1"
  ip_range   = "169.254.1.1/30"
  vpn_tunnel = google_compute_vpn_tunnel.tunnel4.name
}

resource "google_compute_router_peer" "router2_peer2" {
  provider                  = google-beta
  name                      = "router2-peer2"
  router                    = google_compute_router.router2.name
  region                    = "us-central1"
  peer_ip_address           = "169.254.1.2"
  peer_asn                  = 64514
  advertised_route_priority = 100
  interface                 = google_compute_router_interface.router2_interface2.name
}
`, context)
}

func testAccCheckComputeHaVpnGatewayDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_ha_vpn_gateway" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("ComputeHaVpnGateway still exists at %s", url)
			}
		}

		return nil
	}
}
