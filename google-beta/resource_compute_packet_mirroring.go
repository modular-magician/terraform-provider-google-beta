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

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceComputePacketMirroring() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputePacketMirroringCreate,
		Read:   resourceComputePacketMirroringRead,
		Update: resourceComputePacketMirroringUpdate,
		Delete: resourceComputePacketMirroringDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputePacketMirroringImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
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
				ValidateFunc: verify.ValidateGCEName,
				Description:  `The name of the packet mirroring rule`,
			},
			"network": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Description: `Specifies the mirrored VPC network. Only packets in this network
will be mirrored. All mirrored VMs should have a NIC in the given
network. All mirrored subnetworks should belong to the given network.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
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
						"direction": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"INGRESS", "EGRESS", "BOTH", ""}),
							Description:  `Direction of traffic to mirror. Default value: "BOTH" Possible values: ["INGRESS", "EGRESS", "BOTH"]`,
							Default:      "BOTH",
						},
						"ip_protocols": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Possible IP protocols including tcp, udp, icmp and esp`,
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
		UseJSONNumber: true,
	}
}

func resourceComputePacketMirroringCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputePacketMirroringName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputePacketMirroringDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	regionProp, err := expandComputePacketMirroringRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}
	networkProp, err := expandComputePacketMirroringNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	priorityProp, err := expandComputePacketMirroringPriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	collectorIlbProp, err := expandComputePacketMirroringCollectorIlb(d.Get("collector_ilb"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("collector_ilb"); !tpgresource.IsEmptyValue(reflect.ValueOf(collectorIlbProp)) && (ok || !reflect.DeepEqual(v, collectorIlbProp)) {
		obj["collectorIlb"] = collectorIlbProp
	}
	filterProp, err := expandComputePacketMirroringFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(filterProp)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}
	mirroredResourcesProp, err := expandComputePacketMirroringMirroredResources(d.Get("mirrored_resources"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("mirrored_resources"); !tpgresource.IsEmptyValue(reflect.ValueOf(mirroredResourcesProp)) && (ok || !reflect.DeepEqual(v, mirroredResourcesProp)) {
		obj["mirroredResources"] = mirroredResourcesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/packetMirrorings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new PacketMirroring: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PacketMirroring: %s", err)
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
		return fmt.Errorf("Error creating PacketMirroring: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating PacketMirroring", userAgent,
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
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PacketMirroring: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputePacketMirroring %q", d.Id()))
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
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PacketMirroring: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	nameProp, err := expandComputePacketMirroringName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	regionProp, err := expandComputePacketMirroringRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}
	priorityProp, err := expandComputePacketMirroringPriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	collectorIlbProp, err := expandComputePacketMirroringCollectorIlb(d.Get("collector_ilb"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("collector_ilb"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, collectorIlbProp)) {
		obj["collectorIlb"] = collectorIlbProp
	}
	filterProp, err := expandComputePacketMirroringFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}
	mirroredResourcesProp, err := expandComputePacketMirroringMirroredResources(d.Get("mirrored_resources"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("mirrored_resources"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, mirroredResourcesProp)) {
		obj["mirroredResources"] = mirroredResourcesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating PacketMirroring %q: %#v", d.Id(), obj)

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
		return fmt.Errorf("Error updating PacketMirroring %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating PacketMirroring %q: %#v", d.Id(), res)
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Updating PacketMirroring", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputePacketMirroringRead(d, meta)
}

func resourceComputePacketMirroringDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PacketMirroring: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting PacketMirroring %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "PacketMirroring")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting PacketMirroring", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting PacketMirroring %q: %#v", d.Id(), res)
	return nil
}

func resourceComputePacketMirroringImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/packetMirrorings/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputePacketMirroringName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputePacketMirroringDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputePacketMirroringRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenComputePacketMirroringNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
func flattenComputePacketMirroringNetworkUrl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputePacketMirroringPriority(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputePacketMirroringCollectorIlb(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
func flattenComputePacketMirroringCollectorIlbUrl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputePacketMirroringFilter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
	transformed["direction"] =
		flattenComputePacketMirroringFilterDirection(original["direction"], d, config)
	return []interface{}{transformed}
}
func flattenComputePacketMirroringFilterIpProtocols(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputePacketMirroringFilterCidrRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputePacketMirroringFilterDirection(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputePacketMirroringMirroredResources(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
func flattenComputePacketMirroringMirroredResourcesSubnetworks(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
func flattenComputePacketMirroringMirroredResourcesSubnetworksUrl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputePacketMirroringMirroredResourcesInstances(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
func flattenComputePacketMirroringMirroredResourcesInstancesUrl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputePacketMirroringMirroredResourcesTags(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandComputePacketMirroringName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePacketMirroringDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePacketMirroringRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePacketMirroringNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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
	} else if val := reflect.ValueOf(transformedUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["url"] = transformedUrl
	}

	return transformed, nil
}

func expandComputePacketMirroringNetworkUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for url: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputePacketMirroringPriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePacketMirroringCollectorIlb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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
	} else if val := reflect.ValueOf(transformedUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["url"] = transformedUrl
	}

	return transformed, nil
}

func expandComputePacketMirroringCollectorIlbUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseRegionalFieldValue("forwardingRules", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for url: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputePacketMirroringFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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
	} else if val := reflect.ValueOf(transformedIpProtocols); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["IPProtocols"] = transformedIpProtocols
	}

	transformedCidrRanges, err := expandComputePacketMirroringFilterCidrRanges(original["cidr_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCidrRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cidrRanges"] = transformedCidrRanges
	}

	transformedDirection, err := expandComputePacketMirroringFilterDirection(original["direction"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDirection); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["direction"] = transformedDirection
	}

	return transformed, nil
}

func expandComputePacketMirroringFilterIpProtocols(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePacketMirroringFilterCidrRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePacketMirroringFilterDirection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePacketMirroringMirroredResources(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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
	} else if val := reflect.ValueOf(transformedSubnetworks); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subnetworks"] = transformedSubnetworks
	}

	transformedInstances, err := expandComputePacketMirroringMirroredResourcesInstances(original["instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstances); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["instances"] = transformedInstances
	}

	transformedTags, err := expandComputePacketMirroringMirroredResourcesTags(original["tags"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTags); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tags"] = transformedTags
	}

	return transformed, nil
}

func expandComputePacketMirroringMirroredResourcesSubnetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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
		} else if val := reflect.ValueOf(transformedUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["url"] = transformedUrl
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputePacketMirroringMirroredResourcesSubnetworksUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseRegionalFieldValue("subnetworks", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for url: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputePacketMirroringMirroredResourcesInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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
		} else if val := reflect.ValueOf(transformedUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["url"] = transformedUrl
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputePacketMirroringMirroredResourcesInstancesUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseZonalFieldValue("instances", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for url: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputePacketMirroringMirroredResourcesTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
