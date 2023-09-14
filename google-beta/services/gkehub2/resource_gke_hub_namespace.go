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

func ResourceGKEHub2Namespace() *schema.Resource {
	return &schema.Resource{
		Create: resourceGKEHub2NamespaceCreate,
		Read:   resourceGKEHub2NamespaceRead,
		Update: resourceGKEHub2NamespaceUpdate,
		Delete: resourceGKEHub2NamespaceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGKEHub2NamespaceImport,
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
			"scope": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.ProjectNumberDiffSuppress,
				Description:      `The name of the Scope instance.`,
			},
			"scope_namespace_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The client-provided identifier of the namespace.`,
			},
			"scope_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Id of the scope`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Labels for this Namespace.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"namespace_labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Namespace-level cluster namespace labels. These labels are applied
to the related namespace of the member clusters bound to the parent
Scope. Scope-level labels ('namespace_labels' in the Fleet Scope
resource) take precedence over Namespace-level labels if they share
a key. Keys and values must be Kubernetes-conformant.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the Namespace was created in UTC.`,
			},
			"delete_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the Namespace was deleted in UTC.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name for the namespace`,
			},
			"state": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `State of the namespace resource.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Code describes the state of a Namespace resource.`,
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
				Description: `Time the Namespace was updated in UTC.`,
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

func resourceGKEHub2NamespaceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	scopeProp, err := expandGKEHub2NamespaceScope(d.Get("scope"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("scope"); !tpgresource.IsEmptyValue(reflect.ValueOf(scopeProp)) && (ok || !reflect.DeepEqual(v, scopeProp)) {
		obj["scope"] = scopeProp
	}
	namespaceLabelsProp, err := expandGKEHub2NamespaceNamespaceLabels(d.Get("namespace_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("namespace_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(namespaceLabelsProp)) && (ok || !reflect.DeepEqual(v, namespaceLabelsProp)) {
		obj["namespaceLabels"] = namespaceLabelsProp
	}
	labelsProp, err := expandGKEHub2NamespaceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/scopes/{{scope_id}}/namespaces/?scope_namespace_id={{scope_namespace_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Namespace: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Namespace: %s", err)
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
		return fmt.Errorf("Error creating Namespace: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/scopes/{{scope_id}}/namespaces/{{scope_namespace_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = GKEHub2OperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Namespace", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Namespace: %s", err)
	}

	if err := d.Set("name", flattenGKEHub2NamespaceName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/scopes/{{scope_id}}/namespaces/{{scope_namespace_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Namespace %q: %#v", d.Id(), res)

	return resourceGKEHub2NamespaceRead(d, meta)
}

func resourceGKEHub2NamespaceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/scopes/{{scope_id}}/namespaces/{{scope_namespace_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Namespace: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("GKEHub2Namespace %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Namespace: %s", err)
	}

	if err := d.Set("name", flattenGKEHub2NamespaceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Namespace: %s", err)
	}
	if err := d.Set("uid", flattenGKEHub2NamespaceUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Namespace: %s", err)
	}
	if err := d.Set("create_time", flattenGKEHub2NamespaceCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Namespace: %s", err)
	}
	if err := d.Set("update_time", flattenGKEHub2NamespaceUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Namespace: %s", err)
	}
	if err := d.Set("delete_time", flattenGKEHub2NamespaceDeleteTime(res["deleteTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Namespace: %s", err)
	}
	if err := d.Set("scope", flattenGKEHub2NamespaceScope(res["scope"], d, config)); err != nil {
		return fmt.Errorf("Error reading Namespace: %s", err)
	}
	if err := d.Set("state", flattenGKEHub2NamespaceState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Namespace: %s", err)
	}
	if err := d.Set("namespace_labels", flattenGKEHub2NamespaceNamespaceLabels(res["namespaceLabels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Namespace: %s", err)
	}
	if err := d.Set("labels", flattenGKEHub2NamespaceLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Namespace: %s", err)
	}

	return nil
}

func resourceGKEHub2NamespaceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Namespace: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	namespaceLabelsProp, err := expandGKEHub2NamespaceNamespaceLabels(d.Get("namespace_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("namespace_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, namespaceLabelsProp)) {
		obj["namespaceLabels"] = namespaceLabelsProp
	}
	labelsProp, err := expandGKEHub2NamespaceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/scopes/{{scope_id}}/namespaces/{{scope_namespace_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Namespace %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("namespace_labels") {
		updateMask = append(updateMask, "namespaceLabels")
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
		return fmt.Errorf("Error updating Namespace %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Namespace %q: %#v", d.Id(), res)
	}

	err = GKEHub2OperationWaitTime(
		config, res, project, "Updating Namespace", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceGKEHub2NamespaceRead(d, meta)
}

func resourceGKEHub2NamespaceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Namespace: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/scopes/{{scope_id}}/namespaces/{{scope_namespace_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Namespace %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Namespace")
	}

	err = GKEHub2OperationWaitTime(
		config, res, project, "Deleting Namespace", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Namespace %q: %#v", d.Id(), res)
	return nil
}

func resourceGKEHub2NamespaceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/global/scopes/(?P<scope_id>[^/]+)/namespaces/(?P<scope_namespace_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<scope_id>[^/]+)/(?P<scope_namespace_id>[^/]+)",
		"(?P<scope_id>[^/]+)/(?P<scope_namespace_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/scopes/{{scope_id}}/namespaces/{{scope_namespace_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenGKEHub2NamespaceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2NamespaceUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2NamespaceCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2NamespaceUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2NamespaceDeleteTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2NamespaceScope(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenGKEHub2NamespaceState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["code"] =
		flattenGKEHub2NamespaceStateCode(original["code"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2NamespaceStateCode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2NamespaceNamespaceLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2NamespaceLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandGKEHub2NamespaceScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2NamespaceNamespaceLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGKEHub2NamespaceLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
