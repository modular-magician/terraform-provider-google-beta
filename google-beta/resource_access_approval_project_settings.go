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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAccessApprovalProjectSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessApprovalProjectSettingsCreate,
		Read:   resourceAccessApprovalProjectSettingsRead,
		Update: resourceAccessApprovalProjectSettingsUpdate,
		Delete: resourceAccessApprovalProjectSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessApprovalProjectSettingsImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"enrolled_services": {
				Type:     schema.TypeSet,
				Required: true,
				Description: `A list of Google Cloud Services for which the given resource has Access Approval enrolled.
Access requests for the resource given by name against any of these services contained here will be required
to have explicit approval. Enrollment can only be done on an all or nothing basis.

A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.`,
				Elem: accessapprovalProjectSettingsEnrolledServicesSchema(),
				// Default schema.HashSchema is used.
			},
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `ID of the project of the access approval settings.`,
			},
			"notification_emails": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				Description: `A list of email addresses to which notifications relating to approval requests should be sent.
Notifications relating to a resource will be sent to all emails in the settings of ancestor
resources of that resource. A maximum of 50 email addresses are allowed.`,
				MaxItems: 50,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"enrolled_ancestor": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors of the Project.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the settings. Format is "projects/{project_id/accessApprovalSettings"`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func accessapprovalProjectSettingsEnrolledServicesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_product": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The product for which Access Approval will be enrolled. Allowed values are listed (case-sensitive):
  all
  appengine.googleapis.com
  bigquery.googleapis.com
  bigtable.googleapis.com
  cloudkms.googleapis.com
  compute.googleapis.com
  dataflow.googleapis.com
  iam.googleapis.com
  pubsub.googleapis.com
  storage.googleapis.com`,
			},
			"enrollment_level": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"BLOCK_ALL", ""}, false),
				Description:  `The enrollment level of the service. Default value: "BLOCK_ALL" Possible values: ["BLOCK_ALL"]`,
				Default:      "BLOCK_ALL",
			},
		},
	}
}

func resourceAccessApprovalProjectSettingsCreate(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	obj := make(map[string]interface{})
	notificationEmailsProp, err := expandAccessApprovalProjectSettingsNotificationEmails(d.Get("notification_emails"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_emails"); !isEmptyValue(reflect.ValueOf(notificationEmailsProp)) && (ok || !reflect.DeepEqual(v, notificationEmailsProp)) {
		obj["notificationEmails"] = notificationEmailsProp
	}
	enrolledServicesProp, err := expandAccessApprovalProjectSettingsEnrolledServices(d.Get("enrolled_services"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enrolled_services"); !isEmptyValue(reflect.ValueOf(enrolledServicesProp)) && (ok || !reflect.DeepEqual(v, enrolledServicesProp)) {
		obj["enrolledServices"] = enrolledServicesProp
	}

	url, err := replaceVars(d, config, "{{AccessApprovalBasePath}}projects/{{project}}/accessApprovalSettings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ProjectSettings: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	updateMask := []string{}

	if d.HasChange("notification_emails") {
		updateMask = append(updateMask, "notificationEmails")
	}

	if d.HasChange("enrolled_services") {
		updateMask = append(updateMask, "enrolledServices")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ProjectSettings: %s", err)
	}
	if err := d.Set("name", flattenAccessApprovalProjectSettingsName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/accessApprovalSettings")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ProjectSettings %q: %#v", d.Id(), res)

	return resourceAccessApprovalProjectSettingsRead(d, meta)
}

func resourceAccessApprovalProjectSettingsRead(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	url, err := replaceVars(d, config, "{{AccessApprovalBasePath}}projects/{{project}}/accessApprovalSettings")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AccessApprovalProjectSettings %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ProjectSettings: %s", err)
	}

	if err := d.Set("name", flattenAccessApprovalProjectSettingsName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectSettings: %s", err)
	}
	if err := d.Set("notification_emails", flattenAccessApprovalProjectSettingsNotificationEmails(res["notificationEmails"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectSettings: %s", err)
	}
	if err := d.Set("enrolled_services", flattenAccessApprovalProjectSettingsEnrolledServices(res["enrolledServices"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectSettings: %s", err)
	}
	if err := d.Set("enrolled_ancestor", flattenAccessApprovalProjectSettingsEnrolledAncestor(res["enrolledAncestor"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectSettings: %s", err)
	}

	return nil
}

func resourceAccessApprovalProjectSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	obj := make(map[string]interface{})
	notificationEmailsProp, err := expandAccessApprovalProjectSettingsNotificationEmails(d.Get("notification_emails"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_emails"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, notificationEmailsProp)) {
		obj["notificationEmails"] = notificationEmailsProp
	}
	enrolledServicesProp, err := expandAccessApprovalProjectSettingsEnrolledServices(d.Get("enrolled_services"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enrolled_services"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enrolledServicesProp)) {
		obj["enrolledServices"] = enrolledServicesProp
	}

	url, err := replaceVars(d, config, "{{AccessApprovalBasePath}}projects/{{project}}/accessApprovalSettings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ProjectSettings %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("notification_emails") {
		updateMask = append(updateMask, "notificationEmails")
	}

	if d.HasChange("enrolled_services") {
		updateMask = append(updateMask, "enrolledServices")
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

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating ProjectSettings %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ProjectSettings %q: %#v", d.Id(), res)
	}

	return resourceAccessApprovalProjectSettingsRead(d, meta)
}

func resourceAccessApprovalProjectSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	config := meta.(*Config)

	obj := make(map[string]interface{})
	obj["notificationEmails"] = []string{}
	obj["enrolledServices"] = []string{}

	url, err := replaceVars(d, config, "{{AccessApprovalBasePath}}projects/{{project}}/accessApprovalSettings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Emptying ProjectSettings %q: %#v", d.Id(), obj)
	updateMask := []string{}

	updateMask = append(updateMask, "notificationEmails")
	updateMask = append(updateMask, "enrolledServices")

	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	res, err := sendRequestWithTimeout(config, "PATCH", "", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error emptying ProjectSettings %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished emptying ProjectSettings %q: %#v", d.Id(), res)
	}

	return nil
}

func resourceAccessApprovalProjectSettingsImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/accessApprovalSettings",
		"(?P<project>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/accessApprovalSettings")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAccessApprovalProjectSettingsName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsNotificationEmails(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenAccessApprovalProjectSettingsEnrolledServices(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(accessapprovalProjectSettingsEnrolledServicesSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"cloud_product":    flattenAccessApprovalProjectSettingsEnrolledServicesCloudProduct(original["cloudProduct"], d, config),
			"enrollment_level": flattenAccessApprovalProjectSettingsEnrolledServicesEnrollmentLevel(original["enrollmentLevel"], d, config),
		})
	}
	return transformed
}
func flattenAccessApprovalProjectSettingsEnrolledServicesCloudProduct(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsEnrolledServicesEnrollmentLevel(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessApprovalProjectSettingsEnrolledAncestor(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandAccessApprovalProjectSettingsNotificationEmails(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandAccessApprovalProjectSettingsEnrolledServices(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedCloudProduct, err := expandAccessApprovalProjectSettingsEnrolledServicesCloudProduct(original["cloud_product"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCloudProduct); val.IsValid() && !isEmptyValue(val) {
			transformed["cloudProduct"] = transformedCloudProduct
		}

		transformedEnrollmentLevel, err := expandAccessApprovalProjectSettingsEnrolledServicesEnrollmentLevel(original["enrollment_level"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEnrollmentLevel); val.IsValid() && !isEmptyValue(val) {
			transformed["enrollmentLevel"] = transformedEnrollmentLevel
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAccessApprovalProjectSettingsEnrolledServicesCloudProduct(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessApprovalProjectSettingsEnrolledServicesEnrollmentLevel(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
