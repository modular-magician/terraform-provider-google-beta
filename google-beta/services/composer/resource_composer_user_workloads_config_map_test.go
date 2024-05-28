// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package composer_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/composer"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testComposerUserWorkloadsConfigMapPrefix = "tf-test-composer-config-map"

func TestAccComposerUserWorkloadsConfigMap_basic(t *testing.T) {
	t.Parallel()

	envName := fmt.Sprintf("%s-%d", testComposerEnvironmentPrefix, acctest.RandInt(t))
	configMapName := fmt.Sprintf("%s-%d", testComposerUserWorkloadsConfigMapPrefix, acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccComposerEnvironmentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComposerUserWorkloadsConfigMap_basic(envName, configMapName, envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("google_composer_user_workloads_config_map.test", "data.api_host", "apihost:443"),
				),
			},
			{
				ResourceName: "google_composer_user_workloads_config_map.test",
				ImportState:  true,
			},
		},
	})
}

func TestAccComposerUserWorkloadsConfigMap_update(t *testing.T) {
	t.Parallel()

	envName := fmt.Sprintf("%s-%d", testComposerEnvironmentPrefix, acctest.RandInt(t))
	configMapName := fmt.Sprintf("%s-%d", testComposerUserWorkloadsConfigMapPrefix, acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccComposerEnvironmentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComposerUserWorkloadsConfigMap_basic(envName, configMapName, envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv()),
			},
			{
				Config: testAccComposerUserWorkloadsConfigMap_update(envName, configMapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("google_composer_user_workloads_config_map.test", "data.db_host", "dbhost:5432"),
					resource.TestCheckNoResourceAttr("google_composer_user_workloads_config_map.test", "data.api_host"),
				),
			},
		},
	})
}

func TestAccComposerUserWorkloadsConfigMap_delete(t *testing.T) {
	t.Parallel()

	envName := fmt.Sprintf("%s-%d", testComposerEnvironmentPrefix, acctest.RandInt(t))
	configMapName := fmt.Sprintf("%s-%d", testComposerUserWorkloadsConfigMapPrefix, acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccComposerEnvironmentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComposerUserWorkloadsConfigMap_basic(envName, configMapName, envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv()),
			},
			{
				Config: testAccComposerUserWorkloadsConfigMap_delete(envName),
				Check: resource.ComposeTestCheckFunc(
					testAccComposerUserWorkloadsConfigMapDestroyed(t),
				),
			},
		},
	})
}

func testAccComposerUserWorkloadsConfigMap_basic(envName, configMapName, project, region string) string {
	return fmt.Sprintf(`
resource "google_composer_environment" "test" {
  name   = "%s"
  config {
    software_config {
      image_version = "composer-3-airflow-2"
    }
  }
}
resource "google_composer_user_workloads_config_map" "test" {
  environment = google_composer_environment.test.name
  name = "%s"
  project = "%s"
  region = "%s"
  data = {
    api_host: "apihost:443",
  }
}
`, envName, configMapName, project, region)
}

func testAccComposerUserWorkloadsConfigMap_update(envName, configMapName string) string {
	return fmt.Sprintf(`
resource "google_composer_environment" "test" {
  name   = "%s"
  config {
    software_config {
      image_version = "composer-3-airflow-2"
    }
  }
}
resource "google_composer_user_workloads_config_map" "test" {
  environment = google_composer_environment.test.name
  name = "%s"
  data = {
		db_host: "dbhost:5432",
  }
}
`, envName, configMapName)
}

func testAccComposerUserWorkloadsConfigMap_delete(envName string) string {
	return fmt.Sprintf(`
resource "google_composer_environment" "test" {
  name   = "%s"
  config {
    software_config {
      image_version = "composer-3-airflow-2"
    }
  }
}
`, envName)
}

func testAccComposerUserWorkloadsConfigMapDestroyed(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		config := acctest.GoogleProviderConfig(t)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "google_composer_user_workloads_config_map" {
				continue
			}

			idTokens := strings.Split(rs.Primary.ID, "/")
			if len(idTokens) != 8 {
				return fmt.Errorf("Invalid ID %q, expected format projects/{project}/regions/{region}/environments/{environment}/userWorkloadsConfigMaps/{name}", rs.Primary.ID)
			}
			configMapName := &composer.UserWorkloadsConfigMapName{
				Project:     idTokens[1],
				Region:      idTokens[3],
				Environment: idTokens[5],
				ConfigMap:   idTokens[7],
			}

			_, err := config.NewComposerClient(config.UserAgent).Projects.Locations.Environments.UserWorkloadsConfigMaps.Get(configMapName.ResourceName()).Do()
			if err == nil {
				return fmt.Errorf("config map %s still exists", configMapName.ResourceName())
			}
		}

		return nil
	}
}
