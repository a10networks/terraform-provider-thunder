package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_INTERFACE_VE_BFD_RESOURCE = `
resource "vthunder_interface_ve_bfd" "vebfd" {
    ifnum=11
    authentication {
      method = "md5"
      key_id = 0
      password = "joker"
    }
}
`

//Acceptance test
func TestAccInterfaceVeBFD_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_INTERFACE_VE_BFD_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_interface_ve_bfd.vebfd", "ifnum", "11"),
					resource.TestCheckResourceAttr("vthunder_interface_ve_bfd.vebfd", "authentication.0.method", "md5"),
					resource.TestCheckResourceAttr("vthunder_interface_ve_bfd.vebfd", "authentication.0.key_id", "0"),
					resource.TestCheckResourceAttr("vthunder_interface_ve_bfd.vebfd", "authentication.0.password", "joker"),
				),
			},
		},
	})
}
