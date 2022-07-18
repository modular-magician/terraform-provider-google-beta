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
	"log"
	"testing"

	vertexai "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai/beta"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("VertexAiEndpoint", &resource.Sweeper{
		Name: "VertexAiEndpoint",
		F:    testSweepVertexAiEndpoint,
	})
}

func testSweepVertexAiEndpoint(region string) error {
	log.Print("[INFO][SWEEPER_LOG] Starting sweeper for VertexAiEndpoint")

	config, err := sharedConfigForRegion(region)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error getting shared config for region: %s", err)
		return err
	}

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error loading: %s", err)
		return err
	}

	t := &testing.T{}
	billingId := getTestBillingAccountFromEnv(t)

	// Setup variables to be used for Delete arguments.
	d := map[string]string{
		"project":         config.Project,
		"region":          region,
		"location":        region,
		"zone":            "-",
		"billing_account": billingId,
	}

	client := NewDCLVertexAiClient(config, config.userAgent, "", 0)
	err = client.DeleteAllEndpoint(context.Background(), d["project"], d["location"], isDeletableVertexAiEndpoint)
	if err != nil {
		return err
	}
	return nil
}

func isDeletableVertexAiEndpoint(r *vertexai.Endpoint) bool {
	return isSweepableTestResource(*r.Name)
}
