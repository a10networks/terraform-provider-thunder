---
layout: "thunder"
page_title: "thunder: thunder_slb_virtual_server"
sidebar_current: "docs-thunder-resource-slb-virtual-server"
description: |-
    Provides details about thunder slb virtual server resource for A10
---

# thunder\_slb\_virtual\_server

`thunder_slb_virtual_server` Provides details about thunder slb virtual server
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_virtual_server" "resourceSlbVirtualServerTest" {
	name = "string"
ipv6_address = "string"
ip_address = "string"
netmask = "string"
ipv6_acl = "string"
ipv6_acl_shared = "string"
acl_id = 0
acl_name = "string"
acl_id_shared = 0
acl_name_shared = "string"
use_if_ip = 0
ethernet = 0
description = "string"
enable_disable_action = "string"
redistribution_flagged = 0
vport_disable_action = "string"
suppress_internal_loopback = 0
arp_disable = 0
template_policy = "string"
shared_partition_policy_template = 0
template_policy_shared = "string"
template_virtual_server = "string"
shared_partition_vs_template = 0
template_virtual_server_shared = "string"
template_logging = "string"
template_scaleout = "string"
stats_data_action = "string"
extended_stats = 0
vrid = 0
disable_vip_adv = 0
ha_dynamic = 0
redistribute_route_map = "string"
uuid = "string"
user_tag = "string"
migrate_vip {  
 	target_data_cpu =  0 
	target_floating_ipv4 =  "string" 
	target_floating_ipv6 =  "string" 
	cancel_migration =  0 
	finish_migration =  0 
	uuid =  "string" 
	}
port-list {   
	port_number =  0 
	protocol =  "string" 
	range =  0 
	alternate_port =  0 
	proxy_layer =  "string" 
	optimization_level =  "string" 
	support_http2 =  0 
	ip_only_lb =  0 
	name =  "string" 
	conn_limit =  0 
	reset =  0 
	no_logging =  0 
	use_alternate_port =  0 
	alternate_port_number =  0 
	alt_protocol1 =  "string" 
	serv_sel_fail =  0 
	when_down =  0 
	alt_protocol2 =  "string" 
	req_fail =  0 
	when_down_protocol2 =  0 
	action =  "string" 
	l7_service_chain =  0 
	def_selection_if_pref_failed =  "string" 
	ha_conn_mirror =  0 
	on_syn =  0 
	skip_rev_hash =  0 
	message_switching =  0 
	force_routing_mode =  0 
	one_server_conn =  0 
	rate =  0 
	secs =  0 
	reset_on_server_selection_fail =  0 
	clientip_sticky_nat =  0 
	extended_stats =  0 
	gslb_enable =  0 
	view =  0 
	snat_on_vip =  0 
	stats_data_action =  "string" 
	syn_cookie =  0 
	expand =  0 
	attack_detection =  0 
acl-list {   
	acl_id =  0 
	acl_name =  "string" 
	acl_id_shared =  0 
	acl_name_shared =  "string" 
	acl_id_src_nat_pool =  "string" 
	acl_id_seq_num =  0 
	shared_partition_pool_id =  0 
	acl_id_src_nat_pool_shared =  "string" 
	acl_id_seq_num_shared =  0 
	v_acl_id_src_nat_pool =  "string" 
	v_acl_id_seq_num =  0 
	v_shared_partition_pool_id =  0 
	v_acl_id_src_nat_pool_shared =  "string" 
	v_acl_id_seq_num_shared =  0 
	acl_name_src_nat_pool =  "string" 
	acl_name_seq_num =  0 
	shared_partition_pool_name =  0 
	acl_name_src_nat_pool_shared =  "string" 
	acl_name_seq_num_shared =  0 
	v_acl_name_src_nat_pool =  "string" 
	v_acl_name_seq_num =  0 
	v_shared_partition_pool_name =  0 
	v_acl_name_src_nat_pool_shared =  "string" 
	v_acl_name_seq_num_shared =  0 
	}
	template_policy =  "string" 
	shared_partition_policy_template =  0 
	template_policy_shared =  "string" 
aflex-scripts {   
	aflex =  "string" 
	aflex_shared =  "string" 
	}
	no_auto_up_on_aflex =  0 
	enable_scaleout =  0 
	scaleout_bucket_count =  0 
	scaleout_device_group =  0 
	pool =  "string" 
	shared_partition_pool =  0 
	pool_shared =  "string" 
	auto =  0 
	precedence =  0 
	ip_smart_rr =  0 
	use_cgnv6 =  0 
	enable_playerid_check =  0 
	service_group =  "string" 
	ipinip =  0 
	ip_map_list =  "string" 
	rtp_sip_call_id_match =  0 
	use_rcv_hop_for_resp =  0 
	persist_type =  "string" 
	use_rcv_hop_group =  0 
	server_group =  "string" 
	reselection =  "string" 
	eth_fwd =  0 
	trunk_fwd =  0 
	eth_rev =  0 
	trunk_rev =  0 
	template_sip =  "string" 
	p_template_sip_shared =  0 
	template_sip_shared =  "string" 
	template_smpp =  "string" 
	shared_partition_smpp_template =  0 
	template_smpp_shared =  "string" 
	template_dblb =  "string" 
	shared_partition_dblb_template =  0 
	template_dblb_shared =  "string" 
	template_connection_reuse =  "string" 
	shared_partition_connection_reuse_template =  0 
	template_connection_reuse_shared =  "string" 
	template_dns =  "string" 
	shared_partition_dns_template =  0 
	template_dns_shared =  "string" 
	template_dynamic_service =  "string" 
	shared_partition_dynamic_service_template =  0 
	template_dynamic_service_shared =  "string" 
	template_persist_source_ip =  "string" 
	shared_partition_persist_source_ip_template =  0 
	template_persist_source_ip_shared =  "string" 
	template_persist_destination_ip =  "string" 
	shared_partition_persist_destination_ip_template =  0 
	template_persist_destination_ip_shared =  "string" 
	template_persist_ssl_sid =  "string" 
	shared_partition_persist_ssl_sid_template =  0 
	template_persist_ssl_sid_shared =  "string" 
	template_persist_cookie =  "string" 
	shared_partition_persist_cookie_template =  0 
	template_persist_cookie_shared =  "string" 
	template_imap_pop3 =  "string" 
	shared_partition_imap_pop3_template =  0 
	template_imap_pop3_shared =  "string" 
	template_smtp =  "string" 
	shared_partition_smtp_template =  0 
	template_smtp_shared =  "string" 
	template_mqtt =  "string" 
	template_http =  "string" 
	shared_partition_http_template =  0 
	template_http_shared =  "string" 
	template_http_policy =  "string" 
	shared_partition_http_policy_template =  0 
	template_http_policy_shared =  "string" 
	redirect_to_https =  0 
	template_external_service =  "string" 
	shared_partition_external_service_template =  0 
	template_external_service_shared =  "string" 
	template_reqmod_icap =  "string" 
	template_respmod_icap =  "string" 
	template_doh =  "string" 
	shared_partition_doh_template =  0 
	template_doh_shared =  "string" 
	template_server_ssl =  "string" 
	shared_partition_server_ssl_template =  0 
	template_server_ssl_shared =  "string" 
	template_client_ssl =  "string" 
	shared_partition_client_ssl_template =  0 
	template_client_ssl_shared =  "string" 
	template_server_ssh =  "string" 
	template_client_ssh =  "string" 
	template_udp =  "string" 
	shared_partition_udp =  0 
	template_udp_shared =  "string" 
	template_tcp =  "string" 
	shared_partition_tcp =  0 
	template_tcp_shared =  "string" 
	template_virtual_port =  "string" 
	shared_partition_virtual_port_template =  0 
	template_virtual_port_shared =  "string" 
	template_ftp =  "string" 
	template_diameter =  "string" 
	shared_partition_diameter_template =  0 
	template_diameter_shared =  "string" 
	template_cache =  "string" 
	shared_partition_cache_template =  0 
	template_cache_shared =  "string" 
	template_ram_cache =  "string" 
	template_fix =  "string" 
	shared_partition_fix_template =  0 
	template_fix_shared =  "string" 
	waf_template =  "string" 
	template_ssli =  "string" 
	template_tcp_proxy_client =  "string" 
	template_tcp_proxy_server =  "string" 
	template_tcp_proxy =  "string" 
	shared_partition_tcp_proxy_template =  0 
	template_tcp_proxy_shared =  "string" 
	use_default_if_no_server =  0 
	template_scaleout =  "string" 
	no_dest_nat =  0 
	port_translation =  0 
	l7_hardware_assist =  0 
auth_cfg {  
 	aaa_policy =  "string" 
	}
	cpu_compute =  0 
	memory_compute =  0 
	substitute_source_mac =  0 
	ignore_global =  0 
	aflex_table_entry_syn_disable =  0 
	aflex_table_entry_syn_enable =  0 
	gtp_session_lb =  0 
	reply_acme_challenge =  0 
	resolve_web_cat_list =  "string" 
	uuid =  "string" 
	user_tag =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	packet_capture_template =  "string" 
	}
 
}

```

## Argument Reference

* `virtual-server` - Create a Virtual Server
* `name` - SLB Virtual Service Name
* `ipv6-address` - IPV6 address
* `ip-address` - IP Address
* `netmask` - IP subnet mask
* `ipv6-acl` - ipv6 acl name
* `ipv6-acl-shared` - ipv6 acl name
* `acl-id` - ACL id VPORT
* `acl-name` - Apply an access list name (Named Access List)
* `acl-id-shared` - acl id
* `acl-name-shared` - Apply an access list name (Named Access List)
* `use-if-ip` - Use Interface IP
* `ethernet` - Ethernet interface
* `description` - Create a description for VIP
* `enable-disable-action` - 'enable': Enable Virtual Server (default); 'disable': Disable Virtual Server; 'disable-when-all-ports-down': Disable Virtual Server when all member ports are down; 'disable-when-any-port-down': Disable Virtual Server when any member port is down;
* `redistribution-flagged` - Flag VIP for special redistribution handling
* `vport-disable-action` - 'drop-packet': Drop packet for disabled virtual-port;
* `suppress-internal-loopback` - Suppress VIP internal loopback programming
* `arp-disable` - Disable Respond to Virtual Server ARP request
* `template-policy` - Policy Template (Policy template name)
* `shared-partition-policy-template` - Reference a policy template from shared partition
* `template-policy-shared` - Policy Template Name
* `template-virtual-server` - Virtual server template (Virtual server template name)
* `shared-partition-vs-template` - Reference a virtual-server template from shared partition
* `template-virtual-server-shared` - Virtual-Server Template Name
* `template-logging` - NAT Logging template (NAT Logging template name)
* `template-scaleout` - Scaleout template (Scaleout template name)
* `stats-data-action` - 'stats-data-enable': Enable statistical data collection for virtual port; 'stats-data-disable': Disable statistical data collection for virtual port;
* `extended-stats` - Enable extended statistics on virtual port
* `vrid` - Join a vrrp group (Specify ha VRRP-A vrid)
* `disable-vip-adv` - Disable virtual server GARP
* `ha-dynamic` - Dynamic failover based on vip status
* `redistribute-route-map` - Route map reference (Name of route-map)
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `target-data-cpu` - Number of CPUs on the target platform
* `target-floating-ipv4` - Specify IP address
* `target-floating-ipv6` - Specify IPv6 address
* `cancel-migration` - Cancel migration
* `finish-migration` - Complete the migration
* `port-number` - Port
* `protocol` - 'tcp': TCP LB service; 'udp': UDP Port; 'others': for no tcp/udp protocol, do IP load balancing; 'diameter': diameter port; 'dns-tcp': DNS service over TCP; 'dns-udp': DNS service over UDP; 'fast-http': Fast HTTP Port; 'fix': FIX Port; 'ftp': File Transfer Protocol Port; 'ftp-proxy': ftp proxy port; 'http': HTTP Port; 'https': HTTPS port; 'imap': imap proxy port; 'mlb': Message based load balancing; 'mms': Microsoft Multimedia Service Port; 'mysql': mssql port; 'mssql': mssql; 'pop3': pop3 proxy port; 'radius': RADIUS Port; 'rtsp': Real Time Streaming Protocol Port; 'sip': Session initiation protocol over UDP; 'sip-tcp': Session initiation protocol over TCP; 'sips': Session initiation protocol over TLS; 'smpp-tcp': SMPP service over TCP; 'spdy': spdy port; 'spdys': spdys port; 'smtp': SMTP Port; 'mqtt': MQTT Port; 'mqtts': MQTTS Port; 'ssl-proxy': Generic SSL proxy; 'ssli': SSL insight; 'ssh': SSH Port; 'tcp-proxy': Generic TCP proxy; 'tftp': TFTP Port; 'fast-fix': Fast FIX port;
* `range` - Virtual Port range (Virtual Port range value)
* `alternate-port` - Alternate Virtual Port
* `proxy-layer` - 'v1': Force using old proxy; 'v2': Force using new proxy;
* `optimization-level` - '0': No optimization; '1': Optimization level 1 (Experimental);
* `support-http2` - Support HTTP2
* `ip-only-lb` - Enable IP-Only LB mode
* `conn-limit` - Connection Limit
* `reset` - Send client reset when connection number over limit
* `no-logging` - Do not log connection over limit event
* `use-alternate-port` - Use alternate virtual port
* `alternate-port-number` - Virtual Port
* `alt-protocol1` - 'http': HTTP Port;
* `serv-sel-fail` - Use alternate virtual port when server selection failure
* `when-down` - Use alternate virtual port when down
* `alt-protocol2` - 'tcp': TCP LB service;
* `req-fail` - Use alternate virtual port when L7 request fail
* `when-down-protocol2` - Use alternate virtual port when down
* `action` - 'enable': Enable; 'disable': Disable;
* `def-selection-if-pref-failed` - 'def-selection-if-pref-failed': Use default server selection method if prefer method failed; 'def-selection-if-pref-failed-disable': Stop using default server selection method if prefer method failed;
* `ha-conn-mirror` - Enable for HA Conn sync
* `on-syn` - Enable for HA Conn sync for l4 tcp sessions on SYN
* `skip-rev-hash` - Skip rev tuple hash insertion
* `message-switching` - Message switching
* `force-routing-mode` - Force routing mode
* `one-server-conn` - Support server that allow only one connection
* `rate` - Specify the log message rate
* `secs` - Specify the interval in seconds
* `reset-on-server-selection-fail` - Send client reset when server selection fails
* `clientip-sticky-nat` - Prefer to use same source NAT address for a client
* `gslb-enable` - Enable Global Server Load Balancing
* `view` - Specify a GSLB View (ID)
* `snat-on-vip` - Enable source NAT traffic against VIP
* `syn-cookie` - Enable syn-cookie
* `expand` - expand syn-cookie with timestamp and wscale
* `attack-detection` - Enable analytics
* `acl-id-src-nat-pool` - Policy based Source NAT (NAT Pool or Pool Group)
* `acl-id-seq-num` - Specify ACL precedence (sequence-number)
* `shared-partition-pool-id` - Policy based Source NAT from shared partition
* `acl-id-src-nat-pool-shared` - Policy based Source NAT (NAT Pool or Pool Group)
* `acl-id-seq-num-shared` - Specify ACL precedence (sequence-number)
* `v-acl-id-src-nat-pool` - Policy based Source NAT (NAT Pool or Pool Group)
* `v-acl-id-seq-num` - Specify ACL precedence (sequence-number)
* `v-shared-partition-pool-id` - Policy based Source NAT from shared partition
* `v-acl-id-src-nat-pool-shared` - Policy based Source NAT (NAT Pool or Pool Group)
* `v-acl-id-seq-num-shared` - Specify ACL precedence (sequence-number)
* `acl-name-src-nat-pool` - Policy based Source NAT (NAT Pool or Pool Group)
* `acl-name-seq-num` - Specify ACL precedence (sequence-number)
* `shared-partition-pool-name` - Policy based Source NAT from shared partition
* `acl-name-src-nat-pool-shared` - Policy based Source NAT (NAT Pool or Pool Group)
* `acl-name-seq-num-shared` - Specify ACL precedence (sequence-number)
* `v-acl-name-src-nat-pool` - Policy based Source NAT (NAT Pool or Pool Group)
* `v-acl-name-seq-num` - Specify ACL precedence (sequence-number)
* `v-shared-partition-pool-name` - Policy based Source NAT from shared partition
* `v-acl-name-src-nat-pool-shared` - Policy based Source NAT (NAT Pool or Pool Group)
* `v-acl-name-seq-num-shared` - Specify ACL precedence (sequence-number)
* `aflex` - aFleX Script Name
* `aflex-shared` - aFleX Script Name
* `no-auto-up-on-aflex` - Don't automatically mark vport up when aFleX is bound
* `scaleout-bucket-count` - Number of traffic buckets
* `scaleout-device-group` - Device group id
* `pool` - Specify NAT pool or pool group
* `shared-partition-pool` - Specify NAT pool or pool group from shared partition
* `pool-shared` - Specify NAT pool or pool group
* `auto` - Configure auto NAT for the vport
* `precedence` - Set auto NAT pool as higher precedence for source NAT
* `ip-smart-rr` - Use IP address round-robin behavior
* `use-cgnv6` - Follow CGNv6 source NAT configuration
* `enable-playerid-check` - Enable playerid checks on UDP packets once the AX is in active mode
* `service-group` - Bind a Service Group to this Virtual Server (Service Group Name)
* `ipinip` - Enable IP in IP
* `ip-map-list` - Enter name of IP Map List to be bound (IP Map List Name)
* `rtp-sip-call-id-match` - rtp traffic try to match the real server of sip smp call-id session
* `use-rcv-hop-for-resp` - Use receive hop for response to client(For packets on default-vlan, also config "vlan-global enable-def-vlan-l2-forwarding".)
* `persist-type` - 'src-dst-ip-swap-persist': Create persist session after source IP and destination IP swap; 'use-src-ip-for-dst-persist': Use the source IP to create a destination persist session; 'use-dst-ip-for-src-persist': Use the destination IP to create source IP persist session;
* `use-rcv-hop-group` - Set use-rcv-hop group
* `server-group` - Bind a use-rcv-hop-for-resp Server Group to this Virtual Server (Server Group Name)
* `reselection` - 'disable': disable;
* `eth-fwd` - Ethernet interface number
* `trunk-fwd` - Trunk interface number
* `eth-rev` - Ethernet interface number
* `trunk-rev` - Trunk interface number
* `template-sip` - SIP Template Name
* `p-template-sip-shared` - SIP Template Name
* `template-sip-shared` - SIP template
* `template-smpp` - SMPP template
* `shared-partition-smpp-template` - Reference a smpp template from shared partition
* `template-smpp-shared` - SMPP Template Name
* `template-dblb` - DBLB Template (DBLB template name)
* `shared-partition-dblb-template` - Reference a dblb template from shared partition
* `template-dblb-shared` - DBLB Template Name
* `template-connection-reuse` - Connection Reuse Template (Connection Reuse Template Name)
* `shared-partition-connection-reuse-template` - Reference a connection reuse template from shared partition
* `template-connection-reuse-shared` - Connection Reuse Template Name
* `template-dns` - DNS template (DNS template name)
* `shared-partition-dns-template` - Reference a dns template from shared partition
* `template-dns-shared` - DNS Template Name
* `template-dynamic-service` - Dynamic Service Template (dynamic-service template name)
* `shared-partition-dynamic-service-template` - Reference a dynamic service template from shared partition
* `template-dynamic-service-shared` - Dynamic Service Template Name
* `template-persist-source-ip` - Source IP persistence (Source IP persistence template name)
* `shared-partition-persist-source-ip-template` - Reference a persist source ip template from shared partition
* `template-persist-source-ip-shared` - Source IP Persistence Template Name
* `template-persist-destination-ip` - Destination IP persistence (Destination IP persistence template name)
* `shared-partition-persist-destination-ip-template` - Reference a persist destination ip template from shared partition
* `template-persist-destination-ip-shared` - Destination IP Persistence Template Name
* `template-persist-ssl-sid` - SSL SID persistence (SSL SID persistence template name)
* `shared-partition-persist-ssl-sid-template` - Reference a persist SSL SID template from shared partition
* `template-persist-ssl-sid-shared` - SSL SID Persistence Template Name
* `template-persist-cookie` - Cookie persistence (Cookie persistence template name)
* `shared-partition-persist-cookie-template` - Reference a persist cookie template from shared partition
* `template-persist-cookie-shared` - Cookie Persistence Template Name
* `template-imap-pop3` - IMAP/POP3 Template (IMAP/POP3 Config Name)
* `shared-partition-imap-pop3-template` - Reference a IMAP/POP3 template from shared partition
* `template-imap-pop3-shared` - IMAP/POP3 Template Name
* `template-smtp` - SMTP Template (SMTP Config Name)
* `shared-partition-smtp-template` - Reference a SMTP template from shared partition
* `template-smtp-shared` - SMTP Template Name
* `template-mqtt` - MQTT Template (MQTT Config Name)
* `template-http` - HTTP Template Name
* `shared-partition-http-template` - Reference a HTTP template from shared partition
* `template-http-shared` - HTTP Template Name
* `template-http-policy` - http-policy template (http-policy template name)
* `shared-partition-http-policy-template` - Reference a http policy template from shared partition
* `template-http-policy-shared` - Http Policy Template Name
* `redirect-to-https` - Redirect HTTP to HTTPS
* `template-external-service` - External service template (external-service template name)
* `shared-partition-external-service-template` - Reference a external service template from shared partition
* `template-external-service-shared` - External Service Template Name
* `template-reqmod-icap` - ICAP reqmod template (reqmod-icap template name)
* `template-respmod-icap` - ICAP respmod service template (respmod-icap template name)
* `template-doh` - DNS over HTTP(s) Template Name
* `shared-partition-doh-template` - Reference a DNS over HTTP(s) template from shared partition
* `template-doh-shared` - DNS over HTTP(s) Template Name
* `template-server-ssl` - Server Side SSL Template Name
* `shared-partition-server-ssl-template` - Reference a SSL Server template from shared partition
* `template-server-ssl-shared` - Server SSL Template Name
* `template-client-ssl` - Client SSL Template Name
* `shared-partition-client-ssl-template` - Reference a Client SSL template from shared partition
* `template-client-ssl-shared` - Client SSL Template Name
* `template-server-ssh` - Server SSH Template (Server SSH Config Name)
* `template-client-ssh` - Client SSH Template (Client SSH Config Name)
* `template-udp` - L4 UDP Template
* `shared-partition-udp` - Reference a UDP template from shared partition
* `template-udp-shared` - UDP Template Name
* `template-tcp` - TCP Template Name
* `shared-partition-tcp` - Reference a tcp template from shared partition
* `template-tcp-shared` - TCP Template Name
* `template-virtual-port` - Virtual port template (Virtual port template name)
* `shared-partition-virtual-port-template` - Reference a Virtual Port template from shared partition
* `template-virtual-port-shared` - Virtual Port Template Name
* `template-ftp` - FTP port template (Ftp template name)
* `template-diameter` - Diameter Template (diameter template name)
* `shared-partition-diameter-template` - Reference a Diameter template from shared partition
* `template-diameter-shared` - Diameter Template Name
* `template-cache` - RAM caching template (Cache Template Name)
* `shared-partition-cache-template` - Reference a Cache template from shared partition
* `template-cache-shared` - Cache Template Name
* `template-ram-cache` - RAM caching template (Cache Template Name)
* `template-fix` - FIX template (FIX Template Name)
* `shared-partition-fix-template` - Reference a FIX template from shared partition
* `template-fix-shared` - FIX Template Name
* `waf-template` - WAF template (WAF Template Name)
* `template-ssli` - SSLi template (SSLi Template Name)
* `template-tcp-proxy-client` - TCP Proxy Config Client (TCP Proxy Config name)
* `template-tcp-proxy-server` - TCP Proxy Config Server (TCP Proxy Config name)
* `template-tcp-proxy` - TCP Proxy Template Name
* `shared-partition-tcp-proxy-template` - Reference a TCP Proxy template from shared partition
* `template-tcp-proxy-shared` - TCP Proxy Template name
* `use-default-if-no-server` - Use default forwarding if server selection failed
* `no-dest-nat` - Disable destination NAT, this option only supports in wildcard VIP or when a connection is operated in SSLi + EP mode
* `port-translation` - Enable port translation under no-dest-nat
* `l7-hardware-assist` - FPGA assist L7 packet parsing
* `aaa-policy` - Specify AAA policy name to bind to the virtual port
* `cpu-compute` - enable cpu compute on virtual port
* `memory-compute` - enable dynamic memory compute on virtual port
* `substitute-source-mac` - Substitute Source MAC Address to that of the outgoing interface
* `ignore-global` - Ignore global substitute-source-mac
* `aflex-table-entry-syn-disable` - Disable aFlex entry sync for this port
* `aflex-table-entry-syn-enable` - Enable aFlex entry sync for this port
* `gtp-session-lb` - Enable GTP Session Load Balancing
* `reply-acme-challenge` - Reply ACME http-01 challenge. This option only takes effect in HTTP port 80
* `resolve-web-cat-list` - Web Category List name
* `counters1` - 'all': all; 'curr_conn': Current established connections; 'total_l4_conn': Total L4 connections established; 'total_l7_conn': Total L7 connections established; 'total_tcp_conn': Total TCP connections established; 'total_conn': Total connections established; 'total_fwd_bytes': Bytes processed in forward direction; 'total_fwd_pkts': Packets processed in forward direction; 'total_rev_bytes': Bytes processed in reverse direction; 'total_rev_pkts': Packets processed in reverse direction; 'total_dns_pkts': Total DNS packets processed; 'total_mf_dns_pkts': Total MF DNS packets; 'es_total_failure_actions': Total failure actions; 'compression_bytes_before': Data into compression engine; 'compression_bytes_after': Data out of compression engine; 'compression_hit': Number of requests compressed; 'compression_miss': Number of requests NOT compressed; 'compression_miss_no_client': Compression miss no client; 'compression_miss_template_exclusion': Compression miss template exclusion; 'curr_req': Current requests; 'total_req': Total requests; 'total_req_succ': Total successful requests; 'peak_conn': Peak connections; 'curr_conn_rate': Current connection rate; 'last_rsp_time': Last response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time; 'loc_permit': Geo-location Permit count; 'loc_deny': Geo-location Deny count; 'loc_conn': Geo-location Connection count; 'curr_ssl_conn': Current SSL connections; 'total_ssl_conn': Total SSL connections; 'backend-time-to-first-byte': Backend Time from Request to Response First Byte; 'backend-time-to-last-byte': Backend Time from Request to Response Last Byte; 'in-latency': Request Latency at Thunder; 'out-latency': Response Latency at Thunder; 'total_fwd_bytes_out': Bytes processed in forward direction (outbound); 'total_fwd_pkts_out': Packets processed in forward direction (outbound); 'total_rev_bytes_out': Bytes processed in reverse direction (outbound); 'total_rev_pkts_out': Packets processed in reverse direction (outbound); 'curr_req_rate': Current request rate; 'curr_resp': Current response; 'total_resp': Total response; 'total_resp_succ': Total successful responses; 'curr_resp_rate': Current response rate; 'dnsrrl_total_allowed': DNS Response-Rate-Limiting Total Responses Allowed; 'dnsrrl_total_dropped': DNS Response-Rate-Limiting Total Responses Dropped; 'dnsrrl_total_slipped': DNS Response-Rate-Limiting Total Responses Slipped; 'dnsrrl_bad_fqdn': DNS Response-Rate-Limiting Bad FQDN; 'throughput-bits-per-sec': Throughput in bits/sec; 'dynamic-memory-alloc': dynamic memory (bytes) allocated by the vport; 'dynamic-memory-free': dynamic memory (bytes) allocated by the vport; 'dynamic-memory': dynamic memory (bytes) used by the vport(alloc-free); 'ip_only_lb_fwd_bytes': IP-Only-LB Bytes processed in forward direction; 'ip_only_lb_rev_bytes': IP-Only-LB Bytes processed in reverse direction; 'ip_only_lb_fwd_pkts': IP-Only-LB Packets processed in forward direction; 'ip_only_lb_rev_pkts': IP-Only-LB Packets processed in reverse direction; 'total_dns_filter_type_drop': Total DNS Filter Type Drop; 'total_dns_filter_class_drop': Total DNS Filter Class Drop; 'dns_filter_type_a_drop': DNS Filter Type A Drop; 'dns_filter_type_aaaa_drop': DNS Filter Type AAAA Drop; 'dns_filter_type_cname_drop': DNS Filter Type CNAME Drop; 'dns_filter_type_mx_drop': DNS Filter Type MX Drop; 'dns_filter_type_ns_drop': DNS Filter Type NS Drop; 'dns_filter_type_srv_drop': DNS Filter Type SRV Drop; 'dns_filter_type_ptr_drop': DNS Filter Type PTR Drop; 'dns_filter_type_soa_drop': DNS Filter Type SOA Drop; 'dns_filter_type_txt_drop': DNS Filter Type TXT Drop; 'dns_filter_type_any_drop': DNS Filter Type Any Drop; 'dns_filter_type_others_drop': DNS Filter Type OTHERS Drop; 'dns_filter_class_internet_drop': DNS Filter Class INTERNET Drop; 'dns_filter_class_chaos_drop': DNS Filter Class CHAOS Drop; 'dns_filter_class_hesiod_drop': DNS Filter Class HESIOD Drop; 'dns_filter_class_none_drop': DNS Filter Class NONE Drop; 'dns_filter_class_any_drop': DNS Filter Class ANY Drop; 'dns_filter_class_others_drop': DNS Filter Class OTHERS Drop; 'dns_rpz_action_drop': DNS RPZ Action Drop; 'dns_rpz_action_pass_thru': DNS RPZ Action Pass Through; 'dns_rpz_action_tcp_only': DNS RPZ Action TCP Only; 'dns_rpz_action_nxdomain': DNS RPZ Action NXDOMAIN; 'dns_rpz_action_nodata': DNS RPZ Action NODATA; 'dns_rpz_action_local_data': DNS RPZ Action Local Data; 'dns_rpz_trigger_client_ip': DNS RPZ Trigger Client IP; 'dns_rpz_trigger_resp_ip': DNS RPZ Trigger Response IP; 'dns_rpz_trigger_ns_ip': DNS RPZ Trigger NS IP; 'dns_rpz_trigger_qname': DNS RPZ Trigger Qname IP; 'dns_rpz_trigger_ns_name': DNS RPZ Trigger NS Name;
* `packet-capture-template` - Name of the packet capture template to be bind with this object

