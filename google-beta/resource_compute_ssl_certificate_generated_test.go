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

func TestAccComputeSslCertificate_sslCertificateBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeSslCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSslCertificate_sslCertificateBasicExample(context),
			},
			{
				ResourceName:            "google_compute_ssl_certificate.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"private_key", "name_prefix"},
			},
		},
	})
}

func testAccComputeSslCertificate_sslCertificateBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_ssl_certificate" "default" {
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

func TestAccComputeSslCertificate_sslCertificateRandomProviderExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeSslCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSslCertificate_sslCertificateRandomProviderExample(context),
			},
			{
				ResourceName:            "google_compute_ssl_certificate.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"private_key"},
			},
		},
	})
}

func testAccComputeSslCertificate_sslCertificateRandomProviderExample(context map[string]interface{}) string {
	return Nprintf(`
# You may also want to control name generation explicitly:
resource "google_compute_ssl_certificate" "default" {
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

func TestAccComputeSslCertificate_sslCertificateTargetHttpsProxiesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeSslCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSslCertificate_sslCertificateTargetHttpsProxiesExample(context),
			},
			{
				ResourceName:            "google_compute_ssl_certificate.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"private_key", "name_prefix"},
			},
		},
	})
}

func testAccComputeSslCertificate_sslCertificateTargetHttpsProxiesExample(context map[string]interface{}) string {
	return Nprintf(`
// Using with Target HTTPS Proxies
//
// SSL certificates cannot be updated after creation. In order to apply
// the specified configuration, Terraform will destroy the existing
// resource and create a replacement. To effectively use an SSL
// certificate resource with a Target HTTPS Proxy resource, it's
// recommended to specify create_before_destroy in a lifecycle block.
// Either omit the Instance Template name attribute, specify a partial
// name with name_prefix, or use random_id resource. Example:

resource "google_compute_ssl_certificate" "default" {
  name_prefix = "my-certificate-"
  private_key = file("test-fixtures/ssl_cert/test.key")
  certificate = file("test-fixtures/ssl_cert/test.crt")

  lifecycle {
    create_before_destroy = true
  }
}

resource "google_compute_target_https_proxy" "default" {
  name             = "tf-test-test-proxy%{random_suffix}"
  url_map          = google_compute_url_map.default.self_link
  ssl_certificates = [google_compute_ssl_certificate.default.self_link]
}

resource "google_compute_url_map" "default" {
  name        = "tf-test-url-map%{random_suffix}"
  description = "a description"

  default_service = google_compute_backend_service.default.self_link

  host_rule {
    hosts        = ["mysite.com"]
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
  name        = "tf-test-backend-service%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.self_link]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-http-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, context)
}

func testAccCheckComputeSslCertificateDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_ssl_certificate" {
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
			return fmt.Errorf("ComputeSslCertificate still exists at %s", url)
		}
	}

	return nil
}
