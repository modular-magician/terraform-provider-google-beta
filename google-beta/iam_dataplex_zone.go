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

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var DataplexZoneIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"location": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"lake": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"dataplex_zone": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type DataplexZoneIamUpdater struct {
	project      string
	location     string
	lake         string
	dataplexZone string
	d            TerraformResourceData
	Config       *transport_tpg.Config
}

func DataplexZoneIamUpdaterProducer(d TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	location, _ := getLocation(d, config)
	if location != "" {
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	values["location"] = location
	if v, ok := d.GetOk("lake"); ok {
		values["lake"] = v.(string)
	}

	if v, ok := d.GetOk("dataplex_zone"); ok {
		values["dataplex_zone"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/lakes/(?P<lake>[^/]+)/zones/(?P<dataplex_zone>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<lake>[^/]+)/(?P<dataplex_zone>[^/]+)", "(?P<location>[^/]+)/(?P<lake>[^/]+)/(?P<dataplex_zone>[^/]+)", "(?P<dataplex_zone>[^/]+)"}, d, config, d.Get("dataplex_zone").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataplexZoneIamUpdater{
		project:      values["project"],
		location:     values["location"],
		lake:         values["lake"],
		dataplexZone: values["dataplex_zone"],
		d:            d,
		Config:       config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("location", u.location); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("lake", u.lake); err != nil {
		return nil, fmt.Errorf("Error setting lake: %s", err)
	}
	if err := d.Set("dataplex_zone", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting dataplex_zone: %s", err)
	}

	return u, nil
}

func DataplexZoneIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	location, _ := getLocation(d, config)
	if location != "" {
		values["location"] = location
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/lakes/(?P<lake>[^/]+)/zones/(?P<dataplex_zone>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<lake>[^/]+)/(?P<dataplex_zone>[^/]+)", "(?P<location>[^/]+)/(?P<lake>[^/]+)/(?P<dataplex_zone>[^/]+)", "(?P<dataplex_zone>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataplexZoneIamUpdater{
		project:      values["project"],
		location:     values["location"],
		lake:         values["lake"],
		dataplexZone: values["dataplex_zone"],
		d:            d,
		Config:       config,
	}
	if err := d.Set("dataplex_zone", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting dataplex_zone: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *DataplexZoneIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyZoneUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(u.Config, "GET", project, url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *DataplexZoneIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyZoneUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequestWithTimeout(u.Config, "POST", project, url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *DataplexZoneIamUpdater) qualifyZoneUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{DataplexBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s", u.project, u.location, u.lake, u.dataplexZone), methodIdentifier)
	url, err := ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *DataplexZoneIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s", u.project, u.location, u.lake, u.dataplexZone)
}

func (u *DataplexZoneIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-dataplex-zone-%s", u.GetResourceId())
}

func (u *DataplexZoneIamUpdater) DescribeResource() string {
	return fmt.Sprintf("dataplex zone %q", u.GetResourceId())
}
