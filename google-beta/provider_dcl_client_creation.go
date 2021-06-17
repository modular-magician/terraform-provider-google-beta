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

import (
	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"

	dataproc "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/beta"
	eventarc "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/beta"
	gkehub "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
)

func NewDCLDataprocClient(config *Config, userAgent, billingProject string) *dataproc.Client {
	dclClientOptions := dcl.WithHTTPClient(config.client)
	dclUserAgentOptions := dcl.WithUserAgent(userAgent)
	dclLoggerOptions := dcl.WithLogger(dclLogger{})
	var dclConfig *dcl.Config
	if config.UserProjectOverride && billingProject != "" {
		dclBillingProjectHeader := dcl.WithHeader("X-Goog-User-Project", billingProject)
		dclConfig = dcl.NewConfig(
			dclClientOptions,
			dclUserAgentOptions,
			dclLoggerOptions,
			dcl.WithBasePath(config.DataprocBasePath),
			dclBillingProjectHeader,
		)
	} else {
		dclConfig = dcl.NewConfig(
			dclClientOptions,
			dclUserAgentOptions,
			dclLoggerOptions,
			dcl.WithBasePath(config.DataprocBasePath),
		)
	}

	return dataproc.NewClient(dclConfig)
}

func NewDCLEventarcClient(config *Config, userAgent, billingProject string) *eventarc.Client {
	dclClientOptions := dcl.WithHTTPClient(config.client)
	dclUserAgentOptions := dcl.WithUserAgent(userAgent)
	dclLoggerOptions := dcl.WithLogger(dclLogger{})
	var dclConfig *dcl.Config
	if config.UserProjectOverride && billingProject != "" {
		dclBillingProjectHeader := dcl.WithHeader("X-Goog-User-Project", billingProject)
		dclConfig = dcl.NewConfig(
			dclClientOptions,
			dclUserAgentOptions,
			dclLoggerOptions,
			dcl.WithBasePath(config.EventarcBasePath),
			dclBillingProjectHeader,
		)
	} else {
		dclConfig = dcl.NewConfig(
			dclClientOptions,
			dclUserAgentOptions,
			dclLoggerOptions,
			dcl.WithBasePath(config.EventarcBasePath),
		)
	}

	return eventarc.NewClient(dclConfig)
}

func NewDCLGkeHubClient(config *Config, userAgent, billingProject string) *gkehub.Client {
	dclClientOptions := dcl.WithHTTPClient(config.client)
	dclUserAgentOptions := dcl.WithUserAgent(userAgent)
	dclLoggerOptions := dcl.WithLogger(dclLogger{})
	var dclConfig *dcl.Config
	if config.UserProjectOverride && billingProject != "" {
		dclBillingProjectHeader := dcl.WithHeader("X-Goog-User-Project", billingProject)
		dclConfig = dcl.NewConfig(
			dclClientOptions,
			dclUserAgentOptions,
			dclLoggerOptions,
			dcl.WithBasePath(config.GkeHubBasePath),
			dclBillingProjectHeader,
		)
	} else {
		dclConfig = dcl.NewConfig(
			dclClientOptions,
			dclUserAgentOptions,
			dclLoggerOptions,
			dcl.WithBasePath(config.GkeHubBasePath),
		)
	}

	return gkehub.NewClient(dclConfig)
}
