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

func TestAccGkeonpremVmwareCluster_gkeonpremVmwareClusterBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckGkeonpremVmwareClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGkeonpremVmwareCluster_gkeonpremVmwareClusterBasicExample(context),
			},
			{
				ResourceName:            "google_gkeonprem_vmware_cluster.cluster-basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccGkeonpremVmwareCluster_gkeonpremVmwareClusterBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gkeonprem_vmware_cluster" "cluster-basic" {
  provider = google-beta
  name = "cluster-basic%{random_suffix}"
  location = "us-west1"
  admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
  description = "test cluster"
  on_prem_version = "1.13.1-gke.35"
  annotations = {}
  network_config {
    service_address_cidr_blocks = ["10.96.0.0/12"]
    pod_address_cidr_blocks = ["192.168.0.0/16"]
    dhcp_ip_config {
      enabled = true
    }
  }
  control_plane_node {
     cpus = 4
     memory = 8192
     replicas = 1
  }
  load_balancer {
    vip_config {
      control_plane_vip = "10.251.133.5"
      ingress_vip = "10.251.135.19"
    }
    metal_lb_config {
      address_pools {
        pool = "ingress-ip"
        manual_assign = "true"
        addresses = ["10.251.135.19"]
        avoid_buggy_ips = true
      }
      address_pools {
        pool = "lb-test-ip"
        manual_assign = "true"
        addresses = ["10.251.135.19"]
        avoid_buggy_ips = true
      }
    }
  }
}
`, context)
}

func TestAccGkeonpremVmwareCluster_gkeonpremVmwareClusterF5lbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckGkeonpremVmwareClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGkeonpremVmwareCluster_gkeonpremVmwareClusterF5lbExample(context),
			},
			{
				ResourceName:            "google_gkeonprem_vmware_cluster.cluster-f5lb",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccGkeonpremVmwareCluster_gkeonpremVmwareClusterF5lbExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gkeonprem_vmware_cluster" "cluster-f5lb" {
  provider = google-beta  
  name = "cluster-f5lb%{random_suffix}"
  location = "us-west1"
  admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
  description = "test cluster"
  on_prem_version = "1.13.1-gke.35"
  annotations = {}
  network_config {
    service_address_cidr_blocks = ["10.96.0.0/12"]
    pod_address_cidr_blocks = ["192.168.0.0/16"]
    dhcp_ip_config {
      enabled = true
    }
    control_plane_v2_config {
      control_plane_ip_block {
        ips {
          hostname = "test-hostname"
          ip = "10.0.0.1"
        }
        netmask="10.0.0.1/32"
        gateway="test-gateway"
      }
    }
  }
  control_plane_node {
     cpus = 4
     memory = 8192
     replicas = 1
     auto_resize_config {
      enabled = true
     }
  }
  load_balancer {
    vip_config {
      control_plane_vip = "10.251.133.5"
      ingress_vip = "10.251.135.19"
    }
    f5_config {
        address = "10.0.0.1"
        partition = "test-partition"
        snat_pool = "test-snap-pool"
    }
  }
  dataplane_v2 {
    dataplane_v2_enabled = true
    windows_dataplane_v2_enabled = true
    advanced_networking = true
  }
  vm_tracking_enabled = true
  enable_control_plane_v2 = true
  authorization {
    admin_users {
      username = "testuser@gmail.com"
    }
  }
  anti_affinity_groups {
    aag_config_disabled = true
  }
  auto_repair_config {
    enabled = true
  }
  storage {
    vsphere_csi_disabled = true
  }
}
`, context)
}

func TestAccGkeonpremVmwareCluster_gkeonpremVmwareClusterManuallbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckGkeonpremVmwareClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGkeonpremVmwareCluster_gkeonpremVmwareClusterManuallbExample(context),
			},
			{
				ResourceName:            "google_gkeonprem_vmware_cluster.cluster-manuallb",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccGkeonpremVmwareCluster_gkeonpremVmwareClusterManuallbExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gkeonprem_vmware_cluster" "cluster-manuallb" {
  provider = google-beta
  name = "cluster-manuallb%{random_suffix}"
  location = "us-west1"
  admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
  description = "test cluster"
  on_prem_version = "1.13.1-gke.35"
  annotations = {}
  network_config {
    service_address_cidr_blocks = ["10.96.0.0/12"]
    pod_address_cidr_blocks = ["192.168.0.0/16"]
    host_config {
      dns_servers = ["10.254.41.1"]
      ntp_servers = ["216.239.35.8"]
      dns_search_domains = ["test-domain"]
    }
  
    static_ip_config {
      ip_blocks {
        netmask = "255.255.252.0"
        gateway = "10.251.31.254"
        ips {
          ip = "10.251.30.153"
          hostname = "test-hostname1"
        }
        ips {
          ip = "10.251.31.206"
          hostname = "test-hostname2"
        }
        ips {
          ip = "10.251.31.193"
          hostname = "test-hostname3"
        }
        ips { 
          ip = "10.251.30.230"
          hostname = "test-hostname4"
        }
      }
    }
  }
  control_plane_node {
     cpus = 4
     memory = 8192
     replicas = 1
     auto_resize_config {
      enabled = true
     }
  }
  load_balancer {
    vip_config {
      control_plane_vip = "10.251.133.5"
      ingress_vip = "10.251.135.19"
    }
    manual_lb_config {
      ingress_http_node_port = 30005
      ingress_https_node_port = 30006
      control_plane_node_port = 30007
      konnectivity_server_node_port = 30008
    }
  }
  dataplane_v2 {
    dataplane_v2_enabled = true
    windows_dataplane_v2_enabled = true
    advanced_networking = true
  }
  vm_tracking_enabled = true
  enable_control_plane_v2 = true
  authorization {
    admin_users {
      username = "testuser@gmail.com"
    }
  }
  anti_affinity_groups {
    aag_config_disabled = true
  }
  auto_repair_config {
    enabled = true
  }
}
`, context)
}

func testAccCheckGkeonpremVmwareClusterDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_gkeonprem_vmware_cluster" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{GkeonpremBasePath}}projects/{{project}}/locations/{{location}}/vmwareClusters/{{name}}")
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
				return fmt.Errorf("GkeonpremVmwareCluster still exists at %s", url)
			}
		}

		return nil
	}
}
