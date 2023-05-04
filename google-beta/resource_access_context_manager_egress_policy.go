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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceAccessContextManagerEgressPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessContextManagerEgressPolicyCreate,
		Read:   resourceAccessContextManagerEgressPolicyRead,
		Delete: resourceAccessContextManagerEgressPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessContextManagerEgressPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"egress_policy_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The name of the Service Perimeter to add this resource to.`,
			},
			"resource": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `A GCP resource that is inside of the service perimeter.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceAccessContextManagerEgressPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	resourceProp, err := expandNestedAccessContextManagerEgressPolicyResource(d.Get("resource"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("resource"); !isEmptyValue(reflect.ValueOf(resourceProp)) && (ok || !reflect.DeepEqual(v, resourceProp)) {
		obj["resource"] = resourceProp
	}

	url, err := ReplaceVars(d, config, "{{AccessContextManagerBasePath}}{{egress_policy_name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new EgressPolicy: %#v", obj)

	obj, err = resourceAccessContextManagerEgressPolicyPatchCreateEncoder(d, meta, obj)
	if err != nil {
		return err
	}
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": "status.resources"})
	if err != nil {
		return err
	}
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating EgressPolicy: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{egress_policy_name}}/{{resource}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = AccessContextManagerOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating EgressPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create EgressPolicy: %s", err)
	}

	if _, ok := opRes["status"]; ok {
		opRes, err = flattenNestedAccessContextManagerEgressPolicy(d, meta, opRes)
		if err != nil {
			return fmt.Errorf("Error getting nested object from operation response: %s", err)
		}
		if opRes == nil {
			// Object isn't there any more - remove it from the state.
			return fmt.Errorf("Error decoding response from operation, could not find nested object")
		}
	}
	if err := d.Set("resource", flattenNestedAccessContextManagerEgressPolicyResource(opRes["resource"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "{{egress_policy_name}}/{{resource}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating EgressPolicy %q: %#v", d.Id(), res)

	return resourceAccessContextManagerEgressPolicyRead(d, meta)
}

func resourceAccessContextManagerEgressPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{AccessContextManagerBasePath}}{{egress_policy_name}}")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("AccessContextManagerEgressPolicy %q", d.Id()))
	}

	res, err = flattenNestedAccessContextManagerEgressPolicy(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing AccessContextManagerEgressPolicy because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("resource", flattenNestedAccessContextManagerEgressPolicyResource(res["resource"], d, config)); err != nil {
		return fmt.Errorf("Error reading EgressPolicy: %s", err)
	}

	return nil
}

func resourceAccessContextManagerEgressPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := ReplaceVars(d, config, "{{AccessContextManagerBasePath}}{{egress_policy_name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	obj, err = resourceAccessContextManagerEgressPolicyPatchDeleteEncoder(d, meta, obj)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "EgressPolicy")
	}
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": "status.resources"})
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Deleting EgressPolicy %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "EgressPolicy")
	}

	err = AccessContextManagerOperationWaitTime(
		config, res, "Deleting EgressPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting EgressPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceAccessContextManagerEgressPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	parts, err := getImportIdQualifiers([]string{"accessPolicies/(?P<accessPolicy>[^/]+)/servicePerimeters/(?P<perimeter>[^/]+)/(?P<resource>.+)"}, d, config, d.Id())
	if err != nil {
		return nil, err
	}

	if err := d.Set("egress_policy_name", fmt.Sprintf("accessPolicies/%s/servicePerimeters/%s", parts["accessPolicy"], parts["perimeter"])); err != nil {
		return nil, fmt.Errorf("Error setting egress_policy_name: %s", err)
	}
	if err := d.Set("resource", parts["resource"]); err != nil {
		return nil, fmt.Errorf("Error setting resource: %s", err)
	}
	return []*schema.ResourceData{d}, nil
}

func flattenNestedAccessContextManagerEgressPolicyResource(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNestedAccessContextManagerEgressPolicyResource(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func flattenNestedAccessContextManagerEgressPolicy(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["status"]
	if !ok || v == nil {
		return nil, nil
	}
	res = v.(map[string]interface{})

	v, ok = res["resources"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value status.resources. Actual value: %v", v)
	}

	_, item, err := resourceAccessContextManagerEgressPolicyFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceAccessContextManagerEgressPolicyFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedResource, err := expandNestedAccessContextManagerEgressPolicyResource(d.Get("resource"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedResource := flattenNestedAccessContextManagerEgressPolicyResource(expectedResource, d, meta.(*transport_tpg.Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		// List response only contains the ID - construct a response object.
		item := map[string]interface{}{
			"resource": itemRaw,
		}

		itemResource := flattenNestedAccessContextManagerEgressPolicyResource(item["resource"], d, meta.(*transport_tpg.Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemResource)) && isEmptyValue(reflect.ValueOf(expectedFlattenedResource))) && !reflect.DeepEqual(itemResource, expectedFlattenedResource) {
			log.Printf("[DEBUG] Skipping item with resource= %#v, looking for %#v)", itemResource, expectedFlattenedResource)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}

// PatchCreateEncoder handles creating request data to PATCH parent resource
// with list including new object.
func resourceAccessContextManagerEgressPolicyPatchCreateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceAccessContextManagerEgressPolicyListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	_, found, err := resourceAccessContextManagerEgressPolicyFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}

	// Return error if item already created.
	if found != nil {
		return nil, fmt.Errorf("Unable to create EgressPolicy, existing object already found: %+v", found)
	}

	// Return list with the resource to create appended
	res := map[string]interface{}{
		"resources": append(currItems, obj["resource"]),
	}
	wrapped := map[string]interface{}{
		"status": res,
	}
	res = wrapped

	return res, nil
}

// PatchDeleteEncoder handles creating request data to PATCH parent resource
// with list excluding object to delete.
func resourceAccessContextManagerEgressPolicyPatchDeleteEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceAccessContextManagerEgressPolicyListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	idx, item, err := resourceAccessContextManagerEgressPolicyFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}
	if item == nil {
		// Spoof 404 error for proper handling by Delete (i.e. no-op)
		return nil, fake404("nested", "AccessContextManagerEgressPolicy")
	}

	updatedItems := append(currItems[:idx], currItems[idx+1:]...)
	res := map[string]interface{}{
		"resources": updatedItems,
	}
	wrapped := map[string]interface{}{
		"status": res,
	}
	res = wrapped

	return res, nil
}

// ListForPatch handles making API request to get parent resource and
// extracting list of objects.
func resourceAccessContextManagerEgressPolicyListForPatch(d *schema.ResourceData, meta interface{}) ([]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	url, err := ReplaceVars(d, config, "{{AccessContextManagerBasePath}}{{egress_policy_name}}")
	if err != nil {
		return nil, err
	}

	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return nil, err
	}

	res, err := transport_tpg.SendRequest(config, "GET", "", url, userAgent, nil)
	if err != nil {
		return nil, err
	}

	var v interface{}
	var ok bool
	if v, ok = res["status"]; ok && v != nil {
		res = v.(map[string]interface{})
	} else {
		return nil, nil
	}

	v, ok = res["resources"]
	if ok && v != nil {
		ls, lsOk := v.([]interface{})
		if !lsOk {
			return nil, fmt.Errorf(`expected list for nested field "resources"`)
		}
		return ls, nil
	}
	return nil, nil
}
