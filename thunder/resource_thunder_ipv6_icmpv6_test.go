package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_IPV6_ICMPV6_RESOURCE = `
resource "thunder_ipv6_icmpv6" "testname" {
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
					resource.TestCheckResourceAttr("thunder_ipv6_icmpv6.testname", "redirect", "0"),
					resource.TestCheckResourceAttr("thunder_ipv6_icmpv6.testname", "unreachable", "0"),
				),
			},
		},
	})
}
