package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_IP_DNS_SUFFIX_RESOURCE = `
resource "thunder_ip_dns_suffix" "testname" {
	domain_name = "domainName"
}
`

// Acceptance test
func TestAccIpDnsSuffix_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IP_DNS_SUFFIX_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_ip_dns_suffix.testname", "domain_name", "domainName"),
				),
			},
		},
	})
}
