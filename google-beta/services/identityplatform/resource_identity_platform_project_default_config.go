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

package identityplatform

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

func ResourceIdentityPlatformProjectDefaultConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceIdentityPlatformProjectDefaultConfigCreate,
		Read:   resourceIdentityPlatformProjectDefaultConfigRead,
		Update: resourceIdentityPlatformProjectDefaultConfigUpdate,
		Delete: resourceIdentityPlatformProjectDefaultConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIdentityPlatformProjectDefaultConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		DeprecationMessage: "`google_identity_platform_config` is deprecated and will be removed in the next major release of the provider. Use the `google_identity_platform_config` resource instead. It contains a more comprehensive list of fields, and was created before `google_identity_platform_project_default_config` was added.",

		Schema: map[string]*schema.Schema{
			"sign_in": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Configuration related to local sign in methods.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_duplicate_emails": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Whether to allow more than one account to have the same email.`,
						},
						"anonymous": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Configuration options related to authenticating an anonymous user.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: `Whether anonymous user auth is enabled for the project or not.`,
									},
								},
							},
						},
						"email": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Configuration options related to authenticating a user by their email address.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: `Whether email auth is enabled for the project or not.`,
									},
									"password_required": {
										Type:     schema.TypeBool,
										Optional: true,
										Description: `Whether a password is required for email auth or not. If true, both an email and
password must be provided to sign in. If false, a user may sign in via either
email/password or email link.`,
									},
								},
							},
						},
						"phone_number": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Configuration options related to authenticated a user by their phone number.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: `Whether phone number auth is enabled for the project or not.`,
									},
									"test_phone_numbers": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: `A map of <test phone number, fake code> that can be used for phone auth testing.`,
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"hash_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `Output only. Hash config information.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"algorithm": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `Different password hash algorithms used in Identity Toolkit.`,
									},
									"memory_cost": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: `Memory cost for hash calculation. Used by scrypt and other similar password derivation algorithms. See https://tools.ietf.org/html/rfc7914 for explanation of field.`,
									},
									"rounds": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: `How many rounds for hash calculation. Used by scrypt and other similar password derivation algorithms.`,
									},
									"salt_separator": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `Non-printable character to be inserted between the salt and plain text password in base64.`,
									},
									"signer_key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `Signer key in base64.`,
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the Config resource. Example: "projects/my-awesome-project/config"`,
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

func resourceIdentityPlatformProjectDefaultConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	signInProp, err := expandIdentityPlatformProjectDefaultConfigSignIn(d.Get("sign_in"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("sign_in"); !tpgresource.IsEmptyValue(reflect.ValueOf(signInProp)) && (ok || !reflect.DeepEqual(v, signInProp)) {
		obj["signIn"] = signInProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/config")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ProjectDefaultConfig: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectDefaultConfig: %s", err)
	}
	billingProject = project

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
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating ProjectDefaultConfig: %s", err)
	}
	if err := d.Set("name", flattenIdentityPlatformProjectDefaultConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ProjectDefaultConfig %q: %#v", d.Id(), res)

	return resourceIdentityPlatformProjectDefaultConfigRead(d, meta)
}

func resourceIdentityPlatformProjectDefaultConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/config")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectDefaultConfig: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IdentityPlatformProjectDefaultConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ProjectDefaultConfig: %s", err)
	}

	if err := d.Set("name", flattenIdentityPlatformProjectDefaultConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectDefaultConfig: %s", err)
	}
	if err := d.Set("sign_in", flattenIdentityPlatformProjectDefaultConfigSignIn(res["signIn"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectDefaultConfig: %s", err)
	}

	return nil
}

func resourceIdentityPlatformProjectDefaultConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectDefaultConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	signInProp, err := expandIdentityPlatformProjectDefaultConfigSignIn(d.Get("sign_in"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("sign_in"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, signInProp)) {
		obj["signIn"] = signInProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/config")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ProjectDefaultConfig %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("sign_in") {
		updateMask = append(updateMask, "signIn")
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
		})

		if err != nil {
			return fmt.Errorf("Error updating ProjectDefaultConfig %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ProjectDefaultConfig %q: %#v", d.Id(), res)
		}

	}

	return resourceIdentityPlatformProjectDefaultConfigRead(d, meta)
}

func resourceIdentityPlatformProjectDefaultConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectDefaultConfig: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/config")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ProjectDefaultConfig %q", d.Id())

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
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "ProjectDefaultConfig")
	}

	log.Printf("[DEBUG] Finished deleting ProjectDefaultConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceIdentityPlatformProjectDefaultConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/config/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIdentityPlatformProjectDefaultConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformProjectDefaultConfigSignIn(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["email"] =
		flattenIdentityPlatformProjectDefaultConfigSignInEmail(original["email"], d, config)
	transformed["phone_number"] =
		flattenIdentityPlatformProjectDefaultConfigSignInPhoneNumber(original["phoneNumber"], d, config)
	transformed["anonymous"] =
		flattenIdentityPlatformProjectDefaultConfigSignInAnonymous(original["anonymous"], d, config)
	transformed["allow_duplicate_emails"] =
		flattenIdentityPlatformProjectDefaultConfigSignInAllowDuplicateEmails(original["allowDuplicateEmails"], d, config)
	transformed["hash_config"] =
		flattenIdentityPlatformProjectDefaultConfigSignInHashConfig(original["hashConfig"], d, config)
	return []interface{}{transformed}
}
func flattenIdentityPlatformProjectDefaultConfigSignInEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enabled"] =
		flattenIdentityPlatformProjectDefaultConfigSignInEmailEnabled(original["enabled"], d, config)
	transformed["password_required"] =
		flattenIdentityPlatformProjectDefaultConfigSignInEmailPasswordRequired(original["passwordRequired"], d, config)
	return []interface{}{transformed}
}
func flattenIdentityPlatformProjectDefaultConfigSignInEmailEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformProjectDefaultConfigSignInEmailPasswordRequired(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformProjectDefaultConfigSignInPhoneNumber(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enabled"] =
		flattenIdentityPlatformProjectDefaultConfigSignInPhoneNumberEnabled(original["enabled"], d, config)
	transformed["test_phone_numbers"] =
		flattenIdentityPlatformProjectDefaultConfigSignInPhoneNumberTestPhoneNumbers(original["testPhoneNumbers"], d, config)
	return []interface{}{transformed}
}
func flattenIdentityPlatformProjectDefaultConfigSignInPhoneNumberEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformProjectDefaultConfigSignInPhoneNumberTestPhoneNumbers(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformProjectDefaultConfigSignInAnonymous(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enabled"] =
		flattenIdentityPlatformProjectDefaultConfigSignInAnonymousEnabled(original["enabled"], d, config)
	return []interface{}{transformed}
}
func flattenIdentityPlatformProjectDefaultConfigSignInAnonymousEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformProjectDefaultConfigSignInAllowDuplicateEmails(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformProjectDefaultConfigSignInHashConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["algorithm"] =
		flattenIdentityPlatformProjectDefaultConfigSignInHashConfigAlgorithm(original["algorithm"], d, config)
	transformed["signer_key"] =
		flattenIdentityPlatformProjectDefaultConfigSignInHashConfigSignerKey(original["signerKey"], d, config)
	transformed["salt_separator"] =
		flattenIdentityPlatformProjectDefaultConfigSignInHashConfigSaltSeparator(original["saltSeparator"], d, config)
	transformed["rounds"] =
		flattenIdentityPlatformProjectDefaultConfigSignInHashConfigRounds(original["rounds"], d, config)
	transformed["memory_cost"] =
		flattenIdentityPlatformProjectDefaultConfigSignInHashConfigMemoryCost(original["memoryCost"], d, config)
	return []interface{}{transformed}
}
func flattenIdentityPlatformProjectDefaultConfigSignInHashConfigAlgorithm(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformProjectDefaultConfigSignInHashConfigSignerKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformProjectDefaultConfigSignInHashConfigSaltSeparator(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformProjectDefaultConfigSignInHashConfigRounds(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenIdentityPlatformProjectDefaultConfigSignInHashConfigMemoryCost(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func expandIdentityPlatformProjectDefaultConfigSignIn(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEmail, err := expandIdentityPlatformProjectDefaultConfigSignInEmail(original["email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["email"] = transformedEmail
	}

	transformedPhoneNumber, err := expandIdentityPlatformProjectDefaultConfigSignInPhoneNumber(original["phone_number"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPhoneNumber); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["phoneNumber"] = transformedPhoneNumber
	}

	transformedAnonymous, err := expandIdentityPlatformProjectDefaultConfigSignInAnonymous(original["anonymous"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAnonymous); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["anonymous"] = transformedAnonymous
	}

	transformedAllowDuplicateEmails, err := expandIdentityPlatformProjectDefaultConfigSignInAllowDuplicateEmails(original["allow_duplicate_emails"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowDuplicateEmails); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowDuplicateEmails"] = transformedAllowDuplicateEmails
	}

	transformedHashConfig, err := expandIdentityPlatformProjectDefaultConfigSignInHashConfig(original["hash_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHashConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hashConfig"] = transformedHashConfig
	}

	return transformed, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandIdentityPlatformProjectDefaultConfigSignInEmailEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedPasswordRequired, err := expandIdentityPlatformProjectDefaultConfigSignInEmailPasswordRequired(original["password_required"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPasswordRequired); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["passwordRequired"] = transformedPasswordRequired
	}

	return transformed, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInEmailEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInEmailPasswordRequired(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInPhoneNumber(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandIdentityPlatformProjectDefaultConfigSignInPhoneNumberEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedTestPhoneNumbers, err := expandIdentityPlatformProjectDefaultConfigSignInPhoneNumberTestPhoneNumbers(original["test_phone_numbers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTestPhoneNumbers); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["testPhoneNumbers"] = transformedTestPhoneNumbers
	}

	return transformed, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInPhoneNumberEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInPhoneNumberTestPhoneNumbers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInAnonymous(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandIdentityPlatformProjectDefaultConfigSignInAnonymousEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	return transformed, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInAnonymousEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInAllowDuplicateEmails(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInHashConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAlgorithm, err := expandIdentityPlatformProjectDefaultConfigSignInHashConfigAlgorithm(original["algorithm"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAlgorithm); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["algorithm"] = transformedAlgorithm
	}

	transformedSignerKey, err := expandIdentityPlatformProjectDefaultConfigSignInHashConfigSignerKey(original["signer_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSignerKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["signerKey"] = transformedSignerKey
	}

	transformedSaltSeparator, err := expandIdentityPlatformProjectDefaultConfigSignInHashConfigSaltSeparator(original["salt_separator"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSaltSeparator); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["saltSeparator"] = transformedSaltSeparator
	}

	transformedRounds, err := expandIdentityPlatformProjectDefaultConfigSignInHashConfigRounds(original["rounds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRounds); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rounds"] = transformedRounds
	}

	transformedMemoryCost, err := expandIdentityPlatformProjectDefaultConfigSignInHashConfigMemoryCost(original["memory_cost"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMemoryCost); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["memoryCost"] = transformedMemoryCost
	}

	return transformed, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInHashConfigAlgorithm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInHashConfigSignerKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInHashConfigSaltSeparator(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInHashConfigRounds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInHashConfigMemoryCost(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
