---
layout: "thunder"
page_title: "thunder: thunder_slb_virtual_server_port"
sidebar_current: "docs-thunder-resource-slb-virtual-server-port"
description: |-
	Provides details about thunder slb virtual server port resource for A10
---

# thunder\_slb\_virtual\_server\_port

`thunder_slb_virtual_server_port` Provides details about thunder slb virtual server port
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

// Put working JSON here
```

## Argument Reference

* `action` - ‘enable’: Enable; ‘disable’: Disable;
* `alt_protocol1` - ‘http’: HTTP Port;
* `alt_protocol2` - ‘tcp’: TCP LB service;
* `alternate_port` - Alternate Virtual Port
* `alternate_port_number` - Virtual Port
* `auto` - Configure auto NAT for the vport
* `clientip_sticky_nat` - Prefer to use same source NAT address for a client
* `conn_limit` - Connection Limit
* `def_selection_if_pref_failed` - ‘def-selection-if-pref-failed’: Use default server selection method if prefer method failed; ‘def-selection-if-pref-failed-disable’: Stop using default server selection method if prefer method failed;
* `enable_playerid_check` - Enable playerid checks on UDP packets once the AX is in active mode
* `eth_fwd` - Ethernet interface number
* `eth_rev` - Ethernet interface number
* `expand` - expand syn-cookie with timestamp and wscale
* `extended_stats` - Enable extended statistics on virtual port
* `force_routing_mode` - Force routing mode
* `gslb_enable` - Enable Global Server Load Balancing
* `ha_conn_mirror` - Enable for HA Conn sync
* `ip_map_list` - Enter name of IP Map List to be bound (IP Map List Name)
* `ipinip` - Enable IP in IP
* `l7_hardware_assist` - FPGA assist L7 packet parsing
* `message_switching` - Message switching
* `name` - SLB Virtual Service Name
* `no_auto_up_on_aflex` - Don’t automatically mark vport up when aFleX is bound
* `no_dest_nat` - Disable destination NAT, this option only supports in wildcard VIP or when a connection is operated in SSLi + EP mode
* `no_logging` - Do not log connection over limit event
* `on_syn` - Enable for HA Conn sync for l4 tcp sessions on SYN
* `optimization_level` - ‘0’: No optimization; ‘1’: Optimization level 1 (Experimental);
* `persist_type` - ‘src-dst-ip-swap-persist’: Create persist session after source IP and destination IP swap; ‘use-src-ip-for-dst-persist’: Use the source IP to create a destination persist session; ‘use-dst-ip-for-src-persist’: Use the destination IP to create source IP persist session;
* `pool` - Specify NAT pool or pool group
* `pool_shared` - Specify NAT pool or pool group
* `port_number` - Port
* `port_translation` - Enable port translation under no-dest-nat
* `precedence` - Set auto NAT pool as higher precedence for source NAT
* `protocol` - ‘tcp’: TCP LB service; ‘udp’: UDP Port; ‘others’: for no tcp/udp protocol, do IP load balancing; ‘diameter’: diameter port; ‘dns-tcp’: DNS service over TCP; ‘dns-udp’: DNS service over UDP; ‘fast-http’: Fast HTTP Port; ‘fix’: FIX Port; ‘ftp’: File Transfer Protocol Port; ‘ftp-proxy’: ftp proxy port; ‘http’: HTTP Port; ‘https’: HTTPS port; ‘imap’: imap proxy port; ‘mlb’: Message based load balancing; ‘mms’: Microsoft Multimedia Service Port; ‘mysql’: mssql port; ‘mssql’: mssql; ‘pop3’: pop3 proxy port; ‘radius’: RADIUS Port; ‘rtsp’: Real Time Streaming Protocol Port; ‘sip’: Session initiation protocol over UDP; ‘sip-tcp’: Session initiation protocol over TCP; ‘sips’: Session initiation protocol over TLS; ‘smpp-tcp’: SMPP service over TCP; ‘spdy’: spdy port; ‘spdys’: spdys port; ‘smtp’: SMTP Port; ‘ssl-proxy’: Generic SSL proxy; ‘ssli’: SSL insight; ‘tcp-proxy’: Generic TCP proxy; ‘tftp’: TFTP Port;
* `range` - Virtual Port range (Virtual Port range value)
* `rate` - Specify the log message rate
* `redirect_to_https` - Redirect HTTP to HTTPS
* `req_fail` - Use alternate virtual port when L7 request fail
* `reset` - Send client reset when connection number over limit
* `reset_on_server_selection_fail` - Send client reset when server selection fails
* `rtp_sip_call_id_match` - rtp traffic try to match the real server of sip smp call-id session
* `scaleout_bucket_count` - Number of traffic buckets
* `scaleout_device_group` - Device group id
* `secs` - Specify the interval in seconds
* `serv_sel_fail` - Use alternate virtual port when server selection failure
* `service_group` - Bind a Service Group to this Virtual Server (Service Group Name)
* `shared_partition_cache_template` - Reference a Cache template from shared partition
* `shared_partition_client_ssl_template` - Reference a Client SSL template from shared partition
* `shared_partition_connection_reuse_template` - Reference a connection reuse template from shared partition
* `shared_partition_diameter_template` - Reference a Diameter template from shared partition
* `shared_partition_dns_template` - Reference a dns template from shared partition
* `shared_partition_dynamic_service_template` - Reference a dynamic service template from shared partition
* `shared_partition_external_service_template` - Reference a external service template from shared partition
* `shared_partition_http_policy_template` - Reference a http policy template from shared partition
* `shared_partition_http_template` - Reference a HTTP template from shared partition
* `shared_partition_persist_cookie_template` - Reference a persist cookie template from shared partition
* `shared_partition_persist_destination_ip_template` - Reference a persist destination ip template from shared partition
* `shared_partition_persist_source_ip_template` - Reference a persist source ip template from shared partition
* `shared_partition_persist_ssl_sid_template` - Reference a persist SSL SID template from shared partition
* `shared_partition_policy_template` - Reference a policy template from shared partition
* `shared_partition_pool` - Specify NAT pool or pool group from shared partition
* `shared_partition_server_ssl_template` - Reference a SSL Server template from shared partition
* `shared_partition_tcp` - Reference a tcp template from shared partition
* `shared_partition_tcp_proxy_template` - Reference a TCP Proxy template from shared partition
* `shared_partition_udp` - Reference a UDP template from shared partition
* `shared_partition_virtual_port_template` - Reference a Virtual Port template from shared partition
* `skip_rev_hash` - Skip rev tuple hash insertion
* `snat_on_vip` - Enable source NAT traffic against VIP
* `stats_data_action` - ‘stats-data-enable’: Enable statistical data collection for virtual port; ‘stats-data-disable’: Disable statistical data collection for virtual port;
* `support_http2` - Support HTTP2
* `syn_cookie` - Enable syn-cookie
* `template_cache` - RAM caching template (Cache Template Name)
* `template_cache_shared` - Cache Template Name
* `template_client_ssh` - Client SSH Template (Client SSH Config Name)
* `template_client_ssl` - Client SSL Template (Client SSL Config Name)
* `template_client_ssl_shared` - Client SSL Template Name
* `template_connection_reuse` - Connection Reuse Template (Connection Reuse Template Name)
* `template_connection_reuse_shared` - Connection Reuse Template Name
* `template_dblb` - DBLB Template (DBLB template name)
* `template_diameter` - Diameter Template (diameter template name)
* `template_diameter_shared` - Diameter Template Name
* `template_dns` - DNS template (DNS template name)
* `template_dns_shared` - DNS Template Name
* `template_dynamic_service` - Dynamic Service Template (dynamic-service template name)
* `template_dynamic_service_shared` - Dynamic Service Template Name
* `template_external_service` - External service template (external-service template name)
* `template_external_service_shared` - External Service Template Name
* `template_file_inspection` - File Inspection service template (file-inspection template name)
* `template_fix` - FIX template (FIX Template Name)
* `template_ftp` - FTP port template (Ftp template name)
* `template_http` - HTTP Template (HTTP Config Name)
* `template_http_policy` - http-policy template (http-policy template name)
* `template_http_policy_shared` - Http Policy Template Name
* `template_http_shared` - HTTP Template Name
* `template_imap_pop3` - IMAP/POP3 Template (IMAP/POP3 Config Name)
* `template_persist_cookie` - Cookie persistence (Cookie persistence template name)
* `template_persist_cookie_shared` - Cookie Persistence Template Name
* `template_persist_destination_ip` - Destination IP persistence (Destination IP persistence template name)
* `template_persist_destination_ip_shared` - Destination IP Persistence Template Name
* `template_persist_source_ip` - Source IP persistence (Source IP persistence template name)
* `template_persist_source_ip_shared` - Source IP Persistence Template Name
* `template_persist_ssl_sid` - SSL session ID persistence (Source IP Persistence Config name)
* `template_persist_ssl_sid_shared` - SSL SID Persistence Template Name
* `template_policy` - Policy Template (Policy template name)
* `template_policy_shared` - Policy Template Name
* `template_reqmod_icap` - ICAP reqmod template (reqmod-icap template name)
* `template_respmod_icap` - ICAP respmod service template (respmod-icap template name)
* `template_scaleout` - Scaleout template (Scaleout template name)
* `template_server_ssh` - Server SSH Template (Server SSH Config Name)
* `template_server_ssl` - Server Side SSL Template (Server SSL Name)
* `template_server_ssl_shared` - Server SSL Template Name
* `template_sip` - SIP template
* `template_smpp` - SMPP template
* `template_smtp` - SMTP Template (SMTP Config Name)
* `template_ssli` - SSLi template (SSLi Template Name)
* `template_tcp` - L4 TCP Template (TCP Config Name)
* `template_tcp_proxy` - TCP Proxy Template Name
* `template_tcp_proxy_client` - TCP Proxy Config Client (TCP Proxy Config name)
* `template_tcp_proxy_server` - TCP Proxy Config Server (TCP Proxy Config name)
* `template_tcp_proxy_shared` - TCP Proxy Template name
* `template_tcp_shared` - TCP Template Name
* `template_udp` - L4 UDP Template (UDP Config Name)
* `template_udp_shared` - UDP Template Name
* `template_virtual_port` - Virtual port template (Virtual port template name)
* `template_virtual_port_shared` - Virtual Port Template Name
* `trunk_fwd` - Trunk interface number
* `trunk_rev` - Trunk interface number
* `use_alternate_port` - Use alternate virtual port
* `use_cgnv6` - Follow CGNv6 source NAT configuration
* `use_default_if_no_server` - Use default forwarding if server selection failed
* `use_rcv_hop_for_resp` - Use receive hop for response to client(For packets on default-vlan, also config “vlan-global enable-def-vlan-l2-forwarding”.)
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `view` - Specify a GSLB View (ID)
* `waf_template` - WAF template (WAF Template Name)
* `when_down` - Use alternate virtual port when down
* `when_down_protocol2` - Use alternate virtual port when down
* `acl_name` - Apply an access list name (Named Access List)
* `acl_name_seq_num` - Specify ACL precedence (sequence-number)
* `acl_name_seq_num_shared` - Specify ACL precedence (sequence-number)
* `acl_name_src_nat_pool` - Policy based Source NAT (NAT Pool or Pool Group)
* `acl_name_src_nat_pool_shared` - Policy based Source NAT (NAT Pool or Pool Group)
* `shared_partition_pool_name` - Policy based Source NAT from shared partition
* `counters1` - ‘all’: all; ‘curr_conn’: Current connections; ‘total_l4_conn’: Total L4 connections; ‘total_l7_conn’: Total L7 connections; ‘total_tcp_conn’: Total TCP connections; ‘total_conn’: Total connections; ‘total_fwd_bytes’: Total forward bytes; ‘total_fwd_pkts’: Total forward packets; ‘total_rev_bytes’: Total reverse bytes; ‘total_rev_pkts’: Total reverse packets; ‘total_dns_pkts’: Total DNS packets; ‘total_mf_dns_pkts’: Total MF DNS packets; ‘es_total_failure_actions’: Total failure actions; ‘compression_bytes_before’: Data into compression engine; ‘compression_bytes_after’: Data out of compression engine; ‘compression_hit’: Number of requests compressed; ‘compression_miss’: Number of requests NOT compressed; ‘compression_miss_no_client’: Compression miss no client; ‘compression_miss_template_exclusion’: Compression miss template exclusion; ‘curr_req’: Current requests; ‘total_req’: Total requests; ‘total_req_succ’: Total successful requests; ‘peak_conn’: Peak connections; ‘curr_conn_rate’: Current connection rate; ‘last_rsp_time’: Last response time; ‘fastest_rsp_time’: Fastest response time; ‘slowest_rsp_time’: Slowest response time; ‘loc_permit’: Permit number; ‘loc_deny’: Deny number; ‘loc_conn’: Connection number; ‘curr_ssl_conn’: Current SSL connections; ‘total_ssl_conn’: Total SSL connections;
* `acl_id` - ACL id VPORT
* `acl_id_seq_num` - Specify ACL precedence (sequence-number)
* `acl_id_seq_num_shared` - Specify ACL precedence (sequence-number)
* `acl_id_src_nat_pool` - Policy based Source NAT (NAT Pool or Pool Group)
* `acl_id_src_nat_pool_shared` - Policy based Source NAT (NAT Pool or Pool Group)
* `shared_partition_pool_id` - Policy based Source NAT from shared partition
* `aflex` - Bind aFleX Script to the Virtual Port (aFleX Script Name)
* `aflex_shared` - aFleX Script Name
* `aaa_policy` - Specify AAA policy name to bind to the virtual port

