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
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceComputePublicAdvertisedPrefix() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputePublicAdvertisedPrefixCreate,
		Read:   resourceComputePublicAdvertisedPrefixRead,
		Delete: resourceComputePublicAdvertisedPrefixDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputePublicAdvertisedPrefixImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"dns_verification_ip": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The IPv4 address to be used for reverse DNS verification.`,
			},
			"ip_cidr_range": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The IPv4 address range, in CIDR format, represented by this public advertised prefix.`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. The name must be 1-63 characters long, and
comply with RFC1035. Specifically, the name must be 1-63 characters
long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?'
which means the first character must be a lowercase letter, and all
following characters must be a dash, lowercase letter, or digit,
except the last character, which cannot be a dash.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"INITIAL", "PTR_CONFIGURED", "VALIDATED", "REVERSE_DNS_LOOKUP_FAILED", "PREFIX_CONFIGURATION_IN_PROGRESS", "PREFIX_CONFIGURATION_COMPLETE", "PREFIX_REMOVAL_IN_PROGRESS", ""}),
				Description:  `The status of the public advertised prefix. Possible values: ["INITIAL", "PTR_CONFIGURED", "VALIDATED", "REVERSE_DNS_LOOKUP_FAILED", "PREFIX_CONFIGURATION_IN_PROGRESS", "PREFIX_CONFIGURATION_COMPLETE", "PREFIX_REMOVAL_IN_PROGRESS"]`,
			},
			"shared_secret": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output Only. The shared secret to be used for reverse DNS verification.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputePublicAdvertisedPrefixCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputePublicAdvertisedPrefixDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputePublicAdvertisedPrefixName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	dnsVerificationIpProp, err := expandComputePublicAdvertisedPrefixDnsVerificationIp(d.Get("dns_verification_ip"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dns_verification_ip"); !tpgresource.IsEmptyValue(reflect.ValueOf(dnsVerificationIpProp)) && (ok || !reflect.DeepEqual(v, dnsVerificationIpProp)) {
		obj["dnsVerificationIp"] = dnsVerificationIpProp
	}
	ipCidrRangeProp, err := expandComputePublicAdvertisedPrefixIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(ipCidrRangeProp)) && (ok || !reflect.DeepEqual(v, ipCidrRangeProp)) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}
	statusProp, err := expandComputePublicAdvertisedPrefixStatus(d.Get("status"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("status"); !tpgresource.IsEmptyValue(reflect.ValueOf(statusProp)) && (ok || !reflect.DeepEqual(v, statusProp)) {
		obj["status"] = statusProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/publicAdvertisedPrefixes")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new PublicAdvertisedPrefix: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PublicAdvertisedPrefix: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating PublicAdvertisedPrefix: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/publicAdvertisedPrefixes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating PublicAdvertisedPrefix", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create PublicAdvertisedPrefix: %s", err)
	}

	log.Printf("[DEBUG] Finished creating PublicAdvertisedPrefix %q: %#v", d.Id(), res)

	return resourceComputePublicAdvertisedPrefixRead(d, meta)
}

func resourceComputePublicAdvertisedPrefixRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/publicAdvertisedPrefixes/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PublicAdvertisedPrefix: %s", err)
	}
	billingProject = project

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputePublicAdvertisedPrefix %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading PublicAdvertisedPrefix: %s", err)
	}

	if err := d.Set("description", flattenComputePublicAdvertisedPrefixDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicAdvertisedPrefix: %s", err)
	}
	if err := d.Set("name", flattenComputePublicAdvertisedPrefixName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicAdvertisedPrefix: %s", err)
	}
	if err := d.Set("dns_verification_ip", flattenComputePublicAdvertisedPrefixDnsVerificationIp(res["dnsVerificationIp"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicAdvertisedPrefix: %s", err)
	}
	if err := d.Set("ip_cidr_range", flattenComputePublicAdvertisedPrefixIpCidrRange(res["ipCidrRange"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicAdvertisedPrefix: %s", err)
	}
	if err := d.Set("status", flattenComputePublicAdvertisedPrefixStatus(res["status"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicAdvertisedPrefix: %s", err)
	}
	if err := d.Set("shared_secret", flattenComputePublicAdvertisedPrefixSharedSecret(res["sharedSecret"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicAdvertisedPrefix: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading PublicAdvertisedPrefix: %s", err)
	}

	return nil
}

func resourceComputePublicAdvertisedPrefixDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PublicAdvertisedPrefix: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/publicAdvertisedPrefixes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting PublicAdvertisedPrefix %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "PublicAdvertisedPrefix")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting PublicAdvertisedPrefix", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting PublicAdvertisedPrefix %q: %#v", d.Id(), res)
	return nil
}

func resourceComputePublicAdvertisedPrefixImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/global/publicAdvertisedPrefixes/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/publicAdvertisedPrefixes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputePublicAdvertisedPrefixDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputePublicAdvertisedPrefixName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputePublicAdvertisedPrefixDnsVerificationIp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputePublicAdvertisedPrefixIpCidrRange(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputePublicAdvertisedPrefixStatus(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputePublicAdvertisedPrefixSharedSecret(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandComputePublicAdvertisedPrefixDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicAdvertisedPrefixName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicAdvertisedPrefixDnsVerificationIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicAdvertisedPrefixIpCidrRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicAdvertisedPrefixStatus(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
