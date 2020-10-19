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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudIdentityGroupMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudIdentityGroupMembershipCreate,
		Read:   resourceCloudIdentityGroupMembershipRead,
		Update: resourceCloudIdentityGroupMembershipUpdate,
		Delete: resourceCloudIdentityGroupMembershipDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudIdentityGroupMembershipImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"group": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The name of the Group to create this membership in.`,
			},
			"roles": {
				Type:     schema.TypeList,
				Required: true,
				Description: `The MembershipRoles that apply to the Membership.
Must not contain duplicate MembershipRoles with the same name.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"OWNER", "MANAGER", "MEMBER"}, false),
							Description:  `The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER. Possible values: ["OWNER", "MANAGER", "MEMBER"]`,
						},
					},
				},
			},
			"member_key": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `EntityKey of the member.`,
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
				ExactlyOneOf: []string{"member_key", "preferred_member_key"},
			},
			"preferred_member_key": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `EntityKey of the member.`,
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
				ExactlyOneOf: []string{"member_key", "preferred_member_key"},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the Membership was created.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the Membership, of the form groups/{group_id}/memberships/{membership_id}.`,
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The type of the membership.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the Membership was last updated.`,
			},
		},
	}
}

func resourceCloudIdentityGroupMembershipCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	memberKeyProp, err := expandCloudIdentityGroupMembershipMemberKey(d.Get("member_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("member_key"); !isEmptyValue(reflect.ValueOf(memberKeyProp)) && (ok || !reflect.DeepEqual(v, memberKeyProp)) {
		obj["memberKey"] = memberKeyProp
	}
	preferredMemberKeyProp, err := expandCloudIdentityGroupMembershipPreferredMemberKey(d.Get("preferred_member_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("preferred_member_key"); !isEmptyValue(reflect.ValueOf(preferredMemberKeyProp)) && (ok || !reflect.DeepEqual(v, preferredMemberKeyProp)) {
		obj["preferredMemberKey"] = preferredMemberKeyProp
	}
	rolesProp, err := expandCloudIdentityGroupMembershipRoles(d.Get("roles"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("roles"); !isEmptyValue(reflect.ValueOf(rolesProp)) && (ok || !reflect.DeepEqual(v, rolesProp)) {
		obj["roles"] = rolesProp
	}

	url, err := replaceVars(d, config, "{{CloudIdentityBasePath}}{{group}}/memberships")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GroupMembership: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating GroupMembership: %s", err)
	}
	if err := d.Set("name", flattenCloudIdentityGroupMembershipName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating GroupMembership %q: %#v", d.Id(), res)

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

	return resourceCloudIdentityGroupMembershipRead(d, meta)
}

func resourceCloudIdentityGroupMembershipRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{CloudIdentityBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("CloudIdentityGroupMembership %q", d.Id()))
	}

	if err := d.Set("name", flattenCloudIdentityGroupMembershipName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading GroupMembership: %s", err)
	}
	if err := d.Set("member_key", flattenCloudIdentityGroupMembershipMemberKey(res["memberKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading GroupMembership: %s", err)
	}
	if err := d.Set("preferred_member_key", flattenCloudIdentityGroupMembershipPreferredMemberKey(res["preferredMemberKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading GroupMembership: %s", err)
	}
	if err := d.Set("create_time", flattenCloudIdentityGroupMembershipCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading GroupMembership: %s", err)
	}
	if err := d.Set("update_time", flattenCloudIdentityGroupMembershipUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading GroupMembership: %s", err)
	}
	if err := d.Set("roles", flattenCloudIdentityGroupMembershipRoles(res["roles"], d, config)); err != nil {
		return fmt.Errorf("Error reading GroupMembership: %s", err)
	}
	if err := d.Set("type", flattenCloudIdentityGroupMembershipType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading GroupMembership: %s", err)
	}

	return nil
}

func resourceCloudIdentityGroupMembershipUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	memberKeyProp, err := expandCloudIdentityGroupMembershipMemberKey(d.Get("member_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("member_key"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, memberKeyProp)) {
		obj["memberKey"] = memberKeyProp
	}
	preferredMemberKeyProp, err := expandCloudIdentityGroupMembershipPreferredMemberKey(d.Get("preferred_member_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("preferred_member_key"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, preferredMemberKeyProp)) {
		obj["preferredMemberKey"] = preferredMemberKeyProp
	}
	rolesProp, err := expandCloudIdentityGroupMembershipRoles(d.Get("roles"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("roles"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, rolesProp)) {
		obj["roles"] = rolesProp
	}

	url, err := replaceVars(d, config, "{{CloudIdentityBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating GroupMembership %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PUT", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating GroupMembership %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating GroupMembership %q: %#v", d.Id(), res)
	}

	return resourceCloudIdentityGroupMembershipRead(d, meta)
}

func resourceCloudIdentityGroupMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := replaceVars(d, config, "{{CloudIdentityBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting GroupMembership %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "GroupMembership")
	}

	log.Printf("[DEBUG] Finished deleting GroupMembership %q: %#v", d.Id(), res)
	return nil
}

func resourceCloudIdentityGroupMembershipImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
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

func flattenCloudIdentityGroupMembershipName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipMemberKey(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["id"] =
		flattenCloudIdentityGroupMembershipMemberKeyId(original["id"], d, config)
	transformed["namespace"] =
		flattenCloudIdentityGroupMembershipMemberKeyNamespace(original["namespace"], d, config)
	return []interface{}{transformed}
}
func flattenCloudIdentityGroupMembershipMemberKeyId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipMemberKeyNamespace(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipPreferredMemberKey(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["id"] =
		flattenCloudIdentityGroupMembershipPreferredMemberKeyId(original["id"], d, config)
	transformed["namespace"] =
		flattenCloudIdentityGroupMembershipPreferredMemberKeyNamespace(original["namespace"], d, config)
	return []interface{}{transformed}
}
func flattenCloudIdentityGroupMembershipPreferredMemberKeyId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipPreferredMemberKeyNamespace(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipRoles(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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
			"name": flattenCloudIdentityGroupMembershipRolesName(original["name"], d, config),
		})
	}
	return transformed
}
func flattenCloudIdentityGroupMembershipRolesName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdentityGroupMembershipType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandCloudIdentityGroupMembershipMemberKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedId, err := expandCloudIdentityGroupMembershipMemberKeyId(original["id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedId); val.IsValid() && !isEmptyValue(val) {
		transformed["id"] = transformedId
	}

	transformedNamespace, err := expandCloudIdentityGroupMembershipMemberKeyNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !isEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	return transformed, nil
}

func expandCloudIdentityGroupMembershipMemberKeyId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupMembershipMemberKeyNamespace(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupMembershipPreferredMemberKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedId, err := expandCloudIdentityGroupMembershipPreferredMemberKeyId(original["id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedId); val.IsValid() && !isEmptyValue(val) {
		transformed["id"] = transformedId
	}

	transformedNamespace, err := expandCloudIdentityGroupMembershipPreferredMemberKeyNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !isEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	return transformed, nil
}

func expandCloudIdentityGroupMembershipPreferredMemberKeyId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupMembershipPreferredMemberKeyNamespace(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupMembershipRoles(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandCloudIdentityGroupMembershipRolesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudIdentityGroupMembershipRolesName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
