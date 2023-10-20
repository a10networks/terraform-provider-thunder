package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_ROUTER_BGP_ADDRESS_FAMILY_IPV6_NETWORK_IPV6_NETWORK_RESOURCE = `
resource "thunder_router_bgp_address_family_ipv6_network_ipv6_network" "RouterBgpAddressFamilyIpv6NetworkTest" {
    network_ipv6 = "2003::1/64"
    as_number = 1
description = "a"
 
}
`

// Acceptance test
func TestAccRouterBgpAddressFamilyIpv6NetworkIpv6Network_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ROUTER_BGP_ADDRESS_FAMILY_IPV6_NETWORK_IPV6_NETWORK_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_router_bgp_address_family_ipv6_network_ipv6_network.RouterBgpAddressFamilyIpv6NetworkTest", "network_ipv6", "2003::1/64"),
					resource.TestCheckResourceAttr("thunder_router_bgp_address_family_ipv6_network_ipv6_network.RouterBgpAddressFamilyIpv6NetworkTest", "description", "a"),
					resource.TestCheckResourceAttr("thunder_router_bgp_address_family_ipv6_network_ipv6_network.RouterBgpAddressFamilyIpv6NetworkTest", "as_number", "1"),
				),
			},
		},
	})
}
