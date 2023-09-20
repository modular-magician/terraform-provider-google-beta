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

package compute_test

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

func TestAccComputeTargetHttpsProxy_targetHttpsProxyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeTargetHttpsProxyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeTargetHttpsProxy_targetHttpsProxyBasicExample(context),
			},
			{
				ResourceName:            "google_compute_target_https_proxy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ssl_policy", "url_map", "server_tls_policy"},
			},
		},
	})
}

func testAccComputeTargetHttpsProxy_targetHttpsProxyBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_target_https_proxy" "default" {
  name             = "tf-test-test-proxy%{random_suffix}"
  url_map          = google_compute_url_map.default.id
  ssl_certificates = [google_compute_ssl_certificate.default.id]
}

resource "google_compute_ssl_certificate" "default" {
  name        = "tf-test-my-certificate%{random_suffix}"
  private_key = file("test-fixtures/test.key")
  certificate = file("test-fixtures/test.crt")
}

resource "google_compute_url_map" "default" {
  name        = "tf-test-url-map%{random_suffix}"
  description = "a description"

  default_service = google_compute_backend_service.default.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.id
    }
  }
}

resource "google_compute_backend_service" "default" {
  name        = "tf-test-backend-service%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-http-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, context)
}

func TestAccComputeTargetHttpsProxy_targetHttpsProxyHttpKeepAliveTimeoutExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeTargetHttpsProxyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeTargetHttpsProxy_targetHttpsProxyHttpKeepAliveTimeoutExample(context),
			},
			{
				ResourceName:            "google_compute_target_https_proxy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ssl_policy", "url_map", "server_tls_policy"},
			},
		},
	})
}

func testAccComputeTargetHttpsProxy_targetHttpsProxyHttpKeepAliveTimeoutExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_target_https_proxy" "default" {
  name                        = "tf-test-test-http-keep-alive-timeout-proxy%{random_suffix}"
  http_keep_alive_timeout_sec = 610
  url_map                     = google_compute_url_map.default.id
  ssl_certificates            = [google_compute_ssl_certificate.default.id]
}

resource "google_compute_ssl_certificate" "default" {
  name        = "tf-test-my-certificate%{random_suffix}"
  private_key = file("test-fixtures/test.key")
  certificate = file("test-fixtures/test.crt")
}

resource "google_compute_url_map" "default" {
  name        = "tf-test-url-map%{random_suffix}"
  description = "a description"

  default_service = google_compute_backend_service.default.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.id
    }
  }
}

resource "google_compute_backend_service" "default" {
  name                  = "tf-test-backend-service%{random_suffix}"
  port_name             = "http"
  protocol              = "HTTP"
  timeout_sec           = 10
  load_balancing_scheme = "EXTERNAL_MANAGED"

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-http-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, context)
}

func TestAccComputeTargetHttpsProxy_targetHttpsProxyMtlsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeTargetHttpsProxyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeTargetHttpsProxy_targetHttpsProxyMtlsExample(context),
			},
			{
				ResourceName:            "google_compute_target_https_proxy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ssl_policy", "url_map", "server_tls_policy"},
			},
		},
	})
}

func testAccComputeTargetHttpsProxy_targetHttpsProxyMtlsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
  provider = google-beta
}

resource "google_compute_target_https_proxy" "default" {
  provider          = google-beta
  name              = "tf-test-test-mtls-proxy%{random_suffix}"
  url_map           = google_compute_url_map.default.id
  ssl_certificates  = [google_compute_ssl_certificate.default.id]
  server_tls_policy = google_network_security_server_tls_policy.default.id
}

resource "google_certificate_manager_trust_config" "default" {
  provider    = google-beta
  name        = "tf-test-my-trust-config%{random_suffix}"
  description = "sample description for the trust config"
  location    = "global"

  trust_stores {
    trust_anchors {
      pem_certificate = file("test-fixtures/ca_cert.pem")
    }
    intermediate_cas {
      pem_certificate = file("test-fixtures/ca_cert.pem")
    }
  }

  labels = {
    foo = "bar"
  }
}

resource "google_network_security_server_tls_policy" "default" {
  provider               = google-beta
  name                   = "tf-test-my-tls-policy%{random_suffix}"
  description            = "my description"
  location               = "global"
  allow_open             = "false"
  mtls_policy {
    client_validation_mode = "ALLOW_INVALID_OR_MISSING_CLIENT_CERT"
    client_validation_trust_config = "projects/${data.google_project.project.number}/locations/global/trustConfigs/${google_certificate_manager_trust_config.default.name}"
  }
}

resource "google_compute_ssl_certificate" "default" {
  provider    = google-beta
  name        = "tf-test-my-certificate%{random_suffix}"
  private_key = file("test-fixtures/test.key")
  certificate = file("test-fixtures/test.crt")
}

resource "google_compute_url_map" "default" {
  provider    = google-beta
  name        = "tf-test-url-map%{random_suffix}"
  description = "a description"

  default_service = google_compute_backend_service.default.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.id
    }
  }
}

resource "google_compute_backend_service" "default" {
  provider    = google-beta
  name        = "tf-test-backend-service%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  provider           = google-beta
  name               = "tf-test-http-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, context)
}

func testAccCheckComputeTargetHttpsProxyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_target_https_proxy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/targetHttpsProxies/{{name}}")
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
				return fmt.Errorf("ComputeTargetHttpsProxy still exists at %s", url)
			}
		}

		return nil
	}
}
