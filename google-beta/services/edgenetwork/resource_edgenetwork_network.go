// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

package edgenetwork

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceEdgenetworkNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceEdgenetworkNetworkCreate,
		Read:   resourceEdgenetworkNetworkRead,
		Delete: resourceEdgenetworkNetworkDelete,

		Importer: &schema.ResourceImporter{
			State: resourceEdgenetworkNetworkImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The Google Cloud region to which the target Distributed Cloud Edge zone belongs.`,
			},
			"network_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `A unique ID that identifies this network.`,
			},
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the target Distributed Cloud Edge zone.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `A free-text description of the resource. Max length 1024 characters.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: `Labels associated with this resource.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"mtu": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: `IP (L3) MTU value of the network. Default value is '1500'. Possible values are: '1500', '9000'.`,
				Default:     1500,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The time when the subnet was created.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
fractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The canonical name of this resource, with format
'projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}'`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The time when the subnet was last updated.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
fractional digits. Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.`,
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

func resourceEdgenetworkNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandEdgenetworkNetworkLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	descriptionProp, err := expandEdgenetworkNetworkDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	mtuProp, err := expandEdgenetworkNetworkMtu(d.Get("mtu"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("mtu"); !tpgresource.IsEmptyValue(reflect.ValueOf(mtuProp)) && (ok || !reflect.DeepEqual(v, mtuProp)) {
		obj["mtu"] = mtuProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{EdgenetworkBasePath}}projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks?networkId={{network_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Network: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Network: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating Network: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = EdgenetworkOperationWaitTime(
		config, res, project, "Creating Network", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Network: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Network %q: %#v", d.Id(), res)

	return resourceEdgenetworkNetworkRead(d, meta)
}

func resourceEdgenetworkNetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{EdgenetworkBasePath}}projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Network: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("EdgenetworkNetwork %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}

	if err := d.Set("name", flattenEdgenetworkNetworkName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}
	if err := d.Set("labels", flattenEdgenetworkNetworkLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}
	if err := d.Set("description", flattenEdgenetworkNetworkDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}
	if err := d.Set("create_time", flattenEdgenetworkNetworkCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}
	if err := d.Set("update_time", flattenEdgenetworkNetworkUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}
	if err := d.Set("mtu", flattenEdgenetworkNetworkMtu(res["mtu"], d, config)); err != nil {
		return fmt.Errorf("Error reading Network: %s", err)
	}

	return nil
}

func resourceEdgenetworkNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Network: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{EdgenetworkBasePath}}projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Network %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Network")
	}

	err = EdgenetworkOperationWaitTime(
		config, res, project, "Deleting Network", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Network %q: %#v", d.Id(), res)
	return nil
}

func resourceEdgenetworkNetworkImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/zones/(?P<zone>[^/]+)/networks/(?P<network_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<zone>[^/]+)/(?P<network_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<zone>[^/]+)/(?P<network_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<network_id>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenEdgenetworkNetworkName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgenetworkNetworkLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgenetworkNetworkDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgenetworkNetworkCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgenetworkNetworkUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgenetworkNetworkMtu(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
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

func expandEdgenetworkNetworkLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandEdgenetworkNetworkDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgenetworkNetworkMtu(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
