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

var SourceRepoDefaultBasePath = "https://sourcerepo.googleapis.com/v1/"
var SourceRepoCustomEndpointEntryKey = "source_repo_custom_endpoint"
var SourceRepoCustomEndpointEntry = &schema.Schema{
	Type:         schema.TypeString,
	Optional:     true,
	ValidateFunc: validateCustomEndpoint,
	DefaultFunc: schema.MultiEnvDefaultFunc([]string{
		"GOOGLE_SOURCE_REPO_CUSTOM_ENDPOINT",
	}, SourceRepoDefaultBasePath),
}

var GeneratedSourceRepoResourcesMap = map[string]*schema.Resource{
	"google_sourcerepo_repository": resourceSourceRepoRepository(),
}
