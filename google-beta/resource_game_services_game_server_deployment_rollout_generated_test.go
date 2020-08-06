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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGameServicesGameServerDeploymentRollout_gameServiceDeploymentRolloutBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckGameServicesGameServerDeploymentRolloutDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGameServicesGameServerDeploymentRollout_gameServiceDeploymentRolloutBasicExample(context),
			},
		},
	})
}

func testAccGameServicesGameServerDeploymentRollout_gameServiceDeploymentRolloutBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_game_services_game_server_deployment" "default" {
  provider = google-beta

  deployment_id  = "tf-test-tf-test-deployment%{random_suffix}"
  description = "a deployment description"
}

resource "google_game_services_game_server_config" "default" {
  provider = google-beta

  config_id     = "tf-test-tf-test-config%{random_suffix}"
  deployment_id = google_game_services_game_server_deployment.default.deployment_id
  description   = "a config description"

  fleet_configs {
    name       = "some-non-guid"
    fleet_spec = jsonencode({ "replicas" : 1, "scheduling" : "Packed", "template" : { "metadata" : { "name" : "tf-test-game-server-template" }, "spec" : { "template" : { "spec" : { "containers" : [{ "name" : "simple-udp-server", "image" : "gcr.io/agones-images/udp-server:0.14" }] } } } } })

    // Alternate usage:
    // fleet_spec = file(fleet_configs.json)
  }
}

resource "google_game_services_game_server_deployment_rollout" "default" {
  provider = google-beta

  deployment_id              = google_game_services_game_server_deployment.default.deployment_id
  default_game_server_config = google_game_services_game_server_config.default.name
}
`, context)
}

func testAccCheckGameServicesGameServerDeploymentRolloutDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_game_services_game_server_deployment_rollout" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{GameServicesBasePath}}projects/{{project}}/locations/global/gameServerDeployments/{{deployment_id}}/rollout")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("GameServicesGameServerDeploymentRollout still exists at %s", url)
			}
		}

		return nil
	}
}
