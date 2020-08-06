// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFirebaseProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaseProjectCreate,
		Read:   resourceFirebaseProjectRead,
		Delete: resourceFirebaseProjectDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaseProjectImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The GCP project display name`,
			},
			"project_number": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The number of the google project that firebase is enabled on.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceFirebaseProjectCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}:addFirebase")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Project: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Project: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = firebaseOperationWaitTime(
		config, res, project, "Creating Project",
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Project: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Project %q: %#v", d.Id(), res)

	return resourceFirebaseProjectRead(d, meta)
}

func resourceFirebaseProjectRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("FirebaseProject %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Project: %s", err)
	}

	if err := d.Set("project_number", flattenFirebaseProjectProjectNumber(res["projectNumber"], d, config)); err != nil {
		return fmt.Errorf("Error reading Project: %s", err)
	}
	if err := d.Set("display_name", flattenFirebaseProjectDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Project: %s", err)
	}

	return nil
}

func resourceFirebaseProjectDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] Firebase Project resources"+
		" cannot be deleted from GCP. The resource %s will be removed from Terraform"+
		" state, but will still be present on the server.", d.Id())
	d.SetId("")

	return nil
}

func resourceFirebaseProjectImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)",
		"(?P<project>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenFirebaseProjectProjectNumber(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseProjectDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}
