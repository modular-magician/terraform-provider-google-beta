// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"context"
	"testing"

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var TestAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	configs = make(map[string]*transport_tpg.Config)
	fwProviders = make(map[string]*frameworkTestProvider)
	sources = make(map[string]VcrSource)
	testAccProvider = Provider()
	TestAccProviders = map[string]*schema.Provider{
		"google": testAccProvider,
	}
}

func GoogleProviderConfig(t *testing.T) *transport_tpg.Config {
	configsLock.RLock()
	config, ok := configs[t.Name()]
	configsLock.RUnlock()
	if ok {
		return config
	}

	sdkProvider := Provider()
	rc := terraform.ResourceConfig{}
	sdkProvider.Configure(context.Background(), &rc)
	return sdkProvider.Meta().(*transport_tpg.Config)
}
