package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_ROUTER_OSPF_RESOURCE = `
resource "thunder_router_ospf" "RouterTest" {
	process_id = 1
user_tag = "a"
 
}
`

//Acceptance test
func TestAccRouterOspf_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ROUTER_OSPF_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_router_ospf.RouterTest", "process_id", "1"),
					resource.TestCheckResourceAttr("thunder_router_ospf.RouterTest", "user_tag", "a"),
				),
			},
		},
	})
}
