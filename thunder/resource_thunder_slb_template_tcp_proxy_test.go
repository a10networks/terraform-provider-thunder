package thunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var TEST_SLB_TEMPLATE_TCP_PROXY_RESOURCE = `
resource "thunder_slb_template_tcp_proxy" "tcp_proxy" {
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

// Acceptance test
func TestAccSlbTemplateTcpProxy_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_TCP_PROXY_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "name", "testtcpproxy"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "qos", "2"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "init_cwnd", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "idle_timeout", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "fin_timeout", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "half_open_idle_timeout", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "reno", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "early_retransmit", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "server_down_action", "FIN"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "timewait", "5"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "min_rto", "300"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "dynamic_buffer_allocation", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "limited_slowstart", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "disable_sack", "0"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "disable_window_scale", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "alive_if_active", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "mss", "200"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "keepalive_interval", "200"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "retransmit_retries", "2"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "insert_client_ip", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "transmit_buffer", "2"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "nagle", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "initial_window_size", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "keepalive_probes", "3"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "psh_flag_optimization", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "ack_aggressiveness", "low"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "backend_wscale", "5"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "reset_rev", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "maxburst", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "receive_buffer", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "del_session_on_server_down", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "reassembly_timeout", "6"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "reset_fwd", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "disable_tcp_timestamps", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "syn_retries", "6"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "force_delete_timeout", "2"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "reassembly_limit", "5"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "invalid_rate_limit", "4"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "disable_abc", "1"),
					resource.TestCheckResourceAttr("thunder_slb_template_tcp_proxy.tcp_proxy", "half_close_idle_timeout", "100"),
				),
			},
		},
	})
}
