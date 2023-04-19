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
)

func ResourceApigeeNatAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceApigeeNatAddressCreate,
		Read:   resourceApigeeNatAddressRead,
		Delete: resourceApigeeNatAddressDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApigeeNatAddressImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The Apigee instance associated with the Apigee environment,
in the format 'organizations/{{org_name}}/instances/{{instance_name}}'.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Resource ID of the NAT address.`,
			},
			"ip_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The allocated NAT IP address.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `State of the NAT IP address.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApigeeNatAddressCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandApigeeNatAddressName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}{{instance_id}}/natAddresses")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new NatAddress: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating NatAddress: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{instance_id}}/natAddresses/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = ApigeeOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating NatAddress", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create NatAddress: %s", err)
	}

	if err := d.Set("name", flattenApigeeNatAddressName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "{{instance_id}}/natAddresses/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating NatAddress %q: %#v", d.Id(), res)

	return resourceApigeeNatAddressRead(d, meta)
}

func resourceApigeeNatAddressRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}{{instance_id}}/natAddresses/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ApigeeNatAddress %q", d.Id()))
	}

	if err := d.Set("name", flattenApigeeNatAddressName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading NatAddress: %s", err)
	}
	if err := d.Set("ip_address", flattenApigeeNatAddressIpAddress(res["ipAddress"], d, config)); err != nil {
		return fmt.Errorf("Error reading NatAddress: %s", err)
	}
	if err := d.Set("state", flattenApigeeNatAddressState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading NatAddress: %s", err)
	}

	return nil
}

func resourceApigeeNatAddressDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}{{instance_id}}/natAddresses/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting NatAddress %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "NatAddress")
	}

	err = ApigeeOperationWaitTime(
		config, res, "Deleting NatAddress", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting NatAddress %q: %#v", d.Id(), res)
	return nil
}

func resourceApigeeNatAddressImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats cannot import fields with forward slashes in their value
	if err := ParseImportId([]string{
		"(?P<instance_id>.+)/natAddresses/(?P<name>.+)",
		"(?P<instance_id>.+)/(?P<name>.+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "{{instance_id}}/natAddresses/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApigeeNatAddressName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApigeeNatAddressIpAddress(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApigeeNatAddressState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandApigeeNatAddressName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
