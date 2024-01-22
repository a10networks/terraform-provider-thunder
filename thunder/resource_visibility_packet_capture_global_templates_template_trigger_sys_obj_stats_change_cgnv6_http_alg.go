package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlg() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_http_alg`: Configure triggers for cgnv6.http-alg object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgDelete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"radius_requst_dropped": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Query Request Dropped",
						},
						"radius_response_dropped": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Query Response Dropped",
						},
						"out_of_memory_dropped": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out-of-Memory Dropped",
						},
						"queue_len_exceed_dropped": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Queue Length Exceed Dropped",
						},
						"out_of_order_dropped": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Out-of-Order Dropped",
						},
						"header_insertion_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Buff Insertion Failed",
						},
						"header_removal_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Buff Removal Failed",
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
						"radius_requst_dropped": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Query Request Dropped",
						},
						"radius_response_dropped": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Query Response Dropped",
						},
						"out_of_memory_dropped": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out-of-Memory Dropped",
						},
						"queue_len_exceed_dropped": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Queue Length Exceed Dropped",
						},
						"out_of_order_dropped": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Out-of-Order Dropped",
						},
						"header_insertion_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Buff Insertion Failed",
						},
						"header_removal_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Buff Removal Failed",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlg(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlg(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlg(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlg(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsInc1973(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsInc1973 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsInc1973
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RadiusRequstDropped = in["radius_requst_dropped"].(int)
		ret.RadiusResponseDropped = in["radius_response_dropped"].(int)
		ret.OutOfMemoryDropped = in["out_of_memory_dropped"].(int)
		ret.QueueLenExceedDropped = in["queue_len_exceed_dropped"].(int)
		ret.OutOfOrderDropped = in["out_of_order_dropped"].(int)
		ret.HeaderInsertionFailed = in["header_insertion_failed"].(int)
		ret.HeaderRemovalFailed = in["header_removal_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsRate1974(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsRate1974 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsRate1974
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.RadiusRequstDropped = in["radius_requst_dropped"].(int)
		ret.RadiusResponseDropped = in["radius_response_dropped"].(int)
		ret.OutOfMemoryDropped = in["out_of_memory_dropped"].(int)
		ret.QueueLenExceedDropped = in["queue_len_exceed_dropped"].(int)
		ret.OutOfOrderDropped = in["out_of_order_dropped"].(int)
		ret.HeaderInsertionFailed = in["header_insertion_failed"].(int)
		ret.HeaderRemovalFailed = in["header_removal_failed"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlg(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlg {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlg
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsInc1973(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsRate1974(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
