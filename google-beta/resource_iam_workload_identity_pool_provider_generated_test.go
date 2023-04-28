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

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccIAMBetaWorkloadIdentityPoolProvider_iamWorkloadIdentityPoolProviderAwsBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAMBetaWorkloadIdentityPoolProviderDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_iamWorkloadIdentityPoolProviderAwsBasicExample(context),
			},
			{
				ResourceName:            "google_iam_workload_identity_pool_provider.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workload_identity_pool_id", "workload_identity_pool_provider_id"},
			},
		},
	})
}

func testAccIAMBetaWorkloadIdentityPoolProvider_iamWorkloadIdentityPoolProviderAwsBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workload_identity_pool" "pool" {
  workload_identity_pool_id = "tf-test-example-pool%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "example" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "tf-test-example-prvdr%{random_suffix}"
  aws {
    account_id = "999999999999"
  }
}
`, context)
}

func TestAccIAMBetaWorkloadIdentityPoolProvider_iamWorkloadIdentityPoolProviderAwsFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAMBetaWorkloadIdentityPoolProviderDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_iamWorkloadIdentityPoolProviderAwsFullExample(context),
			},
			{
				ResourceName:            "google_iam_workload_identity_pool_provider.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workload_identity_pool_id", "workload_identity_pool_provider_id"},
			},
		},
	})
}

func testAccIAMBetaWorkloadIdentityPoolProvider_iamWorkloadIdentityPoolProviderAwsFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workload_identity_pool" "pool" {
  workload_identity_pool_id = "tf-test-example-pool%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "example" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "tf-test-example-prvdr%{random_suffix}"
  display_name                       = "Name of provider"
  description                        = "AWS identity pool provider for automated test"
  disabled                           = true
  attribute_condition                = "attribute.aws_role==\"arn:aws:sts::999999999999:assumed-role/stack-eu-central-1-lambdaRole\""
  attribute_mapping                  = {
    "google.subject"        = "assertion.arn"
    "attribute.aws_account" = "assertion.account"
    "attribute.environment" = "assertion.arn.contains(\":instance-profile/Production\") ? \"prod\" : \"test\""
  }
  aws {
    account_id = "999999999999"
  }
}
`, context)
}

func TestAccIAMBetaWorkloadIdentityPoolProvider_iamWorkloadIdentityPoolProviderOidcBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAMBetaWorkloadIdentityPoolProviderDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_iamWorkloadIdentityPoolProviderOidcBasicExample(context),
			},
			{
				ResourceName:            "google_iam_workload_identity_pool_provider.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workload_identity_pool_id", "workload_identity_pool_provider_id"},
			},
		},
	})
}

func testAccIAMBetaWorkloadIdentityPoolProvider_iamWorkloadIdentityPoolProviderOidcBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workload_identity_pool" "pool" {
  workload_identity_pool_id = "tf-test-example-pool%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "example" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "tf-test-example-prvdr%{random_suffix}"
  attribute_mapping                  = {
    "google.subject" = "assertion.sub"
  }
  oidc {
    issuer_uri        = "https://sts.windows.net/azure-tenant-id"
  }
}
`, context)
}

func TestAccIAMBetaWorkloadIdentityPoolProvider_iamWorkloadIdentityPoolProviderOidcFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAMBetaWorkloadIdentityPoolProviderDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPoolProvider_iamWorkloadIdentityPoolProviderOidcFullExample(context),
			},
			{
				ResourceName:            "google_iam_workload_identity_pool_provider.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workload_identity_pool_id", "workload_identity_pool_provider_id"},
			},
		},
	})
}

func testAccIAMBetaWorkloadIdentityPoolProvider_iamWorkloadIdentityPoolProviderOidcFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workload_identity_pool" "pool" {
  workload_identity_pool_id = "tf-test-example-pool%{random_suffix}"
}

resource "google_iam_workload_identity_pool_provider" "example" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "tf-test-example-prvdr%{random_suffix}"
  display_name                       = "Name of provider"
  description                        = "OIDC identity pool provider for automated test"
  disabled                           = true
  attribute_condition                = "\"e968c2ef-047c-498d-8d79-16ca1b61e77e\" in assertion.groups"
  attribute_mapping                  = {
    "google.subject"                  = "\"azure::\" + assertion.tid + \"::\" + assertion.sub"
    "attribute.tid"                   = "assertion.tid"
    "attribute.managed_identity_name" = <<EOT
      {
        "8bb39bdb-1cc5-4447-b7db-a19e920eb111":"workload1",
        "55d36609-9bcf-48e0-a366-a3cf19027d2a":"workload2"
      }[assertion.oid]
EOT
  }
  oidc {
    allowed_audiences = ["https://example.com/gcp-oidc-federation", "example.com/gcp-oidc-federation"]
    issuer_uri        = "https://sts.windows.net/azure-tenant-id"
  }
}
`, context)
}

func testAccCheckIAMBetaWorkloadIdentityPoolProviderDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_iam_workload_identity_pool_provider" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{IAMBetaBasePath}}projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}/providers/{{workload_identity_pool_provider_id}}")
			if err != nil {
				return err
			}

			res, err := transport_tpg.SendRequest(config, "GET", "", url, config.UserAgent, nil)
			if err != nil {
				return nil
			}

			if v := res["state"]; v == "DELETED" {
				return nil
			}

			return fmt.Errorf("IAMBetaWorkloadIdentityPoolProvider still exists at %s", url)
		}

		return nil
	}
}
