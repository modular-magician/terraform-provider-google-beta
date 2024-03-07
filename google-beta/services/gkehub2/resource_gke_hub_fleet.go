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
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceGKEHub2Fleet() *schema.Resource {
	return &schema.Resource{
		Create: resourceGKEHub2FleetCreate,
		Read:   resourceGKEHub2FleetRead,
		Update: resourceGKEHub2FleetUpdate,
		Delete: resourceGKEHub2FleetDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGKEHub2FleetImport,
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
			"default_cluster_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The default cluster configurations to apply across the fleet.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"binary_authorization_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Enable/Disable binary authorization features for the cluster.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"evaluation_mode": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: verify.ValidateEnum([]string{"DISABLED", "POLICY_BINDINGS", ""}),
										Description:  `Mode of operation for binauthz policy evaluation. Possible values: ["DISABLED", "POLICY_BINDINGS"]`,
									},
									"policy_bindings": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `Binauthz policies that apply to this cluster.`,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Optional: true,
													Description: `The relative resource name of the binauthz platform policy to audit. GKE
platform policies have the following format:
'projects/{project_number}/platforms/gke/policies/{policy_id}'.`,
												},
											},
										},
									},
								},
							},
						},
						"security_posture_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Enable/Disable Security Posture features for the cluster.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mode": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: verify.ValidateEnum([]string{"DISABLED", "BASIC", ""}),
										Description:  `Sets which mode to use for Security Posture features. Possible values: ["DISABLED", "BASIC"]`,
									},
									"vulnerability_mode": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: verify.ValidateEnum([]string{"VULNERABILITY_DISABLED", "VULNERABILITY_BASIC", "VULNERABILITY_ENTERPRISE", ""}),
										Description:  `Sets which mode to use for vulnerability scanning. Possible values: ["VULNERABILITY_DISABLED", "VULNERABILITY_BASIC", "VULNERABILITY_ENTERPRISE"]`,
									},
								},
							},
						},
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `A user-assigned display name of the Fleet. When present, it must be between 4 to 30 characters.
Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time the fleet was created, in RFC3339 text format.`,
			},
			"delete_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time the fleet was deleted, in RFC3339 text format.`,
			},
			"state": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The state of the fleet resource.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Describes the state of a Fleet resource.`,
						},
					},
				},
			},
			"uid": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Google-generated UUID for this resource. This is unique across all
Fleet resources. If a Fleet resource is deleted and another
resource with the same name is created, it gets a different uid.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time the fleet was last updated, in RFC3339 text format.`,
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

func resourceGKEHub2FleetCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandGKEHub2FleetDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	defaultClusterConfigProp, err := expandGKEHub2FleetDefaultClusterConfig(d.Get("default_cluster_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_cluster_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultClusterConfigProp)) && (ok || !reflect.DeepEqual(v, defaultClusterConfigProp)) {
		obj["defaultClusterConfig"] = defaultClusterConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/fleets")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Fleet: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Fleet: %s", err)
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
		return fmt.Errorf("Error creating Fleet: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/fleets/default")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = GKEHub2OperationWaitTime(
		config, res, project, "Creating Fleet", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Fleet: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Fleet %q: %#v", d.Id(), res)

	return resourceGKEHub2FleetRead(d, meta)
}

func resourceGKEHub2FleetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/fleets/default")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Fleet: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("GKEHub2Fleet %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Fleet: %s", err)
	}

	if err := d.Set("display_name", flattenGKEHub2FleetDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Fleet: %s", err)
	}
	if err := d.Set("create_time", flattenGKEHub2FleetCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Fleet: %s", err)
	}
	if err := d.Set("update_time", flattenGKEHub2FleetUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Fleet: %s", err)
	}
	if err := d.Set("delete_time", flattenGKEHub2FleetDeleteTime(res["deleteTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Fleet: %s", err)
	}
	if err := d.Set("uid", flattenGKEHub2FleetUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Fleet: %s", err)
	}
	if err := d.Set("state", flattenGKEHub2FleetState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Fleet: %s", err)
	}
	if err := d.Set("default_cluster_config", flattenGKEHub2FleetDefaultClusterConfig(res["defaultClusterConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Fleet: %s", err)
	}

	return nil
}

func resourceGKEHub2FleetUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Fleet: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandGKEHub2FleetDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	defaultClusterConfigProp, err := expandGKEHub2FleetDefaultClusterConfig(d.Get("default_cluster_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_cluster_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, defaultClusterConfigProp)) {
		obj["defaultClusterConfig"] = defaultClusterConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/fleets/default")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Fleet %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("default_cluster_config") {
		updateMask = append(updateMask, "defaultClusterConfig")
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
			return fmt.Errorf("Error updating Fleet %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Fleet %q: %#v", d.Id(), res)
		}

		err = GKEHub2OperationWaitTime(
			config, res, project, "Updating Fleet", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceGKEHub2FleetRead(d, meta)
}

func resourceGKEHub2FleetDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Fleet: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/global/fleets/default")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	log.Printf("[DEBUG] Deleting Fleet %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "Fleet")
	}

	err = GKEHub2OperationWaitTime(
		config, res, project, "Deleting Fleet", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Fleet %q: %#v", d.Id(), res)
	return nil
}

func resourceGKEHub2FleetImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/global/fleets/default$",
		"^(?P<project>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/fleets/default")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenGKEHub2FleetDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FleetCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FleetUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FleetDeleteTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FleetUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FleetState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["code"] =
		flattenGKEHub2FleetStateCode(original["code"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FleetStateCode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FleetDefaultClusterConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["binary_authorization_config"] =
		flattenGKEHub2FleetDefaultClusterConfigBinaryAuthorizationConfig(original["binaryAuthorizationConfig"], d, config)
	transformed["security_posture_config"] =
		flattenGKEHub2FleetDefaultClusterConfigSecurityPostureConfig(original["securityPostureConfig"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FleetDefaultClusterConfigBinaryAuthorizationConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["evaluation_mode"] =
		flattenGKEHub2FleetDefaultClusterConfigBinaryAuthorizationConfigEvaluationMode(original["evaluationMode"], d, config)
	transformed["policy_bindings"] =
		flattenGKEHub2FleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindings(original["policyBindings"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FleetDefaultClusterConfigBinaryAuthorizationConfigEvaluationMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindings(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"name": flattenGKEHub2FleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsName(original["name"], d, config),
		})
	}
	return transformed
}
func flattenGKEHub2FleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FleetDefaultClusterConfigSecurityPostureConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["mode"] =
		flattenGKEHub2FleetDefaultClusterConfigSecurityPostureConfigMode(original["mode"], d, config)
	transformed["vulnerability_mode"] =
		flattenGKEHub2FleetDefaultClusterConfigSecurityPostureConfigVulnerabilityMode(original["vulnerabilityMode"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FleetDefaultClusterConfigSecurityPostureConfigMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FleetDefaultClusterConfigSecurityPostureConfigVulnerabilityMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandGKEHub2FleetDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FleetDefaultClusterConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBinaryAuthorizationConfig, err := expandGKEHub2FleetDefaultClusterConfigBinaryAuthorizationConfig(original["binary_authorization_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBinaryAuthorizationConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["binaryAuthorizationConfig"] = transformedBinaryAuthorizationConfig
	}

	transformedSecurityPostureConfig, err := expandGKEHub2FleetDefaultClusterConfigSecurityPostureConfig(original["security_posture_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecurityPostureConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["securityPostureConfig"] = transformedSecurityPostureConfig
	}

	return transformed, nil
}

func expandGKEHub2FleetDefaultClusterConfigBinaryAuthorizationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEvaluationMode, err := expandGKEHub2FleetDefaultClusterConfigBinaryAuthorizationConfigEvaluationMode(original["evaluation_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEvaluationMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["evaluationMode"] = transformedEvaluationMode
	}

	transformedPolicyBindings, err := expandGKEHub2FleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindings(original["policy_bindings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPolicyBindings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["policyBindings"] = transformedPolicyBindings
	}

	return transformed, nil
}

func expandGKEHub2FleetDefaultClusterConfigBinaryAuthorizationConfigEvaluationMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandGKEHub2FleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGKEHub2FleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindingsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FleetDefaultClusterConfigSecurityPostureConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMode, err := expandGKEHub2FleetDefaultClusterConfigSecurityPostureConfigMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	transformedVulnerabilityMode, err := expandGKEHub2FleetDefaultClusterConfigSecurityPostureConfigVulnerabilityMode(original["vulnerability_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVulnerabilityMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vulnerabilityMode"] = transformedVulnerabilityMode
	}

	return transformed, nil
}

func expandGKEHub2FleetDefaultClusterConfigSecurityPostureConfigMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FleetDefaultClusterConfigSecurityPostureConfigVulnerabilityMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
