// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package storagecontrol

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/fwutils"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/registry"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var _ ephemeral.EphemeralResource = &googleEphemeralStorageControlProjectIntelligenceConfig{}

func init() {
	registry.FrameworkEphemeralResource{
		Name:        "google_storage_control_project_intelligence_config",
		ProductName: "storagecontrol",
		Func:        GoogleEphemeralStorageControlProjectIntelligenceConfig,
	}.Register()
}

func GoogleEphemeralStorageControlProjectIntelligenceConfig() ephemeral.EphemeralResource {
	return &googleEphemeralStorageControlProjectIntelligenceConfig{}
}

type googleEphemeralStorageControlProjectIntelligenceConfig struct {
	providerConfig *transport_tpg.Config
}

type ephemeralStorageControlProjectIntelligenceConfigModel struct {
	Name                        types.String `tfsdk:"name"`
	EditionConfig               types.String `tfsdk:"edition_config"`
	Filter                      types.Object `tfsdk:"filter"`
	EffectiveIntelligenceConfig types.Object `tfsdk:"effective_intelligence_config"`
	TrialConfig                 types.Object `tfsdk:"trial_config"`
	UpdateTime                  types.String `tfsdk:"update_time"`
}

func (p *googleEphemeralStorageControlProjectIntelligenceConfig) Metadata(ctx context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	// TerraformName already includes the "google_" prefix and the product name
	// (e.g. "google_storage_control_project_intelligence_config") — unlike
	// req.ProviderTypeName + name, which only works for resources with no
	// product prefix in their type name (the handwritten ephemeral resources
	// this was originally copied from, e.g. google_client_config, don't have one).
	resp.TypeName = "google_storage_control_project_intelligence_config"
}

func (p *googleEphemeralStorageControlProjectIntelligenceConfig) Schema(ctx context.Context, req ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = schema.Schema{

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{

				Required: true,
			},
			"edition_config": schema.StringAttribute{

				Computed: true,
			},
			"filter": schema.ObjectAttribute{
				AttributeTypes: map[string]attr.Type{
					"excluded_cloud_storage_buckets": types.ObjectType{AttrTypes: map[string]attr.Type{
						"bucket_id_regexes": types.ListType{ElemType: types.StringType},
					}},
					"excluded_cloud_storage_locations": types.ObjectType{AttrTypes: map[string]attr.Type{
						"locations": types.ListType{ElemType: types.StringType},
					}},
					"included_cloud_storage_buckets": types.ObjectType{AttrTypes: map[string]attr.Type{
						"bucket_id_regexes": types.ListType{ElemType: types.StringType},
					}},
					"included_cloud_storage_locations": types.ObjectType{AttrTypes: map[string]attr.Type{
						"locations": types.ListType{ElemType: types.StringType},
					}},
				},

				Computed: true,
			},
			"effective_intelligence_config": schema.ObjectAttribute{
				AttributeTypes: map[string]attr.Type{
					"effective_edition":   types.StringType,
					"intelligence_config": types.StringType,
				},

				Computed: true,
			},
			"trial_config": schema.ObjectAttribute{
				AttributeTypes: map[string]attr.Type{
					"expire_time": types.StringType,
				},

				Computed: true,
			},
			"update_time": schema.StringAttribute{

				Computed: true,
			},
		},
	}
}

func (p *googleEphemeralStorageControlProjectIntelligenceConfig) Configure(ctx context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	pd, ok := req.ProviderData.(*transport_tpg.Config)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Ephemeral Resource Configure Type",
			fmt.Sprintf("Expected *transport_tpg.Config, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	p.providerConfig = pd
}

func (p *googleEphemeralStorageControlProjectIntelligenceConfig) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var data ephemeralStorageControlProjectIntelligenceConfigModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	config := p.providerConfig
	userAgent := config.UserAgent

	url := fmt.Sprintf("%sprojects/%s/locations/global/intelligenceConfig", transport_tpg.BaseUrl(Product, config), data.Name.ValueString())

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   config.Project,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		resp.Diagnostics.AddError("Error retrieving ProjectIntelligenceConfig", err.Error())
		return
	}

	editionConfigVal, editionConfigDiags := fwutils.FlattenEphemeralValue(res["editionConfig"], fwutils.EphemeralFieldMapping{
		ApiName:  "editionConfig",
		AttrType: types.StringType,
	})
	resp.Diagnostics.Append(editionConfigDiags...)
	data.EditionConfig = editionConfigVal.(types.String)

	updateTimeVal, updateTimeDiags := fwutils.FlattenEphemeralValue(res["updateTime"], fwutils.EphemeralFieldMapping{
		ApiName:  "updateTime",
		AttrType: types.StringType,
	})
	resp.Diagnostics.Append(updateTimeDiags...)
	data.UpdateTime = updateTimeVal.(types.String)

	filterVal, filterDiags := fwutils.FlattenEphemeralValue(res["filter"], fwutils.EphemeralFieldMapping{
		ApiName: "filter",
		AttrType: types.ObjectType{AttrTypes: map[string]attr.Type{
			"excluded_cloud_storage_buckets": types.ObjectType{AttrTypes: map[string]attr.Type{
				"bucket_id_regexes": types.ListType{ElemType: types.StringType},
			}},
			"excluded_cloud_storage_locations": types.ObjectType{AttrTypes: map[string]attr.Type{
				"locations": types.ListType{ElemType: types.StringType},
			}},
			"included_cloud_storage_buckets": types.ObjectType{AttrTypes: map[string]attr.Type{
				"bucket_id_regexes": types.ListType{ElemType: types.StringType},
			}},
			"included_cloud_storage_locations": types.ObjectType{AttrTypes: map[string]attr.Type{
				"locations": types.ListType{ElemType: types.StringType},
			}},
		}},
		Fields: map[string]fwutils.EphemeralFieldMapping{
			"excluded_cloud_storage_buckets": fwutils.EphemeralFieldMapping{
				ApiName: "excludedCloudStorageBuckets",
				AttrType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"bucket_id_regexes": types.ListType{ElemType: types.StringType},
				}},
				Fields: map[string]fwutils.EphemeralFieldMapping{
					"bucket_id_regexes": fwutils.EphemeralFieldMapping{
						ApiName:  "bucketIdRegexes",
						AttrType: types.ListType{ElemType: types.StringType},
					},
				},
			},
			"excluded_cloud_storage_locations": fwutils.EphemeralFieldMapping{
				ApiName: "excludedCloudStorageLocations",
				AttrType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"locations": types.ListType{ElemType: types.StringType},
				}},
				Fields: map[string]fwutils.EphemeralFieldMapping{
					"locations": fwutils.EphemeralFieldMapping{
						ApiName:  "locations",
						AttrType: types.ListType{ElemType: types.StringType},
					},
				},
			},
			"included_cloud_storage_buckets": fwutils.EphemeralFieldMapping{
				ApiName: "includedCloudStorageBuckets",
				AttrType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"bucket_id_regexes": types.ListType{ElemType: types.StringType},
				}},
				Fields: map[string]fwutils.EphemeralFieldMapping{
					"bucket_id_regexes": fwutils.EphemeralFieldMapping{
						ApiName:  "bucketIdRegexes",
						AttrType: types.ListType{ElemType: types.StringType},
					},
				},
			},
			"included_cloud_storage_locations": fwutils.EphemeralFieldMapping{
				ApiName: "includedCloudStorageLocations",
				AttrType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"locations": types.ListType{ElemType: types.StringType},
				}},
				Fields: map[string]fwutils.EphemeralFieldMapping{
					"locations": fwutils.EphemeralFieldMapping{
						ApiName:  "locations",
						AttrType: types.ListType{ElemType: types.StringType},
					},
				},
			},
		},
	})
	resp.Diagnostics.Append(filterDiags...)
	data.Filter = filterVal.(types.Object)

	effectiveIntelligenceConfigVal, effectiveIntelligenceConfigDiags := fwutils.FlattenEphemeralValue(res["effectiveIntelligenceConfig"], fwutils.EphemeralFieldMapping{
		ApiName: "effectiveIntelligenceConfig",
		AttrType: types.ObjectType{AttrTypes: map[string]attr.Type{
			"effective_edition":   types.StringType,
			"intelligence_config": types.StringType,
		}},
		Fields: map[string]fwutils.EphemeralFieldMapping{
			"effective_edition": fwutils.EphemeralFieldMapping{
				ApiName:  "effectiveEdition",
				AttrType: types.StringType,
			},
			"intelligence_config": fwutils.EphemeralFieldMapping{
				ApiName:  "intelligenceConfig",
				AttrType: types.StringType,
			},
		},
	})
	resp.Diagnostics.Append(effectiveIntelligenceConfigDiags...)
	data.EffectiveIntelligenceConfig = effectiveIntelligenceConfigVal.(types.Object)

	trialConfigVal, trialConfigDiags := fwutils.FlattenEphemeralValue(res["trialConfig"], fwutils.EphemeralFieldMapping{
		ApiName: "trialConfig",
		AttrType: types.ObjectType{AttrTypes: map[string]attr.Type{
			"expire_time": types.StringType,
		}},
		Fields: map[string]fwutils.EphemeralFieldMapping{
			"expire_time": fwutils.EphemeralFieldMapping{
				ApiName:  "expireTime",
				AttrType: types.StringType,
			},
		},
	})
	resp.Diagnostics.Append(trialConfigDiags...)
	data.TrialConfig = trialConfigVal.(types.Object)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.Result.Set(ctx, data)...)
}
