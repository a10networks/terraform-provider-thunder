package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateTcpProgressionTrackingTimeWindowTracking() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_tcp_progression_tracking_time_window_tracking`: Configure and enable TCP Progression Tracking per Time Window\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateTcpProgressionTrackingTimeWindowTrackingCreate,
		UpdateContext: resourceDdosTemplateTcpProgressionTrackingTimeWindowTrackingUpdate,
		ReadContext:   resourceDdosTemplateTcpProgressionTrackingTimeWindowTrackingRead,
		DeleteContext: resourceDdosTemplateTcpProgressionTrackingTimeWindowTrackingDelete,

		Schema: map[string]*schema.Schema{
			"progression_tracking_win_enabled": {
				Type: schema.TypeString, Required: true, Description: "'enable-check': Enable Progression Tracking per Time Window;",
			},
			"progression_tracking_windows_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for progression tracking violation exceed (Default); 'blacklist-src': Blacklist-src for progression tracking violation exceed;",
			},
			"progression_tracking_windows_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take when progression tracking violation exceed",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"window_rcvd_max": {
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum total received byte",
			},
			"window_rcvd_min": {
				Type: schema.TypeInt, Optional: true, Description: "Set the minimum total received byte",
			},
			"window_rcvd_sent_ratio_max": {
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum received to sent ratio (in unit of 0.1% [1:1000])",
			},
			"window_rcvd_sent_ratio_min": {
				Type: schema.TypeInt, Optional: true, Description: "Set the minimum received to sent ratio (in unit of 0.1% [1:1000])",
			},
			"window_sent_max": {
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum total sent byte",
			},
			"window_sent_min": {
				Type: schema.TypeInt, Optional: true, Description: "Set the minimum total sent byte",
			},
			"window_violation": {
				Type: schema.TypeInt, Optional: true, Description: "Set the violation threshold",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceDdosTemplateTcpProgressionTrackingTimeWindowTrackingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingTimeWindowTrackingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingTimeWindowTracking(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateTcpProgressionTrackingTimeWindowTrackingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateTcpProgressionTrackingTimeWindowTrackingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingTimeWindowTrackingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingTimeWindowTracking(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateTcpProgressionTrackingTimeWindowTrackingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateTcpProgressionTrackingTimeWindowTrackingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingTimeWindowTrackingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingTimeWindowTracking(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateTcpProgressionTrackingTimeWindowTrackingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingTimeWindowTrackingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingTimeWindowTracking(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTemplateTcpProgressionTrackingTimeWindowTracking(d *schema.ResourceData) edpt.DdosTemplateTcpProgressionTrackingTimeWindowTracking {
	var ret edpt.DdosTemplateTcpProgressionTrackingTimeWindowTracking
	ret.Inst.ProgressionTrackingWinEnabled = d.Get("progression_tracking_win_enabled").(string)
	ret.Inst.ProgressionTrackingWindowsAction = d.Get("progression_tracking_windows_action").(string)
	ret.Inst.ProgressionTrackingWindowsActionListName = d.Get("progression_tracking_windows_action_list_name").(string)
	//omit uuid
	ret.Inst.WindowRcvdMax = d.Get("window_rcvd_max").(int)
	ret.Inst.WindowRcvdMin = d.Get("window_rcvd_min").(int)
	ret.Inst.WindowRcvdSentRatioMax = d.Get("window_rcvd_sent_ratio_max").(int)
	ret.Inst.WindowRcvdSentRatioMin = d.Get("window_rcvd_sent_ratio_min").(int)
	ret.Inst.WindowSentMax = d.Get("window_sent_max").(int)
	ret.Inst.WindowSentMin = d.Get("window_sent_min").(int)
	ret.Inst.WindowViolation = d.Get("window_violation").(int)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
