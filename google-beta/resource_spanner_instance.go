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
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func deleteSpannerBackups(d *schema.ResourceData, config *transport_tpg.Config, res map[string]interface{}, userAgent string, billingProject string) error {
	var v interface{}
	var ok bool

	v, ok = res["backups"]
	if !ok || v == nil {
		return nil
	}

	// Iterate over the list and delete each backup.
	for _, itemRaw := range v.([]interface{}) {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		backupName := item["name"].(string)

		log.Printf("[DEBUG] Found backups for resource %q: %#v)", d.Id(), item)

		path := "{{SpannerBasePath}}" + backupName

		url, err := tpgresource.ReplaceVars(d, config, path)
		if err != nil {
			return err
		}

		_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "DELETE",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func resourceSpannerInstanceVirtualUpdate(d *schema.ResourceData, resourceSchema map[string]*schema.Schema) bool {
	// force_destroy is the only virtual field
	if d.HasChange("force_destroy") {
		for field := range resourceSchema {
			if field == "force_destroy" {
				continue
			}
			if d.HasChange(field) {
				return false
			}
		}
		return true
	}
	return false
}

func ResourceSpannerInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpannerInstanceCreate,
		Read:   resourceSpannerInstanceRead,
		Update: resourceSpannerInstanceUpdate,
		Delete: resourceSpannerInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSpannerInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"config": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `The name of the instance's configuration (similar but not
quite the same as a region) which defines the geographic placement and
replication of your databases in this instance. It determines where your data
is stored. Values are typically of the form 'regional-europe-west1' , 'us-central' etc.
In order to obtain a valid list please consult the
[Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).`,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The descriptive name for this instance as it appears in UIs. Must be
unique per project and between 4 and 30 characters in length.`,
			},
			"name": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateRegexp(`^[a-z][-a-z0-9]*[a-z0-9]$`),
				Description: `A unique identifier for the instance, which cannot be changed after
the instance is created. The name must be between 6 and 30 characters
in length.


If not provided, a random string starting with 'tf-' will be selected.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `An object containing a list of "key": value pairs.
Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"num_nodes": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
				Description: `The number of nodes allocated to this instance. Exactly one of either node_count or processing_units
must be present in terraform.`,
				ExactlyOneOf: []string{"num_nodes", "processing_units"},
			},
			"processing_units": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
				Description: `The number of processing units allocated to this instance. Exactly one of processing_units
or node_count must be present in terraform.`,
				ExactlyOneOf: []string{"num_nodes", "processing_units"},
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Instance status: 'CREATING' or 'READY'.`,
			},
			"force_destroy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				Description: `When deleting a spanner instance, this boolean option will delete all backups of this instance.
This must be set to true if you created a backup manually in the console.`,
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

func resourceSpannerInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandSpannerInstanceName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	configProp, err := expandSpannerInstanceConfig(d.Get("config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("config"); !tpgresource.IsEmptyValue(reflect.ValueOf(configProp)) && (ok || !reflect.DeepEqual(v, configProp)) {
		obj["config"] = configProp
	}
	displayNameProp, err := expandSpannerInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	nodeCountProp, err := expandSpannerInstanceNumNodes(d.Get("num_nodes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("num_nodes"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeCountProp)) && (ok || !reflect.DeepEqual(v, nodeCountProp)) {
		obj["nodeCount"] = nodeCountProp
	}
	processingUnitsProp, err := expandSpannerInstanceProcessingUnits(d.Get("processing_units"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("processing_units"); !tpgresource.IsEmptyValue(reflect.ValueOf(processingUnitsProp)) && (ok || !reflect.DeepEqual(v, processingUnitsProp)) {
		obj["processingUnits"] = processingUnitsProp
	}
	labelsProp, err := expandSpannerInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	obj, err = resourceSpannerInstanceEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instances")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Instance: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
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
		return fmt.Errorf("Error creating Instance: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = SpannerOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Instance", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Instance: %s", err)
	}

	opRes, err = resourceSpannerInstanceDecoder(d, meta, opRes)
	if err != nil {
		return fmt.Errorf("Error decoding response from operation: %s", err)
	}
	if opRes == nil {
		return fmt.Errorf("Error decoding response from operation, could not find object")
	}

	if err := d.Set("name", flattenSpannerInstanceName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "{{project}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// This is useful if the resource in question doesn't have a perfectly consistent API
	// That is, the Operation for Create might return before the Get operation shows the
	// completed state of the resource.
	time.Sleep(5 * time.Second)

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceSpannerInstanceRead(d, meta)
}

func resourceSpannerInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instances/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SpannerInstance %q", d.Id()))
	}

	res, err = resourceSpannerInstanceDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing SpannerInstance because it no longer exists.")
		d.SetId("")
		return nil
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("force_destroy"); !ok {
		if err := d.Set("force_destroy", false); err != nil {
			return fmt.Errorf("Error setting force_destroy: %s", err)
		}
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	if err := d.Set("name", flattenSpannerInstanceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("config", flattenSpannerInstanceConfig(res["config"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("display_name", flattenSpannerInstanceDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("num_nodes", flattenSpannerInstanceNumNodes(res["nodeCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("processing_units", flattenSpannerInstanceProcessingUnits(res["processingUnits"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("labels", flattenSpannerInstanceLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("state", flattenSpannerInstanceState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	return nil
}

func resourceSpannerInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandSpannerInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	nodeCountProp, err := expandSpannerInstanceNumNodes(d.Get("num_nodes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("num_nodes"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nodeCountProp)) {
		obj["nodeCount"] = nodeCountProp
	}
	processingUnitsProp, err := expandSpannerInstanceProcessingUnits(d.Get("processing_units"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("processing_units"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, processingUnitsProp)) {
		obj["processingUnits"] = processingUnitsProp
	}
	labelsProp, err := expandSpannerInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	obj, err = resourceSpannerInstanceUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instances/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Instance %q: %#v", d.Id(), obj)
	if resourceSpannerInstanceVirtualUpdate(d, ResourceSpannerInstance().Schema) {
		if d.Get("force_destroy") != nil {
			if err := d.Set("force_destroy", d.Get("force_destroy")); err != nil {
				return fmt.Errorf("Error reading Instance: %s", err)
			}
		}
		return nil
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
		return fmt.Errorf("Error updating Instance %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Instance %q: %#v", d.Id(), res)
	}

	err = SpannerOperationWaitTime(
		config, res, project, "Updating Instance", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceSpannerInstanceRead(d, meta)
}

func resourceSpannerInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instances/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	if d.Get("force_destroy").(bool) {
		backupsUrl, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instances/{{name}}/backups")
		if err != nil {
			return err
		}

		resp, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "GET",
			Project:   billingProject,
			RawURL:    backupsUrl,
			UserAgent: userAgent,
		})
		if err != nil {
			// API returns 200 if no backups exist but the instance still exists, hence the error check.
			return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SpannerInstance %q", d.Id()))
		}

		err = deleteSpannerBackups(d, config, resp, billingProject, userAgent)
		if err != nil {
			return err
		}
	}
	log.Printf("[DEBUG] Deleting Instance %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Instance")
	}

	log.Printf("[DEBUG] Finished deleting Instance %q: %#v", d.Id(), res)
	return nil
}

func resourceSpannerInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>.+)/instances/(?P<name>.+)",
		"(?P<project>.+)/(?P<name>.+)",
		"(?P<name>.+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Explicitly set virtual fields to default values on import
	if err := d.Set("force_destroy", false); err != nil {
		return nil, fmt.Errorf("Error setting force_destroy: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenSpannerInstanceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenSpannerInstanceConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenSpannerInstanceDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSpannerInstanceNumNodes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenSpannerInstanceProcessingUnits(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenSpannerInstanceLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSpannerInstanceState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandSpannerInstanceName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	r := regexp.MustCompile("projects/(.+)/instanceConfigs/(.+)")
	if r.MatchString(v.(string)) {
		return v.(string), nil
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("projects/%s/instanceConfigs/%s", project, v.(string)), nil
}

func expandSpannerInstanceDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceNumNodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceProcessingUnits(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func resourceSpannerInstanceEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Temp Logic to accommodate processing_units and num_nodes
	if obj["processingUnits"] == nil && obj["nodeCount"] == nil {
		obj["nodeCount"] = 1
	}
	newObj := make(map[string]interface{})
	newObj["instance"] = obj
	if obj["name"] == nil {
		if err := d.Set("name", resource.PrefixedUniqueId("tfgen-spanid-")[:30]); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
		newObj["instanceId"] = d.Get("name").(string)
	} else {
		newObj["instanceId"] = obj["name"]
	}
	delete(obj, "name")
	return newObj, nil
}

func resourceSpannerInstanceUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	project, err := tpgresource.GetProject(d, meta.(*transport_tpg.Config))
	if err != nil {
		return nil, err
	}
	obj["name"] = fmt.Sprintf("projects/%s/instances/%s", project, obj["name"])
	newObj := make(map[string]interface{})
	newObj["instance"] = obj
	updateMask := make([]string, 0)
	if d.HasChange("num_nodes") {
		updateMask = append(updateMask, "nodeCount")
	}
	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}
	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}
	if d.HasChange("processing_units") {
		updateMask = append(updateMask, "processingUnits")
	}
	newObj["fieldMask"] = strings.Join(updateMask, ",")
	return newObj, nil
}

func resourceSpannerInstanceDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	d.SetId(res["name"].(string))
	if err := tpgresource.ParseImportId([]string{"projects/(?P<project>[^/]+)/instances/(?P<name>[^/]+)"}, d, config); err != nil {
		return nil, err
	}
	res["project"] = d.Get("project").(string)
	res["name"] = d.Get("name").(string)
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{name}}")
	if err != nil {
		return nil, err
	}
	d.SetId(id)
	return res, nil
}
