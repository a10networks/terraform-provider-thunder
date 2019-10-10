package main

/*
func main() {
	id := getAuthHeader()
	post_virtual_server_object(id)
}
*/

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-vthunder/vthunder"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vthunder.Provider})
}
