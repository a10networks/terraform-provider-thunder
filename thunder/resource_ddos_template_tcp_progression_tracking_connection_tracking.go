package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateTcpProgressionTrackingConnectionTracking() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_tcp_progression_tracking_connection_tracking`: Configure and enable TCP Progression Tracking per Connection\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateTcpProgressionTrackingConnectionTrackingCreate,
		UpdateContext: resourceDdosTemplateTcpProgressionTrackingConnectionTrackingUpdate,
		ReadContext:   resourceDdosTemplateTcpProgressionTrackingConnectionTrackingRead,
		DeleteContext: resourceDdosTemplateTcpProgressionTrackingConnectionTrackingDelete,

		Schema: map[string]*schema.Schema{
			"conn_duration_max": {
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum duration time (in unit of 100ms, up to 24 hours)",
			},
			"conn_duration_min": {
				Type: schema.TypeInt, Optional: true, Description: "Set the minimum duration time (in unit of 100ms, up to 24 hours)",
			},
			"conn_rcvd_max": {
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum total received byte",
			},
			"conn_rcvd_min": {
				Type: schema.TypeInt, Optional: true, Description: "Set the minimum total received byte",
			},
			"conn_rcvd_sent_ratio_max": {
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum received to sent ratio (in unit of 0.1% [1:1000])",
			},
			"conn_rcvd_sent_ratio_min": {
				Type: schema.TypeInt, Optional: true, Description: "Set the minimum received to sent ratio (in unit of 0.1% [1:1000])",
			},
			"conn_sent_max": {
				Type: schema.TypeInt, Optional: true, Description: "Set the maximum total sent byte",
			},
			"conn_sent_min": {
				Type: schema.TypeInt, Optional: true, Description: "Set the minimum total sent byte",
			},
			"conn_violation": {
				Type: schema.TypeInt, Optional: true, Description: "Set the violation threshold",
			},
			"progression_tracking_conn_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for progression tracking violation exceed (Default); 'blacklist-src': Blacklist-src for progression tracking violation exceed;",
			},
			"progression_tracking_conn_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take when progression tracking violation exceed",
			},
			"progression_tracking_conn_enabled": {
				Type: schema.TypeString, Required: true, Description: "'enable-check': Enable General Progression Tracking per Connection;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceDdosTemplateTcpProgressionTrackingConnectionTrackingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingConnectionTrackingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingConnectionTracking(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateTcpProgressionTrackingConnectionTrackingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateTcpProgressionTrackingConnectionTrackingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingConnectionTrackingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingConnectionTracking(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateTcpProgressionTrackingConnectionTrackingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateTcpProgressionTrackingConnectionTrackingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingConnectionTrackingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingConnectionTracking(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateTcpProgressionTrackingConnectionTrackingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingConnectionTrackingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingConnectionTracking(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTemplateTcpProgressionTrackingConnectionTracking(d *schema.ResourceData) edpt.DdosTemplateTcpProgressionTrackingConnectionTracking {
	var ret edpt.DdosTemplateTcpProgressionTrackingConnectionTracking
	ret.Inst.ConnDurationMax = d.Get("conn_duration_max").(int)
	ret.Inst.ConnDurationMin = d.Get("conn_duration_min").(int)
	ret.Inst.ConnRcvdMax = d.Get("conn_rcvd_max").(int)
	ret.Inst.ConnRcvdMin = d.Get("conn_rcvd_min").(int)
	ret.Inst.ConnRcvdSentRatioMax = d.Get("conn_rcvd_sent_ratio_max").(int)
	ret.Inst.ConnRcvdSentRatioMin = d.Get("conn_rcvd_sent_ratio_min").(int)
	ret.Inst.ConnSentMax = d.Get("conn_sent_max").(int)
	ret.Inst.ConnSentMin = d.Get("conn_sent_min").(int)
	ret.Inst.ConnViolation = d.Get("conn_violation").(int)
	ret.Inst.ProgressionTrackingConnAction = d.Get("progression_tracking_conn_action").(string)
	ret.Inst.ProgressionTrackingConnActionListName = d.Get("progression_tracking_conn_action_list_name").(string)
	ret.Inst.ProgressionTrackingConnEnabled = d.Get("progression_tracking_conn_enabled").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
