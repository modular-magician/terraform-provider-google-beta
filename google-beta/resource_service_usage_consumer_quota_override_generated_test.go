// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccServiceUsageConsumerQuotaOverride_consumerQuotaOverrideExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckServiceUsageConsumerQuotaOverrideDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceUsageConsumerQuotaOverride_consumerQuotaOverrideExample(context),
			},
		},
	})
}

func testAccServiceUsageConsumerQuotaOverride_consumerQuotaOverrideExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "my_project" {
  provider   = google-beta
  name       = "tf-test-project"
  project_id = "quota%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "google_service_usage_consumer_quota_override" "override" {
  provider       = google-beta
  project        = google_project.my_project.project_id
  service        = "servicemanagement.googleapis.com"
  metric         = "servicemanagement.googleapis.com%2Fdefault_requests"
  limit          = "%2Fmin%2Fproject"
  override_value = "95"
  force          = true
}
`, context)
}

func testAccCheckServiceUsageConsumerQuotaOverrideDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_service_usage_consumer_quota_override" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ServiceUsageBasePath}}projects/{{project}}/services/{{service}}/consumerQuotaMetrics/{{metric}}/limits/{{limit}}/consumerOverrides/")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("ServiceUsageConsumerQuotaOverride still exists at %s", url)
			}
		}

		return nil
	}
}
