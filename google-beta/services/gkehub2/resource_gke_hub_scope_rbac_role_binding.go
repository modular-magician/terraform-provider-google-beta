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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceGKEHub2ScopeRBACRoleBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceGKEHub2ScopeRBACRoleBindingCreate,
		Read:   resourceGKEHub2ScopeRBACRoleBindingRead,
		Update: resourceGKEHub2ScopeRBACRoleBindingUpdate,
		Delete: resourceGKEHub2ScopeRBACRoleBindingDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGKEHub2ScopeRBACRoleBindingImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"role": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Role to bind to the principal.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"predefined_role": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"UNKNOWN", "ADMIN", "EDIT", "VIEW", ""}),
							Description:  `PredefinedRole is an ENUM representation of the default Kubernetes Roles Possible values: ["UNKNOWN", "ADMIN", "EDIT", "VIEW"]`,
						},
					},
				},
			},
			"scope_rbac_role_binding_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The client-provided identifier of the RBAC Role Binding.`,
			},
			"scope_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Id of the scope`,
			},
			"group": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Principal that is be authorized in the cluster (at least of one the oneof
is required). Updating one will unset the other automatically.
group is the group, as seen by the kubernetes cluster.`,
				ExactlyOneOf: []string{"user", "group"},
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Principal that is be authorized in the cluster (at least of one the oneof
is required). Updating one will unset the other automatically.
user is the name of the user as seen by the kubernetes cluster, example
"alice" or "alice@domain.tld"`,
				ExactlyOneOf: []string{"user", "group"},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the RBAC Role Binding was created in UTC.`,
			},
			"delete_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the RBAC Role Binding was deleted in UTC.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name for the RBAC Role Binding`,
			},
			"state": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `State of the RBAC Role Binding resource.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Code describes the state of a RBAC Role Binding resource.`,
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
				Description: `Time the RBAC Role Binding was updated in UTC.`,
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

func resourceGKEHub2ScopeRBACRoleBindingCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	userProp, err := expandGKEHub2ScopeRBACRoleBindingUser(d.Get("user"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user"); !tpgresource.IsEmptyValue(reflect.ValueOf(userProp)) && (ok || !reflect.DeepEqual(v, userProp)) {
		obj["user"] = userProp
	}
	groupProp, err := expandGKEHub2ScopeRBACRoleBindingGroup(d.Get("group"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("group"); !tpgresource.IsEmptyValue(reflect.ValueOf(groupProp)) && (ok || !reflect.DeepEqual(v, groupProp)) {
		obj["group"] = groupProp
	}
	roleProp, err := expandGKEHub2ScopeRBACRoleBindingRole(d.Get("role"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("role"); !tpgresource.IsEmptyValue(reflect.ValueOf(roleProp)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/scopes/{{scope_id}}/rbacrolebindings/?rbacrolebinding_id={{scope_rbac_role_binding_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ScopeRBACRoleBinding: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ScopeRBACRoleBinding: %s", err)
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
		return fmt.Errorf("Error creating ScopeRBACRoleBinding: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/scopes/{{scope_id}}/rbacrolebindings/{{scope_rbac_role_binding_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = GKEHub2OperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating ScopeRBACRoleBinding", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create ScopeRBACRoleBinding: %s", err)
	}

	if err := d.Set("name", flattenGKEHub2ScopeRBACRoleBindingName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/scopes/{{scope_id}}/rbacrolebindings/{{scope_rbac_role_binding_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ScopeRBACRoleBinding %q: %#v", d.Id(), res)

	return resourceGKEHub2ScopeRBACRoleBindingRead(d, meta)
}

func resourceGKEHub2ScopeRBACRoleBindingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/scopes/{{scope_id}}/rbacrolebindings/{{scope_rbac_role_binding_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ScopeRBACRoleBinding: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("GKEHub2ScopeRBACRoleBinding %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ScopeRBACRoleBinding: %s", err)
	}

	if err := d.Set("name", flattenGKEHub2ScopeRBACRoleBindingName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScopeRBACRoleBinding: %s", err)
	}
	if err := d.Set("uid", flattenGKEHub2ScopeRBACRoleBindingUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScopeRBACRoleBinding: %s", err)
	}
	if err := d.Set("create_time", flattenGKEHub2ScopeRBACRoleBindingCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScopeRBACRoleBinding: %s", err)
	}
	if err := d.Set("update_time", flattenGKEHub2ScopeRBACRoleBindingUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScopeRBACRoleBinding: %s", err)
	}
	if err := d.Set("delete_time", flattenGKEHub2ScopeRBACRoleBindingDeleteTime(res["deleteTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScopeRBACRoleBinding: %s", err)
	}
	if err := d.Set("state", flattenGKEHub2ScopeRBACRoleBindingState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScopeRBACRoleBinding: %s", err)
	}
	if err := d.Set("user", flattenGKEHub2ScopeRBACRoleBindingUser(res["user"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScopeRBACRoleBinding: %s", err)
	}
	if err := d.Set("group", flattenGKEHub2ScopeRBACRoleBindingGroup(res["group"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScopeRBACRoleBinding: %s", err)
	}
	if err := d.Set("role", flattenGKEHub2ScopeRBACRoleBindingRole(res["role"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScopeRBACRoleBinding: %s", err)
	}

	return nil
}

func resourceGKEHub2ScopeRBACRoleBindingUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ScopeRBACRoleBinding: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	userProp, err := expandGKEHub2ScopeRBACRoleBindingUser(d.Get("user"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, userProp)) {
		obj["user"] = userProp
	}
	groupProp, err := expandGKEHub2ScopeRBACRoleBindingGroup(d.Get("group"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("group"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, groupProp)) {
		obj["group"] = groupProp
	}
	roleProp, err := expandGKEHub2ScopeRBACRoleBindingRole(d.Get("role"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("role"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/scopes/{{scope_id}}/rbacrolebindings/{{scope_rbac_role_binding_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ScopeRBACRoleBinding %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("user") {
		updateMask = append(updateMask, "user")
	}

	if d.HasChange("group") {
		updateMask = append(updateMask, "group")
	}

	if d.HasChange("role") {
		updateMask = append(updateMask, "role")
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
		return fmt.Errorf("Error updating ScopeRBACRoleBinding %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ScopeRBACRoleBinding %q: %#v", d.Id(), res)
	}

	err = GKEHub2OperationWaitTime(
		config, res, project, "Updating ScopeRBACRoleBinding", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceGKEHub2ScopeRBACRoleBindingRead(d, meta)
}

func resourceGKEHub2ScopeRBACRoleBindingDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ScopeRBACRoleBinding: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/scopes/{{scope_id}}/rbacrolebindings/{{scope_rbac_role_binding_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ScopeRBACRoleBinding %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "ScopeRBACRoleBinding")
	}

	err = GKEHub2OperationWaitTime(
		config, res, project, "Deleting ScopeRBACRoleBinding", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ScopeRBACRoleBinding %q: %#v", d.Id(), res)
	return nil
}

func resourceGKEHub2ScopeRBACRoleBindingImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/global/scopes/(?P<scope_id>[^/]+)/rbacrolebindings/(?P<scope_rbac_role_binding_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<scope_id>[^/]+)/(?P<scope_rbac_role_binding_id>[^/]+)",
		"(?P<scope_id>[^/]+)/(?P<scope_rbac_role_binding_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/scopes/{{scope_id}}/rbacrolebindings/{{scope_rbac_role_binding_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenGKEHub2ScopeRBACRoleBindingName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2ScopeRBACRoleBindingUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2ScopeRBACRoleBindingCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2ScopeRBACRoleBindingUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2ScopeRBACRoleBindingDeleteTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2ScopeRBACRoleBindingState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["code"] =
		flattenGKEHub2ScopeRBACRoleBindingStateCode(original["code"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2ScopeRBACRoleBindingStateCode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2ScopeRBACRoleBindingUser(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2ScopeRBACRoleBindingGroup(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2ScopeRBACRoleBindingRole(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["predefined_role"] =
		flattenGKEHub2ScopeRBACRoleBindingRolePredefinedRole(original["predefinedRole"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2ScopeRBACRoleBindingRolePredefinedRole(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandGKEHub2ScopeRBACRoleBindingUser(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2ScopeRBACRoleBindingGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2ScopeRBACRoleBindingRole(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPredefinedRole, err := expandGKEHub2ScopeRBACRoleBindingRolePredefinedRole(original["predefined_role"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPredefinedRole); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["predefinedRole"] = transformedPredefinedRole
	}

	return transformed, nil
}

func expandGKEHub2ScopeRBACRoleBindingRolePredefinedRole(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
