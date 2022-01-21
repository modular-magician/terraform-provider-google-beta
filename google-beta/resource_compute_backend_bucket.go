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
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceComputeBackendBucket() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeBackendBucketCreate,
		Read:   resourceComputeBackendBucketRead,
		Update: resourceComputeBackendBucketUpdate,
		Delete: resourceComputeBackendBucketDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeBackendBucketImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Cloud Storage bucket name.`,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRegexp(`^(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)$`),
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035.  Specifically, the name must be 1-63 characters long and
match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
the first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the
last character, which cannot be a dash.`,
			},
			"cdn_policy": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: `Cloud CDN configuration for this Backend Bucket.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cache_mode": {
							Type:         schema.TypeString,
							Computed:     true,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"USE_ORIGIN_HEADERS", "FORCE_CACHE_ALL", "CACHE_ALL_STATIC", ""}, false),
							Description: `Specifies the cache setting for all responses from this backend.
The possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL and CACHE_ALL_STATIC Possible values: ["USE_ORIGIN_HEADERS", "FORCE_CACHE_ALL", "CACHE_ALL_STATIC"]`,
						},
						"client_ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: `Specifies the maximum allowed TTL for cached content served by this origin.`,
						},
						"default_ttl": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
							Description: `Specifies the default TTL for cached content served by this origin for responses
that do not have an existing valid TTL (max-age or s-max-age).`,
						},
						"max_ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: `Specifies the maximum allowed TTL for cached content served by this origin.`,
						},
						"negative_caching": {
							Type:        schema.TypeBool,
							Computed:    true,
							Optional:    true,
							Description: `Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects.`,
						},
						"negative_caching_policy": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `Sets a cache TTL for the specified HTTP status code. negativeCaching must be enabled to configure negativeCachingPolicy.
Omitting the policy and leaving negativeCaching enabled will use Cloud CDN's default cache TTLs.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": {
										Type:     schema.TypeInt,
										Optional: true,
										Description: `The HTTP status code to define a TTL against. Only HTTP status codes 300, 301, 308, 404, 405, 410, 421, 451 and 501
can be specified as values, and you cannot specify a status code more than once.`,
									},
									"ttl": {
										Type:     schema.TypeInt,
										Optional: true,
										Description: `The TTL (in seconds) for which to cache responses with the corresponding status code. The maximum allowed value is 1800s
(30 minutes), noting that infrequently accessed objects may be evicted from the cache before the defined TTL.`,
									},
								},
							},
						},
						"serve_while_stale": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: `Serve existing content from the cache (if available) when revalidating content with the origin, or when an error is encountered when refreshing the cache.`,
						},
						"signed_url_cache_max_age_sec": {
							Type:     schema.TypeInt,
							Optional: true,
							Description: `Maximum number of seconds the response to a signed URL request will
be considered fresh. After this time period,
the response will be revalidated before being served.
When serving responses to signed URL requests,
Cloud CDN will internally behave as though
all responses from this backend had a "Cache-Control: public,
max-age=[TTL]" header, regardless of any existing Cache-Control
header. The actual headers served in responses will not be altered.`,
						},
					},
				},
			},
			"custom_response_headers": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Headers that the HTTP/S load balancer should add to proxied responses.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `An optional textual description of the resource; provided by the
client when the resource is created.`,
			},
			"enable_cdn": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `If true, enable Cloud CDN for this BackendBucket.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
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
		UseJSONNumber: true,
	}
}

func resourceComputeBackendBucketCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	bucketNameProp, err := expandComputeBackendBucketBucketName(d.Get("bucket_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bucket_name"); !isEmptyValue(reflect.ValueOf(bucketNameProp)) && (ok || !reflect.DeepEqual(v, bucketNameProp)) {
		obj["bucketName"] = bucketNameProp
	}
	cdnPolicyProp, err := expandComputeBackendBucketCdnPolicy(d.Get("cdn_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cdn_policy"); !isEmptyValue(reflect.ValueOf(cdnPolicyProp)) && (ok || !reflect.DeepEqual(v, cdnPolicyProp)) {
		obj["cdnPolicy"] = cdnPolicyProp
	}
	customResponseHeadersProp, err := expandComputeBackendBucketCustomResponseHeaders(d.Get("custom_response_headers"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("custom_response_headers"); !isEmptyValue(reflect.ValueOf(customResponseHeadersProp)) && (ok || !reflect.DeepEqual(v, customResponseHeadersProp)) {
		obj["customResponseHeaders"] = customResponseHeadersProp
	}
	descriptionProp, err := expandComputeBackendBucketDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	enableCdnProp, err := expandComputeBackendBucketEnableCdn(d.Get("enable_cdn"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_cdn"); !isEmptyValue(reflect.ValueOf(enableCdnProp)) && (ok || !reflect.DeepEqual(v, enableCdnProp)) {
		obj["enableCdn"] = enableCdnProp
	}
	nameProp, err := expandComputeBackendBucketName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/backendBuckets")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BackendBucket: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackendBucket: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating BackendBucket: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating BackendBucket", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create BackendBucket: %s", err)
	}

	log.Printf("[DEBUG] Finished creating BackendBucket %q: %#v", d.Id(), res)

	return resourceComputeBackendBucketRead(d, meta)
}

func resourceComputeBackendBucketRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackendBucket: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeBackendBucket %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}

	if err := d.Set("bucket_name", flattenComputeBackendBucketBucketName(res["bucketName"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}
	if err := d.Set("cdn_policy", flattenComputeBackendBucketCdnPolicy(res["cdnPolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}
	if err := d.Set("custom_response_headers", flattenComputeBackendBucketCustomResponseHeaders(res["customResponseHeaders"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeBackendBucketCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}
	if err := d.Set("description", flattenComputeBackendBucketDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}
	if err := d.Set("enable_cdn", flattenComputeBackendBucketEnableCdn(res["enableCdn"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}
	if err := d.Set("name", flattenComputeBackendBucketName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading BackendBucket: %s", err)
	}

	return nil
}

func resourceComputeBackendBucketUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackendBucket: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	bucketNameProp, err := expandComputeBackendBucketBucketName(d.Get("bucket_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bucket_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, bucketNameProp)) {
		obj["bucketName"] = bucketNameProp
	}
	cdnPolicyProp, err := expandComputeBackendBucketCdnPolicy(d.Get("cdn_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cdn_policy"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, cdnPolicyProp)) {
		obj["cdnPolicy"] = cdnPolicyProp
	}
	customResponseHeadersProp, err := expandComputeBackendBucketCustomResponseHeaders(d.Get("custom_response_headers"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("custom_response_headers"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, customResponseHeadersProp)) {
		obj["customResponseHeaders"] = customResponseHeadersProp
	}
	descriptionProp, err := expandComputeBackendBucketDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	enableCdnProp, err := expandComputeBackendBucketEnableCdn(d.Get("enable_cdn"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_cdn"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableCdnProp)) {
		obj["enableCdn"] = enableCdnProp
	}
	nameProp, err := expandComputeBackendBucketName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating BackendBucket %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PUT", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating BackendBucket %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating BackendBucket %q: %#v", d.Id(), res)
	}

	err = computeOperationWaitTime(
		config, res, project, "Updating BackendBucket", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputeBackendBucketRead(d, meta)
}

func resourceComputeBackendBucketDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackendBucket: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting BackendBucket %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "BackendBucket")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting BackendBucket", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting BackendBucket %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeBackendBucketImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/backendBuckets/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeBackendBucketBucketName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeBackendBucketCdnPolicy(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["signed_url_cache_max_age_sec"] =
		flattenComputeBackendBucketCdnPolicySignedUrlCacheMaxAgeSec(original["signedUrlCacheMaxAgeSec"], d, config)
	transformed["default_ttl"] =
		flattenComputeBackendBucketCdnPolicyDefaultTtl(original["defaultTtl"], d, config)
	transformed["max_ttl"] =
		flattenComputeBackendBucketCdnPolicyMaxTtl(original["maxTtl"], d, config)
	transformed["client_ttl"] =
		flattenComputeBackendBucketCdnPolicyClientTtl(original["clientTtl"], d, config)
	transformed["negative_caching"] =
		flattenComputeBackendBucketCdnPolicyNegativeCaching(original["negativeCaching"], d, config)
	transformed["negative_caching_policy"] =
		flattenComputeBackendBucketCdnPolicyNegativeCachingPolicy(original["negativeCachingPolicy"], d, config)
	transformed["cache_mode"] =
		flattenComputeBackendBucketCdnPolicyCacheMode(original["cacheMode"], d, config)
	transformed["serve_while_stale"] =
		flattenComputeBackendBucketCdnPolicyServeWhileStale(original["serveWhileStale"], d, config)
	return []interface{}{transformed}
}
func flattenComputeBackendBucketCdnPolicySignedUrlCacheMaxAgeSec(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenComputeBackendBucketCdnPolicyDefaultTtl(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenComputeBackendBucketCdnPolicyMaxTtl(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenComputeBackendBucketCdnPolicyClientTtl(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenComputeBackendBucketCdnPolicyNegativeCaching(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeBackendBucketCdnPolicyNegativeCachingPolicy(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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
			"code": flattenComputeBackendBucketCdnPolicyNegativeCachingPolicyCode(original["code"], d, config),
			"ttl":  flattenComputeBackendBucketCdnPolicyNegativeCachingPolicyTtl(original["ttl"], d, config),
		})
	}
	return transformed
}
func flattenComputeBackendBucketCdnPolicyNegativeCachingPolicyCode(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenComputeBackendBucketCdnPolicyNegativeCachingPolicyTtl(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenComputeBackendBucketCdnPolicyCacheMode(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeBackendBucketCdnPolicyServeWhileStale(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenComputeBackendBucketCustomResponseHeaders(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeBackendBucketCreationTimestamp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeBackendBucketDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeBackendBucketEnableCdn(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeBackendBucketName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandComputeBackendBucketBucketName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSignedUrlCacheMaxAgeSec, err := expandComputeBackendBucketCdnPolicySignedUrlCacheMaxAgeSec(original["signed_url_cache_max_age_sec"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["signedUrlCacheMaxAgeSec"] = transformedSignedUrlCacheMaxAgeSec
	}

	transformedDefaultTtl, err := expandComputeBackendBucketCdnPolicyDefaultTtl(original["default_ttl"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["defaultTtl"] = transformedDefaultTtl
	}

	transformedMaxTtl, err := expandComputeBackendBucketCdnPolicyMaxTtl(original["max_ttl"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxTtl); val.IsValid() && !isEmptyValue(val) {
		transformed["maxTtl"] = transformedMaxTtl
	}

	transformedClientTtl, err := expandComputeBackendBucketCdnPolicyClientTtl(original["client_ttl"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["clientTtl"] = transformedClientTtl
	}

	transformedNegativeCaching, err := expandComputeBackendBucketCdnPolicyNegativeCaching(original["negative_caching"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["negativeCaching"] = transformedNegativeCaching
	}

	transformedNegativeCachingPolicy, err := expandComputeBackendBucketCdnPolicyNegativeCachingPolicy(original["negative_caching_policy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNegativeCachingPolicy); val.IsValid() && !isEmptyValue(val) {
		transformed["negativeCachingPolicy"] = transformedNegativeCachingPolicy
	}

	transformedCacheMode, err := expandComputeBackendBucketCdnPolicyCacheMode(original["cache_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCacheMode); val.IsValid() && !isEmptyValue(val) {
		transformed["cacheMode"] = transformedCacheMode
	}

	transformedServeWhileStale, err := expandComputeBackendBucketCdnPolicyServeWhileStale(original["serve_while_stale"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["serveWhileStale"] = transformedServeWhileStale
	}

	return transformed, nil
}

func expandComputeBackendBucketCdnPolicySignedUrlCacheMaxAgeSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyDefaultTtl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyMaxTtl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyClientTtl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyNegativeCaching(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyNegativeCachingPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedCode, err := expandComputeBackendBucketCdnPolicyNegativeCachingPolicyCode(original["code"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCode); val.IsValid() && !isEmptyValue(val) {
			transformed["code"] = transformedCode
		}

		transformedTtl, err := expandComputeBackendBucketCdnPolicyNegativeCachingPolicyTtl(original["ttl"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["ttl"] = transformedTtl
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeBackendBucketCdnPolicyNegativeCachingPolicyCode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyNegativeCachingPolicyTtl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyCacheMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyServeWhileStale(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCustomResponseHeaders(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketEnableCdn(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
