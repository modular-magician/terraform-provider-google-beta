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
)

func ResourceWorkstationsWorkstation() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkstationsWorkstationCreate,
		Read:   resourceWorkstationsWorkstationRead,
		Update: resourceWorkstationsWorkstationUpdate,
		Delete: resourceWorkstationsWorkstationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceWorkstationsWorkstationImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location where the workstation cluster config should reside.`,
			},
			"workstation_cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the workstation cluster.`,
			},
			"workstation_config_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the workstation cluster config.`,
			},
			"workstation_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `ID to use for the workstation.`,
			},
			"annotations": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Client-specified annotations. This is distinct from labels.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Human-readable name for this resource.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the Instance was created in UTC.`,
			},
			"host": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Host to which clients can send HTTPS traffic that will be received by the workstation. 
Authorized traffic will be received to the workstation as HTTP on port 80. 
To send traffic to a different port, clients may prefix the host with the destination port in the format "{port}-{host}".`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the cluster resource.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Current state of the workstation.`,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The system-generated UID of the resource.`,
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

func resourceWorkstationsWorkstationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandWorkstationsWorkstationDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	labelsProp, err := expandWorkstationsWorkstationLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandWorkstationsWorkstationAnnotations(d.Get("annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("annotations"); !isEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	url, err := replaceVars(d, config, "{{WorkstationsBasePath}}projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations?workstationId={{workstation_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Workstation: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Workstation: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Workstation: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = workstationsOperationWaitTime(
		config, res, project, "Creating Workstation", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Workstation: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Workstation %q: %#v", d.Id(), res)

	return resourceWorkstationsWorkstationRead(d, meta)
}

func resourceWorkstationsWorkstationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{WorkstationsBasePath}}projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Workstation: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("WorkstationsWorkstation %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Workstation: %s", err)
	}

	if err := d.Set("name", flattenWorkstationsWorkstationName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workstation: %s", err)
	}
	if err := d.Set("uid", flattenWorkstationsWorkstationUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workstation: %s", err)
	}
	if err := d.Set("display_name", flattenWorkstationsWorkstationDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workstation: %s", err)
	}
	if err := d.Set("labels", flattenWorkstationsWorkstationLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workstation: %s", err)
	}
	if err := d.Set("annotations", flattenWorkstationsWorkstationAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workstation: %s", err)
	}
	if err := d.Set("create_time", flattenWorkstationsWorkstationCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workstation: %s", err)
	}
	if err := d.Set("host", flattenWorkstationsWorkstationHost(res["host"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workstation: %s", err)
	}
	if err := d.Set("state", flattenWorkstationsWorkstationState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workstation: %s", err)
	}

	return nil
}

func resourceWorkstationsWorkstationUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Workstation: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandWorkstationsWorkstationDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	labelsProp, err := expandWorkstationsWorkstationLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandWorkstationsWorkstationAnnotations(d.Get("annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("annotations"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	url, err := replaceVars(d, config, "{{WorkstationsBasePath}}projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Workstation %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("annotations") {
		updateMask = append(updateMask, "annotations")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Workstation %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Workstation %q: %#v", d.Id(), res)
	}

	err = workstationsOperationWaitTime(
		config, res, project, "Updating Workstation", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceWorkstationsWorkstationRead(d, meta)
}

func resourceWorkstationsWorkstationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Workstation: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{WorkstationsBasePath}}projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Workstation %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Workstation")
	}

	err = workstationsOperationWaitTime(
		config, res, project, "Deleting Workstation", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Workstation %q: %#v", d.Id(), res)
	return nil
}

func resourceWorkstationsWorkstationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/workstationClusters/(?P<workstation_cluster_id>[^/]+)/workstationConfigs/(?P<workstation_config_id>[^/]+)/workstations/(?P<workstation_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<workstation_cluster_id>[^/]+)/(?P<workstation_config_id>[^/]+)/(?P<workstation_id>[^/]+)",
		"(?P<location>[^/]+)/(?P<workstation_cluster_id>[^/]+)/(?P<workstation_config_id>[^/]+)/(?P<workstation_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenWorkstationsWorkstationName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationUid(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationAnnotations(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationHost(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenWorkstationsWorkstationState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandWorkstationsWorkstationDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandWorkstationsWorkstationAnnotations(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
