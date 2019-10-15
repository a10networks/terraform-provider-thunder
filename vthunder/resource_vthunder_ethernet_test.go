package vthunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

var TEST_ETH_RESOURCE = `
	resource "vthunder_ethernet" "eth" {
	ethernet_list{	
	ifnum=1
		ip{
			address_list={
			ipv4_address="10.0.2.7"
			ipv4_netmask="255.255.255.0"
		    }
		  }
	action="enable"
	}
	ethernet_list{	
	ifnum=2
		ip{
			address_list={
			ipv4_address="10.0.3.7"
			ipv4_netmask="255.255.255.0"
			}
		  }
	action="enable"
	}
	
}
`

//Acceptance test
func TestAccVthunderEthernet_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_ETH_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_ethernet.eth", "ethernet_list.0.ifnum", "1"),
					resource.TestCheckResourceAttr("vthunder_ethernet.eth", "ethernet_list.0.ip.0.address_list.0.ipv4_address", "10.0.2.7"),
					resource.TestCheckResourceAttr("vthunder_ethernet.eth", "ethernet_list.0.ip.0.address_list.0.ipv4_netmask", "255.255.255.0"),
					resource.TestCheckResourceAttr("vthunder_ethernet.eth", "ethernet_list.0.action", "enable"),
					resource.TestCheckResourceAttr("vthunder_ethernet.eth", "ethernet_list.1.ifnum", "2"),
					resource.TestCheckResourceAttr("vthunder_ethernet.eth", "ethernet_list.1.ip.0.address_list.0.ipv4_address", "10.0.3.7"),
					resource.TestCheckResourceAttr("vthunder_ethernet.eth", "ethernet_list.1.ip.0.address_list.0.ipv4_netmask", "255.255.255.0"),
					resource.TestCheckResourceAttr("vthunder_ethernet.eth", "ethernet_list.1.action", "enable"),
				),
			},
		},
	})
}
