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

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var IapWebTypeComputeIamSchema = map[string]*schema.Schema{
	"project": {
		Type:             schema.TypeString,
		Computed:         true,
		Optional:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type IapWebTypeComputeIamUpdater struct {
	project string
	d       *schema.ResourceData
	Config  *Config
}

func IapWebTypeComputeIamUpdaterProducer(d *schema.ResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}
	values["project"] = project

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/iap_web/compute", "(?P<project>[^/]+)"}, d, config, d.Get("project").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &IapWebTypeComputeIamUpdater{
		project: values["project"],
		d:       d,
		Config:  config,
	}

	d.Set("project", u.project)

	return u, nil
}

func IapWebTypeComputeIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	values["project"] = project

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/iap_web/compute", "(?P<project>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &IapWebTypeComputeIamUpdater{
		project: values["project"],
		d:       d,
		Config:  config,
	}
	d.Set("project", u.project)

	return nil
}

func (u *IapWebTypeComputeIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url := u.qualifyWebTypeComputeUrl("getIamPolicy")

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	policy, err := sendRequest(u.Config, "POST", project, url, obj)
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

func (u *IapWebTypeComputeIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url := u.qualifyWebTypeComputeUrl("setIamPolicy")

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	_, err = sendRequestWithTimeout(u.Config, "POST", project, url, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *IapWebTypeComputeIamUpdater) qualifyWebTypeComputeUrl(methodIdentifier string) string {
	return fmt.Sprintf("https://iap.googleapis.com/v1/%s:%s", fmt.Sprintf("projects/%s/iap_web/compute", u.project), methodIdentifier)
}

func (u *IapWebTypeComputeIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/iap_web/compute", u.project)
}

func (u *IapWebTypeComputeIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-iap-webtypecompute-%s", u.GetResourceId())
}

func (u *IapWebTypeComputeIamUpdater) DescribeResource() string {
	return fmt.Sprintf("iap webtypecompute %q", u.GetResourceId())
}
