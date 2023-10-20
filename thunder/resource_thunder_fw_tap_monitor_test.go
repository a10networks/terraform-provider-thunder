package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_FW_TAP_MONITOR_RESOURCE = `
resource "thunder_fw_tap_monitor" "FwTest" {
	status = "enable" 
}
`

// Acceptance test
func TestAccFwTapMonitor_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_TAP_MONITOR_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_fw_tap_monitor.FwTest", "status", "enable"),
				),
			},
		},
	})
}
