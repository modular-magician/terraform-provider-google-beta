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

func ResourceDataprocMetastoreFederation() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataprocMetastoreFederationCreate,
		Read:   resourceDataprocMetastoreFederationRead,
		Update: resourceDataprocMetastoreFederationUpdate,
		Delete: resourceDataprocMetastoreFederationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataprocMetastoreFederationImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"backend_metastores": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: `A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time. The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a BackendMetastore with a higher number.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rank": {
							Type:     schema.TypeString,
							Required: true,
						},
						"metastore_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateEnum([]string{"METASTORE_TYPE_UNSPECIFIED", "DATAPROC_METASTORE", "BIGQUERY"}),
							Description:  `The type of the backend metastore. Possible values: ["METASTORE_TYPE_UNSPECIFIED", "DATAPROC_METASTORE", "BIGQUERY"]`,
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The relative resource name of the metastore that is being federated. The formats of the relative resource names for the currently supported metastores are listed below: Dataplex: projects/{projectId}/locations/{location}/lakes/{lake_id} BigQuery: projects/{projectId} Dataproc Metastore: projects/{projectId}/locations/{location}/services/{serviceId}`,
						},
					},
				},
			},
			"federation_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The ID of the metastore federation. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
3 and 63 characters.`,
			},
			"version": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the federation version.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `User-defined labels for the metastore federation.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The location where the metastore federation should reside.`,
			},
			"endpoint_uri": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The URI of the endpoint used to access the metastore federation.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The relative resource name of the metastore federation.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The current state of the metastore federation.`,
			},
			"state_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Additional information about the current state of the metastore federation, if available.`,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The globally unique resource identifier of the metastore federation.`,
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

func resourceDataprocMetastoreFederationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandDataprocMetastoreFederationLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	versionProp, err := expandDataprocMetastoreFederationVersion(d.Get("version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("version"); !isEmptyValue(reflect.ValueOf(versionProp)) && (ok || !reflect.DeepEqual(v, versionProp)) {
		obj["version"] = versionProp
	}
	backendMetastoresProp, err := expandDataprocMetastoreFederationBackendMetastores(d.Get("backend_metastores"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backend_metastores"); !isEmptyValue(reflect.ValueOf(backendMetastoresProp)) && (ok || !reflect.DeepEqual(v, backendMetastoresProp)) {
		obj["backendMetastores"] = backendMetastoresProp
	}

	url, err := ReplaceVars(d, config, "{{DataprocMetastoreBasePath}}projects/{{project}}/locations/{{location}}/federations?federationId={{federation_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Federation: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Federation: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Federation: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/federations/{{federation_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = DataprocMetastoreOperationWaitTime(
		config, res, project, "Creating Federation", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Federation: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Federation %q: %#v", d.Id(), res)

	return resourceDataprocMetastoreFederationRead(d, meta)
}

func resourceDataprocMetastoreFederationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{DataprocMetastoreBasePath}}projects/{{project}}/locations/{{location}}/federations/{{federation_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Federation: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DataprocMetastoreFederation %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}

	if err := d.Set("name", flattenDataprocMetastoreFederationName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("labels", flattenDataprocMetastoreFederationLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("endpoint_uri", flattenDataprocMetastoreFederationEndpointUri(res["endpointUri"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("state", flattenDataprocMetastoreFederationState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("state_message", flattenDataprocMetastoreFederationStateMessage(res["stateMessage"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("uid", flattenDataprocMetastoreFederationUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("version", flattenDataprocMetastoreFederationVersion(res["version"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}
	if err := d.Set("backend_metastores", flattenDataprocMetastoreFederationBackendMetastores(res["backendMetastores"], d, config)); err != nil {
		return fmt.Errorf("Error reading Federation: %s", err)
	}

	return nil
}

func resourceDataprocMetastoreFederationUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Federation: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandDataprocMetastoreFederationLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	backendMetastoresProp, err := expandDataprocMetastoreFederationBackendMetastores(d.Get("backend_metastores"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backend_metastores"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, backendMetastoresProp)) {
		obj["backendMetastores"] = backendMetastoresProp
	}

	url, err := ReplaceVars(d, config, "{{DataprocMetastoreBasePath}}projects/{{project}}/locations/{{location}}/federations/{{federation_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Federation %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("backend_metastores") {
		updateMask = append(updateMask, "backendMetastores")
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
		return fmt.Errorf("Error updating Federation %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Federation %q: %#v", d.Id(), res)
	}

	err = DataprocMetastoreOperationWaitTime(
		config, res, project, "Updating Federation", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceDataprocMetastoreFederationRead(d, meta)
}

func resourceDataprocMetastoreFederationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Federation: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{DataprocMetastoreBasePath}}projects/{{project}}/locations/{{location}}/federations/{{federation_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Federation %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Federation")
	}

	err = DataprocMetastoreOperationWaitTime(
		config, res, project, "Deleting Federation", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Federation %q: %#v", d.Id(), res)
	return nil
}

func resourceDataprocMetastoreFederationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/federations/(?P<federation_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<federation_id>[^/]+)",
		"(?P<location>[^/]+)/(?P<federation_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/federations/{{federation_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDataprocMetastoreFederationName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationEndpointUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationStateMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationBackendMetastores(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.(map[string]interface{})
	transformed := make([]interface{}, 0, len(l))
	for k, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"rank":           k,
			"name":           flattenDataprocMetastoreFederationBackendMetastoresName(original["name"], d, config),
			"metastore_type": flattenDataprocMetastoreFederationBackendMetastoresMetastoreType(original["metastoreType"], d, config),
		})
	}
	return transformed
}
func flattenDataprocMetastoreFederationBackendMetastoresName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocMetastoreFederationBackendMetastoresMetastoreType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDataprocMetastoreFederationLabels(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataprocMetastoreFederationVersion(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreFederationBackendMetastores(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDataprocMetastoreFederationBackendMetastoresName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedMetastoreType, err := expandDataprocMetastoreFederationBackendMetastoresMetastoreType(original["metastore_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMetastoreType); val.IsValid() && !isEmptyValue(val) {
			transformed["metastoreType"] = transformedMetastoreType
		}

		transformedRank, err := expandString(original["rank"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedRank] = transformed
	}
	return m, nil
}

func expandDataprocMetastoreFederationBackendMetastoresName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocMetastoreFederationBackendMetastoresMetastoreType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
