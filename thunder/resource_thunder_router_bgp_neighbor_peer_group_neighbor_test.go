package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_ROUTER_BGP_NEIGHBOR_PEER_GROUP_NEIGHBOR_RESOURCE = `
resource "thunder_router_bgp_neighbor_peer_group_neighbor" "RouterBgpNeighborTest" {
    as_number = 1
	peer_group = "a"
activate = 1
peer_group_remote_as = 1
peer_group_key = 1
 
}
`

//Acceptance test
func TestAccRouterBgpNeighborPeerGroupNeighbor_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ROUTER_BGP_NEIGHBOR_PEER_GROUP_NEIGHBOR_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_router_bgp_neighbor_peer_group_neighbor.RouterBgpNeighborTest", "peer_group", "a"),
					resource.TestCheckResourceAttr("thunder_router_bgp_neighbor_peer_group_neighbor.RouterBgpNeighborTest", "activate", "1"),
					resource.TestCheckResourceAttr("thunder_router_bgp_neighbor_peer_group_neighbor.RouterBgpNeighborTest", "peer_group_remote_as", "1"),
					resource.TestCheckResourceAttr("thunder_router_bgp_neighbor_peer_group_neighbor.RouterBgpNeighborTest", "peer_group_key", "1"),
					resource.TestCheckResourceAttr("thunder_router_bgp_neighbor_peer_group_neighbor.RouterBgpNeighborTest", "as_number", "1"),
				),
			},
		},
	})
}
