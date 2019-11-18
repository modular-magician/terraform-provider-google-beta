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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceComputeReservation() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeReservationCreate,
		Read:   resourceComputeReservationRead,
		Delete: resourceComputeReservationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeReservationImport,
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
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"specific_reservation": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: `Reservation for instances with specific machine shapes.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": {
							Type:        schema.TypeInt,
							Required:    true,
							ForceNew:    true,
							Description: `The number of resources that are allocated.`,
						},
						"instance_properties": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Description: `The instance properties for the reservation.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"machine_type": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: `The name of the machine type to reserve.`,
									},
									"guest_accelerators": {
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										Description: `Guest accelerator type and count.`,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accelerator_count": {
													Type:     schema.TypeInt,
													Required: true,
													ForceNew: true,
													Description: `The number of the guest accelerator cards exposed to
this instance.`,
												},
												"accelerator_type": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
													Description: `The full or partial URL of the accelerator type to
attach to this instance. For example:
'projects/my-project/zones/us-central1-c/acceleratorTypes/nvidia-tesla-p100'

If you are creating an instance template, specify only the accelerator name.`,
												},
											},
										},
									},
									"local_ssds": {
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Description: `The amount of local ssd to reserve with each instance. This
reserves disks of type 'local-ssd'.`,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"disk_size_gb": {
													Type:        schema.TypeInt,
													Required:    true,
													ForceNew:    true,
													Description: `The size of the disk in base-2 GB.`,
												},
												"interface": {
													Type:         schema.TypeString,
													Optional:     true,
													ForceNew:     true,
													ValidateFunc: validation.StringInSlice([]string{"SCSI", "NVME", ""}, false),
													Description: `The disk interface to use for attaching this disk, one
of 'SCSI' or 'NVME'. The default is 'SCSI'.`,
													Default: "SCSI",
												},
											},
										},
									},
									"min_cpu_platform": {
										Type:     schema.TypeString,
										Computed: true,
										Optional: true,
										ForceNew: true,
										Description: `The minimum CPU platform for the reservation. For example,
'"Intel Skylake"'. See
the CPU platform availability reference](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform#availablezones)
for information on available CPU platforms.`,
									},
								},
							},
						},
						"in_use_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `How many instances are in use.`,
						},
					},
				},
			},
			"zone": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The zone where the reservation is made.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"specific_reservation_required": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `When set to true, only VMs that target this reservation by name can
consume this reservation. Otherwise, it can be consumed by VMs with
affinity for any reservation. Defaults to false.`,
				Default: false,
			},
			"commitment": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Full or partial URL to a parent commitment. This field displays for
reservations that are tied to a commitment.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The status of the reservation.`,
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

func resourceComputeReservationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeReservationDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeReservationName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	specificReservationRequiredProp, err := expandComputeReservationSpecificReservationRequired(d.Get("specific_reservation_required"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("specific_reservation_required"); !isEmptyValue(reflect.ValueOf(specificReservationRequiredProp)) && (ok || !reflect.DeepEqual(v, specificReservationRequiredProp)) {
		obj["specificReservationRequired"] = specificReservationRequiredProp
	}
	specificReservationProp, err := expandComputeReservationSpecificReservation(d.Get("specific_reservation"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("specific_reservation"); !isEmptyValue(reflect.ValueOf(specificReservationProp)) && (ok || !reflect.DeepEqual(v, specificReservationProp)) {
		obj["specificReservation"] = specificReservationProp
	}
	zoneProp, err := expandComputeReservationZone(d.Get("zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("zone"); !isEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/reservations")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Reservation: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Reservation: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/zones/{{zone}}/reservations/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	waitErr := computeOperationWaitTime(
		config, res, project, "Creating Reservation",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Reservation: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating Reservation %q: %#v", d.Id(), res)

	return resourceComputeReservationRead(d, meta)
}

func resourceComputeReservationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/reservations/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeReservation %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeReservationCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}
	if err := d.Set("description", flattenComputeReservationDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}
	if err := d.Set("name", flattenComputeReservationName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}
	if err := d.Set("commitment", flattenComputeReservationCommitment(res["commitment"], d)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}
	if err := d.Set("specific_reservation_required", flattenComputeReservationSpecificReservationRequired(res["specificReservationRequired"], d)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}
	if err := d.Set("status", flattenComputeReservationStatus(res["status"], d)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}
	if err := d.Set("specific_reservation", flattenComputeReservationSpecificReservation(res["specificReservation"], d)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}
	if err := d.Set("zone", flattenComputeReservationZone(res["zone"], d)); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Reservation: %s", err)
	}

	return nil
}

func resourceComputeReservationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/reservations/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Reservation %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Reservation")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting Reservation",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Reservation %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeReservationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/reservations/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/zones/{{zone}}/reservations/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeReservationCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeReservationDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeReservationName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeReservationCommitment(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeReservationSpecificReservationRequired(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeReservationStatus(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeReservationSpecificReservation(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["count"] =
		flattenComputeReservationSpecificReservationCount(original["count"], d)
	transformed["in_use_count"] =
		flattenComputeReservationSpecificReservationInUseCount(original["inUseCount"], d)
	transformed["instance_properties"] =
		flattenComputeReservationSpecificReservationInstanceProperties(original["instanceProperties"], d)
	return []interface{}{transformed}
}
func flattenComputeReservationSpecificReservationCount(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeReservationSpecificReservationInUseCount(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeReservationSpecificReservationInstanceProperties(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["machine_type"] =
		flattenComputeReservationSpecificReservationInstancePropertiesMachineType(original["machineType"], d)
	transformed["min_cpu_platform"] =
		flattenComputeReservationSpecificReservationInstancePropertiesMinCpuPlatform(original["minCpuPlatform"], d)
	transformed["guest_accelerators"] =
		flattenComputeReservationSpecificReservationInstancePropertiesGuestAccelerators(original["guestAccelerators"], d)
	transformed["local_ssds"] =
		flattenComputeReservationSpecificReservationInstancePropertiesLocalSsds(original["localSsds"], d)
	return []interface{}{transformed}
}
func flattenComputeReservationSpecificReservationInstancePropertiesMachineType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeReservationSpecificReservationInstancePropertiesMinCpuPlatform(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeReservationSpecificReservationInstancePropertiesGuestAccelerators(v interface{}, d *schema.ResourceData) interface{} {
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
			"accelerator_type":  flattenComputeReservationSpecificReservationInstancePropertiesGuestAcceleratorsAcceleratorType(original["acceleratorType"], d),
			"accelerator_count": flattenComputeReservationSpecificReservationInstancePropertiesGuestAcceleratorsAcceleratorCount(original["acceleratorCount"], d),
		})
	}
	return transformed
}
func flattenComputeReservationSpecificReservationInstancePropertiesGuestAcceleratorsAcceleratorType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeReservationSpecificReservationInstancePropertiesGuestAcceleratorsAcceleratorCount(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeReservationSpecificReservationInstancePropertiesLocalSsds(v interface{}, d *schema.ResourceData) interface{} {
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
			"interface":    flattenComputeReservationSpecificReservationInstancePropertiesLocalSsdsInterface(original["interface"], d),
			"disk_size_gb": flattenComputeReservationSpecificReservationInstancePropertiesLocalSsdsDiskSizeGb(original["diskSizeGb"], d),
		})
	}
	return transformed
}
func flattenComputeReservationSpecificReservationInstancePropertiesLocalSsdsInterface(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeReservationSpecificReservationInstancePropertiesLocalSsdsDiskSizeGb(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeReservationZone(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeReservationDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeReservationName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeReservationSpecificReservationRequired(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeReservationSpecificReservation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCount, err := expandComputeReservationSpecificReservationCount(original["count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCount); val.IsValid() && !isEmptyValue(val) {
		transformed["count"] = transformedCount
	}

	transformedInUseCount, err := expandComputeReservationSpecificReservationInUseCount(original["in_use_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInUseCount); val.IsValid() && !isEmptyValue(val) {
		transformed["inUseCount"] = transformedInUseCount
	}

	transformedInstanceProperties, err := expandComputeReservationSpecificReservationInstanceProperties(original["instance_properties"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstanceProperties); val.IsValid() && !isEmptyValue(val) {
		transformed["instanceProperties"] = transformedInstanceProperties
	}

	return transformed, nil
}

func expandComputeReservationSpecificReservationCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeReservationSpecificReservationInUseCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeReservationSpecificReservationInstanceProperties(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMachineType, err := expandComputeReservationSpecificReservationInstancePropertiesMachineType(original["machine_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMachineType); val.IsValid() && !isEmptyValue(val) {
		transformed["machineType"] = transformedMachineType
	}

	transformedMinCpuPlatform, err := expandComputeReservationSpecificReservationInstancePropertiesMinCpuPlatform(original["min_cpu_platform"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinCpuPlatform); val.IsValid() && !isEmptyValue(val) {
		transformed["minCpuPlatform"] = transformedMinCpuPlatform
	}

	transformedGuestAccelerators, err := expandComputeReservationSpecificReservationInstancePropertiesGuestAccelerators(original["guest_accelerators"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGuestAccelerators); val.IsValid() && !isEmptyValue(val) {
		transformed["guestAccelerators"] = transformedGuestAccelerators
	}

	transformedLocalSsds, err := expandComputeReservationSpecificReservationInstancePropertiesLocalSsds(original["local_ssds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocalSsds); val.IsValid() && !isEmptyValue(val) {
		transformed["localSsds"] = transformedLocalSsds
	}

	return transformed, nil
}

func expandComputeReservationSpecificReservationInstancePropertiesMachineType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeReservationSpecificReservationInstancePropertiesMinCpuPlatform(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeReservationSpecificReservationInstancePropertiesGuestAccelerators(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAcceleratorType, err := expandComputeReservationSpecificReservationInstancePropertiesGuestAcceleratorsAcceleratorType(original["accelerator_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAcceleratorType); val.IsValid() && !isEmptyValue(val) {
			transformed["acceleratorType"] = transformedAcceleratorType
		}

		transformedAcceleratorCount, err := expandComputeReservationSpecificReservationInstancePropertiesGuestAcceleratorsAcceleratorCount(original["accelerator_count"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAcceleratorCount); val.IsValid() && !isEmptyValue(val) {
			transformed["acceleratorCount"] = transformedAcceleratorCount
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeReservationSpecificReservationInstancePropertiesGuestAcceleratorsAcceleratorType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeReservationSpecificReservationInstancePropertiesGuestAcceleratorsAcceleratorCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeReservationSpecificReservationInstancePropertiesLocalSsds(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedInterface, err := expandComputeReservationSpecificReservationInstancePropertiesLocalSsdsInterface(original["interface"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedInterface); val.IsValid() && !isEmptyValue(val) {
			transformed["interface"] = transformedInterface
		}

		transformedDiskSizeGb, err := expandComputeReservationSpecificReservationInstancePropertiesLocalSsdsDiskSizeGb(original["disk_size_gb"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDiskSizeGb); val.IsValid() && !isEmptyValue(val) {
			transformed["diskSizeGb"] = transformedDiskSizeGb
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeReservationSpecificReservationInstancePropertiesLocalSsdsInterface(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeReservationSpecificReservationInstancePropertiesLocalSsdsDiskSizeGb(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeReservationZone(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}
