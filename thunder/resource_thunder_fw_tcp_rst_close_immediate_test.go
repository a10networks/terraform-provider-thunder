package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_FW_TCP_RST_CLOSE_IMMEDIATE_RESOURCE = `
resource "thunder_fw_tcp_rst_close_immediate" "FwTest" {
	status = "enable" 
}
`

//Acceptance test
func TestAccFwTcpRstCloseImmediate_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_TCP_RST_CLOSE_IMMEDIATE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_fw_tcp_rst_close_immediate.FwTest", "status", "enable"),
				),
			},
		},
	})
}
