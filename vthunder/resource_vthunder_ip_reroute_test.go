package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_IP_REROUTE_RESOURCE = `
resource "vthunder_ip_reroute" "reRoute" {
	suppress_protocols  {
        static = 0
        ospf = 0
        connected = 0
        ibgp = 0
        isis = 0
        rip = 0
        ebgp = 0
	}
}
`

//Acceptance test
func TestAccIpReroute_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IP_REROUTE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_ip_reroute.reRoute", "suppress_protocols.0.static", "0"),
					resource.TestCheckResourceAttr("vthunder_ip_reroute.reRoute", "suppress_protocols.0.ospf", "0"),
					resource.TestCheckResourceAttr("vthunder_ip_reroute.reRoute", "suppress_protocols.0.connected", "0"),
					resource.TestCheckResourceAttr("vthunder_ip_reroute.reRoute", "suppress_protocols.0.ibgp", "0"),
					resource.TestCheckResourceAttr("vthunder_ip_reroute.reRoute", "suppress_protocols.0.isis", "0"),
					resource.TestCheckResourceAttr("vthunder_ip_reroute.reRoute", "suppress_protocols.0.rip", "0"),
					resource.TestCheckResourceAttr("vthunder_ip_reroute.reRoute", "suppress_protocols.0.ebgp", "0"),
				),
			},
		},
	})
}
