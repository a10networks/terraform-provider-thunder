package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_IP_NAT_ICMP_RESOURCE = `
resource "vthunder_ip_nat_icmp" "NatIcmp" {
	respond_to_ping = 0
	always_source_nat_errors = 0 
}
`

//Acceptance test
func TestAccIpNatIcmp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IP_NAT_ICMP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_ip_nat_icmp.NatIcmp", "respond_to_ping", "0"),
					resource.TestCheckResourceAttr("vthunder_ip_nat_icmp.NatIcmp", "always_source_nat_errors", "0"),
				),
			},
		},
	})
}
