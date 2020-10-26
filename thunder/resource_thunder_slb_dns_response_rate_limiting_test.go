package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_DNS_RESPONSE_RATE_LIMITING_RESOURCE = `
resource "thunder_slb_dns_response_rate_limiting" "response_rate" {
	sampling_enable {
		counters1 = "all"
	}
}
`

//Acceptance test
func TestAccThunderSlbDNSResponseRateLimiting_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_DNS_RESPONSE_RATE_LIMITING_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_dns_response_rate_limiting.response_rate", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
