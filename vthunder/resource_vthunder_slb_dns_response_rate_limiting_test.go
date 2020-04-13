package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_DNS_RESPONSE_RATE_LIMITING_RESOURCE = `
resource "vthunder_slb_dns_response_rate_limiting" "testname" {
	sampling_enable {
		counters1 = "all"
	}
}
`

//Acceptance test
func TestAccVthunderSlbDNSResponseRateLimiting_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_DNS_RESPONSE_RATE_LIMITING_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_dns_response_rate_limiting.testname", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
