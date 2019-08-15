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

// If the base path has changed as a result of your PR, make sure to update
// the provider_reference page!
var IapDefaultBasePath = "https://iap.googleapis.com/v1/"
var IapCustomEndpointEntryKey = "iap_custom_endpoint"
var IapCustomEndpointEntry = &schema.Schema{
	Type:         schema.TypeString,
	Optional:     true,
	ValidateFunc: validateCustomEndpoint,
	DefaultFunc: schema.MultiEnvDefaultFunc([]string{
		"GOOGLE_IAP_CUSTOM_ENDPOINT",
	}, IapDefaultBasePath),
}

var GeneratedIapResourcesMap = map[string]*schema.Resource{
	"google_iap_web_iam_binding":                 ResourceIamBinding(IapWebIamSchema, IapWebIamUpdaterProducer, IapWebIdParseFunc),
	"google_iap_web_iam_member":                  ResourceIamMember(IapWebIamSchema, IapWebIamUpdaterProducer, IapWebIdParseFunc),
	"google_iap_web_iam_policy":                  ResourceIamPolicy(IapWebIamSchema, IapWebIamUpdaterProducer, IapWebIdParseFunc),
	"google_iap_web_type_compute_iam_binding":    ResourceIamBinding(IapWebTypeComputeIamSchema, IapWebTypeComputeIamUpdaterProducer, IapWebTypeComputeIdParseFunc),
	"google_iap_web_type_compute_iam_member":     ResourceIamMember(IapWebTypeComputeIamSchema, IapWebTypeComputeIamUpdaterProducer, IapWebTypeComputeIdParseFunc),
	"google_iap_web_type_compute_iam_policy":     ResourceIamPolicy(IapWebTypeComputeIamSchema, IapWebTypeComputeIamUpdaterProducer, IapWebTypeComputeIdParseFunc),
	"google_iap_web_type_app_engine_iam_binding": ResourceIamBinding(IapWebTypeAppEngineIamSchema, IapWebTypeAppEngineIamUpdaterProducer, IapWebTypeAppEngineIdParseFunc),
	"google_iap_web_type_app_engine_iam_member":  ResourceIamMember(IapWebTypeAppEngineIamSchema, IapWebTypeAppEngineIamUpdaterProducer, IapWebTypeAppEngineIdParseFunc),
	"google_iap_web_type_app_engine_iam_policy":  ResourceIamPolicy(IapWebTypeAppEngineIamSchema, IapWebTypeAppEngineIamUpdaterProducer, IapWebTypeAppEngineIdParseFunc),
	"google_iap_web_backend_service_iam_binding": ResourceIamBinding(IapWebBackendServiceIamSchema, IapWebBackendServiceIamUpdaterProducer, IapWebBackendServiceIdParseFunc),
	"google_iap_web_backend_service_iam_member":  ResourceIamMember(IapWebBackendServiceIamSchema, IapWebBackendServiceIamUpdaterProducer, IapWebBackendServiceIdParseFunc),
	"google_iap_web_backend_service_iam_policy":  ResourceIamPolicy(IapWebBackendServiceIamSchema, IapWebBackendServiceIamUpdaterProducer, IapWebBackendServiceIdParseFunc),
}
