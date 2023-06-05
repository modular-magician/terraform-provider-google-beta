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
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceApigeeSyncAuthorization() *schema.Resource {
	return &schema.Resource{
		Create: resourceApigeeSyncAuthorizationCreate,
		Read:   resourceApigeeSyncAuthorizationRead,
		Update: resourceApigeeSyncAuthorizationUpdate,
		Delete: resourceApigeeSyncAuthorizationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApigeeSyncAuthorizationImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"identities": {
				Type:     schema.TypeList,
				Required: true,
				Description: `Array of service accounts to grant access to control plane resources, each specified using the following format: 'serviceAccount:service-account-name'.

The 'service-account-name' is formatted like an email address. For example: my-synchronizer-manager-serviceAccount@my_project_id.iam.gserviceaccount.com

You might specify multiple service accounts, for example, if you have multiple environments and wish to assign a unique service account to each one.

The service accounts must have **Apigee Synchronizer Manager** role. See also [Create service accounts](https://cloud.google.com/apigee/docs/hybrid/v1.8/sa-about#create-the-service-accounts).`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the Apigee organization.`,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Entity tag (ETag) used for optimistic concurrency control as a way to help prevent simultaneous updates from overwriting each other.
Used internally during updates.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApigeeSyncAuthorizationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	identitiesProp, err := expandApigeeSyncAuthorizationIdentities(d.Get("identities"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("identities"); ok || !reflect.DeepEqual(v, identitiesProp) {
		obj["identities"] = identitiesProp
	}
	etagProp, err := expandApigeeSyncAuthorizationEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !isEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{name}}:setSyncAuthorization")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new SyncAuthorization: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating SyncAuthorization: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "organizations/{{name}}/syncAuthorization")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating SyncAuthorization %q: %#v", d.Id(), res)

	return resourceApigeeSyncAuthorizationRead(d, meta)
}

func resourceApigeeSyncAuthorizationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{name}}:getSyncAuthorization")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "POST", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ApigeeSyncAuthorization %q", d.Id()))
	}

	if err := d.Set("identities", flattenApigeeSyncAuthorizationIdentities(res["identities"], d, config)); err != nil {
		return fmt.Errorf("Error reading SyncAuthorization: %s", err)
	}
	if err := d.Set("etag", flattenApigeeSyncAuthorizationEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading SyncAuthorization: %s", err)
	}

	return nil
}

func resourceApigeeSyncAuthorizationUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	identitiesProp, err := expandApigeeSyncAuthorizationIdentities(d.Get("identities"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("identities"); ok || !reflect.DeepEqual(v, identitiesProp) {
		obj["identities"] = identitiesProp
	}
	etagProp, err := expandApigeeSyncAuthorizationEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}organizations/{{name}}:setSyncAuthorization")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating SyncAuthorization %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating SyncAuthorization %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating SyncAuthorization %q: %#v", d.Id(), res)
	}

	return resourceApigeeSyncAuthorizationRead(d, meta)
}

func resourceApigeeSyncAuthorizationDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] Apigee SyncAuthorization resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceApigeeSyncAuthorizationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"organizations/(?P<name>[^/]+)/syncAuthorization",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "organizations/{{name}}/syncAuthorization")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApigeeSyncAuthorizationIdentities(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeSyncAuthorizationEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandApigeeSyncAuthorizationIdentities(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeSyncAuthorizationEtag(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
