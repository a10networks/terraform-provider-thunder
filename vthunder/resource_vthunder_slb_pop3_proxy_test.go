package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_POP3_PROXY_RESOURCE = `
resource "vthunder_slb_pop3_proxy" "testname" {

	sampling_enable {
	    counters1 = "all"
	}
}
`

//Acceptance test
func TestAccSlbPop3Proxy_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_POP3_PROXY_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_pop3_proxy.testname", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
