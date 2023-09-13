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

package alloydb_test

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

func TestAccAlloydbCluster_alloydbClusterBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckAlloydbClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAlloydbCluster_alloydbClusterBasicExample(context),
			},
			{
				ResourceName:            "google_alloydb_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "initial_user", "restore_backup_source", "restore_continuous_backup_source", "terraform_labels", "cluster_id", "location"},
			},
		},
	})
}

func testAccAlloydbCluster_alloydbClusterBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_alloydb_cluster" "default" {
  cluster_id = "tf-test-alloydb-cluster%{random_suffix}"
  location   = "us-central1"
  network    = google_compute_network.default.id
}

data "google_project" "project" {}

resource "google_compute_network" "default" {
  name = "tf-test-alloydb-cluster%{random_suffix}"
}
`, context)
}

func TestAccAlloydbCluster_alloydbClusterFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckAlloydbClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAlloydbCluster_alloydbClusterFullExample(context),
			},
			{
				ResourceName:            "google_alloydb_cluster.full",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "initial_user", "restore_backup_source", "restore_continuous_backup_source", "terraform_labels", "cluster_id", "location"},
			},
		},
	})
}

func testAccAlloydbCluster_alloydbClusterFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_alloydb_cluster" "full" {
  cluster_id   = "tf-test-alloydb-cluster-full%{random_suffix}"
  location     = "us-central1"
  network      = google_compute_network.default.id

  initial_user {
    user     = "tf-test-alloydb-cluster-full%{random_suffix}"
    password = "tf-test-alloydb-cluster-full%{random_suffix}"
  }

  continuous_backup_config {
    enabled              = true
    recovery_window_days = 14
  }

  automated_backup_policy {
    location      = "us-central1"
    backup_window = "1800s"
    enabled       = true

    weekly_schedule {
      days_of_week = ["MONDAY"]

      start_times {
        hours   = 23
        minutes = 0
        seconds = 0
        nanos   = 0
      }
    }

    quantity_based_retention {
      count = 1
    }

    labels = {
      test = "tf-test-alloydb-cluster-full%{random_suffix}"
    }
  }

  labels = {
    test = "tf-test-alloydb-cluster-full%{random_suffix}"
  }
}

data "google_project" "project" {}

resource "google_compute_network" "default" {
  name = "tf-test-alloydb-cluster-full%{random_suffix}"
}
`, context)
}

func testAccCheckAlloydbClusterDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_alloydb_cluster" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{AlloydbBasePath}}projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}")
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
				return fmt.Errorf("AlloydbCluster still exists at %s", url)
			}
		}

		return nil
	}
}
