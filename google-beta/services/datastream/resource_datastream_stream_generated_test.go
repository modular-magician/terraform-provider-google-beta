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

package datastream_test

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

func TestAccDatastreamStream_datastreamStreamBasicExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckDatastreamStreamDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDatastreamStream_datastreamStreamBasicExample(context),
			},
			{
				ResourceName:            "google_datastream_stream.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "stream_id", "terraform_labels"},
			},
		},
	})
}

func testAccDatastreamStream_datastreamStreamBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
}

resource "google_sql_database_instance" "instance" {
    name             = "tf-test-my-instance%{random_suffix}"
    database_version = "MYSQL_8_0"
    region           = "us-central1"
    settings {
        tier = "db-f1-micro"
        backup_configuration {
            enabled            = true
            binary_log_enabled = true
        }

        ip_configuration {

            // Datastream IPs will vary by region.
            authorized_networks {
                value = "34.71.242.81"
            }

            authorized_networks {
                value = "34.72.28.29"
            }

            authorized_networks {
                value = "34.67.6.157"
            }

            authorized_networks {
                value = "34.67.234.134"
            }

            authorized_networks {
                value = "34.72.239.218"
            }
        }
    }

    deletion_protection  = %{deletion_protection}
}

resource "google_sql_database" "db" {
    instance = google_sql_database_instance.instance.name
    name     = "db"
}

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    name     = "user"
    instance = google_sql_database_instance.instance.name
    host     = "%"
    password = random_password.pwd.result
}

resource "google_datastream_connection_profile" "source_connection_profile" {
    display_name          = "Source connection profile"
    location              = "us-central1"
    connection_profile_id = "tf-test-source-profile%{random_suffix}"

    mysql_profile {
        hostname = google_sql_database_instance.instance.public_ip_address
        username = google_sql_user.user.name
        password = google_sql_user.user.password
    }
}

resource "google_storage_bucket" "bucket" {
  name                        = "tf-test-my-bucket%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_iam_member" "viewer" {
    bucket = google_storage_bucket.bucket.name
    role   = "roles/storage.objectViewer"
    member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-datastream.iam.gserviceaccount.com"
}

resource "google_storage_bucket_iam_member" "creator" {
    bucket = google_storage_bucket.bucket.name
    role   = "roles/storage.objectCreator"
    member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-datastream.iam.gserviceaccount.com"
}

resource "google_storage_bucket_iam_member" "reader" {
    bucket = google_storage_bucket.bucket.name
    role   = "roles/storage.legacyBucketReader"
    member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-datastream.iam.gserviceaccount.com"
}

resource "google_datastream_connection_profile" "destination_connection_profile" {
    display_name          = "Connection profile"
    location              = "us-central1"
    connection_profile_id = "tf-test-destination-profile%{random_suffix}"

    gcs_profile {
        bucket    = google_storage_bucket.bucket.name
        root_path = "/path"
    }
}

resource "google_datastream_stream" "default" {
    stream_id = "tf-test-my-stream%{random_suffix}"
    location = "us-central1"
    display_name = "my stream"
    source_config {
        source_connection_profile = google_datastream_connection_profile.source_connection_profile.id
        mysql_source_config {
          max_concurrent_backfill_tasks = 15
        }
    }
    destination_config {
        destination_connection_profile = google_datastream_connection_profile.destination_connection_profile.id
        gcs_destination_config {
            avro_file_format {}
        }
    }

    backfill_none {
    }
}
`, context)
}

func TestAccDatastreamStream_datastreamStreamFullExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"stream_cmek":         acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckDatastreamStreamDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDatastreamStream_datastreamStreamFullExample(context),
			},
			{
				ResourceName:            "google_datastream_stream.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "stream_id", "terraform_labels"},
			},
		},
	})
}

func testAccDatastreamStream_datastreamStreamFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
}

resource "google_sql_database_instance" "instance" {
    name             = "tf-test-my-instance%{random_suffix}"
    database_version = "MYSQL_8_0"
    region           = "us-central1"
    settings {
        tier = "db-f1-micro"
        backup_configuration {
            enabled            = true
            binary_log_enabled = true
        }

        ip_configuration {

            // Datastream IPs will vary by region.
            authorized_networks {
                value = "34.71.242.81"
            }

            authorized_networks {
                value = "34.72.28.29"
            }

            authorized_networks {
                value = "34.67.6.157"
            }

            authorized_networks {
                value = "34.67.234.134"
            }

            authorized_networks {
                value = "34.72.239.218"
            }
        }
    }

    deletion_protection  = %{deletion_protection}
}

resource "google_sql_database" "db" {
    instance = google_sql_database_instance.instance.name
    name     = "db"
}

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    name     = "user"
    instance = google_sql_database_instance.instance.name
    host     = "%"
    password = random_password.pwd.result
}

resource "google_datastream_connection_profile" "source_connection_profile" {
    display_name          = "Source connection profile"
    location              = "us-central1"
    connection_profile_id = "tf-test-source-profile%{random_suffix}"

    mysql_profile {
        hostname = google_sql_database_instance.instance.public_ip_address
        username = google_sql_user.user.name
        password = google_sql_user.user.password
    }
}

resource "google_storage_bucket" "bucket" {
  name                        = "tf-test-my-bucket%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_iam_member" "viewer" {
    bucket = google_storage_bucket.bucket.name
    role   = "roles/storage.objectViewer"
    member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-datastream.iam.gserviceaccount.com"
}

resource "google_storage_bucket_iam_member" "creator" {
    bucket = google_storage_bucket.bucket.name
    role   = "roles/storage.objectCreator"
    member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-datastream.iam.gserviceaccount.com"
}

resource "google_storage_bucket_iam_member" "reader" {
    bucket = google_storage_bucket.bucket.name
    role   = "roles/storage.legacyBucketReader"
    member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-datastream.iam.gserviceaccount.com"
}

resource "google_kms_crypto_key_iam_member" "key_user" {
    crypto_key_id = "%{stream_cmek}"
    role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
    member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-datastream.iam.gserviceaccount.com"
}

resource "google_datastream_connection_profile" "destination_connection_profile" {
    display_name          = "Connection profile"
    location              = "us-central1"
    connection_profile_id = "tf-test-destination-profile%{random_suffix}"

    gcs_profile {
        bucket    = google_storage_bucket.bucket.name
        root_path = "/path"
    }
}

resource "google_datastream_stream" "default" {
    depends_on = [
        google_kms_crypto_key_iam_member.key_user
    ]
    stream_id = "tf-test-my-stream%{random_suffix}"
    desired_state = "NOT_STARTED"
    location = "us-central1"
    display_name = "my stream"
    labels = {
        key = "value"
    }
    source_config {
        source_connection_profile = google_datastream_connection_profile.source_connection_profile.id
        mysql_source_config {
            include_objects {
                mysql_databases {
                    database = "my-database"
                    mysql_tables {
                        table = "includedTable"
                        mysql_columns {
                            column = "includedColumn"
                            data_type = "VARCHAR"
                            collation = "utf8mb4"
                            primary_key = false
                            nullable = false
                            ordinal_position = 0
                        }
                    }
                    mysql_tables {
                        table = "includedTable_2"
                    }
                }
            }
            exclude_objects {
                mysql_databases {
                    database = "my-database"
                    mysql_tables {
                        table = "excludedTable"
                        mysql_columns {
                            column = "excludedColumn"
                            data_type = "VARCHAR"
                            collation = "utf8mb4"
                            primary_key = false
                            nullable = false
                            ordinal_position = 0
                        }
                    }
                }
            }
            max_concurrent_cdc_tasks = 5
        }
    }
    destination_config {
        destination_connection_profile = google_datastream_connection_profile.destination_connection_profile.id
        gcs_destination_config {
            path = "mydata"
            file_rotation_mb = 200
            file_rotation_interval = "60s"
            json_file_format {
                schema_file_format = "NO_SCHEMA_FILE"
                compression = "GZIP"
            }
        }
    }

    backfill_all {
        mysql_excluded_objects {
            mysql_databases {
                database = "my-database"
                mysql_tables {
                    table = "excludedTable"
                    mysql_columns {
                        column = "excludedColumn"
                        data_type = "VARCHAR"
                        collation = "utf8mb4"
                        primary_key = false
                        nullable = false
                        ordinal_position = 0
                    }
                }
            }
        }
    }

    customer_managed_encryption_key = "%{stream_cmek}"
}
`, context)
}

func TestAccDatastreamStream_datastreamStreamPostgresqlBigqueryDatasetIdExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckDatastreamStreamDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDatastreamStream_datastreamStreamPostgresqlBigqueryDatasetIdExample(context),
			},
			{
				ResourceName:            "google_datastream_stream.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "stream_id", "terraform_labels"},
			},
		},
	})
}

func testAccDatastreamStream_datastreamStreamPostgresqlBigqueryDatasetIdExample(context map[string]interface{}) string {
	return acctest.Nprintf(`

resource "google_bigquery_dataset" "postgres" {
  dataset_id    = "postgres%{random_suffix}"
  friendly_name = "postgres"
  description   = "Database of postgres"
  location      = "us-central1"
}

resource "google_datastream_stream" "default" {
  display_name  = "postgres to bigQuery"
  location      = "us-central1"
  stream_id     = "tf-test-postgres-bigquery%{random_suffix}"

   source_config {
    source_connection_profile = google_datastream_connection_profile.source_connection_profile.id
    mysql_source_config {}
  }

  destination_config {
    destination_connection_profile = google_datastream_connection_profile.destination_connection_profile2.id
    bigquery_destination_config {
      data_freshness = "900s"
      single_target_dataset {
        dataset_id = google_bigquery_dataset.postgres.id
      }
    }
  }

  backfill_all {
  }

}

resource "google_datastream_connection_profile" "destination_connection_profile2" {
    display_name          = "Connection profile"
    location              = "us-central1"
    connection_profile_id = "tf-test-dest-profile%{random_suffix}"
    bigquery_profile {}
}

resource "google_sql_database_instance" "instance" {
    name             = "tf-test-instance-name%{random_suffix}"
    database_version = "MYSQL_8_0"
    region           = "us-central1"
    settings {
        tier = "db-f1-micro"
        backup_configuration {
            enabled            = true
            binary_log_enabled = true
        }

        ip_configuration {
            // Datastream IPs will vary by region.
            authorized_networks {
                value = "34.71.242.81"
            }

            authorized_networks {
                value = "34.72.28.29"
            }

            authorized_networks {
                value = "34.67.6.157"
            }

            authorized_networks {
                value = "34.67.234.134"
            }

            authorized_networks {
                value = "34.72.239.218"
            }
        }
    }

    deletion_protection  = false
}

resource "google_sql_database" "db" {
    instance = google_sql_database_instance.instance.name
    name     = "db"
}

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    name     = "tf-test-my-user%{random_suffix}"
    instance = google_sql_database_instance.instance.name
    host     = "%"
    password = random_password.pwd.result
}

resource "google_datastream_connection_profile" "source_connection_profile" {
    display_name          = "Source connection profile"
    location              = "us-central1"
    connection_profile_id = "tf-test-source-profile%{random_suffix}"

    mysql_profile {
        hostname = google_sql_database_instance.instance.public_ip_address
        username = google_sql_user.user.name
        password = google_sql_user.user.password
    }
}
`, context)
}

func TestAccDatastreamStream_datastreamStreamBigqueryExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection":                     false,
		"bigquery_destination_table_kms_key_name": acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"random_suffix":                           acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckDatastreamStreamDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDatastreamStream_datastreamStreamBigqueryExample(context),
			},
			{
				ResourceName:            "google_datastream_stream.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "stream_id", "terraform_labels"},
			},
		},
	})
}

func testAccDatastreamStream_datastreamStreamBigqueryExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
}

resource "google_sql_database_instance" "instance" {
    name             = "tf-test-my-instance%{random_suffix}"
    database_version = "MYSQL_8_0"
    region           = "us-central1"
    settings {
        tier = "db-f1-micro"
        backup_configuration {
            enabled            = true
            binary_log_enabled = true
        }

        ip_configuration {

            // Datastream IPs will vary by region.
            authorized_networks {
                value = "34.71.242.81"
            }

            authorized_networks {
                value = "34.72.28.29"
            }

            authorized_networks {
                value = "34.67.6.157"
            }

            authorized_networks {
                value = "34.67.234.134"
            }

            authorized_networks {
                value = "34.72.239.218"
            }
        }
    }

    deletion_protection  = %{deletion_protection}
}

resource "google_sql_database" "db" {
    instance = google_sql_database_instance.instance.name
    name     = "db"
}

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    name     = "user"
    instance = google_sql_database_instance.instance.name
    host     = "%"
    password = random_password.pwd.result
}

resource "google_datastream_connection_profile" "source_connection_profile" {
    display_name          = "Source connection profile"
    location              = "us-central1"
    connection_profile_id = "tf-test-source-profile%{random_suffix}"

    mysql_profile {
        hostname = google_sql_database_instance.instance.public_ip_address
        username = google_sql_user.user.name
        password = google_sql_user.user.password
    }
}

data "google_bigquery_default_service_account" "bq_sa" {
}

resource "google_kms_crypto_key_iam_member" "bigquery_key_user" {
  crypto_key_id = "%{bigquery_destination_table_kms_key_name}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:${data.google_bigquery_default_service_account.bq_sa.email}"
}

resource "google_datastream_connection_profile" "destination_connection_profile" {
    display_name          = "Connection profile"
    location              = "us-central1"
    connection_profile_id = "tf-test-destination-profile%{random_suffix}"

    bigquery_profile {}
}

resource "google_datastream_stream" "default" {
    depends_on = [
        google_kms_crypto_key_iam_member.bigquery_key_user
    ]
    stream_id = "tf-test-my-stream%{random_suffix}"
    location = "us-central1"
    display_name = "my stream"
    source_config {
        source_connection_profile = google_datastream_connection_profile.source_connection_profile.id
        mysql_source_config {}
    }
    destination_config {
        destination_connection_profile = google_datastream_connection_profile.destination_connection_profile.id
        bigquery_destination_config {
            source_hierarchy_datasets {
                dataset_template {
                    location = "us-central1"
                    kms_key_name = "%{bigquery_destination_table_kms_key_name}"
                }
            }
        }
    }

    backfill_none {
    }
}
`, context)
}

func testAccCheckDatastreamStreamDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_datastream_stream" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DatastreamBasePath}}projects/{{project}}/locations/{{location}}/streams/{{stream_id}}")
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
				return fmt.Errorf("DatastreamStream still exists at %s", url)
			}
		}

		return nil
	}
}
