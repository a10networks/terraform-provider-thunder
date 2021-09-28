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
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_template_tcp_proxy" "resourceSlbTemplateTcpProxyTest" {
	name = "string"
ack_aggressiveness = "string"
backend_wscale = 0
dynamic_buffer_allocation = 0
fin_timeout = 0
force_delete_timeout = 0
force_delete_timeout_100ms = 0
alive_if_active = 0
idle_timeout = 0
server_down_action = "string"
half_open_idle_timeout = 0
half_close_idle_timeout = 0
init_cwnd = 0
initial_window_size = 0
keepalive_interval = 0
keepalive_probes = 0
mss = 0
psh_flag_optimization = 0
nagle = 0
qos = 0
receive_buffer = 0
reno = 0
transmit_buffer = 0
reset_fwd = 0
reset_rev = 0
disable = 0
down = 0
del_session_on_server_down = 0
retransmit_retries = 0
insert_client_ip = 0
syn_retries = 0
timewait = 0
disable_tcp_timestamps = 0
disable_window_scale = 0
disable_sack = 0
invalid_rate_limit = 0
disable_abc = 0
reassembly_timeout = 0
reassembly_limit = 0
min_rto = 0
limited_slowstart = 0
early_retransmit = 0
maxburst = 0
proxy_header {  
 	proxy_header_action =  "string" 
	version =  "string" 
	}
uuid = "string"
user_tag = "string"
 
}

```

## Argument Reference

* `tcp-proxy` - TCP Proxy
* `name` - TCP Proxy Template Name
* `ack-aggressiveness` - 'low': Delayed ACK; 'medium': Delayed ACK, with ACK on each packet with PUSH flag; 'high': ACK on each packet;
* `backend-wscale` - The TCP window scale used for the server side, default is off (number)
* `dynamic-buffer-allocation` - Optimally adjust the transmit and receive buffer sizes of TCP proxy while keeping their sum constant
* `fin-timeout` - FIN timeout (sec), default is disabled (number)
* `force-delete-timeout` - The maximum time that a session can stay in the system before being deleted, default is off (number (second))
* `force-delete-timeout-100ms` - The maximum time that a session can stay in the system before being deleted, default is off (number in 100ms)
* `alive-if-active` - keep connection alive if active traffic
* `idle-timeout` - Idle Timeout (Interval of 60 seconds), default is 600 (idle timeout in second, default 600)
* `server-down-action` - 'FIN': FIN Connection; 'RST': Reset Connection;
* `half-open-idle-timeout` - TCP Half Open Idle Timeout (sec), default is off (number)
* `half-close-idle-timeout` - TCP Half Close Idle Timeout (sec), default is off (cmd is deprecated, use fin-timeout instead) (number)
* `init-cwnd` - The initial congestion control window size (packets), default is 10 (init-cwnd in packets, default 10)
* `initial-window-size` - Set the initial window size, default is off (number)
* `keepalive-interval` - Interval between keepalive probes (sec), default is off (number (seconds))
* `keepalive-probes` - Number of keepalive probes sent, default is off
* `mss` - Responding MSS to use if client MSS is large, default is off (number)
* `psh-flag-optimization` - Enable Optimized PSH Flag Use
* `nagle` - Enable Nagle Algorithm
* `qos` - QOS level (number)
* `receive-buffer` - TCP Receive Buffer (default 200k) (number default 200000 bytes)
* `reno` - Enable Reno Congestion Control Algorithm
* `transmit-buffer` - TCP Transmit Buffer (default 200k) (number default 200000 bytes)
* `reset-fwd` - send reset to server if error happens
* `reset-rev` - send reset to client if error happens
* `disable` - send reset to client when server is disabled
* `down` - send reset to client when server is down
* `del-session-on-server-down` - Delete session if the server/port goes down (either disabled/hm down)
* `retransmit-retries` - Number of Retries for Retransmit, default is 5
* `insert-client-ip` - Insert client ip into TCP option
* `syn-retries` - SYN Retry Numbers, default is 5
* `timewait` - Timewait Threshold (sec), default 5 (number)
* `disable-tcp-timestamps` - disable TCP Timestamps Option
* `disable-window-scale` - disable TCP Window-Scale Option
* `disable-sack` - disable Selective Ack Option
* `invalid-rate-limit` - Invalid Packet Response Rate Limit (ms), default is 500 (number default 500 challenges)
* `disable-abc` - Appropriate Byte Counting RFC 3465 Disabled, default is enabled (Appropriate Byte Counting (ABC) is enabled by default)
* `reassembly-timeout` - The reassembly timeout, default is 30sec (number)
* `reassembly-limit` - The reassembly queuing limit, default is 25 segments (number)
* `min-rto` - The minmum retransmission timeout, default is 200ms (number)
* `limited-slowstart` - RFC 3742 Limited Slow-Start for TCP with Large Congestion Windows (number)
* `early-retransmit` - Configure the Early-Retransmit Algorithm (RFC 5827) (Early-Retransmit is disabled by default)
* `maxburst` - The max packet count sent per transmission event (number)
* `proxy-header-action` - 'insert': Insert proxy header;
* `version` - 'v1': version 1; 'v2': version 2;
* `uuid` - uuid of the object
* `user-tag` - Customized tag

