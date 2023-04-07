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

func ResourceActiveDirectoryPeering() *schema.Resource {
	return &schema.Resource{
		Create: resourceActiveDirectoryPeeringCreate,
		Read:   resourceActiveDirectoryPeeringRead,
		Update: resourceActiveDirectoryPeeringUpdate,
		Delete: resourceActiveDirectoryPeeringDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"authorized_network": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The full names of the Google Compute Engine networks to which the instance is connected. Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.`,
			},
			"domain_resource": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Full domain resource path for the Managed AD Domain involved in peering. The resource path should be in the form projects/{projectId}/locations/global/domains/{domainName}`,
			},
			"peering_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Resource labels that can contain user-provided metadata`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The current state of this Peering.`,
			},
			"status_message": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Additional information about the current status of this peering, if available.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Unique name of the peering in this scope including projects and location using the form: projects/{projectId}/locations/global/peerings/{peeringId}.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceActiveDirectoryPeeringCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandActiveDirectoryPeeringLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	authorizedNetworkProp, err := expandActiveDirectoryPeeringAuthorizedNetwork(d.Get("authorized_network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authorized_network"); !isEmptyValue(reflect.ValueOf(authorizedNetworkProp)) && (ok || !reflect.DeepEqual(v, authorizedNetworkProp)) {
		obj["authorizedNetwork"] = authorizedNetworkProp
	}
	domainResourceProp, err := expandActiveDirectoryPeeringDomainResource(d.Get("domain_resource"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("domain_resource"); !isEmptyValue(reflect.ValueOf(domainResourceProp)) && (ok || !reflect.DeepEqual(v, domainResourceProp)) {
		obj["domainResource"] = domainResourceProp
	}
	statusMessageProp, err := expandActiveDirectoryPeeringStatusMessage(d.Get("status_message"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("status_message"); !isEmptyValue(reflect.ValueOf(statusMessageProp)) && (ok || !reflect.DeepEqual(v, statusMessageProp)) {
		obj["statusMessage"] = statusMessageProp
	}

	url, err := ReplaceVars(d, config, "{{ActiveDirectoryBasePath}}projects/{{project}}/locations/global/peerings?peeringId={{peering_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Peering: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Peering: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Peering: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/global/domains/{{peering_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = ActiveDirectoryOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Peering", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Peering: %s", err)
	}

	if err := d.Set("name", flattenActiveDirectoryPeeringName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "projects/{{project}}/locations/global/domains/{{peering_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Peering %q: %#v", d.Id(), res)

	return resourceActiveDirectoryPeeringRead(d, meta)
}

func resourceActiveDirectoryPeeringRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ActiveDirectoryBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Peering: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ActiveDirectoryPeering %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Peering: %s", err)
	}

	if err := d.Set("name", flattenActiveDirectoryPeeringName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Peering: %s", err)
	}
	if err := d.Set("labels", flattenActiveDirectoryPeeringLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Peering: %s", err)
	}
	if err := d.Set("authorized_network", flattenActiveDirectoryPeeringAuthorizedNetwork(res["authorizedNetwork"], d, config)); err != nil {
		return fmt.Errorf("Error reading Peering: %s", err)
	}
	if err := d.Set("domain_resource", flattenActiveDirectoryPeeringDomainResource(res["domainResource"], d, config)); err != nil {
		return fmt.Errorf("Error reading Peering: %s", err)
	}

	return nil
}

func resourceActiveDirectoryPeeringUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Peering: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandActiveDirectoryPeeringLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	statusMessageProp, err := expandActiveDirectoryPeeringStatusMessage(d.Get("status_message"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("status_message"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, statusMessageProp)) {
		obj["statusMessage"] = statusMessageProp
	}

	url, err := ReplaceVars(d, config, "{{ActiveDirectoryBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Peering %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Peering %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Peering %q: %#v", d.Id(), res)
	}

	err = ActiveDirectoryOperationWaitTime(
		config, res, project, "Updating Peering", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceActiveDirectoryPeeringRead(d, meta)
}

func resourceActiveDirectoryPeeringDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Peering: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{ActiveDirectoryBasePath}}projects/{{project}}/locations/global/peerings/{{peering_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Peering %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Peering")
	}

	err = ActiveDirectoryOperationWaitTime(
		config, res, project, "Deleting Peering", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Peering %q: %#v", d.Id(), res)
	return nil
}

func flattenActiveDirectoryPeeringName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenActiveDirectoryPeeringLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenActiveDirectoryPeeringAuthorizedNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenActiveDirectoryPeeringDomainResource(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandActiveDirectoryPeeringLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandActiveDirectoryPeeringAuthorizedNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandActiveDirectoryPeeringDomainResource(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandActiveDirectoryPeeringStatusMessage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
