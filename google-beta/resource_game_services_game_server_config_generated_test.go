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

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccGameServicesGameServerConfig_gameServiceConfigBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGameServicesGameServerConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGameServicesGameServerConfig_gameServiceConfigBasicExample(context),
			},
			{
				ResourceName:            "google_game_services_game_server_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"config_id", "location", "deployment_id"},
			},
		},
	})
}

func testAccGameServicesGameServerConfig_gameServiceConfigBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_game_services_game_server_deployment" "default" {
  deployment_id  = "tf-test-tf-test-deployment%{random_suffix}"
  description = "a deployment description"
}

resource "google_game_services_game_server_config" "default" {
  config_id     = "tf-test-tf-test-config%{random_suffix}"
  deployment_id = google_game_services_game_server_deployment.default.deployment_id
  description   = "a config description"

  fleet_configs {
    name       = "something-unique"
    fleet_spec = jsonencode({ "replicas" : 1, "scheduling" : "Packed", "template" : { "metadata" : { "name" : "tf-test-game-server-template" }, "spec" : { "ports": [{"name": "default", "portPolicy": "Dynamic", "containerPort": 7654, "protocol": "UDP"}], "template" : { "spec" : { "containers" : [{ "name" : "simple-udp-server", "image" : "gcr.io/agones-images/udp-server:0.14" }] } } } } })
  }

  scaling_configs {
    name = "scaling-config-name"
    fleet_autoscaler_spec = jsonencode({"policy": {"type": "Webhook","webhook": {"service": {"name": "autoscaler-webhook-service","namespace": "default","path": "scale"}}}})
    selectors {
      labels = {
        "one" : "two"
      }
    }

    schedules {
      cron_job_duration = "3.500s"
      cron_spec         = "0 0 * * 0"
    }
  }
}
`, context)
}

func testAccCheckGameServicesGameServerConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_game_services_game_server_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{GameServicesBasePath}}projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}/configs/{{config_id}}")
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
				return fmt.Errorf("GameServicesGameServerConfig still exists at %s", url)
			}
		}

		return nil
	}
}
