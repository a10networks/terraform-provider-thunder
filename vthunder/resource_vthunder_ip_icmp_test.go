package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_IP_ICMP_RESOURCE = `
resource "vthunder_ip_icmp" "ICMP" {
	redirect = 0
	unreachable = 0 
}
`

//Acceptance test
func TestAccIpIcmp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IP_ICMP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_ip_icmp.ICMP", "redirect", "0"),
					resource.TestCheckResourceAttr("vthunder_ip_icmp.ICMP", "unreachable", "0"),
				),
			},
		},
	})
}
