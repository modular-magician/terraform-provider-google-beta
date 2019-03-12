package google

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

func resourceGoogleProjectOrganizationPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceGoogleProjectOrganizationPolicyCreate,
		Read:   resourceGoogleProjectOrganizationPolicyRead,
		Update: resourceGoogleProjectOrganizationPolicyUpdate,
		Delete: resourceGoogleProjectOrganizationPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceProjectOrgPolicyImporter,
		},

		Schema: mergeSchemas(
			schemaOrganizationPolicy,
			map[string]*schema.Schema{
				"project": {
					Type:     schema.TypeString,
					Required: true,
					ForceNew: true,
				},
			},
		),
	}
}

func resourceProjectOrgPolicyImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("ID must be in the format <projectId>:constraints/<constraint>")
	}
	d.Set("project", parts[0])
	d.Set("constraint", parts[1])

	return []*schema.ResourceData{d}, nil
}

func resourceGoogleProjectOrganizationPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	if err := setProjectOrganizationPolicy(d, meta); err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s:%s", d.Get("project"), d.Get("constraint")))

	return resourceGoogleProjectOrganizationPolicyRead(d, meta)
}

func resourceGoogleProjectOrganizationPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project := prefixedProject(d.Get("project").(string))

	policy, err := config.clientResourceManager.Projects.GetOrgPolicy(project, &cloudresourcemanager.GetOrgPolicyRequest{
		Constraint: canonicalOrgPolicyConstraint(d.Get("constraint").(string)),
	}).Do()

	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("Organization policy for %s", project))
	}

	d.Set("constraint", policy.Constraint)
	d.Set("boolean_policy", flattenBooleanOrganizationPolicy(policy.BooleanPolicy))
	d.Set("list_policy", flattenListOrganizationPolicy(policy.ListPolicy))
	d.Set("restore_policy", flattenRestoreOrganizationPolicy(policy.RestoreDefault))
	d.Set("version", policy.Version)
	d.Set("etag", policy.Etag)
	d.Set("update_time", policy.UpdateTime)

	return nil
}

func resourceGoogleProjectOrganizationPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	if err := setProjectOrganizationPolicy(d, meta); err != nil {
		return err
	}

	return resourceGoogleProjectOrganizationPolicyRead(d, meta)
}

func resourceGoogleProjectOrganizationPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project := prefixedProject(d.Get("project").(string))

	_, err := config.clientResourceManager.Projects.ClearOrgPolicy(project, &cloudresourcemanager.ClearOrgPolicyRequest{
		Constraint: canonicalOrgPolicyConstraint(d.Get("constraint").(string)),
	}).Do()

	if err != nil {
		return err
	}

	return nil
}

func setProjectOrganizationPolicy(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project := prefixedProject(d.Get("project").(string))

	listPolicy, err := expandListOrganizationPolicy(d.Get("list_policy").([]interface{}))
	if err != nil {
		return err
	}

	restore_default, err := expandRestoreOrganizationPolicy(d.Get("restore_policy").([]interface{}))
	if err != nil {
		return err
	}

	_, err = config.clientResourceManager.Projects.SetOrgPolicy(project, &cloudresourcemanager.SetOrgPolicyRequest{
		Policy: &cloudresourcemanager.OrgPolicy{
			Constraint:     canonicalOrgPolicyConstraint(d.Get("constraint").(string)),
			BooleanPolicy:  expandBooleanOrganizationPolicy(d.Get("boolean_policy").([]interface{})),
			ListPolicy:     listPolicy,
			RestoreDefault: restore_default,
			Version:        int64(d.Get("version").(int)),
			Etag:           d.Get("etag").(string),
		},
	}).Do()

	return err
}
