package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_IPV6_NAT_ICMPV6_RESOURCE = `
resource "thunder_ipv6_nat_icmpv6" "natIcmpV6" {
	respond_to_ping = 0 
}
`

//Acceptance test
func TestAccIpv6NatIcmpv6_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IPV6_NAT_ICMPV6_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_ipv6_nat_icmpv6.natIcmpV6", "respond_to_ping", "0"),
				),
			},
		},
	})
}
