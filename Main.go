package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"terraform-provider-thunder/thunder"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: thunder.Provider})
}
