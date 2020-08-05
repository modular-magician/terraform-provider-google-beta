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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceComputePacketMirroring() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputePacketMirroringCreate,
		Read:   resourceComputePacketMirroringRead,
		Update: resourceComputePacketMirroringUpdate,
		Delete: resourceComputePacketMirroringDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputePacketMirroringImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"collector_ilb": {
				Type:     schema.TypeList,
				Required: true,
				Description: `The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL)
that will be used as collector for mirrored traffic. The
specified forwarding rule must have is_mirroring_collector
set to true.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: compareSelfLinkOrResourceName,
							Description:      `The URL of the forwarding rule.`,
						},
					},
				},
			},
			"mirrored_resources": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `A means of specifying which resources to mirror.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instances": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `All the listed instances will be mirrored.  Specify at most 50.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: compareSelfLinkOrResourceName,
										Description:      `The URL of the instances where this rule should be active.`,
									},
								},
							},
							AtLeastOneOf: []string{"mirrored_resources.0.subnetworks", "mirrored_resources.0.instances", "mirrored_resources.0.tags"},
						},
						"subnetworks": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `All instances in one of these subnetworks will be mirrored.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: compareSelfLinkOrResourceName,
										Description:      `The URL of the subnetwork where this rule should be active.`,
									},
								},
							},
							AtLeastOneOf: []string{"mirrored_resources.0.subnetworks", "mirrored_resources.0.instances", "mirrored_resources.0.tags"},
						},
						"tags": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `All instances with these tags will be mirrored.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							AtLeastOneOf: []string{"mirrored_resources.0.subnetworks", "mirrored_resources.0.instances", "mirrored_resources.0.tags"},
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateGCPName,
				Description:  `The name of the packet mirroring rule`,
			},
			"network": {
				Type:     schema.TypeList,
				Required: true,
				Description: `Specifies the mirrored VPC network. Only packets in this network
will be mirrored. All mirrored VMs should have a NIC in the given
network. All mirrored subnetworks should belong to the given network.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: compareSelfLinkOrResourceName,
							Description:      `The full self_link URL of the network where this rule is active.`,
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `A human-readable description of the rule.`,
			},
			"filter": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `A filter for mirrored traffic.  If unset, all traffic is mirrored.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr_ranges": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `IP CIDR ranges that apply as a filter on the source (ingress) or
destination (egress) IP in the IP header. Only IPv4 is supported.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ip_protocols": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Protocols that apply as a filter on mirrored traffic. Possible values: ["tcp", "udp", "icmp"]`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"priority": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
				Description: `Since only one rule can be active at a time, priority is
used to break ties in the case of two rules that apply to
the same instances.`,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				Description: `The Region in which the created address should reside.
If it is not provided, the provider region is used.`,
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

func resourceComputePacketMirroringCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandComputePacketMirroringName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputePacketMirroringDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	regionProp, err := expandComputePacketMirroringRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}
	networkProp, err := expandComputePacketMirroringNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	priorityProp, err := expandComputePacketMirroringPriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !isEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	collectorIlbProp, err := expandComputePacketMirroringCollectorIlb(d.Get("collector_ilb"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("collector_ilb"); !isEmptyValue(reflect.ValueOf(collectorIlbProp)) && (ok || !reflect.DeepEqual(v, collectorIlbProp)) {
		obj["collectorIlb"] = collectorIlbProp
	}
	filterProp, err := expandComputePacketMirroringFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !isEmptyValue(reflect.ValueOf(filterProp)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}
	mirroredResourcesProp, err := expandComputePacketMirroringMirroredResources(d.Get("mirrored_resources"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("mirrored_resources"); !isEmptyValue(reflect.ValueOf(mirroredResourcesProp)) && (ok || !reflect.DeepEqual(v, mirroredResourcesProp)) {
		obj["mirroredResources"] = mirroredResourcesProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/packetMirrorings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new PacketMirroring: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating PacketMirroring: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating PacketMirroring",
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create PacketMirroring: %s", err)
	}

	log.Printf("[DEBUG] Finished creating PacketMirroring %q: %#v", d.Id(), res)

	return resourceComputePacketMirroringRead(d, meta)
}

func resourceComputePacketMirroringRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputePacketMirroring %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading PacketMirroring: %s", err)
	}

	if err := d.Set("name", flattenComputePacketMirroringName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading PacketMirroring: %s", err)
	}
	if err := d.Set("description", flattenComputePacketMirroringDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading PacketMirroring: %s", err)
	}
	if err := d.Set("region", flattenComputePacketMirroringRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading PacketMirroring: %s", err)
	}
	if err := d.Set("network", flattenComputePacketMirroringNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading PacketMirroring: %s", err)
	}
	if err := d.Set("priority", flattenComputePacketMirroringPriority(res["priority"], d, config)); err != nil {
		return fmt.Errorf("Error reading PacketMirroring: %s", err)
	}
	if err := d.Set("collector_ilb", flattenComputePacketMirroringCollectorIlb(res["collectorIlb"], d, config)); err != nil {
		return fmt.Errorf("Error reading PacketMirroring: %s", err)
	}
	if err := d.Set("filter", flattenComputePacketMirroringFilter(res["filter"], d, config)); err != nil {
		return fmt.Errorf("Error reading PacketMirroring: %s", err)
	}
	if err := d.Set("mirrored_resources", flattenComputePacketMirroringMirroredResources(res["mirroredResources"], d, config)); err != nil {
		return fmt.Errorf("Error reading PacketMirroring: %s", err)
	}

	return nil
}

func resourceComputePacketMirroringUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputePacketMirroringName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	regionProp, err := expandComputePacketMirroringRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}
	networkProp, err := expandComputePacketMirroringNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	priorityProp, err := expandComputePacketMirroringPriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	collectorIlbProp, err := expandComputePacketMirroringCollectorIlb(d.Get("collector_ilb"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("collector_ilb"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, collectorIlbProp)) {
		obj["collectorIlb"] = collectorIlbProp
	}
	filterProp, err := expandComputePacketMirroringFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}
	mirroredResourcesProp, err := expandComputePacketMirroringMirroredResources(d.Get("mirrored_resources"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("mirrored_resources"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, mirroredResourcesProp)) {
		obj["mirroredResources"] = mirroredResourcesProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating PacketMirroring %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating PacketMirroring %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating PacketMirroring %q: %#v", d.Id(), res)
	}

	err = computeOperationWaitTime(
		config, res, project, "Updating PacketMirroring",
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputePacketMirroringRead(d, meta)
}

func resourceComputePacketMirroringDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting PacketMirroring %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "PacketMirroring")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting PacketMirroring",
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting PacketMirroring %q: %#v", d.Id(), res)
	return nil
}

func resourceComputePacketMirroringImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/packetMirrorings/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputePacketMirroringName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputePacketMirroringDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputePacketMirroringRegion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenComputePacketMirroringNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["url"] =
		flattenComputePacketMirroringNetworkUrl(original["url"], d, config)
	return []interface{}{transformed}
}
func flattenComputePacketMirroringNetworkUrl(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputePacketMirroringPriority(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
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

func flattenComputePacketMirroringCollectorIlb(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["url"] =
		flattenComputePacketMirroringCollectorIlbUrl(original["url"], d, config)
	return []interface{}{transformed}
}
func flattenComputePacketMirroringCollectorIlbUrl(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputePacketMirroringFilter(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["ip_protocols"] =
		flattenComputePacketMirroringFilterIpProtocols(original["IPProtocols"], d, config)
	transformed["cidr_ranges"] =
		flattenComputePacketMirroringFilterCidrRanges(original["cidrRanges"], d, config)
	return []interface{}{transformed}
}
func flattenComputePacketMirroringFilterIpProtocols(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputePacketMirroringFilterCidrRanges(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputePacketMirroringMirroredResources(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["subnetworks"] =
		flattenComputePacketMirroringMirroredResourcesSubnetworks(original["subnetworks"], d, config)
	transformed["instances"] =
		flattenComputePacketMirroringMirroredResourcesInstances(original["instances"], d, config)
	transformed["tags"] =
		flattenComputePacketMirroringMirroredResourcesTags(original["tags"], d, config)
	return []interface{}{transformed}
}
func flattenComputePacketMirroringMirroredResourcesSubnetworks(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"url": flattenComputePacketMirroringMirroredResourcesSubnetworksUrl(original["url"], d, config),
		})
	}
	return transformed
}
func flattenComputePacketMirroringMirroredResourcesSubnetworksUrl(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputePacketMirroringMirroredResourcesInstances(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"url": flattenComputePacketMirroringMirroredResourcesInstancesUrl(original["url"], d, config),
		})
	}
	return transformed
}
func flattenComputePacketMirroringMirroredResourcesInstancesUrl(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputePacketMirroringMirroredResourcesTags(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandComputePacketMirroringName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputePacketMirroringDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputePacketMirroringRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputePacketMirroringNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUrl, err := expandComputePacketMirroringNetworkUrl(original["url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrl); val.IsValid() && !isEmptyValue(val) {
		transformed["url"] = transformedUrl
	}

	return transformed, nil
}

func expandComputePacketMirroringNetworkUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for url: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputePacketMirroringPriority(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputePacketMirroringCollectorIlb(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUrl, err := expandComputePacketMirroringCollectorIlbUrl(original["url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrl); val.IsValid() && !isEmptyValue(val) {
		transformed["url"] = transformedUrl
	}

	return transformed, nil
}

func expandComputePacketMirroringCollectorIlbUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("forwardingRules", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for url: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputePacketMirroringFilter(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIpProtocols, err := expandComputePacketMirroringFilterIpProtocols(original["ip_protocols"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpProtocols); val.IsValid() && !isEmptyValue(val) {
		transformed["IPProtocols"] = transformedIpProtocols
	}

	transformedCidrRanges, err := expandComputePacketMirroringFilterCidrRanges(original["cidr_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCidrRanges); val.IsValid() && !isEmptyValue(val) {
		transformed["cidrRanges"] = transformedCidrRanges
	}

	return transformed, nil
}

func expandComputePacketMirroringFilterIpProtocols(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputePacketMirroringFilterCidrRanges(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputePacketMirroringMirroredResources(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSubnetworks, err := expandComputePacketMirroringMirroredResourcesSubnetworks(original["subnetworks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubnetworks); val.IsValid() && !isEmptyValue(val) {
		transformed["subnetworks"] = transformedSubnetworks
	}

	transformedInstances, err := expandComputePacketMirroringMirroredResourcesInstances(original["instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstances); val.IsValid() && !isEmptyValue(val) {
		transformed["instances"] = transformedInstances
	}

	transformedTags, err := expandComputePacketMirroringMirroredResourcesTags(original["tags"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTags); val.IsValid() && !isEmptyValue(val) {
		transformed["tags"] = transformedTags
	}

	return transformed, nil
}

func expandComputePacketMirroringMirroredResourcesSubnetworks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedUrl, err := expandComputePacketMirroringMirroredResourcesSubnetworksUrl(original["url"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUrl); val.IsValid() && !isEmptyValue(val) {
			transformed["url"] = transformedUrl
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputePacketMirroringMirroredResourcesSubnetworksUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("subnetworks", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for url: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputePacketMirroringMirroredResourcesInstances(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedUrl, err := expandComputePacketMirroringMirroredResourcesInstancesUrl(original["url"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUrl); val.IsValid() && !isEmptyValue(val) {
			transformed["url"] = transformedUrl
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputePacketMirroringMirroredResourcesInstancesUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseZonalFieldValue("instances", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for url: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputePacketMirroringMirroredResourcesTags(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
