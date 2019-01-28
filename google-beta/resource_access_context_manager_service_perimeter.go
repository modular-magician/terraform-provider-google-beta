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
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	accesscontextmanager "google.golang.org/api/accesscontextmanager/v1beta"
)

func resourceAccessContextManagerServicePerimeter() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessContextManagerServicePerimeterCreate,
		Read:   resourceAccessContextManagerServicePerimeterRead,
		Update: resourceAccessContextManagerServicePerimeterUpdate,
		Delete: resourceAccessContextManagerServicePerimeterDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessContextManagerServicePerimeterImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"perimeter_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"PERIMETER_TYPE_REGULAR", "PERIMETER_TYPE_BRIDGE", ""}, false),
				Default:      "PERIMETER_TYPE_REGULAR",
			},
			"status": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_levels": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"resources": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"restricted_services": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"unrestricted_services": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAccessContextManagerServicePerimeterCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	titleProp, err := expandAccessContextManagerServicePerimeterTitle(d.Get("title"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("title"); !isEmptyValue(reflect.ValueOf(titleProp)) && (ok || !reflect.DeepEqual(v, titleProp)) {
		obj["title"] = titleProp
	}
	descriptionProp, err := expandAccessContextManagerServicePerimeterDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	perimeterTypeProp, err := expandAccessContextManagerServicePerimeterPerimeterType(d.Get("perimeter_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("perimeter_type"); !isEmptyValue(reflect.ValueOf(perimeterTypeProp)) && (ok || !reflect.DeepEqual(v, perimeterTypeProp)) {
		obj["perimeterType"] = perimeterTypeProp
	}
	statusProp, err := expandAccessContextManagerServicePerimeterStatus(d.Get("status"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("status"); !isEmptyValue(reflect.ValueOf(statusProp)) && (ok || !reflect.DeepEqual(v, statusProp)) {
		obj["status"] = statusProp
	}
	parentProp, err := expandAccessContextManagerServicePerimeterParent(d.Get("parent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	nameProp, err := expandAccessContextManagerServicePerimeterName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	obj, err = resourceAccessContextManagerServicePerimeterEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://accesscontextmanager.googleapis.com/v1beta/{{parent}}/servicePerimeters")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ServicePerimeter: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ServicePerimeter: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	op := &accesscontextmanager.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := accessContextManagerOperationWaitTime(
		config.clientAccessContextManager, op, "Creating ServicePerimeter",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ServicePerimeter: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating ServicePerimeter %q: %#v", d.Id(), res)

	return resourceAccessContextManagerServicePerimeterRead(d, meta)
}

func resourceAccessContextManagerServicePerimeterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://accesscontextmanager.googleapis.com/v1beta/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AccessContextManagerServicePerimeter %q", d.Id()))
	}

	if err := d.Set("title", flattenAccessContextManagerServicePerimeterTitle(res["title"], d)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeter: %s", err)
	}
	if err := d.Set("description", flattenAccessContextManagerServicePerimeterDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeter: %s", err)
	}
	if err := d.Set("create_time", flattenAccessContextManagerServicePerimeterCreateTime(res["createTime"], d)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeter: %s", err)
	}
	if err := d.Set("update_time", flattenAccessContextManagerServicePerimeterUpdateTime(res["updateTime"], d)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeter: %s", err)
	}
	if err := d.Set("perimeter_type", flattenAccessContextManagerServicePerimeterPerimeterType(res["perimeterType"], d)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeter: %s", err)
	}
	if err := d.Set("status", flattenAccessContextManagerServicePerimeterStatus(res["status"], d)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeter: %s", err)
	}
	if err := d.Set("update_mask", flattenAccessContextManagerServicePerimeterUpdateMask(res["updateMask"], d)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeter: %s", err)
	}
	if err := d.Set("name", flattenAccessContextManagerServicePerimeterName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeter: %s", err)
	}

	return nil
}

func resourceAccessContextManagerServicePerimeterUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	titleProp, err := expandAccessContextManagerServicePerimeterTitle(d.Get("title"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("title"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, titleProp)) {
		obj["title"] = titleProp
	}
	descriptionProp, err := expandAccessContextManagerServicePerimeterDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	statusProp, err := expandAccessContextManagerServicePerimeterStatus(d.Get("status"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("status"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, statusProp)) {
		obj["status"] = statusProp
	}

	obj, err = resourceAccessContextManagerServicePerimeterEncoder(d, meta, obj)

	url, err := replaceVars(d, config, "https://accesscontextmanager.googleapis.com/v1beta/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ServicePerimeter %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("title") {
		updateMask = append(updateMask, "title")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("status") {
		updateMask = append(updateMask, "status")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating ServicePerimeter %q: %s", d.Id(), err)
	}

	op := &accesscontextmanager.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = accessContextManagerOperationWaitTime(
		config.clientAccessContextManager, op, "Updating ServicePerimeter",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceAccessContextManagerServicePerimeterRead(d, meta)
}

func resourceAccessContextManagerServicePerimeterDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://accesscontextmanager.googleapis.com/v1beta/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ServicePerimeter %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ServicePerimeter")
	}

	op := &accesscontextmanager.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = accessContextManagerOperationWaitTime(
		config.clientAccessContextManager, op, "Deleting ServicePerimeter",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ServicePerimeter %q: %#v", d.Id(), res)
	return nil
}

func resourceAccessContextManagerServicePerimeterImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats can't import ids with forward slashes in them.
	parseImportId([]string{"(?P<name>.+)"}, d, config)
	stringParts := strings.Split(d.Get("name").(string), "/")
	d.Set("parent", fmt.Sprintf("%s/%s", stringParts[0], stringParts[1]))
	return []*schema.ResourceData{d}, nil
}

func flattenAccessContextManagerServicePerimeterTitle(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimeterDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimeterCreateTime(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimeterUpdateTime(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimeterPerimeterType(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil || v.(string) == "" {
		return "PERIMETER_TYPE_REGULAR"
	}
	return v
}

func flattenAccessContextManagerServicePerimeterStatus(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["resources"] =
		flattenAccessContextManagerServicePerimeterStatusResources(original["resources"], d)
	transformed["access_levels"] =
		flattenAccessContextManagerServicePerimeterStatusAccessLevels(original["accessLevels"], d)
	transformed["unrestricted_services"] =
		flattenAccessContextManagerServicePerimeterStatusUnrestrictedServices(original["unrestrictedServices"], d)
	transformed["restricted_services"] =
		flattenAccessContextManagerServicePerimeterStatusRestrictedServices(original["restrictedServices"], d)
	return []interface{}{transformed}
}
func flattenAccessContextManagerServicePerimeterStatusResources(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimeterStatusAccessLevels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimeterStatusUnrestrictedServices(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimeterStatusRestrictedServices(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimeterUpdateMask(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimeterName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandAccessContextManagerServicePerimeterTitle(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterPerimeterType(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterStatus(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResources, err := expandAccessContextManagerServicePerimeterStatusResources(original["resources"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResources); val.IsValid() && !isEmptyValue(val) {
		transformed["resources"] = transformedResources
	}

	transformedAccessLevels, err := expandAccessContextManagerServicePerimeterStatusAccessLevels(original["access_levels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAccessLevels); val.IsValid() && !isEmptyValue(val) {
		transformed["accessLevels"] = transformedAccessLevels
	}

	transformedUnrestrictedServices, err := expandAccessContextManagerServicePerimeterStatusUnrestrictedServices(original["unrestricted_services"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUnrestrictedServices); val.IsValid() && !isEmptyValue(val) {
		transformed["unrestrictedServices"] = transformedUnrestrictedServices
	}

	transformedRestrictedServices, err := expandAccessContextManagerServicePerimeterStatusRestrictedServices(original["restricted_services"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRestrictedServices); val.IsValid() && !isEmptyValue(val) {
		transformed["restrictedServices"] = transformedRestrictedServices
	}

	return transformed, nil
}

func expandAccessContextManagerServicePerimeterStatusResources(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterStatusAccessLevels(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterStatusUnrestrictedServices(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterStatusRestrictedServices(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterParent(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceAccessContextManagerServicePerimeterEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "parent")
	return obj, nil
}
