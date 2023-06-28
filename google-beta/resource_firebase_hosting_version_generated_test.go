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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccFirebaseHostingVersion_firebasehostingVersionRedirectExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseHostingVersion_firebasehostingVersionRedirectExample(context),
			},
			{
				ResourceName:            "google_firebase_hosting_version.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"version_id", "site_id"},
			},
		},
	})
}

func testAccFirebaseHostingVersion_firebasehostingVersionRedirectExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_hosting_site" "default" {
  provider = google-beta
  project  = "%{project_id}"
  site_id  = "tf-test-site-id%{random_suffix}"
}

resource "google_firebase_hosting_version" "default" {
  provider = google-beta
  site_id  = google_firebase_hosting_site.default.site_id
  config {
    redirects {
      glob = "/google/**"
      status_code = 302
      location = "https://www.google.com"
    }
  }
}

resource "google_firebase_hosting_release" "default" {
  provider     = google-beta
  site_id      = google_firebase_hosting_site.default.site_id
  version_name = google_firebase_hosting_version.default.name
  message      = "Redirect to Google"
}
`, context)
}

func TestAccFirebaseHostingVersion_firebasehostingVersionCloudRunExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseHostingVersion_firebasehostingVersionCloudRunExample(context),
			},
			{
				ResourceName:            "google_firebase_hosting_version.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"version_id", "site_id"},
			},
		},
	})
}

func testAccFirebaseHostingVersion_firebasehostingVersionCloudRunExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_hosting_site" "default" {
  provider = google-beta
  project  = "%{project_id}"
  site_id  = "tf-test-site-id%{random_suffix}"
}

resource "google_cloud_run_v2_service" "default" {
  provider = google-beta
  project  = "%{project_id}"
  name     = "tf-test-cloud-run-service-via-hosting%{random_suffix}"
  location = "us-central1"

  # Warning: allows all public traffic
  ingress = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
  }
}

resource "google_firebase_hosting_version" "default" {
  provider = google-beta
  site_id  = google_firebase_hosting_site.default.site_id
  config {
    rewrites {
      glob = "/hello/**"
      run {
        service_id = google_cloud_run_v2_service.default.name
        region = google_cloud_run_v2_service.default.location
      }
    }
  }
}

resource "google_firebase_hosting_release" "default" {
  provider     = google-beta
  site_id      = google_firebase_hosting_site.default.site_id
  version_name = google_firebase_hosting_version.default.name
  message      = "Cloud Run Integration"
}
`, context)
}

func TestAccFirebaseHostingVersion_firebasehostingVersionCloudFunctionsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"zip_path":      "./test-fixtures/cloudfunctions2/function-source.zip",
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseHostingVersion_firebasehostingVersionCloudFunctionsExample(context),
			},
			{
				ResourceName:            "google_firebase_hosting_version.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"version_id", "site_id"},
			},
		},
	})
}

func testAccFirebaseHostingVersion_firebasehostingVersionCloudFunctionsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_hosting_site" "default" {
  provider = google-beta
  project  = "%{project_id}"
  site_id  = "tf-test-site-id%{random_suffix}"
}

resource "google_storage_bucket" "bucket" {
  provider = google-beta
  project  = "%{project_id}"
  name     = "tf-test-site-id%{random_suffix}-function-source"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "object" {
  provider = google-beta
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}

resource "google_cloudfunctions_function" "function" {
  provider = google-beta
  project  = "%{project_id}"

  name        = "tf-test-cloud-function-via-hosting%{random_suffix}"
  description = "A Cloud Function connected to Firebase Hosing"
  runtime     = "nodejs16"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = google_storage_bucket_object.object.name
  trigger_http          = true
  entry_point           = "helloHttp"
}

resource "google_firebase_hosting_version" "default" {
  provider = google-beta
  site_id  = google_firebase_hosting_site.default.site_id
  config {
    rewrites {
      glob = "/hello/**"
      function = google_cloudfunctions_function.function.name
    }
  }
}

resource "google_firebase_hosting_release" "default" {
  provider     = google-beta
  site_id      = google_firebase_hosting_site.default.site_id
  version_name = google_firebase_hosting_version.default.name
  message      = "Cloud Functions Integration"
}
`, context)
}
