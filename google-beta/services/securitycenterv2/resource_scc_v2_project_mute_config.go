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

package securitycenterv2

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceSecurityCenterV2projectMuteConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityCenterV2projectMuteConfigCreate,
		Read:   resourceSecurityCenterV2projectMuteConfigRead,
		Update: resourceSecurityCenterV2projectMuteConfigUpdate,
		Delete: resourceSecurityCenterV2projectMuteConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSecurityCenterV2projectMuteConfigImport,
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
			"filter": {
				Type:     schema.TypeString,
				Required: true,
				Description: `An expression that defines the filter to apply across create/update
events of findings. While creating a filter string, be mindful of
the scope in which the mute configuration is being created. E.g.,
If a filter contains project = X but is created under the
project = Y scope, it might not match any findings.`,
			},
			"mute_config_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Unique identifier provided by the client within the parent scope.`,
			},
			"project_parent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Resource name of the new mute configs's parent. Its format is
"organizations/[organization_id]", "folders/[folder_id]", or
"projects/[project_id]".`,
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidateEnum([]string{"MUTE_CONFIG_TYPE_UNSPECIFIED", "STATIC"}),
				Description: `Required. The type of the mute config,
which determines what type of mute state the config affects. Immutable after creation. Possible values: ["MUTE_CONFIG_TYPE_UNSPECIFIED", "STATIC"]`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A description of the mute config.`,
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `location Id is provided by organization. If not provided, Use global as default.`,
				Default:     "global",
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The time at which the mute config was created. This field is set by
the server and will be ignored if provided on config creation.`,
			},
			"most_recent_editor": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Email address of the user who last edited the mute config. This
field is set by the server and will be ignored if provided on
config creation or update.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Name of the mute config. Its format is
organizations/{organization}/muteConfigs/{configId},
folders/{folder}/muteConfigs/{configId},
or projects/{project}/muteConfigs/{configId}`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The most recent time at which the mute config was
updated. This field is set by the server and will be ignored if
provided on config creation or update.`,
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

func resourceSecurityCenterV2projectMuteConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandSecurityCenterV2projectMuteConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	filterProp, err := expandSecurityCenterV2projectMuteConfigFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(filterProp)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}
	typeProp, err := expandSecurityCenterV2projectMuteConfigType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterV2BasePath}}{{project}}/locations/{{location}}/muteConfigs?muteConfigId={{mute_config_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new projectMuteConfig: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for projectMuteConfig: %s", err)
	}
	billingProject = project

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
		return fmt.Errorf("Error creating projectMuteConfig: %s", err)
	}
	if err := d.Set("name", flattenSecurityCenterV2projectMuteConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating projectMuteConfig %q: %#v", d.Id(), res)

	return resourceSecurityCenterV2projectMuteConfigRead(d, meta)
}

func resourceSecurityCenterV2projectMuteConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterV2BasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for projectMuteConfig: %s", err)
	}
	billingProject = project

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SecurityCenterV2projectMuteConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading projectMuteConfig: %s", err)
	}

	if err := d.Set("name", flattenSecurityCenterV2projectMuteConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading projectMuteConfig: %s", err)
	}
	if err := d.Set("description", flattenSecurityCenterV2projectMuteConfigDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading projectMuteConfig: %s", err)
	}
	if err := d.Set("filter", flattenSecurityCenterV2projectMuteConfigFilter(res["filter"], d, config)); err != nil {
		return fmt.Errorf("Error reading projectMuteConfig: %s", err)
	}
	if err := d.Set("create_time", flattenSecurityCenterV2projectMuteConfigCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading projectMuteConfig: %s", err)
	}
	if err := d.Set("update_time", flattenSecurityCenterV2projectMuteConfigUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading projectMuteConfig: %s", err)
	}
	if err := d.Set("most_recent_editor", flattenSecurityCenterV2projectMuteConfigMostRecentEditor(res["mostRecentEditor"], d, config)); err != nil {
		return fmt.Errorf("Error reading projectMuteConfig: %s", err)
	}
	if err := d.Set("type", flattenSecurityCenterV2projectMuteConfigType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading projectMuteConfig: %s", err)
	}

	return nil
}

func resourceSecurityCenterV2projectMuteConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for projectMuteConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandSecurityCenterV2projectMuteConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	filterProp, err := expandSecurityCenterV2projectMuteConfigFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}
	typeProp, err := expandSecurityCenterV2projectMuteConfigType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterV2BasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating projectMuteConfig %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("filter") {
		updateMask = append(updateMask, "filter")
	}

	if d.HasChange("type") {
		updateMask = append(updateMask, "type")
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
			return fmt.Errorf("Error updating projectMuteConfig %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating projectMuteConfig %q: %#v", d.Id(), res)
		}

	}

	return resourceSecurityCenterV2projectMuteConfigRead(d, meta)
}

func resourceSecurityCenterV2projectMuteConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for projectMuteConfig: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterV2BasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting projectMuteConfig %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "projectMuteConfig")
	}

	log.Printf("[DEBUG] Finished deleting projectMuteConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceSecurityCenterV2projectMuteConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	// current import_formats can't import fields with forward slashes in their value
	name := d.Get("name").(string)

	matched, err := regexp.MatchString("(organizations|folders|projects)/.+/muteConfigs/.+", name)
	if err != nil {
		return nil, fmt.Errorf("error validating import name: %s", err)
	}

	if !matched {
		return nil, fmt.Errorf("error validating import name: %s does not fit naming for muteConfigs. Expected %s",
			name, "organizations/{organization}/muteConfigs/{configId}, folders/{folder}/muteConfigs/{configId} or projects/{project}/muteConfigs/{configId}")
	}

	if err := d.Set("name", name); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}

	// mute_config_id and parent are not returned by the API and therefore need to be set manually
	stringParts := strings.Split(d.Get("name").(string), "/")
	if err := d.Set("mute_config_id", stringParts[3]); err != nil {
		return nil, fmt.Errorf("Error setting mute_config_id: %s", err)
	}

	if err := d.Set("parent", fmt.Sprintf("%s/%s", stringParts[0], stringParts[1])); err != nil {
		return nil, fmt.Errorf("Error setting mute_config_id: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenSecurityCenterV2projectMuteConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterV2projectMuteConfigDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterV2projectMuteConfigFilter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterV2projectMuteConfigCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterV2projectMuteConfigUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterV2projectMuteConfigMostRecentEditor(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterV2projectMuteConfigType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandSecurityCenterV2projectMuteConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterV2projectMuteConfigFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterV2projectMuteConfigType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
