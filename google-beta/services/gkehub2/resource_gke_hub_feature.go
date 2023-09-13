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

func ResourceGKEHub2Feature() *schema.Resource {
	return &schema.Resource{
		Create: resourceGKEHub2FeatureCreate,
		Read:   resourceGKEHub2FeatureRead,
		Update: resourceGKEHub2FeatureUpdate,
		Delete: resourceGKEHub2FeatureDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGKEHub2FeatureImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location for the resource`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `GCP labels for this Feature.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The full, unique name of this Feature resource`,
			},
			"spec": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fleetobservability": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Fleet Observability feature spec.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"logging_config": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `Specified if fleet logging feature is enabled for the entire fleet. If UNSPECIFIED, fleet logging feature is disabled for the entire fleet.`,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"default_config": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: `Specified if applying the default routing config to logs not specified in other configs.`,
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: verify.ValidateEnum([]string{"MODE_UNSPECIFIED", "COPY", "MOVE", ""}),
																Description:  `Specified if fleet logging feature is enabled. Possible values: ["MODE_UNSPECIFIED", "COPY", "MOVE"]`,
															},
														},
													},
												},
												"fleet_scope_logs_config": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: `Specified if applying the routing config to all logs for all fleet scopes.`,
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: verify.ValidateEnum([]string{"MODE_UNSPECIFIED", "COPY", "MOVE", ""}),
																Description:  `Specified if fleet logging feature is enabled. Possible values: ["MODE_UNSPECIFIED", "COPY", "MOVE"]`,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"multiclusteringress": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Multicluster Ingress-specific spec.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"config_membership": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Fully-qualified Membership name which hosts the MultiClusterIngress CRD. Example: 'projects/foo-proj/locations/global/memberships/bar'`,
									},
								},
							},
						},
					},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. When the Feature resource was created.`,
			},
			"delete_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. When the Feature resource was deleted.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"resource_state": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `State of the Feature resource itself.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"has_resources": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: `Whether this Feature has outstanding resources that need to be cleaned up before it can be disabled.`,
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The current state of the Feature resource in the Hub API.`,
						},
					},
				},
			},
			"state": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Output only. The Hub-wide Feature state`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `Output only. The "running state" of the Feature in this Hub.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `The high-level, machine-readable status of this Feature.`,
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `A human-readable description of the current status.`,
									},
									"update_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `The time this status and any related Feature-specific details were updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"`,
									},
								},
							},
						},
					},
				},
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. When the Feature resource was last updated.`,
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

func resourceGKEHub2FeatureCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	specProp, err := expandGKEHub2FeatureSpec(d.Get("spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(specProp)) && (ok || !reflect.DeepEqual(v, specProp)) {
		obj["spec"] = specProp
	}
	labelsProp, err := expandGKEHub2FeatureEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/{{location}}/features?featureId={{name}}")
	if err != nil {
		return err
	}
	url = strings.ReplaceAll(url, "projects/projects/", "projects/")

	log.Printf("[DEBUG] Creating new Feature: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Feature: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

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
		return fmt.Errorf("Error creating Feature: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	id = strings.ReplaceAll(id, "projects/projects/", "projects/")
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = GKEHub2OperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Feature", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Feature: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	id = strings.ReplaceAll(id, "projects/projects/", "projects/")
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Feature %q: %#v", d.Id(), res)

	return resourceGKEHub2FeatureRead(d, meta)
}

func resourceGKEHub2FeatureRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return err
	}
	url = strings.ReplaceAll(url, "projects/projects/", "projects/")

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Feature: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("GKEHub2Feature %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}

	if err := d.Set("labels", flattenGKEHub2FeatureLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}
	if err := d.Set("resource_state", flattenGKEHub2FeatureResourceState(res["resourceState"], d, config)); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}
	if err := d.Set("spec", flattenGKEHub2FeatureSpec(res["spec"], d, config)); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}
	if err := d.Set("state", flattenGKEHub2FeatureState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}
	if err := d.Set("create_time", flattenGKEHub2FeatureCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}
	if err := d.Set("update_time", flattenGKEHub2FeatureUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}
	if err := d.Set("delete_time", flattenGKEHub2FeatureDeleteTime(res["deleteTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}
	if err := d.Set("terraform_labels", flattenGKEHub2FeatureTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}
	if err := d.Set("effective_labels", flattenGKEHub2FeatureEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}

	return nil
}

func resourceGKEHub2FeatureUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Feature: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

	obj := make(map[string]interface{})
	specProp, err := expandGKEHub2FeatureSpec(d.Get("spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, specProp)) {
		obj["spec"] = specProp
	}
	labelsProp, err := expandGKEHub2FeatureEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return err
	}
	url = strings.ReplaceAll(url, "projects/projects/", "projects/")

	log.Printf("[DEBUG] Updating Feature %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("spec") {
		updateMask = append(updateMask, "spec")
	}

	if d.HasChange("effective_labels") {
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
		return fmt.Errorf("Error updating Feature %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Feature %q: %#v", d.Id(), res)
	}

	err = GKEHub2OperationWaitTime(
		config, res, project, "Updating Feature", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceGKEHub2FeatureRead(d, meta)
}

func resourceGKEHub2FeatureDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Feature: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return err
	}
	url = strings.ReplaceAll(url, "projects/projects/", "projects/")

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Feature %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Feature")
	}

	err = GKEHub2OperationWaitTime(
		config, res, project, "Deleting Feature", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Feature %q: %#v", d.Id(), res)
	return nil
}

func resourceGKEHub2FeatureImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/features/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	id = strings.ReplaceAll(id, "projects/projects/", "projects/")
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenGKEHub2FeatureLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenGKEHub2FeatureResourceState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["state"] =
		flattenGKEHub2FeatureResourceStateState(original["state"], d, config)
	transformed["has_resources"] =
		flattenGKEHub2FeatureResourceStateHasResources(original["hasResources"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FeatureResourceStateState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureResourceStateHasResources(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["multiclusteringress"] =
		flattenGKEHub2FeatureSpecMulticlusteringress(original["multiclusteringress"], d, config)
	transformed["fleetobservability"] =
		flattenGKEHub2FeatureSpecFleetobservability(original["fleetobservability"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FeatureSpecMulticlusteringress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["config_membership"] =
		flattenGKEHub2FeatureSpecMulticlusteringressConfigMembership(original["configMembership"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FeatureSpecMulticlusteringressConfigMembership(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureSpecFleetobservability(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["logging_config"] =
		flattenGKEHub2FeatureSpecFleetobservabilityLoggingConfig(original["loggingConfig"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FeatureSpecFleetobservabilityLoggingConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["default_config"] =
		flattenGKEHub2FeatureSpecFleetobservabilityLoggingConfigDefaultConfig(original["defaultConfig"], d, config)
	transformed["fleet_scope_logs_config"] =
		flattenGKEHub2FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(original["fleetScopeLogsConfig"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FeatureSpecFleetobservabilityLoggingConfigDefaultConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["mode"] =
		flattenGKEHub2FeatureSpecFleetobservabilityLoggingConfigDefaultConfigMode(original["mode"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FeatureSpecFleetobservabilityLoggingConfigDefaultConfigMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["mode"] =
		flattenGKEHub2FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigMode(original["mode"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["state"] =
		flattenGKEHub2FeatureStateState(original["state"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FeatureStateState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["code"] =
		flattenGKEHub2FeatureStateStateCode(original["code"], d, config)
	transformed["description"] =
		flattenGKEHub2FeatureStateStateDescription(original["description"], d, config)
	transformed["update_time"] =
		flattenGKEHub2FeatureStateStateUpdateTime(original["updateTime"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FeatureStateStateCode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureStateStateDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureStateStateUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureDeleteTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenGKEHub2FeatureEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandGKEHub2FeatureSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMulticlusteringress, err := expandGKEHub2FeatureSpecMulticlusteringress(original["multiclusteringress"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMulticlusteringress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["multiclusteringress"] = transformedMulticlusteringress
	}

	transformedFleetobservability, err := expandGKEHub2FeatureSpecFleetobservability(original["fleetobservability"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFleetobservability); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fleetobservability"] = transformedFleetobservability
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecMulticlusteringress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedConfigMembership, err := expandGKEHub2FeatureSpecMulticlusteringressConfigMembership(original["config_membership"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConfigMembership); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["configMembership"] = transformedConfigMembership
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecMulticlusteringressConfigMembership(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureSpecFleetobservability(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLoggingConfig, err := expandGKEHub2FeatureSpecFleetobservabilityLoggingConfig(original["logging_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLoggingConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["loggingConfig"] = transformedLoggingConfig
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecFleetobservabilityLoggingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDefaultConfig, err := expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigDefaultConfig(original["default_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["defaultConfig"] = transformedDefaultConfig
	}

	transformedFleetScopeLogsConfig, err := expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(original["fleet_scope_logs_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFleetScopeLogsConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fleetScopeLogsConfig"] = transformedFleetScopeLogsConfig
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigDefaultConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMode, err := expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigDefaultConfigMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigDefaultConfigMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMode, err := expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
