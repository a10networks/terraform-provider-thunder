package thunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var TEST_OVERLAY_TUNNEL_OPTIONS_RESOURCE = `
	resource "thunder_overlay_tunnel_options" "options" {
		tcp_mss_adjust_disable=1
		nvgre_disable_flow_id=1
		ip_dscp_preserve=1
		vxlan_dest_port=100
		uuid="uuid1"
		gateway_mac="00:0a:95:9d:68:16"
		nvgre_key_mode_lower24=1		
}
`

// Acceptance test
func TestAccThunderOverlayTunnelOptions_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_OVERLAY_TUNNEL_OPTIONS_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_overlay_tunnel_options.options", "tcp_mss_adjust_disable", "1"),
					resource.TestCheckResourceAttr("thunder_overlay_tunnel_options.options", "nvgre_disable_flow_id", "1"),
					resource.TestCheckResourceAttr("thunder_overlay_tunnel_options.options", "ip_dscp_preserve", "1"),
					resource.TestCheckResourceAttr("thunder_overlay_tunnel_options.options", "vxlan_dest_port", "100"),
					resource.TestCheckResourceAttr("thunder_overlay_tunnel_options.options", "uuid", "uuid1"),
					resource.TestCheckResourceAttr("thunder_overlay_tunnel_options.options", "gateway_mac", "00:0a:95:9d:68:16"),
					resource.TestCheckResourceAttr("thunder_overlay_tunnel_options.options", "nvgre_key_mode_lower24", "1"),
				),
			},
		},
	})
}
