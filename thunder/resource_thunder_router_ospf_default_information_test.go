package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_ROUTER_OSPF_DEFAULT_INFORMATION_RESOURCE = `
resource "thunder_router_ospf_default_information" "RouterOspfTest" {
	originate = 1
always = 1
 
}
`

//Acceptance test
func TestAccRouterOspfDefaultInformation_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ROUTER_OSPF_DEFAULT_INFORMATION_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_router_ospf_default_information.RouterOspfTest", "originate", "1"),
					resource.TestCheckResourceAttr("thunder_router_ospf_default_information.RouterOspfTest", "always", "1"),
				),
			},
		},
	})
}
