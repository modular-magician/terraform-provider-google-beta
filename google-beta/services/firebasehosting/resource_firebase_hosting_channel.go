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

package firebasehosting

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
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

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
		),

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
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Text labels used for extra metadata and/or filtering

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
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
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The fully-qualified resource name for the channel, in the format:
sites/SITE_ID/channels/CHANNEL_ID`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
		},
		UseJSONNumber: true,
	}
}

func resourceFirebaseHostingChannelCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	retainedReleaseCountProp, err := expandFirebaseHostingChannelRetainedReleaseCount(d.Get("retained_release_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retained_release_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(retainedReleaseCountProp)) && (ok || !reflect.DeepEqual(v, retainedReleaseCountProp)) {
		obj["retainedReleaseCount"] = retainedReleaseCountProp
	}
	expireTimeProp, err := expandFirebaseHostingChannelExpireTime(d.Get("expire_time"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expire_time"); !tpgresource.IsEmptyValue(reflect.ValueOf(expireTimeProp)) && (ok || !reflect.DeepEqual(v, expireTimeProp)) {
		obj["expireTime"] = expireTimeProp
	}
	ttlProp, err := expandFirebaseHostingChannelTtl(d.Get("ttl"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(ttlProp)) && (ok || !reflect.DeepEqual(v, ttlProp)) {
		obj["ttl"] = ttlProp
	}
	labelsProp, err := expandFirebaseHostingChannelEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseHostingBasePath}}sites/{{site_id}}/channels?channelId={{channel_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Channel: %#v", obj)
	billingProject := ""

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
		return fmt.Errorf("Error creating Channel: %s", err)
	}
	if err := d.Set("name", flattenFirebaseHostingChannelName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "sites/{{site_id}}/channels/{{channel_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Channel %q: %#v", d.Id(), res)

	return resourceFirebaseHostingChannelRead(d, meta)
}

func resourceFirebaseHostingChannelRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseHostingBasePath}}sites/{{site_id}}/channels/{{channel_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("FirebaseHostingChannel %q", d.Id()))
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
	if err := d.Set("terraform_labels", flattenFirebaseHostingChannelTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Channel: %s", err)
	}
	if err := d.Set("effective_labels", flattenFirebaseHostingChannelEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Channel: %s", err)
	}

	return nil
}

func resourceFirebaseHostingChannelUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	retainedReleaseCountProp, err := expandFirebaseHostingChannelRetainedReleaseCount(d.Get("retained_release_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retained_release_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, retainedReleaseCountProp)) {
		obj["retainedReleaseCount"] = retainedReleaseCountProp
	}
	expireTimeProp, err := expandFirebaseHostingChannelExpireTime(d.Get("expire_time"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expire_time"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, expireTimeProp)) {
		obj["expireTime"] = expireTimeProp
	}
	labelsProp, err := expandFirebaseHostingChannelEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseHostingBasePath}}sites/{{site_id}}/channels/{{channel_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Channel %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("retained_release_count") {
		updateMask = append(updateMask, "retainedReleaseCount")
	}

	if d.HasChange("expire_time") {
		updateMask = append(updateMask, "expireTime")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
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
		return fmt.Errorf("Error updating Channel %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Channel %q: %#v", d.Id(), res)
	}

	return resourceFirebaseHostingChannelRead(d, meta)
}

func resourceFirebaseHostingChannelDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseHostingBasePath}}sites/{{site_id}}/channels/{{channel_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Channel %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Channel")
	}

	log.Printf("[DEBUG] Finished deleting Channel %q: %#v", d.Id(), res)
	return nil
}

func resourceFirebaseHostingChannelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"sites/(?P<site_id>[^/]+)/channels/(?P<channel_id>[^/]+)",
		"(?P<site_id>[^/]+)/(?P<channel_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "sites/{{site_id}}/channels/{{channel_id}}")
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

func flattenFirebaseHostingChannelLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenFirebaseHostingChannelExpireTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseHostingChannelTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenFirebaseHostingChannelEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandFirebaseHostingChannelRetainedReleaseCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingChannelExpireTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingChannelTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingChannelEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
