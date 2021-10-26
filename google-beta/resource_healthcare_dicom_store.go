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
)

func resourceHealthcareDicomStore() *schema.Resource {
	return &schema.Resource{
		Create: resourceHealthcareDicomStoreCreate,
		Read:   resourceHealthcareDicomStoreRead,
		Update: resourceHealthcareDicomStoreUpdate,
		Delete: resourceHealthcareDicomStoreDelete,

		Importer: &schema.ResourceImporter{
			State: resourceHealthcareDicomStoreImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
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
				Description: `The resource name for the DicomStore.

** Changing this property may recreate the Dicom store (removing all data) **`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `User-supplied key-value pairs used to organize DICOM stores.

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
			"stream_configs": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `To enable streaming to BigQuery, configure the streamConfigs object in your DICOM store.
streamConfigs is an array, so you can specify multiple BigQuery destinations. You can stream metadata from a single DICOM store to up to five BigQuery tables in a BigQuery dataset.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bigquery_destination": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `BigQueryDestination to include a fully qualified BigQuery table URI where DICOM instance metadata will be streamed.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"table_uri": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `a fully qualified BigQuery table URI where DICOM instance metadata will be streamed.`,
									},
								},
							},
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

func resourceHealthcareDicomStoreCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandHealthcareDicomStoreName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	labelsProp, err := expandHealthcareDicomStoreLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	notificationConfigProp, err := expandHealthcareDicomStoreNotificationConfig(d.Get("notification_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_config"); !isEmptyValue(reflect.ValueOf(notificationConfigProp)) && (ok || !reflect.DeepEqual(v, notificationConfigProp)) {
		obj["notificationConfig"] = notificationConfigProp
	}
	streamConfigsProp, err := expandHealthcareDicomStoreStreamConfigs(d.Get("stream_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("stream_configs"); !isEmptyValue(reflect.ValueOf(streamConfigsProp)) && (ok || !reflect.DeepEqual(v, streamConfigsProp)) {
		obj["streamConfigs"] = streamConfigsProp
	}

	url, err := replaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/dicomStores?dicomStoreId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new DicomStore: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating DicomStore: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{dataset}}/dicomStores/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating DicomStore %q: %#v", d.Id(), res)

	return resourceHealthcareDicomStoreRead(d, meta)
}

func resourceHealthcareDicomStoreRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/dicomStores/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("HealthcareDicomStore %q", d.Id()))
	}

	res, err = resourceHealthcareDicomStoreDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing HealthcareDicomStore because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("name", flattenHealthcareDicomStoreName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading DicomStore: %s", err)
	}
	if err := d.Set("labels", flattenHealthcareDicomStoreLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading DicomStore: %s", err)
	}
	if err := d.Set("notification_config", flattenHealthcareDicomStoreNotificationConfig(res["notificationConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading DicomStore: %s", err)
	}
	if err := d.Set("stream_configs", flattenHealthcareDicomStoreStreamConfigs(res["streamConfigs"], d, config)); err != nil {
		return fmt.Errorf("Error reading DicomStore: %s", err)
	}

	return nil
}

func resourceHealthcareDicomStoreUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	labelsProp, err := expandHealthcareDicomStoreLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	notificationConfigProp, err := expandHealthcareDicomStoreNotificationConfig(d.Get("notification_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, notificationConfigProp)) {
		obj["notificationConfig"] = notificationConfigProp
	}
	streamConfigsProp, err := expandHealthcareDicomStoreStreamConfigs(d.Get("stream_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("stream_configs"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, streamConfigsProp)) {
		obj["streamConfigs"] = streamConfigsProp
	}

	url, err := replaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/dicomStores/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating DicomStore %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("notification_config") {
		updateMask = append(updateMask, "notificationConfig")
	}

	if d.HasChange("stream_configs") {
		updateMask = append(updateMask, "streamConfigs")
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
		return fmt.Errorf("Error updating DicomStore %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating DicomStore %q: %#v", d.Id(), res)
	}

	return resourceHealthcareDicomStoreRead(d, meta)
}

func resourceHealthcareDicomStoreDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := replaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/dicomStores/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting DicomStore %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "DicomStore")
	}

	log.Printf("[DEBUG] Finished deleting DicomStore %q: %#v", d.Id(), res)
	return nil
}

func resourceHealthcareDicomStoreImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	dicomStoreId, err := parseHealthcareDicomStoreId(d.Id(), config)
	if err != nil {
		return nil, err
	}

	if err := d.Set("dataset", dicomStoreId.DatasetId.datasetId()); err != nil {
		return nil, fmt.Errorf("Error setting dataset: %s", err)
	}
	if err := d.Set("name", dicomStoreId.Name); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenHealthcareDicomStoreName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenHealthcareDicomStoreLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenHealthcareDicomStoreNotificationConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["pubsub_topic"] =
		flattenHealthcareDicomStoreNotificationConfigPubsubTopic(original["pubsubTopic"], d, config)
	return []interface{}{transformed}
}
func flattenHealthcareDicomStoreNotificationConfigPubsubTopic(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenHealthcareDicomStoreStreamConfigs(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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
			"bigquery_destination": flattenHealthcareDicomStoreStreamConfigsBigqueryDestination(original["bigqueryDestination"], d, config),
		})
	}
	return transformed
}
func flattenHealthcareDicomStoreStreamConfigsBigqueryDestination(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["table_uri"] =
		flattenHealthcareDicomStoreStreamConfigsBigqueryDestinationTableUri(original["tableUri"], d, config)
	return []interface{}{transformed}
}
func flattenHealthcareDicomStoreStreamConfigsBigqueryDestinationTableUri(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandHealthcareDicomStoreName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareDicomStoreLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandHealthcareDicomStoreNotificationConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPubsubTopic, err := expandHealthcareDicomStoreNotificationConfigPubsubTopic(original["pubsub_topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubTopic); val.IsValid() && !isEmptyValue(val) {
		transformed["pubsubTopic"] = transformedPubsubTopic
	}

	return transformed, nil
}

func expandHealthcareDicomStoreNotificationConfigPubsubTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareDicomStoreStreamConfigs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedBigqueryDestination, err := expandHealthcareDicomStoreStreamConfigsBigqueryDestination(original["bigquery_destination"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedBigqueryDestination); val.IsValid() && !isEmptyValue(val) {
			transformed["bigqueryDestination"] = transformedBigqueryDestination
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandHealthcareDicomStoreStreamConfigsBigqueryDestination(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTableUri, err := expandHealthcareDicomStoreStreamConfigsBigqueryDestinationTableUri(original["table_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTableUri); val.IsValid() && !isEmptyValue(val) {
		transformed["tableUri"] = transformedTableUri
	}

	return transformed, nil
}

func expandHealthcareDicomStoreStreamConfigsBigqueryDestinationTableUri(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceHealthcareDicomStoreDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// Modify the name to be the user specified form.
	// We can't just ignore_read on `name` as the linter will
	// complain that the returned `res` is never used afterwards.
	// Some field needs to be actually set, and we chose `name`.
	res["name"] = d.Get("name").(string)
	return res, nil
}
