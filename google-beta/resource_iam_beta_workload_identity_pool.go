// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const workloadIdentityPoolIdRegexp = `^[0-9a-z-]+$`

func validateWorkloadIdentityPoolId(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	if strings.HasPrefix(value, "gcp-") {
		errors = append(errors, fmt.Errorf(
			"%q (%q) can not start with \"gcp-\"", k, value))
	}

	if !regexp.MustCompile(workloadIdentityPoolIdRegexp).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q must contain only lowercase letters (a-z), numbers (0-9), or dashes (-)", k))
	}

	if len(value) < 4 {
		errors = append(errors, fmt.Errorf(
			"%q cannot be smaller than 4 characters", k))
	}

	if len(value) > 32 {
		errors = append(errors, fmt.Errorf(
			"%q cannot be greater than 32 characters", k))
	}

	return
}

func resourceIAMBetaWorkloadIdentityPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceIAMBetaWorkloadIdentityPoolCreate,
		Read:   resourceIAMBetaWorkloadIdentityPoolRead,
		Update: resourceIAMBetaWorkloadIdentityPoolUpdate,
		Delete: resourceIAMBetaWorkloadIdentityPoolDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIAMBetaWorkloadIdentityPoolImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"workload_identity_pool_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateWorkloadIdentityPoolId,
				Description: `The ID to use for the pool, which becomes the final component of the resource name. This
value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
'gcp-' is reserved for use by Google, and may not be specified.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A description of the pool. Cannot exceed 256 characters.`,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
existing tokens to access resources. If the pool is re-enabled, existing tokens grant
access again.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A display name for the pool. Cannot exceed 32 characters.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of the pool as
'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}'.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The state of the pool.
* STATE_UNSPECIFIED: State unspecified.
* ACTIVE: The pool is active, and may be used in Google Cloud policies.
* DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted after
  approximately 30 days. You can restore a soft-deleted pool using
  UndeleteWorkloadIdentityPool. You cannot reuse the ID of a soft-deleted pool until it is
  permanently deleted. While a pool is deleted, you cannot use it to exchange tokens, or
  use existing tokens to access resources. If the pool is undeleted, existing tokens grant
  access again.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIAMBetaWorkloadIdentityPoolCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandIAMBetaWorkloadIdentityPoolDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandIAMBetaWorkloadIdentityPoolDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	disabledProp, err := expandIAMBetaWorkloadIdentityPoolDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !isEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}

	url, err := replaceVars(d, config, "{{IAMBetaBasePath}}projects/{{project}}/locations/global/workloadIdentityPools?workloadIdentityPoolId={{workload_identity_pool_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new WorkloadIdentityPool: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for WorkloadIdentityPool: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating WorkloadIdentityPool: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = iAMBetaOperationWaitTime(
		config, res, project, "Creating WorkloadIdentityPool", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create WorkloadIdentityPool: %s", err)
	}

	log.Printf("[DEBUG] Finished creating WorkloadIdentityPool %q: %#v", d.Id(), res)

	return resourceIAMBetaWorkloadIdentityPoolRead(d, meta)
}

func resourceIAMBetaWorkloadIdentityPoolRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{IAMBetaBasePath}}projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for WorkloadIdentityPool: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("IAMBetaWorkloadIdentityPool %q", d.Id()))
	}

	res, err = resourceIAMBetaWorkloadIdentityPoolDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing IAMBetaWorkloadIdentityPool because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading WorkloadIdentityPool: %s", err)
	}

	if err := d.Set("state", flattenIAMBetaWorkloadIdentityPoolState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkloadIdentityPool: %s", err)
	}
	if err := d.Set("display_name", flattenIAMBetaWorkloadIdentityPoolDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkloadIdentityPool: %s", err)
	}
	if err := d.Set("description", flattenIAMBetaWorkloadIdentityPoolDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkloadIdentityPool: %s", err)
	}
	if err := d.Set("name", flattenIAMBetaWorkloadIdentityPoolName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkloadIdentityPool: %s", err)
	}
	if err := d.Set("disabled", flattenIAMBetaWorkloadIdentityPoolDisabled(res["disabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkloadIdentityPool: %s", err)
	}

	return nil
}

func resourceIAMBetaWorkloadIdentityPoolUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for WorkloadIdentityPool: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandIAMBetaWorkloadIdentityPoolDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandIAMBetaWorkloadIdentityPoolDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	disabledProp, err := expandIAMBetaWorkloadIdentityPoolDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}

	url, err := replaceVars(d, config, "{{IAMBetaBasePath}}projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating WorkloadIdentityPool %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("disabled") {
		updateMask = append(updateMask, "disabled")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating WorkloadIdentityPool %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating WorkloadIdentityPool %q: %#v", d.Id(), res)
	}

	err = iAMBetaOperationWaitTime(
		config, res, project, "Updating WorkloadIdentityPool", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceIAMBetaWorkloadIdentityPoolRead(d, meta)
}

func resourceIAMBetaWorkloadIdentityPoolDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for WorkloadIdentityPool: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{IAMBetaBasePath}}projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting WorkloadIdentityPool %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "WorkloadIdentityPool")
	}

	err = iAMBetaOperationWaitTime(
		config, res, project, "Deleting WorkloadIdentityPool", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting WorkloadIdentityPool %q: %#v", d.Id(), res)
	return nil
}

func resourceIAMBetaWorkloadIdentityPoolImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/global/workloadIdentityPools/(?P<workload_identity_pool_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<workload_identity_pool_id>[^/]+)",
		"(?P<workload_identity_pool_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIAMBetaWorkloadIdentityPoolState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIAMBetaWorkloadIdentityPoolDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIAMBetaWorkloadIdentityPoolDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIAMBetaWorkloadIdentityPoolName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIAMBetaWorkloadIdentityPoolDisabled(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandIAMBetaWorkloadIdentityPoolDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIAMBetaWorkloadIdentityPoolDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIAMBetaWorkloadIdentityPoolDisabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceIAMBetaWorkloadIdentityPoolDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if v := res["state"]; v == "DELETED" {
		return nil, nil
	}

	return res, nil
}
