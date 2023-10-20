package thunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var TEST_VRRP_SESSION_SYNC_RESOURCE = `
	resource "thunder_vrrp_session_sync" "sync" {
		action="enable"
}
`

// Acceptance test
func TestAccThunderVrrpSessionSync_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_VRRP_SESSION_SYNC_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_vrrp_session_sync.sync", "action", "enable"),
				),
			},
		},
	})
}
