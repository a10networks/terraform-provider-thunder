---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_tcp_proxy"
sidebar_current: "docs-vthunder-resource-slb-template-tcp-proxy"
description: |-
    Provides details about vthunder slb template tcp-proxy resource for A10
---

# vthunder\_slb\_template\_tcp_proxy

`vthunder_slb_template_tcp_proxy provides details about slb template tcp-proxy
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

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
```

## Argument Reference

* `name` - TCP Proxy Template Name
* `user_tag` - Customized tag
* `qos` - QOS level (number)
* `init_cwnd` - The initial congestion control window size (packets), default is 10 (init-cwnd in packets, default 10)
* `idle_timeout` - Idle Timeout (Interval of 60 seconds), default is 600 (idle timeout in second, default 600)
* `fin_timeout` - FIN timeout (sec), default is disabled (number)
* `half_open_idle_timeout` - TCP Half Open Idle Timeout (sec), default is off (number)
* `reno` - Enable Reno Congestion Control Algorithm
* `early_retransmit` - Configure the Early-Retransmit Algorithm (RFC 5827) (Early-Retransmit is disabled by default)
* `server_down_action` - 'FIN’: FIN Connection; 'RST’: Reset Connection;
* `timewait` - Timewait Threshold (sec), default 5 (number)
* `min_rto` - The minmum retransmission timeout, default is 200ms (number)
* `dynamic_buffer_allocation` - Optimally adjust the transmit and receive buffer sizes of TCP proxy while keeping their sum constant
* `limited_slowstart` - RFC 3742 Limited Slow-Start for TCP with Large Congestion Windows (number)
* `disable_sack` - disable Selective Ack Option
* `disable_window_scale` - disable TCP Window-Scale Option
* `alive_if_active` - keep connection alive if active traffic
* `mss` - Responding MSS to use if client MSS is large, default is off (number)
* `keepalive_interval` - Interval between keepalive probes (sec), default is off (number (seconds))
* `retransmit_retries` - Number of Retries for Retransmit, default is 5
* `insert_client_ip` - Insert client ip into TCP option
* `transmit_buffer` - TCP Transmit Buffer (default 200k) (number default 200000 bytes)
* `nagle` - Enable Nagle Algorithm
* `initial_window_size` - Set the initial window size, default is off (number)
* `keepalive_probes` - Number of keepalive probes sent, default is off
* `psh_flag_optimization` - Enable Optimized PSH Flag Use
* `ack_aggressiveness` - 'low’: Delayed ACK; 'medium’: Delayed ACK, with ACK on each packet with PUSH flag; 'high’: ACK on each packet;
* `backend_wscale` - The TCP window scale used for the server side, default is off (number)
* `reset_rev` - send reset to client if error happens
* `maxburst` - The max packet count sent per transmission event (number)
* `receive_buffer` - TCP Receive Buffer (default 200k) (number default 200000 bytes)
* `del_session_on_server_down` - Delete session if the server/port goes down (either disabled/hm down)
* `reassembly_timeout` - The reassembly timeout, default is 30sec (number)
* `reset_fwd` - send reset to server if error happens
* `disable_tcp_timestamps` - disable TCP Timestamps Option
* `syn_retries` - SYN Retry Numbers, default is 5
* `force_delete_timeout` - The maximum time that a session can stay in the system before being deleted, default is off (number (second))
* `reassembly_limit` - The reassembly queuing limit, default is 25 segments (number)
* `invalid_rate_limit` - Invalid Packet Response Rate Limit (ms), default is 500 (number default 500 challenges)
* `disable_abc` - Appropriate Byte Counting RFC 3465 Disabled, default is enabled (Appropriate Byte Counting (ABC) is enabled by default)
* `half_close_idle_timeout` - TCP Half Close Idle Timeout (sec), default is off (cmd is deprecated, use fin-timeout instead) (number)

