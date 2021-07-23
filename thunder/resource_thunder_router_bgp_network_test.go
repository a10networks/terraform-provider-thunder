package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_ROUTER_BGP_NETWORK_RESOURCE = `
resource "thunder_router_bgp_network" "RouterBgpTest" {
	synchronization {  
 	network_synchronization =  1 
	}
 
}
`

//Acceptance test
func TestAccRouterBgpNetwork_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ROUTER_BGP_NETWORK_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_router_bgp_network.RouterBgpTest", "synchronization.0.network_synchronization", "1"),
				),
			},
		},
	})
}
