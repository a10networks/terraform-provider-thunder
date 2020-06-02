package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_IP_DNS_Secondary_RESOURCE = `
resource "vthunder_ip_dns_secondary" "dnssecondary" {
    ip_v4_addr = "10.10.10.3"
}
`

//Acceptance test
func TestAccIpDnsSecondary_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IP_DNS_Secondary_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_ip_dns_secondary.dnssecondary", "ip_v4_addr", "10.10.10.3"),
				),
			},
		},
	})
}
