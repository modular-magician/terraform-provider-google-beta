// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	vertexai "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai/beta"
)

func resourceVertexAiMetadataSchema() *schema.Resource {
	return &schema.Resource{
		Create: resourceVertexAiMetadataSchemaCreate,
		Read:   resourceVertexAiMetadataSchemaRead,
		Delete: resourceVertexAiMetadataSchemaDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVertexAiMetadataSchemaImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The location for the resource",
			},

			"metadata_store": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "The metadata store for the resource",
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Output only. The resource name of the MetadataSchema.",
			},

			"schema": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Required. The raw YAML string representation of the MetadataSchema. The combination of [MetadataSchema.version] and the schema name given by `title` in [MetadataSchema.schema] must be unique within a MetadataStore. The schema is defined as an OpenAPI 3.0.2 [MetadataSchema Object](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#schemaObject)",
			},

			"schema_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the MetadataSchema. This is a property that identifies which metadata types will use the MetadataSchema. Possible values: METADATA_SCHEMA_TYPE_UNSPECIFIED, ARTIFACT_TYPE, EXECUTION_TYPE, CONTEXT_TYPE",
			},

			"schema_version": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The version of the MetadataSchema. The version's format must match the following regular expression: `^[0-9]+.+.+$`, which would allow to order/compare different versions. Example: 1.0.0, 1.0.1, etc.",
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "The project for the resource",
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Timestamp when this MetadataSchema was created.",
			},
		},
	}
}

func resourceVertexAiMetadataSchemaCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vertexai.MetadataSchema{
		Location:      dcl.String(d.Get("location").(string)),
		MetadataStore: dcl.String(d.Get("metadata_store").(string)),
		Name:          dcl.String(d.Get("name").(string)),
		Schema:        dcl.String(d.Get("schema").(string)),
		SchemaType:    vertexai.MetadataSchemaSchemaTypeEnumRef(d.Get("schema_type").(string)),
		SchemaVersion: dcl.String(d.Get("schema_version").(string)),
		Project:       dcl.String(project),
	}

	id, err := obj.ID()
	if err != nil {
		return fmt.Errorf("error constructing id: %s", err)
	}
	d.SetId(id)
	directive := CreateDirective
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLVertexAiClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutCreate))
	if bp, err := replaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	res, err := client.ApplyMetadataSchema(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating MetadataSchema: %s", err)
	}

	log.Printf("[DEBUG] Finished creating MetadataSchema %q: %#v", d.Id(), res)

	return resourceVertexAiMetadataSchemaRead(d, meta)
}

func resourceVertexAiMetadataSchemaRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vertexai.MetadataSchema{
		Location:      dcl.String(d.Get("location").(string)),
		MetadataStore: dcl.String(d.Get("metadata_store").(string)),
		Name:          dcl.String(d.Get("name").(string)),
		Schema:        dcl.String(d.Get("schema").(string)),
		SchemaType:    vertexai.MetadataSchemaSchemaTypeEnumRef(d.Get("schema_type").(string)),
		SchemaVersion: dcl.String(d.Get("schema_version").(string)),
		Project:       dcl.String(project),
	}

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLVertexAiClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutRead))
	if bp, err := replaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	res, err := client.GetMetadataSchema(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("VertexAiMetadataSchema %q", d.Id())
		return handleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("metadata_store", res.MetadataStore); err != nil {
		return fmt.Errorf("error setting metadata_store in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("schema", res.Schema); err != nil {
		return fmt.Errorf("error setting schema in state: %s", err)
	}
	if err = d.Set("schema_type", res.SchemaType); err != nil {
		return fmt.Errorf("error setting schema_type in state: %s", err)
	}
	if err = d.Set("schema_version", res.SchemaVersion); err != nil {
		return fmt.Errorf("error setting schema_version in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}

	return nil
}

func resourceVertexAiMetadataSchemaDelete(d *schema.ResourceData, meta interface{}) error {

	log.Printf("[DEBUG] Finished deleting MetadataSchema %q", d.Id())
	return nil
}

func resourceVertexAiMetadataSchemaImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/metadataStores/(?P<metadata_store>[^/]+)/metadataSchemas/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<metadata_store>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<metadata_store>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/metadataStores/{{metadata_store}}/metadataSchemas/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}
