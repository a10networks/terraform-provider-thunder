package thunder

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var TEST_DNS_PRIMARY_RESOURCE = `
	resource "thunder_dns_primary" "dns_primary" {
		ip_v4_addr = "8.8.8.8"
}
`

//Acceptance test
func TestAccThunderDnsPrimary_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_DNS_PRIMARY_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_dns_primary.dns_primary", "ip_v4_addr", "8.8.8.8"),
				),
			},
		},
	})
}
