package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_dns_vport_tmpl_trigger_stats_rate`: Configure stats to triggers packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRateCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRateUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRateRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRateDelete,

		Schema: map[string]*schema.Schema{
			"dns_filter_class_any_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class ANY Drop",
			},
			"dns_filter_class_chaos_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class CHAOS Drop",
			},
			"dns_filter_class_hesiod_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class HESIOD Drop",
			},
			"dns_filter_class_internet_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class INTERNET Drop",
			},
			"dns_filter_class_none_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class NONE Drop",
			},
			"dns_filter_class_others_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Class OTHER Drop",
			},
			"dns_filter_type_a_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type A Drop",
			},
			"dns_filter_type_aaaa_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type AAAA Drop",
			},
			"dns_filter_type_any_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type Any Drop",
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
			"dns_filter_type_others_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type OTHERS Drop",
			},
			"dns_filter_type_ptr_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type PTR Drop",
			},
			"dns_filter_type_soa_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type SOA Drop",
			},
			"dns_filter_type_srv_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type SRV Drop",
			},
			"dns_filter_type_txt_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters DNS Filter Type TXT Drop",
			},
			"dns_rpz_action_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS RPZ Action Drop",
			},
			"dnsrrl_bad_fqdn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS Response-Rate-Limiting Bad FQDN",
			},
			"dnsrrl_total_dropped": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dns rrl drop",
			},
			"duration": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
			},
			"gslb_query_bad": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for gslb query bad",
			},
			"gslb_response_bad": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for gslb response bad",
			},
			"rcode_notimpl_receive": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for response rcode type error receive",
			},
			"rcode_notimpl_response": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for rcode type error response",
			},
			"threshold_exceeded_by": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
			},
			"total_dns_filter_class_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters Total DNS Filter Class Drop",
			},
			"total_dns_filter_type_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for counters Total DNS Filter Type Drop",
			},
			"total_filter_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for query filter drop",
			},
			"total_max_query_len_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for query too long drop",
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
func resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRate(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRate {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsRate
	ret.Inst.Dns_filter_class_any_drop = d.Get("dns_filter_class_any_drop").(int)
	ret.Inst.Dns_filter_class_chaos_drop = d.Get("dns_filter_class_chaos_drop").(int)
	ret.Inst.Dns_filter_class_hesiod_drop = d.Get("dns_filter_class_hesiod_drop").(int)
	ret.Inst.Dns_filter_class_internet_drop = d.Get("dns_filter_class_internet_drop").(int)
	ret.Inst.Dns_filter_class_none_drop = d.Get("dns_filter_class_none_drop").(int)
	ret.Inst.Dns_filter_class_others_drop = d.Get("dns_filter_class_others_drop").(int)
	ret.Inst.Dns_filter_type_a_drop = d.Get("dns_filter_type_a_drop").(int)
	ret.Inst.Dns_filter_type_aaaa_drop = d.Get("dns_filter_type_aaaa_drop").(int)
	ret.Inst.Dns_filter_type_any_drop = d.Get("dns_filter_type_any_drop").(int)
	ret.Inst.Dns_filter_type_cname_drop = d.Get("dns_filter_type_cname_drop").(int)
	ret.Inst.Dns_filter_type_mx_drop = d.Get("dns_filter_type_mx_drop").(int)
	ret.Inst.Dns_filter_type_ns_drop = d.Get("dns_filter_type_ns_drop").(int)
	ret.Inst.Dns_filter_type_others_drop = d.Get("dns_filter_type_others_drop").(int)
	ret.Inst.Dns_filter_type_ptr_drop = d.Get("dns_filter_type_ptr_drop").(int)
	ret.Inst.Dns_filter_type_soa_drop = d.Get("dns_filter_type_soa_drop").(int)
	ret.Inst.Dns_filter_type_srv_drop = d.Get("dns_filter_type_srv_drop").(int)
	ret.Inst.Dns_filter_type_txt_drop = d.Get("dns_filter_type_txt_drop").(int)
	ret.Inst.Dns_rpz_action_drop = d.Get("dns_rpz_action_drop").(int)
	ret.Inst.Dnsrrl_bad_fqdn = d.Get("dnsrrl_bad_fqdn").(int)
	ret.Inst.Dnsrrl_total_dropped = d.Get("dnsrrl_total_dropped").(int)
	ret.Inst.Duration = d.Get("duration").(int)
	ret.Inst.Gslb_query_bad = d.Get("gslb_query_bad").(int)
	ret.Inst.Gslb_response_bad = d.Get("gslb_response_bad").(int)
	ret.Inst.Rcode_notimpl_receive = d.Get("rcode_notimpl_receive").(int)
	ret.Inst.Rcode_notimpl_response = d.Get("rcode_notimpl_response").(int)
	ret.Inst.ThresholdExceededBy = d.Get("threshold_exceeded_by").(int)
	ret.Inst.Total_dns_filter_class_drop = d.Get("total_dns_filter_class_drop").(int)
	ret.Inst.Total_dns_filter_type_drop = d.Get("total_dns_filter_type_drop").(int)
	ret.Inst.Total_filter_drop = d.Get("total_filter_drop").(int)
	ret.Inst.Total_max_query_len_drop = d.Get("total_max_query_len_drop").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
