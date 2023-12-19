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

package gkebackup_test

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

func TestAccGKEBackupRestorePlan_gkebackupRestoreplanAllNamespacesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":             envvar.GetTestProjectFromEnv(),
		"deletion_protection": false,
		"network_name":        acctest.BootstrapSharedTestNetwork(t, "gke-cluster"),
		"subnetwork_name":     acctest.BootstrapSubnet(t, "gke-cluster", acctest.BootstrapSharedTestNetwork(t, "gke-cluster")),
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGKEBackupRestorePlanDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEBackupRestorePlan_gkebackupRestoreplanAllNamespacesExample(context),
			},
			{
				ResourceName:            "google_gke_backup_restore_plan.all_ns",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccGKEBackupRestorePlan_gkebackupRestoreplanAllNamespacesExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-restore-all-ns%{random_suffix}-cluster"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "%{project}.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
  deletion_protection  = "%{deletion_protection}"
  network       = "%{network_name}"
  subnetwork    = "%{subnetwork_name}"
}

resource "google_gke_backup_backup_plan" "basic" {
  name = "tf-test-restore-all-ns%{random_suffix}"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}

resource "google_gke_backup_restore_plan" "all_ns" {
  name = "tf-test-restore-all-ns%{random_suffix}"
  location = "us-central1"
  backup_plan = google_gke_backup_backup_plan.basic.id
  cluster = google_container_cluster.primary.id
  restore_config {
    all_namespaces = true
    namespaced_resource_restore_mode = "FAIL_ON_CONFLICT"
    volume_data_restore_policy = "RESTORE_VOLUME_DATA_FROM_BACKUP"
    cluster_resource_restore_scope {
      all_group_kinds = true
    }
    cluster_resource_conflict_policy = "USE_EXISTING_VERSION"
  }
}
`, context)
}

func TestAccGKEBackupRestorePlan_gkebackupRestoreplanRollbackNamespaceExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":             envvar.GetTestProjectFromEnv(),
		"deletion_protection": false,
		"network_name":        acctest.BootstrapSharedTestNetwork(t, "gke-cluster"),
		"subnetwork_name":     acctest.BootstrapSubnet(t, "gke-cluster", acctest.BootstrapSharedTestNetwork(t, "gke-cluster")),
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGKEBackupRestorePlanDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEBackupRestorePlan_gkebackupRestoreplanRollbackNamespaceExample(context),
			},
			{
				ResourceName:            "google_gke_backup_restore_plan.rollback_ns",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccGKEBackupRestorePlan_gkebackupRestoreplanRollbackNamespaceExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-rollback-ns%{random_suffix}-cluster"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "%{project}.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
  deletion_protection  = "%{deletion_protection}"
  network       = "%{network_name}"
  subnetwork    = "%{subnetwork_name}"
}

resource "google_gke_backup_backup_plan" "basic" {
  name = "tf-test-rollback-ns%{random_suffix}"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}

resource "google_gke_backup_restore_plan" "rollback_ns" {
  name = "tf-test-rollback-ns%{random_suffix}-rp"
  location = "us-central1"
  backup_plan = google_gke_backup_backup_plan.basic.id
  cluster = google_container_cluster.primary.id
  restore_config {
    selected_namespaces {
      namespaces = ["my-ns"]
    }
    namespaced_resource_restore_mode = "DELETE_AND_RESTORE"
    volume_data_restore_policy = "RESTORE_VOLUME_DATA_FROM_BACKUP"
    cluster_resource_restore_scope {
      selected_group_kinds {
        resource_group = "apiextension.k8s.io"
        resource_kind = "CustomResourceDefinition"
      }
      selected_group_kinds {
        resource_group = "storage.k8s.io"
        resource_kind = "StorageClass"
      }
    }
    cluster_resource_conflict_policy = "USE_EXISTING_VERSION"
  }
}
`, context)
}

func TestAccGKEBackupRestorePlan_gkebackupRestoreplanProtectedApplicationExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":             envvar.GetTestProjectFromEnv(),
		"deletion_protection": false,
		"network_name":        acctest.BootstrapSharedTestNetwork(t, "gke-cluster"),
		"subnetwork_name":     acctest.BootstrapSubnet(t, "gke-cluster", acctest.BootstrapSharedTestNetwork(t, "gke-cluster")),
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGKEBackupRestorePlanDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEBackupRestorePlan_gkebackupRestoreplanProtectedApplicationExample(context),
			},
			{
				ResourceName:            "google_gke_backup_restore_plan.rollback_app",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccGKEBackupRestorePlan_gkebackupRestoreplanProtectedApplicationExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-rollback-app%{random_suffix}-cluster"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "%{project}.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
  deletion_protection  = "%{deletion_protection}"
  network       = "%{network_name}"
  subnetwork    = "%{subnetwork_name}"
}

resource "google_gke_backup_backup_plan" "basic" {
  name = "tf-test-rollback-app%{random_suffix}"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}

resource "google_gke_backup_restore_plan" "rollback_app" {
  name = "tf-test-rollback-app%{random_suffix}-rp"
  location = "us-central1"
  backup_plan = google_gke_backup_backup_plan.basic.id
  cluster = google_container_cluster.primary.id
  restore_config {
    selected_applications {
      namespaced_names {
        name = "my-app"
        namespace = "my-ns"
      }
    }
    namespaced_resource_restore_mode = "DELETE_AND_RESTORE"
    volume_data_restore_policy = "REUSE_VOLUME_HANDLE_FROM_BACKUP"
    cluster_resource_restore_scope {
      no_group_kinds = true
    }
  }
}
`, context)
}

func TestAccGKEBackupRestorePlan_gkebackupRestoreplanAllClusterResourcesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":             envvar.GetTestProjectFromEnv(),
		"deletion_protection": false,
		"network_name":        acctest.BootstrapSharedTestNetwork(t, "gke-cluster"),
		"subnetwork_name":     acctest.BootstrapSubnet(t, "gke-cluster", acctest.BootstrapSharedTestNetwork(t, "gke-cluster")),
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGKEBackupRestorePlanDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEBackupRestorePlan_gkebackupRestoreplanAllClusterResourcesExample(context),
			},
			{
				ResourceName:            "google_gke_backup_restore_plan.all_cluster_resources",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccGKEBackupRestorePlan_gkebackupRestoreplanAllClusterResourcesExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-all-groupkinds%{random_suffix}-cluster"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "%{project}.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
  deletion_protection  = "%{deletion_protection}"
  network       = "%{network_name}"
  subnetwork    = "%{subnetwork_name}"
}

resource "google_gke_backup_backup_plan" "basic" {
  name = "tf-test-all-groupkinds%{random_suffix}"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}

resource "google_gke_backup_restore_plan" "all_cluster_resources" {
  name = "tf-test-all-groupkinds%{random_suffix}-rp"
  location = "us-central1"
  backup_plan = google_gke_backup_backup_plan.basic.id
  cluster = google_container_cluster.primary.id
  restore_config {
    no_namespaces = true
    namespaced_resource_restore_mode = "FAIL_ON_CONFLICT"
    cluster_resource_restore_scope {
      all_group_kinds = true
    }
    cluster_resource_conflict_policy = "USE_EXISTING_VERSION"
  }
}
`, context)
}

func TestAccGKEBackupRestorePlan_gkebackupRestoreplanRenameNamespaceExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":             envvar.GetTestProjectFromEnv(),
		"deletion_protection": false,
		"network_name":        acctest.BootstrapSharedTestNetwork(t, "gke-cluster"),
		"subnetwork_name":     acctest.BootstrapSubnet(t, "gke-cluster", acctest.BootstrapSharedTestNetwork(t, "gke-cluster")),
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGKEBackupRestorePlanDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEBackupRestorePlan_gkebackupRestoreplanRenameNamespaceExample(context),
			},
			{
				ResourceName:            "google_gke_backup_restore_plan.rename_ns",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccGKEBackupRestorePlan_gkebackupRestoreplanRenameNamespaceExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-rename-ns%{random_suffix}-cluster"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "%{project}.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
  deletion_protection  = "%{deletion_protection}"
  network       = "%{network_name}"
  subnetwork    = "%{subnetwork_name}"
}

resource "google_gke_backup_backup_plan" "basic" {
  name = "tf-test-rename-ns%{random_suffix}"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}

resource "google_gke_backup_restore_plan" "rename_ns" {
  name = "tf-test-rename-ns%{random_suffix}-rp"
  location = "us-central1"
  backup_plan = google_gke_backup_backup_plan.basic.id
  cluster = google_container_cluster.primary.id
  restore_config {
    selected_namespaces {
      namespaces = ["ns1"]
    }
    namespaced_resource_restore_mode = "FAIL_ON_CONFLICT"
    volume_data_restore_policy = "REUSE_VOLUME_HANDLE_FROM_BACKUP"
    cluster_resource_restore_scope {
      no_group_kinds = true
    }
    transformation_rules {
      description = "rename namespace from ns1 to ns2"
      resource_filter {
        group_kinds {
          resource_kind = "Namespace"
        }
        json_path = ".metadata[?(@.name == 'ns1')]"
      }
      field_actions {
        op = "REPLACE"
        path = "/metadata/name"
        value = "ns2"
      }
    }
    transformation_rules {
      description = "move all resources from ns1 to ns2"
      resource_filter {
        namespaces = ["ns1"]
      }
      field_actions {
        op = "REPLACE"
        path = "/metadata/namespace"
        value = "ns2"
      }
    }
  }
}
`, context)
}

func TestAccGKEBackupRestorePlan_gkebackupRestoreplanSecondTransformationExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":             envvar.GetTestProjectFromEnv(),
		"deletion_protection": false,
		"network_name":        acctest.BootstrapSharedTestNetwork(t, "gke-cluster"),
		"subnetwork_name":     acctest.BootstrapSubnet(t, "gke-cluster", acctest.BootstrapSharedTestNetwork(t, "gke-cluster")),
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGKEBackupRestorePlanDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEBackupRestorePlan_gkebackupRestoreplanSecondTransformationExample(context),
			},
			{
				ResourceName:            "google_gke_backup_restore_plan.transform_rule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccGKEBackupRestorePlan_gkebackupRestoreplanSecondTransformationExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-transform-rule%{random_suffix}-cluster"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "%{project}.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
  deletion_protection  = "%{deletion_protection}"
  network       = "%{network_name}"
  subnetwork    = "%{subnetwork_name}"
}

resource "google_gke_backup_backup_plan" "basic" {
  name = "tf-test-transform-rule%{random_suffix}"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}

resource "google_gke_backup_restore_plan" "transform_rule" {
  name = "tf-test-transform-rule%{random_suffix}-rp"
  description = "copy nginx env variables"
  labels = {
    "app" = "nginx"
  }
  location = "us-central1"
  backup_plan = google_gke_backup_backup_plan.basic.id
  cluster = google_container_cluster.primary.id
  restore_config {
    excluded_namespaces {
      namespaces = ["my-ns"]
    }
    namespaced_resource_restore_mode = "DELETE_AND_RESTORE"
    volume_data_restore_policy = "RESTORE_VOLUME_DATA_FROM_BACKUP"
    cluster_resource_restore_scope {
      excluded_group_kinds {
        resource_group = "apiextension.k8s.io"
        resource_kind = "CustomResourceDefinition"
      }
    }
    cluster_resource_conflict_policy = "USE_EXISTING_VERSION"
    transformation_rules {
      description = "Copy environment variables from the nginx container to the install init container."
      resource_filter {
        group_kinds {
          resource_kind = "Pod"
          resource_group = ""
        }
        json_path = ".metadata[?(@.name == 'nginx')]"
      }
      field_actions {
        op = "COPY"
        path = "/spec/initContainers/0/env"
        from_path = "/spec/containers/0/env"
      }
    }
  }
}
`, context)
}

func testAccCheckGKEBackupRestorePlanDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_gke_backup_restore_plan" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{GKEBackupBasePath}}projects/{{project}}/locations/{{location}}/restorePlans/{{name}}")
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
				return fmt.Errorf("GKEBackupRestorePlan still exists at %s", url)
			}
		}

		return nil
	}
}
