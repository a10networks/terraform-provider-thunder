package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_VTEP_RESOURCE = `
resource "thunder_overlay_tunnel_vtep" "vtep"{
	encap = "ip-encap"
	sampling_enable {
		counters1 = "all"
	}
	local_ip_address {
		     ip_address = "1.2.3.4"
		   }
	remote_ip_address_list {
	ip_address = "1.4.3.2"
	use_lif {
		lif = "174"
	}
	}
		id1 = 36
	}
`

//Acceptance test
func TestAccThunderOverlayTunnelVtep_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_VTEP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_overlay_tunnel_vtep.vtep", "sampling_enable.0.counters1", "all"),
					resource.TestCheckResourceAttr("thunder_overlay_tunnel_vtep.vtep", "local_ip_address.0.ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr("thunder_overlay_tunnel_vtep.vtep", "remote_ip_address_list.0.ip_address", "1.4.3.2"),
					resource.TestCheckResourceAttr("thunder_overlay_tunnel_vtep.vtep", "remote_ip_address_list.0.ip_address.0.use_lif.0.lif", "174"),
					resource.TestCheckResourceAttr("thunder_overlay_tunnel_vtep.vtep", "encap", "ip-encap"),
					resource.TestCheckResourceAttr("thunder_overlay_tunnel_vtep.vtep", "id1", "36"),
				),
			},
		},
	})
}
