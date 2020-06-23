package google

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	dataflow "google.golang.org/api/dataflow/v1b3"
)

// NOTE: resource_dataflow_flex_template currently does not support updating existing jobs.
// Changing any non-computed field will result in the job being deleted (according to its
// on_delete policy) and recreated with the updated parameters.

// resourceDataflowFlexTemplateJob defines the schema for Dataflow FlexTemplate jobs.
func resourceDataflowFlexTemplateJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataflowFlexTemplateJobCreate,
		Read:   resourceDataflowFlexTemplateJobRead,
		Update: resourceDataflowFlexTemplateJobUpdateByReplacement,
		Delete: resourceDataflowJobDelete,
		Schema: map[string]*schema.Schema{

			"container_spec_gcs_path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"on_delete": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"cancel", "drain"}, false),
				Optional:     true,
				Default:      "drain",
			},

			"labels": {
				Type:             schema.TypeMap,
				Optional:         true,
				DiffSuppressFunc: resourceDataflowJobLabelDiffSuppress,
				ForceNew:         true,
			},

			"parameters": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},

			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// ForceNew applies to both stream and batch jobs
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
		},
	}
}

// resourceDataflowFlexTemplateJobCreate creates a Flex Template Job from TF code.
func resourceDataflowFlexTemplateJobCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	region, err := getRegion(d, config)
	if err != nil {
		return err
	}

	request := dataflow.LaunchFlexTemplateRequest{
		LaunchParameter: &dataflow.LaunchFlexTemplateParameter{
			ContainerSpecGcsPath: d.Get("container_spec_gcs_path").(string),
			JobName:              d.Get("name").(string),
			Parameters:           expandStringMap(d, "parameters"),
		},
	}
	response, err := config.clientDataflow.Projects.Locations.FlexTemplates.Launch(project, region, &request).Do()
	if err != nil {
		return err
	}
	job := response.Job
	d.SetId(job.Id)
	d.Set("job_id", job.Id)

	return resourceDataflowFlexTemplateJobRead(d, meta)
}

// resourceDataflowFlexTemplateJobRead reads a Flex Template Job resource.
func resourceDataflowFlexTemplateJobRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	region, err := getRegion(d, config)
	if err != nil {
		return err
	}

	jobId := d.Id()

	job, err := resourceDataflowJobGetJob(config, project, region, jobId)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("Dataflow job %s", jobId))
	}

	d.Set("state", job.CurrentState)
	d.Set("name", job.Name)
	d.Set("project", project)
	d.Set("labels", job.Labels)

	if _, ok := dataflowTerminalStatesMap[job.CurrentState]; ok {
		log.Printf("[DEBUG] Removing resource '%s' because it is in state %s.\n", job.Name, job.CurrentState)
		d.SetId("")
		return nil
	}
	d.SetId(job.Id)
	d.Set("job_id", job.Id)

	return nil
}

// resourceDataflowFlexTemplateJobUpdateByReplacement will be the method for updating Flex-Template jobs
func resourceDataflowFlexTemplateJobUpdateByReplacement(d *schema.ResourceData, meta interface{}) error {
	return nil
}
