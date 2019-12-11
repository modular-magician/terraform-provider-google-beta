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
	"strconv"
	"time"
)

func resourceComputeResourcePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeResourcePolicyCreate,
		Read:   resourceComputeResourcePolicyRead,
		Delete: resourceComputeResourcePolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeResourcePolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The name of the resource, provided by the client when initially creating
the resource. The resource name must be 1-63 characters long, and comply
with RFC1035. Specifically, the name must be 1-63 characters long and
match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])'? which means the
first character must be a lowercase letter, and all following characters
must be a dash, lowercase letter, or digit, except the last character,
which cannot be a dash.`,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `Region where resource policy resides.`,
			},
			"snapshot_schedule_policy": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Policy for creating snapshots of persistent disks.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"schedule": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Description: `Contains one of an 'hourlySchedule', 'dailySchedule', or 'weeklySchedule'.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"daily_schedule": {
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										Description: `The policy will execute every nth day at the specified time.`,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"days_in_cycle": {
													Type:        schema.TypeInt,
													Required:    true,
													ForceNew:    true,
													Description: `The number of days between snapshots.`,
												},
												"start_time": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
													Description: `This must be in UTC format that resolves to one of
00:00, 04:00, 08:00, 12:00, 16:00, or 20:00. For example,
both 13:00-5 and 08:00 are valid.`,
												},
											},
										},
										ExactlyOneOf: []string{"snapshot_schedule_policy.0.schedule.0.hourly_schedule", "snapshot_schedule_policy.0.schedule.0.daily_schedule", "snapshot_schedule_policy.0.schedule.0.weekly_schedule"},
									},
									"hourly_schedule": {
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										Description: `The policy will execute every nth hour starting at the specified time.`,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hours_in_cycle": {
													Type:        schema.TypeInt,
													Required:    true,
													ForceNew:    true,
													Description: `The number of hours between snapshots.`,
												},
												"start_time": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
													Description: `Time within the window to start the operations.
It must be in format "HH:MM",
where HH : [00-23] and MM : [00-00] GMT.`,
												},
											},
										},
										ExactlyOneOf: []string{"snapshot_schedule_policy.0.schedule.0.hourly_schedule", "snapshot_schedule_policy.0.schedule.0.daily_schedule", "snapshot_schedule_policy.0.schedule.0.weekly_schedule"},
									},
									"weekly_schedule": {
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										Description: `Allows specifying a snapshot time for each day of the week.`,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"day_of_weeks": {
													Type:        schema.TypeSet,
													Required:    true,
													ForceNew:    true,
													Description: `May contain up to seven (one for each day of the week) snapshot times.`,
													MinItems:    1,
													MaxItems:    7,
													Elem:        computeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksSchema(),
													// Default schema.HashSchema is used.
												},
											},
										},
										ExactlyOneOf: []string{"snapshot_schedule_policy.0.schedule.0.hourly_schedule", "snapshot_schedule_policy.0.schedule.0.daily_schedule", "snapshot_schedule_policy.0.schedule.0.weekly_schedule"},
									},
								},
							},
						},
						"retention_policy": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `Retention policy applied to snapshots created by this resource policy.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_retention_days": {
										Type:        schema.TypeInt,
										Required:    true,
										ForceNew:    true,
										Description: `Maximum age of the snapshot that is allowed to be kept.`,
									},
									"on_source_disk_delete": {
										Type:         schema.TypeString,
										Optional:     true,
										ForceNew:     true,
										ValidateFunc: validation.StringInSlice([]string{"KEEP_AUTO_SNAPSHOTS", "APPLY_RETENTION_POLICY", ""}, false),
										Description: `Specifies the behavior to apply to scheduled snapshots when
the source disk is deleted.
Valid options are KEEP_AUTO_SNAPSHOTS and APPLY_RETENTION_POLICY`,
										Default: "KEEP_AUTO_SNAPSHOTS",
									},
								},
							},
						},
						"snapshot_properties": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `Properties with which the snapshots are created, such as labels.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"guest_flush": {
										Type:         schema.TypeBool,
										Optional:     true,
										ForceNew:     true,
										Description:  `Whether to perform a 'guest aware' snapshot.`,
										AtLeastOneOf: []string{"snapshot_schedule_policy.0.snapshot_properties.0.labels", "snapshot_schedule_policy.0.snapshot_properties.0.storage_locations", "snapshot_schedule_policy.0.snapshot_properties.0.guest_flush"},
									},
									"labels": {
										Type:         schema.TypeMap,
										Optional:     true,
										ForceNew:     true,
										Description:  `A set of key-value pairs.`,
										Elem:         &schema.Schema{Type: schema.TypeString},
										AtLeastOneOf: []string{"snapshot_schedule_policy.0.snapshot_properties.0.labels", "snapshot_schedule_policy.0.snapshot_properties.0.storage_locations", "snapshot_schedule_policy.0.snapshot_properties.0.guest_flush"},
									},
									"storage_locations": {
										Type:        schema.TypeSet,
										Optional:    true,
										ForceNew:    true,
										Description: `GCS bucket location in which to store the snapshot (regional or multi-regional).`,
										MaxItems:    1,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Set:          schema.HashString,
										AtLeastOneOf: []string{"snapshot_schedule_policy.0.snapshot_properties.0.labels", "snapshot_schedule_policy.0.snapshot_properties.0.storage_locations", "snapshot_schedule_policy.0.snapshot_properties.0.guest_flush"},
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
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func computeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"day": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"}, false),
				Description:  `The day of the week to create the snapshot. e.g. MONDAY`,
			},
			"start_time": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Time within the window to start the operations.
It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.`,
			},
		},
	}
}

func resourceComputeResourcePolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandComputeResourcePolicyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	snapshotSchedulePolicyProp, err := expandComputeResourcePolicySnapshotSchedulePolicy(d.Get("snapshot_schedule_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("snapshot_schedule_policy"); !isEmptyValue(reflect.ValueOf(snapshotSchedulePolicyProp)) && (ok || !reflect.DeepEqual(v, snapshotSchedulePolicyProp)) {
		obj["snapshotSchedulePolicy"] = snapshotSchedulePolicyProp
	}
	regionProp, err := expandComputeResourcePolicyRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/resourcePolicies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ResourcePolicy: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ResourcePolicy: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating ResourcePolicy",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ResourcePolicy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ResourcePolicy %q: %#v", d.Id(), res)

	return resourceComputeResourcePolicyRead(d, meta)
}

func resourceComputeResourcePolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeResourcePolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ResourcePolicy: %s", err)
	}

	if err := d.Set("name", flattenComputeResourcePolicyName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading ResourcePolicy: %s", err)
	}
	if err := d.Set("snapshot_schedule_policy", flattenComputeResourcePolicySnapshotSchedulePolicy(res["snapshotSchedulePolicy"], d)); err != nil {
		return fmt.Errorf("Error reading ResourcePolicy: %s", err)
	}
	if err := d.Set("region", flattenComputeResourcePolicyRegion(res["region"], d)); err != nil {
		return fmt.Errorf("Error reading ResourcePolicy: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading ResourcePolicy: %s", err)
	}

	return nil
}

func resourceComputeResourcePolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ResourcePolicy %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ResourcePolicy")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting ResourcePolicy",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ResourcePolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeResourcePolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/resourcePolicies/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeResourcePolicyName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeResourcePolicySnapshotSchedulePolicy(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["schedule"] =
		flattenComputeResourcePolicySnapshotSchedulePolicySchedule(original["schedule"], d)
	transformed["retention_policy"] =
		flattenComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy(original["retentionPolicy"], d)
	transformed["snapshot_properties"] =
		flattenComputeResourcePolicySnapshotSchedulePolicySnapshotProperties(original["snapshotProperties"], d)
	return []interface{}{transformed}
}
func flattenComputeResourcePolicySnapshotSchedulePolicySchedule(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["hourly_schedule"] =
		flattenComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule(original["hourlySchedule"], d)
	transformed["daily_schedule"] =
		flattenComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule(original["dailySchedule"], d)
	transformed["weekly_schedule"] =
		flattenComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule(original["weeklySchedule"], d)
	return []interface{}{transformed}
}
func flattenComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["hours_in_cycle"] =
		flattenComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleHoursInCycle(original["hoursInCycle"], d)
	transformed["start_time"] =
		flattenComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleStartTime(original["startTime"], d)
	return []interface{}{transformed}
}
func flattenComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleHoursInCycle(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleStartTime(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["days_in_cycle"] =
		flattenComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleDaysInCycle(original["daysInCycle"], d)
	transformed["start_time"] =
		flattenComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleStartTime(original["startTime"], d)
	return []interface{}{transformed}
}
func flattenComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleDaysInCycle(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleStartTime(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["day_of_weeks"] =
		flattenComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeks(original["dayOfWeeks"], d)
	return []interface{}{transformed}
}
func flattenComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeks(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(computeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"start_time": flattenComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksStartTime(original["startTime"], d),
			"day":        flattenComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksDay(original["day"], d),
		})
	}
	return transformed
}
func flattenComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksStartTime(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksDay(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["max_retention_days"] =
		flattenComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyMaxRetentionDays(original["maxRetentionDays"], d)
	transformed["on_source_disk_delete"] =
		flattenComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOnSourceDiskDelete(original["onSourceDiskDelete"], d)
	return []interface{}{transformed}
}
func flattenComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyMaxRetentionDays(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOnSourceDiskDelete(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeResourcePolicySnapshotSchedulePolicySnapshotProperties(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["labels"] =
		flattenComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesLabels(original["labels"], d)
	transformed["storage_locations"] =
		flattenComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesStorageLocations(original["storageLocations"], d)
	transformed["guest_flush"] =
		flattenComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesGuestFlush(original["guestFlush"], d)
	return []interface{}{transformed}
}
func flattenComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesLabels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesStorageLocations(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesGuestFlush(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeResourcePolicyRegion(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeResourcePolicyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSchedule, err := expandComputeResourcePolicySnapshotSchedulePolicySchedule(original["schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSchedule); val.IsValid() && !isEmptyValue(val) {
		transformed["schedule"] = transformedSchedule
	}

	transformedRetentionPolicy, err := expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy(original["retention_policy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRetentionPolicy); val.IsValid() && !isEmptyValue(val) {
		transformed["retentionPolicy"] = transformedRetentionPolicy
	}

	transformedSnapshotProperties, err := expandComputeResourcePolicySnapshotSchedulePolicySnapshotProperties(original["snapshot_properties"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSnapshotProperties); val.IsValid() && !isEmptyValue(val) {
		transformed["snapshotProperties"] = transformedSnapshotProperties
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicySchedule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHourlySchedule, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule(original["hourly_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHourlySchedule); val.IsValid() && !isEmptyValue(val) {
		transformed["hourlySchedule"] = transformedHourlySchedule
	}

	transformedDailySchedule, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule(original["daily_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDailySchedule); val.IsValid() && !isEmptyValue(val) {
		transformed["dailySchedule"] = transformedDailySchedule
	}

	transformedWeeklySchedule, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule(original["weekly_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeeklySchedule); val.IsValid() && !isEmptyValue(val) {
		transformed["weeklySchedule"] = transformedWeeklySchedule
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHoursInCycle, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleHoursInCycle(original["hours_in_cycle"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHoursInCycle); val.IsValid() && !isEmptyValue(val) {
		transformed["hoursInCycle"] = transformedHoursInCycle
	}

	transformedStartTime, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !isEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleHoursInCycle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleStartTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDaysInCycle, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleDaysInCycle(original["days_in_cycle"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDaysInCycle); val.IsValid() && !isEmptyValue(val) {
		transformed["daysInCycle"] = transformedDaysInCycle
	}

	transformedStartTime, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !isEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleDaysInCycle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleStartTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDayOfWeeks, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeks(original["day_of_weeks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDayOfWeeks); val.IsValid() && !isEmptyValue(val) {
		transformed["dayOfWeeks"] = transformedDayOfWeeks
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedStartTime, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksStartTime(original["start_time"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !isEmptyValue(val) {
			transformed["startTime"] = transformedStartTime
		}

		transformedDay, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksDay(original["day"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDay); val.IsValid() && !isEmptyValue(val) {
			transformed["day"] = transformedDay
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksStartTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksDay(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxRetentionDays, err := expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyMaxRetentionDays(original["max_retention_days"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxRetentionDays); val.IsValid() && !isEmptyValue(val) {
		transformed["maxRetentionDays"] = transformedMaxRetentionDays
	}

	transformedOnSourceDiskDelete, err := expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOnSourceDiskDelete(original["on_source_disk_delete"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOnSourceDiskDelete); val.IsValid() && !isEmptyValue(val) {
		transformed["onSourceDiskDelete"] = transformedOnSourceDiskDelete
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyMaxRetentionDays(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOnSourceDiskDelete(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicySnapshotProperties(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLabels, err := expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !isEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	transformedStorageLocations, err := expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesStorageLocations(original["storage_locations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageLocations); val.IsValid() && !isEmptyValue(val) {
		transformed["storageLocations"] = transformedStorageLocations
	}

	transformedGuestFlush, err := expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesGuestFlush(original["guest_flush"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGuestFlush); val.IsValid() && !isEmptyValue(val) {
		transformed["guestFlush"] = transformedGuestFlush
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesStorageLocations(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesGuestFlush(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
