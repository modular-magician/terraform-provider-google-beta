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

package looker_test

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

func TestAccLookerInstance_lookerInstanceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckLookerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccLookerInstance_lookerInstanceBasicExample(context),
			},
			{
				ResourceName:            "google_looker_instance.looker-instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "oauth_config", "region"},
			},
		},
	})
}

func testAccLookerInstance_lookerInstanceBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_looker_instance" "looker-instance" {
  name              = "tf-test-my-instance%{random_suffix}"
  platform_edition  = "LOOKER_CORE_STANDARD"
  region            = "us-central1"
  oauth_config {
    client_id = "tf-test-my-client-id%{random_suffix}"
    client_secret = "tf-test-my-client-secret%{random_suffix}"
  }
}
`, context)
}

func TestAccLookerInstance_lookerInstanceFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckLookerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccLookerInstance_lookerInstanceFullExample(context),
			},
			{
				ResourceName:            "google_looker_instance.looker-instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "oauth_config", "region"},
			},
		},
	})
}

func testAccLookerInstance_lookerInstanceFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_looker_instance" "looker-instance" {
  name               = "tf-test-my-instance%{random_suffix}"
  platform_edition   = "LOOKER_CORE_STANDARD"
  region             = "us-central1"
  public_ip_enabled  = true
  admin_settings {
    allowed_email_domains = ["google.com"]
  }
  // User metadata config is only available when platform edition is LOOKER_CORE_STANDARD.
  user_metadata {
    additional_developer_user_count = 10 
    additional_standard_user_count  = 10
    additional_viewer_user_count    = 10
  }
  maintenance_window {
    day_of_week = "THURSDAY"
    start_time {
      hours   = 22
      minutes = 0
      seconds = 0
      nanos   = 0
    }
  }
  deny_maintenance_period {    
    start_date {
      year = 2050
      month = 1
      day = 1
    }
    end_date {
      year = 2050
      month = 2
      day = 1
    }
    time {
      hours = 10
      minutes = 0
      seconds = 0
      nanos = 0
    }
  }
  oauth_config {
    client_id = "tf-test-my-client-id%{random_suffix}"
    client_secret = "tf-test-my-client-secret%{random_suffix}"
  }  
}
`, context)
}

func TestAccLookerInstance_lookerInstanceEnterpriseFullTestExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"address_name":  acctest.BootstrapSharedTestGlobalAddress(t, "looker-vpc-network-1", acctest.AddressWithPrefixLength(20)),
		"kms_key_name":  acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"network_name":  acctest.BootstrapSharedServiceNetworkingConnection(t, "looker-vpc-network-1", acctest.ServiceNetworkWithPrefixLength(20)),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckLookerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccLookerInstance_lookerInstanceEnterpriseFullTestExample(context),
			},
			{
				ResourceName:            "google_looker_instance.looker-instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "oauth_config", "region"},
			},
		},
	})
}

func testAccLookerInstance_lookerInstanceEnterpriseFullTestExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_looker_instance" "looker-instance" {
  name               = "tf-test-my-instance%{random_suffix}"
  platform_edition   = "LOOKER_CORE_ENTERPRISE_ANNUAL"
  region             = "us-central1"
  private_ip_enabled = true
  public_ip_enabled  = false
  reserved_range     = "${data.google_compute_global_address.looker_range.name}"
  consumer_network   = data.google_compute_network.looker_network.id
  admin_settings {
    allowed_email_domains = ["google.com"]
  }
  encryption_config {
    kms_key_name = "%{kms_key_name}"
  }
  maintenance_window {
    day_of_week = "THURSDAY"
    start_time {
      hours   = 22
      minutes = 0
      seconds = 0
      nanos   = 0
    }
  }
  deny_maintenance_period {
    start_date {
      year = 2050
      month = 1
      day = 1
    }
    end_date {
      year = 2050
      month = 2
      day = 1
    }
    time {
      hours = 10
      minutes = 0
      seconds = 0
      nanos = 0
    }
  }
  oauth_config {
    client_id = "tf-test-my-client-id%{random_suffix}"
    client_secret = "tf-test-my-client-secret%{random_suffix}"
  }
}

data "google_compute_global_address" "looker_range" {
  name          = "%{address_name}"
}

data "google_project" "project" {}

data "google_compute_network" "looker_network" {
  name = "%{network_name}"
}

resource "google_kms_crypto_key_iam_member" "crypto_key" {
  crypto_key_id = "%{kms_key_name}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-looker.iam.gserviceaccount.com"
}
`, context)
}

func testAccCheckLookerInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_looker_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{LookerBasePath}}projects/{{project}}/locations/{{region}}/instances/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:               config,
				Method:               "GET",
				Project:              billingProject,
				RawURL:               url,
				UserAgent:            config.UserAgent,
				ErrorAbortPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.Is429QuotaError},
			})
			if err == nil {
				return fmt.Errorf("LookerInstance still exists at %s", url)
			}
		}

		return nil
	}
}
