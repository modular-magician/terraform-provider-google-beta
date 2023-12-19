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

package gkehub_test

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

func TestAccGKEHubMembership_gkehubMembershipRegionalExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":         envvar.GetTestProjectFromEnv(),
		"location":        envvar.GetTestRegionFromEnv(),
		"network_name":    acctest.BootstrapSharedTestNetwork(t, "gke-cluster"),
		"subnetwork_name": acctest.BootstrapSubnet(t, "gke-cluster", acctest.BootstrapSharedTestNetwork(t, "gke-cluster")),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGKEHubMembershipDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEHubMembership_gkehubMembershipRegionalExample(context),
			},
			{
				ResourceName:            "google_gke_hub_membership.membership",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"membership_id", "location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccGKEHubMembership_gkehubMembershipRegionalExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-basic-cluster%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
  deletion_protection = false
  network       = "%{network_name}"
  subnetwork    = "%{subnetwork_name}"
}

resource "google_gke_hub_membership" "membership" {
  membership_id = "basic%{random_suffix}"
  location = "%{location}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }
}
`, context)
}

func TestAccGKEHubMembership_gkehubMembershipBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"network_name":        acctest.BootstrapSharedTestNetwork(t, "gke-cluster"),
		"subnetwork_name":     acctest.BootstrapSubnet(t, "gke-cluster", acctest.BootstrapSharedTestNetwork(t, "gke-cluster")),
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGKEHubMembershipDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEHubMembership_gkehubMembershipBasicExample(context),
			},
			{
				ResourceName:            "google_gke_hub_membership.membership",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"membership_id", "location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccGKEHubMembership_gkehubMembershipBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-basic-cluster%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
  deletion_protection  = "%{deletion_protection}"
  network       = "%{network_name}"
  subnetwork    = "%{subnetwork_name}"
}

resource "google_gke_hub_membership" "membership" {
  membership_id = "basic%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }

  labels = {
    env = "test"
  }
}
`, context)
}

func TestAccGKEHubMembership_gkehubMembershipIssuerExample(t *testing.T) {
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
		CheckDestroy:             testAccCheckGKEHubMembershipDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEHubMembership_gkehubMembershipIssuerExample(context),
			},
			{
				ResourceName:            "google_gke_hub_membership.membership",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"membership_id", "location", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccGKEHubMembership_gkehubMembershipIssuerExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-basic-cluster%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "%{project}.svc.id.goog"
  }
  deletion_protection  = "%{deletion_protection}"
  network       = "%{network_name}"
  subnetwork    = "%{subnetwork_name}"
}

resource "google_gke_hub_membership" "membership" {
  membership_id = "basic%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = google_container_cluster.primary.id
    }
  }
  authority {
    issuer = "https://container.googleapis.com/v1/${google_container_cluster.primary.id}"
  }
}
`, context)
}

func testAccCheckGKEHubMembershipDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_gke_hub_membership" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{GKEHubBasePath}}projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}")
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
				return fmt.Errorf("GKEHubMembership still exists at %s", url)
			}
		}

		return nil
	}
}
