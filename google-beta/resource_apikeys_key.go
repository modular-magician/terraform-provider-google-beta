// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	apikeys "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apikeys/beta"
)

func resourceApikeysKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceApikeysKeyCreate,
		Read:   resourceApikeysKeyRead,
		Update: resourceApikeysKeyUpdate,
		Delete: resourceApikeysKeyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApikeysKeyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
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

			"restrictions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        ApikeysKeyRestrictionsSchema(),
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"key_string": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},
		},
	}
}

func ApikeysKeyRestrictionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"android_key_restrictions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        ApikeysKeyRestrictionsAndroidKeyRestrictionsSchema(),
			},

			"api_targets": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        ApikeysKeyRestrictionsApiTargetsSchema(),
			},

			"browser_key_restrictions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        ApikeysKeyRestrictionsBrowserKeyRestrictionsSchema(),
			},

			"ios_key_restrictions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        ApikeysKeyRestrictionsIosKeyRestrictionsSchema(),
			},

			"server_key_restrictions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        ApikeysKeyRestrictionsServerKeyRestrictionsSchema(),
			},
		},
	}
}

func ApikeysKeyRestrictionsAndroidKeyRestrictionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allowed_applications": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        ApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsSchema(),
			},
		},
	}
}

func ApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"package_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"sha1_fingerprint": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},
		},
	}
}

func ApikeysKeyRestrictionsApiTargetsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"methods": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"service": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},
		},
	}
}

func ApikeysKeyRestrictionsBrowserKeyRestrictionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allowed_referrers": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ApikeysKeyRestrictionsIosKeyRestrictionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allowed_bundle_ids": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ApikeysKeyRestrictionsServerKeyRestrictionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allowed_ips": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceApikeysKeyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &apikeys.Key{
		DisplayName:  dcl.StringOrNil(d.Get("display_name").(string)),
		Project:      dcl.String(project),
		Restrictions: expandApikeysKeyRestrictions(d.Get("restrictions")),
	}

	id, err := replaceVarsForId(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	createDirective := CreateDirective
	res, err := config.clientApikeysDCL.ApplyKey(context.Background(), obj, createDirective...)
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating Key: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Key %q: %#v", d.Id(), res)

	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	// Id has a server-generated value, set again after creation
	id, err = replaceVarsForId(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return resourceApikeysKeyRead(d, meta)
}

func resourceApikeysKeyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &apikeys.Key{
		DisplayName:  dcl.StringOrNil(d.Get("display_name").(string)),
		Project:      dcl.String(project),
		Restrictions: expandApikeysKeyRestrictions(d.Get("restrictions")),
		Name:         dcl.StringOrNil(d.Get("name").(string)),
	}

	res, err := config.clientApikeysDCL.GetKey(context.Background(), obj)
	if err != nil {
		// Resource not found
		d.SetId("")
		return err
	}

	if err = d.Set("display_name", res.DisplayName); err != nil {
		return fmt.Errorf("error setting display_name in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("restrictions", flattenApikeysKeyRestrictions(res.Restrictions)); err != nil {
		return fmt.Errorf("error setting restrictions in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("etag", res.Etag); err != nil {
		return fmt.Errorf("error setting etag in state: %s", err)
	}
	if err = d.Set("key_string", res.KeyString); err != nil {
		return fmt.Errorf("error setting key_string in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("uid", res.Uid); err != nil {
		return fmt.Errorf("error setting uid in state: %s", err)
	}
	if err = d.Set("update_time", res.UpdateTime); err != nil {
		return fmt.Errorf("error setting update_time in state: %s", err)
	}

	return nil
}
func resourceApikeysKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &apikeys.Key{
		DisplayName:  dcl.StringOrNil(d.Get("display_name").(string)),
		Project:      dcl.String(project),
		Restrictions: expandApikeysKeyRestrictions(d.Get("restrictions")),
		Name:         dcl.StringOrNil(d.Get("name").(string)),
	}
	directive := UpdateDirective
	res, err := config.clientApikeysDCL.ApplyKey(context.Background(), obj, directive...)
	if err != nil {
		return fmt.Errorf("Error updating Key: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Key %q: %#v", d.Id(), res)

	return resourceApikeysKeyRead(d, meta)
}

func resourceApikeysKeyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &apikeys.Key{
		DisplayName:  dcl.StringOrNil(d.Get("display_name").(string)),
		Project:      dcl.String(project),
		Restrictions: expandApikeysKeyRestrictions(d.Get("restrictions")),
		Name:         dcl.StringOrNil(d.Get("name").(string)),
	}

	log.Printf("[DEBUG] Deleting Key %q", d.Id())
	if err := config.clientApikeysDCL.DeleteKey(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting Key: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting Key %q", d.Id())
	return nil
}

func resourceApikeysKeyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandApikeysKeyRestrictions(o interface{}) *apikeys.KeyRestrictions {
	if o == nil {
		return nil
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return nil
	}
	obj := objArr[0].(map[string]interface{})
	return &apikeys.KeyRestrictions{
		AndroidKeyRestrictions: expandApikeysKeyRestrictionsAndroidKeyRestrictions(obj["android_key_restrictions"]),
		ApiTargets:             expandApikeysKeyRestrictionsApiTargetsArray(obj["api_targets"]),
		BrowserKeyRestrictions: expandApikeysKeyRestrictionsBrowserKeyRestrictions(obj["browser_key_restrictions"]),
		IosKeyRestrictions:     expandApikeysKeyRestrictionsIosKeyRestrictions(obj["ios_key_restrictions"]),
		ServerKeyRestrictions:  expandApikeysKeyRestrictionsServerKeyRestrictions(obj["server_key_restrictions"]),
	}
}

func flattenApikeysKeyRestrictions(obj *apikeys.KeyRestrictions) interface{} {
	if obj == nil {
		return nil
	}
	transformed := map[string]interface{}{
		"android_key_restrictions": flattenApikeysKeyRestrictionsAndroidKeyRestrictions(obj.AndroidKeyRestrictions),
		"api_targets":              flattenApikeysKeyRestrictionsApiTargetsArray(obj.ApiTargets),
		"browser_key_restrictions": flattenApikeysKeyRestrictionsBrowserKeyRestrictions(obj.BrowserKeyRestrictions),
		"ios_key_restrictions":     flattenApikeysKeyRestrictionsIosKeyRestrictions(obj.IosKeyRestrictions),
		"server_key_restrictions":  flattenApikeysKeyRestrictionsServerKeyRestrictions(obj.ServerKeyRestrictions),
	}

	return []interface{}{transformed}

}

func expandApikeysKeyRestrictionsAndroidKeyRestrictions(o interface{}) *apikeys.KeyRestrictionsAndroidKeyRestrictions {
	if o == nil {
		return nil
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return nil
	}
	obj := objArr[0].(map[string]interface{})
	return &apikeys.KeyRestrictionsAndroidKeyRestrictions{
		AllowedApplications: expandApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsArray(obj["allowed_applications"]),
	}
}

func flattenApikeysKeyRestrictionsAndroidKeyRestrictions(obj *apikeys.KeyRestrictionsAndroidKeyRestrictions) interface{} {
	if obj == nil {
		return nil
	}
	transformed := map[string]interface{}{
		"allowed_applications": flattenApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsArray(obj.AllowedApplications),
	}

	return []interface{}{transformed}

}
func expandApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsArray(o interface{}) []apikeys.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	if o == nil {
		return nil
	}

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]apikeys.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications, 0, len(objs))
	for _, item := range objs {
		i := expandApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(item)
		items = append(items, *i)
	}

	return items
}

func expandApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(o interface{}) *apikeys.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	if o == nil {
		return nil
	}

	obj := o.(map[string]interface{})
	return &apikeys.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications{
		PackageName:     dcl.StringOrNil(obj["package_name"].(string)),
		Sha1Fingerprint: dcl.StringOrNil(obj["sha1_fingerprint"].(string)),
	}
}

func flattenApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsArray(objs []apikeys.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(&item)
		items = append(items, i)
	}

	return items
}

func flattenApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(obj *apikeys.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications) interface{} {
	if obj == nil {
		return nil
	}
	transformed := map[string]interface{}{
		"package_name":     obj.PackageName,
		"sha1_fingerprint": obj.Sha1Fingerprint,
	}

	return transformed

}
func expandApikeysKeyRestrictionsApiTargetsArray(o interface{}) []apikeys.KeyRestrictionsApiTargets {
	if o == nil {
		return nil
	}

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]apikeys.KeyRestrictionsApiTargets, 0, len(objs))
	for _, item := range objs {
		i := expandApikeysKeyRestrictionsApiTargets(item)
		items = append(items, *i)
	}

	return items
}

func expandApikeysKeyRestrictionsApiTargets(o interface{}) *apikeys.KeyRestrictionsApiTargets {
	if o == nil {
		return nil
	}

	obj := o.(map[string]interface{})
	return &apikeys.KeyRestrictionsApiTargets{
		Methods: expandStringArray(obj["methods"]),
		Service: dcl.StringOrNil(obj["service"].(string)),
	}
}

func flattenApikeysKeyRestrictionsApiTargetsArray(objs []apikeys.KeyRestrictionsApiTargets) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenApikeysKeyRestrictionsApiTargets(&item)
		items = append(items, i)
	}

	return items
}

func flattenApikeysKeyRestrictionsApiTargets(obj *apikeys.KeyRestrictionsApiTargets) interface{} {
	if obj == nil {
		return nil
	}
	transformed := map[string]interface{}{
		"methods": obj.Methods,
		"service": obj.Service,
	}

	return transformed

}

func expandApikeysKeyRestrictionsBrowserKeyRestrictions(o interface{}) *apikeys.KeyRestrictionsBrowserKeyRestrictions {
	if o == nil {
		return nil
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return nil
	}
	obj := objArr[0].(map[string]interface{})
	return &apikeys.KeyRestrictionsBrowserKeyRestrictions{
		AllowedReferrers: expandStringArray(obj["allowed_referrers"]),
	}
}

func flattenApikeysKeyRestrictionsBrowserKeyRestrictions(obj *apikeys.KeyRestrictionsBrowserKeyRestrictions) interface{} {
	if obj == nil {
		return nil
	}
	transformed := map[string]interface{}{
		"allowed_referrers": obj.AllowedReferrers,
	}

	return []interface{}{transformed}

}

func expandApikeysKeyRestrictionsIosKeyRestrictions(o interface{}) *apikeys.KeyRestrictionsIosKeyRestrictions {
	if o == nil {
		return nil
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return nil
	}
	obj := objArr[0].(map[string]interface{})
	return &apikeys.KeyRestrictionsIosKeyRestrictions{
		AllowedBundleIds: expandStringArray(obj["allowed_bundle_ids"]),
	}
}

func flattenApikeysKeyRestrictionsIosKeyRestrictions(obj *apikeys.KeyRestrictionsIosKeyRestrictions) interface{} {
	if obj == nil {
		return nil
	}
	transformed := map[string]interface{}{
		"allowed_bundle_ids": obj.AllowedBundleIds,
	}

	return []interface{}{transformed}

}

func expandApikeysKeyRestrictionsServerKeyRestrictions(o interface{}) *apikeys.KeyRestrictionsServerKeyRestrictions {
	if o == nil {
		return nil
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return nil
	}
	obj := objArr[0].(map[string]interface{})
	return &apikeys.KeyRestrictionsServerKeyRestrictions{
		AllowedIps: expandStringArray(obj["allowed_ips"]),
	}
}

func flattenApikeysKeyRestrictionsServerKeyRestrictions(obj *apikeys.KeyRestrictionsServerKeyRestrictions) interface{} {
	if obj == nil {
		return nil
	}
	transformed := map[string]interface{}{
		"allowed_ips": obj.AllowedIps,
	}

	return []interface{}{transformed}

}
