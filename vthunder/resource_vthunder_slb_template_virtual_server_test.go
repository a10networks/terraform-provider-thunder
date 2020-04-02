package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_VIRTUAL_SERVER_RESOURCE = `
resource "vthunder_slb_template_virtual_server" "testname" {
	name = "testvirtualserver"
	user_tag = "test_tag"
	conn_limit = 1
	conn_limit_reset = 1
	icmpv6_rate_limit = 1
	subnet_gratuitous_arp = 1
	tcp_stack_tfo_backoff_time = 1 
}
`

//Acceptance test
func TestSlbTemplateVirtualServer_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_VIRTUAL_SERVER_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "name", "testvirtualserver"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "conn_limit", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "conn_limit_reset", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "icmpv6_rate_limit", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "subnet_gratuitous_arp", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "tcp_stack_tfo_backoff_time", "1"),
				),
			},
		},
	})
}
