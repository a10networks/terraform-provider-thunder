package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_L7SESSION_RESOURCE = `
resource "thunder_slb_l7session" "l7session" {
	sampling_enable {
	    counters1=  "all"
	}
}
`

//Acceptance test
func TestAccSlbL7session_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_L7SESSION_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_l7session.l7session", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
