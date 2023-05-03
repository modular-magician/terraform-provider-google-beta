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
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceHealthcareHl7V2Store() *schema.Resource {
	return &schema.Resource{
		Create: resourceHealthcareHl7V2StoreCreate,
		Read:   resourceHealthcareHl7V2StoreRead,
		Update: resourceHealthcareHl7V2StoreUpdate,
		Delete: resourceHealthcareHl7V2StoreDelete,

		Importer: &schema.ResourceImporter{
			State: resourceHealthcareHl7V2StoreImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dataset": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `Identifies the dataset addressed by this request. Must be in the format
'projects/{project}/locations/{location}/datasets/{dataset}'`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The resource name for the Hl7V2Store.

** Changing this property may recreate the Hl7v2 store (removing all data) **`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `User-supplied key-value pairs used to organize HL7v2 stores.

Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}

Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}

No more than 64 labels can be associated with a given store.

An object containing a list of "key": value pairs.
Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"notification_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Deprecated:  "This field has been replaced by notificationConfigs",
				Description: `A nested object resource`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pubsub_topic": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
project. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given
Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.`,
						},
					},
				},
			},
			"notification_configs": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A list of notification configs. Each configuration uses a filter to determine whether to publish a
message (both Ingest & Create) on the corresponding notification destination. Only the message name
is sent as part of the notification. Supplied by the client.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pubsub_topic": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
project. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given
Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.

If a notification cannot be published to Cloud Pub/Sub, errors will be logged to Stackdriver`,
						},
						"filter": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Restricts notifications sent for messages matching a filter. If this is empty, all messages
are matched. Syntax: https://cloud.google.com/appengine/docs/standard/python/search/query_strings

Fields/functions available for filtering are:

* messageType, from the MSH-9.1 field. For example, NOT messageType = "ADT".
* send_date or sendDate, the YYYY-MM-DD date the message was sent in the dataset's timeZone, from the MSH-7 segment. For example, send_date < "2017-01-02".
* sendTime, the timestamp when the message was sent, using the RFC3339 time format for comparisons, from the MSH-7 segment. For example, sendTime < "2017-01-02T00:00:00-05:00".
* sendFacility, the care center that the message came from, from the MSH-4 segment. For example, sendFacility = "ABC".
* PatientId(value, type), which matches if the message lists a patient having an ID of the given value and type in the PID-2, PID-3, or PID-4 segments. For example, PatientId("123456", "MRN").
* labels.x, a string value of the label with key x as set using the Message.labels map. For example, labels."priority"="high". The operator :* can be used to assert the existence of a label. For example, labels."priority":*.`,
						},
					},
				},
			},
			"parser_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: `A nested object resource`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_null_header": {
							Type:         schema.TypeBool,
							Optional:     true,
							Description:  `Determines whether messages with no header are allowed.`,
							AtLeastOneOf: []string{"parser_config.0.allow_null_header", "parser_config.0.segment_terminator", "parser_config.0.schema"},
						},
						"schema": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringIsJSON,
							StateFunc:    func(v interface{}) string { s, _ := structure.NormalizeJsonString(v); return s },
							Description: `JSON encoded string for schemas used to parse messages in this
store if schematized parsing is desired.`,
							AtLeastOneOf: []string{"parser_config.0.allow_null_header", "parser_config.0.segment_terminator", "parser_config.0.schema", "parser_config.0.version"},
						},
						"segment_terminator": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateBase64String,
							Description: `Byte(s) to be used as the segment terminator. If this is unset, '\r' will be used as segment terminator.

A base64-encoded string.`,
							AtLeastOneOf: []string{"parser_config.0.allow_null_header", "parser_config.0.segment_terminator", "parser_config.0.schema"},
						},
						"version": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: verify.ValidateEnum([]string{"V1", "V2", "V3", ""}),
							Description:  `The version of the unschematized parser to be used when a custom 'schema' is not set. Default value: "V1" Possible values: ["V1", "V2", "V3"]`,
							Default:      "V1",
						},
					},
				},
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The fully qualified name of this dataset`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceHealthcareHl7V2StoreCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandHealthcareHl7V2StoreName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	parserConfigProp, err := expandHealthcareHl7V2StoreParserConfig(d.Get("parser_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parser_config"); !isEmptyValue(reflect.ValueOf(parserConfigProp)) && (ok || !reflect.DeepEqual(v, parserConfigProp)) {
		obj["parserConfig"] = parserConfigProp
	}
	labelsProp, err := expandHealthcareHl7V2StoreLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	notificationConfigsProp, err := expandHealthcareHl7V2StoreNotificationConfigs(d.Get("notification_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_configs"); !isEmptyValue(reflect.ValueOf(notificationConfigsProp)) && (ok || !reflect.DeepEqual(v, notificationConfigsProp)) {
		obj["notificationConfigs"] = notificationConfigsProp
	}
	notificationConfigProp, err := expandHealthcareHl7V2StoreNotificationConfig(d.Get("notification_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_config"); !isEmptyValue(reflect.ValueOf(notificationConfigProp)) && (ok || !reflect.DeepEqual(v, notificationConfigProp)) {
		obj["notificationConfig"] = notificationConfigProp
	}

	url, err := ReplaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/hl7V2Stores?hl7V2StoreId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Hl7V2Store: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Hl7V2Store: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{dataset}}/hl7V2Stores/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Hl7V2Store %q: %#v", d.Id(), res)

	return resourceHealthcareHl7V2StoreRead(d, meta)
}

func resourceHealthcareHl7V2StoreRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/hl7V2Stores/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("HealthcareHl7V2Store %q", d.Id()))
	}

	res, err = resourceHealthcareHl7V2StoreDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing HealthcareHl7V2Store because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("name", flattenHealthcareHl7V2StoreName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Hl7V2Store: %s", err)
	}
	if err := d.Set("parser_config", flattenHealthcareHl7V2StoreParserConfig(res["parserConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Hl7V2Store: %s", err)
	}
	if err := d.Set("labels", flattenHealthcareHl7V2StoreLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Hl7V2Store: %s", err)
	}
	if err := d.Set("notification_configs", flattenHealthcareHl7V2StoreNotificationConfigs(res["notificationConfigs"], d, config)); err != nil {
		return fmt.Errorf("Error reading Hl7V2Store: %s", err)
	}
	if err := d.Set("notification_config", flattenHealthcareHl7V2StoreNotificationConfig(res["notificationConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Hl7V2Store: %s", err)
	}

	return nil
}

func resourceHealthcareHl7V2StoreUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	parserConfigProp, err := expandHealthcareHl7V2StoreParserConfig(d.Get("parser_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parser_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, parserConfigProp)) {
		obj["parserConfig"] = parserConfigProp
	}
	labelsProp, err := expandHealthcareHl7V2StoreLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	notificationConfigsProp, err := expandHealthcareHl7V2StoreNotificationConfigs(d.Get("notification_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_configs"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, notificationConfigsProp)) {
		obj["notificationConfigs"] = notificationConfigsProp
	}
	notificationConfigProp, err := expandHealthcareHl7V2StoreNotificationConfig(d.Get("notification_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, notificationConfigProp)) {
		obj["notificationConfig"] = notificationConfigProp
	}

	url, err := ReplaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/hl7V2Stores/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Hl7V2Store %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("parser_config") {
		updateMask = append(updateMask, "parser_config.allow_null_header",
			"parser_config.segment_terminator",
			"parser_config.schema")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("notification_configs") {
		updateMask = append(updateMask, "notificationConfigs")
	}

	if d.HasChange("notification_config") {
		updateMask = append(updateMask, "notificationConfig")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Hl7V2Store %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Hl7V2Store %q: %#v", d.Id(), res)
	}

	return resourceHealthcareHl7V2StoreRead(d, meta)
}

func resourceHealthcareHl7V2StoreDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := ReplaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/hl7V2Stores/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Hl7V2Store %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Hl7V2Store")
	}

	log.Printf("[DEBUG] Finished deleting Hl7V2Store %q: %#v", d.Id(), res)
	return nil
}

func resourceHealthcareHl7V2StoreImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*transport_tpg.Config)

	hl7v2StoreId, err := ParseHealthcareHl7V2StoreId(d.Id(), config)
	if err != nil {
		return nil, err
	}

	if err := d.Set("dataset", hl7v2StoreId.DatasetId.DatasetId()); err != nil {
		return nil, fmt.Errorf("Error setting dataset: %s", err)
	}
	if err := d.Set("name", hl7v2StoreId.Name); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenHealthcareHl7V2StoreName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenHealthcareHl7V2StoreParserConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["allow_null_header"] =
		flattenHealthcareHl7V2StoreParserConfigAllowNullHeader(original["allowNullHeader"], d, config)
	transformed["segment_terminator"] =
		flattenHealthcareHl7V2StoreParserConfigSegmentTerminator(original["segmentTerminator"], d, config)
	transformed["schema"] =
		flattenHealthcareHl7V2StoreParserConfigSchema(original["schema"], d, config)
	transformed["version"] =
		flattenHealthcareHl7V2StoreParserConfigVersion(original["version"], d, config)
	return []interface{}{transformed}
}
func flattenHealthcareHl7V2StoreParserConfigAllowNullHeader(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenHealthcareHl7V2StoreParserConfigSegmentTerminator(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenHealthcareHl7V2StoreParserConfigSchema(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		// TODO: return error once https://github.com/GoogleCloudPlatform/magic-modules/issues/3257 is fixed.
		log.Printf("[ERROR] failed to marshal schema to JSON: %v", err)
	}
	return string(b)
}

func flattenHealthcareHl7V2StoreParserConfigVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenHealthcareHl7V2StoreLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenHealthcareHl7V2StoreNotificationConfigs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"pubsub_topic": flattenHealthcareHl7V2StoreNotificationConfigsPubsubTopic(original["pubsubTopic"], d, config),
			"filter":       flattenHealthcareHl7V2StoreNotificationConfigsFilter(original["filter"], d, config),
		})
	}
	return transformed
}
func flattenHealthcareHl7V2StoreNotificationConfigsPubsubTopic(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenHealthcareHl7V2StoreNotificationConfigsFilter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenHealthcareHl7V2StoreNotificationConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["pubsub_topic"] =
		flattenHealthcareHl7V2StoreNotificationConfigPubsubTopic(original["pubsubTopic"], d, config)
	return []interface{}{transformed}
}
func flattenHealthcareHl7V2StoreNotificationConfigPubsubTopic(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandHealthcareHl7V2StoreName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareHl7V2StoreParserConfig(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowNullHeader, err := expandHealthcareHl7V2StoreParserConfigAllowNullHeader(original["allow_null_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowNullHeader); val.IsValid() && !isEmptyValue(val) {
		transformed["allowNullHeader"] = transformedAllowNullHeader
	}

	transformedSegmentTerminator, err := expandHealthcareHl7V2StoreParserConfigSegmentTerminator(original["segment_terminator"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSegmentTerminator); val.IsValid() && !isEmptyValue(val) {
		transformed["segmentTerminator"] = transformedSegmentTerminator
	}

	transformedSchema, err := expandHealthcareHl7V2StoreParserConfigSchema(original["schema"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSchema); val.IsValid() && !isEmptyValue(val) {
		transformed["schema"] = transformedSchema
	}

	transformedVersion, err := expandHealthcareHl7V2StoreParserConfigVersion(original["version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !isEmptyValue(val) {
		transformed["version"] = transformedVersion
	}

	return transformed, nil
}

func expandHealthcareHl7V2StoreParserConfigAllowNullHeader(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareHl7V2StoreParserConfigSegmentTerminator(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareHl7V2StoreParserConfigSchema(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	b := []byte(v.(string))
	if len(b) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func expandHealthcareHl7V2StoreParserConfigVersion(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareHl7V2StoreLabels(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandHealthcareHl7V2StoreNotificationConfigs(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPubsubTopic, err := expandHealthcareHl7V2StoreNotificationConfigsPubsubTopic(original["pubsub_topic"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPubsubTopic); val.IsValid() && !isEmptyValue(val) {
			transformed["pubsubTopic"] = transformedPubsubTopic
		}

		transformedFilter, err := expandHealthcareHl7V2StoreNotificationConfigsFilter(original["filter"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFilter); val.IsValid() && !isEmptyValue(val) {
			transformed["filter"] = transformedFilter
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandHealthcareHl7V2StoreNotificationConfigsPubsubTopic(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareHl7V2StoreNotificationConfigsFilter(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareHl7V2StoreNotificationConfig(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPubsubTopic, err := expandHealthcareHl7V2StoreNotificationConfigPubsubTopic(original["pubsub_topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubTopic); val.IsValid() && !isEmptyValue(val) {
		transformed["pubsubTopic"] = transformedPubsubTopic
	}

	return transformed, nil
}

func expandHealthcareHl7V2StoreNotificationConfigPubsubTopic(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceHealthcareHl7V2StoreDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// Take the returned long form of the name and use it as `self_link`.
	// Then modify the name to be the user specified form.
	// We can't just ignore_read on `name` as the linter will
	// complain that the returned `res` is never used afterwards.
	// Some field needs to be actually set, and we chose `name`.
	if err := d.Set("self_link", res["name"].(string)); err != nil {
		return nil, fmt.Errorf("Error setting self_link: %s", err)
	}
	res["name"] = d.Get("name").(string)
	return res, nil
}
