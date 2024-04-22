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

package monitoring_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccMonitoringUptimeCheckConfig_uptimeCheckConfigHttpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckMonitoringUptimeCheckConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringUptimeCheckConfig_uptimeCheckConfigHttpExample(context),
			},
			{
				ResourceName:      "google_monitoring_uptime_check_config.http",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMonitoringUptimeCheckConfig_uptimeCheckConfigHttpExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_monitoring_uptime_check_config" "http" {
  display_name = "tf-test-http-uptime-check%{random_suffix}"
  timeout      = "60s"
  user_labels  = {
    example-key = "example-value"
  }

  http_check {
    path = "some-path"
    port = "8010"
    request_method = "POST"
    content_type = "USER_PROVIDED"
    custom_content_type = "application/json"
    body = "Zm9vJTI1M0RiYXI="
    ping_config {
      pings_count = 1
    }
  }

  monitored_resource {
    type = "uptime_url"
    labels = {
      project_id = "%{project_id}"
      host       = "192.168.1.1"
    }
  }

  content_matchers {
    content = "\"example\""
    matcher = "MATCHES_JSON_PATH"
    json_path_matcher {
      json_path = "$.path"
      json_matcher = "EXACT_MATCH"
    }
  }

  checker_type = "STATIC_IP_CHECKERS"
}
`, context)
}

func TestAccMonitoringUptimeCheckConfig_uptimeCheckConfigStatusCodeExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckMonitoringUptimeCheckConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringUptimeCheckConfig_uptimeCheckConfigStatusCodeExample(context),
			},
			{
				ResourceName:      "google_monitoring_uptime_check_config.status_code",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMonitoringUptimeCheckConfig_uptimeCheckConfigStatusCodeExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_monitoring_uptime_check_config" "status_code" {
  display_name = "tf-test-http-uptime-check%{random_suffix}"
  timeout      = "60s"

  http_check {
    path = "some-path"
    port = "8010"
    request_method = "POST"
    content_type = "URL_ENCODED"
    body = "Zm9vJTI1M0RiYXI="

    accepted_response_status_codes {
      status_class = "STATUS_CLASS_2XX"
    }
    accepted_response_status_codes {
            status_value = 301
    }
    accepted_response_status_codes {
            status_value = 302
    }
  }

  monitored_resource {
    type = "uptime_url"
    labels = {
      project_id = "%{project_id}"
      host       = "192.168.1.1"
    }
  }

  content_matchers {
    content = "\"example\""
    matcher = "MATCHES_JSON_PATH"
    json_path_matcher {
      json_path = "$.path"
      json_matcher = "EXACT_MATCH"
    }
  }

  checker_type = "STATIC_IP_CHECKERS"
}
`, context)
}

func TestAccMonitoringUptimeCheckConfig_uptimeCheckConfigHttpsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckMonitoringUptimeCheckConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringUptimeCheckConfig_uptimeCheckConfigHttpsExample(context),
			},
			{
				ResourceName:      "google_monitoring_uptime_check_config.https",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMonitoringUptimeCheckConfig_uptimeCheckConfigHttpsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_monitoring_uptime_check_config" "https" {
  display_name = "tf-test-https-uptime-check%{random_suffix}"
  timeout = "60s"

  http_check {
    path = "/some-path"
    port = "443"
    use_ssl = true
    validate_ssl = true
    service_agent_authentication {
      type = "OIDC_TOKEN"
    }
  }

  monitored_resource {
    type = "uptime_url"
    labels = {
      project_id = "%{project_id}"
      host = "192.168.1.1"
    }
  }

  content_matchers {
    content = "example"
    matcher = "MATCHES_JSON_PATH"
    json_path_matcher {
      json_path = "$.path"
      json_matcher = "REGEX_MATCH"
    }
  }
}
`, context)
}

func TestAccMonitoringUptimeCheckConfig_uptimeCheckTcpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckMonitoringUptimeCheckConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringUptimeCheckConfig_uptimeCheckTcpExample(context),
			},
			{
				ResourceName:      "google_monitoring_uptime_check_config.tcp_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMonitoringUptimeCheckConfig_uptimeCheckTcpExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_monitoring_uptime_check_config" "tcp_group" {
  display_name = "tf-test-tcp-uptime-check%{random_suffix}"
  timeout      = "60s"

  tcp_check {
    port = 888
    ping_config {
      pings_count = 2
    }
  }

  resource_group {
    resource_type = "INSTANCE"
    group_id      = google_monitoring_group.check.name
  }
}

resource "google_monitoring_group" "check" {
  display_name = "tf-test-uptime-check-group%{random_suffix}"
  filter       = "resource.metadata.name=has_substring(\"foo\")"
}
`, context)
}

func TestAccMonitoringUptimeCheckConfig_uptimeCheckConfigSyntheticMonitorExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"zip_path":      "./test-fixtures/synthetic-fn-source.zip",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckMonitoringUptimeCheckConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringUptimeCheckConfig_uptimeCheckConfigSyntheticMonitorExample(context),
			},
			{
				ResourceName:      "google_monitoring_uptime_check_config.synthetic_monitor",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMonitoringUptimeCheckConfig_uptimeCheckConfigSyntheticMonitorExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_storage_bucket" "bucket" {
  name     = "%{project_id}-tf-test-gcf-source%{random_suffix}"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}
 
resource "google_cloudfunctions2_function" "function" {
  name = "tf_test_synthetic_function%{random_suffix}"
  location = "us-central1"
 
  build_config {
    runtime = "nodejs16"
    entry_point = "SyntheticFunction"  # Set the entry point 
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }
 
  service_config {
    max_instance_count  = 1
    available_memory    = "256M"
    timeout_seconds     = 60
  }
}

resource "google_monitoring_uptime_check_config" "synthetic_monitor" {
  display_name = "tf_test_synthetic_monitor%{random_suffix}"
  timeout = "60s"

  synthetic_monitor {
    cloud_function_v2 {
      name = google_cloudfunctions2_function.function.id
    }
  }
}
`, context)
}

func testAccCheckMonitoringUptimeCheckConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_monitoring_uptime_check_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{MonitoringBasePath}}v3/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:               config,
				Method:               "GET",
				Project:              billingProject,
				RawURL:               url,
				UserAgent:            config.UserAgent,
				ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsMonitoringConcurrentEditError},
			})
			if err == nil {
				return fmt.Errorf("MonitoringUptimeCheckConfig still exists at %s", url)
			}
		}

		return nil
	}
}
