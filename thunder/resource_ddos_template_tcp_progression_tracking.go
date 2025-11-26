package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateTcpProgressionTracking() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_tcp_progression_tracking`: Configure and enable TCP Progression Tracking\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateTcpProgressionTrackingCreate,
		UpdateContext: resourceDdosTemplateTcpProgressionTrackingUpdate,
		ReadContext:   resourceDdosTemplateTcpProgressionTrackingRead,
		DeleteContext: resourceDdosTemplateTcpProgressionTrackingDelete,

		Schema: map[string]*schema.Schema{
			"ignore_tls_handshake": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore TLS handshake, support SSL-L4 port only",
			},
			"mitigation": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
					},
				},
			},
			"profiling": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"profiling_request_response_model": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable auto-config progression tracking learning for request response model",
						},
						"profiling_connection_life_model": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable auto-config progression tracking learning for connection model",
						},
						"profiling_time_window_model": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable auto-config progression tracking learning for time window model",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"progression_tracking_enabled": {
				Type: schema.TypeString, Required: true, Description: "'enable-check': Enable Progression Tracking Check;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"tcp_name": {
				Type: schema.TypeString, Required: true, Description: "Tcp_name",
			},
		},
	}
}
func resourceDdosTemplateTcpProgressionTrackingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTracking(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateTcpProgressionTrackingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateTcpProgressionTrackingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTracking(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateTcpProgressionTrackingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateTcpProgressionTrackingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTracking(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateTcpProgressionTrackingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTracking(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosTemplateTcpProgressionTrackingMitigation345(d []interface{}) edpt.DdosTemplateTcpProgressionTrackingMitigation345 {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpProgressionTrackingMitigation345
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RequestTracking = getObjectDdosTemplateTcpProgressionTrackingMitigationRequestTracking346(in["request_tracking"].([]interface{}))
		ret.ConnectionTracking = getObjectDdosTemplateTcpProgressionTrackingMitigationConnectionTracking347(in["connection_tracking"].([]interface{}))
		ret.TimeWindowTracking = getObjectDdosTemplateTcpProgressionTrackingMitigationTimeWindowTracking348(in["time_window_tracking"].([]interface{}))
		ret.SlowAttack = getObjectDdosTemplateTcpProgressionTrackingMitigationSlowAttack349(in["slow_attack"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateTcpProgressionTrackingMitigationRequestTracking346(d []interface{}) edpt.DdosTemplateTcpProgressionTrackingMitigationRequestTracking346 {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpProgressionTrackingMitigationRequestTracking346
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

func getObjectDdosTemplateTcpProgressionTrackingMitigationConnectionTracking347(d []interface{}) edpt.DdosTemplateTcpProgressionTrackingMitigationConnectionTracking347 {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpProgressionTrackingMitigationConnectionTracking347
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

func getObjectDdosTemplateTcpProgressionTrackingMitigationTimeWindowTracking348(d []interface{}) edpt.DdosTemplateTcpProgressionTrackingMitigationTimeWindowTracking348 {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpProgressionTrackingMitigationTimeWindowTracking348
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

func getObjectDdosTemplateTcpProgressionTrackingMitigationSlowAttack349(d []interface{}) edpt.DdosTemplateTcpProgressionTrackingMitigationSlowAttack349 {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpProgressionTrackingMitigationSlowAttack349
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

func getObjectDdosTemplateTcpProgressionTrackingProfiling350(d []interface{}) edpt.DdosTemplateTcpProgressionTrackingProfiling350 {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpProgressionTrackingProfiling350
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProfilingRequestResponseModel = in["profiling_request_response_model"].(int)
		ret.ProfilingConnectionLifeModel = in["profiling_connection_life_model"].(int)
		ret.ProfilingTimeWindowModel = in["profiling_time_window_model"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosTemplateTcpProgressionTracking(d *schema.ResourceData) edpt.DdosTemplateTcpProgressionTracking {
	var ret edpt.DdosTemplateTcpProgressionTracking
	ret.Inst.IgnoreTlsHandshake = d.Get("ignore_tls_handshake").(int)
	ret.Inst.Mitigation = getObjectDdosTemplateTcpProgressionTrackingMitigation345(d.Get("mitigation").([]interface{}))
	ret.Inst.Profiling = getObjectDdosTemplateTcpProgressionTrackingProfiling350(d.Get("profiling").([]interface{}))
	ret.Inst.ProgressionTrackingEnabled = d.Get("progression_tracking_enabled").(string)
	//omit uuid
	ret.Inst.Tcp_name = d.Get("tcp_name").(string)
	return ret
}
