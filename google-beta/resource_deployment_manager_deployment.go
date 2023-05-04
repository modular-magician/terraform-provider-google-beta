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
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func customDiffDeploymentManagerDeployment(_ context.Context, d *schema.ResourceDiff, meta interface{}) error {
	if preview := d.Get("preview").(bool); preview {
		log.Printf("[WARN] Deployment preview set to true - Terraform will treat Deployment as recreate-only")

		if d.HasChange("preview") {
			if err := d.ForceNew("preview"); err != nil {
				return err
			}
		}

		if d.HasChange("target") {
			if err := d.ForceNew("target"); err != nil {
				return err
			}
		}

		if d.HasChange("labels") {
			if err := d.ForceNew("labels"); err != nil {
				return err
			}
		}
	}
	return nil
}

func ResourceDeploymentManagerDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeploymentManagerDeploymentCreate,
		Read:   resourceDeploymentManagerDeploymentRead,
		Update: resourceDeploymentManagerDeploymentUpdate,
		Delete: resourceDeploymentManagerDeploymentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDeploymentManagerDeploymentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customDiffDeploymentManagerDeployment,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Unique name for the deployment`,
			},
			"target": {
				Type:     schema.TypeList,
				Required: true,
				Description: `Parameters that define your deployment, including the deployment
configuration and relevant templates.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `The root configuration file to use for this deployment.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"content": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `The full YAML contents of your configuration file.`,
									},
								},
							},
						},
						"imports": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `Specifies import files for this configuration. This can be
used to import templates or other files. For example, you might
import a text file in order to use the file in a template.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"content": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The full contents of the template that you want to import.`,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `The name of the template to import, as declared in the YAML
configuration.`,
									},
								},
							},
						},
					},
				},
			},
			"create_policy": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ACQUIRE", "CREATE_OR_ACQUIRE", ""}),
				Description: `Set the policy to use for creating new resources. Only used on
create and update. Valid values are 'CREATE_OR_ACQUIRE' (default) or
'ACQUIRE'. If set to 'ACQUIRE' and resources do not already exist,
the deployment will fail. Note that updating this field does not
actually affect the deployment, just how it is updated. Default value: "CREATE_OR_ACQUIRE" Possible values: ["ACQUIRE", "CREATE_OR_ACQUIRE"]`,
				Default: "CREATE_OR_ACQUIRE",
			},
			"delete_policy": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ABANDON", "DELETE", ""}),
				Description: `Set the policy to use for deleting new resources on update/delete.
Valid values are 'DELETE' (default) or 'ABANDON'. If 'DELETE',
resource is deleted after removal from Deployment Manager. If
'ABANDON', the resource is only removed from Deployment Manager
and is not actually deleted. Note that updating this field does not
actually change the deployment, just how it is updated. Default value: "DELETE" Possible values: ["ABANDON", "DELETE"]`,
				Default: "DELETE",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Optional user-provided description of deployment.`,
			},
			"labels": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: `Key-value pairs to apply to this labels.`,
				Elem:        deploymentmanagerDeploymentLabelsSchema(),
				// Default schema.HashSchema is used.
			},
			"preview": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `If set to true, a deployment is created with "shell" resources
that are not actually instantiated. This allows you to preview a
deployment. It can be updated to false to actually deploy
with real resources.
 ~>**NOTE:** Deployment Manager does not allow update
of a deployment in preview (unless updating to preview=false). Thus,
Terraform will force-recreate deployments if either preview is updated
to true or if other fields are updated while preview is true.`,
				Default: false,
			},
			"deployment_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Unique identifier for deployment. Output only.`,
			},
			"manifest": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. URL of the manifest representing the last manifest that
was successfully deployed.`,
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Server defined URL for the resource.`,
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

func deploymentmanagerDeploymentLabelsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Key for label.`,
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Value of label.`,
			},
		},
	}
}

func resourceDeploymentManagerDeploymentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandDeploymentManagerDeploymentName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandDeploymentManagerDeploymentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandDeploymentManagerDeploymentLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); ok || !reflect.DeepEqual(v, labelsProp) {
		obj["labels"] = labelsProp
	}
	targetProp, err := expandDeploymentManagerDeploymentTarget(d.Get("target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}

	url, err := ReplaceVars(d, config, "{{DeploymentManagerBasePath}}projects/{{project}}/global/deployments?preview={{preview}}&createPolicy={{create_policy}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Deployment: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Deployment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Deployment: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/deployments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = DeploymentManagerOperationWaitTime(
		config, res, project, "Creating Deployment", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		resourceDeploymentManagerDeploymentPostCreateFailure(d, meta)
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Deployment: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Deployment %q: %#v", d.Id(), res)

	return resourceDeploymentManagerDeploymentRead(d, meta)
}

func resourceDeploymentManagerDeploymentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{DeploymentManagerBasePath}}projects/{{project}}/global/deployments/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Deployment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DeploymentManagerDeployment %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Deployment: %s", err)
	}

	if err := d.Set("name", flattenDeploymentManagerDeploymentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Deployment: %s", err)
	}
	if err := d.Set("description", flattenDeploymentManagerDeploymentDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Deployment: %s", err)
	}
	if err := d.Set("labels", flattenDeploymentManagerDeploymentLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Deployment: %s", err)
	}
	if err := d.Set("deployment_id", flattenDeploymentManagerDeploymentDeploymentId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading Deployment: %s", err)
	}
	if err := d.Set("manifest", flattenDeploymentManagerDeploymentManifest(res["manifest"], d, config)); err != nil {
		return fmt.Errorf("Error reading Deployment: %s", err)
	}
	if err := d.Set("self_link", flattenDeploymentManagerDeploymentSelfLink(res["selfLink"], d, config)); err != nil {
		return fmt.Errorf("Error reading Deployment: %s", err)
	}

	return nil
}

func resourceDeploymentManagerDeploymentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Deployment: %s", err)
	}
	billingProject = project

	d.Partial(true)

	if d.HasChange("preview") {
		obj := make(map[string]interface{})

		getUrl, err := ReplaceVars(d, config, "{{DeploymentManagerBasePath}}projects/{{project}}/global/deployments/{{name}}")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		getRes, err := transport_tpg.SendRequest(config, "GET", billingProject, getUrl, userAgent, nil)
		if err != nil {
			return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DeploymentManagerDeployment %q", d.Id()))
		}

		obj["fingerprint"] = getRes["fingerprint"]

		url, err := ReplaceVars(d, config, "{{DeploymentManagerBasePath}}projects/{{project}}/global/deployments/{{name}}?preview={{preview}}&createPolicy={{create_policy}}&deletePolicy={{delete_policy}}")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating Deployment %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Deployment %q: %#v", d.Id(), res)
		}

		err = DeploymentManagerOperationWaitTime(
			config, res, project, "Updating Deployment", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("description") || d.HasChange("labels") || d.HasChange("target") {
		obj := make(map[string]interface{})

		getUrl, err := ReplaceVars(d, config, "{{DeploymentManagerBasePath}}projects/{{project}}/global/deployments/{{name}}")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		getRes, err := transport_tpg.SendRequest(config, "GET", billingProject, getUrl, userAgent, nil)
		if err != nil {
			return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DeploymentManagerDeployment %q", d.Id()))
		}

		obj["fingerprint"] = getRes["fingerprint"]

		descriptionProp, err := expandDeploymentManagerDeploymentDescription(d.Get("description"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
			obj["description"] = descriptionProp
		}
		labelsProp, err := expandDeploymentManagerDeploymentLabels(d.Get("labels"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("labels"); ok || !reflect.DeepEqual(v, labelsProp) {
			obj["labels"] = labelsProp
		}
		targetProp, err := expandDeploymentManagerDeploymentTarget(d.Get("target"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetProp)) {
			obj["target"] = targetProp
		}

		url, err := ReplaceVars(d, config, "{{DeploymentManagerBasePath}}projects/{{project}}/global/deployments/{{name}}?preview={{preview}}&createPolicy={{create_policy}}&deletePolicy={{delete_policy}}")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating Deployment %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Deployment %q: %#v", d.Id(), res)
		}

		err = DeploymentManagerOperationWaitTime(
			config, res, project, "Updating Deployment", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceDeploymentManagerDeploymentRead(d, meta)
}

func resourceDeploymentManagerDeploymentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Deployment: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{DeploymentManagerBasePath}}projects/{{project}}/global/deployments/{{name}}?deletePolicy={{delete_policy}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Deployment %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Deployment")
	}

	err = DeploymentManagerOperationWaitTime(
		config, res, project, "Deleting Deployment", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Deployment %q: %#v", d.Id(), res)
	return nil
}

func resourceDeploymentManagerDeploymentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/deployments/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/deployments/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDeploymentManagerDeploymentName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDeploymentManagerDeploymentDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDeploymentManagerDeploymentLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(deploymentmanagerDeploymentLabelsSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"key":   flattenDeploymentManagerDeploymentLabelsKey(original["key"], d, config),
			"value": flattenDeploymentManagerDeploymentLabelsValue(original["value"], d, config),
		})
	}
	return transformed
}
func flattenDeploymentManagerDeploymentLabelsKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDeploymentManagerDeploymentLabelsValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDeploymentManagerDeploymentDeploymentId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDeploymentManagerDeploymentManifest(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDeploymentManagerDeploymentSelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDeploymentManagerDeploymentName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDeploymentManagerDeploymentDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDeploymentManagerDeploymentLabels(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKey, err := expandDeploymentManagerDeploymentLabelsKey(original["key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !isEmptyValue(val) {
			transformed["key"] = transformedKey
		}

		transformedValue, err := expandDeploymentManagerDeploymentLabelsValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !isEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDeploymentManagerDeploymentLabelsKey(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDeploymentManagerDeploymentLabelsValue(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDeploymentManagerDeploymentTarget(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedConfig, err := expandDeploymentManagerDeploymentTargetConfig(original["config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["config"] = transformedConfig
	}

	transformedImports, err := expandDeploymentManagerDeploymentTargetImports(original["imports"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImports); val.IsValid() && !isEmptyValue(val) {
		transformed["imports"] = transformedImports
	}

	return transformed, nil
}

func expandDeploymentManagerDeploymentTargetConfig(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedContent, err := expandDeploymentManagerDeploymentTargetConfigContent(original["content"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContent); val.IsValid() && !isEmptyValue(val) {
		transformed["content"] = transformedContent
	}

	return transformed, nil
}

func expandDeploymentManagerDeploymentTargetConfigContent(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDeploymentManagerDeploymentTargetImports(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedContent, err := expandDeploymentManagerDeploymentTargetImportsContent(original["content"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedContent); val.IsValid() && !isEmptyValue(val) {
			transformed["content"] = transformedContent
		}

		transformedName, err := expandDeploymentManagerDeploymentTargetImportsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDeploymentManagerDeploymentTargetImportsContent(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDeploymentManagerDeploymentTargetImportsName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceDeploymentManagerDeploymentPostCreateFailure(d *schema.ResourceData, meta interface{}) {
	log.Printf("[WARN] Attempt to clean up Deployment if it still exists")
	var cleanErr error
	if cleanErr = resourceDeploymentManagerDeploymentRead(d, meta); cleanErr == nil {
		if d.Id() != "" {
			log.Printf("[WARN] Deployment %q still exists, attempting to delete...", d.Id())
			if cleanErr = resourceDeploymentManagerDeploymentDelete(d, meta); cleanErr == nil {
				log.Printf("[WARN] Invalid Deployment was successfully deleted")
				d.SetId("")
			}
		}
	}
	if cleanErr != nil {
		log.Printf("[WARN] Could not confirm cleanup of Deployment if created in error state: %v", cleanErr)
	}
}
