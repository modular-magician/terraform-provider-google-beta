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
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"google.golang.org/api/compute/v1"
)

func resourceComputeGlobalAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeGlobalAddressCreate,
		Read:   resourceComputeGlobalAddressRead,
		Update: resourceComputeGlobalAddressUpdate,
		Delete: resourceComputeGlobalAddressDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeGlobalAddressImport,
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
			"address": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"address_type": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateFunc:     validation.StringInSlice([]string{"EXTERNAL", "INTERNAL", ""}, false),
				DiffSuppressFunc: emptyOrDefaultStringSuppress("EXTERNAL"),
				Default:          "EXTERNAL",
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ip_version": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateFunc:     validation.StringInSlice([]string{"IPV4", "IPV6", ""}, false),
				DiffSuppressFunc: emptyOrDefaultStringSuppress("IPV4"),
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"network": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"prefix_length": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"purpose": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"VPC_PEERING", ""}, false),
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"label_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
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

func resourceComputeGlobalAddressCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	addressProp, err := expandComputeGlobalAddressAddress(d.Get("address"), d, config)
	if err != nil {
		return err
	} else if v := d.Get("address"); !isEmptyValue(reflect.ValueOf(addressProp)) {
		obj["address"] = addressProp
	}
	descriptionProp, err := expandComputeGlobalAddressDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v := d.Get("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeGlobalAddressName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v := d.Get("name"); !isEmptyValue(reflect.ValueOf(nameProp)) {
		obj["name"] = nameProp
	}
	labelsProp, err := expandComputeGlobalAddressLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v := d.Get("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) {
		obj["labels"] = labelsProp
	}
	labelFingerprintProp, err := expandComputeGlobalAddressLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return err
	} else if v := d.Get("label_fingerprint"); !isEmptyValue(reflect.ValueOf(labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	ipVersionProp, err := expandComputeGlobalAddressIpVersion(d.Get("ip_version"), d, config)
	if err != nil {
		return err
	} else if v := d.Get("ip_version"); !isEmptyValue(reflect.ValueOf(ipVersionProp)) {
		obj["ipVersion"] = ipVersionProp
	}
	prefixLengthProp, err := expandComputeGlobalAddressPrefixLength(d.Get("prefix_length"), d, config)
	if err != nil {
		return err
	} else if v := d.Get("prefix_length"); !isEmptyValue(reflect.ValueOf(prefixLengthProp)) {
		obj["prefixLength"] = prefixLengthProp
	}
	addressTypeProp, err := expandComputeGlobalAddressAddressType(d.Get("address_type"), d, config)
	if err != nil {
		return err
	} else if v := d.Get("address_type"); !isEmptyValue(reflect.ValueOf(addressTypeProp)) {
		obj["addressType"] = addressTypeProp
	}
	purposeProp, err := expandComputeGlobalAddressPurpose(d.Get("purpose"), d, config)
	if err != nil {
		return err
	} else if v := d.Get("purpose"); !isEmptyValue(reflect.ValueOf(purposeProp)) {
		obj["purpose"] = purposeProp
	}
	networkProp, err := expandComputeGlobalAddressNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v := d.Get("network"); !isEmptyValue(reflect.ValueOf(networkProp)) {
		obj["network"] = networkProp
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/global/addresses")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GlobalAddress: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating GlobalAddress: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
		config.clientCompute, op, project, "Creating GlobalAddress",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create GlobalAddress: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating GlobalAddress %q: %#v", d.Id(), res)

	if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		// Labels cannot be set in a create.  We'll have to set them here.
		err = resourceComputeGlobalAddressRead(d, meta)
		if err != nil {
			return err
		}

		obj := make(map[string]interface{})
		// d.Get("labels") will have been overridden by the Read call.
		labelsProp, err := expandComputeGlobalAddressLabels(v, d, config)
		if err != nil {
			return err
		}
		obj["labels"] = labelsProp
		labelFingerprintProp := d.Get("label_fingerprint")
		obj["labelFingerprint"] = labelFingerprintProp

		url, err = replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/global/addresses/{{name}}/setLabels")
		if err != nil {
			return err
		}
		res, err = sendRequest(config, "POST", url, obj)
		if err != nil {
			return fmt.Errorf("Error adding labels to ComputeGlobalAddress %q: %s", d.Id(), err)
		}

		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating ComputeGlobalAddress Labels",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

	}

	return resourceComputeGlobalAddressRead(d, meta)
}

func resourceComputeGlobalAddressRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/global/addresses/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeGlobalAddress %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}

	if err := d.Set("address", flattenComputeGlobalAddressAddress(res["address"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeGlobalAddressCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("description", flattenComputeGlobalAddressDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("name", flattenComputeGlobalAddressName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("labels", flattenComputeGlobalAddressLabels(res["labels"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("label_fingerprint", flattenComputeGlobalAddressLabelFingerprint(res["labelFingerprint"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("ip_version", flattenComputeGlobalAddressIpVersion(res["ipVersion"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("prefix_length", flattenComputeGlobalAddressPrefixLength(res["prefixLength"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("address_type", flattenComputeGlobalAddressAddressType(res["addressType"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("purpose", flattenComputeGlobalAddressPurpose(res["purpose"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("network", flattenComputeGlobalAddressNetwork(res["network"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}

	return nil
}

func resourceComputeGlobalAddressUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	d.Partial(true)

	if d.HasChange("labels") || d.HasChange("label_fingerprint") {
		obj := make(map[string]interface{})
		labelsProp, err := expandComputeGlobalAddressLabels(d.Get("labels"), d, config)
		if err != nil {
			return err
		} else if v := d.Get("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) {
			obj["labels"] = labelsProp
		}
		labelFingerprintProp, err := expandComputeGlobalAddressLabelFingerprint(d.Get("label_fingerprint"), d, config)
		if err != nil {
			return err
		} else if v := d.Get("label_fingerprint"); !isEmptyValue(reflect.ValueOf(labelFingerprintProp)) {
			obj["labelFingerprint"] = labelFingerprintProp
		}

		url, err := replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/global/addresses/{{name}}/setLabels")
		if err != nil {
			return err
		}
		res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating GlobalAddress %q: %s", d.Id(), err)
		}

		project, err := getProject(d, config)
		if err != nil {
			return err
		}
		op := &compute.Operation{}
		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating GlobalAddress",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("labels")
		d.SetPartial("label_fingerprint")
	}

	d.Partial(false)

	return resourceComputeGlobalAddressRead(d, meta)
}

func resourceComputeGlobalAddressDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/global/addresses/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting GlobalAddress %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "GlobalAddress")
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting GlobalAddress",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting GlobalAddress %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeGlobalAddressImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{"projects/(?P<project>[^/]+)/global/addresses/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeGlobalAddressAddress(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalAddressCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalAddressDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalAddressName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalAddressLabels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalAddressLabelFingerprint(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalAddressIpVersion(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalAddressPrefixLength(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeGlobalAddressAddressType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalAddressPurpose(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalAddressNetwork(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeGlobalAddressAddress(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressLabels(v interface{}, d *schema.ResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeGlobalAddressLabelFingerprint(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressIpVersion(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressPrefixLength(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressAddressType(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressPurpose(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressNetwork(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}
