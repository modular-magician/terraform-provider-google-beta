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

package vmwareengine

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceVmwareengineSubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceVmwareengineSubnetCreate,
		Read:   resourceVmwareengineSubnetRead,
		Update: resourceVmwareengineSubnetUpdate,
		Delete: resourceVmwareengineSubnetDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVmwareengineSubnetImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"ip_cidr_range": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The IP address range of the subnet in CIDR format.`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The ID of the subnet. For userDefined subnets, this name should be in the format of "service-n",
where n ranges from 1 to 5.`,
			},
			"parent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The resource name of the private cloud to create a new subnet in.
Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
For example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud`,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Creation time of this resource.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"dhcp_address_ranges": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `DHCP address ranges.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"first_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The first IP address of the range.`,
						},
						"last_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The last IP address of the range.`,
						},
					},
				},
			},
			"gateway_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The canonical identifier of the logical router that this subnet is attached to.`,
			},
			"gateway_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The IP address of the gateway of this subnet. Must fall within the IP prefix defined above.`,
			},
			"standard_config": {
				Type:     schema.TypeBool,
				Computed: true,
				Description: `Whether the NSX-T configuration in the backend follows the standard configuration supported by Google Cloud.
If false, the subnet cannot be modified through Google Cloud, only through NSX-T directly.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `State of the subnet.`,
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The type of the subnet.`,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `System-generated unique identifier for the resource.`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Last updated time of this resource.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"vlan_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `VLAN ID of the VLAN on which the subnet is configured.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceVmwareengineSubnetCreate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	ipCidrRangeProp, err := expandVmwareengineSubnetIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); ok || !reflect.DeepEqual(v, ipCidrRangeProp) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VmwareengineBasePath}}{{parent}}/subnets/{{name}}?update_mask=ip_cidr_range")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Subnet: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating Subnet: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/subnets/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = VmwareengineOperationWaitTime(
		config, res, project, "Creating Subnet", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Subnet: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Subnet %q: %#v", d.Id(), res)

	return resourceVmwareengineSubnetRead(d, meta)
}

func resourceVmwareengineSubnetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VmwareengineBasePath}}{{parent}}/subnets/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("VmwareengineSubnet %q", d.Id()))
	}

	if err := d.Set("create_time", flattenVmwareengineSubnetCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnet: %s", err)
	}
	if err := d.Set("update_time", flattenVmwareengineSubnetUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnet: %s", err)
	}
	if err := d.Set("ip_cidr_range", flattenVmwareengineSubnetIpCidrRange(res["ipCidrRange"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnet: %s", err)
	}
	if err := d.Set("gateway_ip", flattenVmwareengineSubnetGatewayIp(res["gatewayIp"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnet: %s", err)
	}
	if err := d.Set("gateway_id", flattenVmwareengineSubnetGatewayId(res["gatewayId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnet: %s", err)
	}
	if err := d.Set("dhcp_address_ranges", flattenVmwareengineSubnetDhcpAddressRanges(res["dhcpAddressRanges"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnet: %s", err)
	}
	if err := d.Set("type", flattenVmwareengineSubnetType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnet: %s", err)
	}
	if err := d.Set("standard_config", flattenVmwareengineSubnetStandardConfig(res["standardConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnet: %s", err)
	}
	if err := d.Set("state", flattenVmwareengineSubnetState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnet: %s", err)
	}
	if err := d.Set("uid", flattenVmwareengineSubnetUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnet: %s", err)
	}
	if err := d.Set("vlan_id", flattenVmwareengineSubnetVlanId(res["vlanId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subnet: %s", err)
	}

	return nil
}

func resourceVmwareengineSubnetUpdate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	ipCidrRangeProp, err := expandVmwareengineSubnetIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); ok || !reflect.DeepEqual(v, ipCidrRangeProp) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VmwareengineBasePath}}{{parent}}/subnets/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Subnet %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("ip_cidr_range") {
		updateMask = append(updateMask, "ipCidrRange")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating Subnet %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Subnet %q: %#v", d.Id(), res)
		}

		err = VmwareengineOperationWaitTime(
			config, res, project, "Updating Subnet", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceVmwareengineSubnetRead(d, meta)
}

func resourceVmwareengineSubnetDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] Vmwareengine Subnet resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceVmwareengineSubnetImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^(?P<parent>.+)/subnets/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/subnets/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenVmwareengineSubnetCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineSubnetUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineSubnetIpCidrRange(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineSubnetGatewayIp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineSubnetGatewayId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineSubnetDhcpAddressRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"first_address": flattenVmwareengineSubnetDhcpAddressRangesFirstAddress(original["firstAddress"], d, config),
			"last_address":  flattenVmwareengineSubnetDhcpAddressRangesLastAddress(original["lastAddress"], d, config),
		})
	}
	return transformed
}
func flattenVmwareengineSubnetDhcpAddressRangesFirstAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineSubnetDhcpAddressRangesLastAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineSubnetType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineSubnetStandardConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineSubnetState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineSubnetUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineSubnetVlanId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func expandVmwareengineSubnetIpCidrRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
