package google

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// For generated resources, base path entries live in product-specific provider
// files. Collect handwritten ones here.

var CloudFunctionsDefaultBasePath = "https://cloudfunctions.googleapis.com/v1/"

var CloudFunctionsBasePathEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
	DefaultFunc: schema.MultiEnvDefaultFunc([]string{
		"GOOGLE_CLOUD_FUNCTIONS_BASE_PATH",
	}, CloudFunctionsDefaultBasePath),
}
