package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosL7HttpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_l7_http_stats`: Statistics for the object l7-http\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosL7HttpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"req_processed": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Processed",
						},
						"req_ofo": {
							Type: schema.TypeInt, Optional: true, Description: "Out-Of-Order Request",
						},
						"ofo_timer_expired": {
							Type: schema.TypeInt, Optional: true, Description: "Out-Of-Order Timeout",
						},
						"ofo_queue_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Out-Of-Order Queue Exceeded",
						},
						"ofo": {
							Type: schema.TypeInt, Optional: true, Description: "Out-Of-Order Packets",
						},
						"partial_hdr": {
							Type: schema.TypeInt, Optional: true, Description: "Partial Header",
						},
						"http_idle_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Http Idle Timeout",
						},
						"new_syn": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN",
						},
						"retrans": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit",
						},
						"retrans_fin": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit FIN",
						},
						"retrans_push": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit PSH",
						},
						"retrans_rst": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit RST",
						},
						"req_retrans": {
							Type: schema.TypeInt, Optional: true, Description: "Retransmit Request",
						},
						"request": {
							Type: schema.TypeInt, Optional: true, Description: "Request Total",
						},
						"req_content_len": {
							Type: schema.TypeInt, Optional: true, Description: "Request Content-Length Received",
						},
						"src_req_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate Exceeded",
						},
						"dst_req_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate Exceeded",
						},
						"lower_than_mss_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Min Payload Size Fail Exceeded",
						},
						"parsereq_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Parse Request Failed",
						},
						"neg_req_remain": {
							Type: schema.TypeInt, Optional: true, Description: "Negative Request Remain",
						},
						"neg_rsp_remain": {
							Type: schema.TypeInt, Optional: true, Description: "Negative Response Remain",
						},
						"invalid_header": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP Header Invalid",
						},
						"too_many_headers": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP Header Too Many",
						},
						"header_name_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP Header Name Too Long",
						},
						"invalid_hdr_name": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP Header Name Invalid",
						},
						"invalid_hdr_val": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP Header Value Invalid",
						},
						"line_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "Line Too Long",
						},
						"client_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Client TCP RST Received",
						},
						"hps_server_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Server TCP RST Received",
						},
						"ddos_policy_violation": {
							Type: schema.TypeInt, Optional: true, Description: "Policy Violation",
						},
						"policy_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Policy Dropped",
						},
						"error_condition": {
							Type: schema.TypeInt, Optional: true, Description: "Error Condition",
						},
						"http11": {
							Type: schema.TypeInt, Optional: true, Description: "Request HTTP 1.1",
						},
						"http10": {
							Type: schema.TypeInt, Optional: true, Description: "Request HTTP 1.0",
						},
						"rsp_chunk": {
							Type: schema.TypeInt, Optional: true, Description: "Response Chunk",
						},
						"http_get": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method GET",
						},
						"http_head": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method HEAD",
						},
						"http_put": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method PUT",
						},
						"http_post": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method POST",
						},
						"http_trace": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method TRACE",
						},
						"http_options": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method OPTIONS",
						},
						"http_connect": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method CONNECT",
						},
						"http_del": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method DELETE",
						},
						"http_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method UNKNOWN",
						},
						"hps_req_sz_1k": {
							Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Less Than or Equal to 1K",
						},
						"hps_req_sz_2k": {
							Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Less Than or Equal to 2K",
						},
						"hps_req_sz_4k": {
							Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Less Than or Equal to 4K",
						},
						"hps_req_sz_8k": {
							Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Less Than or Equal to 8K",
						},
						"hps_req_sz_16k": {
							Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Less Than or Equal to 16K",
						},
						"hps_req_sz_32k": {
							Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Less Than or Equal to 32K",
						},
						"hps_req_sz_64k": {
							Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Less Than or Equal to 64K",
						},
						"hps_req_sz_256k": {
							Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Less Than or Equal to 256K",
						},
						"hps_req_sz_256k_plus": {
							Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Larger Than 256K",
						},
						"hps_rsp_11": {
							Type: schema.TypeInt, Optional: true, Description: "Response HTTP 1.1",
						},
						"hps_rsp_10": {
							Type: schema.TypeInt, Optional: true, Description: "Response HTTP 1.0",
						},
						"hps_rsp_sz_1k": {
							Type: schema.TypeInt, Optional: true, Description: "Response Payload Size Less Than or Equal to 1K",
						},
						"hps_rsp_sz_2k": {
							Type: schema.TypeInt, Optional: true, Description: "Response Payload Size Less Than or Equal to 2K",
						},
						"hps_rsp_sz_4k": {
							Type: schema.TypeInt, Optional: true, Description: "Response Payload Size Less Than or Equal to 4K",
						},
						"hps_rsp_sz_8k": {
							Type: schema.TypeInt, Optional: true, Description: "Response Payload Size Less Than or Equal to 8K",
						},
						"hps_rsp_sz_16k": {
							Type: schema.TypeInt, Optional: true, Description: "Response Payload Size Less Than or Equal to 16K",
						},
						"hps_rsp_sz_32k": {
							Type: schema.TypeInt, Optional: true, Description: "Response Payload Size Less Than or Equal to 32K",
						},
						"hps_rsp_sz_64k": {
							Type: schema.TypeInt, Optional: true, Description: "Response Payload Size Less Than or Equal to 64K",
						},
						"hps_rsp_sz_256k": {
							Type: schema.TypeInt, Optional: true, Description: "Response Payload Size Less Than or Equal to 256K",
						},
						"hps_rsp_sz_256k_plus": {
							Type: schema.TypeInt, Optional: true, Description: "Response Payload Size Larger Than 256K",
						},
						"hps_rsp_status_1xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status Code 1XX",
						},
						"hps_rsp_status_2xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status Code 2XX",
						},
						"hps_rsp_status_3xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status Code 3XX",
						},
						"hps_rsp_status_4xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status Code 4XX",
						},
						"hps_rsp_status_5xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status Code 5XX",
						},
						"hps_rsp_status_504_ax": {
							Type: schema.TypeInt, Optional: true, Description: "Status Code 504 AX-Gen",
						},
						"hps_rsp_status_6xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status Code 6XX",
						},
						"hps_rsp_status_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "Status Code Unknown",
						},
						"chunk_sz_512": {
							Type: schema.TypeInt, Optional: true, Description: "Payload Chunk Size Less Than or Equal to 512",
						},
						"chunk_sz_1k": {
							Type: schema.TypeInt, Optional: true, Description: "Payload Chunk Size Less Than or Equal to 1K",
						},
						"chunk_sz_2k": {
							Type: schema.TypeInt, Optional: true, Description: "Payload Chunk Size Less Than or Equal to 2K",
						},
						"chunk_sz_4k": {
							Type: schema.TypeInt, Optional: true, Description: "Payload Chunk Size Less Than or Equal to 4K",
						},
						"chunk_sz_gt_4k": {
							Type: schema.TypeInt, Optional: true, Description: "Payload Chunk Size Larger Than 4K",
						},
						"chunk_bad": {
							Type: schema.TypeInt, Optional: true, Description: "Bad HTTP Chunk",
						},
						"challenge_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Challenge Failed",
						},
						"challenge_ud_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Challenge URL Redirect Sent",
						},
						"challenge_ud_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Challenge URL Redirect Failed",
						},
						"challenge_js_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Challenge Javascript Sent",
						},
						"challenge_js_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Challenge Javascript Failed",
						},
						"malform_bad_chunk": {
							Type: schema.TypeInt, Optional: true, Description: "Malform Bad Chunk",
						},
						"malform_content_len_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "Malform Content Length Too Long",
						},
						"malform_header_name_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "Malform Header Name Too Long",
						},
						"malform_line_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "Malform Line Too Long",
						},
						"malform_req_line_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "Malform Request Line Too Long",
						},
						"malform_too_many_headers": {
							Type: schema.TypeInt, Optional: true, Description: "Malform Too Many Headers",
						},
						"window_small": {
							Type: schema.TypeInt, Optional: true, Description: "Window Size Small",
						},
						"window_small_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Window Size Small Dropped",
						},
						"alloc_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Alloc Failed",
						},
						"use_hdr_ip_as_source": {
							Type: schema.TypeInt, Optional: true, Description: "Use IP In Header As Src",
						},
						"agent_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Agent Filter Match",
						},
						"agent_filter_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Agent Filter Blacklisted",
						},
						"referer_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Referer Filter Match",
						},
						"referer_filter_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Referer Filter Blacklisted",
						},
						"dst_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Match",
						},
						"dst_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter No Match",
						},
						"dst_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action Blacklist",
						},
						"dst_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action Drop",
						},
						"dst_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action Default Pass",
						},
						"dst_post_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Post Rate Exceeded",
						},
						"src_post_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Src Post Rate Exceeded",
						},
						"dst_resp_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Response Rate Exceeded",
						},
						"dst_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action WL",
						},
						"src_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Match",
						},
						"src_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter No Match",
						},
						"src_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Blacklist",
						},
						"src_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Drop",
						},
						"src_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Default Pass",
						},
						"src_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action WL",
						},
						"src_dst_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Match",
						},
						"src_dst_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter No Match",
						},
						"src_dst_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Action Blacklist",
						},
						"src_dst_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Action Drop",
						},
						"src_dst_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Action Default Pass",
						},
						"src_dst_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Action WL",
						},
						"dst_filter_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Rate Exceed",
						},
						"dst_filter_action_ignore": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action Ignore",
						},
						"dst_filter_action_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action Reset",
						},
						"uri_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "URI Filter Match",
						},
						"http_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP Auth Dropped",
						},
						"http_auth_resp": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP Auth Responded",
						},
						"header_processing_time_0": {
							Type: schema.TypeInt, Optional: true, Description: "Header Process Time Less Than 1s",
						},
						"header_processing_time_1": {
							Type: schema.TypeInt, Optional: true, Description: "Header Process Time Less Than 10s",
						},
						"header_processing_time_2": {
							Type: schema.TypeInt, Optional: true, Description: "Header Process Time Less Than 30s",
						},
						"header_processing_time_3": {
							Type: schema.TypeInt, Optional: true, Description: "Header Process Time Larger or Equal to 30s",
						},
						"header_processing_incomplete": {
							Type: schema.TypeInt, Optional: true, Description: "Header Process Incomplete",
						},
						"malform_req_line_too_small": {
							Type: schema.TypeInt, Optional: true, Description: "Malform Request Line Too Small",
						},
						"malform_req_line_invalid_method": {
							Type: schema.TypeInt, Optional: true, Description: "Malform Request Line Invalid Method",
						},
					},
				},
			},
		},
	}
}

func resourceDdosL7HttpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosL7HttpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosL7HttpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosL7HttpStatsStats := setObjectDdosL7HttpStatsStats(res)
		d.Set("stats", DdosL7HttpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosL7HttpStatsStats(ret edpt.DataDdosL7HttpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"req_processed":                      ret.DtDdosL7HttpStats.Stats.Req_processed,
			"req_ofo":                            ret.DtDdosL7HttpStats.Stats.Req_ofo,
			"ofo_timer_expired":                  ret.DtDdosL7HttpStats.Stats.Ofo_timer_expired,
			"ofo_queue_exceed":                   ret.DtDdosL7HttpStats.Stats.Ofo_queue_exceed,
			"ofo":                                ret.DtDdosL7HttpStats.Stats.Ofo,
			"partial_hdr":                        ret.DtDdosL7HttpStats.Stats.Partial_hdr,
			"http_idle_timeout":                  ret.DtDdosL7HttpStats.Stats.Http_idle_timeout,
			"new_syn":                            ret.DtDdosL7HttpStats.Stats.New_syn,
			"retrans":                            ret.DtDdosL7HttpStats.Stats.Retrans,
			"retrans_fin":                        ret.DtDdosL7HttpStats.Stats.Retrans_fin,
			"retrans_push":                       ret.DtDdosL7HttpStats.Stats.Retrans_push,
			"retrans_rst":                        ret.DtDdosL7HttpStats.Stats.Retrans_rst,
			"req_retrans":                        ret.DtDdosL7HttpStats.Stats.Req_retrans,
			"request":                            ret.DtDdosL7HttpStats.Stats.Request,
			"req_content_len":                    ret.DtDdosL7HttpStats.Stats.Req_content_len,
			"src_req_rate_exceed":                ret.DtDdosL7HttpStats.Stats.Src_req_rate_exceed,
			"dst_req_rate_exceed":                ret.DtDdosL7HttpStats.Stats.Dst_req_rate_exceed,
			"lower_than_mss_exceed":              ret.DtDdosL7HttpStats.Stats.Lower_than_mss_exceed,
			"parsereq_fail":                      ret.DtDdosL7HttpStats.Stats.Parsereq_fail,
			"neg_req_remain":                     ret.DtDdosL7HttpStats.Stats.Neg_req_remain,
			"neg_rsp_remain":                     ret.DtDdosL7HttpStats.Stats.Neg_rsp_remain,
			"invalid_header":                     ret.DtDdosL7HttpStats.Stats.Invalid_header,
			"too_many_headers":                   ret.DtDdosL7HttpStats.Stats.Too_many_headers,
			"header_name_too_long":               ret.DtDdosL7HttpStats.Stats.Header_name_too_long,
			"invalid_hdr_name":                   ret.DtDdosL7HttpStats.Stats.Invalid_hdr_name,
			"invalid_hdr_val":                    ret.DtDdosL7HttpStats.Stats.Invalid_hdr_val,
			"line_too_long":                      ret.DtDdosL7HttpStats.Stats.Line_too_long,
			"client_rst":                         ret.DtDdosL7HttpStats.Stats.Client_rst,
			"hps_server_rst":                     ret.DtDdosL7HttpStats.Stats.Hps_server_rst,
			"ddos_policy_violation":              ret.DtDdosL7HttpStats.Stats.Ddos_policy_violation,
			"policy_drop":                        ret.DtDdosL7HttpStats.Stats.Policy_drop,
			"error_condition":                    ret.DtDdosL7HttpStats.Stats.Error_condition,
			"http11":                             ret.DtDdosL7HttpStats.Stats.Http11,
			"http10":                             ret.DtDdosL7HttpStats.Stats.Http10,
			"rsp_chunk":                          ret.DtDdosL7HttpStats.Stats.Rsp_chunk,
			"http_get":                           ret.DtDdosL7HttpStats.Stats.Http_get,
			"http_head":                          ret.DtDdosL7HttpStats.Stats.Http_head,
			"http_put":                           ret.DtDdosL7HttpStats.Stats.Http_put,
			"http_post":                          ret.DtDdosL7HttpStats.Stats.Http_post,
			"http_trace":                         ret.DtDdosL7HttpStats.Stats.Http_trace,
			"http_options":                       ret.DtDdosL7HttpStats.Stats.Http_options,
			"http_connect":                       ret.DtDdosL7HttpStats.Stats.Http_connect,
			"http_del":                           ret.DtDdosL7HttpStats.Stats.Http_del,
			"http_unknown":                       ret.DtDdosL7HttpStats.Stats.Http_unknown,
			"hps_req_sz_1k":                      ret.DtDdosL7HttpStats.Stats.Hps_req_sz_1k,
			"hps_req_sz_2k":                      ret.DtDdosL7HttpStats.Stats.Hps_req_sz_2k,
			"hps_req_sz_4k":                      ret.DtDdosL7HttpStats.Stats.Hps_req_sz_4k,
			"hps_req_sz_8k":                      ret.DtDdosL7HttpStats.Stats.Hps_req_sz_8k,
			"hps_req_sz_16k":                     ret.DtDdosL7HttpStats.Stats.Hps_req_sz_16k,
			"hps_req_sz_32k":                     ret.DtDdosL7HttpStats.Stats.Hps_req_sz_32k,
			"hps_req_sz_64k":                     ret.DtDdosL7HttpStats.Stats.Hps_req_sz_64k,
			"hps_req_sz_256k":                    ret.DtDdosL7HttpStats.Stats.Hps_req_sz_256k,
			"hps_req_sz_256k_plus":               ret.DtDdosL7HttpStats.Stats.Hps_req_sz_256k_plus,
			"hps_rsp_11":                         ret.DtDdosL7HttpStats.Stats.Hps_rsp_11,
			"hps_rsp_10":                         ret.DtDdosL7HttpStats.Stats.Hps_rsp_10,
			"hps_rsp_sz_1k":                      ret.DtDdosL7HttpStats.Stats.Hps_rsp_sz_1k,
			"hps_rsp_sz_2k":                      ret.DtDdosL7HttpStats.Stats.Hps_rsp_sz_2k,
			"hps_rsp_sz_4k":                      ret.DtDdosL7HttpStats.Stats.Hps_rsp_sz_4k,
			"hps_rsp_sz_8k":                      ret.DtDdosL7HttpStats.Stats.Hps_rsp_sz_8k,
			"hps_rsp_sz_16k":                     ret.DtDdosL7HttpStats.Stats.Hps_rsp_sz_16k,
			"hps_rsp_sz_32k":                     ret.DtDdosL7HttpStats.Stats.Hps_rsp_sz_32k,
			"hps_rsp_sz_64k":                     ret.DtDdosL7HttpStats.Stats.Hps_rsp_sz_64k,
			"hps_rsp_sz_256k":                    ret.DtDdosL7HttpStats.Stats.Hps_rsp_sz_256k,
			"hps_rsp_sz_256k_plus":               ret.DtDdosL7HttpStats.Stats.Hps_rsp_sz_256k_plus,
			"hps_rsp_status_1xx":                 ret.DtDdosL7HttpStats.Stats.Hps_rsp_status_1xx,
			"hps_rsp_status_2xx":                 ret.DtDdosL7HttpStats.Stats.Hps_rsp_status_2xx,
			"hps_rsp_status_3xx":                 ret.DtDdosL7HttpStats.Stats.Hps_rsp_status_3xx,
			"hps_rsp_status_4xx":                 ret.DtDdosL7HttpStats.Stats.Hps_rsp_status_4xx,
			"hps_rsp_status_5xx":                 ret.DtDdosL7HttpStats.Stats.Hps_rsp_status_5xx,
			"hps_rsp_status_504_ax":              ret.DtDdosL7HttpStats.Stats.Hps_rsp_status_504_ax,
			"hps_rsp_status_6xx":                 ret.DtDdosL7HttpStats.Stats.Hps_rsp_status_6xx,
			"hps_rsp_status_unknown":             ret.DtDdosL7HttpStats.Stats.Hps_rsp_status_unknown,
			"chunk_sz_512":                       ret.DtDdosL7HttpStats.Stats.Chunk_sz_512,
			"chunk_sz_1k":                        ret.DtDdosL7HttpStats.Stats.Chunk_sz_1k,
			"chunk_sz_2k":                        ret.DtDdosL7HttpStats.Stats.Chunk_sz_2k,
			"chunk_sz_4k":                        ret.DtDdosL7HttpStats.Stats.Chunk_sz_4k,
			"chunk_sz_gt_4k":                     ret.DtDdosL7HttpStats.Stats.Chunk_sz_gt_4k,
			"chunk_bad":                          ret.DtDdosL7HttpStats.Stats.Chunk_bad,
			"challenge_fail":                     ret.DtDdosL7HttpStats.Stats.Challenge_fail,
			"challenge_ud_sent":                  ret.DtDdosL7HttpStats.Stats.Challenge_ud_sent,
			"challenge_ud_fail":                  ret.DtDdosL7HttpStats.Stats.Challenge_ud_fail,
			"challenge_js_sent":                  ret.DtDdosL7HttpStats.Stats.Challenge_js_sent,
			"challenge_js_fail":                  ret.DtDdosL7HttpStats.Stats.Challenge_js_fail,
			"malform_bad_chunk":                  ret.DtDdosL7HttpStats.Stats.Malform_bad_chunk,
			"malform_content_len_too_long":       ret.DtDdosL7HttpStats.Stats.Malform_content_len_too_long,
			"malform_header_name_too_long":       ret.DtDdosL7HttpStats.Stats.Malform_header_name_too_long,
			"malform_line_too_long":              ret.DtDdosL7HttpStats.Stats.Malform_line_too_long,
			"malform_req_line_too_long":          ret.DtDdosL7HttpStats.Stats.Malform_req_line_too_long,
			"malform_too_many_headers":           ret.DtDdosL7HttpStats.Stats.Malform_too_many_headers,
			"window_small":                       ret.DtDdosL7HttpStats.Stats.Window_small,
			"window_small_drop":                  ret.DtDdosL7HttpStats.Stats.Window_small_drop,
			"alloc_fail":                         ret.DtDdosL7HttpStats.Stats.Alloc_fail,
			"use_hdr_ip_as_source":               ret.DtDdosL7HttpStats.Stats.Use_hdr_ip_as_source,
			"agent_filter_match":                 ret.DtDdosL7HttpStats.Stats.Agent_filter_match,
			"agent_filter_blacklist":             ret.DtDdosL7HttpStats.Stats.Agent_filter_blacklist,
			"referer_filter_match":               ret.DtDdosL7HttpStats.Stats.Referer_filter_match,
			"referer_filter_blacklist":           ret.DtDdosL7HttpStats.Stats.Referer_filter_blacklist,
			"dst_filter_match":                   ret.DtDdosL7HttpStats.Stats.Dst_filter_match,
			"dst_filter_not_match":               ret.DtDdosL7HttpStats.Stats.Dst_filter_not_match,
			"dst_filter_action_blacklist":        ret.DtDdosL7HttpStats.Stats.Dst_filter_action_blacklist,
			"dst_filter_action_drop":             ret.DtDdosL7HttpStats.Stats.Dst_filter_action_drop,
			"dst_filter_action_default_pass":     ret.DtDdosL7HttpStats.Stats.Dst_filter_action_default_pass,
			"dst_post_rate_exceed":               ret.DtDdosL7HttpStats.Stats.Dst_post_rate_exceed,
			"src_post_rate_exceed":               ret.DtDdosL7HttpStats.Stats.Src_post_rate_exceed,
			"dst_resp_rate_exceed":               ret.DtDdosL7HttpStats.Stats.Dst_resp_rate_exceed,
			"dst_filter_action_whitelist":        ret.DtDdosL7HttpStats.Stats.Dst_filter_action_whitelist,
			"src_filter_match":                   ret.DtDdosL7HttpStats.Stats.Src_filter_match,
			"src_filter_not_match":               ret.DtDdosL7HttpStats.Stats.Src_filter_not_match,
			"src_filter_action_blacklist":        ret.DtDdosL7HttpStats.Stats.Src_filter_action_blacklist,
			"src_filter_action_drop":             ret.DtDdosL7HttpStats.Stats.Src_filter_action_drop,
			"src_filter_action_default_pass":     ret.DtDdosL7HttpStats.Stats.Src_filter_action_default_pass,
			"src_filter_action_whitelist":        ret.DtDdosL7HttpStats.Stats.Src_filter_action_whitelist,
			"src_dst_filter_match":               ret.DtDdosL7HttpStats.Stats.Src_dst_filter_match,
			"src_dst_filter_not_match":           ret.DtDdosL7HttpStats.Stats.Src_dst_filter_not_match,
			"src_dst_filter_action_blacklist":    ret.DtDdosL7HttpStats.Stats.Src_dst_filter_action_blacklist,
			"src_dst_filter_action_drop":         ret.DtDdosL7HttpStats.Stats.Src_dst_filter_action_drop,
			"src_dst_filter_action_default_pass": ret.DtDdosL7HttpStats.Stats.Src_dst_filter_action_default_pass,
			"src_dst_filter_action_whitelist":    ret.DtDdosL7HttpStats.Stats.Src_dst_filter_action_whitelist,
			"dst_filter_rate_exceed":             ret.DtDdosL7HttpStats.Stats.Dst_filter_rate_exceed,
			"dst_filter_action_ignore":           ret.DtDdosL7HttpStats.Stats.Dst_filter_action_ignore,
			"dst_filter_action_reset":            ret.DtDdosL7HttpStats.Stats.Dst_filter_action_reset,
			"uri_filter_match":                   ret.DtDdosL7HttpStats.Stats.Uri_filter_match,
			"http_auth_drop":                     ret.DtDdosL7HttpStats.Stats.Http_auth_drop,
			"http_auth_resp":                     ret.DtDdosL7HttpStats.Stats.Http_auth_resp,
			"header_processing_time_0":           ret.DtDdosL7HttpStats.Stats.Header_processing_time_0,
			"header_processing_time_1":           ret.DtDdosL7HttpStats.Stats.Header_processing_time_1,
			"header_processing_time_2":           ret.DtDdosL7HttpStats.Stats.Header_processing_time_2,
			"header_processing_time_3":           ret.DtDdosL7HttpStats.Stats.Header_processing_time_3,
			"header_processing_incomplete":       ret.DtDdosL7HttpStats.Stats.Header_processing_incomplete,
			"malform_req_line_too_small":         ret.DtDdosL7HttpStats.Stats.Malform_req_line_too_small,
			"malform_req_line_invalid_method":    ret.DtDdosL7HttpStats.Stats.Malform_req_line_invalid_method,
		},
	}
}

func getObjectDdosL7HttpStatsStats(d []interface{}) edpt.DdosL7HttpStatsStats {

	count1 := len(d)
	var ret edpt.DdosL7HttpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Req_processed = in["req_processed"].(int)
		ret.Req_ofo = in["req_ofo"].(int)
		ret.Ofo_timer_expired = in["ofo_timer_expired"].(int)
		ret.Ofo_queue_exceed = in["ofo_queue_exceed"].(int)
		ret.Ofo = in["ofo"].(int)
		ret.Partial_hdr = in["partial_hdr"].(int)
		ret.Http_idle_timeout = in["http_idle_timeout"].(int)
		ret.New_syn = in["new_syn"].(int)
		ret.Retrans = in["retrans"].(int)
		ret.Retrans_fin = in["retrans_fin"].(int)
		ret.Retrans_push = in["retrans_push"].(int)
		ret.Retrans_rst = in["retrans_rst"].(int)
		ret.Req_retrans = in["req_retrans"].(int)
		ret.Request = in["request"].(int)
		ret.Req_content_len = in["req_content_len"].(int)
		ret.Src_req_rate_exceed = in["src_req_rate_exceed"].(int)
		ret.Dst_req_rate_exceed = in["dst_req_rate_exceed"].(int)
		ret.Lower_than_mss_exceed = in["lower_than_mss_exceed"].(int)
		ret.Parsereq_fail = in["parsereq_fail"].(int)
		ret.Neg_req_remain = in["neg_req_remain"].(int)
		ret.Neg_rsp_remain = in["neg_rsp_remain"].(int)
		ret.Invalid_header = in["invalid_header"].(int)
		ret.Too_many_headers = in["too_many_headers"].(int)
		ret.Header_name_too_long = in["header_name_too_long"].(int)
		ret.Invalid_hdr_name = in["invalid_hdr_name"].(int)
		ret.Invalid_hdr_val = in["invalid_hdr_val"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Client_rst = in["client_rst"].(int)
		ret.Hps_server_rst = in["hps_server_rst"].(int)
		ret.Ddos_policy_violation = in["ddos_policy_violation"].(int)
		ret.Policy_drop = in["policy_drop"].(int)
		ret.Error_condition = in["error_condition"].(int)
		ret.Http11 = in["http11"].(int)
		ret.Http10 = in["http10"].(int)
		ret.Rsp_chunk = in["rsp_chunk"].(int)
		ret.Http_get = in["http_get"].(int)
		ret.Http_head = in["http_head"].(int)
		ret.Http_put = in["http_put"].(int)
		ret.Http_post = in["http_post"].(int)
		ret.Http_trace = in["http_trace"].(int)
		ret.Http_options = in["http_options"].(int)
		ret.Http_connect = in["http_connect"].(int)
		ret.Http_del = in["http_del"].(int)
		ret.Http_unknown = in["http_unknown"].(int)
		ret.Hps_req_sz_1k = in["hps_req_sz_1k"].(int)
		ret.Hps_req_sz_2k = in["hps_req_sz_2k"].(int)
		ret.Hps_req_sz_4k = in["hps_req_sz_4k"].(int)
		ret.Hps_req_sz_8k = in["hps_req_sz_8k"].(int)
		ret.Hps_req_sz_16k = in["hps_req_sz_16k"].(int)
		ret.Hps_req_sz_32k = in["hps_req_sz_32k"].(int)
		ret.Hps_req_sz_64k = in["hps_req_sz_64k"].(int)
		ret.Hps_req_sz_256k = in["hps_req_sz_256k"].(int)
		ret.Hps_req_sz_256k_plus = in["hps_req_sz_256k_plus"].(int)
		ret.Hps_rsp_11 = in["hps_rsp_11"].(int)
		ret.Hps_rsp_10 = in["hps_rsp_10"].(int)
		ret.Hps_rsp_sz_1k = in["hps_rsp_sz_1k"].(int)
		ret.Hps_rsp_sz_2k = in["hps_rsp_sz_2k"].(int)
		ret.Hps_rsp_sz_4k = in["hps_rsp_sz_4k"].(int)
		ret.Hps_rsp_sz_8k = in["hps_rsp_sz_8k"].(int)
		ret.Hps_rsp_sz_16k = in["hps_rsp_sz_16k"].(int)
		ret.Hps_rsp_sz_32k = in["hps_rsp_sz_32k"].(int)
		ret.Hps_rsp_sz_64k = in["hps_rsp_sz_64k"].(int)
		ret.Hps_rsp_sz_256k = in["hps_rsp_sz_256k"].(int)
		ret.Hps_rsp_sz_256k_plus = in["hps_rsp_sz_256k_plus"].(int)
		ret.Hps_rsp_status_1xx = in["hps_rsp_status_1xx"].(int)
		ret.Hps_rsp_status_2xx = in["hps_rsp_status_2xx"].(int)
		ret.Hps_rsp_status_3xx = in["hps_rsp_status_3xx"].(int)
		ret.Hps_rsp_status_4xx = in["hps_rsp_status_4xx"].(int)
		ret.Hps_rsp_status_5xx = in["hps_rsp_status_5xx"].(int)
		ret.Hps_rsp_status_504_ax = in["hps_rsp_status_504_ax"].(int)
		ret.Hps_rsp_status_6xx = in["hps_rsp_status_6xx"].(int)
		ret.Hps_rsp_status_unknown = in["hps_rsp_status_unknown"].(int)
		ret.Chunk_sz_512 = in["chunk_sz_512"].(int)
		ret.Chunk_sz_1k = in["chunk_sz_1k"].(int)
		ret.Chunk_sz_2k = in["chunk_sz_2k"].(int)
		ret.Chunk_sz_4k = in["chunk_sz_4k"].(int)
		ret.Chunk_sz_gt_4k = in["chunk_sz_gt_4k"].(int)
		ret.Chunk_bad = in["chunk_bad"].(int)
		ret.Challenge_fail = in["challenge_fail"].(int)
		ret.Challenge_ud_sent = in["challenge_ud_sent"].(int)
		ret.Challenge_ud_fail = in["challenge_ud_fail"].(int)
		ret.Challenge_js_sent = in["challenge_js_sent"].(int)
		ret.Challenge_js_fail = in["challenge_js_fail"].(int)
		ret.Malform_bad_chunk = in["malform_bad_chunk"].(int)
		ret.Malform_content_len_too_long = in["malform_content_len_too_long"].(int)
		ret.Malform_header_name_too_long = in["malform_header_name_too_long"].(int)
		ret.Malform_line_too_long = in["malform_line_too_long"].(int)
		ret.Malform_req_line_too_long = in["malform_req_line_too_long"].(int)
		ret.Malform_too_many_headers = in["malform_too_many_headers"].(int)
		ret.Window_small = in["window_small"].(int)
		ret.Window_small_drop = in["window_small_drop"].(int)
		ret.Alloc_fail = in["alloc_fail"].(int)
		ret.Use_hdr_ip_as_source = in["use_hdr_ip_as_source"].(int)
		ret.Agent_filter_match = in["agent_filter_match"].(int)
		ret.Agent_filter_blacklist = in["agent_filter_blacklist"].(int)
		ret.Referer_filter_match = in["referer_filter_match"].(int)
		ret.Referer_filter_blacklist = in["referer_filter_blacklist"].(int)
		ret.Dst_filter_match = in["dst_filter_match"].(int)
		ret.Dst_filter_not_match = in["dst_filter_not_match"].(int)
		ret.Dst_filter_action_blacklist = in["dst_filter_action_blacklist"].(int)
		ret.Dst_filter_action_drop = in["dst_filter_action_drop"].(int)
		ret.Dst_filter_action_default_pass = in["dst_filter_action_default_pass"].(int)
		ret.Dst_post_rate_exceed = in["dst_post_rate_exceed"].(int)
		ret.Src_post_rate_exceed = in["src_post_rate_exceed"].(int)
		ret.Dst_resp_rate_exceed = in["dst_resp_rate_exceed"].(int)
		ret.Dst_filter_action_whitelist = in["dst_filter_action_whitelist"].(int)
		ret.Src_filter_match = in["src_filter_match"].(int)
		ret.Src_filter_not_match = in["src_filter_not_match"].(int)
		ret.Src_filter_action_blacklist = in["src_filter_action_blacklist"].(int)
		ret.Src_filter_action_drop = in["src_filter_action_drop"].(int)
		ret.Src_filter_action_default_pass = in["src_filter_action_default_pass"].(int)
		ret.Src_filter_action_whitelist = in["src_filter_action_whitelist"].(int)
		ret.Src_dst_filter_match = in["src_dst_filter_match"].(int)
		ret.Src_dst_filter_not_match = in["src_dst_filter_not_match"].(int)
		ret.Src_dst_filter_action_blacklist = in["src_dst_filter_action_blacklist"].(int)
		ret.Src_dst_filter_action_drop = in["src_dst_filter_action_drop"].(int)
		ret.Src_dst_filter_action_default_pass = in["src_dst_filter_action_default_pass"].(int)
		ret.Src_dst_filter_action_whitelist = in["src_dst_filter_action_whitelist"].(int)
		ret.Dst_filter_rate_exceed = in["dst_filter_rate_exceed"].(int)
		ret.Dst_filter_action_ignore = in["dst_filter_action_ignore"].(int)
		ret.Dst_filter_action_reset = in["dst_filter_action_reset"].(int)
		ret.Uri_filter_match = in["uri_filter_match"].(int)
		ret.Http_auth_drop = in["http_auth_drop"].(int)
		ret.Http_auth_resp = in["http_auth_resp"].(int)
		ret.Header_processing_time_0 = in["header_processing_time_0"].(int)
		ret.Header_processing_time_1 = in["header_processing_time_1"].(int)
		ret.Header_processing_time_2 = in["header_processing_time_2"].(int)
		ret.Header_processing_time_3 = in["header_processing_time_3"].(int)
		ret.Header_processing_incomplete = in["header_processing_incomplete"].(int)
		ret.Malform_req_line_too_small = in["malform_req_line_too_small"].(int)
		ret.Malform_req_line_invalid_method = in["malform_req_line_invalid_method"].(int)
	}
	return ret
}

func dataToEndpointDdosL7HttpStats(d *schema.ResourceData) edpt.DdosL7HttpStats {
	var ret edpt.DdosL7HttpStats

	ret.Stats = getObjectDdosL7HttpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
