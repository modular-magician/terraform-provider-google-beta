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

package google

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
)

func ResourceCertificateManagerDnsAuthorization() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertificateManagerDnsAuthorizationCreate,
		Read:   resourceCertificateManagerDnsAuthorizationRead,
		Update: resourceCertificateManagerDnsAuthorizationUpdate,
		Delete: resourceCertificateManagerDnsAuthorizationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCertificateManagerDnsAuthorizationImport,
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
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `A domain which is being authorized. A DnsAuthorization resource covers a
single domain and its wildcard, e.g. authorization for "example.com" can
be used to issue certificates for "example.com" and "*.example.com".`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource; provided by the client when the resource is created.
The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
and all following characters must be a dash, underscore, letter or digit.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A human-readable description of the resource.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Set of label tags associated with the DNS Authorization resource.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"dns_resource_record": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `The structure describing the DNS Resource Record that needs to be added
to DNS configuration for the authorization to be usable by
certificate.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Data of the DNS Resource Record.`,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `Fully qualified name of the DNS Resource Record.
E.g. '_acme-challenge.example.com'.`,
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Type of the DNS Resource Record.`,
						},
					},
				},
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

func resourceCertificateManagerDnsAuthorizationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandCertificateManagerDnsAuthorizationDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandCertificateManagerDnsAuthorizationLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	domainProp, err := expandCertificateManagerDnsAuthorizationDomain(d.Get("domain"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("domain"); !tpgresource.IsEmptyValue(reflect.ValueOf(domainProp)) && (ok || !reflect.DeepEqual(v, domainProp)) {
		obj["domain"] = domainProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CertificateManagerBasePath}}projects/{{project}}/locations/global/dnsAuthorizations?dnsAuthorizationId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new DnsAuthorization: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DnsAuthorization: %s", err)
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
		return fmt.Errorf("Error creating DnsAuthorization: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/dnsAuthorizations/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = CertificateManagerOperationWaitTime(
		config, res, project, "Creating DnsAuthorization", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create DnsAuthorization: %s", err)
	}

	log.Printf("[DEBUG] Finished creating DnsAuthorization %q: %#v", d.Id(), res)

	return resourceCertificateManagerDnsAuthorizationRead(d, meta)
}

func resourceCertificateManagerDnsAuthorizationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CertificateManagerBasePath}}projects/{{project}}/locations/global/dnsAuthorizations/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DnsAuthorization: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("CertificateManagerDnsAuthorization %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading DnsAuthorization: %s", err)
	}

	if err := d.Set("description", flattenCertificateManagerDnsAuthorizationDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading DnsAuthorization: %s", err)
	}
	if err := d.Set("labels", flattenCertificateManagerDnsAuthorizationLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading DnsAuthorization: %s", err)
	}
	if err := d.Set("domain", flattenCertificateManagerDnsAuthorizationDomain(res["domain"], d, config)); err != nil {
		return fmt.Errorf("Error reading DnsAuthorization: %s", err)
	}
	if err := d.Set("dns_resource_record", flattenCertificateManagerDnsAuthorizationDnsResourceRecord(res["dnsResourceRecord"], d, config)); err != nil {
		return fmt.Errorf("Error reading DnsAuthorization: %s", err)
	}

	return nil
}

func resourceCertificateManagerDnsAuthorizationUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DnsAuthorization: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandCertificateManagerDnsAuthorizationDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandCertificateManagerDnsAuthorizationLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CertificateManagerBasePath}}projects/{{project}}/locations/global/dnsAuthorizations/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating DnsAuthorization %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
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
		return fmt.Errorf("Error updating DnsAuthorization %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating DnsAuthorization %q: %#v", d.Id(), res)
	}

	err = CertificateManagerOperationWaitTime(
		config, res, project, "Updating DnsAuthorization", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceCertificateManagerDnsAuthorizationRead(d, meta)
}

func resourceCertificateManagerDnsAuthorizationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DnsAuthorization: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{CertificateManagerBasePath}}projects/{{project}}/locations/global/dnsAuthorizations/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting DnsAuthorization %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "DnsAuthorization")
	}

	err = CertificateManagerOperationWaitTime(
		config, res, project, "Deleting DnsAuthorization", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting DnsAuthorization %q: %#v", d.Id(), res)
	return nil
}

func resourceCertificateManagerDnsAuthorizationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/global/dnsAuthorizations/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/dnsAuthorizations/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenCertificateManagerDnsAuthorizationDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerDnsAuthorizationLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerDnsAuthorizationDomain(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerDnsAuthorizationDnsResourceRecord(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["name"] =
		flattenCertificateManagerDnsAuthorizationDnsResourceRecordName(original["name"], d, config)
	transformed["type"] =
		flattenCertificateManagerDnsAuthorizationDnsResourceRecordType(original["type"], d, config)
	transformed["data"] =
		flattenCertificateManagerDnsAuthorizationDnsResourceRecordData(original["data"], d, config)
	return []interface{}{transformed}
}
func flattenCertificateManagerDnsAuthorizationDnsResourceRecordName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerDnsAuthorizationDnsResourceRecordType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerDnsAuthorizationDnsResourceRecordData(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandCertificateManagerDnsAuthorizationDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerDnsAuthorizationLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCertificateManagerDnsAuthorizationDomain(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
