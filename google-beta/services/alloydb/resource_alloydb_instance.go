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

package alloydb

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
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceAlloydbInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlloydbInstanceCreate,
		Read:   resourceAlloydbInstanceRead,
		Update: resourceAlloydbInstanceUpdate,
		Delete: resourceAlloydbInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAlloydbInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(40 * time.Minute),
			Update: schema.DefaultTimeout(40 * time.Minute),
			Delete: schema.DefaultTimeout(40 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.SetAnnotationsDiff,
		),

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `Identifies the alloydb cluster. Must be in the format
'projects/{project}/locations/{location}/clusters/{cluster_id}'`,
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the alloydb instance.`,
			},
			"instance_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"PRIMARY", "READ_POOL"}),
				Description:  `The type of the instance. If the instance type is READ_POOL, provide the associated PRIMARY instance in the 'depends_on' meta-data attribute. Possible values: ["PRIMARY", "READ_POOL"]`,
			},
			"annotations": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"availability_type": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"AVAILABILITY_TYPE_UNSPECIFIED", "ZONAL", "REGIONAL", ""}),
				Description: `'Availability type of an Instance. Defaults to REGIONAL for both primary and read instances.
Note that primary and read instances can have different availability types.
Only READ_POOL instance supports ZONAL type. Users can't specify the zone for READ_POOL instance.
Zone is automatically chosen from the list of zones in the region specified.
Read pool of size 1 can only have zonal availability. Read pools with node count of 2 or more
can have regional availability (nodes are present in 2 or more zones in a region).' Possible values: ["AVAILABILITY_TYPE_UNSPECIFIED", "ZONAL", "REGIONAL"]`,
			},
			"database_flags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `User-settable and human-readable display name for the Instance.`,
			},
			"gce_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `User-defined labels for the alloydb instance.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"machine_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: `Configurations for the machines that host the underlying database engine.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: `The number of CPU's in the VM instance.`,
						},
					},
				},
			},
			"read_pool_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Read pool specific config. If the instance type is READ_POOL, this configuration must be provided.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: `Read capacity, i.e. number of nodes in a read pool instance.`,
						},
					},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the Instance was created in UTC.`,
			},
			"effective_annotations": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"ip_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The IP address for the Instance. This is the connection endpoint for an end-user application.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the instance resource.`,
			},
			"reconciling": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `Set to true if the current state of Instance does not match the user's intended state, and the service is actively updating the resource to reconcile them. This can happen due to user-triggered updates or system actions like failover or maintenance.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The current state of the alloydb instance.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The system-generated UID of the resource.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the Instance was updated in UTC.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceAlloydbInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandAlloydbInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	gceZoneProp, err := expandAlloydbInstanceGceZone(d.Get("gce_zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gce_zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(gceZoneProp)) && (ok || !reflect.DeepEqual(v, gceZoneProp)) {
		obj["gceZone"] = gceZoneProp
	}
	databaseFlagsProp, err := expandAlloydbInstanceDatabaseFlags(d.Get("database_flags"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("database_flags"); !tpgresource.IsEmptyValue(reflect.ValueOf(databaseFlagsProp)) && (ok || !reflect.DeepEqual(v, databaseFlagsProp)) {
		obj["databaseFlags"] = databaseFlagsProp
	}
	availabilityTypeProp, err := expandAlloydbInstanceAvailabilityType(d.Get("availability_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("availability_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(availabilityTypeProp)) && (ok || !reflect.DeepEqual(v, availabilityTypeProp)) {
		obj["availabilityType"] = availabilityTypeProp
	}
	instanceTypeProp, err := expandAlloydbInstanceInstanceType(d.Get("instance_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("instance_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(instanceTypeProp)) && (ok || !reflect.DeepEqual(v, instanceTypeProp)) {
		obj["instanceType"] = instanceTypeProp
	}
	readPoolConfigProp, err := expandAlloydbInstanceReadPoolConfig(d.Get("read_pool_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("read_pool_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(readPoolConfigProp)) && (ok || !reflect.DeepEqual(v, readPoolConfigProp)) {
		obj["readPoolConfig"] = readPoolConfigProp
	}
	machineConfigProp, err := expandAlloydbInstanceMachineConfig(d.Get("machine_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("machine_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(machineConfigProp)) && (ok || !reflect.DeepEqual(v, machineConfigProp)) {
		obj["machineConfig"] = machineConfigProp
	}
	labelsProp, err := expandAlloydbInstanceEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandAlloydbInstanceEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/instances?instanceId={{instance_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Instance: %#v", obj)
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
		return fmt.Errorf("Error creating Instance: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{cluster}}/instances/{{instance_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = AlloydbOperationWaitTime(
		config, res, project, "Creating Instance", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Instance: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceAlloydbInstanceRead(d, meta)
}

func resourceAlloydbInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/instances/{{instance_id}}")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("AlloydbInstance %q", d.Id()))
	}

	if err := d.Set("name", flattenAlloydbInstanceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("create_time", flattenAlloydbInstanceCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("update_time", flattenAlloydbInstanceUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("uid", flattenAlloydbInstanceUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("labels", flattenAlloydbInstanceLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("annotations", flattenAlloydbInstanceAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("state", flattenAlloydbInstanceState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("gce_zone", flattenAlloydbInstanceGceZone(res["gceZone"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("reconciling", flattenAlloydbInstanceReconciling(res["reconciling"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("database_flags", flattenAlloydbInstanceDatabaseFlags(res["databaseFlags"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("availability_type", flattenAlloydbInstanceAvailabilityType(res["availabilityType"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("instance_type", flattenAlloydbInstanceInstanceType(res["instanceType"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("ip_address", flattenAlloydbInstanceIpAddress(res["ipAddress"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("read_pool_config", flattenAlloydbInstanceReadPoolConfig(res["readPoolConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("machine_config", flattenAlloydbInstanceMachineConfig(res["machineConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("terraform_labels", flattenAlloydbInstanceTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("effective_labels", flattenAlloydbInstanceEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("effective_annotations", flattenAlloydbInstanceEffectiveAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	return nil
}

func resourceAlloydbInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	displayNameProp, err := expandAlloydbInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	gceZoneProp, err := expandAlloydbInstanceGceZone(d.Get("gce_zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gce_zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, gceZoneProp)) {
		obj["gceZone"] = gceZoneProp
	}
	databaseFlagsProp, err := expandAlloydbInstanceDatabaseFlags(d.Get("database_flags"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("database_flags"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, databaseFlagsProp)) {
		obj["databaseFlags"] = databaseFlagsProp
	}
	availabilityTypeProp, err := expandAlloydbInstanceAvailabilityType(d.Get("availability_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("availability_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, availabilityTypeProp)) {
		obj["availabilityType"] = availabilityTypeProp
	}
	readPoolConfigProp, err := expandAlloydbInstanceReadPoolConfig(d.Get("read_pool_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("read_pool_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, readPoolConfigProp)) {
		obj["readPoolConfig"] = readPoolConfigProp
	}
	machineConfigProp, err := expandAlloydbInstanceMachineConfig(d.Get("machine_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("machine_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, machineConfigProp)) {
		obj["machineConfig"] = machineConfigProp
	}
	labelsProp, err := expandAlloydbInstanceEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandAlloydbInstanceEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/instances/{{instance_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Instance %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("gce_zone") {
		updateMask = append(updateMask, "gceZone")
	}

	if d.HasChange("database_flags") {
		updateMask = append(updateMask, "databaseFlags")
	}

	if d.HasChange("availability_type") {
		updateMask = append(updateMask, "availabilityType")
	}

	if d.HasChange("read_pool_config") {
		updateMask = append(updateMask, "readPoolConfig")
	}

	if d.HasChange("machine_config") {
		updateMask = append(updateMask, "machineConfig")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("effective_annotations") {
		updateMask = append(updateMask, "annotations")
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
		return fmt.Errorf("Error updating Instance %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Instance %q: %#v", d.Id(), res)
	}

	err = AlloydbOperationWaitTime(
		config, res, project, "Updating Instance", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceAlloydbInstanceRead(d, meta)
}

func resourceAlloydbInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/instances/{{instance_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Instance %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Instance")
	}

	err = AlloydbOperationWaitTime(
		config, res, project, "Deleting Instance", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Instance %q: %#v", d.Id(), res)
	return nil
}

func resourceAlloydbInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{
		"(?P<cluster>.+)/instances/(?P<instance_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{cluster}}/instances/{{instance_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAlloydbInstanceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbInstanceCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbInstanceUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbInstanceUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbInstanceLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenAlloydbInstanceAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("annotations"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenAlloydbInstanceState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbInstanceGceZone(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbInstanceReconciling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbInstanceDatabaseFlags(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbInstanceAvailabilityType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbInstanceInstanceType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbInstanceIpAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbInstanceReadPoolConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["node_count"] =
		flattenAlloydbInstanceReadPoolConfigNodeCount(original["nodeCount"], d, config)
	return []interface{}{transformed}
}
func flattenAlloydbInstanceReadPoolConfigNodeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenAlloydbInstanceMachineConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["cpu_count"] =
		flattenAlloydbInstanceMachineConfigCpuCount(original["cpuCount"], d, config)
	return []interface{}{transformed}
}
func flattenAlloydbInstanceMachineConfigCpuCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenAlloydbInstanceTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenAlloydbInstanceEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbInstanceEffectiveAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandAlloydbInstanceDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceGceZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceDatabaseFlags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAlloydbInstanceAvailabilityType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceInstanceType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceReadPoolConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNodeCount, err := expandAlloydbInstanceReadPoolConfigNodeCount(original["node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nodeCount"] = transformedNodeCount
	}

	return transformed, nil
}

func expandAlloydbInstanceReadPoolConfigNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceMachineConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCpuCount, err := expandAlloydbInstanceMachineConfigCpuCount(original["cpu_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCpuCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cpuCount"] = transformedCpuCount
	}

	return transformed, nil
}

func expandAlloydbInstanceMachineConfigCpuCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAlloydbInstanceEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
