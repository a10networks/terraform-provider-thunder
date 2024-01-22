package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_interface_ethernet_tmpl`: Configure template for interface.ethernet\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplDelete,

		Schema: map[string]*schema.Schema{
			"capture_config": {
				Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
			},
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"input_errors": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Input errors",
						},
						"crc": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRC",
						},
						"runts": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Runts",
						},
						"giants": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Giants",
						},
						"output_errors": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Output errors",
						},
						"collisions": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Collisions",
						},
						"giants_output": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Output Giants",
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
						"input_errors": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Input errors",
						},
						"crc": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRC",
						},
						"runts": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Runts",
						},
						"giants": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Giants",
						},
						"output_errors": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Output errors",
						},
						"collisions": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Collisions",
						},
						"giants_output": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Output Giants",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"trigger_stats_severity": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
						},
						"error_alert": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
						},
						"error_warning": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
						},
						"error_critical": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
						},
						"drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
						},
						"drop_alert": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
						},
						"drop_warning": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
						},
						"drop_critical": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsInc2693(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsInc2693 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsInc2693
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Input_errors = in["input_errors"].(int)
		ret.Crc = in["crc"].(int)
		ret.Runts = in["runts"].(int)
		ret.Giants = in["giants"].(int)
		ret.Output_errors = in["output_errors"].(int)
		ret.Collisions = in["collisions"].(int)
		ret.Giants_output = in["giants_output"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsRate2694(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsRate2694 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsRate2694
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Input_errors = in["input_errors"].(int)
		ret.Crc = in["crc"].(int)
		ret.Runts = in["runts"].(int)
		ret.Giants = in["giants"].(int)
		ret.Output_errors = in["output_errors"].(int)
		ret.Collisions = in["collisions"].(int)
		ret.Giants_output = in["giants_output"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsSeverity2695(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsSeverity2695 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsSeverity2695
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsInc2693(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsRate2694(d.Get("trigger_stats_rate").([]interface{}))
	ret.Inst.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsSeverity2695(d.Get("trigger_stats_severity").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
