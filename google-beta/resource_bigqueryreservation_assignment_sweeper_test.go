package google

import (
	"context"
	"log"
	"testing"

	bigqueryreservation "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/beta"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("BigqueryReservationAssignment", &resource.Sweeper{
		Name: "BigqueryReservationAssignment",
		F:    testSweepBigqueryReservationAssignment,
	})
}

func testSweepBigqueryReservationAssignment(region string) error {
	resourceName := "BigqueryReservationAssignment"
	log.Printf("[INFO][SWEEPER_LOG] Starting sweeper for %s", resourceName)

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

	err = config.clientBigqueryReservationDCL.DeleteAllAssignment(context.Background(), d["project"], d["location"], d["reservation"], isDeletableBigqueryReservationAssignment)
	if err != nil {
		return err
	}
	return nil
}

func isDeletableBigqueryReservationAssignment(r *bigqueryreservation.Assignment) bool {
	return isSweepableTestResource(*r.Name)
}
