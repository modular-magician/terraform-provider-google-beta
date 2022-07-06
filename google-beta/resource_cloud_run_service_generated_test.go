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

func TestAccCloudRunService_cloudRunServiceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceBasicExample(context),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceBasicExample(context map[string]interface{}) string {
	return Nprintf(`
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

  traffic {
    percent         = 100
    latest_revision = true
  }
}
`, context)
}

func TestAccCloudRunService_cloudRunServiceSqlExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceSqlExample(context),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "autogenerate_revision_name"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceSqlExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }

    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale"      = "1000"
        "run.googleapis.com/cloudsql-instances" = google_sql_database_instance.instance.connection_name
        "run.googleapis.com/client-name"        = "terraform"
      }
    }
  }
  autogenerate_revision_name = true
}

resource "google_sql_database_instance" "instance" {
  name             = "tf-test-cloudrun-sql%{random_suffix}"
  region           = "us-east1"
  database_version = "MYSQL_5_7"
  settings {
    tier = "db-f1-micro"
  }

  deletion_protection  = "%{deletion_protection}"
}
`, context)
}

func TestAccCloudRunService_cloudRunServiceNoauthExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceNoauthExample(context),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceNoauthExample(context map[string]interface{}) string {
	return Nprintf(`
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
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location    = google_cloud_run_service.default.location
  project     = google_cloud_run_service.default.project
  service     = google_cloud_run_service.default.name

  policy_data = data.google_iam_policy.noauth.policy_data
}
`, context)
}

func TestAccCloudRunService_cloudRunServiceMultipleEnvironmentVariablesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceMultipleEnvironmentVariablesExample(context),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "autogenerate_revision_name"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceMultipleEnvironmentVariablesExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
        env {
          name = "SOURCE"
          value = "remote"
        }
        env {
          name = "TARGET"
          value = "home"
        }
      }
    }
  }

  metadata {
    annotations = {
      generated-by = "magic-modules"
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
  autogenerate_revision_name = true

  lifecycle {
    ignore_changes = [
        metadata.0.annotations,
    ]
  }
}
`, context)
}

func TestAccCloudRunService_cloudRunServiceScheduledExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckCloudRunServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceScheduledExample(context),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceScheduledExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project_service" "run_api" {
  project                    = "%{project}"
  service                    = "run.googleapis.com"
  disable_dependent_services = true
  disable_on_destroy         = false
}

resource "google_project_service" "iam_api" {
  project                    = "%{project}"
  service                    = "iam.googleapis.com"
  disable_on_destroy         = false
}

resource "google_project_service" "resource_manager_api" {
  project                    = "%{project}"
  service                    = "cloudresourcemanager.googleapis.com"
  disable_on_destroy         = false
}

resource "google_project_service" "scheduler_api" {
  project                    = "%{project}"
  service                    = "cloudscheduler.googleapis.com"
  disable_on_destroy         = false
}

resource "google_cloud_run_service" "default" {
  project  = "%{project}"
  name     = "tf-test-my-scheduled-service%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  # Use an explicit depends_on clause to wait until API is enabled
  depends_on = [
    google_project_service.run_api
  ]
}

resource "google_service_account" "default" {
  project      = "%{project}"
  account_id   = "tf-test-scheduler-sa%{random_suffix}"
  description  = "Cloud Scheduler service account; used to trigger scheduled Cloud Run jobs."
  display_name = "scheduler-sa"

  # Use an explicit depends_on clause to wait until API is enabled
  depends_on = [
    google_project_service.iam_api
  ]
}

resource "google_cloud_scheduler_job" "default" {
  name             = "tf-test-scheduled-cloud-run-job%{random_suffix}"
  description      = "Invoke a Cloud Run container on a schedule."
  schedule         = "*/8 * * * *"
  time_zone        = "America/New_York"
  attempt_deadline = "320s"

  retry_config {
    retry_count = 1
  }

  http_target {
    http_method = "POST"
    uri         = google_cloud_run_service.default.status[0].url

    oidc_token {
      service_account_email = google_service_account.default.email
    }
  }

  # Use an explicit depends_on clause to wait until API is enabled
  depends_on = [
    google_project_service.scheduler_api
  ]
}

resource "google_cloud_run_service_iam_member" "default" {
  project = "%{project}"
  location = google_cloud_run_service.default.location
  service = google_cloud_run_service.default.name
  role = "roles/run.invoker"
  member = "serviceAccount:${google_service_account.default.email}"
}
`, context)
}

func TestAccCloudRunService_cloudRunServiceSecretEnvironmentVariablesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceSecretEnvironmentVariablesExample(context),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "autogenerate_revision_name"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceSecretEnvironmentVariablesExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_project" "project" {
}

resource "google_secret_manager_secret" "secret" {
  secret_id = "secret%{random_suffix}"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret-version-data" {
  secret = google_secret_manager_secret.secret.name
  secret_data = "secret-data"
}

resource "google_secret_manager_secret_iam_member" "secret-access" {
  secret_id = google_secret_manager_secret.secret.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  depends_on = [google_secret_manager_secret.secret]
}

resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        env {
          name = "SECRET_ENV_VAR"
          value_from {
            secret_key_ref {
              name = google_secret_manager_secret.secret.secret_id
              key = "1"
            }
          }
        }
      }
    }
  }

  metadata {
    annotations = {
      generated-by = "magic-modules"
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
  autogenerate_revision_name = true

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }

  depends_on = [google_secret_manager_secret_version.secret-version-data]
}
`, context)
}

func TestAccCloudRunService_cloudRunServiceSecretVolumesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceSecretVolumesExample(context),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "autogenerate_revision_name"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceSecretVolumesExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_project" "project" {
}

resource "google_secret_manager_secret" "secret" {
  secret_id = "secret%{random_suffix}"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret-version-data" {
  secret = google_secret_manager_secret.secret.name
  secret_data = "secret-data"
}

resource "google_secret_manager_secret_iam_member" "secret-access" {
  secret_id = google_secret_manager_secret.secret.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  depends_on = [google_secret_manager_secret.secret]
}

resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv%{random_suffix}"
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
          secret_name = google_secret_manager_secret.secret.secret_id
          default_mode = 292 # 0444
          items {
            key = "1"
            path = "my-secret"
            mode = 256 # 0400
          }
        }
      }
    }
  }

  metadata {
    annotations = {
      generated-by = "magic-modules"
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
  autogenerate_revision_name = true

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }

  depends_on = [google_secret_manager_secret_version.secret-version-data]
}
`, context)
}

func TestAccCloudRunService_cloudRunServiceIngressExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckCloudRunServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceIngressExample(context),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceIngressExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_service" "default" {
    name     = "tf-test-ingress-service%{random_suffix}"
    location = "us-central1"

    template {
      spec {
        containers {
          image = "gcr.io/cloudrun/hello" #public image for your service
        }
      }
    }
    traffic {
      percent         = 100
      latest_revision = true
    }
    metadata {
      annotations = {
        # For valid annotation values and descriptions, see
        # https://cloud.google.com/sdk/gcloud/reference/run/deploy#--ingress
        "run.googleapis.com/ingress" = "internal"
      }
    }
}
`, context)
}

func TestAccCloudRunService_eventarcBasicTfExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckCloudRunServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_eventarcBasicTfExample(context),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccCloudRunService_eventarcBasicTfExample(context map[string]interface{}) string {
	return Nprintf(`
# Used to retrieve project_number later
data "google_project" "project" {
  provider = google-beta
}
        
# Enable Cloud Run API
resource "google_project_service" "run" {
  provider = google-beta
  service            = "run.googleapis.com"
  disable_on_destroy = false
}
    
# Enable Eventarc API
resource "google_project_service" "eventarc" {
  provider = google-beta
  service            = "eventarc.googleapis.com"
  disable_on_destroy = false
}


  
# Deploy Cloud Run service
resource "google_cloud_run_service" "default" {
  provider = google-beta
  name     = "tf-test-cloudrun-hello-tf%{random_suffix}"
  location = "us-east1"

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
      }
    }
  }
  
  traffic {
    percent         = 100
    latest_revision = true
  }
  
  depends_on = [google_project_service.run]
}
      
# Make Cloud Run service publicly accessible
resource "google_cloud_run_service_iam_member" "allUsers" {
  provider = google-beta
  service  = google_cloud_run_service.default.name
  location = google_cloud_run_service.default.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}


  
      
# Create a Pub/Sub trigger
resource "google_eventarc_trigger" "tf-test-trigger-pubsub-tf%{random_suffix}" {
  provider = google-beta
  name     = "tf-test-trigger-pubsub-tf%{random_suffix}"
  location = google_cloud_run_service.default.location
  matching_criteria {
    attribute = "type"   
    value     = "google.cloud.pubsub.topic.v1.messagePublished"
  }
  destination {
    cloud_run_service {
      service = google_cloud_run_service.default.name
      region  = google_cloud_run_service.default.location
    }
  }

  depends_on = [google_project_service.eventarc]
}


# Give default Compute service account eventarc.eventReceiver role
resource "google_project_iam_binding" "project" {
  provider = google-beta
  project = data.google_project.project.id
  role    = "roles/eventarc.eventReceiver"

  members = [
    "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  ]
}
    
# Create an AuditLog for Cloud Storage trigger
resource "google_eventarc_trigger" "tf-test-trigger-auditlog-tf%{random_suffix}" {
  provider = google-beta
  name     = "tf-test-trigger-auditlog-tf%{random_suffix}"
  location = google_cloud_run_service.default.location
  project  = data.google_project.project.id
  matching_criteria {
    attribute = "type"
    value     = "google.cloud.audit.log.v1.written"
  }
  matching_criteria {
    attribute = "serviceName"
    value     = "storage.googleapis.com"
  }
  matching_criteria {
    attribute = "methodName"
    value     = "storage.objects.create"
  }
  destination {
    cloud_run_service {
      service = google_cloud_run_service.default.name
      region  = google_cloud_run_service.default.location
    }
  } 
  service_account = "${data.google_project.project.number}-compute@developer.gserviceaccount.com"

  depends_on = [google_project_service.eventarc]
}

`, context)
}

func TestAccCloudRunService_cloudRunServiceMultipleRegionsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckCloudRunServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceMultipleRegionsExample(context),
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceMultipleRegionsExample(context map[string]interface{}) string {
	return Nprintf(`
# Cloud Run service replicated across multiple GCP regions

resource "google_project_service" "compute_api" {
  project                    = "%{project}"
  service                    = "compute.googleapis.com"
  disable_dependent_services = true
  disable_on_destroy         = false
}

resource "google_project_service" "run_api" {
  project                    = "%{project}"
  service                    = "run.googleapis.com"
  disable_dependent_services = true
  disable_on_destroy         = false
}

variable "domain_name" {
  type    = string
  default = "example.com"
}

variable "run_regions" {
  type    = list(string)
  default = ["us-central1", "europe-west1"]
}

resource "google_compute_global_address" "lb_default" {
  name    = "tf-test-myservice-service-ip%{random_suffix}"
  project = "%{project}"

  # Use an explicit depends_on clause to wait until API is enabled
  depends_on = [
    google_project_service.compute_api
  ]
}

resource "google_compute_backend_service" "lb_default" {
  name                  = "tf-test-myservice-backend%{random_suffix}"
  project               = "%{project}"
  load_balancing_scheme = "EXTERNAL_MANAGED"

  backend {
    balancing_mode  = "UTILIZATION"
    capacity_scaler = 0.85
    group           = google_compute_region_network_endpoint_group.lb_default[0].id
  }

  backend {
    balancing_mode  = "UTILIZATION"
    capacity_scaler = 0.85
    group           = google_compute_region_network_endpoint_group.lb_default[1].id
  }

  # Use an explicit depends_on clause to wait until API is enabled
  depends_on = [
    google_project_service.compute_api,
  ]
}


resource "google_compute_url_map" "lb_default" {
  name            = "tf-test-myservice-lb-urlmap%{random_suffix}"
  project         = "%{project}"
  default_service = google_compute_backend_service.lb_default.id

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.lb_default.id
    route_rules {
      priority = 1
      url_redirect {
        https_redirect         = true
        redirect_response_code = "MOVED_PERMANENTLY_DEFAULT"
      }
    }
  }
}

resource "google_compute_managed_ssl_certificate" "lb_default" {
  name    = "tf-test-myservice-ssl-cert%{random_suffix}"
  project = "%{project}"

  managed {
    domains = [var.domain_name]
  }
}

resource "google_compute_target_https_proxy" "lb_default" {
  name    = "tf-test-myservice-https-proxy%{random_suffix}"
  project = "%{project}"
  url_map = google_compute_url_map.lb_default.id
  ssl_certificates = [
    google_compute_managed_ssl_certificate.lb_default.name
  ]
  depends_on = [
    google_compute_managed_ssl_certificate.lb_default
  ]
}

resource "google_compute_global_forwarding_rule" "lb_default" {
  name                  = "tf-test-myservice-lb-fr%{random_suffix}"
  project               = "%{project}"
  load_balancing_scheme = "EXTERNAL_MANAGED"
  target                = google_compute_target_https_proxy.lb_default.id
  ip_address            = google_compute_global_address.lb_default.id
  port_range            = "443"
  depends_on            = [google_compute_target_https_proxy.lb_default]
}

resource "google_compute_region_network_endpoint_group" "lb_default" {
  count                 = length(var.run_regions)
  project               = "%{project}"
  name                  = "tf-test-myservice-neg%{random_suffix}"
  network_endpoint_type = "SERVERLESS"
  region                = var.run_regions[count.index]
  cloud_run {
    service = google_cloud_run_service.run_default[count.index].name
  }
}

output "load_balancer_ip_addr" {
  value = google_compute_global_address.lb_default.address
}

resource "google_cloud_run_service" "run_default" {
  count    = length(var.run_regions)
  project  = "%{project}"
  name     = "tf-test-myservice-run-app%{random_suffix}-${var.run_regions[count.index]}"
  location = var.run_regions[count.index]

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  # Use an explicit depends_on clause to wait until API is enabled
  depends_on = [
    google_project_service.run_api
  ]
}

resource "google_cloud_run_service_iam_member" "run_allow_unauthenticated" {
  count    = length(var.run_regions)
  project  = "%{project}"
  location = google_cloud_run_service.run_default[count.index].location
  service  = google_cloud_run_service.run_default[count.index].name
  role     = "roles/run.invoker"
  member   = "allUsers"
}

resource "google_compute_url_map" "https_default" {
  name    = "tf-test-myservice-https-urlmap%{random_suffix}"
  project = "%{project}"

  default_url_redirect {
    redirect_response_code = "MOVED_PERMANENTLY_DEFAULT"
    https_redirect         = true
    strip_query            = false
  }
}

resource "google_compute_target_http_proxy" "https_default" {
  name    = "tf-test-myservice-http-proxy%{random_suffix}"
  project = "%{project}"
  url_map = google_compute_url_map.https_default.id

  depends_on = [
    google_compute_url_map.https_default
  ]
}

resource "google_compute_global_forwarding_rule" "https_default" {
  name       = "tf-test-myservice-https-fr%{random_suffix}"
  project    = "%{project}"
  target     = google_compute_target_http_proxy.https_default.id
  ip_address = google_compute_global_address.lb_default.id
  port_range = "80"
  depends_on = [google_compute_target_http_proxy.https_default]
}
`, context)
}

func TestAccCloudRunService_cloudRunSystemPackagesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckCloudRunServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunSystemPackagesExample(context),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunSystemPackagesExample(context map[string]interface{}) string {
	return Nprintf(`
# Example of how to deploy a Cloud Run application with system packages

resource "google_cloud_run_service" "default" {
  name     = "tf-test-graphviz-example%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        # Replace with the URL of your graphviz image
        #   gcr.io/<YOUR_GCP_PROJECT_ID>/graphviz
        image = "gcr.io/cloudrun/hello"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

# Make Cloud Run service publicly accessible
resource "google_cloud_run_service_iam_member" "allow_unauthenticated" {
  service  = google_cloud_run_service.default.name
  location = google_cloud_run_service.default.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}
`, context)
}

func testAccCheckCloudRunServiceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloud_run_service" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{CloudRunBasePath}}apis/serving.knative.dev/v1/namespaces/{{project}}/services/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil, isCloudRunCreationConflict)
			if err == nil {
				return fmt.Errorf("CloudRunService still exists at %s", url)
			}
		}

		return nil
	}
}
