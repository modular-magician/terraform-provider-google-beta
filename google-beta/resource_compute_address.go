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
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"google.golang.org/api/compute/v1"
)

func resourceComputeAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeAddressCreate,
		Read:   resourceComputeAddressRead,
		Update: resourceComputeAddressUpdate,
		Delete: resourceComputeAddressDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeAddressImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRegexp(`^(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)$`),
			},
			"address": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"address_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"INTERNAL", "EXTERNAL", ""}, false),
				Default:      "EXTERNAL",
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"network_tier": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"PREMIUM", "STANDARD", ""}, false),
			},
			"purpose": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"GCE_ENDPOINT", ""}, false),
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"subnetwork": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"label_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"users": {
				Type:     schema.TypeList,
				Computed: true,
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
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeAddressCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	addressProp, err := expandComputeAddressAddress(d.Get("address"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("address"); !isEmptyValue(reflect.ValueOf(addressProp)) && (ok || !reflect.DeepEqual(v, addressProp)) {
		obj["address"] = addressProp
	}
	addressTypeProp, err := expandComputeAddressAddressType(d.Get("address_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("address_type"); !isEmptyValue(reflect.ValueOf(addressTypeProp)) && (ok || !reflect.DeepEqual(v, addressTypeProp)) {
		obj["addressType"] = addressTypeProp
	}
	descriptionProp, err := expandComputeAddressDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeAddressName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	purposeProp, err := expandComputeAddressPurpose(d.Get("purpose"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("purpose"); !isEmptyValue(reflect.ValueOf(purposeProp)) && (ok || !reflect.DeepEqual(v, purposeProp)) {
		obj["purpose"] = purposeProp
	}
	networkTierProp, err := expandComputeAddressNetworkTier(d.Get("network_tier"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_tier"); !isEmptyValue(reflect.ValueOf(networkTierProp)) && (ok || !reflect.DeepEqual(v, networkTierProp)) {
		obj["networkTier"] = networkTierProp
	}
	subnetworkProp, err := expandComputeAddressSubnetwork(d.Get("subnetwork"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("subnetwork"); !isEmptyValue(reflect.ValueOf(subnetworkProp)) && (ok || !reflect.DeepEqual(v, subnetworkProp)) {
		obj["subnetwork"] = subnetworkProp
	}
	labelsProp, err := expandComputeAddressLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	labelFingerprintProp, err := expandComputeAddressLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("label_fingerprint"); !isEmptyValue(reflect.ValueOf(labelFingerprintProp)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	regionProp, err := expandComputeAddressRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/addresses")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Address: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Address: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/addresses/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
		config.clientCompute, op, project, "Creating Address",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Address: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating Address %q: %#v", d.Id(), res)

	if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		// Labels cannot be set in a create.  We'll have to set them here.
		err = resourceComputeAddressRead(d, meta)
		if err != nil {
			return err
		}

		obj := make(map[string]interface{})
		// d.Get("labels") will have been overridden by the Read call.
		labelsProp, err := expandComputeAddressLabels(v, d, config)
		if err != nil {
			return err
		}
		obj["labels"] = labelsProp
		labelFingerprintProp := d.Get("label_fingerprint")
		obj["labelFingerprint"] = labelFingerprintProp

		url, err = replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/addresses/{{name}}/setLabels")
		if err != nil {
			return err
		}
		res, err = sendRequest(config, "POST", project, url, obj)
		if err != nil {
			return fmt.Errorf("Error adding labels to ComputeAddress %q: %s", d.Id(), err)
		}

		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating ComputeAddress Labels",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

	}

	return resourceComputeAddressRead(d, meta)
}

func resourceComputeAddressRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/addresses/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeAddress %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Address: %s", err)
	}

	if err := d.Set("address", flattenComputeAddressAddress(res["address"], d)); err != nil {
		return fmt.Errorf("Error reading Address: %s", err)
	}
	if err := d.Set("address_type", flattenComputeAddressAddressType(res["addressType"], d)); err != nil {
		return fmt.Errorf("Error reading Address: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeAddressCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading Address: %s", err)
	}
	if err := d.Set("description", flattenComputeAddressDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading Address: %s", err)
	}
	if err := d.Set("name", flattenComputeAddressName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading Address: %s", err)
	}
	if err := d.Set("purpose", flattenComputeAddressPurpose(res["purpose"], d)); err != nil {
		return fmt.Errorf("Error reading Address: %s", err)
	}
	if err := d.Set("network_tier", flattenComputeAddressNetworkTier(res["networkTier"], d)); err != nil {
		return fmt.Errorf("Error reading Address: %s", err)
	}
	if err := d.Set("subnetwork", flattenComputeAddressSubnetwork(res["subnetwork"], d)); err != nil {
		return fmt.Errorf("Error reading Address: %s", err)
	}
	if err := d.Set("users", flattenComputeAddressUsers(res["users"], d)); err != nil {
		return fmt.Errorf("Error reading Address: %s", err)
	}
	if err := d.Set("labels", flattenComputeAddressLabels(res["labels"], d)); err != nil {
		return fmt.Errorf("Error reading Address: %s", err)
	}
	if err := d.Set("label_fingerprint", flattenComputeAddressLabelFingerprint(res["labelFingerprint"], d)); err != nil {
		return fmt.Errorf("Error reading Address: %s", err)
	}
	if err := d.Set("region", flattenComputeAddressRegion(res["region"], d)); err != nil {
		return fmt.Errorf("Error reading Address: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Address: %s", err)
	}

	return nil
}

func resourceComputeAddressUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	d.Partial(true)

	if d.HasChange("labels") || d.HasChange("label_fingerprint") {
		obj := make(map[string]interface{})
		labelsProp, err := expandComputeAddressLabels(d.Get("labels"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
			obj["labels"] = labelsProp
		}
		labelFingerprintProp, err := expandComputeAddressLabelFingerprint(d.Get("label_fingerprint"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("label_fingerprint"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
			obj["labelFingerprint"] = labelFingerprintProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/addresses/{{name}}/setLabels")
		if err != nil {
			return err
		}
		res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating Address %q: %s", d.Id(), err)
		}

		op := &compute.Operation{}
		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating Address",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("labels")
		d.SetPartial("label_fingerprint")
	}

	d.Partial(false)

	return resourceComputeAddressRead(d, meta)
}

func resourceComputeAddressDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/addresses/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Address %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Address")
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting Address",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Address %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeAddressImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/addresses/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/addresses/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeAddressAddress(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeAddressAddressType(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil || v.(string) == "" {
		return "EXTERNAL"
	}
	return v
}

func flattenComputeAddressCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeAddressDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeAddressName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeAddressPurpose(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeAddressNetworkTier(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeAddressSubnetwork(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeAddressUsers(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeAddressLabels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeAddressLabelFingerprint(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeAddressRegion(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func expandComputeAddressAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAddressAddressType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAddressDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAddressName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAddressPurpose(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAddressNetworkTier(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAddressSubnetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("subnetworks", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for subnetwork: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeAddressLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeAddressLabelFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAddressRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
