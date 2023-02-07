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

func TestAccApiGatewayGateway_apigatewayGatewayBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckApiGatewayGatewayDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayGateway_apigatewayGatewayBasicExample(context),
			},
			{
				ResourceName:            "google_api_gateway_gateway.api_gw",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "gateway_id"},
			},
		},
	})
}

func testAccApiGatewayGateway_apigatewayGatewayBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_gw" {
  provider = google-beta
  api_id = "tf-test-api-%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_gw" {
  provider = google-beta
  api = google_api_gateway_api.api_gw.api_id
  api_config_id = "tf-test-config-%{random_suffix}"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "google_api_gateway_gateway" "api_gw" {
  provider = google-beta
  api_config = google_api_gateway_api_config.api_gw.id
  gateway_id = "tf-test-gateway-%{random_suffix}"
}
`, context)
}

func TestAccApiGatewayGateway_apigatewayGatewayFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckApiGatewayGatewayDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayGateway_apigatewayGatewayFullExample(context),
			},
			{
				ResourceName:            "google_api_gateway_gateway.api_gw",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "gateway_id"},
			},
		},
	})
}

func testAccApiGatewayGateway_apigatewayGatewayFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_gw" {
  provider = google-beta
  api_id = "tf-test-api-%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_gw" {
  provider = google-beta
  api = google_api_gateway_api.api_gw.api_id
  api_config_id = "tf-test-config-%{random_suffix}"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
}

resource "google_api_gateway_gateway" "api_gw" {
  provider = google-beta
  region     = "us-central1"
  api_config = google_api_gateway_api_config.api_gw.id
  gateway_id = "tf-test-gateway-%{random_suffix}"
  display_name = "MM Dev API Gateway"
  labels = {
    environment = "dev"
  }
}
`, context)
}

func testAccCheckApiGatewayGatewayDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_api_gateway_gateway" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ApiGatewayBasePath}}projects/{{project}}/locations/{{region}}/gateways/{{gateway_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("ApiGatewayGateway still exists at %s", url)
			}
		}

		return nil
	}
}
