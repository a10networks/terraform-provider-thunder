package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_ROUTE_MAP_RESOURCE = `
resource "thunder_route_map" "routemapetest" {
	action = "permit"
    tag= "a"
    sequence= 1
 
}
`

//Acceptance test
func TestAccRouteMap_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ROUTE_MAP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_route_map.routemapetest", "action", "permit"),
					resource.TestCheckResourceAttr("thunder_route_map.routemapetest", "tag", "a"),
					resource.TestCheckResourceAttr("thunder_route_map.routemapetest", "sequence", "1"),
				),
			},
		},
	})
}
