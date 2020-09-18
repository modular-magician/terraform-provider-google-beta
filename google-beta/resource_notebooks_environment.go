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
)

func resourceNotebooksEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceNotebooksEnvironmentCreate,
		Read:   resourceNotebooksEnvironmentRead,
		Update: resourceNotebooksEnvironmentUpdate,
		Delete: resourceNotebooksEnvironmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNotebooksEnvironmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `A reference to the zone where the machine resides.`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The name specified for the Environment instance.
Format: projects/{project_id}/locations/{location}/environments/{environmentId}`,
			},
			"container_image": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Use a container image to start the notebook instance.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repository": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The path to the container image repository. 
For example: gcr.io/{project_id}/{imageName}`,
						},
						"tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The tag of the container image. If not specified, this defaults to the latest tag.`,
						},
					},
				},
				ExactlyOneOf: []string{"vm_image", "container_image"},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A brief description of this environment.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Display name of this environment for the UI.`,
			},
			"post_startup_script": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Path to a Bash script that automatically runs after a notebook instance fully boots up. 
The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"`,
			},
			"vm_image": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Use a Compute Engine VM image to start the notebook instance.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The name of the Google Cloud project that this VM image belongs to. 
Format: projects/{project_id}`,
						},
						"image_family": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Use this VM image family to find the image; the newest image in this family will be used.`,
						},
						"image_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Use VM image name to find the image.`,
						},
					},
				},
				ExactlyOneOf: []string{"vm_image", "container_image"},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Instance creation time`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceNotebooksEnvironmentCreate(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	obj := make(map[string]interface{})
	displayNameProp, err := expandNotebooksEnvironmentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandNotebooksEnvironmentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	postStartupScriptProp, err := expandNotebooksEnvironmentPostStartupScript(d.Get("post_startup_script"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("post_startup_script"); !isEmptyValue(reflect.ValueOf(postStartupScriptProp)) && (ok || !reflect.DeepEqual(v, postStartupScriptProp)) {
		obj["postStartupScript"] = postStartupScriptProp
	}
	vmImageProp, err := expandNotebooksEnvironmentVmImage(d.Get("vm_image"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vm_image"); !isEmptyValue(reflect.ValueOf(vmImageProp)) && (ok || !reflect.DeepEqual(v, vmImageProp)) {
		obj["vmImage"] = vmImageProp
	}
	containerImageProp, err := expandNotebooksEnvironmentContainerImage(d.Get("container_image"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("container_image"); !isEmptyValue(reflect.ValueOf(containerImageProp)) && (ok || !reflect.DeepEqual(v, containerImageProp)) {
		obj["containerImage"] = containerImageProp
	}

	url, err := replaceVars(d, config, "{{NotebooksBasePath}}projects/{{project}}/locations/{{location}}/environments?environmentId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Environment: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Environment: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/environments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = notebooksOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Environment",
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Environment: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "projects/{{project}}/locations/{{location}}/environments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Environment %q: %#v", d.Id(), res)

	return resourceNotebooksEnvironmentRead(d, meta)
}

func resourceNotebooksEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	url, err := replaceVars(d, config, "{{NotebooksBasePath}}projects/{{project}}/locations/{{location}}/environments/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("NotebooksEnvironment %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Environment: %s", err)
	}

	if err := d.Set("display_name", flattenNotebooksEnvironmentDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Environment: %s", err)
	}
	if err := d.Set("description", flattenNotebooksEnvironmentDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Environment: %s", err)
	}
	if err := d.Set("post_startup_script", flattenNotebooksEnvironmentPostStartupScript(res["postStartupScript"], d, config)); err != nil {
		return fmt.Errorf("Error reading Environment: %s", err)
	}
	if err := d.Set("create_time", flattenNotebooksEnvironmentCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Environment: %s", err)
	}
	if err := d.Set("vm_image", flattenNotebooksEnvironmentVmImage(res["vmImage"], d, config)); err != nil {
		return fmt.Errorf("Error reading Environment: %s", err)
	}
	if err := d.Set("container_image", flattenNotebooksEnvironmentContainerImage(res["containerImage"], d, config)); err != nil {
		return fmt.Errorf("Error reading Environment: %s", err)
	}

	return nil
}

func resourceNotebooksEnvironmentUpdate(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandNotebooksEnvironmentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandNotebooksEnvironmentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	postStartupScriptProp, err := expandNotebooksEnvironmentPostStartupScript(d.Get("post_startup_script"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("post_startup_script"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, postStartupScriptProp)) {
		obj["postStartupScript"] = postStartupScriptProp
	}
	vmImageProp, err := expandNotebooksEnvironmentVmImage(d.Get("vm_image"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vm_image"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, vmImageProp)) {
		obj["vmImage"] = vmImageProp
	}
	containerImageProp, err := expandNotebooksEnvironmentContainerImage(d.Get("container_image"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("container_image"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, containerImageProp)) {
		obj["containerImage"] = containerImageProp
	}

	url, err := replaceVars(d, config, "{{NotebooksBasePath}}projects/{{project}}/locations/{{location}}/environments/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Environment %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PUT", billingProject, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Environment %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Environment %q: %#v", d.Id(), res)
	}

	err = notebooksOperationWaitTime(
		config, res, project, "Updating Environment",
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceNotebooksEnvironmentRead(d, meta)
}

func resourceNotebooksEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{NotebooksBasePath}}projects/{{project}}/locations/{{location}}/environments/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Environment %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Environment")
	}

	err = notebooksOperationWaitTime(
		config, res, project, "Deleting Environment",
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Environment %q: %#v", d.Id(), res)
	return nil
}

func resourceNotebooksEnvironmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/environments/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/environments/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNotebooksEnvironmentDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksEnvironmentDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksEnvironmentPostStartupScript(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksEnvironmentCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksEnvironmentVmImage(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["project"] =
		flattenNotebooksEnvironmentVmImageProject(original["project"], d, config)
	transformed["image_name"] =
		flattenNotebooksEnvironmentVmImageImageName(original["imageName"], d, config)
	transformed["image_family"] =
		flattenNotebooksEnvironmentVmImageImageFamily(original["imageFamily"], d, config)
	return []interface{}{transformed}
}
func flattenNotebooksEnvironmentVmImageProject(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksEnvironmentVmImageImageName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksEnvironmentVmImageImageFamily(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksEnvironmentContainerImage(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["repository"] =
		flattenNotebooksEnvironmentContainerImageRepository(original["repository"], d, config)
	transformed["tag"] =
		flattenNotebooksEnvironmentContainerImageTag(original["tag"], d, config)
	return []interface{}{transformed}
}
func flattenNotebooksEnvironmentContainerImageRepository(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksEnvironmentContainerImageTag(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandNotebooksEnvironmentDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksEnvironmentDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksEnvironmentPostStartupScript(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksEnvironmentVmImage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProject, err := expandNotebooksEnvironmentVmImageProject(original["project"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProject); val.IsValid() && !isEmptyValue(val) {
		transformed["project"] = transformedProject
	}

	transformedImageName, err := expandNotebooksEnvironmentVmImageImageName(original["image_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImageName); val.IsValid() && !isEmptyValue(val) {
		transformed["imageName"] = transformedImageName
	}

	transformedImageFamily, err := expandNotebooksEnvironmentVmImageImageFamily(original["image_family"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImageFamily); val.IsValid() && !isEmptyValue(val) {
		transformed["imageFamily"] = transformedImageFamily
	}

	return transformed, nil
}

func expandNotebooksEnvironmentVmImageProject(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksEnvironmentVmImageImageName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksEnvironmentVmImageImageFamily(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksEnvironmentContainerImage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRepository, err := expandNotebooksEnvironmentContainerImageRepository(original["repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepository); val.IsValid() && !isEmptyValue(val) {
		transformed["repository"] = transformedRepository
	}

	transformedTag, err := expandNotebooksEnvironmentContainerImageTag(original["tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !isEmptyValue(val) {
		transformed["tag"] = transformedTag
	}

	return transformed, nil
}

func expandNotebooksEnvironmentContainerImageRepository(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksEnvironmentContainerImageTag(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
