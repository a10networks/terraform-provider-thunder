package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_TEMPLATE_DNS_RESOURCE = `
resource "thunder_slb_template_dns" "template_dns" {
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

// Acceptance test
func TestAccSlbTemplateDNS_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_TEMPLATE_DNS_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_dns.template_dns", "name", "testdns"),
					resource.TestCheckResourceAttr("thunder_slb_template_dns.template_dns", "response_rate_limiting.0.enable_log", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_dns.template_dns", "response_rate_limiting.0.filter_response_rate", "3"),
					resource.TestCheckResourceAttr("thunder_slb_template_dns.template_dns", "response_rate_limiting.0.slip_rate", "4"),
					resource.TestCheckResourceAttr("thunder_slb_template_dns.template_dns", "response_rate_limiting.0.response_rate", "5"),
					resource.TestCheckResourceAttr("thunder_slb_template_dns.template_dns", "response_rate_limiting.0.window", "6"),
					resource.TestCheckResourceAttr("thunder_slb_template_dns.template_dns", "response_rate_limiting.0.action", "log-only"),
				),
			},
		},
	})
}
