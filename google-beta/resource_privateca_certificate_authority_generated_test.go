// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
)

func TestAccPrivatecaCertificateAuthority_privatecaCertificateAuthorityBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckPrivatecaCertificateAuthorityDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCertificateAuthority_privatecaCertificateAuthorityBasicExample(context),
			},
		},
	})
}

func testAccPrivatecaCertificateAuthority_privatecaCertificateAuthorityBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_authority" "default" {
  provider = google-beta
  certificate_authority_id = "tf-test-my-certificate-authority%{random_suffix}"
  location = "us-central1"
  config {
    subject_config {
      subject {
        organization = "HashiCorp"
      }
      common_name = "my-certificate-authority"
      subject_alt_name {
        dns_names = ["hashicorp.com"]
      }
    }
    reusable_config {
      reusable_config = "projects/568668481468/locations/us-central1/reusableConfigs/root-unconstrained"
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }
  disable_on_delete = true
}
`, context)
}

func TestAccPrivatecaCertificateAuthority_privatecaCertificateAuthorityFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckPrivatecaCertificateAuthorityDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCertificateAuthority_privatecaCertificateAuthorityFullExample(context),
			},
		},
	})
}

func testAccPrivatecaCertificateAuthority_privatecaCertificateAuthorityFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_authority" "default" {
  provider = google-beta
  certificate_authority_id = "tf-test-my-certificate-authority%{random_suffix}"
  location = "us-central1"
  tier = "DEVOPS"
  config {
    subject_config {
      subject {
        country_code = "US"
        organization = "HashiCorp"
        organizational_unit = "Terraform"
        locality = "San Francisco"
        province = "CA"
        street_address = "101 2nd St #700"
        postal_code = "94105"
      }
      common_name = "my-certificate-authority"
      subject_alt_name {
        dns_names = ["hashicorp.com"]
        email_addresses = ["email@example.com"]
        ip_addresses = ["127.0.0.1"]
        uris = ["http://www.ietf.org/rfc/rfc3986.txt"]
      }
    }
    reusable_config {
      reusable_config = "projects/568668481468/locations/us-central1/reusableConfigs/root-unconstrained"
    }
  }
  lifetime = "86400s"
  issuing_options {
    include_ca_cert_url = true
    include_crl_access_url = false
  }
  key_spec {
    algorithm = "EC_P256_SHA256"
  }
  disable_on_delete = true
}
`, context)
}

func TestAccPrivatecaCertificateAuthority_privatecaCertificateAuthorityCmekExample(t *testing.T) {
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"kms_key_name":  BootstrapKMSKeyWithPurposeInLocation(t, "ASYMMETRIC_SIGN", "us-central1").CryptoKey.Name,
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckPrivatecaCertificateAuthorityDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCertificateAuthority_privatecaCertificateAuthorityCmekExample(context),
			},
		},
	})
}

func testAccPrivatecaCertificateAuthority_privatecaCertificateAuthorityCmekExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project_service_identity" "privateca_sa" {
  provider = google-beta
  service  = "privateca.googleapis.com"
}

resource "google_kms_crypto_key_iam_binding" "privateca_sa_keyuser_signerverifier" {
  provider      = google-beta
  crypto_key_id = "%{kms_key_name}"
  role          = "roles/cloudkms.signerVerifier"

  members = [
    "serviceAccount:${google_project_service_identity.privateca_sa.email}",
  ]
}

resource "google_kms_crypto_key_iam_binding" "privateca_sa_keyuser_viewer" {
  provider      = google-beta
  crypto_key_id = "%{kms_key_name}"
  role          = "roles/viewer"
  members = [
    "serviceAccount:${google_project_service_identity.privateca_sa.email}",
  ]
}

resource "google_privateca_certificate_authority" "default" {
  provider                 = google-beta
  certificate_authority_id = "tf-test%{random_suffix}"
  location                 = "us-central1"

  key_spec {
    cloud_kms_key_version = "%{kms_key_name}/cryptoKeyVersions/1"
  }

  config  {
    subject_config  {
      common_name = "Example Authority"
      subject {
        organization = "Example, Org."
      }
    }

    reusable_config {
      reusable_config= "root-unconstrained"
    }
  }

  depends_on = [
    google_kms_crypto_key_iam_binding.privateca_sa_keyuser_signerverifier,
    google_kms_crypto_key_iam_binding.privateca_sa_keyuser_viewer,
  ]

  disable_on_delete = true
}
`, context)
}

func testAccCheckPrivatecaCertificateAuthorityDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_privateca_certificate_authority" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{PrivatecaBasePath}}projects/{{project}}/locations/{{location}}/certificateAuthorities/{{certificate_authority_id}}")
			if err != nil {
				return err
			}

			res, err := sendRequest(config, "GET", "", url, config.userAgent, nil)
			if err != nil {
				return nil
			}

			if s := res["state"]; s != "PENDING_DELETION" {
				return fmt.Errorf("CertificateAuthority %s got %s, want PENDING_DELETION", url, s)
			}
		}

		return nil
	}
}
