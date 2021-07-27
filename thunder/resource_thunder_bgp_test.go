package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_BGP_RESOURCE = `
resource "thunder_bgp" "Test" {
	extended_asn_cap = 1
 
}
`

//Acceptance test
func TestAccBgp_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_BGP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_bgp.Test", "extended_asn_cap", "1"),
				),
			},
		},
	})
}
