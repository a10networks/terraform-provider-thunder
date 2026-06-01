package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_alg_pptp`: Configure triggers for fw.alg.pptp object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpDelete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"call_req_pns_call_id_mismatch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Call ID Mismatch on Call Request",
						},
						"call_reply_pns_call_id_mismatch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Call ID Mismatch on Call Reply",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"trigger_stats_rate": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threshold_exceeded_by": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
						},
						"duration": {
							Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
						},
						"call_req_pns_call_id_mismatch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Call ID Mismatch on Call Request",
						},
						"call_reply_pns_call_id_mismatch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Call ID Mismatch on Call Reply",
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
			"template_name": {
				Type: schema.TypeString, Required: true, Description: "Template_name",
			},
		},
	}
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsInc2123(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsInc2123 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsInc2123
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CallReqPnsCallIdMismatch = in["call_req_pns_call_id_mismatch"].(int)
		ret.CallReplyPnsCallIdMismatch = in["call_reply_pns_call_id_mismatch"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsRate2124(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsRate2124 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsRate2124
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.CallReqPnsCallIdMismatch = in["call_req_pns_call_id_mismatch"].(int)
		ret.CallReplyPnsCallIdMismatch = in["call_reply_pns_call_id_mismatch"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptp(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptp {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptp
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsInc2123(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsRate2124(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Template_name = d.Get("template_name").(string)
	return ret
}
