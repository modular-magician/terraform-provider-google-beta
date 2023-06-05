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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceNetworkSecurityAddressGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecurityAddressGroupCreate,
		Read:   resourceNetworkSecurityAddressGroupRead,
		Update: resourceNetworkSecurityAddressGroupUpdate,
		Delete: resourceNetworkSecurityAddressGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkSecurityAddressGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"capacity": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Capacity of the Address Group.`,
			},
			"location": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The location of the gateway security policy.
The default value is 'global'.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Name of the AddressGroup resource.`,
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidateEnum([]string{"IPV4", "IPV6"}),
				Description:  `The type of the Address Group. Possible values are "IPV4" or "IPV6". Possible values: ["IPV4", "IPV6"]`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Free-text description of the resource.`,
			},
			"items": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `List of items.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Set of label tags associated with the AddressGroup resource.
An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"parent": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The name of the parent this address group belongs to. Format: organizations/{organization_id} or projects/{project_id}.`,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The timestamp when the resource was created.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The timestamp when the resource was updated.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceNetworkSecurityAddressGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityAddressGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetworkSecurityAddressGroupLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	typeProp, err := expandNetworkSecurityAddressGroupType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	itemsProp, err := expandNetworkSecurityAddressGroupItems(d.Get("items"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("items"); !isEmptyValue(reflect.ValueOf(itemsProp)) && (ok || !reflect.DeepEqual(v, itemsProp)) {
		obj["items"] = itemsProp
	}
	capacityProp, err := expandNetworkSecurityAddressGroupCapacity(d.Get("capacity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("capacity"); !isEmptyValue(reflect.ValueOf(capacityProp)) && (ok || !reflect.DeepEqual(v, capacityProp)) {
		obj["capacity"] = capacityProp
	}

	url, err := ReplaceVars(d, config, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/addressGroups?addressGroupId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AddressGroup: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating AddressGroup: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{parent}}/locations/{{location}}/addressGroups/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkSecurityAddressGroupOperationWaitTime(
		config, res, "Creating AddressGroup", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create AddressGroup: %s", err)
	}

	log.Printf("[DEBUG] Finished creating AddressGroup %q: %#v", d.Id(), res)

	return resourceNetworkSecurityAddressGroupRead(d, meta)
}

func resourceNetworkSecurityAddressGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/addressGroups/{{name}}")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkSecurityAddressGroup %q", d.Id()))
	}

	if err := d.Set("description", flattenNetworkSecurityAddressGroupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading AddressGroup: %s", err)
	}
	if err := d.Set("create_time", flattenNetworkSecurityAddressGroupCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading AddressGroup: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkSecurityAddressGroupUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading AddressGroup: %s", err)
	}
	if err := d.Set("labels", flattenNetworkSecurityAddressGroupLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading AddressGroup: %s", err)
	}
	if err := d.Set("type", flattenNetworkSecurityAddressGroupType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading AddressGroup: %s", err)
	}
	if err := d.Set("items", flattenNetworkSecurityAddressGroupItems(res["items"], d, config)); err != nil {
		return fmt.Errorf("Error reading AddressGroup: %s", err)
	}
	if err := d.Set("capacity", flattenNetworkSecurityAddressGroupCapacity(res["capacity"], d, config)); err != nil {
		return fmt.Errorf("Error reading AddressGroup: %s", err)
	}

	return nil
}

func resourceNetworkSecurityAddressGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityAddressGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetworkSecurityAddressGroupLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	typeProp, err := expandNetworkSecurityAddressGroupType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	itemsProp, err := expandNetworkSecurityAddressGroupItems(d.Get("items"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("items"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, itemsProp)) {
		obj["items"] = itemsProp
	}
	capacityProp, err := expandNetworkSecurityAddressGroupCapacity(d.Get("capacity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("capacity"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, capacityProp)) {
		obj["capacity"] = capacityProp
	}

	url, err := ReplaceVars(d, config, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/addressGroups/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AddressGroup %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("type") {
		updateMask = append(updateMask, "type")
	}

	if d.HasChange("items") {
		updateMask = append(updateMask, "items")
	}

	if d.HasChange("capacity") {
		updateMask = append(updateMask, "capacity")
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
		return fmt.Errorf("Error updating AddressGroup %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating AddressGroup %q: %#v", d.Id(), res)
	}

	err = NetworkSecurityAddressGroupOperationWaitTime(
		config, res, "Updating AddressGroup", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}
	return resourceNetworkSecurityAddressGroupRead(d, meta)
}

func resourceNetworkSecurityAddressGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := ReplaceVars(d, config, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/addressGroups/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting AddressGroup %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "AddressGroup")
	}

	err = NetworkSecurityAddressGroupOperationWaitTime(
		config, res, "Deleting AddressGroup", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting AddressGroup %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkSecurityAddressGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"(?P<parent>.+)/locations/(?P<location>[^/]+)/addressGroups/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "{{parent}}/locations/{{location}}/addressGroups/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkSecurityAddressGroupDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAddressGroupCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAddressGroupUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAddressGroupLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAddressGroupType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAddressGroupItems(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAddressGroupCapacity(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func expandNetworkSecurityAddressGroupDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityAddressGroupLabels(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandNetworkSecurityAddressGroupType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityAddressGroupItems(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityAddressGroupCapacity(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
