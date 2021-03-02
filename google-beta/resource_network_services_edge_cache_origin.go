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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetworkServicesEdgeCacheOrigin() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkServicesEdgeCacheOriginCreate,
		Read:   resourceNetworkServicesEdgeCacheOriginRead,
		Update: resourceNetworkServicesEdgeCacheOriginUpdate,
		Delete: resourceNetworkServicesEdgeCacheOriginDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkServicesEdgeCacheOriginImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource; provided by the client when the resource is created.
The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
and all following characters must be a dash, underscore, letter or digit.`,
			},
			"origin_address": {
				Type:     schema.TypeString,
				Required: true,
				Description: `A fully qualified domain name (FQDN) or IP address reachable over the public Internet, or the address of a Google Cloud Storage bucket.

This address will be used as the origin for cache requests - e.g. FQDN: media-backend.example.com IPv4:35.218.1.1 IPv6:[2607:f8b0:4012:809::200e] Cloud Storage: gs://bucketname

When providing an FQDN (hostname), it must be publicly resolvable (e.g. via Google public DNS) and IP addresses must be publicly routable.
If a Cloud Storage bucket is provided, it must be in the canonical "gs://bucketname" format. Other forms, such as "storage.googleapis.com", will be rejected.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A human-readable description of the resource.`,
			},
			"failover_origin": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The Origin resource to try when the current origin cannot be reached.
After maxAttempts is reached, the configured failoverOrigin will be used to fulfil the request.

The value of timeout.maxAttemptsTimeout dictates the timeout across all origins.
A reference to a Topic resource.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Set of label tags associated with the EdgeCache resource.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"max_attempts": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(0, 4),
				Description: `The maximum number of attempts to cache fill from this origin. Another attempt is made when a cache fill fails with one of the retryConditions.

Once maxAttempts to this origin have failed the failoverOrigin will be used, if one is specified. That failoverOrigin may specify its own maxAttempts,
retryConditions and failoverOrigin to control its own cache fill failures.

The total number of allowed attempts to cache fill across this and failover origins is limited to four.
The total time allowed for cache fill attempts across this and failover origins can be controlled with maxAttemptsTimeout.

The last valid response from an origin will be returned to the client.
If no origin returns a valid response, an HTTP 503 will be returned to the client.

Defaults to 1. Must be a value greater than 0 and less than 4.`,
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
				Description: `The port to connect to the origin on.
Defaults to port 443 for HTTP2 and HTTPS protocols, and port 80 for HTTP.`,
			},
			"protocol": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"HTTP2", "HTTPS", "HTTP", ""}, false),
				Description: `The protocol to use to connect to the configured origin. Defaults to HTTP2, and it is strongly recommended that users use HTTP2 for both security & performance.

When using HTTP2 or HTTPS as the protocol, a valid, publicly-signed, unexpired TLS (SSL) certificate must be presented by the origin server. Possible values: ["HTTP2", "HTTPS", "HTTP"]`,
			},
			"retry_conditions": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Description: `Specifies one or more retry conditions for the configured origin.

If the failure mode during a connection attempt to the origin matches the configured retryCondition(s),
the origin request will be retried up to maxAttempts times. The failoverOrigin, if configured, will then be used to satisfy the request.

The default retryCondition is "CONNECT_FAILURE".

retryConditions apply to this origin, and not subsequent failoverOrigin(s),
which may specify their own retryConditions and maxAttempts.

Valid values are:

- CONNECT_FAILURE: Retry on failures connecting to origins, for example due to connection timeouts.
- HTTP_5XX: Retry if the origin responds with any 5xx response code, or if the origin does not respond at all, example: disconnects, reset, read timeout, connection failure, and refused streams.
- GATEWAY_ERROR: Similar to 5xx, but only applies to response codes 502, 503 or 504.
- RETRIABLE_4XX: Retry for retriable 4xx response codes, which include HTTP 409 (Conflict) and HTTP 429 (Too Many Requests)
- NOT_FOUND: Retry if the origin returns a HTTP 404 (Not Found). This can be useful when generating video content, and the segment is not available yet. Possible values: ["CONNECT_FAILURE", "HTTP_5XX", "GATEWAY_ERROR", "RETRIABLE_4XX", "NOT_FOUND"]`,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"CONNECT_FAILURE", "HTTP_5XX", "GATEWAY_ERROR", "RETRIABLE_4XX", "NOT_FOUND"}, false),
				},
			},
			"timeout": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The connection and HTTP timeout configuration for this origin.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connect_timeout": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							Description: `The maximum duration to wait for the origin connection to be established, including DNS lookup, TLS handshake and TCP/QUIC connection establishment.

Defaults to 5 seconds. The timeout must be a value between 1s and 15s.`,
							AtLeastOneOf: []string{"timeout.0.connect_timeout", "timeout.0.max_attempts_timeout", "timeout.0.response_timeout"},
						},
						"max_attempts_timeout": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							Description: `The maximum time across all connection attempts to the origin, including failover origins, before returning an error to the client. A HTTP 503 will be returned if the timeout is reached before a response is returned.

Defaults to 5 seconds. The timeout must be a value between 1s and 15s.`,
							AtLeastOneOf: []string{"timeout.0.connect_timeout", "timeout.0.max_attempts_timeout", "timeout.0.response_timeout"},
						},
						"response_timeout": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							Description: `The maximum duration to wait for data to arrive when reading from the HTTP connection/stream.

Defaults to 5 seconds. The timeout must be a value between 1s and 30s.`,
							AtLeastOneOf: []string{"timeout.0.connect_timeout", "timeout.0.max_attempts_timeout", "timeout.0.response_timeout"},
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

func resourceNetworkServicesEdgeCacheOriginCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesEdgeCacheOriginDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetworkServicesEdgeCacheOriginLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	originAddressProp, err := expandNetworkServicesEdgeCacheOriginOriginAddress(d.Get("origin_address"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("origin_address"); !isEmptyValue(reflect.ValueOf(originAddressProp)) && (ok || !reflect.DeepEqual(v, originAddressProp)) {
		obj["originAddress"] = originAddressProp
	}
	protocolProp, err := expandNetworkServicesEdgeCacheOriginProtocol(d.Get("protocol"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("protocol"); !isEmptyValue(reflect.ValueOf(protocolProp)) && (ok || !reflect.DeepEqual(v, protocolProp)) {
		obj["protocol"] = protocolProp
	}
	portProp, err := expandNetworkServicesEdgeCacheOriginPort(d.Get("port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port"); !isEmptyValue(reflect.ValueOf(portProp)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}
	maxAttemptsProp, err := expandNetworkServicesEdgeCacheOriginMaxAttempts(d.Get("max_attempts"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("max_attempts"); !isEmptyValue(reflect.ValueOf(maxAttemptsProp)) && (ok || !reflect.DeepEqual(v, maxAttemptsProp)) {
		obj["maxAttempts"] = maxAttemptsProp
	}
	failoverOriginProp, err := expandNetworkServicesEdgeCacheOriginFailoverOrigin(d.Get("failover_origin"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("failover_origin"); !isEmptyValue(reflect.ValueOf(failoverOriginProp)) && (ok || !reflect.DeepEqual(v, failoverOriginProp)) {
		obj["failoverOrigin"] = failoverOriginProp
	}
	retryConditionsProp, err := expandNetworkServicesEdgeCacheOriginRetryConditions(d.Get("retry_conditions"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retry_conditions"); !isEmptyValue(reflect.ValueOf(retryConditionsProp)) && (ok || !reflect.DeepEqual(v, retryConditionsProp)) {
		obj["retryConditions"] = retryConditionsProp
	}
	timeoutProp, err := expandNetworkServicesEdgeCacheOriginTimeout(d.Get("timeout"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("timeout"); !isEmptyValue(reflect.ValueOf(timeoutProp)) && (ok || !reflect.DeepEqual(v, timeoutProp)) {
		obj["timeout"] = timeoutProp
	}

	url, err := replaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/edgeCacheOrigins?edgeCacheOriginId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new EdgeCacheOrigin: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EdgeCacheOrigin: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating EdgeCacheOrigin: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/global/edgeCacheOrigins/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = networkServicesOperationWaitTime(
		config, res, project, "Creating EdgeCacheOrigin", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create EdgeCacheOrigin: %s", err)
	}

	log.Printf("[DEBUG] Finished creating EdgeCacheOrigin %q: %#v", d.Id(), res)

	return resourceNetworkServicesEdgeCacheOriginRead(d, meta)
}

func resourceNetworkServicesEdgeCacheOriginRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/edgeCacheOrigins/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EdgeCacheOrigin: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("NetworkServicesEdgeCacheOrigin %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading EdgeCacheOrigin: %s", err)
	}

	if err := d.Set("description", flattenNetworkServicesEdgeCacheOriginDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading EdgeCacheOrigin: %s", err)
	}
	if err := d.Set("labels", flattenNetworkServicesEdgeCacheOriginLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading EdgeCacheOrigin: %s", err)
	}
	if err := d.Set("origin_address", flattenNetworkServicesEdgeCacheOriginOriginAddress(res["originAddress"], d, config)); err != nil {
		return fmt.Errorf("Error reading EdgeCacheOrigin: %s", err)
	}
	if err := d.Set("protocol", flattenNetworkServicesEdgeCacheOriginProtocol(res["protocol"], d, config)); err != nil {
		return fmt.Errorf("Error reading EdgeCacheOrigin: %s", err)
	}
	if err := d.Set("port", flattenNetworkServicesEdgeCacheOriginPort(res["port"], d, config)); err != nil {
		return fmt.Errorf("Error reading EdgeCacheOrigin: %s", err)
	}
	if err := d.Set("max_attempts", flattenNetworkServicesEdgeCacheOriginMaxAttempts(res["maxAttempts"], d, config)); err != nil {
		return fmt.Errorf("Error reading EdgeCacheOrigin: %s", err)
	}
	if err := d.Set("failover_origin", flattenNetworkServicesEdgeCacheOriginFailoverOrigin(res["failoverOrigin"], d, config)); err != nil {
		return fmt.Errorf("Error reading EdgeCacheOrigin: %s", err)
	}
	if err := d.Set("retry_conditions", flattenNetworkServicesEdgeCacheOriginRetryConditions(res["retryConditions"], d, config)); err != nil {
		return fmt.Errorf("Error reading EdgeCacheOrigin: %s", err)
	}
	if err := d.Set("timeout", flattenNetworkServicesEdgeCacheOriginTimeout(res["timeout"], d, config)); err != nil {
		return fmt.Errorf("Error reading EdgeCacheOrigin: %s", err)
	}

	return nil
}

func resourceNetworkServicesEdgeCacheOriginUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EdgeCacheOrigin: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesEdgeCacheOriginDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetworkServicesEdgeCacheOriginLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	originAddressProp, err := expandNetworkServicesEdgeCacheOriginOriginAddress(d.Get("origin_address"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("origin_address"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, originAddressProp)) {
		obj["originAddress"] = originAddressProp
	}
	protocolProp, err := expandNetworkServicesEdgeCacheOriginProtocol(d.Get("protocol"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("protocol"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, protocolProp)) {
		obj["protocol"] = protocolProp
	}
	portProp, err := expandNetworkServicesEdgeCacheOriginPort(d.Get("port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}
	maxAttemptsProp, err := expandNetworkServicesEdgeCacheOriginMaxAttempts(d.Get("max_attempts"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("max_attempts"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, maxAttemptsProp)) {
		obj["maxAttempts"] = maxAttemptsProp
	}
	failoverOriginProp, err := expandNetworkServicesEdgeCacheOriginFailoverOrigin(d.Get("failover_origin"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("failover_origin"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, failoverOriginProp)) {
		obj["failoverOrigin"] = failoverOriginProp
	}
	retryConditionsProp, err := expandNetworkServicesEdgeCacheOriginRetryConditions(d.Get("retry_conditions"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retry_conditions"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, retryConditionsProp)) {
		obj["retryConditions"] = retryConditionsProp
	}
	timeoutProp, err := expandNetworkServicesEdgeCacheOriginTimeout(d.Get("timeout"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("timeout"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, timeoutProp)) {
		obj["timeout"] = timeoutProp
	}

	url, err := replaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/edgeCacheOrigins/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating EdgeCacheOrigin %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("origin_address") {
		updateMask = append(updateMask, "originAddress")
	}

	if d.HasChange("protocol") {
		updateMask = append(updateMask, "protocol")
	}

	if d.HasChange("port") {
		updateMask = append(updateMask, "port")
	}

	if d.HasChange("max_attempts") {
		updateMask = append(updateMask, "maxAttempts")
	}

	if d.HasChange("failover_origin") {
		updateMask = append(updateMask, "failoverOrigin")
	}

	if d.HasChange("retry_conditions") {
		updateMask = append(updateMask, "retryConditions")
	}

	if d.HasChange("timeout") {
		updateMask = append(updateMask, "timeout")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating EdgeCacheOrigin %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating EdgeCacheOrigin %q: %#v", d.Id(), res)
	}

	err = networkServicesOperationWaitTime(
		config, res, project, "Updating EdgeCacheOrigin", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceNetworkServicesEdgeCacheOriginRead(d, meta)
}

func resourceNetworkServicesEdgeCacheOriginDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EdgeCacheOrigin: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/edgeCacheOrigins/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting EdgeCacheOrigin %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "EdgeCacheOrigin")
	}

	err = networkServicesOperationWaitTime(
		config, res, project, "Deleting EdgeCacheOrigin", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting EdgeCacheOrigin %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkServicesEdgeCacheOriginImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/global/edgeCacheOrigins/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/global/edgeCacheOrigins/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkServicesEdgeCacheOriginDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNetworkServicesEdgeCacheOriginLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNetworkServicesEdgeCacheOriginOriginAddress(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNetworkServicesEdgeCacheOriginProtocol(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNetworkServicesEdgeCacheOriginPort(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenNetworkServicesEdgeCacheOriginMaxAttempts(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenNetworkServicesEdgeCacheOriginFailoverOrigin(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNetworkServicesEdgeCacheOriginRetryConditions(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNetworkServicesEdgeCacheOriginTimeout(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["connect_timeout"] =
		flattenNetworkServicesEdgeCacheOriginTimeoutConnectTimeout(original["connectTimeout"], d, config)
	transformed["max_attempts_timeout"] =
		flattenNetworkServicesEdgeCacheOriginTimeoutMaxAttemptsTimeout(original["maxAttemptsTimeout"], d, config)
	transformed["response_timeout"] =
		flattenNetworkServicesEdgeCacheOriginTimeoutResponseTimeout(original["responseTimeout"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkServicesEdgeCacheOriginTimeoutConnectTimeout(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNetworkServicesEdgeCacheOriginTimeoutMaxAttemptsTimeout(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNetworkServicesEdgeCacheOriginTimeoutResponseTimeout(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandNetworkServicesEdgeCacheOriginDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEdgeCacheOriginLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandNetworkServicesEdgeCacheOriginOriginAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEdgeCacheOriginProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEdgeCacheOriginPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEdgeCacheOriginMaxAttempts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEdgeCacheOriginFailoverOrigin(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEdgeCacheOriginRetryConditions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEdgeCacheOriginTimeout(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedConnectTimeout, err := expandNetworkServicesEdgeCacheOriginTimeoutConnectTimeout(original["connect_timeout"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConnectTimeout); val.IsValid() && !isEmptyValue(val) {
		transformed["connectTimeout"] = transformedConnectTimeout
	}

	transformedMaxAttemptsTimeout, err := expandNetworkServicesEdgeCacheOriginTimeoutMaxAttemptsTimeout(original["max_attempts_timeout"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxAttemptsTimeout); val.IsValid() && !isEmptyValue(val) {
		transformed["maxAttemptsTimeout"] = transformedMaxAttemptsTimeout
	}

	transformedResponseTimeout, err := expandNetworkServicesEdgeCacheOriginTimeoutResponseTimeout(original["response_timeout"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponseTimeout); val.IsValid() && !isEmptyValue(val) {
		transformed["responseTimeout"] = transformedResponseTimeout
	}

	return transformed, nil
}

func expandNetworkServicesEdgeCacheOriginTimeoutConnectTimeout(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEdgeCacheOriginTimeoutMaxAttemptsTimeout(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEdgeCacheOriginTimeoutResponseTimeout(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
