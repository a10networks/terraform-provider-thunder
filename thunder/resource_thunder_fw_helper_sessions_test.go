package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_HELPER_SESSIONS_RESOURCE = `
resource "thunder_fw_helper_sessions" "FwTest" {
	mode = "disable"
}
`

//Acceptance test
func TestAccFwHelperSessions_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_HELPER_SESSIONS_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_fw_helper_sessions.FwTest", "mode", "disable"),
				),
			},
		},
	})
}
