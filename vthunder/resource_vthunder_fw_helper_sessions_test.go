package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_HELPER_SESSIONS_RESOURCE = `
{'resource': {'vthunder_fw_helper_sessions': {'helper_sessions': {'mode': 'disable'}}}}
`

//Acceptance test
func TestAccFwHelperSessions_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_HELPER_SESSIONS_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_fw_helper_sessions.FwTest", "helper_sessions", "{'mode': 'disable'}"),
				),
			},
		},
	})
}
