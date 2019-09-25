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

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccComputeRegionUrlMap_regionUrlMapBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeRegionUrlMapDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionUrlMap_regionUrlMapBasicExample(context),
			},
		},
	})
}

func testAccComputeRegionUrlMap_regionUrlMapBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_url_map" "regionurlmap" {
  provider = "google-beta"

  region      = "us-central1"

  name        = "regionurlmap%{random_suffix}"
  description = "a description"

  default_service = "${google_compute_region_backend_service.home.self_link}"

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = "${google_compute_region_backend_service.home.self_link}"

    path_rule {
      paths   = ["/home"]
      service = "${google_compute_region_backend_service.home.self_link}"
    }

    path_rule {
      paths   = ["/login"]
      service = "${google_compute_region_backend_service.login.self_link}"
    }
  }

  test {
    service = "${google_compute_region_backend_service.home.self_link}"
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_region_backend_service" "login" {
  provider = "google-beta"

  region      = "us-central1"

  name        = "login%{random_suffix}"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = ["${google_compute_region_health_check.default.self_link}"]
}

resource "google_compute_region_backend_service" "home" {
  provider = "google-beta"

  region      = "us-central1"

  name        = "home%{random_suffix}"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = ["${google_compute_region_health_check.default.self_link}"]
}

resource "google_compute_region_health_check" "default" {
  provider = "google-beta"

  region	     = "us-central1"

  name               = "health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1
  http_health_check  {
    port         = 80
    request_path = "/"
  }
}
`, context)
}

func testAccCheckComputeRegionUrlMapDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_region_url_map" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/urlMaps/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeRegionUrlMap still exists at %s", url)
		}
	}

	return nil
}
