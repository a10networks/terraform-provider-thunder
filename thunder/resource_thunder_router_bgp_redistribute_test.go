package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_ROUTER_BGP_REDISTRIBUTE_RESOURCE = `
resource "thunder_router_bgp_redistribute" "RouterBgpTest" {
	connected_cfg {  
 	connected =  1 
	route_map =  "a" 
	}
 
}
`

//Acceptance test
func TestAccRouterBgpRedistribute_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ROUTER_BGP_REDISTRIBUTE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_router_bgp_redistribute.RouterBgpTest", "connected_cfg.0.connected", "1"),
					resource.TestCheckResourceAttr("thunder_router_bgp_redistribute.RouterBgpTest", "connected_cfg.0.route_map", "a"),
				),
			},
		},
	})
}
