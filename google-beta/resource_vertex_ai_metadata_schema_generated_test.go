// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"fmt"
	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	vertexai "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai/beta"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

func TestAccVertexAiMetadataSchema_BasicHandWritten(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  getTestProjectFromEnv(),
		"region":        getTestRegionFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVertexAiMetadataSchemaDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAiMetadataSchema_BasicHandWritten(context),
			},
			{
				ResourceName:      "google_vertex_ai_metadata_schema.primary",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccVertexAiMetadataSchema_BasicHandWritten(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vertex_ai_metadata_schema" "primary" {
  location       = "%{region}"
  metadata_store = google_vertex_ai_metadata_store.minimal.name
  name           = "tf-test-schema%{random_suffix}"
  schema         = "title: sample.Schema\ntype: object"
  schema_type    = "ARTIFACT_TYPE"
  schema_version = "1.0.0"
  project        = "%{project_name}"
}

resource "google_vertex_ai_metadata_store" "minimal" {
  region  = "%{region}"
  name    = "tf-test-store%{random_suffix}"
  project = "%{project_name}"
}

`, context)
}

func testAccCheckVertexAiMetadataSchemaDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "rs.google_vertex_ai_metadata_schema" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			billingProject := ""
			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			obj := &vertexai.MetadataSchema{
				Location:      dcl.String(rs.Primary.Attributes["location"]),
				MetadataStore: dcl.String(rs.Primary.Attributes["metadata_store"]),
				Name:          dcl.String(rs.Primary.Attributes["name"]),
				Schema:        dcl.String(rs.Primary.Attributes["schema"]),
				SchemaType:    vertexai.MetadataSchemaSchemaTypeEnumRef(rs.Primary.Attributes["schema_type"]),
				SchemaVersion: dcl.String(rs.Primary.Attributes["schema_version"]),
				Project:       dcl.StringOrNil(rs.Primary.Attributes["project"]),
				CreateTime:    dcl.StringOrNil(rs.Primary.Attributes["create_time"]),
			}

			client := NewDCLVertexAiClient(config, config.userAgent, billingProject, 0)
			_, err := client.GetMetadataSchema(context.Background(), obj)
			if err == nil {
				return fmt.Errorf("google_vertex_ai_metadata_schema still exists %v", obj)
			}
		}
		return nil
	}
}
