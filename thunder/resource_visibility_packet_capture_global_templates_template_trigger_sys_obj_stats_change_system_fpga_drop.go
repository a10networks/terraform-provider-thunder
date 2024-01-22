package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_fpga_drop`: Configure triggers for system.fpga-drop object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropDelete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mrx_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total MRX Drop",
						},
						"hrx_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total HRX Drop",
						},
						"siz_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Size Drop",
						},
						"fcs_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total FCS Drop",
						},
						"land_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total LAND Attack Drop",
						},
						"empty_frag_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Empty frag Drop",
						},
						"mic_frag_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Micro frag Drop",
						},
						"ipv4_opt_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IPv4 opt Drop",
						},
						"ipv4_frag": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IP frag Drop",
						},
						"bad_ip_hdr_len": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP hdr len Drop",
						},
						"bad_ip_flags_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP Flags Drop",
						},
						"bad_ip_ttl_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP TTL Drop",
						},
						"no_ip_payload_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total No IP Payload Drop",
						},
						"oversize_ip_payload": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Oversize IP PL Drop",
						},
						"bad_ip_payload_len": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP PL len Drop",
						},
						"bad_ip_frag_offset": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP frag off Drop",
						},
						"bad_ip_chksum_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP csum Drop",
						},
						"icmp_pod_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total ICMP POD Drop",
						},
						"tcp_bad_urg_offet": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP bad urg off Drop",
						},
						"tcp_short_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP short hdr Drop",
						},
						"tcp_bad_ip_len": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP Bad IP Len Drop",
						},
						"tcp_null_flags": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP null flags Drop",
						},
						"tcp_null_scan": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP null scan Drop",
						},
						"tcp_fin_sin": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP SYN&FIN Drop",
						},
						"tcp_xmas_flags": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP XMAS FLAGS Drop",
						},
						"tcp_xmas_scan": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP XMAS scan Drop",
						},
						"tcp_syn_frag": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP SYN frag Drop",
						},
						"tcp_frag_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP frag header Drop",
						},
						"tcp_bad_chksum": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP bad csum Drop",
						},
						"udp_short_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP short hdr Drop",
						},
						"udp_bad_ip_len": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP bad leng Drop",
						},
						"udp_kb_frags": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP KB frag Drop",
						},
						"udp_port_lb": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP port LB Drop",
						},
						"udp_bad_chksum": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP bad csum Drop",
						},
						"runt_ip_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Runt IP hdr Drop",
						},
						"runt_tcpudp_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Runt TCPUDP hdr Drop",
						},
						"tun_mismatch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Tun mismatch Drop",
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
						"mrx_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total MRX Drop",
						},
						"hrx_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total HRX Drop",
						},
						"siz_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Size Drop",
						},
						"fcs_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total FCS Drop",
						},
						"land_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total LAND Attack Drop",
						},
						"empty_frag_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Empty frag Drop",
						},
						"mic_frag_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Micro frag Drop",
						},
						"ipv4_opt_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IPv4 opt Drop",
						},
						"ipv4_frag": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IP frag Drop",
						},
						"bad_ip_hdr_len": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP hdr len Drop",
						},
						"bad_ip_flags_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP Flags Drop",
						},
						"bad_ip_ttl_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP TTL Drop",
						},
						"no_ip_payload_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total No IP Payload Drop",
						},
						"oversize_ip_payload": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Oversize IP PL Drop",
						},
						"bad_ip_payload_len": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP PL len Drop",
						},
						"bad_ip_frag_offset": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP frag off Drop",
						},
						"bad_ip_chksum_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP csum Drop",
						},
						"icmp_pod_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total ICMP POD Drop",
						},
						"tcp_bad_urg_offet": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP bad urg off Drop",
						},
						"tcp_short_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP short hdr Drop",
						},
						"tcp_bad_ip_len": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP Bad IP Len Drop",
						},
						"tcp_null_flags": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP null flags Drop",
						},
						"tcp_null_scan": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP null scan Drop",
						},
						"tcp_fin_sin": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP SYN&FIN Drop",
						},
						"tcp_xmas_flags": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP XMAS FLAGS Drop",
						},
						"tcp_xmas_scan": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP XMAS scan Drop",
						},
						"tcp_syn_frag": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP SYN frag Drop",
						},
						"tcp_frag_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP frag header Drop",
						},
						"tcp_bad_chksum": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP bad csum Drop",
						},
						"udp_short_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP short hdr Drop",
						},
						"udp_bad_ip_len": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP bad leng Drop",
						},
						"udp_kb_frags": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP KB frag Drop",
						},
						"udp_port_lb": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP port LB Drop",
						},
						"udp_bad_chksum": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP bad csum Drop",
						},
						"runt_ip_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Runt IP hdr Drop",
						},
						"runt_tcpudp_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Runt TCPUDP hdr Drop",
						},
						"tun_mismatch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Tun mismatch Drop",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc2093(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc2093 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc2093
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MrxDrop = in["mrx_drop"].(int)
		ret.HrxDrop = in["hrx_drop"].(int)
		ret.SizDrop = in["siz_drop"].(int)
		ret.FcsDrop = in["fcs_drop"].(int)
		ret.LandDrop = in["land_drop"].(int)
		ret.EmptyFragDrop = in["empty_frag_drop"].(int)
		ret.MicFragDrop = in["mic_frag_drop"].(int)
		ret.Ipv4OptDrop = in["ipv4_opt_drop"].(int)
		ret.Ipv4Frag = in["ipv4_frag"].(int)
		ret.BadIpHdrLen = in["bad_ip_hdr_len"].(int)
		ret.BadIpFlagsDrop = in["bad_ip_flags_drop"].(int)
		ret.BadIpTtlDrop = in["bad_ip_ttl_drop"].(int)
		ret.NoIpPayloadDrop = in["no_ip_payload_drop"].(int)
		ret.OversizeIpPayload = in["oversize_ip_payload"].(int)
		ret.BadIpPayloadLen = in["bad_ip_payload_len"].(int)
		ret.BadIpFragOffset = in["bad_ip_frag_offset"].(int)
		ret.BadIpChksumDrop = in["bad_ip_chksum_drop"].(int)
		ret.IcmpPodDrop = in["icmp_pod_drop"].(int)
		ret.TcpBadUrgOffet = in["tcp_bad_urg_offet"].(int)
		ret.TcpShortHdr = in["tcp_short_hdr"].(int)
		ret.TcpBadIpLen = in["tcp_bad_ip_len"].(int)
		ret.TcpNullFlags = in["tcp_null_flags"].(int)
		ret.TcpNullScan = in["tcp_null_scan"].(int)
		ret.TcpFinSin = in["tcp_fin_sin"].(int)
		ret.TcpXmasFlags = in["tcp_xmas_flags"].(int)
		ret.TcpXmasScan = in["tcp_xmas_scan"].(int)
		ret.TcpSynFrag = in["tcp_syn_frag"].(int)
		ret.TcpFragHdr = in["tcp_frag_hdr"].(int)
		ret.TcpBadChksum = in["tcp_bad_chksum"].(int)
		ret.UdpShortHdr = in["udp_short_hdr"].(int)
		ret.UdpBadIpLen = in["udp_bad_ip_len"].(int)
		ret.UdpKbFrags = in["udp_kb_frags"].(int)
		ret.UdpPortLb = in["udp_port_lb"].(int)
		ret.UdpBadChksum = in["udp_bad_chksum"].(int)
		ret.RuntIpHdr = in["runt_ip_hdr"].(int)
		ret.RuntTcpudpHdr = in["runt_tcpudp_hdr"].(int)
		ret.TunMismatch = in["tun_mismatch"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate2094(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate2094 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate2094
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.MrxDrop = in["mrx_drop"].(int)
		ret.HrxDrop = in["hrx_drop"].(int)
		ret.SizDrop = in["siz_drop"].(int)
		ret.FcsDrop = in["fcs_drop"].(int)
		ret.LandDrop = in["land_drop"].(int)
		ret.EmptyFragDrop = in["empty_frag_drop"].(int)
		ret.MicFragDrop = in["mic_frag_drop"].(int)
		ret.Ipv4OptDrop = in["ipv4_opt_drop"].(int)
		ret.Ipv4Frag = in["ipv4_frag"].(int)
		ret.BadIpHdrLen = in["bad_ip_hdr_len"].(int)
		ret.BadIpFlagsDrop = in["bad_ip_flags_drop"].(int)
		ret.BadIpTtlDrop = in["bad_ip_ttl_drop"].(int)
		ret.NoIpPayloadDrop = in["no_ip_payload_drop"].(int)
		ret.OversizeIpPayload = in["oversize_ip_payload"].(int)
		ret.BadIpPayloadLen = in["bad_ip_payload_len"].(int)
		ret.BadIpFragOffset = in["bad_ip_frag_offset"].(int)
		ret.BadIpChksumDrop = in["bad_ip_chksum_drop"].(int)
		ret.IcmpPodDrop = in["icmp_pod_drop"].(int)
		ret.TcpBadUrgOffet = in["tcp_bad_urg_offet"].(int)
		ret.TcpShortHdr = in["tcp_short_hdr"].(int)
		ret.TcpBadIpLen = in["tcp_bad_ip_len"].(int)
		ret.TcpNullFlags = in["tcp_null_flags"].(int)
		ret.TcpNullScan = in["tcp_null_scan"].(int)
		ret.TcpFinSin = in["tcp_fin_sin"].(int)
		ret.TcpXmasFlags = in["tcp_xmas_flags"].(int)
		ret.TcpXmasScan = in["tcp_xmas_scan"].(int)
		ret.TcpSynFrag = in["tcp_syn_frag"].(int)
		ret.TcpFragHdr = in["tcp_frag_hdr"].(int)
		ret.TcpBadChksum = in["tcp_bad_chksum"].(int)
		ret.UdpShortHdr = in["udp_short_hdr"].(int)
		ret.UdpBadIpLen = in["udp_bad_ip_len"].(int)
		ret.UdpKbFrags = in["udp_kb_frags"].(int)
		ret.UdpPortLb = in["udp_port_lb"].(int)
		ret.UdpBadChksum = in["udp_bad_chksum"].(int)
		ret.RuntIpHdr = in["runt_ip_hdr"].(int)
		ret.RuntTcpudpHdr = in["runt_tcpudp_hdr"].(int)
		ret.TunMismatch = in["tun_mismatch"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc2093(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate2094(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
