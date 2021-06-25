---
layout: "thunder"
page_title: "thunder: thunder_slb_template_tcp_proxy"
sidebar_current: "docs-thunder-resource-slb-template-tcp-proxy"
description: |-
	Provides details about thunder slb template tcp proxy resource for A10
---

# thunder\_slb\_template\_tcp\_proxy

`thunder_slb_template_tcp_proxy` Provides details about thunder slb template tcp proxy
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_template_tcp_proxy" "Slb_Template_Tcp_Proxy_Test" {

qos = 0
init_cwnd = 0
idle_timeout = 0
fin_timeout = 0
half_open_idle_timeout = 0
reno = 0
down = 0
early_retransmit = 0
server_down_action = "string"
timewait = 0
min_rto = 0
dynamic_buffer_allocation = 0
limited_slowstart = 0
disable_sack = 0
disable_window_scale = 0
alive_if_active = 0
mss = 0
keepalive_interval = 0
retransmit_retries = 0
insert_client_ip = 0
transmit_buffer = 0
nagle = 0
force_delete_timeout_100ms = 0
initial_window_size = 0
keepalive_probes = 0
psh_flag_optimization = 0
ack_aggressiveness = "string"
backend_wscale = 0
disable = 0
reset_rev = 0
maxburst = 0
uuid = "string"
receive_buffer = 0
del_session_on_server_down = 0
name = "string"
reassembly_timeout = 0
reset_fwd = 0
disable_tcp_timestamps = 0
syn_retries = 0
force_delete_timeout = 0
user_tag = "string"
reassembly_limit = 0
invalid_rate_limit = 0
disable_abc = 0
half_close_idle_timeout = 0
 
}
```

## Argument Reference

* `ack_aggressiveness` - ‘low’: Delayed ACK; ‘medium’: Delayed ACK, with ACK on each packet with PUSH flag; ‘high’: ACK on each packet;
* `alive_if_active` - keep connection alive if active traffic
* `backend_wscale` - The TCP window scale used for the server side, default is off (number)
* `del_session_on_server_down` - Delete session if the server/port goes down (either disabled/hm down)
* `disable` - send reset to client when server is disabled
* `disable_abc` - Appropriate Byte Counting RFC 3465 Disabled, default is enabled (Appropriate Byte Counting (ABC) is enabled by default)
* `disable_sack` - disable Selective Ack Option
* `disable_tcp_timestamps` - disable TCP Timestamps Option
* `disable_window_scale` - disable TCP Window-Scale Option
* `down` - send reset to client when server is down
* `dynamic_buffer_allocation` - Optimally adjust the transmit and receive buffer sizes of TCP proxy while keeping their sum constant
* `early_retransmit` - Configure the Early-Retransmit Algorithm (RFC 5827) (Early-Retransmit is disabled by default)
* `fin_timeout` - FIN timeout (sec), default is 5 (number)
* `force_delete_timeout` - The maximum time that a session can stay in the system before being deleted, default is off (number (second))
* `force_delete_timeout_100ms` - The maximum time that a session can stay in the system before being deleted, default is off (number in 100ms)
* `half_close_idle_timeout` - TCP Half Close Idle Timeout (sec), default is off (number)
* `half_open_idle_timeout` - TCP Half Open Idle Timeout (sec), default is off (number)
* `idle_timeout` - Idle Timeout (Interval of 60 seconds), default is 600 (idle timeout in second, default 600)
* `init_cwnd` - The initial congestion control window size (packets), default is 10 (number)
* `initial_window_size` - Set the initial window size, default is off (number)
* `insert_client_ip` - Insert client ip into TCP option
* `invalid_rate_limit` - Invalid Packet Response Rate Limit (ms), default is 500 (number default 500 challenges)
* `keepalive_interval` - Interval between keepalive probes (sec), default is off (number)
* `keepalive_probes` - Number of keepalive probes sent, default is off
* `limited_slowstart` - RFC 3742 Limited Slow-Start for TCP with Large Congestion Windows (number)
* `maxburst` - The max packet count sent per transmission event (number)
* `min_rto` - The minmum retransmission timeout, default is 200ms (number)
* `mss` - Responding MSS to use if client MSS is large, default is off (number)
* `nagle` - Enable Nagle Algorithm
* `name` - TCP Proxy Template Name
* `psh_flag_optimization` - Enable Optimized PSH Flag Use
* `qos` - QOS level (number)
* `reassembly_limit` - The reassembly queuing limit, default is 25 segments (number)
* `reassembly_timeout` - The reassembly timeout, default is 30sec (number)
* `receive_buffer` - TCP Receive Buffer (default 200k) (number)
* `reno` - Enable Reno Congestion Control Algorithm
* `reset_fwd` - send reset to server if error happens
* `reset_rev` - send reset to client if error happens
* `retransmit_retries` - Number of Retries for Retransmit, default is 5
* `server_down_action` - ‘FIN’: FIN Connection; ‘RST’: Reset Connection;
* `syn_retries` - SYN Retry Numbers, default is 5
* `timewait` - Timewait Threshold (sec), default 5 (number)
* `transmit_buffer` - TCP Transmit Buffer (default 200k) (number)
* `user_tag` - Customized tag
* `uuid` - uuid of the object
