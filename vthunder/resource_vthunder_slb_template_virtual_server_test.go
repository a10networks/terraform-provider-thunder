package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_VIRTUAL_SERVER_RESOURCE = `
resource "vthunder_slb_template_virtual_server" "testname" {
	name = "testvirtualserver"
	user_tag = "test_tag"
	conn_limit = 10000
	conn_rate_limit_no_logging = 0
	icmp_lockup_period = 10
	conn_limit_reset = 0
	rate_interval = "100ms"
	icmpv6_rate_limit = 50
	subnet_gratuitous_arp = 0
	icmpv6_lockup = 2000
	conn_rate_limit_reset = 1
	tcp_stack_tfo_backoff_time = 60
	tcp_stack_tfo_cookie_time_limit = 60
	conn_limit_no_logging = 0
	icmpv6_lockup_period = 50
	conn_rate_limit = 50
	tcp_stack_tfo_active_conn_limit = 50
	icmp_lockup = 50
	icmp_rate_limit = 50
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
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "conn_limit", 10000),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "conn_rate_limit_no_logging", 0),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "icmp_lockup_period", 10),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "conn_limit_reset", 0),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "rate_interval", "100ms"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "icmpv6_rate_limit", 50),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "subnet_gratuitous_arp", 0),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "icmpv6_lockup", 2000),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "conn_rate_limit_reset", 1),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "tcp_stack_tfo_backoff_time", 60),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "tcp_stack_tfo_cookie_time_limit", 60),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "conn_limit_no_logging", 0),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "icmpv6_lockup_period", 50),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "conn_rate_limit", 50),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "tcp_stack_tfo_active_conn_limit", 50),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "icmp_lockup", 50),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_server.testname", "icmp_rate_limit", 50),
				),
			},
		},
	})
}
