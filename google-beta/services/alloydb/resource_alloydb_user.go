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

package alloydb

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceAlloydbUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlloydbUserCreate,
		Read:   resourceAlloydbUserRead,
		Update: resourceAlloydbUserUpdate,
		Delete: resourceAlloydbUserDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAlloydbUserImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `Identifies the alloydb cluster. Must be in the format
'projects/{project}/locations/{location}/clusters/{cluster_id}'`,
			},
			"user_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the alloydb user.`,
			},
			"database_roles": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Optional. List of database roles this user has. The database role strings are subject to the PostgreSQL naming conventions.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Input only. Password for the user. This field is required but shouldn't be set if user_type is set to 'ALLOYDB_IAM_USER'`,
				Sensitive:   true,
			},
			"user_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ALLOYDB_BUILT_IN", "ALLOYDB_IAM_USER", ""}),
				Description:  `Optional. Type of this user. Default value: "ALLOYDB_BUILT_IN" Possible values: ["ALLOYDB_BUILT_IN", "ALLOYDB_IAM_USER"]`,
				Default:      "ALLOYDB_BUILT_IN",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Name of the resource in the form of projects/{project}/locations/{location}/cluster/{cluster}/users/{user}.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceAlloydbUserCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	passwordProp, err := expandAlloydbUserPassword(d.Get("password"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("password"); !tpgresource.IsEmptyValue(reflect.ValueOf(passwordProp)) && (ok || !reflect.DeepEqual(v, passwordProp)) {
		obj["password"] = passwordProp
	}
	databaseRolesProp, err := expandAlloydbUserDatabaseRoles(d.Get("database_roles"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("database_roles"); !tpgresource.IsEmptyValue(reflect.ValueOf(databaseRolesProp)) && (ok || !reflect.DeepEqual(v, databaseRolesProp)) {
		obj["databaseRoles"] = databaseRolesProp
	}
	userTypeProp, err := expandAlloydbUserUserType(d.Get("user_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(userTypeProp)) && (ok || !reflect.DeepEqual(v, userTypeProp)) {
		obj["userType"] = userTypeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/users?userId={{user_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new User: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
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
		return fmt.Errorf("Error creating User: %s", err)
	}
	if err := d.Set("name", flattenAlloydbUserName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{cluster}}/users/{{user_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = transport_tpg.PollingWaitTime(resourceAlloydbUserPollRead(d, meta), transport_tpg.PollCheckForExistence, "Creating User", d.Timeout(schema.TimeoutCreate), 10)
	if err != nil {
		return fmt.Errorf("Error waiting to create User: %s", err)
	}

	log.Printf("[DEBUG] Finished creating User %q: %#v", d.Id(), res)

	return resourceAlloydbUserRead(d, meta)
}

func resourceAlloydbUserPollRead(d *schema.ResourceData, meta interface{}) transport_tpg.PollReadFunc {
	return func() (map[string]interface{}, error) {
		config := meta.(*transport_tpg.Config)

		url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/users/{{user_id}}")
		if err != nil {
			return nil, err
		}

		billingProject := ""

		// err == nil indicates that the billing_project value was found
		if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
			billingProject = bp
		}

		userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
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

func resourceAlloydbUserRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/users/{{user_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("AlloydbUser %q", d.Id()))
	}

	if err := d.Set("name", flattenAlloydbUserName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading User: %s", err)
	}
	if err := d.Set("database_roles", flattenAlloydbUserDatabaseRoles(res["databaseRoles"], d, config)); err != nil {
		return fmt.Errorf("Error reading User: %s", err)
	}
	if err := d.Set("user_type", flattenAlloydbUserUserType(res["userType"], d, config)); err != nil {
		return fmt.Errorf("Error reading User: %s", err)
	}

	return nil
}

func resourceAlloydbUserUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	passwordProp, err := expandAlloydbUserPassword(d.Get("password"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("password"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, passwordProp)) {
		obj["password"] = passwordProp
	}
	databaseRolesProp, err := expandAlloydbUserDatabaseRoles(d.Get("database_roles"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("database_roles"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, databaseRolesProp)) {
		obj["databaseRoles"] = databaseRolesProp
	}
	userTypeProp, err := expandAlloydbUserUserType(d.Get("user_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, userTypeProp)) {
		obj["userType"] = userTypeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/users/{{user_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating User %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("password") {
		updateMask = append(updateMask, "password")
	}

	if d.HasChange("database_roles") {
		updateMask = append(updateMask, "databaseRoles")
	}

	if d.HasChange("user_type") {
		updateMask = append(updateMask, "userType")
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

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating User %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating User %q: %#v", d.Id(), res)
	}

	err = transport_tpg.PollingWaitTime(resourceAlloydbUserPollRead(d, meta), transport_tpg.PollCheckForExistence, "Updating User", d.Timeout(schema.TimeoutUpdate), 10)
	if err != nil {
		return err
	}

	return resourceAlloydbUserRead(d, meta)
}

func resourceAlloydbUserDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/users/{{user_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting User %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
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
		return transport_tpg.HandleNotFoundError(err, d, "User")
	}

	err = transport_tpg.PollingWaitTime(resourceAlloydbUserPollRead(d, meta), transport_tpg.PollCheckForAbsence, "Deleting User", d.Timeout(schema.TimeoutCreate), 10)
	if err != nil {
		return fmt.Errorf("Error waiting to delete User: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting User %q: %#v", d.Id(), res)
	return nil
}

func resourceAlloydbUserImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{
		"(?P<cluster>.+)/users/(?P<user_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{cluster}}/users/{{user_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAlloydbUserName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbUserDatabaseRoles(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbUserUserType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandAlloydbUserPassword(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbUserDatabaseRoles(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbUserUserType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
