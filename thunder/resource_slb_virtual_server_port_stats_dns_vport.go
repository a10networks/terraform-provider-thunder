package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbVirtualServerPortStats49() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_virtual_server_port_stats_dns_vport`: Statistics for the object port\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbVirtualServerPortStats49Create,
		UpdateContext: resourceSlbVirtualServerPortStats49Update,
		ReadContext:   resourceSlbVirtualServerPortStats49Read,
		DeleteContext: resourceSlbVirtualServerPortStats49Delete,

		Schema: map[string]*schema.Schema{
			"port_number": {
				Type: schema.TypeInt, Required: true, Description: "Port",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': TCP LB service; 'udp': UDP Port; 'others': for no tcp/udp protocol, do IP load balancing; 'diameter': diameter port; 'dns-tcp': DNS service over TCP; 'dns-udp': DNS service over UDP; 'fast-http': Fast HTTP Port; 'fix': FIX Port; 'ftp': File Transfer Protocol Port; 'ftp-proxy': ftp proxy port; 'http': HTTP Port; 'https': HTTPS port; 'imap': imap proxy port; 'mlb': Message based load balancing; 'mms': Microsoft Multimedia Service Port; 'mysql': mssql port; 'mssql': mssql; 'pop3': pop3 proxy port; 'radius': RADIUS Port; 'rtsp': Real Time Streaming Protocol Port; 'sip': Session initiation protocol over UDP; 'sip-tcp': Session initiation protocol over TCP; 'sips': Session initiation protocol over TLS; 'smpp-tcp': SMPP service over TCP; 'spdy': spdy port; 'spdys': spdys port; 'smtp': SMTP Port; 'mqtt': MQTT Port; 'mqtts': MQTTS Port; 'ssl-proxy': Generic SSL proxy; 'ssli': SSL insight; 'ssh': SSH Port; 'tcp-proxy': Generic TCP proxy; 'tftp': TFTP Port; 'fast-fix': Fast FIX port; 'http-over-quic': HTTP3-over-quic port;",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_vport": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dns_total_request": {
										Type: schema.TypeInt, Optional: true, Description: "total request",
									},
									"dns_total_response": {
										Type: schema.TypeInt, Optional: true, Description: "total response",
									},
									"dns_total_drop": {
										Type: schema.TypeInt, Optional: true, Description: "total drop",
									},
									"dns_request_response": {
										Type: schema.TypeInt, Optional: true, Description: "the request response by the device",
									},
									"dns_request_send": {
										Type: schema.TypeInt, Optional: true, Description: "request send to backend server",
									},
									"dns_request_drop": {
										Type: schema.TypeInt, Optional: true, Description: "request drop",
									},
									"dns_response_drop": {
										Type: schema.TypeInt, Optional: true, Description: "response drop",
									},
									"dns_response_send": {
										Type: schema.TypeInt, Optional: true, Description: "response send to client",
									},
									"dns_request_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "request time out",
									},
									"dns_request_rexmit": {
										Type: schema.TypeInt, Optional: true, Description: "request retransmit",
									},
									"dns_cache_hit": {
										Type: schema.TypeInt, Optional: true, Description: "request cache hit",
									},
									"dnsrrl_total_dropped": {
										Type: schema.TypeInt, Optional: true, Description: "dns rrl drop",
									},
									"total_mf_dns_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "total malform dns pkt",
									},
									"total_filter_drop": {
										Type: schema.TypeInt, Optional: true, Description: "query filter drop",
									},
									"total_max_query_len_drop": {
										Type: schema.TypeInt, Optional: true, Description: "query too long drop",
									},
									"rcode_formerr_receive": {
										Type: schema.TypeInt, Optional: true, Description: "response rcode form error receive",
									},
									"rcode_serverr_receive": {
										Type: schema.TypeInt, Optional: true, Description: "response rcode server error receive",
									},
									"rcode_nxdomain_receive": {
										Type: schema.TypeInt, Optional: true, Description: "response rcode nxdomain receive",
									},
									"rcode_notimpl_receive": {
										Type: schema.TypeInt, Optional: true, Description: "response rcode type error receive",
									},
									"rcode_refuse_receive": {
										Type: schema.TypeInt, Optional: true, Description: "response rcode refuse receive",
									},
									"rcode_yxdomain_receive": {
										Type: schema.TypeInt, Optional: true, Description: "response rcode yxdomain receive",
									},
									"rcode_yxrrset_receive": {
										Type: schema.TypeInt, Optional: true, Description: "response rcode yxrrset receive",
									},
									"rcode_nxrrset_receive": {
										Type: schema.TypeInt, Optional: true, Description: "response rcode nxrrset receive",
									},
									"rcode_notauth_receive": {
										Type: schema.TypeInt, Optional: true, Description: "response rcode not auth receive",
									},
									"rcode_dsotypen_receive": {
										Type: schema.TypeInt, Optional: true, Description: "response rcode dsotypen receive",
									},
									"rcode_badver_receive": {
										Type: schema.TypeInt, Optional: true, Description: "edns response rcode bad version receive",
									},
									"rcode_badkey_receive": {
										Type: schema.TypeInt, Optional: true, Description: "edns response rcode bad key receive",
									},
									"rcode_badtime_receive": {
										Type: schema.TypeInt, Optional: true, Description: "edns response rcode bad time receive",
									},
									"rcode_badmode_receive": {
										Type: schema.TypeInt, Optional: true, Description: "ends response rcode bad name receive",
									},
									"rcode_badname_receive": {
										Type: schema.TypeInt, Optional: true, Description: "ends response rcode bad name receive",
									},
									"rcode_badalg_receive": {
										Type: schema.TypeInt, Optional: true, Description: "edns response rcode bad alg receive",
									},
									"rcode_badtranc_receive": {
										Type: schema.TypeInt, Optional: true, Description: "edns response rcode bad trancate receive",
									},
									"rcode_badcookie_receive": {
										Type: schema.TypeInt, Optional: true, Description: "ends response rcode bad cookie receive",
									},
									"rcode_other_receive": {
										Type: schema.TypeInt, Optional: true, Description: "other rcode receive",
									},
									"rcode_noerror_generate": {
										Type: schema.TypeInt, Optional: true, Description: "rcode no error generate",
									},
									"rcode_formerr_response": {
										Type: schema.TypeInt, Optional: true, Description: "rcode format error response",
									},
									"rcode_serverr_response": {
										Type: schema.TypeInt, Optional: true, Description: "rcode server error response",
									},
									"rcode_nxdomain_response": {
										Type: schema.TypeInt, Optional: true, Description: "rcode name error response",
									},
									"rcode_notimpl_response": {
										Type: schema.TypeInt, Optional: true, Description: "rcode type error response",
									},
									"rcode_refuse_response": {
										Type: schema.TypeInt, Optional: true, Description: "rcode refuse response",
									},
									"rcode_yxdomain_response": {
										Type: schema.TypeInt, Optional: true, Description: "rcode yxdomain response",
									},
									"rcode_yxrrset_response": {
										Type: schema.TypeInt, Optional: true, Description: "rcode yxrrset response",
									},
									"rcode_nxrrset_response": {
										Type: schema.TypeInt, Optional: true, Description: "rcode nxrrset response",
									},
									"rcode_notauth_response": {
										Type: schema.TypeInt, Optional: true, Description: "rcode not auth response",
									},
									"rcode_dsotypen_response": {
										Type: schema.TypeInt, Optional: true, Description: "rcode dsotypen response",
									},
									"rcode_badver_response": {
										Type: schema.TypeInt, Optional: true, Description: "edns bad var response",
									},
									"rcode_badkey_response": {
										Type: schema.TypeInt, Optional: true, Description: "ends bad key response",
									},
									"rcode_badtime_response": {
										Type: schema.TypeInt, Optional: true, Description: "edns bad time response",
									},
									"rcode_badmode_response": {
										Type: schema.TypeInt, Optional: true, Description: "edns bad mode response",
									},
									"rcode_badname_response": {
										Type: schema.TypeInt, Optional: true, Description: "edns bad name response",
									},
									"rcode_badalg_response": {
										Type: schema.TypeInt, Optional: true, Description: "ends bad alg response",
									},
									"rcode_badtranc_response": {
										Type: schema.TypeInt, Optional: true, Description: "edns bad trancate response",
									},
									"rcode_badcookie_response": {
										Type: schema.TypeInt, Optional: true, Description: "edns bad cookie response",
									},
									"rcode_other_response": {
										Type: schema.TypeInt, Optional: true, Description: "other rcode response",
									},
									"gslb_drop": {
										Type: schema.TypeInt, Optional: true, Description: "gslb drop",
									},
									"gslb_query_drop": {
										Type: schema.TypeInt, Optional: true, Description: "gslb query drop",
									},
									"gslb_query_bad": {
										Type: schema.TypeInt, Optional: true, Description: "gslb query bad",
									},
									"gslb_response_drop": {
										Type: schema.TypeInt, Optional: true, Description: "gslb response drop",
									},
									"gslb_response_bad": {
										Type: schema.TypeInt, Optional: true, Description: "gslb response bad",
									},
									"gslb_query_fwd": {
										Type: schema.TypeInt, Optional: true, Description: "gslb query forward",
									},
									"gslb_response_rvs": {
										Type: schema.TypeInt, Optional: true, Description: "gslb response reverse",
									},
									"gslb_response_good": {
										Type: schema.TypeInt, Optional: true, Description: "gslb response good",
									},
									"type_a_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns A query",
									},
									"type_aaaa_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns  AAAA query",
									},
									"type_cname_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns  CNAME query",
									},
									"type_mx_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns  MX query",
									},
									"type_ns_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns NS query",
									},
									"type_srv_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns SRV query",
									},
									"type_ptr_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns PTR query",
									},
									"type_soa_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns SOA query",
									},
									"type_txt_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns TXT query",
									},
									"type_any_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns ANY query",
									},
									"type_other_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns other type query",
									},
									"type_nsid_query": {
										Type: schema.TypeInt, Optional: true, Description: "edns NSID query",
									},
									"type_dau_query": {
										Type: schema.TypeInt, Optional: true, Description: "edns DAU query",
									},
									"type_n3u_query": {
										Type: schema.TypeInt, Optional: true, Description: "edns N3U query",
									},
									"type_expire_query": {
										Type: schema.TypeInt, Optional: true, Description: "edns EXPIRE query",
									},
									"type_cookie_query": {
										Type: schema.TypeInt, Optional: true, Description: "edns COOKIE query",
									},
									"type_keepalive_query": {
										Type: schema.TypeInt, Optional: true, Description: "edns KEEPALIVE query",
									},
									"type_padding_query": {
										Type: schema.TypeInt, Optional: true, Description: "edns PADDING query",
									},
									"type_chain_query": {
										Type: schema.TypeInt, Optional: true, Description: "edns CHAIN query",
									},
									"total_dns_filter_type_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters Total DNS Filter Type Drop",
									},
									"total_dns_filter_class_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters Total DNS Filter Class Drop",
									},
									"dns_filter_type_a_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Type A Drop",
									},
									"dns_filter_type_aaaa_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Type AAAA Drop",
									},
									"dns_filter_type_cname_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Type CNAME Drop",
									},
									"dns_filter_type_mx_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Type MX Drop",
									},
									"dns_filter_type_ns_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Type NS Drop",
									},
									"dns_filter_type_srv_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Type SRV Drop",
									},
									"dns_filter_type_ptr_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Type PTR Drop",
									},
									"dns_filter_type_soa_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Type SOA Drop",
									},
									"dns_filter_type_txt_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Type TXT Drop",
									},
									"dns_filter_type_any_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Type Any Drop",
									},
									"dns_filter_type_others_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Type OTHERS Drop",
									},
									"dns_filter_class_internet_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Class INTERNET Drop",
									},
									"dns_filter_class_chaos_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Class CHAOS Drop",
									},
									"dns_filter_class_hesiod_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Class HESIOD Drop",
									},
									"dns_filter_class_none_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Class NONE Drop",
									},
									"dns_filter_class_any_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Class ANY Drop",
									},
									"dns_filter_class_others_drop": {
										Type: schema.TypeInt, Optional: true, Description: "counters DNS Filter Class OTHER Drop",
									},
									"dns_recursive_resolution_started": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Started",
									},
									"dns_recursive_resolution_succeeded": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Succeeded",
									},
									"dns_recursive_resolution_send_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Send Failed",
									},
									"dns_recursive_resolution_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Timed Out",
									},
									"dns_recursive_resolution_retransmit_sent": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Retransmit Sent",
									},
									"dns_recursive_resolution_retransmit_exceeded": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Retransmit Exceeded",
									},
									"dns_recursive_resolution_buff_alloc_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Buffer Allocation Failed",
									},
									"dns_recursive_resolution_ongoing_client_retransmit": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Ongoing Client Retransmit",
									},
									"slb_dns_client_ssl_succ": {
										Type: schema.TypeInt, Optional: true, Description: "No. of client ssl success",
									},
									"slb_dns_server_ssl_succ": {
										Type: schema.TypeInt, Optional: true, Description: "No. of server ssl success",
									},
									"slb_dns_udp_conn": {
										Type: schema.TypeInt, Optional: true, Description: "No. of backend udp connections",
									},
									"slb_dns_udp_conn_succ": {
										Type: schema.TypeInt, Optional: true, Description: "No. of backend udp conn established",
									},
									"slb_dns_padding_to_server_removed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"slb_dns_padding_to_client_added": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"slb_dns_edns_subnet_to_server_removed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"slb_dns_udp_retransmit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"slb_dns_udp_retransmit_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dns_rpz_action_drop": {
										Type: schema.TypeInt, Optional: true, Description: "DNS RPZ Action Drop",
									},
									"dns_rpz_action_pass_thru": {
										Type: schema.TypeInt, Optional: true, Description: "DNS RPZ Action Pass Through",
									},
									"dns_rpz_action_tcp_only": {
										Type: schema.TypeInt, Optional: true, Description: "DNS RPZ Action TCP Only",
									},
									"dns_rpz_action_nxdomain": {
										Type: schema.TypeInt, Optional: true, Description: "DNS RPZ Action NXDOMAIN",
									},
									"dns_rpz_action_nodata": {
										Type: schema.TypeInt, Optional: true, Description: "DNS RPZ Action NODATA",
									},
									"dns_rpz_action_local_data": {
										Type: schema.TypeInt, Optional: true, Description: "DNS RPZ Action Local Data",
									},
									"dns_rpz_trigger_client_ip": {
										Type: schema.TypeInt, Optional: true, Description: "DNS RPZ Trigger Client IP",
									},
									"dns_rpz_trigger_resp_ip": {
										Type: schema.TypeInt, Optional: true, Description: "DNS RPZ Trigger Response IP",
									},
									"dns_rpz_trigger_ns_ip": {
										Type: schema.TypeInt, Optional: true, Description: "DNS RPZ Trigger NS IP",
									},
									"dns_rpz_trigger_qname": {
										Type: schema.TypeInt, Optional: true, Description: "DNS RPZ Trigger Qname IP",
									},
									"dns_rpz_trigger_ns_name": {
										Type: schema.TypeInt, Optional: true, Description: "DNS RPZ Trigger NS Name",
									},
									"dnsrrl_total_allowed": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Response-Rate-Limiting Total Responses Allowed",
									},
									"dnsrrl_total_slipped": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Response-Rate-Limiting Total Responses Slipped",
									},
									"dnsrrl_bad_fqdn": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Response-Rate-Limiting Bad FQDN",
									},
									"total_mf_dns_pkt_detect": {
										Type: schema.TypeInt, Optional: true, Description: "total malform dns pkt detected",
									},
									"type_rrsig_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns RRSIG query",
									},
									"type_tsig_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns  TSIG query",
									},
									"type_dnskey_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns  DNSKEY query",
									},
									"type_axfr_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns  AXFR query",
									},
									"type_ixfr_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns IXFR query",
									},
									"type_caa_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns CAA query",
									},
									"type_naptr_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns NAPTR query",
									},
									"type_ds_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns DS query",
									},
									"type_cert_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns CERT query",
									},
									"dns_recursive_resolution_not_dplane": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Not Dplane",
									},
									"dns_recursive_resolution_no_resolver": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution No Resolver",
									},
									"dns_recursive_resolution_max_trials_exceeded": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Max Trials Exceeded",
									},
									"dns_recursive_resolution_no_hints": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution No Hints",
									},
									"dns_recursive_resolution_res_submit_err": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Resolver Submit Err",
									},
									"dns_recursive_resolution_res_check_err": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Resolver Check Err",
									},
									"dns_recursive_resolution_udp_conn_err": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution UDP Conn Err",
									},
									"dns_recursive_resolution_tcp_conn_err": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution TCP Conn Err",
									},
									"dns_recursive_resolution_udp_send_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution UDP Send Failed",
									},
									"dns_recursive_resolution_icmp_err": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution ICMP Err",
									},
									"dns_recursive_resolution_query_not_sent": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Query Not Sent",
									},
									"dns_tcp_pipeline_request_drop": {
										Type: schema.TypeInt, Optional: true, Description: "DNS TCP Pipeline Request Drop",
									},
									"dns_recursive_resolution_resp_truncated": {
										Type: schema.TypeInt, Optional: true, Description: "Dns Recursive Resolution Resp Truncated",
									},
									"dns_recursive_resolution_full_response": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Serving Full Response",
									},
									"dns_full_response_from_cache": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Full Response from Cache",
									},
									"dns_recursive_resolution_missing_glue": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution missing Glue Record",
									},
									"dns_recursive_resolution_ns_cache_hit": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Hit NS Cache entry with valid glue record",
									},
									"dns_recursive_resolution_ns_cache_miss": {
										Type: schema.TypeInt, Optional: true, Description: "DNS DNS Recursive Resolution cache miss for NS Cache entry",
									},
									"dns_recursive_resolution_lookup_ip_proto_switch_46": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Lookup IP Protocol Switch 4to6",
									},
									"dns_recursive_resolution_lookup_ip_proto_switch_64": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Lookup IP Protocol Switch 6to4",
									},
									"slb_dns_edns_ecs_received": {
										Type: schema.TypeInt, Optional: true, Description: "Number of ecs from client received",
									},
									"slb_dns_edns_ecs_inserted": {
										Type: schema.TypeInt, Optional: true, Description: "Number of ecs inserted",
									},
									"slb_dns_edns_ecs_insertion_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Number of ecs insertion failure",
									},
									"dns_recursive_resolution_invalid_hints": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Invalid Hints",
									},
									"dns_recursive_resolution_pending_resolution": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Pending Resolution",
									},
									"dns_recursive_resolution_query_dropped": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Query Dropped",
									},
									"dns_recursive_resolution_respond_with_servfail": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Respond With Servfail",
									},
									"dns_recursive_resolution_total_trials_1": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Total trials upto 1",
									},
									"dns_recursive_resolution_total_trials_3": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Total trials upto 3",
									},
									"dns_recursive_resolution_total_trials_6": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Total trials upto 6",
									},
									"dns_recursive_resolution_total_trials_12": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Total trials upto 12",
									},
									"dns_recursive_resolution_total_trials_24": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Total trials upto 24",
									},
									"dns_recursive_resolution_total_trials_48": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Total trials upto 48",
									},
									"dns_recursive_resolution_total_trials_64": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Total trials upto 64",
									},
									"dns_recursive_resolution_total_trials_128": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Total trials upto 128",
									},
									"dns_recursive_resolution_total_trials_max": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Total trials upto 255",
									},
									"type_https_query": {
										Type: schema.TypeInt, Optional: true, Description: "dns HTTPS query",
									},
									"empty_response": {
										Type: schema.TypeInt, Optional: true, Description: "empty response",
									},
									"dnsrrl_total_tc": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Response-Rate-Limiting Total Responses Replied With Truncated",
									},
									"dns_negative_served": {
										Type: schema.TypeInt, Optional: true, Description: "DNS negative cache return to client",
									},
									"dns_recursive_resolution_reach_max_depth": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution Reach Max Depth",
									},
									"dns_rr_dnssec_req_received": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution DNSSEC-OK Request Received",
									},
									"dns_rr_dnssec_resp_served": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution DNSSEC-OK Response Served",
									},
									"dns_rr_dnssec_validation_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution DNSSEC Validation Failed",
									},
									"dns_rr_dnssec_val_alg_not_supported": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution DNSSEC Validation Algorithm not Supported",
									},
									"dns_rr_dnssec_val_dgst_not_supported": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution DNSSEC Validation Digest Type not Supported",
									},
									"dns_rr_dnssec_val_rrsig_signer_err": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution DNSSEC Validation RRSIG Signer's Name Error",
									},
									"dns_rr_dnssec_val_rrsig_labels_err": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution DNSSEC Validation RRSIG Labels Error",
									},
									"dns_rr_dnssec_val_rrsig_non_validity_period": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution DNSSEC Validation RRSIG non Validity Period",
									},
									"dns_rr_dnssec_val_dnskey_proto_err": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution DNSSEC Validation DNSKEY Protocol Error",
									},
									"dns_rr_dnssec_val_incorrect_sig": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution DNSSEC Validation Incorrect Signature",
									},
									"dns_rr_dnssec_val_incorrect_key_dgst": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution DNSSEC Validation Incorrect Key Digest",
									},
									"dns_rr_dnssec_val_with_trust_anchor_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution DNSSEC Validation with Trust Anchor Failed",
									},
									"dns_rr_dnssec_val_rrset_size_exceed_limit": {
										Type: schema.TypeInt, Optional: true, Description: "DNS Recursive Resolution DNSSEC Validation RRSET Size Exceed Limit",
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbVirtualServerPortStats49Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats49Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats49(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortStats49Read(ctx, d, meta)
	}
	return diags
}

func resourceSlbVirtualServerPortStats49Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats49Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats49(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortStats49Read(ctx, d, meta)
	}
	return diags
}
func resourceSlbVirtualServerPortStats49Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats49Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats49(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbVirtualServerPortStats49Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats49Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats49(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbVirtualServerPortStats49Stats(d []interface{}) edpt.SlbVirtualServerPortStats49Stats {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats49Stats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dns_vport = getObjectSlbVirtualServerPortStats49StatsDns_vport(in["dns_vport"].([]interface{}))
	}
	return ret
}

func getObjectSlbVirtualServerPortStats49StatsDns_vport(d []interface{}) edpt.SlbVirtualServerPortStats49StatsDns_vport {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats49StatsDns_vport
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dns_total_request = in["dns_total_request"].(int)
		ret.Dns_total_response = in["dns_total_response"].(int)
		ret.Dns_total_drop = in["dns_total_drop"].(int)
		ret.Dns_request_response = in["dns_request_response"].(int)
		ret.Dns_request_send = in["dns_request_send"].(int)
		ret.Dns_request_drop = in["dns_request_drop"].(int)
		ret.Dns_response_drop = in["dns_response_drop"].(int)
		ret.Dns_response_send = in["dns_response_send"].(int)
		ret.Dns_request_timeout = in["dns_request_timeout"].(int)
		ret.Dns_request_rexmit = in["dns_request_rexmit"].(int)
		ret.Dns_cache_hit = in["dns_cache_hit"].(int)
		ret.Dnsrrl_total_dropped = in["dnsrrl_total_dropped"].(int)
		ret.Total_mf_dns_pkt = in["total_mf_dns_pkt"].(int)
		ret.Total_filter_drop = in["total_filter_drop"].(int)
		ret.Total_max_query_len_drop = in["total_max_query_len_drop"].(int)
		ret.Rcode_formerr_receive = in["rcode_formerr_receive"].(int)
		ret.Rcode_serverr_receive = in["rcode_serverr_receive"].(int)
		ret.Rcode_nxdomain_receive = in["rcode_nxdomain_receive"].(int)
		ret.Rcode_notimpl_receive = in["rcode_notimpl_receive"].(int)
		ret.Rcode_refuse_receive = in["rcode_refuse_receive"].(int)
		ret.Rcode_yxdomain_receive = in["rcode_yxdomain_receive"].(int)
		ret.Rcode_yxrrset_receive = in["rcode_yxrrset_receive"].(int)
		ret.Rcode_nxrrset_receive = in["rcode_nxrrset_receive"].(int)
		ret.Rcode_notauth_receive = in["rcode_notauth_receive"].(int)
		ret.Rcode_dsotypen_receive = in["rcode_dsotypen_receive"].(int)
		ret.Rcode_badver_receive = in["rcode_badver_receive"].(int)
		ret.Rcode_badkey_receive = in["rcode_badkey_receive"].(int)
		ret.Rcode_badtime_receive = in["rcode_badtime_receive"].(int)
		ret.Rcode_badmode_receive = in["rcode_badmode_receive"].(int)
		ret.Rcode_badname_receive = in["rcode_badname_receive"].(int)
		ret.Rcode_badalg_receive = in["rcode_badalg_receive"].(int)
		ret.Rcode_badtranc_receive = in["rcode_badtranc_receive"].(int)
		ret.Rcode_badcookie_receive = in["rcode_badcookie_receive"].(int)
		ret.Rcode_other_receive = in["rcode_other_receive"].(int)
		ret.Rcode_noerror_generate = in["rcode_noerror_generate"].(int)
		ret.Rcode_formerr_response = in["rcode_formerr_response"].(int)
		ret.Rcode_serverr_response = in["rcode_serverr_response"].(int)
		ret.Rcode_nxdomain_response = in["rcode_nxdomain_response"].(int)
		ret.Rcode_notimpl_response = in["rcode_notimpl_response"].(int)
		ret.Rcode_refuse_response = in["rcode_refuse_response"].(int)
		ret.Rcode_yxdomain_response = in["rcode_yxdomain_response"].(int)
		ret.Rcode_yxrrset_response = in["rcode_yxrrset_response"].(int)
		ret.Rcode_nxrrset_response = in["rcode_nxrrset_response"].(int)
		ret.Rcode_notauth_response = in["rcode_notauth_response"].(int)
		ret.Rcode_dsotypen_response = in["rcode_dsotypen_response"].(int)
		ret.Rcode_badver_response = in["rcode_badver_response"].(int)
		ret.Rcode_badkey_response = in["rcode_badkey_response"].(int)
		ret.Rcode_badtime_response = in["rcode_badtime_response"].(int)
		ret.Rcode_badmode_response = in["rcode_badmode_response"].(int)
		ret.Rcode_badname_response = in["rcode_badname_response"].(int)
		ret.Rcode_badalg_response = in["rcode_badalg_response"].(int)
		ret.Rcode_badtranc_response = in["rcode_badtranc_response"].(int)
		ret.Rcode_badcookie_response = in["rcode_badcookie_response"].(int)
		ret.Rcode_other_response = in["rcode_other_response"].(int)
		ret.Gslb_drop = in["gslb_drop"].(int)
		ret.Gslb_query_drop = in["gslb_query_drop"].(int)
		ret.Gslb_query_bad = in["gslb_query_bad"].(int)
		ret.Gslb_response_drop = in["gslb_response_drop"].(int)
		ret.Gslb_response_bad = in["gslb_response_bad"].(int)
		ret.Gslb_query_fwd = in["gslb_query_fwd"].(int)
		ret.Gslb_response_rvs = in["gslb_response_rvs"].(int)
		ret.Gslb_response_good = in["gslb_response_good"].(int)
		ret.Type_a_query = in["type_a_query"].(int)
		ret.Type_aaaa_query = in["type_aaaa_query"].(int)
		ret.Type_cname_query = in["type_cname_query"].(int)
		ret.Type_mx_query = in["type_mx_query"].(int)
		ret.Type_ns_query = in["type_ns_query"].(int)
		ret.Type_srv_query = in["type_srv_query"].(int)
		ret.Type_ptr_query = in["type_ptr_query"].(int)
		ret.Type_soa_query = in["type_soa_query"].(int)
		ret.Type_txt_query = in["type_txt_query"].(int)
		ret.Type_any_query = in["type_any_query"].(int)
		ret.Type_other_query = in["type_other_query"].(int)
		ret.Type_nsid_query = in["type_nsid_query"].(int)
		ret.Type_dau_query = in["type_dau_query"].(int)
		ret.Type_n3u_query = in["type_n3u_query"].(int)
		ret.Type_expire_query = in["type_expire_query"].(int)
		ret.Type_cookie_query = in["type_cookie_query"].(int)
		ret.Type_keepalive_query = in["type_keepalive_query"].(int)
		ret.Type_padding_query = in["type_padding_query"].(int)
		ret.Type_chain_query = in["type_chain_query"].(int)
		ret.Total_dns_filter_type_drop = in["total_dns_filter_type_drop"].(int)
		ret.Total_dns_filter_class_drop = in["total_dns_filter_class_drop"].(int)
		ret.Dns_filter_type_a_drop = in["dns_filter_type_a_drop"].(int)
		ret.Dns_filter_type_aaaa_drop = in["dns_filter_type_aaaa_drop"].(int)
		ret.Dns_filter_type_cname_drop = in["dns_filter_type_cname_drop"].(int)
		ret.Dns_filter_type_mx_drop = in["dns_filter_type_mx_drop"].(int)
		ret.Dns_filter_type_ns_drop = in["dns_filter_type_ns_drop"].(int)
		ret.Dns_filter_type_srv_drop = in["dns_filter_type_srv_drop"].(int)
		ret.Dns_filter_type_ptr_drop = in["dns_filter_type_ptr_drop"].(int)
		ret.Dns_filter_type_soa_drop = in["dns_filter_type_soa_drop"].(int)
		ret.Dns_filter_type_txt_drop = in["dns_filter_type_txt_drop"].(int)
		ret.Dns_filter_type_any_drop = in["dns_filter_type_any_drop"].(int)
		ret.Dns_filter_type_others_drop = in["dns_filter_type_others_drop"].(int)
		ret.Dns_filter_class_internet_drop = in["dns_filter_class_internet_drop"].(int)
		ret.Dns_filter_class_chaos_drop = in["dns_filter_class_chaos_drop"].(int)
		ret.Dns_filter_class_hesiod_drop = in["dns_filter_class_hesiod_drop"].(int)
		ret.Dns_filter_class_none_drop = in["dns_filter_class_none_drop"].(int)
		ret.Dns_filter_class_any_drop = in["dns_filter_class_any_drop"].(int)
		ret.Dns_filter_class_others_drop = in["dns_filter_class_others_drop"].(int)
		ret.Dns_recursive_resolution_started = in["dns_recursive_resolution_started"].(int)
		ret.Dns_recursive_resolution_succeeded = in["dns_recursive_resolution_succeeded"].(int)
		ret.Dns_recursive_resolution_send_failed = in["dns_recursive_resolution_send_failed"].(int)
		ret.Dns_recursive_resolution_timeout = in["dns_recursive_resolution_timeout"].(int)
		ret.Dns_recursive_resolution_retransmit_sent = in["dns_recursive_resolution_retransmit_sent"].(int)
		ret.Dns_recursive_resolution_retransmit_exceeded = in["dns_recursive_resolution_retransmit_exceeded"].(int)
		ret.Dns_recursive_resolution_buff_alloc_failed = in["dns_recursive_resolution_buff_alloc_failed"].(int)
		ret.Dns_recursive_resolution_ongoing_client_retransmit = in["dns_recursive_resolution_ongoing_client_retransmit"].(int)
		ret.Slb_dns_client_ssl_succ = in["slb_dns_client_ssl_succ"].(int)
		ret.Slb_dns_server_ssl_succ = in["slb_dns_server_ssl_succ"].(int)
		ret.Slb_dns_udp_conn = in["slb_dns_udp_conn"].(int)
		ret.Slb_dns_udp_conn_succ = in["slb_dns_udp_conn_succ"].(int)
		ret.Slb_dns_padding_to_server_removed = in["slb_dns_padding_to_server_removed"].(int)
		ret.Slb_dns_padding_to_client_added = in["slb_dns_padding_to_client_added"].(int)
		ret.Slb_dns_edns_subnet_to_server_removed = in["slb_dns_edns_subnet_to_server_removed"].(int)
		ret.Slb_dns_udp_retransmit = in["slb_dns_udp_retransmit"].(int)
		ret.Slb_dns_udp_retransmit_fail = in["slb_dns_udp_retransmit_fail"].(int)
		ret.Dns_rpz_action_drop = in["dns_rpz_action_drop"].(int)
		ret.Dns_rpz_action_pass_thru = in["dns_rpz_action_pass_thru"].(int)
		ret.Dns_rpz_action_tcp_only = in["dns_rpz_action_tcp_only"].(int)
		ret.Dns_rpz_action_nxdomain = in["dns_rpz_action_nxdomain"].(int)
		ret.Dns_rpz_action_nodata = in["dns_rpz_action_nodata"].(int)
		ret.Dns_rpz_action_local_data = in["dns_rpz_action_local_data"].(int)
		ret.Dns_rpz_trigger_client_ip = in["dns_rpz_trigger_client_ip"].(int)
		ret.Dns_rpz_trigger_resp_ip = in["dns_rpz_trigger_resp_ip"].(int)
		ret.Dns_rpz_trigger_ns_ip = in["dns_rpz_trigger_ns_ip"].(int)
		ret.Dns_rpz_trigger_qname = in["dns_rpz_trigger_qname"].(int)
		ret.Dns_rpz_trigger_ns_name = in["dns_rpz_trigger_ns_name"].(int)
		ret.Dnsrrl_total_allowed = in["dnsrrl_total_allowed"].(int)
		ret.Dnsrrl_total_slipped = in["dnsrrl_total_slipped"].(int)
		ret.Dnsrrl_bad_fqdn = in["dnsrrl_bad_fqdn"].(int)
		ret.Total_mf_dns_pkt_detect = in["total_mf_dns_pkt_detect"].(int)
		ret.Type_rrsig_query = in["type_rrsig_query"].(int)
		ret.Type_tsig_query = in["type_tsig_query"].(int)
		ret.Type_dnskey_query = in["type_dnskey_query"].(int)
		ret.Type_axfr_query = in["type_axfr_query"].(int)
		ret.Type_ixfr_query = in["type_ixfr_query"].(int)
		ret.Type_caa_query = in["type_caa_query"].(int)
		ret.Type_naptr_query = in["type_naptr_query"].(int)
		ret.Type_ds_query = in["type_ds_query"].(int)
		ret.Type_cert_query = in["type_cert_query"].(int)
		ret.Dns_recursive_resolution_not_dplane = in["dns_recursive_resolution_not_dplane"].(int)
		ret.Dns_recursive_resolution_no_resolver = in["dns_recursive_resolution_no_resolver"].(int)
		ret.Dns_recursive_resolution_max_trials_exceeded = in["dns_recursive_resolution_max_trials_exceeded"].(int)
		ret.Dns_recursive_resolution_no_hints = in["dns_recursive_resolution_no_hints"].(int)
		ret.Dns_recursive_resolution_res_submit_err = in["dns_recursive_resolution_res_submit_err"].(int)
		ret.Dns_recursive_resolution_res_check_err = in["dns_recursive_resolution_res_check_err"].(int)
		ret.Dns_recursive_resolution_udp_conn_err = in["dns_recursive_resolution_udp_conn_err"].(int)
		ret.Dns_recursive_resolution_tcp_conn_err = in["dns_recursive_resolution_tcp_conn_err"].(int)
		ret.Dns_recursive_resolution_udp_send_failed = in["dns_recursive_resolution_udp_send_failed"].(int)
		ret.Dns_recursive_resolution_icmp_err = in["dns_recursive_resolution_icmp_err"].(int)
		ret.Dns_recursive_resolution_query_not_sent = in["dns_recursive_resolution_query_not_sent"].(int)
		ret.Dns_tcp_pipeline_request_drop = in["dns_tcp_pipeline_request_drop"].(int)
		ret.Dns_recursive_resolution_resp_truncated = in["dns_recursive_resolution_resp_truncated"].(int)
		ret.Dns_recursive_resolution_full_response = in["dns_recursive_resolution_full_response"].(int)
		ret.Dns_full_response_from_cache = in["dns_full_response_from_cache"].(int)
		ret.Dns_recursive_resolution_missing_glue = in["dns_recursive_resolution_missing_glue"].(int)
		ret.Dns_recursive_resolution_ns_cache_hit = in["dns_recursive_resolution_ns_cache_hit"].(int)
		ret.Dns_recursive_resolution_ns_cache_miss = in["dns_recursive_resolution_ns_cache_miss"].(int)
		ret.Dns_recursive_resolution_lookup_ip_proto_switch_46 = in["dns_recursive_resolution_lookup_ip_proto_switch_46"].(int)
		ret.Dns_recursive_resolution_lookup_ip_proto_switch_64 = in["dns_recursive_resolution_lookup_ip_proto_switch_64"].(int)
		ret.Slb_dns_edns_ecs_received = in["slb_dns_edns_ecs_received"].(int)
		ret.Slb_dns_edns_ecs_inserted = in["slb_dns_edns_ecs_inserted"].(int)
		ret.Slb_dns_edns_ecs_insertion_fail = in["slb_dns_edns_ecs_insertion_fail"].(int)
		ret.Dns_recursive_resolution_invalid_hints = in["dns_recursive_resolution_invalid_hints"].(int)
		ret.Dns_recursive_resolution_pending_resolution = in["dns_recursive_resolution_pending_resolution"].(int)
		ret.Dns_recursive_resolution_query_dropped = in["dns_recursive_resolution_query_dropped"].(int)
		ret.Dns_recursive_resolution_respond_with_servfail = in["dns_recursive_resolution_respond_with_servfail"].(int)
		ret.Dns_recursive_resolution_total_trials_1 = in["dns_recursive_resolution_total_trials_1"].(int)
		ret.Dns_recursive_resolution_total_trials_3 = in["dns_recursive_resolution_total_trials_3"].(int)
		ret.Dns_recursive_resolution_total_trials_6 = in["dns_recursive_resolution_total_trials_6"].(int)
		ret.Dns_recursive_resolution_total_trials_12 = in["dns_recursive_resolution_total_trials_12"].(int)
		ret.Dns_recursive_resolution_total_trials_24 = in["dns_recursive_resolution_total_trials_24"].(int)
		ret.Dns_recursive_resolution_total_trials_48 = in["dns_recursive_resolution_total_trials_48"].(int)
		ret.Dns_recursive_resolution_total_trials_64 = in["dns_recursive_resolution_total_trials_64"].(int)
		ret.Dns_recursive_resolution_total_trials_128 = in["dns_recursive_resolution_total_trials_128"].(int)
		ret.Dns_recursive_resolution_total_trials_max = in["dns_recursive_resolution_total_trials_max"].(int)
		ret.Type_https_query = in["type_https_query"].(int)
		ret.Empty_response = in["empty_response"].(int)
		ret.Dnsrrl_total_tc = in["dnsrrl_total_tc"].(int)
		ret.Dns_negative_served = in["dns_negative_served"].(int)
		ret.Dns_recursive_resolution_reach_max_depth = in["dns_recursive_resolution_reach_max_depth"].(int)
		ret.Dns_rr_dnssec_req_received = in["dns_rr_dnssec_req_received"].(int)
		ret.Dns_rr_dnssec_resp_served = in["dns_rr_dnssec_resp_served"].(int)
		ret.Dns_rr_dnssec_validation_failed = in["dns_rr_dnssec_validation_failed"].(int)
		ret.Dns_rr_dnssec_val_alg_not_supported = in["dns_rr_dnssec_val_alg_not_supported"].(int)
		ret.Dns_rr_dnssec_val_dgst_not_supported = in["dns_rr_dnssec_val_dgst_not_supported"].(int)
		ret.Dns_rr_dnssec_val_rrsig_signer_err = in["dns_rr_dnssec_val_rrsig_signer_err"].(int)
		ret.Dns_rr_dnssec_val_rrsig_labels_err = in["dns_rr_dnssec_val_rrsig_labels_err"].(int)
		ret.Dns_rr_dnssec_val_rrsig_non_validity_period = in["dns_rr_dnssec_val_rrsig_non_validity_period"].(int)
		ret.Dns_rr_dnssec_val_dnskey_proto_err = in["dns_rr_dnssec_val_dnskey_proto_err"].(int)
		ret.Dns_rr_dnssec_val_incorrect_sig = in["dns_rr_dnssec_val_incorrect_sig"].(int)
		ret.Dns_rr_dnssec_val_incorrect_key_dgst = in["dns_rr_dnssec_val_incorrect_key_dgst"].(int)
		ret.Dns_rr_dnssec_val_with_trust_anchor_failed = in["dns_rr_dnssec_val_with_trust_anchor_failed"].(int)
		ret.Dns_rr_dnssec_val_rrset_size_exceed_limit = in["dns_rr_dnssec_val_rrset_size_exceed_limit"].(int)
	}
	return ret
}

func dataToEndpointSlbVirtualServerPortStats49(d *schema.ResourceData) edpt.SlbVirtualServerPortStats49 {
	var ret edpt.SlbVirtualServerPortStats49
	ret.Inst.PortNumber = d.Get("port_number").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Stats = getObjectSlbVirtualServerPortStats49Stats(d.Get("stats").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
