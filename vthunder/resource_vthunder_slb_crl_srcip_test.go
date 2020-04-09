package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_CRL_SRCIP_RESOURCE = `
resource "vthunder_slb_crl_srcip" "testname" {
	sampling_enable {
		counters1 = "all"
	}
}
`

//Acceptance test
func TestAccVthunderCrlSrcip_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_CRL_SRCIP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_crl_srcip.testname", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
