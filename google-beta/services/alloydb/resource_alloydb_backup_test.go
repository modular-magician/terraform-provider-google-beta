// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package alloydb_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccAlloydbBackup_update(t *testing.T) {
	t.Parallel()

	random_suffix := acctest.RandString(t, 10)
	context := map[string]interface{}{
		"network_name":  acctest.BootstrapSharedServiceNetworkingConnection(t, "alloydb-backup-update-1"),
		"random_suffix": random_suffix,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckAlloydbBackupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAlloydbBackup_alloydbBackupBasic(context),
			},
			{
				ResourceName:            "google_alloydb_backup.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backup_id", "location", "reconciling", "update_time", "labels", "terraform_labels"},
			},
			{
				Config: testAccAlloydbBackup_update(context),
			},
			{
				ResourceName:            "google_alloydb_backup.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backup_id", "location", "reconciling", "update_time", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccAlloydbBackup_alloydbBackupBasic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_alloydb_backup" "default" {
  location     = "us-central1"
  backup_id    = "tf-test-alloydb-backup%{random_suffix}"
  cluster_name = google_alloydb_cluster.default.name

  description = "example description"
  labels = {
    "label" = "key"
  }
  depends_on = [google_alloydb_instance.default]
}

resource "google_alloydb_cluster" "default" {
  cluster_id = "tf-test-alloydb-cluster%{random_suffix}"
  location   = "us-central1"
  network_config {
    network = data.google_compute_network.default.id
  }
}

resource "google_alloydb_instance" "default" {
  cluster       = google_alloydb_cluster.default.name
  instance_id   = "tf-test-alloydb-instance%{random_suffix}"
  instance_type = "PRIMARY"
}

data "google_compute_network" "default" {
  name = "%{network_name}"
}
`, context)
}

// Updates "label" field
func testAccAlloydbBackup_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_alloydb_backup" "default" {
  location     = "us-central1"
  backup_id    = "tf-test-alloydb-backup%{random_suffix}"
  cluster_name = google_alloydb_cluster.default.name

  description = "example description"
  labels = {
    "label" = "updated_key"
    "label2" = "updated_key2"
  }
  depends_on = [google_alloydb_instance.default]
}

resource "google_alloydb_cluster" "default" {
  cluster_id = "tf-test-alloydb-cluster%{random_suffix}"
  location   = "us-central1"
  network_config {
    network = data.google_compute_network.default.id
  }
}

resource "google_alloydb_instance" "default" {
  cluster       = google_alloydb_cluster.default.name
  instance_id   = "tf-test-alloydb-instance%{random_suffix}"
  instance_type = "PRIMARY"
}

data "google_compute_network" "default" {
  name = "%{network_name}"
}
`, context)
}

// Test to create on-demand backup with mandatory fields.
func TestAccAlloydbBackup_createBackupWithMandatoryFields(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"network_name":  acctest.BootstrapSharedServiceNetworkingConnection(t, "alloydb-backup-mandatory-1"),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckAlloydbBackupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAlloydbBackup_createBackupWithMandatoryFields(context),
			},
		},
	})
}

func testAccAlloydbBackup_createBackupWithMandatoryFields(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_alloydb_backup" "default" {
  backup_id    = "tf-test-alloydb-backup%{random_suffix}"
  location = "us-central1"
  cluster_name = google_alloydb_cluster.default.name
  depends_on = [google_alloydb_instance.default]
}

resource "google_alloydb_cluster" "default" {
  location = "us-central1"
  cluster_id = "tf-test-alloydb-cluster%{random_suffix}"
  network_config {
    network = data.google_compute_network.default.id
  }
}

data "google_project" "project" { }

data "google_compute_network" "default" {
  name = "%{network_name}"
}

resource "google_alloydb_instance" "default" {
  cluster       = google_alloydb_cluster.default.name
  instance_id   = "tf-test-alloydb-instance%{random_suffix}"
  instance_type = "PRIMARY"
}
`, context)
}

func TestAccAlloydbBackup_usingCMEK(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"network_name":  acctest.BootstrapSharedServiceNetworkingConnection(t, "alloydb-backup-cmek-1"),
		"random_suffix": acctest.RandString(t, 10),
		"key_name":      "tf-test-key-" + acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckAlloydbBackupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAlloydbBackup_usingCMEK(context),
			},
			{
				ResourceName:            "google_alloydb_backup.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backup_id", "location", "reconciling", "update_time", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccAlloydbBackup_usingCMEK(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_alloydb_backup" "default" {
	location     = "us-central1"
	backup_id    = "tf-test-alloydb-backup%{random_suffix}"
	cluster_name = google_alloydb_cluster.default.name
	description = "example description"
	labels = {
		"label" = "updated_key"
		"label2" = "updated_key2"
	}
	encryption_config {
		kms_key_name = google_kms_crypto_key.key.id
	}
	depends_on = [
		google_alloydb_instance.default,
		google_kms_crypto_key_iam_member.crypto_key
	]
}
	  
resource "google_alloydb_cluster" "default" {
	cluster_id = "tf-test-alloydb-cluster%{random_suffix}"
	location   = "us-central1"
  network_config {
    network = data.google_compute_network.default.id
  }
}
	  
resource "google_alloydb_instance" "default" {
	cluster       = google_alloydb_cluster.default.name
	instance_id   = "tf-test-alloydb-instance%{random_suffix}"
	instance_type = "PRIMARY"
}
	  
data "google_compute_network" "default" {
	name = "%{network_name}"
}
data "google_project" "project" {}

resource "google_kms_key_ring" "keyring" {
  name     = "%{key_name}"
  location = "us-central1"
}

resource "google_kms_crypto_key" "key" {
  name     = "%{key_name}"
  key_ring = google_kms_key_ring.keyring.id
}

resource "google_kms_crypto_key_iam_member" "crypto_key" {
  crypto_key_id = google_kms_crypto_key.key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-alloydb.iam.gserviceaccount.com"
}
`, context)
}

func TestAccAlloydbBackup_tags(t *testing.T) {
	t.Parallel()
	
	random_suffix := acctest.RandString(t, 10)
	context := map[string]interface{}{
		"network_name":  acctest.BootstrapSharedServiceNetworkingConnection(t, "alloydb-backup-update-1"),
		"random_suffix": random_suffix,
	}
	tagKey := acctest.BootstrapSharedTestTagKey(t, "alloydb-backups-tagkey")
	tagValue := acctest.BootstrapSharedTestTagValue(t, "alloydb-backups-tagvalue", tagKey)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckAlloydbBackupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAlloydbBackupTags(context, map[string]string{tagKey: tagValue}),
			},
			{
				ResourceName:            "google_alloydb_backup.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backup_id", "location", "reconciling", "update_time", "labels", "terraform_labels", "tags"},
			},
		},
	})
}

func testAccAlloydbBackupTags(context map[string]interface{}, tags map[string]string) string {

	return acctest.Nprintf(`

resource "google_alloydb_cluster" "default" {
  cluster_id = "a12-alloydb-cluster"
  location   = "us-west2"
  network_config {
    network = google_compute_network.default.id
  }
}

resource "google_alloydb_instance" "default" {
  cluster       = google_alloydb_cluster.default.name
  instance_id   = "a12-alloydb-instance"
  instance_type = "PRIMARY"

  depends_on = [google_service_networking_connection.vpc_connection]
}

resource "google_compute_global_address" "private_ip_alloc" {
  name          =  "a12-alloydb-cluster"
  address_type  = "INTERNAL"
  purpose       = "VPC_PEERING"
  prefix_length = 16
  network       = google_compute_network.default.id
}

resource "google_service_networking_connection" "vpc_connection" {
  network                 = google_compute_network.default.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_alloc.name]
}

resource "google_compute_network" "default" {
  name = "a12-alloydb-network"
}
resource "google_alloydb_backup" "default" {
  location     = "us-west2"
  backup_id    = "a12-alloydb-backup"
  cluster_name = google_alloydb_cluster.default.name
  depends_on = [google_alloydb_instance.default]
  tags = {
	"tagKeys/281478409127147" = "tagValues/281479442205542"
}
}
`, context)
}
