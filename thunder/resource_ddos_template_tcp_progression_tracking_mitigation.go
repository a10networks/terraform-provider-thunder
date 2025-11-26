package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateTcpProgressionTrackingMitigation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_tcp_progression_tracking_mitigation`: Configure and enable TCP Progression Tracking Mitigation\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateTcpProgressionTrackingMitigationCreate,
		UpdateContext: resourceDdosTemplateTcpProgressionTrackingMitigationUpdate,
		ReadContext:   resourceDdosTemplateTcpProgressionTrackingMitigationRead,
		DeleteContext: resourceDdosTemplateTcpProgressionTrackingMitigationDelete,

		Schema: map[string]*schema.Schema{
			"connection_tracking": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"progression_tracking_conn_enabled": {
							Type: schema.TypeString, Optional: true, Description: "'enable-check': Enable General Progression Tracking per Connection;",
						},
						"conn_sent_max": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum total sent byte",
						},
						"conn_sent_min": {
							Type: schema.TypeInt, Optional: true, Description: "Set the minimum total sent byte",
						},
						"conn_rcvd_max": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum total received byte",
						},
						"conn_rcvd_min": {
							Type: schema.TypeInt, Optional: true, Description: "Set the minimum total received byte",
						},
						"conn_rcvd_sent_ratio_min": {
							Type: schema.TypeInt, Optional: true, Description: "Set the minimum received to sent ratio (in unit of 0.1% [1:1000])",
						},
						"conn_rcvd_sent_ratio_max": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum received to sent ratio (in unit of 0.1% [1:1000])",
						},
						"conn_duration_max": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum duration time (in unit of 100ms, up to 24 hours)",
						},
						"conn_duration_min": {
							Type: schema.TypeInt, Optional: true, Description: "Set the minimum duration time (in unit of 100ms, up to 24 hours)",
						},
						"conn_violation": {
							Type: schema.TypeInt, Optional: true, Description: "Set the violation threshold",
						},
						"progression_tracking_conn_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take when progression tracking violation exceed",
						},
						"progression_tracking_conn_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for progression tracking violation exceed (Default); 'blacklist-src': Blacklist-src for progression tracking violation exceed;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"request_tracking": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"progression_tracking_req_enabled": {
							Type: schema.TypeString, Optional: true, Description: "'enable-check': Enable General Progression Tracking per Request Response;",
						},
						"request_response_model": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable Request Response Model; 'disable': Disable Request Response Model;",
						},
						"response_length_max": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum response length",
						},
						"response_length_min": {
							Type: schema.TypeInt, Optional: true, Description: "Set the minimum response length",
						},
						"request_length_min": {
							Type: schema.TypeInt, Optional: true, Description: "Set the minimum request length",
						},
						"request_length_max": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum request length",
						},
						"request_to_response_max_time": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum request to response time (in unit of 100ms)",
						},
						"response_to_request_max_time": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum response to request time (in unit of 100ms)",
						},
						"first_request_max_time": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum idle time before the first request (in unit of 100ms)",
						},
						"progression_tracking_req_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take when progression tracking violation exceed",
						},
						"violation": {
							Type: schema.TypeInt, Optional: true, Description: "Set the violation threshold",
						},
						"progression_tracking_req_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for progression tracking violation exceed (Default); 'blacklist-src': Blacklist-src for progression tracking violation exceed;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"slow_attack": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"response_pkt_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "Set the transferred packets per response",
						},
						"init_response_max_time": {
							Type: schema.TypeInt, Optional: true, Description: "Set server think time (in unit of 100ms). Suggested value larger than 45 secs",
						},
						"init_request_max_time": {
							Type: schema.TypeInt, Optional: true, Description: "Set client query time (in unit of 100ms). Suggested value larger than 30 secs",
						},
						"progression_tracking_slow_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take when progression tracking violation exceed",
						},
						"progression_tracking_slow_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for progression tracking violation exceed (Default); 'reset': Reset client connection; 'blacklist-src': Blacklist-src for progression tracking violation exceed;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"time_window_tracking": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"progression_tracking_win_enabled": {
							Type: schema.TypeString, Optional: true, Description: "'enable-check': Enable Progression Tracking per Time Window;",
						},
						"window_sent_max": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum total sent byte",
						},
						"window_sent_min": {
							Type: schema.TypeInt, Optional: true, Description: "Set the minimum total sent byte",
						},
						"window_rcvd_max": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum total received byte",
						},
						"window_rcvd_min": {
							Type: schema.TypeInt, Optional: true, Description: "Set the minimum total received byte",
						},
						"window_rcvd_sent_ratio_min": {
							Type: schema.TypeInt, Optional: true, Description: "Set the minimum received to sent ratio (in unit of 0.1% [1:1000])",
						},
						"window_rcvd_sent_ratio_max": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum received to sent ratio (in unit of 0.1% [1:1000])",
						},
						"window_violation": {
							Type: schema.TypeInt, Optional: true, Description: "Set the violation threshold",
						},
						"progression_tracking_windows_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take when progression tracking violation exceed",
						},
						"progression_tracking_windows_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for progression tracking violation exceed (Default); 'blacklist-src': Blacklist-src for progression tracking violation exceed;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"tcp_name": {
				Type: schema.TypeString, Required: true, Description: "Tcp_name",
			},
		},
	}
}
func resourceDdosTemplateTcpProgressionTrackingMitigationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingMitigationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingMitigation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateTcpProgressionTrackingMitigationRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateTcpProgressionTrackingMitigationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingMitigationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingMitigation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateTcpProgressionTrackingMitigationRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateTcpProgressionTrackingMitigationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingMitigationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingMitigation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateTcpProgressionTrackingMitigationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingMitigationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingMitigation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosTemplateTcpProgressionTrackingMitigationConnectionTracking341(d []interface{}) edpt.DdosTemplateTcpProgressionTrackingMitigationConnectionTracking341 {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpProgressionTrackingMitigationConnectionTracking341
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProgressionTrackingConnEnabled = in["progression_tracking_conn_enabled"].(string)
		ret.ConnSentMax = in["conn_sent_max"].(int)
		ret.ConnSentMin = in["conn_sent_min"].(int)
		ret.ConnRcvdMax = in["conn_rcvd_max"].(int)
		ret.ConnRcvdMin = in["conn_rcvd_min"].(int)
		ret.ConnRcvdSentRatioMin = in["conn_rcvd_sent_ratio_min"].(int)
		ret.ConnRcvdSentRatioMax = in["conn_rcvd_sent_ratio_max"].(int)
		ret.ConnDurationMax = in["conn_duration_max"].(int)
		ret.ConnDurationMin = in["conn_duration_min"].(int)
		ret.ConnViolation = in["conn_violation"].(int)
		ret.ProgressionTrackingConnActionListName = in["progression_tracking_conn_action_list_name"].(string)
		ret.ProgressionTrackingConnAction = in["progression_tracking_conn_action"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosTemplateTcpProgressionTrackingMitigationRequestTracking342(d []interface{}) edpt.DdosTemplateTcpProgressionTrackingMitigationRequestTracking342 {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpProgressionTrackingMitigationRequestTracking342
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProgressionTrackingReqEnabled = in["progression_tracking_req_enabled"].(string)
		ret.RequestResponseModel = in["request_response_model"].(string)
		ret.ResponseLengthMax = in["response_length_max"].(int)
		ret.ResponseLengthMin = in["response_length_min"].(int)
		ret.RequestLengthMin = in["request_length_min"].(int)
		ret.RequestLengthMax = in["request_length_max"].(int)
		ret.RequestToResponseMaxTime = in["request_to_response_max_time"].(int)
		ret.ResponseToRequestMaxTime = in["response_to_request_max_time"].(int)
		ret.FirstRequestMaxTime = in["first_request_max_time"].(int)
		ret.ProgressionTrackingReqActionListName = in["progression_tracking_req_action_list_name"].(string)
		ret.Violation = in["violation"].(int)
		ret.ProgressionTrackingReqAction = in["progression_tracking_req_action"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosTemplateTcpProgressionTrackingMitigationSlowAttack343(d []interface{}) edpt.DdosTemplateTcpProgressionTrackingMitigationSlowAttack343 {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpProgressionTrackingMitigationSlowAttack343
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ResponsePktRateMax = in["response_pkt_rate_max"].(int)
		ret.InitResponseMaxTime = in["init_response_max_time"].(int)
		ret.InitRequestMaxTime = in["init_request_max_time"].(int)
		ret.ProgressionTrackingSlowActionListName = in["progression_tracking_slow_action_list_name"].(string)
		ret.ProgressionTrackingSlowAction = in["progression_tracking_slow_action"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosTemplateTcpProgressionTrackingMitigationTimeWindowTracking344(d []interface{}) edpt.DdosTemplateTcpProgressionTrackingMitigationTimeWindowTracking344 {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpProgressionTrackingMitigationTimeWindowTracking344
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProgressionTrackingWinEnabled = in["progression_tracking_win_enabled"].(string)
		ret.WindowSentMax = in["window_sent_max"].(int)
		ret.WindowSentMin = in["window_sent_min"].(int)
		ret.WindowRcvdMax = in["window_rcvd_max"].(int)
		ret.WindowRcvdMin = in["window_rcvd_min"].(int)
		ret.WindowRcvdSentRatioMin = in["window_rcvd_sent_ratio_min"].(int)
		ret.WindowRcvdSentRatioMax = in["window_rcvd_sent_ratio_max"].(int)
		ret.WindowViolation = in["window_violation"].(int)
		ret.ProgressionTrackingWindowsActionListName = in["progression_tracking_windows_action_list_name"].(string)
		ret.ProgressionTrackingWindowsAction = in["progression_tracking_windows_action"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosTemplateTcpProgressionTrackingMitigation(d *schema.ResourceData) edpt.DdosTemplateTcpProgressionTrackingMitigation {
	var ret edpt.DdosTemplateTcpProgressionTrackingMitigation
	ret.Inst.ConnectionTracking = getObjectDdosTemplateTcpProgressionTrackingMitigationConnectionTracking341(d.Get("connection_tracking").([]interface{}))
	ret.Inst.RequestTracking = getObjectDdosTemplateTcpProgressionTrackingMitigationRequestTracking342(d.Get("request_tracking").([]interface{}))
	ret.Inst.SlowAttack = getObjectDdosTemplateTcpProgressionTrackingMitigationSlowAttack343(d.Get("slow_attack").([]interface{}))
	ret.Inst.TimeWindowTracking = getObjectDdosTemplateTcpProgressionTrackingMitigationTimeWindowTracking344(d.Get("time_window_tracking").([]interface{}))
	ret.Inst.Tcp_name = d.Get("tcp_name").(string)
	return ret
}
