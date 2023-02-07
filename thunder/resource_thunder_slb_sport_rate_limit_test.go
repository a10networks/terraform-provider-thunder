package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_SPORT_RATE_LIMIT_RESOURCE = `
resource "thunder_slb_sport_rate_limit" "sport_rate1" {
	sampling_enable  {
	counters1 = "all"
	}
}
`

//Acceptance test
func TestAccSlbSportRateLimit_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_SPORT_RATE_LIMIT_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_sport_rate_limit.sport_rate1", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
