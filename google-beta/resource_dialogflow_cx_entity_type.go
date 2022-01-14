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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDialogflowCXEntityType() *schema.Resource {
	return &schema.Resource{
		Create: resourceDialogflowCXEntityTypeCreate,
		Read:   resourceDialogflowCXEntityTypeRead,
		Update: resourceDialogflowCXEntityTypeUpdate,
		Delete: resourceDialogflowCXEntityTypeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDialogflowCXEntityTypeImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(40 * time.Minute),
			Update: schema.DefaultTimeout(40 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Description:  `The human-readable name of the entity type, unique within the agent.`,
			},
			"entities": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `The collection of entity entries associated with the entity type.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"synonyms": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `A collection of value synonyms. For example, if the entity type is vegetable, and value is scallions, a synonym could be green onions.
For KIND_LIST entity types: This collection must contain exactly one synonym equal to value.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The primary value associated with this entity entry. For example, if the entity type is vegetable, the value could be scallions.
For KIND_MAP entity types: A canonical value to be used in place of synonyms.
For KIND_LIST entity types: A string that can contain references to other entity types (with or without aliases).`,
						},
					},
				},
			},
			"kind": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateEnum([]string{"KIND_MAP", "KIND_LIST", "KIND_REGEXP"}),
				Description: `Indicates whether the entity type can be automatically expanded.
* KIND_MAP: Map entity types allow mapping of a group of synonyms to a canonical value.
* KIND_LIST: List entity types contain a set of entries that do not map to canonical values. However, list entity types can contain references to other entity types (with or without aliases).
* KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values. Possible values: ["KIND_MAP", "KIND_LIST", "KIND_REGEXP"]`,
			},
			"auto_expansion_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateEnum([]string{"AUTO_EXPANSION_MODE_DEFAULT", "AUTO_EXPANSION_MODE_UNSPECIFIED", ""}),
				Description: `Represents kinds of entities.
* AUTO_EXPANSION_MODE_UNSPECIFIED: Auto expansion disabled for the entity.
* AUTO_EXPANSION_MODE_DEFAULT: Allows an agent to recognize values that have not been explicitly listed in the entity. Possible values: ["AUTO_EXPANSION_MODE_DEFAULT", "AUTO_EXPANSION_MODE_UNSPECIFIED"]`,
			},
			"enable_fuzzy_extraction": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Enables fuzzy entity extraction during classification.`,
			},
			"excluded_phrases": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry giant(an adjective), you might consider adding giants(a noun) as an exclusion. 
If the kind of entity type is KIND_MAP, then the phrases specified by entities and excluded phrases should be mutually exclusive.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The word or phrase to be excluded.`,
						},
					},
				},
			},
			"language_code": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The language of the following fields in entityType:
EntityType.entities.value
EntityType.entities.synonyms
EntityType.excluded_phrases.value
If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.`,
			},
			"parent": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The agent to create a entity type for. 
Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.`,
			},
			"redact": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The unique identifier of the entity type.
Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/entityTypes/<Entity Type ID>.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceDialogflowCXEntityTypeCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXEntityTypeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	kindProp, err := expandDialogflowCXEntityTypeKind(d.Get("kind"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kind"); !isEmptyValue(reflect.ValueOf(kindProp)) && (ok || !reflect.DeepEqual(v, kindProp)) {
		obj["kind"] = kindProp
	}
	autoExpansionModeProp, err := expandDialogflowCXEntityTypeAutoExpansionMode(d.Get("auto_expansion_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("auto_expansion_mode"); !isEmptyValue(reflect.ValueOf(autoExpansionModeProp)) && (ok || !reflect.DeepEqual(v, autoExpansionModeProp)) {
		obj["autoExpansionMode"] = autoExpansionModeProp
	}
	entitiesProp, err := expandDialogflowCXEntityTypeEntities(d.Get("entities"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("entities"); !isEmptyValue(reflect.ValueOf(entitiesProp)) && (ok || !reflect.DeepEqual(v, entitiesProp)) {
		obj["entities"] = entitiesProp
	}
	excludedPhrasesProp, err := expandDialogflowCXEntityTypeExcludedPhrases(d.Get("excluded_phrases"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("excluded_phrases"); !isEmptyValue(reflect.ValueOf(excludedPhrasesProp)) && (ok || !reflect.DeepEqual(v, excludedPhrasesProp)) {
		obj["excludedPhrases"] = excludedPhrasesProp
	}
	enableFuzzyExtractionProp, err := expandDialogflowCXEntityTypeEnableFuzzyExtraction(d.Get("enable_fuzzy_extraction"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_fuzzy_extraction"); !isEmptyValue(reflect.ValueOf(enableFuzzyExtractionProp)) && (ok || !reflect.DeepEqual(v, enableFuzzyExtractionProp)) {
		obj["enableFuzzyExtraction"] = enableFuzzyExtractionProp
	}
	redactProp, err := expandDialogflowCXEntityTypeRedact(d.Get("redact"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redact"); !isEmptyValue(reflect.ValueOf(redactProp)) && (ok || !reflect.DeepEqual(v, redactProp)) {
		obj["redact"] = redactProp
	}

	url, err := replaceVars(d, config, "{{DialogflowCXBasePath}}{{parent}}/entityTypes")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new EntityType: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating EntityType: %s", err)
	}
	if err := d.Set("name", flattenDialogflowCXEntityTypeName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{parent}}/entityTypes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating EntityType %q: %#v", d.Id(), res)

	return resourceDialogflowCXEntityTypeRead(d, meta)
}

func resourceDialogflowCXEntityTypeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{DialogflowCXBasePath}}{{parent}}/entityTypes/{{name}}")
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
		return handleNotFoundError(err, d, fmt.Sprintf("DialogflowCXEntityType %q", d.Id()))
	}

	if err := d.Set("name", flattenDialogflowCXEntityTypeName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntityType: %s", err)
	}
	if err := d.Set("display_name", flattenDialogflowCXEntityTypeDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntityType: %s", err)
	}
	if err := d.Set("kind", flattenDialogflowCXEntityTypeKind(res["kind"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntityType: %s", err)
	}
	if err := d.Set("auto_expansion_mode", flattenDialogflowCXEntityTypeAutoExpansionMode(res["autoExpansionMode"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntityType: %s", err)
	}
	if err := d.Set("entities", flattenDialogflowCXEntityTypeEntities(res["entities"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntityType: %s", err)
	}
	if err := d.Set("excluded_phrases", flattenDialogflowCXEntityTypeExcludedPhrases(res["excludedPhrases"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntityType: %s", err)
	}
	if err := d.Set("enable_fuzzy_extraction", flattenDialogflowCXEntityTypeEnableFuzzyExtraction(res["enableFuzzyExtraction"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntityType: %s", err)
	}
	if err := d.Set("redact", flattenDialogflowCXEntityTypeRedact(res["redact"], d, config)); err != nil {
		return fmt.Errorf("Error reading EntityType: %s", err)
	}

	return nil
}

func resourceDialogflowCXEntityTypeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXEntityTypeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	kindProp, err := expandDialogflowCXEntityTypeKind(d.Get("kind"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kind"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, kindProp)) {
		obj["kind"] = kindProp
	}
	autoExpansionModeProp, err := expandDialogflowCXEntityTypeAutoExpansionMode(d.Get("auto_expansion_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("auto_expansion_mode"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, autoExpansionModeProp)) {
		obj["autoExpansionMode"] = autoExpansionModeProp
	}
	entitiesProp, err := expandDialogflowCXEntityTypeEntities(d.Get("entities"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("entities"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, entitiesProp)) {
		obj["entities"] = entitiesProp
	}
	excludedPhrasesProp, err := expandDialogflowCXEntityTypeExcludedPhrases(d.Get("excluded_phrases"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("excluded_phrases"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, excludedPhrasesProp)) {
		obj["excludedPhrases"] = excludedPhrasesProp
	}
	enableFuzzyExtractionProp, err := expandDialogflowCXEntityTypeEnableFuzzyExtraction(d.Get("enable_fuzzy_extraction"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_fuzzy_extraction"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableFuzzyExtractionProp)) {
		obj["enableFuzzyExtraction"] = enableFuzzyExtractionProp
	}
	redactProp, err := expandDialogflowCXEntityTypeRedact(d.Get("redact"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redact"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, redactProp)) {
		obj["redact"] = redactProp
	}

	url, err := replaceVars(d, config, "{{DialogflowCXBasePath}}{{parent}}/entityTypes/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating EntityType %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("kind") {
		updateMask = append(updateMask, "kind")
	}

	if d.HasChange("auto_expansion_mode") {
		updateMask = append(updateMask, "autoExpansionMode")
	}

	if d.HasChange("entities") {
		updateMask = append(updateMask, "entities")
	}

	if d.HasChange("excluded_phrases") {
		updateMask = append(updateMask, "excludedPhrases")
	}

	if d.HasChange("enable_fuzzy_extraction") {
		updateMask = append(updateMask, "enableFuzzyExtraction")
	}

	if d.HasChange("redact") {
		updateMask = append(updateMask, "redact")
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
		return fmt.Errorf("Error updating EntityType %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating EntityType %q: %#v", d.Id(), res)
	}

	return resourceDialogflowCXEntityTypeRead(d, meta)
}

func resourceDialogflowCXEntityTypeDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := replaceVars(d, config, "{{DialogflowCXBasePath}}{{parent}}/entityTypes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting EntityType %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "EntityType")
	}

	log.Printf("[DEBUG] Finished deleting EntityType %q: %#v", d.Id(), res)
	return nil
}

func resourceDialogflowCXEntityTypeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value and parent contains slashes
	if err := parseImportId([]string{
		"(?P<parent>.+)/entityTypes/(?P<name>[^/]+)",
		"(?P<parent>.+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{parent}}/entityTypes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDialogflowCXEntityTypeName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenDialogflowCXEntityTypeDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXEntityTypeKind(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXEntityTypeAutoExpansionMode(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXEntityTypeEntities(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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
			"value":    flattenDialogflowCXEntityTypeEntitiesValue(original["value"], d, config),
			"synonyms": flattenDialogflowCXEntityTypeEntitiesSynonyms(original["synonyms"], d, config),
		})
	}
	return transformed
}
func flattenDialogflowCXEntityTypeEntitiesValue(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXEntityTypeEntitiesSynonyms(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXEntityTypeExcludedPhrases(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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
			"value": flattenDialogflowCXEntityTypeExcludedPhrasesValue(original["value"], d, config),
		})
	}
	return transformed
}
func flattenDialogflowCXEntityTypeExcludedPhrasesValue(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXEntityTypeEnableFuzzyExtraction(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDialogflowCXEntityTypeRedact(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandDialogflowCXEntityTypeDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXEntityTypeKind(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXEntityTypeAutoExpansionMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXEntityTypeEntities(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedValue, err := expandDialogflowCXEntityTypeEntitiesValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !isEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		transformedSynonyms, err := expandDialogflowCXEntityTypeEntitiesSynonyms(original["synonyms"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSynonyms); val.IsValid() && !isEmptyValue(val) {
			transformed["synonyms"] = transformedSynonyms
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXEntityTypeEntitiesValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXEntityTypeEntitiesSynonyms(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXEntityTypeExcludedPhrases(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedValue, err := expandDialogflowCXEntityTypeExcludedPhrasesValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !isEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXEntityTypeExcludedPhrasesValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXEntityTypeEnableFuzzyExtraction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXEntityTypeRedact(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
