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

package networkservices_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccNetworkServicesServiceBinding_networkServicesServiceBindingBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesServiceBindingDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesServiceBinding_networkServicesServiceBindingBasicExample(context),
			},
			{
				ResourceName:            "google_network_services_service_binding.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "name", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkServicesServiceBinding_networkServicesServiceBindingBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_service_directory_namespace" "default" {
  provider     = google-beta
  namespace_id = "tf-test-my-namespace%{random_suffix}"
  location     = "us-central1"
}

resource "google_service_directory_service" "default" {
  provider   = google-beta
  service_id = "tf-test-my-service%{random_suffix}"
  namespace  = google_service_directory_namespace.default.id

  metadata = {
    stage  = "prod"
    region = "us-central1"
  }
}

resource "google_network_services_service_binding" "default" {
  provider    = google-beta
  name        = "tf-test-my-service-binding%{random_suffix}"
  labels      = {
    foo = "bar"
  }
  description = "my description"
  service = google_service_directory_service.default.id
}
`, context)
}

func testAccCheckNetworkServicesServiceBindingDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_services_service_binding" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/serviceBindings/{{name}}")
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
				return fmt.Errorf("NetworkServicesServiceBinding still exists at %s", url)
			}
		}

		return nil
	}
}
