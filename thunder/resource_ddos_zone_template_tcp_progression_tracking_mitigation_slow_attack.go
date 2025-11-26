package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_tcp_progression_tracking_mitigation_slow_attack`: Configure and enable TCP progression Tracking Mitigation for slow attack (identify slow attacker)\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackCreate,
		UpdateContext: resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackUpdate,
		ReadContext:   resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackRead,
		DeleteContext: resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackDelete,

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
			"slow_attacker_identification": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_identification": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Progression tracking will identify slow attacker and blacklist it based on the config value",
						},
						"active_connection": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Set the minimum tracking active connection to start identifying slow attacker, default value is 3",
						},
						"bad_connection": {
							Type: schema.TypeInt, Optional: true, Default: 75, Description: "Set the maximum percentage of slow connection (per source), default value is 75",
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
			"tcp_name": {
				Type: schema.TypeString, Required: true, Description: "Tcp_name",
			},
		},
	}
}
func resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification373(d []interface{}) edpt.DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification373 {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification373
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EnableIdentification = in["enable_identification"].(int)
		ret.ActiveConnection = in["active_connection"].(int)
		ret.BadConnection = in["bad_connection"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack(d *schema.ResourceData) edpt.DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack {
	var ret edpt.DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttack
	ret.Inst.InitRequestMaxTime = d.Get("init_request_max_time").(int)
	ret.Inst.InitResponseMaxTime = d.Get("init_response_max_time").(int)
	ret.Inst.ProgressionTrackingSlowAction = d.Get("progression_tracking_slow_action").(string)
	ret.Inst.ProgressionTrackingSlowActionListName = d.Get("progression_tracking_slow_action_list_name").(string)
	ret.Inst.ResponsePktRateMax = d.Get("response_pkt_rate_max").(int)
	ret.Inst.SlowAttackerIdentification = getObjectDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification373(d.Get("slow_attacker_identification").([]interface{}))
	//omit uuid
	ret.Inst.Tcp_name = d.Get("tcp_name").(string)
	return ret
}
