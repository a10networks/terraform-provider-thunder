package thunder

import (
    "testing"

    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_ROUTER_BGP_NETWORK_IP_CIDR_RESOURCE = `
resource "thunder_router_bgp_network_ip_cidr" "RouterBgpNetworkTest" {
    as_number = 1
	network_ipv4_cidr = "6.0.0.0/8"
route_map = "a"
 
}
`

//Acceptance test
func TestAccRouterBgp{bgpAsNumber}NetworkIpCidr_create(t *testing.T) {
    resource.Test(t, resource.TestCase{
        Providers: testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: TEST_ROUTER_BGP_NETWORK_IP_CIDR_RESOURCE,
                Check: resource.ComposeTestCheckFunc(
                    resource.TestCheckResourceAttr("thunder_router_bgp_network_ip_cidr.RouterBgpNetworkTest","ip_cidr.0.network_ipv4_cidr","6.0.0.0/8"),
                    resource.TestCheckResourceAttr("thunder_router_bgp_network_ip_cidr.RouterBgpNetworkTest","ip_cidr.0.route_map","a"),
                    resource.TestCheckResourceAttr("thunder_router_bgp_network_ip_cidr.RouterBgpNetworkTest","as_number","1"),
                ),
            },
        },
    })
}
