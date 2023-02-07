package thunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var TEST_VRRP_COMMON_RESOURCE = `
	resource "thunder_vrrp_common" "vrrp_common" {
		set_id = 1
		device_id = 1
		action = "enable"
}
`

//Acceptance test
func TestAccThunderVrrpCommon_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_VRRP_COMMON_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_vrrp_common.vrrp_common", "set_id", "1"),
					resource.TestCheckResourceAttr("thunder_vrrp_common.vrrp_common", "device_id", "1"),
					resource.TestCheckResourceAttr("thunder_vrrp_common.vrrp_common", "action", "enable"),
				),
			},
		},
	})
}
