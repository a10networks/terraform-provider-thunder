package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_RATE_LIMIT_LOG_RESOURCE = `
resource "thunder_slb_rate_limit_log" "rate_limit_log" {

	sampling_enable {
		counters1 = "all"
	} 
}
`

//Acceptance test
func TestAccSlbRateLimitLog_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_RATE_LIMIT_LOG_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_rate_limit_log.rate_limit_log", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
