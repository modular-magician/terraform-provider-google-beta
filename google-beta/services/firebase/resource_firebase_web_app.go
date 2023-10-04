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

package firebase

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

func ResourceFirebaseWebApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaseWebAppCreate,
		Read:   resourceFirebaseWebAppRead,
		Update: resourceFirebaseWebAppUpdate,
		Delete: resourceFirebaseWebAppDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaseWebAppImport,
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
				Description: `The user-assigned display name of the App.`,
			},
			"api_key_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				Description: `The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the WebApp.
If apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the WebApp.
This auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.`,
			},
			"app_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The globally unique, Firebase-assigned identifier of the App.
This identifier should be treated as an opaque token, as the data format is not specified.`,
			},
			"app_urls": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The URLs where the 'WebApp' is hosted.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The fully qualified resource name of the App, for example:
projects/projectId/webApps/appId`,
			},
			"deletion_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DELETE",
				Description: `Set to 'ABANDON' to allow the WebApp to be untracked from terraform state
rather than deleted upon 'terraform destroy'. This is useful becaue the WebApp may be
serving traffic. Set to 'DELETE' to delete the WebApp. Default to 'DELETE'`,
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

func resourceFirebaseWebAppCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandFirebaseWebAppDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	apiKeyIdProp, err := expandFirebaseWebAppApiKeyId(d.Get("api_key_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("api_key_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(apiKeyIdProp)) && (ok || !reflect.DeepEqual(v, apiKeyIdProp)) {
		obj["apiKeyId"] = apiKeyIdProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}/webApps")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new WebApp: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for WebApp: %s", err)
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
		return fmt.Errorf("Error creating WebApp: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/webApps/{{app_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = FirebaseOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating WebApp", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create WebApp: %s", err)
	}

	if err := d.Set("name", flattenFirebaseWebAppName(opRes["name"], d, config)); err != nil {
		return err
	}
	if err := d.Set("app_id", flattenFirebaseWebAppAppId(opRes["appId"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/webApps/{{app_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating WebApp %q: %#v", d.Id(), res)

	return resourceFirebaseWebAppRead(d, meta)
}

func resourceFirebaseWebAppRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}/webApps/{{app_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for WebApp: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("FirebaseWebApp %q", d.Id()))
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("deletion_policy"); !ok {
		if err := d.Set("deletion_policy", "DELETE"); err != nil {
			return fmt.Errorf("Error setting deletion_policy: %s", err)
		}
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading WebApp: %s", err)
	}

	if err := d.Set("name", flattenFirebaseWebAppName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading WebApp: %s", err)
	}
	if err := d.Set("display_name", flattenFirebaseWebAppDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading WebApp: %s", err)
	}
	if err := d.Set("app_id", flattenFirebaseWebAppAppId(res["appId"], d, config)); err != nil {
		return fmt.Errorf("Error reading WebApp: %s", err)
	}
	if err := d.Set("app_urls", flattenFirebaseWebAppAppUrls(res["appUrls"], d, config)); err != nil {
		return fmt.Errorf("Error reading WebApp: %s", err)
	}
	if err := d.Set("api_key_id", flattenFirebaseWebAppApiKeyId(res["apiKeyId"], d, config)); err != nil {
		return fmt.Errorf("Error reading WebApp: %s", err)
	}

	return nil
}

func resourceFirebaseWebAppUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for WebApp: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandFirebaseWebAppDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	apiKeyIdProp, err := expandFirebaseWebAppApiKeyId(d.Get("api_key_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("api_key_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, apiKeyIdProp)) {
		obj["apiKeyId"] = apiKeyIdProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}/webApps/{{app_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating WebApp %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("api_key_id") {
		updateMask = append(updateMask, "apiKeyId")
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
			return fmt.Errorf("Error updating WebApp %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating WebApp %q: %#v", d.Id(), res)
		}

	}

	return resourceFirebaseWebAppRead(d, meta)
}

func resourceFirebaseWebAppDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	// Handwritten
	obj := make(map[string]interface{})
	if d.Get("deletion_policy") == "DELETE" {
		obj["immediate"] = true
	} else {
		fmt.Printf("Skip deleting App %q due to deletion_policy: %q\n", d.Id(), d.Get("deletion_policy"))
		return nil
	}
	// End of Handwritten
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for App: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseBasePath}}{{name}}:remove")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting App %q", d.Id())

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
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "App")
	}

	err = FirebaseOperationWaitTime(
		config, res, project, "Deleting App", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting App %q: %#v", d.Id(), res)

	// This is useful if the Delete operation returns before the Get operation
	// during post-test destroy shows the completed state of the resource.
	time.Sleep(5 * time.Second)

	return nil
}

func resourceFirebaseWebAppImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^(?P<project>[^/]+) projects/(?P<project>[^/]+)/webApps/(?P<app_id>[^/]+)$",
		"^projects/(?P<project>[^/]+)/webApps/(?P<app_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<project>[^/]+)/(?P<app_id>[^/]+)$",
		"^webApps/(?P<app_id>[^/]+)$",
		"^(?P<app_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/webApps/{{app_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Explicitly set virtual fields to default values on import
	if err := d.Set("deletion_policy", "DELETE"); err != nil {
		return nil, fmt.Errorf("Error setting deletion_policy: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenFirebaseWebAppName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseWebAppDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseWebAppAppId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseWebAppAppUrls(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseWebAppApiKeyId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandFirebaseWebAppDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseWebAppApiKeyId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
