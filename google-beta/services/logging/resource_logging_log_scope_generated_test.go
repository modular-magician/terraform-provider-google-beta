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

package logging_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccLoggingLogScope_loggingLogScopeBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckLoggingLogScopeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccLoggingLogScope_loggingLogScopeBasicExample(context),
			},
			{
				ResourceName:            "google_logging_log_scope.logging_log_scope",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name", "parent"},
			},
		},
	})
}

func testAccLoggingLogScope_loggingLogScopeBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_logging_log_scope" "logging_log_scope" {
    parent         = "projects/%{project}"
    location       = "global"
    name           = "projects/%{project}/locations/global/logScopes/tf-test-my-log-scope%{random_suffix}"
    resource_names = [
        "projects/%{project}",
        "projects/%{project}/locations/global/buckets/_Default/views/view1%{random_suffix}",
        "projects/%{project}/locations/global/buckets/_Default/views/view2%{random_suffix}"
    ]
    description    = "A log scope configured with Terraform"
}
`, context)
}

func testAccCheckLoggingLogScopeDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_logging_log_scope" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{LoggingBasePath}}{{parent}}/locations/{{location}}/logScopes/{{name}}")
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
				return fmt.Errorf("LoggingLogScope still exists at %s", url)
			}
		}

		return nil
	}
}
