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

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceApiGatewayGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayGatewayCreate,
		Read:   resourceApiGatewayGatewayRead,
		Update: resourceApiGatewayGatewayUpdate,
		Delete: resourceApiGatewayGatewayDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApiGatewayGatewayImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"api_config": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tpgresource.CompareResourceNames,
				Description: `Resource name of the API Config for this Gateway. Format: projects/{project}/locations/global/apis/{api}/configs/{apiConfig}.
When changing api configs please ensure the new config is a new resource and the lifecycle rule 'create_before_destroy' is set.`,
			},
			"gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Identifier to assign to the Gateway. Must be unique within scope of the parent resource(project).`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: `A user-visible name for the API.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Resource labels to represent user-provided metadata.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The region of the gateway for the API.`,
			},
			"default_hostname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The default API Gateway host name of the form {gatewayId}-{hash}.{region_code}.gateway.dev.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Resource name of the Gateway. Format: projects/{project}/locations/{region}/gateways/{gateway}`,
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

func resourceApiGatewayGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandApiGatewayGatewayDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	apiConfigProp, err := expandApiGatewayGatewayApiConfig(d.Get("api_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("api_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(apiConfigProp)) && (ok || !reflect.DeepEqual(v, apiConfigProp)) {
		obj["apiConfig"] = apiConfigProp
	}
	labelsProp, err := expandApiGatewayGatewayLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApiGatewayBasePath}}projects/{{project}}/locations/{{region}}/gateways?gatewayId={{gateway_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Gateway: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Gateway: %s", err)
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
		return fmt.Errorf("Error creating Gateway: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/gateways/{{gateway_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = ApiGatewayOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Gateway", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Gateway: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/gateways/{{gateway_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Gateway %q: %#v", d.Id(), res)

	return resourceApiGatewayGatewayRead(d, meta)
}

func resourceApiGatewayGatewayRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApiGatewayBasePath}}projects/{{project}}/locations/{{region}}/gateways/{{gateway_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Gateway: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ApiGatewayGateway %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}

	if err := d.Set("name", flattenApiGatewayGatewayName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("display_name", flattenApiGatewayGatewayDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("api_config", flattenApiGatewayGatewayApiConfig(res["apiConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("default_hostname", flattenApiGatewayGatewayDefaultHostname(res["defaultHostname"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("labels", flattenApiGatewayGatewayLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}

	return nil
}

func resourceApiGatewayGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Gateway: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandApiGatewayGatewayDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	apiConfigProp, err := expandApiGatewayGatewayApiConfig(d.Get("api_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("api_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, apiConfigProp)) {
		obj["apiConfig"] = apiConfigProp
	}
	labelsProp, err := expandApiGatewayGatewayLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApiGatewayBasePath}}projects/{{project}}/locations/{{region}}/gateways/{{gateway_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Gateway %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("api_config") {
		updateMask = append(updateMask, "apiConfig")
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
		return fmt.Errorf("Error updating Gateway %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Gateway %q: %#v", d.Id(), res)
	}

	err = ApiGatewayOperationWaitTime(
		config, res, project, "Updating Gateway", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceApiGatewayGatewayRead(d, meta)
}

func resourceApiGatewayGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Gateway: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ApiGatewayBasePath}}projects/{{project}}/locations/{{region}}/gateways/{{gateway_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Gateway %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Gateway")
	}

	err = ApiGatewayOperationWaitTime(
		config, res, project, "Deleting Gateway", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Gateway %q: %#v", d.Id(), res)
	return nil
}

func resourceApiGatewayGatewayImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>.+)/locations/(?P<region>.+)/gateways/(?P<gateway_id>.+)",
		"(?P<project>.+)/(?P<region>.+)/(?P<gateway_id>.+)",
		"(?P<region>.+)/(?P<gateway_id>.+)",
		"(?P<gateway_id>.+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/gateways/{{gateway_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApiGatewayGatewayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApiGatewayGatewayDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApiGatewayGatewayApiConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApiGatewayGatewayDefaultHostname(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApiGatewayGatewayLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandApiGatewayGatewayDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApiGatewayGatewayApiConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApiGatewayGatewayLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
