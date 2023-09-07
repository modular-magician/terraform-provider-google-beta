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

package cloudscheduler_test

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

func TestAccCloudSchedulerJob_schedulerJobPubsubExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudSchedulerJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudSchedulerJob_schedulerJobPubsubExample(context),
			},
			{
				ResourceName:            "google_cloud_scheduler_job.job",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccCloudSchedulerJob_schedulerJobPubsubExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_topic" "topic" {
  name = "tf-test-job-topic%{random_suffix}"
}

resource "google_cloud_scheduler_job" "job" {
  name        = "tf-test-test-job%{random_suffix}"
  description = "test job"
  schedule    = "*/2 * * * *"

  pubsub_target {
    # topic.id is the topic's full resource name.
    topic_name = google_pubsub_topic.topic.id
    data       = base64encode("test")
  }
}
`, context)
}

func TestAccCloudSchedulerJob_schedulerJobHttpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudSchedulerJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudSchedulerJob_schedulerJobHttpExample(context),
			},
			{
				ResourceName:            "google_cloud_scheduler_job.job",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccCloudSchedulerJob_schedulerJobHttpExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_scheduler_job" "job" {
  name             = "tf-test-test-job%{random_suffix}"
  description      = "test http job"
  schedule         = "*/8 * * * *"
  time_zone        = "America/New_York"
  attempt_deadline = "320s"

  retry_config {
    retry_count = 1
  }

  http_target {
    http_method = "POST"
    uri         = "https://example.com/"
    body        = base64encode("{\"foo\":\"bar\"}")
  }
}
`, context)
}

func TestAccCloudSchedulerJob_schedulerJobPausedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudSchedulerJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudSchedulerJob_schedulerJobPausedExample(context),
			},
			{
				ResourceName:            "google_cloud_scheduler_job.job",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccCloudSchedulerJob_schedulerJobPausedExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_scheduler_job" "job" {
  paused           = true
  name             = "tf-test-test-job%{random_suffix}"
  description      = "test http job with updated fields"
  schedule         = "*/8 * * * *"
  time_zone        = "America/New_York"
  attempt_deadline = "320s"

  retry_config {
    retry_count = 1
  }

  http_target {
    http_method = "POST"
    uri         = "https://example.com/ping"
    body        = base64encode("{\"foo\":\"bar\"}")
  }
  headers = {
    "Content-Type" = "application/json"
  }
}
`, context)
}

func TestAccCloudSchedulerJob_schedulerJobAppEngineExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudSchedulerJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudSchedulerJob_schedulerJobAppEngineExample(context),
			},
			{
				ResourceName:            "google_cloud_scheduler_job.job",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccCloudSchedulerJob_schedulerJobAppEngineExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_scheduler_job" "job" {
  name             = "tf-test-test-job%{random_suffix}"
  schedule         = "*/4 * * * *"
  description      = "test app engine job"
  time_zone        = "Europe/London"
  attempt_deadline = "320s"

  retry_config {
    min_backoff_duration = "1s"
    max_retry_duration = "10s"
    max_doublings = 2
    retry_count = 3
  }

  app_engine_http_target {
    http_method = "POST"

    app_engine_routing {
      service  = "web"
      version  = "prod"
      instance = "my-instance-001"
    }

    relative_uri = "/ping"
  }
}
`, context)
}

func TestAccCloudSchedulerJob_schedulerJobOauthExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"region":        envvar.GetTestRegionFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudSchedulerJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudSchedulerJob_schedulerJobOauthExample(context),
			},
			{
				ResourceName:            "google_cloud_scheduler_job.job",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccCloudSchedulerJob_schedulerJobOauthExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_compute_default_service_account" "default" {
}

resource "google_cloud_scheduler_job" "job" {
  name             = "tf-test-test-job%{random_suffix}"
  description      = "test http job"
  schedule         = "*/8 * * * *"
  time_zone        = "America/New_York"
  attempt_deadline = "320s"

  http_target {
    http_method = "GET"
    uri         = "https://cloudscheduler.googleapis.com/v1/projects/%{project_name}/locations/%{region}/jobs"

    oauth_token {
      service_account_email = data.google_compute_default_service_account.default.email
    }
  }
}
`, context)
}

func TestAccCloudSchedulerJob_schedulerJobOidcExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudSchedulerJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudSchedulerJob_schedulerJobOidcExample(context),
			},
			{
				ResourceName:            "google_cloud_scheduler_job.job",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccCloudSchedulerJob_schedulerJobOidcExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_compute_default_service_account" "default" {
}

resource "google_cloud_scheduler_job" "job" {
  name             = "tf-test-test-job%{random_suffix}"
  description      = "test http job"
  schedule         = "*/8 * * * *"
  time_zone        = "America/New_York"
  attempt_deadline = "320s"

  http_target {
    http_method = "GET"
    uri         = "https://example.com/ping"

    oidc_token {
      service_account_email = data.google_compute_default_service_account.default.email
    }
  }
}
`, context)
}

func testAccCheckCloudSchedulerJobDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloud_scheduler_job" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{CloudSchedulerBasePath}}projects/{{project}}/locations/{{region}}/jobs/{{name}}")
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
				return fmt.Errorf("CloudSchedulerJob still exists at %s", url)
			}
		}

		return nil
	}
}
