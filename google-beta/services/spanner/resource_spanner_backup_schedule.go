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

package spanner

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceSpannerBackupSchedule() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpannerBackupScheduleCreate,
		Read:   resourceSpannerBackupScheduleRead,
		Update: resourceSpannerBackupScheduleUpdate,
		Delete: resourceSpannerBackupScheduleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSpannerBackupScheduleImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"database": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The database to create the backup schedule on.`,
			},
			"instance": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The instance to create the database on.`,
			},
			"retention_duration": {
				Type:     schema.TypeString,
				Required: true,
				Description: `At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
A duration in seconds with up to nine fractional digits, ending with 's'. Example: '3.5s'.
You can set this to a value up to 366 days.`,
			},
			"full_backup_spec": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `The schedule creates only full backups..`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
				ExactlyOneOf: []string{"full_backup_spec", "incremental_backup_spec"},
			},
			"incremental_backup_spec": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `The schedule creates incremental backup chains.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
				ExactlyOneOf: []string{"full_backup_spec", "incremental_backup_spec"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateRegexp(`^[a-z][a-z0-9_\-]*[a-z0-9]$`),
				Description: `A unique identifier for the backup schedule, which cannot be changed after
the backup schedule is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].`,
			},
			"spec": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Defines specifications of the backup schedule.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cron_spec": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Cron style schedule specification..`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"text": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `Textual representation of the crontab. User can customize the
backup frequency and the backup version time using the cron
expression. The version time must be in UTC timzeone.
The backup will contain an externally consistent copy of the
database at the version time. Allowed frequencies are 12 hour, 1 day,
1 week and 1 month. Examples of valid cron specifications:
  0 2/12 * * * : every 12 hours at (2, 14) hours past midnight in UTC.
  0 2,14 * * * : every 12 hours at (2,14) hours past midnight in UTC.
  0 2 * * *    : once a day at 2 past midnight in UTC.
  0 2 * * 0    : once a week every Sunday at 2 past midnight in UTC.
  0 2 8 * *    : once a month on 8th day at 2 past midnight in UTC.`,
									},
								},
							},
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

func resourceSpannerBackupScheduleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandSpannerBackupScheduleName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	retentionDurationProp, err := expandSpannerBackupScheduleRetentionDuration(d.Get("retention_duration"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retention_duration"); !tpgresource.IsEmptyValue(reflect.ValueOf(retentionDurationProp)) && (ok || !reflect.DeepEqual(v, retentionDurationProp)) {
		obj["retentionDuration"] = retentionDurationProp
	}
	specProp, err := expandSpannerBackupScheduleSpec(d.Get("spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("spec"); ok || !reflect.DeepEqual(v, specProp) {
		obj["spec"] = specProp
	}
	fullBackupSpecProp, err := expandSpannerBackupScheduleFullBackupSpec(d.Get("full_backup_spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("full_backup_spec"); ok || !reflect.DeepEqual(v, fullBackupSpecProp) {
		obj["fullBackupSpec"] = fullBackupSpecProp
	}
	incrementalBackupSpecProp, err := expandSpannerBackupScheduleIncrementalBackupSpec(d.Get("incremental_backup_spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("incremental_backup_spec"); ok || !reflect.DeepEqual(v, incrementalBackupSpecProp) {
		obj["incrementalBackupSpec"] = incrementalBackupSpecProp
	}

	obj, err = resourceSpannerBackupScheduleEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instances/{{instance}}/databases/{{database}}/backupSchedules?backup_schedule_id={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BackupSchedule: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackupSchedule: %s", err)
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
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating BackupSchedule: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/instances/{{instance}}/databases/{{database}}/backupSchedules/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating BackupSchedule %q: %#v", d.Id(), res)

	return resourceSpannerBackupScheduleRead(d, meta)
}

func resourceSpannerBackupScheduleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instances/{{instance}}/databases/{{database}}/backupSchedules/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackupSchedule: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SpannerBackupSchedule %q", d.Id()))
	}

	res, err = resourceSpannerBackupScheduleDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing SpannerBackupSchedule because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading BackupSchedule: %s", err)
	}

	if err := d.Set("name", flattenSpannerBackupScheduleName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupSchedule: %s", err)
	}
	if err := d.Set("retention_duration", flattenSpannerBackupScheduleRetentionDuration(res["retentionDuration"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupSchedule: %s", err)
	}
	if err := d.Set("spec", flattenSpannerBackupScheduleSpec(res["spec"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupSchedule: %s", err)
	}
	if err := d.Set("full_backup_spec", flattenSpannerBackupScheduleFullBackupSpec(res["fullBackupSpec"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupSchedule: %s", err)
	}
	if err := d.Set("incremental_backup_spec", flattenSpannerBackupScheduleIncrementalBackupSpec(res["incrementalBackupSpec"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupSchedule: %s", err)
	}

	return nil
}

func resourceSpannerBackupScheduleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackupSchedule: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	retentionDurationProp, err := expandSpannerBackupScheduleRetentionDuration(d.Get("retention_duration"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retention_duration"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, retentionDurationProp)) {
		obj["retentionDuration"] = retentionDurationProp
	}
	specProp, err := expandSpannerBackupScheduleSpec(d.Get("spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("spec"); ok || !reflect.DeepEqual(v, specProp) {
		obj["spec"] = specProp
	}

	obj, err = resourceSpannerBackupScheduleEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instances/{{instance}}/databases/{{database}}/backupSchedules/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating BackupSchedule %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("retention_duration") {
		updateMask = append(updateMask, "retentionDuration")
	}

	if d.HasChange("spec") {
		updateMask = append(updateMask, "spec")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	// The generated code sets the wrong masks for the following fields.
	newUpdateMask := []string{}
	if d.HasChange("spec.0.cron_spec.0.text") {
		newUpdateMask = append(newUpdateMask, "spec.cron_spec.text")
	}
	// Pull out any other set fields from the generated mask.
	for _, mask := range updateMask {
		if mask == "spec" {
			continue
		}
		newUpdateMask = append(newUpdateMask, mask)
	}
	// Overwrite the previously set mask.
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(newUpdateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating BackupSchedule %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating BackupSchedule %q: %#v", d.Id(), res)
		}

	}

	return resourceSpannerBackupScheduleRead(d, meta)
}

func resourceSpannerBackupScheduleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackupSchedule: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{SpannerBasePath}}projects/{{project}}/instances/{{instance}}/databases/{{database}}/backupSchedules/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting BackupSchedule %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "BackupSchedule")
	}

	log.Printf("[DEBUG] Finished deleting BackupSchedule %q: %#v", d.Id(), res)
	return nil
}

func resourceSpannerBackupScheduleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/instances/(?P<instance>[^/]+)/databases/(?P<database>[^/]+)/backupSchedules/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<instance>[^/]+)/(?P<database>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<instance>[^/]+)/(?P<database>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/instances/{{instance}}/databases/{{database}}/backupSchedules/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenSpannerBackupScheduleName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenSpannerBackupScheduleRetentionDuration(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSpannerBackupScheduleSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["cron_spec"] =
		flattenSpannerBackupScheduleSpecCronSpec(original["cronSpec"], d, config)
	return []interface{}{transformed}
}
func flattenSpannerBackupScheduleSpecCronSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["text"] =
		flattenSpannerBackupScheduleSpecCronSpecText(original["text"], d, config)
	return []interface{}{transformed}
}
func flattenSpannerBackupScheduleSpecCronSpecText(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSpannerBackupScheduleFullBackupSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	transformed := make(map[string]interface{})
	return []interface{}{transformed}
}

func flattenSpannerBackupScheduleIncrementalBackupSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	transformed := make(map[string]interface{})
	return []interface{}{transformed}
}

func expandSpannerBackupScheduleName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerBackupScheduleRetentionDuration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerBackupScheduleSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCronSpec, err := expandSpannerBackupScheduleSpecCronSpec(original["cron_spec"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCronSpec); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cronSpec"] = transformedCronSpec
	}

	return transformed, nil
}

func expandSpannerBackupScheduleSpecCronSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedText, err := expandSpannerBackupScheduleSpecCronSpecText(original["text"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedText); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["text"] = transformedText
	}

	return transformed, nil
}

func expandSpannerBackupScheduleSpecCronSpecText(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerBackupScheduleFullBackupSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	transformed := make(map[string]interface{})

	return transformed, nil
}

func expandSpannerBackupScheduleIncrementalBackupSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	transformed := make(map[string]interface{})

	return transformed, nil
}

func resourceSpannerBackupScheduleEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	obj["name"] = d.Get("name").(string)
	if obj["name"] == nil || obj["name"] == "" {
		if err := d.Set("name", id.PrefixedUniqueId("tfgen-spanid-")[:30]); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
	}
	config := meta.(*transport_tpg.Config)
	var err error
	obj["name"], err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/instances/{{instance}}/databases/{{database}}/backupSchedules/{{name}}")
	if err != nil {
		return obj, err
	}
	delete(obj, "instance")
	delete(obj, "database")
	return obj, nil
}

func resourceSpannerBackupScheduleDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	d.SetId(res["name"].(string))
	if err := tpgresource.ParseImportId([]string{"projects/(?P<project>[^/]+)/instances/(?P<instance>[^/]+)/databases/(?P<database>[^/]+)/backupSchedules/(?P<name>[^/]+)"}, d, config); err != nil {
		return nil, err
	}
	res["project"] = d.Get("project").(string)
	res["instance"] = d.Get("instance").(string)
	res["database"] = d.Get("database").(string)
	res["name"] = d.Get("name").(string)
	id, err := tpgresource.ReplaceVars(d, config, "{{instance}}/{{database}}/{{name}}")
	if err != nil {
		return nil, err
	}
	d.SetId(id)
	return res, nil
}
