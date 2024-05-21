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

package compute

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

func ResourceComputeStoragePool() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeStoragePoolCreate,
		Read:   resourceComputeStoragePoolRead,
		Update: resourceComputeStoragePoolUpdate,
		Delete: resourceComputeStoragePoolDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeStoragePoolImport,
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
			"pool_provisioned_capacity_gb": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Size, in GiB, of the storage pool. Choose between 10,240 and 1,048,576 GiB.`,
			},
			"pool_provisioned_throughput": {
				Type:     schema.TypeInt,
				Required: true,
				Description: `Provisioned throughput of the storage pool. For hyperdisk-balanced storage pool type,
provision between 0 and 10,240, must be a multiple of 1,024 MB/s. For hyperdisk-throughput
storage pool type, provision between 100 and 180, must be a multiple of 10 MB/s.`,
			},
			"storage_pool_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareResourceNames,
				Description: `URL of the storage pool type resource describing which storage pool type to use to
create the storage pool. Provide this when creating the storage pool.`,
			},
			"capacity_provisioning_type": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ADVANCED", "STANDARD", ""}),
				Description:  `Provisioning type of the byte capacity of the pool. Possible values: ["ADVANCED", "STANDARD"]`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `An optional description of this resource. Provide this property when
you create the resource.`,
			},
			"performance_provisioning_type": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ADVANCED", "STANDARD", ""}),
				Description: `Provisioning type of the performance-related parameters of the pool,
such as throughput and IOPS. Possible values: ["ADVANCED", "STANDARD"]`,
			},
			"pool_provisioned_iops": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
				Description: `Provisioned IOPS of the storage pool. Only relevant if the storage pool type is
hyperdisk-balanced. Provision between 0 and 40,000, must be a multiple of 10,000.`,
			},
			"zone": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `A reference to the zone where the storage pool resides.`,
			},
			"storage_pool_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The unique identifier for the resource. This identifier is defined by the server.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"disk_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Number of disks used.`,
			},
			"last_resize_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Timestamp of the last successful resize in RFC3339 text format.`,
			},
			"max_total_provisioned_disk_capacity_gb": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Maximum allowed aggregate disk size in gigabytes.`,
			},
			"pool_used_capacity_bytes": {
				Type:     schema.TypeInt,
				Computed: true,
				Description: `Space used by data stored in disks within the storage pool (in bytes).
This will reflect the total number of bytes written to the disks in the pool,
in contrast to the capacity of those disks.`,
			},
			"pool_used_iops": {
				Type:     schema.TypeInt,
				Computed: true,
				Description: `Sum of all the disks' provisioned IOPS, minus some amount that is allowed
per disk that is not counted towards pool's IOPS capacity.`,
			},
			"pool_used_throughput": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Sum of all the disks' provisioned throughput in MB/s.`,
			},
			"pool_user_written_bytes": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Amount of data written into the pool, before it is compacted.`,
			},
			"total_provisioned_disk_capacity_gb": {
				Type:     schema.TypeInt,
				Computed: true,
				Description: `Sum of all the capacity provisioned in disks in this storage pool.
A disk's provisioned capacity is the same as its total capacity.`,
			},
			"total_provisioned_disk_iops": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Sum of all the disks' provisioned IOPS.`,
			},
			"total_provisioned_disk_throughput": {
				Type:     schema.TypeInt,
				Computed: true,
				Description: `Sum of all the disks' provisioned throughput in MB/s, minus
some amount that is allowed per disk that is not counted
toward pool's throughput capacity.`,
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

func resourceComputeStoragePoolCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeStoragePoolName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeStoragePoolDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	poolProvisionedCapacityGbProp, err := expandComputeStoragePoolPoolProvisionedCapacityGb(d.Get("pool_provisioned_capacity_gb"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("pool_provisioned_capacity_gb"); !tpgresource.IsEmptyValue(reflect.ValueOf(poolProvisionedCapacityGbProp)) && (ok || !reflect.DeepEqual(v, poolProvisionedCapacityGbProp)) {
		obj["poolProvisionedCapacityGb"] = poolProvisionedCapacityGbProp
	}
	poolProvisionedIopsProp, err := expandComputeStoragePoolPoolProvisionedIops(d.Get("pool_provisioned_iops"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("pool_provisioned_iops"); !tpgresource.IsEmptyValue(reflect.ValueOf(poolProvisionedIopsProp)) && (ok || !reflect.DeepEqual(v, poolProvisionedIopsProp)) {
		obj["poolProvisionedIops"] = poolProvisionedIopsProp
	}
	poolProvisionedThroughputProp, err := expandComputeStoragePoolPoolProvisionedThroughput(d.Get("pool_provisioned_throughput"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("pool_provisioned_throughput"); !tpgresource.IsEmptyValue(reflect.ValueOf(poolProvisionedThroughputProp)) && (ok || !reflect.DeepEqual(v, poolProvisionedThroughputProp)) {
		obj["poolProvisionedThroughput"] = poolProvisionedThroughputProp
	}
	storagePoolTypeProp, err := expandComputeStoragePoolStoragePoolType(d.Get("storage_pool_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("storage_pool_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(storagePoolTypeProp)) && (ok || !reflect.DeepEqual(v, storagePoolTypeProp)) {
		obj["storagePoolType"] = storagePoolTypeProp
	}
	capacityProvisioningTypeProp, err := expandComputeStoragePoolCapacityProvisioningType(d.Get("capacity_provisioning_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("capacity_provisioning_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(capacityProvisioningTypeProp)) && (ok || !reflect.DeepEqual(v, capacityProvisioningTypeProp)) {
		obj["capacityProvisioningType"] = capacityProvisioningTypeProp
	}
	performanceProvisioningTypeProp, err := expandComputeStoragePoolPerformanceProvisioningType(d.Get("performance_provisioning_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("performance_provisioning_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(performanceProvisioningTypeProp)) && (ok || !reflect.DeepEqual(v, performanceProvisioningTypeProp)) {
		obj["performanceProvisioningType"] = performanceProvisioningTypeProp
	}
	zoneProp, err := expandComputeStoragePoolZone(d.Get("zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/storagePools")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new StoragePool: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for StoragePool: %s", err)
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
		return fmt.Errorf("Error creating StoragePool: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/zones/{{zone}}/storagePools/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating StoragePool", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create StoragePool: %s", err)
	}

	log.Printf("[DEBUG] Finished creating StoragePool %q: %#v", d.Id(), res)

	return resourceComputeStoragePoolRead(d, meta)
}

func resourceComputeStoragePoolRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/storagePools/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for StoragePool: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeStoragePool %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}

	if err := d.Set("storage_pool_id", flattenComputeStoragePoolStoragePoolId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeStoragePoolCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("name", flattenComputeStoragePoolName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("description", flattenComputeStoragePoolDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("pool_provisioned_capacity_gb", flattenComputeStoragePoolPoolProvisionedCapacityGb(res["poolProvisionedCapacityGb"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("pool_provisioned_iops", flattenComputeStoragePoolPoolProvisionedIops(res["poolProvisionedIops"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("pool_provisioned_throughput", flattenComputeStoragePoolPoolProvisionedThroughput(res["poolProvisionedThroughput"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("last_resize_timestamp", flattenComputeStoragePoolLastResizeTimestamp(res["lastResizeTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("disk_count", flattenComputeStoragePoolDiskCount(res["diskCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("pool_used_capacity_bytes", flattenComputeStoragePoolPoolUsedCapacityBytes(res["poolUsedCapacityBytes"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("pool_user_written_bytes", flattenComputeStoragePoolPoolUserWrittenBytes(res["poolUserWrittenBytes"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("total_provisioned_disk_capacity_gb", flattenComputeStoragePoolTotalProvisionedDiskCapacityGb(res["totalProvisionedDiskCapacityGb"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("max_total_provisioned_disk_capacity_gb", flattenComputeStoragePoolMaxTotalProvisionedDiskCapacityGb(res["maxTotalProvisionedDiskCapacityGb"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("pool_used_iops", flattenComputeStoragePoolPoolUsedIops(res["poolUsedIops"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("total_provisioned_disk_iops", flattenComputeStoragePoolTotalProvisionedDiskIops(res["totalProvisionedDiskIops"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("pool_used_throughput", flattenComputeStoragePoolPoolUsedThroughput(res["poolUsedThroughput"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("total_provisioned_disk_throughput", flattenComputeStoragePoolTotalProvisionedDiskThroughput(res["totalProvisionedDiskThroughput"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("storage_pool_type", flattenComputeStoragePoolStoragePoolType(res["storagePoolType"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("capacity_provisioning_type", flattenComputeStoragePoolCapacityProvisioningType(res["capacityProvisioningType"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("performance_provisioning_type", flattenComputeStoragePoolPerformanceProvisioningType(res["performanceProvisioningType"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("zone", flattenComputeStoragePoolZone(res["zone"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading StoragePool: %s", err)
	}

	return nil
}

func resourceComputeStoragePoolUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for StoragePool: %s", err)
	}
	billingProject = project

	d.Partial(true)

	if d.HasChange("pool_provisioned_capacity_gb") {
		obj := make(map[string]interface{})

		poolProvisionedCapacityGbProp, err := expandComputeStoragePoolPoolProvisionedCapacityGb(d.Get("pool_provisioned_capacity_gb"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("pool_provisioned_capacity_gb"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, poolProvisionedCapacityGbProp)) {
			obj["poolProvisionedCapacityGb"] = poolProvisionedCapacityGbProp
		}

		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/storagePools/{{name}}?updateMask=poolProvisionedCapacityGb")
		if err != nil {
			return err
		}

		headers := make(http.Header)

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
			Headers:   headers,
		})
		if err != nil {
			return fmt.Errorf("Error updating StoragePool %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating StoragePool %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating StoragePool", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("pool_provisioned_iops") {
		obj := make(map[string]interface{})

		poolProvisionedIopsProp, err := expandComputeStoragePoolPoolProvisionedIops(d.Get("pool_provisioned_iops"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("pool_provisioned_iops"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, poolProvisionedIopsProp)) {
			obj["poolProvisionedIops"] = poolProvisionedIopsProp
		}

		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/storagePools/{{name}}?updateMask=poolProvisionedIops")
		if err != nil {
			return err
		}

		headers := make(http.Header)

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
			Headers:   headers,
		})
		if err != nil {
			return fmt.Errorf("Error updating StoragePool %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating StoragePool %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating StoragePool", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("pool_provisioned_throughput") {
		obj := make(map[string]interface{})

		poolProvisionedThroughputProp, err := expandComputeStoragePoolPoolProvisionedThroughput(d.Get("pool_provisioned_throughput"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("pool_provisioned_throughput"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, poolProvisionedThroughputProp)) {
			obj["poolProvisionedThroughput"] = poolProvisionedThroughputProp
		}

		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/storagePools/{{name}}?updateMask=poolProvisionedThroughput")
		if err != nil {
			return err
		}

		headers := make(http.Header)

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
			Headers:   headers,
		})
		if err != nil {
			return fmt.Errorf("Error updating StoragePool %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating StoragePool %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating StoragePool", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceComputeStoragePoolRead(d, meta)
}

func resourceComputeStoragePoolDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for StoragePool: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/storagePools/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting StoragePool %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "StoragePool")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting StoragePool", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting StoragePool %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeStoragePoolImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/storagePools/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<zone>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/zones/{{zone}}/storagePools/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeStoragePoolStoragePoolId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeStoragePoolCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeStoragePoolName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeStoragePoolDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeStoragePoolPoolProvisionedCapacityGb(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeStoragePoolPoolProvisionedIops(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeStoragePoolPoolProvisionedThroughput(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeStoragePoolLastResizeTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeStoragePoolDiskCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeStoragePoolPoolUsedCapacityBytes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeStoragePoolPoolUserWrittenBytes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeStoragePoolTotalProvisionedDiskCapacityGb(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeStoragePoolMaxTotalProvisionedDiskCapacityGb(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeStoragePoolPoolUsedIops(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeStoragePoolTotalProvisionedDiskIops(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeStoragePoolPoolUsedThroughput(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeStoragePoolTotalProvisionedDiskThroughput(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeStoragePoolStoragePoolType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenComputeStoragePoolCapacityProvisioningType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeStoragePoolPerformanceProvisioningType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeStoragePoolZone(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func expandComputeStoragePoolName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeStoragePoolDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeStoragePoolPoolProvisionedCapacityGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeStoragePoolPoolProvisionedIops(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeStoragePoolPoolProvisionedThroughput(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeStoragePoolStoragePoolType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseZonalFieldValue("storagePoolTypes", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for storage_pool_type: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeStoragePoolCapacityProvisioningType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeStoragePoolPerformanceProvisioningType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeStoragePoolZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}
