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

package securityposture_test

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

func TestAccSecurityposturePostureDeployment_securityposturePostureDeploymentBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":         envvar.GetTestOrgFromEnv(t),
		"project_number": envvar.GetTestProjectNumberFromEnv(),
		"random_suffix":  acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecurityposturePostureDeploymentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityposturePostureDeployment_securityposturePostureDeploymentBasicExample(context),
			},
			{
				ResourceName:            "google_securityposture_posture_deployment.postureDeployment",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent", "location", "posture_deployment_id"},
			},
		},
	})
}

func testAccSecurityposturePostureDeployment_securityposturePostureDeploymentBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_securityposture_posture" "posture_1" {
    posture_id  = "tf_test_posture_1%{random_suffix}"
    parent      = "organizations/%{org_id}"
    location    = "global"
    state       = "ACTIVE"
    description = "a new posture"
    policy_sets {
        policy_set_id = "org_policy_set"
        description   = "set of org policies"
        policies {
            policy_id = "policy_1"
            constraint {
                org_policy_constraint {
                    canned_constraint_id = "storage.uniformBucketLevelAccess"
                    policy_rules {
                        enforce = true
                    }
                }
            }
        }
    }
}

resource "google_securityposture_posture_deployment" "postureDeployment" {
    posture_deployment_id = "tf_test_posture_deployment_1%{random_suffix}"
    parent                = "organizations/%{org_id}"
    location              = "global"
    description           = "a new posture deployment"
    target_resource       = "projects/%{project_number}"
    posture_id            = google_securityposture_posture.posture_1.name
    posture_revision_id   = google_securityposture_posture.posture_1.revision_id
}
`, context)
}

func testAccCheckSecurityposturePostureDeploymentDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_securityposture_posture_deployment" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{SecuritypostureBasePath}}{{parent}}/locations/{{location}}/postureDeployments/{{posture_deployment_id}}")
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
				return fmt.Errorf("SecurityposturePostureDeployment still exists at %s", url)
			}
		}

		return nil
	}
}
