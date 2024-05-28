// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package composer

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
	"log"
	"time"

	composer "google.golang.org/api/composer/v1beta1"
)

func ResourceComposerUserWorkloadsConfigMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceComposerUserWorkloadsConfigMapCreate,
		Read:   resourceComposerUserWorkloadsConfigMapRead,
		Update: resourceComposerUserWorkloadsConfigMapUpdate,
		Delete: resourceComposerUserWorkloadsConfigMapDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComposerUserWorkloadsConfigMapImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(time.Minute),
			Update: schema.DefaultTimeout(time.Minute),
			Delete: schema.DefaultTimeout(time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
			tpgresource.DefaultProviderRegion,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateGCEName,
				Description:  `Name of the config map.`,
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The location or Compute Engine region for the environment.`,
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
			},
			"environment": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateGCEName,
				Description:  `Name of the environment.`,
			},
			"data": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    false,
				Description: `The "data" field of Kubernetes ConfigMap, organized in key-value pairs.`,
			},
		},
	}
}

func resourceComposerUserWorkloadsConfigMapCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	confiMapName, err := resourceComposerUserWorkloadsConfigMapName(d, config)
	if err != nil {
		return err
	}

	configMap := &composer.UserWorkloadsConfigMap{
		Name: confiMapName.ResourceName(),
		Data: tpgresource.ConvertStringMap(d.Get("data").(map[string]interface{})),
	}

	log.Printf("[DEBUG] Creating new UserWorkloadsConfigMap %q", confiMapName.ParentName())
	resp, err := config.NewComposerClient(userAgent).Projects.Locations.Environments.UserWorkloadsConfigMaps.Create(confiMapName.ParentName(), configMap).Do()
	if err != nil {
		return fmt.Errorf("Error creating UserWorkloadsConfigMap: %s", err)
	}

	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/environments/{{environment}}/userWorkloadsConfigMaps/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	respJson, _ := resp.MarshalJSON()
	log.Printf("[DEBUG] Finished creating UserWorkloadsConfigMap %q: %#v", d.Id(), string(respJson))

	return resourceComposerUserWorkloadsConfigMapRead(d, meta)
}

func resourceComposerUserWorkloadsConfigMapRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	confiMapName, err := resourceComposerUserWorkloadsConfigMapName(d, config)
	if err != nil {
		return err
	}

	res, err := config.NewComposerClient(userAgent).Projects.Locations.Environments.UserWorkloadsConfigMaps.Get(confiMapName.ResourceName()).Do()
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("UserWorkloadsConfigMap %q", d.Id()))
	}

	if err := d.Set("project", confiMapName.Project); err != nil {
		return fmt.Errorf("Error setting UserWorkloadsConfigMap Project: %s", err)
	}
	if err := d.Set("region", confiMapName.Region); err != nil {
		return fmt.Errorf("Error setting UserWorkloadsConfigMap Region: %s", err)
	}
	if err := d.Set("environment", confiMapName.Environment); err != nil {
		return fmt.Errorf("Error setting UserWorkloadsConfigMap Environment: %s", err)
	}
	if err := d.Set("name", tpgresource.GetResourceNameFromSelfLink(res.Name)); err != nil {
		return fmt.Errorf("Error setting UserWorkloadsConfigMap Name: %s", err)
	}
	if err := d.Set("data", res.Data); err != nil {
		return fmt.Errorf("Error setting UserWorkloadsConfigMap Data: %s", err)
	}
	return nil
}

func resourceComposerUserWorkloadsConfigMapUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	confiMapName, err := resourceComposerUserWorkloadsConfigMapName(d, config)
	if err != nil {
		return err
	}

	if d.HasChange("data") {
		configMap := &composer.UserWorkloadsConfigMap{
			Name: confiMapName.ResourceName(),
			Data: tpgresource.ConvertStringMap(d.Get("data").(map[string]interface{})),
		}

		configMapJson, _ := configMap.MarshalJSON()
		log.Printf("[DEBUG] Updating UserWorkloadsConfigMap %q: %s", d.Id(), string(configMapJson))

		resp, err := config.NewComposerClient(userAgent).Projects.Locations.Environments.UserWorkloadsConfigMaps.Update(confiMapName.ResourceName(), configMap).Do()
		if err != nil {
			return err
		}

		respJson, _ := resp.MarshalJSON()
		log.Printf("[DEBUG] Finished updating UserWorkloadsConfigMap %q: %s", d.Id(), string(respJson))
	}

	return nil
}

func resourceComposerUserWorkloadsConfigMapDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	confiMapName, err := resourceComposerUserWorkloadsConfigMapName(d, config)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting UserWorkloadsConfigMap %q", d.Id())
	_, err = config.NewComposerClient(userAgent).Projects.Locations.Environments.UserWorkloadsConfigMaps.Delete(confiMapName.ResourceName()).Do()
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Finished deleting UserWorkloadsConfigMap %q", d.Id())

	return nil
}

func resourceComposerUserWorkloadsConfigMapImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/environments/(?P<environment>[^/]+)/userWorkloadsConfigMaps/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<environment>[^/]+)/(?P<name>[^/]+)", "(?P<environment>[^/]+)/(?P<name>[^/]+)"}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/environments/{{environment}}/userWorkloadsConfigMaps/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func resourceComposerUserWorkloadsConfigMapName(d *schema.ResourceData, config *transport_tpg.Config) (*UserWorkloadsConfigMapName, error) {
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return nil, err
	}

	region, err := tpgresource.GetRegion(d, config)
	if err != nil {
		return nil, err
	}

	return &UserWorkloadsConfigMapName{
		Project:     project,
		Region:      region,
		Environment: d.Get("environment").(string),
		ConfigMap:   d.Get("name").(string),
	}, nil
}

type UserWorkloadsConfigMapName struct {
	Project     string
	Region      string
	Environment string
	ConfigMap   string
}

func (n *UserWorkloadsConfigMapName) ResourceName() string {
	return fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsConfigMaps/%s", n.Project, n.Region, n.Environment, n.ConfigMap)
}

func (n *UserWorkloadsConfigMapName) ParentName() string {
	return fmt.Sprintf("projects/%s/locations/%s/environments/%s", n.Project, n.Region, n.Environment)
}
