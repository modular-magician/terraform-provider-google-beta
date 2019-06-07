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

import "github.com/hashicorp/terraform/helper/schema"

var CloudBuildDefaultBasePath = "https://cloudbuild.googleapis.com/v1/"

var CloudBuildCustomEndpointEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
	DefaultFunc: schema.MultiEnvDefaultFunc([]string{
		"GOOGLE_CLOUD_BUILD_CUSTOM_ENDPOINT",
	}, CloudBuildDefaultBasePath),
}

var GeneratedCloudBuildResourcesMap = map[string]*schema.Resource{
	"google_cloudbuild_trigger": resourceCloudBuildTrigger(),
}
