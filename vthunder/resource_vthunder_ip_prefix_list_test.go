package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_IP_PREFIX_LIST_RESOURCE = `
resource "vthunder_ip_prefix_list" "IpPrefixList" {
	name = "testprefixlist"
	rules {
        any= 1
        action = "deny"
        seq = 1
        le = 32
        ipaddr = "23.0.0.0/8"
        description = "string"
        ge = 24
	}
	uuid = "string"
}
`

//Acceptance test
func TestAccIpPrefixList_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_IP_PREFIX_LIST_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_ip_prefix_list.IpPrefixList", "name", "testprefixlist"),
					resource.TestCheckResourceAttr("vthunder_ip_prefix_list.IpPrefixList", "rules.0.any", "1"),
					resource.TestCheckResourceAttr("vthunder_ip_prefix_list.IpPrefixList", "rules.0.action", "deny"),
					resource.TestCheckResourceAttr("vthunder_ip_prefix_list.IpPrefixList", "rules.0.seq", "1"),
					resource.TestCheckResourceAttr("vthunder_ip_prefix_list.IpPrefixList", "rules.0.le", "32"),
					resource.TestCheckResourceAttr("vthunder_ip_prefix_list.IpPrefixList", "rules.0.ipaddr", "23.0.0.0/8"),
					resource.TestCheckResourceAttr("vthunder_ip_prefix_list.IpPrefixList", "rules.0.description", "string"),
					resource.TestCheckResourceAttr("vthunder_ip_prefix_list.IpPrefixList", "rules.0.ge", "24"),
				),
			},
		},
	})
}
