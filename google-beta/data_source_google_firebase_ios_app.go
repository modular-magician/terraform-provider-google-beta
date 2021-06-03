package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGoogleFirebaseIosApp() *schema.Resource {
	// Generate datasource schema from resource
	dsSchema := datasourceSchemaFromResourceSchema(resourceFirebaseIosApp().Schema)

	// Set 'Required' schema elements
	addRequiredFieldsToSchema(dsSchema, "app_id")

	// Set 'Optional' schema elements
	addOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceGoogleFirebaseIosAppRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleFirebaseIosAppRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	// Blindly use zero-value if type assertion fails
	appId, ok := d.Get("app_id").(string)

	// Since an `app_id` is a unique identifier, the Unique Resource from
	// Sub-Collection access pattern may be used here, in the format:
	// `projects/-/iosApps/{{app_id}}`
	project, ok := d.Get("project").(string)
	if project == "" {
		project = "-"
	}

	name := fmt.Sprintf("projects/%s/iosApps/%s", project, appId)

	d.SetId(name)
	// if err := d.Set("name", name); err != nil {
	// 	return fmt.Errorf("Error setting name: %s", err)
	// }
	return resourceFirebaseIosAppRead(d, meta)
}
