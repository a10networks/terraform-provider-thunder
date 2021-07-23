package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/a10networks/terraform-provider-thunder/thunder"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: thunder.Provider})
}
