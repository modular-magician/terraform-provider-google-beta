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
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccGKEHubMembership_gkehubMembershipBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGKEHubMembershipDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEHubMembership_gkehubMembershipBasicExample(context),
			},
			{
				ResourceName:            "google_gke_hub_membership.membership",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"membership_id"},
			},
		},
	})
}

func testAccGKEHubMembership_gkehubMembershipBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "basiccluster%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
}

resource "google_gke_hub_membership" "membership" {
  membership_id = "basic%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }
}
`, context)
}

func TestAccGKEHubMembership_gkehubMembershipIssuerExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       acctest.GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGKEHubMembershipDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEHubMembership_gkehubMembershipIssuerExample(context),
			},
			{
				ResourceName:            "google_gke_hub_membership.membership",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"membership_id"},
			},
		},
	})
}

func testAccGKEHubMembership_gkehubMembershipIssuerExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "basiccluster%{random_suffix}"
  location           = "us-central1-a"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "%{project}.svc.id.goog"
  }
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

			config := GoogleProviderConfig(t)

			url, err := acctest.ReplaceVarsForTest(config, rs, "{{GKEHubBasePath}}projects/{{project}}/locations/global/memberships/{{membership_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("GKEHubMembership still exists at %s", url)
			}
		}

		return nil
	}
}
