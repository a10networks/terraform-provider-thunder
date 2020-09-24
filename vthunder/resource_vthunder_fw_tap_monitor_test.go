package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_TAP_MONITOR_RESOURCE = `
resource "vthunder_fw_tap_monitor" "FwTest" {
	status = "enable" 
}
`

//Acceptance test
func TestAccFwTapMonitor_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_TAP_MONITOR_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_fw_tap_monitor.FwTest", "status", "enable"),
				),
			},
		},
	})
}
