package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_INTERFACE_ETHERNET_LLDP_RESOURCE = `
resource "vthunder_interface_ethernet_lldp" "ethernetlldp" {
	ifnum=1
	enable_cfg {
	  rt_enable = 1
	  rx = 1
	  tx = 1
	}
  }
`

//Acceptance test
func TestAccInterfaceEthernetLLDP_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_INTERFACE_ETHERNET_LLDP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_interface_ethernet_lldp.ethernetlldp", "ifnum", "1"),
					resource.TestCheckResourceAttr("vthunder_interface_ethernet_lldp.ethernetlldp", "enable_cfg.0.rt_enable", "1"),
					resource.TestCheckResourceAttr("vthunder_interface_ethernet_lldp.ethernetlldp", "enable_cfg.0.rx", "1"),
					resource.TestCheckResourceAttr("vthunder_interface_ethernet_lldp.ethernetlldp", "enable_cfg.0.tx", "1")),
			},
		},
	})
}
