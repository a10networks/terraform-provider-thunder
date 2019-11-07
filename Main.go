package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"terraform-provider-vthunder/vthunder"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vthunder.Provider})
}
