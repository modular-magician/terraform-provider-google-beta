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

package compute

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceComputeRegionSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRegionSecurityPolicyCreate,
		Read:   resourceComputeRegionSecurityPolicyRead,
		Update: resourceComputeRegionSecurityPolicyUpdate,
		Delete: resourceComputeRegionSecurityPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRegionSecurityPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035.
Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.`,
			},
			"ddos_protection_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Configuration for Google Cloud Armor DDOS Proctection Config.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ddos_protection": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: verify.ValidateEnum([]string{"ADVANCED", "ADVANCED_PREVIEW", "STANDARD"}),
							Description: `Google Cloud Armor offers the following options to help protect systems against DDoS attacks:
- STANDARD: basic always-on protection for network load balancers, protocol forwarding, or VMs with public IP addresses.
- ADVANCED: additional protections for Managed Protection Plus subscribers who use network load balancers, protocol forwarding, or VMs with public IP addresses.
- ADVANCED_PREVIEW: flag to enable the security policy in preview mode. Possible values: ["ADVANCED", "ADVANCED_PREVIEW", "STANDARD"]`,
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of this resource. Provide this property when you create the resource.`,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `The Region in which the created Region Security Policy should reside.
If it is not provided, the provider region is used.`,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"CLOUD_ARMOR", "CLOUD_ARMOR_EDGE", "CLOUD_ARMOR_NETWORK", ""}),
				Description: `The type indicates the intended use of the security policy.
- CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers.
- CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache.
- CLOUD_ARMOR_NETWORK: Cloud Armor network policies can be configured to filter packets targeting network load balancing resources such as backend services, target pools, target instances, and instances with external IPs. They filter requests before the request is served from the application.
This field can be set only at resource creation time. Possible values: ["CLOUD_ARMOR", "CLOUD_ARMOR_EDGE", "CLOUD_ARMOR_NETWORK"]`,
			},
			"user_defined_fields": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Definitions of user-defined fields for CLOUD_ARMOR_NETWORK policies.
A user-defined field consists of up to 4 bytes extracted from a fixed offset in the packet, relative to the IPv4, IPv6, TCP, or UDP header, with an optional mask to select certain bits.
Rules may then specify matching values for these fields.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"base": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: verify.ValidateEnum([]string{"IPV4", "IPV6", "TCP", "UDP"}),
							Description: `The base relative to which 'offset' is measured. Possible values are:
- IPV4: Points to the beginning of the IPv4 header.
- IPV6: Points to the beginning of the IPv6 header.
- TCP: Points to the beginning of the TCP header, skipping over any IPv4 options or IPv6 extension headers. Not present for non-first fragments.
- UDP: Points to the beginning of the UDP header, skipping over any IPv4 options or IPv6 extension headers. Not present for non-first fragments. Possible values: ["IPV4", "IPV6", "TCP", "UDP"]`,
						},
						"mask": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `If specified, apply this mask (bitwise AND) to the field to ignore bits before matching.
Encoded as a hexadecimal number (starting with "0x").
The last byte of the field (in network byte order) corresponds to the least significant byte of the mask.`,
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The name of this field. Must be unique within the policy.`,
						},
						"offset": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: `Offset of the first byte of the field (in network byte order) relative to 'base'.`,
						},
						"size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: `Size of the field in bytes. Valid values: 1-4.`,
						},
					},
				},
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Fingerprint of this resource. This field is used internally during
updates of this resource.`,
			},
			"policy_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The unique identifier for the resource. This identifier is defined by the server.`,
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URL for the resource.`,
			},
			"self_link_with_policy_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URL for this resource with the resource id.`,
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

func resourceComputeRegionSecurityPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionSecurityPolicyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeRegionSecurityPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	fingerprintProp, err := expandComputeRegionSecurityPolicyFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	typeProp, err := expandComputeRegionSecurityPolicyType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	ddosProtectionConfigProp, err := expandComputeRegionSecurityPolicyDdosProtectionConfig(d.Get("ddos_protection_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ddos_protection_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(ddosProtectionConfigProp)) && (ok || !reflect.DeepEqual(v, ddosProtectionConfigProp)) {
		obj["ddosProtectionConfig"] = ddosProtectionConfigProp
	}
	userDefinedFieldsProp, err := expandComputeRegionSecurityPolicyUserDefinedFields(d.Get("user_defined_fields"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_defined_fields"); !tpgresource.IsEmptyValue(reflect.ValueOf(userDefinedFieldsProp)) && (ok || !reflect.DeepEqual(v, userDefinedFieldsProp)) {
		obj["userDefinedFields"] = userDefinedFieldsProp
	}
	regionProp, err := expandComputeRegionSecurityPolicyRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/securityPolicies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionSecurityPolicy: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionSecurityPolicy: %s", err)
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
		return fmt.Errorf("Error creating RegionSecurityPolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/securityPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating RegionSecurityPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RegionSecurityPolicy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating RegionSecurityPolicy %q: %#v", d.Id(), res)

	return resourceComputeRegionSecurityPolicyRead(d, meta)
}

func resourceComputeRegionSecurityPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/securityPolicies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionSecurityPolicy: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeRegionSecurityPolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicy: %s", err)
	}

	if err := d.Set("policy_id", flattenComputeRegionSecurityPolicyPolicyId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicy: %s", err)
	}
	if err := d.Set("name", flattenComputeRegionSecurityPolicyName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicy: %s", err)
	}
	if err := d.Set("description", flattenComputeRegionSecurityPolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicy: %s", err)
	}
	if err := d.Set("fingerprint", flattenComputeRegionSecurityPolicyFingerprint(res["fingerprint"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicy: %s", err)
	}
	if err := d.Set("type", flattenComputeRegionSecurityPolicyType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicy: %s", err)
	}
	if err := d.Set("ddos_protection_config", flattenComputeRegionSecurityPolicyDdosProtectionConfig(res["ddosProtectionConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicy: %s", err)
	}
	if err := d.Set("self_link", flattenComputeRegionSecurityPolicySelfLink(res["selfLink"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicy: %s", err)
	}
	if err := d.Set("self_link_with_policy_id", flattenComputeRegionSecurityPolicySelfLinkWithPolicyId(res["selfLinkWithId"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicy: %s", err)
	}
	if err := d.Set("user_defined_fields", flattenComputeRegionSecurityPolicyUserDefinedFields(res["userDefinedFields"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicy: %s", err)
	}
	if err := d.Set("region", flattenComputeRegionSecurityPolicyRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSecurityPolicy: %s", err)
	}

	return nil
}

func resourceComputeRegionSecurityPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionSecurityPolicy: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeRegionSecurityPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	fingerprintProp, err := expandComputeRegionSecurityPolicyFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	ddosProtectionConfigProp, err := expandComputeRegionSecurityPolicyDdosProtectionConfig(d.Get("ddos_protection_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ddos_protection_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, ddosProtectionConfigProp)) {
		obj["ddosProtectionConfig"] = ddosProtectionConfigProp
	}
	userDefinedFieldsProp, err := expandComputeRegionSecurityPolicyUserDefinedFields(d.Get("user_defined_fields"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_defined_fields"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, userDefinedFieldsProp)) {
		obj["userDefinedFields"] = userDefinedFieldsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/securityPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RegionSecurityPolicy %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("fingerprint") {
		updateMask = append(updateMask, "fingerprint")
	}

	if d.HasChange("ddos_protection_config") {
		updateMask = append(updateMask, "ddosProtectionConfig")
	}

	if d.HasChange("user_defined_fields") {
		updateMask = append(updateMask, "userDefinedFields")
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

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating RegionSecurityPolicy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating RegionSecurityPolicy %q: %#v", d.Id(), res)
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Updating RegionSecurityPolicy", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputeRegionSecurityPolicyRead(d, meta)
}

func resourceComputeRegionSecurityPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionSecurityPolicy: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/securityPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting RegionSecurityPolicy %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "RegionSecurityPolicy")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting RegionSecurityPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RegionSecurityPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRegionSecurityPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/securityPolicies/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/securityPolicies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRegionSecurityPolicyPolicyId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyFingerprint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyDdosProtectionConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["ddos_protection"] =
		flattenComputeRegionSecurityPolicyDdosProtectionConfigDdosProtection(original["ddosProtection"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRegionSecurityPolicyDdosProtectionConfigDdosProtection(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicySelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicySelfLinkWithPolicyId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyUserDefinedFields(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"name":   flattenComputeRegionSecurityPolicyUserDefinedFieldsName(original["name"], d, config),
			"base":   flattenComputeRegionSecurityPolicyUserDefinedFieldsBase(original["base"], d, config),
			"offset": flattenComputeRegionSecurityPolicyUserDefinedFieldsOffset(original["offset"], d, config),
			"size":   flattenComputeRegionSecurityPolicyUserDefinedFieldsSize(original["size"], d, config),
			"mask":   flattenComputeRegionSecurityPolicyUserDefinedFieldsMask(original["mask"], d, config),
		})
	}
	return transformed
}
func flattenComputeRegionSecurityPolicyUserDefinedFieldsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyUserDefinedFieldsBase(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyUserDefinedFieldsOffset(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeRegionSecurityPolicyUserDefinedFieldsSize(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeRegionSecurityPolicyUserDefinedFieldsMask(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionSecurityPolicyRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func expandComputeRegionSecurityPolicyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyDdosProtectionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDdosProtection, err := expandComputeRegionSecurityPolicyDdosProtectionConfigDdosProtection(original["ddos_protection"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDdosProtection); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ddosProtection"] = transformedDdosProtection
	}

	return transformed, nil
}

func expandComputeRegionSecurityPolicyDdosProtectionConfigDdosProtection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyUserDefinedFields(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandComputeRegionSecurityPolicyUserDefinedFieldsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedBase, err := expandComputeRegionSecurityPolicyUserDefinedFieldsBase(original["base"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedBase); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["base"] = transformedBase
		}

		transformedOffset, err := expandComputeRegionSecurityPolicyUserDefinedFieldsOffset(original["offset"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOffset); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["offset"] = transformedOffset
		}

		transformedSize, err := expandComputeRegionSecurityPolicyUserDefinedFieldsSize(original["size"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSize); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["size"] = transformedSize
		}

		transformedMask, err := expandComputeRegionSecurityPolicyUserDefinedFieldsMask(original["mask"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMask); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["mask"] = transformedMask
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRegionSecurityPolicyUserDefinedFieldsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyUserDefinedFieldsBase(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyUserDefinedFieldsOffset(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyUserDefinedFieldsSize(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyUserDefinedFieldsMask(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSecurityPolicyRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
