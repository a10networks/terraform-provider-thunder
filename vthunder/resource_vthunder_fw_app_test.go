package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_APP_RESOURCE = `
resource "vthunder_fw_app" "FwTest" {
	sampling_enable {
		counters1 = "all" 
	}
}
`

//Acceptance test
func TestAccFwApp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_APP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_fw_app.FwTest", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
