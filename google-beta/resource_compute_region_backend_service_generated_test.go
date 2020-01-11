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

func TestAccComputeRegionBackendService_regionBackendServiceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeRegionBackendServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_regionBackendServiceBasicExample(context),
			},
			{
				ResourceName:      "google_compute_region_backend_service.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionBackendService_regionBackendServiceBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_backend_service" "default" {
  name                            = "region-backend-service%{random_suffix}"
  region                          = "us-central1"
  health_checks                   = [google_compute_health_check.default.self_link]
  connection_draining_timeout_sec = 10
  session_affinity                = "CLIENT_IP"
}

resource "google_compute_health_check" "default" {
  name               = "health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}
`, context)
}

func TestAccComputeRegionBackendService_regionBackendServiceIlbRoundRobinExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeRegionBackendServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_regionBackendServiceIlbRoundRobinExample(context),
			},
		},
	})
}

func testAccComputeRegionBackendService_regionBackendServiceIlbRoundRobinExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_backend_service" "default" {
  provider = "google-beta"

  region = "us-central1"
  name = "region-backend-service%{random_suffix}"
  health_checks = ["${google_compute_health_check.health_check.self_link}"]
  protocol = "HTTP"
  load_balancing_scheme = "INTERNAL_MANAGED"
  locality_lb_policy = "ROUND_ROBIN"
}

resource "google_compute_health_check" "health_check" {
  provider = "google-beta"

  name               = "health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeRegionBackendService_regionBackendServiceIlbRingHashExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeRegionBackendServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_regionBackendServiceIlbRingHashExample(context),
			},
		},
	})
}

func testAccComputeRegionBackendService_regionBackendServiceIlbRingHashExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_backend_service" "default" {
  provider = "google-beta"

  region = "us-central1"
  name = "region-backend-service%{random_suffix}"
  health_checks = ["${google_compute_health_check.health_check.self_link}"]
  load_balancing_scheme = "INTERNAL_MANAGED"
  locality_lb_policy = "RING_HASH"
  session_affinity = "HTTP_COOKIE"
  protocol = "HTTP"
  circuit_breakers {
    max_connections = 10
  }
  consistent_hash {
    http_cookie {
      ttl {
        seconds = 11
        nanos = 1111
      }
      name = "mycookie"
    }
  }
  outlier_detection {
    consecutive_errors = 2
  }
}

resource "google_compute_health_check" "health_check" {
  provider = "google-beta"

  name               = "health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeRegionBackendService_regionBackendServiceBalancingModeExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeRegionBackendServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_regionBackendServiceBalancingModeExample(context),
			},
		},
	})
}

func testAccComputeRegionBackendService_regionBackendServiceBalancingModeExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_backend_service" "default" {
  provider = google-beta

  load_balancing_scheme = "INTERNAL_MANAGED"

  backend {
    group = google_compute_region_instance_group_manager.rigm.instance_group
    balancing_mode = "UTILIZATION"
  }

  region      = "us-central1"
  name        = "region-backend-service%{random_suffix}"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.self_link]
}

data "google_compute_image" "debian_image" {
  provider = google-beta
  family   = "debian-9"
  project  = "debian-cloud"
}

resource "google_compute_region_instance_group_manager" "rigm" {
  provider = google-beta
  region   = "us-central1"
  name     = "rigm-internal"
  version {
    instance_template = google_compute_instance_template.instance_template.self_link
    name              = "primary"
  }
  base_instance_name = "internal-glb"
  target_size        = 1
}

resource "google_compute_instance_template" "instance_template" {
  provider     = google-beta
  name         = "template-region-backend-service%{random_suffix}"
  machine_type = "n1-standard-1"

  network_interface {
    network = google_compute_network.default.self_link
    subnetwork = google_compute_subnetwork.default.self_link
  }

  disk {
    source_image = data.google_compute_image.debian_image.self_link
    auto_delete  = true
    boot         = true
  }

  tags = ["allow-ssh", "load-balanced-backend"]
}

resource "google_compute_region_health_check" "default" {
  provider = google-beta

  region = "us-central1"
  name   = "health-check%{random_suffix}"
  http_health_check {
    port_specification = "USE_SERVING_PORT"
  }
}

resource "google_compute_network" "default" {
  provider = google-beta
  name                    = "net%{random_suffix}"
  auto_create_subnetworks = false
  routing_mode = "REGIONAL"
}

resource "google_compute_subnetwork" "default" {
  provider = google-beta
  name          = "net%{random_suffix}-default"
  ip_cidr_range = "10.1.2.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.self_link
}
`, context)
}

func testAccCheckComputeRegionBackendServiceDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_region_backend_service" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeRegionBackendService still exists at %s", url)
		}
	}

	return nil
}
