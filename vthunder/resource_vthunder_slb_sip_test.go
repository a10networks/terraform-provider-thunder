package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_SIP_RESOURCE = `
resource "vthunder_slb_sip" "testname" {

	sampling_enable {
	    counters1 = "all"
	}
}
`

//Acceptance test
func TestSlbSip_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_SIP_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_sip.testname", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
