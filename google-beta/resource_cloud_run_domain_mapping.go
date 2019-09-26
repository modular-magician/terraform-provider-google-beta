// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceCloudRunDomainMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudRunDomainMappingCreate,
		Read:   resourceCloudRunDomainMappingRead,
		Update: resourceCloudRunDomainMappingUpdate,
		Delete: resourceCloudRunDomainMappingDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudRunDomainMappingImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metadata": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
						},
						"annotations": {
							Type:     schema.TypeMap,
							Computed: true,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"labels": {
							Type:     schema.TypeMap,
							Computed: true,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"generation": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"resource_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"self_link": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uid": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"spec": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"route_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"certificate_mode": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"NONE", "AUTOMATIC", ""}, false),
						},
						"force_override": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_records": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rrdata": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"A", "AAAA", "CNAME", ""}, false),
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"conditions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"reason": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"mapped_route_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"observed_generation": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCloudRunDomainMappingCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	specProp, err := expandCloudRunDomainMappingSpec(d.Get("spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("spec"); !isEmptyValue(reflect.ValueOf(specProp)) && (ok || !reflect.DeepEqual(v, specProp)) {
		obj["spec"] = specProp
	}
	metadataProp, err := expandCloudRunDomainMappingMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(metadataProp)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}

	obj, err = resourceCloudRunDomainMappingEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{CloudRunBasePath}}projects/{{project}}/locations/{{location}}/domainmappings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new DomainMapping: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating DomainMapping: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating DomainMapping %q: %#v", d.Id(), res)

	return resourceCloudRunDomainMappingRead(d, meta)
}

func resourceCloudRunDomainMappingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{CloudRunBasePath}}projects/{{project}}/locations/{{location}}/domainmappings/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("CloudRunDomainMapping %q", d.Id()))
	}

	res, err = resourceCloudRunDomainMappingDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing CloudRunDomainMapping because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}

	if err := d.Set("status", flattenCloudRunDomainMappingStatus(res["status"], d)); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}
	if err := d.Set("spec", flattenCloudRunDomainMappingSpec(res["spec"], d)); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}
	if err := d.Set("metadata", flattenCloudRunDomainMappingMetadata(res["metadata"], d)); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}

	return nil
}

func resourceCloudRunDomainMappingUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	specProp, err := expandCloudRunDomainMappingSpec(d.Get("spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("spec"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, specProp)) {
		obj["spec"] = specProp
	}
	metadataProp, err := expandCloudRunDomainMappingMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}

	obj, err = resourceCloudRunDomainMappingEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{CloudRunBasePath}}projects/{{project}}/locations/{{location}}/domainmappings/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating DomainMapping %q: %#v", d.Id(), obj)
	_, err = sendRequestWithTimeout(config, "PUT", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating DomainMapping %q: %s", d.Id(), err)
	}

	return resourceCloudRunDomainMappingRead(d, meta)
}

func resourceCloudRunDomainMappingDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{CloudRunBasePath}}projects/{{project}}/locations/{{location}}/domainmappings/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting DomainMapping %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "DomainMapping")
	}

	log.Printf("[DEBUG] Finished deleting DomainMapping %q: %#v", d.Id(), res)
	return nil
}

func resourceCloudRunDomainMappingImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/domainmappings/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenCloudRunDomainMappingStatus(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["conditions"] =
		flattenCloudRunDomainMappingStatusConditions(original["conditions"], d)
	transformed["observed_generation"] =
		flattenCloudRunDomainMappingStatusObservedGeneration(original["observedGeneration"], d)
	transformed["resource_records"] =
		flattenCloudRunDomainMappingStatusResourceRecords(original["resourceRecords"], d)
	transformed["mapped_route_name"] =
		flattenCloudRunDomainMappingStatusMappedRouteName(original["mappedRouteName"], d)
	return []interface{}{transformed}
}
func flattenCloudRunDomainMappingStatusConditions(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"message": flattenCloudRunDomainMappingStatusConditionsMessage(original["message"], d),
			"status":  flattenCloudRunDomainMappingStatusConditionsStatus(original["status"], d),
			"reason":  flattenCloudRunDomainMappingStatusConditionsReason(original["reason"], d),
			"type":    flattenCloudRunDomainMappingStatusConditionsType(original["type"], d),
		})
	}
	return transformed
}
func flattenCloudRunDomainMappingStatusConditionsMessage(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudRunDomainMappingStatusConditionsStatus(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudRunDomainMappingStatusConditionsReason(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudRunDomainMappingStatusConditionsType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudRunDomainMappingStatusObservedGeneration(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenCloudRunDomainMappingStatusResourceRecords(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"type":   flattenCloudRunDomainMappingStatusResourceRecordsType(original["type"], d),
			"rrdata": flattenCloudRunDomainMappingStatusResourceRecordsRrdata(original["rrdata"], d),
			"name":   flattenCloudRunDomainMappingStatusResourceRecordsName(original["name"], d),
		})
	}
	return transformed
}
func flattenCloudRunDomainMappingStatusResourceRecordsType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudRunDomainMappingStatusResourceRecordsRrdata(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudRunDomainMappingStatusResourceRecordsName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudRunDomainMappingStatusMappedRouteName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudRunDomainMappingSpec(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["force_override"] =
		flattenCloudRunDomainMappingSpecForceOverride(original["forceOverride"], d)
	transformed["route_name"] =
		flattenCloudRunDomainMappingSpecRouteName(original["routeName"], d)
	transformed["certificate_mode"] =
		flattenCloudRunDomainMappingSpecCertificateMode(original["certificateMode"], d)
	return []interface{}{transformed}
}
func flattenCloudRunDomainMappingSpecForceOverride(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudRunDomainMappingSpecRouteName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudRunDomainMappingSpecCertificateMode(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudRunDomainMappingMetadata(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["labels"] =
		flattenCloudRunDomainMappingMetadataLabels(original["labels"], d)
	transformed["generation"] =
		flattenCloudRunDomainMappingMetadataGeneration(original["generation"], d)
	transformed["resource_version"] =
		flattenCloudRunDomainMappingMetadataResourceVersion(original["resourceVersion"], d)
	transformed["self_link"] =
		flattenCloudRunDomainMappingMetadataSelfLink(original["selfLink"], d)
	transformed["uid"] =
		flattenCloudRunDomainMappingMetadataUid(original["uid"], d)
	transformed["namespace"] =
		flattenCloudRunDomainMappingMetadataNamespace(original["namespace"], d)
	transformed["annotations"] =
		flattenCloudRunDomainMappingMetadataAnnotations(original["annotations"], d)
	return []interface{}{transformed}
}
func flattenCloudRunDomainMappingMetadataLabels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudRunDomainMappingMetadataGeneration(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenCloudRunDomainMappingMetadataResourceVersion(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudRunDomainMappingMetadataSelfLink(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudRunDomainMappingMetadataUid(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudRunDomainMappingMetadataNamespace(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudRunDomainMappingMetadataAnnotations(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandCloudRunDomainMappingSpec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedForceOverride, err := expandCloudRunDomainMappingSpecForceOverride(original["force_override"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedForceOverride); val.IsValid() && !isEmptyValue(val) {
		transformed["forceOverride"] = transformedForceOverride
	}

	transformedRouteName, err := expandCloudRunDomainMappingSpecRouteName(original["route_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRouteName); val.IsValid() && !isEmptyValue(val) {
		transformed["routeName"] = transformedRouteName
	}

	transformedCertificateMode, err := expandCloudRunDomainMappingSpecCertificateMode(original["certificate_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertificateMode); val.IsValid() && !isEmptyValue(val) {
		transformed["certificateMode"] = transformedCertificateMode
	}

	return transformed, nil
}

func expandCloudRunDomainMappingSpecForceOverride(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunDomainMappingSpecRouteName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunDomainMappingSpecCertificateMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunDomainMappingMetadata(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLabels, err := expandCloudRunDomainMappingMetadataLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !isEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	transformedGeneration, err := expandCloudRunDomainMappingMetadataGeneration(original["generation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGeneration); val.IsValid() && !isEmptyValue(val) {
		transformed["generation"] = transformedGeneration
	}

	transformedResourceVersion, err := expandCloudRunDomainMappingMetadataResourceVersion(original["resource_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceVersion); val.IsValid() && !isEmptyValue(val) {
		transformed["resourceVersion"] = transformedResourceVersion
	}

	transformedSelfLink, err := expandCloudRunDomainMappingMetadataSelfLink(original["self_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSelfLink); val.IsValid() && !isEmptyValue(val) {
		transformed["selfLink"] = transformedSelfLink
	}

	transformedUid, err := expandCloudRunDomainMappingMetadataUid(original["uid"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUid); val.IsValid() && !isEmptyValue(val) {
		transformed["uid"] = transformedUid
	}

	transformedNamespace, err := expandCloudRunDomainMappingMetadataNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !isEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	transformedAnnotations, err := expandCloudRunDomainMappingMetadataAnnotations(original["annotations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAnnotations); val.IsValid() && !isEmptyValue(val) {
		transformed["annotations"] = transformedAnnotations
	}

	return transformed, nil
}

func expandCloudRunDomainMappingMetadataLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudRunDomainMappingMetadataGeneration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunDomainMappingMetadataResourceVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunDomainMappingMetadataSelfLink(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunDomainMappingMetadataUid(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunDomainMappingMetadataNamespace(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunDomainMappingMetadataAnnotations(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func resourceCloudRunDomainMappingEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	name := d.Get("name").(string)
	metadata := obj["metadata"].(map[string]interface{})
	metadata["name"] = name

	// The only acceptable version/kind right now
	obj["apiVersion"] = "serving.knative.dev/v1alpha1"
	obj["kind"] = "DomainMapping"
	return obj, nil
}

func resourceCloudRunDomainMappingDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// metadata is not present if the API returns an error
	if obj, ok := res["metadata"]; ok {
		if meta, ok := obj.(map[string]interface{}); ok {
			res["name"] = meta["name"]
		} else {
			return nil, fmt.Errorf("Unable to decode 'metadata' block from API response.")
		}
	}
	return res, nil
}
