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

func resourceVertexAiModel() *schema.Resource {
	return &schema.Resource{
		Create: resourceVertexAiModelCreate,
		Read:   resourceVertexAiModelRead,
		Update: resourceVertexAiModelUpdate,
		Delete: resourceVertexAiModelDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVertexAiModelImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"container_spec": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "The specification of the container that is to be used when deploying this Model. The specification is ingested upon ModelService.UploadModel, and all binaries it contains are copied and stored internally by Vertex AI. Not present for AutoML Models.",
				MaxItems:    1,
				Elem:        VertexAiModelContainerSpecSchema(),
			},

			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Required. The display name of the Model. The name can be up to 128 characters long and can be consist of any UTF-8 characters.",
			},

			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The location for the resource",
			},

			"artifact_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Immutable. The path to the directory containing the Model artifact and any of its supporting files. Not present for AutoML Models.",
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the Model.",
			},

			"encryption_spec": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Customer-managed encryption key spec for a Model. If set, this Model and all sub-resources of this Model will be secured by this key.",
				MaxItems:    1,
				Elem:        VertexAiModelEncryptionSpecSchema(),
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The labels with user-defined metadata to organize your Models. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "The project for the resource",
			},

			"version_aliases": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "User provided version aliases so that a model version can be referenced via alias (i.e. projects/{project}/locations/{location}/models/{model_id}@{version_alias} instead of auto-generated version id (i.e. projects/{project}/locations/{location}/models/{model_id}@{version_id}). The format is a-z{0,126}[a-z0-9] to distinguish from version_id. A default version alias will be created for the first version of the model, and there must be exactly one default version alias for a model.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"version_description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The description of this version.",
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Timestamp when this Model was uploaded into Vertex AI.",
			},

			"deployed_models": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. The pointers to DeployedModels created from this Model. Note that Model could have been deployed to Endpoints in different Locations.",
				Elem:        VertexAiModelDeployedModelsSchema(),
			},

			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Used to perform consistent read-modify-write updates. If not set, a blind \"overwrite\" update happens.",
			},

			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource name of the Model.",
			},

			"original_model_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. If this Model is a copy of another Model, this contains info about the original.",
				Elem:        VertexAiModelOriginalModelInfoSchema(),
			},

			"supported_deployment_resources_types": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. When this Model is deployed, its prediction resources are described by the `prediction_resources` field of the Endpoint.deployed_models object. Because not all Models support all resource configuration types, the configuration types this Model supports are listed here. If no configuration types are listed, the Model cannot be deployed to an Endpoint and does not support online predictions (PredictionService.Predict or PredictionService.Explain). Such a Model can serve predictions by using a BatchPredictionJob, if it has at least one entry each in supported_input_storage_formats and supported_output_storage_formats.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"supported_export_formats": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. The formats in which this Model may be exported. If empty, this Model is not available for export.",
				Elem:        VertexAiModelSupportedExportFormatsSchema(),
			},

			"supported_input_storage_formats": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. The formats this Model supports in BatchPredictionJob.input_config. If PredictSchemata.instance_schema_uri exists, the instances should be given as per that schema. The possible formats are: * `jsonl` The JSON Lines format, where each instance is a single line. Uses GcsSource. * `csv` The CSV format, where each instance is a single comma-separated line. The first line in the file is the header, containing comma-separated field names. Uses GcsSource. * `tf-record` The TFRecord format, where each instance is a single record in tfrecord syntax. Uses GcsSource. * `tf-record-gzip` Similar to `tf-record`, but the file is gzipped. Uses GcsSource. * `bigquery` Each instance is a single row in BigQuery. Uses BigQuerySource. * `file-list` Each line of the file is the location of an instance to process, uses `gcs_source` field of the InputConfig object. If this Model doesn't support any of these formats it means it cannot be used with a BatchPredictionJob. However, if it has supported_deployment_resources_types, it could serve online predictions by using PredictionService.Predict or PredictionService.Explain. TODO(rsurowka): Give a link describing how OpenAPI schema instances are expressed in JSONL and BigQuery. TODO(rsurowka): Should we provide a schema for TFRecord? Or maybe say that at least for now TFRecord input is not supported via schemata (that would also simplify giving them back as part of predictions). TODO(rsurowka): Define CSV format (decide how much we want to support). E.g. no nesting? Or no arrays, or no nested arrays? E.g. https://json-csv.com/ seems to be able to do pretty advanced conversions, but we may decide to make it relatively simple for now.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"supported_output_storage_formats": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. The formats this Model supports in BatchPredictionJob.output_config. If both PredictSchemata.instance_schema_uri and PredictSchemata.prediction_schema_uri exist, the predictions are returned together with their instances. In other words, the prediction has the original instance data first, followed by the actual prediction content (as per the schema). The possible formats are: * `jsonl` The JSON Lines format, where each prediction is a single line. Uses GcsDestination. * `csv` The CSV format, where each prediction is a single comma-separated line. The first line in the file is the header, containing comma-separated field names. Uses GcsDestination. * `bigquery` Each prediction is a single row in a BigQuery table, uses BigQueryDestination . If this Model doesn't support any of these formats it means it cannot be used with a BatchPredictionJob. However, if it has supported_deployment_resources_types, it could serve online predictions by using PredictionService.Predict or PredictionService.Explain. TODO(rsurowka): Analogous TODOs as for instances field above.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"training_pipeline": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The resource name of the TrainingPipeline that uploaded this Model, if any.",
			},

			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Timestamp when this Model was most recently updated.",
			},

			"version_create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Timestamp when this version was created.",
			},

			"version_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Immutable. The version ID of the model. A new version is committed when a new model version is uploaded or trained under an existing model id. It is an auto-incrementing decimal number in string representation.",
			},

			"version_update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Timestamp when this version was most recently updated.",
			},
		},
	}
}

func VertexAiModelContainerSpecSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image_uri": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Required. Immutable. URI of the Docker image to be used as the custom container for serving predictions. This URI must identify an image in Artifact Registry or Container Registry. Learn more about the [container publishing requirements](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#publishing), including permissions requirements for the Vertex AI Service Agent. The container image is ingested upon ModelService.UploadModel, stored internally, and this original path is afterwards not used. To learn about the requirements for the Docker image itself, see [Custom container requirements](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#). You can use the URI to one of Vertex AI's [pre-built container images for prediction](https://cloud.google.com/vertex-ai/docs/predictions/pre-built-containers) in this field.",
			},

			"args": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Immutable. Specifies arguments for the command that runs when the container starts. This overrides the container's [`CMD`](https://docs.docker.com/engine/reference/builder/#cmd). Specify this field as an array of executable and arguments, similar to a Docker `CMD`'s \"default parameters\" form. If you don't specify this field but do specify the command field, then the command from the `command` field runs without any additional arguments. See the [Kubernetes documentation about how the `command` and `args` fields interact with a container's `ENTRYPOINT` and `CMD`](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#notes). If you don't specify this field and don't specify the `command` field, then the container's [`ENTRYPOINT`](https://docs.docker.com/engine/reference/builder/#cmd) and `CMD` determine what runs based on their default behavior. See the Docker documentation about [how `CMD` and `ENTRYPOINT` interact](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). In this field, you can reference [environment variables set by Vertex AI](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables) and environment variables set in the env field. You cannot reference environment variables set in the Docker image. In order for environment variables to be expanded, reference them by using the following syntax: `$(VARIABLE_NAME)` Note that this differs from Bash variable expansion, which does not use parentheses. If a variable cannot be resolved, the reference in the input string is used unchanged. To avoid variable expansion, you can escape this syntax with `$$`; for example: `$$(VARIABLE_NAME)` This field corresponds to the `args` field of the Kubernetes Containers [v1 core API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"command": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Immutable. Specifies the command that runs when the container starts. This overrides the container's [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint). Specify this field as an array of executable and arguments, similar to a Docker `ENTRYPOINT`'s \"exec\" form, not its \"shell\" form. If you do not specify this field, then the container's `ENTRYPOINT` runs, in conjunction with the args field or the container's [`CMD`](https://docs.docker.com/engine/reference/builder/#cmd), if either exists. If this field is not specified and the container does not have an `ENTRYPOINT`, then refer to the Docker documentation about [how `CMD` and `ENTRYPOINT` interact](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). If you specify this field, then you can also specify the `args` field to provide additional arguments for this command. However, if you specify this field, then the container's `CMD` is ignored. See the [Kubernetes documentation about how the `command` and `args` fields interact with a container's `ENTRYPOINT` and `CMD`](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#notes). In this field, you can reference [environment variables set by Vertex AI](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables) and environment variables set in the env field. You cannot reference environment variables set in the Docker image. In order for environment variables to be expanded, reference them by using the following syntax: `$(VARIABLE_NAME)` Note that this differs from Bash variable expansion, which does not use parentheses. If a variable cannot be resolved, the reference in the input string is used unchanged. To avoid variable expansion, you can escape this syntax with `$$`; for example: `$$(VARIABLE_NAME)` This field corresponds to the `command` field of the Kubernetes Containers [v1 core API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"env": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Immutable. List of environment variables to set in the container. After the container starts running, code running in the container can read these environment variables. Additionally, the command and args fields can reference these variables. Later entries in this list can also reference earlier entries. For example, the following example sets the variable `VAR_2` to have the value `foo bar`: ```json [ { \"name\": \"VAR_1\", \"value\": \"foo\" }, { \"name\": \"VAR_2\", \"value\": \"$(VAR_1) bar\" } ] ``` If you switch the order of the variables in the example, then the expansion does not occur. This field corresponds to the `env` field of the Kubernetes Containers [v1 core API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).",
				Elem:        VertexAiModelContainerSpecEnvSchema(),
			},

			"health_route": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Immutable. HTTP path on the container to send health checks to. Vertex AI intermittently sends GET requests to this path on the container's IP address and port to check that the container is healthy. Read more about [health checks](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#health). For example, if you set this field to `/bar`, then Vertex AI intermittently sends a GET request to the `/bar` path on the port of your container specified by the first value of this `ModelContainerSpec`'s ports field. If you don't specify this field, it defaults to the following value when you deploy this Model to an Endpoint: `/v1/endpoints/ENDPOINT/deployedModels/DEPLOYED_MODEL:predict` The placeholders in this value are replaced as follows: * ENDPOINT: The last segment (following `endpoints/`)of the Endpoint.name][] field of the Endpoint where this Model has been deployed. (Vertex AI makes this value available to your container code as the [`AIP_ENDPOINT_ID` environment variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).) * DEPLOYED_MODEL: DeployedModel.id of the `DeployedModel`. (Vertex AI makes this value available to your container code as the [`AIP_DEPLOYED_MODEL_ID` environment variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)",
			},

			"ports": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Immutable. List of ports to expose from the container. Vertex AI sends any prediction requests that it receives to the first port on this list. Vertex AI also sends [liveness and health checks](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#liveness) to this port. If you do not specify this field, it defaults to following value: ```json [ { \"containerPort\": 8080 } ] ``` Vertex AI does not use ports other than the first one listed. This field corresponds to the `ports` field of the Kubernetes Containers [v1 core API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).",
				Elem:        VertexAiModelContainerSpecPortsSchema(),
			},

			"predict_route": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Immutable. HTTP path on the container to send prediction requests to. Vertex AI forwards requests sent using projects.locations.endpoints.predict to this path on the container's IP address and port. Vertex AI then returns the container's response in the API response. For example, if you set this field to `/foo`, then when Vertex AI receives a prediction request, it forwards the request body in a POST request to the `/foo` path on the port of your container specified by the first value of this `ModelContainerSpec`'s ports field. If you don't specify this field, it defaults to the following value when you deploy this Model to an Endpoint: `/v1/endpoints/ENDPOINT/deployedModels/DEPLOYED_MODEL:predict` The placeholders in this value are replaced as follows: * ENDPOINT: The last segment (following `endpoints/`)of the Endpoint.name][] field of the Endpoint where this Model has been deployed. (Vertex AI makes this value available to your container code as the [`AIP_ENDPOINT_ID` environment variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).) * DEPLOYED_MODEL: DeployedModel.id of the `DeployedModel`. (Vertex AI makes this value available to your container code as the [`AIP_DEPLOYED_MODEL_ID` environment variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)",
			},
		},
	}
}

func VertexAiModelContainerSpecEnvSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Required. Name of the environment variable. Must be a valid C identifier.",
			},

			"value": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Required. Variables that reference a $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not.",
			},
		},
	}
}

func VertexAiModelContainerSpecPortsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"container_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The number of the port to expose on the pod's IP address. Must be a valid port number, between 1 and 65535 inclusive.",
			},
		},
	}
}

func VertexAiModelEncryptionSpecSchema() *schema.Resource {
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

func VertexAiModelDeployedModelsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deployed_model_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Immutable. An ID of a DeployedModel in the above Endpoint.",
			},

			"endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Immutable. A resource name of an Endpoint.",
			},
		},
	}
}

func VertexAiModelOriginalModelInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"model": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The resource name of the Model this Model is a copy of, including the revision. Format: `projects/{project}/locations/{location}/models/{model_id}@{version_id}`",
			},
		},
	}
}

func VertexAiModelSupportedExportFormatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"exportable_contents": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. The content of this Model that may be exported.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The ID of the export format. The possible format IDs are: * `tflite` Used for Android mobile devices. * `edgetpu-tflite` Used for [Edge TPU](https://cloud.google.com/edge-tpu/) devices. * `tf-saved-model` A tensorflow model in SavedModel format. * `tf-js` A [TensorFlow.js](https://www.tensorflow.org/js) model that can be used in the browser and in Node.js using JavaScript. * `core-ml` Used for iOS mobile devices. * `custom-trained` A Model that was uploaded or trained by custom code.",
			},
		},
	}
}

func resourceVertexAiModelCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vertexai.Model{
		ContainerSpec:      expandVertexAiModelContainerSpec(d.Get("container_spec")),
		DisplayName:        dcl.String(d.Get("display_name").(string)),
		Location:           dcl.String(d.Get("location").(string)),
		ArtifactUri:        dcl.String(d.Get("artifact_uri").(string)),
		Description:        dcl.String(d.Get("description").(string)),
		EncryptionSpec:     expandVertexAiModelEncryptionSpec(d.Get("encryption_spec")),
		Labels:             checkStringMap(d.Get("labels")),
		Project:            dcl.String(project),
		VersionAliases:     expandStringArray(d.Get("version_aliases")),
		VersionDescription: dcl.String(d.Get("version_description").(string)),
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
	res, err := client.ApplyModel(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating Model: %s", err)
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

	log.Printf("[DEBUG] Finished creating Model %q: %#v", d.Id(), res)

	return resourceVertexAiModelRead(d, meta)
}

func resourceVertexAiModelRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vertexai.Model{
		ContainerSpec:      expandVertexAiModelContainerSpec(d.Get("container_spec")),
		DisplayName:        dcl.String(d.Get("display_name").(string)),
		Location:           dcl.String(d.Get("location").(string)),
		ArtifactUri:        dcl.String(d.Get("artifact_uri").(string)),
		Description:        dcl.String(d.Get("description").(string)),
		EncryptionSpec:     expandVertexAiModelEncryptionSpec(d.Get("encryption_spec")),
		Labels:             checkStringMap(d.Get("labels")),
		Project:            dcl.String(project),
		VersionAliases:     expandStringArray(d.Get("version_aliases")),
		VersionDescription: dcl.String(d.Get("version_description").(string)),
		Name:               dcl.StringOrNil(d.Get("name").(string)),
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
	res, err := client.GetModel(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("VertexAiModel %q", d.Id())
		return handleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("container_spec", flattenVertexAiModelContainerSpec(res.ContainerSpec)); err != nil {
		return fmt.Errorf("error setting container_spec in state: %s", err)
	}
	if err = d.Set("display_name", res.DisplayName); err != nil {
		return fmt.Errorf("error setting display_name in state: %s", err)
	}
	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("artifact_uri", res.ArtifactUri); err != nil {
		return fmt.Errorf("error setting artifact_uri in state: %s", err)
	}
	if err = d.Set("description", res.Description); err != nil {
		return fmt.Errorf("error setting description in state: %s", err)
	}
	if err = d.Set("encryption_spec", flattenVertexAiModelEncryptionSpec(res.EncryptionSpec)); err != nil {
		return fmt.Errorf("error setting encryption_spec in state: %s", err)
	}
	if err = d.Set("labels", res.Labels); err != nil {
		return fmt.Errorf("error setting labels in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("version_aliases", res.VersionAliases); err != nil {
		return fmt.Errorf("error setting version_aliases in state: %s", err)
	}
	if err = d.Set("version_description", res.VersionDescription); err != nil {
		return fmt.Errorf("error setting version_description in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("deployed_models", flattenVertexAiModelDeployedModelsArray(res.DeployedModels)); err != nil {
		return fmt.Errorf("error setting deployed_models in state: %s", err)
	}
	if err = d.Set("etag", res.Etag); err != nil {
		return fmt.Errorf("error setting etag in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("original_model_info", flattenVertexAiModelOriginalModelInfo(res.OriginalModelInfo)); err != nil {
		return fmt.Errorf("error setting original_model_info in state: %s", err)
	}
	if err = d.Set("supported_deployment_resources_types", flattenVertexAiModelSupportedDeploymentResourcesTypesArray(res.SupportedDeploymentResourcesTypes)); err != nil {
		return fmt.Errorf("error setting supported_deployment_resources_types in state: %s", err)
	}
	if err = d.Set("supported_export_formats", flattenVertexAiModelSupportedExportFormatsArray(res.SupportedExportFormats)); err != nil {
		return fmt.Errorf("error setting supported_export_formats in state: %s", err)
	}
	if err = d.Set("supported_input_storage_formats", res.SupportedInputStorageFormats); err != nil {
		return fmt.Errorf("error setting supported_input_storage_formats in state: %s", err)
	}
	if err = d.Set("supported_output_storage_formats", res.SupportedOutputStorageFormats); err != nil {
		return fmt.Errorf("error setting supported_output_storage_formats in state: %s", err)
	}
	if err = d.Set("training_pipeline", res.TrainingPipeline); err != nil {
		return fmt.Errorf("error setting training_pipeline in state: %s", err)
	}
	if err = d.Set("update_time", res.UpdateTime); err != nil {
		return fmt.Errorf("error setting update_time in state: %s", err)
	}
	if err = d.Set("version_create_time", res.VersionCreateTime); err != nil {
		return fmt.Errorf("error setting version_create_time in state: %s", err)
	}
	if err = d.Set("version_id", res.VersionId); err != nil {
		return fmt.Errorf("error setting version_id in state: %s", err)
	}
	if err = d.Set("version_update_time", res.VersionUpdateTime); err != nil {
		return fmt.Errorf("error setting version_update_time in state: %s", err)
	}

	return nil
}
func resourceVertexAiModelUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vertexai.Model{
		ContainerSpec:      expandVertexAiModelContainerSpec(d.Get("container_spec")),
		DisplayName:        dcl.String(d.Get("display_name").(string)),
		Location:           dcl.String(d.Get("location").(string)),
		ArtifactUri:        dcl.String(d.Get("artifact_uri").(string)),
		Description:        dcl.String(d.Get("description").(string)),
		EncryptionSpec:     expandVertexAiModelEncryptionSpec(d.Get("encryption_spec")),
		Labels:             checkStringMap(d.Get("labels")),
		Project:            dcl.String(project),
		VersionAliases:     expandStringArray(d.Get("version_aliases")),
		VersionDescription: dcl.String(d.Get("version_description").(string)),
		Name:               dcl.StringOrNil(d.Get("name").(string)),
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
	res, err := client.ApplyModel(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error updating Model: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Model %q: %#v", d.Id(), res)

	return resourceVertexAiModelRead(d, meta)
}

func resourceVertexAiModelDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vertexai.Model{
		ContainerSpec:      expandVertexAiModelContainerSpec(d.Get("container_spec")),
		DisplayName:        dcl.String(d.Get("display_name").(string)),
		Location:           dcl.String(d.Get("location").(string)),
		ArtifactUri:        dcl.String(d.Get("artifact_uri").(string)),
		Description:        dcl.String(d.Get("description").(string)),
		EncryptionSpec:     expandVertexAiModelEncryptionSpec(d.Get("encryption_spec")),
		Labels:             checkStringMap(d.Get("labels")),
		Project:            dcl.String(project),
		VersionAliases:     expandStringArray(d.Get("version_aliases")),
		VersionDescription: dcl.String(d.Get("version_description").(string)),
		Name:               dcl.StringOrNil(d.Get("name").(string)),
	}

	log.Printf("[DEBUG] Deleting Model %q", d.Id())
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
	if err := client.DeleteModel(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting Model: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting Model %q", d.Id())
	return nil
}

func resourceVertexAiModelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/models/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/models/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandVertexAiModelContainerSpec(o interface{}) *vertexai.ModelContainerSpec {
	if o == nil {
		return vertexai.EmptyModelContainerSpec
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return vertexai.EmptyModelContainerSpec
	}
	obj := objArr[0].(map[string]interface{})
	return &vertexai.ModelContainerSpec{
		ImageUri:     dcl.String(obj["image_uri"].(string)),
		Args:         expandStringArray(obj["args"]),
		Command:      expandStringArray(obj["command"]),
		Env:          expandVertexAiModelContainerSpecEnvArray(obj["env"]),
		HealthRoute:  dcl.String(obj["health_route"].(string)),
		Ports:        expandVertexAiModelContainerSpecPortsArray(obj["ports"]),
		PredictRoute: dcl.String(obj["predict_route"].(string)),
	}
}

func flattenVertexAiModelContainerSpec(obj *vertexai.ModelContainerSpec) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"image_uri":     obj.ImageUri,
		"args":          obj.Args,
		"command":       obj.Command,
		"env":           flattenVertexAiModelContainerSpecEnvArray(obj.Env),
		"health_route":  obj.HealthRoute,
		"ports":         flattenVertexAiModelContainerSpecPortsArray(obj.Ports),
		"predict_route": obj.PredictRoute,
	}

	return []interface{}{transformed}

}
func expandVertexAiModelContainerSpecEnvArray(o interface{}) []vertexai.ModelContainerSpecEnv {
	if o == nil {
		return make([]vertexai.ModelContainerSpecEnv, 0)
	}

	objs := o.([]interface{})
	if len(objs) == 0 || objs[0] == nil {
		return make([]vertexai.ModelContainerSpecEnv, 0)
	}

	items := make([]vertexai.ModelContainerSpecEnv, 0, len(objs))
	for _, item := range objs {
		i := expandVertexAiModelContainerSpecEnv(item)
		items = append(items, *i)
	}

	return items
}

func expandVertexAiModelContainerSpecEnv(o interface{}) *vertexai.ModelContainerSpecEnv {
	if o == nil {
		return vertexai.EmptyModelContainerSpecEnv
	}

	obj := o.(map[string]interface{})
	return &vertexai.ModelContainerSpecEnv{
		Name:  dcl.String(obj["name"].(string)),
		Value: dcl.String(obj["value"].(string)),
	}
}

func flattenVertexAiModelContainerSpecEnvArray(objs []vertexai.ModelContainerSpecEnv) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenVertexAiModelContainerSpecEnv(&item)
		items = append(items, i)
	}

	return items
}

func flattenVertexAiModelContainerSpecEnv(obj *vertexai.ModelContainerSpecEnv) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"name":  obj.Name,
		"value": obj.Value,
	}

	return transformed

}
func expandVertexAiModelContainerSpecPortsArray(o interface{}) []vertexai.ModelContainerSpecPorts {
	if o == nil {
		return make([]vertexai.ModelContainerSpecPorts, 0)
	}

	objs := o.([]interface{})
	if len(objs) == 0 || objs[0] == nil {
		return make([]vertexai.ModelContainerSpecPorts, 0)
	}

	items := make([]vertexai.ModelContainerSpecPorts, 0, len(objs))
	for _, item := range objs {
		i := expandVertexAiModelContainerSpecPorts(item)
		items = append(items, *i)
	}

	return items
}

func expandVertexAiModelContainerSpecPorts(o interface{}) *vertexai.ModelContainerSpecPorts {
	if o == nil {
		return vertexai.EmptyModelContainerSpecPorts
	}

	obj := o.(map[string]interface{})
	return &vertexai.ModelContainerSpecPorts{
		ContainerPort: dcl.Int64(int64(obj["container_port"].(int))),
	}
}

func flattenVertexAiModelContainerSpecPortsArray(objs []vertexai.ModelContainerSpecPorts) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenVertexAiModelContainerSpecPorts(&item)
		items = append(items, i)
	}

	return items
}

func flattenVertexAiModelContainerSpecPorts(obj *vertexai.ModelContainerSpecPorts) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"container_port": obj.ContainerPort,
	}

	return transformed

}

func expandVertexAiModelEncryptionSpec(o interface{}) *vertexai.ModelEncryptionSpec {
	if o == nil {
		return vertexai.EmptyModelEncryptionSpec
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return vertexai.EmptyModelEncryptionSpec
	}
	obj := objArr[0].(map[string]interface{})
	return &vertexai.ModelEncryptionSpec{
		KmsKeyName: dcl.String(obj["kms_key_name"].(string)),
	}
}

func flattenVertexAiModelEncryptionSpec(obj *vertexai.ModelEncryptionSpec) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"kms_key_name": obj.KmsKeyName,
	}

	return []interface{}{transformed}

}

func flattenVertexAiModelDeployedModelsArray(objs []vertexai.ModelDeployedModels) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenVertexAiModelDeployedModels(&item)
		items = append(items, i)
	}

	return items
}

func flattenVertexAiModelDeployedModels(obj *vertexai.ModelDeployedModels) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"deployed_model_id": obj.DeployedModelId,
		"endpoint":          obj.Endpoint,
	}

	return transformed

}

func flattenVertexAiModelOriginalModelInfo(obj *vertexai.ModelOriginalModelInfo) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"model": obj.Model,
	}

	return []interface{}{transformed}

}

func flattenVertexAiModelSupportedExportFormatsArray(objs []vertexai.ModelSupportedExportFormats) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenVertexAiModelSupportedExportFormats(&item)
		items = append(items, i)
	}

	return items
}

func flattenVertexAiModelSupportedExportFormats(obj *vertexai.ModelSupportedExportFormats) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"exportable_contents": flattenVertexAiModelSupportedExportFormatsExportableContentsArray(obj.ExportableContents),
		"id":                  obj.Id,
	}

	return transformed

}
func flattenVertexAiModelSupportedDeploymentResourcesTypesArray(obj []vertexai.ModelSupportedDeploymentResourcesTypesEnum) interface{} {
	if obj == nil {
		return nil
	}
	items := []string{}
	for _, item := range obj {
		items = append(items, string(item))
	}
	return items
}
func flattenVertexAiModelSupportedExportFormatsExportableContentsArray(obj []vertexai.ModelSupportedExportFormatsExportableContentsEnum) interface{} {
	if obj == nil {
		return nil
	}
	items := []string{}
	for _, item := range obj {
		items = append(items, string(item))
	}
	return items
}
