package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_DNS_CACHE_RESOURCE = `
resource "vthunder_slb_dns_cache" "dns_cache" {
	sampling_enable {
		counters1 = "all"
	}
}

`

//Acceptance test
func TestAccVthunderDNSCache_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_DNS_CACHE_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_dns_cache.dns_cache", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
