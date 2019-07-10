package google

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var SourceRepoRepositoryIamSchema = map[string]*schema.Schema{
	"project": {
		Type:             schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew:         true,
	},
	"repository": {
		Type:             schema.TypeString,
		Required: true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type SourceRepoRepositoryIamUpdater struct {
	project string
	repository string
	d       *schema.ResourceData
	Config  *Config
}

func SourceRepoRepositoryIamUpdaterProducer(d *schema.ResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)
	
	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}

	// While this may be overridden by the "project" value from getImportIdQualifiers below,
	// setting project here ensures the value is set even if the value set in config is the short
	// name or otherwise doesn't include the project.
	values["project"] = project

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/repos/(?P<repository>[^/]+)","(?P<project>[^/]+)/(?P<repository>[^/]+)","(?P<repository>[^/]+)"}, d, config, d.Get("repository").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &SourceRepoRepositoryIamUpdater{
		project: values["project"],
		repository: values["repository"],
		d:       d,
		Config:  config,
	}
	d.SetId(u.GetResourceId())

	return u, nil
}

func SourceRepoRepositoryIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)
	
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	values["project"] = project

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/repos/(?P<repository>[^/]+)","(?P<project>[^/]+)/(?P<repository>[^/]+)","(?P<repository>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
    values[k] = v
	}

	u := &SourceRepoRepositoryIamUpdater{
		project: values["project"],
		repository: values["repository"],
		d:       d,
		Config:  config,
	}
	d.Set("repository", u.GetResourceId())
	d.SetId(u.GetResourceId())
	return nil
}

func (u *SourceRepoRepositoryIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url := u.qualifyRepositoryUrl("getIamPolicy")

	policy, err := sendRequest(u.Config, "GET", url, nil)
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

func (u *SourceRepoRepositoryIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url := u.qualifyRepositoryUrl("setIamPolicy")
	
	_, err = sendRequestWithTimeout(u.Config, "POST", url, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *<+ resource_name -%>IamUpdater) qualifyRepositoryUrl(methodIdentifier string) string {
	return fmt.Sprintf("https://sourcerepo.googleapis.com/v1/%s:%s", fmt.Sprintf("projects/%s/repos/%s", u.project, u.repository), methodIdentifier)
}

func (u *SourceRepoRepositoryIamUpdater) GetResourceId() string {
	return fmt.Sprintf("%s/%s", u.project, u.repository)
}

func (u *SourceRepoRepositoryIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-sourcerepo-repository-%s", u.GetResourceId())
}

func (u *SourceRepoRepositoryIamUpdater) DescribeResource() string {
	return fmt.Sprintf("sourcerepo repository %q", u.GetResourceId())
}
