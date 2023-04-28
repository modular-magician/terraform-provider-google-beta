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
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceGameServicesGameServerDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceGameServicesGameServerDeploymentCreate,
		Read:   resourceGameServicesGameServerDeploymentRead,
		Update: resourceGameServicesGameServerDeploymentUpdate,
		Delete: resourceGameServicesGameServerDeploymentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGameServicesGameServerDeploymentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"deployment_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `A unique id for the deployment.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Human readable description of the game server deployment.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `The labels associated with this game server deployment. Each label is a
key-value pair.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Location of the Deployment.`,
				Default:     "global",
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource id of the game server deployment, eg:

'projects/{project_id}/locations/{location}/gameServerDeployments/{deployment_id}'.
For example,

'projects/my-project/locations/{location}/gameServerDeployments/my-deployment'.`,
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

func resourceGameServicesGameServerDeploymentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandGameServicesGameServerDeploymentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandGameServicesGameServerDeploymentLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := ReplaceVars(d, config, "{{GameServicesBasePath}}projects/{{project}}/locations/{{location}}/gameServerDeployments?deploymentId={{deployment_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GameServerDeployment: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GameServerDeployment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating GameServerDeployment: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = GameServicesOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating GameServerDeployment", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create GameServerDeployment: %s", err)
	}

	if err := d.Set("name", flattenGameServicesGameServerDeploymentName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating GameServerDeployment %q: %#v", d.Id(), res)

	return resourceGameServicesGameServerDeploymentRead(d, meta)
}

func resourceGameServicesGameServerDeploymentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{GameServicesBasePath}}projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GameServerDeployment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("GameServicesGameServerDeployment %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading GameServerDeployment: %s", err)
	}

	if err := d.Set("name", flattenGameServicesGameServerDeploymentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading GameServerDeployment: %s", err)
	}
	if err := d.Set("description", flattenGameServicesGameServerDeploymentDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading GameServerDeployment: %s", err)
	}
	if err := d.Set("labels", flattenGameServicesGameServerDeploymentLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading GameServerDeployment: %s", err)
	}

	return nil
}

func resourceGameServicesGameServerDeploymentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GameServerDeployment: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandGameServicesGameServerDeploymentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandGameServicesGameServerDeploymentLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := ReplaceVars(d, config, "{{GameServicesBasePath}}projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating GameServerDeployment %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
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
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating GameServerDeployment %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating GameServerDeployment %q: %#v", d.Id(), res)
	}

	err = GameServicesOperationWaitTime(
		config, res, project, "Updating GameServerDeployment", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceGameServicesGameServerDeploymentRead(d, meta)
}

func resourceGameServicesGameServerDeploymentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GameServerDeployment: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{GameServicesBasePath}}projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting GameServerDeployment %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "GameServerDeployment")
	}

	err = GameServicesOperationWaitTime(
		config, res, project, "Deleting GameServerDeployment", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting GameServerDeployment %q: %#v", d.Id(), res)
	return nil
}

func resourceGameServicesGameServerDeploymentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/gameServerDeployments/(?P<deployment_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<deployment_id>[^/]+)",
		"(?P<location>[^/]+)/(?P<deployment_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenGameServicesGameServerDeploymentName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGameServicesGameServerDeploymentDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGameServicesGameServerDeploymentLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandGameServicesGameServerDeploymentDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerDeploymentLabels(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
