package google

import (
	"context"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func init() {
	resource.AddTestSweepers("ApigeeSharedFlow", &resource.Sweeper{
		Name: "ApigeeSharedFlow",
		F:    testSweepApigeeSharedFlow,
	})
}

// At the time of writing, the CI only passes us-central1 as the region
func testSweepApigeeSharedFlow(region string) error {
	resourceName := "ApigeeSharedFlow"
	log.Printf("[INFO][SWEEPER_LOG] Starting sweeper for %s", resourceName)

	config, err := acctest.SharedConfigForRegion(region)
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
	billingId := acctest.GetTestBillingAccountFromEnv(t)

	// Setup variables to replace in list template
	d := &ResourceDataMock{
		FieldsInSchema: map[string]interface{}{
			"project":         config.Project,
			"region":          region,
			"location":        region,
			"zone":            "-",
			"billing_account": billingId,
		},
	}

	listTemplate := strings.Split("https://apigee.googleapis.com/v1/organizations/{{org_id}}/sharedflows/{{name}}", "?")[0]
	listUrl, err := ReplaceVars(d, config, listTemplate)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error preparing sweeper list url: %s", err)
		return nil
	}

	res, err := transport_tpg.SendRequest(config, "GET", config.Project, listUrl, config.UserAgent, nil)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] Error in response from request %s: %s", listUrl, err)
		return nil
	}

	resourceList, ok := res["sharedFlows"]
	if !ok {
		log.Printf("[INFO][SWEEPER_LOG] Nothing found in response.")
		return nil
	}

	rl := resourceList.([]interface{})

	log.Printf("[INFO][SWEEPER_LOG] Found %d items in %s list response.", len(rl), resourceName)
	// Keep count of items that aren't sweepable for logging.
	nonPrefixCount := 0
	for _, ri := range rl {
		obj := ri.(map[string]interface{})
		var name string
		// Id detected in the delete URL, attempt to use id.
		if obj["id"] != nil {
			name = GetResourceNameFromSelfLink(obj["id"].(string))
		} else if obj["name"] != nil {
			name = GetResourceNameFromSelfLink(obj["name"].(string))
		} else {
			log.Printf("[INFO][SWEEPER_LOG] %s resource name and id were nil", resourceName)
			return nil
		}
		// Skip resources that shouldn't be sweeped
		if !acctest.IsSweepableTestResource(name) {
			nonPrefixCount++
			continue
		}

		deleteTemplate := "https://apigee.googleapis.com/v1/organizations/{{org_id}}/sharedflows/{{name}}"
		deleteUrl, err := ReplaceVars(d, config, deleteTemplate)
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] error preparing delete url: %s", err)
			return nil
		}
		deleteUrl = deleteUrl + name

		// Don't wait on operations as we may have a lot to delete
		_, err = transport_tpg.SendRequest(config, "DELETE", config.Project, deleteUrl, config.UserAgent, nil)
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] Error deleting for url %s : %s", deleteUrl, err)
		} else {
			log.Printf("[INFO][SWEEPER_LOG] Sent delete request for %s resource: %s", resourceName, name)
		}
	}

	if nonPrefixCount > 0 {
		log.Printf("[INFO][SWEEPER_LOG] %d items were non-sweepable and skipped.", nonPrefixCount)
	}

	return nil
}
