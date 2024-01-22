package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosL7SipStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_l7_sip_stats`: Statistics for the object l7-sip\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosL7SipStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Policy Drop",
						},
						"policy_violation": {
							Type: schema.TypeInt, Optional: true, Description: "Policy Violation",
						},
						"idle_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Idle Timeout",
						},
						"ofo_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Out-of-Order Timeout",
						},
						"seq_check_ofo": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence Check Out-Of-Order",
						},
						"pkts_ofo_total": {
							Type: schema.TypeInt, Optional: true, Description: "Packet Out-Of-Order Total",
						},
						"ofo_queue_size_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Out-Of-Order Queue Size Exceed",
						},
						"seq_check_retrans_fin": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence Check Retransmit Fin",
						},
						"seq_check_retrans_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence Check Retransmit Rst",
						},
						"seq_check_retrans_push": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence Check Retransmit Push",
						},
						"seq_check_retrans_other": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence Check Retransmit Other",
						},
						"pkts_retrans_total": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Retransmit Total",
						},
						"client_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Client Rst",
						},
						"error_condition": {
							Type: schema.TypeInt, Optional: true, Description: "Error Condition",
						},
						"request_method_ack": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method ACK",
						},
						"request_method_bye": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method BYE",
						},
						"request_method_cancel": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method CANCEL",
						},
						"request_method_invite": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method INVITE",
						},
						"request_method_info": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method INFO",
						},
						"request_method_message": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method MESSAGE",
						},
						"request_method_notify": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method NOTIFY",
						},
						"request_method_options": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method OPTIONS",
						},
						"request_method_prack": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method PRACK",
						},
						"request_method_publish": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method PUBLISH",
						},
						"request_method_register": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method REGISTER",
						},
						"request_method_refer": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method REFER",
						},
						"request_method_subscribe": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method SUBSCRIBE",
						},
						"request_method_update": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method UPDATE",
						},
						"request_method_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown Request Method",
						},
						"request_unknown_version": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown Request Version",
						},
						"keep_alive_msg": {
							Type: schema.TypeInt, Optional: true, Description: "KeepAlive Message",
						},
						"rate1_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate 1 Limit Exceed",
						},
						"rate2_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate 2 Limit Exceed",
						},
						"src_rate1_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate 1 Limit Exceed",
						},
						"src_rate2_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate 2 Limit Exceed",
						},
						"response_1xx": {
							Type: schema.TypeInt, Optional: true, Description: "Response Status Code 1xx",
						},
						"response_2xx": {
							Type: schema.TypeInt, Optional: true, Description: "Response Status Code 2xx",
						},
						"response_3xx": {
							Type: schema.TypeInt, Optional: true, Description: "Response Status Code 3xx",
						},
						"response_4xx": {
							Type: schema.TypeInt, Optional: true, Description: "Response Status Code 4xx",
						},
						"response_5xx": {
							Type: schema.TypeInt, Optional: true, Description: "Response Status Code 5xx",
						},
						"response_6xx": {
							Type: schema.TypeInt, Optional: true, Description: "Response Status Code 6xx",
						},
						"response_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown Response Status Code",
						},
						"response_unknown_version": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown Response Version",
						},
						"read_start_line_error": {
							Type: schema.TypeInt, Optional: true, Description: "Read Start Line Read Error",
						},
						"invalid_start_line_error": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Start Line",
						},
						"parse_start_line_error": {
							Type: schema.TypeInt, Optional: true, Description: "Start Line Parse Error",
						},
						"line_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "Line Too Long",
						},
						"line_mem_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "Line Memory Allocated",
						},
						"line_mem_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Line Memory Freed",
						},
						"max_uri_len_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Max URI Length Exceed",
						},
						"too_many_header": {
							Type: schema.TypeInt, Optional: true, Description: "Max Header Count Exceed",
						},
						"invalid_header": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Header",
						},
						"header_name_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "Max Header Name Length Exceed",
						},
						"parse_header_fail_error": {
							Type: schema.TypeInt, Optional: true, Description: "Header Parse Fail",
						},
						"max_header_value_len_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Max Header Value Length Exceed",
						},
						"max_call_id_len_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Max Call ID Length Exceed",
						},
						"header_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter Match",
						},
						"header_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter Not Match",
						},
						"header_filter_none_match": {
							Type: schema.TypeInt, Optional: true, Description: "None Header Filter Match",
						},
						"header_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter Action Drop",
						},
						"header_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter Action Blacklist",
						},
						"header_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter Action Whitelist",
						},
						"header_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter Action Default Pass",
						},
						"max_sdp_len_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Max SDP Length Exceed",
						},
						"body_too_big": {
							Type: schema.TypeInt, Optional: true, Description: "Body Too Big",
						},
						"get_content_fail_error": {
							Type: schema.TypeInt, Optional: true, Description: "Get Content Fail",
						},
						"concatenate_msg": {
							Type: schema.TypeInt, Optional: true, Description: "Concatenate Msessage",
						},
						"mem_alloc_fail_error": {
							Type: schema.TypeInt, Optional: true, Description: "Memory Allocate Fail",
						},
						"malform_request": {
							Type: schema.TypeInt, Optional: true, Description: "Malformed Request",
						},
						"src_header_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Header Filter Match",
						},
						"src_header_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Header Filter Not Match",
						},
						"src_header_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Header Filter Action Drop",
						},
						"src_header_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Src Header Filter Action Blacklist",
						},
						"src_header_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Src Header Filter Action Whitelist",
						},
						"src_header_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Src Header Filter Action Default Pass",
						},
						"src_dst_header_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Header Filter Match",
						},
						"src_dst_header_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Header Filter Not Match",
						},
						"src_dst_header_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Header Filter Action Drop",
						},
						"src_dst_header_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Src Header Filter Action Blacklist",
						},
						"src_dst_header_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Src Header Filter Action Whitelist",
						},
						"src_dst_header_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Src Header Filter Action Default Pass",
						},
					},
				},
			},
		},
	}
}

func resourceDdosL7SipStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosL7SipStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosL7SipStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosL7SipStatsStats := setObjectDdosL7SipStatsStats(res)
		d.Set("stats", DdosL7SipStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosL7SipStatsStats(ret edpt.DataDdosL7SipStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"policy_drop":                               ret.DtDdosL7SipStats.Stats.Policy_drop,
			"policy_violation":                          ret.DtDdosL7SipStats.Stats.Policy_violation,
			"idle_timeout":                              ret.DtDdosL7SipStats.Stats.Idle_timeout,
			"ofo_timeout":                               ret.DtDdosL7SipStats.Stats.Ofo_timeout,
			"seq_check_ofo":                             ret.DtDdosL7SipStats.Stats.Seq_check_ofo,
			"pkts_ofo_total":                            ret.DtDdosL7SipStats.Stats.Pkts_ofo_total,
			"ofo_queue_size_exceed":                     ret.DtDdosL7SipStats.Stats.Ofo_queue_size_exceed,
			"seq_check_retrans_fin":                     ret.DtDdosL7SipStats.Stats.Seq_check_retrans_fin,
			"seq_check_retrans_rst":                     ret.DtDdosL7SipStats.Stats.Seq_check_retrans_rst,
			"seq_check_retrans_push":                    ret.DtDdosL7SipStats.Stats.Seq_check_retrans_push,
			"seq_check_retrans_other":                   ret.DtDdosL7SipStats.Stats.Seq_check_retrans_other,
			"pkts_retrans_total":                        ret.DtDdosL7SipStats.Stats.Pkts_retrans_total,
			"client_rst":                                ret.DtDdosL7SipStats.Stats.Client_rst,
			"error_condition":                           ret.DtDdosL7SipStats.Stats.Error_condition,
			"request_method_ack":                        ret.DtDdosL7SipStats.Stats.Request_method_ack,
			"request_method_bye":                        ret.DtDdosL7SipStats.Stats.Request_method_bye,
			"request_method_cancel":                     ret.DtDdosL7SipStats.Stats.Request_method_cancel,
			"request_method_invite":                     ret.DtDdosL7SipStats.Stats.Request_method_invite,
			"request_method_info":                       ret.DtDdosL7SipStats.Stats.Request_method_info,
			"request_method_message":                    ret.DtDdosL7SipStats.Stats.Request_method_message,
			"request_method_notify":                     ret.DtDdosL7SipStats.Stats.Request_method_notify,
			"request_method_options":                    ret.DtDdosL7SipStats.Stats.Request_method_options,
			"request_method_prack":                      ret.DtDdosL7SipStats.Stats.Request_method_prack,
			"request_method_publish":                    ret.DtDdosL7SipStats.Stats.Request_method_publish,
			"request_method_register":                   ret.DtDdosL7SipStats.Stats.Request_method_register,
			"request_method_refer":                      ret.DtDdosL7SipStats.Stats.Request_method_refer,
			"request_method_subscribe":                  ret.DtDdosL7SipStats.Stats.Request_method_subscribe,
			"request_method_update":                     ret.DtDdosL7SipStats.Stats.Request_method_update,
			"request_method_unknown":                    ret.DtDdosL7SipStats.Stats.Request_method_unknown,
			"request_unknown_version":                   ret.DtDdosL7SipStats.Stats.Request_unknown_version,
			"keep_alive_msg":                            ret.DtDdosL7SipStats.Stats.Keep_alive_msg,
			"rate1_limit_exceed":                        ret.DtDdosL7SipStats.Stats.Rate1_limit_exceed,
			"rate2_limit_exceed":                        ret.DtDdosL7SipStats.Stats.Rate2_limit_exceed,
			"src_rate1_limit_exceed":                    ret.DtDdosL7SipStats.Stats.Src_rate1_limit_exceed,
			"src_rate2_limit_exceed":                    ret.DtDdosL7SipStats.Stats.Src_rate2_limit_exceed,
			"response_1xx":                              ret.DtDdosL7SipStats.Stats.Response_1xx,
			"response_2xx":                              ret.DtDdosL7SipStats.Stats.Response_2xx,
			"response_3xx":                              ret.DtDdosL7SipStats.Stats.Response_3xx,
			"response_4xx":                              ret.DtDdosL7SipStats.Stats.Response_4xx,
			"response_5xx":                              ret.DtDdosL7SipStats.Stats.Response_5xx,
			"response_6xx":                              ret.DtDdosL7SipStats.Stats.Response_6xx,
			"response_unknown":                          ret.DtDdosL7SipStats.Stats.Response_unknown,
			"response_unknown_version":                  ret.DtDdosL7SipStats.Stats.Response_unknown_version,
			"read_start_line_error":                     ret.DtDdosL7SipStats.Stats.Read_start_line_error,
			"invalid_start_line_error":                  ret.DtDdosL7SipStats.Stats.Invalid_start_line_error,
			"parse_start_line_error":                    ret.DtDdosL7SipStats.Stats.Parse_start_line_error,
			"line_too_long":                             ret.DtDdosL7SipStats.Stats.Line_too_long,
			"line_mem_allocated":                        ret.DtDdosL7SipStats.Stats.Line_mem_allocated,
			"line_mem_freed":                            ret.DtDdosL7SipStats.Stats.Line_mem_freed,
			"max_uri_len_exceed":                        ret.DtDdosL7SipStats.Stats.Max_uri_len_exceed,
			"too_many_header":                           ret.DtDdosL7SipStats.Stats.Too_many_header,
			"invalid_header":                            ret.DtDdosL7SipStats.Stats.Invalid_header,
			"header_name_too_long":                      ret.DtDdosL7SipStats.Stats.Header_name_too_long,
			"parse_header_fail_error":                   ret.DtDdosL7SipStats.Stats.Parse_header_fail_error,
			"max_header_value_len_exceed":               ret.DtDdosL7SipStats.Stats.Max_header_value_len_exceed,
			"max_call_id_len_exceed":                    ret.DtDdosL7SipStats.Stats.Max_call_id_len_exceed,
			"header_filter_match":                       ret.DtDdosL7SipStats.Stats.Header_filter_match,
			"header_filter_not_match":                   ret.DtDdosL7SipStats.Stats.Header_filter_not_match,
			"header_filter_none_match":                  ret.DtDdosL7SipStats.Stats.Header_filter_none_match,
			"header_filter_action_drop":                 ret.DtDdosL7SipStats.Stats.Header_filter_action_drop,
			"header_filter_action_blacklist":            ret.DtDdosL7SipStats.Stats.Header_filter_action_blacklist,
			"header_filter_action_whitelist":            ret.DtDdosL7SipStats.Stats.Header_filter_action_whitelist,
			"header_filter_action_default_pass":         ret.DtDdosL7SipStats.Stats.Header_filter_action_default_pass,
			"max_sdp_len_exceed":                        ret.DtDdosL7SipStats.Stats.Max_sdp_len_exceed,
			"body_too_big":                              ret.DtDdosL7SipStats.Stats.Body_too_big,
			"get_content_fail_error":                    ret.DtDdosL7SipStats.Stats.Get_content_fail_error,
			"concatenate_msg":                           ret.DtDdosL7SipStats.Stats.Concatenate_msg,
			"mem_alloc_fail_error":                      ret.DtDdosL7SipStats.Stats.Mem_alloc_fail_error,
			"malform_request":                           ret.DtDdosL7SipStats.Stats.Malform_request,
			"src_header_filter_match":                   ret.DtDdosL7SipStats.Stats.Src_header_filter_match,
			"src_header_filter_not_match":               ret.DtDdosL7SipStats.Stats.Src_header_filter_not_match,
			"src_header_filter_action_drop":             ret.DtDdosL7SipStats.Stats.Src_header_filter_action_drop,
			"src_header_filter_action_blacklist":        ret.DtDdosL7SipStats.Stats.Src_header_filter_action_blacklist,
			"src_header_filter_action_whitelist":        ret.DtDdosL7SipStats.Stats.Src_header_filter_action_whitelist,
			"src_header_filter_action_default_pass":     ret.DtDdosL7SipStats.Stats.Src_header_filter_action_default_pass,
			"src_dst_header_filter_match":               ret.DtDdosL7SipStats.Stats.Src_dst_header_filter_match,
			"src_dst_header_filter_not_match":           ret.DtDdosL7SipStats.Stats.Src_dst_header_filter_not_match,
			"src_dst_header_filter_action_drop":         ret.DtDdosL7SipStats.Stats.Src_dst_header_filter_action_drop,
			"src_dst_header_filter_action_blacklist":    ret.DtDdosL7SipStats.Stats.Src_dst_header_filter_action_blacklist,
			"src_dst_header_filter_action_whitelist":    ret.DtDdosL7SipStats.Stats.Src_dst_header_filter_action_whitelist,
			"src_dst_header_filter_action_default_pass": ret.DtDdosL7SipStats.Stats.Src_dst_header_filter_action_default_pass,
		},
	}
}

func getObjectDdosL7SipStatsStats(d []interface{}) edpt.DdosL7SipStatsStats {

	count1 := len(d)
	var ret edpt.DdosL7SipStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Policy_drop = in["policy_drop"].(int)
		ret.Policy_violation = in["policy_violation"].(int)
		ret.Idle_timeout = in["idle_timeout"].(int)
		ret.Ofo_timeout = in["ofo_timeout"].(int)
		ret.Seq_check_ofo = in["seq_check_ofo"].(int)
		ret.Pkts_ofo_total = in["pkts_ofo_total"].(int)
		ret.Ofo_queue_size_exceed = in["ofo_queue_size_exceed"].(int)
		ret.Seq_check_retrans_fin = in["seq_check_retrans_fin"].(int)
		ret.Seq_check_retrans_rst = in["seq_check_retrans_rst"].(int)
		ret.Seq_check_retrans_push = in["seq_check_retrans_push"].(int)
		ret.Seq_check_retrans_other = in["seq_check_retrans_other"].(int)
		ret.Pkts_retrans_total = in["pkts_retrans_total"].(int)
		ret.Client_rst = in["client_rst"].(int)
		ret.Error_condition = in["error_condition"].(int)
		ret.Request_method_ack = in["request_method_ack"].(int)
		ret.Request_method_bye = in["request_method_bye"].(int)
		ret.Request_method_cancel = in["request_method_cancel"].(int)
		ret.Request_method_invite = in["request_method_invite"].(int)
		ret.Request_method_info = in["request_method_info"].(int)
		ret.Request_method_message = in["request_method_message"].(int)
		ret.Request_method_notify = in["request_method_notify"].(int)
		ret.Request_method_options = in["request_method_options"].(int)
		ret.Request_method_prack = in["request_method_prack"].(int)
		ret.Request_method_publish = in["request_method_publish"].(int)
		ret.Request_method_register = in["request_method_register"].(int)
		ret.Request_method_refer = in["request_method_refer"].(int)
		ret.Request_method_subscribe = in["request_method_subscribe"].(int)
		ret.Request_method_update = in["request_method_update"].(int)
		ret.Request_method_unknown = in["request_method_unknown"].(int)
		ret.Request_unknown_version = in["request_unknown_version"].(int)
		ret.Keep_alive_msg = in["keep_alive_msg"].(int)
		ret.Rate1_limit_exceed = in["rate1_limit_exceed"].(int)
		ret.Rate2_limit_exceed = in["rate2_limit_exceed"].(int)
		ret.Src_rate1_limit_exceed = in["src_rate1_limit_exceed"].(int)
		ret.Src_rate2_limit_exceed = in["src_rate2_limit_exceed"].(int)
		ret.Response_1xx = in["response_1xx"].(int)
		ret.Response_2xx = in["response_2xx"].(int)
		ret.Response_3xx = in["response_3xx"].(int)
		ret.Response_4xx = in["response_4xx"].(int)
		ret.Response_5xx = in["response_5xx"].(int)
		ret.Response_6xx = in["response_6xx"].(int)
		ret.Response_unknown = in["response_unknown"].(int)
		ret.Response_unknown_version = in["response_unknown_version"].(int)
		ret.Read_start_line_error = in["read_start_line_error"].(int)
		ret.Invalid_start_line_error = in["invalid_start_line_error"].(int)
		ret.Parse_start_line_error = in["parse_start_line_error"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Line_mem_allocated = in["line_mem_allocated"].(int)
		ret.Line_mem_freed = in["line_mem_freed"].(int)
		ret.Max_uri_len_exceed = in["max_uri_len_exceed"].(int)
		ret.Too_many_header = in["too_many_header"].(int)
		ret.Invalid_header = in["invalid_header"].(int)
		ret.Header_name_too_long = in["header_name_too_long"].(int)
		ret.Parse_header_fail_error = in["parse_header_fail_error"].(int)
		ret.Max_header_value_len_exceed = in["max_header_value_len_exceed"].(int)
		ret.Max_call_id_len_exceed = in["max_call_id_len_exceed"].(int)
		ret.Header_filter_match = in["header_filter_match"].(int)
		ret.Header_filter_not_match = in["header_filter_not_match"].(int)
		ret.Header_filter_none_match = in["header_filter_none_match"].(int)
		ret.Header_filter_action_drop = in["header_filter_action_drop"].(int)
		ret.Header_filter_action_blacklist = in["header_filter_action_blacklist"].(int)
		ret.Header_filter_action_whitelist = in["header_filter_action_whitelist"].(int)
		ret.Header_filter_action_default_pass = in["header_filter_action_default_pass"].(int)
		ret.Max_sdp_len_exceed = in["max_sdp_len_exceed"].(int)
		ret.Body_too_big = in["body_too_big"].(int)
		ret.Get_content_fail_error = in["get_content_fail_error"].(int)
		ret.Concatenate_msg = in["concatenate_msg"].(int)
		ret.Mem_alloc_fail_error = in["mem_alloc_fail_error"].(int)
		ret.Malform_request = in["malform_request"].(int)
		ret.Src_header_filter_match = in["src_header_filter_match"].(int)
		ret.Src_header_filter_not_match = in["src_header_filter_not_match"].(int)
		ret.Src_header_filter_action_drop = in["src_header_filter_action_drop"].(int)
		ret.Src_header_filter_action_blacklist = in["src_header_filter_action_blacklist"].(int)
		ret.Src_header_filter_action_whitelist = in["src_header_filter_action_whitelist"].(int)
		ret.Src_header_filter_action_default_pass = in["src_header_filter_action_default_pass"].(int)
		ret.Src_dst_header_filter_match = in["src_dst_header_filter_match"].(int)
		ret.Src_dst_header_filter_not_match = in["src_dst_header_filter_not_match"].(int)
		ret.Src_dst_header_filter_action_drop = in["src_dst_header_filter_action_drop"].(int)
		ret.Src_dst_header_filter_action_blacklist = in["src_dst_header_filter_action_blacklist"].(int)
		ret.Src_dst_header_filter_action_whitelist = in["src_dst_header_filter_action_whitelist"].(int)
		ret.Src_dst_header_filter_action_default_pass = in["src_dst_header_filter_action_default_pass"].(int)
	}
	return ret
}

func dataToEndpointDdosL7SipStats(d *schema.ResourceData) edpt.DdosL7SipStats {
	var ret edpt.DdosL7SipStats

	ret.Stats = getObjectDdosL7SipStatsStats(d.Get("stats").([]interface{}))
	return ret
}
