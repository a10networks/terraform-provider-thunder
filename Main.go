package main

import (
	"github.com/a10networks/terraform-provider-thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: thunder.Provider})
}
