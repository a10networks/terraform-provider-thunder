package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_l4`: Configure triggers for slb.l4 object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4Create,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4Update,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4Read,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4Delete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"syncookiessentfailed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP SYN cookie snt fail",
						},
						"svrselfail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server sel failure",
						},
						"snat_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT failure",
						},
						"snat_no_fwd_route": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT no fwd route",
						},
						"snat_no_rev_route": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT no rev route",
						},
						"snat_icmp_error_process": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT ICMP Process",
						},
						"snat_icmp_no_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT ICMP No Match",
						},
						"smart_nat_id_mismatch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Auto NAT id mismatch",
						},
						"syncookiescheckfailed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP SYN cookie failed",
						},
						"connlimit_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Conn Limit drops",
						},
						"conn_rate_limit_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Conn rate limit drops",
						},
						"conn_rate_limit_reset": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Conn rate limit resets",
						},
						"dns_policy_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS Policy Drop",
						},
						"no_resourse_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No resource drop",
						},
						"bw_rate_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for BW-Limit Exceed drop",
						},
						"l4_cps_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 CPS exceed drop",
						},
						"nat_cps_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT CPS exceed drop",
						},
						"l7_cps_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L7 CPS exceed drop",
						},
						"ssl_cps_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SSL CPS exceed drop",
						},
						"ssl_tpt_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SSL TPT exceed drop",
						},
						"concurrent_conn_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3V Conn Limit Drop",
						},
						"svr_syn_handshake_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 server handshake fail",
						},
						"synattack": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 SYN attack",
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
						"syncookiessentfailed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP SYN cookie snt fail",
						},
						"svrselfail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server sel failure",
						},
						"snat_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT failure",
						},
						"snat_no_fwd_route": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT no fwd route",
						},
						"snat_no_rev_route": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT no rev route",
						},
						"snat_icmp_error_process": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT ICMP Process",
						},
						"snat_icmp_no_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT ICMP No Match",
						},
						"smart_nat_id_mismatch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Auto NAT id mismatch",
						},
						"syncookiescheckfailed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP SYN cookie failed",
						},
						"connlimit_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Conn Limit drops",
						},
						"conn_rate_limit_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Conn rate limit drops",
						},
						"conn_rate_limit_reset": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Conn rate limit resets",
						},
						"dns_policy_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS Policy Drop",
						},
						"no_resourse_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No resource drop",
						},
						"bw_rate_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for BW-Limit Exceed drop",
						},
						"l4_cps_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 CPS exceed drop",
						},
						"nat_cps_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT CPS exceed drop",
						},
						"l7_cps_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L7 CPS exceed drop",
						},
						"ssl_cps_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SSL CPS exceed drop",
						},
						"ssl_tpt_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SSL TPT exceed drop",
						},
						"concurrent_conn_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3V Conn Limit Drop",
						},
						"svr_syn_handshake_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 server handshake fail",
						},
						"synattack": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 SYN attack",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4Read(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4Read(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsInc2045(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsInc2045 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsInc2045
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Syncookiessentfailed = in["syncookiessentfailed"].(int)
		ret.Svrselfail = in["svrselfail"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Snat_no_fwd_route = in["snat_no_fwd_route"].(int)
		ret.Snat_no_rev_route = in["snat_no_rev_route"].(int)
		ret.Snat_icmp_error_process = in["snat_icmp_error_process"].(int)
		ret.Snat_icmp_no_match = in["snat_icmp_no_match"].(int)
		ret.Smart_nat_id_mismatch = in["smart_nat_id_mismatch"].(int)
		ret.Syncookiescheckfailed = in["syncookiescheckfailed"].(int)
		ret.Connlimit_drop = in["connlimit_drop"].(int)
		ret.Conn_rate_limit_drop = in["conn_rate_limit_drop"].(int)
		ret.Conn_rate_limit_reset = in["conn_rate_limit_reset"].(int)
		ret.Dns_policy_drop = in["dns_policy_drop"].(int)
		ret.No_resourse_drop = in["no_resourse_drop"].(int)
		ret.Bw_rate_limit_exceed = in["bw_rate_limit_exceed"].(int)
		ret.L4_cps_exceed = in["l4_cps_exceed"].(int)
		ret.Nat_cps_exceed = in["nat_cps_exceed"].(int)
		ret.L7_cps_exceed = in["l7_cps_exceed"].(int)
		ret.Ssl_cps_exceed = in["ssl_cps_exceed"].(int)
		ret.Ssl_tpt_exceed = in["ssl_tpt_exceed"].(int)
		ret.Concurrent_conn_exceed = in["concurrent_conn_exceed"].(int)
		ret.Svr_syn_handshake_fail = in["svr_syn_handshake_fail"].(int)
		ret.Synattack = in["synattack"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsRate2046(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsRate2046 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsRate2046
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Syncookiessentfailed = in["syncookiessentfailed"].(int)
		ret.Svrselfail = in["svrselfail"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Snat_no_fwd_route = in["snat_no_fwd_route"].(int)
		ret.Snat_no_rev_route = in["snat_no_rev_route"].(int)
		ret.Snat_icmp_error_process = in["snat_icmp_error_process"].(int)
		ret.Snat_icmp_no_match = in["snat_icmp_no_match"].(int)
		ret.Smart_nat_id_mismatch = in["smart_nat_id_mismatch"].(int)
		ret.Syncookiescheckfailed = in["syncookiescheckfailed"].(int)
		ret.Connlimit_drop = in["connlimit_drop"].(int)
		ret.Conn_rate_limit_drop = in["conn_rate_limit_drop"].(int)
		ret.Conn_rate_limit_reset = in["conn_rate_limit_reset"].(int)
		ret.Dns_policy_drop = in["dns_policy_drop"].(int)
		ret.No_resourse_drop = in["no_resourse_drop"].(int)
		ret.Bw_rate_limit_exceed = in["bw_rate_limit_exceed"].(int)
		ret.L4_cps_exceed = in["l4_cps_exceed"].(int)
		ret.Nat_cps_exceed = in["nat_cps_exceed"].(int)
		ret.L7_cps_exceed = in["l7_cps_exceed"].(int)
		ret.Ssl_cps_exceed = in["ssl_cps_exceed"].(int)
		ret.Ssl_tpt_exceed = in["ssl_tpt_exceed"].(int)
		ret.Concurrent_conn_exceed = in["concurrent_conn_exceed"].(int)
		ret.Svr_syn_handshake_fail = in["svr_syn_handshake_fail"].(int)
		ret.Synattack = in["synattack"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4 {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsInc2045(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsRate2046(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
