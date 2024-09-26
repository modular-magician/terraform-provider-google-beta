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

package networkconnectivity

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceNetworkConnectivityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkConnectivityGroupCreate,
		Read:   resourceNetworkConnectivityGroupRead,
		Update: resourceNetworkConnectivityGroupUpdate,
		Delete: resourceNetworkConnectivityGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkConnectivityGroupImport,
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
			"hub": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The name of the hub. Hub names must be unique. They use the following form: projects/{projectNumber}/locations/global/hubs/{hubId}`,
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateFunc:     verify.ValidateEnum([]string{"default", "center", "edge"}),
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The name of the group. Group names must be unique. They use the following form: 'projects/{projectNumber}/locations/global/hubs/{hub}/groups/{groupId}' Possible values: ["default", "center", "edge"]`,
			},
			"auto_accept": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Optional. The auto-accept setting for this group.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_accept_projects": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Description: `A list of project ids or project numbers for which you want to enable auto-accept. The auto-accept setting is applied to spokes being created or updated in these projects.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of the hub.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The time the hub was created.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"route_table": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The name of the route table that corresponds to this group. They use the following form: 'projects/{projectNumber}/locations/global/hubs/{hubId}/routeTables/{route_table_id}'`,
			},
			"state": {
				Type:         schema.TypeString,
				Computed:     true,
				Description:  `Output only. The current lifecycle state of this hub.`,
				ExactlyOneOf: []string{},
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The Google-generated UUID for the group. This value is unique across all group resources. If a group is deleted and another with the same name is created, the new route table is assigned a different uniqueId.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The time the hub was last updated.`,
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

func resourceNetworkConnectivityGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandNetworkConnectivityGroupName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandNetworkConnectivityGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	autoAcceptProp, err := expandNetworkConnectivityGroupAutoAccept(d.Get("auto_accept"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("auto_accept"); !tpgresource.IsEmptyValue(reflect.ValueOf(autoAcceptProp)) && (ok || !reflect.DeepEqual(v, autoAcceptProp)) {
		obj["autoAccept"] = autoAcceptProp
	}
	labelsProp, err := expandNetworkConnectivityGroupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/global/hubs/{{hub}}/groups/{{name}}?updateMask=autoAccept.autoAcceptProjects")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Group: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Group: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating Group: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVarsForId(d, config, "projects/{{project}}/locations/global/hubs/{{hub}}/groups/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkConnectivityOperationWaitTime(
		config, res, tpgresource.GetResourceNameFromSelfLink(project), "Creating Group", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Group: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Group %q: %#v", d.Id(), res)

	return resourceNetworkConnectivityGroupRead(d, meta)
}

func resourceNetworkConnectivityGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/global/hubs/{{hub}}/groups/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Group: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkConnectivityGroup %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}

	if err := d.Set("name", flattenNetworkConnectivityGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("create_time", flattenNetworkConnectivityGroupCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkConnectivityGroupUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("labels", flattenNetworkConnectivityGroupLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("description", flattenNetworkConnectivityGroupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("uid", flattenNetworkConnectivityGroupUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("state", flattenNetworkConnectivityGroupState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("auto_accept", flattenNetworkConnectivityGroupAutoAccept(res["autoAccept"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("route_table", flattenNetworkConnectivityGroupRouteTable(res["routeTable"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkConnectivityGroupTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkConnectivityGroupEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}

	return nil
}

func resourceNetworkConnectivityGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Group: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkConnectivityGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetworkConnectivityGroupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/global/hubs/{{hub}}/groups/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Group %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
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
			return fmt.Errorf("Error updating Group %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Group %q: %#v", d.Id(), res)
		}

		err = NetworkConnectivityOperationWaitTime(
			config, res, tpgresource.GetResourceNameFromSelfLink(project), "Updating Group", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkConnectivityGroupRead(d, meta)
}

func resourceNetworkConnectivityGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Group: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/global/hubs/{{hub}}/groups/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Group %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "Group")
	}

	err = NetworkConnectivityOperationWaitTime(
		config, res, tpgresource.GetResourceNameFromSelfLink(project), "Deleting Group", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Group %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkConnectivityGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/global/hubs/(?P<hub>[^/]+)/groups/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<hub>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<hub>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVarsForId(d, config, "projects/{{project}}/locations/global/hubs/{{hub}}/groups/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkConnectivityGroupName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityGroupCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityGroupUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityGroupLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkConnectivityGroupDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityGroupUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityGroupState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityGroupAutoAccept(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["auto_accept_projects"] =
		flattenNetworkConnectivityGroupAutoAcceptAutoAcceptProjects(original["autoAcceptProjects"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkConnectivityGroupAutoAcceptAutoAcceptProjects(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityGroupRouteTable(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityGroupTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkConnectivityGroupEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkConnectivityGroupName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityGroupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityGroupAutoAccept(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAutoAcceptProjects, err := expandNetworkConnectivityGroupAutoAcceptAutoAcceptProjects(original["auto_accept_projects"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAutoAcceptProjects); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["autoAcceptProjects"] = transformedAutoAcceptProjects
	}

	return transformed, nil
}

func expandNetworkConnectivityGroupAutoAcceptAutoAcceptProjects(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityGroupEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
