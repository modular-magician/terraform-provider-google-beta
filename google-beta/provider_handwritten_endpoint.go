package google

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// For generated resources, endpoint entries live in product-specific provider
// files. Collect handwritten ones here.

var CloudFunctionsDefaultBasePath = "https://cloudfunctions.googleapis.com/v1/"

var CloudFunctionsCustomEndpointEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
	DefaultFunc: schema.MultiEnvDefaultFunc([]string{
		"GOOGLE_CLOUD_FUNCTIONS_CUSTOM_ENDPOINT",
	}, CloudFunctionsDefaultBasePath),
}
