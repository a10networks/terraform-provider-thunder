package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_FW_APP_RESOURCE = `
resource "thunder_fw_app" "FwTest" {
	sampling_enable {
		counters1 = "all" 
	}
}
`

// Acceptance test
func TestAccFwApp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_APP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_fw_app.FwTest", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
