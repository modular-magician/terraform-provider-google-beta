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
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const workforcePoolIdRegexp = `^[a-z][a-z0-9-]{4,61}[a-z0-9]$`

func ValidateWorkforcePoolId(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	if strings.HasPrefix(value, "gcp-") {
		errors = append(errors, fmt.Errorf(
			"%q (%q) can not start with \"gcp-\". "+
				"The prefix `gcp-` is reserved for use by Google, and may not be specified.", k, value))
	}

	if !regexp.MustCompile(workforcePoolIdRegexp).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must contain only lowercase letters [a-z], digits [0-9], and hyphens "+
				"[-]. The WorkforcePool ID must be between 6 and 63 characters, begin "+
				"with a letter, and cannot have a trailing hyphen.", k, value))
	}

	return
}

func ResourceIAMWorkforcePoolWorkforcePool() *schema.Resource {
	return &schema.Resource{
		Create: resourceIAMWorkforcePoolWorkforcePoolCreate,
		Read:   resourceIAMWorkforcePoolWorkforcePoolRead,
		Update: resourceIAMWorkforcePoolWorkforcePoolUpdate,
		Delete: resourceIAMWorkforcePoolWorkforcePoolDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIAMWorkforcePoolWorkforcePoolImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location for the resource.`,
			},
			"parent": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Immutable. The resource name of the parent. Format: 'organizations/{org-id}'.`,
			},
			"workforce_pool_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: ValidateWorkforcePoolId,
				Description: `The name of the pool. The ID must be a globally unique string of 6 to 63 lowercase letters,
digits, or hyphens. It must start with a letter, and cannot have a trailing hyphen.
The prefix 'gcp-' is reserved for use by Google, and may not be specified.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A user-specified description of the pool. Cannot exceed 256 characters.`,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Whether the pool is disabled. You cannot use a disabled pool to exchange tokens,
or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A user-specified display name of the pool in Google Cloud Console. Cannot exceed 32 characters.`,
			},
			"session_duration": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Duration that the Google Cloud access tokens, console sign-in sessions,
and 'gcloud' sign-in sessions from this pool are valid.
Must be greater than 15 minutes (900s) and less than 12 hours (43200s).
If 'sessionDuration' is not configured, minted credentials have a default duration of one hour (3600s).
A duration in seconds with up to nine fractional digits, ending with ''s''. Example: "'3.5s'".`,
				Default: "3600s",
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The resource name of the pool.
Format: 'locations/{location}/workforcePools/{workforcePoolId}'`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The state of the pool.
 * STATE_UNSPECIFIED: State unspecified.
 * ACTIVE: The pool is active, and may be used in Google Cloud policies.
 * DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted
   after approximately 30 days. You can restore a soft-deleted pool using
   [workforcePools.undelete](https://cloud.google.com/iam/docs/reference/rest/v1/locations.workforcePools/undelete#google.iam.admin.v1.WorkforcePools.UndeleteWorkforcePool).
   You cannot reuse the ID of a soft-deleted pool until it is permanently deleted.
   While a pool is deleted, you cannot use it to exchange tokens, or use
   existing tokens to access resources. If the pool is undeleted, existing
   tokens grant access again.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceIAMWorkforcePoolWorkforcePoolCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	parentProp, err := expandIAMWorkforcePoolWorkforcePoolParent(d.Get("parent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	displayNameProp, err := expandIAMWorkforcePoolWorkforcePoolDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandIAMWorkforcePoolWorkforcePoolDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	disabledProp, err := expandIAMWorkforcePoolWorkforcePoolDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !isEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	sessionDurationProp, err := expandIAMWorkforcePoolWorkforcePoolSessionDuration(d.Get("session_duration"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("session_duration"); !isEmptyValue(reflect.ValueOf(sessionDurationProp)) && (ok || !reflect.DeepEqual(v, sessionDurationProp)) {
		obj["sessionDuration"] = sessionDurationProp
	}

	url, err := ReplaceVars(d, config, "{{IAMWorkforcePoolBasePath}}locations/{{location}}/workforcePools?workforcePoolId={{workforce_pool_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new WorkforcePool: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating WorkforcePool: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "locations/{{location}}/workforcePools/{{workforce_pool_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = IAMWorkforcePoolOperationWaitTime(
		config, res, "Creating WorkforcePool", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create WorkforcePool: %s", err)
	}

	log.Printf("[DEBUG] Finished creating WorkforcePool %q: %#v", d.Id(), res)

	return resourceIAMWorkforcePoolWorkforcePoolRead(d, meta)
}

func resourceIAMWorkforcePoolWorkforcePoolRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{IAMWorkforcePoolBasePath}}locations/{{location}}/workforcePools/{{workforce_pool_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IAMWorkforcePoolWorkforcePool %q", d.Id()))
	}

	res, err = resourceIAMWorkforcePoolWorkforcePoolDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing IAMWorkforcePoolWorkforcePool because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("name", flattenIAMWorkforcePoolWorkforcePoolName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkforcePool: %s", err)
	}
	if err := d.Set("parent", flattenIAMWorkforcePoolWorkforcePoolParent(res["parent"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkforcePool: %s", err)
	}
	if err := d.Set("display_name", flattenIAMWorkforcePoolWorkforcePoolDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkforcePool: %s", err)
	}
	if err := d.Set("description", flattenIAMWorkforcePoolWorkforcePoolDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkforcePool: %s", err)
	}
	if err := d.Set("state", flattenIAMWorkforcePoolWorkforcePoolState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkforcePool: %s", err)
	}
	if err := d.Set("disabled", flattenIAMWorkforcePoolWorkforcePoolDisabled(res["disabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkforcePool: %s", err)
	}
	if err := d.Set("session_duration", flattenIAMWorkforcePoolWorkforcePoolSessionDuration(res["sessionDuration"], d, config)); err != nil {
		return fmt.Errorf("Error reading WorkforcePool: %s", err)
	}

	return nil
}

func resourceIAMWorkforcePoolWorkforcePoolUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	displayNameProp, err := expandIAMWorkforcePoolWorkforcePoolDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandIAMWorkforcePoolWorkforcePoolDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	disabledProp, err := expandIAMWorkforcePoolWorkforcePoolDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	sessionDurationProp, err := expandIAMWorkforcePoolWorkforcePoolSessionDuration(d.Get("session_duration"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("session_duration"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sessionDurationProp)) {
		obj["sessionDuration"] = sessionDurationProp
	}

	url, err := ReplaceVars(d, config, "{{IAMWorkforcePoolBasePath}}locations/{{location}}/workforcePools/{{workforce_pool_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating WorkforcePool %q: %#v", d.Id(), obj)
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

	if d.HasChange("session_duration") {
		updateMask = append(updateMask, "sessionDuration")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating WorkforcePool %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating WorkforcePool %q: %#v", d.Id(), res)
	}

	err = IAMWorkforcePoolOperationWaitTime(
		config, res, "Updating WorkforcePool", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceIAMWorkforcePoolWorkforcePoolRead(d, meta)
}

func resourceIAMWorkforcePoolWorkforcePoolDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := ReplaceVars(d, config, "{{IAMWorkforcePoolBasePath}}locations/{{location}}/workforcePools/{{workforce_pool_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting WorkforcePool %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "WorkforcePool")
	}

	err = IAMWorkforcePoolOperationWaitTime(
		config, res, "Deleting WorkforcePool", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting WorkforcePool %q: %#v", d.Id(), res)
	return nil
}

func resourceIAMWorkforcePoolWorkforcePoolImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"locations/(?P<location>[^/]+)/workforcePools/(?P<workforce_pool_id>[^/]+)",
		"(?P<location>[^/]+)/(?P<workforce_pool_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "locations/{{location}}/workforcePools/{{workforce_pool_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIAMWorkforcePoolWorkforcePoolName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolParent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolDisabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolWorkforcePoolSessionDuration(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandIAMWorkforcePoolWorkforcePoolParent(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolDisplayName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolDisabled(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolSessionDuration(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceIAMWorkforcePoolWorkforcePoolDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if v := res["state"]; v == "DELETED" {
		return nil, nil
	}

	return res, nil
}
