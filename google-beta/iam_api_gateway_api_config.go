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

var ApiGatewayApiConfigIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"api": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"api_config": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type ApiGatewayApiConfigIamUpdater struct {
	project   string
	api       string
	apiConfig string
	d         TerraformResourceData
	Config    *transport_tpg.Config
}

func ApiGatewayApiConfigIamUpdaterProducer(d TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	if v, ok := d.GetOk("api"); ok {
		values["api"] = v.(string)
	}

	if v, ok := d.GetOk("api_config"); ok {
		values["api_config"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/global/apis/(?P<api>[^/]+)/configs/(?P<api_config>[^/]+)", "(?P<project>[^/]+)/(?P<api>[^/]+)/(?P<api_config>[^/]+)", "(?P<api>[^/]+)/(?P<api_config>[^/]+)", "(?P<api_config>[^/]+)"}, d, config, d.Get("api_config").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ApiGatewayApiConfigIamUpdater{
		project:   values["project"],
		api:       values["api"],
		apiConfig: values["api_config"],
		d:         d,
		Config:    config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("api", u.api); err != nil {
		return nil, fmt.Errorf("Error setting api: %s", err)
	}
	if err := d.Set("api_config", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting api_config: %s", err)
	}

	return u, nil
}

func ApiGatewayApiConfigIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/global/apis/(?P<api>[^/]+)/configs/(?P<api_config>[^/]+)", "(?P<project>[^/]+)/(?P<api>[^/]+)/(?P<api_config>[^/]+)", "(?P<api>[^/]+)/(?P<api_config>[^/]+)", "(?P<api_config>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ApiGatewayApiConfigIamUpdater{
		project:   values["project"],
		api:       values["api"],
		apiConfig: values["api_config"],
		d:         d,
		Config:    config,
	}
	if err := d.Set("api_config", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting api_config: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *ApiGatewayApiConfigIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyApiConfigUrl("getIamPolicy")
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

	policy, err := SendRequest(u.Config, "GET", project, url, userAgent, obj)
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

func (u *ApiGatewayApiConfigIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyApiConfigUrl("setIamPolicy")
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

	_, err = SendRequestWithTimeout(u.Config, "POST", project, url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *ApiGatewayApiConfigIamUpdater) qualifyApiConfigUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{ApiGatewayBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/global/apis/%s/configs/%s", u.project, u.api, u.apiConfig), methodIdentifier)
	url, err := ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *ApiGatewayApiConfigIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/global/apis/%s/configs/%s", u.project, u.api, u.apiConfig)
}

func (u *ApiGatewayApiConfigIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-apigateway-apiconfig-%s", u.GetResourceId())
}

func (u *ApiGatewayApiConfigIamUpdater) DescribeResource() string {
	return fmt.Sprintf("apigateway apiconfig %q", u.GetResourceId())
}
