// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package cloudrun_test

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudRunService_cloudRunServiceUpdate(t *testing.T) {
	t.Parallel()

	project := envvar.GetTestProjectFromEnv()
	name := "tftest-cloudrun-" + acctest.RandString(t, 6)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceUpdate(name, project, "10", "600"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_cloudRunServiceUpdate(name, project, "50", "300"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
		},
	})
}

// test that the status fields are propagated correctly
func TestAccCloudRunService_cloudRunServiceCreateHasStatus(t *testing.T) {
	t.Parallel()

	project := envvar.GetTestProjectFromEnv()
	name := "tftest-cloudrun-" + acctest.RandString(t, 6)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceUpdate(name, project, "10", "600"),
				Check:  resource.TestCheckResourceAttrSet("google_cloud_run_service.default", "status.0.traffic.0.revision_name"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels"},
			},
		},
	})
}

// this test checks that Terraform does not fail with a 409 recreating the same service
func TestAccCloudRunService_foregroundDeletion(t *testing.T) {
	t.Parallel()

	project := envvar.GetTestProjectFromEnv()
	name := "tftest-cloudrun-" + acctest.RandString(t, 6)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceUpdate(name, project, "10", "600"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
			{
				Config: " ", // very explicitly add a space, as the test runner fails if this is just ""
			},
			{
				Config: testAccCloudRunService_cloudRunServiceUpdate(name, project, "10", "600"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceUpdate(name, project, concurrency, timeoutSeconds string) string {
	return fmt.Sprintf(`
resource "google_cloud_run_service" "default" {
  name     = "%s"
  location = "us-central1"

  metadata {
    namespace = "%s"
    annotations = {
      generated-by = "magic-modules"
    }
    labels = {
      env                   = "foo"
      default_expiration_ms = 3600000
    }
  }

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        ports {
          container_port = 8080
        }
      }
      container_concurrency = %s
      timeout_seconds = %s
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
    tag             = "magic-module"
  }

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }
}
`, name, project, concurrency, timeoutSeconds)
}

func TestAccCloudRunService_secretVolume(t *testing.T) {
	t.Parallel()

	project := envvar.GetTestProjectFromEnv()
	name := "tftest-cloudrun-" + acctest.RandString(t, 6)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceUpdateWithSecretVolume(name, project, "secret-"+acctest.RandString(t, 5), "secret-"+acctest.RandString(t, 6), "google_secret_manager_secret.secret1.secret_id"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_cloudRunServiceUpdateWithSecretVolume(name, project, "secret-"+acctest.RandString(t, 10), "secret-"+acctest.RandString(t, 11), "google_secret_manager_secret.secret2.secret_id"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceUpdateWithSecretVolume(name, project, secretName1, secretName2, secretRef string) string {
	return fmt.Sprintf(`
data "google_project" "project" {
}

resource "google_secret_manager_secret" "secret1" {
  secret_id = "%s"
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret" "secret2" {
  secret_id = "%s"
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "secret1-version-data" {
  secret = google_secret_manager_secret.secret1.name
  secret_data = "secret-data1"
}

resource "google_secret_manager_secret_version" "secret2-version-data" {
  secret = google_secret_manager_secret.secret2.name
  secret_data = "secret-data2"
}

resource "google_secret_manager_secret_iam_member" "secret1-access" {
  secret_id = google_secret_manager_secret.secret1.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  depends_on = [google_secret_manager_secret.secret1]
}

resource "google_secret_manager_secret_iam_member" "secret2-access" {
  secret_id = google_secret_manager_secret.secret2.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  depends_on = [google_secret_manager_secret.secret2]
}

resource "google_cloud_run_service" "default" {
  name     = "%s"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        volume_mounts {
          name = "a-volume"
          mount_path = "/secrets"
        }
      }
      volumes {
        name = "a-volume"
        secret {
          secret_name = %s
          items {
            key = "1"
            path = "my-secret"
          }
        }
      }
    }
  }

  metadata {
    namespace = "%s"
    annotations = {
      generated-by = "magic-modules"
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }

  depends_on = [google_secret_manager_secret_version.secret1-version-data, google_secret_manager_secret_version.secret2-version-data]
}
`, secretName1, secretName2, name, secretRef, project)
}

func TestAccCloudRunService_secretEnvironmentVariable(t *testing.T) {
	t.Parallel()

	project := envvar.GetTestProjectFromEnv()
	name := "tftest-cloudrun-" + acctest.RandString(t, 6)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceUpdateWithSecretEnvVar(name, project, "secret-"+acctest.RandString(t, 5), "secret-"+acctest.RandString(t, 6), "google_secret_manager_secret.secret1.secret_id"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_cloudRunServiceUpdateWithSecretEnvVar(name, project, "secret-"+acctest.RandString(t, 10), "secret-"+acctest.RandString(t, 11), "google_secret_manager_secret.secret2.secret_id"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceUpdateWithSecretEnvVar(name, project, secretName1, secretName2, secretRef string) string {
	return fmt.Sprintf(`
data "google_project" "project" {
}

resource "google_secret_manager_secret" "secret1" {
  secret_id = "%s"
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret" "secret2" {
  secret_id = "%s"
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "secret1-version-data" {
  secret = google_secret_manager_secret.secret1.name
  secret_data = "secret-data1"
}

resource "google_secret_manager_secret_version" "secret2-version-data" {
  secret = google_secret_manager_secret.secret2.name
  secret_data = "secret-data2"
}

resource "google_secret_manager_secret_iam_member" "secret1-access" {
  secret_id = google_secret_manager_secret.secret1.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  depends_on = [google_secret_manager_secret.secret1]
}

resource "google_secret_manager_secret_iam_member" "secret2-access" {
  secret_id = google_secret_manager_secret.secret2.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  depends_on = [google_secret_manager_secret.secret2]
}

resource "google_cloud_run_service" "default" {
  name     = "%s"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        env {
          name = "SECRET_ENV_VAR"
          value_from {
            secret_key_ref {
              name = %s
              key = "1"
            }
          }
        }
      }
    }
  }

  metadata {
    namespace = "%s"
    annotations = {
      generated-by = "magic-modules"
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }

  depends_on = [google_secret_manager_secret_version.secret1-version-data, google_secret_manager_secret_version.secret2-version-data]
}
`, secretName1, secretName2, name, secretRef, project)
}

func TestAccCloudRunService_withProviderDefaultLabels(t *testing.T) {
	// The test failed if VCR testing is enabled, because the cached provider config is used.
	// With the cached provider config, any changes in the provider default labels will not be applied.
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_withProviderDefaultLabels(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.%", "2"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.env", "foo"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.default_expiration_ms", "3600000"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.%", "3"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.default_key1", "default_value1"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.env", "foo"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.default_expiration_ms", "3600000"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.%", "4"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.annotations.%", "1"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.annotations.generated-by", "magic-modules"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_annotations.%", "6"),
				),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_resourceLabelsOverridesProviderDefaultLabels(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.%", "3"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.env", "foo"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.default_expiration_ms", "3600000"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.default_key1", "value1"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.%", "3"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.default_key1", "value1"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.env", "foo"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.default_expiration_ms", "3600000"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.%", "4"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.annotations.%", "1"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.annotations.generated-by", "magic-modules-update"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_annotations.%", "6"),
				),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_moveResourceLabelToProviderDefaultLabels(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.%", "2"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.default_expiration_ms", "3600000"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.default_key1", "value1"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.%", "3"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.default_key1", "value1"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.env", "foo"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.default_expiration_ms", "3600000"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.%", "4"),
				),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_resourceLabelsOverridesProviderDefaultLabels(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.%", "3"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.env", "foo"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.default_expiration_ms", "3600000"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.default_key1", "value1"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.%", "3"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.default_key1", "value1"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.env", "foo"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.default_expiration_ms", "3600000"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.%", "4"),
				),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_cloudRunServiceBasic(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("google_cloud_run_service.default", "metadata.0.labels.%"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.%", "1"),

					resource.TestCheckNoResourceAttr("google_cloud_run_service.default", "metadata.0.annotations.%"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_annotations.%", "5"),
				),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
		},
	})
}

func TestAccCloudRunServiceMigration_withLabels(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	name := "tftest-cloudrun-" + acctest.RandString(t, 6)
	project := envvar.GetTestProjectFromEnv()
	oldVersion := map[string]resource.ExternalProvider{
		"google": {
			VersionConstraint: "4.83.0", // a version that doesn't separate user defined labels and system labels
			Source:            "registry.terraform.io/hashicorp/google",
		},
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck: func() { acctest.AccTestPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config:            testAccCloudRunService_cloudRunServiceUpdate(name, project, "10", "600"),
				ExternalProviders: oldVersion,
			},
			{
				Config:                   testAccCloudRunService_cloudRunServiceUpdate(name, project, "10", "600"),
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.%", "2"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.%", "3"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.annotations.%", "1"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_annotations.%", "6"),
				),
			},
		},
	})
}

func TestAccCloudRunService_withComputedLabels(t *testing.T) {
	// Skip it in VCR test because of the randomness of uuid in "labels" field
	// which causes the replaying mode after recording mode failing in VCR test
	acctest.SkipIfVcr(t)
	t.Parallel()

	name := "tftest-cloudrun-" + acctest.RandString(t, 6)
	project := envvar.GetTestProjectFromEnv()

	acctest.VcrTest(t, resource.TestCase{
		PreCheck: func() { acctest.AccTestPreCheck(t) },
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_withComputedLabels(name, project, "10", "600"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
		},
	})
}

func testAccCloudRunService_withComputedLabels(name, project, concurrency, timeoutSeconds string) string {
	return fmt.Sprintf(`
resource "random_uuid" "test" {
}

resource "google_cloud_run_service" "default" {
  name     = "%s"
  location = "us-central1"

  metadata {
    namespace = "%s"
    annotations = {
      env = "${random_uuid.test.result}"
    }
    labels = {
      key1 = "${random_uuid.test.result}"
    }
  }

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        ports {
          container_port = 8080
        }
      }
      container_concurrency = %s
      timeout_seconds = %s
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
    tag             = "magic-module"
  }
}
`, name, project, concurrency, timeoutSeconds)
}

func testAccCloudRunService_withProviderDefaultLabels(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google" {
  default_labels = {
    default_key1 = "default_value1"
  }
  add_terraform_attribution_label = false
}

resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  metadata {
    namespace = "%{project}"
    annotations = {
      generated-by = "magic-modules"
    }
    labels = {
      env                   = "foo"
      default_expiration_ms = 3600000
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
`, context)
}

func testAccCloudRunService_resourceLabelsOverridesProviderDefaultLabels(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google" {
  default_labels = {
    default_key1 = "default_value1"
  }
  add_terraform_attribution_label = false
}

resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  metadata {
    namespace = "%{project}"
    annotations = {
      generated-by = "magic-modules-update"
    }
    labels = {
      env                   = "foo"
      default_expiration_ms = 3600000
      default_key1          = "value1"
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
`, context)
}

func testAccCloudRunService_moveResourceLabelToProviderDefaultLabels(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google" {
  default_labels = {
    default_key1 = "default_value1"
    env          = "foo"
  }
  add_terraform_attribution_label = false
}

resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  metadata {
    namespace = "%{project}"
    annotations = {
      generated-by = "magic-modules"
    }
    labels = {
      default_expiration_ms = 3600000
      default_key1          = "value1"
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
`, context)
}

func testAccCloudRunService_cloudRunServiceBasic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  metadata {
    namespace = "%{project}"
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
`, context)
}

func TestAccCloudRunService_probes(t *testing.T) {
	t.Parallel()

	project := envvar.GetTestProjectFromEnv()
	name := "tftest-cloudrun-" + acctest.RandString(t, 6)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceWithEmptyTCPStartupProbeAndHTTPLivenessProbe(name, project),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_cloudRunServiceUpdateWithTCPStartupProbeAndHTTPLivenessProbe(name, project, "2", "1", "5", "2"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_cloudRunServiceUpdateWithEmptyHTTPStartupProbe(name, project),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_cloudRunServiceUpdateWithHTTPStartupProbe(name, project),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_cloudRunServiceUpdateWithEmptyGRPCLivenessProbe(name, project),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_cloudRunServiceUpdateWithGRPCLivenessProbe(name, project),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceWithEmptyTCPStartupProbeAndHTTPLivenessProbe(name, project string) string {
	return fmt.Sprintf(`
resource "google_cloud_run_service" "default" {
  name     = "%s"
  location = "us-central1"

  metadata {
    namespace = "%s"
    annotations = {
      generated-by = "magic-modules"
    }
  }

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        ports {
          container_port = 8080
        }
        startup_probe {
          tcp_socket {}
        }
        liveness_probe {
          http_get {}
        }
      }
    }
  }

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }
}
`, name, project)
}

func testAccCloudRunService_cloudRunServiceUpdateWithTCPStartupProbeAndHTTPLivenessProbe(name, project, delay, timeout, peroid, failure_threshold string) string {
	return fmt.Sprintf(`
resource "google_cloud_run_service" "default" {
  name     = "%s"
  location = "us-central1"

  metadata {
    namespace = "%s"
    annotations = {
      generated-by = "magic-modules"
    }
  }

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        ports {
          container_port = 8080
        }
        startup_probe {
          initial_delay_seconds = %s
          period_seconds = %s
          timeout_seconds = %s
          failure_threshold = %s
          tcp_socket {
            port = 8080
          }
        }
        liveness_probe {
          initial_delay_seconds = %s
          period_seconds = %s
          timeout_seconds = %s
          failure_threshold = %s
          http_get {
            path = "/some-path"
            port = 8080
            http_headers {
              name = "User-Agent"
              value = "magic-modules"
            }
            http_headers {
              name = "Some-Name"
            }
          }
        }
      }
    }
  }

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }
}
`, name, project, delay, peroid, timeout, failure_threshold, delay, peroid, timeout, failure_threshold)
}

func testAccCloudRunService_cloudRunServiceUpdateWithEmptyHTTPStartupProbe(name, project string) string {
	return fmt.Sprintf(`
resource "google_cloud_run_service" "default" {
  name     = "%s"
  location = "us-central1"

  metadata {
    namespace = "%s"
    annotations = {
      generated-by = "magic-modules"
    }
  }

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        startup_probe {
          http_get {}
        }
      }
    }
  }

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }
}
`, name, project)
}

func testAccCloudRunService_cloudRunServiceUpdateWithHTTPStartupProbe(name, project string) string {
	return fmt.Sprintf(`
resource "google_cloud_run_service" "default" {
  name     = "%s"
  location = "us-central1"

  metadata {
    namespace = "%s"
    annotations = {
      generated-by = "magic-modules"
    }
  }

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        startup_probe {
          http_get {
            path = "/some-path"
            port = 8080
            http_headers {
              name = "User-Agent"
              value = "magic-modules"
            }
            http_headers {
              name = "Some-Name"
            }
          }
        }
      }
    }
  }

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }
}
`, name, project)
}

func testAccCloudRunService_cloudRunServiceUpdateWithEmptyGRPCLivenessProbe(name, project string) string {
	return fmt.Sprintf(`
resource "google_cloud_run_service" "default" {
  name     = "%s"
  location = "us-central1"

  metadata {
    namespace = "%s"
    annotations = {
      generated-by = "magic-modules"
    }
  }

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        liveness_probe {
          grpc {}
        }
      }
    }
  }

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }
}
`, name, project)
}

func testAccCloudRunService_cloudRunServiceUpdateWithGRPCLivenessProbe(name, project string) string {
	return fmt.Sprintf(`
resource "google_cloud_run_service" "default" {
  name     = "%s"
  location = "us-central1"

  metadata {
    namespace = "%s"
    annotations = {
      generated-by = "magic-modules"
    }
  }

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        liveness_probe {
          grpc {
            port = 8080
            service = "health"
          }
        }
      }
    }
  }

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }
}
`, name, project)
}

func TestAccCloudRunService_withCreationOnlyAttribution(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":              envvar.GetTestProjectFromEnv(),
		"random_suffix":        acctest.RandString(t, 10),
		"add_attribution":      "true",
		"attribution_strategy": "CREATION_ONLY",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_withAttributionLabelCreate(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.%", "1"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.user_label", "foo"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.%", "2"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.user_label", "foo"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.goog-terraform-provisioned", "true"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.%", "3"), // Includes one label generated by Cloud Run
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.user_label", "foo"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.goog-terraform-provisioned", "true"),
				),
			},
			{
				Config: testAccCloudRunService_withAttributionLabelUpdate(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.%", "1"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.user_label", "bar"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.%", "2"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.user_label", "bar"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.goog-terraform-provisioned", "true"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.%", "3"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.user_label", "bar"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.goog-terraform-provisioned", "true"),
				),
			},
			{
				Config: testAccCloudRunService_withAttributionLabelClear(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("google_cloud_run_service.default", "metadata.0.labels.%"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.%", "1"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.goog-terraform-provisioned", "true"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.%", "2"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.goog-terraform-provisioned", "true"),
				),
			},
		},
	})
}

func TestAccCloudRunService_withProactiveAttribution(t *testing.T) {
	// VCR tests cache provider configuration between steps, this test changes provider configuration and fails under VCR.
	acctest.SkipIfVcr(t)
	t.Parallel()

	suffix := acctest.RandString(t, 10)
	createContext := map[string]interface{}{
		"project":              envvar.GetTestProjectFromEnv(),
		"random_suffix":        suffix,
		"add_attribution":      "false",
		"attribution_strategy": "PROACTIVE",
	}
	updateContext := map[string]interface{}{
		"project":              envvar.GetTestProjectFromEnv(),
		"random_suffix":        suffix,
		"add_attribution":      "true",
		"attribution_strategy": "PROACTIVE",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_withAttributionLabelCreate(createContext),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.%", "1"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.user_label", "foo"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.%", "1"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.user_label", "foo"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.%", "2"), // Includes one label generated by Cloud Run
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.user_label", "foo"),
				),
			},
			{
				Config: testAccCloudRunService_withAttributionLabelUpdate(updateContext),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.%", "1"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.labels.user_label", "bar"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.%", "2"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.user_label", "bar"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.goog-terraform-provisioned", "true"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.%", "3"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.user_label", "bar"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.goog-terraform-provisioned", "true"),
				),
			},
			{
				Config: testAccCloudRunService_withAttributionLabelClear(updateContext),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("google_cloud_run_service.default", "metadata.0.labels.%"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.%", "1"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.terraform_labels.goog-terraform-provisioned", "true"),

					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.%", "2"),
					resource.TestCheckResourceAttr("google_cloud_run_service.default", "metadata.0.effective_labels.goog-terraform-provisioned", "true"),
				),
			},
		},
	})
}

func testAccCloudRunService_withAttributionLabelCreate(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google" {
  add_terraform_attribution_label               = %{add_attribution}
  terraform_attribution_label_addition_strategy = "%{attribution_strategy}"
}

resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  metadata {
    namespace = "%{project}"
    labels = {
      user_label = "foo"
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
`, context)
}

func testAccCloudRunService_withAttributionLabelUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google" {
  add_terraform_attribution_label               = %{add_attribution}
  terraform_attribution_label_addition_strategy = "%{attribution_strategy}"
}

resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  metadata {
    namespace = "%{project}"
    labels = {
      user_label = "bar"
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
`, context)
}

func testAccCloudRunService_withAttributionLabelClear(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google" {
  add_terraform_attribution_label               = %{add_attribution}
  terraform_attribution_label_addition_strategy = "%{attribution_strategy}"
}

resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  metadata {
    namespace = "%{project}"
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
`, context)
}

func TestAccCloudRunService_csiVolume(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	project := envvar.GetTestProjectFromEnv()
	name := "tftest-cloudrun-" + acctest.RandString(t, 6)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceWithNoVolume(name, project),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_cloudRunServiceUpdateWithGcsVolume(name, project),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceWithNoVolume(name, project string) string {
	return fmt.Sprintf(`
resource "google_cloud_run_service" "default" {
  name     = "%s"
  location = "us-central1"

  metadata {
    namespace = "%s"
    annotations = {
      generated-by = "magic-modules"
    }
  }

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
      }
    }
  }

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }
}
`, name, project)
}

func testAccCloudRunService_cloudRunServiceUpdateWithGcsVolume(name, project string) string {
	return fmt.Sprintf(`
resource "google_cloud_run_service" "default" {
  name     = "%s"
  location = "us-central1"

  metadata {
    namespace = "%s"
    annotations = {
      generated-by = "magic-modules"
    }
  }

  template {
    metadata {
      annotations = {
        "run.googleapis.com/execution-environment" = "gen2"
      }
    }
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        volume_mounts {
          name = "vol1"
          mount_path = "/mnt/vol1"
        }
      }
      volumes {
        name = "vol1"
        csi {
          driver = "gcsfuse.run.googleapis.com"
          read_only = true
          volume_attributes = {
            bucketName = "gcp-public-data-landsat"
          }
        }
      }
    }
  }

}
`, name, project)
}

func TestAccCloudRunService_emptyDirVolume(t *testing.T) {
	t.Parallel()

	project := envvar.GetTestProjectFromEnv()
	name := "tftest-cloudrun-" + acctest.RandString(t, 6)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceWithEmptyDirVolume(name, project),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceWithEmptyDirVolume(name, project string) string {
	return fmt.Sprintf(`
resource "google_cloud_run_service" "default" {
  provider = google-beta
  name     = "%s"
  location = "us-central1"

  metadata {
    namespace = "%s"
    annotations = {
      generated-by = "magic-modules"
      "run.googleapis.com/launch-stage" = "BETA"
    }
  }

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        volume_mounts {
          name = "vol1"
          mount_path = "/mnt/vol1"
        }
      }
      volumes {
        name = "vol1"
        empty_dir { size_limit = "256Mi" }
      }
    }
  }

}
`, name, project)
}

func TestAccCloudRunService_resourcesRequirements(t *testing.T) {
	t.Parallel()

	project := envvar.GetTestProjectFromEnv()
	name := "tftest-cloudrun-" + acctest.RandString(t, 6)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunV2Service_cloudrunServiceWithoutGpu(name, project),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunV2Service_cloudrunServiceWithGpu(name, project),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunV2Service_cloudrunServiceWithoutGpu(name, project),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "metadata.0.annotations", "metadata.0.labels", "metadata.0.terraform_labels", "status.0.conditions"},
			},
		},
	})
}

func testAccCloudRunV2Service_cloudrunServiceWithoutGpu(name, project string) string {
	return fmt.Sprintf(`
resource "google_cloud_run_service" "default" {
  provider = google-beta
  name     = "%s"
  location = "us-central1"

  metadata {
    namespace = "%s"
  }

  template {
    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale": "1"
        "run.googleapis.com/cpu-throttling": "false"
      }
    }
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        resources {
          limits = {
            "cpu" = "4"
            "memory" = "16Gi"
          }
        }
      }
    }
  }
}
`, name, project)
}

func testAccCloudRunV2Service_cloudrunServiceWithGpu(name, project string) string {
	return fmt.Sprintf(`
resource "google_cloud_run_service" "default" {
  provider = google-beta
  name     = "%s"
  location = "us-central1"

  metadata {
    namespace = "%s"
    annotations = {
      "run.googleapis.com/launch-stage" = "BETA"
    }
  }

  template {
    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale": "1"
        "run.googleapis.com/cpu-throttling": "false"
      }
    }
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        resources {
          limits = {
            "cpu" = "4"
            "memory" = "16Gi"
            "nvidia.com/gpu" = "1"
          }
        }
      }
      node_selector = {
        "run.googleapis.com/accelerator" = "nvidia-l4"
      }
    }
  }
}
`, name, project)
}
