// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package functions_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccProviderFunction_location_from_id(t *testing.T) {
	t.Parallel()

	location := envvar.GetTestRegionFromEnv()
	locationRegex := regexp.MustCompile(fmt.Sprintf("^%s$", location))

	validId := fmt.Sprintf("projects/my-project/locations/%s/services/my-cloudrun-service", location)
	repetitiveInput := fmt.Sprintf("projects/my-project/locations/%s/locations/not-this-one/services/my-cloudrun-service", location)
	invalidId := "projects/my-project/zones/us-central1-c/instances/my-instance"

	context := map[string]interface{}{
		"function_name": "location_from_id",
		"output_name":   "location",
		"location":      location,
		"resource_name": fmt.Sprintf("tf-test-location-func-%s", acctest.RandString(t, 10)),
		"input":         "", // overridden in test cases
	}

	acctest.VcrTest(t, resource.TestCase{
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Given valid resource id input, the output value matches the expected value
				Config: testProviderFunction_generic_element_from_id(context, validId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchOutput(context["output_name"].(string), locationRegex),
				),
			},
			{
				// Given repetitive input, the output value is the left-most match in the input
				Config: testProviderFunction_generic_element_from_id(context, repetitiveInput),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchOutput(context["output_name"].(string), locationRegex),
				),
			},
			{
				// Given invalid input, an error occurs
				Config:      testProviderFunction_generic_element_from_id(context, invalidId),
				ExpectError: regexp.MustCompile("Error in function call"), // ExpectError doesn't inspect the specific error messages
			},
			{
				// Can get the location from a resource's id in one step
				// Uses google_compute_subnetwork resource's id attribute with format projects/{{project}}/locations/{{location}}/subnetworks/{{name}}
				Config: testProviderFunction_get_location_from_resource_id(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchOutput(context["output_name"].(string), locationRegex),
				),
			},
			// No test case for self link - are self links solely zonal/regional?
		},
	})
}

func testProviderFunction_get_location_from_resource_id(context map[string]interface{}) string {
	return acctest.Nprintf(`
# terraform block required for provider function to be found
terraform {
	required_providers {
		google = {
			source = "hashicorp/google"
		}
	}
}

resource "google_cloud_run_v2_service" "default" {
  name     = "%{resource_name}"
  location = "%{location}"
  ingress = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
  }
}

output "%{output_name}" {
	value = provider::google::%{function_name}(google_cloud_run_v2_service.default.id)
}
`, context)
}
