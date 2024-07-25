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

package workstations_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccWorkstationsWorkstationConfigIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":    acctest.RandString(t, 10),
		"role":             "roles/viewer",
		"key_short_name":   "tf-test-key-" + acctest.RandString(t, 10),
		"value_short_name": "tf-test-value-" + acctest.RandString(t, 10),
		"org_id":           envvar.GetTestOrgFromEnv(t),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfigIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_workstations_workstation_config_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-workstation-cluster%s", context["random_suffix"]), fmt.Sprintf("tf-test-workstation-config%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccWorkstationsWorkstationConfigIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_workstations_workstation_config_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-workstation-cluster%s", context["random_suffix"]), fmt.Sprintf("tf-test-workstation-config%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccWorkstationsWorkstationConfigIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":    acctest.RandString(t, 10),
		"role":             "roles/viewer",
		"key_short_name":   "tf-test-key-" + acctest.RandString(t, 10),
		"value_short_name": "tf-test-value-" + acctest.RandString(t, 10),
		"org_id":           envvar.GetTestOrgFromEnv(t),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccWorkstationsWorkstationConfigIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_workstations_workstation_config_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-workstation-cluster%s", context["random_suffix"]), fmt.Sprintf("tf-test-workstation-config%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccWorkstationsWorkstationConfigIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":    acctest.RandString(t, 10),
		"role":             "roles/viewer",
		"key_short_name":   "tf-test-key-" + acctest.RandString(t, 10),
		"value_short_name": "tf-test-value-" + acctest.RandString(t, 10),
		"org_id":           envvar.GetTestOrgFromEnv(t),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfigIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_workstations_workstation_config_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_workstations_workstation_config_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-workstation-cluster%s", context["random_suffix"]), fmt.Sprintf("tf-test-workstation-config%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccWorkstationsWorkstationConfigIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_workstations_workstation_config_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-workstation-cluster%s", context["random_suffix"]), fmt.Sprintf("tf-test-workstation-config%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccWorkstationsWorkstationConfigIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_tags_tag_key" "tag_key1" {
  provider   = google-beta
  parent     = "organizations/%{org_id}"
  short_name = "%{key_short_name}"
}

resource "google_tags_tag_value" "tag_value1" {
  provider   = google-beta
  parent     = "tagKeys/${google_tags_tag_key.tag_key1.name}"
  short_name = "%{value_short_name}"
}

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

  idle_timeout = "600s"
  running_timeout = "21600s"

  replica_zones = ["us-central1-a", "us-central1-b"]
  annotations = {
    label-one = "value-one"
  }

  labels = {
    "label" = "key"
  }

  host {
    gce_instance {
      machine_type                = "e2-standard-4"
      boot_disk_size_gb           = 35
      disable_public_ip_addresses = true
      disable_ssh                 = false
      vm_tags = {
        "tagKeys/${google_tags_tag_key.tag_key1.name}" = "tagValues/${google_tags_tag_value.tag_value1.name}"
      }
    }
  }
  allowed_ports = [
    {
      first: 22
      last:22
    },
    {
      first: 80
      last: 80
    },
    {
      first: 8000
      last: 9000
    }
  ]
}

resource "google_workstations_workstation_config_iam_member" "foo" {
  provider = google-beta
  project = google_workstations_workstation_config.default.project
  location = google_workstations_workstation_config.default.location
  workstation_cluster_id = google_workstations_workstation_config.default.workstation_cluster_id
  workstation_config_id = google_workstations_workstation_config.default.workstation_config_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccWorkstationsWorkstationConfigIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_tags_tag_key" "tag_key1" {
  provider   = google-beta
  parent     = "organizations/%{org_id}"
  short_name = "%{key_short_name}"
}

resource "google_tags_tag_value" "tag_value1" {
  provider   = google-beta
  parent     = "tagKeys/${google_tags_tag_key.tag_key1.name}"
  short_name = "%{value_short_name}"
}

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

  idle_timeout = "600s"
  running_timeout = "21600s"

  replica_zones = ["us-central1-a", "us-central1-b"]
  annotations = {
    label-one = "value-one"
  }

  labels = {
    "label" = "key"
  }

  host {
    gce_instance {
      machine_type                = "e2-standard-4"
      boot_disk_size_gb           = 35
      disable_public_ip_addresses = true
      disable_ssh                 = false
      vm_tags = {
        "tagKeys/${google_tags_tag_key.tag_key1.name}" = "tagValues/${google_tags_tag_value.tag_value1.name}"
      }
    }
  }
  allowed_ports = [
    {
      first: 22
      last:22
    },
    {
      first: 80
      last: 80
    },
    {
      first: 8000
      last: 9000
    }
  ]
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_workstations_workstation_config_iam_policy" "foo" {
  provider = google-beta
  project = google_workstations_workstation_config.default.project
  location = google_workstations_workstation_config.default.location
  workstation_cluster_id = google_workstations_workstation_config.default.workstation_cluster_id
  workstation_config_id = google_workstations_workstation_config.default.workstation_config_id
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_workstations_workstation_config_iam_policy" "foo" {
  provider = google-beta
  project = google_workstations_workstation_config.default.project
  location = google_workstations_workstation_config.default.location
  workstation_cluster_id = google_workstations_workstation_config.default.workstation_cluster_id
  workstation_config_id = google_workstations_workstation_config.default.workstation_config_id
  depends_on = [
    google_workstations_workstation_config_iam_policy.foo
  ]
}
`, context)
}

func testAccWorkstationsWorkstationConfigIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_tags_tag_key" "tag_key1" {
  provider   = google-beta
  parent     = "organizations/%{org_id}"
  short_name = "%{key_short_name}"
}

resource "google_tags_tag_value" "tag_value1" {
  provider   = google-beta
  parent     = "tagKeys/${google_tags_tag_key.tag_key1.name}"
  short_name = "%{value_short_name}"
}

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

  idle_timeout = "600s"
  running_timeout = "21600s"

  replica_zones = ["us-central1-a", "us-central1-b"]
  annotations = {
    label-one = "value-one"
  }

  labels = {
    "label" = "key"
  }

  host {
    gce_instance {
      machine_type                = "e2-standard-4"
      boot_disk_size_gb           = 35
      disable_public_ip_addresses = true
      disable_ssh                 = false
      vm_tags = {
        "tagKeys/${google_tags_tag_key.tag_key1.name}" = "tagValues/${google_tags_tag_value.tag_value1.name}"
      }
    }
  }
  allowed_ports = [
    {
      first: 22
      last:22
    },
    {
      first: 80
      last: 80
    },
    {
      first: 8000
      last: 9000
    }
  ]
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_workstations_workstation_config_iam_policy" "foo" {
  provider = google-beta
  project = google_workstations_workstation_config.default.project
  location = google_workstations_workstation_config.default.location
  workstation_cluster_id = google_workstations_workstation_config.default.workstation_cluster_id
  workstation_config_id = google_workstations_workstation_config.default.workstation_config_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccWorkstationsWorkstationConfigIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_tags_tag_key" "tag_key1" {
  provider   = google-beta
  parent     = "organizations/%{org_id}"
  short_name = "%{key_short_name}"
}

resource "google_tags_tag_value" "tag_value1" {
  provider   = google-beta
  parent     = "tagKeys/${google_tags_tag_key.tag_key1.name}"
  short_name = "%{value_short_name}"
}

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

  idle_timeout = "600s"
  running_timeout = "21600s"

  replica_zones = ["us-central1-a", "us-central1-b"]
  annotations = {
    label-one = "value-one"
  }

  labels = {
    "label" = "key"
  }

  host {
    gce_instance {
      machine_type                = "e2-standard-4"
      boot_disk_size_gb           = 35
      disable_public_ip_addresses = true
      disable_ssh                 = false
      vm_tags = {
        "tagKeys/${google_tags_tag_key.tag_key1.name}" = "tagValues/${google_tags_tag_value.tag_value1.name}"
      }
    }
  }
  allowed_ports = [
    {
      first: 22
      last:22
    },
    {
      first: 80
      last: 80
    },
    {
      first: 8000
      last: 9000
    }
  ]
}

resource "google_workstations_workstation_config_iam_binding" "foo" {
  provider = google-beta
  project = google_workstations_workstation_config.default.project
  location = google_workstations_workstation_config.default.location
  workstation_cluster_id = google_workstations_workstation_config.default.workstation_cluster_id
  workstation_config_id = google_workstations_workstation_config.default.workstation_config_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccWorkstationsWorkstationConfigIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_tags_tag_key" "tag_key1" {
  provider   = google-beta
  parent     = "organizations/%{org_id}"
  short_name = "%{key_short_name}"
}

resource "google_tags_tag_value" "tag_value1" {
  provider   = google-beta
  parent     = "tagKeys/${google_tags_tag_key.tag_key1.name}"
  short_name = "%{value_short_name}"
}

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

  idle_timeout = "600s"
  running_timeout = "21600s"

  replica_zones = ["us-central1-a", "us-central1-b"]
  annotations = {
    label-one = "value-one"
  }

  labels = {
    "label" = "key"
  }

  host {
    gce_instance {
      machine_type                = "e2-standard-4"
      boot_disk_size_gb           = 35
      disable_public_ip_addresses = true
      disable_ssh                 = false
      vm_tags = {
        "tagKeys/${google_tags_tag_key.tag_key1.name}" = "tagValues/${google_tags_tag_value.tag_value1.name}"
      }
    }
  }
  allowed_ports = [
    {
      first: 22
      last:22
    },
    {
      first: 80
      last: 80
    },
    {
      first: 8000
      last: 9000
    }
  ]
}

resource "google_workstations_workstation_config_iam_binding" "foo" {
  provider = google-beta
  project = google_workstations_workstation_config.default.project
  location = google_workstations_workstation_config.default.location
  workstation_cluster_id = google_workstations_workstation_config.default.workstation_cluster_id
  workstation_config_id = google_workstations_workstation_config.default.workstation_config_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
