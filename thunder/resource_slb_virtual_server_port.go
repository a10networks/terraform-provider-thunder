package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbVirtualServerPort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_virtual_server_port`: Virtual Port\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbVirtualServerPortCreate,
		UpdateContext: resourceSlbVirtualServerPortUpdate,
		ReadContext:   resourceSlbVirtualServerPortRead,
		DeleteContext: resourceSlbVirtualServerPortDelete,

		Schema: map[string]*schema.Schema{
			"acl_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_id": {
							Type: schema.TypeInt, Optional: true, Description: "ACL id VPORT",
						},
						"acl_name": {
							Type: schema.TypeString, Optional: true, Description: "Apply an access list name (Named Access List)",
						},
						"acl_id_shared": {
							Type: schema.TypeInt, Optional: true, Description: "acl id",
						},
						"acl_name_shared": {
							Type: schema.TypeString, Optional: true, Description: "Apply an access list name (Named Access List)",
						},
						"acl_id_src_nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "Policy based Source NAT (NAT Pool or Pool Group)",
						},
						"acl_id_seq_num": {
							Type: schema.TypeInt, Optional: true, Description: "Specify ACL precedence (sequence-number)",
						},
						"shared_partition_pool_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Policy based Source NAT from shared partition",
						},
						"acl_id_src_nat_pool_shared": {
							Type: schema.TypeString, Optional: true, Description: "Policy based Source NAT (NAT Pool or Pool Group)",
						},
						"acl_id_seq_num_shared": {
							Type: schema.TypeInt, Optional: true, Description: "Specify ACL precedence (sequence-number)",
						},
						"v_acl_id_src_nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "Policy based Source NAT (NAT Pool or Pool Group)",
						},
						"v_acl_id_seq_num": {
							Type: schema.TypeInt, Optional: true, Description: "Specify ACL precedence (sequence-number)",
						},
						"v_shared_partition_pool_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Policy based Source NAT from shared partition",
						},
						"v_acl_id_src_nat_pool_shared": {
							Type: schema.TypeString, Optional: true, Description: "Policy based Source NAT (NAT Pool or Pool Group)",
						},
						"v_acl_id_seq_num_shared": {
							Type: schema.TypeInt, Optional: true, Description: "Specify ACL precedence (sequence-number)",
						},
						"acl_name_src_nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "Policy based Source NAT (NAT Pool or Pool Group)",
						},
						"acl_name_seq_num": {
							Type: schema.TypeInt, Optional: true, Description: "Specify ACL precedence (sequence-number)",
						},
						"shared_partition_pool_name": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Policy based Source NAT from shared partition",
						},
						"acl_name_src_nat_pool_shared": {
							Type: schema.TypeString, Optional: true, Description: "Policy based Source NAT (NAT Pool or Pool Group)",
						},
						"acl_name_seq_num_shared": {
							Type: schema.TypeInt, Optional: true, Description: "Specify ACL precedence (sequence-number)",
						},
						"v_acl_name_src_nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "Policy based Source NAT (NAT Pool or Pool Group)",
						},
						"v_acl_name_seq_num": {
							Type: schema.TypeInt, Optional: true, Description: "Specify ACL precedence (sequence-number)",
						},
						"v_shared_partition_pool_name": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Policy based Source NAT from shared partition",
						},
						"v_acl_name_src_nat_pool_shared": {
							Type: schema.TypeString, Optional: true, Description: "Policy based Source NAT (NAT Pool or Pool Group)",
						},
						"v_acl_name_seq_num_shared": {
							Type: schema.TypeInt, Optional: true, Description: "Specify ACL precedence (sequence-number)",
						},
					},
				},
			},
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable; 'disable': Disable;",
			},
			"aflex_scripts": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aflex": {
							Type: schema.TypeString, Optional: true, Description: "aFleX Script Name",
						},
						"aflex_shared": {
							Type: schema.TypeString, Optional: true, Description: "aFleX Script Name",
						},
					},
				},
			},
			"aflex_table_entry_syn_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable aFlex entry sync for this port",
			},
			"aflex_table_entry_syn_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable aFlex entry sync for this port",
			},
			"alt_protocol1": {
				Type: schema.TypeString, Optional: true, Description: "'http': HTTP Port;",
			},
			"alt_protocol2": {
				Type: schema.TypeString, Optional: true, Description: "'tcp': TCP LB service;",
			},
			"alternate_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Alternate Virtual Port",
			},
			"alternate_port_number": {
				Type: schema.TypeInt, Optional: true, Description: "Virtual Port",
			},
			"attack_detection": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable analytics",
			},
			"auth_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aaa_policy": {
							Type: schema.TypeString, Optional: true, Description: "Specify AAA policy name to bind to the virtual port",
						},
					},
				},
			},
			"auto": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure auto NAT for the vport",
			},
			"clientip_sticky_nat": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Prefer to use same source NAT address for a client",
			},
			"conn_limit": {
				Type: schema.TypeInt, Optional: true, Default: 64000000, Description: "Connection Limit",
			},
			"cpu_compute": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable cpu compute on virtual port",
			},
			"def_selection_if_pref_failed": {
				Type: schema.TypeString, Optional: true, Default: "def-selection-if-pref-failed", Description: "'def-selection-if-pref-failed': Use default server selection method if prefer method failed; 'def-selection-if-pref-failed-disable': Stop using default server selection method if prefer method failed;",
			},
			"enable_playerid_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable playerid checks on UDP packets once the AX is in active mode",
			},
			"enable_scaleout": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"eth_fwd": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet interface number",
			},
			"eth_rev": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet interface number",
			},
			"expand": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "expand syn-cookie with timestamp and wscale",
			},
			"extended_stats": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable extended statistics on virtual port",
			},
			"fast_path": {
				Type: schema.TypeString, Optional: true, Description: "'force': Force fast path in SLB processing; 'disable': Disable fast path in SLB processing;",
			},
			"force_routing_mode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force routing mode",
			},
			"gslb_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Global Server Load Balancing",
			},
			"gtp_session_lb": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable GTP Session Load Balancing",
			},
			"ha_conn_mirror": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable for HA Conn sync",
			},
			"ignore_global": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore global substitute-source-mac",
			},
			"ip_map_list": {
				Type: schema.TypeString, Optional: true, Description: "Enter name of IP Map List to be bound (IP Map List Name)",
			},
			"ip_only_lb": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IP-Only LB mode",
			},
			"ip_smart_rr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use IP address round-robin behavior",
			},
			"ipinip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IP in IP",
			},
			"l7_hardware_assist": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "FPGA assist L7 packet parsing",
			},
			"l7_service_chain": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"memory_compute": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable dynamic memory compute on virtual port",
			},
			"message_switching": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Message switching",
			},
			"name": {
				Type: schema.TypeString, Optional: true, Description: "SLB Virtual Service Name",
			},
			"ng_waf": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Next-gen WAF",
			},
			"no_auto_up_on_aflex": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't automatically mark vport up when aFleX is bound",
			},
			"no_dest_nat": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable destination NAT, this option only supports in wildcard VIP or when a connection is operated in SSLi + EP mode",
			},
			"no_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not log connection over limit event",
			},
			"on_syn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable for HA Conn sync for l4 tcp sessions on SYN",
			},
			"one_server_conn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Support server that allow only one connection",
			},
			"optimization_level": {
				Type: schema.TypeString, Optional: true, Default: "0", Description: "'0': No optimization; '1': Optimization level 1 (Experimental);",
			},
			"p_template_sip_shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SIP Template Name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"persist_type": {
				Type: schema.TypeString, Optional: true, Description: "'src-dst-ip-swap-persist': Create persist session after source IP and destination IP swap; 'use-src-ip-for-dst-persist': Use the source IP to create a destination persist session; 'use-dst-ip-for-src-persist': Use the destination IP to create source IP persist session;",
			},
			"pool": {
				Type: schema.TypeString, Optional: true, Description: "Specify NAT pool or pool group",
			},
			"pool_shared": {
				Type: schema.TypeString, Optional: true, Description: "Specify NAT pool or pool group",
			},
			"port_number": {
				Type: schema.TypeInt, Required: true, Description: "Port",
			},
			"port_translation": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable port translation under no-dest-nat",
			},
			"precedence": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set auto NAT pool as higher precedence for source NAT",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': TCP LB service; 'udp': UDP Port; 'others': for no tcp/udp protocol, do IP load balancing; 'diameter': diameter port; 'dns-tcp': DNS service over TCP; 'dns-udp': DNS service over UDP; 'fast-http': Fast HTTP Port; 'fix': FIX Port; 'ftp': File Transfer Protocol Port; 'ftp-proxy': ftp proxy port; 'http': HTTP Port; 'https': HTTPS port; 'imap': imap proxy port; 'mlb': Message based load balancing; 'mms': Microsoft Multimedia Service Port; 'mysql': mssql port; 'mssql': mssql; 'pop3': pop3 proxy port; 'radius': RADIUS Port; 'rtsp': Real Time Streaming Protocol Port; 'sip': Session initiation protocol over UDP; 'sip-tcp': Session initiation protocol over TCP; 'sips': Session initiation protocol over TLS; 'smpp-tcp': SMPP service over TCP; 'spdy': spdy port; 'spdys': spdys port; 'smtp': SMTP Port; 'mqtt': MQTT Port; 'mqtts': MQTTS Port; 'ssl-proxy': Generic SSL proxy; 'ssli': SSL insight; 'ssh': SSH Port; 'tcp-proxy': Generic TCP proxy; 'tftp': TFTP Port; 'fast-fix': Fast FIX port; 'http-over-quic': HTTP3-over-quic port;",
			},
			"proxy_layer": {
				Type: schema.TypeString, Optional: true, Description: "'v1': Force using old proxy; 'v2': Force using new proxy;",
			},
			"range": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Virtual Port range (Virtual Port range value)",
			},
			"rate": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the log message rate",
			},
			"redirect_to_https": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Redirect HTTP to HTTPS",
			},
			"reply_acme_challenge": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reply ACME http-01 challenge. This option only takes effect in HTTP port 80",
			},
			"req_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use alternate virtual port when L7 request fail",
			},
			"reselection": {
				Type: schema.TypeString, Optional: true, Description: "'disable': disable;",
			},
			"reset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send client reset when connection number over limit",
			},
			"reset_on_server_selection_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send client reset when server selection fails",
			},
			"resolve_web_cat_list": {
				Type: schema.TypeString, Optional: true, Description: "Web Category List name",
			},
			"rtp_sip_call_id_match": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "rtp traffic try to match the real server of sip smp call-id session",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_conn': Current established connections; 'total_l4_conn': Total L4 connections established; 'total_l7_conn': Total L7 connections established; 'total_tcp_conn': Total TCP connections established; 'total_conn': Total connections established; 'total_fwd_bytes': Bytes processed in forward direction; 'total_fwd_pkts': Packets processed in forward direction; 'total_rev_bytes': Bytes processed in reverse direction; 'total_rev_pkts': Packets processed in reverse direction; 'total_dns_pkts': Total DNS packets processed; 'total_mf_dns_pkts': Total MF DNS packets; 'es_total_failure_actions': Total failure actions; 'compression_bytes_before': Data into gzip compression engine; 'compression_bytes_after': Data out of gzip compression engine; 'compression_hit': Number of requests compressed; 'compression_miss': Number of requests NOT compressed; 'compression_miss_no_client': Compression miss no client; 'compression_miss_template_exclusion': Compression miss template exclusion; 'curr_req': Current requests; 'total_req': Total requests; 'total_req_succ': Total successful requests; 'peak_conn': Peak connections; 'curr_conn_rate': Current connection rate; 'last_rsp_time': Last response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time; 'loc_permit': Geo-location Permit count; 'loc_deny': Geo-location Deny count; 'loc_conn': Geo-location Connection count; 'curr_ssl_conn': Current SSL connections; 'total_ssl_conn': Total SSL connections; 'backend-time-to-first-byte': Backend Time from Request to Response First Byte; 'backend-time-to-last-byte': Backend Time from Request to Response Last Byte; 'in-latency': Request Latency at Thunder; 'out-latency': Response Latency at Thunder; 'total_fwd_bytes_out': Bytes processed in forward direction (outbound); 'total_fwd_pkts_out': Packets processed in forward direction (outbound); 'total_rev_bytes_out': Bytes processed in reverse direction (outbound); 'total_rev_pkts_out': Packets processed in reverse direction (outbound); 'curr_req_rate': Current request rate; 'curr_resp': Current response; 'total_resp': Total response; 'total_resp_succ': Total successful responses; 'curr_resp_rate': Current response rate; 'dnsrrl_total_allowed': DNS Response-Rate-Limiting Total Responses Allowed; 'dnsrrl_total_dropped': DNS Response-Rate-Limiting Total Responses Dropped; 'dnsrrl_total_slipped': DNS Response-Rate-Limiting Total Responses Slipped; 'dnsrrl_bad_fqdn': DNS Response-Rate-Limiting Bad FQDN; 'throughput-bits-per-sec': Throughput in bits/sec; 'dynamic-memory-alloc': dynamic memory (bytes) allocated by the vport; 'dynamic-memory-free': dynamic memory (bytes) allocated by the vport; 'dynamic-memory': dynamic memory (bytes) used by the vport(alloc-free); 'ip_only_lb_fwd_bytes': IP-Only-LB Bytes processed in forward direction; 'ip_only_lb_rev_bytes': IP-Only-LB Bytes processed in reverse direction; 'ip_only_lb_fwd_pkts': IP-Only-LB Packets processed in forward direction; 'ip_only_lb_rev_pkts': IP-Only-LB Packets processed in reverse direction; 'total_dns_filter_type_drop': Total DNS Filter Type Drop; 'total_dns_filter_class_drop': Total DNS Filter Class Drop; 'dns_filter_type_a_drop': DNS Filter Type A Drop; 'dns_filter_type_aaaa_drop': DNS Filter Type AAAA Drop; 'dns_filter_type_cname_drop': DNS Filter Type CNAME Drop; 'dns_filter_type_mx_drop': DNS Filter Type MX Drop; 'dns_filter_type_ns_drop': DNS Filter Type NS Drop; 'dns_filter_type_srv_drop': DNS Filter Type SRV Drop; 'dns_filter_type_ptr_drop': DNS Filter Type PTR Drop; 'dns_filter_type_soa_drop': DNS Filter Type SOA Drop; 'dns_filter_type_txt_drop': DNS Filter Type TXT Drop; 'dns_filter_type_any_drop': DNS Filter Type Any Drop; 'dns_filter_type_others_drop': DNS Filter Type OTHERS Drop; 'dns_filter_class_internet_drop': DNS Filter Class INTERNET Drop; 'dns_filter_class_chaos_drop': DNS Filter Class CHAOS Drop; 'dns_filter_class_hesiod_drop': DNS Filter Class HESIOD Drop; 'dns_filter_class_none_drop': DNS Filter Class NONE Drop; 'dns_filter_class_any_drop': DNS Filter Class ANY Drop; 'dns_filter_class_others_drop': DNS Filter Class OTHERS Drop; 'dns_rpz_action_drop': DNS RPZ Action Drop; 'dns_rpz_action_pass_thru': DNS RPZ Action Pass Through; 'dns_rpz_action_tcp_only': DNS RPZ Action TCP Only; 'dns_rpz_action_nxdomain': DNS RPZ Action NXDOMAIN; 'dns_rpz_action_nodata': DNS RPZ Action NODATA; 'dns_rpz_action_local_data': DNS RPZ Action Local Data; 'dns_rpz_trigger_client_ip': DNS RPZ Trigger Client IP; 'dns_rpz_trigger_resp_ip': DNS RPZ Trigger Response IP; 'dns_rpz_trigger_ns_ip': DNS RPZ Trigger NS IP; 'dns_rpz_trigger_qname': DNS RPZ Trigger Qname IP; 'dns_rpz_trigger_ns_name': DNS RPZ Trigger NS Name; 'compression_bytes_before_br': Data into brotli compression engine; 'compression_bytes_after_br': Data out of brotli compression engine; 'compression_bytes_before_total': Data into compression engine; 'compression_bytes_after_total': Data out of compression engine; 'compression_hit_br': Number of requests compressed with brotli; 'compression_miss_br': Number of requests NOT compressed with brotli; 'compression_hit_total': Number of requests compressed; 'compression_miss_total': Number of requests NOT compressed; 'dnsrrl_total_tc': DNS Response-Rate-Limiting Total Responses Replied With Truncated; 'http1_client_idle_timeout': HTTP1 Client Idle Timeout; 'http2_client_idle_timeout': HTTP2 Client Idle Timeout;",
						},
					},
				},
			},
			"secs": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the interval in seconds",
			},
			"serv_sel_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use alternate virtual port when server selection failure",
			},
			"server_group": {
				Type: schema.TypeString, Optional: true, Description: "Bind a use-rcv-hop-for-resp Server Group to this Virtual Server (Server Group Name)",
			},
			"service_group": {
				Type: schema.TypeString, Optional: true, Description: "Bind a Service Group to this Virtual Server (Service Group Name)",
			},
			"shared_partition_cache_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a Cache template from shared partition",
			},
			"shared_partition_client_ssl_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a Client SSL template from shared partition",
			},
			"shared_partition_connection_reuse_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a connection reuse template from shared partition",
			},
			"shared_partition_dblb_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a dblb template from shared partition",
			},
			"shared_partition_diameter_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a Diameter template from shared partition",
			},
			"shared_partition_dns_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a dns template from shared partition",
			},
			"shared_partition_doh_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a DNS over HTTP(s) template from shared partition",
			},
			"shared_partition_dynamic_service_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a dynamic service template from shared partition",
			},
			"shared_partition_external_service_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a external service template from shared partition",
			},
			"shared_partition_fix_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a FIX template from shared partition",
			},
			"shared_partition_http_policy_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a http policy template from shared partition",
			},
			"shared_partition_http_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a HTTP template from shared partition",
			},
			"shared_partition_imap_pop3_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a IMAP/POP3 template from shared partition",
			},
			"shared_partition_persist_cookie_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a persist cookie template from shared partition",
			},
			"shared_partition_persist_destination_ip_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a persist destination ip template from shared partition",
			},
			"shared_partition_persist_source_ip_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a persist source ip template from shared partition",
			},
			"shared_partition_persist_ssl_sid_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a persist SSL SID template from shared partition",
			},
			"shared_partition_policy_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a policy template from shared partition",
			},
			"shared_partition_pool": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify NAT pool or pool group from shared partition",
			},
			"shared_partition_quic_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a QUIC template from shared partition",
			},
			"shared_partition_server_ssl_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a SSL Server template from shared partition",
			},
			"shared_partition_smpp_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a smpp template from shared partition",
			},
			"shared_partition_smtp_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a SMTP template from shared partition",
			},
			"shared_partition_tcp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a tcp template from shared partition",
			},
			"shared_partition_tcp_proxy_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a TCP Proxy template from shared partition",
			},
			"shared_partition_udp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a UDP template from shared partition",
			},
			"shared_partition_virtual_port_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a Virtual Port template from shared partition",
			},
			"showtech_print_extended_stats": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable print extended stats in showtech",
			},
			"skip_rev_hash": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Skip rev tuple hash insertion",
			},
			"snat_on_vip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable source NAT traffic against VIP",
			},
			"stats_data_action": {
				Type: schema.TypeString, Optional: true, Default: "stats-data-enable", Description: "'stats-data-enable': Enable statistical data collection for virtual port; 'stats-data-disable': Disable statistical data collection for virtual port;",
			},
			"substitute_source_mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Substitute Source MAC Address to that of the outgoing interface",
			},
			"support_http2": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Support HTTP2",
			},
			"syn_cookie": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable syn-cookie",
			},
			"template_cache": {
				Type: schema.TypeString, Optional: true, Description: "RAM caching template (Cache Template Name)",
			},
			"template_cache_shared": {
				Type: schema.TypeString, Optional: true, Description: "Cache Template Name",
			},
			"template_client_ssh": {
				Type: schema.TypeString, Optional: true, Description: "Client SSH Template (Client SSH Config Name)",
			},
			"template_client_ssl": {
				Type: schema.TypeString, Optional: true, Description: "Client SSL Template Name",
			},
			"template_client_ssl_shared": {
				Type: schema.TypeString, Optional: true, Description: "Client SSL Template Name",
			},
			"template_connection_reuse": {
				Type: schema.TypeString, Optional: true, Description: "Connection Reuse Template (Connection Reuse Template Name)",
			},
			"template_connection_reuse_shared": {
				Type: schema.TypeString, Optional: true, Description: "Connection Reuse Template Name",
			},
			"template_dblb": {
				Type: schema.TypeString, Optional: true, Description: "DBLB Template (DBLB template name)",
			},
			"template_dblb_shared": {
				Type: schema.TypeString, Optional: true, Description: "DBLB Template Name",
			},
			"template_diameter": {
				Type: schema.TypeString, Optional: true, Description: "Diameter Template (diameter template name)",
			},
			"template_diameter_shared": {
				Type: schema.TypeString, Optional: true, Description: "Diameter Template Name",
			},
			"template_dns": {
				Type: schema.TypeString, Optional: true, Description: "DNS template (DNS template name)",
			},
			"template_dns_shared": {
				Type: schema.TypeString, Optional: true, Description: "DNS Template Name",
			},
			"template_doh": {
				Type: schema.TypeString, Optional: true, Description: "DNS over HTTP(s) Template Name",
			},
			"template_doh_shared": {
				Type: schema.TypeString, Optional: true, Description: "DNS over HTTP(s) Template Name",
			},
			"template_dynamic_service": {
				Type: schema.TypeString, Optional: true, Description: "Dynamic Service Template (dynamic-service template name)",
			},
			"template_dynamic_service_shared": {
				Type: schema.TypeString, Optional: true, Description: "Dynamic Service Template Name",
			},
			"template_external_service": {
				Type: schema.TypeString, Optional: true, Description: "External service template (external-service template name)",
			},
			"template_external_service_shared": {
				Type: schema.TypeString, Optional: true, Description: "External Service Template Name",
			},
			"template_fix": {
				Type: schema.TypeString, Optional: true, Description: "FIX template (FIX Template Name)",
			},
			"template_fix_shared": {
				Type: schema.TypeString, Optional: true, Description: "FIX Template Name",
			},
			"template_ftp": {
				Type: schema.TypeString, Optional: true, Description: "FTP port template (Ftp template name)",
			},
			"template_http": {
				Type: schema.TypeString, Optional: true, Description: "HTTP Template Name",
			},
			"template_http_policy": {
				Type: schema.TypeString, Optional: true, Description: "http-policy template (http-policy template name)",
			},
			"template_http_policy_shared": {
				Type: schema.TypeString, Optional: true, Description: "Http Policy Template Name",
			},
			"template_http_shared": {
				Type: schema.TypeString, Optional: true, Description: "HTTP Template Name",
			},
			"template_imap_pop3": {
				Type: schema.TypeString, Optional: true, Description: "IMAP/POP3 Template (IMAP/POP3 Config Name)",
			},
			"template_imap_pop3_shared": {
				Type: schema.TypeString, Optional: true, Description: "IMAP/POP3 Template Name",
			},
			"template_mqtt": {
				Type: schema.TypeString, Optional: true, Description: "MQTT Template (MQTT Config Name)",
			},
			"template_persist_cookie": {
				Type: schema.TypeString, Optional: true, Description: "Cookie persistence (Cookie persistence template name)",
			},
			"template_persist_cookie_shared": {
				Type: schema.TypeString, Optional: true, Description: "Cookie Persistence Template Name",
			},
			"template_persist_destination_ip": {
				Type: schema.TypeString, Optional: true, Description: "Destination IP persistence (Destination IP persistence template name)",
			},
			"template_persist_destination_ip_shared": {
				Type: schema.TypeString, Optional: true, Description: "Destination IP Persistence Template Name",
			},
			"template_persist_source_ip": {
				Type: schema.TypeString, Optional: true, Description: "Source IP persistence (Source IP persistence template name)",
			},
			"template_persist_source_ip_shared": {
				Type: schema.TypeString, Optional: true, Description: "Source IP Persistence Template Name",
			},
			"template_persist_ssl_sid": {
				Type: schema.TypeString, Optional: true, Description: "SSL SID persistence (SSL SID persistence template name)",
			},
			"template_persist_ssl_sid_shared": {
				Type: schema.TypeString, Optional: true, Description: "SSL SID Persistence Template Name",
			},
			"template_policy": {
				Type: schema.TypeString, Optional: true, Description: "Policy Template (Policy template name)",
			},
			"template_policy_shared": {
				Type: schema.TypeString, Optional: true, Description: "Policy Template Name",
			},
			"template_quic": {
				Type: schema.TypeString, Optional: true, Description: "QUIC Template Name",
			},
			"template_quic_client": {
				Type: schema.TypeString, Optional: true, Description: "QUIC Config Client (QUIC Config name)",
			},
			"template_quic_server": {
				Type: schema.TypeString, Optional: true, Description: "QUIC Config Server (QUIC Config name)",
			},
			"template_quic_shared": {
				Type: schema.TypeString, Optional: true, Description: "QUIC Template name",
			},
			"template_ram_cache": {
				Type: schema.TypeString, Optional: true, Description: "RAM caching template (Cache Template Name)",
			},
			"template_reqmod_icap": {
				Type: schema.TypeString, Optional: true, Description: "ICAP reqmod template (reqmod-icap template name)",
			},
			"template_respmod_icap": {
				Type: schema.TypeString, Optional: true, Description: "ICAP respmod service template (respmod-icap template name)",
			},
			"template_scaleout": {
				Type: schema.TypeString, Optional: true, Description: "Scaleout template (Scaleout template name)",
			},
			"template_server_ssh": {
				Type: schema.TypeString, Optional: true, Description: "Server SSH Template (Server SSH Config Name)",
			},
			"template_server_ssl": {
				Type: schema.TypeString, Optional: true, Description: "Server Side SSL Template Name",
			},
			"template_server_ssl_shared": {
				Type: schema.TypeString, Optional: true, Description: "Server SSL Template Name",
			},
			"template_sip": {
				Type: schema.TypeString, Optional: true, Description: "SIP Template Name",
			},
			"template_sip_shared": {
				Type: schema.TypeString, Optional: true, Description: "SIP template",
			},
			"template_smpp": {
				Type: schema.TypeString, Optional: true, Description: "SMPP template",
			},
			"template_smpp_shared": {
				Type: schema.TypeString, Optional: true, Description: "SMPP Template Name",
			},
			"template_smtp": {
				Type: schema.TypeString, Optional: true, Description: "SMTP Template (SMTP Config Name)",
			},
			"template_smtp_shared": {
				Type: schema.TypeString, Optional: true, Description: "SMTP Template Name",
			},
			"template_ssli": {
				Type: schema.TypeString, Optional: true, Description: "SSLi template (SSLi Template Name)",
			},
			"template_tcp": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "TCP Template Name",
			},
			"template_tcp_proxy": {
				Type: schema.TypeString, Optional: true, Description: "TCP Proxy Template Name",
			},
			"template_tcp_proxy_client": {
				Type: schema.TypeString, Optional: true, Description: "TCP Proxy Config Client (TCP Proxy Config name)",
			},
			"template_tcp_proxy_server": {
				Type: schema.TypeString, Optional: true, Description: "TCP Proxy Config Server (TCP Proxy Config name)",
			},
			"template_tcp_proxy_shared": {
				Type: schema.TypeString, Optional: true, Description: "TCP Proxy Template name",
			},
			"template_tcp_shared": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "TCP Template Name",
			},
			"template_udp": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "L4 UDP Template",
			},
			"template_udp_shared": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "UDP Template Name",
			},
			"template_virtual_port": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "Virtual port template (Virtual port template name)",
			},
			"template_virtual_port_shared": {
				Type: schema.TypeString, Optional: true, Description: "Virtual Port Template Name",
			},
			"trunk_fwd": {
				Type: schema.TypeInt, Optional: true, Description: "Trunk interface number",
			},
			"trunk_rev": {
				Type: schema.TypeInt, Optional: true, Description: "Trunk interface number",
			},
			"use_alternate_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use alternate virtual port",
			},
			"use_cgnv6": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Follow CGNv6 source NAT configuration",
			},
			"use_default_if_no_server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use default forwarding if server selection failed",
			},
			"use_rcv_hop_for_resp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use receive hop for response to client(For packets on default-vlan, also config \"vlan-global enable-def-vlan-l2-forwarding\".)",
			},
			"use_rcv_hop_group": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set use-rcv-hop group",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"view": {
				Type: schema.TypeInt, Optional: true, Description: "Specify a GSLB View (ID)",
			},
			"when_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use alternate virtual port when down",
			},
			"when_down_protocol2": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use alternate virtual port when down",
			},
		},
	}
}
func resourceSlbVirtualServerPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbVirtualServerPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbVirtualServerPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbVirtualServerPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbVirtualServerPortAclList(d []interface{}) []edpt.SlbVirtualServerPortAclList {

	count1 := len(d)
	ret := make([]edpt.SlbVirtualServerPortAclList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbVirtualServerPortAclList
		oi.AclId = in["acl_id"].(int)
		oi.AclName = in["acl_name"].(string)
		oi.AclIdShared = in["acl_id_shared"].(int)
		oi.AclNameShared = in["acl_name_shared"].(string)
		oi.AclIdSrcNatPool = in["acl_id_src_nat_pool"].(string)
		oi.AclIdSeqNum = in["acl_id_seq_num"].(int)
		oi.SharedPartitionPoolId = in["shared_partition_pool_id"].(int)
		oi.AclIdSrcNatPoolShared = in["acl_id_src_nat_pool_shared"].(string)
		oi.AclIdSeqNumShared = in["acl_id_seq_num_shared"].(int)
		oi.VAclIdSrcNatPool = in["v_acl_id_src_nat_pool"].(string)
		oi.VAclIdSeqNum = in["v_acl_id_seq_num"].(int)
		oi.VSharedPartitionPoolId = in["v_shared_partition_pool_id"].(int)
		oi.VAclIdSrcNatPoolShared = in["v_acl_id_src_nat_pool_shared"].(string)
		oi.VAclIdSeqNumShared = in["v_acl_id_seq_num_shared"].(int)
		oi.AclNameSrcNatPool = in["acl_name_src_nat_pool"].(string)
		oi.AclNameSeqNum = in["acl_name_seq_num"].(int)
		oi.SharedPartitionPoolName = in["shared_partition_pool_name"].(int)
		oi.AclNameSrcNatPoolShared = in["acl_name_src_nat_pool_shared"].(string)
		oi.AclNameSeqNumShared = in["acl_name_seq_num_shared"].(int)
		oi.VAclNameSrcNatPool = in["v_acl_name_src_nat_pool"].(string)
		oi.VAclNameSeqNum = in["v_acl_name_seq_num"].(int)
		oi.VSharedPartitionPoolName = in["v_shared_partition_pool_name"].(int)
		oi.VAclNameSrcNatPoolShared = in["v_acl_name_src_nat_pool_shared"].(string)
		oi.VAclNameSeqNumShared = in["v_acl_name_seq_num_shared"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbVirtualServerPortAflexScripts(d []interface{}) []edpt.SlbVirtualServerPortAflexScripts {

	count1 := len(d)
	ret := make([]edpt.SlbVirtualServerPortAflexScripts, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbVirtualServerPortAflexScripts
		oi.Aflex = in["aflex"].(string)
		oi.AflexShared = in["aflex_shared"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbVirtualServerPortAuthCfg(d []interface{}) edpt.SlbVirtualServerPortAuthCfg {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortAuthCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AaaPolicy = in["aaa_policy"].(string)
	}
	return ret
}

func getSliceSlbVirtualServerPortSamplingEnable(d []interface{}) []edpt.SlbVirtualServerPortSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbVirtualServerPortSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbVirtualServerPortSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbVirtualServerPort(d *schema.ResourceData) edpt.SlbVirtualServerPort {
	var ret edpt.SlbVirtualServerPort
	ret.Inst.AclList = getSliceSlbVirtualServerPortAclList(d.Get("acl_list").([]interface{}))
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.AflexScripts = getSliceSlbVirtualServerPortAflexScripts(d.Get("aflex_scripts").([]interface{}))
	ret.Inst.AflexTableEntrySynDisable = d.Get("aflex_table_entry_syn_disable").(int)
	ret.Inst.AflexTableEntrySynEnable = d.Get("aflex_table_entry_syn_enable").(int)
	ret.Inst.AltProtocol1 = d.Get("alt_protocol1").(string)
	ret.Inst.AltProtocol2 = d.Get("alt_protocol2").(string)
	ret.Inst.AlternatePort = d.Get("alternate_port").(int)
	ret.Inst.AlternatePortNumber = d.Get("alternate_port_number").(int)
	ret.Inst.AttackDetection = d.Get("attack_detection").(int)
	ret.Inst.AuthCfg = getObjectSlbVirtualServerPortAuthCfg(d.Get("auth_cfg").([]interface{}))
	ret.Inst.Auto = d.Get("auto").(int)
	ret.Inst.ClientipStickyNat = d.Get("clientip_sticky_nat").(int)
	ret.Inst.ConnLimit = d.Get("conn_limit").(int)
	ret.Inst.CpuCompute = d.Get("cpu_compute").(int)
	ret.Inst.DefSelectionIfPrefFailed = d.Get("def_selection_if_pref_failed").(string)
	ret.Inst.EnablePlayeridCheck = d.Get("enable_playerid_check").(int)
	ret.Inst.EnableScaleout = d.Get("enable_scaleout").(int)
	ret.Inst.EthFwd = d.Get("eth_fwd").(int)
	ret.Inst.EthRev = d.Get("eth_rev").(int)
	ret.Inst.Expand = d.Get("expand").(int)
	ret.Inst.ExtendedStats = d.Get("extended_stats").(int)
	ret.Inst.FastPath = d.Get("fast_path").(string)
	ret.Inst.ForceRoutingMode = d.Get("force_routing_mode").(int)
	ret.Inst.GslbEnable = d.Get("gslb_enable").(int)
	ret.Inst.GtpSessionLb = d.Get("gtp_session_lb").(int)
	ret.Inst.HaConnMirror = d.Get("ha_conn_mirror").(int)
	ret.Inst.IgnoreGlobal = d.Get("ignore_global").(int)
	ret.Inst.IpMapList = d.Get("ip_map_list").(string)
	ret.Inst.IpOnlyLb = d.Get("ip_only_lb").(int)
	ret.Inst.IpSmartRr = d.Get("ip_smart_rr").(int)
	ret.Inst.Ipinip = d.Get("ipinip").(int)
	ret.Inst.L7HardwareAssist = d.Get("l7_hardware_assist").(int)
	ret.Inst.L7ServiceChain = d.Get("l7_service_chain").(int)
	ret.Inst.MemoryCompute = d.Get("memory_compute").(int)
	ret.Inst.MessageSwitching = d.Get("message_switching").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NgWaf = d.Get("ng_waf").(int)
	ret.Inst.NoAutoUpOnAflex = d.Get("no_auto_up_on_aflex").(int)
	ret.Inst.NoDestNat = d.Get("no_dest_nat").(int)
	ret.Inst.NoLogging = d.Get("no_logging").(int)
	ret.Inst.OnSyn = d.Get("on_syn").(int)
	ret.Inst.OneServerConn = d.Get("one_server_conn").(int)
	ret.Inst.OptimizationLevel = d.Get("optimization_level").(string)
	ret.Inst.PTemplateSipShared = d.Get("p_template_sip_shared").(int)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.PersistType = d.Get("persist_type").(string)
	ret.Inst.Pool = d.Get("pool").(string)
	ret.Inst.PoolShared = d.Get("pool_shared").(string)
	ret.Inst.PortNumber = d.Get("port_number").(int)
	ret.Inst.PortTranslation = d.Get("port_translation").(int)
	ret.Inst.Precedence = d.Get("precedence").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.ProxyLayer = d.Get("proxy_layer").(string)
	ret.Inst.Range = d.Get("range").(int)
	ret.Inst.Rate = d.Get("rate").(int)
	ret.Inst.RedirectToHttps = d.Get("redirect_to_https").(int)
	ret.Inst.ReplyAcmeChallenge = d.Get("reply_acme_challenge").(int)
	ret.Inst.ReqFail = d.Get("req_fail").(int)
	ret.Inst.Reselection = d.Get("reselection").(string)
	ret.Inst.Reset = d.Get("reset").(int)
	ret.Inst.ResetOnServerSelectionFail = d.Get("reset_on_server_selection_fail").(int)
	ret.Inst.ResolveWebCatList = d.Get("resolve_web_cat_list").(string)
	ret.Inst.RtpSipCallIdMatch = d.Get("rtp_sip_call_id_match").(int)
	ret.Inst.SamplingEnable = getSliceSlbVirtualServerPortSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Secs = d.Get("secs").(int)
	ret.Inst.ServSelFail = d.Get("serv_sel_fail").(int)
	ret.Inst.ServerGroup = d.Get("server_group").(string)
	ret.Inst.ServiceGroup = d.Get("service_group").(string)
	ret.Inst.SharedPartitionCacheTemplate = d.Get("shared_partition_cache_template").(int)
	ret.Inst.SharedPartitionClientSslTemplate = d.Get("shared_partition_client_ssl_template").(int)
	ret.Inst.SharedPartitionConnectionReuseTemplate = d.Get("shared_partition_connection_reuse_template").(int)
	ret.Inst.SharedPartitionDblbTemplate = d.Get("shared_partition_dblb_template").(int)
	ret.Inst.SharedPartitionDiameterTemplate = d.Get("shared_partition_diameter_template").(int)
	ret.Inst.SharedPartitionDnsTemplate = d.Get("shared_partition_dns_template").(int)
	ret.Inst.SharedPartitionDohTemplate = d.Get("shared_partition_doh_template").(int)
	ret.Inst.SharedPartitionDynamicServiceTemplate = d.Get("shared_partition_dynamic_service_template").(int)
	ret.Inst.SharedPartitionExternalServiceTemplate = d.Get("shared_partition_external_service_template").(int)
	ret.Inst.SharedPartitionFixTemplate = d.Get("shared_partition_fix_template").(int)
	ret.Inst.SharedPartitionHttpPolicyTemplate = d.Get("shared_partition_http_policy_template").(int)
	ret.Inst.SharedPartitionHttpTemplate = d.Get("shared_partition_http_template").(int)
	ret.Inst.SharedPartitionImapPop3Template = d.Get("shared_partition_imap_pop3_template").(int)
	ret.Inst.SharedPartitionPersistCookieTemplate = d.Get("shared_partition_persist_cookie_template").(int)
	ret.Inst.SharedPartitionPersistDestinationIpTemplate = d.Get("shared_partition_persist_destination_ip_template").(int)
	ret.Inst.SharedPartitionPersistSourceIpTemplate = d.Get("shared_partition_persist_source_ip_template").(int)
	ret.Inst.SharedPartitionPersistSslSidTemplate = d.Get("shared_partition_persist_ssl_sid_template").(int)
	ret.Inst.SharedPartitionPolicyTemplate = d.Get("shared_partition_policy_template").(int)
	ret.Inst.SharedPartitionPool = d.Get("shared_partition_pool").(int)
	ret.Inst.SharedPartitionQuicTemplate = d.Get("shared_partition_quic_template").(int)
	ret.Inst.SharedPartitionServerSslTemplate = d.Get("shared_partition_server_ssl_template").(int)
	ret.Inst.SharedPartitionSmppTemplate = d.Get("shared_partition_smpp_template").(int)
	ret.Inst.SharedPartitionSmtpTemplate = d.Get("shared_partition_smtp_template").(int)
	ret.Inst.SharedPartitionTcp = d.Get("shared_partition_tcp").(int)
	ret.Inst.SharedPartitionTcpProxyTemplate = d.Get("shared_partition_tcp_proxy_template").(int)
	ret.Inst.SharedPartitionUdp = d.Get("shared_partition_udp").(int)
	ret.Inst.SharedPartitionVirtualPortTemplate = d.Get("shared_partition_virtual_port_template").(int)
	ret.Inst.ShowtechPrintExtendedStats = d.Get("showtech_print_extended_stats").(int)
	ret.Inst.SkipRevHash = d.Get("skip_rev_hash").(int)
	ret.Inst.SnatOnVip = d.Get("snat_on_vip").(int)
	ret.Inst.StatsDataAction = d.Get("stats_data_action").(string)
	ret.Inst.SubstituteSourceMac = d.Get("substitute_source_mac").(int)
	ret.Inst.SupportHttp2 = d.Get("support_http2").(int)
	ret.Inst.SynCookie = d.Get("syn_cookie").(int)
	ret.Inst.TemplateCache = d.Get("template_cache").(string)
	ret.Inst.TemplateCacheShared = d.Get("template_cache_shared").(string)
	ret.Inst.TemplateClientSsh = d.Get("template_client_ssh").(string)
	ret.Inst.TemplateClientSsl = d.Get("template_client_ssl").(string)
	ret.Inst.TemplateClientSslShared = d.Get("template_client_ssl_shared").(string)
	ret.Inst.TemplateConnectionReuse = d.Get("template_connection_reuse").(string)
	ret.Inst.TemplateConnectionReuseShared = d.Get("template_connection_reuse_shared").(string)
	ret.Inst.TemplateDblb = d.Get("template_dblb").(string)
	ret.Inst.TemplateDblbShared = d.Get("template_dblb_shared").(string)
	ret.Inst.TemplateDiameter = d.Get("template_diameter").(string)
	ret.Inst.TemplateDiameterShared = d.Get("template_diameter_shared").(string)
	ret.Inst.TemplateDns = d.Get("template_dns").(string)
	ret.Inst.TemplateDnsShared = d.Get("template_dns_shared").(string)
	ret.Inst.TemplateDoh = d.Get("template_doh").(string)
	ret.Inst.TemplateDohShared = d.Get("template_doh_shared").(string)
	ret.Inst.TemplateDynamicService = d.Get("template_dynamic_service").(string)
	ret.Inst.TemplateDynamicServiceShared = d.Get("template_dynamic_service_shared").(string)
	ret.Inst.TemplateExternalService = d.Get("template_external_service").(string)
	ret.Inst.TemplateExternalServiceShared = d.Get("template_external_service_shared").(string)
	ret.Inst.TemplateFix = d.Get("template_fix").(string)
	ret.Inst.TemplateFixShared = d.Get("template_fix_shared").(string)
	ret.Inst.TemplateFtp = d.Get("template_ftp").(string)
	ret.Inst.TemplateHttp = d.Get("template_http").(string)
	ret.Inst.TemplateHttpPolicy = d.Get("template_http_policy").(string)
	ret.Inst.TemplateHttpPolicyShared = d.Get("template_http_policy_shared").(string)
	ret.Inst.TemplateHttpShared = d.Get("template_http_shared").(string)
	ret.Inst.TemplateImapPop3 = d.Get("template_imap_pop3").(string)
	ret.Inst.TemplateImapPop3Shared = d.Get("template_imap_pop3_shared").(string)
	ret.Inst.TemplateMqtt = d.Get("template_mqtt").(string)
	ret.Inst.TemplatePersistCookie = d.Get("template_persist_cookie").(string)
	ret.Inst.TemplatePersistCookieShared = d.Get("template_persist_cookie_shared").(string)
	ret.Inst.TemplatePersistDestinationIp = d.Get("template_persist_destination_ip").(string)
	ret.Inst.TemplatePersistDestinationIpShared = d.Get("template_persist_destination_ip_shared").(string)
	ret.Inst.TemplatePersistSourceIp = d.Get("template_persist_source_ip").(string)
	ret.Inst.TemplatePersistSourceIpShared = d.Get("template_persist_source_ip_shared").(string)
	ret.Inst.TemplatePersistSslSid = d.Get("template_persist_ssl_sid").(string)
	ret.Inst.TemplatePersistSslSidShared = d.Get("template_persist_ssl_sid_shared").(string)
	ret.Inst.TemplatePolicy = d.Get("template_policy").(string)
	ret.Inst.TemplatePolicyShared = d.Get("template_policy_shared").(string)
	ret.Inst.TemplateQuic = d.Get("template_quic").(string)
	ret.Inst.TemplateQuicClient = d.Get("template_quic_client").(string)
	ret.Inst.TemplateQuicServer = d.Get("template_quic_server").(string)
	ret.Inst.TemplateQuicShared = d.Get("template_quic_shared").(string)
	ret.Inst.TemplateRamCache = d.Get("template_ram_cache").(string)
	ret.Inst.TemplateReqmodIcap = d.Get("template_reqmod_icap").(string)
	ret.Inst.TemplateRespmodIcap = d.Get("template_respmod_icap").(string)
	ret.Inst.TemplateScaleout = d.Get("template_scaleout").(string)
	ret.Inst.TemplateServerSsh = d.Get("template_server_ssh").(string)
	ret.Inst.TemplateServerSsl = d.Get("template_server_ssl").(string)
	ret.Inst.TemplateServerSslShared = d.Get("template_server_ssl_shared").(string)
	ret.Inst.TemplateSip = d.Get("template_sip").(string)
	ret.Inst.TemplateSipShared = d.Get("template_sip_shared").(string)
	ret.Inst.TemplateSmpp = d.Get("template_smpp").(string)
	ret.Inst.TemplateSmppShared = d.Get("template_smpp_shared").(string)
	ret.Inst.TemplateSmtp = d.Get("template_smtp").(string)
	ret.Inst.TemplateSmtpShared = d.Get("template_smtp_shared").(string)
	ret.Inst.TemplateSsli = d.Get("template_ssli").(string)
	ret.Inst.TemplateTcp = d.Get("template_tcp").(string)
	ret.Inst.TemplateTcpProxy = d.Get("template_tcp_proxy").(string)
	ret.Inst.TemplateTcpProxyClient = d.Get("template_tcp_proxy_client").(string)
	ret.Inst.TemplateTcpProxyServer = d.Get("template_tcp_proxy_server").(string)
	ret.Inst.TemplateTcpProxyShared = d.Get("template_tcp_proxy_shared").(string)
	ret.Inst.TemplateTcpShared = d.Get("template_tcp_shared").(string)
	ret.Inst.TemplateUdp = d.Get("template_udp").(string)
	ret.Inst.TemplateUdpShared = d.Get("template_udp_shared").(string)
	ret.Inst.TemplateVirtualPort = d.Get("template_virtual_port").(string)
	ret.Inst.TemplateVirtualPortShared = d.Get("template_virtual_port_shared").(string)
	ret.Inst.TrunkFwd = d.Get("trunk_fwd").(int)
	ret.Inst.TrunkRev = d.Get("trunk_rev").(int)
	ret.Inst.UseAlternatePort = d.Get("use_alternate_port").(int)
	ret.Inst.UseCgnv6 = d.Get("use_cgnv6").(int)
	ret.Inst.UseDefaultIfNoServer = d.Get("use_default_if_no_server").(int)
	ret.Inst.UseRcvHopForResp = d.Get("use_rcv_hop_for_resp").(int)
	ret.Inst.UseRcvHopGroup = d.Get("use_rcv_hop_group").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.View = d.Get("view").(int)
	ret.Inst.WhenDown = d.Get("when_down").(int)
	ret.Inst.WhenDownProtocol2 = d.Get("when_down_protocol2").(int)
	return ret
}
