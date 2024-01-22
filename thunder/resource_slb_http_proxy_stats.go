package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHttpProxyStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_http_proxy_stats`: Statistics for the object http-proxy\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbHttpProxyStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"curr_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Curr Proxy Conns",
						},
						"total_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Total Proxy Conns",
						},
						"req": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP requests",
						},
						"req_succ": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP requests(succ)",
						},
						"noproxy": {
							Type: schema.TypeInt, Optional: true, Description: "No proxy error",
						},
						"client_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Client RST",
						},
						"server_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Server RST",
						},
						"notuple": {
							Type: schema.TypeInt, Optional: true, Description: "No tuple error",
						},
						"parsereq_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Parse req fail",
						},
						"svrsel_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Server selection fail",
						},
						"fwdreq_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Fwd req fail",
						},
						"fwdreqdata_fail": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"req_retran": {
							Type: schema.TypeInt, Optional: true, Description: "Packets retrans",
						},
						"req_ofo": {
							Type: schema.TypeInt, Optional: true, Description: "Packets ofo",
						},
						"server_resel": {
							Type: schema.TypeInt, Optional: true, Description: "Server reselection",
						},
						"svr_prem_close": {
							Type: schema.TypeInt, Optional: true, Description: "Server premature close",
						},
						"new_svrconn": {
							Type: schema.TypeInt, Optional: true, Description: "Server conn made",
						},
						"snat_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Source NAT failure",
						},
						"req_over_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Request over limit",
						},
						"req_rate_over_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Request rate over limit",
						},
						"compression_before": {
							Type: schema.TypeInt, Optional: true, Description: "Tot data before compress",
						},
						"compression_after": {
							Type: schema.TypeInt, Optional: true, Description: "Tot data after compress",
						},
						"response_1xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 1XX",
						},
						"response_100": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 100",
						},
						"response_101": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 101",
						},
						"response_102": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 102",
						},
						"response_2xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 2XX",
						},
						"response_200": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 200",
						},
						"response_201": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 201",
						},
						"response_202": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 202",
						},
						"response_203": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 203",
						},
						"response_204": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 204",
						},
						"response_205": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 205",
						},
						"response_206": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 206",
						},
						"response_207": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 207",
						},
						"response_3xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 3XX",
						},
						"response_300": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 300",
						},
						"response_301": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 301",
						},
						"response_302": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 302",
						},
						"response_303": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 303",
						},
						"response_304": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 304",
						},
						"response_305": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 305",
						},
						"response_306": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 306",
						},
						"response_307": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 307",
						},
						"response_4xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 4XX",
						},
						"response_400": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 400",
						},
						"response_401": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 401",
						},
						"response_402": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 402",
						},
						"response_403": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 403",
						},
						"response_404": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 404",
						},
						"response_405": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 405",
						},
						"response_406": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 406",
						},
						"response_407": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 407",
						},
						"response_408": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 408",
						},
						"response_409": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 409",
						},
						"response_410": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 410",
						},
						"response_411": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 411",
						},
						"response_412": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 412",
						},
						"response_413": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 413",
						},
						"response_414": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 414",
						},
						"response_415": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 415",
						},
						"response_416": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 416",
						},
						"response_417": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 417",
						},
						"response_418": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 418",
						},
						"response_422": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 422",
						},
						"response_423": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 423",
						},
						"response_424": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 424",
						},
						"response_425": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 425",
						},
						"response_426": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 426",
						},
						"response_449": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 449",
						},
						"response_450": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 450",
						},
						"response_5xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 5XX",
						},
						"response_500": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 500",
						},
						"response_501": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 501",
						},
						"response_502": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 502",
						},
						"response_503": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 503",
						},
						"response_504": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 504",
						},
						"response_505": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 505",
						},
						"response_506": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 506",
						},
						"response_507": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 507",
						},
						"response_508": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 508",
						},
						"response_509": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 509",
						},
						"response_510": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 510",
						},
						"response_6xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status code 6XX",
						},
						"response_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "Status code unknown",
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
						"req_content_len": {
							Type: schema.TypeInt, Optional: true, Description: "Req content len",
						},
						"rsp_content_len": {
							Type: schema.TypeInt, Optional: true, Description: "Resp content len",
						},
						"rsp_chunk": {
							Type: schema.TypeInt, Optional: true, Description: "Resp chunk encoding",
						},
						"cache_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP req (cache succ)",
						},
						"close_on_ddos": {
							Type: schema.TypeInt, Optional: true, Description: "Close on DDoS",
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
							Type: schema.TypeInt, Optional: true, Description: "Rsp time less than 50m",
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
						"req_track": {
							Type: schema.TypeInt, Optional: true, Description: "Method TRACK",
						},
						"connect_req": {
							Type: schema.TypeInt, Optional: true, Description: "Total HTTP CONNECT requests",
						},
						"req_enter_ssli": {
							Type: schema.TypeInt, Optional: true, Description: "Total HTTP requests enter SSLi",
						},
						"decompression_before": {
							Type: schema.TypeInt, Optional: true, Description: "Tot data before decompress",
						},
						"decompression_after": {
							Type: schema.TypeInt, Optional: true, Description: "Tot data after decompress",
						},
						"req_http2": {
							Type: schema.TypeInt, Optional: true, Description: "Request 2.0",
						},
						"response_http2": {
							Type: schema.TypeInt, Optional: true, Description: "Resp 2.0",
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
						"compression_before_br": {
							Type: schema.TypeInt, Optional: true, Description: "Tot data before brotli compress",
						},
						"compression_after_br": {
							Type: schema.TypeInt, Optional: true, Description: "Tot data after brotli compress",
						},
						"compression_before_total": {
							Type: schema.TypeInt, Optional: true, Description: "Tot data before compress",
						},
						"compression_after_total": {
							Type: schema.TypeInt, Optional: true, Description: "Tot data after compress",
						},
						"decompression_before_br": {
							Type: schema.TypeInt, Optional: true, Description: "Tot data before brotli decompress",
						},
						"decompression_after_br": {
							Type: schema.TypeInt, Optional: true, Description: "Tot data after brotli decompress",
						},
						"decompression_before_total": {
							Type: schema.TypeInt, Optional: true, Description: "Tot data before decompress",
						},
						"decompression_after_total": {
							Type: schema.TypeInt, Optional: true, Description: "Tot data after decompress",
						},
						"req_http3": {
							Type: schema.TypeInt, Optional: true, Description: "Request 3.0",
						},
						"response_http3": {
							Type: schema.TypeInt, Optional: true, Description: "Resp 3.0",
						},
					},
				},
			},
		},
	}
}

func resourceSlbHttpProxyStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttpProxyStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttpProxyStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbHttpProxyStatsStats := setObjectSlbHttpProxyStatsStats(res)
		d.Set("stats", SlbHttpProxyStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHttpProxyStatsStats(ret edpt.DataSlbHttpProxyStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_proxy":                      ret.DtSlbHttpProxyStats.Stats.Curr_proxy,
			"total_proxy":                     ret.DtSlbHttpProxyStats.Stats.Total_proxy,
			"req":                             ret.DtSlbHttpProxyStats.Stats.Req,
			"req_succ":                        ret.DtSlbHttpProxyStats.Stats.Req_succ,
			"noproxy":                         ret.DtSlbHttpProxyStats.Stats.Noproxy,
			"client_rst":                      ret.DtSlbHttpProxyStats.Stats.Client_rst,
			"server_rst":                      ret.DtSlbHttpProxyStats.Stats.Server_rst,
			"notuple":                         ret.DtSlbHttpProxyStats.Stats.Notuple,
			"parsereq_fail":                   ret.DtSlbHttpProxyStats.Stats.Parsereq_fail,
			"svrsel_fail":                     ret.DtSlbHttpProxyStats.Stats.Svrsel_fail,
			"fwdreq_fail":                     ret.DtSlbHttpProxyStats.Stats.Fwdreq_fail,
			"fwdreqdata_fail":                 ret.DtSlbHttpProxyStats.Stats.Fwdreqdata_fail,
			"req_retran":                      ret.DtSlbHttpProxyStats.Stats.Req_retran,
			"req_ofo":                         ret.DtSlbHttpProxyStats.Stats.Req_ofo,
			"server_resel":                    ret.DtSlbHttpProxyStats.Stats.Server_resel,
			"svr_prem_close":                  ret.DtSlbHttpProxyStats.Stats.Svr_prem_close,
			"new_svrconn":                     ret.DtSlbHttpProxyStats.Stats.New_svrconn,
			"snat_fail":                       ret.DtSlbHttpProxyStats.Stats.Snat_fail,
			"req_over_limit":                  ret.DtSlbHttpProxyStats.Stats.Req_over_limit,
			"req_rate_over_limit":             ret.DtSlbHttpProxyStats.Stats.Req_rate_over_limit,
			"compression_before":              ret.DtSlbHttpProxyStats.Stats.Compression_before,
			"compression_after":               ret.DtSlbHttpProxyStats.Stats.Compression_after,
			"response_1xx":                    ret.DtSlbHttpProxyStats.Stats.Response_1xx,
			"response_100":                    ret.DtSlbHttpProxyStats.Stats.Response_100,
			"response_101":                    ret.DtSlbHttpProxyStats.Stats.Response_101,
			"response_102":                    ret.DtSlbHttpProxyStats.Stats.Response_102,
			"response_2xx":                    ret.DtSlbHttpProxyStats.Stats.Response_2xx,
			"response_200":                    ret.DtSlbHttpProxyStats.Stats.Response_200,
			"response_201":                    ret.DtSlbHttpProxyStats.Stats.Response_201,
			"response_202":                    ret.DtSlbHttpProxyStats.Stats.Response_202,
			"response_203":                    ret.DtSlbHttpProxyStats.Stats.Response_203,
			"response_204":                    ret.DtSlbHttpProxyStats.Stats.Response_204,
			"response_205":                    ret.DtSlbHttpProxyStats.Stats.Response_205,
			"response_206":                    ret.DtSlbHttpProxyStats.Stats.Response_206,
			"response_207":                    ret.DtSlbHttpProxyStats.Stats.Response_207,
			"response_3xx":                    ret.DtSlbHttpProxyStats.Stats.Response_3xx,
			"response_300":                    ret.DtSlbHttpProxyStats.Stats.Response_300,
			"response_301":                    ret.DtSlbHttpProxyStats.Stats.Response_301,
			"response_302":                    ret.DtSlbHttpProxyStats.Stats.Response_302,
			"response_303":                    ret.DtSlbHttpProxyStats.Stats.Response_303,
			"response_304":                    ret.DtSlbHttpProxyStats.Stats.Response_304,
			"response_305":                    ret.DtSlbHttpProxyStats.Stats.Response_305,
			"response_306":                    ret.DtSlbHttpProxyStats.Stats.Response_306,
			"response_307":                    ret.DtSlbHttpProxyStats.Stats.Response_307,
			"response_4xx":                    ret.DtSlbHttpProxyStats.Stats.Response_4xx,
			"response_400":                    ret.DtSlbHttpProxyStats.Stats.Response_400,
			"response_401":                    ret.DtSlbHttpProxyStats.Stats.Response_401,
			"response_402":                    ret.DtSlbHttpProxyStats.Stats.Response_402,
			"response_403":                    ret.DtSlbHttpProxyStats.Stats.Response_403,
			"response_404":                    ret.DtSlbHttpProxyStats.Stats.Response_404,
			"response_405":                    ret.DtSlbHttpProxyStats.Stats.Response_405,
			"response_406":                    ret.DtSlbHttpProxyStats.Stats.Response_406,
			"response_407":                    ret.DtSlbHttpProxyStats.Stats.Response_407,
			"response_408":                    ret.DtSlbHttpProxyStats.Stats.Response_408,
			"response_409":                    ret.DtSlbHttpProxyStats.Stats.Response_409,
			"response_410":                    ret.DtSlbHttpProxyStats.Stats.Response_410,
			"response_411":                    ret.DtSlbHttpProxyStats.Stats.Response_411,
			"response_412":                    ret.DtSlbHttpProxyStats.Stats.Response_412,
			"response_413":                    ret.DtSlbHttpProxyStats.Stats.Response_413,
			"response_414":                    ret.DtSlbHttpProxyStats.Stats.Response_414,
			"response_415":                    ret.DtSlbHttpProxyStats.Stats.Response_415,
			"response_416":                    ret.DtSlbHttpProxyStats.Stats.Response_416,
			"response_417":                    ret.DtSlbHttpProxyStats.Stats.Response_417,
			"response_418":                    ret.DtSlbHttpProxyStats.Stats.Response_418,
			"response_422":                    ret.DtSlbHttpProxyStats.Stats.Response_422,
			"response_423":                    ret.DtSlbHttpProxyStats.Stats.Response_423,
			"response_424":                    ret.DtSlbHttpProxyStats.Stats.Response_424,
			"response_425":                    ret.DtSlbHttpProxyStats.Stats.Response_425,
			"response_426":                    ret.DtSlbHttpProxyStats.Stats.Response_426,
			"response_449":                    ret.DtSlbHttpProxyStats.Stats.Response_449,
			"response_450":                    ret.DtSlbHttpProxyStats.Stats.Response_450,
			"response_5xx":                    ret.DtSlbHttpProxyStats.Stats.Response_5xx,
			"response_500":                    ret.DtSlbHttpProxyStats.Stats.Response_500,
			"response_501":                    ret.DtSlbHttpProxyStats.Stats.Response_501,
			"response_502":                    ret.DtSlbHttpProxyStats.Stats.Response_502,
			"response_503":                    ret.DtSlbHttpProxyStats.Stats.Response_503,
			"response_504":                    ret.DtSlbHttpProxyStats.Stats.Response_504,
			"response_505":                    ret.DtSlbHttpProxyStats.Stats.Response_505,
			"response_506":                    ret.DtSlbHttpProxyStats.Stats.Response_506,
			"response_507":                    ret.DtSlbHttpProxyStats.Stats.Response_507,
			"response_508":                    ret.DtSlbHttpProxyStats.Stats.Response_508,
			"response_509":                    ret.DtSlbHttpProxyStats.Stats.Response_509,
			"response_510":                    ret.DtSlbHttpProxyStats.Stats.Response_510,
			"response_6xx":                    ret.DtSlbHttpProxyStats.Stats.Response_6xx,
			"response_unknown":                ret.DtSlbHttpProxyStats.Stats.Response_unknown,
			"req_get":                         ret.DtSlbHttpProxyStats.Stats.Req_get,
			"req_head":                        ret.DtSlbHttpProxyStats.Stats.Req_head,
			"req_put":                         ret.DtSlbHttpProxyStats.Stats.Req_put,
			"req_post":                        ret.DtSlbHttpProxyStats.Stats.Req_post,
			"req_trace":                       ret.DtSlbHttpProxyStats.Stats.Req_trace,
			"req_options":                     ret.DtSlbHttpProxyStats.Stats.Req_options,
			"req_connect":                     ret.DtSlbHttpProxyStats.Stats.Req_connect,
			"req_delete":                      ret.DtSlbHttpProxyStats.Stats.Req_delete,
			"req_unknown":                     ret.DtSlbHttpProxyStats.Stats.Req_unknown,
			"req_content_len":                 ret.DtSlbHttpProxyStats.Stats.Req_content_len,
			"rsp_content_len":                 ret.DtSlbHttpProxyStats.Stats.Rsp_content_len,
			"rsp_chunk":                       ret.DtSlbHttpProxyStats.Stats.Rsp_chunk,
			"cache_rsp":                       ret.DtSlbHttpProxyStats.Stats.Cache_rsp,
			"close_on_ddos":                   ret.DtSlbHttpProxyStats.Stats.Close_on_ddos,
			"req_sz_1k":                       ret.DtSlbHttpProxyStats.Stats.Req_sz_1k,
			"req_sz_2k":                       ret.DtSlbHttpProxyStats.Stats.Req_sz_2k,
			"req_sz_4k":                       ret.DtSlbHttpProxyStats.Stats.Req_sz_4k,
			"req_sz_8k":                       ret.DtSlbHttpProxyStats.Stats.Req_sz_8k,
			"req_sz_16k":                      ret.DtSlbHttpProxyStats.Stats.Req_sz_16k,
			"req_sz_32k":                      ret.DtSlbHttpProxyStats.Stats.Req_sz_32k,
			"req_sz_64k":                      ret.DtSlbHttpProxyStats.Stats.Req_sz_64k,
			"req_sz_256k":                     ret.DtSlbHttpProxyStats.Stats.Req_sz_256k,
			"req_sz_gt_256k":                  ret.DtSlbHttpProxyStats.Stats.Req_sz_gt_256k,
			"rsp_sz_1k":                       ret.DtSlbHttpProxyStats.Stats.Rsp_sz_1k,
			"rsp_sz_2k":                       ret.DtSlbHttpProxyStats.Stats.Rsp_sz_2k,
			"rsp_sz_4k":                       ret.DtSlbHttpProxyStats.Stats.Rsp_sz_4k,
			"rsp_sz_8k":                       ret.DtSlbHttpProxyStats.Stats.Rsp_sz_8k,
			"rsp_sz_16k":                      ret.DtSlbHttpProxyStats.Stats.Rsp_sz_16k,
			"rsp_sz_32k":                      ret.DtSlbHttpProxyStats.Stats.Rsp_sz_32k,
			"rsp_sz_64k":                      ret.DtSlbHttpProxyStats.Stats.Rsp_sz_64k,
			"rsp_sz_256k":                     ret.DtSlbHttpProxyStats.Stats.Rsp_sz_256k,
			"rsp_sz_gt_256k":                  ret.DtSlbHttpProxyStats.Stats.Rsp_sz_gt_256k,
			"chunk_sz_512":                    ret.DtSlbHttpProxyStats.Stats.Chunk_sz_512,
			"chunk_sz_1k":                     ret.DtSlbHttpProxyStats.Stats.Chunk_sz_1k,
			"chunk_sz_2k":                     ret.DtSlbHttpProxyStats.Stats.Chunk_sz_2k,
			"chunk_sz_4k":                     ret.DtSlbHttpProxyStats.Stats.Chunk_sz_4k,
			"chunk_sz_gt_4k":                  ret.DtSlbHttpProxyStats.Stats.Chunk_sz_gt_4k,
			"req_10u":                         ret.DtSlbHttpProxyStats.Stats.Req_10u,
			"req_20u":                         ret.DtSlbHttpProxyStats.Stats.Req_20u,
			"req_50u":                         ret.DtSlbHttpProxyStats.Stats.Req_50u,
			"req_100u":                        ret.DtSlbHttpProxyStats.Stats.Req_100u,
			"req_200u":                        ret.DtSlbHttpProxyStats.Stats.Req_200u,
			"req_500u":                        ret.DtSlbHttpProxyStats.Stats.Req_500u,
			"req_1m":                          ret.DtSlbHttpProxyStats.Stats.Req_1m,
			"req_2m":                          ret.DtSlbHttpProxyStats.Stats.Req_2m,
			"req_5m":                          ret.DtSlbHttpProxyStats.Stats.Req_5m,
			"req_10m":                         ret.DtSlbHttpProxyStats.Stats.Req_10m,
			"req_20m":                         ret.DtSlbHttpProxyStats.Stats.Req_20m,
			"req_50m":                         ret.DtSlbHttpProxyStats.Stats.Req_50m,
			"req_100m":                        ret.DtSlbHttpProxyStats.Stats.Req_100m,
			"req_200m":                        ret.DtSlbHttpProxyStats.Stats.Req_200m,
			"req_500m":                        ret.DtSlbHttpProxyStats.Stats.Req_500m,
			"req_1s":                          ret.DtSlbHttpProxyStats.Stats.Req_1s,
			"req_2s":                          ret.DtSlbHttpProxyStats.Stats.Req_2s,
			"req_5s":                          ret.DtSlbHttpProxyStats.Stats.Req_5s,
			"req_over_5s":                     ret.DtSlbHttpProxyStats.Stats.Req_over_5s,
			"req_track":                       ret.DtSlbHttpProxyStats.Stats.Req_track,
			"connect_req":                     ret.DtSlbHttpProxyStats.Stats.Connect_req,
			"req_enter_ssli":                  ret.DtSlbHttpProxyStats.Stats.Req_enter_ssli,
			"decompression_before":            ret.DtSlbHttpProxyStats.Stats.Decompression_before,
			"decompression_after":             ret.DtSlbHttpProxyStats.Stats.Decompression_after,
			"req_http2":                       ret.DtSlbHttpProxyStats.Stats.Req_http2,
			"response_http2":                  ret.DtSlbHttpProxyStats.Stats.Response_http2,
			"doh_req":                         ret.DtSlbHttpProxyStats.Stats.Doh_req,
			"doh_req_get":                     ret.DtSlbHttpProxyStats.Stats.Doh_req_get,
			"doh_req_post":                    ret.DtSlbHttpProxyStats.Stats.Doh_req_post,
			"doh_non_doh_req":                 ret.DtSlbHttpProxyStats.Stats.Doh_non_doh_req,
			"doh_non_doh_req_get":             ret.DtSlbHttpProxyStats.Stats.Doh_non_doh_req_get,
			"doh_non_doh_req_post":            ret.DtSlbHttpProxyStats.Stats.Doh_non_doh_req_post,
			"doh_resp":                        ret.DtSlbHttpProxyStats.Stats.Doh_resp,
			"doh_tc_resp":                     ret.DtSlbHttpProxyStats.Stats.Doh_tc_resp,
			"doh_udp_dns_req":                 ret.DtSlbHttpProxyStats.Stats.Doh_udp_dns_req,
			"doh_udp_dns_resp":                ret.DtSlbHttpProxyStats.Stats.Doh_udp_dns_resp,
			"doh_tcp_dns_req":                 ret.DtSlbHttpProxyStats.Stats.Doh_tcp_dns_req,
			"doh_tcp_dns_resp":                ret.DtSlbHttpProxyStats.Stats.Doh_tcp_dns_resp,
			"doh_req_send_failed":             ret.DtSlbHttpProxyStats.Stats.Doh_req_send_failed,
			"doh_resp_send_failed":            ret.DtSlbHttpProxyStats.Stats.Doh_resp_send_failed,
			"doh_malloc_fail":                 ret.DtSlbHttpProxyStats.Stats.Doh_malloc_fail,
			"doh_req_udp_retry":               ret.DtSlbHttpProxyStats.Stats.Doh_req_udp_retry,
			"doh_req_udp_retry_fail":          ret.DtSlbHttpProxyStats.Stats.Doh_req_udp_retry_fail,
			"doh_req_tcp_retry":               ret.DtSlbHttpProxyStats.Stats.Doh_req_tcp_retry,
			"doh_req_tcp_retry_fail":          ret.DtSlbHttpProxyStats.Stats.Doh_req_tcp_retry_fail,
			"doh_snat_failed":                 ret.DtSlbHttpProxyStats.Stats.Doh_snat_failed,
			"doh_path_not_found":              ret.DtSlbHttpProxyStats.Stats.Doh_path_not_found,
			"doh_get_dns_arg_failed":          ret.DtSlbHttpProxyStats.Stats.Doh_get_dns_arg_failed,
			"doh_get_base64_decode_failed":    ret.DtSlbHttpProxyStats.Stats.Doh_get_base64_decode_failed,
			"doh_post_content_type_mismatch":  ret.DtSlbHttpProxyStats.Stats.Doh_post_content_type_mismatch,
			"doh_post_payload_not_found":      ret.DtSlbHttpProxyStats.Stats.Doh_post_payload_not_found,
			"doh_post_payload_extract_failed": ret.DtSlbHttpProxyStats.Stats.Doh_post_payload_extract_failed,
			"doh_non_doh_method":              ret.DtSlbHttpProxyStats.Stats.Doh_non_doh_method,
			"doh_tcp_send_failed":             ret.DtSlbHttpProxyStats.Stats.Doh_tcp_send_failed,
			"doh_udp_send_failed":             ret.DtSlbHttpProxyStats.Stats.Doh_udp_send_failed,
			"doh_query_time_out":              ret.DtSlbHttpProxyStats.Stats.Doh_query_time_out,
			"doh_dns_query_type_a":            ret.DtSlbHttpProxyStats.Stats.Doh_dns_query_type_a,
			"doh_dns_query_type_aaaa":         ret.DtSlbHttpProxyStats.Stats.Doh_dns_query_type_aaaa,
			"doh_dns_query_type_ns":           ret.DtSlbHttpProxyStats.Stats.Doh_dns_query_type_ns,
			"doh_dns_query_type_cname":        ret.DtSlbHttpProxyStats.Stats.Doh_dns_query_type_cname,
			"doh_dns_query_type_any":          ret.DtSlbHttpProxyStats.Stats.Doh_dns_query_type_any,
			"doh_dns_query_type_srv":          ret.DtSlbHttpProxyStats.Stats.Doh_dns_query_type_srv,
			"doh_dns_query_type_mx":           ret.DtSlbHttpProxyStats.Stats.Doh_dns_query_type_mx,
			"doh_dns_query_type_soa":          ret.DtSlbHttpProxyStats.Stats.Doh_dns_query_type_soa,
			"doh_dns_query_type_others":       ret.DtSlbHttpProxyStats.Stats.Doh_dns_query_type_others,
			"doh_resp_setup_failed":           ret.DtSlbHttpProxyStats.Stats.Doh_resp_setup_failed,
			"doh_resp_header_alloc_failed":    ret.DtSlbHttpProxyStats.Stats.Doh_resp_header_alloc_failed,
			"doh_resp_que_failed":             ret.DtSlbHttpProxyStats.Stats.Doh_resp_que_failed,
			"doh_resp_udp_frags":              ret.DtSlbHttpProxyStats.Stats.Doh_resp_udp_frags,
			"doh_resp_tcp_frags":              ret.DtSlbHttpProxyStats.Stats.Doh_resp_tcp_frags,
			"doh_serv_sel_failed":             ret.DtSlbHttpProxyStats.Stats.Doh_serv_sel_failed,
			"doh_retry_w_tcp":                 ret.DtSlbHttpProxyStats.Stats.Doh_retry_w_tcp,
			"doh_get_uri_too_long":            ret.DtSlbHttpProxyStats.Stats.Doh_get_uri_too_long,
			"doh_post_payload_too_large":      ret.DtSlbHttpProxyStats.Stats.Doh_post_payload_too_large,
			"doh_dns_malformed_query":         ret.DtSlbHttpProxyStats.Stats.Doh_dns_malformed_query,
			"doh_dns_resp_rcode_err_format":   ret.DtSlbHttpProxyStats.Stats.Doh_dns_resp_rcode_err_format,
			"doh_dns_resp_rcode_err_server":   ret.DtSlbHttpProxyStats.Stats.Doh_dns_resp_rcode_err_server,
			"doh_dns_resp_rcode_err_name":     ret.DtSlbHttpProxyStats.Stats.Doh_dns_resp_rcode_err_name,
			"doh_dns_resp_rcode_err_type":     ret.DtSlbHttpProxyStats.Stats.Doh_dns_resp_rcode_err_type,
			"doh_dns_resp_rcode_refuse":       ret.DtSlbHttpProxyStats.Stats.Doh_dns_resp_rcode_refuse,
			"doh_dns_resp_rcode_yxdomain":     ret.DtSlbHttpProxyStats.Stats.Doh_dns_resp_rcode_yxdomain,
			"doh_dns_resp_rcode_yxrrset":      ret.DtSlbHttpProxyStats.Stats.Doh_dns_resp_rcode_yxrrset,
			"doh_dns_resp_rcode_nxrrset":      ret.DtSlbHttpProxyStats.Stats.Doh_dns_resp_rcode_nxrrset,
			"doh_dns_resp_rcode_notauth":      ret.DtSlbHttpProxyStats.Stats.Doh_dns_resp_rcode_notauth,
			"doh_dns_resp_rcode_notzone":      ret.DtSlbHttpProxyStats.Stats.Doh_dns_resp_rcode_notzone,
			"doh_dns_resp_rcode_other":        ret.DtSlbHttpProxyStats.Stats.Doh_dns_resp_rcode_other,
			"compression_before_br":           ret.DtSlbHttpProxyStats.Stats.Compression_before_br,
			"compression_after_br":            ret.DtSlbHttpProxyStats.Stats.Compression_after_br,
			"compression_before_total":        ret.DtSlbHttpProxyStats.Stats.Compression_before_total,
			"compression_after_total":         ret.DtSlbHttpProxyStats.Stats.Compression_after_total,
			"decompression_before_br":         ret.DtSlbHttpProxyStats.Stats.Decompression_before_br,
			"decompression_after_br":          ret.DtSlbHttpProxyStats.Stats.Decompression_after_br,
			"decompression_before_total":      ret.DtSlbHttpProxyStats.Stats.Decompression_before_total,
			"decompression_after_total":       ret.DtSlbHttpProxyStats.Stats.Decompression_after_total,
			"req_http3":                       ret.DtSlbHttpProxyStats.Stats.Req_http3,
			"response_http3":                  ret.DtSlbHttpProxyStats.Stats.Response_http3,
		},
	}
}

func getObjectSlbHttpProxyStatsStats(d []interface{}) edpt.SlbHttpProxyStatsStats {

	count1 := len(d)
	var ret edpt.SlbHttpProxyStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_proxy = in["curr_proxy"].(int)
		ret.Total_proxy = in["total_proxy"].(int)
		ret.Req = in["req"].(int)
		ret.Req_succ = in["req_succ"].(int)
		ret.Noproxy = in["noproxy"].(int)
		ret.Client_rst = in["client_rst"].(int)
		ret.Server_rst = in["server_rst"].(int)
		ret.Notuple = in["notuple"].(int)
		ret.Parsereq_fail = in["parsereq_fail"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.Fwdreq_fail = in["fwdreq_fail"].(int)
		ret.Fwdreqdata_fail = in["fwdreqdata_fail"].(int)
		ret.Req_retran = in["req_retran"].(int)
		ret.Req_ofo = in["req_ofo"].(int)
		ret.Server_resel = in["server_resel"].(int)
		ret.Svr_prem_close = in["svr_prem_close"].(int)
		ret.New_svrconn = in["new_svrconn"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Req_over_limit = in["req_over_limit"].(int)
		ret.Req_rate_over_limit = in["req_rate_over_limit"].(int)
		ret.Compression_before = in["compression_before"].(int)
		ret.Compression_after = in["compression_after"].(int)
		ret.Response_1xx = in["response_1xx"].(int)
		ret.Response_100 = in["response_100"].(int)
		ret.Response_101 = in["response_101"].(int)
		ret.Response_102 = in["response_102"].(int)
		ret.Response_2xx = in["response_2xx"].(int)
		ret.Response_200 = in["response_200"].(int)
		ret.Response_201 = in["response_201"].(int)
		ret.Response_202 = in["response_202"].(int)
		ret.Response_203 = in["response_203"].(int)
		ret.Response_204 = in["response_204"].(int)
		ret.Response_205 = in["response_205"].(int)
		ret.Response_206 = in["response_206"].(int)
		ret.Response_207 = in["response_207"].(int)
		ret.Response_3xx = in["response_3xx"].(int)
		ret.Response_300 = in["response_300"].(int)
		ret.Response_301 = in["response_301"].(int)
		ret.Response_302 = in["response_302"].(int)
		ret.Response_303 = in["response_303"].(int)
		ret.Response_304 = in["response_304"].(int)
		ret.Response_305 = in["response_305"].(int)
		ret.Response_306 = in["response_306"].(int)
		ret.Response_307 = in["response_307"].(int)
		ret.Response_4xx = in["response_4xx"].(int)
		ret.Response_400 = in["response_400"].(int)
		ret.Response_401 = in["response_401"].(int)
		ret.Response_402 = in["response_402"].(int)
		ret.Response_403 = in["response_403"].(int)
		ret.Response_404 = in["response_404"].(int)
		ret.Response_405 = in["response_405"].(int)
		ret.Response_406 = in["response_406"].(int)
		ret.Response_407 = in["response_407"].(int)
		ret.Response_408 = in["response_408"].(int)
		ret.Response_409 = in["response_409"].(int)
		ret.Response_410 = in["response_410"].(int)
		ret.Response_411 = in["response_411"].(int)
		ret.Response_412 = in["response_412"].(int)
		ret.Response_413 = in["response_413"].(int)
		ret.Response_414 = in["response_414"].(int)
		ret.Response_415 = in["response_415"].(int)
		ret.Response_416 = in["response_416"].(int)
		ret.Response_417 = in["response_417"].(int)
		ret.Response_418 = in["response_418"].(int)
		ret.Response_422 = in["response_422"].(int)
		ret.Response_423 = in["response_423"].(int)
		ret.Response_424 = in["response_424"].(int)
		ret.Response_425 = in["response_425"].(int)
		ret.Response_426 = in["response_426"].(int)
		ret.Response_449 = in["response_449"].(int)
		ret.Response_450 = in["response_450"].(int)
		ret.Response_5xx = in["response_5xx"].(int)
		ret.Response_500 = in["response_500"].(int)
		ret.Response_501 = in["response_501"].(int)
		ret.Response_502 = in["response_502"].(int)
		ret.Response_503 = in["response_503"].(int)
		ret.Response_504 = in["response_504"].(int)
		ret.Response_505 = in["response_505"].(int)
		ret.Response_506 = in["response_506"].(int)
		ret.Response_507 = in["response_507"].(int)
		ret.Response_508 = in["response_508"].(int)
		ret.Response_509 = in["response_509"].(int)
		ret.Response_510 = in["response_510"].(int)
		ret.Response_6xx = in["response_6xx"].(int)
		ret.Response_unknown = in["response_unknown"].(int)
		ret.Req_get = in["req_get"].(int)
		ret.Req_head = in["req_head"].(int)
		ret.Req_put = in["req_put"].(int)
		ret.Req_post = in["req_post"].(int)
		ret.Req_trace = in["req_trace"].(int)
		ret.Req_options = in["req_options"].(int)
		ret.Req_connect = in["req_connect"].(int)
		ret.Req_delete = in["req_delete"].(int)
		ret.Req_unknown = in["req_unknown"].(int)
		ret.Req_content_len = in["req_content_len"].(int)
		ret.Rsp_content_len = in["rsp_content_len"].(int)
		ret.Rsp_chunk = in["rsp_chunk"].(int)
		ret.Cache_rsp = in["cache_rsp"].(int)
		ret.Close_on_ddos = in["close_on_ddos"].(int)
		ret.Req_sz_1k = in["req_sz_1k"].(int)
		ret.Req_sz_2k = in["req_sz_2k"].(int)
		ret.Req_sz_4k = in["req_sz_4k"].(int)
		ret.Req_sz_8k = in["req_sz_8k"].(int)
		ret.Req_sz_16k = in["req_sz_16k"].(int)
		ret.Req_sz_32k = in["req_sz_32k"].(int)
		ret.Req_sz_64k = in["req_sz_64k"].(int)
		ret.Req_sz_256k = in["req_sz_256k"].(int)
		ret.Req_sz_gt_256k = in["req_sz_gt_256k"].(int)
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
		ret.Req_track = in["req_track"].(int)
		ret.Connect_req = in["connect_req"].(int)
		ret.Req_enter_ssli = in["req_enter_ssli"].(int)
		ret.Decompression_before = in["decompression_before"].(int)
		ret.Decompression_after = in["decompression_after"].(int)
		ret.Req_http2 = in["req_http2"].(int)
		ret.Response_http2 = in["response_http2"].(int)
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
		ret.Compression_before_br = in["compression_before_br"].(int)
		ret.Compression_after_br = in["compression_after_br"].(int)
		ret.Compression_before_total = in["compression_before_total"].(int)
		ret.Compression_after_total = in["compression_after_total"].(int)
		ret.Decompression_before_br = in["decompression_before_br"].(int)
		ret.Decompression_after_br = in["decompression_after_br"].(int)
		ret.Decompression_before_total = in["decompression_before_total"].(int)
		ret.Decompression_after_total = in["decompression_after_total"].(int)
		ret.Req_http3 = in["req_http3"].(int)
		ret.Response_http3 = in["response_http3"].(int)
	}
	return ret
}

func dataToEndpointSlbHttpProxyStats(d *schema.ResourceData) edpt.SlbHttpProxyStats {
	var ret edpt.SlbHttpProxyStats

	ret.Stats = getObjectSlbHttpProxyStatsStats(d.Get("stats").([]interface{}))
	return ret
}
