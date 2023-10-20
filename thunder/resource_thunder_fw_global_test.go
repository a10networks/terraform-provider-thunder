package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_FW_GLOBAL_RESOURCE = `
resource "thunder_fw_global" "FwTest" {
	disable_ip_fw_sessions = "1" 
}
`

// Acceptance test
func TestAccFwGlobal_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_GLOBAL_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_fw_global.FwTest", "disable_ip_fw_sessions", "1"),
				),
			},
		},
	})
}
