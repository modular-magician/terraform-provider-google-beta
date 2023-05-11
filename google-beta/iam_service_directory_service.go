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

var ServiceDirectoryServiceIamSchema = map[string]*schema.Schema{
	"name": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type ServiceDirectoryServiceIamUpdater struct {
	name   string
	d      tpgresource.TerraformResourceData
	Config *transport_tpg.Config
}

func ServiceDirectoryServiceIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	if v, ok := d.GetOk("name"); ok {
		values["name"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/namespaces/(?P<namespace_id>[^/]+)/services/(?P<service_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<namespace_id>[^/]+)/(?P<service_id>[^/]+)", "(?P<location>[^/]+)/(?P<namespace_id>[^/]+)/(?P<service_id>[^/]+)"}, d, config, d.Get("name").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ServiceDirectoryServiceIamUpdater{
		name:   values["name"],
		d:      d,
		Config: config,
	}

	if err := d.Set("name", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}

	return u, nil
}

func ServiceDirectoryServiceIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/namespaces/(?P<namespace_id>[^/]+)/services/(?P<service_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<namespace_id>[^/]+)/(?P<service_id>[^/]+)", "(?P<location>[^/]+)/(?P<namespace_id>[^/]+)/(?P<service_id>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ServiceDirectoryServiceIamUpdater{
		name:   values["name"],
		d:      d,
		Config: config,
	}
	if err := d.Set("name", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *ServiceDirectoryServiceIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyServiceUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	var obj map[string]interface{}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(u.Config, "POST", "", url, userAgent, obj)
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

func (u *ServiceDirectoryServiceIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyServiceUrl("setIamPolicy")
	if err != nil {
		return err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequestWithTimeout(u.Config, "POST", "", url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *ServiceDirectoryServiceIamUpdater) qualifyServiceUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{ServiceDirectoryBasePath}}%s:%s", fmt.Sprintf("%s", u.name), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *ServiceDirectoryServiceIamUpdater) GetResourceId() string {
	return fmt.Sprintf("%s", u.name)
}

func (u *ServiceDirectoryServiceIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-servicedirectory-service-%s", u.GetResourceId())
}

func (u *ServiceDirectoryServiceIamUpdater) DescribeResource() string {
	return fmt.Sprintf("servicedirectory service %q", u.GetResourceId())
}
