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

func resourceVertexAiEndpointTrafficSplit() *schema.Resource {
	return &schema.Resource{
		Create: resourceVertexAiEndpointTrafficSplitCreate,
		Read:   resourceVertexAiEndpointTrafficSplitRead,
		Update: resourceVertexAiEndpointTrafficSplitUpdate,
		Delete: resourceVertexAiEndpointTrafficSplitDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVertexAiEndpointTrafficSplitImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "The endpoint for the resource",
			},

			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The location for the resource",
			},

			"traffic_split": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "A map from a DeployedModel's ID to the percentage of this Endpoint's traffic that should be forwarded to that DeployedModel. If a DeployedModel's ID is not listed in this map, then it receives no traffic. The traffic percentage values must add up to 100, or map must be empty if the Endpoint is to not accept any traffic at a moment.",
				Elem:        VertexAiEndpointTrafficSplitTrafficSplitSchema(),
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "The project for the resource",
			},

			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Used to perform consistent read-modify-write updates. If not set, a blind \"overwrite\" update happens.",
			},
		},
	}
}

func VertexAiEndpointTrafficSplitTrafficSplitSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deployed_model_id": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "A deployed model's id.",
			},

			"traffic_percentage": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The percentage of this Endpoint's traffic that should be forwarded to the DeployedModel.",
			},
		},
	}
}

func resourceVertexAiEndpointTrafficSplitCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vertexai.EndpointTrafficSplit{
		Endpoint:     dcl.String(d.Get("endpoint").(string)),
		Location:     dcl.String(d.Get("location").(string)),
		TrafficSplit: expandVertexAiEndpointTrafficSplitTrafficSplitArray(d.Get("traffic_split")),
		Project:      dcl.String(project),
	}

	id, err := obj.ID()
	if err != nil {
		return fmt.Errorf("error constructing id: %s", err)
	}
	d.SetId(id)

	directive := UpdateDirective
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
	res, err := client.ApplyEndpointTrafficSplit(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating EndpointTrafficSplit: %s", err)
	}

	log.Printf("[DEBUG] Finished creating EndpointTrafficSplit %q: %#v", d.Id(), res)

	return resourceVertexAiEndpointTrafficSplitRead(d, meta)
}

func resourceVertexAiEndpointTrafficSplitRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vertexai.EndpointTrafficSplit{
		Endpoint:     dcl.String(d.Get("endpoint").(string)),
		Location:     dcl.String(d.Get("location").(string)),
		TrafficSplit: expandVertexAiEndpointTrafficSplitTrafficSplitArray(d.Get("traffic_split")),
		Project:      dcl.String(project),
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
	res, err := client.GetEndpointTrafficSplit(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("VertexAiEndpointTrafficSplit %q", d.Id())
		return handleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("endpoint", res.Endpoint); err != nil {
		return fmt.Errorf("error setting endpoint in state: %s", err)
	}
	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("traffic_split", flattenVertexAiEndpointTrafficSplitTrafficSplitArray(res.TrafficSplit)); err != nil {
		return fmt.Errorf("error setting traffic_split in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("etag", res.Etag); err != nil {
		return fmt.Errorf("error setting etag in state: %s", err)
	}

	return nil
}
func resourceVertexAiEndpointTrafficSplitUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vertexai.EndpointTrafficSplit{
		Endpoint:     dcl.String(d.Get("endpoint").(string)),
		Location:     dcl.String(d.Get("location").(string)),
		TrafficSplit: expandVertexAiEndpointTrafficSplitTrafficSplitArray(d.Get("traffic_split")),
		Project:      dcl.String(project),
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
	res, err := client.ApplyEndpointTrafficSplit(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error updating EndpointTrafficSplit: %s", err)
	}

	log.Printf("[DEBUG] Finished creating EndpointTrafficSplit %q: %#v", d.Id(), res)

	return resourceVertexAiEndpointTrafficSplitRead(d, meta)
}

func resourceVertexAiEndpointTrafficSplitDelete(d *schema.ResourceData, meta interface{}) error {

	log.Printf("[DEBUG] Finished deleting EndpointTrafficSplit %q", d.Id())
	return nil
}

func resourceVertexAiEndpointTrafficSplitImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/endpoints/(?P<endpoint>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<endpoint>[^/]+)",
		"(?P<location>[^/]+)/(?P<endpoint>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/endpoints/{{endpoint}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandVertexAiEndpointTrafficSplitTrafficSplitArray(o interface{}) []vertexai.EndpointTrafficSplitTrafficSplit {
	if o == nil {
		return make([]vertexai.EndpointTrafficSplitTrafficSplit, 0)
	}

	objs := o.([]interface{})
	if len(objs) == 0 || objs[0] == nil {
		return make([]vertexai.EndpointTrafficSplitTrafficSplit, 0)
	}

	items := make([]vertexai.EndpointTrafficSplitTrafficSplit, 0, len(objs))
	for _, item := range objs {
		i := expandVertexAiEndpointTrafficSplitTrafficSplit(item)
		items = append(items, *i)
	}

	return items
}

func expandVertexAiEndpointTrafficSplitTrafficSplit(o interface{}) *vertexai.EndpointTrafficSplitTrafficSplit {
	if o == nil {
		return vertexai.EmptyEndpointTrafficSplitTrafficSplit
	}

	obj := o.(map[string]interface{})
	return &vertexai.EndpointTrafficSplitTrafficSplit{
		DeployedModelId:   dcl.String(obj["deployed_model_id"].(string)),
		TrafficPercentage: dcl.Int64(int64(obj["traffic_percentage"].(int))),
	}
}

func flattenVertexAiEndpointTrafficSplitTrafficSplitArray(objs []vertexai.EndpointTrafficSplitTrafficSplit) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenVertexAiEndpointTrafficSplitTrafficSplit(&item)
		items = append(items, i)
	}

	return items
}

func flattenVertexAiEndpointTrafficSplitTrafficSplit(obj *vertexai.EndpointTrafficSplitTrafficSplit) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"deployed_model_id":  obj.DeployedModelId,
		"traffic_percentage": obj.TrafficPercentage,
	}

	return transformed

}
