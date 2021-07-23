package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_POP3_PROXY_RESOURCE = `
resource "thunder_slb_pop3_proxy" "pop3_proxy1" {

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
					resource.TestCheckResourceAttr("thunder_slb_pop3_proxy.pop3_proxy1", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
