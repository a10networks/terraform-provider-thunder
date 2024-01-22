package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_ddos_proc`: Configure triggers for cgnv6.ddos-protection object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcDelete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l3_entry_match_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Entry match drop",
						},
						"l3_entry_match_drop_hw": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 HW entry match drop",
						},
						"l3_entry_drop_max_hw_exceeded": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Entry Drop due to HW Limit Exceeded",
						},
						"l4_entry_match_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 Entry match drop",
						},
						"l4_entry_match_drop_hw": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 HW Entry match drop",
						},
						"l4_entry_drop_max_hw_exceeded": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 Entry Drop due to HW Limit Exceeded",
						},
						"l4_entry_list_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 Entry list alloc failures",
						},
						"ip_node_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Node alloc failures",
						},
						"ip_port_block_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port block alloc failure",
						},
						"ip_other_block_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Other block alloc failure",
						},
						"l3_entry_add_to_bgp_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Entry BGP add failures",
						},
						"l3_entry_remove_from_bgp_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 entry BGP remove failures",
						},
						"l3_entry_add_to_hw_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 entry HW add failure",
						},
						"syn_cookie_verification_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SYN cookie verification failed",
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
						"l3_entry_match_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Entry match drop",
						},
						"l3_entry_match_drop_hw": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 HW entry match drop",
						},
						"l3_entry_drop_max_hw_exceeded": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Entry Drop due to HW Limit Exceeded",
						},
						"l4_entry_match_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 Entry match drop",
						},
						"l4_entry_match_drop_hw": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 HW Entry match drop",
						},
						"l4_entry_drop_max_hw_exceeded": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 Entry Drop due to HW Limit Exceeded",
						},
						"l4_entry_list_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 Entry list alloc failures",
						},
						"ip_node_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Node alloc failures",
						},
						"ip_port_block_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port block alloc failure",
						},
						"ip_other_block_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Other block alloc failure",
						},
						"l3_entry_add_to_bgp_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Entry BGP add failures",
						},
						"l3_entry_remove_from_bgp_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 entry BGP remove failures",
						},
						"l3_entry_add_to_hw_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 entry HW add failure",
						},
						"syn_cookie_verification_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SYN cookie verification failed",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsInc1955(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsInc1955 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsInc1955
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L3_entry_match_drop = in["l3_entry_match_drop"].(int)
		ret.L3_entry_match_drop_hw = in["l3_entry_match_drop_hw"].(int)
		ret.L3_entry_drop_max_hw_exceeded = in["l3_entry_drop_max_hw_exceeded"].(int)
		ret.L4_entry_match_drop = in["l4_entry_match_drop"].(int)
		ret.L4_entry_match_drop_hw = in["l4_entry_match_drop_hw"].(int)
		ret.L4_entry_drop_max_hw_exceeded = in["l4_entry_drop_max_hw_exceeded"].(int)
		ret.L4_entry_list_alloc_failure = in["l4_entry_list_alloc_failure"].(int)
		ret.Ip_node_alloc_failure = in["ip_node_alloc_failure"].(int)
		ret.Ip_port_block_alloc_failure = in["ip_port_block_alloc_failure"].(int)
		ret.Ip_other_block_alloc_failure = in["ip_other_block_alloc_failure"].(int)
		ret.L3_entry_add_to_bgp_failure = in["l3_entry_add_to_bgp_failure"].(int)
		ret.L3_entry_remove_from_bgp_failure = in["l3_entry_remove_from_bgp_failure"].(int)
		ret.L3_entry_add_to_hw_failure = in["l3_entry_add_to_hw_failure"].(int)
		ret.Syn_cookie_verification_failed = in["syn_cookie_verification_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsRate1956(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsRate1956 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsRate1956
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.L3_entry_match_drop = in["l3_entry_match_drop"].(int)
		ret.L3_entry_match_drop_hw = in["l3_entry_match_drop_hw"].(int)
		ret.L3_entry_drop_max_hw_exceeded = in["l3_entry_drop_max_hw_exceeded"].(int)
		ret.L4_entry_match_drop = in["l4_entry_match_drop"].(int)
		ret.L4_entry_match_drop_hw = in["l4_entry_match_drop_hw"].(int)
		ret.L4_entry_drop_max_hw_exceeded = in["l4_entry_drop_max_hw_exceeded"].(int)
		ret.L4_entry_list_alloc_failure = in["l4_entry_list_alloc_failure"].(int)
		ret.Ip_node_alloc_failure = in["ip_node_alloc_failure"].(int)
		ret.Ip_port_block_alloc_failure = in["ip_port_block_alloc_failure"].(int)
		ret.Ip_other_block_alloc_failure = in["ip_other_block_alloc_failure"].(int)
		ret.L3_entry_add_to_bgp_failure = in["l3_entry_add_to_bgp_failure"].(int)
		ret.L3_entry_remove_from_bgp_failure = in["l3_entry_remove_from_bgp_failure"].(int)
		ret.L3_entry_add_to_hw_failure = in["l3_entry_add_to_hw_failure"].(int)
		ret.Syn_cookie_verification_failed = in["syn_cookie_verification_failed"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProc(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProc {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProc
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsInc1955(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsRate1956(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
