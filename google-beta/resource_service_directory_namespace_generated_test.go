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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccServiceDirectoryNamespace_serviceDirectoryNamespaceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckServiceDirectoryNamespaceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceDirectoryNamespace_serviceDirectoryNamespaceBasicExample(context),
			},
			{
				ResourceName:            "google_service_directory_namespace.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "namespace_id"},
			},
		},
	})
}

func testAccServiceDirectoryNamespace_serviceDirectoryNamespaceBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_service_directory_namespace" "example" {
  provider     = google-beta
  namespace_id = "tf-test-example-namespace%{random_suffix}"
  location     = "us-central1"

  labels = {
    key = "value"
    foo = "bar"
  }
}
`, context)
}

func testAccCheckServiceDirectoryNamespaceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_service_directory_namespace" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ServiceDirectoryBasePath}}{{name}}")
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
				return fmt.Errorf("ServiceDirectoryNamespace still exists at %s", url)
			}
		}

		return nil
	}
}
