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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceComputeNodeTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeNodeTemplateCreate,
		Read:   resourceComputeNodeTemplateRead,
		Delete: resourceComputeNodeTemplateDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeNodeTemplateImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cpu_overcommit_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ENABLED", "NONE", ""}),
				Description:  `CPU overcommit. Default value: "NONE" Possible values: ["ENABLED", "NONE"]`,
				Default:      "NONE",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional textual description of the resource.`,
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Name of the resource.`,
			},
			"node_affinity_labels": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Description: `Labels to use for node affinity, which will be used in
instance scheduling.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"node_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Node type to use for nodes group that are created from this template.
Only one of nodeTypeFlexibility and nodeType can be specified.`,
				ConflictsWith: []string{"node_type_flexibility"},
			},
			"node_type_flexibility": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Flexible properties for the desired node type. Node groups that
use this node template will create nodes of a type that matches
these properties. Only one of nodeTypeFlexibility and nodeType can
be specified.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpus": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							Description:  `Number of virtual CPUs to use.`,
							AtLeastOneOf: []string{"node_type_flexibility.0.cpus", "node_type_flexibility.0.memory"},
						},
						"memory": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							Description:  `Physical memory available to the node, defined in MB.`,
							AtLeastOneOf: []string{"node_type_flexibility.0.cpus", "node_type_flexibility.0.memory"},
						},
						"local_ssd": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Use local SSD`,
						},
					},
				},
				ConflictsWith: []string{"node_type"},
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `Region where nodes using the node template will be created.
If it is not provided, the provider region is used.`,
			},
			"server_binding": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `The server binding policy for nodes using this template. Determines
where the nodes should restart following a maintenance event.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: verify.ValidateEnum([]string{"RESTART_NODE_ON_ANY_SERVER", "RESTART_NODE_ON_MINIMAL_SERVERS"}),
							Description: `Type of server binding policy. If 'RESTART_NODE_ON_ANY_SERVER',
nodes using this template will restart on any physical server
following a maintenance event.

If 'RESTART_NODE_ON_MINIMAL_SERVER', nodes using this template
will restart on the same physical server following a maintenance
event, instead of being live migrated to or restarted on a new
physical server. This option may be useful if you are using
software licenses tied to the underlying server characteristics
such as physical sockets or cores, to avoid the need for
additional licenses when maintenance occurs. However, VMs on such
nodes will experience outages while maintenance is applied. Possible values: ["RESTART_NODE_ON_ANY_SERVER", "RESTART_NODE_ON_MINIMAL_SERVERS"]`,
						},
					},
				},
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeNodeTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeNodeTemplateDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeNodeTemplateName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	nodeAffinityLabelsProp, err := expandComputeNodeTemplateNodeAffinityLabels(d.Get("node_affinity_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_affinity_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeAffinityLabelsProp)) && (ok || !reflect.DeepEqual(v, nodeAffinityLabelsProp)) {
		obj["nodeAffinityLabels"] = nodeAffinityLabelsProp
	}
	nodeTypeProp, err := expandComputeNodeTemplateNodeType(d.Get("node_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeTypeProp)) && (ok || !reflect.DeepEqual(v, nodeTypeProp)) {
		obj["nodeType"] = nodeTypeProp
	}
	nodeTypeFlexibilityProp, err := expandComputeNodeTemplateNodeTypeFlexibility(d.Get("node_type_flexibility"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_type_flexibility"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeTypeFlexibilityProp)) && (ok || !reflect.DeepEqual(v, nodeTypeFlexibilityProp)) {
		obj["nodeTypeFlexibility"] = nodeTypeFlexibilityProp
	}
	serverBindingProp, err := expandComputeNodeTemplateServerBinding(d.Get("server_binding"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("server_binding"); !tpgresource.IsEmptyValue(reflect.ValueOf(serverBindingProp)) && (ok || !reflect.DeepEqual(v, serverBindingProp)) {
		obj["serverBinding"] = serverBindingProp
	}
	cpuOvercommitTypeProp, err := expandComputeNodeTemplateCpuOvercommitType(d.Get("cpu_overcommit_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cpu_overcommit_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(cpuOvercommitTypeProp)) && (ok || !reflect.DeepEqual(v, cpuOvercommitTypeProp)) {
		obj["cpuOvercommitType"] = cpuOvercommitTypeProp
	}
	regionProp, err := expandComputeNodeTemplateRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/nodeTemplates")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new NodeTemplate: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NodeTemplate: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating NodeTemplate: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating NodeTemplate", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create NodeTemplate: %s", err)
	}

	log.Printf("[DEBUG] Finished creating NodeTemplate %q: %#v", d.Id(), res)

	return resourceComputeNodeTemplateRead(d, meta)
}

func resourceComputeNodeTemplateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NodeTemplate: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeNodeTemplate %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading NodeTemplate: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeNodeTemplateCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodeTemplate: %s", err)
	}
	if err := d.Set("description", flattenComputeNodeTemplateDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodeTemplate: %s", err)
	}
	if err := d.Set("name", flattenComputeNodeTemplateName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodeTemplate: %s", err)
	}
	if err := d.Set("node_affinity_labels", flattenComputeNodeTemplateNodeAffinityLabels(res["nodeAffinityLabels"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodeTemplate: %s", err)
	}
	if err := d.Set("node_type", flattenComputeNodeTemplateNodeType(res["nodeType"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodeTemplate: %s", err)
	}
	if err := d.Set("node_type_flexibility", flattenComputeNodeTemplateNodeTypeFlexibility(res["nodeTypeFlexibility"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodeTemplate: %s", err)
	}
	if err := d.Set("server_binding", flattenComputeNodeTemplateServerBinding(res["serverBinding"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodeTemplate: %s", err)
	}
	if err := d.Set("cpu_overcommit_type", flattenComputeNodeTemplateCpuOvercommitType(res["cpuOvercommitType"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodeTemplate: %s", err)
	}
	if err := d.Set("region", flattenComputeNodeTemplateRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading NodeTemplate: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading NodeTemplate: %s", err)
	}

	return nil
}

func resourceComputeNodeTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NodeTemplate: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting NodeTemplate %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "NodeTemplate")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting NodeTemplate", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting NodeTemplate %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeNodeTemplateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/nodeTemplates/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeNodeTemplateCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNodeTemplateDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNodeTemplateName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNodeTemplateNodeAffinityLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNodeTemplateNodeType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNodeTemplateNodeTypeFlexibility(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["cpus"] =
		flattenComputeNodeTemplateNodeTypeFlexibilityCpus(original["cpus"], d, config)
	transformed["memory"] =
		flattenComputeNodeTemplateNodeTypeFlexibilityMemory(original["memory"], d, config)
	transformed["local_ssd"] =
		flattenComputeNodeTemplateNodeTypeFlexibilityLocalSsd(original["localSsd"], d, config)
	return []interface{}{transformed}
}
func flattenComputeNodeTemplateNodeTypeFlexibilityCpus(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNodeTemplateNodeTypeFlexibilityMemory(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNodeTemplateNodeTypeFlexibilityLocalSsd(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNodeTemplateServerBinding(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["type"] =
		flattenComputeNodeTemplateServerBindingType(original["type"], d, config)
	return []interface{}{transformed}
}
func flattenComputeNodeTemplateServerBindingType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNodeTemplateCpuOvercommitType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeNodeTemplateRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func expandComputeNodeTemplateDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeTemplateName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeTemplateNodeAffinityLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeNodeTemplateNodeType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeTemplateNodeTypeFlexibility(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCpus, err := expandComputeNodeTemplateNodeTypeFlexibilityCpus(original["cpus"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCpus); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cpus"] = transformedCpus
	}

	transformedMemory, err := expandComputeNodeTemplateNodeTypeFlexibilityMemory(original["memory"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMemory); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["memory"] = transformedMemory
	}

	transformedLocalSsd, err := expandComputeNodeTemplateNodeTypeFlexibilityLocalSsd(original["local_ssd"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocalSsd); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["localSsd"] = transformedLocalSsd
	}

	return transformed, nil
}

func expandComputeNodeTemplateNodeTypeFlexibilityCpus(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeTemplateNodeTypeFlexibilityMemory(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeTemplateNodeTypeFlexibilityLocalSsd(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeTemplateServerBinding(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedType, err := expandComputeNodeTemplateServerBindingType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["type"] = transformedType
	}

	return transformed, nil
}

func expandComputeNodeTemplateServerBindingType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeTemplateCpuOvercommitType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeTemplateRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
