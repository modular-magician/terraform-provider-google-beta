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
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccComputeRegionSslCertificate_regionSslCertificateBasicExample(t *testing.T) {
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeRegionSslCertificateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionSslCertificate_regionSslCertificateBasicExample(context),
			},
			{
				ResourceName:            "google_compute_region_ssl_certificate.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"private_key", "region", "name_prefix"},
			},
		},
	})
}

func testAccComputeRegionSslCertificate_regionSslCertificateBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_ssl_certificate" "default" {
  region      = "us-central1"
  name_prefix = "my-certificate-"
  description = "a description"
  private_key = file("test-fixtures/ssl_cert/test.key")
  certificate = file("test-fixtures/ssl_cert/test.crt")

  lifecycle {
    create_before_destroy = true
  }
}
`, context)
}

func TestAccComputeRegionSslCertificate_regionSslCertificateRandomProviderExample(t *testing.T) {
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeRegionSslCertificateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionSslCertificate_regionSslCertificateRandomProviderExample(context),
			},
			{
				ResourceName:            "google_compute_region_ssl_certificate.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"private_key", "region"},
			},
		},
	})
}

func testAccComputeRegionSslCertificate_regionSslCertificateRandomProviderExample(context map[string]interface{}) string {
	return Nprintf(`
# You may also want to control name generation explicitly:
resource "google_compute_region_ssl_certificate" "default" {
  region   = "us-central1"

  # The name will contain 8 random hex digits,
  # e.g. "my-certificate-48ab27cd2a"
  name        = random_id.certificate.hex
  private_key = file("test-fixtures/ssl_cert/test.key")
  certificate = file("test-fixtures/ssl_cert/test.crt")

  lifecycle {
    create_before_destroy = true
  }
}

resource "random_id" "certificate" {
  byte_length = 4
  prefix      = "my-certificate-"

  # For security, do not expose raw certificate values in the output
  keepers = {
    private_key = filebase64sha256("test-fixtures/ssl_cert/test.key")
    certificate = filebase64sha256("test-fixtures/ssl_cert/test.crt")
  }
}
`, context)
}

func TestAccComputeRegionSslCertificate_regionSslCertificateTargetHttpsProxiesExample(t *testing.T) {
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeRegionSslCertificateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionSslCertificate_regionSslCertificateTargetHttpsProxiesExample(context),
			},
			{
				ResourceName:            "google_compute_region_ssl_certificate.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"private_key", "region", "name_prefix"},
			},
		},
	})
}

func testAccComputeRegionSslCertificate_regionSslCertificateTargetHttpsProxiesExample(context map[string]interface{}) string {
	return Nprintf(`
// Using with Region Target HTTPS Proxies
//
// SSL certificates cannot be updated after creation. In order to apply
// the specified configuration, Terraform will destroy the existing
// resource and create a replacement. To effectively use an SSL
// certificate resource with a Target HTTPS Proxy resource, it's
// recommended to specify create_before_destroy in a lifecycle block.
// Either omit the Instance Template name attribute, specify a partial
// name with name_prefix, or use random_id resource. Example:

resource "google_compute_region_ssl_certificate" "default" {
  region      = "us-central1"
  name_prefix = "my-certificate-"
  private_key = file("test-fixtures/ssl_cert/test.key")
  certificate = file("test-fixtures/ssl_cert/test.crt")

  lifecycle {
    create_before_destroy = true
  }
}

resource "google_compute_region_target_https_proxy" "default" {
  region           = "us-central1"
  name             = "tf-test-test-proxy%{random_suffix}"
  url_map          = google_compute_region_url_map.default.id
  ssl_certificates = [google_compute_region_ssl_certificate.default.id]
}

resource "google_compute_region_url_map" "default" {
  region      = "us-central1"
  name        = "tf-test-url-map%{random_suffix}"
  description = "a description"

  default_service = google_compute_region_backend_service.default.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_region_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_region_backend_service.default.id
    }
  }
}

resource "google_compute_region_backend_service" "default" {
  region      = "us-central1"
  name        = "tf-test-backend-service%{random_suffix}"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.id]
}

resource "google_compute_region_health_check" "default" {
  region   = "us-central1"
  name     = "tf-test-http-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func testAccCheckComputeRegionSslCertificateDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_region_ssl_certificate" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/sslCertificates/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("ComputeRegionSslCertificate still exists at %s", url)
			}
		}

		return nil
	}
}
