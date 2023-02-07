package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_ROUTER_BGP_NEIGHBOR_IPV4_NEIGHBOR_RESOURCE = `
resource "thunder_router_bgp_neighbor_ipv4_neighbor" "RouterBgpNeighborTest" {
    as_number = 1
	neighbor_ipv4 = "3.3.3.3"
activate = 1
peer_group_name = "a"
nbr_remote_as = 2
 
}
`

//Acceptance test
func TestAccRouterBgpNeighborIpv4Neighbor_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ROUTER_BGP_NEIGHBOR_IPV4_NEIGHBOR_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_router_bgp_neighbor_ipv4_neighbor.RouterBgpNeighborTest", "neighbor_ipv4", "3.3.3.3"),
					resource.TestCheckResourceAttr("thunder_router_bgp_neighbor_ipv4_neighbor.RouterBgpNeighborTest", "activate", "1"),
					resource.TestCheckResourceAttr("thunder_router_bgp_neighbor_ipv4_neighbor.RouterBgpNeighborTest", "peer_group_name", "a"),
					resource.TestCheckResourceAttr("thunder_router_bgp_neighbor_ipv4_neighbor.RouterBgpNeighborTest", "nbr_remote_as", "2"),
					resource.TestCheckResourceAttr("thunder_router_bgp_neighbor_ipv4_neighbor.RouterBgpNeighborTest", "as_number", "1"),
				),
			},
		},
	})
}
