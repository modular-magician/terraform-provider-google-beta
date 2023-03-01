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
)

func TestAccWorkstationsWorkstationConfig_workstationConfigBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_workstationConfigBasicExample(context),
			},
			{
				ResourceName:            "google_workstations_workstation_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workstation_config_id", "workstation_cluster_id", "location"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_workstationConfigBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider               = google-beta
  workstation_cluster_id = "tf-test-workstation-cluster%{random_suffix}"
  network                = google_compute_network.default.id
  subnetwork             = google_compute_subnetwork.default.id
  location               = "us-central1"
  
  labels = {
    "label" = "key"
  }

  annotations = {
    label-one = "value-one"
  }
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		         = "us-central1"
  
  host {
    gce_instance {
      machine_type                = "e2-standard-4"
      boot_disk_size_gb           = 35
      disable_public_ip_addresses = true
    }
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_workstationConfigContainerExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_workstationConfigContainerExample(context),
			},
			{
				ResourceName:            "google_workstations_workstation_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workstation_config_id", "workstation_cluster_id", "location"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_workstationConfigContainerExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider               = google-beta
  workstation_cluster_id = "tf-test-workstation-cluster%{random_suffix}"
  network                = google_compute_network.default.id
  subnetwork             = google_compute_subnetwork.default.id
  location               = "us-central1"
  
  labels = {
    "label" = "key"
  }

  annotations = {
    label-one = "value-one"
  }
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		         = "us-central1"
  
  host {
    gce_instance {
      machine_type                = "e2-standard-4"
      boot_disk_size_gb           = 35
      disable_public_ip_addresses = true
    }
  }

  container {
    image = "intellij"
    env = {
      NAME = "FOO"
      BABE = "bar"
    }
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_workstationConfigPersistentDirectoriesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_workstationConfigPersistentDirectoriesExample(context),
			},
			{
				ResourceName:            "google_workstations_workstation_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workstation_config_id", "workstation_cluster_id", "location"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_workstationConfigPersistentDirectoriesExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider               = google-beta
  workstation_cluster_id = "tf-test-workstation-cluster%{random_suffix}"
  network                = google_compute_network.default.id
  subnetwork             = google_compute_subnetwork.default.id
  location               = "us-central1"
  
  labels = {
    "label" = "key"
  }

  annotations = {
    label-one = "value-one"
  }
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		         = "us-central1"

  host {
    gce_instance {
      machine_type                = "e2-standard-4"
      boot_disk_size_gb           = 35
      disable_public_ip_addresses = true
      shielded_instance_config {
        enable_secure_boot = true
        enable_vtpm        = true
      }
    }
  }

  persistent_directories {
    mount_path = "/home"
    gce_pd {
      size_gb        = 200
      reclaim_policy = "DELETE"
    }
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_workstationConfigShieldedInstanceConfigExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_workstationConfigShieldedInstanceConfigExample(context),
			},
			{
				ResourceName:            "google_workstations_workstation_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workstation_config_id", "workstation_cluster_id", "location"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_workstationConfigShieldedInstanceConfigExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider               = google-beta
  workstation_cluster_id = "tf-test-workstation-cluster%{random_suffix}"
  network                = google_compute_network.default.id
  subnetwork             = google_compute_subnetwork.default.id
  location               = "us-central1"
  
  labels = {
    "label" = "key"
  }

  annotations = {
    label-one = "value-one"
  }
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		         = "us-central1"

  host {
    gce_instance {
      machine_type                = "e2-standard-4"
      boot_disk_size_gb           = 35
      disable_public_ip_addresses = true
      shielded_instance_config {
        enable_secure_boot = true
        enable_vtpm        = true
      }
    }
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_workstationConfigEncryptionKeyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_workstationConfigEncryptionKeyExample(context),
			},
			{
				ResourceName:            "google_workstations_workstation_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workstation_config_id", "workstation_cluster_id", "location"},
			},
		},
	})
}

func testAccWorkstationsWorkstationConfig_workstationConfigEncryptionKeyExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider               = google-beta
  workstation_cluster_id = "tf-test-workstation-cluster%{random_suffix}"
  network                = google_compute_network.default.id
  subnetwork             = google_compute_subnetwork.default.id
  location               = "us-central1"
  
  labels = {
    "label" = "key"
  }

  annotations = {
    label-one = "value-one"
  }
}

resource "google_kms_key_ring" "default" {
  name     = "tf-test-workstation-cluster%{random_suffix}"
  location = "global"
  provider = google-beta
}

resource "google_kms_crypto_key" "default" {
  name            = "tf-test-workstation-cluster%{random_suffix}"
  key_ring        = google_kms_key_ring.default.id
  rotation_period = "100000s"
  provider        = google-beta
}

resource "google_service_account" "default" {
  account_id   = "tf-test-my-account%{random_suffix}"
  display_name = "Service Account"
  provider = google-beta
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		         = "us-central1"

  host {
    gce_instance {
      machine_type                = "e2-standard-4"
      boot_disk_size_gb           = 35
      disable_public_ip_addresses = true
      shielded_instance_config {
        enable_secure_boot = true
        enable_vtpm        = true
      }
    }
  }

  encryption_key {
    kms_key                 = google_kms_crypto_key.default.id
    kms_key_service_account = google_service_account.default.email
  }
}
`, context)
}

func testAccCheckWorkstationsWorkstationConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_workstations_workstation_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{WorkstationsBasePath}}projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("WorkstationsWorkstationConfig still exists at %s", url)
			}
		}

		return nil
	}
}
