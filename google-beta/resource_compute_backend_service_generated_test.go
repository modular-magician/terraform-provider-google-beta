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

func TestAccComputeBackendService_backendServiceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_backendServiceBasicExample(context),
			},
			{
				ResourceName:      "google_compute_backend_service.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendService_backendServiceBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, context)
}

func TestAccComputeBackendService_backendServiceCacheSimpleExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_backendServiceCacheSimpleExample(context),
			},
			{
				ResourceName:      "google_compute_backend_service.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendService_backendServiceCacheSimpleExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
  enable_cdn  = true
  cdn_policy {
    signed_url_cache_max_age_sec = 7200
  }
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, context)
}

func TestAccComputeBackendService_backendServiceCacheIncludeHttpHeadersExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_backendServiceCacheIncludeHttpHeadersExample(context),
			},
			{
				ResourceName:      "google_compute_backend_service.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendService_backendServiceCacheIncludeHttpHeadersExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  enable_cdn  = true
  cdn_policy {
    cache_mode = "USE_ORIGIN_HEADERS"
    cache_key_policy {
      include_host = true
      include_protocol = true
      include_query_string = true
      include_http_headers = ["X-My-Header-Field"]
    }
  }
}
`, context)
}

func TestAccComputeBackendService_backendServiceCacheIncludeNamedCookiesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_backendServiceCacheIncludeNamedCookiesExample(context),
			},
			{
				ResourceName:      "google_compute_backend_service.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendService_backendServiceCacheIncludeNamedCookiesExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  enable_cdn  = true
  cdn_policy {
    cache_mode = "CACHE_ALL_STATIC"
    default_ttl = 3600
    client_ttl  = 7200
    max_ttl     = 10800
    cache_key_policy {
      include_host = true
      include_protocol = true
      include_query_string = true
      include_named_cookies = ["__next_preview_data", "__prerender_bypass"]
    }
  }
}
`, context)
}

func TestAccComputeBackendService_backendServiceCacheExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_backendServiceCacheExample(context),
			},
			{
				ResourceName:      "google_compute_backend_service.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendService_backendServiceCacheExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
  enable_cdn  = true
  cdn_policy {
    cache_mode = "CACHE_ALL_STATIC"
    default_ttl = 3600
    client_ttl  = 7200
    max_ttl     = 10800
    negative_caching = true
    signed_url_cache_max_age_sec = 7200
  }
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, context)
}

func TestAccComputeBackendService_backendServiceCacheBypassCacheOnRequestHeadersExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_backendServiceCacheBypassCacheOnRequestHeadersExample(context),
			},
			{
				ResourceName:      "google_compute_backend_service.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendService_backendServiceCacheBypassCacheOnRequestHeadersExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
  enable_cdn  = true
  cdn_policy {
    cache_mode = "CACHE_ALL_STATIC"
    default_ttl = 3600
    client_ttl  = 7200
    max_ttl     = 10800
    negative_caching = true
    signed_url_cache_max_age_sec = 7200

    bypass_cache_on_request_headers {
      header_name = "Authorization"
    }

    bypass_cache_on_request_headers {
      header_name = "Proxy-Authorization"
    }
  }
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, context)
}

func TestAccComputeBackendService_backendServiceTrafficDirectorRoundRobinExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_backendServiceTrafficDirectorRoundRobinExample(context),
			},
			{
				ResourceName:      "google_compute_backend_service.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendService_backendServiceTrafficDirectorRoundRobinExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  provider = google-beta

  name                  = "tf-test-backend-service%{random_suffix}"
  health_checks         = [google_compute_health_check.health_check.id]
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
  locality_lb_policy    = "ROUND_ROBIN"
}

resource "google_compute_health_check" "health_check" {
  provider = google-beta

  name = "tf-test-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeBackendService_backendServiceTrafficDirectorRingHashExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_backendServiceTrafficDirectorRingHashExample(context),
			},
			{
				ResourceName:      "google_compute_backend_service.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendService_backendServiceTrafficDirectorRingHashExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  provider = google-beta

  name                  = "tf-test-backend-service%{random_suffix}"
  health_checks         = [google_compute_health_check.health_check.id]
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
  locality_lb_policy    = "RING_HASH"
  session_affinity      = "HTTP_COOKIE"
  circuit_breakers {
    max_connections = 10
  }
  consistent_hash {
    http_cookie {
      ttl {
        seconds = 11
        nanos   = 1111
      }
      name = "mycookie"
    }
  }
  outlier_detection {
    consecutive_errors = 2
  }
}

resource "google_compute_health_check" "health_check" {
  provider = google-beta

  name = "tf-test-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeBackendService_backendServiceNetworkEndpointExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_backendServiceNetworkEndpointExample(context),
			},
			{
				ResourceName:      "google_compute_backend_service.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendService_backendServiceNetworkEndpointExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_global_network_endpoint_group" "external_proxy" {
  provider = google-beta
  name                  = "tf-test-network-endpoint%{random_suffix}"
  network_endpoint_type = "INTERNET_FQDN_PORT"
  default_port          = "443"
}

resource "google_compute_global_network_endpoint" "proxy" {
  provider = google-beta
  global_network_endpoint_group = google_compute_global_network_endpoint_group.external_proxy.id
  fqdn                          = "test.example.com"
  port                          = google_compute_global_network_endpoint_group.external_proxy.default_port
}

resource "google_compute_backend_service" "default" {
  provider = google-beta
  name                            = "tf-test-backend-service%{random_suffix}"
  enable_cdn                      = true
  timeout_sec                     = 10
  connection_draining_timeout_sec = 10

  custom_request_headers          = ["host: ${google_compute_global_network_endpoint.proxy.fqdn}"]
  custom_response_headers         = ["X-Cache-Hit: {cdn_cache_status}"]

  backend {
    group = google_compute_global_network_endpoint_group.external_proxy.id
  }
}
`, context)
}

func TestAccComputeBackendService_backendServiceExternalManagedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendService_backendServiceExternalManagedExample(context),
			},
			{
				ResourceName:      "google_compute_backend_service.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendService_backendServiceExternalManagedExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_backend_service" "default" {
  name          = "tf-test-backend-service%{random_suffix}"
  health_checks = [google_compute_health_check.default.id]
  load_balancing_scheme = "EXTERNAL_MANAGED"
}

resource "google_compute_health_check" "default" {
  name = "tf-test-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func testAccCheckComputeBackendServiceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_backend_service" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/backendServices/{{name}}")
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
				return fmt.Errorf("ComputeBackendService still exists at %s", url)
			}
		}

		return nil
	}
}
