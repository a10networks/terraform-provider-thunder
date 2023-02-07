package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_INTERFACE_MANAGEMENT_RESOURCE = `
resource "thunder_interface_management" "InterfaceManagement" {
	lldp {
    enable_cfg {
      rt_enable = 1
      rx = 1
      tx = 1
    }
  }
}
`

//Acceptance test
func TestAccInterfaceManagement_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_INTERFACE_MANAGEMENT_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_interface_management.InterfaceManagement", "lldp.0.enable_cfg.0.rt_enable", "1"),
					resource.TestCheckResourceAttr("thunder_interface_management.InterfaceManagement", "lldp.0.enable_cfg.0.rx", "1"),
					resource.TestCheckResourceAttr("thunder_interface_management.InterfaceManagement", "lldp.0.enable_cfg.0.tx", "1"),
				),
			},
		},
	})
}
