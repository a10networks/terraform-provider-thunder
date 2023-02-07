package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_ROUTER_BGP_ADDRESS_FAMILY_RESOURCE = `
resource "thunder_router_bgp_address_family" "RouterBgpTest" {
	ipv6 {  
 network {  
 synchronization {  
 	network_synchronization =  1 
	}
	}
bgp {  
 	dampening =  1 
	dampening_half =  1 
	dampening_start_reuse =  1 
	dampening_max_supress =  1 
	dampening_start_supress =  1 
	}
	}
 
}
`

//Acceptance test
func TestAccRouterBgpAddressFamily_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ROUTER_BGP_ADDRESS_FAMILY_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_router_bgp_address_family.RouterBgpTest", "ipv6.0.network.0.synchronization.0.network_synchronization", "1"),
					resource.TestCheckResourceAttr("thunder_router_bgp_address_family.RouterBgpTest", "ipv6.0.bgp.0.dampening", "1"),
					resource.TestCheckResourceAttr("thunder_router_bgp_address_family.RouterBgpTest", "bgp.0.dampening_half", "1"),
					resource.TestCheckResourceAttr("thunder_router_bgp_address_family.RouterBgpTest", "bgp.0.dampening_start_reuse", "1"),
					resource.TestCheckResourceAttr("thunder_router_bgp_address_family.RouterBgpTest", "bgp.0.dampening_max_supress", "1"),
					resource.TestCheckResourceAttr("thunder_router_bgp_address_family.RouterBgpTest", "bgp.0.dampening_start_supress", "1"),
				),
			},
		},
	})
}
