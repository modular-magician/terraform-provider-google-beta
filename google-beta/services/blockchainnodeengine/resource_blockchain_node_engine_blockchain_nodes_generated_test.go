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

package blockchainnodeengine_test

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

func TestAccBlockchainNodeEngineBlockchainNodes_blockchainNodesBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBlockchainNodeEngineBlockchainNodesDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBlockchainNodeEngineBlockchainNodes_blockchainNodesBasicExample(context),
			},
			{
				ResourceName:            "google_blockchain_node_engine_blockchain_nodes.default_node",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"blockchain_node_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccBlockchainNodeEngineBlockchainNodes_blockchainNodesBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_blockchain_node_engine_blockchain_nodes" "default_node" {
  location = "us-central1"
  blockchain_type = "ETHEREUM"
  blockchain_node_id = "tf_test_blockchain_basic_node%{random_suffix}"
  ethereum_details {
    api_enable_admin = true
    api_enable_debug = true
    validator_config {
      mev_relay_urls = ["https://mev1.example.org/","https://mev2.example.org/"]
    }
    node_type = "ARCHIVE"
    consensus_client = "LIGHTHOUSE"
    execution_client = "ERIGON"
    network = "MAINNET"
  }
  
  labels = {
    environment = "dev"
  }
}
`, context)
}

func TestAccBlockchainNodeEngineBlockchainNodes_blockchainNodesGethDetailsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBlockchainNodeEngineBlockchainNodesDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBlockchainNodeEngineBlockchainNodes_blockchainNodesGethDetailsExample(context),
			},
			{
				ResourceName:            "google_blockchain_node_engine_blockchain_nodes.default_node_geth",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"blockchain_node_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccBlockchainNodeEngineBlockchainNodes_blockchainNodesGethDetailsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_blockchain_node_engine_blockchain_nodes" "default_node_geth" {
  location = "us-central1"
  blockchain_type = "ETHEREUM"
  blockchain_node_id = "tf_test_blockchain_geth_node%{random_suffix}"
  ethereum_details {
    api_enable_admin = true
    api_enable_debug = true
    validator_config {
      mev_relay_urls = ["https://mev1.example.org/","https://mev2.example.org/"]
    }
    node_type = "FULL"
    consensus_client = "LIGHTHOUSE"
    execution_client = "GETH"
    network = "MAINNET"
    geth_details {
      garbage_collection_mode = "FULL"
    }
  }
  
  labels = {
    environment = "dev"
  }
}
`, context)
}

func testAccCheckBlockchainNodeEngineBlockchainNodesDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_blockchain_node_engine_blockchain_nodes" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{BlockchainNodeEngineBasePath}}projects/{{project}}/locations/{{location}}/blockchainNodes/{{blockchain_node_id}}")
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
				return fmt.Errorf("BlockchainNodeEngineBlockchainNodes still exists at %s", url)
			}
		}

		return nil
	}
}
