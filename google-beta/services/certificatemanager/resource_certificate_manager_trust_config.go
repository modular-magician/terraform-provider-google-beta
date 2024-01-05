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

package certificatemanager

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

func ResourceCertificateManagerTrustConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertificateManagerTrustConfigCreate,
		Read:   resourceCertificateManagerTrustConfigRead,
		Update: resourceCertificateManagerTrustConfigUpdate,
		Delete: resourceCertificateManagerTrustConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCertificateManagerTrustConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The trust config location.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `A user-defined name of the trust config. Trust config names must be unique globally.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `One or more paragraphs of text description of a trust config.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Set of label tags associated with the trust config.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"trust_stores": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Set of trust stores to perform validation against.
This field is supported when TrustConfig is configured with Load Balancers, currently not supported for SPIFFE certificate validation.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"intermediate_cas": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `Set of intermediate CA certificates used for the path building phase of chain validation.
The field is currently not supported if trust config is used for the workload certificate feature.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pem_certificate": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `PEM intermediate certificate used for building up paths for validation.
Each certificate provided in PEM format may occupy up to 5kB.`,
										Sensitive: true,
									},
								},
							},
						},
						"trust_anchors": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `List of Trust Anchors to be used while performing validation against a given TrustStore.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pem_certificate": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `PEM root certificate of the PKI used for validation.
Each certificate provided in PEM format may occupy up to 5kB.`,
										Sensitive: true,
									},
								},
							},
						},
					},
				},
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The creation timestamp of a TrustConfig.

A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				ForceNew:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The last update timestamp of a TrustConfig.

A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
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

func resourceCertificateManagerTrustConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandCertificateManagerTrustConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	trustStoresProp, err := expandCertificateManagerTrustConfigTrustStores(d.Get("trust_stores"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("trust_stores"); !tpgresource.IsEmptyValue(reflect.ValueOf(trustStoresProp)) && (ok || !reflect.DeepEqual(v, trustStoresProp)) {
		obj["trustStores"] = trustStoresProp
	}
	labelsProp, err := expandCertificateManagerTrustConfigEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CertificateManagerBasePath}}projects/{{project}}/locations/{{location}}/trustConfigs?trustConfigId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TrustConfig: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TrustConfig: %s", err)
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
		return fmt.Errorf("Error creating TrustConfig: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/trustConfigs/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = CertificateManagerOperationWaitTime(
		config, res, project, "Creating TrustConfig", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create TrustConfig: %s", err)
	}

	log.Printf("[DEBUG] Finished creating TrustConfig %q: %#v", d.Id(), res)

	return resourceCertificateManagerTrustConfigRead(d, meta)
}

func resourceCertificateManagerTrustConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CertificateManagerBasePath}}projects/{{project}}/locations/{{location}}/trustConfigs/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TrustConfig: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("CertificateManagerTrustConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading TrustConfig: %s", err)
	}

	if err := d.Set("create_time", flattenCertificateManagerTrustConfigCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading TrustConfig: %s", err)
	}
	if err := d.Set("update_time", flattenCertificateManagerTrustConfigUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading TrustConfig: %s", err)
	}
	if err := d.Set("labels", flattenCertificateManagerTrustConfigLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading TrustConfig: %s", err)
	}
	if err := d.Set("description", flattenCertificateManagerTrustConfigDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading TrustConfig: %s", err)
	}
	if err := d.Set("trust_stores", flattenCertificateManagerTrustConfigTrustStores(res["trustStores"], d, config)); err != nil {
		return fmt.Errorf("Error reading TrustConfig: %s", err)
	}
	if err := d.Set("terraform_labels", flattenCertificateManagerTrustConfigTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading TrustConfig: %s", err)
	}
	if err := d.Set("effective_labels", flattenCertificateManagerTrustConfigEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading TrustConfig: %s", err)
	}

	return nil
}

func resourceCertificateManagerTrustConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TrustConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandCertificateManagerTrustConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	trustStoresProp, err := expandCertificateManagerTrustConfigTrustStores(d.Get("trust_stores"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("trust_stores"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, trustStoresProp)) {
		obj["trustStores"] = trustStoresProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CertificateManagerBasePath}}projects/{{project}}/locations/{{location}}/trustConfigs/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating TrustConfig %q: %#v", d.Id(), obj)
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": "*"})
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
		return fmt.Errorf("Error updating TrustConfig %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating TrustConfig %q: %#v", d.Id(), res)
	}

	err = CertificateManagerOperationWaitTime(
		config, res, project, "Updating TrustConfig", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceCertificateManagerTrustConfigRead(d, meta)
}

func resourceCertificateManagerTrustConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TrustConfig: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{CertificateManagerBasePath}}projects/{{project}}/locations/{{location}}/trustConfigs/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting TrustConfig %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "TrustConfig")
	}

	err = CertificateManagerOperationWaitTime(
		config, res, project, "Deleting TrustConfig", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting TrustConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceCertificateManagerTrustConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/trustConfigs/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/trustConfigs/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenCertificateManagerTrustConfigCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerTrustConfigUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerTrustConfigLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenCertificateManagerTrustConfigDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerTrustConfigTrustStores(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"trust_anchors":    flattenCertificateManagerTrustConfigTrustStoresTrustAnchors(original["trustAnchors"], d, config),
			"intermediate_cas": flattenCertificateManagerTrustConfigTrustStoresIntermediateCas(original["intermediateCas"], d, config),
		})
	}
	return transformed
}
func flattenCertificateManagerTrustConfigTrustStoresTrustAnchors(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"pem_certificate": flattenCertificateManagerTrustConfigTrustStoresTrustAnchorsPemCertificate(original["pemCertificate"], d, config),
		})
	}
	return transformed
}
func flattenCertificateManagerTrustConfigTrustStoresTrustAnchorsPemCertificate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerTrustConfigTrustStoresIntermediateCas(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"pem_certificate": flattenCertificateManagerTrustConfigTrustStoresIntermediateCasPemCertificate(original["pemCertificate"], d, config),
		})
	}
	return transformed
}
func flattenCertificateManagerTrustConfigTrustStoresIntermediateCasPemCertificate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCertificateManagerTrustConfigTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenCertificateManagerTrustConfigEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandCertificateManagerTrustConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerTrustConfigTrustStores(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedTrustAnchors, err := expandCertificateManagerTrustConfigTrustStoresTrustAnchors(original["trust_anchors"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTrustAnchors); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["trustAnchors"] = transformedTrustAnchors
		}

		transformedIntermediateCas, err := expandCertificateManagerTrustConfigTrustStoresIntermediateCas(original["intermediate_cas"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIntermediateCas); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["intermediateCas"] = transformedIntermediateCas
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCertificateManagerTrustConfigTrustStoresTrustAnchors(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPemCertificate, err := expandCertificateManagerTrustConfigTrustStoresTrustAnchorsPemCertificate(original["pem_certificate"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPemCertificate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["pemCertificate"] = transformedPemCertificate
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCertificateManagerTrustConfigTrustStoresTrustAnchorsPemCertificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerTrustConfigTrustStoresIntermediateCas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPemCertificate, err := expandCertificateManagerTrustConfigTrustStoresIntermediateCasPemCertificate(original["pem_certificate"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPemCertificate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["pemCertificate"] = transformedPemCertificate
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCertificateManagerTrustConfigTrustStoresIntermediateCasPemCertificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerTrustConfigEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
