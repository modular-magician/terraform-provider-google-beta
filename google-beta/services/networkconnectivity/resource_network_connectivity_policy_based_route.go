// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

package networkconnectivity

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceNetworkConnectivityPolicyBasedRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkConnectivityPolicyBasedRouteCreate,
		Read:   resourceNetworkConnectivityPolicyBasedRouteRead,
		Update: resourceNetworkConnectivityPolicyBasedRouteUpdate,
		Delete: resourceNetworkConnectivityPolicyBasedRouteDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkConnectivityPolicyBasedRouteImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"filter": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: `The filter to match L4 traffic.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol_version": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: verify.ValidateEnum([]string{"IPV4"}),
							Description:  `Internet protocol versions this policy-based route applies to. Possible values: ["IPV4"]`,
						},
						"dest_range": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `The destination IP range of outgoing packets that this policy-based route applies to. Default is "0.0.0.0/0" if protocol version is IPv4.`,
							Default:     "0.0.0.0/0",
						},
						"ip_protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `The IP protocol that this policy-based route applies to. Valid values are 'TCP', 'UDP', and 'ALL'. Default is 'ALL'.`,
							Default:     "ALL",
						},
						"src_range": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `The source IP range of outgoing packets that this policy-based route applies to. Default is "0.0.0.0/0" if protocol version is IPv4.`,
							Default:     "0.0.0.0/0",
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the policy based route.`,
			},
			"network": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Fully-qualified URL of the network that this route applies to, for example: projects/my-project/global/networks/my-network.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"interconnect_attachment": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `The interconnect attachments that this policy-based route applies to.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `Cloud region to install this policy-based route on for Interconnect attachments. Use 'all' to install it on all Interconnect attachments.`,
						},
					},
				},
				ConflictsWith: []string{"virtual_machine"},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `User-defined labels.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"next_hop_ilb_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  `The IP address of a global-access-enabled L4 ILB that is the next hop for matching packets.`,
				ExactlyOneOf: []string{"next_hop_ilb_ip", "next_hop_other_routes"},
			},
			"next_hop_other_routes": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"DEFAULT_ROUTING", ""}),
				Description:  `Other routes that will be referenced to determine the next hop of the packet. Possible values: ["DEFAULT_ROUTING"]`,
				ExactlyOneOf: []string{"next_hop_ilb_ip", "next_hop_other_routes"},
			},
			"priority": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: `The priority of this policy-based route. Priority is used to break ties in cases where there are more than one matching policy-based routes found. In cases where multiple policy-based routes are matched, the one with the lowest-numbered priority value wins. The default value is 1000. The priority value must be from 1 to 65535, inclusive.`,
				Default:     1000,
			},
			"virtual_machine": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `VM instances to which this policy-based route applies to.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tags": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Description: `A list of VM instance tags that this policy-based route applies to. VM instances that have ANY of tags specified here will install this PBR.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
				ConflictsWith: []string{"interconnect_attachment"},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time when the policy-based route was created.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				ForceNew:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"kind": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Type of this resource.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time when the policy-based route was created.`,
			},
			"warnings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `If potential misconfigurations are detected for this route, this field will be populated with warning messages.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `A warning code, if applicable.`,
						},
						"data": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: `Metadata about this warning in key: value format. The key should provides more detail on the warning being returned. For example, for warnings where there are no results in a list request for a particular zone, this key might be scope and the key value might be the zone name. Other examples might be a key indicating a deprecated resource and a suggested replacement.`,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"warning_message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `A human-readable description of the warning code.`,
						},
					},
				},
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

func resourceNetworkConnectivityPolicyBasedRouteCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkConnectivityPolicyBasedRouteDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networkProp, err := expandNetworkConnectivityPolicyBasedRouteNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	filterProp, err := expandNetworkConnectivityPolicyBasedRouteFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(filterProp)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}
	nextHopOtherRoutesProp, err := expandNetworkConnectivityPolicyBasedRouteNextHopOtherRoutes(d.Get("next_hop_other_routes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("next_hop_other_routes"); !tpgresource.IsEmptyValue(reflect.ValueOf(nextHopOtherRoutesProp)) && (ok || !reflect.DeepEqual(v, nextHopOtherRoutesProp)) {
		obj["nextHopOtherRoutes"] = nextHopOtherRoutesProp
	}
	nextHopIlbIpProp, err := expandNetworkConnectivityPolicyBasedRouteNextHopIlbIp(d.Get("next_hop_ilb_ip"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("next_hop_ilb_ip"); !tpgresource.IsEmptyValue(reflect.ValueOf(nextHopIlbIpProp)) && (ok || !reflect.DeepEqual(v, nextHopIlbIpProp)) {
		obj["nextHopIlbIp"] = nextHopIlbIpProp
	}
	priorityProp, err := expandNetworkConnectivityPolicyBasedRoutePriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	virtualMachineProp, err := expandNetworkConnectivityPolicyBasedRouteVirtualMachine(d.Get("virtual_machine"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("virtual_machine"); !tpgresource.IsEmptyValue(reflect.ValueOf(virtualMachineProp)) && (ok || !reflect.DeepEqual(v, virtualMachineProp)) {
		obj["virtualMachine"] = virtualMachineProp
	}
	interconnectAttachmentProp, err := expandNetworkConnectivityPolicyBasedRouteInterconnectAttachment(d.Get("interconnect_attachment"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("interconnect_attachment"); !tpgresource.IsEmptyValue(reflect.ValueOf(interconnectAttachmentProp)) && (ok || !reflect.DeepEqual(v, interconnectAttachmentProp)) {
		obj["interconnectAttachment"] = interconnectAttachmentProp
	}
	labelsProp, err := expandNetworkConnectivityPolicyBasedRouteEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/global/policyBasedRoutes?policyBasedRouteId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new PolicyBasedRoute: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PolicyBasedRoute: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating PolicyBasedRoute: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/policyBasedRoutes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkConnectivityOperationWaitTime(
		config, res, project, "Creating PolicyBasedRoute", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create PolicyBasedRoute: %s", err)
	}

	log.Printf("[DEBUG] Finished creating PolicyBasedRoute %q: %#v", d.Id(), res)

	return resourceNetworkConnectivityPolicyBasedRouteRead(d, meta)
}

func resourceNetworkConnectivityPolicyBasedRouteRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/global/policyBasedRoutes/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PolicyBasedRoute: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkConnectivityPolicyBasedRoute %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading PolicyBasedRoute: %s", err)
	}

	if err := d.Set("description", flattenNetworkConnectivityPolicyBasedRouteDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyBasedRoute: %s", err)
	}
	if err := d.Set("labels", flattenNetworkConnectivityPolicyBasedRouteLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyBasedRoute: %s", err)
	}
	if err := d.Set("network", flattenNetworkConnectivityPolicyBasedRouteNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyBasedRoute: %s", err)
	}
	if err := d.Set("filter", flattenNetworkConnectivityPolicyBasedRouteFilter(res["filter"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyBasedRoute: %s", err)
	}
	if err := d.Set("next_hop_other_routes", flattenNetworkConnectivityPolicyBasedRouteNextHopOtherRoutes(res["nextHopOtherRoutes"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyBasedRoute: %s", err)
	}
	if err := d.Set("next_hop_ilb_ip", flattenNetworkConnectivityPolicyBasedRouteNextHopIlbIp(res["nextHopIlbIp"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyBasedRoute: %s", err)
	}
	if err := d.Set("priority", flattenNetworkConnectivityPolicyBasedRoutePriority(res["priority"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyBasedRoute: %s", err)
	}
	if err := d.Set("virtual_machine", flattenNetworkConnectivityPolicyBasedRouteVirtualMachine(res["virtualMachine"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyBasedRoute: %s", err)
	}
	if err := d.Set("interconnect_attachment", flattenNetworkConnectivityPolicyBasedRouteInterconnectAttachment(res["interconnectAttachment"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyBasedRoute: %s", err)
	}
	if err := d.Set("create_time", flattenNetworkConnectivityPolicyBasedRouteCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyBasedRoute: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkConnectivityPolicyBasedRouteUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyBasedRoute: %s", err)
	}
	if err := d.Set("kind", flattenNetworkConnectivityPolicyBasedRouteKind(res["kind"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyBasedRoute: %s", err)
	}
	if err := d.Set("warnings", flattenNetworkConnectivityPolicyBasedRouteWarnings(res["warnings"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyBasedRoute: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkConnectivityPolicyBasedRouteTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyBasedRoute: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkConnectivityPolicyBasedRouteEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading PolicyBasedRoute: %s", err)
	}

	return nil
}

func resourceNetworkConnectivityPolicyBasedRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	// Only the root field "labels" and "terraform_labels" are mutable
	return resourceNetworkConnectivityPolicyBasedRouteRead(d, meta)
}

func resourceNetworkConnectivityPolicyBasedRouteDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PolicyBasedRoute: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/global/policyBasedRoutes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting PolicyBasedRoute %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "PolicyBasedRoute")
	}

	err = NetworkConnectivityOperationWaitTime(
		config, res, project, "Deleting PolicyBasedRoute", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting PolicyBasedRoute %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkConnectivityPolicyBasedRouteImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/global/policyBasedRoutes/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/policyBasedRoutes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkConnectivityPolicyBasedRouteDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityPolicyBasedRouteLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenNetworkConnectivityPolicyBasedRouteNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityPolicyBasedRouteFilter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["protocol_version"] =
		flattenNetworkConnectivityPolicyBasedRouteFilterProtocolVersion(original["protocolVersion"], d, config)
	transformed["ip_protocol"] =
		flattenNetworkConnectivityPolicyBasedRouteFilterIpProtocol(original["ipProtocol"], d, config)
	transformed["src_range"] =
		flattenNetworkConnectivityPolicyBasedRouteFilterSrcRange(original["srcRange"], d, config)
	transformed["dest_range"] =
		flattenNetworkConnectivityPolicyBasedRouteFilterDestRange(original["destRange"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkConnectivityPolicyBasedRouteFilterProtocolVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityPolicyBasedRouteFilterIpProtocol(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityPolicyBasedRouteFilterSrcRange(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityPolicyBasedRouteFilterDestRange(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityPolicyBasedRouteNextHopOtherRoutes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityPolicyBasedRouteNextHopIlbIp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityPolicyBasedRoutePriority(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
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

func flattenNetworkConnectivityPolicyBasedRouteVirtualMachine(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["tags"] =
		flattenNetworkConnectivityPolicyBasedRouteVirtualMachineTags(original["tags"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkConnectivityPolicyBasedRouteVirtualMachineTags(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityPolicyBasedRouteInterconnectAttachment(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["region"] =
		flattenNetworkConnectivityPolicyBasedRouteInterconnectAttachmentRegion(original["region"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkConnectivityPolicyBasedRouteInterconnectAttachmentRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityPolicyBasedRouteCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityPolicyBasedRouteUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityPolicyBasedRouteKind(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityPolicyBasedRouteWarnings(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"code":            flattenNetworkConnectivityPolicyBasedRouteWarningsCode(original["code"], d, config),
			"data":            flattenNetworkConnectivityPolicyBasedRouteWarningsData(original["data"], d, config),
			"warning_message": flattenNetworkConnectivityPolicyBasedRouteWarningsWarningMessage(original["warningMessage"], d, config),
		})
	}
	return transformed
}
func flattenNetworkConnectivityPolicyBasedRouteWarningsCode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityPolicyBasedRouteWarningsData(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityPolicyBasedRouteWarningsWarningMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityPolicyBasedRouteTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenNetworkConnectivityPolicyBasedRouteEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkConnectivityPolicyBasedRouteDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityPolicyBasedRouteNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityPolicyBasedRouteFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProtocolVersion, err := expandNetworkConnectivityPolicyBasedRouteFilterProtocolVersion(original["protocol_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProtocolVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["protocolVersion"] = transformedProtocolVersion
	}

	transformedIpProtocol, err := expandNetworkConnectivityPolicyBasedRouteFilterIpProtocol(original["ip_protocol"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpProtocol); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ipProtocol"] = transformedIpProtocol
	}

	transformedSrcRange, err := expandNetworkConnectivityPolicyBasedRouteFilterSrcRange(original["src_range"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcRange); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcRange"] = transformedSrcRange
	}

	transformedDestRange, err := expandNetworkConnectivityPolicyBasedRouteFilterDestRange(original["dest_range"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestRange); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destRange"] = transformedDestRange
	}

	return transformed, nil
}

func expandNetworkConnectivityPolicyBasedRouteFilterProtocolVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityPolicyBasedRouteFilterIpProtocol(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityPolicyBasedRouteFilterSrcRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityPolicyBasedRouteFilterDestRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityPolicyBasedRouteNextHopOtherRoutes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityPolicyBasedRouteNextHopIlbIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityPolicyBasedRoutePriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityPolicyBasedRouteVirtualMachine(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTags, err := expandNetworkConnectivityPolicyBasedRouteVirtualMachineTags(original["tags"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTags); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tags"] = transformedTags
	}

	return transformed, nil
}

func expandNetworkConnectivityPolicyBasedRouteVirtualMachineTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityPolicyBasedRouteInterconnectAttachment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRegion, err := expandNetworkConnectivityPolicyBasedRouteInterconnectAttachmentRegion(original["region"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRegion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["region"] = transformedRegion
	}

	return transformed, nil
}

func expandNetworkConnectivityPolicyBasedRouteInterconnectAttachmentRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityPolicyBasedRouteEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
