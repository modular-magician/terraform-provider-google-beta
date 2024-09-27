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

package compute

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceComputeFirewallPolicyRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeFirewallPolicyRuleCreate,
		Read:   resourceComputeFirewallPolicyRuleRead,
		Delete: resourceComputeFirewallPolicyRuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeFirewallPolicyRuleImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"action": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny", "goto_next" and "apply_security_profile_group".`,
			},
			"direction": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"INGRESS", "EGRESS"}),
				Description:  `The direction in which this rule applies. Possible values: ["INGRESS", "EGRESS"]`,
			},
			"firewall_policy": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The firewall policy of the resource.`,
			},
			"match": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: `A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"layer4_configs": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Description: `Pairs of IP protocols and ports that the rule should match.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_protocol": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
										Description: `The IP protocol to which this rule applies. The protocol type is required when creating a firewall rule.
This value can either be one of the following well known protocol strings (tcp, udp, icmp, esp, ah, ipip, sctp), or the IP protocol number.`,
									},
									"ports": {
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Description: `An optional list of ports to which this rule applies. This field is only applicable for UDP or TCP protocol. Each entry must be either an integer or a range. If not specified, this rule applies to connections through any port.
Example inputs include: ["22"], ["80","443"], and ["12345-12349"].`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"dest_address_groups": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `Address groups which should be matched against the traffic destination. Maximum number of destination address groups is 10.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"dest_fqdns": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `Fully Qualified Domain Name (FQDN) which should be matched against traffic destination. Maximum number of destination fqdn allowed is 100.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"dest_ip_ranges": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 5000.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"dest_region_codes": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `Region codes whose IP addresses will be used to match for destination of traffic. Should be specified as 2 letter country code defined as per ISO 3166 alpha-2 country codes. ex."US" Maximum number of dest region codes allowed is 5000.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"dest_threat_intelligences": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `Names of Network Threat Intelligence lists. The IPs in these lists will be matched against traffic destination.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"src_address_groups": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `Address groups which should be matched against the traffic source. Maximum number of source address groups is 10.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"src_fqdns": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `Fully Qualified Domain Name (FQDN) which should be matched against traffic source. Maximum number of source fqdn allowed is 100.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"src_ip_ranges": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 5000.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"src_region_codes": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `Region codes whose IP addresses will be used to match for source of traffic. Should be specified as 2 letter country code defined as per ISO 3166 alpha-2 country codes. ex."US" Maximum number of source region codes allowed is 5000.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"src_threat_intelligences": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `Names of Network Threat Intelligence lists. The IPs in these lists will be matched against traffic source.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
				Description: `An integer indicating the priority of a rule in the list.
The priority must be a positive value between 0 and 2147483647.
Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description for this resource.`,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `Denotes whether the firewall policy rule is disabled.
When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist.
If this is unspecified, the firewall policy rule will be enabled.`,
			},
			"enable_logging": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `Denotes whether to enable logging for a particular rule.
If logging is enabled, logs will be exported to the configured export destination in Stackdriver.
Logs may be exported to BigQuery or Pub/Sub.
Note: you cannot enable logging on "goto_next" rules.`,
			},
			"security_profile_group": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `A fully-qualified URL of a SecurityProfile resource instance.
Example: https://networksecurity.googleapis.com/v1/projects/{project}/locations/{location}/securityProfileGroups/my-security-profile-group
Must be specified if action = 'apply_security_profile_group' and cannot be specified for other actions.`,
			},
			"target_resources": {
				Type:             schema.TypeList,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `A list of network resource URLs to which this rule applies.
This field allows you to control which network's VMs get this rule.
If this field is left blank, all VMs within the organization will receive the rule.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_service_accounts": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `A list of service accounts indicating the sets of instances that are applied with this rule.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tls_inspect": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `Boolean flag indicating if the traffic should be TLS decrypted.
Can be set only if action = 'apply_security_profile_group' and cannot be set for other actions.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"kind": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Type of the resource. Always 'compute#firewallPolicyRule' for firewall policy rules`,
			},
			"rule_tuple_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Calculation of the complexity of a single firewall policy rule.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeFirewallPolicyRuleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeFirewallPolicyRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	priorityProp, err := expandComputeFirewallPolicyRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	matchProp, err := expandComputeFirewallPolicyRuleMatch(d.Get("match"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("match"); !tpgresource.IsEmptyValue(reflect.ValueOf(matchProp)) && (ok || !reflect.DeepEqual(v, matchProp)) {
		obj["match"] = matchProp
	}
	actionProp, err := expandComputeFirewallPolicyRuleAction(d.Get("action"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("action"); !tpgresource.IsEmptyValue(reflect.ValueOf(actionProp)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	securityProfileGroupProp, err := expandComputeFirewallPolicyRuleSecurityProfileGroup(d.Get("security_profile_group"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("security_profile_group"); !tpgresource.IsEmptyValue(reflect.ValueOf(securityProfileGroupProp)) && (ok || !reflect.DeepEqual(v, securityProfileGroupProp)) {
		obj["securityProfileGroup"] = securityProfileGroupProp
	}
	tlsInspectProp, err := expandComputeFirewallPolicyRuleTlsInspect(d.Get("tls_inspect"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tls_inspect"); !tpgresource.IsEmptyValue(reflect.ValueOf(tlsInspectProp)) && (ok || !reflect.DeepEqual(v, tlsInspectProp)) {
		obj["tlsInspect"] = tlsInspectProp
	}
	directionProp, err := expandComputeFirewallPolicyRuleDirection(d.Get("direction"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("direction"); !tpgresource.IsEmptyValue(reflect.ValueOf(directionProp)) && (ok || !reflect.DeepEqual(v, directionProp)) {
		obj["direction"] = directionProp
	}
	targetResourcesProp, err := expandComputeFirewallPolicyRuleTargetResources(d.Get("target_resources"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_resources"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetResourcesProp)) && (ok || !reflect.DeepEqual(v, targetResourcesProp)) {
		obj["targetResources"] = targetResourcesProp
	}
	enableLoggingProp, err := expandComputeFirewallPolicyRuleEnableLogging(d.Get("enable_logging"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_logging"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableLoggingProp)) && (ok || !reflect.DeepEqual(v, enableLoggingProp)) {
		obj["enableLogging"] = enableLoggingProp
	}
	targetServiceAccountsProp, err := expandComputeFirewallPolicyRuleTargetServiceAccounts(d.Get("target_service_accounts"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_service_accounts"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetServiceAccountsProp)) && (ok || !reflect.DeepEqual(v, targetServiceAccountsProp)) {
		obj["targetServiceAccounts"] = targetServiceAccountsProp
	}
	disabledProp, err := expandComputeFirewallPolicyRuleDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}{{firewall_policy}}/addRule")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FirewallPolicyRule: %#v", obj)
	billingProject := ""

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
		return fmt.Errorf("Error creating FirewallPolicyRule: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{firewall_policy}}/rules/{{priority}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	parent, err := tpgresource.ReplaceVars(d, config, "{{firewall_policy}}")
	if err != nil {
		return fmt.Errorf("Error constructing parent: %s", err)
	}

	var opRes map[string]interface{}
	err = ComputeOrgOperationWaitTimeWithResponse(
		config, res, &opRes, parent, "FirewallPolicyRule operation", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting for FirewallPolicyRule operation: %s", err)
	}

	log.Printf("[DEBUG] Finished creating FirewallPolicyRule %q: %#v", d.Id(), res)

	return resourceComputeFirewallPolicyRuleRead(d, meta)
}

func resourceComputeFirewallPolicyRuleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}{{firewall_policy}}/getRule?priority={{priority}}")
	if err != nil {
		return err
	}

	billingProject := ""

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeFirewallPolicyRule %q", d.Id()))
	}

	if err := d.Set("creation_timestamp", flattenComputeFirewallPolicyRuleCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyRule: %s", err)
	}
	if err := d.Set("kind", flattenComputeFirewallPolicyRuleKind(res["kind"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyRule: %s", err)
	}
	if err := d.Set("description", flattenComputeFirewallPolicyRuleDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyRule: %s", err)
	}
	if err := d.Set("priority", flattenComputeFirewallPolicyRulePriority(res["priority"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyRule: %s", err)
	}
	if err := d.Set("match", flattenComputeFirewallPolicyRuleMatch(res["match"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyRule: %s", err)
	}
	if err := d.Set("action", flattenComputeFirewallPolicyRuleAction(res["action"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyRule: %s", err)
	}
	if err := d.Set("security_profile_group", flattenComputeFirewallPolicyRuleSecurityProfileGroup(res["securityProfileGroup"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyRule: %s", err)
	}
	if err := d.Set("tls_inspect", flattenComputeFirewallPolicyRuleTlsInspect(res["tlsInspect"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyRule: %s", err)
	}
	if err := d.Set("direction", flattenComputeFirewallPolicyRuleDirection(res["direction"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyRule: %s", err)
	}
	if err := d.Set("target_resources", flattenComputeFirewallPolicyRuleTargetResources(res["targetResources"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyRule: %s", err)
	}
	if err := d.Set("enable_logging", flattenComputeFirewallPolicyRuleEnableLogging(res["enableLogging"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyRule: %s", err)
	}
	if err := d.Set("rule_tuple_count", flattenComputeFirewallPolicyRuleRuleTupleCount(res["ruleTupleCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyRule: %s", err)
	}
	if err := d.Set("target_service_accounts", flattenComputeFirewallPolicyRuleTargetServiceAccounts(res["targetServiceAccounts"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyRule: %s", err)
	}
	if err := d.Set("disabled", flattenComputeFirewallPolicyRuleDisabled(res["disabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyRule: %s", err)
	}

	return nil
}

func resourceComputeFirewallPolicyRuleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}{{firewall_policy}}/removeRule?priority={{priority}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting FirewallPolicyRule %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "FirewallPolicyRule")
	}

	parent, err := tpgresource.ReplaceVars(d, config, "{{firewall_policy}}")
	if err != nil {
		return fmt.Errorf("Error constructing parent: %s", err)
	}

	var opRes map[string]interface{}
	err = ComputeOrgOperationWaitTimeWithResponse(
		config, res, &opRes, parent, "FirewallPolicyRule operation", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting for FirewallPolicyRule operation: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting FirewallPolicyRule %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeFirewallPolicyRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^locations/global/firewallPolicies/(?P<firewall_policy>[^/]+)/rules/(?P<priority>[^/]+)$",
		"^(?P<firewall_policy>[^/]+)/(?P<priority>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "locations/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Current import uses only the Firewall Policy Id
	// Replace the field to format "locations/global/firewallPolicies/{{firewallPolicyId}}"
	stringParts := strings.Split(d.Get("firewall_policy").(string), "/")
	if len(stringParts) == 1 {
		if err := d.Set("firewall_policy", fmt.Sprintf("locations/global/firewallPolicies/%s", d.Get("firewall_policy").(string))); err != nil {
			return nil, fmt.Errorf("Error setting firewall_policy, %s", err)
		}
	}

	return []*schema.ResourceData{d}, nil
}

func flattenComputeFirewallPolicyRuleCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleKind(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRulePriority(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeFirewallPolicyRuleMatch(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["src_ip_ranges"] =
		flattenComputeFirewallPolicyRuleMatchSrcIpRanges(original["srcIpRanges"], d, config)
	transformed["dest_ip_ranges"] =
		flattenComputeFirewallPolicyRuleMatchDestIpRanges(original["destIpRanges"], d, config)
	transformed["layer4_configs"] =
		flattenComputeFirewallPolicyRuleMatchLayer4Configs(original["layer4Configs"], d, config)
	transformed["dest_address_groups"] =
		flattenComputeFirewallPolicyRuleMatchDestAddressGroups(original["destAddressGroups"], d, config)
	transformed["src_address_groups"] =
		flattenComputeFirewallPolicyRuleMatchSrcAddressGroups(original["srcAddressGroups"], d, config)
	transformed["src_fqdns"] =
		flattenComputeFirewallPolicyRuleMatchSrcFqdns(original["srcFqdns"], d, config)
	transformed["dest_fqdns"] =
		flattenComputeFirewallPolicyRuleMatchDestFqdns(original["destFqdns"], d, config)
	transformed["src_region_codes"] =
		flattenComputeFirewallPolicyRuleMatchSrcRegionCodes(original["srcRegionCodes"], d, config)
	transformed["dest_region_codes"] =
		flattenComputeFirewallPolicyRuleMatchDestRegionCodes(original["destRegionCodes"], d, config)
	transformed["dest_threat_intelligences"] =
		flattenComputeFirewallPolicyRuleMatchDestThreatIntelligences(original["destThreatIntelligences"], d, config)
	transformed["src_threat_intelligences"] =
		flattenComputeFirewallPolicyRuleMatchSrcThreatIntelligences(original["srcThreatIntelligences"], d, config)
	return []interface{}{transformed}
}
func flattenComputeFirewallPolicyRuleMatchSrcIpRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleMatchDestIpRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleMatchLayer4Configs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"ip_protocol": flattenComputeFirewallPolicyRuleMatchLayer4ConfigsIpProtocol(original["ipProtocol"], d, config),
			"ports":       flattenComputeFirewallPolicyRuleMatchLayer4ConfigsPorts(original["ports"], d, config),
		})
	}
	return transformed
}
func flattenComputeFirewallPolicyRuleMatchLayer4ConfigsIpProtocol(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleMatchLayer4ConfigsPorts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleMatchDestAddressGroups(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleMatchSrcAddressGroups(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleMatchSrcFqdns(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleMatchDestFqdns(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleMatchSrcRegionCodes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleMatchDestRegionCodes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleMatchDestThreatIntelligences(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleMatchSrcThreatIntelligences(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleAction(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleSecurityProfileGroup(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleTlsInspect(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleDirection(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleTargetResources(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleEnableLogging(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleRuleTupleCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeFirewallPolicyRuleTargetServiceAccounts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyRuleDisabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandComputeFirewallPolicyRuleDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRulePriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleMatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSrcIpRanges, err := expandComputeFirewallPolicyRuleMatchSrcIpRanges(original["src_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcIpRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcIpRanges"] = transformedSrcIpRanges
	}

	transformedDestIpRanges, err := expandComputeFirewallPolicyRuleMatchDestIpRanges(original["dest_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestIpRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destIpRanges"] = transformedDestIpRanges
	}

	transformedLayer4Configs, err := expandComputeFirewallPolicyRuleMatchLayer4Configs(original["layer4_configs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLayer4Configs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["layer4Configs"] = transformedLayer4Configs
	}

	transformedDestAddressGroups, err := expandComputeFirewallPolicyRuleMatchDestAddressGroups(original["dest_address_groups"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestAddressGroups); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destAddressGroups"] = transformedDestAddressGroups
	}

	transformedSrcAddressGroups, err := expandComputeFirewallPolicyRuleMatchSrcAddressGroups(original["src_address_groups"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcAddressGroups); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcAddressGroups"] = transformedSrcAddressGroups
	}

	transformedSrcFqdns, err := expandComputeFirewallPolicyRuleMatchSrcFqdns(original["src_fqdns"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcFqdns); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcFqdns"] = transformedSrcFqdns
	}

	transformedDestFqdns, err := expandComputeFirewallPolicyRuleMatchDestFqdns(original["dest_fqdns"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestFqdns); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destFqdns"] = transformedDestFqdns
	}

	transformedSrcRegionCodes, err := expandComputeFirewallPolicyRuleMatchSrcRegionCodes(original["src_region_codes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcRegionCodes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcRegionCodes"] = transformedSrcRegionCodes
	}

	transformedDestRegionCodes, err := expandComputeFirewallPolicyRuleMatchDestRegionCodes(original["dest_region_codes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestRegionCodes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destRegionCodes"] = transformedDestRegionCodes
	}

	transformedDestThreatIntelligences, err := expandComputeFirewallPolicyRuleMatchDestThreatIntelligences(original["dest_threat_intelligences"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestThreatIntelligences); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destThreatIntelligences"] = transformedDestThreatIntelligences
	}

	transformedSrcThreatIntelligences, err := expandComputeFirewallPolicyRuleMatchSrcThreatIntelligences(original["src_threat_intelligences"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcThreatIntelligences); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcThreatIntelligences"] = transformedSrcThreatIntelligences
	}

	return transformed, nil
}

func expandComputeFirewallPolicyRuleMatchSrcIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleMatchDestIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleMatchLayer4Configs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpProtocol, err := expandComputeFirewallPolicyRuleMatchLayer4ConfigsIpProtocol(original["ip_protocol"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpProtocol); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipProtocol"] = transformedIpProtocol
		}

		transformedPorts, err := expandComputeFirewallPolicyRuleMatchLayer4ConfigsPorts(original["ports"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPorts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ports"] = transformedPorts
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeFirewallPolicyRuleMatchLayer4ConfigsIpProtocol(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleMatchLayer4ConfigsPorts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleMatchDestAddressGroups(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleMatchSrcAddressGroups(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleMatchSrcFqdns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleMatchDestFqdns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleMatchSrcRegionCodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleMatchDestRegionCodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleMatchDestThreatIntelligences(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleMatchSrcThreatIntelligences(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleSecurityProfileGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleTlsInspect(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleDirection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleTargetResources(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleEnableLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleTargetServiceAccounts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyRuleDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
