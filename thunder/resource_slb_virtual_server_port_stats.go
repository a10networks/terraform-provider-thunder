package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbVirtualServerPortStats56() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_virtual_server_port_stats`: Statistics for the object port\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbVirtualServerPortStats56Read,

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
						"curr_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Current established connections",
						},
						"total_l4_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total L4 connections established",
						},
						"total_l7_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total L7 connections established",
						},
						"total_tcp_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP connections established",
						},
						"total_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total connections established",
						},
						"total_fwd_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes processed in forward direction",
						},
						"total_fwd_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Packets processed in forward direction",
						},
						"total_rev_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes processed in reverse direction",
						},
						"total_rev_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Packets processed in reverse direction",
						},
						"total_dns_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Total DNS packets processed",
						},
						"total_mf_dns_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Total MF DNS packets",
						},
						"es_total_failure_actions": {
							Type: schema.TypeInt, Optional: true, Description: "Total failure actions",
						},
						"compression_bytes_before": {
							Type: schema.TypeInt, Optional: true, Description: "Data into gzip compression engine",
						},
						"compression_bytes_after": {
							Type: schema.TypeInt, Optional: true, Description: "Data out of gzip compression engine",
						},
						"compression_hit": {
							Type: schema.TypeInt, Optional: true, Description: "Number of requests compressed",
						},
						"compression_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Number of requests NOT compressed",
						},
						"compression_miss_no_client": {
							Type: schema.TypeInt, Optional: true, Description: "Compression miss no client",
						},
						"compression_miss_template_exclusion": {
							Type: schema.TypeInt, Optional: true, Description: "Compression miss template exclusion",
						},
						"curr_req": {
							Type: schema.TypeInt, Optional: true, Description: "Current requests",
						},
						"total_req": {
							Type: schema.TypeInt, Optional: true, Description: "Total requests",
						},
						"total_req_succ": {
							Type: schema.TypeInt, Optional: true, Description: "Total successful requests",
						},
						"peak_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Peak connections",
						},
						"curr_conn_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Current connection rate",
						},
						"last_rsp_time": {
							Type: schema.TypeInt, Optional: true, Description: "Last response time",
						},
						"fastest_rsp_time": {
							Type: schema.TypeInt, Optional: true, Description: "Fastest response time",
						},
						"slowest_rsp_time": {
							Type: schema.TypeInt, Optional: true, Description: "Slowest response time",
						},
						"loc_permit": {
							Type: schema.TypeInt, Optional: true, Description: "Geo-location Permit count",
						},
						"loc_deny": {
							Type: schema.TypeInt, Optional: true, Description: "Geo-location Deny count",
						},
						"loc_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Geo-location Connection count",
						},
						"curr_ssl_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Current SSL connections",
						},
						"total_ssl_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total SSL connections",
						},
						"backend_time_to_first_byte": {
							Type: schema.TypeInt, Optional: true, Description: "Backend Time from Request to Response First Byte",
						},
						"backend_time_to_last_byte": {
							Type: schema.TypeInt, Optional: true, Description: "Backend Time from Request to Response Last Byte",
						},
						"in_latency": {
							Type: schema.TypeInt, Optional: true, Description: "Request Latency at Thunder",
						},
						"out_latency": {
							Type: schema.TypeInt, Optional: true, Description: "Response Latency at Thunder",
						},
						"total_fwd_bytes_out": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes processed in forward direction (outbound)",
						},
						"total_fwd_pkts_out": {
							Type: schema.TypeInt, Optional: true, Description: "Packets processed in forward direction (outbound)",
						},
						"total_rev_bytes_out": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes processed in reverse direction (outbound)",
						},
						"total_rev_pkts_out": {
							Type: schema.TypeInt, Optional: true, Description: "Packets processed in reverse direction (outbound)",
						},
						"curr_req_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Current request rate",
						},
						"curr_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Current response",
						},
						"total_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Total response",
						},
						"total_resp_succ": {
							Type: schema.TypeInt, Optional: true, Description: "Total successful responses",
						},
						"curr_resp_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Current response rate",
						},
						"dnsrrl_total_allowed": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Response-Rate-Limiting Total Responses Allowed",
						},
						"dnsrrl_total_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Response-Rate-Limiting Total Responses Dropped",
						},
						"dnsrrl_total_slipped": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Response-Rate-Limiting Total Responses Slipped",
						},
						"dnsrrl_bad_fqdn": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Response-Rate-Limiting Bad FQDN",
						},
						"throughput_bits_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "Throughput in bits/sec",
						},
						"dynamic_memory": {
							Type: schema.TypeInt, Optional: true, Description: "dynamic memory (bytes) used by the vport(alloc-free)",
						},
						"ip_only_lb_fwd_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "IP-Only-LB Bytes processed in forward direction",
						},
						"ip_only_lb_rev_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "IP-Only-LB Bytes processed in reverse direction",
						},
						"ip_only_lb_fwd_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "IP-Only-LB Packets processed in forward direction",
						},
						"ip_only_lb_rev_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "IP-Only-LB Packets processed in reverse direction",
						},
						"total_dns_filter_type_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total DNS Filter Type Drop",
						},
						"total_dns_filter_class_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total DNS Filter Class Drop",
						},
						"dns_filter_type_a_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Type A Drop",
						},
						"dns_filter_type_aaaa_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Type AAAA Drop",
						},
						"dns_filter_type_cname_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Type CNAME Drop",
						},
						"dns_filter_type_mx_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Type MX Drop",
						},
						"dns_filter_type_ns_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Type NS Drop",
						},
						"dns_filter_type_srv_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Type SRV Drop",
						},
						"dns_filter_type_ptr_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Type PTR Drop",
						},
						"dns_filter_type_soa_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Type SOA Drop",
						},
						"dns_filter_type_txt_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Type TXT Drop",
						},
						"dns_filter_type_any_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Type Any Drop",
						},
						"dns_filter_type_others_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Type OTHERS Drop",
						},
						"dns_filter_class_internet_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Class INTERNET Drop",
						},
						"dns_filter_class_chaos_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Class CHAOS Drop",
						},
						"dns_filter_class_hesiod_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Class HESIOD Drop",
						},
						"dns_filter_class_none_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Class NONE Drop",
						},
						"dns_filter_class_any_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Class ANY Drop",
						},
						"dns_filter_class_others_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Filter Class OTHERS Drop",
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
						"compression_bytes_before_br": {
							Type: schema.TypeInt, Optional: true, Description: "Data into brotli compression engine",
						},
						"compression_bytes_after_br": {
							Type: schema.TypeInt, Optional: true, Description: "Data out of brotli compression engine",
						},
						"compression_bytes_before_total": {
							Type: schema.TypeInt, Optional: true, Description: "Data into compression engine",
						},
						"compression_bytes_after_total": {
							Type: schema.TypeInt, Optional: true, Description: "Data out of compression engine",
						},
						"compression_hit_br": {
							Type: schema.TypeInt, Optional: true, Description: "Number of requests compressed with brotli",
						},
						"compression_miss_br": {
							Type: schema.TypeInt, Optional: true, Description: "Number of requests NOT compressed with brotli",
						},
						"compression_hit_total": {
							Type: schema.TypeInt, Optional: true, Description: "Number of requests compressed",
						},
						"compression_miss_total": {
							Type: schema.TypeInt, Optional: true, Description: "Number of requests NOT compressed",
						},
						"dnsrrl_total_tc": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Response-Rate-Limiting Total Responses Replied With Truncated",
						},
						"http1_client_idle_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP1 Client Idle Timeout",
						},
						"http2_client_idle_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP2 Client Idle Timeout",
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

func resourceSlbVirtualServerPortStats56Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats56Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats56(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbVirtualServerPortStats56Stats := setObjectSlbVirtualServerPortStats56Stats(res)
		d.Set("stats", SlbVirtualServerPortStats56Stats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbVirtualServerPortStats56Stats(ret edpt.DataSlbVirtualServerPortStats56) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_conn":                           ret.DtSlbVirtualServerPortStats56.Stats.Curr_conn,
			"total_l4_conn":                       ret.DtSlbVirtualServerPortStats56.Stats.Total_l4_conn,
			"total_l7_conn":                       ret.DtSlbVirtualServerPortStats56.Stats.Total_l7_conn,
			"total_tcp_conn":                      ret.DtSlbVirtualServerPortStats56.Stats.Total_tcp_conn,
			"total_conn":                          ret.DtSlbVirtualServerPortStats56.Stats.Total_conn,
			"total_fwd_bytes":                     ret.DtSlbVirtualServerPortStats56.Stats.Total_fwd_bytes,
			"total_fwd_pkts":                      ret.DtSlbVirtualServerPortStats56.Stats.Total_fwd_pkts,
			"total_rev_bytes":                     ret.DtSlbVirtualServerPortStats56.Stats.Total_rev_bytes,
			"total_rev_pkts":                      ret.DtSlbVirtualServerPortStats56.Stats.Total_rev_pkts,
			"total_dns_pkts":                      ret.DtSlbVirtualServerPortStats56.Stats.Total_dns_pkts,
			"total_mf_dns_pkts":                   ret.DtSlbVirtualServerPortStats56.Stats.Total_mf_dns_pkts,
			"es_total_failure_actions":            ret.DtSlbVirtualServerPortStats56.Stats.Es_total_failure_actions,
			"compression_bytes_before":            ret.DtSlbVirtualServerPortStats56.Stats.Compression_bytes_before,
			"compression_bytes_after":             ret.DtSlbVirtualServerPortStats56.Stats.Compression_bytes_after,
			"compression_hit":                     ret.DtSlbVirtualServerPortStats56.Stats.Compression_hit,
			"compression_miss":                    ret.DtSlbVirtualServerPortStats56.Stats.Compression_miss,
			"compression_miss_no_client":          ret.DtSlbVirtualServerPortStats56.Stats.Compression_miss_no_client,
			"compression_miss_template_exclusion": ret.DtSlbVirtualServerPortStats56.Stats.Compression_miss_template_exclusion,
			"curr_req":                            ret.DtSlbVirtualServerPortStats56.Stats.Curr_req,
			"total_req":                           ret.DtSlbVirtualServerPortStats56.Stats.Total_req,
			"total_req_succ":                      ret.DtSlbVirtualServerPortStats56.Stats.Total_req_succ,
			"peak_conn":                           ret.DtSlbVirtualServerPortStats56.Stats.Peak_conn,
			"curr_conn_rate":                      ret.DtSlbVirtualServerPortStats56.Stats.Curr_conn_rate,
			"last_rsp_time":                       ret.DtSlbVirtualServerPortStats56.Stats.Last_rsp_time,
			"fastest_rsp_time":                    ret.DtSlbVirtualServerPortStats56.Stats.Fastest_rsp_time,
			"slowest_rsp_time":                    ret.DtSlbVirtualServerPortStats56.Stats.Slowest_rsp_time,
			"loc_permit":                          ret.DtSlbVirtualServerPortStats56.Stats.Loc_permit,
			"loc_deny":                            ret.DtSlbVirtualServerPortStats56.Stats.Loc_deny,
			"loc_conn":                            ret.DtSlbVirtualServerPortStats56.Stats.Loc_conn,
			"curr_ssl_conn":                       ret.DtSlbVirtualServerPortStats56.Stats.Curr_ssl_conn,
			"total_ssl_conn":                      ret.DtSlbVirtualServerPortStats56.Stats.Total_ssl_conn,
			"backend_time_to_first_byte":          ret.DtSlbVirtualServerPortStats56.Stats.BackendTimeToFirstByte,
			"backend_time_to_last_byte":           ret.DtSlbVirtualServerPortStats56.Stats.BackendTimeToLastByte,
			"in_latency":                          ret.DtSlbVirtualServerPortStats56.Stats.InLatency,
			"out_latency":                         ret.DtSlbVirtualServerPortStats56.Stats.OutLatency,
			"total_fwd_bytes_out":                 ret.DtSlbVirtualServerPortStats56.Stats.Total_fwd_bytes_out,
			"total_fwd_pkts_out":                  ret.DtSlbVirtualServerPortStats56.Stats.Total_fwd_pkts_out,
			"total_rev_bytes_out":                 ret.DtSlbVirtualServerPortStats56.Stats.Total_rev_bytes_out,
			"total_rev_pkts_out":                  ret.DtSlbVirtualServerPortStats56.Stats.Total_rev_pkts_out,
			"curr_req_rate":                       ret.DtSlbVirtualServerPortStats56.Stats.Curr_req_rate,
			"curr_resp":                           ret.DtSlbVirtualServerPortStats56.Stats.Curr_resp,
			"total_resp":                          ret.DtSlbVirtualServerPortStats56.Stats.Total_resp,
			"total_resp_succ":                     ret.DtSlbVirtualServerPortStats56.Stats.Total_resp_succ,
			"curr_resp_rate":                      ret.DtSlbVirtualServerPortStats56.Stats.Curr_resp_rate,
			"dnsrrl_total_allowed":                ret.DtSlbVirtualServerPortStats56.Stats.Dnsrrl_total_allowed,
			"dnsrrl_total_dropped":                ret.DtSlbVirtualServerPortStats56.Stats.Dnsrrl_total_dropped,
			"dnsrrl_total_slipped":                ret.DtSlbVirtualServerPortStats56.Stats.Dnsrrl_total_slipped,
			"dnsrrl_bad_fqdn":                     ret.DtSlbVirtualServerPortStats56.Stats.Dnsrrl_bad_fqdn,
			"throughput_bits_per_sec":             ret.DtSlbVirtualServerPortStats56.Stats.ThroughputBitsPerSec,
			"dynamic_memory":                      ret.DtSlbVirtualServerPortStats56.Stats.DynamicMemory,
			"ip_only_lb_fwd_bytes":                ret.DtSlbVirtualServerPortStats56.Stats.Ip_only_lb_fwd_bytes,
			"ip_only_lb_rev_bytes":                ret.DtSlbVirtualServerPortStats56.Stats.Ip_only_lb_rev_bytes,
			"ip_only_lb_fwd_pkts":                 ret.DtSlbVirtualServerPortStats56.Stats.Ip_only_lb_fwd_pkts,
			"ip_only_lb_rev_pkts":                 ret.DtSlbVirtualServerPortStats56.Stats.Ip_only_lb_rev_pkts,
			"total_dns_filter_type_drop":          ret.DtSlbVirtualServerPortStats56.Stats.Total_dns_filter_type_drop,
			"total_dns_filter_class_drop":         ret.DtSlbVirtualServerPortStats56.Stats.Total_dns_filter_class_drop,
			"dns_filter_type_a_drop":              ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_type_a_drop,
			"dns_filter_type_aaaa_drop":           ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_type_aaaa_drop,
			"dns_filter_type_cname_drop":          ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_type_cname_drop,
			"dns_filter_type_mx_drop":             ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_type_mx_drop,
			"dns_filter_type_ns_drop":             ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_type_ns_drop,
			"dns_filter_type_srv_drop":            ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_type_srv_drop,
			"dns_filter_type_ptr_drop":            ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_type_ptr_drop,
			"dns_filter_type_soa_drop":            ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_type_soa_drop,
			"dns_filter_type_txt_drop":            ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_type_txt_drop,
			"dns_filter_type_any_drop":            ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_type_any_drop,
			"dns_filter_type_others_drop":         ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_type_others_drop,
			"dns_filter_class_internet_drop":      ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_class_internet_drop,
			"dns_filter_class_chaos_drop":         ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_class_chaos_drop,
			"dns_filter_class_hesiod_drop":        ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_class_hesiod_drop,
			"dns_filter_class_none_drop":          ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_class_none_drop,
			"dns_filter_class_any_drop":           ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_class_any_drop,
			"dns_filter_class_others_drop":        ret.DtSlbVirtualServerPortStats56.Stats.Dns_filter_class_others_drop,
			"dns_rpz_action_drop":                 ret.DtSlbVirtualServerPortStats56.Stats.Dns_rpz_action_drop,
			"dns_rpz_action_pass_thru":            ret.DtSlbVirtualServerPortStats56.Stats.Dns_rpz_action_pass_thru,
			"dns_rpz_action_tcp_only":             ret.DtSlbVirtualServerPortStats56.Stats.Dns_rpz_action_tcp_only,
			"dns_rpz_action_nxdomain":             ret.DtSlbVirtualServerPortStats56.Stats.Dns_rpz_action_nxdomain,
			"dns_rpz_action_nodata":               ret.DtSlbVirtualServerPortStats56.Stats.Dns_rpz_action_nodata,
			"dns_rpz_action_local_data":           ret.DtSlbVirtualServerPortStats56.Stats.Dns_rpz_action_local_data,
			"dns_rpz_trigger_client_ip":           ret.DtSlbVirtualServerPortStats56.Stats.Dns_rpz_trigger_client_ip,
			"dns_rpz_trigger_resp_ip":             ret.DtSlbVirtualServerPortStats56.Stats.Dns_rpz_trigger_resp_ip,
			"dns_rpz_trigger_ns_ip":               ret.DtSlbVirtualServerPortStats56.Stats.Dns_rpz_trigger_ns_ip,
			"dns_rpz_trigger_qname":               ret.DtSlbVirtualServerPortStats56.Stats.Dns_rpz_trigger_qname,
			"dns_rpz_trigger_ns_name":             ret.DtSlbVirtualServerPortStats56.Stats.Dns_rpz_trigger_ns_name,
			"compression_bytes_before_br":         ret.DtSlbVirtualServerPortStats56.Stats.Compression_bytes_before_br,
			"compression_bytes_after_br":          ret.DtSlbVirtualServerPortStats56.Stats.Compression_bytes_after_br,
			"compression_bytes_before_total":      ret.DtSlbVirtualServerPortStats56.Stats.Compression_bytes_before_total,
			"compression_bytes_after_total":       ret.DtSlbVirtualServerPortStats56.Stats.Compression_bytes_after_total,
			"compression_hit_br":                  ret.DtSlbVirtualServerPortStats56.Stats.Compression_hit_br,
			"compression_miss_br":                 ret.DtSlbVirtualServerPortStats56.Stats.Compression_miss_br,
			"compression_hit_total":               ret.DtSlbVirtualServerPortStats56.Stats.Compression_hit_total,
			"compression_miss_total":              ret.DtSlbVirtualServerPortStats56.Stats.Compression_miss_total,
			"dnsrrl_total_tc":                     ret.DtSlbVirtualServerPortStats56.Stats.Dnsrrl_total_tc,
			"http1_client_idle_timeout":           ret.DtSlbVirtualServerPortStats56.Stats.Http1_client_idle_timeout,
			"http2_client_idle_timeout":           ret.DtSlbVirtualServerPortStats56.Stats.Http2_client_idle_timeout,
		},
	}
}

func getObjectSlbVirtualServerPortStats56Stats(d []interface{}) edpt.SlbVirtualServerPortStats56Stats {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats56Stats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_conn = in["curr_conn"].(int)
		ret.Total_l4_conn = in["total_l4_conn"].(int)
		ret.Total_l7_conn = in["total_l7_conn"].(int)
		ret.Total_tcp_conn = in["total_tcp_conn"].(int)
		ret.Total_conn = in["total_conn"].(int)
		ret.Total_fwd_bytes = in["total_fwd_bytes"].(int)
		ret.Total_fwd_pkts = in["total_fwd_pkts"].(int)
		ret.Total_rev_bytes = in["total_rev_bytes"].(int)
		ret.Total_rev_pkts = in["total_rev_pkts"].(int)
		ret.Total_dns_pkts = in["total_dns_pkts"].(int)
		ret.Total_mf_dns_pkts = in["total_mf_dns_pkts"].(int)
		ret.Es_total_failure_actions = in["es_total_failure_actions"].(int)
		ret.Compression_bytes_before = in["compression_bytes_before"].(int)
		ret.Compression_bytes_after = in["compression_bytes_after"].(int)
		ret.Compression_hit = in["compression_hit"].(int)
		ret.Compression_miss = in["compression_miss"].(int)
		ret.Compression_miss_no_client = in["compression_miss_no_client"].(int)
		ret.Compression_miss_template_exclusion = in["compression_miss_template_exclusion"].(int)
		ret.Curr_req = in["curr_req"].(int)
		ret.Total_req = in["total_req"].(int)
		ret.Total_req_succ = in["total_req_succ"].(int)
		ret.Peak_conn = in["peak_conn"].(int)
		ret.Curr_conn_rate = in["curr_conn_rate"].(int)
		ret.Last_rsp_time = in["last_rsp_time"].(int)
		ret.Fastest_rsp_time = in["fastest_rsp_time"].(int)
		ret.Slowest_rsp_time = in["slowest_rsp_time"].(int)
		ret.Loc_permit = in["loc_permit"].(int)
		ret.Loc_deny = in["loc_deny"].(int)
		ret.Loc_conn = in["loc_conn"].(int)
		ret.Curr_ssl_conn = in["curr_ssl_conn"].(int)
		ret.Total_ssl_conn = in["total_ssl_conn"].(int)
		ret.BackendTimeToFirstByte = in["backend_time_to_first_byte"].(int)
		ret.BackendTimeToLastByte = in["backend_time_to_last_byte"].(int)
		ret.InLatency = in["in_latency"].(int)
		ret.OutLatency = in["out_latency"].(int)
		ret.Total_fwd_bytes_out = in["total_fwd_bytes_out"].(int)
		ret.Total_fwd_pkts_out = in["total_fwd_pkts_out"].(int)
		ret.Total_rev_bytes_out = in["total_rev_bytes_out"].(int)
		ret.Total_rev_pkts_out = in["total_rev_pkts_out"].(int)
		ret.Curr_req_rate = in["curr_req_rate"].(int)
		ret.Curr_resp = in["curr_resp"].(int)
		ret.Total_resp = in["total_resp"].(int)
		ret.Total_resp_succ = in["total_resp_succ"].(int)
		ret.Curr_resp_rate = in["curr_resp_rate"].(int)
		ret.Dnsrrl_total_allowed = in["dnsrrl_total_allowed"].(int)
		ret.Dnsrrl_total_dropped = in["dnsrrl_total_dropped"].(int)
		ret.Dnsrrl_total_slipped = in["dnsrrl_total_slipped"].(int)
		ret.Dnsrrl_bad_fqdn = in["dnsrrl_bad_fqdn"].(int)
		ret.ThroughputBitsPerSec = in["throughput_bits_per_sec"].(int)
		ret.DynamicMemory = in["dynamic_memory"].(int)
		ret.Ip_only_lb_fwd_bytes = in["ip_only_lb_fwd_bytes"].(int)
		ret.Ip_only_lb_rev_bytes = in["ip_only_lb_rev_bytes"].(int)
		ret.Ip_only_lb_fwd_pkts = in["ip_only_lb_fwd_pkts"].(int)
		ret.Ip_only_lb_rev_pkts = in["ip_only_lb_rev_pkts"].(int)
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
		ret.Compression_bytes_before_br = in["compression_bytes_before_br"].(int)
		ret.Compression_bytes_after_br = in["compression_bytes_after_br"].(int)
		ret.Compression_bytes_before_total = in["compression_bytes_before_total"].(int)
		ret.Compression_bytes_after_total = in["compression_bytes_after_total"].(int)
		ret.Compression_hit_br = in["compression_hit_br"].(int)
		ret.Compression_miss_br = in["compression_miss_br"].(int)
		ret.Compression_hit_total = in["compression_hit_total"].(int)
		ret.Compression_miss_total = in["compression_miss_total"].(int)
		ret.Dnsrrl_total_tc = in["dnsrrl_total_tc"].(int)
		ret.Http1_client_idle_timeout = in["http1_client_idle_timeout"].(int)
		ret.Http2_client_idle_timeout = in["http2_client_idle_timeout"].(int)
	}
	return ret
}

func dataToEndpointSlbVirtualServerPortStats56(d *schema.ResourceData) edpt.SlbVirtualServerPortStats56 {
	var ret edpt.SlbVirtualServerPortStats56

	ret.PortNumber = d.Get("port_number").(int)

	ret.Protocol = d.Get("protocol").(string)

	ret.Stats = getObjectSlbVirtualServerPortStats56Stats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
