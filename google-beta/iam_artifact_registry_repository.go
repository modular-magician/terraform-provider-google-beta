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

var ArtifactRegistryRepositoryIamSchema = map[string]*schema.Schema{
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
	"repository": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type ArtifactRegistryRepositoryIamUpdater struct {
	project    string
	location   string
	repository string
	d          tpgresource.TerraformResourceData
	Config     *transport_tpg.Config
}

func ArtifactRegistryRepositoryIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
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
	if v, ok := d.GetOk("repository"); ok {
		values["repository"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/repositories/(?P<repository>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<repository>[^/]+)", "(?P<location>[^/]+)/(?P<repository>[^/]+)", "(?P<repository>[^/]+)"}, d, config, d.Get("repository").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ArtifactRegistryRepositoryIamUpdater{
		project:    values["project"],
		location:   values["location"],
		repository: values["repository"],
		d:          d,
		Config:     config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("location", u.location); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("repository", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting repository: %s", err)
	}

	return u, nil
}

func ArtifactRegistryRepositoryIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	location, _ := getLocation(d, config)
	if location != "" {
		values["location"] = location
	}

	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/repositories/(?P<repository>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<repository>[^/]+)", "(?P<location>[^/]+)/(?P<repository>[^/]+)", "(?P<repository>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ArtifactRegistryRepositoryIamUpdater{
		project:    values["project"],
		location:   values["location"],
		repository: values["repository"],
		d:          d,
		Config:     config,
	}
	if err := d.Set("repository", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting repository: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *ArtifactRegistryRepositoryIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyRepositoryUrl("getIamPolicy")
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

	policy, err := transport_tpg.SendRequest(u.Config, "GET", project, url, userAgent, obj)
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

func (u *ArtifactRegistryRepositoryIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyRepositoryUrl("setIamPolicy")
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

func (u *ArtifactRegistryRepositoryIamUpdater) qualifyRepositoryUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{ArtifactRegistryBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/repositories/%s", u.project, u.location, u.repository), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *ArtifactRegistryRepositoryIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/repositories/%s", u.project, u.location, u.repository)
}

func (u *ArtifactRegistryRepositoryIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-artifactregistry-repository-%s", u.GetResourceId())
}

func (u *ArtifactRegistryRepositoryIamUpdater) DescribeResource() string {
	return fmt.Sprintf("artifactregistry repository %q", u.GetResourceId())
}
