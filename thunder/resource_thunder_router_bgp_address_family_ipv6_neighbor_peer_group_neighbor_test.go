package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_ROUTER_BGP_ADDRESS_FAMILY_IPV6_NEIGHBOR_PEER_GROUP_NEIGHBOR_RESOURCE = `
resource "thunder_router_bgp_address_family_ipv6_neighbor_peer_group_neighbor" "RouterBgpAddressFamilyIpv6NeighborTest" {
    peer_group = "ab"
    as_number = 1
activate = 1
allowas_in = 1
 
}
`

//Acceptance test
func TestAccRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ROUTER_BGP_ADDRESS_FAMILY_IPV6_NEIGHBOR_PEER_GROUP_NEIGHBOR_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_router_bgp_address_family_ipv6_neighbor_peer_group_neighbor.RouterBgpAddressFamilyIpv6NeighborTest", "peer_group", "ab"),
					resource.TestCheckResourceAttr("thunder_router_bgp_address_family_ipv6_neighbor_peer_group_neighbor.RouterBgpAddressFamilyIpv6NeighborTest", "activate", "1"),
					resource.TestCheckResourceAttr("thunder_router_bgp_address_family_ipv6_neighbor_peer_group_neighbor.RouterBgpAddressFamilyIpv6NeighborTest", "allowas_in", "1"),
					resource.TestCheckResourceAttr("thunder_router_bgp_address_family_ipv6_neighbor_peer_group_neighbor.RouterBgpAddressFamilyIpv6NeighborTest", "as_number", "1"),
				),
			},
		},
	})
}
