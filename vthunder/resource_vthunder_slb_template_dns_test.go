package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_TEMPLATE_DNS_RESOURCE = `
resource "vthunder_slb_template_dns" "testname" {
	name = "testdns"
	response_rate_limiting {
        enable_log=1
        filter_response_rate=3
        slip_rate=4
        response_rate=5
        window=6
        action="log-only"
    }
}
`

//Acceptance test
func TestAccSlbTemplateDNS_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TEMPLATE_DNS_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_dns.testname", "name", "testdns"),
					resource.TestCheckResourceAttr("vthunder_slb_template_dns.testname", "response_rate_limiting.0.enable_log", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_dns.testname", "response_rate_limiting.0.filter_response_rate", "3"),
					resource.TestCheckResourceAttr("vthunder_slb_template_dns.testname", "response_rate_limiting.0.slip_rate", "4"),
					resource.TestCheckResourceAttr("vthunder_slb_template_dns.testname", "response_rate_limiting.0.response_rate", "5"),
					resource.TestCheckResourceAttr("vthunder_slb_template_dns.testname", "response_rate_limiting.0.window", "6"),
					resource.TestCheckResourceAttr("vthunder_slb_template_dns.testname", "response_rate_limiting.0.action", "log-only"),
				),
			},
		},
	})
}
