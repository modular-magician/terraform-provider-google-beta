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

package apigee

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceApigeeKeystoresAliasesSelfSignedCert() *schema.Resource {
	return &schema.Resource{
		Create: resourceApigeeKeystoresAliasesSelfSignedCertCreate,
		Read:   resourceApigeeKeystoresAliasesSelfSignedCertRead,
		Delete: resourceApigeeKeystoresAliasesSelfSignedCertDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApigeeKeystoresAliasesSelfSignedCertImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"alias": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Alias for the key/certificate pair. Values must match the regular expression [\w\s-.]{1,255}.
This must be provided for all formats except selfsignedcert; self-signed certs may specify the alias in either
this parameter or the JSON body.`,
			},
			"environment": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The Apigee environment name`,
			},
			"keystore": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The Apigee keystore name associated in an Apigee environment`,
			},
			"org_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The Apigee Organization name associated with the Apigee environment`,
			},
			"sig_alg": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Signature algorithm to generate private key. Valid values are SHA512withRSA, SHA384withRSA, and SHA256withRSA`,
			},
			"subject": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: `Subject details.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"common_name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `Common name of the organization. Maximum length is 64 characters.`,
						},
						"country_code": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `Two-letter country code. Example, IN for India, US for United States of America.`,
						},
						"email": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `Email address. Max 255 characters.`,
						},
						"locality": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `City or town name. Maximum length is 128 characters.`,
						},
						"org": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `Organization name. Maximum length is 64 characters.`,
						},
						"org_unit": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `Organization team name. Maximum length is 64 characters.`,
						},
						"state": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `State or district name. Maximum length is 128 characters.`,
						},
					},
				},
			},
			"cert_validity_in_days": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: `Validity duration of certificate, in days. Accepts positive non-zero value. Defaults to 365.`,
			},
			"key_size": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Key size. Default and maximum value is 2048 bits.`,
			},
			"subject_alternative_dns_names": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `List of alternative host names. Maximum length is 255 characters for each value.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subject_alternative_name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `Subject Alternative Name`,
						},
					},
				},
			},
			"certs_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Chain of certificates under this alias.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_info": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `List of all properties in the object.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"basic_constraints": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `X.509 basic constraints extension.`,
									},
									"expiry_date": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `X.509 notAfter validity period in milliseconds since epoch.`,
									},
									"is_valid": {
										Type:     schema.TypeString,
										Computed: true,
										Description: `Flag that specifies whether the certificate is valid.
Flag is set to Yes if the certificate is valid, No if expired, or Not yet if not yet valid.`,
									},
									"issuer": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `X.509 issuer.`,
									},
									"public_key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `Public key component of the X.509 subject public key info.`,
									},
									"serial_number": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `X.509 serial number.`,
									},
									"sig_alg_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `X.509 signatureAlgorithm.`,
									},
									"subject": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `X.509 subject.`,
									},
									"subject_alternative_names": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: `X.509 subject alternative names (SANs) extension.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"valid_from": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `X.509 notBefore validity period in milliseconds since epoch.`,
									},
									"version": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: `X.509 version.`,
									},
								},
							},
						},
					},
				},
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Optional.Type of Alias`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApigeeKeystoresAliasesSelfSignedCertCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	aliasProp, err := expandApigeeKeystoresAliasesSelfSignedCertAlias(d.Get("alias"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("alias"); !tpgresource.IsEmptyValue(reflect.ValueOf(aliasProp)) && (ok || !reflect.DeepEqual(v, aliasProp)) {
		obj["alias"] = aliasProp
	}
	subjectAlternativeDnsNamesProp, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames(d.Get("subject_alternative_dns_names"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("subject_alternative_dns_names"); !tpgresource.IsEmptyValue(reflect.ValueOf(subjectAlternativeDnsNamesProp)) && (ok || !reflect.DeepEqual(v, subjectAlternativeDnsNamesProp)) {
		obj["subjectAlternativeDnsNames"] = subjectAlternativeDnsNamesProp
	}
	keySizeProp, err := expandApigeeKeystoresAliasesSelfSignedCertKeySize(d.Get("key_size"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("key_size"); !tpgresource.IsEmptyValue(reflect.ValueOf(keySizeProp)) && (ok || !reflect.DeepEqual(v, keySizeProp)) {
		obj["keySize"] = keySizeProp
	}
	sigAlgProp, err := expandApigeeKeystoresAliasesSelfSignedCertSigAlg(d.Get("sig_alg"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("sig_alg"); !tpgresource.IsEmptyValue(reflect.ValueOf(sigAlgProp)) && (ok || !reflect.DeepEqual(v, sigAlgProp)) {
		obj["sigAlg"] = sigAlgProp
	}
	subjectProp, err := expandApigeeKeystoresAliasesSelfSignedCertSubject(d.Get("subject"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("subject"); !tpgresource.IsEmptyValue(reflect.ValueOf(subjectProp)) && (ok || !reflect.DeepEqual(v, subjectProp)) {
		obj["subject"] = subjectProp
	}
	certValidityInDaysProp, err := expandApigeeKeystoresAliasesSelfSignedCertCertValidityInDays(d.Get("cert_validity_in_days"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cert_validity_in_days"); !tpgresource.IsEmptyValue(reflect.ValueOf(certValidityInDaysProp)) && (ok || !reflect.DeepEqual(v, certValidityInDaysProp)) {
		obj["certValidityInDays"] = certValidityInDaysProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases?alias={{alias}}&format=selfsignedcert")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new KeystoresAliasesSelfSignedCert: %#v", obj)
	billingProject := ""

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
		return fmt.Errorf("Error creating KeystoresAliasesSelfSignedCert: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating KeystoresAliasesSelfSignedCert %q: %#v", d.Id(), res)

	return resourceApigeeKeystoresAliasesSelfSignedCertRead(d, meta)
}

func resourceApigeeKeystoresAliasesSelfSignedCertRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}")
	if err != nil {
		return err
	}

	billingProject := ""

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ApigeeKeystoresAliasesSelfSignedCert %q", d.Id()))
	}

	if err := d.Set("certs_info", flattenApigeeKeystoresAliasesSelfSignedCertCertsInfo(res["certsInfo"], d, config)); err != nil {
		return fmt.Errorf("Error reading KeystoresAliasesSelfSignedCert: %s", err)
	}
	if err := d.Set("type", flattenApigeeKeystoresAliasesSelfSignedCertType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading KeystoresAliasesSelfSignedCert: %s", err)
	}
	if err := d.Set("alias", flattenApigeeKeystoresAliasesSelfSignedCertAlias(res["alias"], d, config)); err != nil {
		return fmt.Errorf("Error reading KeystoresAliasesSelfSignedCert: %s", err)
	}
	if err := d.Set("subject_alternative_dns_names", flattenApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames(res["subjectAlternativeDnsNames"], d, config)); err != nil {
		return fmt.Errorf("Error reading KeystoresAliasesSelfSignedCert: %s", err)
	}

	return nil
}

func resourceApigeeKeystoresAliasesSelfSignedCertDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	log.Printf("[DEBUG] Deleting KeystoresAliasesSelfSignedCert %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "KeystoresAliasesSelfSignedCert")
	}

	log.Printf("[DEBUG] Finished deleting KeystoresAliasesSelfSignedCert %q: %#v", d.Id(), res)
	return nil
}

func resourceApigeeKeystoresAliasesSelfSignedCertImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats cannot import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{
		"organizations/(?P<org_id>[^/]+)/environments/(?P<environment>[^/]+)/keystores/(?P<keystore>[^/]+)/aliases/(?P<alias>[^/]+)",
		"(?P<org_id>[^/]+)/(?P<environment>[^/]+)/(?P<keystore>[^/]+)/(?P<alias>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApigeeKeystoresAliasesSelfSignedCertCertsInfo(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["cert_info"] =
		flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfo(original["certInfo"], d, config)
	return []interface{}{transformed}
}
func flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfo(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"version":                   flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoVersion(original["version"], d, config),
			"subject":                   flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoSubject(original["subject"], d, config),
			"issuer":                    flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoIssuer(original["issuer"], d, config),
			"expiry_date":               flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoExpiryDate(original["expiryDate"], d, config),
			"valid_from":                flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoValidFrom(original["validFrom"], d, config),
			"is_valid":                  flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoIsValid(original["isValid"], d, config),
			"subject_alternative_names": flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoSubjectAlternativeNames(original["subjectAlternativeNames"], d, config),
			"sig_alg_name":              flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoSigAlgName(original["sigAlgName"], d, config),
			"public_key":                flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoPublicKey(original["publicKey"], d, config),
			"basic_constraints":         flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoBasicConstraints(original["basicConstraints"], d, config),
			"serial_number":             flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoSerialNumber(original["serialNumber"], d, config),
		})
	}
	return transformed
}
func flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoSubject(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoIssuer(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoExpiryDate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoValidFrom(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoIsValid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoSubjectAlternativeNames(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoSigAlgName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoPublicKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoBasicConstraints(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeKeystoresAliasesSelfSignedCertCertsInfoCertInfoSerialNumber(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeKeystoresAliasesSelfSignedCertType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeKeystoresAliasesSelfSignedCertAlias(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["subject_alternative_name"] =
		flattenApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesSubjectAlternativeName(original["subjectAlternativeName"], d, config)
	return []interface{}{transformed}
}
func flattenApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesSubjectAlternativeName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandApigeeKeystoresAliasesSelfSignedCertAlias(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSubjectAlternativeName, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesSubjectAlternativeName(original["subject_alternative_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubjectAlternativeName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subjectAlternativeName"] = transformedSubjectAlternativeName
	}

	return transformed, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesSubjectAlternativeName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertKeySize(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSigAlg(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCountryCode, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectCountryCode(original["country_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCountryCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["countryCode"] = transformedCountryCode
	}

	transformedState, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["state"] = transformedState
	}

	transformedLocality, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectLocality(original["locality"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocality); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["locality"] = transformedLocality
	}

	transformedOrg, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectOrg(original["org"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOrg); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["org"] = transformedOrg
	}

	transformedOrgUnit, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectOrgUnit(original["org_unit"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOrgUnit); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["orgUnit"] = transformedOrgUnit
	}

	transformedCommonName, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectCommonName(original["common_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCommonName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["commonName"] = transformedCommonName
	}

	transformedEmail, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectEmail(original["email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["email"] = transformedEmail
	}

	return transformed, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectCountryCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectLocality(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectOrg(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectOrgUnit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectCommonName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertCertValidityInDays(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
