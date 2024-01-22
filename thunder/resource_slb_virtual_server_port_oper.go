package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbVirtualServerPortOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_virtual_server_port_oper`: Operational Status for the object port\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbVirtualServerPortOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"real_curr_conn": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"int_curr_conn": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"curr_conn_overflow": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"print_extended_stats": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"loc_list": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"geo_location": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"level_str": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"group_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"loc_max_depth": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"loc_success": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"loc_error": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"loc_override": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"loc_last": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"http_hits_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "Name",
									},
									"hits_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"http_vport_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status_200": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_201": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_202": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_203": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_204": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_205": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_206": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_207": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_100": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_101": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_102": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_300": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_301": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_302": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_303": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_304": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_305": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_306": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_307": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_400": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_401": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_402": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_403": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_404": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_405": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_406": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_407": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_408": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_409": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_410": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_411": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_412": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_413": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_414": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_415": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_416": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_417": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_418": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_422": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_423": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_424": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_425": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_426": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_449": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_450": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_500": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_501": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_502": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_503": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_504": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_504_ax": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_505": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_506": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_507": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_508": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_509": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_510": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_1xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_2xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_3xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_4xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_5xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_6xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_unknown": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ws_handshake_request": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ws_handshake_success": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ws_client_switch": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ws_server_switch": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_10u": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_20u": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_50u": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_100u": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_200u": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_500u": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_1m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_2m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_5m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_10m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_20m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_50m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_100m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_200m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_500m": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_1s": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_2s": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_5s": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_over_5s": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"curr_http2_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_http2_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"peak_http2_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_http2_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http2_control_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http2_header_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http2_data_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http2_reset_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http2_reset_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http2_goaway_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http2_goaway_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"stream_closed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"jsi_requests": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"jsi_responses": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"jsi_pri_requests": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"jsi_api_requests": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"jsi_api_responses": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"jsi_api_no_auth_hdr": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"jsi_api_no_token": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"jsi_skip_no_fi": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"jsi_skip_no_ua": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"jsi_skip_not_browser": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"jsi_hash_add_fails": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"jsi_hash_lookup_fails": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_length_long": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_get": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_head": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_put": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_post": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_trace": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_options": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_connect": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_delete": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_unknown": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_track": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_1k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_2k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_4k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_8k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_16k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_32k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_64k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_256k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_sz_gt_256k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"chunk_sz_512": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"chunk_sz_1k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"chunk_sz_2k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"chunk_sz_4k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"chunk_sz_gt_4k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_1k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_2k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_4k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_8k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_16k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_32k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_64k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_256k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_sz_gt_256k": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_content_len": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsp_chunk": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_req_get": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_req_post": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_non_doh_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_non_doh_req_get": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_non_doh_req_post": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_resp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_tc_resp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_udp_dns_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_udp_dns_resp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_tcp_dns_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_tcp_dns_resp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_req_send_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_resp_send_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_malloc_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_req_udp_retry": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_req_udp_retry_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_req_tcp_retry": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_req_tcp_retry_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_snat_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_path_not_found": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_get_dns_arg_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_get_base64_decode_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_post_content_type_mismatch": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_post_payload_not_found": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_post_payload_extract_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_non_doh_method": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_tcp_send_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_udp_send_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_query_time_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_a": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_aaaa": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_ns": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_cname": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_any": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_srv": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_mx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_soa": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_query_type_others": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_resp_setup_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_resp_header_alloc_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_resp_que_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_resp_udp_frags": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_resp_tcp_frags": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_serv_sel_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_retry_w_tcp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_get_uri_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_post_payload_too_large": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_malformed_query": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_err_format": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_err_server": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_err_name": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_err_type": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_refuse": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_yxdomain": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_yxrrset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_nxrrset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_notauth": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_notzone": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"doh_dns_resp_rcode_other": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"http_host_hits": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"http_url_hits": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"http_vport": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"clear_curr_conn": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"port_number": {
				Type: schema.TypeInt, Required: true, Description: "Port",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': TCP LB service; 'udp': UDP Port; 'others': for no tcp/udp protocol, do IP load balancing; 'diameter': diameter port; 'dns-tcp': DNS service over TCP; 'dns-udp': DNS service over UDP; 'fast-http': Fast HTTP Port; 'fix': FIX Port; 'ftp': File Transfer Protocol Port; 'ftp-proxy': ftp proxy port; 'http': HTTP Port; 'https': HTTPS port; 'imap': imap proxy port; 'mlb': Message based load balancing; 'mms': Microsoft Multimedia Service Port; 'mysql': mssql port; 'mssql': mssql; 'pop3': pop3 proxy port; 'radius': RADIUS Port; 'rtsp': Real Time Streaming Protocol Port; 'sip': Session initiation protocol over UDP; 'sip-tcp': Session initiation protocol over TCP; 'sips': Session initiation protocol over TLS; 'smpp-tcp': SMPP service over TCP; 'spdy': spdy port; 'spdys': spdys port; 'smtp': SMTP Port; 'mqtt': MQTT Port; 'mqtts': MQTTS Port; 'ssl-proxy': Generic SSL proxy; 'ssli': SSL insight; 'ssh': SSH Port; 'tcp-proxy': Generic TCP proxy; 'tftp': TFTP Port; 'fast-fix': Fast FIX port; 'http-over-quic': HTTP3-over-quic port;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceSlbVirtualServerPortOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbVirtualServerPortOperOper := setObjectSlbVirtualServerPortOperOper(res)
		d.Set("oper", SlbVirtualServerPortOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbVirtualServerPortOperOper(ret edpt.DataSlbVirtualServerPortOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":                ret.DtSlbVirtualServerPortOper.Oper.State,
			"real_curr_conn":       ret.DtSlbVirtualServerPortOper.Oper.Real_curr_conn,
			"int_curr_conn":        ret.DtSlbVirtualServerPortOper.Oper.Int_curr_conn,
			"curr_conn_overflow":   ret.DtSlbVirtualServerPortOper.Oper.Curr_conn_overflow,
			"print_extended_stats": ret.DtSlbVirtualServerPortOper.Oper.Print_extended_stats,
			"loc_list":             ret.DtSlbVirtualServerPortOper.Oper.Loc_list,
			"geo_location":         ret.DtSlbVirtualServerPortOper.Oper.Geo_location,
			"level_str":            ret.DtSlbVirtualServerPortOper.Oper.Level_str,
			"group_id":             ret.DtSlbVirtualServerPortOper.Oper.Group_id,
			"loc_max_depth":        ret.DtSlbVirtualServerPortOper.Oper.Loc_max_depth,
			"loc_success":          ret.DtSlbVirtualServerPortOper.Oper.Loc_success,
			"loc_error":            ret.DtSlbVirtualServerPortOper.Oper.Loc_error,
			"loc_override":         ret.DtSlbVirtualServerPortOper.Oper.Loc_override,
			"loc_last":             ret.DtSlbVirtualServerPortOper.Oper.Loc_last,
			"http_hits_list":       setSliceSlbVirtualServerPortOperOperHttpHitsList(ret.DtSlbVirtualServerPortOper.Oper.HttpHitsList),
			"http_vport_cpu_list":  setSliceSlbVirtualServerPortOperOperHttpVportCpuList(ret.DtSlbVirtualServerPortOper.Oper.HttpVportCpuList),
			"cpu_count":            ret.DtSlbVirtualServerPortOper.Oper.CpuCount,
			"http_host_hits":       ret.DtSlbVirtualServerPortOper.Oper.Http_host_hits,
			"http_url_hits":        ret.DtSlbVirtualServerPortOper.Oper.Http_url_hits,
			"http_vport":           ret.DtSlbVirtualServerPortOper.Oper.Http_vport,
			"clear_curr_conn":      ret.DtSlbVirtualServerPortOper.Oper.Clear_curr_conn,
		},
	}
}

func setSliceSlbVirtualServerPortOperOperHttpHitsList(d []edpt.SlbVirtualServerPortOperOperHttpHitsList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["hits_count"] = item.HitsCount
		result = append(result, in)
	}
	return result
}

func setSliceSlbVirtualServerPortOperOperHttpVportCpuList(d []edpt.SlbVirtualServerPortOperOperHttpVportCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["status_200"] = item.Status_200
		in["status_201"] = item.Status_201
		in["status_202"] = item.Status_202
		in["status_203"] = item.Status_203
		in["status_204"] = item.Status_204
		in["status_205"] = item.Status_205
		in["status_206"] = item.Status_206
		in["status_207"] = item.Status_207
		in["status_100"] = item.Status_100
		in["status_101"] = item.Status_101
		in["status_102"] = item.Status_102
		in["status_300"] = item.Status_300
		in["status_301"] = item.Status_301
		in["status_302"] = item.Status_302
		in["status_303"] = item.Status_303
		in["status_304"] = item.Status_304
		in["status_305"] = item.Status_305
		in["status_306"] = item.Status_306
		in["status_307"] = item.Status_307
		in["status_400"] = item.Status_400
		in["status_401"] = item.Status_401
		in["status_402"] = item.Status_402
		in["status_403"] = item.Status_403
		in["status_404"] = item.Status_404
		in["status_405"] = item.Status_405
		in["status_406"] = item.Status_406
		in["status_407"] = item.Status_407
		in["status_408"] = item.Status_408
		in["status_409"] = item.Status_409
		in["status_410"] = item.Status_410
		in["status_411"] = item.Status_411
		in["status_412"] = item.Status_412
		in["status_413"] = item.Status_413
		in["status_414"] = item.Status_414
		in["status_415"] = item.Status_415
		in["status_416"] = item.Status_416
		in["status_417"] = item.Status_417
		in["status_418"] = item.Status_418
		in["status_422"] = item.Status_422
		in["status_423"] = item.Status_423
		in["status_424"] = item.Status_424
		in["status_425"] = item.Status_425
		in["status_426"] = item.Status_426
		in["status_449"] = item.Status_449
		in["status_450"] = item.Status_450
		in["status_500"] = item.Status_500
		in["status_501"] = item.Status_501
		in["status_502"] = item.Status_502
		in["status_503"] = item.Status_503
		in["status_504"] = item.Status_504
		in["status_504_ax"] = item.Status_504_ax
		in["status_505"] = item.Status_505
		in["status_506"] = item.Status_506
		in["status_507"] = item.Status_507
		in["status_508"] = item.Status_508
		in["status_509"] = item.Status_509
		in["status_510"] = item.Status_510
		in["status_1xx"] = item.Status_1xx
		in["status_2xx"] = item.Status_2xx
		in["status_3xx"] = item.Status_3xx
		in["status_4xx"] = item.Status_4xx
		in["status_5xx"] = item.Status_5xx
		in["status_6xx"] = item.Status_6xx
		in["status_unknown"] = item.Status_unknown
		in["ws_handshake_request"] = item.Ws_handshake_request
		in["ws_handshake_success"] = item.Ws_handshake_success
		in["ws_client_switch"] = item.Ws_client_switch
		in["ws_server_switch"] = item.Ws_server_switch
		in["req_10u"] = item.Req_10u
		in["req_20u"] = item.Req_20u
		in["req_50u"] = item.Req_50u
		in["req_100u"] = item.Req_100u
		in["req_200u"] = item.Req_200u
		in["req_500u"] = item.Req_500u
		in["req_1m"] = item.Req_1m
		in["req_2m"] = item.Req_2m
		in["req_5m"] = item.Req_5m
		in["req_10m"] = item.Req_10m
		in["req_20m"] = item.Req_20m
		in["req_50m"] = item.Req_50m
		in["req_100m"] = item.Req_100m
		in["req_200m"] = item.Req_200m
		in["req_500m"] = item.Req_500m
		in["req_1s"] = item.Req_1s
		in["req_2s"] = item.Req_2s
		in["req_5s"] = item.Req_5s
		in["req_over_5s"] = item.Req_over_5s
		in["curr_http2_conn"] = item.Curr_http2_conn
		in["total_http2_conn"] = item.Total_http2_conn
		in["peak_http2_conn"] = item.Peak_http2_conn
		in["total_http2_bytes"] = item.Total_http2_bytes
		in["http2_control_bytes"] = item.Http2_control_bytes
		in["http2_header_bytes"] = item.Http2_header_bytes
		in["http2_data_bytes"] = item.Http2_data_bytes
		in["http2_reset_received"] = item.Http2_reset_received
		in["http2_reset_sent"] = item.Http2_reset_sent
		in["http2_goaway_received"] = item.Http2_goaway_received
		in["http2_goaway_sent"] = item.Http2_goaway_sent
		in["stream_closed"] = item.Stream_closed
		in["jsi_requests"] = item.Jsi_requests
		in["jsi_responses"] = item.Jsi_responses
		in["jsi_pri_requests"] = item.Jsi_pri_requests
		in["jsi_api_requests"] = item.Jsi_api_requests
		in["jsi_api_responses"] = item.Jsi_api_responses
		in["jsi_api_no_auth_hdr"] = item.Jsi_api_no_auth_hdr
		in["jsi_api_no_token"] = item.Jsi_api_no_token
		in["jsi_skip_no_fi"] = item.Jsi_skip_no_fi
		in["jsi_skip_no_ua"] = item.Jsi_skip_no_ua
		in["jsi_skip_not_browser"] = item.Jsi_skip_not_browser
		in["jsi_hash_add_fails"] = item.Jsi_hash_add_fails
		in["jsi_hash_lookup_fails"] = item.Jsi_hash_lookup_fails
		in["header_length_long"] = item.Header_length_long
		in["req_get"] = item.Req_get
		in["req_head"] = item.Req_head
		in["req_put"] = item.Req_put
		in["req_post"] = item.Req_post
		in["req_trace"] = item.Req_trace
		in["req_options"] = item.Req_options
		in["req_connect"] = item.Req_connect
		in["req_delete"] = item.Req_delete
		in["req_unknown"] = item.Req_unknown
		in["req_track"] = item.Req_track
		in["rsp_sz_1k"] = item.Rsp_sz_1k
		in["rsp_sz_2k"] = item.Rsp_sz_2k
		in["rsp_sz_4k"] = item.Rsp_sz_4k
		in["rsp_sz_8k"] = item.Rsp_sz_8k
		in["rsp_sz_16k"] = item.Rsp_sz_16k
		in["rsp_sz_32k"] = item.Rsp_sz_32k
		in["rsp_sz_64k"] = item.Rsp_sz_64k
		in["rsp_sz_256k"] = item.Rsp_sz_256k
		in["rsp_sz_gt_256k"] = item.Rsp_sz_gt_256k
		in["chunk_sz_512"] = item.Chunk_sz_512
		in["chunk_sz_1k"] = item.Chunk_sz_1k
		in["chunk_sz_2k"] = item.Chunk_sz_2k
		in["chunk_sz_4k"] = item.Chunk_sz_4k
		in["chunk_sz_gt_4k"] = item.Chunk_sz_gt_4k
		in["req_sz_1k"] = item.Req_sz_1k
		in["req_sz_2k"] = item.Req_sz_2k
		in["req_sz_4k"] = item.Req_sz_4k
		in["req_sz_8k"] = item.Req_sz_8k
		in["req_sz_16k"] = item.Req_sz_16k
		in["req_sz_32k"] = item.Req_sz_32k
		in["req_sz_64k"] = item.Req_sz_64k
		in["req_sz_256k"] = item.Req_sz_256k
		in["req_sz_gt_256k"] = item.Req_sz_gt_256k
		in["req_content_len"] = item.Req_content_len
		in["rsp_chunk"] = item.Rsp_chunk
		in["doh_req"] = item.Doh_req
		in["doh_req_get"] = item.Doh_req_get
		in["doh_req_post"] = item.Doh_req_post
		in["doh_non_doh_req"] = item.Doh_non_doh_req
		in["doh_non_doh_req_get"] = item.Doh_non_doh_req_get
		in["doh_non_doh_req_post"] = item.Doh_non_doh_req_post
		in["doh_resp"] = item.Doh_resp
		in["doh_tc_resp"] = item.Doh_tc_resp
		in["doh_udp_dns_req"] = item.Doh_udp_dns_req
		in["doh_udp_dns_resp"] = item.Doh_udp_dns_resp
		in["doh_tcp_dns_req"] = item.Doh_tcp_dns_req
		in["doh_tcp_dns_resp"] = item.Doh_tcp_dns_resp
		in["doh_req_send_failed"] = item.Doh_req_send_failed
		in["doh_resp_send_failed"] = item.Doh_resp_send_failed
		in["doh_malloc_fail"] = item.Doh_malloc_fail
		in["doh_req_udp_retry"] = item.Doh_req_udp_retry
		in["doh_req_udp_retry_fail"] = item.Doh_req_udp_retry_fail
		in["doh_req_tcp_retry"] = item.Doh_req_tcp_retry
		in["doh_req_tcp_retry_fail"] = item.Doh_req_tcp_retry_fail
		in["doh_snat_failed"] = item.Doh_snat_failed
		in["doh_path_not_found"] = item.Doh_path_not_found
		in["doh_get_dns_arg_failed"] = item.Doh_get_dns_arg_failed
		in["doh_get_base64_decode_failed"] = item.Doh_get_base64_decode_failed
		in["doh_post_content_type_mismatch"] = item.Doh_post_content_type_mismatch
		in["doh_post_payload_not_found"] = item.Doh_post_payload_not_found
		in["doh_post_payload_extract_failed"] = item.Doh_post_payload_extract_failed
		in["doh_non_doh_method"] = item.Doh_non_doh_method
		in["doh_tcp_send_failed"] = item.Doh_tcp_send_failed
		in["doh_udp_send_failed"] = item.Doh_udp_send_failed
		in["doh_query_time_out"] = item.Doh_query_time_out
		in["doh_dns_query_type_a"] = item.Doh_dns_query_type_a
		in["doh_dns_query_type_aaaa"] = item.Doh_dns_query_type_aaaa
		in["doh_dns_query_type_ns"] = item.Doh_dns_query_type_ns
		in["doh_dns_query_type_cname"] = item.Doh_dns_query_type_cname
		in["doh_dns_query_type_any"] = item.Doh_dns_query_type_any
		in["doh_dns_query_type_srv"] = item.Doh_dns_query_type_srv
		in["doh_dns_query_type_mx"] = item.Doh_dns_query_type_mx
		in["doh_dns_query_type_soa"] = item.Doh_dns_query_type_soa
		in["doh_dns_query_type_others"] = item.Doh_dns_query_type_others
		in["doh_resp_setup_failed"] = item.Doh_resp_setup_failed
		in["doh_resp_header_alloc_failed"] = item.Doh_resp_header_alloc_failed
		in["doh_resp_que_failed"] = item.Doh_resp_que_failed
		in["doh_resp_udp_frags"] = item.Doh_resp_udp_frags
		in["doh_resp_tcp_frags"] = item.Doh_resp_tcp_frags
		in["doh_serv_sel_failed"] = item.Doh_serv_sel_failed
		in["doh_retry_w_tcp"] = item.Doh_retry_w_tcp
		in["doh_get_uri_too_long"] = item.Doh_get_uri_too_long
		in["doh_post_payload_too_large"] = item.Doh_post_payload_too_large
		in["doh_dns_malformed_query"] = item.Doh_dns_malformed_query
		in["doh_dns_resp_rcode_err_format"] = item.Doh_dns_resp_rcode_err_format
		in["doh_dns_resp_rcode_err_server"] = item.Doh_dns_resp_rcode_err_server
		in["doh_dns_resp_rcode_err_name"] = item.Doh_dns_resp_rcode_err_name
		in["doh_dns_resp_rcode_err_type"] = item.Doh_dns_resp_rcode_err_type
		in["doh_dns_resp_rcode_refuse"] = item.Doh_dns_resp_rcode_refuse
		in["doh_dns_resp_rcode_yxdomain"] = item.Doh_dns_resp_rcode_yxdomain
		in["doh_dns_resp_rcode_yxrrset"] = item.Doh_dns_resp_rcode_yxrrset
		in["doh_dns_resp_rcode_nxrrset"] = item.Doh_dns_resp_rcode_nxrrset
		in["doh_dns_resp_rcode_notauth"] = item.Doh_dns_resp_rcode_notauth
		in["doh_dns_resp_rcode_notzone"] = item.Doh_dns_resp_rcode_notzone
		in["doh_dns_resp_rcode_other"] = item.Doh_dns_resp_rcode_other
		result = append(result, in)
	}
	return result
}

func getObjectSlbVirtualServerPortOperOper(d []interface{}) edpt.SlbVirtualServerPortOperOper {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Real_curr_conn = in["real_curr_conn"].(int)
		ret.Int_curr_conn = in["int_curr_conn"].(int)
		ret.Curr_conn_overflow = in["curr_conn_overflow"].(int)
		ret.Print_extended_stats = in["print_extended_stats"].(int)
		ret.Loc_list = in["loc_list"].(string)
		ret.Geo_location = in["geo_location"].(string)
		ret.Level_str = in["level_str"].(string)
		ret.Group_id = in["group_id"].(int)
		ret.Loc_max_depth = in["loc_max_depth"].(int)
		ret.Loc_success = in["loc_success"].(int)
		ret.Loc_error = in["loc_error"].(int)
		ret.Loc_override = in["loc_override"].(int)
		ret.Loc_last = in["loc_last"].(string)
		ret.HttpHitsList = getSliceSlbVirtualServerPortOperOperHttpHitsList(in["http_hits_list"].([]interface{}))
		ret.HttpVportCpuList = getSliceSlbVirtualServerPortOperOperHttpVportCpuList(in["http_vport_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
		ret.Http_host_hits = in["http_host_hits"].(int)
		ret.Http_url_hits = in["http_url_hits"].(int)
		ret.Http_vport = in["http_vport"].(int)
		ret.Clear_curr_conn = in["clear_curr_conn"].(int)
	}
	return ret
}

func getSliceSlbVirtualServerPortOperOperHttpHitsList(d []interface{}) []edpt.SlbVirtualServerPortOperOperHttpHitsList {

	count1 := len(d)
	ret := make([]edpt.SlbVirtualServerPortOperOperHttpHitsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbVirtualServerPortOperOperHttpHitsList
		oi.Name = in["name"].(string)
		oi.HitsCount = in["hits_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbVirtualServerPortOperOperHttpVportCpuList(d []interface{}) []edpt.SlbVirtualServerPortOperOperHttpVportCpuList {

	count1 := len(d)
	ret := make([]edpt.SlbVirtualServerPortOperOperHttpVportCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbVirtualServerPortOperOperHttpVportCpuList
		oi.Status_200 = in["status_200"].(int)
		oi.Status_201 = in["status_201"].(int)
		oi.Status_202 = in["status_202"].(int)
		oi.Status_203 = in["status_203"].(int)
		oi.Status_204 = in["status_204"].(int)
		oi.Status_205 = in["status_205"].(int)
		oi.Status_206 = in["status_206"].(int)
		oi.Status_207 = in["status_207"].(int)
		oi.Status_100 = in["status_100"].(int)
		oi.Status_101 = in["status_101"].(int)
		oi.Status_102 = in["status_102"].(int)
		oi.Status_300 = in["status_300"].(int)
		oi.Status_301 = in["status_301"].(int)
		oi.Status_302 = in["status_302"].(int)
		oi.Status_303 = in["status_303"].(int)
		oi.Status_304 = in["status_304"].(int)
		oi.Status_305 = in["status_305"].(int)
		oi.Status_306 = in["status_306"].(int)
		oi.Status_307 = in["status_307"].(int)
		oi.Status_400 = in["status_400"].(int)
		oi.Status_401 = in["status_401"].(int)
		oi.Status_402 = in["status_402"].(int)
		oi.Status_403 = in["status_403"].(int)
		oi.Status_404 = in["status_404"].(int)
		oi.Status_405 = in["status_405"].(int)
		oi.Status_406 = in["status_406"].(int)
		oi.Status_407 = in["status_407"].(int)
		oi.Status_408 = in["status_408"].(int)
		oi.Status_409 = in["status_409"].(int)
		oi.Status_410 = in["status_410"].(int)
		oi.Status_411 = in["status_411"].(int)
		oi.Status_412 = in["status_412"].(int)
		oi.Status_413 = in["status_413"].(int)
		oi.Status_414 = in["status_414"].(int)
		oi.Status_415 = in["status_415"].(int)
		oi.Status_416 = in["status_416"].(int)
		oi.Status_417 = in["status_417"].(int)
		oi.Status_418 = in["status_418"].(int)
		oi.Status_422 = in["status_422"].(int)
		oi.Status_423 = in["status_423"].(int)
		oi.Status_424 = in["status_424"].(int)
		oi.Status_425 = in["status_425"].(int)
		oi.Status_426 = in["status_426"].(int)
		oi.Status_449 = in["status_449"].(int)
		oi.Status_450 = in["status_450"].(int)
		oi.Status_500 = in["status_500"].(int)
		oi.Status_501 = in["status_501"].(int)
		oi.Status_502 = in["status_502"].(int)
		oi.Status_503 = in["status_503"].(int)
		oi.Status_504 = in["status_504"].(int)
		oi.Status_504_ax = in["status_504_ax"].(int)
		oi.Status_505 = in["status_505"].(int)
		oi.Status_506 = in["status_506"].(int)
		oi.Status_507 = in["status_507"].(int)
		oi.Status_508 = in["status_508"].(int)
		oi.Status_509 = in["status_509"].(int)
		oi.Status_510 = in["status_510"].(int)
		oi.Status_1xx = in["status_1xx"].(int)
		oi.Status_2xx = in["status_2xx"].(int)
		oi.Status_3xx = in["status_3xx"].(int)
		oi.Status_4xx = in["status_4xx"].(int)
		oi.Status_5xx = in["status_5xx"].(int)
		oi.Status_6xx = in["status_6xx"].(int)
		oi.Status_unknown = in["status_unknown"].(int)
		oi.Ws_handshake_request = in["ws_handshake_request"].(int)
		oi.Ws_handshake_success = in["ws_handshake_success"].(int)
		oi.Ws_client_switch = in["ws_client_switch"].(int)
		oi.Ws_server_switch = in["ws_server_switch"].(int)
		oi.Req_10u = in["req_10u"].(int)
		oi.Req_20u = in["req_20u"].(int)
		oi.Req_50u = in["req_50u"].(int)
		oi.Req_100u = in["req_100u"].(int)
		oi.Req_200u = in["req_200u"].(int)
		oi.Req_500u = in["req_500u"].(int)
		oi.Req_1m = in["req_1m"].(int)
		oi.Req_2m = in["req_2m"].(int)
		oi.Req_5m = in["req_5m"].(int)
		oi.Req_10m = in["req_10m"].(int)
		oi.Req_20m = in["req_20m"].(int)
		oi.Req_50m = in["req_50m"].(int)
		oi.Req_100m = in["req_100m"].(int)
		oi.Req_200m = in["req_200m"].(int)
		oi.Req_500m = in["req_500m"].(int)
		oi.Req_1s = in["req_1s"].(int)
		oi.Req_2s = in["req_2s"].(int)
		oi.Req_5s = in["req_5s"].(int)
		oi.Req_over_5s = in["req_over_5s"].(int)
		oi.Curr_http2_conn = in["curr_http2_conn"].(int)
		oi.Total_http2_conn = in["total_http2_conn"].(int)
		oi.Peak_http2_conn = in["peak_http2_conn"].(int)
		oi.Total_http2_bytes = in["total_http2_bytes"].(int)
		oi.Http2_control_bytes = in["http2_control_bytes"].(int)
		oi.Http2_header_bytes = in["http2_header_bytes"].(int)
		oi.Http2_data_bytes = in["http2_data_bytes"].(int)
		oi.Http2_reset_received = in["http2_reset_received"].(int)
		oi.Http2_reset_sent = in["http2_reset_sent"].(int)
		oi.Http2_goaway_received = in["http2_goaway_received"].(int)
		oi.Http2_goaway_sent = in["http2_goaway_sent"].(int)
		oi.Stream_closed = in["stream_closed"].(int)
		oi.Jsi_requests = in["jsi_requests"].(int)
		oi.Jsi_responses = in["jsi_responses"].(int)
		oi.Jsi_pri_requests = in["jsi_pri_requests"].(int)
		oi.Jsi_api_requests = in["jsi_api_requests"].(int)
		oi.Jsi_api_responses = in["jsi_api_responses"].(int)
		oi.Jsi_api_no_auth_hdr = in["jsi_api_no_auth_hdr"].(int)
		oi.Jsi_api_no_token = in["jsi_api_no_token"].(int)
		oi.Jsi_skip_no_fi = in["jsi_skip_no_fi"].(int)
		oi.Jsi_skip_no_ua = in["jsi_skip_no_ua"].(int)
		oi.Jsi_skip_not_browser = in["jsi_skip_not_browser"].(int)
		oi.Jsi_hash_add_fails = in["jsi_hash_add_fails"].(int)
		oi.Jsi_hash_lookup_fails = in["jsi_hash_lookup_fails"].(int)
		oi.Header_length_long = in["header_length_long"].(int)
		oi.Req_get = in["req_get"].(int)
		oi.Req_head = in["req_head"].(int)
		oi.Req_put = in["req_put"].(int)
		oi.Req_post = in["req_post"].(int)
		oi.Req_trace = in["req_trace"].(int)
		oi.Req_options = in["req_options"].(int)
		oi.Req_connect = in["req_connect"].(int)
		oi.Req_delete = in["req_delete"].(int)
		oi.Req_unknown = in["req_unknown"].(int)
		oi.Req_track = in["req_track"].(int)
		oi.Rsp_sz_1k = in["rsp_sz_1k"].(int)
		oi.Rsp_sz_2k = in["rsp_sz_2k"].(int)
		oi.Rsp_sz_4k = in["rsp_sz_4k"].(int)
		oi.Rsp_sz_8k = in["rsp_sz_8k"].(int)
		oi.Rsp_sz_16k = in["rsp_sz_16k"].(int)
		oi.Rsp_sz_32k = in["rsp_sz_32k"].(int)
		oi.Rsp_sz_64k = in["rsp_sz_64k"].(int)
		oi.Rsp_sz_256k = in["rsp_sz_256k"].(int)
		oi.Rsp_sz_gt_256k = in["rsp_sz_gt_256k"].(int)
		oi.Chunk_sz_512 = in["chunk_sz_512"].(int)
		oi.Chunk_sz_1k = in["chunk_sz_1k"].(int)
		oi.Chunk_sz_2k = in["chunk_sz_2k"].(int)
		oi.Chunk_sz_4k = in["chunk_sz_4k"].(int)
		oi.Chunk_sz_gt_4k = in["chunk_sz_gt_4k"].(int)
		oi.Req_sz_1k = in["req_sz_1k"].(int)
		oi.Req_sz_2k = in["req_sz_2k"].(int)
		oi.Req_sz_4k = in["req_sz_4k"].(int)
		oi.Req_sz_8k = in["req_sz_8k"].(int)
		oi.Req_sz_16k = in["req_sz_16k"].(int)
		oi.Req_sz_32k = in["req_sz_32k"].(int)
		oi.Req_sz_64k = in["req_sz_64k"].(int)
		oi.Req_sz_256k = in["req_sz_256k"].(int)
		oi.Req_sz_gt_256k = in["req_sz_gt_256k"].(int)
		oi.Req_content_len = in["req_content_len"].(int)
		oi.Rsp_chunk = in["rsp_chunk"].(int)
		oi.Doh_req = in["doh_req"].(int)
		oi.Doh_req_get = in["doh_req_get"].(int)
		oi.Doh_req_post = in["doh_req_post"].(int)
		oi.Doh_non_doh_req = in["doh_non_doh_req"].(int)
		oi.Doh_non_doh_req_get = in["doh_non_doh_req_get"].(int)
		oi.Doh_non_doh_req_post = in["doh_non_doh_req_post"].(int)
		oi.Doh_resp = in["doh_resp"].(int)
		oi.Doh_tc_resp = in["doh_tc_resp"].(int)
		oi.Doh_udp_dns_req = in["doh_udp_dns_req"].(int)
		oi.Doh_udp_dns_resp = in["doh_udp_dns_resp"].(int)
		oi.Doh_tcp_dns_req = in["doh_tcp_dns_req"].(int)
		oi.Doh_tcp_dns_resp = in["doh_tcp_dns_resp"].(int)
		oi.Doh_req_send_failed = in["doh_req_send_failed"].(int)
		oi.Doh_resp_send_failed = in["doh_resp_send_failed"].(int)
		oi.Doh_malloc_fail = in["doh_malloc_fail"].(int)
		oi.Doh_req_udp_retry = in["doh_req_udp_retry"].(int)
		oi.Doh_req_udp_retry_fail = in["doh_req_udp_retry_fail"].(int)
		oi.Doh_req_tcp_retry = in["doh_req_tcp_retry"].(int)
		oi.Doh_req_tcp_retry_fail = in["doh_req_tcp_retry_fail"].(int)
		oi.Doh_snat_failed = in["doh_snat_failed"].(int)
		oi.Doh_path_not_found = in["doh_path_not_found"].(int)
		oi.Doh_get_dns_arg_failed = in["doh_get_dns_arg_failed"].(int)
		oi.Doh_get_base64_decode_failed = in["doh_get_base64_decode_failed"].(int)
		oi.Doh_post_content_type_mismatch = in["doh_post_content_type_mismatch"].(int)
		oi.Doh_post_payload_not_found = in["doh_post_payload_not_found"].(int)
		oi.Doh_post_payload_extract_failed = in["doh_post_payload_extract_failed"].(int)
		oi.Doh_non_doh_method = in["doh_non_doh_method"].(int)
		oi.Doh_tcp_send_failed = in["doh_tcp_send_failed"].(int)
		oi.Doh_udp_send_failed = in["doh_udp_send_failed"].(int)
		oi.Doh_query_time_out = in["doh_query_time_out"].(int)
		oi.Doh_dns_query_type_a = in["doh_dns_query_type_a"].(int)
		oi.Doh_dns_query_type_aaaa = in["doh_dns_query_type_aaaa"].(int)
		oi.Doh_dns_query_type_ns = in["doh_dns_query_type_ns"].(int)
		oi.Doh_dns_query_type_cname = in["doh_dns_query_type_cname"].(int)
		oi.Doh_dns_query_type_any = in["doh_dns_query_type_any"].(int)
		oi.Doh_dns_query_type_srv = in["doh_dns_query_type_srv"].(int)
		oi.Doh_dns_query_type_mx = in["doh_dns_query_type_mx"].(int)
		oi.Doh_dns_query_type_soa = in["doh_dns_query_type_soa"].(int)
		oi.Doh_dns_query_type_others = in["doh_dns_query_type_others"].(int)
		oi.Doh_resp_setup_failed = in["doh_resp_setup_failed"].(int)
		oi.Doh_resp_header_alloc_failed = in["doh_resp_header_alloc_failed"].(int)
		oi.Doh_resp_que_failed = in["doh_resp_que_failed"].(int)
		oi.Doh_resp_udp_frags = in["doh_resp_udp_frags"].(int)
		oi.Doh_resp_tcp_frags = in["doh_resp_tcp_frags"].(int)
		oi.Doh_serv_sel_failed = in["doh_serv_sel_failed"].(int)
		oi.Doh_retry_w_tcp = in["doh_retry_w_tcp"].(int)
		oi.Doh_get_uri_too_long = in["doh_get_uri_too_long"].(int)
		oi.Doh_post_payload_too_large = in["doh_post_payload_too_large"].(int)
		oi.Doh_dns_malformed_query = in["doh_dns_malformed_query"].(int)
		oi.Doh_dns_resp_rcode_err_format = in["doh_dns_resp_rcode_err_format"].(int)
		oi.Doh_dns_resp_rcode_err_server = in["doh_dns_resp_rcode_err_server"].(int)
		oi.Doh_dns_resp_rcode_err_name = in["doh_dns_resp_rcode_err_name"].(int)
		oi.Doh_dns_resp_rcode_err_type = in["doh_dns_resp_rcode_err_type"].(int)
		oi.Doh_dns_resp_rcode_refuse = in["doh_dns_resp_rcode_refuse"].(int)
		oi.Doh_dns_resp_rcode_yxdomain = in["doh_dns_resp_rcode_yxdomain"].(int)
		oi.Doh_dns_resp_rcode_yxrrset = in["doh_dns_resp_rcode_yxrrset"].(int)
		oi.Doh_dns_resp_rcode_nxrrset = in["doh_dns_resp_rcode_nxrrset"].(int)
		oi.Doh_dns_resp_rcode_notauth = in["doh_dns_resp_rcode_notauth"].(int)
		oi.Doh_dns_resp_rcode_notzone = in["doh_dns_resp_rcode_notzone"].(int)
		oi.Doh_dns_resp_rcode_other = in["doh_dns_resp_rcode_other"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbVirtualServerPortOper(d *schema.ResourceData) edpt.SlbVirtualServerPortOper {
	var ret edpt.SlbVirtualServerPortOper

	ret.Oper = getObjectSlbVirtualServerPortOperOper(d.Get("oper").([]interface{}))

	ret.PortNumber = d.Get("port_number").(int)

	ret.Protocol = d.Get("protocol").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
