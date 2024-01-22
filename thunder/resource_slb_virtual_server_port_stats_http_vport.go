package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbVirtualServerPortStats50() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_virtual_server_port_stats_http_vport`: Statistics for the object port\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbVirtualServerPortStats50Create,
		UpdateContext: resourceSlbVirtualServerPortStats50Update,
		ReadContext:   resourceSlbVirtualServerPortStats50Read,
		DeleteContext: resourceSlbVirtualServerPortStats50Delete,

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
						"http_vport": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status_200": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 200",
									},
									"status_201": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 201",
									},
									"status_202": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 202",
									},
									"status_203": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 203",
									},
									"status_204": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 204",
									},
									"status_205": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 205",
									},
									"status_206": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 206",
									},
									"status_207": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 207",
									},
									"status_100": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 100",
									},
									"status_101": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 101",
									},
									"status_102": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 102",
									},
									"status_103": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 103",
									},
									"status_300": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 300",
									},
									"status_301": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 301",
									},
									"status_302": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 302",
									},
									"status_303": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 303",
									},
									"status_304": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 304",
									},
									"status_305": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 305",
									},
									"status_306": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 306",
									},
									"status_307": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 307",
									},
									"status_400": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 400",
									},
									"status_401": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 401",
									},
									"status_402": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 402",
									},
									"status_403": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 403",
									},
									"status_404": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 404",
									},
									"status_405": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 405",
									},
									"status_406": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 406",
									},
									"status_407": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 407",
									},
									"status_408": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 408",
									},
									"status_409": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 409",
									},
									"status_410": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 410",
									},
									"status_411": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 411",
									},
									"status_412": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 412",
									},
									"status_413": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 413",
									},
									"status_414": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 414",
									},
									"status_415": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 415",
									},
									"status_416": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 416",
									},
									"status_417": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 417",
									},
									"status_418": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 418",
									},
									"status_422": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 422",
									},
									"status_423": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 423",
									},
									"status_424": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 424",
									},
									"status_425": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 425",
									},
									"status_426": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 426",
									},
									"status_449": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 449",
									},
									"status_450": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 450",
									},
									"status_500": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 500",
									},
									"status_501": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 501",
									},
									"status_502": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 502",
									},
									"status_503": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 503",
									},
									"status_504": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 504",
									},
									"status_504_ax": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 504 AX-gen",
									},
									"status_505": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 505",
									},
									"status_506": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 506",
									},
									"status_507": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 507",
									},
									"status_508": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 508",
									},
									"status_509": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 509",
									},
									"status_510": {
										Type: schema.TypeInt, Optional: true, Description: "Status code 510",
									},
									"status_1xx": {
										Type: schema.TypeInt, Optional: true, Description: "status code 1XX",
									},
									"status_2xx": {
										Type: schema.TypeInt, Optional: true, Description: "status code 2XX",
									},
									"status_3xx": {
										Type: schema.TypeInt, Optional: true, Description: "status code 3XX",
									},
									"status_4xx": {
										Type: schema.TypeInt, Optional: true, Description: "status code 4XX",
									},
									"status_5xx": {
										Type: schema.TypeInt, Optional: true, Description: "status code 5XX",
									},
									"status_6xx": {
										Type: schema.TypeInt, Optional: true, Description: "status code 6XX",
									},
									"status_unknown": {
										Type: schema.TypeInt, Optional: true, Description: "Status code unknown",
									},
									"ws_handshake_request": {
										Type: schema.TypeInt, Optional: true, Description: "WS Handshake Req",
									},
									"ws_handshake_success": {
										Type: schema.TypeInt, Optional: true, Description: "WS Handshake Res",
									},
									"ws_client_switch": {
										Type: schema.TypeInt, Optional: true, Description: "WS Client Pkts",
									},
									"ws_server_switch": {
										Type: schema.TypeInt, Optional: true, Description: "WS Server Pkts",
									},
									"req_10u": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 10u",
									},
									"req_20u": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 20u",
									},
									"req_50u": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 50u",
									},
									"req_100u": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 100u",
									},
									"req_200u": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 200u",
									},
									"req_500u": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 500u",
									},
									"req_1m": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 1m",
									},
									"req_2m": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 2m",
									},
									"req_5m": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 5m",
									},
									"req_10m": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 10m",
									},
									"req_20m": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 20m",
									},
									"req_50m": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 5m",
									},
									"req_100m": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 100m",
									},
									"req_200m": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 200m",
									},
									"req_500m": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 500m",
									},
									"req_1s": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 1s",
									},
									"req_2s": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 2s",
									},
									"req_5s": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 5s",
									},
									"req_over_5s": {
										Type: schema.TypeInt, Optional: true, Description: "Rsp time greater than equal to 5s",
									},
									"total_requests": {
										Type: schema.TypeInt, Optional: true, Description: "Total number of Requests",
									},
									"curr_http2_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Current number of HTTP/2 connections",
									},
									"total_http2_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Total number of HTTP/2 connections",
									},
									"peak_http2_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Largest number of concurrent HTTP/2 connections",
									},
									"total_http2_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Total bytes in HTTP/2 frames",
									},
									"http2_control_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Bytes in HTTP/2 control frames",
									},
									"http2_header_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Bytes in HTTP/2 header frames",
									},
									"http2_data_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Bytes in HTTP/2 data frames",
									},
									"http2_reset_received": {
										Type: schema.TypeInt, Optional: true, Description: "Reset frames received",
									},
									"http2_reset_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Reset frames sent",
									},
									"http2_goaway_received": {
										Type: schema.TypeInt, Optional: true, Description: "GoAway frames received",
									},
									"http2_goaway_sent": {
										Type: schema.TypeInt, Optional: true, Description: "GoAway frames sent",
									},
									"stream_closed": {
										Type: schema.TypeInt, Optional: true, Description: "HTTP/2 streams closed",
									},
									"jsi_requests": {
										Type: schema.TypeInt, Optional: true, Description: "JSI requests",
									},
									"jsi_responses": {
										Type: schema.TypeInt, Optional: true, Description: "JSI responses",
									},
									"jsi_pri_requests": {
										Type: schema.TypeInt, Optional: true, Description: "JSI primary requests (AJAX)",
									},
									"jsi_api_requests": {
										Type: schema.TypeInt, Optional: true, Description: "JSI API requests (AJAX)",
									},
									"jsi_api_responses": {
										Type: schema.TypeInt, Optional: true, Description: "JSI API responses",
									},
									"jsi_api_no_auth_hdr": {
										Type: schema.TypeInt, Optional: true, Description: "JSI API no auth hdr",
									},
									"jsi_api_no_token": {
										Type: schema.TypeInt, Optional: true, Description: "JSI API no token in auth hdr",
									},
									"jsi_skip_no_fi": {
										Type: schema.TypeInt, Optional: true, Description: "JSI skipped - File Inspection not configured",
									},
									"jsi_skip_no_ua": {
										Type: schema.TypeInt, Optional: true, Description: "JSI skipped - no User-Agent",
									},
									"jsi_skip_not_browser": {
										Type: schema.TypeInt, Optional: true, Description: "JSI skipped - not fromm browser",
									},
									"jsi_hash_add_fails": {
										Type: schema.TypeInt, Optional: true, Description: "JSI hash addition failures",
									},
									"jsi_hash_lookup_fails": {
										Type: schema.TypeInt, Optional: true, Description: "JSI hash lookup failures",
									},
									"header_length_long": {
										Type: schema.TypeInt, Optional: true, Description: "HTTP Header length too long",
									},
									"req_get": {
										Type: schema.TypeInt, Optional: true, Description: "Method GET",
									},
									"req_head": {
										Type: schema.TypeInt, Optional: true, Description: "Method HEAD",
									},
									"req_put": {
										Type: schema.TypeInt, Optional: true, Description: "Method PUT",
									},
									"req_post": {
										Type: schema.TypeInt, Optional: true, Description: "Method POST",
									},
									"req_trace": {
										Type: schema.TypeInt, Optional: true, Description: "Method TRACE",
									},
									"req_options": {
										Type: schema.TypeInt, Optional: true, Description: "Method OPTIONS",
									},
									"req_connect": {
										Type: schema.TypeInt, Optional: true, Description: "Method CONNECT",
									},
									"req_delete": {
										Type: schema.TypeInt, Optional: true, Description: "Method DELETE",
									},
									"req_unknown": {
										Type: schema.TypeInt, Optional: true, Description: "Method UNKNOWN",
									},
									"req_track": {
										Type: schema.TypeInt, Optional: true, Description: "Method TRACK",
									},
									"rsp_sz_1k": {
										Type: schema.TypeInt, Optional: true, Description: "Resp less than equal to 1K",
									},
									"rsp_sz_2k": {
										Type: schema.TypeInt, Optional: true, Description: "Resp less than equal to 2K",
									},
									"rsp_sz_4k": {
										Type: schema.TypeInt, Optional: true, Description: "Resp less than equal to 4K",
									},
									"rsp_sz_8k": {
										Type: schema.TypeInt, Optional: true, Description: "Resp less than equal to 8K",
									},
									"rsp_sz_16k": {
										Type: schema.TypeInt, Optional: true, Description: "Resp less than equal to 16K",
									},
									"rsp_sz_32k": {
										Type: schema.TypeInt, Optional: true, Description: "Resp less than equal to 32K",
									},
									"rsp_sz_64k": {
										Type: schema.TypeInt, Optional: true, Description: "Resp less than equal to 64K",
									},
									"rsp_sz_256k": {
										Type: schema.TypeInt, Optional: true, Description: "Resp less than equal to 256K",
									},
									"rsp_sz_gt_256k": {
										Type: schema.TypeInt, Optional: true, Description: "Resp greater than 256K",
									},
									"chunk_sz_512": {
										Type: schema.TypeInt, Optional: true, Description: "Chunk less than equal to 512",
									},
									"chunk_sz_1k": {
										Type: schema.TypeInt, Optional: true, Description: "Chunk less than equal to 1K",
									},
									"chunk_sz_2k": {
										Type: schema.TypeInt, Optional: true, Description: "Chunk less than equal to 2K",
									},
									"chunk_sz_4k": {
										Type: schema.TypeInt, Optional: true, Description: "Chunk less than equal to 4K",
									},
									"chunk_sz_gt_4k": {
										Type: schema.TypeInt, Optional: true, Description: "Chunk greater than 4K",
									},
									"req_sz_1k": {
										Type: schema.TypeInt, Optional: true, Description: "Req less than equal to 1K",
									},
									"req_sz_2k": {
										Type: schema.TypeInt, Optional: true, Description: "Req less than equal to 2K",
									},
									"req_sz_4k": {
										Type: schema.TypeInt, Optional: true, Description: "Req less than equal to 4K",
									},
									"req_sz_8k": {
										Type: schema.TypeInt, Optional: true, Description: "Req less than equal to 8K",
									},
									"req_sz_16k": {
										Type: schema.TypeInt, Optional: true, Description: "Req less than equal to 16K",
									},
									"req_sz_32k": {
										Type: schema.TypeInt, Optional: true, Description: "Req less than equal to 32K",
									},
									"req_sz_64k": {
										Type: schema.TypeInt, Optional: true, Description: "Req less than equal to 64K",
									},
									"req_sz_256k": {
										Type: schema.TypeInt, Optional: true, Description: "Req less than equal to 256K",
									},
									"req_sz_gt_256k": {
										Type: schema.TypeInt, Optional: true, Description: "Req greater than 256K",
									},
									"req_content_len": {
										Type: schema.TypeInt, Optional: true, Description: "Req content len",
									},
									"rsp_chunk": {
										Type: schema.TypeInt, Optional: true, Description: "Resp chunk encoding",
									},
									"doh_req": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Requests",
									},
									"doh_req_get": {
										Type: schema.TypeInt, Optional: true, Description: "DoH GET Requests",
									},
									"doh_req_post": {
										Type: schema.TypeInt, Optional: true, Description: "DoH POST Requests",
									},
									"doh_non_doh_req": {
										Type: schema.TypeInt, Optional: true, Description: "DoH non DoH Requests",
									},
									"doh_non_doh_req_get": {
										Type: schema.TypeInt, Optional: true, Description: "DoH non DoH GET Requests",
									},
									"doh_non_doh_req_post": {
										Type: schema.TypeInt, Optional: true, Description: "DoH non DoH POST Requests",
									},
									"doh_resp": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Responses",
									},
									"doh_tc_resp": {
										Type: schema.TypeInt, Optional: true, Description: "DoH TC Responses",
									},
									"doh_udp_dns_req": {
										Type: schema.TypeInt, Optional: true, Description: "DoH UDP DNS Requests",
									},
									"doh_udp_dns_resp": {
										Type: schema.TypeInt, Optional: true, Description: "DoH UDP DNS Responses",
									},
									"doh_tcp_dns_req": {
										Type: schema.TypeInt, Optional: true, Description: "DoH TCP DNS Requests",
									},
									"doh_tcp_dns_resp": {
										Type: schema.TypeInt, Optional: true, Description: "DoH TCP DNS Responses",
									},
									"doh_req_send_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Request Send Failed",
									},
									"doh_resp_send_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Response Send Failed",
									},
									"doh_malloc_fail": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Memory alloc failed",
									},
									"doh_req_udp_retry": {
										Type: schema.TypeInt, Optional: true, Description: "DoH UDP Retry",
									},
									"doh_req_udp_retry_fail": {
										Type: schema.TypeInt, Optional: true, Description: "DoH UDP Retry failed",
									},
									"doh_req_tcp_retry": {
										Type: schema.TypeInt, Optional: true, Description: "DoH TCP Retry",
									},
									"doh_req_tcp_retry_fail": {
										Type: schema.TypeInt, Optional: true, Description: "DoH TCP Retry failed",
									},
									"doh_snat_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Source NAT failed",
									},
									"doh_path_not_found": {
										Type: schema.TypeInt, Optional: true, Description: "DoH URI Path not found",
									},
									"doh_get_dns_arg_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DoH GET dns arg not found in uri",
									},
									"doh_get_base64_decode_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DoH GET base64url decode failed",
									},
									"doh_post_content_type_mismatch": {
										Type: schema.TypeInt, Optional: true, Description: "DoH POST content-type not found",
									},
									"doh_post_payload_not_found": {
										Type: schema.TypeInt, Optional: true, Description: "DoH POST payload not found",
									},
									"doh_post_payload_extract_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DoH POST payload extract failed",
									},
									"doh_non_doh_method": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Non DoH HTTP request method rcvd",
									},
									"doh_tcp_send_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DoH serv TCP DNS send failed",
									},
									"doh_udp_send_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DoH serv UDP DNS send failed",
									},
									"doh_query_time_out": {
										Type: schema.TypeInt, Optional: true, Description: "DoH serv Query timed out",
									},
									"doh_dns_query_type_a": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Query type A",
									},
									"doh_dns_query_type_aaaa": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Query type AAAA",
									},
									"doh_dns_query_type_ns": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Query type NS",
									},
									"doh_dns_query_type_cname": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Query type CNAME",
									},
									"doh_dns_query_type_any": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Query type ANY",
									},
									"doh_dns_query_type_srv": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Query type SRV",
									},
									"doh_dns_query_type_mx": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Query type MX",
									},
									"doh_dns_query_type_soa": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Query type SOA",
									},
									"doh_dns_query_type_others": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Query type Others",
									},
									"doh_resp_setup_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Response setup failed",
									},
									"doh_resp_header_alloc_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Resp hdr alloc failed",
									},
									"doh_resp_que_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Resp queue failed",
									},
									"doh_resp_udp_frags": {
										Type: schema.TypeInt, Optional: true, Description: "DoH UDP Frags Rcvd",
									},
									"doh_resp_tcp_frags": {
										Type: schema.TypeInt, Optional: true, Description: "DoH TCP Frags Rcvd",
									},
									"doh_serv_sel_failed": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Server Select Failed",
									},
									"doh_retry_w_tcp": {
										Type: schema.TypeInt, Optional: true, Description: "DoH Retry with TCP SG",
									},
									"doh_get_uri_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "DoH GET URI too long",
									},
									"doh_post_payload_too_large": {
										Type: schema.TypeInt, Optional: true, Description: "DoH POST Payload too large",
									},
									"doh_dns_malformed_query": {
										Type: schema.TypeInt, Optional: true, Description: "DoH DNS Malformed Query",
									},
									"doh_dns_resp_rcode_err_format": {
										Type: schema.TypeInt, Optional: true, Description: "DoH DNS Response rcode ERR_FORMAT",
									},
									"doh_dns_resp_rcode_err_server": {
										Type: schema.TypeInt, Optional: true, Description: "DoH DNS Response rcode ERR_SERVER",
									},
									"doh_dns_resp_rcode_err_name": {
										Type: schema.TypeInt, Optional: true, Description: "DoH DNS Response rcode ERR_NAME",
									},
									"doh_dns_resp_rcode_err_type": {
										Type: schema.TypeInt, Optional: true, Description: "DoH DNS Response rcode ERR_TYPE",
									},
									"doh_dns_resp_rcode_refuse": {
										Type: schema.TypeInt, Optional: true, Description: "DoH DNS Response rcode REFUSE",
									},
									"doh_dns_resp_rcode_yxdomain": {
										Type: schema.TypeInt, Optional: true, Description: "DoH DNS Response rcode YXDOMAIN",
									},
									"doh_dns_resp_rcode_yxrrset": {
										Type: schema.TypeInt, Optional: true, Description: "DoH DNS Response rcode YXRRSET",
									},
									"doh_dns_resp_rcode_nxrrset": {
										Type: schema.TypeInt, Optional: true, Description: "DoH DNS Response rcode NXRRSET",
									},
									"doh_dns_resp_rcode_notauth": {
										Type: schema.TypeInt, Optional: true, Description: "DoH DNS Response rcode NOTAUTH",
									},
									"doh_dns_resp_rcode_notzone": {
										Type: schema.TypeInt, Optional: true, Description: "DoH DNS Response rcode NOTZONE",
									},
									"doh_dns_resp_rcode_other": {
										Type: schema.TypeInt, Optional: true, Description: "DoH DNS Response rcode OTHER",
									},
									"curr_http3_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Current number of HTTP/3 connections",
									},
									"total_http3_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Total number of HTTP/3 connections",
									},
									"peak_http3_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Largest number of concurrent HTTP/3 connections",
									},
									"total_http3_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Total bytes in HTTP/3 frames",
									},
									"http3_control_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Bytes in HTTP/3 control frames",
									},
									"http3_header_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Bytes in HTTP/3 header frames",
									},
									"http3_data_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Bytes in HTTP/3 data frames",
									},
									"http3_goaway_received": {
										Type: schema.TypeInt, Optional: true, Description: "GoAway frames received",
									},
									"http3_goaway_sent": {
										Type: schema.TypeInt, Optional: true, Description: "GoAway frames sent",
									},
									"http3_stream_closed": {
										Type: schema.TypeInt, Optional: true, Description: "HTTP/3 streams closed",
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
func resourceSlbVirtualServerPortStats50Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats50Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats50(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortStats50Read(ctx, d, meta)
	}
	return diags
}

func resourceSlbVirtualServerPortStats50Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats50Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats50(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortStats50Read(ctx, d, meta)
	}
	return diags
}
func resourceSlbVirtualServerPortStats50Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats50Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats50(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbVirtualServerPortStats50Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats50Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats50(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbVirtualServerPortStats50Stats(d []interface{}) edpt.SlbVirtualServerPortStats50Stats {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats50Stats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Http_vport = getObjectSlbVirtualServerPortStats50StatsHttp_vport(in["http_vport"].([]interface{}))
	}
	return ret
}

func getObjectSlbVirtualServerPortStats50StatsHttp_vport(d []interface{}) edpt.SlbVirtualServerPortStats50StatsHttp_vport {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats50StatsHttp_vport
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Status_200 = in["status_200"].(int)
		ret.Status_201 = in["status_201"].(int)
		ret.Status_202 = in["status_202"].(int)
		ret.Status_203 = in["status_203"].(int)
		ret.Status_204 = in["status_204"].(int)
		ret.Status_205 = in["status_205"].(int)
		ret.Status_206 = in["status_206"].(int)
		ret.Status_207 = in["status_207"].(int)
		ret.Status_100 = in["status_100"].(int)
		ret.Status_101 = in["status_101"].(int)
		ret.Status_102 = in["status_102"].(int)
		ret.Status_103 = in["status_103"].(int)
		ret.Status_300 = in["status_300"].(int)
		ret.Status_301 = in["status_301"].(int)
		ret.Status_302 = in["status_302"].(int)
		ret.Status_303 = in["status_303"].(int)
		ret.Status_304 = in["status_304"].(int)
		ret.Status_305 = in["status_305"].(int)
		ret.Status_306 = in["status_306"].(int)
		ret.Status_307 = in["status_307"].(int)
		ret.Status_400 = in["status_400"].(int)
		ret.Status_401 = in["status_401"].(int)
		ret.Status_402 = in["status_402"].(int)
		ret.Status_403 = in["status_403"].(int)
		ret.Status_404 = in["status_404"].(int)
		ret.Status_405 = in["status_405"].(int)
		ret.Status_406 = in["status_406"].(int)
		ret.Status_407 = in["status_407"].(int)
		ret.Status_408 = in["status_408"].(int)
		ret.Status_409 = in["status_409"].(int)
		ret.Status_410 = in["status_410"].(int)
		ret.Status_411 = in["status_411"].(int)
		ret.Status_412 = in["status_412"].(int)
		ret.Status_413 = in["status_413"].(int)
		ret.Status_414 = in["status_414"].(int)
		ret.Status_415 = in["status_415"].(int)
		ret.Status_416 = in["status_416"].(int)
		ret.Status_417 = in["status_417"].(int)
		ret.Status_418 = in["status_418"].(int)
		ret.Status_422 = in["status_422"].(int)
		ret.Status_423 = in["status_423"].(int)
		ret.Status_424 = in["status_424"].(int)
		ret.Status_425 = in["status_425"].(int)
		ret.Status_426 = in["status_426"].(int)
		ret.Status_449 = in["status_449"].(int)
		ret.Status_450 = in["status_450"].(int)
		ret.Status_500 = in["status_500"].(int)
		ret.Status_501 = in["status_501"].(int)
		ret.Status_502 = in["status_502"].(int)
		ret.Status_503 = in["status_503"].(int)
		ret.Status_504 = in["status_504"].(int)
		ret.Status_504_ax = in["status_504_ax"].(int)
		ret.Status_505 = in["status_505"].(int)
		ret.Status_506 = in["status_506"].(int)
		ret.Status_507 = in["status_507"].(int)
		ret.Status_508 = in["status_508"].(int)
		ret.Status_509 = in["status_509"].(int)
		ret.Status_510 = in["status_510"].(int)
		ret.Status_1xx = in["status_1xx"].(int)
		ret.Status_2xx = in["status_2xx"].(int)
		ret.Status_3xx = in["status_3xx"].(int)
		ret.Status_4xx = in["status_4xx"].(int)
		ret.Status_5xx = in["status_5xx"].(int)
		ret.Status_6xx = in["status_6xx"].(int)
		ret.Status_unknown = in["status_unknown"].(int)
		ret.Ws_handshake_request = in["ws_handshake_request"].(int)
		ret.Ws_handshake_success = in["ws_handshake_success"].(int)
		ret.Ws_client_switch = in["ws_client_switch"].(int)
		ret.Ws_server_switch = in["ws_server_switch"].(int)
		ret.Req_10u = in["req_10u"].(int)
		ret.Req_20u = in["req_20u"].(int)
		ret.Req_50u = in["req_50u"].(int)
		ret.Req_100u = in["req_100u"].(int)
		ret.Req_200u = in["req_200u"].(int)
		ret.Req_500u = in["req_500u"].(int)
		ret.Req_1m = in["req_1m"].(int)
		ret.Req_2m = in["req_2m"].(int)
		ret.Req_5m = in["req_5m"].(int)
		ret.Req_10m = in["req_10m"].(int)
		ret.Req_20m = in["req_20m"].(int)
		ret.Req_50m = in["req_50m"].(int)
		ret.Req_100m = in["req_100m"].(int)
		ret.Req_200m = in["req_200m"].(int)
		ret.Req_500m = in["req_500m"].(int)
		ret.Req_1s = in["req_1s"].(int)
		ret.Req_2s = in["req_2s"].(int)
		ret.Req_5s = in["req_5s"].(int)
		ret.Req_over_5s = in["req_over_5s"].(int)
		ret.Total_requests = in["total_requests"].(int)
		ret.Curr_http2_conn = in["curr_http2_conn"].(int)
		ret.Total_http2_conn = in["total_http2_conn"].(int)
		ret.Peak_http2_conn = in["peak_http2_conn"].(int)
		ret.Total_http2_bytes = in["total_http2_bytes"].(int)
		ret.Http2_control_bytes = in["http2_control_bytes"].(int)
		ret.Http2_header_bytes = in["http2_header_bytes"].(int)
		ret.Http2_data_bytes = in["http2_data_bytes"].(int)
		ret.Http2_reset_received = in["http2_reset_received"].(int)
		ret.Http2_reset_sent = in["http2_reset_sent"].(int)
		ret.Http2_goaway_received = in["http2_goaway_received"].(int)
		ret.Http2_goaway_sent = in["http2_goaway_sent"].(int)
		ret.Stream_closed = in["stream_closed"].(int)
		ret.Jsi_requests = in["jsi_requests"].(int)
		ret.Jsi_responses = in["jsi_responses"].(int)
		ret.Jsi_pri_requests = in["jsi_pri_requests"].(int)
		ret.Jsi_api_requests = in["jsi_api_requests"].(int)
		ret.Jsi_api_responses = in["jsi_api_responses"].(int)
		ret.Jsi_api_no_auth_hdr = in["jsi_api_no_auth_hdr"].(int)
		ret.Jsi_api_no_token = in["jsi_api_no_token"].(int)
		ret.Jsi_skip_no_fi = in["jsi_skip_no_fi"].(int)
		ret.Jsi_skip_no_ua = in["jsi_skip_no_ua"].(int)
		ret.Jsi_skip_not_browser = in["jsi_skip_not_browser"].(int)
		ret.Jsi_hash_add_fails = in["jsi_hash_add_fails"].(int)
		ret.Jsi_hash_lookup_fails = in["jsi_hash_lookup_fails"].(int)
		ret.Header_length_long = in["header_length_long"].(int)
		ret.Req_get = in["req_get"].(int)
		ret.Req_head = in["req_head"].(int)
		ret.Req_put = in["req_put"].(int)
		ret.Req_post = in["req_post"].(int)
		ret.Req_trace = in["req_trace"].(int)
		ret.Req_options = in["req_options"].(int)
		ret.Req_connect = in["req_connect"].(int)
		ret.Req_delete = in["req_delete"].(int)
		ret.Req_unknown = in["req_unknown"].(int)
		ret.Req_track = in["req_track"].(int)
		ret.Rsp_sz_1k = in["rsp_sz_1k"].(int)
		ret.Rsp_sz_2k = in["rsp_sz_2k"].(int)
		ret.Rsp_sz_4k = in["rsp_sz_4k"].(int)
		ret.Rsp_sz_8k = in["rsp_sz_8k"].(int)
		ret.Rsp_sz_16k = in["rsp_sz_16k"].(int)
		ret.Rsp_sz_32k = in["rsp_sz_32k"].(int)
		ret.Rsp_sz_64k = in["rsp_sz_64k"].(int)
		ret.Rsp_sz_256k = in["rsp_sz_256k"].(int)
		ret.Rsp_sz_gt_256k = in["rsp_sz_gt_256k"].(int)
		ret.Chunk_sz_512 = in["chunk_sz_512"].(int)
		ret.Chunk_sz_1k = in["chunk_sz_1k"].(int)
		ret.Chunk_sz_2k = in["chunk_sz_2k"].(int)
		ret.Chunk_sz_4k = in["chunk_sz_4k"].(int)
		ret.Chunk_sz_gt_4k = in["chunk_sz_gt_4k"].(int)
		ret.Req_sz_1k = in["req_sz_1k"].(int)
		ret.Req_sz_2k = in["req_sz_2k"].(int)
		ret.Req_sz_4k = in["req_sz_4k"].(int)
		ret.Req_sz_8k = in["req_sz_8k"].(int)
		ret.Req_sz_16k = in["req_sz_16k"].(int)
		ret.Req_sz_32k = in["req_sz_32k"].(int)
		ret.Req_sz_64k = in["req_sz_64k"].(int)
		ret.Req_sz_256k = in["req_sz_256k"].(int)
		ret.Req_sz_gt_256k = in["req_sz_gt_256k"].(int)
		ret.Req_content_len = in["req_content_len"].(int)
		ret.Rsp_chunk = in["rsp_chunk"].(int)
		ret.Doh_req = in["doh_req"].(int)
		ret.Doh_req_get = in["doh_req_get"].(int)
		ret.Doh_req_post = in["doh_req_post"].(int)
		ret.Doh_non_doh_req = in["doh_non_doh_req"].(int)
		ret.Doh_non_doh_req_get = in["doh_non_doh_req_get"].(int)
		ret.Doh_non_doh_req_post = in["doh_non_doh_req_post"].(int)
		ret.Doh_resp = in["doh_resp"].(int)
		ret.Doh_tc_resp = in["doh_tc_resp"].(int)
		ret.Doh_udp_dns_req = in["doh_udp_dns_req"].(int)
		ret.Doh_udp_dns_resp = in["doh_udp_dns_resp"].(int)
		ret.Doh_tcp_dns_req = in["doh_tcp_dns_req"].(int)
		ret.Doh_tcp_dns_resp = in["doh_tcp_dns_resp"].(int)
		ret.Doh_req_send_failed = in["doh_req_send_failed"].(int)
		ret.Doh_resp_send_failed = in["doh_resp_send_failed"].(int)
		ret.Doh_malloc_fail = in["doh_malloc_fail"].(int)
		ret.Doh_req_udp_retry = in["doh_req_udp_retry"].(int)
		ret.Doh_req_udp_retry_fail = in["doh_req_udp_retry_fail"].(int)
		ret.Doh_req_tcp_retry = in["doh_req_tcp_retry"].(int)
		ret.Doh_req_tcp_retry_fail = in["doh_req_tcp_retry_fail"].(int)
		ret.Doh_snat_failed = in["doh_snat_failed"].(int)
		ret.Doh_path_not_found = in["doh_path_not_found"].(int)
		ret.Doh_get_dns_arg_failed = in["doh_get_dns_arg_failed"].(int)
		ret.Doh_get_base64_decode_failed = in["doh_get_base64_decode_failed"].(int)
		ret.Doh_post_content_type_mismatch = in["doh_post_content_type_mismatch"].(int)
		ret.Doh_post_payload_not_found = in["doh_post_payload_not_found"].(int)
		ret.Doh_post_payload_extract_failed = in["doh_post_payload_extract_failed"].(int)
		ret.Doh_non_doh_method = in["doh_non_doh_method"].(int)
		ret.Doh_tcp_send_failed = in["doh_tcp_send_failed"].(int)
		ret.Doh_udp_send_failed = in["doh_udp_send_failed"].(int)
		ret.Doh_query_time_out = in["doh_query_time_out"].(int)
		ret.Doh_dns_query_type_a = in["doh_dns_query_type_a"].(int)
		ret.Doh_dns_query_type_aaaa = in["doh_dns_query_type_aaaa"].(int)
		ret.Doh_dns_query_type_ns = in["doh_dns_query_type_ns"].(int)
		ret.Doh_dns_query_type_cname = in["doh_dns_query_type_cname"].(int)
		ret.Doh_dns_query_type_any = in["doh_dns_query_type_any"].(int)
		ret.Doh_dns_query_type_srv = in["doh_dns_query_type_srv"].(int)
		ret.Doh_dns_query_type_mx = in["doh_dns_query_type_mx"].(int)
		ret.Doh_dns_query_type_soa = in["doh_dns_query_type_soa"].(int)
		ret.Doh_dns_query_type_others = in["doh_dns_query_type_others"].(int)
		ret.Doh_resp_setup_failed = in["doh_resp_setup_failed"].(int)
		ret.Doh_resp_header_alloc_failed = in["doh_resp_header_alloc_failed"].(int)
		ret.Doh_resp_que_failed = in["doh_resp_que_failed"].(int)
		ret.Doh_resp_udp_frags = in["doh_resp_udp_frags"].(int)
		ret.Doh_resp_tcp_frags = in["doh_resp_tcp_frags"].(int)
		ret.Doh_serv_sel_failed = in["doh_serv_sel_failed"].(int)
		ret.Doh_retry_w_tcp = in["doh_retry_w_tcp"].(int)
		ret.Doh_get_uri_too_long = in["doh_get_uri_too_long"].(int)
		ret.Doh_post_payload_too_large = in["doh_post_payload_too_large"].(int)
		ret.Doh_dns_malformed_query = in["doh_dns_malformed_query"].(int)
		ret.Doh_dns_resp_rcode_err_format = in["doh_dns_resp_rcode_err_format"].(int)
		ret.Doh_dns_resp_rcode_err_server = in["doh_dns_resp_rcode_err_server"].(int)
		ret.Doh_dns_resp_rcode_err_name = in["doh_dns_resp_rcode_err_name"].(int)
		ret.Doh_dns_resp_rcode_err_type = in["doh_dns_resp_rcode_err_type"].(int)
		ret.Doh_dns_resp_rcode_refuse = in["doh_dns_resp_rcode_refuse"].(int)
		ret.Doh_dns_resp_rcode_yxdomain = in["doh_dns_resp_rcode_yxdomain"].(int)
		ret.Doh_dns_resp_rcode_yxrrset = in["doh_dns_resp_rcode_yxrrset"].(int)
		ret.Doh_dns_resp_rcode_nxrrset = in["doh_dns_resp_rcode_nxrrset"].(int)
		ret.Doh_dns_resp_rcode_notauth = in["doh_dns_resp_rcode_notauth"].(int)
		ret.Doh_dns_resp_rcode_notzone = in["doh_dns_resp_rcode_notzone"].(int)
		ret.Doh_dns_resp_rcode_other = in["doh_dns_resp_rcode_other"].(int)
		ret.Curr_http3_conn = in["curr_http3_conn"].(int)
		ret.Total_http3_conn = in["total_http3_conn"].(int)
		ret.Peak_http3_conn = in["peak_http3_conn"].(int)
		ret.Total_http3_bytes = in["total_http3_bytes"].(int)
		ret.Http3_control_bytes = in["http3_control_bytes"].(int)
		ret.Http3_header_bytes = in["http3_header_bytes"].(int)
		ret.Http3_data_bytes = in["http3_data_bytes"].(int)
		ret.Http3_goaway_received = in["http3_goaway_received"].(int)
		ret.Http3_goaway_sent = in["http3_goaway_sent"].(int)
		ret.Http3_stream_closed = in["http3_stream_closed"].(int)
	}
	return ret
}

func dataToEndpointSlbVirtualServerPortStats50(d *schema.ResourceData) edpt.SlbVirtualServerPortStats50 {
	var ret edpt.SlbVirtualServerPortStats50
	ret.Inst.PortNumber = d.Get("port_number").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Stats = getObjectSlbVirtualServerPortStats50Stats(d.Get("stats").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
