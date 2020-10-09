package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_SERVER_RESOURCE = `
resource "thunder_fw_server" "FwTest" {
	name = "a"
	server_ipv6_addr = "2003::1" 
}
`

//Acceptance test
func TestAccFwServer_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_SERVER_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_fw_server.FwTest", "name", "a"),
					resource.TestCheckResourceAttr("thunder_fw_server.FwTest", "server_ipv6_addr", "2003::1"),
				),
			},
		},
	})
}
