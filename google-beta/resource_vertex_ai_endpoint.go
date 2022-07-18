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

func resourceVertexAiEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceVertexAiEndpointCreate,
		Read:   resourceVertexAiEndpointRead,
		Update: resourceVertexAiEndpointUpdate,
		Delete: resourceVertexAiEndpointDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVertexAiEndpointImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Required. The display name of the Endpoint. The name can be up to 128 characters long and can be consist of any UTF-8 characters.",
			},

			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The location for the resource",
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the Endpoint.",
			},

			"encryption_spec": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.",
				MaxItems:    1,
				Elem:        VertexAiEndpointEncryptionSpecSchema(),
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"network": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.",
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "The project for the resource",
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Timestamp when this Endpoint was created.",
			},

			"deployed_models": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively.",
				Elem:        VertexAiEndpointDeployedModelsSchema(),
			},

			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Used to perform consistent read-modify-write updates. If not set, a blind \"overwrite\" update happens.",
			},

			"model_deployment_monitoring_job": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: `projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}`",
			},

			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The resource name of the Endpoint.",
			},

			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Timestamp when this Endpoint was last updated.",
			},
		},
	}
}

func VertexAiEndpointEncryptionSpecSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"kms_key_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. The key needs to be in the same region as where the compute resource is created.",
			},
		},
	}
}

func VertexAiEndpointDeployedModelsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"automatic_resources": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A description of resources that to large degree are decided by Vertex AI, and require only a modest additional configuration.",
				Elem:        VertexAiEndpointDeployedModelsAutomaticResourcesSchema(),
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Timestamp when the DeployedModel was created.",
			},

			"dedicated_resources": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A description of resources that are dedicated to the DeployedModel, and that need a higher degree of manual configuration.",
				Elem:        VertexAiEndpointDeployedModelsDedicatedResourcesSchema(),
			},

			"disable_container_logging": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "For custom-trained Models and AutoML Tabular Models, the container of the DeployedModel instances will send `stderr` and `stdout` streams to Stackdriver Logging by default. Please note that the logs incur cost, which are subject to [Cloud Logging pricing](https://cloud.google.com/stackdriver/pricing). User can disable container logging by setting this flag to true.",
			},

			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The display name of the DeployedModel. If not provided upon creation, the Model's display_name is used.",
			},

			"enable_access_logging": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "These logs are like standard server access logs, containing information like timestamp and latency for each prediction request. Note that Stackdriver logs may incur a cost, especially if your project receives prediction requests at a high queries per second rate (QPS). Estimate your costs before enabling this option.",
			},

			"enable_container_logging": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If true, the container of the DeployedModel instances will send `stderr` and `stdout` streams to Stackdriver Logging. Only supported for custom-trained Models and AutoML Tabular Models.",
			},

			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the DeployedModel. If not provided upon deployment, Vertex AI will generate a value for this ID. This value should be 1-10 characters, and valid characters are /[0-9]/.",
			},

			"model": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the Model that this is the deployment of. Note that the Model may be in a different location than the DeployedModel's Endpoint.",
			},

			"model_version_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The version ID of the model that is deployed.",
			},

			"private_endpoints": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. Provide paths for users to send predict/explain/health requests directly to the deployed model services running on Cloud via private services access. This field is populated if network is configured.",
				Elem:        VertexAiEndpointDeployedModelsPrivateEndpointsSchema(),
			},

			"service_account": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The service account that the DeployedModel's container runs as. Specify the email address of the service account. If this service account is not specified, the container runs as a service account that doesn't have access to the resource project. Users deploying the Model must have the `iam.serviceAccounts.actAs` permission on this service account.",
			},

			"shared_resources": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource name of the shared DeploymentResourcePool to deploy on. Format: projects/{project}/locations/{location}/deploymentResourcePools/{deployment_resource_pool}",
			},
		},
	}
}

func VertexAiEndpointDeployedModelsAutomaticResourcesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_replica_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum number of replicas this DeployedModel may be deployed on when the traffic against it increases. If the requested value is too large, the deployment will error, but if deployment succeeds then the ability to scale the model to that many replicas is guaranteed (barring service outages). If traffic against the DeployedModel increases beyond what its replicas at maximum may handle, a portion of the traffic will be dropped. If this value is not provided, a no upper bound for scaling under heavy traffic will be assume, though Vertex AI may be unable to scale beyond certain replica number.",
			},

			"min_replica_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The minimum number of replicas this DeployedModel will be always deployed on. If traffic against it increases, it may dynamically be deployed onto more replicas up to max_replica_count, and as traffic decreases, some of these extra replicas may be freed. If the requested value is too large, the deployment will error.",
			},
		},
	}
}

func VertexAiEndpointDeployedModelsDedicatedResourcesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"autoscaling_metric_specs": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The metric specifications that overrides a resource utilization metric (CPU utilization, accelerator's duty cycle, and so on) target value (default to 60 if not set). At most one entry is allowed per metric. If machine_spec.accelerator_count is above 0, the autoscaling will be based on both CPU utilization and accelerator's duty cycle metrics and scale up when either metrics exceeds its target value while scale down if both metrics are under their target value. The default target value is 60 for both metrics. If machine_spec.accelerator_count is 0, the autoscaling will be based on CPU utilization metric only with default target value 60 if not explicitly set. For example, in the case of Online Prediction, if you want to override target CPU utilization to 80, you should set autoscaling_metric_specs.metric_name to `aiplatform.googleapis.com/prediction/online/cpu/utilization` and autoscaling_metric_specs.target to `80`.",
				Elem:        VertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsSchema(),
			},

			"machine_spec": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The specification of a single machine used by the prediction.",
				Elem:        VertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecSchema(),
			},

			"max_replica_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum number of replicas this DeployedModel may be deployed on when the traffic against it increases. If the requested value is too large, the deployment will error, but if deployment succeeds then the ability to scale the model to that many replicas is guaranteed (barring service outages). If traffic against the DeployedModel increases beyond what its replicas at maximum may handle, a portion of the traffic will be dropped. If this value is not provided, will use min_replica_count as the default value. The value of this field impacts the charge against Vertex CPU and GPU quotas. Specifically, you will be charged for max_replica_count * number of cores in the selected machine type) and (max_replica_count * number of GPUs per replica in the selected machine type).",
			},

			"min_replica_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The minimum number of machine replicas this DeployedModel will be always deployed on. This value must be greater than or equal to 1. If traffic against the DeployedModel increases, it may dynamically be deployed onto more replicas, and as traffic decreases, some of these extra replicas may be freed.",
			},
		},
	}
}

func VertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metric_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource metric name. Supported metrics: * For Online Prediction: * `aiplatform.googleapis.com/prediction/online/accelerator/duty_cycle` * `aiplatform.googleapis.com/prediction/online/cpu/utilization`",
			},

			"target": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The target resource utilization in percentage (1% - 100%) for the given metric; once the real usage deviates from the target by a certain percentage, the machine replicas change. The default value is 60 (representing 60%) if not provided.",
			},
		},
	}
}

func VertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"accelerator_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of accelerators to attach to the machine.",
			},

			"accelerator_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of accelerator(s) that may be attached to the machine as per accelerator_count. Possible values: ACCELERATOR_TYPE_UNSPECIFIED, NVIDIA_TESLA_K80, NVIDIA_TESLA_P100, NVIDIA_TESLA_V100, NVIDIA_TESLA_P4, NVIDIA_TESLA_T4, NVIDIA_TESLA_A100, TPU_V2, TPU_V3",
			},

			"machine_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the machine. See the [list of machine types supported for prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types) See the [list of machine types supported for custom training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types). For DeployedModel this field is optional, and the default value is `n1-standard-2`. For BatchPredictionJob or as part of WorkerPoolSpec this field is required. TODO(rsurowka): Try to better unify the required vs optional.",
			},
		},
	}
}

func VertexAiEndpointDeployedModelsPrivateEndpointsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"explain_http_uri": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Http(s) path to send explain requests.",
			},

			"health_http_uri": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Http(s) path to send health check requests.",
			},

			"predict_http_uri": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Http(s) path to send prediction requests.",
			},

			"service_attachment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The name of the service attachment resource. Populated if private service connect is enabled.",
			},
		},
	}
}

func resourceVertexAiEndpointCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vertexai.Endpoint{
		DisplayName:    dcl.String(d.Get("display_name").(string)),
		Location:       dcl.String(d.Get("location").(string)),
		Description:    dcl.String(d.Get("description").(string)),
		EncryptionSpec: expandVertexAiEndpointEncryptionSpec(d.Get("encryption_spec")),
		Labels:         checkStringMap(d.Get("labels")),
		Network:        dcl.String(d.Get("network").(string)),
		Project:        dcl.String(project),
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
	res, err := client.ApplyEndpoint(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating Endpoint: %s", err)
	}

	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	// ID has a server-generated value, set again after creation.

	id, err = res.ID()
	if err != nil {
		return fmt.Errorf("error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Endpoint %q: %#v", d.Id(), res)

	return resourceVertexAiEndpointRead(d, meta)
}

func resourceVertexAiEndpointRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vertexai.Endpoint{
		DisplayName:    dcl.String(d.Get("display_name").(string)),
		Location:       dcl.String(d.Get("location").(string)),
		Description:    dcl.String(d.Get("description").(string)),
		EncryptionSpec: expandVertexAiEndpointEncryptionSpec(d.Get("encryption_spec")),
		Labels:         checkStringMap(d.Get("labels")),
		Network:        dcl.String(d.Get("network").(string)),
		Project:        dcl.String(project),
		Name:           dcl.StringOrNil(d.Get("name").(string)),
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
	res, err := client.GetEndpoint(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("VertexAiEndpoint %q", d.Id())
		return handleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("display_name", res.DisplayName); err != nil {
		return fmt.Errorf("error setting display_name in state: %s", err)
	}
	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("description", res.Description); err != nil {
		return fmt.Errorf("error setting description in state: %s", err)
	}
	if err = d.Set("encryption_spec", flattenVertexAiEndpointEncryptionSpec(res.EncryptionSpec)); err != nil {
		return fmt.Errorf("error setting encryption_spec in state: %s", err)
	}
	if err = d.Set("labels", res.Labels); err != nil {
		return fmt.Errorf("error setting labels in state: %s", err)
	}
	if err = d.Set("network", res.Network); err != nil {
		return fmt.Errorf("error setting network in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("deployed_models", flattenVertexAiEndpointDeployedModelsArray(res.DeployedModels)); err != nil {
		return fmt.Errorf("error setting deployed_models in state: %s", err)
	}
	if err = d.Set("etag", res.Etag); err != nil {
		return fmt.Errorf("error setting etag in state: %s", err)
	}
	if err = d.Set("model_deployment_monitoring_job", res.ModelDeploymentMonitoringJob); err != nil {
		return fmt.Errorf("error setting model_deployment_monitoring_job in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("update_time", res.UpdateTime); err != nil {
		return fmt.Errorf("error setting update_time in state: %s", err)
	}

	return nil
}
func resourceVertexAiEndpointUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vertexai.Endpoint{
		DisplayName:    dcl.String(d.Get("display_name").(string)),
		Location:       dcl.String(d.Get("location").(string)),
		Description:    dcl.String(d.Get("description").(string)),
		EncryptionSpec: expandVertexAiEndpointEncryptionSpec(d.Get("encryption_spec")),
		Labels:         checkStringMap(d.Get("labels")),
		Network:        dcl.String(d.Get("network").(string)),
		Project:        dcl.String(project),
		Name:           dcl.StringOrNil(d.Get("name").(string)),
	}
	directive := UpdateDirective
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLVertexAiClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutUpdate))
	if bp, err := replaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	res, err := client.ApplyEndpoint(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error updating Endpoint: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Endpoint %q: %#v", d.Id(), res)

	return resourceVertexAiEndpointRead(d, meta)
}

func resourceVertexAiEndpointDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vertexai.Endpoint{
		DisplayName:    dcl.String(d.Get("display_name").(string)),
		Location:       dcl.String(d.Get("location").(string)),
		Description:    dcl.String(d.Get("description").(string)),
		EncryptionSpec: expandVertexAiEndpointEncryptionSpec(d.Get("encryption_spec")),
		Labels:         checkStringMap(d.Get("labels")),
		Network:        dcl.String(d.Get("network").(string)),
		Project:        dcl.String(project),
		Name:           dcl.StringOrNil(d.Get("name").(string)),
	}

	log.Printf("[DEBUG] Deleting Endpoint %q", d.Id())
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
	if err := client.DeleteEndpoint(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting Endpoint: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting Endpoint %q", d.Id())
	return nil
}

func resourceVertexAiEndpointImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/endpoints/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/endpoints/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandVertexAiEndpointEncryptionSpec(o interface{}) *vertexai.EndpointEncryptionSpec {
	if o == nil {
		return vertexai.EmptyEndpointEncryptionSpec
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return vertexai.EmptyEndpointEncryptionSpec
	}
	obj := objArr[0].(map[string]interface{})
	return &vertexai.EndpointEncryptionSpec{
		KmsKeyName: dcl.String(obj["kms_key_name"].(string)),
	}
}

func flattenVertexAiEndpointEncryptionSpec(obj *vertexai.EndpointEncryptionSpec) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"kms_key_name": obj.KmsKeyName,
	}

	return []interface{}{transformed}

}

func flattenVertexAiEndpointDeployedModelsArray(objs []vertexai.EndpointDeployedModels) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenVertexAiEndpointDeployedModels(&item)
		items = append(items, i)
	}

	return items
}

func flattenVertexAiEndpointDeployedModels(obj *vertexai.EndpointDeployedModels) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"automatic_resources":       flattenVertexAiEndpointDeployedModelsAutomaticResources(obj.AutomaticResources),
		"create_time":               obj.CreateTime,
		"dedicated_resources":       flattenVertexAiEndpointDeployedModelsDedicatedResources(obj.DedicatedResources),
		"disable_container_logging": obj.DisableContainerLogging,
		"display_name":              obj.DisplayName,
		"enable_access_logging":     obj.EnableAccessLogging,
		"enable_container_logging":  obj.EnableContainerLogging,
		"id":                        obj.Id,
		"model":                     obj.Model,
		"model_version_id":          obj.ModelVersionId,
		"private_endpoints":         flattenVertexAiEndpointDeployedModelsPrivateEndpoints(obj.PrivateEndpoints),
		"service_account":           obj.ServiceAccount,
		"shared_resources":          obj.SharedResources,
	}

	return transformed

}

func flattenVertexAiEndpointDeployedModelsAutomaticResources(obj *vertexai.EndpointDeployedModelsAutomaticResources) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"max_replica_count": obj.MaxReplicaCount,
		"min_replica_count": obj.MinReplicaCount,
	}

	return []interface{}{transformed}

}

func flattenVertexAiEndpointDeployedModelsDedicatedResources(obj *vertexai.EndpointDeployedModelsDedicatedResources) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"autoscaling_metric_specs": flattenVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsArray(obj.AutoscalingMetricSpecs),
		"machine_spec":             flattenVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpec(obj.MachineSpec),
		"max_replica_count":        obj.MaxReplicaCount,
		"min_replica_count":        obj.MinReplicaCount,
	}

	return []interface{}{transformed}

}

func flattenVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsArray(objs []vertexai.EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(&item)
		items = append(items, i)
	}

	return items
}

func flattenVertexAiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(obj *vertexai.EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"metric_name": obj.MetricName,
		"target":      obj.Target,
	}

	return transformed

}

func flattenVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpec(obj *vertexai.EndpointDeployedModelsDedicatedResourcesMachineSpec) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"accelerator_count": obj.AcceleratorCount,
		"accelerator_type":  obj.AcceleratorType,
		"machine_type":      obj.MachineType,
	}

	return []interface{}{transformed}

}

func flattenVertexAiEndpointDeployedModelsPrivateEndpoints(obj *vertexai.EndpointDeployedModelsPrivateEndpoints) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"explain_http_uri":   obj.ExplainHttpUri,
		"health_http_uri":    obj.HealthHttpUri,
		"predict_http_uri":   obj.PredictHttpUri,
		"service_attachment": obj.ServiceAttachment,
	}

	return []interface{}{transformed}

}
