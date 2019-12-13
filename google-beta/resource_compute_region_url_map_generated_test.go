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
  provider = google-beta

  region = "us-central1"

  name        = "regionurlmap%<random_suffix>s"
  description = "a description"

  default_service = google_compute_region_backend_service.home.self_link

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_region_backend_service.home.self_link

    path_rule {
      paths   = ["/home"]
      service = google_compute_region_backend_service.home.self_link
    }

    path_rule {
      paths   = ["/login"]
      service = google_compute_region_backend_service.login.self_link
    }
  }

  test {
    service = google_compute_region_backend_service.home.self_link
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_region_backend_service" "login" {
  provider = google-beta

  region = "us-central1"

  name        = "login%<random_suffix>s"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.self_link]
}

resource "google_compute_region_backend_service" "home" {
  provider = google-beta

  region = "us-central1"

  name        = "home%<random_suffix>s"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.self_link]
}

resource "google_compute_region_health_check" "default" {
  provider = google-beta

  region = "us-central1"

  name               = "health-check%<random_suffix>s"
  check_interval_sec = 1
  timeout_sec        = 1
  http_health_check {
    port         = 80
    request_path = "/"
  }
}
`, context)
}

func TestAccComputeRegionUrlMap_regionUrlMapL7IlbPathExample(t *testing.T) {
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
				Config: testAccComputeRegionUrlMap_regionUrlMapL7IlbPathExample(context),
			},
		},
	})
}

func testAccComputeRegionUrlMap_regionUrlMapL7IlbPathExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_url_map" "regionurlmap" {
  provider = "google-beta"
  name        = "regionurlmap%<random_suffix>s"
  description = "a description"
  default_service = google_compute_region_backend_service.home.self_link

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = google_compute_region_backend_service.home.self_link

    path_rule {
      paths   = ["/home"]
      route_action {
        cors_policy {
          allow_credentials = true
          allow_headers = ["Allowed content"]
          allow_methods = ["GET"]
          allow_origins = ["Allowed origin"]
          expose_headers = ["Exposed header"]
          max_age = 30
          disabled = false
        }
        fault_injection_policy {
          abort {
            http_status = 234
            percentage = 5.6
          }
          delay {
            fixed_delay {
              seconds = 0
              nanos = 50000
            }
            percentage = 7.8
          }
        }
        request_mirror_policy {
          backend_service = google_compute_region_backend_service.home.self_link
        }
        retry_policy {
          num_retries = 4
          per_try_timeout {
            seconds = 30
          }
          retry_conditions = ["5xx", "deadline-exceeded"]
        }
        timeout {
          seconds = 20
          nanos = 750000000
        }
        url_rewrite {
          host_rewrite = "A replacement header"
          path_prefix_rewrite = "A replacement path"
        }
        weighted_backend_services {
          backend_service = google_compute_region_backend_service.home.self_link
          weight = 400
          header_action {
            request_headers_to_remove = ["RemoveMe"]
            request_headers_to_add {
              header_name = "AddMe"
              header_value = "MyValue"
              replace = true
            }
            response_headers_to_remove = ["RemoveMe"]
            response_headers_to_add {
              header_name = "AddMe"
              header_value = "MyValue"
              replace = false
            }
          }
        }
      }
    }
  }

  test {
    service = google_compute_region_backend_service.home.self_link
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_region_backend_service" "home" {
  provider = "google-beta"
  name        = "home%<random_suffix>s"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.self_link]
  load_balancing_scheme = "INTERNAL_MANAGED"
}

resource "google_compute_region_health_check" "default" {
  provider = "google-beta"
  name               = "health-check%<random_suffix>s"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeRegionUrlMap_regionUrlMapL7IlbPathPartialExample(t *testing.T) {
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
				Config: testAccComputeRegionUrlMap_regionUrlMapL7IlbPathPartialExample(context),
			},
		},
	})
}

func testAccComputeRegionUrlMap_regionUrlMapL7IlbPathPartialExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_url_map" "regionurlmap" {
  provider = "google-beta"
  name        = "regionurlmap%<random_suffix>s"
  description = "a description"
  default_service = google_compute_region_backend_service.home.self_link

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = google_compute_region_backend_service.home.self_link

    path_rule {
      paths   = ["/home"]
      route_action {
        retry_policy {
          num_retries = 4
          per_try_timeout {
            seconds = 30
          }
          retry_conditions = ["5xx", "deadline-exceeded"]
        }
        timeout {
          seconds = 20
          nanos = 750000000
        }
        url_rewrite {
          host_rewrite = "A replacement header"
          path_prefix_rewrite = "A replacement path"
        }
        weighted_backend_services {
          backend_service = google_compute_region_backend_service.home.self_link
          weight = 400
          header_action {
            response_headers_to_add {
              header_name = "AddMe"
              header_value = "MyValue"
              replace = false
            }
          }
        }
      }
    }
  }

  test {
    service = google_compute_region_backend_service.home.self_link
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_region_backend_service" "home" {
  provider = "google-beta"
  name        = "home%<random_suffix>s"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.self_link]
  load_balancing_scheme = "INTERNAL_MANAGED"
}

resource "google_compute_region_health_check" "default" {
  provider = "google-beta"
  name               = "health-check%<random_suffix>s"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeRegionUrlMap_regionUrlMapL7IlbRouteExample(t *testing.T) {
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
				Config: testAccComputeRegionUrlMap_regionUrlMapL7IlbRouteExample(context),
			},
		},
	})
}

func testAccComputeRegionUrlMap_regionUrlMapL7IlbRouteExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_url_map" "regionurlmap" {
  provider        = "google-beta"
  name            = "regionurlmap%<random_suffix>s"
  description     = "a description"
  default_service = google_compute_region_backend_service.home.self_link

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = google_compute_region_backend_service.home.self_link

    route_rules {
      priority = 1
      header_action {
        request_headers_to_remove = ["RemoveMe2"]
        request_headers_to_add {
          header_name = "AddSomethingElse"
          header_value = "MyOtherValue"
          replace = true
        }
        response_headers_to_remove = ["RemoveMe3"]
        response_headers_to_add {
          header_name = "AddMe"
          header_value = "MyValue"
          replace = false
        }
      }
      match_rules {
        full_path_match = "a full path"
        header_matches {
          header_name = "someheader"
          exact_match = "match this exactly"
          invert_match = true
        }
        ignore_case = true
        metadata_filters {
          filter_match_criteria = "MATCH_ANY"
          filter_labels {
            name = "PLANET"
            value = "MARS"
          }
        }
        query_parameter_matches {
          name = "a query parameter"
          present_match = true
        }
      }
      url_redirect {
        host_redirect = "A host"
        https_redirect = false
        path_redirect = "some/path"
        redirect_response_code = "TEMPORARY_REDIRECT"
        strip_query = true
      }
    }
  }

  test {
    service = google_compute_region_backend_service.home.self_link
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_region_backend_service" "home" {
  provider    = "google-beta"
  name        = "home%<random_suffix>s"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.self_link]
  load_balancing_scheme = "INTERNAL_MANAGED"
}

resource "google_compute_region_health_check" "default" {
  provider = "google-beta"
  name     = "health-check%<random_suffix>s"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeRegionUrlMap_regionUrlMapL7IlbRoutePartialExample(t *testing.T) {
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
				Config: testAccComputeRegionUrlMap_regionUrlMapL7IlbRoutePartialExample(context),
			},
		},
	})
}

func testAccComputeRegionUrlMap_regionUrlMapL7IlbRoutePartialExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_url_map" "regionurlmap" {
  provider = "google-beta"
  name        = "regionurlmap%<random_suffix>s"
  description = "a description"
  default_service = google_compute_region_backend_service.home.self_link

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = google_compute_region_backend_service.home.self_link

    route_rules {
      priority = 1
      service = google_compute_region_backend_service.home.self_link
      header_action {
        request_headers_to_remove = ["RemoveMe2"]
      }
      match_rules {
        full_path_match = "a full path"
        header_matches {
          header_name = "someheader"
          exact_match = "match this exactly"
          invert_match = true
        }
        query_parameter_matches {
          name = "a query parameter"
          present_match = true
        }
      }
    }
  }

  test {
    service = google_compute_region_backend_service.home.self_link
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_region_backend_service" "home" {
  provider = "google-beta"
  name        = "home%<random_suffix>s"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.self_link]
  load_balancing_scheme = "INTERNAL_MANAGED"
}

resource "google_compute_region_health_check" "default" {
  provider = "google-beta"
  name               = "health-check%<random_suffix>s"
  http_health_check {
    port = 80
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
