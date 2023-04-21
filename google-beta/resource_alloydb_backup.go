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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceAlloydbBackup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlloydbBackupCreate,
		Read:   resourceAlloydbBackupRead,
		Update: resourceAlloydbBackupUpdate,
		Delete: resourceAlloydbBackupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAlloydbBackupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"backup_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the alloydb backup.`,
			},
			"cluster_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: ProjectNumberDiffSuppress,
				Description:      `The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location where the alloydb backup should reside.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `User-provided description of the backup.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `User-defined labels for the alloydb backup.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the Backup was created in UTC.`,
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `A hash of the resource.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}`,
			},
			"reconciling": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `If true, indicates that the service is actively updating the resource. This can happen due to user-triggered updates or system actions like failover or maintenance.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The current state of the backup.`,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the Backup was updated in UTC.`,
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

func resourceAlloydbBackupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	clusterNameProp, err := expandAlloydbBackupClusterName(d.Get("cluster_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cluster_name"); !isEmptyValue(reflect.ValueOf(clusterNameProp)) && (ok || !reflect.DeepEqual(v, clusterNameProp)) {
		obj["clusterName"] = clusterNameProp
	}
	labelsProp, err := expandAlloydbBackupLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	descriptionProp, err := expandAlloydbBackupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	obj, err = resourceAlloydbBackupEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{AlloydbBasePath}}projects/{{project}}/locations/{{location}}/backups?backupId={{backup_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Backup: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Backup: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Backup: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backups/{{backup_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = AlloydbOperationWaitTime(
		config, res, project, "Creating Backup", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Backup: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Backup %q: %#v", d.Id(), res)

	return resourceAlloydbBackupRead(d, meta)
}

func resourceAlloydbBackupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{AlloydbBasePath}}projects/{{project}}/locations/{{location}}/backups/{{backup_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Backup: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AlloydbBackup %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}

	if err := d.Set("name", flattenAlloydbBackupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("uid", flattenAlloydbBackupUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("cluster_name", flattenAlloydbBackupClusterName(res["clusterName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("labels", flattenAlloydbBackupLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("create_time", flattenAlloydbBackupCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("update_time", flattenAlloydbBackupUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("state", flattenAlloydbBackupState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("description", flattenAlloydbBackupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("reconciling", flattenAlloydbBackupReconciling(res["reconciling"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("etag", flattenAlloydbBackupEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}

	return nil
}

func resourceAlloydbBackupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Backup: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandAlloydbBackupLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	obj, err = resourceAlloydbBackupEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{AlloydbBasePath}}projects/{{project}}/locations/{{location}}/backups/{{backup_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Backup %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Backup %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Backup %q: %#v", d.Id(), res)
	}

	err = AlloydbOperationWaitTime(
		config, res, project, "Updating Backup", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceAlloydbBackupRead(d, meta)
}

func resourceAlloydbBackupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Backup: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{AlloydbBasePath}}projects/{{project}}/locations/{{location}}/backups/{{backup_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Backup %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Backup")
	}

	err = AlloydbOperationWaitTime(
		config, res, project, "Deleting Backup", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Backup %q: %#v", d.Id(), res)
	return nil
}

func resourceAlloydbBackupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/backups/(?P<backup_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<backup_id>[^/]+)",
		"(?P<location>[^/]+)/(?P<backup_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backups/{{backup_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAlloydbBackupName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupClusterName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupReconciling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandAlloydbBackupClusterName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbBackupLabels(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAlloydbBackupDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceAlloydbBackupEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// The only other available type is AUTOMATED which cannot be set manually
	obj["type"] = "ON_DEMAND"
	return obj, nil
}
