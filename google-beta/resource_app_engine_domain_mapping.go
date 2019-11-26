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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func sslSettingsDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	// If certificate id is empty, and ssl management type is `MANUAL`, then
	// ssl settings will not be configured, and ssl_settings block is not returned

	if k == "ssl_settings.#" &&
		old == "0" && new == "1" &&
		d.Get("ssl_settings.0.certificate_id") == "" &&
		d.Get("ssl_settings.0.ssl_management_type") == "MANUAL" {
		return true
	}

	return false
}

func resourceAppEngineDomainMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppEngineDomainMappingCreate,
		Read:   resourceAppEngineDomainMappingRead,
		Update: resourceAppEngineDomainMappingUpdate,
		Delete: resourceAppEngineDomainMappingDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAppEngineDomainMappingImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Relative name of the domain serving the application. Example: example.com.`,
			},
			"override_strategy": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"STRICT", "OVERRIDE", ""}, false),
				Description: `Whether the domain creation should override any existing mappings for this domain.
By default, overrides are rejected.`,
				Default: "STRICT",
			},
			"ssl_settings": {
				Type:             schema.TypeList,
				Optional:         true,
				DiffSuppressFunc: sslSettingsDiffSuppress,
				Description:      `SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.`,
				MaxItems:         1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ssl_management_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"AUTOMATIC", "MANUAL"}, false),
							Description: `SSL management type for this domain. If 'AUTOMATIC', a managed certificate is automatically provisioned.
If 'MANUAL', 'certificateId' must be manually specified in order to configure SSL for this domain.`,
						},
						"certificate_id": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							Description: `ID of the AuthorizedCertificate resource configuring SSL for the application. Clearing this field will
remove SSL support.
By default, a managed certificate is automatically created for every domain mapping. To omit SSL support
or to configure SSL manually, specify 'SslManagementType.MANUAL' on a 'CREATE' or 'UPDATE' request. You must be
authorized to administer the 'AuthorizedCertificate' resource to manually map it to a DomainMapping resource.
Example: 12345.`,
						},
						"pending_managed_certificate_id": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `ID of the managed 'AuthorizedCertificate' resource currently being provisioned, if applicable. Until the new
managed certificate has been successfully provisioned, the previous SSL state will be preserved. Once the
provisioning process completes, the 'certificateId' field will reflect the new managed certificate and this
field will be left empty. To remove SSL support while there is still a pending managed certificate, clear the
'certificateId' field with an update request.`,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.`,
			},
			"resource_records": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `The resource records required to configure this domain mapping. These records must be added to the domain's DNS
configuration in order to serve the application via this domain mapping.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Relative name of the object affected by this record. Only applicable for CNAME records. Example: 'www'.`,
						},
						"rrdata": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Data for this record. Values vary by record type, as defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1).`,
						},
						"type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"A", "AAAA", "CNAME", ""}, false),
							Description:  `Resource record type. Example: 'AAAA'.`,
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

func resourceAppEngineDomainMappingCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	sslSettingsProp, err := expandAppEngineDomainMappingSslSettings(d.Get("ssl_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ssl_settings"); !isEmptyValue(reflect.ValueOf(sslSettingsProp)) && (ok || !reflect.DeepEqual(v, sslSettingsProp)) {
		obj["sslSettings"] = sslSettingsProp
	}
	idProp, err := expandAppEngineDomainMappingDomainName(d.Get("domain_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("domain_name"); !isEmptyValue(reflect.ValueOf(idProp)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}

	url, err := replaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/domainMappings")
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
	id, err := replaceVars(d, config, "apps/{{project}}/domainMappings/{{domain_name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	waitErr := appEngineOperationWaitTime(
		config, res, project, "Creating DomainMapping",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create DomainMapping: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating DomainMapping %q: %#v", d.Id(), res)

	return resourceAppEngineDomainMappingRead(d, meta)
}

func resourceAppEngineDomainMappingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/domainMappings/{{domain_name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AppEngineDomainMapping %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}

	if err := d.Set("name", flattenAppEngineDomainMappingName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}
	if err := d.Set("ssl_settings", flattenAppEngineDomainMappingSslSettings(res["sslSettings"], d)); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}
	if err := d.Set("resource_records", flattenAppEngineDomainMappingResourceRecords(res["resourceRecords"], d)); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}
	if err := d.Set("domain_name", flattenAppEngineDomainMappingDomainName(res["id"], d)); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}

	return nil
}

func resourceAppEngineDomainMappingUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	sslSettingsProp, err := expandAppEngineDomainMappingSslSettings(d.Get("ssl_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ssl_settings"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sslSettingsProp)) {
		obj["sslSettings"] = sslSettingsProp
	}

	url, err := replaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/domainMappings/{{domain_name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating DomainMapping %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("ssl_settings") {
		updateMask = append(updateMask, "ssl_settings.certificate_id,ssl_settings.ssl_management_type")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating DomainMapping %q: %s", d.Id(), err)
	}

	err = appEngineOperationWaitTime(
		config, res, project, "Updating DomainMapping",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))
	if err != nil {
		return err
	}

	return resourceAppEngineDomainMappingRead(d, meta)
}

func resourceAppEngineDomainMappingDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/domainMappings/{{domain_name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting DomainMapping %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "DomainMapping")
	}

	err = appEngineOperationWaitTime(
		config, res, project, "Deleting DomainMapping",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting DomainMapping %q: %#v", d.Id(), res)
	return nil
}

func resourceAppEngineDomainMappingImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"apps/(?P<project>[^/]+)/domainMappings/(?P<domain_name>[^/]+)",
		"(?P<project>[^/]+)/(?P<domain_name>[^/]+)",
		"(?P<domain_name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "apps/{{project}}/domainMappings/{{domain_name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAppEngineDomainMappingName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineDomainMappingSslSettings(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["certificate_id"] =
		flattenAppEngineDomainMappingSslSettingsCertificateId(original["certificateId"], d)
	transformed["ssl_management_type"] =
		flattenAppEngineDomainMappingSslSettingsSslManagementType(original["sslManagementType"], d)
	transformed["pending_managed_certificate_id"] =
		flattenAppEngineDomainMappingSslSettingsPendingManagedCertificateId(original["pendingManagedCertificateId"], d)
	return []interface{}{transformed}
}
func flattenAppEngineDomainMappingSslSettingsCertificateId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineDomainMappingSslSettingsSslManagementType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineDomainMappingSslSettingsPendingManagedCertificateId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineDomainMappingResourceRecords(v interface{}, d *schema.ResourceData) interface{} {
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
			"name":   flattenAppEngineDomainMappingResourceRecordsName(original["name"], d),
			"rrdata": flattenAppEngineDomainMappingResourceRecordsRrdata(original["rrdata"], d),
			"type":   flattenAppEngineDomainMappingResourceRecordsType(original["type"], d),
		})
	}
	return transformed
}
func flattenAppEngineDomainMappingResourceRecordsName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineDomainMappingResourceRecordsRrdata(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineDomainMappingResourceRecordsType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAppEngineDomainMappingDomainName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandAppEngineDomainMappingSslSettings(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCertificateId, err := expandAppEngineDomainMappingSslSettingsCertificateId(original["certificate_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertificateId); val.IsValid() && !isEmptyValue(val) {
		transformed["certificateId"] = transformedCertificateId
	}

	transformedSslManagementType, err := expandAppEngineDomainMappingSslSettingsSslManagementType(original["ssl_management_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSslManagementType); val.IsValid() && !isEmptyValue(val) {
		transformed["sslManagementType"] = transformedSslManagementType
	}

	transformedPendingManagedCertificateId, err := expandAppEngineDomainMappingSslSettingsPendingManagedCertificateId(original["pending_managed_certificate_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPendingManagedCertificateId); val.IsValid() && !isEmptyValue(val) {
		transformed["pendingManagedCertificateId"] = transformedPendingManagedCertificateId
	}

	return transformed, nil
}

func expandAppEngineDomainMappingSslSettingsCertificateId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineDomainMappingSslSettingsSslManagementType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineDomainMappingSslSettingsPendingManagedCertificateId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineDomainMappingDomainName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
