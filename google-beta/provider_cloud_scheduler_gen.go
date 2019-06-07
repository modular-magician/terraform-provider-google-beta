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

var CloudSchedulerDefaultBasePath = "https://cloudscheduler.googleapis.com/v1/"

var CloudSchedulerCustomEndpointEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
	DefaultFunc: schema.MultiEnvDefaultFunc([]string{
		"GOOGLE_CLOUD_SCHEDULER_CUSTOM_ENDPOINT",
	}, CloudSchedulerDefaultBasePath),
}

var GeneratedCloudSchedulerResourcesMap = map[string]*schema.Resource{
	"google_cloud_scheduler_job": resourceCloudSchedulerJob(),
}
