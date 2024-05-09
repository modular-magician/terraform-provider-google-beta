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

package netapp_test

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

func TestAccNetappkmsconfig_kmsConfigCreateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetappkmsconfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetappkmsconfig_kmsConfigCreateExample(context),
			},
			{
				ResourceName:            "google_netapp_kmsconfig.kmsConfig",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "name", "terraform_labels"},
			},
		},
	})
}

func testAccNetappkmsconfig_kmsConfigCreateExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_kms_key_ring" "keyring" {
  name     = "tf-test-key-ring%{random_suffix}"
  location = "us-central1"
}

resource "google_kms_crypto_key" "crypto_key" {
  name            = "tf-test-crypto-name%{random_suffix}"
  key_ring        = google_kms_key_ring.keyring.id
  # rotation_period = "7776000s"
}

resource "google_netapp_kmsconfig" "kmsConfig" {
  name = "tf-test-kms-test%{random_suffix}"
  description="this is a test description"
  crypto_key_name=google_kms_crypto_key.crypto_key.id
  location="us-central1"
}
`, context)
}

func testAccCheckNetappkmsconfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_netapp_kmsconfig" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetappBasePath}}projects/{{project}}/locations/{{location}}/kmsConfigs/{{name}}")
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
				return fmt.Errorf("Netappkmsconfig still exists at %s", url)
			}
		}

		return nil
	}
}
