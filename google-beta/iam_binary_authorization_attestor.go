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

var BinaryAuthorizationAttestorIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"attestor": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type BinaryAuthorizationAttestorIamUpdater struct {
	project  string
	attestor string
	d        TerraformResourceData
	Config   *transport_tpg.Config
}

func BinaryAuthorizationAttestorIamUpdaterProducer(d TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	if v, ok := d.GetOk("attestor"); ok {
		values["attestor"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/attestors/(?P<attestor>[^/]+)", "(?P<project>[^/]+)/(?P<attestor>[^/]+)", "(?P<attestor>[^/]+)"}, d, config, d.Get("attestor").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &BinaryAuthorizationAttestorIamUpdater{
		project:  values["project"],
		attestor: values["attestor"],
		d:        d,
		Config:   config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("attestor", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting attestor: %s", err)
	}

	return u, nil
}

func BinaryAuthorizationAttestorIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/attestors/(?P<attestor>[^/]+)", "(?P<project>[^/]+)/(?P<attestor>[^/]+)", "(?P<attestor>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &BinaryAuthorizationAttestorIamUpdater{
		project:  values["project"],
		attestor: values["attestor"],
		d:        d,
		Config:   config,
	}
	if err := d.Set("attestor", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting attestor: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *BinaryAuthorizationAttestorIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyAttestorUrl("getIamPolicy")
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

func (u *BinaryAuthorizationAttestorIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyAttestorUrl("setIamPolicy")
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

func (u *BinaryAuthorizationAttestorIamUpdater) qualifyAttestorUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{BinaryAuthorizationBasePath}}%s:%s", fmt.Sprintf("projects/%s/attestors/%s", u.project, u.attestor), methodIdentifier)
	url, err := ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *BinaryAuthorizationAttestorIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/attestors/%s", u.project, u.attestor)
}

func (u *BinaryAuthorizationAttestorIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-binaryauthorization-attestor-%s", u.GetResourceId())
}

func (u *BinaryAuthorizationAttestorIamUpdater) DescribeResource() string {
	return fmt.Sprintf("binaryauthorization attestor %q", u.GetResourceId())
}
