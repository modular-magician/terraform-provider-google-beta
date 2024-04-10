// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

package databasemigrationservice

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceDatabaseMigrationServicePrivateConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabaseMigrationServicePrivateConnectionCreate,
		Read:   resourceDatabaseMigrationServicePrivateConnectionRead,
		Update: resourceDatabaseMigrationServicePrivateConnectionUpdate,
		Delete: resourceDatabaseMigrationServicePrivateConnectionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDatabaseMigrationServicePrivateConnectionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the location this private connection is located in.`,
			},
			"private_connection_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The private connectivity identifier.`,
			},
			"vpc_peering_config": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Description: `The VPC Peering configuration is used to create VPC peering
between databasemigrationservice and the consumer's VPC.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `A free subnet for peering. (CIDR of /29)`,
						},
						"vpc_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `Fully qualified name of the VPC that Database Migration Service will peer to.
Format: projects/{project}/global/{networks}/{name}`,
						},
					},
				},
			},
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `Display name.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Labels.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				ForceNew:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"error": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The PrivateConnection error in case of failure.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"details": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: `A list of messages that carry the error details.`,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"message": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `A message containing more information about the error that occurred.`,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource's name.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `State of the PrivateConnection.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceDatabaseMigrationServicePrivateConnectionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandDatabaseMigrationServicePrivateConnectionDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	vpcPeeringConfigProp, err := expandDatabaseMigrationServicePrivateConnectionVpcPeeringConfig(d.Get("vpc_peering_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vpc_peering_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(vpcPeeringConfigProp)) && (ok || !reflect.DeepEqual(v, vpcPeeringConfigProp)) {
		obj["vpcPeeringConfig"] = vpcPeeringConfigProp
	}
	labelsProp, err := expandDatabaseMigrationServicePrivateConnectionEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DatabaseMigrationServiceBasePath}}projects/{{project}}/locations/{{location}}/privateConnections?privateConnectionId={{private_connection_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new PrivateConnection: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PrivateConnection: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating PrivateConnection: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = DatabaseMigrationServiceOperationWaitTime(
		config, res, project, "Creating PrivateConnection", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create PrivateConnection: %s", err)
	}

	log.Printf("[DEBUG] Finished creating PrivateConnection %q: %#v", d.Id(), res)

	return resourceDatabaseMigrationServicePrivateConnectionRead(d, meta)
}

func resourceDatabaseMigrationServicePrivateConnectionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DatabaseMigrationServiceBasePath}}projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PrivateConnection: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DatabaseMigrationServicePrivateConnection %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading PrivateConnection: %s", err)
	}

	if err := d.Set("name", flattenDatabaseMigrationServicePrivateConnectionName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateConnection: %s", err)
	}
	if err := d.Set("labels", flattenDatabaseMigrationServicePrivateConnectionLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateConnection: %s", err)
	}
	if err := d.Set("display_name", flattenDatabaseMigrationServicePrivateConnectionDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateConnection: %s", err)
	}
	if err := d.Set("state", flattenDatabaseMigrationServicePrivateConnectionState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateConnection: %s", err)
	}
	if err := d.Set("error", flattenDatabaseMigrationServicePrivateConnectionError(res["error"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateConnection: %s", err)
	}
	if err := d.Set("vpc_peering_config", flattenDatabaseMigrationServicePrivateConnectionVpcPeeringConfig(res["vpcPeeringConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateConnection: %s", err)
	}
	if err := d.Set("terraform_labels", flattenDatabaseMigrationServicePrivateConnectionTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateConnection: %s", err)
	}
	if err := d.Set("effective_labels", flattenDatabaseMigrationServicePrivateConnectionEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateConnection: %s", err)
	}

	return nil
}

func resourceDatabaseMigrationServicePrivateConnectionUpdate(d *schema.ResourceData, meta interface{}) error {
	// Only the root field "labels" and "terraform_labels" are mutable
	return resourceDatabaseMigrationServicePrivateConnectionRead(d, meta)
}

func resourceDatabaseMigrationServicePrivateConnectionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PrivateConnection: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DatabaseMigrationServiceBasePath}}projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting PrivateConnection %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "PrivateConnection")
	}

	err = DatabaseMigrationServiceOperationWaitTime(
		config, res, project, "Deleting PrivateConnection", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting PrivateConnection %q: %#v", d.Id(), res)
	return nil
}

func resourceDatabaseMigrationServicePrivateConnectionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/privateConnections/(?P<private_connection_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<private_connection_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<private_connection_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDatabaseMigrationServicePrivateConnectionName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatabaseMigrationServicePrivateConnectionLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenDatabaseMigrationServicePrivateConnectionDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatabaseMigrationServicePrivateConnectionState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatabaseMigrationServicePrivateConnectionError(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["message"] =
		flattenDatabaseMigrationServicePrivateConnectionErrorMessage(original["message"], d, config)
	transformed["details"] =
		flattenDatabaseMigrationServicePrivateConnectionErrorDetails(original["details"], d, config)
	return []interface{}{transformed}
}
func flattenDatabaseMigrationServicePrivateConnectionErrorMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatabaseMigrationServicePrivateConnectionErrorDetails(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatabaseMigrationServicePrivateConnectionVpcPeeringConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["vpc_name"] =
		flattenDatabaseMigrationServicePrivateConnectionVpcPeeringConfigVpcName(original["vpcName"], d, config)
	transformed["subnet"] =
		flattenDatabaseMigrationServicePrivateConnectionVpcPeeringConfigSubnet(original["subnet"], d, config)
	return []interface{}{transformed}
}
func flattenDatabaseMigrationServicePrivateConnectionVpcPeeringConfigVpcName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatabaseMigrationServicePrivateConnectionVpcPeeringConfigSubnet(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatabaseMigrationServicePrivateConnectionTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenDatabaseMigrationServicePrivateConnectionEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDatabaseMigrationServicePrivateConnectionDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServicePrivateConnectionVpcPeeringConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedVpcName, err := expandDatabaseMigrationServicePrivateConnectionVpcPeeringConfigVpcName(original["vpc_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVpcName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vpcName"] = transformedVpcName
	}

	transformedSubnet, err := expandDatabaseMigrationServicePrivateConnectionVpcPeeringConfigSubnet(original["subnet"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubnet); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subnet"] = transformedSubnet
	}

	return transformed, nil
}

func expandDatabaseMigrationServicePrivateConnectionVpcPeeringConfigVpcName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServicePrivateConnectionVpcPeeringConfigSubnet(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServicePrivateConnectionEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
