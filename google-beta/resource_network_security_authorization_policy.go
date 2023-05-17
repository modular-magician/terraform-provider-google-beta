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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceNetworkSecurityAuthorizationPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecurityAuthorizationPolicyCreate,
		Read:   resourceNetworkSecurityAuthorizationPolicyRead,
		Update: resourceNetworkSecurityAuthorizationPolicyUpdate,
		Delete: resourceNetworkSecurityAuthorizationPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkSecurityAuthorizationPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"action": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ALLOW", "DENY"}),
				Description:  `The action to take when a rule match is found. Possible values are "ALLOW" or "DENY". Possible values: ["ALLOW", "DENY"]`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the AuthorizationPolicy resource.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A free-text description of the resource. Max length 1024 characters.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Set of label tags associated with the AuthorizationPolicy resource.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The location of the authorization policy.
The default value is 'global'.`,
				Default: "global",
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `List of rules to match. Note that at least one of the rules must match in order for the action specified in the 'action' field to be taken.
A rule is a match if there is a matching source and destination. If left blank, the action specified in the action field will be applied on every request.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destinations": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `List of attributes for the traffic destination. All of the destinations must match. A destination is a match if a request matches all the specified hosts, ports, methods and headers.
If not set, the action specified in the 'action' field will be applied without any rule checks for the destination.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hosts": {
										Type:        schema.TypeList,
										Required:    true,
										Description: `List of host names to match. Matched against the ":authority" header in http requests. At least one host should match. Each host can be an exact match, or a prefix match (example "mydomain.*") or a suffix match (example "*.myorg.com") or a presence (any) match "*".`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"methods": {
										Type:        schema.TypeList,
										Required:    true,
										Description: `A list of HTTP methods to match. At least one method should match. Should not be set for gRPC services.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ports": {
										Type:        schema.TypeList,
										Required:    true,
										Description: `List of destination ports to match. At least one port should match.`,
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"http_header_match": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `Match against key:value pair in http header. Provides a flexible match based on HTTP headers, for potentially advanced use cases. At least one header should match.
Avoid using header matches to make authorization decisions unless there is a strong guarantee that requests arrive through a trusted client or proxy.`,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"header_name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: `The name of the HTTP header to match. For matching against the HTTP request's authority, use a headerMatch with the header name ":authority". For matching a request's method, use the headerName ":method".`,
												},
												"regex_match": {
													Type:        schema.TypeString,
													Required:    true,
													Description: `The value of the header must match the regular expression specified in regexMatch. For regular expression grammar, please see: en.cppreference.com/w/cpp/regex/ecmascript For matching against a port specified in the HTTP request, use a headerMatch with headerName set to Host and a regular expression that satisfies the RFC2616 Host header's port specifier.`,
												},
											},
										},
									},
								},
							},
						},
						"sources": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `List of attributes for the traffic source. All of the sources must match. A source is a match if both principals and ipBlocks match.
If not set, the action specified in the 'action' field will be applied without any rule checks for the source.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_blocks": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `List of CIDR ranges to match based on source IP address. At least one IP block should match. Single IP (e.g., "1.2.3.4") and CIDR (e.g., "1.2.3.0/24") are supported. Authorization based on source IP alone should be avoided.
The IP addresses of any load balancers or proxies should be considered untrusted.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"principals": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `List of peer identities to match for authorization. At least one principal should match. Each peer can be an exact match, or a prefix match (example, "namespace/*") or a suffix match (example, "*/service-account") or a presence match "*".
Authorization based on the principal name without certificate validation (configured by ServerTlsPolicy resource) is considered insecure.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the AuthorizationPolicy was created in UTC.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the AuthorizationPolicy was updated in UTC.`,
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

func resourceNetworkSecurityAuthorizationPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandNetworkSecurityAuthorizationPolicyLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	descriptionProp, err := expandNetworkSecurityAuthorizationPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	actionProp, err := expandNetworkSecurityAuthorizationPolicyAction(d.Get("action"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("action"); !tpgresource.IsEmptyValue(reflect.ValueOf(actionProp)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	rulesProp, err := expandNetworkSecurityAuthorizationPolicyRules(d.Get("rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rules"); !tpgresource.IsEmptyValue(reflect.ValueOf(rulesProp)) && (ok || !reflect.DeepEqual(v, rulesProp)) {
		obj["rules"] = rulesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/authorizationPolicies?authorizationPolicyId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AuthorizationPolicy: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AuthorizationPolicy: %s", err)
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
		return fmt.Errorf("Error creating AuthorizationPolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Creating AuthorizationPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create AuthorizationPolicy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating AuthorizationPolicy %q: %#v", d.Id(), res)

	return resourceNetworkSecurityAuthorizationPolicyRead(d, meta)
}

func resourceNetworkSecurityAuthorizationPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AuthorizationPolicy: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkSecurityAuthorizationPolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading AuthorizationPolicy: %s", err)
	}

	if err := d.Set("create_time", flattenNetworkSecurityAuthorizationPolicyCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading AuthorizationPolicy: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkSecurityAuthorizationPolicyUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading AuthorizationPolicy: %s", err)
	}
	if err := d.Set("labels", flattenNetworkSecurityAuthorizationPolicyLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading AuthorizationPolicy: %s", err)
	}
	if err := d.Set("description", flattenNetworkSecurityAuthorizationPolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading AuthorizationPolicy: %s", err)
	}
	if err := d.Set("action", flattenNetworkSecurityAuthorizationPolicyAction(res["action"], d, config)); err != nil {
		return fmt.Errorf("Error reading AuthorizationPolicy: %s", err)
	}
	if err := d.Set("rules", flattenNetworkSecurityAuthorizationPolicyRules(res["rules"], d, config)); err != nil {
		return fmt.Errorf("Error reading AuthorizationPolicy: %s", err)
	}

	return nil
}

func resourceNetworkSecurityAuthorizationPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AuthorizationPolicy: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandNetworkSecurityAuthorizationPolicyLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	descriptionProp, err := expandNetworkSecurityAuthorizationPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	actionProp, err := expandNetworkSecurityAuthorizationPolicyAction(d.Get("action"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("action"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	rulesProp, err := expandNetworkSecurityAuthorizationPolicyRules(d.Get("rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rules"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, rulesProp)) {
		obj["rules"] = rulesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AuthorizationPolicy %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("action") {
		updateMask = append(updateMask, "action")
	}

	if d.HasChange("rules") {
		updateMask = append(updateMask, "rules")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

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
		return fmt.Errorf("Error updating AuthorizationPolicy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating AuthorizationPolicy %q: %#v", d.Id(), res)
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Updating AuthorizationPolicy", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceNetworkSecurityAuthorizationPolicyRead(d, meta)
}

func resourceNetworkSecurityAuthorizationPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AuthorizationPolicy: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting AuthorizationPolicy %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "AuthorizationPolicy")
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Deleting AuthorizationPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting AuthorizationPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkSecurityAuthorizationPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/authorizationPolicies/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkSecurityAuthorizationPolicyCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAuthorizationPolicyUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAuthorizationPolicyLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAuthorizationPolicyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAuthorizationPolicyAction(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAuthorizationPolicyRules(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"sources":      flattenNetworkSecurityAuthorizationPolicyRulesSources(original["sources"], d, config),
			"destinations": flattenNetworkSecurityAuthorizationPolicyRulesDestinations(original["destinations"], d, config),
		})
	}
	return transformed
}
func flattenNetworkSecurityAuthorizationPolicyRulesSources(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"principals": flattenNetworkSecurityAuthorizationPolicyRulesSourcesPrincipals(original["principals"], d, config),
			"ip_blocks":  flattenNetworkSecurityAuthorizationPolicyRulesSourcesIpBlocks(original["ipBlocks"], d, config),
		})
	}
	return transformed
}
func flattenNetworkSecurityAuthorizationPolicyRulesSourcesPrincipals(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAuthorizationPolicyRulesSourcesIpBlocks(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAuthorizationPolicyRulesDestinations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"hosts":             flattenNetworkSecurityAuthorizationPolicyRulesDestinationsHosts(original["hosts"], d, config),
			"ports":             flattenNetworkSecurityAuthorizationPolicyRulesDestinationsPorts(original["ports"], d, config),
			"methods":           flattenNetworkSecurityAuthorizationPolicyRulesDestinationsMethods(original["methods"], d, config),
			"http_header_match": flattenNetworkSecurityAuthorizationPolicyRulesDestinationsHttpHeaderMatch(original["httpHeaderMatch"], d, config),
		})
	}
	return transformed
}
func flattenNetworkSecurityAuthorizationPolicyRulesDestinationsHosts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAuthorizationPolicyRulesDestinationsPorts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAuthorizationPolicyRulesDestinationsMethods(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAuthorizationPolicyRulesDestinationsHttpHeaderMatch(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["header_name"] =
		flattenNetworkSecurityAuthorizationPolicyRulesDestinationsHttpHeaderMatchHeaderName(original["headerName"], d, config)
	transformed["regex_match"] =
		flattenNetworkSecurityAuthorizationPolicyRulesDestinationsHttpHeaderMatchRegexMatch(original["regexMatch"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkSecurityAuthorizationPolicyRulesDestinationsHttpHeaderMatchHeaderName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityAuthorizationPolicyRulesDestinationsHttpHeaderMatchRegexMatch(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkSecurityAuthorizationPolicyLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandNetworkSecurityAuthorizationPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityAuthorizationPolicyAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityAuthorizationPolicyRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedSources, err := expandNetworkSecurityAuthorizationPolicyRulesSources(original["sources"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSources); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["sources"] = transformedSources
		}

		transformedDestinations, err := expandNetworkSecurityAuthorizationPolicyRulesDestinations(original["destinations"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDestinations); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["destinations"] = transformedDestinations
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkSecurityAuthorizationPolicyRulesSources(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPrincipals, err := expandNetworkSecurityAuthorizationPolicyRulesSourcesPrincipals(original["principals"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPrincipals); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["principals"] = transformedPrincipals
		}

		transformedIpBlocks, err := expandNetworkSecurityAuthorizationPolicyRulesSourcesIpBlocks(original["ip_blocks"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpBlocks); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipBlocks"] = transformedIpBlocks
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkSecurityAuthorizationPolicyRulesSourcesPrincipals(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityAuthorizationPolicyRulesSourcesIpBlocks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityAuthorizationPolicyRulesDestinations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedHosts, err := expandNetworkSecurityAuthorizationPolicyRulesDestinationsHosts(original["hosts"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedHosts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["hosts"] = transformedHosts
		}

		transformedPorts, err := expandNetworkSecurityAuthorizationPolicyRulesDestinationsPorts(original["ports"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPorts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ports"] = transformedPorts
		}

		transformedMethods, err := expandNetworkSecurityAuthorizationPolicyRulesDestinationsMethods(original["methods"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMethods); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["methods"] = transformedMethods
		}

		transformedHttpHeaderMatch, err := expandNetworkSecurityAuthorizationPolicyRulesDestinationsHttpHeaderMatch(original["http_header_match"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedHttpHeaderMatch); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["httpHeaderMatch"] = transformedHttpHeaderMatch
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkSecurityAuthorizationPolicyRulesDestinationsHosts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityAuthorizationPolicyRulesDestinationsPorts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityAuthorizationPolicyRulesDestinationsMethods(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityAuthorizationPolicyRulesDestinationsHttpHeaderMatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHeaderName, err := expandNetworkSecurityAuthorizationPolicyRulesDestinationsHttpHeaderMatchHeaderName(original["header_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHeaderName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["headerName"] = transformedHeaderName
	}

	transformedRegexMatch, err := expandNetworkSecurityAuthorizationPolicyRulesDestinationsHttpHeaderMatchRegexMatch(original["regex_match"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRegexMatch); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["regexMatch"] = transformedRegexMatch
	}

	return transformed, nil
}

func expandNetworkSecurityAuthorizationPolicyRulesDestinationsHttpHeaderMatchHeaderName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityAuthorizationPolicyRulesDestinationsHttpHeaderMatchRegexMatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
