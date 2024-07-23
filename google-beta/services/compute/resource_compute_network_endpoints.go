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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

type NetworkEndpointsNetworkEndpoint struct {
	IPAddress string
	Port      int
	Instance  string
}

func NetworkEndpointsNetworkEndpointConvertToStruct(endpoint interface{}) NetworkEndpointsNetworkEndpoint {
	e := endpoint.(map[string]interface{})
	ipAddress := e["ip_address"].(string)
	port := e["port"].(int)
	instance, _ := e["instance"].(string)
	return NetworkEndpointsNetworkEndpoint{
		IPAddress: ipAddress,
		Port:      port,
		Instance:  instance,
	}
}

func NetworkEndpointsNetworkEndpointConvertToAny(endpoint NetworkEndpointsNetworkEndpoint) interface{} {
	m := make(map[string]interface{})
	m["ip_address"] = endpoint.IPAddress
	m["port"] = endpoint.Port
	m["instance"] = endpoint.Instance
	return m
}

// Continues to read network endpoints as long as there are unread pages remaining
func networkEndpointsPaginatedRead(d *schema.ResourceData, config *transport_tpg.Config, userAgent, url, project, billingProject, pt string) ([]interface{}, error) {
	var allEndpoints []interface{}
	for len(pt) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   billingProject,
			RawURL:    fmt.Sprintf("%s?pageToken=%s", url, pt),
			UserAgent: userAgent,
		})
		if err != nil {
			return nil, transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeNetworkEndpoints %q", d.Id()))
		}
		resEndpoints := res["items"].([]interface{})
		allEndpoints = append(allEndpoints, resEndpoints...)
		pt, _ = res["nextPageToken"].(string)
	}
	return allEndpoints, nil
}

// Mutates the parent NEG by attaching or detaching endpoints in chunks. `url` determines if endpoints are attached or detached.
// The last page is not processed, but instead returned for the Create/Delete functions to write.
func networkEndpointsPaginatedMutate(d *schema.ResourceData, endpoints []interface{}, config *transport_tpg.Config, userAgent, url, project, billingProject string, chunkSize int, returnLastPage bool) ([]interface{}, error) {
	// Pull out what this mutation is doing - either attachNetworkEndpoints or detachNetworkEndpoints
	verb := url[len(url)-len("attachNetworkEndpoints"):]
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{zone}}/{{network_endpoint_group}}/endpoints")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	i := 0
	for ; i < len(endpoints); i += chunkSize {
		j := i + chunkSize
		if j > len(endpoints) {
			if returnLastPage {
				break
			}
			j = len(endpoints)
		}
		timeoutType := schema.TimeoutCreate
		if verb != "attachNetworkEndpoints" {
			timeoutType = schema.TimeoutDelete
		}
		body := map[string]interface{}{"networkEndpoints": endpoints[i:j]}
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      body,
			Timeout:   d.Timeout(timeoutType),
		})
		if err != nil {
			return nil, fmt.Errorf("Error during %s: %s", verb, err)
		}

		err = ComputeOperationWaitTime(
			config, res, project, verb, userAgent,
			d.Timeout(schema.TimeoutDefault))

		if err != nil {
			// The mutation wasn't applied
			return nil, fmt.Errorf("Error in %s operation: %s", verb, err)
		}

		log.Printf("[DEBUG] Finished %s %q: %#v", verb, id, res)
	}
	if returnLastPage {
		return endpoints[i:], nil
	}
	return nil, nil
}

func ResourceComputeNetworkEndpoints() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeNetworkEndpointsCreate,
		Read:   resourceComputeNetworkEndpointsRead,
		Update: resourceComputeNetworkEndpointsUpdate,
		Delete: resourceComputeNetworkEndpointsDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeNetworkEndpointsImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
			tpgresource.DefaultProviderZone,
		),

		Schema: map[string]*schema.Schema{
			"network_endpoint_group": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tpgresource.CompareResourceNames,
				Description:      `The network endpoint group these endpoints are part of.`,
			},
			"network_endpoints": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: `The network endpoints to be added to the enclosing network endpoint group
(NEG). Each endpoint specifies an IP address and port, along with
additional information depending on the NEG type.`,
				Elem: computeNetworkEndpointsNetworkEndpointsSchema(),
				// Default schema.HashSchema is used.
			},
			"zone": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `Zone where the containing network endpoint group is located.`,
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

func computeNetworkEndpointsNetworkEndpointsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type:     schema.TypeString,
				Required: true,
				Description: `IPv4 address of network endpoint. The IP address must belong
to a VM in GCE (either the primary IP or as part of an aliased IP
range).`,
			},
			"instance": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `The name for a specific VM instance that the IP address belongs to.
This is required for network endpoints of type GCE_VM_IP_PORT.
The instance must be in the same zone as the network endpoint group.`,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `Port number of network endpoint.
**Note** 'port' is required unless the Network Endpoint Group is created
with the type of 'GCE_VM_IP'`,
			},
		},
	}
}

func resourceComputeNetworkEndpointsCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	networkEndpointsProp, err := expandComputeNetworkEndpointsNetworkEndpoints(d.Get("network_endpoints"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_endpoints"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkEndpointsProp)) && (ok || !reflect.DeepEqual(v, networkEndpointsProp)) {
		obj["networkEndpoints"] = networkEndpointsProp
	}

	obj, err = resourceComputeNetworkEndpointsEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "networkEndpoint/{{project}}/{{zone}}/{{network_endpoint_group}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/attachNetworkEndpoints")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new NetworkEndpoints: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkEndpoints: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	chunkSize := 500 // API only accepts 500 endpoints at a time
	lastPage, err := networkEndpointsPaginatedMutate(d, obj["networkEndpoints"].([]interface{}), config, userAgent, url, project, billingProject, chunkSize, true)
	if err != nil {
		// networkEndpointsPaginatedMutate already adds error description
		return err
	}
	obj["networkEndpoints"] = lastPage
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
		return fmt.Errorf("Error creating NetworkEndpoints: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{zone}}/{{network_endpoint_group}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating NetworkEndpoints", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create NetworkEndpoints: %s", err)
	}

	log.Printf("[DEBUG] Finished creating NetworkEndpoints %q: %#v", d.Id(), res)

	return resourceComputeNetworkEndpointsRead(d, meta)
}

func resourceComputeNetworkEndpointsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/listNetworkEndpoints")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkEndpoints: %s", err)
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
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeNetworkEndpoints %q", d.Id()))
	}

	res, err = resourceComputeNetworkEndpointsDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing ComputeNetworkEndpoints because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading NetworkEndpoints: %s", err)
	}

	zone, err := tpgresource.GetZone(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("zone", zone); err != nil {
		return fmt.Errorf("Error reading NetworkEndpoints: %s", err)
	}

	if err := d.Set("network_endpoints", flattenComputeNetworkEndpointsNetworkEndpoints(res["networkEndpoints"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpoints: %s", err)
	}

	return nil
}

func resourceComputeNetworkEndpointsUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkEndpoints: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	networkEndpointsProp, err := expandComputeNetworkEndpointsNetworkEndpoints(d.Get("network_endpoints"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_endpoints"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, networkEndpointsProp)) {
		obj["networkEndpoints"] = networkEndpointsProp
	}

	obj, err = resourceComputeNetworkEndpointsEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "networkEndpoint/{{project}}/{{zone}}/{{network_endpoint_group}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/attachNetworkEndpoints")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating NetworkEndpoints %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	detachUrl, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/detachNetworkEndpoints")
	o, n := d.GetChange("network_endpoints")

	oldEndpoints := make(map[NetworkEndpointsNetworkEndpoint]struct{})
	newEndpoints := make(map[NetworkEndpointsNetworkEndpoint]struct{})

	for _, e := range o.(*schema.Set).List() {
		oldEndpoints[NetworkEndpointsNetworkEndpointConvertToStruct(e)] = struct{}{}
	}

	for _, e := range n.(*schema.Set).List() {
		newEndpoints[NetworkEndpointsNetworkEndpointConvertToStruct(e)] = struct{}{}
	}

	// We want to ignore any endpoints that are shared between the two.
	endpointsToKeep := make(map[NetworkEndpointsNetworkEndpoint]struct{})
	for e := range oldEndpoints {
		if _, ok := newEndpoints[e]; ok {
			endpointsToKeep[e] = struct{}{}
		}
	}
	log.Printf("number of old endpoints: %v\n", len(oldEndpoints))
	log.Printf("number of new endpoints: %v\n", len(newEndpoints))
	log.Printf("number of shared endpoints: %v\n", len(endpointsToKeep))

	for e := range endpointsToKeep {
		// Removing all shared endpoints from the old endpoints yields the list of endpoints to detach.
		delete(oldEndpoints, e)
		// Removing all shared endpoints from the new endpoints yields the list of endpoints to attch.
		delete(newEndpoints, e)
	}

	var endpointsToDetach []interface{}
	for e := range oldEndpoints {
		endpointsToDetach = append(endpointsToDetach, NetworkEndpointsNetworkEndpointConvertToAny(e))
	}
	var endpointsToAttach []interface{}
	for e := range newEndpoints {
		endpointsToAttach = append(endpointsToAttach, NetworkEndpointsNetworkEndpointConvertToAny(e))
	}

	log.Printf("number of endpoints to detach: %v\n", len(endpointsToDetach))
	log.Printf("number of endpoints to attach: %v\n", len(endpointsToAttach))

	chunkSize := 500 // API only accepts 500 endpoints at a time

	_, err = networkEndpointsPaginatedMutate(d, endpointsToDetach, config, userAgent, detachUrl, project, billingProject, chunkSize, false)
	if err != nil {
		// networkEndpointsPaginatedMutate already adds error description
		return err
	}

	lastPage, err := networkEndpointsPaginatedMutate(d, endpointsToAttach, config, userAgent, url, project, billingProject, chunkSize, true)
	if err != nil {
		// networkEndpointsPaginatedMutate already adds error description
		return err
	}

	obj = map[string]interface{}{
		"networkEndpoints": lastPage,
	}

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
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return fmt.Errorf("Error updating NetworkEndpoints %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating NetworkEndpoints %q: %#v", d.Id(), res)
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Updating NetworkEndpoints", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputeNetworkEndpointsRead(d, meta)
}

func resourceComputeNetworkEndpointsDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkEndpoints: %s", err)
	}
	billingProject = project

	lockName, err := tpgresource.ReplaceVars(d, config, "networkEndpoint/{{project}}/{{zone}}/{{network_endpoint_group}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/detachNetworkEndpoints")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	var endpointsToDelete []interface{}

	endpoints := d.Get("network_endpoints").(*schema.Set).List()

	for _, e := range endpoints {
		endpoint := e.(map[string]interface{})
		toDelete := make(map[string]interface{})
		instanceProp, err := expandNestedComputeNetworkEndpointInstance(endpoint["instance"], d, config)
		if err != nil {
			return err
		}
		if instanceProp != "" {
			toDelete["instance"] = instanceProp
		}

		portProp, err := expandNestedComputeNetworkEndpointPort(endpoint["port"], d, config)
		if err != nil {
			return err
		}
		if portProp != 0 {
			toDelete["port"] = portProp
		}

		ipAddressProp, err := expandNestedComputeNetworkEndpointIpAddress(endpoint["ip_address"], d, config)
		if err != nil {
			return err
		}
		toDelete["ipAddress"] = ipAddressProp
		endpointsToDelete = append(endpointsToDelete, toDelete)
	}

	chunkSize := 500 // API only accepts 500 endpoints at a time
	lastPage, err := networkEndpointsPaginatedMutate(d, endpointsToDelete, config, userAgent, url, project, billingProject, chunkSize, true)
	if err != nil {
		// networkEndpointsPaginatedMutate already adds error description
		return err
	}

	obj = map[string]interface{}{
		"networkEndpoints": lastPage,
	}

	log.Printf("[DEBUG] Deleting NetworkEndpoints %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "NetworkEndpoints")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting NetworkEndpoints", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting NetworkEndpoints %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeNetworkEndpointsImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/networkEndpointGroups/(?P<network_endpoint_group>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<network_endpoint_group>[^/]+)$",
		"^(?P<zone>[^/]+)/(?P<network_endpoint_group>[^/]+)$",
		"^(?P<network_endpoint_group>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{zone}}/{{network_endpoint_group}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeNetworkEndpointsNetworkEndpoints(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(computeNetworkEndpointsNetworkEndpointsSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"instance":   flattenComputeNetworkEndpointsNetworkEndpointsInstance(original["instance"], d, config),
			"port":       flattenComputeNetworkEndpointsNetworkEndpointsPort(original["port"], d, config),
			"ip_address": flattenComputeNetworkEndpointsNetworkEndpointsIpAddress(original["ipAddress"], d, config),
		})
	}
	return transformed
}
func flattenComputeNetworkEndpointsNetworkEndpointsInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeNetworkEndpointsNetworkEndpointsPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles int given in float64 format
	if floatVal, ok := v.(float64); ok {
		return int(floatVal)
	}
	return v
}

func flattenComputeNetworkEndpointsNetworkEndpointsIpAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandComputeNetworkEndpointsNetworkEndpoints(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedInstance, err := expandComputeNetworkEndpointsNetworkEndpointsInstance(original["instance"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["instance"] = transformedInstance
		}

		transformedPort, err := expandComputeNetworkEndpointsNetworkEndpointsPort(original["port"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["port"] = transformedPort
		}

		transformedIpAddress, err := expandComputeNetworkEndpointsNetworkEndpointsIpAddress(original["ip_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipAddress"] = transformedIpAddress
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeNetworkEndpointsNetworkEndpointsInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.GetResourceNameFromSelfLink(v.(string)), nil
}

func expandComputeNetworkEndpointsNetworkEndpointsPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkEndpointsNetworkEndpointsIpAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceComputeNetworkEndpointsEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Network Endpoint Group is a URL parameter only, so replace self-link/path with resource name only.
	if err := d.Set("network_endpoint_group", tpgresource.GetResourceNameFromSelfLink(d.Get("network_endpoint_group").(string))); err != nil {
		return nil, fmt.Errorf("Error setting network_endpoint_group: %s", err)
	}

	return obj, nil
}

func resourceComputeNetworkEndpointsDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return nil, err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/listNetworkEndpoints")
	if err != nil {
		return nil, err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return nil, fmt.Errorf("Error fetching project for NetworkEndpoint: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}
	// Read past the first page to get all endpoints.
	pt, _ := res["nextPageToken"].(string)
	allEndpoints, err := networkEndpointsPaginatedRead(d, config, userAgent, url, project, billingProject, pt)
	if err != nil {
		// networkEndpointsPaginatedRead already adds error description
		return nil, err
	}
	firstPage := res["items"].([]interface{})
	allEndpoints = append(firstPage, allEndpoints...)

	// listNetworkEndpoints returns data in a different structure, so we need to
	// convert to the Terraform schema.
	var transformed []interface{}
	for _, e := range allEndpoints {
		t := e.(map[string]interface{})["networkEndpoint"]
		transformed = append(transformed, t)
	}

	return map[string]interface{}{"networkEndpoints": transformed}, nil
}
