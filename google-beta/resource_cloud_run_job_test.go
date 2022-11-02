package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudRunJob_cloudRunJobUpdate(t *testing.T) {
	t.Parallel()

	project := getTestProjectFromEnv()
	name := "tftest-cloudrun-" + randString(t, 6)

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunJob_cloudRunJobUpdate(name, project, "10", "600"),
			},
			{
				ResourceName:            "google_cloud_run_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunJob_cloudRunJobUpdate(name, project, "50", "300"),
			},
			{
				ResourceName:            "google_cloud_run_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
		},
	})
}

// this test checks that Terraform does not fail with a 409 recreating the same job
func TestAccCloudRunJob_foregroundDeletion(t *testing.T) {
	t.Parallel()

	project := getTestProjectFromEnv()
	name := "tftest-cloudrun-" + randString(t, 6)

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunJob_cloudRunJobUpdate(name, project, "10", "600"),
			},
			{
				ResourceName:            "google_cloud_run_Job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
			{
				Config: " ", // very explicitly add a space, as the test runner fails if this is just ""
			},
			{
				Config: testAccCloudRunJob_cloudRunJobUpdate(name, project, "10", "600"),
			},
			{
				ResourceName:            "google_cloud_run_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
		},
	})
}

func testAccCloudRunJob_cloudRunJobUpdate(name, project, maxRetries, timeoutSeconds string) string {
	return fmt.Sprintf(`
resource "google_cloud_run_job" "default" {
  name     = "%s"
  location = "us-central1"
  provider = google-beta

  metadata {
  namespace = "%s"
  annotations = {
      generated-by = "magic-modules"
    }
  }
  template {
    spec {
      template {
        spec {
          containers {
            image = "us-docker.pkg.dev/cloudrun/container/hello"
          }
          max_retries = %s
          timeout_seconds = %s
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
`, name, project, concurrency, timeoutSeconds)
}

func TestAccCloudRunJob_secretVolume(t *testing.T) {
	t.Parallel()

	project := getTestProjectFromEnv()
	name := "tftest-cloudrun-" + randString(t, 6)

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunJob_cloudRunJobUpdateWithSecretVolume(name, project, "secret-"+randString(t, 5), "secret-"+randString(t, 6), "google_secret_manager_secret.secret1.secret_id"),
			},
			{
				ResourceName:            "google_cloud_run_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunJob_cloudRunJobUpdateWithSecretVolume(name, project, "secret-"+randString(t, 10), "secret-"+randString(t, 11), "google_secret_manager_secret.secret2.secret_id"),
			},
			{
				ResourceName:            "google_cloud_run_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
		},
	})
}

func testAccCloudRunJob_cloudRunJobUpdateWithSecretVolume(name, project, secretName1, secretName2, secretRef string) string {
	return fmt.Sprintf(`
data "google_project" "project" {
}

resource "google_secret_manager_secret" "secret1" {
  secret_id = "%s"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret" "secret2" {
  secret_id = "%s"
  replication {
    automatic = true
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

resource "google_cloud_run_job" "default" {
  name     = "%s"
  location = "us-central1"
  provider = google-beta

  metadata {
    namespace = "%s"
    annotations = {
      "run.googleapis.com/launch-stage" = "BETA"
      generated-by = "magic-modules"
    }
  }

  template {
    spec {
      template {
        spec {
          containers {
            image = "us-docker.pkg.dev/cloudrun/container/hello"
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
    }
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

func TestAccCloudRunJob_secretEnvironmentVariable(t *testing.T) {
	t.Parallel()

	project := getTestProjectFromEnv()
	name := "tftest-cloudrun-" + randString(t, 6)

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunJob_cloudRunJobUpdateWithSecretEnvVar(name, project, "secret-"+randString(t, 5), "secret-"+randString(t, 6), "google_secret_manager_secret.secret1.secret_id"),
			},
			{
				ResourceName:            "google_cloud_run_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunJob_cloudRunJobUpdateWithSecretEnvVar(name, project, "secret-"+randString(t, 10), "secret-"+randString(t, 11), "google_secret_manager_secret.secret2.secret_id"),
			},
			{
				ResourceName:            "google_cloud_run_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
		},
	})
}

func testAccCloudRunJob_cloudRunJobUpdateWithSecretEnvVar(name, project, secretName1, secretName2, secretRef string) string {
	return fmt.Sprintf(`
data "google_project" "project" {
}

resource "google_secret_manager_secret" "secret1" {
  secret_id = "%s"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret" "secret2" {
  secret_id = "%s"
  replication {
    automatic = true
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

resource "google_cloud_run_job" "default" {
  name     = "%s"
  location = "us-central1"
  provider = google-beta

  metadata {
    namespace = "%s"
    annotations = {
      "run.googleapis.com/launch-stage" = "BETA"
      generated-by = "magic-modules"
    }
  }
  template {
    spec {
      template {
        spec {
          containers {
            image = "us-docker.pkg.dev/cloudrun/container/hello"
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
    }
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
