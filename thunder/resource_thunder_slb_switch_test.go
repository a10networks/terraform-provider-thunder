package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_SWITCH_RESOURCE = `
resource "thunder_slb_switch" "slb_switch1" {
	sampling_enable  {
	   counters1 = "all"
	} 
}
`

// Acceptance test
func TestAccSlbSwitch_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_SWITCH_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_switch.slb_switch1", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
