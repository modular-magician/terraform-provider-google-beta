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

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var DataCatalogTaxonomyIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"region": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"taxonomy": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type DataCatalogTaxonomyIamUpdater struct {
	project  string
	region   string
	taxonomy string
	d        tpgresource.TerraformResourceData
	Config   *transport_tpg.Config
}

func DataCatalogTaxonomyIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	region, _ := getRegion(d, config)
	if region != "" {
		if err := d.Set("region", region); err != nil {
			return nil, fmt.Errorf("Error setting region: %s", err)
		}
	}
	values["region"] = region
	if v, ok := d.GetOk("taxonomy"); ok {
		values["taxonomy"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/taxonomies/(?P<taxonomy>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<taxonomy>[^/]+)", "(?P<region>[^/]+)/(?P<taxonomy>[^/]+)", "(?P<taxonomy>[^/]+)"}, d, config, d.Get("taxonomy").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataCatalogTaxonomyIamUpdater{
		project:  values["project"],
		region:   values["region"],
		taxonomy: values["taxonomy"],
		d:        d,
		Config:   config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("region", u.region); err != nil {
		return nil, fmt.Errorf("Error setting region: %s", err)
	}
	if err := d.Set("taxonomy", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting taxonomy: %s", err)
	}

	return u, nil
}

func DataCatalogTaxonomyIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	region, _ := getRegion(d, config)
	if region != "" {
		values["region"] = region
	}

	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/taxonomies/(?P<taxonomy>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<taxonomy>[^/]+)", "(?P<region>[^/]+)/(?P<taxonomy>[^/]+)", "(?P<taxonomy>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataCatalogTaxonomyIamUpdater{
		project:  values["project"],
		region:   values["region"],
		taxonomy: values["taxonomy"],
		d:        d,
		Config:   config,
	}
	if err := d.Set("taxonomy", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting taxonomy: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *DataCatalogTaxonomyIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyTaxonomyUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(u.Config, "POST", project, url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = tpgresource.Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *DataCatalogTaxonomyIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyTaxonomyUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequestWithTimeout(u.Config, "POST", project, url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *DataCatalogTaxonomyIamUpdater) qualifyTaxonomyUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{DataCatalogBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", u.project, u.region, u.taxonomy), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *DataCatalogTaxonomyIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", u.project, u.region, u.taxonomy)
}

func (u *DataCatalogTaxonomyIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-datacatalog-taxonomy-%s", u.GetResourceId())
}

func (u *DataCatalogTaxonomyIamUpdater) DescribeResource() string {
	return fmt.Sprintf("datacatalog taxonomy %q", u.GetResourceId())
}
