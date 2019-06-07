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

var SpannerDefaultBasePath = "https://spanner.googleapis.com/v1/"

var SpannerCustomEndpointEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
	DefaultFunc: schema.MultiEnvDefaultFunc([]string{
		"GOOGLE_SPANNER_CUSTOM_ENDPOINT",
	}, SpannerDefaultBasePath),
}

var GeneratedSpannerResourcesMap = map[string]*schema.Resource{
	"google_spanner_instance": resourceSpannerInstance(),
	"google_spanner_database": resourceSpannerDatabase(),
}
