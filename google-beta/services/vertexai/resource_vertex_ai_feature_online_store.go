// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

package vertexai

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceVertexAIFeatureOnlineStore() *schema.Resource {
	return &schema.Resource{
		Create: resourceVertexAIFeatureOnlineStoreCreate,
		Read:   resourceVertexAIFeatureOnlineStoreRead,
		Update: resourceVertexAIFeatureOnlineStoreUpdate,
		Delete: resourceVertexAIFeatureOnlineStoreDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVertexAIFeatureOnlineStoreImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource name of the Feature Online Store. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.`,
			},
			"bigtable": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Settings for Cloud Bigtable instance that will be created to serve featureValues for all FeatureViews under this FeatureOnlineStore.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_scaling": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `Autoscaling config applied to Bigtable Instance.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_node_count": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: `The maximum number of nodes to scale up to. Must be greater than or equal to minNodeCount, and less than or equal to 10 times of 'minNodeCount'.`,
									},
									"min_node_count": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: `The minimum number of nodes to scale down to. Must be greater than or equal to 1.`,
									},
									"cpu_utilization_target": {
										Type:        schema.TypeInt,
										Computed:    true,
										Optional:    true,
										Description: `A percentage of the cluster's CPU capacity. Can be from 10% to 80%. When a cluster's CPU utilization exceeds the target that you have set, Bigtable immediately adds nodes to the cluster. When CPU utilization is substantially lower than the target, Bigtable removes nodes. If not set will default to 50%.`,
									},
								},
							},
						},
					},
				},
				ExactlyOneOf: []string{"bigtable", "optimized"},
			},
			"dedicated_serving_endpoint": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: `The dedicated serving endpoint for this FeatureOnlineStore, which is different from common vertex service endpoint. Only need to set when you choose Optimized storage type or enable EmbeddingManagement. Will use public endpoint by default.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"private_service_connect_config": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Private service connect config.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_private_service_connect": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: `If set to true, customers will use private service connection to send request. Otherwise, the connection will set to public endpoint.`,
									},
									"project_allowlist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `A list of Projects from which the forwarding rule will target the service attachment.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"public_endpoint_domain_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Domain name to use for this FeatureOnlineStore`,
						},
						"service_attachment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Name of the service attachment resource. Applicable only if private service connect is enabled and after FeatureViewSync is created.`,
						},
					},
				},
			},
			"embedding_management": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The settings for embedding management in FeatureOnlineStore. Embedding management can only be used with BigTable.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Description: `Enable embedding management.`,
						},
					},
				},
				ConflictsWith: []string{"optimized"},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `The labels with user-defined metadata to organize your feature online stores.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"optimized": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Settings for the Optimized store that will be created to serve featureValues for all FeatureViews under this FeatureOnlineStore`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
				ConflictsWith: []string{"embedding_management"},
				ExactlyOneOf:  []string{"bigtable", "optimized"},
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The region of feature online store. eg us-central1`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the feature online store was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Used to perform consistent read-modify-write updates.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The state of the Feature Online Store. See the possible states in [this link](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featureOnlineStores#state).`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the feature online store was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"force_destroy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: `If set to true, any FeatureViews and Features for this FeatureOnlineStore will also be deleted.`,
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

func resourceVertexAIFeatureOnlineStoreCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	bigtableProp, err := expandVertexAIFeatureOnlineStoreBigtable(d.Get("bigtable"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bigtable"); !tpgresource.IsEmptyValue(reflect.ValueOf(bigtableProp)) && (ok || !reflect.DeepEqual(v, bigtableProp)) {
		obj["bigtable"] = bigtableProp
	}
	optimizedProp, err := expandVertexAIFeatureOnlineStoreOptimized(d.Get("optimized"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("optimized"); ok || !reflect.DeepEqual(v, optimizedProp) {
		obj["optimized"] = optimizedProp
	}
	dedicatedServingEndpointProp, err := expandVertexAIFeatureOnlineStoreDedicatedServingEndpoint(d.Get("dedicated_serving_endpoint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dedicated_serving_endpoint"); !tpgresource.IsEmptyValue(reflect.ValueOf(dedicatedServingEndpointProp)) && (ok || !reflect.DeepEqual(v, dedicatedServingEndpointProp)) {
		obj["dedicatedServingEndpoint"] = dedicatedServingEndpointProp
	}
	embeddingManagementProp, err := expandVertexAIFeatureOnlineStoreEmbeddingManagement(d.Get("embedding_management"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("embedding_management"); !tpgresource.IsEmptyValue(reflect.ValueOf(embeddingManagementProp)) && (ok || !reflect.DeepEqual(v, embeddingManagementProp)) {
		obj["embeddingManagement"] = embeddingManagementProp
	}
	labelsProp, err := expandVertexAIFeatureOnlineStoreEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featureOnlineStores?featureOnlineStoreId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FeatureOnlineStore: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for FeatureOnlineStore: %s", err)
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
		return fmt.Errorf("Error creating FeatureOnlineStore: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = VertexAIOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating FeatureOnlineStore", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create FeatureOnlineStore: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating FeatureOnlineStore %q: %#v", d.Id(), res)

	return resourceVertexAIFeatureOnlineStoreRead(d, meta)
}

func resourceVertexAIFeatureOnlineStoreRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for FeatureOnlineStore: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("VertexAIFeatureOnlineStore %q", d.Id()))
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("force_destroy"); !ok {
		if err := d.Set("force_destroy", false); err != nil {
			return fmt.Errorf("Error setting force_destroy: %s", err)
		}
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}

	if err := d.Set("create_time", flattenVertexAIFeatureOnlineStoreCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}
	if err := d.Set("update_time", flattenVertexAIFeatureOnlineStoreUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}
	if err := d.Set("labels", flattenVertexAIFeatureOnlineStoreLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}
	if err := d.Set("state", flattenVertexAIFeatureOnlineStoreState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}
	if err := d.Set("bigtable", flattenVertexAIFeatureOnlineStoreBigtable(res["bigtable"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}
	if err := d.Set("optimized", flattenVertexAIFeatureOnlineStoreOptimized(res["optimized"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}
	if err := d.Set("dedicated_serving_endpoint", flattenVertexAIFeatureOnlineStoreDedicatedServingEndpoint(res["dedicatedServingEndpoint"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}
	if err := d.Set("embedding_management", flattenVertexAIFeatureOnlineStoreEmbeddingManagement(res["embeddingManagement"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}
	if err := d.Set("terraform_labels", flattenVertexAIFeatureOnlineStoreTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}
	if err := d.Set("effective_labels", flattenVertexAIFeatureOnlineStoreEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}

	return nil
}

func resourceVertexAIFeatureOnlineStoreUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for FeatureOnlineStore: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	bigtableProp, err := expandVertexAIFeatureOnlineStoreBigtable(d.Get("bigtable"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bigtable"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, bigtableProp)) {
		obj["bigtable"] = bigtableProp
	}
	optimizedProp, err := expandVertexAIFeatureOnlineStoreOptimized(d.Get("optimized"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("optimized"); ok || !reflect.DeepEqual(v, optimizedProp) {
		obj["optimized"] = optimizedProp
	}
	dedicatedServingEndpointProp, err := expandVertexAIFeatureOnlineStoreDedicatedServingEndpoint(d.Get("dedicated_serving_endpoint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dedicated_serving_endpoint"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, dedicatedServingEndpointProp)) {
		obj["dedicatedServingEndpoint"] = dedicatedServingEndpointProp
	}
	embeddingManagementProp, err := expandVertexAIFeatureOnlineStoreEmbeddingManagement(d.Get("embedding_management"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("embedding_management"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, embeddingManagementProp)) {
		obj["embeddingManagement"] = embeddingManagementProp
	}
	labelsProp, err := expandVertexAIFeatureOnlineStoreEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating FeatureOnlineStore %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("bigtable") {
		updateMask = append(updateMask, "bigtable")
	}

	if d.HasChange("optimized") {
		updateMask = append(updateMask, "optimized")
	}

	if d.HasChange("dedicated_serving_endpoint") {
		updateMask = append(updateMask, "dedicatedServingEndpoint")
	}

	if d.HasChange("embedding_management") {
		updateMask = append(updateMask, "embeddingManagement")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})

		if err != nil {
			return fmt.Errorf("Error updating FeatureOnlineStore %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating FeatureOnlineStore %q: %#v", d.Id(), res)
		}

		err = VertexAIOperationWaitTime(
			config, res, project, "Updating FeatureOnlineStore", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceVertexAIFeatureOnlineStoreRead(d, meta)
}

func resourceVertexAIFeatureOnlineStoreDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for FeatureOnlineStore: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	if v, ok := d.GetOk("force_destroy"); ok {
		url, err = transport_tpg.AddQueryParams(url, map[string]string{"force": fmt.Sprintf("%v", v)})
		if err != nil {
			return err
		}
	}

	log.Printf("[DEBUG] Deleting FeatureOnlineStore %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "FeatureOnlineStore")
	}

	err = VertexAIOperationWaitTime(
		config, res, project, "Deleting FeatureOnlineStore", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting FeatureOnlineStore %q: %#v", d.Id(), res)
	return nil
}

func resourceVertexAIFeatureOnlineStoreImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/featureOnlineStores/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Explicitly set virtual fields to default values on import
	if err := d.Set("force_destroy", false); err != nil {
		return nil, fmt.Errorf("Error setting force_destroy: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenVertexAIFeatureOnlineStoreCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeatureOnlineStoreUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeatureOnlineStoreLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenVertexAIFeatureOnlineStoreState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeatureOnlineStoreBigtable(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["auto_scaling"] =
		flattenVertexAIFeatureOnlineStoreBigtableAutoScaling(original["autoScaling"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeatureOnlineStoreBigtableAutoScaling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["min_node_count"] =
		flattenVertexAIFeatureOnlineStoreBigtableAutoScalingMinNodeCount(original["minNodeCount"], d, config)
	transformed["max_node_count"] =
		flattenVertexAIFeatureOnlineStoreBigtableAutoScalingMaxNodeCount(original["maxNodeCount"], d, config)
	transformed["cpu_utilization_target"] =
		flattenVertexAIFeatureOnlineStoreBigtableAutoScalingCpuUtilizationTarget(original["cpuUtilizationTarget"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeatureOnlineStoreBigtableAutoScalingMinNodeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
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

func flattenVertexAIFeatureOnlineStoreBigtableAutoScalingMaxNodeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
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

func flattenVertexAIFeatureOnlineStoreBigtableAutoScalingCpuUtilizationTarget(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
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

func flattenVertexAIFeatureOnlineStoreOptimized(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	transformed := make(map[string]interface{})
	return []interface{}{transformed}
}

func flattenVertexAIFeatureOnlineStoreDedicatedServingEndpoint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["public_endpoint_domain_name"] =
		flattenVertexAIFeatureOnlineStoreDedicatedServingEndpointPublicEndpointDomainName(original["publicEndpointDomainName"], d, config)
	transformed["service_attachment"] =
		flattenVertexAIFeatureOnlineStoreDedicatedServingEndpointServiceAttachment(original["serviceAttachment"], d, config)
	transformed["private_service_connect_config"] =
		flattenVertexAIFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfig(original["privateServiceConnectConfig"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeatureOnlineStoreDedicatedServingEndpointPublicEndpointDomainName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeatureOnlineStoreDedicatedServingEndpointServiceAttachment(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enable_private_service_connect"] =
		flattenVertexAIFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfigEnablePrivateServiceConnect(original["enablePrivateServiceConnect"], d, config)
	transformed["project_allowlist"] =
		flattenVertexAIFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfigProjectAllowlist(original["projectAllowlist"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfigEnablePrivateServiceConnect(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfigProjectAllowlist(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeatureOnlineStoreEmbeddingManagement(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enabled"] =
		flattenVertexAIFeatureOnlineStoreEmbeddingManagementEnabled(original["enabled"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeatureOnlineStoreEmbeddingManagementEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeatureOnlineStoreTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenVertexAIFeatureOnlineStoreEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandVertexAIFeatureOnlineStoreBigtable(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAutoScaling, err := expandVertexAIFeatureOnlineStoreBigtableAutoScaling(original["auto_scaling"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAutoScaling); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["autoScaling"] = transformedAutoScaling
	}

	return transformed, nil
}

func expandVertexAIFeatureOnlineStoreBigtableAutoScaling(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinNodeCount, err := expandVertexAIFeatureOnlineStoreBigtableAutoScalingMinNodeCount(original["min_node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minNodeCount"] = transformedMinNodeCount
	}

	transformedMaxNodeCount, err := expandVertexAIFeatureOnlineStoreBigtableAutoScalingMaxNodeCount(original["max_node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxNodeCount"] = transformedMaxNodeCount
	}

	transformedCpuUtilizationTarget, err := expandVertexAIFeatureOnlineStoreBigtableAutoScalingCpuUtilizationTarget(original["cpu_utilization_target"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCpuUtilizationTarget); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cpuUtilizationTarget"] = transformedCpuUtilizationTarget
	}

	return transformed, nil
}

func expandVertexAIFeatureOnlineStoreBigtableAutoScalingMinNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureOnlineStoreBigtableAutoScalingMaxNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureOnlineStoreBigtableAutoScalingCpuUtilizationTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureOnlineStoreOptimized(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	transformed := make(map[string]interface{})

	return transformed, nil
}

func expandVertexAIFeatureOnlineStoreDedicatedServingEndpoint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPublicEndpointDomainName, err := expandVertexAIFeatureOnlineStoreDedicatedServingEndpointPublicEndpointDomainName(original["public_endpoint_domain_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublicEndpointDomainName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["publicEndpointDomainName"] = transformedPublicEndpointDomainName
	}

	transformedServiceAttachment, err := expandVertexAIFeatureOnlineStoreDedicatedServingEndpointServiceAttachment(original["service_attachment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAttachment); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAttachment"] = transformedServiceAttachment
	}

	transformedPrivateServiceConnectConfig, err := expandVertexAIFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfig(original["private_service_connect_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrivateServiceConnectConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["privateServiceConnectConfig"] = transformedPrivateServiceConnectConfig
	}

	return transformed, nil
}

func expandVertexAIFeatureOnlineStoreDedicatedServingEndpointPublicEndpointDomainName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureOnlineStoreDedicatedServingEndpointServiceAttachment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnablePrivateServiceConnect, err := expandVertexAIFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfigEnablePrivateServiceConnect(original["enable_private_service_connect"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnablePrivateServiceConnect); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enablePrivateServiceConnect"] = transformedEnablePrivateServiceConnect
	}

	transformedProjectAllowlist, err := expandVertexAIFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfigProjectAllowlist(original["project_allowlist"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectAllowlist); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectAllowlist"] = transformedProjectAllowlist
	}

	return transformed, nil
}

func expandVertexAIFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfigEnablePrivateServiceConnect(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfigProjectAllowlist(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureOnlineStoreEmbeddingManagement(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandVertexAIFeatureOnlineStoreEmbeddingManagementEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	return transformed, nil
}

func expandVertexAIFeatureOnlineStoreEmbeddingManagementEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureOnlineStoreEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
