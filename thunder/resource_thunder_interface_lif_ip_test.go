package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_INTERFACE_LIF_IP_RESOURCE = `
resource "thunder_interface_lif_ip" "InterfaceLifTest" {
	ifname = 1
	router {  
 isis {  
 	tag =  "a" 
	}
	}
address_list {   
	ipv4_address =  "3.3.3.3" 
	ipv4_netmask =  "255.255.0.0" 
	}
 
}
`

// Acceptance test
func TestAccInterfaceLifIp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_INTERFACE_LIF_IP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_interface_lif_ip.InterfaceLifTest", "ifname", "1"),
					resource.TestCheckResourceAttr("thunder_interface_lif_ip.InterfaceLifTest", "router.0.isis.0.tag", "a"),
					resource.TestCheckResourceAttr("thunder_interface_lif_ip.InterfaceLifTest", "address_list.0.ipv4_address", "3.3.3.3"),
					resource.TestCheckResourceAttr("thunder_interface_lif_ip.InterfaceLifTest", "address_list.0.ipv4_netmask", "255.255.0.0"),
				),
			},
		},
	})
}
