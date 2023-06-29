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

func TestAccNetworkSecurityTlsInspectionPolicy_networkSecurityTlsInspectionPolicyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityTlsInspectionPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityTlsInspectionPolicy_networkSecurityTlsInspectionPolicyBasicExample(context),
			},
			{
				ResourceName:            "google_network_security_tls_inspection_policy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccNetworkSecurityTlsInspectionPolicy_networkSecurityTlsInspectionPolicyBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  provider = google-beta
  name      = "tf-test-my-basic-ca-pool%{random_suffix}"
  location  = "us-central1"
  tier     = "DEVOPS"
  publishing_options {
    publish_ca_cert = false
    publish_crl = false
  }
  issuance_policy {
    maximum_lifetime = "1209600s"
    baseline_values {
      ca_options {
        is_ca = false
      }
      key_usage {
        base_key_usage {}
        extended_key_usage {
          server_auth = true
        }
      }
    }
  }
}


resource "google_privateca_certificate_authority" "default" {
  provider = google-beta
  pool = google_privateca_ca_pool.default.name
  certificate_authority_id = "tf-test-my-basic-certificate-authority%{random_suffix}"
  location = "us-central1"
  lifetime = "86400s"
  type = "SELF_SIGNED"
  deletion_protection = false
  skip_grace_period = true
  ignore_active_certificates_on_deletion = true
  config {
    subject_config {
      subject {
        organization = "Test LLC"
        common_name = "my-ca"
      }
    }
    x509_config {
      ca_options {
        is_ca = true
      }
      key_usage {
        base_key_usage {
          cert_sign = true
          crl_sign = true
        }
        extended_key_usage {
          server_auth = false
        }
      }
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }
}

resource "google_project_service_identity" "ns_sa" {
  provider = google-beta

  service = "networksecurity.googleapis.com"
}

resource "google_privateca_ca_pool_iam_member" "tls_inspection_permission" {
  provider = google-beta

  ca_pool = google_privateca_ca_pool.default.id
  role = "roles/privateca.certificateManager"
  member = "serviceAccount:${google_project_service_identity.ns_sa.email}"
}

resource "google_network_security_tls_inspection_policy" "default" {
  provider = google-beta
  name     = "tf-test-my-tls-inspection-policy%{random_suffix}"
  location = "us-central1"
  ca_pool  = google_privateca_ca_pool.default.id
  exclude_public_ca_set = false
  depends_on = [google_privateca_ca_pool.default, google_privateca_certificate_authority.default, google_privateca_ca_pool_iam_member.tls_inspection_permission]
}
`, context)
}

func testAccCheckNetworkSecurityTlsInspectionPolicyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_security_tls_inspection_policy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/tlsInspectionPolicies/{{name}}")
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
				return fmt.Errorf("NetworkSecurityTlsInspectionPolicy still exists at %s", url)
			}
		}

		return nil
	}
}
