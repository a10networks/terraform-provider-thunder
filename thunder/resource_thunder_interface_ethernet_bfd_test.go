package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_INTERFACE_ETHERNET_BFD_RESOURCE = `
resource "thunder_interface_ethernet_bfd" "ethernetbfd" {
	ifnum=1
	authentication {
	  method = "md5"
	  key_id = 0
	  password = "joker"
	}
  }
`

//Acceptance test
func TestAccInterfaceEthernetBFD_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_INTERFACE_ETHERNET_BFD_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_interface_ethernet_bfd.ethernetbfd", "ifnum", "1"),
					resource.TestCheckResourceAttr("thunder_interface_ethernet_bfd.ethernetbfd", "authentication.0.method", "md5"),
					resource.TestCheckResourceAttr("thunder_interface_ethernet_bfd.ethernetbfd", "authentication.0.key_id", "0"),
					resource.TestCheckResourceAttr("thunder_interface_ethernet_bfd.ethernetbfd", "authentication.0.password", "joker"),
				),
			},
		},
	})
}
