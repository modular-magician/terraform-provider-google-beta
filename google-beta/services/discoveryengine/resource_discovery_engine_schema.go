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

package discoveryengine

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceDiscoveryEngineSchema() *schema.Resource {
	return &schema.Resource{
		Create: resourceDiscoveryEngineSchemaCreate,
		Read:   resourceDiscoveryEngineSchemaRead,
		Delete: resourceDiscoveryEngineSchemaDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDiscoveryEngineSchemaImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"data_store_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The unique id of the data store.`,
			},
			"location": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The geographic location where the data store should reside. The value can
only be one of "global", "us" and "eu".`,
			},
			"schema_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The unique id of the schema.`,
			},
			"json_schema": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsJSON,
				StateFunc:    func(v interface{}) string { s, _ := structure.NormalizeJsonString(v); return s },
				Description:  `The JSON representation of the schema.`,
				ExactlyOneOf: []string{"json_schema"},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The unique full resource name of the schema. Values are of the format
'projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}/schemas/{schema_id}'.
This field must be a UTF-8 encoded string with a length limit of 1024
characters.`,
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

func resourceDiscoveryEngineSchemaCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	jsonSchemaProp, err := expandDiscoveryEngineSchemaJsonSchema(d.Get("json_schema"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("json_schema"); !tpgresource.IsEmptyValue(reflect.ValueOf(jsonSchemaProp)) && (ok || !reflect.DeepEqual(v, jsonSchemaProp)) {
		obj["jsonSchema"] = jsonSchemaProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}/schemas?schemaId={{schema_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Schema: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Schema: %s", err)
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
		return fmt.Errorf("Error creating Schema: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}/schemas/{{schema_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = DiscoveryEngineOperationWaitTime(
		config, res, project, "Creating Schema", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Schema: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Schema %q: %#v", d.Id(), res)

	return resourceDiscoveryEngineSchemaRead(d, meta)
}

func resourceDiscoveryEngineSchemaRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}/schemas/{{schema_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Schema: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DiscoveryEngineSchema %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Schema: %s", err)
	}

	if err := d.Set("name", flattenDiscoveryEngineSchemaName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Schema: %s", err)
	}
	if err := d.Set("json_schema", flattenDiscoveryEngineSchemaJsonSchema(res["jsonSchema"], d, config)); err != nil {
		return fmt.Errorf("Error reading Schema: %s", err)
	}

	return nil
}

func resourceDiscoveryEngineSchemaDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Schema: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}/schemas/{{schema_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Schema %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "Schema")
	}

	err = DiscoveryEngineOperationWaitTime(
		config, res, project, "Deleting Schema", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Schema %q: %#v", d.Id(), res)
	return nil
}

func resourceDiscoveryEngineSchemaImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/collections/default_collection/dataStores/(?P<data_store_id>[^/]+)/schemas/(?P<schema_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<data_store_id>[^/]+)/(?P<schema_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<data_store_id>[^/]+)/(?P<schema_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}/schemas/{{schema_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDiscoveryEngineSchemaName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineSchemaJsonSchema(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	s, err := structure.NormalizeJsonString(v)
	if err != nil {
		log.Printf("[ERROR] failed to normalize JSON string: %v", err)
	}
	return s
}

func expandDiscoveryEngineSchemaJsonSchema(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
