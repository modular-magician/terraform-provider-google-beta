package google

import (
	"fmt"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"log"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	compute "google.golang.org/api/compute/v0.beta"
)

func ResourceComputeDiskAsyncReplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceDiskAsyncReplicationCreate,
		Read:   resourceDiskAsyncReplicationRead,
		Delete: resourceDiskAsyncReplicationDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"primary_disk": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      `Primary disk for asynchronous replication.`,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"secondary_disk": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: `Secondary disk for asynchronous replication.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disk": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							Description:      `Secondary disk for asynchronous replication.`,
							DiffSuppressFunc: compareSelfLinkOrResourceName,
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Output-only. Status of replication on the secondary disk.`,
						},
					},
				},
			},
		},
		UseJSONNumber: true,
	}
}

func asyncReplicationGetComputeClient(d *schema.ResourceData, meta interface{}) (*compute.Service, error) {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return nil, err
	}

	clientCompute := config.NewComputeClient(userAgent)
	return clientCompute, nil
}

func asyncReplicationGetDiskFromConfig(disk string, d *schema.ResourceData, meta interface{}) (zv *ZonalFieldValue, rv *RegionalFieldValue, resourceId string, err error) {
	config := meta.(*transport_tpg.Config)

	var zonalMatch bool
	zonalMatch, err = regexp.MatchString(fmt.Sprintf(zonalLinkBasePattern, "disks"), disk)
	if err != nil {
		return
	}
	zv, parseErr := ParseDiskFieldValue(disk, d, config)
	if !zonalMatch || parseErr != nil {
		rv, err = ParseRegionDiskFieldValue(disk, d, config)
		if err != nil {
			return
		}
		var regionalMatch bool
		regionalMatch, err = regexp.MatchString(fmt.Sprintf(regionalLinkBasePattern, "disks"), disk)
		if !regionalMatch || err != nil {
			err = fmt.Errorf("regional disk expected: %s", disk)
			return
		}
		resourceId = fmt.Sprintf(regionalLinkTemplate, rv.Project, rv.Region, "disks", rv.Name)
	} else {
		resourceId = fmt.Sprintf(zonalLinkTemplate, zv.Project, zv.Zone, "disks", zv.Name)
	}
	return
}

func asyncReplicationGetDiskStatus(client *compute.Service, zv *ZonalFieldValue, rv *RegionalFieldValue) (diskStatus *compute.Disk, err error) {
	if rv == nil { // Zonal disk
		diskStatus, err = client.Disks.Get(zv.Project, zv.Zone, zv.Name).Do()
		log.Printf("[DEBUG] Get disk zones/%s/%s: %v", zv.Zone, zv.Name, diskStatus)
	} else {
		diskStatus, err = client.RegionDisks.Get(rv.Project, rv.Region, rv.Name).Do()
		log.Printf("[DEBUG] Get disk regions/%s/%s: %v", rv.Region, rv.Name, diskStatus)
	}
	return
}

func resourceDiskAsyncReplicationCreate(d *schema.ResourceData, meta interface{}) error {
	clientCompute, err := asyncReplicationGetComputeClient(d, meta)
	if err != nil {
		return err
	}

	zv, rv, resourceId, err := asyncReplicationGetDiskFromConfig(d.Get("primary_disk").(string), d, meta)
	if err != nil {
		return err
	}

	secondaryDiskList := d.Get("secondary_disk").([]interface{})
	secondaryDiskMap := secondaryDiskList[0].(map[string]interface{})
	secondaryDisk := secondaryDiskMap["disk"].(string)
	if rv == nil { // Zonal disk
		replicationRequest := compute.DisksStartAsyncReplicationRequest{
			AsyncSecondaryDisk: secondaryDisk,
		}
		_, err = clientCompute.Disks.StartAsyncReplication(zv.Project, zv.Zone, zv.Name, &replicationRequest).Do()
		if err != nil {
			return err
		}
	} else {
		replicationRequest := compute.RegionDisksStartAsyncReplicationRequest{
			AsyncSecondaryDisk: secondaryDisk,
		}
		_, err = clientCompute.RegionDisks.StartAsyncReplication(rv.Project, rv.Region, rv.Name, &replicationRequest).Do()
		if err != nil {
			return err
		}
	}
	err = resource.Retry(time.Minute*time.Duration(5), func() *resource.RetryError {
		diskStatus, err := asyncReplicationGetDiskStatus(clientCompute, zv, rv)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		if diskStatus.ResourceStatus == nil {
			return resource.NonRetryableError(fmt.Errorf("no resource status for disk: %s", resourceId))
		}
		if secondaryState, ok := diskStatus.ResourceStatus.AsyncSecondaryDisks[secondaryDisk]; ok {
			if secondaryState.State != "ACTIVE" {
				time.Sleep(5 * time.Second)
				return resource.RetryableError(fmt.Errorf("secondary disk %s state (%s) is not: ACTIVE", secondaryDisk, secondaryState))
			}
			return nil
		}
		time.Sleep(5 * time.Second)
		return resource.RetryableError(fmt.Errorf("secondary disk %s state not available", secondaryDisk))
	})
	if err != nil {
		return err
	}
	d.SetId(resourceId)
	return resourceDiskAsyncReplicationRead(d, meta)
}

func resourceDiskAsyncReplicationRead(d *schema.ResourceData, meta interface{}) error {
	clientCompute, err := asyncReplicationGetComputeClient(d, meta)
	if err != nil {
		return err
	}

	primaryDisk := d.Get("primary_disk").(string)
	if primaryDisk == "" {
		primaryDisk = d.Id()
		d.Set("primary_disk", primaryDisk)
	}

	zv, rv, resourceId, err := asyncReplicationGetDiskFromConfig(primaryDisk, d, meta)
	if err != nil {
		return err
	}

	diskStatus, err := asyncReplicationGetDiskStatus(clientCompute, zv, rv)
	if err != nil {
		return err
	}

	secondaryDisks := make([]map[string]string, 0)
	existingSecondaryDisks := make(map[string]bool, 0)
	for _, disk := range diskStatus.AsyncSecondaryDisks {
		secondaryDisk := make(map[string]string)

		_, _, resourceName, err := asyncReplicationGetDiskFromConfig(disk.AsyncReplicationDisk.Disk, d, meta)
		if err != nil {
			return err
		}

		if diskStatus.ResourceStatus == nil {
			return fmt.Errorf("no resource status for disk: %s", resourceId)
		}

		secondaryDisk["disk"] = resourceName
		existingSecondaryDisks[resourceName] = true
		if secondaryState, ok := diskStatus.ResourceStatus.AsyncSecondaryDisks[resourceName]; ok {
			// Note this might be other than ACTIVE or STOPPED, but we wait for proper state
			// on replication start/stop so it shouldnt affect Terraform
			log.Printf("[DEBUG] Secondary disk %s is in state: %s", resourceName, secondaryState.State)
			secondaryDisk["state"] = secondaryState.State
		}
		secondaryDisks = append(secondaryDisks, secondaryDisk)
	}

	log.Printf("[DEBUG] Secondary disks: %v", secondaryDisks)
	if err = d.Set("secondary_disk", secondaryDisks); err != nil {
		return fmt.Errorf("Error setting secondary_disk: %s", err)
	}
	d.SetId(resourceId)
	return nil
}

func resourceDiskAsyncReplicationDelete(d *schema.ResourceData, meta interface{}) error {
	clientCompute, err := asyncReplicationGetComputeClient(d, meta)
	if err != nil {
		return err
	}

	zv, rv, _, err := asyncReplicationGetDiskFromConfig(d.Get("primary_disk").(string), d, meta)
	if err != nil {
		return err
	}

	var replicationStopped bool = false
	secondaryDiskList := d.Get("secondary_disk").([]interface{})
	secondaryDiskMap := secondaryDiskList[0].(map[string]interface{})
	secondaryDisk := secondaryDiskMap["disk"].(string)
	_, _, resourceName, err := asyncReplicationGetDiskFromConfig(secondaryDisk, d, meta)
	if err != nil {
		return err
	}

	diskStatus, err := asyncReplicationGetDiskStatus(clientCompute, zv, rv)
	if err != nil {
		return err
	}

	if diskStatus.ResourceStatus == nil {
		// Nothing to do, replication not running
		return nil
	}

	if secondaryState, ok := diskStatus.ResourceStatus.AsyncSecondaryDisks[resourceName]; ok {
		if secondaryState.State != "STOPPED" {
			replicationStopped = true
			if rv == nil { // Zonal disk
				_, err = clientCompute.Disks.StopAsyncReplication(zv.Project, zv.Zone, zv.Name).Do()
				if err != nil {
					return err
				}
			} else {
				_, err = clientCompute.RegionDisks.StopAsyncReplication(rv.Project, rv.Region, rv.Name).Do()
				if err != nil {
					return err
				}
			}
			err = resource.Retry(time.Minute*time.Duration(5), func() *resource.RetryError {
				diskStatus, err := asyncReplicationGetDiskStatus(clientCompute, zv, rv)
				if err != nil {
					return resource.NonRetryableError(err)
				}
				if secondaryState, ok := diskStatus.ResourceStatus.AsyncSecondaryDisks[resourceName]; ok {
					if secondaryState.State != "STOPPED" {
						time.Sleep(5 * time.Second)
						return resource.RetryableError(fmt.Errorf("secondary disk %s state (%s) is not STOPPED", secondaryDisk, secondaryState))
					}
					return nil
				}
				return resource.NonRetryableError(fmt.Errorf("secondary disk %s state not available", secondaryDisk))
			})
			if err != nil {
				return err
			}
		}
	} else {
		return fmt.Errorf("could not find secondary disk: %s", secondaryDisk)
	}

	if replicationStopped {
		// Allow the replication to quiescence
		time.Sleep(5000 * time.Millisecond)
	}
	return nil
}
