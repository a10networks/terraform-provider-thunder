package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateTcpProgressionTrackingMitigationRequestTracking() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_tcp_progression_tracking_mitigation_request_tracking`: Configure and enable TCP Progression Tracking Mitigation per Request response\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateTcpProgressionTrackingMitigationRequestTrackingCreate,
		UpdateContext: resourceDdosTemplateTcpProgressionTrackingMitigationRequestTrackingUpdate,
		ReadContext:   resourceDdosTemplateTcpProgressionTrackingMitigationRequestTrackingRead,
		DeleteContext: resourceDdosTemplateTcpProgressionTrackingMitigationRequestTrackingDelete,

		Schema: map[string]*schema.Schema{
			"first_request_max_time": {
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum idle time before the first request (in unit of 100ms)",
			},
			"progression_tracking_req_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for progression tracking violation exceed (Default); 'blacklist-src': Blacklist-src for progression tracking violation exceed;",
			},
			"progression_tracking_req_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take when progression tracking violation exceed",
			},
			"progression_tracking_req_enabled": {
				Type: schema.TypeString, Required: true, Description: "'enable-check': Enable General Progression Tracking per Request Response;",
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
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum request to response time (in unit of 100ms)",
			},
			"response_length_max": {
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum response length",
			},
			"response_length_min": {
				Type: schema.TypeInt, Optional: true, Description: "Set the minimum response length",
			},
			"response_to_request_max_time": {
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum response to request time (in unit of 100ms)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"violation": {
				Type: schema.TypeInt, Optional: true, Description: "Set the violation threshold",
			},
			"tcp_name": {
				Type: schema.TypeString, Required: true, Description: "Tcp_name",
			},
		},
	}
}
func resourceDdosTemplateTcpProgressionTrackingMitigationRequestTrackingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingMitigationRequestTrackingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingMitigationRequestTracking(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateTcpProgressionTrackingMitigationRequestTrackingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateTcpProgressionTrackingMitigationRequestTrackingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingMitigationRequestTrackingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingMitigationRequestTracking(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateTcpProgressionTrackingMitigationRequestTrackingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateTcpProgressionTrackingMitigationRequestTrackingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingMitigationRequestTrackingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingMitigationRequestTracking(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateTcpProgressionTrackingMitigationRequestTrackingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingMitigationRequestTrackingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingMitigationRequestTracking(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTemplateTcpProgressionTrackingMitigationRequestTracking(d *schema.ResourceData) edpt.DdosTemplateTcpProgressionTrackingMitigationRequestTracking {
	var ret edpt.DdosTemplateTcpProgressionTrackingMitigationRequestTracking
	ret.Inst.FirstRequestMaxTime = d.Get("first_request_max_time").(int)
	ret.Inst.ProgressionTrackingReqAction = d.Get("progression_tracking_req_action").(string)
	ret.Inst.ProgressionTrackingReqActionListName = d.Get("progression_tracking_req_action_list_name").(string)
	ret.Inst.ProgressionTrackingReqEnabled = d.Get("progression_tracking_req_enabled").(string)
	ret.Inst.RequestLengthMax = d.Get("request_length_max").(int)
	ret.Inst.RequestLengthMin = d.Get("request_length_min").(int)
	ret.Inst.RequestResponseModel = d.Get("request_response_model").(string)
	ret.Inst.RequestToResponseMaxTime = d.Get("request_to_response_max_time").(int)
	ret.Inst.ResponseLengthMax = d.Get("response_length_max").(int)
	ret.Inst.ResponseLengthMin = d.Get("response_length_min").(int)
	ret.Inst.ResponseToRequestMaxTime = d.Get("response_to_request_max_time").(int)
	//omit uuid
	ret.Inst.Violation = d.Get("violation").(int)
	ret.Inst.Tcp_name = d.Get("tcp_name").(string)
	return ret
}
