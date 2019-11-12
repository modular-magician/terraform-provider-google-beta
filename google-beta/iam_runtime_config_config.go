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

var RuntimeConfigConfigIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"config": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type RuntimeConfigConfigIamUpdater struct {
	project string
	config  string
	d       *schema.ResourceData
	Config  *Config
}

func RuntimeConfigConfigIamUpdaterProducer(d *schema.ResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}
	values["project"] = project
	if v, ok := d.GetOk("config"); ok {
		values["config"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/configs/(?P<config>[^/]+)", "(?P<project>[^/]+)/(?P<config>[^/]+)", "(?P<config>[^/]+)"}, d, config, d.Get("config").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &RuntimeConfigConfigIamUpdater{
		project: values["project"],
		config:  values["config"],
		d:       d,
		Config:  config,
	}

	d.Set("project", u.project)
	d.Set("config", u.GetResourceId())

	return u, nil
}

func RuntimeConfigConfigIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	values["project"] = project

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/configs/(?P<config>[^/]+)", "(?P<project>[^/]+)/(?P<config>[^/]+)", "(?P<config>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &RuntimeConfigConfigIamUpdater{
		project: values["project"],
		config:  values["config"],
		d:       d,
		Config:  config,
	}
	d.Set("config", u.GetResourceId())

	return nil
}

func (u *RuntimeConfigConfigIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url := u.qualifyConfigUrl("getIamPolicy")

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	policy, err := sendRequest(u.Config, "GET", project, url, obj)
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

func (u *RuntimeConfigConfigIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url := u.qualifyConfigUrl("setIamPolicy")

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

func (u *RuntimeConfigConfigIamUpdater) qualifyConfigUrl(methodIdentifier string) string {
	return fmt.Sprintf("https://runtimeconfig.googleapis.com/v1beta1/%s:%s", fmt.Sprintf("projects/%s/configs/%s", u.project, u.config), methodIdentifier)
}

func (u *RuntimeConfigConfigIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/configs/%s", u.project, u.config)
}

func (u *RuntimeConfigConfigIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-runtimeconfig-config-%s", u.GetResourceId())
}

func (u *RuntimeConfigConfigIamUpdater) DescribeResource() string {
	return fmt.Sprintf("runtimeconfig config %q", u.GetResourceId())
}
