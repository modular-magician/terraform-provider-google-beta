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

func TestAccFirebaseDatabaseInstance_firebaseDatabaseInstanceExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    getTestProjectFromEnv(),
		"region":        getTestRegionFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckFirebaseDatabaseInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseDatabaseInstance_firebaseDatabaseInstanceExample(context),
			},
			{
				ResourceName:            "google_firebase_database_instance.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "instance_id"},
			},
		},
	})
}

func testAccFirebaseDatabaseInstance_firebaseDatabaseInstanceExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_firebase_database_instance" "basic" {
  provider = google-beta
  project  = "%{project_id}"
  region   = "%{region}"
  instance_id = "tf-test-active-db%{random_suffix}"
}
`, context)
}

func TestAccFirebaseDatabaseInstance_firebaseDatabaseInstanceDisabledExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    getTestProjectFromEnv(),
		"region":        getTestRegionFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckFirebaseDatabaseInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseDatabaseInstance_firebaseDatabaseInstanceDisabledExample(context),
			},
			{
				ResourceName:            "google_firebase_database_instance.full",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "instance_id", "desired_state"},
			},
		},
	})
}

func testAccFirebaseDatabaseInstance_firebaseDatabaseInstanceDisabledExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_firebase_database_instance" "full" {
  provider = google-beta
  project  = "%{project_id}"
  region   = "%{region}"
  instance_id = "tf-test-disabled-db%{random_suffix}"
  type     = "USER_DATABASE"
  desired_state   = "DISABLED"
}
`, context)
}

func TestAccFirebaseDatabaseInstance_firebaseDatabaseInstanceDefaultExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"region":        getTestRegionFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckFirebaseDatabaseInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseDatabaseInstance_firebaseDatabaseInstanceDefaultExample(context),
			},
			{
				ResourceName:            "google_firebase_database_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "instance_id"},
			},
		},
	})
}

func testAccFirebaseDatabaseInstance_firebaseDatabaseInstanceDefaultExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "default" {
  provider = google-beta
  project_id = "tf-test-rtdb-project%{random_suffix}"
  name       = "tf-test-rtdb-project%{random_suffix}"
  org_id     = "%{org_id}"
  labels     = {
    "firebase" = "enabled"
  }
}

resource "google_firebase_project" "default" {
  provider = google-beta
  project  = google_project.default.project_id
}

resource "google_project_service" "firebase_database" {
  provider = google-beta
  project  = google_firebase_project.default.project
  service  = "firebasedatabase.googleapis.com"
}

resource "google_firebase_database_instance" "default" {
  provider = google-beta
  project  = google_firebase_project.default.project
  region   = "%{region}"
  instance_id = "tf-test-rtdb-project%{random_suffix}-default-rtdb"
  type     = "DEFAULT_DATABASE"
}
`, context)
}

func testAccCheckFirebaseDatabaseInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_firebase_database_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{FirebaseDatabaseBasePath}}projects/{{project}}/locations/{{region}}/instances/{{instance_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			res, err := sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err != nil {
				return err // RTDB only supports soft-delete.
			}

			dbState := res["state"]
			if dbState == "DELETED" {
				return nil // USER_DATABASE soft deleted.
			}
			dbType := res["type"]
			if dbState == "DISABLED" && dbType == "DEFAULT_DATABASE" {
				return nil // DEFAULT_DATABASE is left in a DISABLED state because it cannot be deleted.
			}

			return fmt.Errorf("firebase_database_instance %s got state=%s, want DELETED", url, s)
		}

		return nil
	}
}
