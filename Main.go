package main

/*
func main() {
	id := getAuthHeader()
	post_virtual_server_object(id)
}
*/

import (
	"vthunder"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vthunder.Provider})
}
