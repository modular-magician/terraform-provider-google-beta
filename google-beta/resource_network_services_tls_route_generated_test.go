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
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccNetworkServicesTlsRoute_networkServicesTlsRouteBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesTlsRouteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesTlsRoute_networkServicesTlsRouteBasicExample(context),
			},
			{
				ResourceName:            "google_network_services_tls_route.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesTlsRoute_networkServicesTlsRouteBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  provider               = google-beta
  name          = "tf-test-my-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  provider               = google-beta
  name               = "tf-test-backend-service-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_network_services_tls_route" "default" {
  provider               = google-beta
  name                   = "tf-test-my-tls-route%{random_suffix}"
  description             = "my description"
  rules                   {
    matches {
      sni_host = ["example.com"]
      alpn = ["http/1.1"]
    }
    action {
      destinations {
        service_name = google_compute_backend_service.default.id
        weight = 1
      }
    }
  }
}
`, context)
}

func TestAccNetworkServicesTlsRoute_networkServicesTlsRouteMeshBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesTlsRouteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesTlsRoute_networkServicesTlsRouteMeshBasicExample(context),
			},
			{
				ResourceName:            "google_network_services_tls_route.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesTlsRoute_networkServicesTlsRouteMeshBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  provider               = google-beta
  name          = "tf-test-my-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  provider               = google-beta
  name               = "tf-test-backend-service-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_network_services_mesh" "default" {
  provider    = google-beta
  name        = "tf-test-my-tls-route%{random_suffix}"
  labels      = {
    foo = "bar"
  }
  description = "my description"
}


resource "google_network_services_tls_route" "default" {
  provider               = google-beta
  name                   = "tf-test-my-tls-route%{random_suffix}"
  description             = "my description"
  meshes = [
    google_network_services_mesh.default.id
  ]
  rules                   {
    matches {
      sni_host = ["example.com"]
      alpn = ["http/1.1"]
    }
    action {
      destinations {
        service_name = google_compute_backend_service.default.id
        weight = 1
      }
    }
  }
}
`, context)
}

func TestAccNetworkServicesTlsRoute_networkServicesTlsRouteGatewayBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesTlsRouteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesTlsRoute_networkServicesTlsRouteGatewayBasicExample(context),
			},
			{
				ResourceName:            "google_network_services_tls_route.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesTlsRoute_networkServicesTlsRouteGatewayBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_service" "default" {
  provider               = google-beta
  name          = "tf-test-my-backend-service%{random_suffix}"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  provider               = google-beta
  name               = "tf-test-backend-service-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_network_services_gateway" "default" {
  provider    = google-beta
  name        = "tf-test-my-tls-route%{random_suffix}"
  labels      = {
    foo = "bar"
  }
  description = "my description"
  scope = "my-scope"
  type = "OPEN_MESH"
  ports = [443]
}

resource "google_network_services_tls_route" "default" {
  provider               = google-beta
  name                   = "tf-test-my-tls-route%{random_suffix}"
  description             = "my description"
  gateways = [
    google_network_services_gateway.default.id
  ]
  rules                   {
    matches {
      sni_host = ["example.com"]
      alpn = ["http/1.1"]
    }
    action {
      destinations {
        service_name = google_compute_backend_service.default.id
        weight = 1
      }
    }
  }
}
`, context)
}

func testAccCheckNetworkServicesTlsRouteDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_services_tls_route" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := acctest.ReplaceVarsForTest(config, rs, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/tlsRoutes/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("NetworkServicesTlsRoute still exists at %s", url)
			}
		}

		return nil
	}
}
