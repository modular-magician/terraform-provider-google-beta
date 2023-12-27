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

package securesourcemanager

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceSecureSourceManagerRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecureSourceManagerRepositoryCreate,
		Read:   resourceSecureSourceManagerRepositoryRead,
		Update: resourceSecureSourceManagerRepositoryUpdate,
		Delete: resourceSecureSourceManagerRepositoryDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSecureSourceManagerRepositoryImport,
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
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The ID to use for the repository, which will become the final component of the repository's resource name.
This value should be 4-63 characters, and valid characters are /[a-z][0-9]-/.`,
			},
			"allow_missing": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `If set to true, and the repository is not found, the request will succeed but no action will be taken on the server.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Description of the repository.`,
			},
			"etag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.`,
			},
			"initial_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Initial configurations for the repository.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_branch": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Default branch name of the repository.`,
						},
						"gitignores": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `List of gitignore template names user can choose from.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"license": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `License template name user can choose from.`,
						},
						"readme": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `README template name.`,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `A unique identifier for a repository.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Create timestamp.`,
			},
			"instance": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the instance in which the repository is hosted.`,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Unique identifier of the repository.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Update timestamp.`,
			},
			"uris": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `URIs for the repository.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `API is the URI for API access.`,
						},
						"git_http": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `gitHttps is the git HTTPS URI for git operations.`,
						},
						"html": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `HTML is the URI for user to view the repository in a browser.`,
						},
					},
				},
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

func resourceSecureSourceManagerRepositoryCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandSecureSourceManagerRepositoryName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandSecureSourceManagerRepositoryDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	etagProp, err := expandSecureSourceManagerRepositoryEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	initialConfigProp, err := expandSecureSourceManagerRepositoryInitialConfig(d.Get("initial_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("initial_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(initialConfigProp)) && (ok || !reflect.DeepEqual(v, initialConfigProp)) {
		obj["initialConfig"] = initialConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecureSourceManagerBasePath}}projects/{{project}}/locations/{{location}}/repositories?repositoryId={{respository_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Repository: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Repository: %s", err)
	}
	billingProject = project

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
		return fmt.Errorf("Error creating Repository: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = SecureSourceManagerOperationWaitTime(
		config, res, project, "Creating Repository", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Repository: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Repository %q: %#v", d.Id(), res)

	return resourceSecureSourceManagerRepositoryRead(d, meta)
}

func resourceSecureSourceManagerRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecureSourceManagerBasePath}}projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Repository: %s", err)
	}
	billingProject = project

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SecureSourceManagerRepository %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}

	if err := d.Set("name", flattenSecureSourceManagerRepositoryName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("description", flattenSecureSourceManagerRepositoryDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("instance", flattenSecureSourceManagerRepositoryInstance(res["instance"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("uid", flattenSecureSourceManagerRepositoryUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("create_time", flattenSecureSourceManagerRepositoryCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("update_time", flattenSecureSourceManagerRepositoryUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("etag", flattenSecureSourceManagerRepositoryEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("uris", flattenSecureSourceManagerRepositoryUris(res["uris"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("initial_config", flattenSecureSourceManagerRepositoryInitialConfig(res["initialConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}

	return nil
}

func resourceSecureSourceManagerRepositoryUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Repository: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	nameProp, err := expandSecureSourceManagerRepositoryName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandSecureSourceManagerRepositoryDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	etagProp, err := expandSecureSourceManagerRepositoryEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	initialConfigProp, err := expandSecureSourceManagerRepositoryInitialConfig(d.Get("initial_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("initial_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, initialConfigProp)) {
		obj["initialConfig"] = initialConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecureSourceManagerBasePath}}projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Repository %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PUT",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating Repository %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Repository %q: %#v", d.Id(), res)
	}

	return resourceSecureSourceManagerRepositoryRead(d, meta)
}

func resourceSecureSourceManagerRepositoryDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Repository: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{SecureSourceManagerBasePath}}projects/{{project}}/locations/{{location}}/repositories/{{respository_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Repository %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Repository")
	}

	err = SecureSourceManagerOperationWaitTime(
		config, res, project, "Deleting Repository", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Repository %q: %#v", d.Id(), res)
	return nil
}

func resourceSecureSourceManagerRepositoryImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/repositories/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenSecureSourceManagerRepositoryName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerRepositoryDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerRepositoryInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerRepositoryUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerRepositoryCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerRepositoryUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerRepositoryEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerRepositoryUris(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["html"] =
		flattenSecureSourceManagerRepositoryUrisHtml(original["html"], d, config)
	transformed["git_http"] =
		flattenSecureSourceManagerRepositoryUrisGitHttp(original["gitHttp"], d, config)
	transformed["api"] =
		flattenSecureSourceManagerRepositoryUrisApi(original["api"], d, config)
	return []interface{}{transformed}
}
func flattenSecureSourceManagerRepositoryUrisHtml(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerRepositoryUrisGitHttp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerRepositoryUrisApi(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerRepositoryInitialConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["default_branch"] =
		flattenSecureSourceManagerRepositoryInitialConfigDefaultBranch(original["defaultBranch"], d, config)
	transformed["gitignores"] =
		flattenSecureSourceManagerRepositoryInitialConfigGitignores(original["gitignores"], d, config)
	transformed["license"] =
		flattenSecureSourceManagerRepositoryInitialConfigLicense(original["license"], d, config)
	transformed["readme"] =
		flattenSecureSourceManagerRepositoryInitialConfigReadme(original["readme"], d, config)
	return []interface{}{transformed}
}
func flattenSecureSourceManagerRepositoryInitialConfigDefaultBranch(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerRepositoryInitialConfigGitignores(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerRepositoryInitialConfigLicense(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecureSourceManagerRepositoryInitialConfigReadme(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandSecureSourceManagerRepositoryName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerRepositoryDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerRepositoryEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerRepositoryInitialConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDefaultBranch, err := expandSecureSourceManagerRepositoryInitialConfigDefaultBranch(original["default_branch"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultBranch); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["defaultBranch"] = transformedDefaultBranch
	}

	transformedGitignores, err := expandSecureSourceManagerRepositoryInitialConfigGitignores(original["gitignores"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGitignores); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gitignores"] = transformedGitignores
	}

	transformedLicense, err := expandSecureSourceManagerRepositoryInitialConfigLicense(original["license"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLicense); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["license"] = transformedLicense
	}

	transformedReadme, err := expandSecureSourceManagerRepositoryInitialConfigReadme(original["readme"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReadme); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["readme"] = transformedReadme
	}

	return transformed, nil
}

func expandSecureSourceManagerRepositoryInitialConfigDefaultBranch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerRepositoryInitialConfigGitignores(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerRepositoryInitialConfigLicense(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerRepositoryInitialConfigReadme(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
