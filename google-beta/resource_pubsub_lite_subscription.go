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
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourcePubsubLiteSubscription() *schema.Resource {
	return &schema.Resource{
		Create: resourcePubsubLiteSubscriptionCreate,
		Read:   resourcePubsubLiteSubscriptionRead,
		Update: resourcePubsubLiteSubscriptionUpdate,
		Delete: resourcePubsubLiteSubscriptionDelete,

		Importer: &schema.ResourceImporter{
			State: resourcePubsubLiteSubscriptionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `Name of the subscription.`,
			},
			"topic": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `A reference to a Topic resource.`,
			},
			"delivery_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The settings for this subscription's message delivery.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delivery_requirement": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateEnum([]string{"DELIVER_IMMEDIATELY", "DELIVER_AFTER_STORED", "DELIVERY_REQUIREMENT_UNSPECIFIED"}),
							Description:  `When this subscription should send messages to subscribers relative to messages persistence in storage. Possible values: ["DELIVER_IMMEDIATELY", "DELIVER_AFTER_STORED", "DELIVERY_REQUIREMENT_UNSPECIFIED"]`,
						},
					},
				},
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The region of the pubsub lite topic.`,
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The zone of the pubsub lite topic.`,
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

func resourcePubsubLiteSubscriptionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	topicProp, err := expandPubsubLiteSubscriptionTopic(d.Get("topic"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("topic"); !isEmptyValue(reflect.ValueOf(topicProp)) && (ok || !reflect.DeepEqual(v, topicProp)) {
		obj["topic"] = topicProp
	}
	deliveryConfigProp, err := expandPubsubLiteSubscriptionDeliveryConfig(d.Get("delivery_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("delivery_config"); !isEmptyValue(reflect.ValueOf(deliveryConfigProp)) && (ok || !reflect.DeepEqual(v, deliveryConfigProp)) {
		obj["deliveryConfig"] = deliveryConfigProp
	}

	obj, err = resourcePubsubLiteSubscriptionEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{PubsubLiteBasePath}}projects/{{project}}/locations/{{zone}}/subscriptions?subscriptionId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Subscription: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Subscription: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Subscription: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Subscription %q: %#v", d.Id(), res)

	return resourcePubsubLiteSubscriptionRead(d, meta)
}

func resourcePubsubLiteSubscriptionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{PubsubLiteBasePath}}projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Subscription: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("PubsubLiteSubscription %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}

	if err := d.Set("topic", flattenPubsubLiteSubscriptionTopic(res["topic"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}
	if err := d.Set("delivery_config", flattenPubsubLiteSubscriptionDeliveryConfig(res["deliveryConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Subscription: %s", err)
	}

	return nil
}

func resourcePubsubLiteSubscriptionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Subscription: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	deliveryConfigProp, err := expandPubsubLiteSubscriptionDeliveryConfig(d.Get("delivery_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("delivery_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, deliveryConfigProp)) {
		obj["deliveryConfig"] = deliveryConfigProp
	}

	obj, err = resourcePubsubLiteSubscriptionEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{PubsubLiteBasePath}}projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Subscription %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("delivery_config") {
		updateMask = append(updateMask, "deliveryConfig")
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

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Subscription %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Subscription %q: %#v", d.Id(), res)
	}

	return resourcePubsubLiteSubscriptionRead(d, meta)
}

func resourcePubsubLiteSubscriptionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Subscription: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{PubsubLiteBasePath}}projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Subscription %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Subscription")
	}

	log.Printf("[DEBUG] Finished deleting Subscription %q: %#v", d.Id(), res)
	return nil
}

func resourcePubsubLiteSubscriptionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<zone>[^/]+)/subscriptions/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenPubsubLiteSubscriptionTopic(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenPubsubLiteSubscriptionDeliveryConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["delivery_requirement"] =
		flattenPubsubLiteSubscriptionDeliveryConfigDeliveryRequirement(original["deliveryRequirement"], d, config)
	return []interface{}{transformed}
}
func flattenPubsubLiteSubscriptionDeliveryConfigDeliveryRequirement(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandPubsubLiteSubscriptionTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	project, err := getProject(d, config)
	if err != nil {
		return "", err
	}

	zone, err := getZone(d, config)
	if err != nil {
		return nil, err
	}

	if zone == "" {
		return nil, fmt.Errorf("zone must be non-empty - set in resource or at provider-level")
	}

	topic := d.Get("topic").(string)

	re := regexp.MustCompile(`projects\/(.*)\/locations\/(.*)\/topics\/(.*)`)
	match := re.FindStringSubmatch(topic)
	if len(match) == 4 {
		return topic, nil
	} else {
		// If no full topic given, we expand it to a full topic on the same project
		fullTopic := fmt.Sprintf("projects/%s/locations/%s/topics/%s", project, zone, topic)
		if err := d.Set("topic", fullTopic); err != nil {
			return nil, fmt.Errorf("Error setting topic: %s", err)
		}
		return fullTopic, nil
	}
}

func expandPubsubLiteSubscriptionDeliveryConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDeliveryRequirement, err := expandPubsubLiteSubscriptionDeliveryConfigDeliveryRequirement(original["delivery_requirement"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDeliveryRequirement); val.IsValid() && !isEmptyValue(val) {
		transformed["deliveryRequirement"] = transformedDeliveryRequirement
	}

	return transformed, nil
}

func expandPubsubLiteSubscriptionDeliveryConfigDeliveryRequirement(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourcePubsubLiteSubscriptionEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)

	zone, err := getZone(d, config)
	if err != nil {
		return nil, err
	}

	if zone == "" {
		return nil, fmt.Errorf("zone must be non-empty - set in resource or at provider-level")
	}

	// API Endpoint requires region in the URL. We infer it from the zone.

	region := getRegionFromZone(zone)

	if region == "" {
		return nil, fmt.Errorf("invalid zone %q, unable to infer region from zone", zone)
	}

	return obj, nil
}
