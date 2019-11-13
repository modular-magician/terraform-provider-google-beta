// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceMLEngineModel() *schema.Resource {
	return &schema.Resource{
		Create: resourceMLEngineModelCreate,
		Read:   resourceMLEngineModelRead,
		Delete: resourceMLEngineModelDelete,

		Importer: &schema.ResourceImporter{
			State: resourceMLEngineModelImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name specified for the model.`,
			},
			"default_version": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `The default version of the model. This version will be used to handle
prediction requests that do not specify a version.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `The name specified for the version when it was created.`,
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The description specified for the model when it was created.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: `One or more labels that you can add, to organize your models.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"online_prediction_console_logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging`,
			},
			"online_prediction_logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `If true, online prediction access logs are sent to StackDriver Logging.`,
			},
			"regions": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `The list of regions where the model is going to be deployed.
Currently only one region per model is supported`,
				MaxItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceMLEngineModelCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandMLEngineModelName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandMLEngineModelDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	defaultVersionProp, err := expandMLEngineModelDefaultVersion(d.Get("default_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_version"); !isEmptyValue(reflect.ValueOf(defaultVersionProp)) && (ok || !reflect.DeepEqual(v, defaultVersionProp)) {
		obj["defaultVersion"] = defaultVersionProp
	}
	regionsProp, err := expandMLEngineModelRegions(d.Get("regions"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("regions"); !isEmptyValue(reflect.ValueOf(regionsProp)) && (ok || !reflect.DeepEqual(v, regionsProp)) {
		obj["regions"] = regionsProp
	}
	onlinePredictionLoggingProp, err := expandMLEngineModelOnlinePredictionLogging(d.Get("online_prediction_logging"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("online_prediction_logging"); !isEmptyValue(reflect.ValueOf(onlinePredictionLoggingProp)) && (ok || !reflect.DeepEqual(v, onlinePredictionLoggingProp)) {
		obj["onlinePredictionLogging"] = onlinePredictionLoggingProp
	}
	onlinePredictionConsoleLoggingProp, err := expandMLEngineModelOnlinePredictionConsoleLogging(d.Get("online_prediction_console_logging"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("online_prediction_console_logging"); !isEmptyValue(reflect.ValueOf(onlinePredictionConsoleLoggingProp)) && (ok || !reflect.DeepEqual(v, onlinePredictionConsoleLoggingProp)) {
		obj["onlinePredictionConsoleLogging"] = onlinePredictionConsoleLoggingProp
	}
	labelsProp, err := expandMLEngineModelLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := replaceVars(d, config, "{{MLEngineBasePath}}projects/{{project}}/models")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Model: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Model: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/models/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Model %q: %#v", d.Id(), res)

	return resourceMLEngineModelRead(d, meta)
}

func resourceMLEngineModelRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{MLEngineBasePath}}projects/{{project}}/models/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("MLEngineModel %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Model: %s", err)
	}

	if err := d.Set("name", flattenMLEngineModelName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading Model: %s", err)
	}
	if err := d.Set("description", flattenMLEngineModelDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading Model: %s", err)
	}
	if err := d.Set("default_version", flattenMLEngineModelDefaultVersion(res["defaultVersion"], d)); err != nil {
		return fmt.Errorf("Error reading Model: %s", err)
	}
	if err := d.Set("regions", flattenMLEngineModelRegions(res["regions"], d)); err != nil {
		return fmt.Errorf("Error reading Model: %s", err)
	}
	if err := d.Set("online_prediction_logging", flattenMLEngineModelOnlinePredictionLogging(res["onlinePredictionLogging"], d)); err != nil {
		return fmt.Errorf("Error reading Model: %s", err)
	}
	if err := d.Set("online_prediction_console_logging", flattenMLEngineModelOnlinePredictionConsoleLogging(res["onlinePredictionConsoleLogging"], d)); err != nil {
		return fmt.Errorf("Error reading Model: %s", err)
	}
	if err := d.Set("labels", flattenMLEngineModelLabels(res["labels"], d)); err != nil {
		return fmt.Errorf("Error reading Model: %s", err)
	}

	return nil
}

func resourceMLEngineModelDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{MLEngineBasePath}}projects/{{project}}/models/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Model %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Model")
	}

	err = mLEngineOperationWaitTime(
		config, res, project, "Deleting Model",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Model %q: %#v", d.Id(), res)
	return nil
}

func resourceMLEngineModelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/models/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/models/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenMLEngineModelName(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenMLEngineModelDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMLEngineModelDefaultVersion(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["name"] =
		flattenMLEngineModelDefaultVersionName(original["name"], d)
	return []interface{}{transformed}
}
func flattenMLEngineModelDefaultVersionName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMLEngineModelRegions(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMLEngineModelOnlinePredictionLogging(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMLEngineModelOnlinePredictionConsoleLogging(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMLEngineModelLabels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandMLEngineModelName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelDefaultVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandMLEngineModelDefaultVersionName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
		transformed["name"] = transformedName
	}

	return transformed, nil
}

func expandMLEngineModelDefaultVersionName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelRegions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelOnlinePredictionLogging(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelOnlinePredictionConsoleLogging(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
