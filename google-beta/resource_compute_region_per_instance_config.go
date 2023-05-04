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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceComputeRegionPerInstanceConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRegionPerInstanceConfigCreate,
		Read:   resourceComputeRegionPerInstanceConfigRead,
		Update: resourceComputeRegionPerInstanceConfigUpdate,
		Delete: resourceComputeRegionPerInstanceConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRegionPerInstanceConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name for this per-instance config and its corresponding instance.`,
			},
			"region_instance_group_manager": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The region instance group manager this instance config is part of.`,
			},
			"preserved_state": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The preserved state for this instance.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disk": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: `Stateful disks for the instance.`,
							Elem:        computeRegionPerInstanceConfigPreservedStateDiskSchema(),
							// Default schema.HashSchema is used.
						},
						"external_ip": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: `Preserved external IPs defined for this instance. This map is keyed with the name of the network interface.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"auto_delete": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: verify.ValidateEnum([]string{"NEVER", "ON_PERMANENT_INSTANCE_DELETION", ""}),
										Description:  `These stateful IPs will never be released during autohealing, update or VM instance recreate operations. This flag is used to configure if the IP reservation should be deleted after it is no longer used by the group, e.g. when the given instance or the whole group is deleted. Default value: "NEVER" Possible values: ["NEVER", "ON_PERMANENT_INSTANCE_DELETION"]`,
										Default:      "NEVER",
									},
									"ip_address": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `Ip address representation`,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": {
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: compareSelfLinkOrResourceName,
													Description:      `The URL of the reservation for this IP address.`,
												},
											},
										},
									},
								},
							},
						},
						"internal_ip": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: `Preserved internal IPs defined for this instance. This map is keyed with the name of the network interface.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"auto_delete": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: verify.ValidateEnum([]string{"NEVER", "ON_PERMANENT_INSTANCE_DELETION", ""}),
										Description:  `These stateful IPs will never be released during autohealing, update or VM instance recreate operations. This flag is used to configure if the IP reservation should be deleted after it is no longer used by the group, e.g. when the given instance or the whole group is deleted. Default value: "NEVER" Possible values: ["NEVER", "ON_PERMANENT_INSTANCE_DELETION"]`,
										Default:      "NEVER",
									},
									"ip_address": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `Ip address representation`,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": {
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: compareSelfLinkOrResourceName,
													Description:      `The URL of the reservation for this IP address.`,
												},
											},
										},
									},
								},
							},
						},
						"metadata": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: `Preserved metadata defined for this instance. This is a list of key->value pairs.`,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `Region where the containing instance group manager is located`,
			},
			"minimal_action": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NONE",
				Description: `The minimal action to perform on the instance during an update.
Default is 'NONE'. Possible values are:
* REPLACE
* RESTART
* REFRESH
* NONE`,
			},
			"most_disruptive_allowed_action": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "REPLACE",
				Description: `The most disruptive action to perform on the instance during an update.
Default is 'REPLACE'. Possible values are:
* REPLACE
* RESTART
* REFRESH
* NONE`,
			},
			"remove_instance_state_on_destroy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				Description: `When true, deleting this config will immediately remove any specified state from the underlying instance.
When false, deleting this config will *not* immediately remove any state from the underlying instance.
State will be removed on the next instance recreation or update.`,
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

func computeRegionPerInstanceConfigPreservedStateDiskSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"device_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance.`,
			},
			"source": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The URI of an existing persistent disk to attach under the specified device-name in the format
'projects/project-id/zones/zone/disks/disk-name'.`,
			},
			"delete_rule": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"NEVER", "ON_PERMANENT_INSTANCE_DELETION", ""}),
				Description: `A value that prescribes what should happen to the stateful disk when the VM instance is deleted.
The available options are 'NEVER' and 'ON_PERMANENT_INSTANCE_DELETION'.
'NEVER' - detach the disk when the VM is deleted, but do not delete the disk.
'ON_PERMANENT_INSTANCE_DELETION' will delete the stateful disk when the VM is permanently
deleted from the instance group. Default value: "NEVER" Possible values: ["NEVER", "ON_PERMANENT_INSTANCE_DELETION"]`,
				Default: "NEVER",
			},
			"mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"READ_ONLY", "READ_WRITE", ""}),
				Description:  `The mode of the disk. Default value: "READ_WRITE" Possible values: ["READ_ONLY", "READ_WRITE"]`,
				Default:      "READ_WRITE",
			},
		},
	}
}

func resourceComputeRegionPerInstanceConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandNestedComputeRegionPerInstanceConfigName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	preservedStateProp, err := expandNestedComputeRegionPerInstanceConfigPreservedState(d.Get("preserved_state"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("preserved_state"); !isEmptyValue(reflect.ValueOf(preservedStateProp)) && (ok || !reflect.DeepEqual(v, preservedStateProp)) {
		obj["preservedState"] = preservedStateProp
	}

	obj, err = resourceComputeRegionPerInstanceConfigEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := ReplaceVars(d, config, "instanceGroupManager/{{project}}/{{region}}/{{region_instance_group_manager}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/instanceGroupManagers/{{region_instance_group_manager}}/createInstances")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionPerInstanceConfig: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionPerInstanceConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating RegionPerInstanceConfig: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{project}}/{{region}}/{{region_instance_group_manager}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating RegionPerInstanceConfig", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RegionPerInstanceConfig: %s", err)
	}

	log.Printf("[DEBUG] Finished creating RegionPerInstanceConfig %q: %#v", d.Id(), res)

	return resourceComputeRegionPerInstanceConfigRead(d, meta)
}

func resourceComputeRegionPerInstanceConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/instanceGroupManagers/{{region_instance_group_manager}}/listPerInstanceConfigs")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionPerInstanceConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "POST", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeRegionPerInstanceConfig %q", d.Id()))
	}

	res, err = flattenNestedComputeRegionPerInstanceConfig(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ComputeRegionPerInstanceConfig because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("minimal_action"); !ok {
		if err := d.Set("minimal_action", "NONE"); err != nil {
			return fmt.Errorf("Error setting minimal_action: %s", err)
		}
	}
	if _, ok := d.GetOkExists("most_disruptive_allowed_action"); !ok {
		if err := d.Set("most_disruptive_allowed_action", "REPLACE"); err != nil {
			return fmt.Errorf("Error setting most_disruptive_allowed_action: %s", err)
		}
	}
	if _, ok := d.GetOkExists("remove_instance_state_on_destroy"); !ok {
		if err := d.Set("remove_instance_state_on_destroy", false); err != nil {
			return fmt.Errorf("Error setting remove_instance_state_on_destroy: %s", err)
		}
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionPerInstanceConfig: %s", err)
	}

	region, err := getRegion(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("region", region); err != nil {
		return fmt.Errorf("Error reading RegionPerInstanceConfig: %s", err)
	}

	if err := d.Set("name", flattenNestedComputeRegionPerInstanceConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionPerInstanceConfig: %s", err)
	}
	if err := d.Set("preserved_state", flattenNestedComputeRegionPerInstanceConfigPreservedState(res["preservedState"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionPerInstanceConfig: %s", err)
	}

	return nil
}

func resourceComputeRegionPerInstanceConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionPerInstanceConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	nameProp, err := expandNestedComputeRegionPerInstanceConfigName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	preservedStateProp, err := expandNestedComputeRegionPerInstanceConfigPreservedState(d.Get("preserved_state"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("preserved_state"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, preservedStateProp)) {
		obj["preservedState"] = preservedStateProp
	}

	obj, err = resourceComputeRegionPerInstanceConfigUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := ReplaceVars(d, config, "instanceGroupManager/{{project}}/{{region}}/{{region_instance_group_manager}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/instanceGroupManagers/{{region_instance_group_manager}}/updatePerInstanceConfigs")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RegionPerInstanceConfig %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating RegionPerInstanceConfig %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating RegionPerInstanceConfig %q: %#v", d.Id(), res)
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Updating RegionPerInstanceConfig", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	// Instance name in applyUpdatesToInstances request must include zone
	instanceName, err := findInstanceName(d, config)
	if err != nil {
		return err
	}

	obj = make(map[string]interface{})
	obj["instances"] = []string{instanceName}

	minAction := d.Get("minimal_action")
	if minAction == "" {
		minAction = "NONE"
	}
	obj["minimalAction"] = minAction

	mostDisruptiveAction := d.Get("most_disruptive_allowed_action")
	if isEmptyValue(reflect.ValueOf(mostDisruptiveAction)) {
		mostDisruptiveAction = "REPLACE"
	}
	obj["mostDisruptiveAllowedAction"] = mostDisruptiveAction

	url, err = ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/instanceGroupManagers/{{region_instance_group_manager}}/applyUpdatesToInstances")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Applying updates to PerInstanceConfig %q: %#v", d.Id(), obj)
	res, err = transport_tpg.SendRequestWithTimeout(config, "POST", project, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating PerInstanceConfig %q: %s", d.Id(), err)
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Applying update to PerInstanceConfig", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}
	return resourceComputeRegionPerInstanceConfigRead(d, meta)
}

func resourceComputeRegionPerInstanceConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	lockName, err := ReplaceVars(d, config, "instanceGroupManager/{{project}}/{{region}}/{{region_instance_group_manager}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/instanceGroupManagers/{{region_instance_group_manager}}/deletePerInstanceConfigs")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	obj = map[string]interface{}{
		"names": [1]string{d.Get("name").(string)},
	}
	log.Printf("[DEBUG] Deleting RegionPerInstanceConfig %q", d.Id())

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", project, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "RegionPerInstanceConfig")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting RegionPerInstanceConfig", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	// Potentially delete the state managed by this config
	if d.Get("remove_instance_state_on_destroy").(bool) {
		// Instance name in applyUpdatesToInstances request must include zone
		instanceName, err := findInstanceName(d, config)
		if err != nil {
			return err
		}

		obj = make(map[string]interface{})
		obj["instances"] = []string{instanceName}

		// Updates must be applied to the instance after deleting the PerInstanceConfig
		url, err = ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/instanceGroupManagers/{{region_instance_group_manager}}/applyUpdatesToInstances")
		if err != nil {
			return err
		}

		log.Printf("[DEBUG] Applying updates to PerInstanceConfig %q: %#v", d.Id(), obj)
		res, err = transport_tpg.SendRequestWithTimeout(config, "POST", project, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return fmt.Errorf("Error updating PerInstanceConfig %q: %s", d.Id(), err)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Applying update to PerInstanceConfig", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return fmt.Errorf("Error deleting PerInstanceConfig %q: %s", d.Id(), err)
		}

		// RegionPerInstanceConfig goes into "DELETING" state while the instance is actually deleted
		err = PollingWaitTime(resourceComputeRegionPerInstanceConfigPollRead(d, meta), PollCheckInstanceConfigDeleted, "Deleting RegionPerInstanceConfig", d.Timeout(schema.TimeoutDelete), 1)
		if err != nil {
			return fmt.Errorf("Error waiting for delete on RegionPerInstanceConfig %q: %s", d.Id(), err)
		}
	}

	log.Printf("[DEBUG] Finished deleting RegionPerInstanceConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRegionPerInstanceConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/instanceGroupManagers/(?P<region_instance_group_manager>[^/]+)/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<region_instance_group_manager>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<region_instance_group_manager>[^/]+)/(?P<name>[^/]+)",
		"(?P<region_instance_group_manager>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "{{project}}/{{region}}/{{region_instance_group_manager}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Explicitly set virtual fields to default values on import
	if err := d.Set("minimal_action", "NONE"); err != nil {
		return nil, fmt.Errorf("Error setting minimal_action: %s", err)
	}
	if err := d.Set("most_disruptive_allowed_action", "REPLACE"); err != nil {
		return nil, fmt.Errorf("Error setting most_disruptive_allowed_action: %s", err)
	}
	if err := d.Set("remove_instance_state_on_destroy", false); err != nil {
		return nil, fmt.Errorf("Error setting remove_instance_state_on_destroy: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenNestedComputeRegionPerInstanceConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeRegionPerInstanceConfigPreservedState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["metadata"] =
		flattenNestedComputeRegionPerInstanceConfigPreservedStateMetadata(original["metadata"], d, config)
	transformed["disk"] =
		flattenNestedComputeRegionPerInstanceConfigPreservedStateDisk(original["disks"], d, config)
	transformed["internal_ip"] =
		flattenNestedComputeRegionPerInstanceConfigPreservedStateInternalIp(original["internalIPs"], d, config)
	transformed["external_ip"] =
		flattenNestedComputeRegionPerInstanceConfigPreservedStateExternalIp(original["externalIPs"], d, config)
	return []interface{}{transformed}
}
func flattenNestedComputeRegionPerInstanceConfigPreservedStateMetadata(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeRegionPerInstanceConfigPreservedStateDisk(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	disks := v.(map[string]interface{})
	transformed := make([]interface{}, 0, len(disks))
	for devName, deleteRuleRaw := range disks {
		diskObj := deleteRuleRaw.(map[string]interface{})
		source, err := tpgresource.GetRelativePath(diskObj["source"].(string))
		if err != nil {
			source = diskObj["source"].(string)
		}
		transformed = append(transformed, map[string]interface{}{
			"device_name": devName,
			"delete_rule": diskObj["autoDelete"],
			"source":      source,
			"mode":        diskObj["mode"],
		})
	}
	return transformed
}

func flattenNestedComputeRegionPerInstanceConfigPreservedStateInternalIp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.(map[string]interface{})
	transformed := make([]interface{}, 0, len(l))
	for k, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"interface_name": k,
			"auto_delete":    flattenNestedComputeRegionPerInstanceConfigPreservedStateInternalIpAutoDelete(original["autoDelete"], d, config),
			"ip_address":     flattenNestedComputeRegionPerInstanceConfigPreservedStateInternalIpIpAddress(original["ipAddress"], d, config),
		})
	}
	return transformed
}
func flattenNestedComputeRegionPerInstanceConfigPreservedStateInternalIpAutoDelete(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeRegionPerInstanceConfigPreservedStateInternalIpIpAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["address"] =
		flattenNestedComputeRegionPerInstanceConfigPreservedStateInternalIpIpAddressAddress(original["address"], d, config)
	return []interface{}{transformed}
}
func flattenNestedComputeRegionPerInstanceConfigPreservedStateInternalIpIpAddressAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenNestedComputeRegionPerInstanceConfigPreservedStateExternalIp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.(map[string]interface{})
	transformed := make([]interface{}, 0, len(l))
	for k, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"interface_name": k,
			"auto_delete":    flattenNestedComputeRegionPerInstanceConfigPreservedStateExternalIpAutoDelete(original["autoDelete"], d, config),
			"ip_address":     flattenNestedComputeRegionPerInstanceConfigPreservedStateExternalIpIpAddress(original["ipAddress"], d, config),
		})
	}
	return transformed
}
func flattenNestedComputeRegionPerInstanceConfigPreservedStateExternalIpAutoDelete(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeRegionPerInstanceConfigPreservedStateExternalIpIpAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["address"] =
		flattenNestedComputeRegionPerInstanceConfigPreservedStateExternalIpIpAddressAddress(original["address"], d, config)
	return []interface{}{transformed}
}
func flattenNestedComputeRegionPerInstanceConfigPreservedStateExternalIpIpAddressAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func expandNestedComputeRegionPerInstanceConfigName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedComputeRegionPerInstanceConfigPreservedState(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMetadata, err := expandNestedComputeRegionPerInstanceConfigPreservedStateMetadata(original["metadata"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMetadata); val.IsValid() && !isEmptyValue(val) {
		transformed["metadata"] = transformedMetadata
	}

	transformedDisk, err := expandNestedComputeRegionPerInstanceConfigPreservedStateDisk(original["disk"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisk); val.IsValid() && !isEmptyValue(val) {
		transformed["disks"] = transformedDisk
	}

	transformedInternalIp, err := expandNestedComputeRegionPerInstanceConfigPreservedStateInternalIp(original["internal_ip"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInternalIp); val.IsValid() && !isEmptyValue(val) {
		transformed["internalIPs"] = transformedInternalIp
	}

	transformedExternalIp, err := expandNestedComputeRegionPerInstanceConfigPreservedStateExternalIp(original["external_ip"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExternalIp); val.IsValid() && !isEmptyValue(val) {
		transformed["externalIPs"] = transformedExternalIp
	}

	return transformed, nil
}

func expandNestedComputeRegionPerInstanceConfigPreservedStateMetadata(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandNestedComputeRegionPerInstanceConfigPreservedStateDisk(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	l := v.(*schema.Set).List()
	req := make(map[string]interface{})
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		deviceName := original["device_name"].(string)
		diskObj := make(map[string]interface{})
		deleteRule := original["delete_rule"].(string)
		if deleteRule != "" {
			diskObj["autoDelete"] = deleteRule
		}
		source := original["source"]
		if source != "" {
			diskObj["source"] = source
		}
		mode := original["mode"]
		if source != "" {
			diskObj["mode"] = mode
		}
		req[deviceName] = diskObj
	}
	return req, nil
}

func expandNestedComputeRegionPerInstanceConfigPreservedStateInternalIp(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAutoDelete, err := expandNestedComputeRegionPerInstanceConfigPreservedStateInternalIpAutoDelete(original["auto_delete"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAutoDelete); val.IsValid() && !isEmptyValue(val) {
			transformed["autoDelete"] = transformedAutoDelete
		}

		transformedIpAddress, err := expandNestedComputeRegionPerInstanceConfigPreservedStateInternalIpIpAddress(original["ip_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !isEmptyValue(val) {
			transformed["ipAddress"] = transformedIpAddress
		}

		transformedInterfaceName, err := expandString(original["interface_name"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedInterfaceName] = transformed
	}
	return m, nil
}

func expandNestedComputeRegionPerInstanceConfigPreservedStateInternalIpAutoDelete(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedComputeRegionPerInstanceConfigPreservedStateInternalIpIpAddress(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAddress, err := expandNestedComputeRegionPerInstanceConfigPreservedStateInternalIpIpAddressAddress(original["address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAddress); val.IsValid() && !isEmptyValue(val) {
		transformed["address"] = transformedAddress
	}

	return transformed, nil
}

func expandNestedComputeRegionPerInstanceConfigPreservedStateInternalIpIpAddressAddress(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("addresses", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for address: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandNestedComputeRegionPerInstanceConfigPreservedStateExternalIp(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAutoDelete, err := expandNestedComputeRegionPerInstanceConfigPreservedStateExternalIpAutoDelete(original["auto_delete"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAutoDelete); val.IsValid() && !isEmptyValue(val) {
			transformed["autoDelete"] = transformedAutoDelete
		}

		transformedIpAddress, err := expandNestedComputeRegionPerInstanceConfigPreservedStateExternalIpIpAddress(original["ip_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !isEmptyValue(val) {
			transformed["ipAddress"] = transformedIpAddress
		}

		transformedInterfaceName, err := expandString(original["interface_name"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedInterfaceName] = transformed
	}
	return m, nil
}

func expandNestedComputeRegionPerInstanceConfigPreservedStateExternalIpAutoDelete(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedComputeRegionPerInstanceConfigPreservedStateExternalIpIpAddress(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAddress, err := expandNestedComputeRegionPerInstanceConfigPreservedStateExternalIpIpAddressAddress(original["address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAddress); val.IsValid() && !isEmptyValue(val) {
		transformed["address"] = transformedAddress
	}

	return transformed, nil
}

func expandNestedComputeRegionPerInstanceConfigPreservedStateExternalIpIpAddressAddress(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("addresses", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for address: %s", err)
	}
	return f.RelativeLink(), nil
}

func resourceComputeRegionPerInstanceConfigEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	wrappedReq := map[string]interface{}{
		"instances": []interface{}{obj},
	}
	return wrappedReq, nil
}

func resourceComputeRegionPerInstanceConfigUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// updates and creates use different wrapping object names
	wrappedReq := map[string]interface{}{
		"perInstanceConfigs": []interface{}{obj},
	}
	return wrappedReq, nil
}

func flattenNestedComputeRegionPerInstanceConfig(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["items"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value items. Actual value: %v", v)
	}

	_, item, err := resourceComputeRegionPerInstanceConfigFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceComputeRegionPerInstanceConfigFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedName, err := expandNestedComputeRegionPerInstanceConfigName(d.Get("name"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedName := flattenNestedComputeRegionPerInstanceConfigName(expectedName, d, meta.(*transport_tpg.Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		itemName := flattenNestedComputeRegionPerInstanceConfigName(item["name"], d, meta.(*transport_tpg.Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemName)) && isEmptyValue(reflect.ValueOf(expectedFlattenedName))) && !reflect.DeepEqual(itemName, expectedFlattenedName) {
			log.Printf("[DEBUG] Skipping item with name= %#v, looking for %#v)", itemName, expectedFlattenedName)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
