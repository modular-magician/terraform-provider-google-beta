// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package backupdr_test

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"strings"
	"testing"
)

func TestAccDataSourceGoogleBackupDRBackupPlan_basic(t *testing.T) {
	t.Parallel()
	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleBackupDRBackupPlan_basic(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceState("data.google_backup_dr_backup_plan.foo", "google_backup_dr_backup_plan.foo"),
				),
			},
		},
	})
}

func testAccCheckBackupDRBackupPlanDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_backup_dr_backup_plan" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{BackupDRBasePath}}projects/{{project}}/locations/{{location}}/backupPlans/{{backup_plan_id}}")
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
				return fmt.Errorf("Backup Plan still exists at %s", url)
			}
		}

		return nil
	}
}

func testAccDataSourceGoogleBackupDRBackupPlan_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_backup_dr_backup_vault" "my-backup-vault" {
    location ="us-central1"
    backup_vault_id    = "bv-1"
    description = "This is a second backup vault built by Terraform."
    backup_minimum_enforced_retention_duration = "100000s"
    labels = {
      foo = "bar1"
      bar = "baz1"
    }
    annotations = {
      annotations1 = "bar1"
      annotations2 = "baz1"
    }
    force_update = "true"
    force_delete = "true"
    allow_missing = "true" 
}


resource "google_backup_dr_backup_plan" "foo" { 
  location = "us-central1" 
  backup_plan_id = "bp-test-tf1"
  resource_type= "compute.googleapis.com/Instance"
  backup_vault = google_backup_dr_backup_vault.my-backup-vault.name
  depends_on=[ google_backup_dr_backup_vault.my-backup-vault ]
  backup_rules {
	rule_id = "rule-1"
	backup_retention_days = 5
	standard_schedule {
	  recurrence_type = "HOURLY"
	   hourly_frequency = 6
	    time_zone = "UTC"
	     backup_window{
		start_hour_of_day = 0
		end_hour_of_day = 24
      }
    }
	}
}

data "google_backup_dr_backup_plan" "foo" {
  location =  "us-central1"
  backup_plan_id="bp-test-tf1"
  depends_on= [ google_backup_dr_backup_plan.foo ]
  }
`, context)
}
