// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceVmwareenginePrivateCloud() *schema.Resource {
	return &schema.Resource{
		Create: resourceVmwareenginePrivateCloudCreate,
		Read:   resourceVmwareenginePrivateCloudRead,
		Update: resourceVmwareenginePrivateCloudUpdate,
		Delete: resourceVmwareenginePrivateCloudDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVmwareenginePrivateCloudImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(180 * time.Minute),
			Update: schema.DefaultTimeout(160 * time.Minute),
			Delete: schema.DefaultTimeout(120 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The location where the PrivateCloud should reside.`,
			},
			"management_cluster": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: false,
				Description: `The management cluster for this private cloud. This field is required during creation of the private cloud to provide details for the default cluster.
The following fields can't be changed after private cloud creation: ManagementCluster.clusterId, ManagementCluster.nodeTypeId.`,
				MaxItems: 1,
				Elem:     getManagementClusterSchema(),
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the PrivateCloud.`,
			},
			"network_config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Network configuration in the consumer project with which the peering has to be done.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"management_cidr": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Management CIDR used by VMware management appliances.`,
						},
						"vmware_engine_network": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The relative resource name of the VMware Engine network attached to the private cloud.
Specify the name in the following form: projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId}
where {project} can either be a project number or a project ID.`,
						},
						"management_ip_address_layout_version": {
							Type:     schema.TypeInt,
							Computed: true,
							Description: `The IP address layout version of the management IP address range.
Possible versions include: * managementIpAddressLayoutVersion=1: Indicates the legacy
IP address layout used by some existing private clouds. This is no longer supported for new private clouds
as it does not support all features. * managementIpAddressLayoutVersion=2: Indicates the latest IP address layout
used by all newly created private clouds. This version supports all current features.`,
						},
						"vmware_engine_network_canonical": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The canonical name of the VMware Engine network in
the form: projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId}`,
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `User-provided description for this private cloud.`,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Creation time of this resource.

A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"delete_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Time when the resource was scheduled for deletion.

A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"expire_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Time when the resource will be irreversibly deleted.

A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"hcx": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Details about a HCX Cloud Manager appliance.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fqdn": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Fully qualified domain name of the appliance.`,
						},
						"internal_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Internal IP address of the appliance.`,
						},
						"state": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateEnum([]string{"ACTIVE", "CREATING", ""}),
							Description:  `State of the appliance. Possible values: ["ACTIVE", "CREATING"]`,
						},
						"version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Version of the appliance.`,
						},
					},
				},
			},
			"nsx": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Details about a NSX Manager appliance.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fqdn": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Fully qualified domain name of the appliance.`,
						},
						"internal_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Internal IP address of the appliance.`,
						},
						"state": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateEnum([]string{"ACTIVE", "CREATING", ""}),
							Description:  `State of the appliance. Possible values: ["ACTIVE", "CREATING"]`,
						},
						"version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Version of the appliance.`,
						},
					},
				},
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `State of the resource. New values may be added to this enum when appropriate.`,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `System-generated unique identifier for the resource.`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Last update time of this resource.

A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"vcenter": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Details about a vCenter Server management appliance.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fqdn": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Fully qualified domain name of the appliance.`,
						},
						"internal_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Internal IP address of the appliance.`,
						},
						"state": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateEnum([]string{"ACTIVE", "CREATING", ""}),
							Description:  `State of the appliance. Possible values: ["ACTIVE", "CREATING"]`,
						},
						"version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Version of the appliance.`,
						},
					},
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceVmwareenginePrivateCloudCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandVmwareenginePrivateCloudDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networkConfigProp, err := expandVmwareenginePrivateCloudNetworkConfig(d.Get("network_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_config"); !isEmptyValue(reflect.ValueOf(networkConfigProp)) && (ok || !reflect.DeepEqual(v, networkConfigProp)) {
		obj["networkConfig"] = networkConfigProp
	}
	managementClusterProp, err := expandVmwareenginePrivateCloudManagementCluster(d.Get("management_cluster"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("management_cluster"); !isEmptyValue(reflect.ValueOf(managementClusterProp)) && (ok || !reflect.DeepEqual(v, managementClusterProp)) {
		obj["managementCluster"] = managementClusterProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VmwareengineBasePath}}projects/{{project}}/locations/{{location}}/privateClouds?privateCloudId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new PrivateCloud: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PrivateCloud: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})

	if err != nil {
		return fmt.Errorf("Error creating PrivateCloud: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/privateClouds/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = VmwareengineOperationWaitTime(
		config, res, project, "Creating PrivateCloud", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create PrivateCloud: %s", err)
	}

	log.Printf("[DEBUG] Finished creating PrivateCloud %q: %#v", d.Id(), res)

	return resourceVmwareenginePrivateCloudRead(d, meta)
}

func resourceVmwareenginePrivateCloudRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VmwareengineBasePath}}projects/{{project}}/locations/{{location}}/privateClouds/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PrivateCloud: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})

	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("VmwareenginePrivateCloud %q", d.Id()))
	}

	// We are only interested in management cluster of the PrivateCloud. It should always exist
	// `management` property will be true only for ManagementCluster and there can only be one.
	mgmtCluster, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url + "/clusters?filter=management=true",
		UserAgent: userAgent,
	})

	if err != nil {
		return fmt.Errorf("Error reading management cluster of PrivateCloud: %s", err)
	}

	// There can only be 1 management cluster and if the PC read is successfuly and
	// we got response from cluster API then it should be present.
	mgmtClusterObj := mgmtCluster["clusters"].([]interface{})[0]

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading PrivateCloud: %s", err)
	}

	if err := d.Set("description", flattenVmwareenginePrivateCloudDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateCloud: %s", err)
	}
	if err := d.Set("create_time", flattenVmwareenginePrivateCloudCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateCloud: %s", err)
	}
	if err := d.Set("update_time", flattenVmwareenginePrivateCloudUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateCloud: %s", err)
	}
	if err := d.Set("delete_time", flattenVmwareenginePrivateCloudDeleteTime(res["deleteTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateCloud: %s", err)
	}
	if err := d.Set("expire_time", flattenVmwareenginePrivateCloudExpireTime(res["expireTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateCloud: %s", err)
	}
	if err := d.Set("uid", flattenVmwareenginePrivateCloudUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateCloud: %s", err)
	}
	if err := d.Set("state", flattenVmwareenginePrivateCloudState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateCloud: %s", err)
	}
	if err := d.Set("network_config", flattenVmwareenginePrivateCloudNetworkConfig(res["networkConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateCloud: %s", err)
	}
	if err := d.Set("management_cluster", flattenVmwareenginePrivateCloudManagementCluster(mgmtClusterObj, d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateCloud: %s", err)
	}
	if err := d.Set("hcx", flattenVmwareenginePrivateCloudHcx(res["hcx"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateCloud: %s", err)
	}
	if err := d.Set("nsx", flattenVmwareenginePrivateCloudNsx(res["nsx"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateCloud: %s", err)
	}
	if err := d.Set("vcenter", flattenVmwareenginePrivateCloudVcenter(res["vcenter"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateCloud: %s", err)
	}

	return nil
}

func resourceVmwareenginePrivateCloudUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PrivateCloud: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandVmwareenginePrivateCloudDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networkConfigProp, err := expandVmwareenginePrivateCloudNetworkConfig(d.Get("network_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, networkConfigProp)) {
		obj["networkConfig"] = networkConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VmwareengineBasePath}}projects/{{project}}/locations/{{location}}/privateClouds/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating PrivateCloud %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		log.Printf("[DEBUG] description changed %q", d.Id())
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("network_config") {
		log.Printf("[DEBUG] network_config changed %q", d.Id())
		updateMask = append(updateMask, "networkConfig")
	}

	if d.HasChange("management_cluster[0].node_type_configs") {
		log.Printf("[DEBUG] node_type_config changed %q", d.Id())
	}

	if d.HasChange("management_cluster") {
		log.Printf("[DEBUG] management_cluster changed %q", d.Id())
	}

	// this function will be triggered when either PC needs core attributes
	// needs to be updated or management cluster
	// update is required. Checking this avoids PC API call when only management cluster
	// is updated.
	if len(updateMask) > 0 {
		// updateMask is a URL parameter but not present in the schema, so tpgresource.ReplaceVars
		// won't set it
		// PS: don't update the url directly, its used for deriving cluster URL too and updating it will have impact.
		patchUrl, err := transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    patchUrl,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})

		if err != nil {
			return fmt.Errorf("Error updating PrivateCloud %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating PrivateCloud %q: %#v", d.Id(), res)
		}

		err = VmwareengineOperationWaitTime(
			config, res, project, "Updating PrivateCloud", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	// updates management cluster if required.
	err = updateManagementCluster(d, config, url, billingProject, userAgent, project)
	if err != nil {
		return err
	}

	return resourceVmwareenginePrivateCloudRead(d, meta)
}

func updateManagementCluster(d *schema.ResourceData, config *transport_tpg.Config, parentUrl string, billingProject string, userAgent string, project string) error {
	clusterObj := make(map[string]interface{})
	managementClusterProp, err := expandVmwareenginePrivateCloudManagementCluster(d.Get("management_cluster"), d, config)
	mgmtMap := managementClusterProp.(map[string]interface{})
	clusterUrl := fmt.Sprintf("%s/clusters/%s", parentUrl, mgmtMap["clusterId"].(string))
	clusterUpdateMask := []string{}
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("management_cluster"); !isEmptyValue(reflect.ValueOf(managementClusterProp)) && (ok || !reflect.DeepEqual(v, managementClusterProp)) {
		clusterObj["nodeTypeConfigs"] = mgmtMap["nodeTypeConfigs"]
	}

	if d.HasChange("management_cluster") {
		log.Printf("[DEBUG] managment_cluster changed %q", d.Id())
		log.Printf("[DEBUG] clusterURL %s", clusterUrl)
		clusterUpdateMask = append(clusterUpdateMask, "nodeTypeConfigs.*.nodeCount")
	}

	patchUrl, err := transport_tpg.AddQueryParams(clusterUrl, map[string]string{"updateMask": strings.Join(clusterUpdateMask, ",")})
	if err != nil {
		return err
	}

	// nothing to update
	if len(clusterUpdateMask) == 0 {
		return nil
	}
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    patchUrl,
		UserAgent: userAgent,
		Body:      clusterObj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating magament cluster %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating magament cluster %q: %#v", d.Id(), res)
	}

	err = VmwareengineOperationWaitTime(
		config, res, project, "Updating Managment Cluster", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return nil
}

func resourceVmwareenginePrivateCloudDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PrivateCloud: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{VmwareengineBasePath}}projects/{{project}}/locations/{{location}}/privateClouds/{{name}}?delay_hours=0")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting PrivateCloud %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})

	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "PrivateCloud")
	}

	err = VmwareengineOperationWaitTime(
		config, res, project, "Deleting PrivateCloud", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}
	privateCloudPollRead := func(d *schema.ResourceData, meta interface{}) transport_tpg.PollReadFunc {
		return func() (map[string]interface{}, error) {
			config := meta.(*transport_tpg.Config)
			url, err := tpgresource.ReplaceVars(d, config, "{{VmwareengineBasePath}}projects/{{project}}/locations/{{location}}/privateClouds/{{name}}")
			if err != nil {
				return nil, err
			}
			billingProject := ""
			project, err := getProject(d, config)
			if err != nil {
				return nil, fmt.Errorf("Error fetching project for PrivateCloud: %s", err)
			}
			billingProject = project
			// err == nil indicates that the billing_project value was found
			if bp, err := getBillingProject(d, config); err == nil {
				billingProject = bp
			}
			userAgent, err := generateUserAgentString(d, config.UserAgent)
			if err != nil {
				return nil, err
			}
			res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: userAgent,
			})
			if err != nil {
				return res, err
			}
			return res, nil
		}
	}

	err = PollingWaitTime(privateCloudPollRead(d, meta), PollCheckForAbsence, "Deleting PrivateCloud", d.Timeout(schema.TimeoutDelete), 10)
	if err != nil {
		return fmt.Errorf("Error waiting to delete PrivateCloud: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting PrivateCloud %q: %#v", d.Id(), res)
	return nil
}

func resourceVmwareenginePrivateCloudImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/privateClouds/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/privateClouds/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenVmwareenginePrivateCloudDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudDeleteTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudExpireTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudNetworkConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["management_cidr"] =
		flattenVmwareenginePrivateCloudNetworkConfigManagementCidr(original["managementCidr"], d, config)
	transformed["vmware_engine_network"] =
		flattenVmwareenginePrivateCloudNetworkConfigVmwareEngineNetwork(original["vmwareEngineNetwork"], d, config)
	transformed["vmware_engine_network_canonical"] =
		flattenVmwareenginePrivateCloudNetworkConfigVmwareEngineNetworkCanonical(original["vmwareEngineNetworkCanonical"], d, config)
	transformed["management_ip_address_layout_version"] =
		flattenVmwareenginePrivateCloudNetworkConfigManagementIpAddressLayoutVersion(original["managementIpAddressLayoutVersion"], d, config)
	return []interface{}{transformed}
}

func flattenVmwareenginePrivateCloudNetworkConfigManagementCidr(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudNetworkConfigVmwareEngineNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudNetworkConfigVmwareEngineNetworkCanonical(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudNetworkConfigManagementIpAddressLayoutVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenVmwareenginePrivateCloudManagementCluster(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["cluster_id"] =
		flattenVmwareenginePrivateCloudManagementClusterClusterId(original["name"], d, config)
	transformed["node_type_configs"] =
		flattenVmwareenginePrivateCloudManagementClusterNodeTypeConfigs(original["nodeTypeConfigs"], d, config)
	transformed["create_time"] = flattenVmwareengineClusterCreateTime(original["createTime"], d, config)

	transformed["update_time"] = flattenVmwareengineClusterUpdateTime(original["updateTime"], d, config)
	transformed["management"] = flattenVmwareengineClusterManagement(original["management"], d, config)
	transformed["uid"] = flattenVmwareengineClusterUid(original["uid"], d, config)
	transformed["state"] = flattenVmwareengineClusterState(original["state"], d, config)

	return []interface{}{transformed}
}
func flattenVmwareenginePrivateCloudManagementClusterClusterId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	oldvalue := v.(string)
	return oldvalue[strings.LastIndex(oldvalue, "/")+1:]
}

func flattenVmwareenginePrivateCloudManagementClusterNodeTypeConfigs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return flattenVmwareengineClusterNodeTypeConfigs(v, d, config)
}

func flattenVmwareenginePrivateCloudHcx(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["internal_ip"] =
		flattenVmwareenginePrivateCloudHcxInternalIp(original["internalIp"], d, config)
	transformed["version"] =
		flattenVmwareenginePrivateCloudHcxVersion(original["version"], d, config)
	transformed["state"] =
		flattenVmwareenginePrivateCloudHcxState(original["state"], d, config)
	transformed["fqdn"] =
		flattenVmwareenginePrivateCloudHcxFqdn(original["fqdn"], d, config)
	return []interface{}{transformed}
}

func flattenVmwareenginePrivateCloudHcxInternalIp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudHcxVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudHcxState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudHcxFqdn(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudNsx(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["internal_ip"] =
		flattenVmwareenginePrivateCloudNsxInternalIp(original["internalIp"], d, config)
	transformed["version"] =
		flattenVmwareenginePrivateCloudNsxVersion(original["version"], d, config)
	transformed["state"] =
		flattenVmwareenginePrivateCloudNsxState(original["state"], d, config)
	transformed["fqdn"] =
		flattenVmwareenginePrivateCloudNsxFqdn(original["fqdn"], d, config)
	return []interface{}{transformed}
}
func flattenVmwareenginePrivateCloudNsxInternalIp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudNsxVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudNsxState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudNsxFqdn(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudVcenter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["internal_ip"] =
		flattenVmwareenginePrivateCloudVcenterInternalIp(original["internalIp"], d, config)
	transformed["version"] =
		flattenVmwareenginePrivateCloudVcenterVersion(original["version"], d, config)
	transformed["state"] =
		flattenVmwareenginePrivateCloudVcenterState(original["state"], d, config)
	transformed["fqdn"] =
		flattenVmwareenginePrivateCloudVcenterFqdn(original["fqdn"], d, config)
	return []interface{}{transformed}
}
func flattenVmwareenginePrivateCloudVcenterInternalIp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudVcenterVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudVcenterState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareenginePrivateCloudVcenterFqdn(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandVmwareenginePrivateCloudDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareenginePrivateCloudNetworkConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedManagementCidr, err := expandVmwareenginePrivateCloudNetworkConfigManagementCidr(original["management_cidr"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedManagementCidr); val.IsValid() && !isEmptyValue(val) {
		transformed["managementCidr"] = transformedManagementCidr
	}

	transformedVmwareEngineNetwork, err := expandVmwareenginePrivateCloudNetworkConfigVmwareEngineNetwork(original["vmware_engine_network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVmwareEngineNetwork); val.IsValid() && !isEmptyValue(val) {
		transformed["vmwareEngineNetwork"] = transformedVmwareEngineNetwork
	}

	transformedVmwareEngineNetworkCanonical, err := expandVmwareenginePrivateCloudNetworkConfigVmwareEngineNetworkCanonical(original["vmware_engine_network_canonical"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVmwareEngineNetworkCanonical); val.IsValid() && !isEmptyValue(val) {
		transformed["vmwareEngineNetworkCanonical"] = transformedVmwareEngineNetworkCanonical
	}

	transformedManagementIpAddressLayoutVersion, err := expandVmwareenginePrivateCloudNetworkConfigManagementIpAddressLayoutVersion(original["management_ip_address_layout_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedManagementIpAddressLayoutVersion); val.IsValid() && !isEmptyValue(val) {
		transformed["managementIpAddressLayoutVersion"] = transformedManagementIpAddressLayoutVersion
	}

	return transformed, nil
}

func expandVmwareenginePrivateCloudNetworkConfigManagementCidr(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareenginePrivateCloudNetworkConfigVmwareEngineNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareenginePrivateCloudNetworkConfigVmwareEngineNetworkCanonical(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareenginePrivateCloudNetworkConfigManagementIpAddressLayoutVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareenginePrivateCloudManagementCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedClusterId, err := expandVmwareenginePrivateCloudManagementClusterClusterId(original["cluster_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClusterId); val.IsValid() && !isEmptyValue(val) {
		transformed["clusterId"] = transformedClusterId
	}

	transformedNodeTypeConfigs, err := expandVmwareenginePrivateCloudManagementClusterNodeTypeConfigs(original["node_type_configs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNodeTypeConfigs); val.IsValid() && !isEmptyValue(val) {
		transformed["nodeTypeConfigs"] = transformedNodeTypeConfigs
	}

	return transformed, nil
}

func expandVmwareenginePrivateCloudManagementClusterClusterId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareenginePrivateCloudManagementClusterNodeTypeConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	return expandVmwareengineClusterNodeTypeConfigs(v, d, config)
}

func getManagementClusterSchema() *schema.Resource {
	clusterResource := ResourceVmwareengineCluster()

	// renaming `name` to cluster_id
	if _, ok := clusterResource.Schema["name"]; ok {
		clusterResource.Schema["cluster_id"] = clusterResource.Schema["name"]
		delete(clusterResource.Schema, "name")
	}

	// remove parent as its not required for management cluster.
	if _, ok := clusterResource.Schema["parent"]; ok {
		delete(clusterResource.Schema, "parent")
	}

	return clusterResource
}
