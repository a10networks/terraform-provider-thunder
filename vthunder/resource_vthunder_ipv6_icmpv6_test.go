package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_IPV6_ICMPV6_RESOURCE = `
resource "vthunder_ipv6_icmpv6" "testname" {
    redirect = 0
	unreachable = 0 
}
`

//Acceptance test
func TestAccIpv6Icmpv6_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IPV6_ICMPV6_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_ipv6_icmpv6.testname", "redirect", "0"),
					resource.TestCheckResourceAttr("vthunder_ipv6_icmpv6.testname", "unreachable", "0"),
				),
			},
		},
	})
}
