package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_INTERFACE_LIF_RESOURCE = `
resource "thunder_interface_lif" "InterfaceTest" {
	ifname = 34
user_tag = "mylif"
 
}
`

//Acceptance test
func TestAccInterfaceLif_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_INTERFACE_LIF_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_interface_lif.InterfaceTest", "ifname", "34"),
					resource.TestCheckResourceAttr("thunder_interface_lif.InterfaceTest", "user_tag", "mylif"),
				),
			},
		},
	})
}
