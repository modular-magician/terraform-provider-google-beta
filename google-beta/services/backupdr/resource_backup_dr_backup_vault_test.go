// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package backupdr_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccBackupDRBackupVault_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBackupDRBackupVault_basic(context),
			},
			{
				ResourceName:            "google_backup_dr_backup_vault.backup-vault-test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "backup_vault_id", "labels", "location", "terraform_labels"},
			},
			{
				Config: testAccBackupDRBackupVault_update(context),
			},
			{
				ResourceName:            "google_backup_dr_backup_vault.backup-vault-test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "backup_vault_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccBackupDRBackupVault_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
	resource "google_backup_dr_backup_vault" "backup-vault-test" {
	provider = google-beta
	location = "us-central1"
	backup_vault_id    = "tf-test-backup-vault-test%{random_suffix}"
	description = "This is a backup vault built by Terraform."
	backup_minimum_enforced_retention_duration = "100000s"
	}
	`, context)
}

func testAccBackupDRBackupVault_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
	resource "google_backup_dr_backup_vault" "backup-vault-test" {
	provider = google-beta
	location = "us-central1"
	backup_vault_id    = "tf-test-backup-vault-test%{random_suffix}"
	description = "Terraform BV updated."
	backup_minimum_enforced_retention_duration = "200000s"
	}
	`, context)
}
