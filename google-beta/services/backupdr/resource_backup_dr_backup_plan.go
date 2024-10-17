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

package backupdr

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceBackupDRBackupPlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceBackupDRBackupPlanCreate,
		Read:   resourceBackupDRBackupPlanRead,
		Delete: resourceBackupDRBackupPlanDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBackupDRBackupPlanImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"backup_rules": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: `The backup rules for this 'BackupPlan'. There must be at least one 'BackupRule' message.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_retention_days": {
							Type:        schema.TypeInt,
							Required:    true,
							ForceNew:    true,
							Description: `Configures the duration for which backup data will be kept. The value should be greater than or equal to minimum enforced retention of the backup vault.`,
						},
						"rule_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `The unique ID of this 'BackupRule'. The 'rule_id' is unique per 'BackupPlan'.`,
						},
						"standard_schedule": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Description: `StandardSchedule defines a schedule that runs within the confines of a defined window of days.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"recurrence_type": {
										Type:         schema.TypeString,
										Required:     true,
										ForceNew:     true,
										ValidateFunc: verify.ValidateEnum([]string{"HOURLY", "DAILY", "WEEKLY", "MONTHLY", "YEARLY"}),
										Description:  `RecurrenceType enumerates the applicable periodicity for the schedule. Possible values: ["HOURLY", "DAILY", "WEEKLY", "MONTHLY", "YEARLY"]`,
									},
									"time_zone": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: `The time zone to be used when interpreting the schedule.`,
									},
									"backup_window": {
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Description: `A BackupWindow defines the window of the day during which backup jobs will run. Jobs are queued at the beginning of the window and will be marked as
'NOT_RUN' if they do not start by the end of the window.`,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"start_hour_of_day": {
													Type:        schema.TypeInt,
													Required:    true,
													ForceNew:    true,
													Description: `The hour of the day (0-23) when the window starts, for example, if the value of the start hour of the day is 6, that means the backup window starts at 6:00.`,
												},
												"end_hour_of_day": {
													Type:     schema.TypeInt,
													Optional: true,
													ForceNew: true,
													Description: `The hour of the day (1-24) when the window ends, for example, if the value of end hour of the day is 10, that means the backup window end time is 10:00.
The end hour of the day should be greater than the start`,
												},
											},
										},
									},
									"days_of_month": {
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										Description: `Specifies days of months like 1, 5, or 14 on which jobs will run.`,
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"days_of_week": {
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										Description: `Specifies days of week like MONDAY or TUESDAY, on which jobs will run. This is required for 'recurrence_type', 'WEEKLY' and is not applicable otherwise. Possible values: ["DAY_OF_WEEK_UNSPECIFIED", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY"]`,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: verify.ValidateEnum([]string{"DAY_OF_WEEK_UNSPECIFIED", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY"}),
										},
									},
									"hourly_frequency": {
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Description: `Specifies frequency for hourly backups. An hourly frequency of 2 means jobs will run every 2 hours from start time till end time defined.
This is required for 'recurrence_type', 'HOURLY' and is not applicable otherwise.`,
									},
									"months": {
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										Description: `Specifies values of months Possible values: ["MONTH_UNSPECIFIED", "JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE", "JULY", "AUGUST", "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER"]`,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: verify.ValidateEnum([]string{"MONTH_UNSPECIFIED", "JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE", "JULY", "AUGUST", "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER"}),
										},
									},
									"week_day_of_month": {
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										Description: `Specifies a week day of the month like FIRST SUNDAY or LAST MONDAY, on which jobs will run.`,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"day_of_week": {
													Type:         schema.TypeString,
													Required:     true,
													ForceNew:     true,
													ValidateFunc: verify.ValidateEnum([]string{"DAY_OF_WEEK_UNSPECIFIED", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"}),
													Description:  `Specifies the day of the week. Possible values: ["DAY_OF_WEEK_UNSPECIFIED", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]`,
												},
												"week_of_month": {
													Type:         schema.TypeString,
													Required:     true,
													ForceNew:     true,
													ValidateFunc: verify.ValidateEnum([]string{"WEEK_OF_MONTH_UNSPECIFIED", "FIRST", "SECOND", "THIRD", "FOURTH", "LAST"}),
													Description:  `WeekOfMonth enumerates possible weeks in the month, e.g. the first, third, or last week of the month. Possible values: ["WEEK_OF_MONTH_UNSPECIFIED", "FIRST", "SECOND", "THIRD", "FOURTH", "LAST"]`,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"backup_vault": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Backup vault where the backups gets stored using this Backup plan.`,
			},
			"backup_plan_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the backup plan`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location for the backup plan`,
			},
			"resource_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource type to which the 'BackupPlan' will be applied. Examples include, "compute.googleapis.com/Instance" and "storage.googleapis.com/Bucket".`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The description allows for additional details about 'BackupPlan' and its use cases to be provided.`,
			},
			"backup_vault_service_account": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The Google Cloud Platform Service Account to be used by the BackupVault for taking backups.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `When the 'BackupPlan' was created.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of backup plan resource created`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `When the 'BackupPlan' was last updated.`,
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

func resourceBackupDRBackupPlanCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandBackupDRBackupPlanDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	backupVaultProp, err := expandBackupDRBackupPlanBackupVault(d.Get("backup_vault"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backup_vault"); !tpgresource.IsEmptyValue(reflect.ValueOf(backupVaultProp)) && (ok || !reflect.DeepEqual(v, backupVaultProp)) {
		obj["backupVault"] = backupVaultProp
	}
	resourceTypeProp, err := expandBackupDRBackupPlanResourceType(d.Get("resource_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("resource_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourceTypeProp)) && (ok || !reflect.DeepEqual(v, resourceTypeProp)) {
		obj["resourceType"] = resourceTypeProp
	}
	backupRulesProp, err := expandBackupDRBackupPlanBackupRules(d.Get("backup_rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backup_rules"); !tpgresource.IsEmptyValue(reflect.ValueOf(backupRulesProp)) && (ok || !reflect.DeepEqual(v, backupRulesProp)) {
		obj["backupRules"] = backupRulesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BackupDRBasePath}}projects/{{project}}/locations/{{location}}/backupPlans/?backup_plan_id={{backup_plan_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BackupPlan: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackupPlan: %s", err)
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
		return fmt.Errorf("Error creating BackupPlan: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backupPlans/{{backup_plan_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = BackupDROperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating BackupPlan", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create BackupPlan: %s", err)
	}

	if err := d.Set("name", flattenBackupDRBackupPlanName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backupPlans/{{backup_plan_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating BackupPlan %q: %#v", d.Id(), res)

	return resourceBackupDRBackupPlanRead(d, meta)
}

func resourceBackupDRBackupPlanRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BackupDRBasePath}}projects/{{project}}/locations/{{location}}/backupPlans/{{backup_plan_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackupPlan: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("BackupDRBackupPlan %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading BackupPlan: %s", err)
	}

	if err := d.Set("name", flattenBackupDRBackupPlanName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupPlan: %s", err)
	}
	if err := d.Set("description", flattenBackupDRBackupPlanDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupPlan: %s", err)
	}
	if err := d.Set("backup_vault", flattenBackupDRBackupPlanBackupVault(res["backupVault"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupPlan: %s", err)
	}
	if err := d.Set("backup_vault_service_account", flattenBackupDRBackupPlanBackupVaultServiceAccount(res["backupVaultServiceAccount"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupPlan: %s", err)
	}
	if err := d.Set("resource_type", flattenBackupDRBackupPlanResourceType(res["resourceType"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupPlan: %s", err)
	}
	if err := d.Set("create_time", flattenBackupDRBackupPlanCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupPlan: %s", err)
	}
	if err := d.Set("update_time", flattenBackupDRBackupPlanUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupPlan: %s", err)
	}
	if err := d.Set("backup_rules", flattenBackupDRBackupPlanBackupRules(res["backupRules"], d, config)); err != nil {
		return fmt.Errorf("Error reading BackupPlan: %s", err)
	}

	return nil
}

func resourceBackupDRBackupPlanDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BackupPlan: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{BackupDRBasePath}}projects/{{project}}/locations/{{location}}/backupPlans/{{backup_plan_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting BackupPlan %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "BackupPlan")
	}

	err = BackupDROperationWaitTime(
		config, res, project, "Deleting BackupPlan", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting BackupPlan %q: %#v", d.Id(), res)
	return nil
}

func resourceBackupDRBackupPlanImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/backupPlans/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backupPlans/{{backup_plan_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBackupDRBackupPlanName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanBackupVault(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanBackupVaultServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanResourceType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanBackupRules(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"rule_id":               flattenBackupDRBackupPlanBackupRulesRuleId(original["ruleId"], d, config),
			"backup_retention_days": flattenBackupDRBackupPlanBackupRulesBackupRetentionDays(original["backupRetentionDays"], d, config),
			"standard_schedule":     flattenBackupDRBackupPlanBackupRulesStandardSchedule(original["standardSchedule"], d, config),
		})
	}
	return transformed
}
func flattenBackupDRBackupPlanBackupRulesRuleId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanBackupRulesBackupRetentionDays(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenBackupDRBackupPlanBackupRulesStandardSchedule(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["recurrence_type"] =
		flattenBackupDRBackupPlanBackupRulesStandardScheduleRecurrenceType(original["recurrenceType"], d, config)
	transformed["hourly_frequency"] =
		flattenBackupDRBackupPlanBackupRulesStandardScheduleHourlyFrequency(original["hourlyFrequency"], d, config)
	transformed["days_of_week"] =
		flattenBackupDRBackupPlanBackupRulesStandardScheduleDaysOfWeek(original["daysOfWeek"], d, config)
	transformed["days_of_month"] =
		flattenBackupDRBackupPlanBackupRulesStandardScheduleDaysOfMonth(original["daysOfMonth"], d, config)
	transformed["week_day_of_month"] =
		flattenBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonth(original["weekDayOfMonth"], d, config)
	transformed["months"] =
		flattenBackupDRBackupPlanBackupRulesStandardScheduleMonths(original["months"], d, config)
	transformed["time_zone"] =
		flattenBackupDRBackupPlanBackupRulesStandardScheduleTimeZone(original["timeZone"], d, config)
	transformed["backup_window"] =
		flattenBackupDRBackupPlanBackupRulesStandardScheduleBackupWindow(original["backupWindow"], d, config)
	return []interface{}{transformed}
}
func flattenBackupDRBackupPlanBackupRulesStandardScheduleRecurrenceType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanBackupRulesStandardScheduleHourlyFrequency(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenBackupDRBackupPlanBackupRulesStandardScheduleDaysOfWeek(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanBackupRulesStandardScheduleDaysOfMonth(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonth(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["week_of_month"] =
		flattenBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonthWeekOfMonth(original["weekOfMonth"], d, config)
	transformed["day_of_week"] =
		flattenBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonthDayOfWeek(original["dayOfWeek"], d, config)
	return []interface{}{transformed}
}
func flattenBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonthWeekOfMonth(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonthDayOfWeek(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanBackupRulesStandardScheduleMonths(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanBackupRulesStandardScheduleTimeZone(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRBackupPlanBackupRulesStandardScheduleBackupWindow(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["start_hour_of_day"] =
		flattenBackupDRBackupPlanBackupRulesStandardScheduleBackupWindowStartHourOfDay(original["startHourOfDay"], d, config)
	transformed["end_hour_of_day"] =
		flattenBackupDRBackupPlanBackupRulesStandardScheduleBackupWindowEndHourOfDay(original["endHourOfDay"], d, config)
	return []interface{}{transformed}
}
func flattenBackupDRBackupPlanBackupRulesStandardScheduleBackupWindowStartHourOfDay(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenBackupDRBackupPlanBackupRulesStandardScheduleBackupWindowEndHourOfDay(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func expandBackupDRBackupPlanDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupVault(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanResourceType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedRuleId, err := expandBackupDRBackupPlanBackupRulesRuleId(original["rule_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRuleId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ruleId"] = transformedRuleId
		}

		transformedBackupRetentionDays, err := expandBackupDRBackupPlanBackupRulesBackupRetentionDays(original["backup_retention_days"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedBackupRetentionDays); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["backupRetentionDays"] = transformedBackupRetentionDays
		}

		transformedStandardSchedule, err := expandBackupDRBackupPlanBackupRulesStandardSchedule(original["standard_schedule"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStandardSchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["standardSchedule"] = transformedStandardSchedule
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBackupDRBackupPlanBackupRulesRuleId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesBackupRetentionDays(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRecurrenceType, err := expandBackupDRBackupPlanBackupRulesStandardScheduleRecurrenceType(original["recurrence_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecurrenceType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recurrenceType"] = transformedRecurrenceType
	}

	transformedHourlyFrequency, err := expandBackupDRBackupPlanBackupRulesStandardScheduleHourlyFrequency(original["hourly_frequency"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHourlyFrequency); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hourlyFrequency"] = transformedHourlyFrequency
	}

	transformedDaysOfWeek, err := expandBackupDRBackupPlanBackupRulesStandardScheduleDaysOfWeek(original["days_of_week"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDaysOfWeek); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["daysOfWeek"] = transformedDaysOfWeek
	}

	transformedDaysOfMonth, err := expandBackupDRBackupPlanBackupRulesStandardScheduleDaysOfMonth(original["days_of_month"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDaysOfMonth); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["daysOfMonth"] = transformedDaysOfMonth
	}

	transformedWeekDayOfMonth, err := expandBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonth(original["week_day_of_month"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeekDayOfMonth); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["weekDayOfMonth"] = transformedWeekDayOfMonth
	}

	transformedMonths, err := expandBackupDRBackupPlanBackupRulesStandardScheduleMonths(original["months"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMonths); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["months"] = transformedMonths
	}

	transformedTimeZone, err := expandBackupDRBackupPlanBackupRulesStandardScheduleTimeZone(original["time_zone"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeZone); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["timeZone"] = transformedTimeZone
	}

	transformedBackupWindow, err := expandBackupDRBackupPlanBackupRulesStandardScheduleBackupWindow(original["backup_window"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBackupWindow); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["backupWindow"] = transformedBackupWindow
	}

	return transformed, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleRecurrenceType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleHourlyFrequency(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleDaysOfWeek(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleDaysOfMonth(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonth(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedWeekOfMonth, err := expandBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonthWeekOfMonth(original["week_of_month"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeekOfMonth); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["weekOfMonth"] = transformedWeekOfMonth
	}

	transformedDayOfWeek, err := expandBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonthDayOfWeek(original["day_of_week"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDayOfWeek); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dayOfWeek"] = transformedDayOfWeek
	}

	return transformed, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonthWeekOfMonth(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonthDayOfWeek(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleMonths(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleTimeZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleBackupWindow(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStartHourOfDay, err := expandBackupDRBackupPlanBackupRulesStandardScheduleBackupWindowStartHourOfDay(original["start_hour_of_day"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartHourOfDay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startHourOfDay"] = transformedStartHourOfDay
	}

	transformedEndHourOfDay, err := expandBackupDRBackupPlanBackupRulesStandardScheduleBackupWindowEndHourOfDay(original["end_hour_of_day"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEndHourOfDay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["endHourOfDay"] = transformedEndHourOfDay
	}

	return transformed, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleBackupWindowStartHourOfDay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleBackupWindowEndHourOfDay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
