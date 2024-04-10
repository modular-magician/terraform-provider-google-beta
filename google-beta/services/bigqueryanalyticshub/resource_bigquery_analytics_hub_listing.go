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

package bigqueryanalyticshub

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
)

func ResourceBigqueryAnalyticsHubListing() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigqueryAnalyticsHubListingCreate,
		Read:   resourceBigqueryAnalyticsHubListingRead,
		Update: resourceBigqueryAnalyticsHubListingUpdate,
		Delete: resourceBigqueryAnalyticsHubListingDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBigqueryAnalyticsHubListingImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"bigquery_dataset": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Shared dataset i.e. BigQuery dataset source.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dataset": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tpgresource.ProjectNumberDiffSuppress,
							Description:      `Resource name of the dataset source for this listing. e.g. projects/myproject/datasets/123`,
						},
					},
				},
			},
			"data_exchange_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and can't start or end with spaces.`,
			},
			"listing_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the location this data exchange listing.`,
			},
			"categories": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Categories of the listing. Up to two categories are allowed.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"data_provider": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Details of the data provider who owns the source data.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Name of the data provider.`,
						},
						"primary_contact": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Email or URL of the data provider.`,
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).`,
			},
			"documentation": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Documentation describing the listing.`,
			},
			"icon": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Base64 encoded image representing the listing.`,
			},
			"primary_contact": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Email or URL of the primary point of contact of the listing.`,
			},
			"publisher": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Details of the publisher who owns the listing and who can share the source data.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Name of the listing publisher.`,
						},
						"primary_contact": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Email or URL of the listing publisher.`,
						},
					},
				},
			},
			"request_access": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Email or URL of the request access of the listing. Subscribers can use this reference to request access.`,
			},
			"restricted_export_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `If set, restricted export configuration will be propagated and enforced on the linked dataset.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `If true, enable restricted export.`,
						},
						"restrict_query_result": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `If true, restrict export of query result derived from restricted linked dataset table.`,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the listing. e.g. "projects/myproject/locations/US/dataExchanges/123/listings/456"`,
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

func resourceBigqueryAnalyticsHubListingCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandBigqueryAnalyticsHubListingDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandBigqueryAnalyticsHubListingDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	primaryContactProp, err := expandBigqueryAnalyticsHubListingPrimaryContact(d.Get("primary_contact"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("primary_contact"); !tpgresource.IsEmptyValue(reflect.ValueOf(primaryContactProp)) && (ok || !reflect.DeepEqual(v, primaryContactProp)) {
		obj["primaryContact"] = primaryContactProp
	}
	documentationProp, err := expandBigqueryAnalyticsHubListingDocumentation(d.Get("documentation"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("documentation"); !tpgresource.IsEmptyValue(reflect.ValueOf(documentationProp)) && (ok || !reflect.DeepEqual(v, documentationProp)) {
		obj["documentation"] = documentationProp
	}
	iconProp, err := expandBigqueryAnalyticsHubListingIcon(d.Get("icon"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("icon"); !tpgresource.IsEmptyValue(reflect.ValueOf(iconProp)) && (ok || !reflect.DeepEqual(v, iconProp)) {
		obj["icon"] = iconProp
	}
	requestAccessProp, err := expandBigqueryAnalyticsHubListingRequestAccess(d.Get("request_access"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("request_access"); !tpgresource.IsEmptyValue(reflect.ValueOf(requestAccessProp)) && (ok || !reflect.DeepEqual(v, requestAccessProp)) {
		obj["requestAccess"] = requestAccessProp
	}
	dataProviderProp, err := expandBigqueryAnalyticsHubListingDataProvider(d.Get("data_provider"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("data_provider"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataProviderProp)) && (ok || !reflect.DeepEqual(v, dataProviderProp)) {
		obj["dataProvider"] = dataProviderProp
	}
	publisherProp, err := expandBigqueryAnalyticsHubListingPublisher(d.Get("publisher"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("publisher"); !tpgresource.IsEmptyValue(reflect.ValueOf(publisherProp)) && (ok || !reflect.DeepEqual(v, publisherProp)) {
		obj["publisher"] = publisherProp
	}
	categoriesProp, err := expandBigqueryAnalyticsHubListingCategories(d.Get("categories"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("categories"); !tpgresource.IsEmptyValue(reflect.ValueOf(categoriesProp)) && (ok || !reflect.DeepEqual(v, categoriesProp)) {
		obj["categories"] = categoriesProp
	}
	bigqueryDatasetProp, err := expandBigqueryAnalyticsHubListingBigqueryDataset(d.Get("bigquery_dataset"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bigquery_dataset"); !tpgresource.IsEmptyValue(reflect.ValueOf(bigqueryDatasetProp)) && (ok || !reflect.DeepEqual(v, bigqueryDatasetProp)) {
		obj["bigqueryDataset"] = bigqueryDatasetProp
	}
	restrictedExportConfigProp, err := expandBigqueryAnalyticsHubListingRestrictedExportConfig(d.Get("restricted_export_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("restricted_export_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(restrictedExportConfigProp)) && (ok || !reflect.DeepEqual(v, restrictedExportConfigProp)) {
		obj["restrictedExportConfig"] = restrictedExportConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BigqueryAnalyticsHubBasePath}}projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings?listing_id={{listing_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Listing: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Listing: %s", err)
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
		return fmt.Errorf("Error creating Listing: %s", err)
	}
	if err := d.Set("name", flattenBigqueryAnalyticsHubListingName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Listing %q: %#v", d.Id(), res)

	return resourceBigqueryAnalyticsHubListingRead(d, meta)
}

func resourceBigqueryAnalyticsHubListingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BigqueryAnalyticsHubBasePath}}projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Listing: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("BigqueryAnalyticsHubListing %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Listing: %s", err)
	}

	if err := d.Set("name", flattenBigqueryAnalyticsHubListingName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Listing: %s", err)
	}
	if err := d.Set("display_name", flattenBigqueryAnalyticsHubListingDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Listing: %s", err)
	}
	if err := d.Set("description", flattenBigqueryAnalyticsHubListingDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Listing: %s", err)
	}
	if err := d.Set("primary_contact", flattenBigqueryAnalyticsHubListingPrimaryContact(res["primaryContact"], d, config)); err != nil {
		return fmt.Errorf("Error reading Listing: %s", err)
	}
	if err := d.Set("documentation", flattenBigqueryAnalyticsHubListingDocumentation(res["documentation"], d, config)); err != nil {
		return fmt.Errorf("Error reading Listing: %s", err)
	}
	if err := d.Set("icon", flattenBigqueryAnalyticsHubListingIcon(res["icon"], d, config)); err != nil {
		return fmt.Errorf("Error reading Listing: %s", err)
	}
	if err := d.Set("request_access", flattenBigqueryAnalyticsHubListingRequestAccess(res["requestAccess"], d, config)); err != nil {
		return fmt.Errorf("Error reading Listing: %s", err)
	}
	if err := d.Set("data_provider", flattenBigqueryAnalyticsHubListingDataProvider(res["dataProvider"], d, config)); err != nil {
		return fmt.Errorf("Error reading Listing: %s", err)
	}
	if err := d.Set("publisher", flattenBigqueryAnalyticsHubListingPublisher(res["publisher"], d, config)); err != nil {
		return fmt.Errorf("Error reading Listing: %s", err)
	}
	if err := d.Set("categories", flattenBigqueryAnalyticsHubListingCategories(res["categories"], d, config)); err != nil {
		return fmt.Errorf("Error reading Listing: %s", err)
	}
	if err := d.Set("bigquery_dataset", flattenBigqueryAnalyticsHubListingBigqueryDataset(res["bigqueryDataset"], d, config)); err != nil {
		return fmt.Errorf("Error reading Listing: %s", err)
	}
	if err := d.Set("restricted_export_config", flattenBigqueryAnalyticsHubListingRestrictedExportConfig(res["restrictedExportConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Listing: %s", err)
	}

	return nil
}

func resourceBigqueryAnalyticsHubListingUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Listing: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandBigqueryAnalyticsHubListingDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandBigqueryAnalyticsHubListingDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	primaryContactProp, err := expandBigqueryAnalyticsHubListingPrimaryContact(d.Get("primary_contact"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("primary_contact"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, primaryContactProp)) {
		obj["primaryContact"] = primaryContactProp
	}
	documentationProp, err := expandBigqueryAnalyticsHubListingDocumentation(d.Get("documentation"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("documentation"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, documentationProp)) {
		obj["documentation"] = documentationProp
	}
	iconProp, err := expandBigqueryAnalyticsHubListingIcon(d.Get("icon"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("icon"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, iconProp)) {
		obj["icon"] = iconProp
	}
	requestAccessProp, err := expandBigqueryAnalyticsHubListingRequestAccess(d.Get("request_access"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("request_access"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, requestAccessProp)) {
		obj["requestAccess"] = requestAccessProp
	}
	dataProviderProp, err := expandBigqueryAnalyticsHubListingDataProvider(d.Get("data_provider"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("data_provider"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, dataProviderProp)) {
		obj["dataProvider"] = dataProviderProp
	}
	publisherProp, err := expandBigqueryAnalyticsHubListingPublisher(d.Get("publisher"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("publisher"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, publisherProp)) {
		obj["publisher"] = publisherProp
	}
	categoriesProp, err := expandBigqueryAnalyticsHubListingCategories(d.Get("categories"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("categories"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, categoriesProp)) {
		obj["categories"] = categoriesProp
	}
	bigqueryDatasetProp, err := expandBigqueryAnalyticsHubListingBigqueryDataset(d.Get("bigquery_dataset"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bigquery_dataset"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, bigqueryDatasetProp)) {
		obj["bigqueryDataset"] = bigqueryDatasetProp
	}
	restrictedExportConfigProp, err := expandBigqueryAnalyticsHubListingRestrictedExportConfig(d.Get("restricted_export_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("restricted_export_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, restrictedExportConfigProp)) {
		obj["restrictedExportConfig"] = restrictedExportConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BigqueryAnalyticsHubBasePath}}projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Listing %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("primary_contact") {
		updateMask = append(updateMask, "primaryContact")
	}

	if d.HasChange("documentation") {
		updateMask = append(updateMask, "documentation")
	}

	if d.HasChange("icon") {
		updateMask = append(updateMask, "icon")
	}

	if d.HasChange("request_access") {
		updateMask = append(updateMask, "requestAccess")
	}

	if d.HasChange("data_provider") {
		updateMask = append(updateMask, "dataProvider")
	}

	if d.HasChange("publisher") {
		updateMask = append(updateMask, "publisher")
	}

	if d.HasChange("categories") {
		updateMask = append(updateMask, "categories")
	}

	if d.HasChange("bigquery_dataset") {
		updateMask = append(updateMask, "bigqueryDataset")
	}

	if d.HasChange("restricted_export_config") {
		updateMask = append(updateMask, "restrictedExportConfig")
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
			return fmt.Errorf("Error updating Listing %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Listing %q: %#v", d.Id(), res)
		}

	}

	return resourceBigqueryAnalyticsHubListingRead(d, meta)
}

func resourceBigqueryAnalyticsHubListingDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Listing: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{BigqueryAnalyticsHubBasePath}}projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Listing %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "Listing")
	}

	log.Printf("[DEBUG] Finished deleting Listing %q: %#v", d.Id(), res)
	return nil
}

func resourceBigqueryAnalyticsHubListingImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/dataExchanges/(?P<data_exchange_id>[^/]+)/listings/(?P<listing_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<data_exchange_id>[^/]+)/(?P<listing_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<data_exchange_id>[^/]+)/(?P<listing_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBigqueryAnalyticsHubListingName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryAnalyticsHubListingDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryAnalyticsHubListingDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryAnalyticsHubListingPrimaryContact(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryAnalyticsHubListingDocumentation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryAnalyticsHubListingIcon(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryAnalyticsHubListingRequestAccess(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryAnalyticsHubListingDataProvider(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["name"] =
		flattenBigqueryAnalyticsHubListingDataProviderName(original["name"], d, config)
	transformed["primary_contact"] =
		flattenBigqueryAnalyticsHubListingDataProviderPrimaryContact(original["primaryContact"], d, config)
	return []interface{}{transformed}
}
func flattenBigqueryAnalyticsHubListingDataProviderName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryAnalyticsHubListingDataProviderPrimaryContact(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryAnalyticsHubListingPublisher(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["name"] =
		flattenBigqueryAnalyticsHubListingPublisherName(original["name"], d, config)
	transformed["primary_contact"] =
		flattenBigqueryAnalyticsHubListingPublisherPrimaryContact(original["primaryContact"], d, config)
	return []interface{}{transformed}
}
func flattenBigqueryAnalyticsHubListingPublisherName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryAnalyticsHubListingPublisherPrimaryContact(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryAnalyticsHubListingCategories(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryAnalyticsHubListingBigqueryDataset(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataset"] =
		flattenBigqueryAnalyticsHubListingBigqueryDatasetDataset(original["dataset"], d, config)
	return []interface{}{transformed}
}
func flattenBigqueryAnalyticsHubListingBigqueryDatasetDataset(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryAnalyticsHubListingRestrictedExportConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enabled"] =
		flattenBigqueryAnalyticsHubListingRestrictedExportConfigEnabled(original["enabled"], d, config)
	transformed["restrict_query_result"] =
		flattenBigqueryAnalyticsHubListingRestrictedExportConfigRestrictQueryResult(original["restrictQueryResult"], d, config)
	return []interface{}{transformed}
}
func flattenBigqueryAnalyticsHubListingRestrictedExportConfigEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryAnalyticsHubListingRestrictedExportConfigRestrictQueryResult(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandBigqueryAnalyticsHubListingDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingPrimaryContact(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingDocumentation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingIcon(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingRequestAccess(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingDataProvider(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandBigqueryAnalyticsHubListingDataProviderName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["name"] = transformedName
	}

	transformedPrimaryContact, err := expandBigqueryAnalyticsHubListingDataProviderPrimaryContact(original["primary_contact"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrimaryContact); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["primaryContact"] = transformedPrimaryContact
	}

	return transformed, nil
}

func expandBigqueryAnalyticsHubListingDataProviderName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingDataProviderPrimaryContact(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingPublisher(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandBigqueryAnalyticsHubListingPublisherName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["name"] = transformedName
	}

	transformedPrimaryContact, err := expandBigqueryAnalyticsHubListingPublisherPrimaryContact(original["primary_contact"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrimaryContact); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["primaryContact"] = transformedPrimaryContact
	}

	return transformed, nil
}

func expandBigqueryAnalyticsHubListingPublisherName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingPublisherPrimaryContact(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingCategories(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingBigqueryDataset(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDataset, err := expandBigqueryAnalyticsHubListingBigqueryDatasetDataset(original["dataset"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDataset); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dataset"] = transformedDataset
	}

	return transformed, nil
}

func expandBigqueryAnalyticsHubListingBigqueryDatasetDataset(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingRestrictedExportConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandBigqueryAnalyticsHubListingRestrictedExportConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedRestrictQueryResult, err := expandBigqueryAnalyticsHubListingRestrictedExportConfigRestrictQueryResult(original["restrict_query_result"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRestrictQueryResult); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["restrictQueryResult"] = transformedRestrictQueryResult
	}

	return transformed, nil
}

func expandBigqueryAnalyticsHubListingRestrictedExportConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingRestrictedExportConfigRestrictQueryResult(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
