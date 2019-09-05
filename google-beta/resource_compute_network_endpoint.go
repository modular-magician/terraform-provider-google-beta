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

	"github.com/hashicorp/terraform/helper/schema"
	"google.golang.org/api/compute/v1"
)

func resourceComputeNetworkEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeNetworkEndpointCreate,
		Read:   resourceComputeNetworkEndpointRead,
		Delete: resourceComputeNetworkEndpointDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeNetworkEndpointImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"instance": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_endpoint_group": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"zone": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
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

func resourceComputeNetworkEndpointCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	instanceProp, err := expandComputeNetworkEndpointInstance(d.Get("instance"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("instance"); !isEmptyValue(reflect.ValueOf(instanceProp)) && (ok || !reflect.DeepEqual(v, instanceProp)) {
		obj["instance"] = instanceProp
	}
	portProp, err := expandComputeNetworkEndpointPort(d.Get("port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port"); !isEmptyValue(reflect.ValueOf(portProp)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}
	ipAddressProp, err := expandComputeNetworkEndpointIpAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_address"); !isEmptyValue(reflect.ValueOf(ipAddressProp)) && (ok || !reflect.DeepEqual(v, ipAddressProp)) {
		obj["ipAddress"] = ipAddressProp
	}

	obj, err = resourceComputeNetworkEndpointEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := replaceVars(d, config, "networkEndpoint/{{project}}/{{zone}}/{{network_endpoint_group}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/attachNetworkEndpoints")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new NetworkEndpoint: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating NetworkEndpoint: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{project}}/{{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}")
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
		config.clientCompute, op, project, "Creating NetworkEndpoint",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create NetworkEndpoint: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating NetworkEndpoint %q: %#v", d.Id(), res)

	return resourceComputeNetworkEndpointRead(d, meta)
}

func resourceComputeNetworkEndpointRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/listNetworkEndpoints")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "POST", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeNetworkEndpoint %q", d.Id()))
	}

	res, err = flattenNestedComputeNetworkEndpoint(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ComputeNetworkEndpoint because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	res, err = resourceComputeNetworkEndpointDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing ComputeNetworkEndpoint because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading NetworkEndpoint: %s", err)
	}

	if err := d.Set("instance", flattenComputeNetworkEndpointInstance(res["instance"], d)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpoint: %s", err)
	}
	if err := d.Set("port", flattenComputeNetworkEndpointPort(res["port"], d)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpoint: %s", err)
	}
	if err := d.Set("ip_address", flattenComputeNetworkEndpointIpAddress(res["ipAddress"], d)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpoint: %s", err)
	}

	return nil
}

func resourceComputeNetworkEndpointDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	lockName, err := replaceVars(d, config, "networkEndpoint/{{project}}/{{zone}}/{{network_endpoint_group}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/detachNetworkEndpoints")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	toDelete := make(map[string]interface{})
	instanceProp, err := expandComputeNetworkEndpointInstance(d.Get("instance"), d, config)
	if err != nil {
		return err
	}
	toDelete["instance"] = instanceProp

	portProp, err := expandComputeNetworkEndpointPort(d.Get("port"), d, config)
	if err != nil {
		return err
	}
	toDelete["port"] = portProp

	ipAddressProp, err := expandComputeNetworkEndpointIpAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return err
	}
	toDelete["ipAddress"] = ipAddressProp

	obj = map[string]interface{}{
		"networkEndpoints": []map[string]interface{}{toDelete},
	}
	log.Printf("[DEBUG] Deleting NetworkEndpoint %q", d.Id())

	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "NetworkEndpoint")
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting NetworkEndpoint",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting NetworkEndpoint %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeNetworkEndpointImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/networkEndpointGroups/(?P<network_endpoint_group>[^/]+)/(?P<instance>[^/]+)/(?P<ip_address>[^/]+)/(?P<port>[^/]+)",
		"(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<network_endpoint_group>[^/]+)/(?P<instance>[^/]+)/(?P<ip_address>[^/]+)/(?P<port>[^/]+)",
		"(?P<zone>[^/]+)/(?P<network_endpoint_group>[^/]+)/(?P<instance>[^/]+)/(?P<ip_address>[^/]+)/(?P<port>[^/]+)",
		"(?P<network_endpoint_group>[^/]+)/(?P<instance>[^/]+)/(?P<ip_address>[^/]+)/(?P<port>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{project}}/{{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeNetworkEndpointInstance(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeNetworkEndpointPort(v interface{}, d *schema.ResourceData) interface{} {
	// Handles int given in float64 format
	if floatVal, ok := v.(float64); ok {
		return int(floatVal)
	}
	return v
}

func flattenComputeNetworkEndpointIpAddress(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandComputeNetworkEndpointInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return GetResourceNameFromSelfLink(v.(string)), nil
}

func expandComputeNetworkEndpointPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkEndpointIpAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceComputeNetworkEndpointEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	wrappedReq := map[string]interface{}{
		"networkEndpoints": []interface{}{obj},
	}
	return wrappedReq, nil
}

func flattenNestedComputeNetworkEndpoint(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["items"]
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
		return nil, fmt.Errorf("expected list or map for value items. Actual value: %v", v)
	}

	_, item, err := resourceComputeNetworkEndpointFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceComputeNetworkEndpointFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	config := meta.(*Config)
	expectedInstance, err := expandComputeNetworkEndpointInstance(d.Get("instance"), d, config)
	if err != nil {
		return -1, nil, err
	}
	expectedIpAddress, err := expandComputeNetworkEndpointIpAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return -1, nil, err
	}
	expectedPort, err := expandComputeNetworkEndpointPort(d.Get("port"), d, config)
	if err != nil {
		return -1, nil, err
	}

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		// Decode list item before comparing.
		item, err := resourceComputeNetworkEndpointDecoder(d, meta, item)
		if err != nil {
			return -1, nil, err
		}

		itemInstance := flattenComputeNetworkEndpointInstance(item["instance"], d)
		if !reflect.DeepEqual(itemInstance, expectedInstance) {
			log.Printf("[DEBUG] Skipping item with instance= %#v, looking for %#v)", itemInstance, expectedInstance)
			continue
		}
		itemIpAddress := flattenComputeNetworkEndpointIpAddress(item["ipAddress"], d)
		if !reflect.DeepEqual(itemIpAddress, expectedIpAddress) {
			log.Printf("[DEBUG] Skipping item with ipAddress= %#v, looking for %#v)", itemIpAddress, expectedIpAddress)
			continue
		}
		itemPort := flattenComputeNetworkEndpointPort(item["port"], d)
		if !reflect.DeepEqual(itemPort, expectedPort) {
			log.Printf("[DEBUG] Skipping item with port= %#v, looking for %#v)", itemPort, expectedPort)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
func resourceComputeNetworkEndpointDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	v, ok := res["networkEndpoint"]
	if !ok || v == nil {
		return res, nil
	}

	return v.(map[string]interface{}), nil
}
