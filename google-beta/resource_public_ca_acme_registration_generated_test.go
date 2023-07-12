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
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccPublicCAAcmeRegistration_publicCaAcmeRegistrationExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
			"tls":    {},
		},
		CheckDestroy: testAccCheckPublicCAAcmeRegistrationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPublicCAAcmeRegistration_publicCaAcmeRegistrationExample(context),
			},
		},
	})
}

func testAccPublicCAAcmeRegistration_publicCaAcmeRegistrationExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "tls_private_key" "default" {
  algorithm = "RSA"
  rsa_bits  = 2048
}

resource "google_public_ca_external_account_key" "account_key" {
  project = "%{project}"
}

data "google_client_openid_userinfo" "user" {}

resource "google_public_ca_acme_registration" "test" {
  depends_on = [
    tls_private_key.default,
    data.google_client_openid_userinfo.user,
    google_public_ca_external_account_key.account_key
  ]
  name            = "test"
  email           = data.google_client_openid_userinfo.user.email
  private_key_pem = tls_private_key.default.private_key_pem_pkcs8
  eab_key_id      = google_public_ca_external_account_key.account_key.key_id
  eab_hmac_key    = google_public_ca_external_account_key.account_key.b64_mac_key
  project         = "%{project}"
}
`, context)
}

func testAccCheckPublicCAAcmeRegistrationDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_public_ca_acme_registration" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{PublicCABasePath}}projects/{{project}}/locations/global/acmeRegistrations/{{name}}/{{name}}")
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
				return fmt.Errorf("PublicCAAcmeRegistration still exists at %s", url)
			}
		}

		return nil
	}
}
