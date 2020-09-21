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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudIdentityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudIdentityGroupCreate,
		Read:   resourceCloudIdentityGroupRead,
		Update: resourceCloudIdentityGroupUpdate,
		Delete: resourceCloudIdentityGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudIdentityGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"group_key": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: `EntityKey of the Group.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `The ID of the entity.

For Google-managed entities, the id must be the email address of an existing
group or user.

For external-identity-mapped entities, the id must be a string conforming
to the Identity Source's requirements.

Must be unique within a namespace.`,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `The namespace in which the entity exists.

If not specified, the EntityKey represents a Google-managed entity
such as a Google user or a Google Group.

If specified, the EntityKey represents an external-identity-mapped group.
The namespace must correspond to an identity source created in Admin Console
and must be in the form of 'identitysources/{identity_source_id}'.`,
						},
					},
				},
			},
			"labels": {
				Type:     schema.TypeMap,
				Required: true,
				ForceNew: true,
				Description: `The labels that apply to the Group.

Must not contain more than one entry. Must contain the entry
'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
'system/groups/external': '' if the Group is an external-identity-mapped group.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"parent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The resource name of the entity under which this Group resides in the
Cloud Identity resource hierarchy.

Must be of the form identitysources/{identity_source_id} for external-identity-mapped
groups or customers/{customer_id} for Google Groups.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `An extended description to help users determine the purpose of a Group.
Must not be longer than 4,096 characters.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The display name of the Group.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the Group was created.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Resource name of the Group in the format: groups/{group_id}, where group_id
is the unique ID assigned to the Group.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the Group was last updated.`,
			},
		},
	}
}

func resourceCloudIdentityGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	groupKeyProp, err := expandCloudIdentityGroupGroupKey(d.Get("group_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("group_key"); !isEmptyValue(reflect.ValueOf(groupKeyProp)) && (ok || !reflect.DeepEqual(v, groupKeyProp)) {
		obj["groupKey"] = groupKeyProp
	}
	parentProp, err := expandCloudIdentityGroupParent(d.Get("parent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	displayNameProp, err := expandCloudIdentityGroupDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandCloudIdentityGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandCloudIdentityGroupLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := replaceVars(d, config, "{{CloudIdentityBasePath}}groups?initialGroupConfig=EMPTY")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Group: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Group: %s", err)
	}
	if err := d.Set("name", flattenCloudIdentityGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Group %q: %#v", d.Id(), res)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		respBody, ok := res["response"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}

		name, ok = respBody.(map[string]interface{})["name"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}
	}
	if err := d.Set("name", name.(string)); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name.(string))

	return resourceCloudIdentityGroupRead(d, meta)
}

func resourceCloudIdentityGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{CloudIdentityBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("CloudIdentityGroup %q", d.Id()))
	}

	if err := d.Set("name", flattenCloudIdentityGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("group_key", flattenCloudIdentityGroupGroupKey(res["groupKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("parent", flattenCloudIdentityGroupParent(res["parent"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("display_name", flattenCloudIdentityGroupDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("description", flattenCloudIdentityGroupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("create_time", flattenCloudIdentityGroupCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("update_time", flattenCloudIdentityGroupUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("labels", flattenCloudIdentityGroupLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}

	return nil
}

func resourceCloudIdentityGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	billingProject := ""

	obj := make(map[string]interface{})
	displayNameProp, err := expandCloudIdentityGroupDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandCloudIdentityGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	url, err := replaceVars(d, config, "{{CloudIdentityBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Group %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Group %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Group %q: %#v", d.Id(), res)
	}

	return resourceCloudIdentityGroupRead(d, meta)
}

func resourceCloudIdentityGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	billingProject := ""

	url, err := replaceVars(d, config, "{{CloudIdentityBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Group %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Group")
	}

	log.Printf("[DEBUG] Finished deleting Group %q: %#v", d.Id(), res)
	return nil
}

func resourceCloudIdentityGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	name := d.Get("name").(string)

	if err := d.Set("name", name); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name)
	return []*schema.ResourceData{d}, nil
}

func flattenCloudIdentityGroupName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdentityGroupGroupKey(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["id"] =
		flattenCloudIdentityGroupGroupKeyId(original["id"], d, config)
	transformed["namespace"] =
		flattenCloudIdentityGroupGroupKeyNamespace(original["namespace"], d, config)
	return []interface{}{transformed}
}
func flattenCloudIdentityGroupGroupKeyId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdentityGroupGroupKeyNamespace(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdentityGroupParent(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdentityGroupDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdentityGroupDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdentityGroupCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdentityGroupUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdentityGroupLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandCloudIdentityGroupGroupKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedId, err := expandCloudIdentityGroupGroupKeyId(original["id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedId); val.IsValid() && !isEmptyValue(val) {
		transformed["id"] = transformedId
	}

	transformedNamespace, err := expandCloudIdentityGroupGroupKeyNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !isEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	return transformed, nil
}

func expandCloudIdentityGroupGroupKeyId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupGroupKeyNamespace(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupParent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
