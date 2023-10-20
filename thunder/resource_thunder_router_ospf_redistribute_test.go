package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_ROUTER_OSPF_REDISTRIBUTE_RESOURCE = `
resource "thunder_router_ospf_redistribute" "RouterOspfTest" {
	redist-list {   
	type =  "bgp" 
	metric =  0 
	}
 
}
`

// Acceptance test
func TestAccRouterOspfRedistribute_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ROUTER_OSPF_REDISTRIBUTE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_router_ospf_redistribute.RouterOspfTest", "redist_list.0.type", "bgp"),
					resource.TestCheckResourceAttr("thunder_router_ospf_redistribute.RouterOspfTest", "redist_list.0.metric", "0"),
				),
			},
		},
	})
}
