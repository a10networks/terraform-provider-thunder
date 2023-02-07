package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_ROUTER_BGP_NEIGHBOR_RESOURCE = `
resource "thunder_router_bgp_neighbor" "RouterBgpTest" {
	peer-group-neighbor-list {   
	peer_group =  "a" 
	peer_group_key =  1 
	}
peer_group_key = 1
 
}
`

//Acceptance test
func TestAccRouterBgpNeighbor_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ROUTER_BGP_NEIGHBOR_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_router_bgp_neighbor.RouterBgpTest", "peer_group_neighbor_list.0.peer_group", "a"),
					resource.TestCheckResourceAttr("thunder_router_bgp_neighbor.RouterBgpTest", "peer_group_neighbor_list.0.peer_group_key", "1"),
					resource.TestCheckResourceAttr("thunder_router_bgp_neighbor.RouterBgpTest", "peer_group_key", "1"),
				),
			},
		},
	})
}
