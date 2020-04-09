package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_TCP_PROXY_RESOURCE = `
resource "vthunder_slb_template_tcp_proxy" "testname" {
	name = "testtcpproxy"
	user_tag = "test_tag"
	qos = 2
	init_cwnd = 1
	idle_timeout = 1
	fin_timeout = 1
	half_open_idle_timeout = 1
	reno = 1
	early_retransmit = 1
	server_down_action = "FIN"
	timewait = 5
	min_rto = 300
	dynamic_buffer_allocation = 1
	limited_slowstart = 1
	disable_sack = 0
	disable_window_scale = 1
	alive_if_active = 1
	mss = 200
	keepalive_interval = 200
	retransmit_retries = 2
	insert_client_ip = 1
	transmit_buffer = 2
	nagle = 1
	initial_window_size = 1
	keepalive_probes = 3
	psh_flag_optimization = 1
	ack_aggressiveness = "low"
	backend_wscale = 5
	reset_rev = 1
	maxburst = 1
	receive_buffer = 1
	del_session_on_server_down = 1
	reassembly_timeout = 6
	reset_fwd = 1
	disable_tcp_timestamps = 1
	syn_retries = 6
	force_delete_timeout = 2
	reassembly_limit = 5
	invalid_rate_limit = 4
	disable_abc = 1
	half_close_idle_timeout = 100	

}
`

//Acceptance test
func TestAccSlbTemplateTcpProxy_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_TCP_PROXY_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "name", "testtcpproxy"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "qos", "2"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "init_cwnd", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "idle_timeout", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "fin_timeout", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "half_open_idle_timeout", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "reno", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "early_retransmit", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "server_down_action", "FIN"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "timewait", "5"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "min_rto", "300"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "dynamic_buffer_allocation", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "limited_slowstart", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "disable_sack", "0"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "disable_window_scale", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "alive_if_active", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "mss", "200"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "keepalive_interval", "200"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "retransmit_retries", "2"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "insert_client_ip", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "transmit_buffer", "2"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "nagle", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "initial_window_size", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "keepalive_probes", "3"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "psh_flag_optimization", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "ack_aggressiveness", "low"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "backend_wscale", "5"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "reset_rev", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "maxburst", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "receive_buffer", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "del_session_on_server_down", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "reassembly_timeout", "6"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "reset_fwd", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "disable_tcp_timestamps", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "syn_retries", "6"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "force_delete_timeout", "2"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "reassembly_limit", "5"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "invalid_rate_limit", "4"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "disable_abc", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_tcp_proxy.testname", "half_close_idle_timeout", "100"),
				),
			},
		},
	})
}
