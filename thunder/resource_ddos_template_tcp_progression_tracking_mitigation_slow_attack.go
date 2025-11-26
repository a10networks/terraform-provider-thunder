package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateTcpProgressionTrackingMitigationSlowAttack() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_tcp_progression_tracking_mitigation_slow_attack`: Configure and enable TCP progression Tracking Mitigation for slow attack (identify slow attacker)\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateTcpProgressionTrackingMitigationSlowAttackCreate,
		UpdateContext: resourceDdosTemplateTcpProgressionTrackingMitigationSlowAttackUpdate,
		ReadContext:   resourceDdosTemplateTcpProgressionTrackingMitigationSlowAttackRead,
		DeleteContext: resourceDdosTemplateTcpProgressionTrackingMitigationSlowAttackDelete,

		Schema: map[string]*schema.Schema{
			"init_request_max_time": {
				Type: schema.TypeInt, Optional: true, Description: "Set client query time (in unit of 100ms). Suggested value larger than 30 secs",
			},
			"init_response_max_time": {
				Type: schema.TypeInt, Optional: true, Description: "Set server think time (in unit of 100ms). Suggested value larger than 45 secs",
			},
			"progression_tracking_slow_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for progression tracking violation exceed (Default); 'reset': Reset client connection; 'blacklist-src': Blacklist-src for progression tracking violation exceed;",
			},
			"progression_tracking_slow_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take when progression tracking violation exceed",
			},
			"response_pkt_rate_max": {
				Type: schema.TypeInt, Optional: true, Description: "Set the transferred packets per response",
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
func resourceDdosTemplateTcpProgressionTrackingMitigationSlowAttackCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingMitigationSlowAttackCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingMitigationSlowAttack(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateTcpProgressionTrackingMitigationSlowAttackRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateTcpProgressionTrackingMitigationSlowAttackUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingMitigationSlowAttackUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingMitigationSlowAttack(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateTcpProgressionTrackingMitigationSlowAttackRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateTcpProgressionTrackingMitigationSlowAttackDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingMitigationSlowAttackDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingMitigationSlowAttack(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateTcpProgressionTrackingMitigationSlowAttackRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateTcpProgressionTrackingMitigationSlowAttackRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateTcpProgressionTrackingMitigationSlowAttack(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTemplateTcpProgressionTrackingMitigationSlowAttack(d *schema.ResourceData) edpt.DdosTemplateTcpProgressionTrackingMitigationSlowAttack {
	var ret edpt.DdosTemplateTcpProgressionTrackingMitigationSlowAttack
	ret.Inst.InitRequestMaxTime = d.Get("init_request_max_time").(int)
	ret.Inst.InitResponseMaxTime = d.Get("init_response_max_time").(int)
	ret.Inst.ProgressionTrackingSlowAction = d.Get("progression_tracking_slow_action").(string)
	ret.Inst.ProgressionTrackingSlowActionListName = d.Get("progression_tracking_slow_action_list_name").(string)
	ret.Inst.ResponsePktRateMax = d.Get("response_pkt_rate_max").(int)
	//omit uuid
	ret.Inst.Tcp_name = d.Get("tcp_name").(string)
	return ret
}
