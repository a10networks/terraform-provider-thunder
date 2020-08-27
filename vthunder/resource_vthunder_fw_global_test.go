package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_FW_GLOBAL_RESOURCE = `
{'resource': {'vthunder_fw_global': {'global': {'disable_ip_fw_sessions': 1}}}}
`

//Acceptance test
func TestAccFwGlobal_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_FW_GLOBAL_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_fw_global.FwTest", "global", "{'disable-ip-fw-sessions': 1}"),
				),
			},
		},
	})
}
