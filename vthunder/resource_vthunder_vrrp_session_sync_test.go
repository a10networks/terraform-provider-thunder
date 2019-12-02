package vthunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

var TEST_VRRP_SESSION_SYNC_RESOURCE = `
	resource "vthunder_vrrp_session_sync" "sync" {
		action="enable"
}
`

//Acceptance test
func TestAccVthunderVrrpSessionSync_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_VRRP_SESSION_SYNC_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_vrrp_session_sync.sync", "action", "enable"),
				),
			},
		},
	})
}
