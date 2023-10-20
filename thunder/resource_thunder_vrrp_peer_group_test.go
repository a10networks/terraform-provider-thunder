package thunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var TEST_VRRP_PEER_GROUP_RESOURCE = `
	resource "thunder_vrrp_peer_group" "peer_group" {
		peer{
			ip_peer_address_cfg{
			ip_peer_address = "10.0.2.2"
		}
			ip_peer_address_cfg{
			ip_peer_address = "10.0.2.3"
		}
	}
}
`

// Acceptance test
func TestAccThunderVrrpPeerGroup_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_VRRP_PEER_GROUP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_vrrp_peer_group.peer_group", "peer.0.ip_peer_address_cfg.0.ip_peer_address", "10.0.2.2"),
					resource.TestCheckResourceAttr("thunder_vrrp_peer_group.peer_group", "peer.0.ip_peer_address_cfg.1.ip_peer_address", "10.0.2.3"),
				),
			},
		},
	})
}
