package vthunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)


var TEST_VTEP_RESOURCE = `
resource "vthunder_overlay_tunnel_vtep" "vtep"{
source_ip_address{
ip_address="1.2.3.4"
vni_list{
segment=1
partition="part7"
lif=1
}
}
encap="nvgre"
host_list{
destination_vtep="1.4.3.2"
ip_addr="1.4.3.5"
overlay_mac_addr="00:1B:44:11:3A:B7"
vni=1
}
id2=3
destination_ip_address_list{
ip_address="2.3.4.5"
vni_list{
segment=1
}
encap="nvgre"
}
}
`

//Acceptance test
func TestAccVthunderOverlayTunnelVtep_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_VTEP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_vtep.vtep", "source_ip_address.0.ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_vtep.vtep", "source_ip_address.0.vni_list.0.segment", "1"),
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_vtep.vtep", "source_ip_address.0.vni_list.0.partition", "part7"),
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_vtep.vtep", "source_ip_address.0.vni_list.0.lif", "1"),
					
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_vtep.vtep", "encap", "nvgre"),
					
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_vtep.vtep", "host_list.0.destination_vtep", "1.4.3.2"),
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_vtep.vtep", "host_list.0.ip_addr", "1.4.3.5"),
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_vtep.vtep", "host_list.0.overlay_mac_addr", "00:1B:44:11:3A:B7"),
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_vtep.vtep", "host_list.0.vni", "1"),
					
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_vtep.vtep", "id2", "3"),
					
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_vtep.vtep", "destination_ip_address_list.0.ip_address", "2.3.4.5"),
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_vtep.vtep", "destination_ip_address_list.0.vni_list.0.segment", "1"),
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_vtep.vtep", "destination_ip_address_list.0.encap", "nvgre"),
					
				),
			},
		},
	})
}
