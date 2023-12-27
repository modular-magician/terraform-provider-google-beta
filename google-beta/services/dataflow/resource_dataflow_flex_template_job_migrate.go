// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package dataflow

import (
	"context"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDataflowFlexTemplateJobResourceV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			"container_spec_gcs_path": {
				Type:     schema.TypeString,
				Required: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: `The region in which the created job should run.`,
			},

			"transform_name_mapping": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job.`,
			},

			"launch_options": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Options to be used to configure Flex Template launcher instance.`,
			},

			"on_delete": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"cancel", "drain"}, false),
				Optional:     true,
				Default:      "cancel",
			},

			"labels": {
				Type:             schema.TypeMap,
				Optional:         true,
				DiffSuppressFunc: resourceDataflowJobLabelDiffSuppress,
				Description:      `User labels to be specified for the job. Keys and values should follow the restrictions specified in the labeling restrictions page. NOTE: Google-provided Dataflow templates often provide default labels that begin with goog-dataflow-provided. Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.`,
			},

			"parameters": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"job_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The type of this job, selected from the JobType enum.`,
			},

			"num_workers": {
				Type:     schema.TypeInt,
				Optional: true,
				// ForceNew applies to both stream and batch jobs
				ForceNew:    true,
				Description: `The initial number of Google Compute Engine instances for the job.`,
			},

			"max_workers": {
				Type:     schema.TypeInt,
				Optional: true,
				// ForceNew applies to both stream and batch jobs
				ForceNew:    true,
				Description: `The maximum number of Google Compute Engine instances to be made available to your pipeline during execution, from 1 to 1000.`,
			},

			"service_account_email": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The Service Account email used to create the job.`,
			},

			"temp_location": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The Cloud Storage path to use for temporary files. Must be a valid Cloud Storage URL, beginning with gs://.`,
			},

			"staging_location": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The Cloud Storage path to use for staging files. Must be a valid Cloud Storage URL, beginning with gs://.`,
			},

			"sdk_container_image": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Docker registry location of container image to use for the 'worker harness. Default is the container for the version of the SDK. Note this field is only valid for portable pipelines.`,
			},

			"network": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The network to which VMs will be assigned. If it is not provided, "default" will be used.`,
			},

			"subnetwork": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".`,
			},

			"machine_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The machine type to use for the job.`,
			},

			"kms_key_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The name for the Cloud KMS key for the job. Key format is: projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY`,
			},

			"ip_configuration": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  `The configuration for VM IPs. Options are "WORKER_IP_PUBLIC" or "WORKER_IP_PRIVATE".`,
				ValidateFunc: validation.StringInSlice([]string{"WORKER_IP_PUBLIC", "WORKER_IP_PRIVATE"}, false),
			},

			"additional_experiments": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: `List of experiments that should be used by the job. An example value is ["enable_stackdriver_agent_metrics"].`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"enable_streaming_engine": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `Indicates if the job should use the streaming engine feature.`,
			},

			"autoscaling_algorithm": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The algorithm to use for autoscaling`,
			},

			"launcher_machine_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The machine type to use for launching the job. The default is n1-standard-1.`,
			},

			"skip_wait_on_job_termination": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: `If true, treat DRAINING and CANCELLING as terminal job states and do not wait for further changes before removing from terraform state and moving on. WARNING: this will lead to job name conflicts if you do not ensure that the job names are different, e.g. by embedding a release ID or by using a random_id.`,
			},
		},
		UseJSONNumber: true,
	}
}

func ResourceDataflowFlexTemplateJobStateUpgradeV0(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	return tpgresource.LabelsStateUpgrade(rawState, resourceDataflowJobGoogleLabelPrefix)
}
