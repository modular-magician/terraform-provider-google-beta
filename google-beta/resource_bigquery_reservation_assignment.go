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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	bigqueryreservation "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/beta"
)

func resourceBigqueryReservationAssignment() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigqueryReservationAssignmentCreate,
		Read:   resourceBigqueryReservationAssignmentRead,
		Delete: resourceBigqueryReservationAssignmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBigqueryReservationAssignmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"assignee": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"job_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  ``,
				ValidateFunc: validation.StringInSlice([]string{"JOB_TYPE_UNSPECIFIED", "PIPELINE", "QUERY", ""}, false),
			},

			"reservation": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},
		},
	}
}

func resourceBigqueryReservationAssignmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProjectField(d, config)
	if err != nil {
		return err
	}

	obj := &bigqueryreservation.Assignment{
		Assignee:    dcl.String(d.Get("assignee").(string)),
		JobType:     bigqueryreservation.AssignmentJobTypeEnumRef(d.Get("job_type").(string)),
		Reservation: dcl.String(d.Get("reservation").(string)),
		Location:    dcl.String(d.Get("location").(string)),
		Project:     dcl.String(project),
	}

	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	createDirective := CreateDirective
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLBigqueryReservationClient(config, userAgent, billingProject)
	res, err := client.ApplyAssignment(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating Assignment: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Assignment %q: %#v", d.Id(), res)

	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	// Id has a server-generated value, set again after creation
	id, err = replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return resourceBigqueryReservationAssignmentRead(d, meta)
}

func resourceBigqueryReservationAssignmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProjectField(d, config)
	if err != nil {
		return err
	}

	obj := &bigqueryreservation.Assignment{
		Assignee:    dcl.String(d.Get("assignee").(string)),
		JobType:     bigqueryreservation.AssignmentJobTypeEnumRef(d.Get("job_type").(string)),
		Reservation: dcl.String(d.Get("reservation").(string)),
		Location:    dcl.String(d.Get("location").(string)),
		Project:     dcl.String(project),
		Name:        dcl.StringOrNil(d.Get("name").(string)),
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
	client := NewDCLBigqueryReservationClient(config, userAgent, billingProject)
	res, err := client.GetAssignment(context.Background(), obj)
	if err != nil {
		// Resource not found
		d.SetId("")
		return err
	}

	if err = d.Set("assignee", res.Assignee); err != nil {
		return fmt.Errorf("error setting assignee in state: %s", err)
	}
	if err = d.Set("job_type", res.JobType); err != nil {
		return fmt.Errorf("error setting job_type in state: %s", err)
	}
	if err = d.Set("reservation", res.Reservation); err != nil {
		return fmt.Errorf("error setting reservation in state: %s", err)
	}
	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("state", res.State); err != nil {
		return fmt.Errorf("error setting state in state: %s", err)
	}

	return nil
}

func resourceBigqueryReservationAssignmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProjectField(d, config)
	if err != nil {
		return err
	}

	obj := &bigqueryreservation.Assignment{
		Assignee:    dcl.String(d.Get("assignee").(string)),
		JobType:     bigqueryreservation.AssignmentJobTypeEnumRef(d.Get("job_type").(string)),
		Reservation: dcl.String(d.Get("reservation").(string)),
		Location:    dcl.String(d.Get("location").(string)),
		Project:     dcl.String(project),
		Name:        dcl.StringOrNil(d.Get("name").(string)),
	}

	log.Printf("[DEBUG] Deleting Assignment %q", d.Id())
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLBigqueryReservationClient(config, userAgent, billingProject)
	if err := client.DeleteAssignment(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting Assignment: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting Assignment %q", d.Id())
	return nil
}

func resourceBigqueryReservationAssignmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/reservations/(?P<reservation>[^/]+)/assignments/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<reservation>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<reservation>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}
