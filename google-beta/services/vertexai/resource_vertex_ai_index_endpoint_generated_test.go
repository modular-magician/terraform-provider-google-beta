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

package vertexai_test

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

func TestAccVertexAIIndexEndpoint_vertexAiIndexEndpointTestExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"network_name":  acctest.BootstrapSharedServiceNetworkingConnection(t, "vpc-network-1"),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckVertexAIIndexEndpointDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIIndexEndpoint_vertexAiIndexEndpointTestExample(context),
			},
			{
				ResourceName:            "google_vertex_ai_index_endpoint.index_endpoint",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "labels", "public_endpoint_enabled", "region", "terraform_labels"},
			},
		},
	})
}

func testAccVertexAIIndexEndpoint_vertexAiIndexEndpointTestExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_index_endpoint" "index_endpoint" {
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  region       = "us-central1"
  labels       = {
    label-one = "value-one"
  }
  network      = "projects/${data.google_project.project.number}/global/networks/${data.google_compute_network.vertex_network.name}"
}

data "google_compute_network" "vertex_network" {
  name       = "%{network_name}"
}

data "google_project" "project" {}
`, context)
}

func TestAccVertexAIIndexEndpoint_vertexAiIndexEndpointWithPscExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckVertexAIIndexEndpointDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIIndexEndpoint_vertexAiIndexEndpointWithPscExample(context),
			},
			{
				ResourceName:            "google_vertex_ai_index_endpoint.index_endpoint",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "labels", "public_endpoint_enabled", "region", "terraform_labels"},
			},
		},
	})
}

func testAccVertexAIIndexEndpoint_vertexAiIndexEndpointWithPscExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_index_endpoint" "index_endpoint" {
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  region       = "us-central1"
  labels       = {
    label-one = "value-one"
  }

  private_service_connect_config {
    enable_private_service_connect = true
    project_allowlist = [
        data.google_project.project.number,
    ]
  }
}

data "google_project" "project" {}
`, context)
}

func TestAccVertexAIIndexEndpoint_vertexAiIndexEndpointWithFalsePscExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckVertexAIIndexEndpointDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIIndexEndpoint_vertexAiIndexEndpointWithFalsePscExample(context),
			},
		},
	})
}

func testAccVertexAIIndexEndpoint_vertexAiIndexEndpointWithFalsePscExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_index_endpoint" "index_endpoint" {
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  region       = "us-central1"
  labels       = {
    label-one = "value-one"
  }

  private_service_connect_config {
    enable_private_service_connect = false
  }
}
`, context)
}

func TestAccVertexAIIndexEndpoint_vertexAiIndexEndpointWithPublicEndpointExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"network_name":  acctest.BootstrapSharedTestNetwork(t, "vertex-ai-index-endpoint"),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckVertexAIIndexEndpointDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIIndexEndpoint_vertexAiIndexEndpointWithPublicEndpointExample(context),
			},
			{
				ResourceName:            "google_vertex_ai_index_endpoint.index_endpoint",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "labels", "public_endpoint_enabled", "region", "terraform_labels"},
			},
		},
	})
}

func testAccVertexAIIndexEndpoint_vertexAiIndexEndpointWithPublicEndpointExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_index_endpoint" "index_endpoint" {
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint with an public endpoint"
  region       = "us-central1"
  labels       = {
    label-one = "value-one"
  }

  public_endpoint_enabled = true
}
`, context)
}

func testAccCheckVertexAIIndexEndpointDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_vertex_ai_index_endpoint" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/indexEndpoints/{{name}}")
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
				return fmt.Errorf("VertexAIIndexEndpoint still exists at %s", url)
			}
		}

		return nil
	}
}
