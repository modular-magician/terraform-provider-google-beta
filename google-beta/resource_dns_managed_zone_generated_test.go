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
)

func TestAccDNSManagedZone_dnsManagedZoneQuickstartExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsManagedZoneQuickstartExample(context),
			},
			{
				ResourceName:            "google_dns_managed_zone.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_destroy"},
			},
		},
	})
}

func testAccDNSManagedZone_dnsManagedZoneQuickstartExample(context map[string]interface{}) string {
	return Nprintf(`
# to setup a web-server
resource "google_compute_instance" "default" {
  name         = "tf-test-dns-compute-instance%{random_suffix}"
  machine_type = "g1-small"
  zone         = "us-central1-b"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
    access_config {
      // Ephemeral public IP
    }
  }
  metadata_startup_script = <<-EOF
  sudo apt-get update && \
  sudo apt-get install apache2 -y && \
  echo "<!doctype html><html><body><h1>Hello World!</h1></body></html>" > /var/www/html/index.html
  EOF
}

# to allow http traffic
resource "google_compute_firewall" "default" {
  name    = "tf-test-allow-http-traffic%{random_suffix}"
  network = "default"
  allow {
    ports    = ["80"]
    protocol = "tcp"
  }
  source_ranges = ["0.0.0.0/0"]
}

# to create a DNS zone
resource "google_dns_managed_zone" "default" {
  name          = "tf-test-example-zone-googlecloudexample%{random_suffix}"
  dns_name      = "googlecloudexample.com."
  description   = "Example DNS zone"
  force_destroy = "true"
}

# to register web-server's ip address in DNS
resource "google_dns_record_set" "default" {
  name         = google_dns_managed_zone.default.dns_name
  managed_zone = google_dns_managed_zone.default.name
  type         = "A"
  ttl          = 300
  rrdatas = [
    google_compute_instance.default.network_interface.0.access_config.0.nat_ip
  ]
}
`, context)
}

func TestAccDNSManagedZone_dnsRecordSetBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsRecordSetBasicExample(context),
			},
			{
				ResourceName:      "google_dns_managed_zone.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDNSManagedZone_dnsRecordSetBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dns_managed_zone" "parent-zone" {
  provider = "google-beta"
  name        = "tf-test-sample-zone%{random_suffix}"
  dns_name    = "tf-test-sample-zone%{random_suffix}.hashicorptest.com."
  description = "Test Description"
}

resource "google_dns_record_set" "default" {
  provider = "google-beta"
  managed_zone = google_dns_managed_zone.parent-zone.name
  name         = "test-record.tf-test-sample-zone%{random_suffix}.hashicorptest.com."
  type         = "A"
  rrdatas      = ["10.0.0.1", "10.1.0.1"]
  ttl          = 86400
}
`, context)
}

func TestAccDNSManagedZone_dnsManagedZoneBasicExample(t *testing.T) {
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsManagedZoneBasicExample(context),
			},
			{
				ResourceName:      "google_dns_managed_zone.example-zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDNSManagedZone_dnsManagedZoneBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dns_managed_zone" "example-zone" {
  name        = "example-zone"
  dns_name    = "example-${random_id.rnd.hex}.com."
  description = "Example DNS zone"
  labels = {
    foo = "bar"
  }
}

resource "random_id" "rnd" {
  byte_length = 4
}
`, context)
}

func TestAccDNSManagedZone_dnsManagedZonePrivateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsManagedZonePrivateExample(context),
			},
			{
				ResourceName:      "google_dns_managed_zone.private-zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDNSManagedZone_dnsManagedZonePrivateExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dns_managed_zone" "private-zone" {
  name        = "tf-test-private-zone%{random_suffix}"
  dns_name    = "private.example.com."
  description = "Example private DNS zone"
  labels = {
    foo = "bar"
  }

  visibility = "private"

  private_visibility_config {
    networks {
      network_url = google_compute_network.network-1.id
    }
    networks {
      network_url = google_compute_network.network-2.id
    }
  }
}

resource "google_compute_network" "network-1" {
  name                    = "tf-test-network-1%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_network" "network-2" {
  name                    = "tf-test-network-2%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccDNSManagedZone_dnsManagedZonePrivatePeeringExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsManagedZonePrivatePeeringExample(context),
			},
			{
				ResourceName:      "google_dns_managed_zone.peering-zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDNSManagedZone_dnsManagedZonePrivatePeeringExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dns_managed_zone" "peering-zone" {
  name        = "tf-test-peering-zone%{random_suffix}"
  dns_name    = "peering.example.com."
  description = "Example private DNS peering zone"

  visibility = "private"

  private_visibility_config {
    networks {
      network_url = google_compute_network.network-source.id
    }
  }

  peering_config {
    target_network {
      network_url = google_compute_network.network-target.id
    }
  }
}

resource "google_compute_network" "network-source" {
  name                    = "tf-test-network-source%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_network" "network-target" {
  name                    = "tf-test-network-target%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccDNSManagedZone_dnsManagedZoneServiceDirectoryExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsManagedZoneServiceDirectoryExample(context),
			},
			{
				ResourceName:      "google_dns_managed_zone.sd-zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDNSManagedZone_dnsManagedZoneServiceDirectoryExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dns_managed_zone" "sd-zone" {
  provider = google-beta

  name        = "tf-test-peering-zone%{random_suffix}"
  dns_name    = "services.example.com."
  description = "Example private DNS Service Directory zone"

  visibility = "private"

  service_directory_config {
    namespace {
      namespace_url = google_service_directory_namespace.example.id
    }
  }
}

resource "google_service_directory_namespace" "example" {
  provider = google-beta

  namespace_id = "example"
  location     = "us-central1"
}

resource "google_compute_network" "network" {
  provider = google-beta

  name                    = "network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func testAccCheckDNSManagedZoneDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dns_managed_zone" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{DNSBasePath}}projects/{{project}}/managedZones/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("DNSManagedZone still exists at %s", url)
			}
		}

		return nil
	}
}
