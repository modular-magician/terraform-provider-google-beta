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
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	compute "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

func resourceComputeNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeNetworkCreate,
		Read:   resourceComputeNetworkRead,
		Update: resourceComputeNetworkUpdate,
		Delete: resourceComputeNetworkDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeNetworkImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"auto_create_subnetworks": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Default:     true,
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"mtu": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: ``,
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"routing_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: ``,
			},

			"gateway_ipv4": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"delete_default_routes_on_create": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
				Default:     false,
			},
		},
	}
}

func resourceComputeNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &compute.Network{
		Name:                  dcl.String(d.Get("name").(string)),
		AutoCreateSubnetworks: dcl.Bool(d.Get("auto_create_subnetworks").(bool)),
		Description:           dcl.String(d.Get("description").(string)),
		Mtu:                   dcl.Int64(int64(d.Get("mtu").(int))),
		Project:               dcl.String(project),
		RoutingConfig:         expandComputeNetworkRoutingConfigCollapsed(d),
	}

	id, err := replaceVarsForId(d, config, "projects/{{project}}/global/networks/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	createDirective := CreateDirective
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLComputeClient(config, userAgent, billingProject)
	res, err := client.ApplyNetwork(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating Network: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Network %q: %#v", d.Id(), res)

	if err := deleteComputeNetworkDefaultRoutes(d, client, config, res); err != nil {
		return fmt.Errorf("error encountered in post-create: %v", err)
	}

	return resourceComputeNetworkRead(d, meta)
}

func resourceComputeNetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &compute.Network{
		Name:                  dcl.String(d.Get("name").(string)),
		AutoCreateSubnetworks: dcl.Bool(d.Get("auto_create_subnetworks").(bool)),
		Description:           dcl.String(d.Get("description").(string)),
		Mtu:                   dcl.Int64(int64(d.Get("mtu").(int))),
		Project:               dcl.String(project),
		RoutingConfig:         expandComputeNetworkRoutingConfigCollapsed(d),
	}

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLComputeClient(config, userAgent, billingProject)
	res, err := client.GetNetwork(context.Background(), obj)
	if err != nil {
		// Resource not found
		d.SetId("")
		return err
	}

	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("auto_create_subnetworks", res.AutoCreateSubnetworks); err != nil {
		return fmt.Errorf("error setting auto_create_subnetworks in state: %s", err)
	}
	if err = d.Set("description", res.Description); err != nil {
		return fmt.Errorf("error setting description in state: %s", err)
	}
	if err = d.Set("mtu", res.Mtu); err != nil {
		return fmt.Errorf("error setting mtu in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = setStateForCollapsedObject(d, flattenComputeNetworkRoutingConfig(res.RoutingConfig)); err != nil {
		return fmt.Errorf("error setting routing_config in state: %s", err)
	}
	if err = d.Set("gateway_ipv4", res.GatewayIPv4); err != nil {
		return fmt.Errorf("error setting gateway_ipv4 in state: %s", err)
	}
	if err = d.Set("self_link", res.SelfLink); err != nil {
		return fmt.Errorf("error setting self_link in state: %s", err)
	}

	return nil
}
func resourceComputeNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &compute.Network{
		Name:                  dcl.String(d.Get("name").(string)),
		AutoCreateSubnetworks: dcl.Bool(d.Get("auto_create_subnetworks").(bool)),
		Description:           dcl.String(d.Get("description").(string)),
		Mtu:                   dcl.Int64(int64(d.Get("mtu").(int))),
		Project:               dcl.String(project),
		RoutingConfig:         expandComputeNetworkRoutingConfigCollapsed(d),
	}
	directive := UpdateDirective
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLComputeClient(config, userAgent, billingProject)
	res, err := client.ApplyNetwork(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error updating Network: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Network %q: %#v", d.Id(), res)

	return resourceComputeNetworkRead(d, meta)
}

func resourceComputeNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &compute.Network{
		Name:                  dcl.String(d.Get("name").(string)),
		AutoCreateSubnetworks: dcl.Bool(d.Get("auto_create_subnetworks").(bool)),
		Description:           dcl.String(d.Get("description").(string)),
		Mtu:                   dcl.Int64(int64(d.Get("mtu").(int))),
		Project:               dcl.String(project),
		RoutingConfig:         expandComputeNetworkRoutingConfigCollapsed(d),
	}

	log.Printf("[DEBUG] Deleting Network %q", d.Id())
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLComputeClient(config, userAgent, billingProject)
	if err := client.DeleteNetwork(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting Network: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting Network %q", d.Id())
	return nil
}

func resourceComputeNetworkImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/networks/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/global/networks/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	d.Set("delete_default_routes_on_create", false)

	return []*schema.ResourceData{d}, nil
}

func expandComputeNetworkRoutingConfigCollapsed(d *schema.ResourceData) *compute.NetworkRoutingConfig {
	collapsed := compute.NetworkRoutingConfig{
		RoutingMode: compute.NetworkRoutingConfigRoutingModeEnumRef(d.Get("routing_mode").(string)),
	}
	// Return nil if empty
	if (compute.NetworkRoutingConfig{}) == collapsed {
		return nil
	}
	return &collapsed
}

func flattenComputeNetworkRoutingConfig(obj *compute.NetworkRoutingConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"routing_mode": obj.RoutingMode,
	}

	return []interface{}{transformed}

}
