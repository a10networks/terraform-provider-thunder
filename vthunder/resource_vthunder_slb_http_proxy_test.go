package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_HTTP_PROXY_RESOURCE = `
resource "vthunder_slb_http_proxy" "http_proxy1" {
	sampling_enable {
		counters1 = "all"
		counters2 = "req_sz_8k"
	}
}

`

//Acceptance test
func TestAccVthunderHttpProxy_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_HTTP_PROXY_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_http_proxy.http_proxy1", "sampling_enable.0.counters1", "all"),
					resource.TestCheckResourceAttr("vthunder_slb_http_proxy.http_proxy1", "sampling_enable.0.counters2", "req_sz_8k"),
				),
			},
		},
	})
}
