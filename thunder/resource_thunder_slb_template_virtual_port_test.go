package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TEMPLATE_VIRTUAL_PORT_RESOURCE = `
resource "thunder_slb_template_virtual_port" "virtual_port" {
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
func TestAccSlbTemplateVirtualPort_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_VIRTUAL_PORT_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "name", "testvirtualport"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "reset_unknown_conn", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "ignore_tcp_msl", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "rate", "5"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "snat_msl", "10"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "allow_syn_otherflags", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "aflow", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "conn_limit", "50"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "drop_unknown_conn", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "reset_l7_on_failover", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "pkt_rate_type", "src-port"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "rate_interval", "100ms"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "snat_port_preserve", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "conn_rate_limit_reset", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "when_rr_enable", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "non_syn_initiation", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "conn_limit_reset", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "dscp", "50"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "pkt_rate_limit_reset", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "conn_limit_no_logging", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "conn_rate_limit_no_logging", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "log_options", "no-logging"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "allow_vip_to_rport_mapping", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "pkt_rate_interval", "100ms"),
					resource.TestCheckResourceAttr("thunder_slb_template_virtual_port.virtual_port", "conn_rate_limit", "40"),
				),
			},
		},
	})
}
