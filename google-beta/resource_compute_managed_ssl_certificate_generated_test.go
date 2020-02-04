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

func TestAccComputeManagedSslCertificate_managedSslCertificateBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeManagedSslCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeManagedSslCertificate_managedSslCertificateBasicExample(context),
			},
		},
	})
}

func testAccComputeManagedSslCertificate_managedSslCertificateBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_managed_ssl_certificate" "default" {
  provider = google-beta

  name = "tf-test-test-cert%{random_suffix}"

  managed {
    domains = ["sslcert.tf-test.club."]
  }
}

resource "google_compute_target_https_proxy" "default" {
  provider = google-beta

  name             = "tf-test-test-proxy%{random_suffix}"
  url_map          = google_compute_url_map.default.self_link
  ssl_certificates = [google_compute_managed_ssl_certificate.default.self_link]
}

resource "google_compute_url_map" "default" {
  provider = google-beta

  name        = "tf-test-url-map%{random_suffix}"
  description = "a description"

  default_service = google_compute_backend_service.default.self_link

  host_rule {
    hosts        = ["sslcert.tf-test.club"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.self_link

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.self_link
    }
  }
}

resource "google_compute_backend_service" "default" {
  provider = google-beta

  name        = "tf-test-backend-service%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.self_link]
}

resource "google_compute_http_health_check" "default" {
  provider = google-beta

  name               = "tf-test-http-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_dns_managed_zone" "zone" {
  provider = google-beta

  name     = "tf-test-dnszone%{random_suffix}"
  dns_name = "sslcert.tf-test.club."
}

resource "google_compute_global_forwarding_rule" "default" {
  provider = google-beta

  name       = "tf-test-forwarding-rule%{random_suffix}"
  target     = google_compute_target_https_proxy.default.self_link
  port_range = 443
}

resource "google_dns_record_set" "set" {
  provider = google-beta

  name         = "sslcert.tf-test.club."
  type         = "A"
  ttl          = 3600
  managed_zone = google_dns_managed_zone.zone.name
  rrdatas      = [google_compute_global_forwarding_rule.default.ip_address]
}

provider "google-beta" {
  region = "us-central1"
  zone   = "us-central1-a"
}
`, context)
}

func testAccCheckComputeManagedSslCertificateDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_managed_ssl_certificate" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/sslCertificates/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeManagedSslCertificate still exists at %s", url)
		}
	}

	return nil
}
