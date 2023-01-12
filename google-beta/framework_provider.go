package google

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &frameworkProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) provider.Provider {
	return &frameworkProvider{
		version: version,
	}
}

// frameworkProvider is the provider implementation.
type frameworkProvider struct {
	version string
}

// Metadata returns the provider type name.
func (p *frameworkProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "google-beta"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *frameworkProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

// Configure prepares an API client for data sources and resources.
func (p *frameworkProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

// DataSources defines the data sources implemented in the provider.
func (p *frameworkProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *frameworkProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
