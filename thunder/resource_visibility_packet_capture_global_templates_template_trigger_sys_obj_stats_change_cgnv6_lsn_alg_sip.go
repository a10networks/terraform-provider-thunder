package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_sip`: Configure triggers for cgnv6.lsn.alg.sip object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipDelete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"method_unknown": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP Method UNKNOWN",
						},
						"parse_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP Message Parse Error",
						},
						"tcp_out_of_order_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Out-of-Order Drop",
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
						"method_unknown": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP Method UNKNOWN",
						},
						"parse_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP Message Parse Error",
						},
						"tcp_out_of_order_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Out-of-Order Drop",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsInc2113(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsInc2113 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsInc2113
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MethodUnknown = in["method_unknown"].(int)
		ret.ParseError = in["parse_error"].(int)
		ret.TcpOutOfOrderDrop = in["tcp_out_of_order_drop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsRate2114(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsRate2114 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsRate2114
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.MethodUnknown = in["method_unknown"].(int)
		ret.ParseError = in["parse_error"].(int)
		ret.TcpOutOfOrderDrop = in["tcp_out_of_order_drop"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSip(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSip {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSip
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsInc2113(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsRate2114(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Template_name = d.Get("template_name").(string)
	return ret
}
