// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	vertexai "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai/beta"
)

func resourceVertexAiModelDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceVertexAiModelDeploymentCreate,
		Read:   resourceVertexAiModelDeploymentRead,
		Delete: resourceVertexAiModelDeploymentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVertexAiModelDeploymentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dedicated_resources": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "A description of resources that are dedicated to the DeployedModel, and that need a higher degree of manual configuration.",
				MaxItems:    1,
				Elem:        VertexAiModelDeploymentDedicatedResourcesSchema(),
			},

			"endpoint": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "The name of the endpoint to deploy to",
			},

			"model": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "The name of the model to deploy",
			},

			"location": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "The location of the endpoint",
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "The project of the endpoint",
			},

			"deployed_model_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The deployed ID of the model in the endpoint",
			},
		},
	}
}

func VertexAiModelDeploymentDedicatedResourcesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"machine_spec": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Required. Immutable. The specification of a single machine used by the prediction.",
				MaxItems:    1,
				Elem:        VertexAiModelDeploymentDedicatedResourcesMachineSpecSchema(),
			},

			"min_replica_count": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Required. Immutable. The minimum number of machine replicas this DeployedModel will be always deployed on. This value must be greater than or equal to 1. If traffic against the DeployedModel increases, it may dynamically be deployed onto more replicas, and as traffic decreases, some of these extra replicas may be freed.",
			},

			"max_replica_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Immutable. The maximum number of replicas this DeployedModel may be deployed on when the traffic against it increases. If the requested value is too large, the deployment will error, but if deployment succeeds then the ability to scale the model to that many replicas is guaranteed (barring service outages). If traffic against the DeployedModel increases beyond what its replicas at maximum may handle, a portion of the traffic will be dropped. If this value is not provided, will use min_replica_count as the default value. The value of this field impacts the charge against Vertex CPU and GPU quotas. Specifically, you will be charged for max_replica_count * number of cores in the selected machine type) and (max_replica_count * number of GPUs per replica in the selected machine type).",
			},
		},
	}
}

func VertexAiModelDeploymentDedicatedResourcesMachineSpecSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"machine_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Immutable. The type of the machine. See the [list of machine types supported for prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types) See the [list of machine types supported for custom training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types). For DeployedModel this field is optional, and the default value is `n1-standard-2`. For BatchPredictionJob or as part of WorkerPoolSpec this field is required. TODO(rsurowka): Try to better unify the required vs optional.",
			},
		},
	}
}

func resourceVertexAiModelDeploymentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vertexai.ModelDeployment{
		DedicatedResources: expandVertexAiModelDeploymentDedicatedResources(d.Get("dedicated_resources")),
		Endpoint:           dcl.String(d.Get("endpoint").(string)),
		Model:              dcl.String(d.Get("model").(string)),
		Location:           dcl.StringOrNil(d.Get("location").(string)),
		Project:            dcl.String(project),
	}

	id, err := obj.ID()
	if err != nil {
		return fmt.Errorf("error constructing id: %s", err)
	}
	d.SetId(id)
	directive := CreateDirective
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLVertexAiClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutCreate))
	if bp, err := replaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	res, err := client.ApplyModelDeployment(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating ModelDeployment: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ModelDeployment %q: %#v", d.Id(), res)

	return resourceVertexAiModelDeploymentRead(d, meta)
}

func resourceVertexAiModelDeploymentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vertexai.ModelDeployment{
		DedicatedResources: expandVertexAiModelDeploymentDedicatedResources(d.Get("dedicated_resources")),
		Endpoint:           dcl.String(d.Get("endpoint").(string)),
		Model:              dcl.String(d.Get("model").(string)),
		Location:           dcl.StringOrNil(d.Get("location").(string)),
		Project:            dcl.String(project),
	}

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLVertexAiClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutRead))
	if bp, err := replaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	res, err := client.GetModelDeployment(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("VertexAiModelDeployment %q", d.Id())
		return handleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("dedicated_resources", flattenVertexAiModelDeploymentDedicatedResources(res.DedicatedResources)); err != nil {
		return fmt.Errorf("error setting dedicated_resources in state: %s", err)
	}
	if err = d.Set("endpoint", res.Endpoint); err != nil {
		return fmt.Errorf("error setting endpoint in state: %s", err)
	}
	if err = d.Set("model", res.Model); err != nil {
		return fmt.Errorf("error setting model in state: %s", err)
	}
	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("deployed_model_id", res.DeployedModelId); err != nil {
		return fmt.Errorf("error setting deployed_model_id in state: %s", err)
	}

	return nil
}

func resourceVertexAiModelDeploymentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vertexai.ModelDeployment{
		DedicatedResources: expandVertexAiModelDeploymentDedicatedResources(d.Get("dedicated_resources")),
		Endpoint:           dcl.String(d.Get("endpoint").(string)),
		Model:              dcl.String(d.Get("model").(string)),
		Location:           dcl.StringOrNil(d.Get("location").(string)),
		Project:            dcl.String(project),
	}

	log.Printf("[DEBUG] Deleting ModelDeployment %q", d.Id())
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLVertexAiClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutDelete))
	if bp, err := replaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	if err := client.DeleteModelDeployment(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting ModelDeployment: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting ModelDeployment %q", d.Id())
	return nil
}

func resourceVertexAiModelDeploymentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/endpoints/(?P<endpoint>[^/]+)/models/(?P<model>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<endpoint>[^/]+)/(?P<model>[^/]+)",
		"(?P<location>[^/]+)/(?P<endpoint>[^/]+)/(?P<model>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}/models/{{model}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandVertexAiModelDeploymentDedicatedResources(o interface{}) *vertexai.ModelDeploymentDedicatedResources {
	if o == nil {
		return vertexai.EmptyModelDeploymentDedicatedResources
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return vertexai.EmptyModelDeploymentDedicatedResources
	}
	obj := objArr[0].(map[string]interface{})
	return &vertexai.ModelDeploymentDedicatedResources{
		MachineSpec:     expandVertexAiModelDeploymentDedicatedResourcesMachineSpec(obj["machine_spec"]),
		MinReplicaCount: dcl.Int64(int64(obj["min_replica_count"].(int))),
		MaxReplicaCount: dcl.Int64OrNil(int64(obj["max_replica_count"].(int))),
	}
}

func flattenVertexAiModelDeploymentDedicatedResources(obj *vertexai.ModelDeploymentDedicatedResources) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"machine_spec":      flattenVertexAiModelDeploymentDedicatedResourcesMachineSpec(obj.MachineSpec),
		"min_replica_count": obj.MinReplicaCount,
		"max_replica_count": obj.MaxReplicaCount,
	}

	return []interface{}{transformed}

}

func expandVertexAiModelDeploymentDedicatedResourcesMachineSpec(o interface{}) *vertexai.ModelDeploymentDedicatedResourcesMachineSpec {
	if o == nil {
		return vertexai.EmptyModelDeploymentDedicatedResourcesMachineSpec
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return vertexai.EmptyModelDeploymentDedicatedResourcesMachineSpec
	}
	obj := objArr[0].(map[string]interface{})
	return &vertexai.ModelDeploymentDedicatedResourcesMachineSpec{
		MachineType: dcl.String(obj["machine_type"].(string)),
	}
}

func flattenVertexAiModelDeploymentDedicatedResourcesMachineSpec(obj *vertexai.ModelDeploymentDedicatedResourcesMachineSpec) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"machine_type": obj.MachineType,
	}

	return []interface{}{transformed}

}
