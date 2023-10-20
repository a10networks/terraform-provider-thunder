package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_FW_TCP_WINDOW_CHECK_RESOURCE = `
resource "thunder_fw_tcp_window_check" "FwTest" {
        status = "enable" 
}
`

// Acceptance test
func TestAccFwTcpWindowCheck_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_TCP_WINDOW_CHECK_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_fw_tcp_window_check.FwTest", "status", "enable"),
				),
			},
		},
	})
}
