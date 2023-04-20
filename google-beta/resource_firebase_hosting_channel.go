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
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceFirebaseHostingChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaseHostingChannelCreate,
		Read:   resourceFirebaseHostingChannelRead,
		Update: resourceFirebaseHostingChannelUpdate,
		Delete: resourceFirebaseHostingChannelDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaseHostingChannelImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"channel_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Required. Immutable. A unique ID within the site that identifies the channel.`,
			},
			"site_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Required. The ID of the site in which to create this channel.`,
			},
			"expire_time": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				Description: `The time at which the channel will be automatically deleted. If null, the channel
will not be automatically deleted. This field is present in the output whether it's
set directly or via the 'ttl' field.`,
				ConflictsWith: []string{"ttl"},
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Text labels used for extra metadata and/or filtering`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"retained_release_count": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
				Description: `The number of previous releases to retain on the channel for rollback or other
purposes. Must be a number between 1-100. Defaults to 10 for new channels.`,
			},
			"ttl": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Input only. A time-to-live for this channel. Sets 'expire_time' to the provided
duration past the time of the request. A duration in seconds with up to nine fractional
digits, terminated by 's'. Example: "86400s" (one day).`,
				ConflictsWith: []string{"expire_time"},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The fully-qualified resource name for the channel, in the format:
sites/SITE_ID/channels/CHANNEL_ID`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceFirebaseHostingChannelCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	retainedReleaseCountProp, err := expandFirebaseHostingChannelRetainedReleaseCount(d.Get("retained_release_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retained_release_count"); !isEmptyValue(reflect.ValueOf(retainedReleaseCountProp)) && (ok || !reflect.DeepEqual(v, retainedReleaseCountProp)) {
		obj["retainedReleaseCount"] = retainedReleaseCountProp
	}
	labelsProp, err := expandFirebaseHostingChannelLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	expireTimeProp, err := expandFirebaseHostingChannelExpireTime(d.Get("expire_time"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expire_time"); !isEmptyValue(reflect.ValueOf(expireTimeProp)) && (ok || !reflect.DeepEqual(v, expireTimeProp)) {
		obj["expireTime"] = expireTimeProp
	}
	ttlProp, err := expandFirebaseHostingChannelTtl(d.Get("ttl"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ttl"); !isEmptyValue(reflect.ValueOf(ttlProp)) && (ok || !reflect.DeepEqual(v, ttlProp)) {
		obj["ttl"] = ttlProp
	}

	url, err := ReplaceVars(d, config, "{{FirebaseHostingBasePath}}sites/{{site_id}}/channels?channelId={{channel_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Channel: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Channel: %s", err)
	}
	if err := d.Set("name", flattenFirebaseHostingChannelName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "sites/{{site_id}}/channels/{{channel_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Channel %q: %#v", d.Id(), res)

	return resourceFirebaseHostingChannelRead(d, meta)
}

func resourceFirebaseHostingChannelRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{FirebaseHostingBasePath}}sites/{{site_id}}/channels/{{channel_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("FirebaseHostingChannel %q", d.Id()))
	}

	if err := d.Set("name", flattenFirebaseHostingChannelName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Channel: %s", err)
	}
	if err := d.Set("retained_release_count", flattenFirebaseHostingChannelRetainedReleaseCount(res["retainedReleaseCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading Channel: %s", err)
	}
	if err := d.Set("labels", flattenFirebaseHostingChannelLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Channel: %s", err)
	}
	if err := d.Set("expire_time", flattenFirebaseHostingChannelExpireTime(res["expireTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Channel: %s", err)
	}

	return nil
}

func resourceFirebaseHostingChannelUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	retainedReleaseCountProp, err := expandFirebaseHostingChannelRetainedReleaseCount(d.Get("retained_release_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retained_release_count"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, retainedReleaseCountProp)) {
		obj["retainedReleaseCount"] = retainedReleaseCountProp
	}
	labelsProp, err := expandFirebaseHostingChannelLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	expireTimeProp, err := expandFirebaseHostingChannelExpireTime(d.Get("expire_time"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expire_time"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, expireTimeProp)) {
		obj["expireTime"] = expireTimeProp
	}

	url, err := ReplaceVars(d, config, "{{FirebaseHostingBasePath}}sites/{{site_id}}/channels/{{channel_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Channel %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("retained_release_count") {
		updateMask = append(updateMask, "retainedReleaseCount")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("expire_time") {
		updateMask = append(updateMask, "expireTime")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Channel %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Channel %q: %#v", d.Id(), res)
	}

	return resourceFirebaseHostingChannelRead(d, meta)
}

func resourceFirebaseHostingChannelDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := ReplaceVars(d, config, "{{FirebaseHostingBasePath}}sites/{{site_id}}/channels/{{channel_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Channel %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Channel")
	}

	log.Printf("[DEBUG] Finished deleting Channel %q: %#v", d.Id(), res)
	return nil
}

func resourceFirebaseHostingChannelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"sites/(?P<site_id>[^/]+)/channels/(?P<channel_id>[^/]+)",
		"(?P<site_id>[^/]+)/(?P<channel_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "sites/{{site_id}}/channels/{{channel_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenFirebaseHostingChannelName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseHostingChannelRetainedReleaseCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenFirebaseHostingChannelLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseHostingChannelExpireTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandFirebaseHostingChannelRetainedReleaseCount(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingChannelLabels(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandFirebaseHostingChannelExpireTime(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingChannelTtl(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
