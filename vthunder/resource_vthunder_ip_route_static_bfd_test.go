package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_IP_ROUTE_STATIC_BFD_RESOURCE = `
resource "vthunder_ip_route_static_bfd" "ipStaticBFD" {
	local_ip="3.3.3.3"
	nexthop_ip="3.3.3.3"
}  
`

//Acceptance test
func TestAccIPRouteStaticBfd_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IP_ROUTE_STATIC_BFD_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_ip_route_static_bfd.ipStaticBFD", "local_ip", "3.3.3.3"),
					resource.TestCheckResourceAttr("vthunder_ip_route_static_bfd.ipStaticBFD", "nexthop_ip", "3.3.3.3"),
				),
			},
		},
	})
}
