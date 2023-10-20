package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_FW_RADIUS_SERVER_RESOURCE = `
resource "thunder_fw_radius_server" "FwRadiusTest" {
	listen_port = "1024" 
}
`

// Acceptance test
func TestAccFwRadiusServer_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_RADIUS_SERVER_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_fw_radius_server.FwRadiusTest", "listen_port", "1024"),
				),
			},
		},
	})
}
