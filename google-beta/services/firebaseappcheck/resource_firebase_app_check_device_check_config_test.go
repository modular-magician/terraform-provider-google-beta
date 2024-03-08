// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package firebaseappcheck_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccFirebaseAppCheckDeviceCheckConfig_firebaseAppCheckDeviceCheckConfigUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":       envvar.GetTestProjectFromEnv(),
		"team_id":          "9987654321",
		"private_key_path": "test-fixtures/private-key.p8",
		"token_ttl":        "3900s",
		"random_suffix":    acctest.RandString(t, 10),
	}

	contextUpdated := map[string]interface{}{
		"project_id":       envvar.GetTestProjectFromEnv(),
		"team_id":          "9987654321",
		"private_key_path": "test-fixtures/private-key-2.p8",
		"token_ttl":        "7200s",
		// Bundle ID needs to be the same between updates but different between tests
		"random_suffix": context["random_suffix"],
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAppCheckDeviceCheckConfig_firebaseAppCheckDeviceCheckConfigFullExample(context),
			},
			{
				ResourceName:            "google_firebase_app_check_device_check_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"private_key", "app_id"},
			},
			{
				Config: testAccFirebaseAppCheckDeviceCheckConfig_firebaseAppCheckDeviceCheckConfigFullExample(contextUpdated),
			},
			{
				ResourceName:            "google_firebase_app_check_device_check_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"private_key", "app_id"},
			},
		},
	})
}
