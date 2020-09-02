package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_RADIUS_SERVER_RESOURCE = `
resource "vthunder_fw_radius_server" "FwRadiusTest" {
	listen_port = "1024" 
}
`

//Acceptance test
func TestAccFwRadiusServer_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_RADIUS_SERVER_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_fw_radius_server.FwRadiusTest", "listen_port", "1024"),
				),
			},
		},
	})
}
