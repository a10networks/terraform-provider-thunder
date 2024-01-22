package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_cgnv6_encap_domain_tmpl`: Configure template for cgnv6.map.encapsulation.domain\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplDelete,

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
						"inbound_addr_port_validation_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv4 Destination Address Port Validation Failed",
						},
						"inbound_rev_lookup_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv4 Reverse Route Lookup Failed",
						},
						"inbound_dest_unreachable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv6 Destination Address Unreachable",
						},
						"outbound_addr_validation_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv6 Source Address Validation Failed",
						},
						"outbound_rev_lookup_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv6 Reverse Route Lookup Failed",
						},
						"outbound_dest_unreachable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv4 Destination Address Unreachable",
						},
						"packet_mtu_exceeded": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Exceeded MTU",
						},
						"interface_not_configured": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Interfaces not Configured Dropped",
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
						"inbound_addr_port_validation_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv4 Destination Address Port Validation Failed",
						},
						"inbound_rev_lookup_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv4 Reverse Route Lookup Failed",
						},
						"inbound_dest_unreachable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inbound IPv6 Destination Address Unreachable",
						},
						"outbound_addr_validation_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv6 Source Address Validation Failed",
						},
						"outbound_rev_lookup_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv6 Reverse Route Lookup Failed",
						},
						"outbound_dest_unreachable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Outbound IPv4 Destination Address Unreachable",
						},
						"packet_mtu_exceeded": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Exceeded MTU",
						},
						"interface_not_configured": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Interfaces not Configured Dropped",
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
func resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsInc2669(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsInc2669 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsInc2669
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Inbound_addr_port_validation_failed = in["inbound_addr_port_validation_failed"].(int)
		ret.Inbound_rev_lookup_failed = in["inbound_rev_lookup_failed"].(int)
		ret.Inbound_dest_unreachable = in["inbound_dest_unreachable"].(int)
		ret.Outbound_addr_validation_failed = in["outbound_addr_validation_failed"].(int)
		ret.Outbound_rev_lookup_failed = in["outbound_rev_lookup_failed"].(int)
		ret.Outbound_dest_unreachable = in["outbound_dest_unreachable"].(int)
		ret.Packet_mtu_exceeded = in["packet_mtu_exceeded"].(int)
		ret.Interface_not_configured = in["interface_not_configured"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsRate2670(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsRate2670 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsRate2670
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Inbound_addr_port_validation_failed = in["inbound_addr_port_validation_failed"].(int)
		ret.Inbound_rev_lookup_failed = in["inbound_rev_lookup_failed"].(int)
		ret.Inbound_dest_unreachable = in["inbound_dest_unreachable"].(int)
		ret.Outbound_addr_validation_failed = in["outbound_addr_validation_failed"].(int)
		ret.Outbound_rev_lookup_failed = in["outbound_rev_lookup_failed"].(int)
		ret.Outbound_dest_unreachable = in["outbound_dest_unreachable"].(int)
		ret.Packet_mtu_exceeded = in["packet_mtu_exceeded"].(int)
		ret.Interface_not_configured = in["interface_not_configured"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsSeverity2671(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsSeverity2671 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsSeverity2671
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

func dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsInc2669(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsRate2670(d.Get("trigger_stats_rate").([]interface{}))
	ret.Inst.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsSeverity2671(d.Get("trigger_stats_severity").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
