package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGoogleComputeResourcePolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGoogleComputeResourcePolicyRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"project": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"region": {
				Type:     schema.TypeString,
				Required: true,
			},

			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGoogleComputeResourcePolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	region, err := getRegion(d, config)
	if err != nil {
		return err
	}
	name := d.Get("name").(string)
	resourcePolicy, err := config.NewComputeClient(userAgent).ResourcePolicies.Get(project, region, name).Do()
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ResourcePolicy Not Found : %s", name))
	}
	if err := d.Set("self_link", resourcePolicy.SelfLink); err != nil {
		return fmt.Errorf("Error setting self_link: %s", err)
	}
	if err := d.Set("description", resourcePolicy.Description); err != nil {
		return fmt.Errorf("Error setting description: %s", err)
	}
	d.SetId(fmt.Sprintf("projects/%s/regions/%s/resourcePolicies/%s", project, region, name))
	return nil
}
