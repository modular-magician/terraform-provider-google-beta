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
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

type DataFusionOperationWaiter struct {
	Config    *transport_tpg.Config
	UserAgent string
	Project   string
	tpgresource.CommonOperationWaiter
}

func (w *DataFusionOperationWaiter) QueryOp() (interface{}, error) {
	if w == nil {
		return nil, fmt.Errorf("Cannot query operation, it's unset or nil.")
	}
	// Returns the proper get.
	url := fmt.Sprintf("%s%s", w.Config.DataFusionBasePath, w.CommonOperationWaiter.Op.Name)

	return transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    w.Config,
		Method:    "GET",
		Project:   w.Project,
		RawURL:    url,
		UserAgent: w.UserAgent,
	})
}

func createDataFusionWaiter(config *transport_tpg.Config, op map[string]interface{}, project, activity, userAgent string) (*DataFusionOperationWaiter, error) {
	w := &DataFusionOperationWaiter{
		Config:    config,
		UserAgent: userAgent,
		Project:   project,
	}
	if err := w.CommonOperationWaiter.SetOp(op); err != nil {
		return nil, err
	}
	return w, nil
}

// nolint: deadcode,unused
func DataFusionOperationWaitTimeWithResponse(config *transport_tpg.Config, op map[string]interface{}, response *map[string]interface{}, project, activity, userAgent string, timeout time.Duration) error {
	w, err := createDataFusionWaiter(config, op, project, activity, userAgent)
	if err != nil {
		return err
	}
	if err := tpgresource.OperationWait(w, activity, timeout, config.PollInterval); err != nil {
		return err
	}
	return json.Unmarshal([]byte(w.CommonOperationWaiter.Op.Response), response)
}

func DataFusionOperationWaitTime(config *transport_tpg.Config, op map[string]interface{}, project, activity, userAgent string, timeout time.Duration) error {
	if val, ok := op["name"]; !ok || val == "" {
		// This was a synchronous call - there is no operation to wait for.
		return nil
	}
	w, err := createDataFusionWaiter(config, op, project, activity, userAgent)
	if err != nil {
		// If w is nil, the op was synchronous.
		return err
	}
	return tpgresource.OperationWait(w, activity, timeout, config.PollInterval)
}
