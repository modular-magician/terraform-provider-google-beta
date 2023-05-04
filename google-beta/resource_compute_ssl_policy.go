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
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func sslPolicyCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	profile := diff.Get("profile")
	customFeaturesCount := diff.Get("custom_features.#")

	// Validate that policy configs aren't incompatible during all phases
	// CUSTOM profile demands non-zero custom_features, and other profiles (i.e., not CUSTOM) demand zero custom_features
	if diff.HasChange("profile") || diff.HasChange("custom_features") {
		if profile.(string) == "CUSTOM" {
			if customFeaturesCount.(int) == 0 {
				return fmt.Errorf("Error in SSL Policy %s: the profile is set to %s but no custom_features are set.", diff.Get("name"), profile.(string))
			}
		} else {
			if customFeaturesCount != 0 {
				return fmt.Errorf("Error in SSL Policy %s: the profile is set to %s but using custom_features requires the profile to be CUSTOM.", diff.Get("name"), profile.(string))
			}
		}
		return nil
	}
	return nil
}

func ResourceComputeSslPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeSslPolicyCreate,
		Read:   resourceComputeSslPolicyRead,
		Update: resourceComputeSslPolicyUpdate,
		Delete: resourceComputeSslPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeSslPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: sslPolicyCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"custom_features": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: `Profile specifies the set of SSL features that can be used by the
load balancer when negotiating SSL with clients. This can be one of
'COMPATIBLE', 'MODERN', 'RESTRICTED', or 'CUSTOM'. If using 'CUSTOM',
the set of SSL features to enable must be specified in the
'customFeatures' field.

See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
for which ciphers are available to use. **Note**: this argument
*must* be present when using the 'CUSTOM' profile. This argument
*must not* be present when using any other profile.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"min_tls_version": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"TLS_1_0", "TLS_1_1", "TLS_1_2", ""}),
				Description: `The minimum version of SSL protocol that can be used by the clients
to establish a connection with the load balancer. Default value: "TLS_1_0" Possible values: ["TLS_1_0", "TLS_1_1", "TLS_1_2"]`,
				Default: "TLS_1_0",
			},
			"profile": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM", ""}),
				Description: `Profile specifies the set of SSL features that can be used by the
load balancer when negotiating SSL with clients. If using 'CUSTOM',
the set of SSL features to enable must be specified in the
'customFeatures' field.

See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
for information on what cipher suites each profile provides. If
'CUSTOM' is used, the 'custom_features' attribute **must be set**. Default value: "COMPATIBLE" Possible values: ["COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM"]`,
				Default: "COMPATIBLE",
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"enabled_features": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: `The list of features enabled in the SSL policy.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Fingerprint of this resource. A hash of the contents stored in this
object. This field is used in optimistic locking.`,
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

func resourceComputeSslPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeSslPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeSslPolicyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	profileProp, err := expandComputeSslPolicyProfile(d.Get("profile"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("profile"); !isEmptyValue(reflect.ValueOf(profileProp)) && (ok || !reflect.DeepEqual(v, profileProp)) {
		obj["profile"] = profileProp
	}
	minTlsVersionProp, err := expandComputeSslPolicyMinTlsVersion(d.Get("min_tls_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("min_tls_version"); !isEmptyValue(reflect.ValueOf(minTlsVersionProp)) && (ok || !reflect.DeepEqual(v, minTlsVersionProp)) {
		obj["minTlsVersion"] = minTlsVersionProp
	}
	customFeaturesProp, err := expandComputeSslPolicyCustomFeatures(d.Get("custom_features"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("custom_features"); !isEmptyValue(reflect.ValueOf(customFeaturesProp)) && (ok || !reflect.DeepEqual(v, customFeaturesProp)) {
		obj["customFeatures"] = customFeaturesProp
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/sslPolicies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new SslPolicy: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SslPolicy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating SslPolicy: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/global/sslPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating SslPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create SslPolicy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating SslPolicy %q: %#v", d.Id(), res)

	return resourceComputeSslPolicyRead(d, meta)
}

func resourceComputeSslPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/sslPolicies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SslPolicy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeSslPolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeSslPolicyCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("description", flattenComputeSslPolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("name", flattenComputeSslPolicyName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("profile", flattenComputeSslPolicyProfile(res["profile"], d, config)); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("min_tls_version", flattenComputeSslPolicyMinTlsVersion(res["minTlsVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("enabled_features", flattenComputeSslPolicyEnabledFeatures(res["enabledFeatures"], d, config)); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("custom_features", flattenComputeSslPolicyCustomFeatures(res["customFeatures"], d, config)); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("fingerprint", flattenComputeSslPolicyFingerprint(res["fingerprint"], d, config)); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}

	return nil
}

func resourceComputeSslPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SslPolicy: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	profileProp, err := expandComputeSslPolicyProfile(d.Get("profile"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("profile"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, profileProp)) {
		obj["profile"] = profileProp
	}
	minTlsVersionProp, err := expandComputeSslPolicyMinTlsVersion(d.Get("min_tls_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("min_tls_version"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, minTlsVersionProp)) {
		obj["minTlsVersion"] = minTlsVersionProp
	}
	customFeaturesProp, err := expandComputeSslPolicyCustomFeatures(d.Get("custom_features"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("custom_features"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, customFeaturesProp)) {
		obj["customFeatures"] = customFeaturesProp
	}

	obj, err = resourceComputeSslPolicyUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/sslPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating SslPolicy %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating SslPolicy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating SslPolicy %q: %#v", d.Id(), res)
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Updating SslPolicy", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputeSslPolicyRead(d, meta)
}

func resourceComputeSslPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SslPolicy: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/sslPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting SslPolicy %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "SslPolicy")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting SslPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting SslPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeSslPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/global/sslPolicies/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/global/sslPolicies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeSslPolicyCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSslPolicyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSslPolicyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSslPolicyProfile(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSslPolicyMinTlsVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeSslPolicyEnabledFeatures(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeSslPolicyCustomFeatures(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeSslPolicyFingerprint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandComputeSslPolicyDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSslPolicyName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSslPolicyProfile(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSslPolicyMinTlsVersion(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSslPolicyCustomFeatures(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func resourceComputeSslPolicyUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// TODO(https://github.com/GoogleCloudPlatform/magic-modules/issues/184): Handle fingerprint consistently
	obj["fingerprint"] = d.Get("fingerprint")

	// TODO(https://github.com/GoogleCloudPlatform/magic-modules/issues/183): Can we generalize this
	// Send a null fields if customFeatures is empty.
	if v, ok := obj["customFeatures"]; ok && len(v.([]interface{})) == 0 {
		obj["customFeatures"] = nil
	}

	return obj, nil
}
