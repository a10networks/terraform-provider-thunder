package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbVirtualServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_virtual_server`: Create a Virtual Server\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbVirtualServerCreate,
		UpdateContext: resourceSlbVirtualServerUpdate,
		ReadContext:   resourceSlbVirtualServerRead,
		DeleteContext: resourceSlbVirtualServerDelete,

		Schema: map[string]*schema.Schema{
			"acl_id": {
				Type: schema.TypeInt, Optional: true, Description: "acl id",
			},
			"acl_id_shared": {
				Type: schema.TypeInt, Optional: true, Description: "acl id",
			},
			"acl_name": {
				Type: schema.TypeString, Optional: true, Description: "Access List name (IPv4 Access List Name)",
			},
			"acl_name_shared": {
				Type: schema.TypeString, Optional: true, Description: "Access List name (IPv4 Access List Name)",
			},
			"arp_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Respond to Virtual Server ARP request",
			},
			"description": {
				Type: schema.TypeString, Optional: true, Description: "Create a description for VIP",
			},
			"disable_vip_adv": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable virtual server GARP",
			},
			"enable_disable_action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable Virtual Server (default); 'disable': Disable Virtual Server; 'disable-when-all-ports-down': Disable Virtual Server when all member ports are down; 'disable-when-any-port-down': Disable Virtual Server when any member port is down;",
			},
			"ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
			},
			"extended_stats": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable extended statistics on virtual server",
			},
			"gaming_protocol_compliance": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Gaming Protocol Compliance Check",
			},
			"ha_dynamic": {
				Type: schema.TypeInt, Optional: true, Description: "Dynamic failover based on vip status",
			},
			"ip_address": {
				Type: schema.TypeString, Optional: true, Description: "IP Address",
			},
			"ipv6_acl": {
				Type: schema.TypeString, Optional: true, Description: "ipv6 acl name",
			},
			"ipv6_acl_shared": {
				Type: schema.TypeString, Optional: true, Description: "ipv6 acl name",
			},
			"ipv6_address": {
				Type: schema.TypeString, Optional: true, Description: "IPV6 address",
			},
			"migrate_vip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_data_cpu": {
							Type: schema.TypeInt, Optional: true, Description: "Number of CPUs on the target platform",
						},
						"target_floating_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "Specify IP address",
						},
						"target_floating_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address",
						},
						"cancel_migration": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cancel migration",
						},
						"finish_migration": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Complete the migration",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "SLB Virtual Server Name",
			},
			"netmask": {
				Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
			},
			"port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_number": {
							Type: schema.TypeInt, Required: true, Description: "Port",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'tcp': TCP LB service; 'udp': UDP Port; 'others': for no tcp/udp protocol, do IP load balancing; 'diameter': diameter port; 'dns-tcp': DNS service over TCP; 'dns-udp': DNS service over UDP; 'fast-http': Fast HTTP Port; 'fix': FIX Port; 'ftp': File Transfer Protocol Port; 'ftp-proxy': ftp proxy port; 'http': HTTP Port; 'https': HTTPS port; 'imap': imap proxy port; 'mlb': Message based load balancing; 'mms': Microsoft Multimedia Service Port; 'mysql': mssql port; 'mssql': mssql; 'pop3': pop3 proxy port; 'radius': RADIUS Port; 'rtsp': Real Time Streaming Protocol Port; 'sip': Session initiation protocol over UDP; 'sip-tcp': Session initiation protocol over TCP; 'sips': Session initiation protocol over TLS; 'smpp-tcp': SMPP service over TCP; 'spdy': spdy port; 'spdys': spdys port; 'smtp': SMTP Port; 'mqtt': MQTT Port; 'mqtts': MQTTS Port; 'ssl-proxy': Generic SSL proxy; 'ssli': SSL insight; 'ssh': SSH Port; 'tcp-proxy': Generic TCP proxy; 'tftp': TFTP Port; 'fast-fix': Fast FIX port; 'http-over-quic': HTTP3-over-quic port;",
						},
						"range": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Virtual Port range (Virtual Port range value)",
						},
						"alternate_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Alternate Virtual Port",
						},
						"proxy_layer": {
							Type: schema.TypeString, Optional: true, Description: "'v1': Force using old proxy; 'v2': Force using new proxy;",
						},
						"optimization_level": {
							Type: schema.TypeString, Optional: true, Default: "0", Description: "'0': No optimization; '1': Optimization level 1 (Experimental);",
						},
						"support_http2": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Support HTTP2",
						},
						"ip_only_lb": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IP-Only LB mode",
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "SLB Virtual Service Name",
						},
						"conn_limit": {
							Type: schema.TypeInt, Optional: true, Default: 64000000, Description: "Connection Limit",
						},
						"reset": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send client reset when connection number over limit",
						},
						"no_logging": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not log connection over limit event",
						},
						"use_alternate_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use alternate virtual port",
						},
						"alternate_port_number": {
							Type: schema.TypeInt, Optional: true, Description: "Virtual Port",
						},
						"alt_protocol1": {
							Type: schema.TypeString, Optional: true, Description: "'http': HTTP Port;",
						},
						"serv_sel_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use alternate virtual port when server selection failure",
						},
						"when_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use alternate virtual port when down",
						},
						"alt_protocol2": {
							Type: schema.TypeString, Optional: true, Description: "'tcp': TCP LB service;",
						},
						"req_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use alternate virtual port when L7 request fail",
						},
						"when_down_protocol2": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use alternate virtual port when down",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable; 'disable': Disable;",
						},
						"l7_service_chain": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"def_selection_if_pref_failed": {
							Type: schema.TypeString, Optional: true, Default: "def-selection-if-pref-failed", Description: "'def-selection-if-pref-failed': Use default server selection method if prefer method failed; 'def-selection-if-pref-failed-disable': Stop using default server selection method if prefer method failed;",
						},
						"ha_conn_mirror": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable for HA Conn sync",
						},
						"on_syn": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable for HA Conn sync for l4 tcp sessions on SYN",
						},
						"skip_rev_hash": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Skip rev tuple hash insertion",
						},
						"message_switching": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Message switching",
						},
						"force_routing_mode": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force routing mode",
						},
						"one_server_conn": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Support server that allow only one connection",
						},
						"rate": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the log message rate",
						},
						"secs": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the interval in seconds",
						},
						"reset_on_server_selection_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send client reset when server selection fails",
						},
						"clientip_sticky_nat": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Prefer to use same source NAT address for a client",
						},
						"extended_stats": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable extended statistics on virtual port",
						},
						"gslb_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Global Server Load Balancing",
						},
						"view": {
							Type: schema.TypeInt, Optional: true, Description: "Specify a GSLB View (ID)",
						},
						"snat_on_vip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable source NAT traffic against VIP",
						},
						"stats_data_action": {
							Type: schema.TypeString, Optional: true, Default: "stats-data-enable", Description: "'stats-data-enable': Enable statistical data collection for virtual port; 'stats-data-disable': Disable statistical data collection for virtual port;",
						},
						"syn_cookie": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable syn-cookie",
						},
						"showtech_print_extended_stats": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable print extended stats in showtech",
						},
						"expand": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "expand syn-cookie with timestamp and wscale",
						},
						"attack_detection": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable analytics",
						},
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
						"template_policy": {
							Type: schema.TypeString, Optional: true, Description: "Policy Template (Policy template name)",
						},
						"shared_partition_policy_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a policy template from shared partition",
						},
						"template_policy_shared": {
							Type: schema.TypeString, Optional: true, Description: "Policy Template Name",
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
						"no_auto_up_on_aflex": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't automatically mark vport up when aFleX is bound",
						},
						"enable_scaleout": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"pool": {
							Type: schema.TypeString, Optional: true, Description: "Specify NAT pool or pool group",
						},
						"shared_partition_pool": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify NAT pool or pool group from shared partition",
						},
						"pool_shared": {
							Type: schema.TypeString, Optional: true, Description: "Specify NAT pool or pool group",
						},
						"auto": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure auto NAT for the vport",
						},
						"precedence": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set auto NAT pool as higher precedence for source NAT",
						},
						"ip_smart_rr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use IP address round-robin behavior",
						},
						"use_cgnv6": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Follow CGNv6 source NAT configuration",
						},
						"enable_playerid_check": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable playerid checks on UDP packets once the AX is in active mode",
						},
						"service_group": {
							Type: schema.TypeString, Optional: true, Description: "Bind a Service Group to this Virtual Server (Service Group Name)",
						},
						"ipinip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IP in IP",
						},
						"ip_map_list": {
							Type: schema.TypeString, Optional: true, Description: "Enter name of IP Map List to be bound (IP Map List Name)",
						},
						"rtp_sip_call_id_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "rtp traffic try to match the real server of sip smp call-id session",
						},
						"use_rcv_hop_for_resp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use receive hop for response to client(For packets on default-vlan, also config \"vlan-global enable-def-vlan-l2-forwarding\".)",
						},
						"persist_type": {
							Type: schema.TypeString, Optional: true, Description: "'src-dst-ip-swap-persist': Create persist session after source IP and destination IP swap; 'use-src-ip-for-dst-persist': Use the source IP to create a destination persist session; 'use-dst-ip-for-src-persist': Use the destination IP to create source IP persist session;",
						},
						"use_rcv_hop_group": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set use-rcv-hop group",
						},
						"server_group": {
							Type: schema.TypeString, Optional: true, Description: "Bind a use-rcv-hop-for-resp Server Group to this Virtual Server (Server Group Name)",
						},
						"reselection": {
							Type: schema.TypeString, Optional: true, Description: "'disable': disable;",
						},
						"eth_fwd": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface number",
						},
						"trunk_fwd": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk interface number",
						},
						"eth_rev": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface number",
						},
						"trunk_rev": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk interface number",
						},
						"template_sip": {
							Type: schema.TypeString, Optional: true, Description: "SIP Template Name",
						},
						"p_template_sip_shared": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "SIP Template Name",
						},
						"template_sip_shared": {
							Type: schema.TypeString, Optional: true, Description: "SIP template",
						},
						"template_smpp": {
							Type: schema.TypeString, Optional: true, Description: "SMPP template",
						},
						"shared_partition_smpp_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a smpp template from shared partition",
						},
						"template_smpp_shared": {
							Type: schema.TypeString, Optional: true, Description: "SMPP Template Name",
						},
						"template_dblb": {
							Type: schema.TypeString, Optional: true, Description: "DBLB Template (DBLB template name)",
						},
						"shared_partition_dblb_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a dblb template from shared partition",
						},
						"template_dblb_shared": {
							Type: schema.TypeString, Optional: true, Description: "DBLB Template Name",
						},
						"template_connection_reuse": {
							Type: schema.TypeString, Optional: true, Description: "Connection Reuse Template (Connection Reuse Template Name)",
						},
						"shared_partition_connection_reuse_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a connection reuse template from shared partition",
						},
						"template_connection_reuse_shared": {
							Type: schema.TypeString, Optional: true, Description: "Connection Reuse Template Name",
						},
						"template_dns": {
							Type: schema.TypeString, Optional: true, Description: "DNS template (DNS template name)",
						},
						"shared_partition_dns_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a dns template from shared partition",
						},
						"template_dns_shared": {
							Type: schema.TypeString, Optional: true, Description: "DNS Template Name",
						},
						"template_dynamic_service": {
							Type: schema.TypeString, Optional: true, Description: "Dynamic Service Template (dynamic-service template name)",
						},
						"shared_partition_dynamic_service_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a dynamic service template from shared partition",
						},
						"template_dynamic_service_shared": {
							Type: schema.TypeString, Optional: true, Description: "Dynamic Service Template Name",
						},
						"template_persist_source_ip": {
							Type: schema.TypeString, Optional: true, Description: "Source IP persistence (Source IP persistence template name)",
						},
						"shared_partition_persist_source_ip_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a persist source ip template from shared partition",
						},
						"template_persist_source_ip_shared": {
							Type: schema.TypeString, Optional: true, Description: "Source IP Persistence Template Name",
						},
						"template_persist_destination_ip": {
							Type: schema.TypeString, Optional: true, Description: "Destination IP persistence (Destination IP persistence template name)",
						},
						"shared_partition_persist_destination_ip_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a persist destination ip template from shared partition",
						},
						"template_persist_destination_ip_shared": {
							Type: schema.TypeString, Optional: true, Description: "Destination IP Persistence Template Name",
						},
						"template_persist_ssl_sid": {
							Type: schema.TypeString, Optional: true, Description: "SSL SID persistence (SSL SID persistence template name)",
						},
						"shared_partition_persist_ssl_sid_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a persist SSL SID template from shared partition",
						},
						"template_persist_ssl_sid_shared": {
							Type: schema.TypeString, Optional: true, Description: "SSL SID Persistence Template Name",
						},
						"template_persist_cookie": {
							Type: schema.TypeString, Optional: true, Description: "Cookie persistence (Cookie persistence template name)",
						},
						"shared_partition_persist_cookie_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a persist cookie template from shared partition",
						},
						"template_persist_cookie_shared": {
							Type: schema.TypeString, Optional: true, Description: "Cookie Persistence Template Name",
						},
						"template_imap_pop3": {
							Type: schema.TypeString, Optional: true, Description: "IMAP/POP3 Template (IMAP/POP3 Config Name)",
						},
						"shared_partition_imap_pop3_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a IMAP/POP3 template from shared partition",
						},
						"template_imap_pop3_shared": {
							Type: schema.TypeString, Optional: true, Description: "IMAP/POP3 Template Name",
						},
						"template_smtp": {
							Type: schema.TypeString, Optional: true, Description: "SMTP Template (SMTP Config Name)",
						},
						"shared_partition_smtp_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a SMTP template from shared partition",
						},
						"template_smtp_shared": {
							Type: schema.TypeString, Optional: true, Description: "SMTP Template Name",
						},
						"template_mqtt": {
							Type: schema.TypeString, Optional: true, Description: "MQTT Template (MQTT Config Name)",
						},
						"template_http": {
							Type: schema.TypeString, Optional: true, Description: "HTTP Template Name",
						},
						"shared_partition_http_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a HTTP template from shared partition",
						},
						"template_http_shared": {
							Type: schema.TypeString, Optional: true, Description: "HTTP Template Name",
						},
						"template_http_policy": {
							Type: schema.TypeString, Optional: true, Description: "http-policy template (http-policy template name)",
						},
						"shared_partition_http_policy_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a http policy template from shared partition",
						},
						"template_http_policy_shared": {
							Type: schema.TypeString, Optional: true, Description: "Http Policy Template Name",
						},
						"redirect_to_https": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Redirect HTTP to HTTPS",
						},
						"template_external_service": {
							Type: schema.TypeString, Optional: true, Description: "External service template (external-service template name)",
						},
						"shared_partition_external_service_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a external service template from shared partition",
						},
						"template_external_service_shared": {
							Type: schema.TypeString, Optional: true, Description: "External Service Template Name",
						},
						"template_reqmod_icap": {
							Type: schema.TypeString, Optional: true, Description: "ICAP reqmod template (reqmod-icap template name)",
						},
						"template_respmod_icap": {
							Type: schema.TypeString, Optional: true, Description: "ICAP respmod service template (respmod-icap template name)",
						},
						"template_doh": {
							Type: schema.TypeString, Optional: true, Description: "DNS over HTTP(s) Template Name",
						},
						"shared_partition_doh_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a DNS over HTTP(s) template from shared partition",
						},
						"template_doh_shared": {
							Type: schema.TypeString, Optional: true, Description: "DNS over HTTP(s) Template Name",
						},
						"template_server_ssl": {
							Type: schema.TypeString, Optional: true, Description: "Server Side SSL Template Name",
						},
						"shared_partition_server_ssl_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a SSL Server template from shared partition",
						},
						"template_server_ssl_shared": {
							Type: schema.TypeString, Optional: true, Description: "Server SSL Template Name",
						},
						"template_client_ssl": {
							Type: schema.TypeString, Optional: true, Description: "Client SSL Template Name",
						},
						"shared_partition_client_ssl_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a Client SSL template from shared partition",
						},
						"template_client_ssl_shared": {
							Type: schema.TypeString, Optional: true, Description: "Client SSL Template Name",
						},
						"template_server_ssh": {
							Type: schema.TypeString, Optional: true, Description: "Server SSH Template (Server SSH Config Name)",
						},
						"template_client_ssh": {
							Type: schema.TypeString, Optional: true, Description: "Client SSH Template (Client SSH Config Name)",
						},
						"template_udp": {
							Type: schema.TypeString, Optional: true, Default: "default", Description: "L4 UDP Template",
						},
						"shared_partition_udp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a UDP template from shared partition",
						},
						"template_udp_shared": {
							Type: schema.TypeString, Optional: true, Default: "default", Description: "UDP Template Name",
						},
						"template_tcp": {
							Type: schema.TypeString, Optional: true, Default: "default", Description: "TCP Template Name",
						},
						"shared_partition_tcp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a tcp template from shared partition",
						},
						"template_tcp_shared": {
							Type: schema.TypeString, Optional: true, Default: "default", Description: "TCP Template Name",
						},
						"template_virtual_port": {
							Type: schema.TypeString, Optional: true, Default: "default", Description: "Virtual port template (Virtual port template name)",
						},
						"shared_partition_virtual_port_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a Virtual Port template from shared partition",
						},
						"template_virtual_port_shared": {
							Type: schema.TypeString, Optional: true, Description: "Virtual Port Template Name",
						},
						"template_ftp": {
							Type: schema.TypeString, Optional: true, Description: "FTP port template (Ftp template name)",
						},
						"template_diameter": {
							Type: schema.TypeString, Optional: true, Description: "Diameter Template (diameter template name)",
						},
						"shared_partition_diameter_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a Diameter template from shared partition",
						},
						"template_diameter_shared": {
							Type: schema.TypeString, Optional: true, Description: "Diameter Template Name",
						},
						"template_cache": {
							Type: schema.TypeString, Optional: true, Description: "RAM caching template (Cache Template Name)",
						},
						"shared_partition_cache_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a Cache template from shared partition",
						},
						"template_cache_shared": {
							Type: schema.TypeString, Optional: true, Description: "Cache Template Name",
						},
						"template_ram_cache": {
							Type: schema.TypeString, Optional: true, Description: "RAM caching template (Cache Template Name)",
						},
						"template_fix": {
							Type: schema.TypeString, Optional: true, Description: "FIX template (FIX Template Name)",
						},
						"shared_partition_fix_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a FIX template from shared partition",
						},
						"template_fix_shared": {
							Type: schema.TypeString, Optional: true, Description: "FIX Template Name",
						},
						"template_ssli": {
							Type: schema.TypeString, Optional: true, Description: "SSLi template (SSLi Template Name)",
						},
						"template_tcp_proxy_client": {
							Type: schema.TypeString, Optional: true, Description: "TCP Proxy Config Client (TCP Proxy Config name)",
						},
						"template_tcp_proxy_server": {
							Type: schema.TypeString, Optional: true, Description: "TCP Proxy Config Server (TCP Proxy Config name)",
						},
						"template_tcp_proxy": {
							Type: schema.TypeString, Optional: true, Description: "TCP Proxy Template Name",
						},
						"shared_partition_tcp_proxy_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a TCP Proxy template from shared partition",
						},
						"template_tcp_proxy_shared": {
							Type: schema.TypeString, Optional: true, Description: "TCP Proxy Template name",
						},
						"template_quic_client": {
							Type: schema.TypeString, Optional: true, Description: "QUIC Config Client (QUIC Config name)",
						},
						"template_quic_server": {
							Type: schema.TypeString, Optional: true, Description: "QUIC Config Server (QUIC Config name)",
						},
						"template_quic": {
							Type: schema.TypeString, Optional: true, Description: "QUIC Template Name",
						},
						"shared_partition_quic_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a QUIC template from shared partition",
						},
						"template_quic_shared": {
							Type: schema.TypeString, Optional: true, Description: "QUIC Template name",
						},
						"use_default_if_no_server": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use default forwarding if server selection failed",
						},
						"template_scaleout": {
							Type: schema.TypeString, Optional: true, Description: "Scaleout template (Scaleout template name)",
						},
						"no_dest_nat": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable destination NAT, this option only supports in wildcard VIP or when a connection is operated in SSLi + EP mode",
						},
						"port_translation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable port translation under no-dest-nat",
						},
						"l7_hardware_assist": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "FPGA assist L7 packet parsing",
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
						"cpu_compute": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable cpu compute on virtual port",
						},
						"memory_compute": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable dynamic memory compute on virtual port",
						},
						"substitute_source_mac": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Substitute Source MAC Address to that of the outgoing interface",
						},
						"ignore_global": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore global substitute-source-mac",
						},
						"aflex_table_entry_syn_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable aFlex entry sync for this port",
						},
						"aflex_table_entry_syn_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable aFlex entry sync for this port",
						},
						"gtp_session_lb": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable GTP Session Load Balancing",
						},
						"reply_acme_challenge": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reply ACME http-01 challenge. This option only takes effect in HTTP port 80",
						},
						"resolve_web_cat_list": {
							Type: schema.TypeString, Optional: true, Description: "Web Category List name",
						},
						"ng_waf": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Next-gen WAF",
						},
						"fast_path": {
							Type: schema.TypeString, Optional: true, Description: "'force': Force fast path in SLB processing; 'disable': Disable fast path in SLB processing;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
						"packet_capture_template": {
							Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
						},
					},
				},
			},
			"redistribute_route_map": {
				Type: schema.TypeString, Optional: true, Description: "Route map reference (Name of route-map)",
			},
			"redistribution_flagged": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Flag VIP for special redistribution handling",
			},
			"shared_partition_policy_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a policy template from shared partition",
			},
			"shared_partition_vs_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a virtual-server template from shared partition",
			},
			"stats_data_action": {
				Type: schema.TypeString, Optional: true, Default: "stats-data-enable", Description: "'stats-data-enable': Enable statistical data collection for virtual server; 'stats-data-disable': Disable statistical data collection for virtual server;",
			},
			"suppress_internal_loopback": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Suppress VIP internal loopback programming",
			},
			"template_logging": {
				Type: schema.TypeString, Optional: true, Description: "NAT Logging template (NAT Logging template name)",
			},
			"template_policy": {
				Type: schema.TypeString, Optional: true, Description: "Policy template (Policy template name)",
			},
			"template_policy_shared": {
				Type: schema.TypeString, Optional: true, Description: "Policy Template Name",
			},
			"template_scaleout": {
				Type: schema.TypeString, Optional: true, Description: "Scaleout template (Scaleout template name)",
			},
			"template_virtual_server": {
				Type: schema.TypeString, Optional: true, Description: "Virtual server template (Virtual server template name)",
			},
			"template_virtual_server_shared": {
				Type: schema.TypeString, Optional: true, Description: "Virtual-Server Template Name",
			},
			"use_if_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use Interface IP",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vport_disable_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop-packet': Drop packet for disabled virtual-port;",
			},
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "Join a vrrp group (Specify ha VRRP-A vrid)",
			},
		},
	}
}
func resourceSlbVirtualServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbVirtualServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbVirtualServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbVirtualServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbVirtualServerMigrateVip1476(d []interface{}) edpt.SlbVirtualServerMigrateVip1476 {

	count1 := len(d)
	var ret edpt.SlbVirtualServerMigrateVip1476
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TargetDataCpu = in["target_data_cpu"].(int)
		ret.TargetFloatingIpv4 = in["target_floating_ipv4"].(string)
		ret.TargetFloatingIpv6 = in["target_floating_ipv6"].(string)
		ret.CancelMigration = in["cancel_migration"].(int)
		ret.FinishMigration = in["finish_migration"].(int)
		//omit uuid
	}
	return ret
}

func getSliceSlbVirtualServerPortList(d []interface{}) []edpt.SlbVirtualServerPortList {

	count1 := len(d)
	ret := make([]edpt.SlbVirtualServerPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbVirtualServerPortList
		oi.PortNumber = in["port_number"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Range = in["range"].(int)
		oi.AlternatePort = in["alternate_port"].(int)
		oi.ProxyLayer = in["proxy_layer"].(string)
		oi.OptimizationLevel = in["optimization_level"].(string)
		oi.SupportHttp2 = in["support_http2"].(int)
		oi.IpOnlyLb = in["ip_only_lb"].(int)
		oi.Name = in["name"].(string)
		oi.ConnLimit = in["conn_limit"].(int)
		oi.Reset = in["reset"].(int)
		oi.NoLogging = in["no_logging"].(int)
		oi.UseAlternatePort = in["use_alternate_port"].(int)
		oi.AlternatePortNumber = in["alternate_port_number"].(int)
		oi.AltProtocol1 = in["alt_protocol1"].(string)
		oi.ServSelFail = in["serv_sel_fail"].(int)
		oi.WhenDown = in["when_down"].(int)
		oi.AltProtocol2 = in["alt_protocol2"].(string)
		oi.ReqFail = in["req_fail"].(int)
		oi.WhenDownProtocol2 = in["when_down_protocol2"].(int)
		oi.Action = in["action"].(string)
		oi.L7ServiceChain = in["l7_service_chain"].(int)
		oi.DefSelectionIfPrefFailed = in["def_selection_if_pref_failed"].(string)
		oi.HaConnMirror = in["ha_conn_mirror"].(int)
		oi.OnSyn = in["on_syn"].(int)
		oi.SkipRevHash = in["skip_rev_hash"].(int)
		oi.MessageSwitching = in["message_switching"].(int)
		oi.ForceRoutingMode = in["force_routing_mode"].(int)
		oi.OneServerConn = in["one_server_conn"].(int)
		oi.Rate = in["rate"].(int)
		oi.Secs = in["secs"].(int)
		oi.ResetOnServerSelectionFail = in["reset_on_server_selection_fail"].(int)
		oi.ClientipStickyNat = in["clientip_sticky_nat"].(int)
		oi.ExtendedStats = in["extended_stats"].(int)
		oi.GslbEnable = in["gslb_enable"].(int)
		oi.View = in["view"].(int)
		oi.SnatOnVip = in["snat_on_vip"].(int)
		oi.StatsDataAction = in["stats_data_action"].(string)
		oi.SynCookie = in["syn_cookie"].(int)
		oi.ShowtechPrintExtendedStats = in["showtech_print_extended_stats"].(int)
		oi.Expand = in["expand"].(int)
		oi.AttackDetection = in["attack_detection"].(int)
		oi.AclList = getSliceSlbVirtualServerPortListAclList(in["acl_list"].([]interface{}))
		oi.TemplatePolicy = in["template_policy"].(string)
		oi.SharedPartitionPolicyTemplate = in["shared_partition_policy_template"].(int)
		oi.TemplatePolicyShared = in["template_policy_shared"].(string)
		oi.AflexScripts = getSliceSlbVirtualServerPortListAflexScripts(in["aflex_scripts"].([]interface{}))
		oi.NoAutoUpOnAflex = in["no_auto_up_on_aflex"].(int)
		oi.EnableScaleout = in["enable_scaleout"].(int)
		oi.Pool = in["pool"].(string)
		oi.SharedPartitionPool = in["shared_partition_pool"].(int)
		oi.PoolShared = in["pool_shared"].(string)
		oi.Auto = in["auto"].(int)
		oi.Precedence = in["precedence"].(int)
		oi.IpSmartRr = in["ip_smart_rr"].(int)
		oi.UseCgnv6 = in["use_cgnv6"].(int)
		oi.EnablePlayeridCheck = in["enable_playerid_check"].(int)
		oi.ServiceGroup = in["service_group"].(string)
		oi.Ipinip = in["ipinip"].(int)
		oi.IpMapList = in["ip_map_list"].(string)
		oi.RtpSipCallIdMatch = in["rtp_sip_call_id_match"].(int)
		oi.UseRcvHopForResp = in["use_rcv_hop_for_resp"].(int)
		oi.PersistType = in["persist_type"].(string)
		oi.UseRcvHopGroup = in["use_rcv_hop_group"].(int)
		oi.ServerGroup = in["server_group"].(string)
		oi.Reselection = in["reselection"].(string)
		oi.EthFwd = in["eth_fwd"].(int)
		oi.TrunkFwd = in["trunk_fwd"].(int)
		oi.EthRev = in["eth_rev"].(int)
		oi.TrunkRev = in["trunk_rev"].(int)
		oi.TemplateSip = in["template_sip"].(string)
		oi.PTemplateSipShared = in["p_template_sip_shared"].(int)
		oi.TemplateSipShared = in["template_sip_shared"].(string)
		oi.TemplateSmpp = in["template_smpp"].(string)
		oi.SharedPartitionSmppTemplate = in["shared_partition_smpp_template"].(int)
		oi.TemplateSmppShared = in["template_smpp_shared"].(string)
		oi.TemplateDblb = in["template_dblb"].(string)
		oi.SharedPartitionDblbTemplate = in["shared_partition_dblb_template"].(int)
		oi.TemplateDblbShared = in["template_dblb_shared"].(string)
		oi.TemplateConnectionReuse = in["template_connection_reuse"].(string)
		oi.SharedPartitionConnectionReuseTemplate = in["shared_partition_connection_reuse_template"].(int)
		oi.TemplateConnectionReuseShared = in["template_connection_reuse_shared"].(string)
		oi.TemplateDns = in["template_dns"].(string)
		oi.SharedPartitionDnsTemplate = in["shared_partition_dns_template"].(int)
		oi.TemplateDnsShared = in["template_dns_shared"].(string)
		oi.TemplateDynamicService = in["template_dynamic_service"].(string)
		oi.SharedPartitionDynamicServiceTemplate = in["shared_partition_dynamic_service_template"].(int)
		oi.TemplateDynamicServiceShared = in["template_dynamic_service_shared"].(string)
		oi.TemplatePersistSourceIp = in["template_persist_source_ip"].(string)
		oi.SharedPartitionPersistSourceIpTemplate = in["shared_partition_persist_source_ip_template"].(int)
		oi.TemplatePersistSourceIpShared = in["template_persist_source_ip_shared"].(string)
		oi.TemplatePersistDestinationIp = in["template_persist_destination_ip"].(string)
		oi.SharedPartitionPersistDestinationIpTemplate = in["shared_partition_persist_destination_ip_template"].(int)
		oi.TemplatePersistDestinationIpShared = in["template_persist_destination_ip_shared"].(string)
		oi.TemplatePersistSslSid = in["template_persist_ssl_sid"].(string)
		oi.SharedPartitionPersistSslSidTemplate = in["shared_partition_persist_ssl_sid_template"].(int)
		oi.TemplatePersistSslSidShared = in["template_persist_ssl_sid_shared"].(string)
		oi.TemplatePersistCookie = in["template_persist_cookie"].(string)
		oi.SharedPartitionPersistCookieTemplate = in["shared_partition_persist_cookie_template"].(int)
		oi.TemplatePersistCookieShared = in["template_persist_cookie_shared"].(string)
		oi.TemplateImapPop3 = in["template_imap_pop3"].(string)
		oi.SharedPartitionImapPop3Template = in["shared_partition_imap_pop3_template"].(int)
		oi.TemplateImapPop3Shared = in["template_imap_pop3_shared"].(string)
		oi.TemplateSmtp = in["template_smtp"].(string)
		oi.SharedPartitionSmtpTemplate = in["shared_partition_smtp_template"].(int)
		oi.TemplateSmtpShared = in["template_smtp_shared"].(string)
		oi.TemplateMqtt = in["template_mqtt"].(string)
		oi.TemplateHttp = in["template_http"].(string)
		oi.SharedPartitionHttpTemplate = in["shared_partition_http_template"].(int)
		oi.TemplateHttpShared = in["template_http_shared"].(string)
		oi.TemplateHttpPolicy = in["template_http_policy"].(string)
		oi.SharedPartitionHttpPolicyTemplate = in["shared_partition_http_policy_template"].(int)
		oi.TemplateHttpPolicyShared = in["template_http_policy_shared"].(string)
		oi.RedirectToHttps = in["redirect_to_https"].(int)
		oi.TemplateExternalService = in["template_external_service"].(string)
		oi.SharedPartitionExternalServiceTemplate = in["shared_partition_external_service_template"].(int)
		oi.TemplateExternalServiceShared = in["template_external_service_shared"].(string)
		oi.TemplateReqmodIcap = in["template_reqmod_icap"].(string)
		oi.TemplateRespmodIcap = in["template_respmod_icap"].(string)
		oi.TemplateDoh = in["template_doh"].(string)
		oi.SharedPartitionDohTemplate = in["shared_partition_doh_template"].(int)
		oi.TemplateDohShared = in["template_doh_shared"].(string)
		oi.TemplateServerSsl = in["template_server_ssl"].(string)
		oi.SharedPartitionServerSslTemplate = in["shared_partition_server_ssl_template"].(int)
		oi.TemplateServerSslShared = in["template_server_ssl_shared"].(string)
		oi.TemplateClientSsl = in["template_client_ssl"].(string)
		oi.SharedPartitionClientSslTemplate = in["shared_partition_client_ssl_template"].(int)
		oi.TemplateClientSslShared = in["template_client_ssl_shared"].(string)
		oi.TemplateServerSsh = in["template_server_ssh"].(string)
		oi.TemplateClientSsh = in["template_client_ssh"].(string)
		oi.TemplateUdp = in["template_udp"].(string)
		oi.SharedPartitionUdp = in["shared_partition_udp"].(int)
		oi.TemplateUdpShared = in["template_udp_shared"].(string)
		oi.TemplateTcp = in["template_tcp"].(string)
		oi.SharedPartitionTcp = in["shared_partition_tcp"].(int)
		oi.TemplateTcpShared = in["template_tcp_shared"].(string)
		oi.TemplateVirtualPort = in["template_virtual_port"].(string)
		oi.SharedPartitionVirtualPortTemplate = in["shared_partition_virtual_port_template"].(int)
		oi.TemplateVirtualPortShared = in["template_virtual_port_shared"].(string)
		oi.TemplateFtp = in["template_ftp"].(string)
		oi.TemplateDiameter = in["template_diameter"].(string)
		oi.SharedPartitionDiameterTemplate = in["shared_partition_diameter_template"].(int)
		oi.TemplateDiameterShared = in["template_diameter_shared"].(string)
		oi.TemplateCache = in["template_cache"].(string)
		oi.SharedPartitionCacheTemplate = in["shared_partition_cache_template"].(int)
		oi.TemplateCacheShared = in["template_cache_shared"].(string)
		oi.TemplateRamCache = in["template_ram_cache"].(string)
		oi.TemplateFix = in["template_fix"].(string)
		oi.SharedPartitionFixTemplate = in["shared_partition_fix_template"].(int)
		oi.TemplateFixShared = in["template_fix_shared"].(string)
		oi.TemplateSsli = in["template_ssli"].(string)
		oi.TemplateTcpProxyClient = in["template_tcp_proxy_client"].(string)
		oi.TemplateTcpProxyServer = in["template_tcp_proxy_server"].(string)
		oi.TemplateTcpProxy = in["template_tcp_proxy"].(string)
		oi.SharedPartitionTcpProxyTemplate = in["shared_partition_tcp_proxy_template"].(int)
		oi.TemplateTcpProxyShared = in["template_tcp_proxy_shared"].(string)
		oi.TemplateQuicClient = in["template_quic_client"].(string)
		oi.TemplateQuicServer = in["template_quic_server"].(string)
		oi.TemplateQuic = in["template_quic"].(string)
		oi.SharedPartitionQuicTemplate = in["shared_partition_quic_template"].(int)
		oi.TemplateQuicShared = in["template_quic_shared"].(string)
		oi.UseDefaultIfNoServer = in["use_default_if_no_server"].(int)
		oi.TemplateScaleout = in["template_scaleout"].(string)
		oi.NoDestNat = in["no_dest_nat"].(int)
		oi.PortTranslation = in["port_translation"].(int)
		oi.L7HardwareAssist = in["l7_hardware_assist"].(int)
		oi.AuthCfg = getObjectSlbVirtualServerPortListAuthCfg(in["auth_cfg"].([]interface{}))
		oi.CpuCompute = in["cpu_compute"].(int)
		oi.MemoryCompute = in["memory_compute"].(int)
		oi.SubstituteSourceMac = in["substitute_source_mac"].(int)
		oi.IgnoreGlobal = in["ignore_global"].(int)
		oi.AflexTableEntrySynDisable = in["aflex_table_entry_syn_disable"].(int)
		oi.AflexTableEntrySynEnable = in["aflex_table_entry_syn_enable"].(int)
		oi.GtpSessionLb = in["gtp_session_lb"].(int)
		oi.ReplyAcmeChallenge = in["reply_acme_challenge"].(int)
		oi.ResolveWebCatList = in["resolve_web_cat_list"].(string)
		oi.NgWaf = in["ng_waf"].(int)
		oi.FastPath = in["fast_path"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceSlbVirtualServerPortListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.PacketCaptureTemplate = in["packet_capture_template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbVirtualServerPortListAclList(d []interface{}) []edpt.SlbVirtualServerPortListAclList {

	count1 := len(d)
	ret := make([]edpt.SlbVirtualServerPortListAclList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbVirtualServerPortListAclList
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

func getSliceSlbVirtualServerPortListAflexScripts(d []interface{}) []edpt.SlbVirtualServerPortListAflexScripts {

	count1 := len(d)
	ret := make([]edpt.SlbVirtualServerPortListAflexScripts, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbVirtualServerPortListAflexScripts
		oi.Aflex = in["aflex"].(string)
		oi.AflexShared = in["aflex_shared"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbVirtualServerPortListAuthCfg(d []interface{}) edpt.SlbVirtualServerPortListAuthCfg {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortListAuthCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AaaPolicy = in["aaa_policy"].(string)
	}
	return ret
}

func getSliceSlbVirtualServerPortListSamplingEnable(d []interface{}) []edpt.SlbVirtualServerPortListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbVirtualServerPortListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbVirtualServerPortListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbVirtualServer(d *schema.ResourceData) edpt.SlbVirtualServer {
	var ret edpt.SlbVirtualServer
	ret.Inst.AclId = d.Get("acl_id").(int)
	ret.Inst.AclIdShared = d.Get("acl_id_shared").(int)
	ret.Inst.AclName = d.Get("acl_name").(string)
	ret.Inst.AclNameShared = d.Get("acl_name_shared").(string)
	ret.Inst.ArpDisable = d.Get("arp_disable").(int)
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.DisableVipAdv = d.Get("disable_vip_adv").(int)
	ret.Inst.EnableDisableAction = d.Get("enable_disable_action").(string)
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.ExtendedStats = d.Get("extended_stats").(int)
	ret.Inst.GamingProtocolCompliance = d.Get("gaming_protocol_compliance").(int)
	ret.Inst.HaDynamic = d.Get("ha_dynamic").(int)
	ret.Inst.IpAddress = d.Get("ip_address").(string)
	ret.Inst.Ipv6Acl = d.Get("ipv6_acl").(string)
	ret.Inst.Ipv6AclShared = d.Get("ipv6_acl_shared").(string)
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	ret.Inst.MigrateVip = getObjectSlbVirtualServerMigrateVip1476(d.Get("migrate_vip").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Netmask = d.Get("netmask").(string)
	ret.Inst.PortList = getSliceSlbVirtualServerPortList(d.Get("port_list").([]interface{}))
	ret.Inst.RedistributeRouteMap = d.Get("redistribute_route_map").(string)
	ret.Inst.RedistributionFlagged = d.Get("redistribution_flagged").(int)
	ret.Inst.SharedPartitionPolicyTemplate = d.Get("shared_partition_policy_template").(int)
	ret.Inst.SharedPartitionVsTemplate = d.Get("shared_partition_vs_template").(int)
	ret.Inst.StatsDataAction = d.Get("stats_data_action").(string)
	ret.Inst.SuppressInternalLoopback = d.Get("suppress_internal_loopback").(int)
	ret.Inst.TemplateLogging = d.Get("template_logging").(string)
	ret.Inst.TemplatePolicy = d.Get("template_policy").(string)
	ret.Inst.TemplatePolicyShared = d.Get("template_policy_shared").(string)
	ret.Inst.TemplateScaleout = d.Get("template_scaleout").(string)
	ret.Inst.TemplateVirtualServer = d.Get("template_virtual_server").(string)
	ret.Inst.TemplateVirtualServerShared = d.Get("template_virtual_server_shared").(string)
	ret.Inst.UseIfIp = d.Get("use_if_ip").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VportDisableAction = d.Get("vport_disable_action").(string)
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
