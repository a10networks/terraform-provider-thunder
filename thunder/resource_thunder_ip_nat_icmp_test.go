package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_IP_NAT_ICMP_RESOURCE = `
resource "thunder_ip_nat_icmp" "NatIcmp" {
	respond_to_ping = 0
	always_source_nat_errors = 0 
}
`

// Acceptance test
func TestAccIpNatIcmp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IP_NAT_ICMP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_ip_nat_icmp.NatIcmp", "respond_to_ping", "0"),
					resource.TestCheckResourceAttr("thunder_ip_nat_icmp.NatIcmp", "always_source_nat_errors", "0"),
				),
			},
		},
	})
}
