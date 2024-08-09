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

package kms

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func folderPrefixSuppress(_, old, new string, d *schema.ResourceData) bool {
	prefix := "folders/"
	return prefix+old == new || prefix+new == old
}

func ResourceKMSAutokeyConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceKMSAutokeyConfigCreate,
		Read:   resourceKMSAutokeyConfigRead,
		Update: resourceKMSAutokeyConfigUpdate,
		Delete: resourceKMSAutokeyConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceKMSAutokeyConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"folder": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: folderPrefixSuppress,
				Description:      `The folder for which to retrieve config.`,
			},
			"key_project": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The target key project for a given folder where KMS Autokey will provision a
CryptoKey for any new KeyHandle the Developer creates. Should have the form
'projects/<project_id_or_number>'.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceKMSAutokeyConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	keyProjectProp, err := expandKMSAutokeyConfigKeyProject(d.Get("key_project"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("key_project"); !tpgresource.IsEmptyValue(reflect.ValueOf(keyProjectProp)) && (ok || !reflect.DeepEqual(v, keyProjectProp)) {
		obj["keyProject"] = keyProjectProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{KMSBasePath}}folders/{{folder}}/autokeyConfig?updateMask=keyProject")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AutokeyConfig: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	url = strings.Replace(url, "folders/folders/", "folders/", 1)
	folderValue := d.Get("folder").(string)
	folderValue = strings.Replace(folderValue, "folders/", "", 1)
	d.Set("folder", folderValue)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating AutokeyConfig: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "folders/{{folder}}/autokeyConfig")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating AutokeyConfig %q: %#v", d.Id(), res)

	return resourceKMSAutokeyConfigRead(d, meta)
}

func resourceKMSAutokeyConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{KMSBasePath}}folders/{{folder}}/autokeyConfig")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	url = strings.Replace(url, "folders/folders/", "folders/", 1)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("KMSAutokeyConfig %q", d.Id()))
	}

	if err := d.Set("key_project", flattenKMSAutokeyConfigKeyProject(res["keyProject"], d, config)); err != nil {
		return fmt.Errorf("Error reading AutokeyConfig: %s", err)
	}

	return nil
}

func resourceKMSAutokeyConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	keyProjectProp, err := expandKMSAutokeyConfigKeyProject(d.Get("key_project"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("key_project"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, keyProjectProp)) {
		obj["keyProject"] = keyProjectProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{KMSBasePath}}folders/{{folder}}/autokeyConfig?updateMask=keyProject")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AutokeyConfig %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	url = strings.Replace(url, "folders/folders/", "folders/", 1)
	folderValue := d.Get("folder").(string)
	folderValue = strings.Replace(folderValue, "folders/", "", 1)
	d.Set("folder", folderValue)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

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
		return fmt.Errorf("Error updating AutokeyConfig %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating AutokeyConfig %q: %#v", d.Id(), res)
	}

	return resourceKMSAutokeyConfigRead(d, meta)
}

func resourceKMSAutokeyConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{KMSBasePath}}folders/{{folder}}/autokeyConfig?updateMask=keyProject")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	url = strings.Replace(url, "folders/folders/", "folders/", 1)

	log.Printf("[DEBUG] Deleting AutokeyConfig %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "AutokeyConfig")
	}

	log.Printf("[DEBUG] Finished deleting AutokeyConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceKMSAutokeyConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^folders/(?P<folder>[^/]+)/autokeyConfig$",
		"^(?P<folder>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "folders/{{folder}}/autokeyConfig")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenKMSAutokeyConfigKeyProject(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandKMSAutokeyConfigKeyProject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
