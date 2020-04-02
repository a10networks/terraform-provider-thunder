package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_VIRTUAL_PORT_RESOURCE = `
resource "vthunder_slb_template_virtual_port" "testname" {
	name = "testvirtualport"
	user_tag = "test_tag"
	reset_unknown_conn = 0
	ignore_tcp_msl = 0
	rate = 5
	snat_msl = 10
	allow_syn_otherflags = 0
	aflow = 0
	conn_limit = 50
	drop_unknown_conn = 0
	reset_l7_on_failover = 0
	pkt_rate_type = "src-port"
	rate_interval = "100ms"
	snat_port_preserve = 0
	conn_rate_limit_reset = 0
	when_rr_enable = 0
	non_syn_initiation = 0
	conn_limit_reset = 0
	dscp = 50
	pkt_rate_limit_reset = 0
	conn_limit_no_logging = 0
	conn_rate_limit_no_logging = 0
	log_options = "no-logging"
	allow_vip_to_rport_mapping = 0
	pkt_rate_interval = "100ms"
	conn_rate_limit = 40
}
`

//Acceptance test
func TestSlbTemplateVirtualPort_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_VIRTUAL_PORT_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "name", "testvirtualport"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "reset_unknown_conn", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "ignore_tcp_msl", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "rate", "5"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "snat_msl", "10"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "allow_syn_otherflags", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "aflow", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "conn_limit", "50"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "drop_unknown_conn", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "reset_l7_on_failover", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "pkt_rate_type", "src-port"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "rate_interval", "100ms"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "snat_port_preserve", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "conn_rate_limit_reset", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "when_rr_enable", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "non_syn_initiation", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "conn_limit_reset", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "dscp", "50"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "pkt_rate_limit_reset", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "conn_limit_no_logging", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "conn_rate_limit_no_logging", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "log_options", "no-logging"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "allow_vip_to_rport_mapping", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "pkt_rate_interval", "100ms"),
					resource.TestCheckResourceAttr("vthunder_slb_template_virtual_port.testname", "conn_rate_limit", "40"),
				),
			},
		},
	})
}
