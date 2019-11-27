package vthunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

var TEST_VRRP_COMMON_RESOURCE = `
	resource "vthunder_vrrp_common" "vrrp_common" {
		set_id = 1
		device_id = 1
		action = "enable"
}
`

//Acceptance test
func TestAccVthunderVrrpCommon_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_VRRP_COMMON_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_vrrp_common.vrrp_common", "set_id", "1"),
					resource.TestCheckResourceAttr("vthunder_vrrp_common.vrrp_common", "device_id", "1"),
					resource.TestCheckResourceAttr("vthunder_vrrp_common.vrrp_common", "action", "enable"),
				),
			},
		},
	})
}
