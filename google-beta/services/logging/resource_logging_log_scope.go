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

package logging

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

func ResourceLoggingLogScope() *schema.Resource {
	return &schema.Resource{
		Create: resourceLoggingLogScopeCreate,
		Read:   resourceLoggingLogScopeRead,
		Update: resourceLoggingLogScopeUpdate,
		Delete: resourceLoggingLogScopeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceLoggingLogScopeImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareResourceNames,
				Description:      `The resource name of the log scope. For example: \'projects/my-project/locations/global/logScopes/my-log-scope\'`,
			},
			"resource_names": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Names of one or more parent resources : *  \'projects/[PROJECT_ID]\' May alternatively be one or more views : * \'projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/views/[VIEW_ID]\' A log scope can include a maximum of 50 projects and a maximum of 100 resources in total.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Describes this log scopes.`,
			},
			"location": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The location of the resource. The only supported location is global so far.`,
			},
			"parent": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The parent of the resource.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The creation timestamp of the log scopes.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The last update timestamp of the log scopes.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceLoggingLogScopeCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandLoggingLogScopeName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	resourceNamesProp, err := expandLoggingLogScopeResourceNames(d.Get("resource_names"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("resource_names"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourceNamesProp)) && (ok || !reflect.DeepEqual(v, resourceNamesProp)) {
		obj["resourceNames"] = resourceNamesProp
	}
	descriptionProp, err := expandLoggingLogScopeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	obj, err = resourceLoggingLogScopeEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{LoggingBasePath}}{{parent}}/locations/{{location}}/logScopes?logScopeId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new LogScope: %#v", obj)
	billingProject := ""

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
		return fmt.Errorf("Error creating LogScope: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/locations/{{location}}/logScopes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating LogScope %q: %#v", d.Id(), res)

	return resourceLoggingLogScopeRead(d, meta)
}

func resourceLoggingLogScopeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{LoggingBasePath}}{{parent}}/locations/{{location}}/logScopes/{{name}}")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("LoggingLogScope %q", d.Id()))
	}

	if err := d.Set("resource_names", flattenLoggingLogScopeResourceNames(res["resourceNames"], d, config)); err != nil {
		return fmt.Errorf("Error reading LogScope: %s", err)
	}
	if err := d.Set("description", flattenLoggingLogScopeDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading LogScope: %s", err)
	}
	if err := d.Set("create_time", flattenLoggingLogScopeCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading LogScope: %s", err)
	}
	if err := d.Set("update_time", flattenLoggingLogScopeUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading LogScope: %s", err)
	}

	return nil
}

func resourceLoggingLogScopeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	resourceNamesProp, err := expandLoggingLogScopeResourceNames(d.Get("resource_names"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("resource_names"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, resourceNamesProp)) {
		obj["resourceNames"] = resourceNamesProp
	}
	descriptionProp, err := expandLoggingLogScopeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	obj, err = resourceLoggingLogScopeEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{LoggingBasePath}}{{parent}}/locations/{{location}}/logScopes/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating LogScope %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("resource_names") {
		updateMask = append(updateMask, "resourceNames")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
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
			return fmt.Errorf("Error updating LogScope %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating LogScope %q: %#v", d.Id(), res)
		}

	}

	return resourceLoggingLogScopeRead(d, meta)
}

func resourceLoggingLogScopeDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{LoggingBasePath}}{{parent}}/locations/{{location}}/logScopes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting LogScope %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "LogScope")
	}

	log.Printf("[DEBUG] Finished deleting LogScope %q: %#v", d.Id(), res)
	return nil
}

func resourceLoggingLogScopeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^(?P<parent>.+)/locations/(?P<location>[^/]+)/logScopes/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/locations/{{location}}/logScopes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenLoggingLogScopeResourceNames(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenLoggingLogScopeDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenLoggingLogScopeCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenLoggingLogScopeUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandLoggingLogScopeName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingLogScopeResourceNames(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingLogScopeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceLoggingLogScopeEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Extract any empty fields from the bucket field.
	// Extract parent, location from name
	parent := d.Get("parent").(string)
	name := d.Get("name").(string)
	parent, err := tpgresource.ExtractFieldByPattern("parent", parent, name, "((projects|folders|organizations|billingAccounts)/[a-z0-9A-Z-]*)/locations/.*")
	if err != nil {
		return nil, fmt.Errorf("error extracting parent field: %s", err)
	}
	location := d.Get("location").(string)
	location, err = tpgresource.ExtractFieldByPattern("location", location, name, "[a-zA-Z]*/[a-z0-9A-Z-]*/locations/([a-z0-9-]*)/logScopes/.*")
	if err != nil {
		return nil, fmt.Errorf("error extracting location field: %s", err)
	}
	// Set parent to the extracted value.
	d.Set("parent", parent)
	// Set all the other fields to their short forms before forming url and setting ID.
	name = tpgresource.GetResourceNameFromSelfLink(name)
	d.Set("location", location)
	d.Set("name", name)
	return obj, nil
}
