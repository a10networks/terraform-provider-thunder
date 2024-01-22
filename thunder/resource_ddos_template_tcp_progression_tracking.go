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
			"first_request_max_time": {
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum wait time from connection creation until the first data is transmitted over the connection (100 ms)",
			},
			"profiling_connection_life_model": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable auto-config progression tracking learning for connection model",
			},
			"profiling_request_response_model": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable auto-config progression tracking learning for request response model",
			},
			"profiling_time_window_model": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable auto-config progression tracking learning for time window model",
			},
			"progression_tracking_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for progression tracking violation exceed (Default); 'blacklist-src': Blacklist-src for progression tracking violation exceed;",
			},
			"progression_tracking_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take when progression tracking violation exceed",
			},
			"progression_tracking_enabled": {
				Type: schema.TypeString, Required: true, Description: "'enable-check': Enable Progression Tracking Check;",
			},
			"request_length_max": {
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum request length",
			},
			"request_length_min": {
				Type: schema.TypeInt, Optional: true, Description: "Set the minimum request length",
			},
			"request_response_model": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable Request Response Model; 'disable': Disable Request Response Model;",
			},
			"request_to_response_max_time": {
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum request to response time (100 ms)",
			},
			"response_length_max": {
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum response length",
			},
			"response_length_min": {
				Type: schema.TypeInt, Optional: true, Description: "Set the minimum response length",
			},
			"response_request_max_ratio": {
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum response to request ratio (in unit of 0.1% [1:1000])",
			},
			"response_request_min_ratio": {
				Type: schema.TypeInt, Optional: true, Description: "Set the minimum response to request ratio (in unit of 0.1% [1:1000])",
			},
			"response_to_request_max_time": {
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum response to request time (100 ms)",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"violation": {
				Type: schema.TypeInt, Optional: true, Description: "Set the violation threshold",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
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

func getObjectDdosTemplateTcpProgressionTrackingConnectionTracking300(d []interface{}) edpt.DdosTemplateTcpProgressionTrackingConnectionTracking300 {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpProgressionTrackingConnectionTracking300
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

func getObjectDdosTemplateTcpProgressionTrackingTimeWindowTracking301(d []interface{}) edpt.DdosTemplateTcpProgressionTrackingTimeWindowTracking301 {

	count1 := len(d)
	var ret edpt.DdosTemplateTcpProgressionTrackingTimeWindowTracking301
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

func dataToEndpointDdosTemplateTcpProgressionTracking(d *schema.ResourceData) edpt.DdosTemplateTcpProgressionTracking {
	var ret edpt.DdosTemplateTcpProgressionTracking
	ret.Inst.ConnectionTracking = getObjectDdosTemplateTcpProgressionTrackingConnectionTracking300(d.Get("connection_tracking").([]interface{}))
	ret.Inst.FirstRequestMaxTime = d.Get("first_request_max_time").(int)
	ret.Inst.ProfilingConnectionLifeModel = d.Get("profiling_connection_life_model").(int)
	ret.Inst.ProfilingRequestResponseModel = d.Get("profiling_request_response_model").(int)
	ret.Inst.ProfilingTimeWindowModel = d.Get("profiling_time_window_model").(int)
	ret.Inst.ProgressionTrackingAction = d.Get("progression_tracking_action").(string)
	ret.Inst.ProgressionTrackingActionListName = d.Get("progression_tracking_action_list_name").(string)
	ret.Inst.ProgressionTrackingEnabled = d.Get("progression_tracking_enabled").(string)
	ret.Inst.RequestLengthMax = d.Get("request_length_max").(int)
	ret.Inst.RequestLengthMin = d.Get("request_length_min").(int)
	ret.Inst.RequestResponseModel = d.Get("request_response_model").(string)
	ret.Inst.RequestToResponseMaxTime = d.Get("request_to_response_max_time").(int)
	ret.Inst.ResponseLengthMax = d.Get("response_length_max").(int)
	ret.Inst.ResponseLengthMin = d.Get("response_length_min").(int)
	ret.Inst.ResponseRequestMaxRatio = d.Get("response_request_max_ratio").(int)
	ret.Inst.ResponseRequestMinRatio = d.Get("response_request_min_ratio").(int)
	ret.Inst.ResponseToRequestMaxTime = d.Get("response_to_request_max_time").(int)
	ret.Inst.TimeWindowTracking = getObjectDdosTemplateTcpProgressionTrackingTimeWindowTracking301(d.Get("time_window_tracking").([]interface{}))
	//omit uuid
	ret.Inst.Violation = d.Get("violation").(int)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
