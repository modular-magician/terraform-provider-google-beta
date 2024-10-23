// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
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

package storage

import (
	"fmt"
	"log"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/gammazero/workerpool"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/storage/v1"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceStorageFolder() *schema.Resource {
	return &schema.Resource{
		Create: resourceStorageFolderCreate,
		Read:   resourceStorageFolderRead,
		Update: resourceStorageFolderUpdate,
		Delete: resourceStorageFolderDelete,

		Importer: &schema.ResourceImporter{
			State: resourceStorageFolderImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the bucket that contains the folder.`,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidateRegexp(`/$`),
				Description: `The name of the folder expressed as a path. Must include
trailing '/'. For example, 'example_dir/example_dir2/'.`,
			},
			"recursive": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Set to true if parent folder and child folders need to be created`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp at which this folder was created.`,
			},
			"metageneration": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The metadata generation of the folder.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp at which this folder was most recently updated.`,
			},
			"force_destroy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Set to true if force destroy folder`,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceStorageFolderCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	bucket := d.Get("bucket").(string)
	recursive := d.Get("recursive").(bool)
	forceDestroy := d.Get("forces_destroy").(bool)
	folder := &storage.Folder{
		Name: d.Get("name").(string),
	}

	var res *storage.Folder

	err = transport_tpg.Retry(transport_tpg.RetryOptions{
		RetryFunc: func() error {
			res, err = config.NewStorageClient(userAgent).Folders.Insert(bucket, folder).Recursive(recursive).Do()
			return err
		},
		Timeout:              d.Timeout(schema.TimeoutCreate),
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.Is429RetryableQuotaError},
	})

	log.Printf("[DEBUG] Creating new Folder: %#v", folder.Name)
	if err != nil {
		return fmt.Errorf("Error creating Folder: %s", err)
	}

	// Store the ID now
	d.SetId(res.Id)
	if err := d.Set("recursive", recursive); err != nil {
		return fmt.Errorf("Error setting recursive flag: %s", err)
	}

	if err := d.Set("force_destroy", forceDestroy); err != nil {
		return fmt.Errorf("Error setting force destroy flag: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Folder %q: %#v", d.Id(), res)

	return resourceStorageFolderRead(d, meta)
}

func resourceStorageFolderRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}
	bucket := d.Get("bucket").(string)
	folder := d.Get("name").(string)
	var res *storage.Folder
	err = transport_tpg.Retry(transport_tpg.RetryOptions{
		RetryFunc: func() (operr error) {
			var retryErr error
			res, retryErr = config.NewStorageClient(userAgent).Folders.Get(bucket, folder).Do()
			return retryErr
		},
		Timeout:              d.Timeout(schema.TimeoutRead),
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsNotFoundRetryableError("folder read")},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("StorageFolder %q", d.Id()))
	}

	if err := d.Set("create_time", flattenStorageFolderCreateTime(res.CreateTime, d, config)); err != nil {
		return fmt.Errorf("Error reading Folder: %s", err)
	}
	if err := d.Set("update_time", flattenStorageFolderUpdateTime(res.UpdateTime, d, config)); err != nil {
		return fmt.Errorf("Error reading Folder: %s", err)
	}
	if err := d.Set("metageneration", flattenStorageFolderMetageneration(res.Metageneration, d, config)); err != nil {
		return fmt.Errorf("Error reading Folder: %s", err)
	}
	if err := d.Set("bucket", flattenStorageFolderBucket(res.Bucket, d, config)); err != nil {
		return fmt.Errorf("Error reading Folder: %s", err)
	}
	if err := d.Set("name", flattenStorageFolderName(res.Name, d, config)); err != nil {
		return fmt.Errorf("Error reading Folder: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res.SelfLink)); err != nil {
		return fmt.Errorf("Error reading Folder: %s", err)
	}

	return nil
}

func resourceStorageFolderUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}
	oldName, newName := d.GetChange("name")

	bucket := d.Get("bucket").(string)

	if oldName != newName {
		objectsService := storage.NewFoldersService(config.NewStorageClientWithTimeoutOverride(userAgent, d.Timeout(schema.TimeoutUpdate)))

		rename := objectsService.Rename(bucket, oldName.(string), newName.(string))
		_, err = rename.Do()

		if err != nil {
			return fmt.Errorf("Error updating folder %s: %s", oldName, err)
		}
	}

	return resourceStorageFolderRead(d, meta)
}

func resourceStorageFolderDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	bucket := d.Get("bucket").(string)
	name := d.Get("name").(string)
	re := regexp.MustCompile(`(\w*\/)`)
	folderName := re.FindAllString(name, -1)
	var listError, deleteObjectError error
	for deleteObjectError == nil {
		res, err := config.NewStorageClient(userAgent).Objects.List(bucket + folderName[0]).Versions(true).Do()
		if err != nil {
			log.Printf("Error listing contents of bucket %s: %v", bucket, err)
			// If we can't list the contents, try deleting the bucket anyway in case it's empty
			listError = err
			break
		}

		if len(res.Items) == 0 {
			break // 0 items, folder empty
		}

		if !d.Get("force_destroy").(bool) {
			deleteErr := fmt.Errorf("Error trying to delete folder %s containing objects without `force_destroy` set to true", bucket)
			log.Printf("Error! %s : %s\n\n", bucket, deleteErr)
			return deleteErr
		}
		// GCS requires that a bucket be empty (have no objects or object
		// versions) before it can be deleted.
		log.Printf("[DEBUG] GCS Folder attempting to forceDestroy\n\n")

		// Create a workerpool for parallel deletion of resources. In the
		// future, it would be great to expose Terraform's global parallelism
		// flag here, but that's currently reserved for core use. Testing
		// shows that NumCPUs-1 is the most performant on average networks.
		//
		// The challenge with making this user-configurable is that the
		// configuration would reside in the Terraform configuration file,
		// decreasing its portability. Ideally we'd want this to connect to
		// Terraform's top-level -parallelism flag, but that's not plumbed nor
		// is it scheduled to be plumbed to individual providers.
		wp := workerpool.New(runtime.NumCPU() - 1)

		for _, object := range res.Items {
			log.Printf("[DEBUG] Found %s", object.Name)
			object := object

			wp.Submit(func() {
				log.Printf("[TRACE] Attempting to delete %s", object.Name)
				if err := config.NewStorageClient(userAgent).Objects.Delete(bucket, object.Name).Generation(object.Generation).Do(); err != nil {
					deleteObjectError = err
					log.Printf("[ERR] Failed to delete storage object %s: %s", object.Name, err)
				} else {
					log.Printf("[TRACE] Successfully deleted %s", object.Name)
				}
			})
		}

		// Wait for everything to finish.
		wp.StopWait()
	}

	if gerr, ok := err.(*googleapi.Error); ok && gerr.Code == 409 && strings.Contains(gerr.Message, "not empty") && listError != nil {
		return fmt.Errorf("could not delete non-empty folder due to error when listing contents: %v", listError)
	}

	log.Printf("[DEBUG] folder names to delete: %#v", folderName)
	foldersList, err := config.NewStorageClient(userAgent).Folders.List(bucket).Prefix(folderName[0]).Do()
	if err != nil {
		return err
	}
	items := foldersList.Items
	for i := len(items) - 1; i >= 0; i-- {
		err = transport_tpg.Retry(transport_tpg.RetryOptions{
			RetryFunc: func() error {
				err = config.NewStorageClient(userAgent).Folders.Delete(bucket, items[i].Name).Do()
				return err
			},
			Timeout:              d.Timeout(schema.TimeoutDelete),
			ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.Is429RetryableQuotaError},
		})
		if err != nil {
			return err
		}
	}

	log.Printf("[DEBUG] Finished deleting Folder %q: %#v", d.Id(), name)
	return nil
}

func resourceStorageFolderImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := d.Set("force_destroy", false); err != nil {
		return nil, fmt.Errorf("Error setting force_destroy: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenStorageFolderCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenStorageFolderUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenStorageFolderMetageneration(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenStorageFolderBucket(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenStorageFolderName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenStorageFolderRecursive(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandStorageFolderBucket(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageFolderName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageFolderRecursive(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
