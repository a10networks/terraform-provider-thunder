package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmpl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_dns_vport_tmpl`: Configure template for counter.dns_vport\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplDelete,

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
						"dnsrrl_total_dropped": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dns rrl drop",
						},
						"total_filter_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for query filter drop",
						},
						"total_max_query_len_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for query too long drop",
						},
						"rcode_notimpl_receive": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for response rcode type error receive",
						},
						"rcode_notimpl_response": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for rcode type error response",
						},
						"gslb_query_bad": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for gslb query bad",
						},
						"gslb_response_bad": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for gslb response bad",
						},
						"total_dns_filter_type_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters Total DNS Filter Type Drop",
						},
						"total_dns_filter_class_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters Total DNS Filter Class Drop",
						},
						"dns_filter_type_a_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type A Drop",
						},
						"dns_filter_type_aaaa_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type AAAA Drop",
						},
						"dns_filter_type_cname_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type CNAME Drop",
						},
						"dns_filter_type_mx_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type MX Drop",
						},
						"dns_filter_type_ns_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type NS Drop",
						},
						"dns_filter_type_srv_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type SRV Drop",
						},
						"dns_filter_type_ptr_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type PTR Drop",
						},
						"dns_filter_type_soa_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type SOA Drop",
						},
						"dns_filter_type_txt_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type TXT Drop",
						},
						"dns_filter_type_any_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type Any Drop",
						},
						"dns_filter_type_others_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type OTHERS Drop",
						},
						"dns_filter_class_internet_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class INTERNET Drop",
						},
						"dns_filter_class_chaos_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class CHAOS Drop",
						},
						"dns_filter_class_hesiod_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class HESIOD Drop",
						},
						"dns_filter_class_none_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class NONE Drop",
						},
						"dns_filter_class_any_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class ANY Drop",
						},
						"dns_filter_class_others_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class OTHER Drop",
						},
						"dns_rpz_action_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS RPZ Action Drop",
						},
						"dnsrrl_bad_fqdn": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS Response-Rate-Limiting Bad FQDN",
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
						"dnsrrl_total_dropped": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dns rrl drop",
						},
						"total_filter_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for query filter drop",
						},
						"total_max_query_len_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for query too long drop",
						},
						"rcode_notimpl_receive": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for response rcode type error receive",
						},
						"rcode_notimpl_response": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for rcode type error response",
						},
						"gslb_query_bad": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for gslb query bad",
						},
						"gslb_response_bad": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for gslb response bad",
						},
						"total_dns_filter_type_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters Total DNS Filter Type Drop",
						},
						"total_dns_filter_class_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters Total DNS Filter Class Drop",
						},
						"dns_filter_type_a_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type A Drop",
						},
						"dns_filter_type_aaaa_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type AAAA Drop",
						},
						"dns_filter_type_cname_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type CNAME Drop",
						},
						"dns_filter_type_mx_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type MX Drop",
						},
						"dns_filter_type_ns_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type NS Drop",
						},
						"dns_filter_type_srv_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type SRV Drop",
						},
						"dns_filter_type_ptr_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type PTR Drop",
						},
						"dns_filter_type_soa_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type SOA Drop",
						},
						"dns_filter_type_txt_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type TXT Drop",
						},
						"dns_filter_type_any_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type Any Drop",
						},
						"dns_filter_type_others_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type OTHERS Drop",
						},
						"dns_filter_class_internet_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class INTERNET Drop",
						},
						"dns_filter_class_chaos_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class CHAOS Drop",
						},
						"dns_filter_class_hesiod_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class HESIOD Drop",
						},
						"dns_filter_class_none_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class NONE Drop",
						},
						"dns_filter_class_any_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class ANY Drop",
						},
						"dns_filter_class_others_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class OTHER Drop",
						},
						"dns_rpz_action_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS RPZ Action Drop",
						},
						"dnsrrl_bad_fqdn": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS Response-Rate-Limiting Bad FQDN",
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
func resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesDnsVportTmpl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesDnsVportTmpl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesDnsVportTmpl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesDnsVportTmpl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsInc2678(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsInc2678 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsInc2678
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dnsrrl_total_dropped = in["dnsrrl_total_dropped"].(int)
		ret.Total_filter_drop = in["total_filter_drop"].(int)
		ret.Total_max_query_len_drop = in["total_max_query_len_drop"].(int)
		ret.Rcode_notimpl_receive = in["rcode_notimpl_receive"].(int)
		ret.Rcode_notimpl_response = in["rcode_notimpl_response"].(int)
		ret.Gslb_query_bad = in["gslb_query_bad"].(int)
		ret.Gslb_response_bad = in["gslb_response_bad"].(int)
		ret.Total_dns_filter_type_drop = in["total_dns_filter_type_drop"].(int)
		ret.Total_dns_filter_class_drop = in["total_dns_filter_class_drop"].(int)
		ret.Dns_filter_type_a_drop = in["dns_filter_type_a_drop"].(int)
		ret.Dns_filter_type_aaaa_drop = in["dns_filter_type_aaaa_drop"].(int)
		ret.Dns_filter_type_cname_drop = in["dns_filter_type_cname_drop"].(int)
		ret.Dns_filter_type_mx_drop = in["dns_filter_type_mx_drop"].(int)
		ret.Dns_filter_type_ns_drop = in["dns_filter_type_ns_drop"].(int)
		ret.Dns_filter_type_srv_drop = in["dns_filter_type_srv_drop"].(int)
		ret.Dns_filter_type_ptr_drop = in["dns_filter_type_ptr_drop"].(int)
		ret.Dns_filter_type_soa_drop = in["dns_filter_type_soa_drop"].(int)
		ret.Dns_filter_type_txt_drop = in["dns_filter_type_txt_drop"].(int)
		ret.Dns_filter_type_any_drop = in["dns_filter_type_any_drop"].(int)
		ret.Dns_filter_type_others_drop = in["dns_filter_type_others_drop"].(int)
		ret.Dns_filter_class_internet_drop = in["dns_filter_class_internet_drop"].(int)
		ret.Dns_filter_class_chaos_drop = in["dns_filter_class_chaos_drop"].(int)
		ret.Dns_filter_class_hesiod_drop = in["dns_filter_class_hesiod_drop"].(int)
		ret.Dns_filter_class_none_drop = in["dns_filter_class_none_drop"].(int)
		ret.Dns_filter_class_any_drop = in["dns_filter_class_any_drop"].(int)
		ret.Dns_filter_class_others_drop = in["dns_filter_class_others_drop"].(int)
		ret.Dns_rpz_action_drop = in["dns_rpz_action_drop"].(int)
		ret.Dnsrrl_bad_fqdn = in["dnsrrl_bad_fqdn"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRate2679(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRate2679 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRate2679
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Dnsrrl_total_dropped = in["dnsrrl_total_dropped"].(int)
		ret.Total_filter_drop = in["total_filter_drop"].(int)
		ret.Total_max_query_len_drop = in["total_max_query_len_drop"].(int)
		ret.Rcode_notimpl_receive = in["rcode_notimpl_receive"].(int)
		ret.Rcode_notimpl_response = in["rcode_notimpl_response"].(int)
		ret.Gslb_query_bad = in["gslb_query_bad"].(int)
		ret.Gslb_response_bad = in["gslb_response_bad"].(int)
		ret.Total_dns_filter_type_drop = in["total_dns_filter_type_drop"].(int)
		ret.Total_dns_filter_class_drop = in["total_dns_filter_class_drop"].(int)
		ret.Dns_filter_type_a_drop = in["dns_filter_type_a_drop"].(int)
		ret.Dns_filter_type_aaaa_drop = in["dns_filter_type_aaaa_drop"].(int)
		ret.Dns_filter_type_cname_drop = in["dns_filter_type_cname_drop"].(int)
		ret.Dns_filter_type_mx_drop = in["dns_filter_type_mx_drop"].(int)
		ret.Dns_filter_type_ns_drop = in["dns_filter_type_ns_drop"].(int)
		ret.Dns_filter_type_srv_drop = in["dns_filter_type_srv_drop"].(int)
		ret.Dns_filter_type_ptr_drop = in["dns_filter_type_ptr_drop"].(int)
		ret.Dns_filter_type_soa_drop = in["dns_filter_type_soa_drop"].(int)
		ret.Dns_filter_type_txt_drop = in["dns_filter_type_txt_drop"].(int)
		ret.Dns_filter_type_any_drop = in["dns_filter_type_any_drop"].(int)
		ret.Dns_filter_type_others_drop = in["dns_filter_type_others_drop"].(int)
		ret.Dns_filter_class_internet_drop = in["dns_filter_class_internet_drop"].(int)
		ret.Dns_filter_class_chaos_drop = in["dns_filter_class_chaos_drop"].(int)
		ret.Dns_filter_class_hesiod_drop = in["dns_filter_class_hesiod_drop"].(int)
		ret.Dns_filter_class_none_drop = in["dns_filter_class_none_drop"].(int)
		ret.Dns_filter_class_any_drop = in["dns_filter_class_any_drop"].(int)
		ret.Dns_filter_class_others_drop = in["dns_filter_class_others_drop"].(int)
		ret.Dns_rpz_action_drop = in["dns_rpz_action_drop"].(int)
		ret.Dnsrrl_bad_fqdn = in["dnsrrl_bad_fqdn"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsSeverity2680(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsSeverity2680 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsSeverity2680
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

func dataToEndpointVisibilityPacketCaptureObjectTemplatesDnsVportTmpl(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesDnsVportTmpl {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesDnsVportTmpl
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsInc2678(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRate2679(d.Get("trigger_stats_rate").([]interface{}))
	ret.Inst.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsSeverity2680(d.Get("trigger_stats_severity").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
