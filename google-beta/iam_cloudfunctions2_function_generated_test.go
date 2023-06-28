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
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccCloudfunctions2functionIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
		"project":       envvar.GetTestProjectFromEnv(),

		"zip_path": "./test-fixtures/cloudfunctions2/function-source.zip",
		"location": "us-central1",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudfunctions2functionIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_cloudfunctions2_function_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/functions/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-function-v2%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccCloudfunctions2functionIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_cloudfunctions2_function_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/functions/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-function-v2%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudfunctions2functionIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
		"project":       envvar.GetTestProjectFromEnv(),

		"zip_path": "./test-fixtures/cloudfunctions2/function-source.zip",
		"location": "us-central1",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccCloudfunctions2functionIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_cloudfunctions2_function_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/functions/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-function-v2%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudfunctions2functionIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
		"project":       envvar.GetTestProjectFromEnv(),

		"zip_path": "./test-fixtures/cloudfunctions2/function-source.zip",
		"location": "us-central1",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudfunctions2functionIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_cloudfunctions2_function_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_cloudfunctions2_function_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/functions/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-function-v2%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCloudfunctions2functionIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_cloudfunctions2_function_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/functions/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-function-v2%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudfunctions2functionIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
# [START functions_v2_basic]
locals {
  project = "%{project}" # Google Cloud Platform Project ID
}

resource "google_storage_bucket" "bucket" {
  name     = "${local.project}-tf-test-gcf-source%{random_suffix}"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}
 
resource "google_cloudfunctions2_function" "function" {
  name = "tf-test-function-v2%{random_suffix}"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime = "nodejs16"
    entry_point = "helloHttp"  # Set the entry point 
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

output "function_uri" { 
  value = google_cloudfunctions2_function.function.service_config[0].uri
}
# [END functions_v2_basic]

resource "google_cloudfunctions2_function_iam_member" "foo" {
  project = google_cloudfunctions2_function.function.project
  location = google_cloudfunctions2_function.function.location
  cloud_function = google_cloudfunctions2_function.function.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccCloudfunctions2functionIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
# [START functions_v2_basic]
locals {
  project = "%{project}" # Google Cloud Platform Project ID
}

resource "google_storage_bucket" "bucket" {
  name     = "${local.project}-tf-test-gcf-source%{random_suffix}"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}
 
resource "google_cloudfunctions2_function" "function" {
  name = "tf-test-function-v2%{random_suffix}"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime = "nodejs16"
    entry_point = "helloHttp"  # Set the entry point 
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

output "function_uri" { 
  value = google_cloudfunctions2_function.function.service_config[0].uri
}
# [END functions_v2_basic]

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_cloudfunctions2_function_iam_policy" "foo" {
  project = google_cloudfunctions2_function.function.project
  location = google_cloudfunctions2_function.function.location
  cloud_function = google_cloudfunctions2_function.function.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_cloudfunctions2_function_iam_policy" "foo" {
  project = google_cloudfunctions2_function.function.project
  location = google_cloudfunctions2_function.function.location
  cloud_function = google_cloudfunctions2_function.function.name
  depends_on = [
    google_cloudfunctions2_function_iam_policy.foo
  ]
}
`, context)
}

func testAccCloudfunctions2functionIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
# [START functions_v2_basic]
locals {
  project = "%{project}" # Google Cloud Platform Project ID
}

resource "google_storage_bucket" "bucket" {
  name     = "${local.project}-tf-test-gcf-source%{random_suffix}"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}
 
resource "google_cloudfunctions2_function" "function" {
  name = "tf-test-function-v2%{random_suffix}"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime = "nodejs16"
    entry_point = "helloHttp"  # Set the entry point 
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

output "function_uri" { 
  value = google_cloudfunctions2_function.function.service_config[0].uri
}
# [END functions_v2_basic]

data "google_iam_policy" "foo" {
}

resource "google_cloudfunctions2_function_iam_policy" "foo" {
  project = google_cloudfunctions2_function.function.project
  location = google_cloudfunctions2_function.function.location
  cloud_function = google_cloudfunctions2_function.function.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccCloudfunctions2functionIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
# [START functions_v2_basic]
locals {
  project = "%{project}" # Google Cloud Platform Project ID
}

resource "google_storage_bucket" "bucket" {
  name     = "${local.project}-tf-test-gcf-source%{random_suffix}"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}
 
resource "google_cloudfunctions2_function" "function" {
  name = "tf-test-function-v2%{random_suffix}"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime = "nodejs16"
    entry_point = "helloHttp"  # Set the entry point 
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

output "function_uri" { 
  value = google_cloudfunctions2_function.function.service_config[0].uri
}
# [END functions_v2_basic]

resource "google_cloudfunctions2_function_iam_binding" "foo" {
  project = google_cloudfunctions2_function.function.project
  location = google_cloudfunctions2_function.function.location
  cloud_function = google_cloudfunctions2_function.function.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccCloudfunctions2functionIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
# [START functions_v2_basic]
locals {
  project = "%{project}" # Google Cloud Platform Project ID
}

resource "google_storage_bucket" "bucket" {
  name     = "${local.project}-tf-test-gcf-source%{random_suffix}"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}
 
resource "google_cloudfunctions2_function" "function" {
  name = "tf-test-function-v2%{random_suffix}"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime = "nodejs16"
    entry_point = "helloHttp"  # Set the entry point 
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

output "function_uri" { 
  value = google_cloudfunctions2_function.function.service_config[0].uri
}
# [END functions_v2_basic]

resource "google_cloudfunctions2_function_iam_binding" "foo" {
  project = google_cloudfunctions2_function.function.project
  location = google_cloudfunctions2_function.function.location
  cloud_function = google_cloudfunctions2_function.function.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
