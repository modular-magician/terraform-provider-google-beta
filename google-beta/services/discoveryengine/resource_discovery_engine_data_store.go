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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceDiscoveryEngineDataStore() *schema.Resource {
	return &schema.Resource{
		Create: resourceDiscoveryEngineDataStoreCreate,
		Read:   resourceDiscoveryEngineDataStoreRead,
		Update: resourceDiscoveryEngineDataStoreUpdate,
		Delete: resourceDiscoveryEngineDataStoreDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDiscoveryEngineDataStoreImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"content_config": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"NO_CONTENT", "CONTENT_REQUIRED", "PUBLIC_WEBSITE"}),
				Description:  `The content config of the data store. Possible values: ["NO_CONTENT", "CONTENT_REQUIRED", "PUBLIC_WEBSITE"]`,
			},
			"data_store_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The unique id of the data store.`,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The display name of the data store. This field must be a UTF-8 encoded
string with a length limit of 128 characters.`,
			},
			"industry_vertical": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"GENERIC", "MEDIA"}),
				Description:  `The industry vertical that the data store registers. Possible values: ["GENERIC", "MEDIA"]`,
			},
			"location": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The geographic location where the data store should reside. The value can
only be one of "global", "us" and "eu".`,
			},
			"create_advanced_site_search": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `If true, an advanced data store for site search will be created. If the
data store is not configured as site search (GENERIC vertical and
PUBLIC_WEBSITE contentConfig), this flag will be ignored.`,
				Default: false,
			},
			"solution_types": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `The solutions that the data store enrolls. Possible values: ["SOLUTION_TYPE_RECOMMENDATION", "SOLUTION_TYPE_SEARCH", "SOLUTION_TYPE_CHAT"]`,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: verify.ValidateEnum([]string{"SOLUTION_TYPE_RECOMMENDATION", "SOLUTION_TYPE_SEARCH", "SOLUTION_TYPE_CHAT"}),
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Timestamp when the DataStore was created.`,
			},
			"default_schema_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The id of the default Schema associated with this data store.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The unique full resource name of the data store. Values are of the format
'projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}'.
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

func resourceDiscoveryEngineDataStoreCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandDiscoveryEngineDataStoreDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	industryVerticalProp, err := expandDiscoveryEngineDataStoreIndustryVertical(d.Get("industry_vertical"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("industry_vertical"); !tpgresource.IsEmptyValue(reflect.ValueOf(industryVerticalProp)) && (ok || !reflect.DeepEqual(v, industryVerticalProp)) {
		obj["industryVertical"] = industryVerticalProp
	}
	solutionTypesProp, err := expandDiscoveryEngineDataStoreSolutionTypes(d.Get("solution_types"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("solution_types"); !tpgresource.IsEmptyValue(reflect.ValueOf(solutionTypesProp)) && (ok || !reflect.DeepEqual(v, solutionTypesProp)) {
		obj["solutionTypes"] = solutionTypesProp
	}
	contentConfigProp, err := expandDiscoveryEngineDataStoreContentConfig(d.Get("content_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("content_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(contentConfigProp)) && (ok || !reflect.DeepEqual(v, contentConfigProp)) {
		obj["contentConfig"] = contentConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores?dataStoreId={{data_store_id}}&createAdvancedSiteSearch={{create_advanced_site_search}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new DataStore: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DataStore: %s", err)
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
		return fmt.Errorf("Error creating DataStore: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = DiscoveryEngineOperationWaitTime(
		config, res, project, "Creating DataStore", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create DataStore: %s", err)
	}

	log.Printf("[DEBUG] Finished creating DataStore %q: %#v", d.Id(), res)

	return resourceDiscoveryEngineDataStoreRead(d, meta)
}

func resourceDiscoveryEngineDataStoreRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DataStore: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DiscoveryEngineDataStore %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading DataStore: %s", err)
	}

	if err := d.Set("name", flattenDiscoveryEngineDataStoreName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataStore: %s", err)
	}
	if err := d.Set("display_name", flattenDiscoveryEngineDataStoreDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataStore: %s", err)
	}
	if err := d.Set("industry_vertical", flattenDiscoveryEngineDataStoreIndustryVertical(res["industryVertical"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataStore: %s", err)
	}
	if err := d.Set("solution_types", flattenDiscoveryEngineDataStoreSolutionTypes(res["solutionTypes"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataStore: %s", err)
	}
	if err := d.Set("default_schema_id", flattenDiscoveryEngineDataStoreDefaultSchemaId(res["defaultSchemaId"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataStore: %s", err)
	}
	if err := d.Set("content_config", flattenDiscoveryEngineDataStoreContentConfig(res["contentConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataStore: %s", err)
	}
	if err := d.Set("create_time", flattenDiscoveryEngineDataStoreCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataStore: %s", err)
	}

	return nil
}

func resourceDiscoveryEngineDataStoreUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DataStore: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandDiscoveryEngineDataStoreDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating DataStore %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating DataStore %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating DataStore %q: %#v", d.Id(), res)
		}

		err = DiscoveryEngineOperationWaitTime(
			config, res, project, "Updating DataStore", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceDiscoveryEngineDataStoreRead(d, meta)
}

func resourceDiscoveryEngineDataStoreDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DataStore: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting DataStore %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "DataStore")
	}

	err = DiscoveryEngineOperationWaitTime(
		config, res, project, "Deleting DataStore", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting DataStore %q: %#v", d.Id(), res)
	return nil
}

func resourceDiscoveryEngineDataStoreImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/collections/default_collection/dataStores/(?P<data_store_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<data_store_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<data_store_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDiscoveryEngineDataStoreName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineDataStoreDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineDataStoreIndustryVertical(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineDataStoreSolutionTypes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineDataStoreDefaultSchemaId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineDataStoreContentConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineDataStoreCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDiscoveryEngineDataStoreDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineDataStoreIndustryVertical(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineDataStoreSolutionTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineDataStoreContentConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
