package vthunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

var TEST_DNS_PRIMARY_RESOURCE = `
	resource "vthunder_dns_primary" "dns_primary" {
		ip_v4_addr = "8.8.8.8"
}
`

//Acceptance test
func TestAccVthunderDnsPrimary_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_DNS_PRIMARY_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_dns_primary.dns_primary", "ip_v4_addr", "8.8.8.8"),
				),
			},
		},
	})
}
