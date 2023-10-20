package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_HEALTH_GATEWAY_RESOURCE = `
resource "thunder_slb_health_gateway" "health_gateway1" {
	sampling_enable {
		counters1 = "all"
	}
}
`

// Acceptance test
func TestAccThunderHealthGateway_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_HEALTH_GATEWAY_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_health_gateway.health_gateway1", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
