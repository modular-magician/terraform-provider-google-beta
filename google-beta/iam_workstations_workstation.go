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

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var WorkstationsWorkstationIamSchema = map[string]*schema.Schema{
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
	"workstation_cluster_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"workstation_config_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"workstation_id": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type WorkstationsWorkstationIamUpdater struct {
	project              string
	location             string
	workstationClusterId string
	workstationConfigId  string
	workstationId        string
	d                    tpgresource.TerraformResourceData
	Config               *transport_tpg.Config
}

func WorkstationsWorkstationIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
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
	if v, ok := d.GetOk("workstation_cluster_id"); ok {
		values["workstation_cluster_id"] = v.(string)
	}

	if v, ok := d.GetOk("workstation_config_id"); ok {
		values["workstation_config_id"] = v.(string)
	}

	if v, ok := d.GetOk("workstation_id"); ok {
		values["workstation_id"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/workstationClusters/(?P<workstation_cluster_id>[^/]+)/workstationConfigs/(?P<workstation_config_id>[^/]+)/workstations/(?P<workstation_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<workstation_cluster_id>[^/]+)/(?P<workstation_config_id>[^/]+)/(?P<workstation_id>[^/]+)", "(?P<location>[^/]+)/(?P<workstation_cluster_id>[^/]+)/(?P<workstation_config_id>[^/]+)/(?P<workstation_id>[^/]+)", "(?P<workstation_id>[^/]+)"}, d, config, d.Get("workstation_id").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &WorkstationsWorkstationIamUpdater{
		project:              values["project"],
		location:             values["location"],
		workstationClusterId: values["workstation_cluster_id"],
		workstationConfigId:  values["workstation_config_id"],
		workstationId:        values["workstation_id"],
		d:                    d,
		Config:               config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("location", u.location); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("workstation_cluster_id", u.workstationClusterId); err != nil {
		return nil, fmt.Errorf("Error setting workstation_cluster_id: %s", err)
	}
	if err := d.Set("workstation_config_id", u.workstationConfigId); err != nil {
		return nil, fmt.Errorf("Error setting workstation_config_id: %s", err)
	}
	if err := d.Set("workstation_id", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting workstation_id: %s", err)
	}

	return u, nil
}

func WorkstationsWorkstationIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	location, _ := getLocation(d, config)
	if location != "" {
		values["location"] = location
	}

	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/workstationClusters/(?P<workstation_cluster_id>[^/]+)/workstationConfigs/(?P<workstation_config_id>[^/]+)/workstations/(?P<workstation_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<workstation_cluster_id>[^/]+)/(?P<workstation_config_id>[^/]+)/(?P<workstation_id>[^/]+)", "(?P<location>[^/]+)/(?P<workstation_cluster_id>[^/]+)/(?P<workstation_config_id>[^/]+)/(?P<workstation_id>[^/]+)", "(?P<workstation_id>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &WorkstationsWorkstationIamUpdater{
		project:              values["project"],
		location:             values["location"],
		workstationClusterId: values["workstation_cluster_id"],
		workstationConfigId:  values["workstation_config_id"],
		workstationId:        values["workstation_id"],
		d:                    d,
		Config:               config,
	}
	if err := d.Set("workstation_id", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting workstation_id: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *WorkstationsWorkstationIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyWorkstationUrl("getIamPolicy")
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

	policy, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "GET",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
	})
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

func (u *WorkstationsWorkstationIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyWorkstationUrl("setIamPolicy")
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

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   u.d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *WorkstationsWorkstationIamUpdater) qualifyWorkstationUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{WorkstationsBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s/workstations/%s", u.project, u.location, u.workstationClusterId, u.workstationConfigId, u.workstationId), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *WorkstationsWorkstationIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s/workstations/%s", u.project, u.location, u.workstationClusterId, u.workstationConfigId, u.workstationId)
}

func (u *WorkstationsWorkstationIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-workstations-workstation-%s", u.GetResourceId())
}

func (u *WorkstationsWorkstationIamUpdater) DescribeResource() string {
	return fmt.Sprintf("workstations workstation %q", u.GetResourceId())
}
