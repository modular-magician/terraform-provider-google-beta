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

package networkservices

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceNetworkServicesTlsRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkServicesTlsRouteCreate,
		Read:   resourceNetworkServicesTlsRouteRead,
		Update: resourceNetworkServicesTlsRouteUpdate,
		Delete: resourceNetworkServicesTlsRouteDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkServicesTlsRouteImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the TlsRoute resource.`,
			},
			"rules": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Rules that define how traffic is routed and handled.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `Required. A detailed rule defining how to route traffic.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destinations": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `The destination to which traffic should be forwarded.`,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"service_name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: `The URL of a BackendService to route traffic to.`,
												},
												"weight": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: `Specifies the proportion of requests forwarded to the backend referenced by the serviceName field.`,
												},
											},
										},
									},
								},
							},
						},
						"matches": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `Matches define the predicate used to match requests to a given action.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alpn": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `ALPN (Application-Layer Protocol Negotiation) to match against. Examples: "http/1.1", "h2". At least one of sniHost and alpn is required. Up to 5 alpns across all matches can be set.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"sni_host": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `SNI (server name indicator) to match against. SNI will be matched against all wildcard domains, i.e. www.example.com will be first matched against www.example.com, then *.example.com, then *.com.
Partial wildcards are not supported, and values like *w.example.com are invalid. At least one of sniHost and alpn is required. Up to 5 sni hosts across all matches can be set.`,
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
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A free-text description of the resource. Max length 1024 characters.`,
			},
			"gateways": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests served by the gateway.
Each gateway reference should match the pattern: projects/*/locations/global/gateways/<gateway_name>`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"meshes": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served by the mesh.
Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name>
The attached Mesh should be of a type SIDECAR`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the TlsRoute was created in UTC.`,
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URL of this resource.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the TlsRoute was updated in UTC.`,
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

func resourceNetworkServicesTlsRouteCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesTlsRouteDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	meshesProp, err := expandNetworkServicesTlsRouteMeshes(d.Get("meshes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("meshes"); ok || !reflect.DeepEqual(v, meshesProp) {
		obj["meshes"] = meshesProp
	}
	gatewaysProp, err := expandNetworkServicesTlsRouteGateways(d.Get("gateways"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gateways"); ok || !reflect.DeepEqual(v, gatewaysProp) {
		obj["gateways"] = gatewaysProp
	}
	rulesProp, err := expandNetworkServicesTlsRouteRules(d.Get("rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rules"); ok || !reflect.DeepEqual(v, rulesProp) {
		obj["rules"] = rulesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/tlsRoutes?tlsRouteId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TlsRoute: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TlsRoute: %s", err)
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
		return fmt.Errorf("Error creating TlsRoute: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/tlsRoutes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkServicesOperationWaitTime(
		config, res, project, "Creating TlsRoute", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create TlsRoute: %s", err)
	}

	log.Printf("[DEBUG] Finished creating TlsRoute %q: %#v", d.Id(), res)

	return resourceNetworkServicesTlsRouteRead(d, meta)
}

func resourceNetworkServicesTlsRouteRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/tlsRoutes/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TlsRoute: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkServicesTlsRoute %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading TlsRoute: %s", err)
	}

	if err := d.Set("self_link", flattenNetworkServicesTlsRouteSelfLink(res["selfLink"], d, config)); err != nil {
		return fmt.Errorf("Error reading TlsRoute: %s", err)
	}
	if err := d.Set("create_time", flattenNetworkServicesTlsRouteCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading TlsRoute: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkServicesTlsRouteUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading TlsRoute: %s", err)
	}
	if err := d.Set("description", flattenNetworkServicesTlsRouteDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading TlsRoute: %s", err)
	}
	if err := d.Set("meshes", flattenNetworkServicesTlsRouteMeshes(res["meshes"], d, config)); err != nil {
		return fmt.Errorf("Error reading TlsRoute: %s", err)
	}
	if err := d.Set("gateways", flattenNetworkServicesTlsRouteGateways(res["gateways"], d, config)); err != nil {
		return fmt.Errorf("Error reading TlsRoute: %s", err)
	}
	if err := d.Set("rules", flattenNetworkServicesTlsRouteRules(res["rules"], d, config)); err != nil {
		return fmt.Errorf("Error reading TlsRoute: %s", err)
	}

	return nil
}

func resourceNetworkServicesTlsRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TlsRoute: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesTlsRouteDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	meshesProp, err := expandNetworkServicesTlsRouteMeshes(d.Get("meshes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("meshes"); ok || !reflect.DeepEqual(v, meshesProp) {
		obj["meshes"] = meshesProp
	}
	gatewaysProp, err := expandNetworkServicesTlsRouteGateways(d.Get("gateways"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gateways"); ok || !reflect.DeepEqual(v, gatewaysProp) {
		obj["gateways"] = gatewaysProp
	}
	rulesProp, err := expandNetworkServicesTlsRouteRules(d.Get("rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rules"); ok || !reflect.DeepEqual(v, rulesProp) {
		obj["rules"] = rulesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/tlsRoutes/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating TlsRoute %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("meshes") {
		updateMask = append(updateMask, "meshes")
	}

	if d.HasChange("gateways") {
		updateMask = append(updateMask, "gateways")
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

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating TlsRoute %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating TlsRoute %q: %#v", d.Id(), res)
		}

		err = NetworkServicesOperationWaitTime(
			config, res, project, "Updating TlsRoute", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkServicesTlsRouteRead(d, meta)
}

func resourceNetworkServicesTlsRouteDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TlsRoute: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/tlsRoutes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting TlsRoute %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "TlsRoute")
	}

	err = NetworkServicesOperationWaitTime(
		config, res, project, "Deleting TlsRoute", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting TlsRoute %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkServicesTlsRouteImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/global/tlsRoutes/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/tlsRoutes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkServicesTlsRouteSelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTlsRouteCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTlsRouteUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTlsRouteDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTlsRouteMeshes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTlsRouteGateways(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTlsRouteRules(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"matches": flattenNetworkServicesTlsRouteRulesMatches(original["matches"], d, config),
			"action":  flattenNetworkServicesTlsRouteRulesAction(original["action"], d, config),
		})
	}
	return transformed
}
func flattenNetworkServicesTlsRouteRulesMatches(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"sni_host": flattenNetworkServicesTlsRouteRulesMatchesSniHost(original["sniHost"], d, config),
			"alpn":     flattenNetworkServicesTlsRouteRulesMatchesAlpn(original["alpn"], d, config),
		})
	}
	return transformed
}
func flattenNetworkServicesTlsRouteRulesMatchesSniHost(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTlsRouteRulesMatchesAlpn(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTlsRouteRulesAction(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["destinations"] =
		flattenNetworkServicesTlsRouteRulesActionDestinations(original["destinations"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkServicesTlsRouteRulesActionDestinations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"service_name": flattenNetworkServicesTlsRouteRulesActionDestinationsServiceName(original["serviceName"], d, config),
			"weight":       flattenNetworkServicesTlsRouteRulesActionDestinationsWeight(original["weight"], d, config),
		})
	}
	return transformed
}
func flattenNetworkServicesTlsRouteRulesActionDestinationsServiceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTlsRouteRulesActionDestinationsWeight(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func expandNetworkServicesTlsRouteDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTlsRouteMeshes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTlsRouteGateways(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTlsRouteRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMatches, err := expandNetworkServicesTlsRouteRulesMatches(original["matches"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMatches); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["matches"] = transformedMatches
		}

		transformedAction, err := expandNetworkServicesTlsRouteRulesAction(original["action"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["action"] = transformedAction
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesTlsRouteRulesMatches(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedSniHost, err := expandNetworkServicesTlsRouteRulesMatchesSniHost(original["sni_host"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["sniHost"] = transformedSniHost
		}

		transformedAlpn, err := expandNetworkServicesTlsRouteRulesMatchesAlpn(original["alpn"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["alpn"] = transformedAlpn
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesTlsRouteRulesMatchesSniHost(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTlsRouteRulesMatchesAlpn(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTlsRouteRulesAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDestinations, err := expandNetworkServicesTlsRouteRulesActionDestinations(original["destinations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestinations); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destinations"] = transformedDestinations
	}

	return transformed, nil
}

func expandNetworkServicesTlsRouteRulesActionDestinations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedServiceName, err := expandNetworkServicesTlsRouteRulesActionDestinationsServiceName(original["service_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedServiceName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["serviceName"] = transformedServiceName
		}

		transformedWeight, err := expandNetworkServicesTlsRouteRulesActionDestinationsWeight(original["weight"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedWeight); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["weight"] = transformedWeight
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesTlsRouteRulesActionDestinationsServiceName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTlsRouteRulesActionDestinationsWeight(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
