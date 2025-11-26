package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_tcp_progression_tracking_mitigation_slow_attack_slow_attacker_identification`: Configure and enable TCP Progression Tracking Identification for Slow Attacker\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentificationCreate,
		UpdateContext: resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentificationUpdate,
		ReadContext:   resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentificationRead,
		DeleteContext: resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentificationDelete,

		Schema: map[string]*schema.Schema{
			"active_connection": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Set the minimum tracking active connection to start identifying slow attacker, default value is 3",
			},
			"bad_connection": {
				Type: schema.TypeInt, Optional: true, Default: 75, Description: "Set the maximum percentage of slow connection (per source), default value is 75",
			},
			"enable_identification": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Progression tracking will identify slow attacker and blacklist it based on the config value",
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
func resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentificationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentificationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentificationRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentificationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentificationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentificationRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentificationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentificationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentificationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentificationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification(d *schema.ResourceData) edpt.DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification {
	var ret edpt.DdosZoneTemplateTcpProgressionTrackingMitigationSlowAttackSlowAttackerIdentification
	ret.Inst.ActiveConnection = d.Get("active_connection").(int)
	ret.Inst.BadConnection = d.Get("bad_connection").(int)
	ret.Inst.EnableIdentification = d.Get("enable_identification").(int)
	//omit uuid
	ret.Inst.Tcp_name = d.Get("tcp_name").(string)
	return ret
}
