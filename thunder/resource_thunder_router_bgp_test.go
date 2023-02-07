package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_ROUTER_BGP_RESOURCE = `
resource "thunder_router_bgp" "RouterTest" {
	as_number = 1
user_tag = "a"
 
}
`

//Acceptance test
func TestAccRouterBgp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ROUTER_BGP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_router_bgp.RouterTest", "as_number", "1"),
					resource.TestCheckResourceAttr("thunder_router_bgp.RouterTest", "user_tag", "a"),
				),
			},
		},
	})
}
