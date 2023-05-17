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

func TestAccComputeMachineImage_machineImageBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeMachineImage_machineImageBasicExample(context),
			},
			{
				ResourceName:            "google_compute_machine_image.image",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"source_instance"},
			},
		},
	})
}

func testAccComputeMachineImage_machineImageBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "tf-test-my-vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "tf-test-my-image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
}
`, context)
}

func TestAccComputeMachineImage_computeMachineImageKmsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"policyChanged": BootstrapPSARole(t, "service-", "compute-system", "roles/cloudkms.cryptoKeyEncrypterDecrypter"),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeMachineImage_computeMachineImageKmsExample(context),
			},
			{
				ResourceName:            "google_compute_machine_image.image",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"source_instance"},
			},
		},
	})
}

func testAccComputeMachineImage_computeMachineImageKmsExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "tf-test-my-vm%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "tf-test-my-image%{random_suffix}"
  source_instance = google_compute_instance.vm.self_link
  machine_image_encryption_key {
    kms_key_name = google_kms_crypto_key.crypto_key.id
  }
}

resource "google_kms_crypto_key" "crypto_key" {
  provider = google-beta
  name     = "key%{random_suffix}"
  key_ring = google_kms_key_ring.key_ring.id
}

resource "google_kms_key_ring" "key_ring" {
  provider = google-beta
  name     = "keyring%{random_suffix}"
  location = "us"
}
`, context)
}

func testAccCheckComputeMachineImageDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_machine_image" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/machineImages/{{name}}")
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
				return fmt.Errorf("ComputeMachineImage still exists at %s", url)
			}
		}

		return nil
	}
}
