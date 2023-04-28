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
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceBigqueryReservationCapacityCommitment() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigqueryReservationCapacityCommitmentCreate,
		Read:   resourceBigqueryReservationCapacityCommitmentRead,
		Update: resourceBigqueryReservationCapacityCommitmentUpdate,
		Delete: resourceBigqueryReservationCapacityCommitmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBigqueryReservationCapacityCommitmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"plan": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Capacity commitment plan. Valid values are at https://cloud.google.com/bigquery/docs/reference/reservations/rpc/google.cloud.bigquery.reservation.v1#commitmentplan`,
			},
			"slot_count": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: `Number of slots in this commitment.`,
			},
			"capacity_commitment_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The optional capacity commitment ID. Capacity commitment name will be generated automatically if this field is
empty. This field must only contain lower case alphanumeric characters or dashes. The first and last character
cannot be a dash. Max length is 64 characters. NOTE: this ID won't be kept if the capacity commitment is split
or merged.`,
			},
			"edition": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS`,
			},
			"enforce_single_admin_project_per_org": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `If true, fail the request if another project in the organization has a capacity commitment.`,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The geographic location where the transfer config should reside.
Examples: US, EU, asia-northeast1. The default value is US.`,
				Default: "US",
			},
			"renewal_plan": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The plan this capacity commitment is converted to after commitmentEndTime passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable some commitment plans.`,
			},
			"commitment_end_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The start of the current commitment period. It is applicable only for ACTIVE capacity commitments.`,
			},
			"commitment_start_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The start of the current commitment period. It is applicable only for ACTIVE capacity commitments.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the capacity commitment, e.g., projects/myproject/locations/US/capacityCommitments/123`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `State of the commitment`,
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

func resourceBigqueryReservationCapacityCommitmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	slotCountProp, err := expandBigqueryReservationCapacityCommitmentSlotCount(d.Get("slot_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("slot_count"); !isEmptyValue(reflect.ValueOf(slotCountProp)) && (ok || !reflect.DeepEqual(v, slotCountProp)) {
		obj["slotCount"] = slotCountProp
	}
	planProp, err := expandBigqueryReservationCapacityCommitmentPlan(d.Get("plan"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("plan"); !isEmptyValue(reflect.ValueOf(planProp)) && (ok || !reflect.DeepEqual(v, planProp)) {
		obj["plan"] = planProp
	}
	renewalPlanProp, err := expandBigqueryReservationCapacityCommitmentRenewalPlan(d.Get("renewal_plan"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("renewal_plan"); !isEmptyValue(reflect.ValueOf(renewalPlanProp)) && (ok || !reflect.DeepEqual(v, renewalPlanProp)) {
		obj["renewalPlan"] = renewalPlanProp
	}
	editionProp, err := expandBigqueryReservationCapacityCommitmentEdition(d.Get("edition"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("edition"); !isEmptyValue(reflect.ValueOf(editionProp)) && (ok || !reflect.DeepEqual(v, editionProp)) {
		obj["edition"] = editionProp
	}

	url, err := ReplaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/capacityCommitments?capacityCommitmentId={{capacity_commitment_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new CapacityCommitment: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for CapacityCommitment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating CapacityCommitment: %s", err)
	}
	if err := d.Set("name", flattenBigqueryReservationCapacityCommitmentName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/capacityCommitments/{{capacity_commitment_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating CapacityCommitment %q: %#v", d.Id(), res)

	return resourceBigqueryReservationCapacityCommitmentRead(d, meta)
}

func resourceBigqueryReservationCapacityCommitmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/capacityCommitments/{{capacity_commitment_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for CapacityCommitment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("BigqueryReservationCapacityCommitment %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading CapacityCommitment: %s", err)
	}

	if err := d.Set("name", flattenBigqueryReservationCapacityCommitmentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading CapacityCommitment: %s", err)
	}
	if err := d.Set("slot_count", flattenBigqueryReservationCapacityCommitmentSlotCount(res["slotCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading CapacityCommitment: %s", err)
	}
	if err := d.Set("plan", flattenBigqueryReservationCapacityCommitmentPlan(res["plan"], d, config)); err != nil {
		return fmt.Errorf("Error reading CapacityCommitment: %s", err)
	}
	if err := d.Set("state", flattenBigqueryReservationCapacityCommitmentState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading CapacityCommitment: %s", err)
	}
	if err := d.Set("commitment_start_time", flattenBigqueryReservationCapacityCommitmentCommitmentStartTime(res["commitmentStartTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading CapacityCommitment: %s", err)
	}
	if err := d.Set("commitment_end_time", flattenBigqueryReservationCapacityCommitmentCommitmentEndTime(res["commitmentEndTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading CapacityCommitment: %s", err)
	}
	if err := d.Set("renewal_plan", flattenBigqueryReservationCapacityCommitmentRenewalPlan(res["renewalPlan"], d, config)); err != nil {
		return fmt.Errorf("Error reading CapacityCommitment: %s", err)
	}
	if err := d.Set("edition", flattenBigqueryReservationCapacityCommitmentEdition(res["edition"], d, config)); err != nil {
		return fmt.Errorf("Error reading CapacityCommitment: %s", err)
	}

	return nil
}

func resourceBigqueryReservationCapacityCommitmentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for CapacityCommitment: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	planProp, err := expandBigqueryReservationCapacityCommitmentPlan(d.Get("plan"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("plan"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, planProp)) {
		obj["plan"] = planProp
	}
	renewalPlanProp, err := expandBigqueryReservationCapacityCommitmentRenewalPlan(d.Get("renewal_plan"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("renewal_plan"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, renewalPlanProp)) {
		obj["renewalPlan"] = renewalPlanProp
	}

	url, err := ReplaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/capacityCommitments/{{capacity_commitment_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating CapacityCommitment %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("plan") {
		updateMask = append(updateMask, "plan")
	}

	if d.HasChange("renewal_plan") {
		updateMask = append(updateMask, "renewalPlan")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating CapacityCommitment %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating CapacityCommitment %q: %#v", d.Id(), res)
	}

	return resourceBigqueryReservationCapacityCommitmentRead(d, meta)
}

func resourceBigqueryReservationCapacityCommitmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for CapacityCommitment: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/capacityCommitments/{{capacity_commitment_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting CapacityCommitment %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "CapacityCommitment")
	}

	log.Printf("[DEBUG] Finished deleting CapacityCommitment %q: %#v", d.Id(), res)
	return nil
}

func resourceBigqueryReservationCapacityCommitmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/capacityCommitments/(?P<capacity_commitment_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<capacity_commitment_id>[^/]+)",
		"(?P<location>[^/]+)/(?P<capacity_commitment_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/capacityCommitments/{{capacity_commitment_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBigqueryReservationCapacityCommitmentName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryReservationCapacityCommitmentSlotCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenBigqueryReservationCapacityCommitmentPlan(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryReservationCapacityCommitmentState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryReservationCapacityCommitmentCommitmentStartTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryReservationCapacityCommitmentCommitmentEndTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryReservationCapacityCommitmentRenewalPlan(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryReservationCapacityCommitmentEdition(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandBigqueryReservationCapacityCommitmentSlotCount(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationCapacityCommitmentPlan(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationCapacityCommitmentRenewalPlan(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationCapacityCommitmentEdition(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
