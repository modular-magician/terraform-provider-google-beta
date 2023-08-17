// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package vertexai_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccVertexAIEndpointIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIEndpointIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_endpoint_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/endpoints/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv()),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccVertexAIEndpointIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_endpoint_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/endpoints/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv()),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVertexAIEndpointIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccVertexAIEndpointIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_vertex_ai_endpoint_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/endpoints/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv()),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVertexAIEndpointIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIEndpointIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_vertex_ai_endpoint_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_vertex_ai_endpoint_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv()),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccVertexAIEndpointIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_vertex_ai_endpoint_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv()),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccVertexAIEndpointIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_endpoint" "endpoint" {
  name         = "tf-test-endpoint-name%{random_suffix}"
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  location     = "us-central1"
  region       = "us-central1"
  labels       = {
    label-one = "value-one"
  }
  network      = "projects/${data.google_project.project.number}/global/networks/${data.google_compute_network.vertex_network.name}"
  encryption_spec {
    kms_key_name = "tf-test-kms-name%{random_suffix}"
  }
  depends_on   = [
    google_service_networking_connection.vertex_vpc_connection
  ]
}

resource "google_service_networking_connection" "vertex_vpc_connection" {
  network                 = data.google_compute_network.vertex_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.vertex_range.name]
}

resource "google_compute_global_address" "vertex_range" {
  name          = "tf-test-address-name%{random_suffix}"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 24
  network       = data.google_compute_network.vertex_network.id
}

data "google_compute_network" "vertex_network" {
  name       = "tf-test-network-name%{random_suffix}"
}

resource "google_kms_crypto_key_iam_member" "crypto_key" {
  crypto_key_id = "tf-test-kms-name%{random_suffix}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-aiplatform.iam.gserviceaccount.com"
}

data "google_project" "project" {}

resource "google_vertex_ai_endpoint_iam_member" "foo" {
  project = google_vertex_ai_endpoint.endpoint.project
  region = google_vertex_ai_endpoint.endpoint.region
  endpoint = google_vertex_ai_endpoint.endpoint.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccVertexAIEndpointIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_endpoint" "endpoint" {
  name         = "tf-test-endpoint-name%{random_suffix}"
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  location     = "us-central1"
  region       = "us-central1"
  labels       = {
    label-one = "value-one"
  }
  network      = "projects/${data.google_project.project.number}/global/networks/${data.google_compute_network.vertex_network.name}"
  encryption_spec {
    kms_key_name = "tf-test-kms-name%{random_suffix}"
  }
  depends_on   = [
    google_service_networking_connection.vertex_vpc_connection
  ]
}

resource "google_service_networking_connection" "vertex_vpc_connection" {
  network                 = data.google_compute_network.vertex_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.vertex_range.name]
}

resource "google_compute_global_address" "vertex_range" {
  name          = "tf-test-address-name%{random_suffix}"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 24
  network       = data.google_compute_network.vertex_network.id
}

data "google_compute_network" "vertex_network" {
  name       = "tf-test-network-name%{random_suffix}"
}

resource "google_kms_crypto_key_iam_member" "crypto_key" {
  crypto_key_id = "tf-test-kms-name%{random_suffix}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-aiplatform.iam.gserviceaccount.com"
}

data "google_project" "project" {}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_vertex_ai_endpoint_iam_policy" "foo" {
  project = google_vertex_ai_endpoint.endpoint.project
  region = google_vertex_ai_endpoint.endpoint.region
  endpoint = google_vertex_ai_endpoint.endpoint.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_vertex_ai_endpoint_iam_policy" "foo" {
  project = google_vertex_ai_endpoint.endpoint.project
  region = google_vertex_ai_endpoint.endpoint.region
  endpoint = google_vertex_ai_endpoint.endpoint.name
  depends_on = [
    google_vertex_ai_endpoint_iam_policy.foo
  ]
}
`, context)
}

func testAccVertexAIEndpointIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_endpoint" "endpoint" {
  name         = "tf-test-endpoint-name%{random_suffix}"
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  location     = "us-central1"
  region       = "us-central1"
  labels       = {
    label-one = "value-one"
  }
  network      = "projects/${data.google_project.project.number}/global/networks/${data.google_compute_network.vertex_network.name}"
  encryption_spec {
    kms_key_name = "tf-test-kms-name%{random_suffix}"
  }
  depends_on   = [
    google_service_networking_connection.vertex_vpc_connection
  ]
}

resource "google_service_networking_connection" "vertex_vpc_connection" {
  network                 = data.google_compute_network.vertex_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.vertex_range.name]
}

resource "google_compute_global_address" "vertex_range" {
  name          = "tf-test-address-name%{random_suffix}"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 24
  network       = data.google_compute_network.vertex_network.id
}

data "google_compute_network" "vertex_network" {
  name       = "tf-test-network-name%{random_suffix}"
}

resource "google_kms_crypto_key_iam_member" "crypto_key" {
  crypto_key_id = "tf-test-kms-name%{random_suffix}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-aiplatform.iam.gserviceaccount.com"
}

data "google_project" "project" {}

data "google_iam_policy" "foo" {
}

resource "google_vertex_ai_endpoint_iam_policy" "foo" {
  project = google_vertex_ai_endpoint.endpoint.project
  region = google_vertex_ai_endpoint.endpoint.region
  endpoint = google_vertex_ai_endpoint.endpoint.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccVertexAIEndpointIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_endpoint" "endpoint" {
  name         = "tf-test-endpoint-name%{random_suffix}"
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  location     = "us-central1"
  region       = "us-central1"
  labels       = {
    label-one = "value-one"
  }
  network      = "projects/${data.google_project.project.number}/global/networks/${data.google_compute_network.vertex_network.name}"
  encryption_spec {
    kms_key_name = "tf-test-kms-name%{random_suffix}"
  }
  depends_on   = [
    google_service_networking_connection.vertex_vpc_connection
  ]
}

resource "google_service_networking_connection" "vertex_vpc_connection" {
  network                 = data.google_compute_network.vertex_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.vertex_range.name]
}

resource "google_compute_global_address" "vertex_range" {
  name          = "tf-test-address-name%{random_suffix}"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 24
  network       = data.google_compute_network.vertex_network.id
}

data "google_compute_network" "vertex_network" {
  name       = "tf-test-network-name%{random_suffix}"
}

resource "google_kms_crypto_key_iam_member" "crypto_key" {
  crypto_key_id = "tf-test-kms-name%{random_suffix}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-aiplatform.iam.gserviceaccount.com"
}

data "google_project" "project" {}

resource "google_vertex_ai_endpoint_iam_binding" "foo" {
  project = google_vertex_ai_endpoint.endpoint.project
  region = google_vertex_ai_endpoint.endpoint.region
  endpoint = google_vertex_ai_endpoint.endpoint.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccVertexAIEndpointIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_endpoint" "endpoint" {
  name         = "tf-test-endpoint-name%{random_suffix}"
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  location     = "us-central1"
  region       = "us-central1"
  labels       = {
    label-one = "value-one"
  }
  network      = "projects/${data.google_project.project.number}/global/networks/${data.google_compute_network.vertex_network.name}"
  encryption_spec {
    kms_key_name = "tf-test-kms-name%{random_suffix}"
  }
  depends_on   = [
    google_service_networking_connection.vertex_vpc_connection
  ]
}

resource "google_service_networking_connection" "vertex_vpc_connection" {
  network                 = data.google_compute_network.vertex_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.vertex_range.name]
}

resource "google_compute_global_address" "vertex_range" {
  name          = "tf-test-address-name%{random_suffix}"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 24
  network       = data.google_compute_network.vertex_network.id
}

data "google_compute_network" "vertex_network" {
  name       = "tf-test-network-name%{random_suffix}"
}

resource "google_kms_crypto_key_iam_member" "crypto_key" {
  crypto_key_id = "tf-test-kms-name%{random_suffix}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-aiplatform.iam.gserviceaccount.com"
}

data "google_project" "project" {}

resource "google_vertex_ai_endpoint_iam_binding" "foo" {
  project = google_vertex_ai_endpoint.endpoint.project
  region = google_vertex_ai_endpoint.endpoint.region
  endpoint = google_vertex_ai_endpoint.endpoint.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
