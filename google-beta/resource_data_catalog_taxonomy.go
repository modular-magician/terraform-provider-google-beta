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
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceDataCatalogTaxonomy() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataCatalogTaxonomyCreate,
		Read:   resourceDataCatalogTaxonomyRead,
		Update: resourceDataCatalogTaxonomyUpdate,
		Delete: resourceDataCatalogTaxonomyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataCatalogTaxonomyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				Description: `User defined name of this taxonomy.
It must: contain only unicode letters, numbers, underscores, dashes
and spaces; not start or end with spaces; and be at most 200 bytes
long when encoded in UTF-8.`,
			},
			"activated_policy_types": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A list of policy types that are activated for this taxonomy. If not set,
defaults to an empty list. Possible values: ["POLICY_TYPE_UNSPECIFIED", "FINE_GRAINED_ACCESS_CONTROL"]`,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validateEnum([]string{"POLICY_TYPE_UNSPECIFIED", "FINE_GRAINED_ACCESS_CONTROL"}),
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Description of this taxonomy. It must: contain only unicode characters,
tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes
long when encoded in UTF-8. If not set, defaults to an empty description.`,
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `Taxonomy location region.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Resource name of this taxonomy, whose format is:
"projects/{project}/locations/{region}/taxonomies/{taxonomy}".`,
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

func resourceDataCatalogTaxonomyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandDataCatalogTaxonomyDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandDataCatalogTaxonomyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	activatedPolicyTypesProp, err := expandDataCatalogTaxonomyActivatedPolicyTypes(d.Get("activated_policy_types"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("activated_policy_types"); !isEmptyValue(reflect.ValueOf(activatedPolicyTypesProp)) && (ok || !reflect.DeepEqual(v, activatedPolicyTypesProp)) {
		obj["activatedPolicyTypes"] = activatedPolicyTypesProp
	}

	url, err := ReplaceVars(d, config, "{{DataCatalogBasePath}}projects/{{project}}/locations/{{region}}/taxonomies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Taxonomy: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Taxonomy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Taxonomy: %s", err)
	}
	if err := d.Set("name", flattenDataCatalogTaxonomyName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Taxonomy %q: %#v", d.Id(), res)

	return resourceDataCatalogTaxonomyRead(d, meta)
}

func resourceDataCatalogTaxonomyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Taxonomy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("DataCatalogTaxonomy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Taxonomy: %s", err)
	}

	if err := d.Set("name", flattenDataCatalogTaxonomyName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Taxonomy: %s", err)
	}
	if err := d.Set("display_name", flattenDataCatalogTaxonomyDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Taxonomy: %s", err)
	}
	if err := d.Set("description", flattenDataCatalogTaxonomyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Taxonomy: %s", err)
	}
	if err := d.Set("activated_policy_types", flattenDataCatalogTaxonomyActivatedPolicyTypes(res["activatedPolicyTypes"], d, config)); err != nil {
		return fmt.Errorf("Error reading Taxonomy: %s", err)
	}

	return nil
}

func resourceDataCatalogTaxonomyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Taxonomy: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandDataCatalogTaxonomyDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandDataCatalogTaxonomyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	activatedPolicyTypesProp, err := expandDataCatalogTaxonomyActivatedPolicyTypes(d.Get("activated_policy_types"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("activated_policy_types"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, activatedPolicyTypesProp)) {
		obj["activatedPolicyTypes"] = activatedPolicyTypesProp
	}

	url, err := ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Taxonomy %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("activated_policy_types") {
		updateMask = append(updateMask, "activatedPolicyTypes")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Taxonomy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Taxonomy %q: %#v", d.Id(), res)
	}

	return resourceDataCatalogTaxonomyRead(d, meta)
}

func resourceDataCatalogTaxonomyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Taxonomy: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Taxonomy %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Taxonomy")
	}

	log.Printf("[DEBUG] Finished deleting Taxonomy %q: %#v", d.Id(), res)
	return nil
}

func resourceDataCatalogTaxonomyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	name := d.Get("name").(string)
	d.SetId(name)

	re := regexp.MustCompile("projects/(.+)/(?:locations|regions)/(.+)/taxonomies/(.+)")
	if matches := re.FindStringSubmatch(name); matches != nil {
		d.Set("project", matches[1])
		d.Set("region", matches[2])
	}

	return []*schema.ResourceData{d}, nil
}

func flattenDataCatalogTaxonomyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogTaxonomyDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogTaxonomyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogTaxonomyActivatedPolicyTypes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDataCatalogTaxonomyDisplayName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTaxonomyDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTaxonomyActivatedPolicyTypes(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
