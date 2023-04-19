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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceServiceUsageConsumerQuotaOverride() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceUsageConsumerQuotaOverrideCreate,
		Read:   resourceServiceUsageConsumerQuotaOverrideRead,
		Update: resourceServiceUsageConsumerQuotaOverrideUpdate,
		Delete: resourceServiceUsageConsumerQuotaOverrideDelete,

		Importer: &schema.ResourceImporter{
			State: resourceServiceUsageConsumerQuotaOverrideImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"limit": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The limit on the metric, e.g. '/project/region'.

~> Make sure that 'limit' is in a format that doesn't start with '1/' or contain curly braces.
E.g. use '/project/user' instead of '1/{project}/{user}'.`,
			},
			"metric": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The metric that should be limited, e.g. 'compute.googleapis.com/cpus'.`,
			},
			"override_value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The overriding quota limit value. Can be any nonnegative integer, or -1 (unlimited quota).`,
			},
			"service": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The service that the metrics belong to, e.g. 'compute.googleapis.com'.`,
			},
			"dimensions": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: `If this map is nonempty, then this override applies only to specific values for dimensions defined in the limit unit.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"force": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `If the new quota would decrease the existing quota by more than 10%, the request is rejected.
If 'force' is 'true', that safety check is ignored.`,
				Default: false,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The server-generated name of the quota override.`,
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

func resourceServiceUsageConsumerQuotaOverrideCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	overrideValueProp, err := expandNestedServiceUsageConsumerQuotaOverrideOverrideValue(d.Get("override_value"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("override_value"); !isEmptyValue(reflect.ValueOf(overrideValueProp)) && (ok || !reflect.DeepEqual(v, overrideValueProp)) {
		obj["overrideValue"] = overrideValueProp
	}
	dimensionsProp, err := expandNestedServiceUsageConsumerQuotaOverrideDimensions(d.Get("dimensions"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dimensions"); !isEmptyValue(reflect.ValueOf(dimensionsProp)) && (ok || !reflect.DeepEqual(v, dimensionsProp)) {
		obj["dimensions"] = dimensionsProp
	}

	url, err := ReplaceVars(d, config, "{{ServiceUsageBasePath}}projects/{{project}}/services/{{service}}/consumerQuotaMetrics/{{metric}}/limits/{{limit}}/consumerOverrides/{{name}}?force={{force}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ConsumerQuotaOverride: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ConsumerQuotaOverride: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ConsumerQuotaOverride: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/services/{{service}}/consumerQuotaMetrics/{{metric}}/limits/{{limit}}/consumerOverrides/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = ServiceUsageOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating ConsumerQuotaOverride", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ConsumerQuotaOverride: %s", err)
	}

	if _, ok := opRes["overrides"]; ok {
		opRes, err = flattenNestedServiceUsageConsumerQuotaOverride(d, meta, opRes)
		if err != nil {
			return fmt.Errorf("Error getting nested object from operation response: %s", err)
		}
		if opRes == nil {
			// Object isn't there any more - remove it from the state.
			return fmt.Errorf("Error decoding response from operation, could not find nested object")
		}
	}
	if err := d.Set("name", flattenNestedServiceUsageConsumerQuotaOverrideName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "projects/{{project}}/services/{{service}}/consumerQuotaMetrics/{{metric}}/limits/{{limit}}/consumerOverrides/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ConsumerQuotaOverride %q: %#v", d.Id(), res)

	return resourceServiceUsageConsumerQuotaOverrideRead(d, meta)
}

func resourceServiceUsageConsumerQuotaOverrideRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ServiceUsageBasePath}}projects/{{project}}/services/{{service}}/consumerQuotaMetrics/{{metric}}/limits/{{limit}}/consumerOverrides/")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ConsumerQuotaOverride: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ServiceUsageConsumerQuotaOverride %q", d.Id()))
	}

	res, err = flattenNestedServiceUsageConsumerQuotaOverride(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ServiceUsageConsumerQuotaOverride because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ConsumerQuotaOverride: %s", err)
	}

	if err := d.Set("override_value", flattenNestedServiceUsageConsumerQuotaOverrideOverrideValue(res["overrideValue"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConsumerQuotaOverride: %s", err)
	}
	if err := d.Set("dimensions", flattenNestedServiceUsageConsumerQuotaOverrideDimensions(res["dimensions"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConsumerQuotaOverride: %s", err)
	}
	if err := d.Set("name", flattenNestedServiceUsageConsumerQuotaOverrideName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConsumerQuotaOverride: %s", err)
	}

	return nil
}

func resourceServiceUsageConsumerQuotaOverrideUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ConsumerQuotaOverride: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	overrideValueProp, err := expandNestedServiceUsageConsumerQuotaOverrideOverrideValue(d.Get("override_value"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("override_value"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, overrideValueProp)) {
		obj["overrideValue"] = overrideValueProp
	}

	url, err := ReplaceVars(d, config, "{{ServiceUsageBasePath}}projects/{{project}}/services/{{service}}/consumerQuotaMetrics/{{metric}}/limits/{{limit}}/consumerOverrides/{{name}}?force={{force}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ConsumerQuotaOverride %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating ConsumerQuotaOverride %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ConsumerQuotaOverride %q: %#v", d.Id(), res)
	}

	err = ServiceUsageOperationWaitTime(
		config, res, project, "Updating ConsumerQuotaOverride", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceServiceUsageConsumerQuotaOverrideRead(d, meta)
}

func resourceServiceUsageConsumerQuotaOverrideDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ConsumerQuotaOverride: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{ServiceUsageBasePath}}projects/{{project}}/services/{{service}}/consumerQuotaMetrics/{{metric}}/limits/{{limit}}/consumerOverrides/{{name}}?force={{force}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ConsumerQuotaOverride %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ConsumerQuotaOverride")
	}

	err = ServiceUsageOperationWaitTime(
		config, res, project, "Deleting ConsumerQuotaOverride", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ConsumerQuotaOverride %q: %#v", d.Id(), res)
	return nil
}

func resourceServiceUsageConsumerQuotaOverrideImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/services/(?P<service>[^/]+)/consumerQuotaMetrics/(?P<metric>[^/]+)/limits/(?P<limit>[^/]+)/consumerOverrides/(?P<name>[^/]+)",
		"services/(?P<service>[^/]+)/consumerQuotaMetrics/(?P<metric>[^/]+)/limits/(?P<limit>[^/]+)/consumerOverrides/(?P<name>[^/]+)",
		"(?P<service>[^/]+)/(?P<metric>[^/]+)/(?P<limit>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/services/{{service}}/consumerQuotaMetrics/{{metric}}/limits/{{limit}}/consumerOverrides/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNestedServiceUsageConsumerQuotaOverrideOverrideValue(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return "0"
	}
	return v
}

func flattenNestedServiceUsageConsumerQuotaOverrideDimensions(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedServiceUsageConsumerQuotaOverrideName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func expandNestedServiceUsageConsumerQuotaOverrideOverrideValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedServiceUsageConsumerQuotaOverrideDimensions(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func flattenNestedServiceUsageConsumerQuotaOverride(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["overrides"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value overrides. Actual value: %v", v)
	}

	_, item, err := resourceServiceUsageConsumerQuotaOverrideFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceServiceUsageConsumerQuotaOverrideFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedName := d.Get("name")
	expectedFlattenedName := flattenNestedServiceUsageConsumerQuotaOverrideName(expectedName, d, meta.(*Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		itemName := flattenNestedServiceUsageConsumerQuotaOverrideName(item["name"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemName)) && isEmptyValue(reflect.ValueOf(expectedFlattenedName))) && !reflect.DeepEqual(itemName, expectedFlattenedName) {
			log.Printf("[DEBUG] Skipping item with name= %#v, looking for %#v)", itemName, expectedFlattenedName)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
