package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_CRL_SRCIP_RESOURCE = `
resource "thunder_slb_crl_srcip" "crl_srcip" {
	sampling_enable {
		counters1 = "all"
	}
}
`

//Acceptance test
func TestAccThunderCrlSrcip_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_CRL_SRCIP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_crl_srcip.crl_srcip", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
