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

package gkehub2

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

func ResourceGKEHub2Scope() *schema.Resource {
	return &schema.Resource{
		Create: resourceGKEHub2ScopeCreate,
		Read:   resourceGKEHub2ScopeRead,
		Update: resourceGKEHub2ScopeUpdate,
		Delete: resourceGKEHub2ScopeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGKEHub2ScopeImport,
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
			"scope_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The client-provided identifier of the scope.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Labels for this Scope.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the Scope was created in UTC.`,
			},
			"delete_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the Scope was deleted in UTC.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The unique identifier of the scope`,
			},
			"state": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `State of the scope resource.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Code describes the state of a Scope resource.`,
						},
					},
				},
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Google-generated UUID for this resource.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the Scope was updated in UTC.`,
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

func resourceGKEHub2ScopeCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandGKEHub2ScopeLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/scopes?scopeId={{scope_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Scope: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Scope: %s", err)
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
		return fmt.Errorf("Error creating Scope: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/scopes/{{scope_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = GKEHub2OperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Scope", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Scope: %s", err)
	}

	if err := d.Set("name", flattenGKEHub2ScopeName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/scopes/{{scope_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Scope %q: %#v", d.Id(), res)

	return resourceGKEHub2ScopeRead(d, meta)
}

func resourceGKEHub2ScopeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/scopes/{{scope_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Scope: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("GKEHub2Scope %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Scope: %s", err)
	}

	if err := d.Set("name", flattenGKEHub2ScopeName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Scope: %s", err)
	}
	if err := d.Set("uid", flattenGKEHub2ScopeUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Scope: %s", err)
	}
	if err := d.Set("create_time", flattenGKEHub2ScopeCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Scope: %s", err)
	}
	if err := d.Set("update_time", flattenGKEHub2ScopeUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Scope: %s", err)
	}
	if err := d.Set("delete_time", flattenGKEHub2ScopeDeleteTime(res["deleteTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Scope: %s", err)
	}
	if err := d.Set("state", flattenGKEHub2ScopeState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Scope: %s", err)
	}
	if err := d.Set("labels", flattenGKEHub2ScopeLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Scope: %s", err)
	}

	return nil
}

func resourceGKEHub2ScopeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Scope: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandGKEHub2ScopeLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/scopes/{{scope_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Scope %q: %#v", d.Id(), obj)
	updateMask := []string{}

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
		return fmt.Errorf("Error updating Scope %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Scope %q: %#v", d.Id(), res)
	}

	err = GKEHub2OperationWaitTime(
		config, res, project, "Updating Scope", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceGKEHub2ScopeRead(d, meta)
}

func resourceGKEHub2ScopeDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Scope: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/scopes/{{scope_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Scope %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Scope")
	}

	err = GKEHub2OperationWaitTime(
		config, res, project, "Deleting Scope", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Scope %q: %#v", d.Id(), res)
	return nil
}

func resourceGKEHub2ScopeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/global/scopes/(?P<scope_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<scope_id>[^/]+)",
		"(?P<scope_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/scopes/{{scope_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenGKEHub2ScopeName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2ScopeUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2ScopeCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2ScopeUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2ScopeDeleteTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2ScopeState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["code"] =
		flattenGKEHub2ScopeStateCode(original["code"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2ScopeStateCode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2ScopeLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandGKEHub2ScopeLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
