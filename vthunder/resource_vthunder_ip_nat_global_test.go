package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_IP_NAT_GLOBAL_RESOURCE = `
resource "vthunder_ip_nat_global" "IpNatGlobal" {
	reset_idle_tcp_conn = 0
}
`

//Acceptance test
func TestAccIpNatGlobal_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IP_NAT_GLOBAL_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_ip_nat_global.IpNatGlobal", "reset_idle_tcp_conn", "0"),
				),
			},
		},
	})
}
