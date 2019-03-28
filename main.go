package main

import (
	"github.com/hashicorp/terraform/plugin"
	google "github.com/terraform-providers/terraform-provider-google-beta/google-beta"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: google.Provider})
}
