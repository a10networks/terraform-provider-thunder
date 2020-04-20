package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_SSL_FORWARD_PROXY_RESOURCE = `
resource "vthunder_slb_ssl_forward_proxy" "SSLForwardProxy" {
	sampling_enable {
	    counters1 = "all"
	}
}
`

//Acceptance test
func TestSlbSSLForwardProxy_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_SSL_FORWARD_PROXY_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_ssl_forward_proxy.SSLForwardProxy", "sampling_enable.0.counters1", "all"),
				),
			},
		},
	})
}
