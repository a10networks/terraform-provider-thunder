package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SYSTEM_VE_MAC_SCHEME_RESOURCE = `
resource "thunder_system_ve_mac_scheme" "SystemTest" {
	ve_mac_scheme_val = "hash-based"
 
}
`

//Acceptance test
func TestAccSystemVeMacScheme_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SYSTEM_VE_MAC_SCHEME_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_system_ve_mac_scheme.SystemTest", "ve_mac_scheme_val", "hash-based"),
				),
			},
		},
	})
}
