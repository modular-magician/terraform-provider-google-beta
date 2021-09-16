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
)

func resourceDataLossPreventionStoredInfoType() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataLossPreventionStoredInfoTypeCreate,
		Read:   resourceDataLossPreventionStoredInfoTypeRead,
		Update: resourceDataLossPreventionStoredInfoTypeUpdate,
		Delete: resourceDataLossPreventionStoredInfoTypeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataLossPreventionStoredInfoTypeImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"parent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The parent of the info type in any of the following formats:

* 'projects/{{project}}'
* 'projects/{{project}}/locations/{{location}}'
* 'organizations/{{organization_id}}'
* 'organizations/{{organization_id}}/locations/{{location}}'`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A description of the info type.`,
			},
			"dictionary": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Dictionary which defines the rule.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloud_storage_path": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Newline-delimited file of words in Cloud Storage. Only a single file is accepted.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"path": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `A url representing a file or path (no wildcards) in Cloud Storage. Example: 'gs://[BUCKET_NAME]/dictionary.txt'`,
									},
								},
							},
							ExactlyOneOf: []string{"dictionary.0.word_list", "dictionary.0.cloud_storage_path"},
						},
						"word_list": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `List of words or phrases to search for.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"words": {
										Type:     schema.TypeList,
										Required: true,
										Description: `Words or phrases defining the dictionary. The dictionary must contain at least one
phrase and every phrase must contain at least 2 characters that are letters or digits.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
							ExactlyOneOf: []string{"dictionary.0.word_list", "dictionary.0.cloud_storage_path"},
						},
					},
				},
				ExactlyOneOf: []string{"dictionary", "regex", "large_custom_dictionary"},
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `User set display name of the info type.`,
			},
			"large_custom_dictionary": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Dictionary which defines the rule.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"output_path": {
							Type:     schema.TypeList,
							Required: true,
							Description: `Location to store dictionary artifacts in Google Cloud Storage. These files will only be accessible by project owners and the DLP API.
If any of these artifacts are modified, the dictionary is considered invalid and can no longer be used.`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"path": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `A url representing a file or path (no wildcards) in Cloud Storage. Example: 'gs://[BUCKET_NAME]/dictionary.txt'`,
									},
								},
							},
						},
						"big_query_field": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Field in a BigQuery table where each cell represents a dictionary phrase.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"field": {
										Type:        schema.TypeList,
										Required:    true,
										Description: `Designated field in the BigQuery table.`,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: `Name describing the field.`,
												},
											},
										},
									},
									"table": {
										Type:        schema.TypeList,
										Required:    true,
										Description: `Field in a BigQuery table where each cell represents a dictionary phrase.`,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dataset_id": {
													Type:        schema.TypeString,
													Required:    true,
													Description: `The dataset ID of the table.`,
												},
												"project_id": {
													Type:        schema.TypeString,
													Required:    true,
													Description: `The Google Cloud Platform project ID of the project containing the table.`,
												},
												"table_id": {
													Type:        schema.TypeString,
													Required:    true,
													Description: `The name of the table.`,
												},
											},
										},
									},
								},
							},
							ExactlyOneOf: []string{"large_custom_dictionary.0.cloud_storage_file_set", "large_custom_dictionary.0.big_query_field"},
						},
						"cloud_storage_file_set": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Set of files containing newline-delimited lists of dictionary phrases.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `The url, in the format 'gs://<bucket>/<path>'. Trailing wildcard in the path is allowed.`,
									},
								},
							},
							ExactlyOneOf: []string{"large_custom_dictionary.0.cloud_storage_file_set", "large_custom_dictionary.0.big_query_field"},
						},
					},
				},
				ExactlyOneOf: []string{"dictionary", "regex", "large_custom_dictionary"},
			},
			"regex": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Regular expression which defines the rule.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pattern": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Pattern defining the regular expression.
Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.`,
						},
						"group_indexes": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.`,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
					},
				},
				ExactlyOneOf: []string{"dictionary", "regex", "large_custom_dictionary"},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the info type. Set by the server.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceDataLossPreventionStoredInfoTypeCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandDataLossPreventionStoredInfoTypeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandDataLossPreventionStoredInfoTypeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	regexProp, err := expandDataLossPreventionStoredInfoTypeRegex(d.Get("regex"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("regex"); !isEmptyValue(reflect.ValueOf(regexProp)) && (ok || !reflect.DeepEqual(v, regexProp)) {
		obj["regex"] = regexProp
	}
	dictionaryProp, err := expandDataLossPreventionStoredInfoTypeDictionary(d.Get("dictionary"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dictionary"); !isEmptyValue(reflect.ValueOf(dictionaryProp)) && (ok || !reflect.DeepEqual(v, dictionaryProp)) {
		obj["dictionary"] = dictionaryProp
	}
	largeCustomDictionaryProp, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionary(d.Get("large_custom_dictionary"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("large_custom_dictionary"); !isEmptyValue(reflect.ValueOf(largeCustomDictionaryProp)) && (ok || !reflect.DeepEqual(v, largeCustomDictionaryProp)) {
		obj["largeCustomDictionary"] = largeCustomDictionaryProp
	}

	obj, err = resourceDataLossPreventionStoredInfoTypeEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{DataLossPreventionBasePath}}{{parent}}/storedInfoTypes")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new StoredInfoType: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating StoredInfoType: %s", err)
	}
	if err := d.Set("name", flattenDataLossPreventionStoredInfoTypeName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{parent}}/storedInfoTypes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = PollingWaitTime(resourceDataLossPreventionStoredInfoTypePollRead(d, meta), PollCheckForExistence, "Creating StoredInfoType", d.Timeout(schema.TimeoutCreate), 1)
	if err != nil {
		return fmt.Errorf("Error waiting to create StoredInfoType: %s", err)
	}

	log.Printf("[DEBUG] Finished creating StoredInfoType %q: %#v", d.Id(), res)

	return resourceDataLossPreventionStoredInfoTypeRead(d, meta)
}

func resourceDataLossPreventionStoredInfoTypePollRead(d *schema.ResourceData, meta interface{}) PollReadFunc {
	return func() (map[string]interface{}, error) {
		config := meta.(*Config)

		url, err := replaceVars(d, config, "{{DataLossPreventionBasePath}}{{parent}}/storedInfoTypes/{{name}}")
		if err != nil {
			return nil, err
		}

		billingProject := ""

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		userAgent, err := generateUserAgentString(d, config.userAgent)
		if err != nil {
			return nil, err
		}

		res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
		if err != nil {
			return res, err
		}
		res, err = resourceDataLossPreventionStoredInfoTypeDecoder(d, meta, res)
		if err != nil {
			return nil, err
		}
		if res == nil {
			// Decoded object not found, spoof a 404 error for poll
			return nil, &googleapi.Error{
				Code:    404,
				Message: "could not find object DataLossPreventionStoredInfoType",
			}
		}

		return res, nil
	}
}

func resourceDataLossPreventionStoredInfoTypeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{DataLossPreventionBasePath}}{{parent}}/storedInfoTypes/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("DataLossPreventionStoredInfoType %q", d.Id()))
	}

	res, err = resourceDataLossPreventionStoredInfoTypeDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing DataLossPreventionStoredInfoType because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("name", flattenDataLossPreventionStoredInfoTypeName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoredInfoType: %s", err)
	}
	if err := d.Set("description", flattenDataLossPreventionStoredInfoTypeDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoredInfoType: %s", err)
	}
	if err := d.Set("display_name", flattenDataLossPreventionStoredInfoTypeDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoredInfoType: %s", err)
	}
	if err := d.Set("regex", flattenDataLossPreventionStoredInfoTypeRegex(res["regex"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoredInfoType: %s", err)
	}
	if err := d.Set("dictionary", flattenDataLossPreventionStoredInfoTypeDictionary(res["dictionary"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoredInfoType: %s", err)
	}
	if err := d.Set("large_custom_dictionary", flattenDataLossPreventionStoredInfoTypeLargeCustomDictionary(res["largeCustomDictionary"], d, config)); err != nil {
		return fmt.Errorf("Error reading StoredInfoType: %s", err)
	}

	return nil
}

func resourceDataLossPreventionStoredInfoTypeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	descriptionProp, err := expandDataLossPreventionStoredInfoTypeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandDataLossPreventionStoredInfoTypeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	obj, err = resourceDataLossPreventionStoredInfoTypeEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{DataLossPreventionBasePath}}{{parent}}/storedInfoTypes/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating StoredInfoType %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating StoredInfoType %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating StoredInfoType %q: %#v", d.Id(), res)
	}

	return resourceDataLossPreventionStoredInfoTypeRead(d, meta)
}

func resourceDataLossPreventionStoredInfoTypeDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := replaceVars(d, config, "{{DataLossPreventionBasePath}}{{parent}}/storedInfoTypes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting StoredInfoType %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "StoredInfoType")
	}

	log.Printf("[DEBUG] Finished deleting StoredInfoType %q: %#v", d.Id(), res)
	return nil
}

func resourceDataLossPreventionStoredInfoTypeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// Custom import to handle parent possibilities
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}
	parts := strings.Split(d.Get("name").(string), "/")
	if len(parts) == 6 {
		if err := d.Set("name", parts[5]); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
	} else if len(parts) == 4 {
		if err := d.Set("name", parts[3]); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
	} else {
		return nil, fmt.Errorf("Unexpected import id: %s, expected form {{parent}}/storedInfoType/{{name}}", d.Get("name").(string))
	}
	// Remove "/storedInfoType/{{name}}" from the id
	parts = parts[:len(parts)-2]
	if err := d.Set("parent", strings.Join(parts, "/")); err != nil {
		return nil, fmt.Errorf("Error setting parent: %s", err)
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{parent}}/storedInfoTypes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDataLossPreventionStoredInfoTypeName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenDataLossPreventionStoredInfoTypeDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataLossPreventionStoredInfoTypeDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataLossPreventionStoredInfoTypeRegex(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["pattern"] =
		flattenDataLossPreventionStoredInfoTypeRegexPattern(original["pattern"], d, config)
	transformed["group_indexes"] =
		flattenDataLossPreventionStoredInfoTypeRegexGroupIndexes(original["groupIndexes"], d, config)
	return []interface{}{transformed}
}
func flattenDataLossPreventionStoredInfoTypeRegexPattern(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataLossPreventionStoredInfoTypeRegexGroupIndexes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataLossPreventionStoredInfoTypeDictionary(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["word_list"] =
		flattenDataLossPreventionStoredInfoTypeDictionaryWordList(original["wordList"], d, config)
	transformed["cloud_storage_path"] =
		flattenDataLossPreventionStoredInfoTypeDictionaryCloudStoragePath(original["cloudStoragePath"], d, config)
	return []interface{}{transformed}
}
func flattenDataLossPreventionStoredInfoTypeDictionaryWordList(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["words"] =
		flattenDataLossPreventionStoredInfoTypeDictionaryWordListWords(original["words"], d, config)
	return []interface{}{transformed}
}
func flattenDataLossPreventionStoredInfoTypeDictionaryWordListWords(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataLossPreventionStoredInfoTypeDictionaryCloudStoragePath(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["path"] =
		flattenDataLossPreventionStoredInfoTypeDictionaryCloudStoragePathPath(original["path"], d, config)
	return []interface{}{transformed}
}
func flattenDataLossPreventionStoredInfoTypeDictionaryCloudStoragePathPath(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataLossPreventionStoredInfoTypeLargeCustomDictionary(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["output_path"] =
		flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath(original["outputPath"], d, config)
	transformed["cloud_storage_file_set"] =
		flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(original["cloudStorageFileSet"], d, config)
	transformed["big_query_field"] =
		flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField(original["bigQueryField"], d, config)
	return []interface{}{transformed}
}
func flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["path"] =
		flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPathPath(original["path"], d, config)
	return []interface{}{transformed}
}
func flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPathPath(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["url"] =
		flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetUrl(original["url"], d, config)
	return []interface{}{transformed}
}
func flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetUrl(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["table"] =
		flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(original["table"], d, config)
	transformed["field"] =
		flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(original["field"], d, config)
	return []interface{}{transformed}
}
func flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["project_id"] =
		flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableProjectId(original["projectId"], d, config)
	transformed["dataset_id"] =
		flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableDatasetId(original["datasetId"], d, config)
	transformed["table_id"] =
		flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableTableId(original["tableId"], d, config)
	return []interface{}{transformed}
}
func flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableProjectId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableDatasetId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableTableId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["name"] =
		flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldName(original["name"], d, config)
	return []interface{}{transformed}
}
func flattenDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandDataLossPreventionStoredInfoTypeDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeRegex(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPattern, err := expandDataLossPreventionStoredInfoTypeRegexPattern(original["pattern"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPattern); val.IsValid() && !isEmptyValue(val) {
		transformed["pattern"] = transformedPattern
	}

	transformedGroupIndexes, err := expandDataLossPreventionStoredInfoTypeRegexGroupIndexes(original["group_indexes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGroupIndexes); val.IsValid() && !isEmptyValue(val) {
		transformed["groupIndexes"] = transformedGroupIndexes
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeRegexPattern(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeRegexGroupIndexes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeDictionary(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedWordList, err := expandDataLossPreventionStoredInfoTypeDictionaryWordList(original["word_list"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWordList); val.IsValid() && !isEmptyValue(val) {
		transformed["wordList"] = transformedWordList
	}

	transformedCloudStoragePath, err := expandDataLossPreventionStoredInfoTypeDictionaryCloudStoragePath(original["cloud_storage_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCloudStoragePath); val.IsValid() && !isEmptyValue(val) {
		transformed["cloudStoragePath"] = transformedCloudStoragePath
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeDictionaryWordList(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedWords, err := expandDataLossPreventionStoredInfoTypeDictionaryWordListWords(original["words"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWords); val.IsValid() && !isEmptyValue(val) {
		transformed["words"] = transformedWords
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeDictionaryWordListWords(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeDictionaryCloudStoragePath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPath, err := expandDataLossPreventionStoredInfoTypeDictionaryCloudStoragePathPath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeDictionaryCloudStoragePathPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionary(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedOutputPath, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath(original["output_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOutputPath); val.IsValid() && !isEmptyValue(val) {
		transformed["outputPath"] = transformedOutputPath
	}

	transformedCloudStorageFileSet, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(original["cloud_storage_file_set"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCloudStorageFileSet); val.IsValid() && !isEmptyValue(val) {
		transformed["cloudStorageFileSet"] = transformedCloudStorageFileSet
	}

	transformedBigQueryField, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField(original["big_query_field"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBigQueryField); val.IsValid() && !isEmptyValue(val) {
		transformed["bigQueryField"] = transformedBigQueryField
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPath, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPathPath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPathPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUrl, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetUrl(original["url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrl); val.IsValid() && !isEmptyValue(val) {
		transformed["url"] = transformedUrl
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTable, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(original["table"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTable); val.IsValid() && !isEmptyValue(val) {
		transformed["table"] = transformedTable
	}

	transformedField, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(original["field"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedField); val.IsValid() && !isEmptyValue(val) {
		transformed["field"] = transformedField
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProjectId, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !isEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedDatasetId, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableDatasetId(original["dataset_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !isEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	transformedTableId, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableTableId(original["table_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTableId); val.IsValid() && !isEmptyValue(val) {
		transformed["tableId"] = transformedTableId
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableProjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableDatasetId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableTableId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
		transformed["name"] = transformedName
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceDataLossPreventionStoredInfoTypeEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	newObj := make(map[string]interface{})
	newObj["config"] = obj
	return newObj, nil
}

func resourceDataLossPreventionStoredInfoTypeDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// Stored info types come back nested with previous versions. We only want the current
	// version in the unwrapped form
	name := res["name"].(string)
	v, ok := res["currentVersion"]
	if !ok || v == nil {
		return nil, nil
	}

	current := v.(map[string]interface{})
	configRaw, ok := current["config"]
	if !ok || configRaw == nil {
		return nil, nil
	}

	config := configRaw.(map[string]interface{})
	// Name comes back on the top level, so set here
	config["name"] = name

	return config, nil
}
