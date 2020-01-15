package vthunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

var TEST_OVERLAY_TUNNEL_OPTIONS_RESOURCE = `
	resource "vthunder_overlay_tunnel_options" "options" {
		tcp_mss_adjust_disable=1
		nvgre_disable_flow_id=1
		ip_dscp_preserve=1
		vxlan_dest_port=100
		uuid="uuid1"
		gateway_mac="00:0a:95:9d:68:16"
		nvgre_key_mode_lower24=1		
}
`

//Acceptance test
func TestAccVthunderOverlayTunnelOptions_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_OVERLAY_TUNNEL_OPTIONS_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_options.options", "tcp_mss_adjust_disable", "1"),
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_options.options", "nvgre_disable_flow_id", "1"),
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_options.options", "ip_dscp_preserve", "1"),
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_options.options", "vxlan_dest_port", "100"),
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_options.options", "uuid", "uuid1"),
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_options.options", "gateway_mac", "00:0a:95:9d:68:16"),
					resource.TestCheckResourceAttr("vthunder_overlay_tunnel_options.options", "nvgre_key_mode_lower24", "1"),
				),
			},
		},
	})
}
