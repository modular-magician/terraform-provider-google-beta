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

func ResourceIdentityPlatformTenant() *schema.Resource {
	return &schema.Resource{
		Create: resourceIdentityPlatformTenantCreate,
		Read:   resourceIdentityPlatformTenantRead,
		Update: resourceIdentityPlatformTenantUpdate,
		Delete: resourceIdentityPlatformTenantDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIdentityPlatformTenantImport,
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
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Human friendly display name of the tenant.`,
			},
			"allow_password_signup": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Whether to allow email/password user authentication.`,
			},
			"disable_auth": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Whether authentication is disabled for the tenant. If true, the users under
the disabled tenant are not allowed to sign-in. Admins of the disabled tenant
are not able to manage its users.`,
			},
			"enable_email_link_signin": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Whether to enable email link user authentication.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the tenant that is generated by the server`,
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

func resourceIdentityPlatformTenantCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandIdentityPlatformTenantDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	allowPasswordSignupProp, err := expandIdentityPlatformTenantAllowPasswordSignup(d.Get("allow_password_signup"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("allow_password_signup"); !tpgresource.IsEmptyValue(reflect.ValueOf(allowPasswordSignupProp)) && (ok || !reflect.DeepEqual(v, allowPasswordSignupProp)) {
		obj["allowPasswordSignup"] = allowPasswordSignupProp
	}
	enableEmailLinkSigninProp, err := expandIdentityPlatformTenantEnableEmailLinkSignin(d.Get("enable_email_link_signin"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_email_link_signin"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableEmailLinkSigninProp)) && (ok || !reflect.DeepEqual(v, enableEmailLinkSigninProp)) {
		obj["enableEmailLinkSignin"] = enableEmailLinkSigninProp
	}
	disableAuthProp, err := expandIdentityPlatformTenantDisableAuth(d.Get("disable_auth"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disable_auth"); !tpgresource.IsEmptyValue(reflect.ValueOf(disableAuthProp)) && (ok || !reflect.DeepEqual(v, disableAuthProp)) {
		obj["disableAuth"] = disableAuthProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Tenant: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Tenant: %s", err)
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
		return fmt.Errorf("Error creating Tenant: %s", err)
	}
	if err := d.Set("name", flattenIdentityPlatformTenantName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/tenants/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
	}
	if err := d.Set("name", tpgresource.GetResourceNameFromSelfLink(name.(string))); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	// Store the ID now that we have set the computed name
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/tenants/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Tenant %q: %#v", d.Id(), res)

	return resourceIdentityPlatformTenantRead(d, meta)
}

func resourceIdentityPlatformTenantRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Tenant: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IdentityPlatformTenant %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Tenant: %s", err)
	}

	if err := d.Set("name", flattenIdentityPlatformTenantName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tenant: %s", err)
	}
	if err := d.Set("display_name", flattenIdentityPlatformTenantDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tenant: %s", err)
	}
	if err := d.Set("allow_password_signup", flattenIdentityPlatformTenantAllowPasswordSignup(res["allowPasswordSignup"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tenant: %s", err)
	}
	if err := d.Set("enable_email_link_signin", flattenIdentityPlatformTenantEnableEmailLinkSignin(res["enableEmailLinkSignin"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tenant: %s", err)
	}
	if err := d.Set("disable_auth", flattenIdentityPlatformTenantDisableAuth(res["disableAuth"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tenant: %s", err)
	}

	return nil
}

func resourceIdentityPlatformTenantUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Tenant: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandIdentityPlatformTenantDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	allowPasswordSignupProp, err := expandIdentityPlatformTenantAllowPasswordSignup(d.Get("allow_password_signup"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("allow_password_signup"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, allowPasswordSignupProp)) {
		obj["allowPasswordSignup"] = allowPasswordSignupProp
	}
	enableEmailLinkSigninProp, err := expandIdentityPlatformTenantEnableEmailLinkSignin(d.Get("enable_email_link_signin"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_email_link_signin"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableEmailLinkSigninProp)) {
		obj["enableEmailLinkSignin"] = enableEmailLinkSigninProp
	}
	disableAuthProp, err := expandIdentityPlatformTenantDisableAuth(d.Get("disable_auth"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disable_auth"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, disableAuthProp)) {
		obj["disableAuth"] = disableAuthProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Tenant %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("allow_password_signup") {
		updateMask = append(updateMask, "allowPasswordSignup")
	}

	if d.HasChange("enable_email_link_signin") {
		updateMask = append(updateMask, "enableEmailLinkSignin")
	}

	if d.HasChange("disable_auth") {
		updateMask = append(updateMask, "disableAuth")
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
		return fmt.Errorf("Error updating Tenant %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Tenant %q: %#v", d.Id(), res)
	}

	return resourceIdentityPlatformTenantRead(d, meta)
}

func resourceIdentityPlatformTenantDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Tenant: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Tenant %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Tenant")
	}

	log.Printf("[DEBUG] Finished deleting Tenant %q: %#v", d.Id(), res)
	return nil
}

func resourceIdentityPlatformTenantImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/tenants/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/tenants/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIdentityPlatformTenantName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenIdentityPlatformTenantDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantAllowPasswordSignup(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantEnableEmailLinkSignin(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantDisableAuth(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandIdentityPlatformTenantDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantAllowPasswordSignup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantEnableEmailLinkSignin(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantDisableAuth(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
