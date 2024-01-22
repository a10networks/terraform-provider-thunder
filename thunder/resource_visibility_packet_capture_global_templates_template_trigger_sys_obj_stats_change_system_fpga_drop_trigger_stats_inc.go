package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_system_fpga_drop_trigger_stats_inc`: Configure stats to trigger packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"bad_ip_chksum_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP csum Drop",
			},
			"bad_ip_flags_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP Flags Drop",
			},
			"bad_ip_frag_offset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP frag off Drop",
			},
			"bad_ip_hdr_len": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP hdr len Drop",
			},
			"bad_ip_payload_len": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP PL len Drop",
			},
			"bad_ip_ttl_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP TTL Drop",
			},
			"empty_frag_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Empty frag Drop",
			},
			"fcs_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total FCS Drop",
			},
			"hrx_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total HRX Drop",
			},
			"icmp_pod_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total ICMP POD Drop",
			},
			"ipv4_frag": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IP frag Drop",
			},
			"ipv4_opt_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IPv4 opt Drop",
			},
			"land_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total LAND Attack Drop",
			},
			"mic_frag_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Micro frag Drop",
			},
			"mrx_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total MRX Drop",
			},
			"no_ip_payload_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total No IP Payload Drop",
			},
			"oversize_ip_payload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Oversize IP PL Drop",
			},
			"runt_ip_hdr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Runt IP hdr Drop",
			},
			"runt_tcpudp_hdr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Runt TCPUDP hdr Drop",
			},
			"siz_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Size Drop",
			},
			"tcp_bad_chksum": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP bad csum Drop",
			},
			"tcp_bad_ip_len": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP Bad IP Len Drop",
			},
			"tcp_bad_urg_offet": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP bad urg off Drop",
			},
			"tcp_fin_sin": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP SYN&FIN Drop",
			},
			"tcp_frag_hdr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP frag header Drop",
			},
			"tcp_null_flags": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP null flags Drop",
			},
			"tcp_null_scan": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP null scan Drop",
			},
			"tcp_short_hdr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP short hdr Drop",
			},
			"tcp_syn_frag": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP SYN frag Drop",
			},
			"tcp_xmas_flags": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP XMAS FLAGS Drop",
			},
			"tcp_xmas_scan": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP XMAS scan Drop",
			},
			"tun_mismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Tun mismatch Drop",
			},
			"udp_bad_chksum": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP bad csum Drop",
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
			"udp_short_hdr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP short hdr Drop",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc
	ret.Inst.BadIpChksumDrop = d.Get("bad_ip_chksum_drop").(int)
	ret.Inst.BadIpFlagsDrop = d.Get("bad_ip_flags_drop").(int)
	ret.Inst.BadIpFragOffset = d.Get("bad_ip_frag_offset").(int)
	ret.Inst.BadIpHdrLen = d.Get("bad_ip_hdr_len").(int)
	ret.Inst.BadIpPayloadLen = d.Get("bad_ip_payload_len").(int)
	ret.Inst.BadIpTtlDrop = d.Get("bad_ip_ttl_drop").(int)
	ret.Inst.EmptyFragDrop = d.Get("empty_frag_drop").(int)
	ret.Inst.FcsDrop = d.Get("fcs_drop").(int)
	ret.Inst.HrxDrop = d.Get("hrx_drop").(int)
	ret.Inst.IcmpPodDrop = d.Get("icmp_pod_drop").(int)
	ret.Inst.Ipv4Frag = d.Get("ipv4_frag").(int)
	ret.Inst.Ipv4OptDrop = d.Get("ipv4_opt_drop").(int)
	ret.Inst.LandDrop = d.Get("land_drop").(int)
	ret.Inst.MicFragDrop = d.Get("mic_frag_drop").(int)
	ret.Inst.MrxDrop = d.Get("mrx_drop").(int)
	ret.Inst.NoIpPayloadDrop = d.Get("no_ip_payload_drop").(int)
	ret.Inst.OversizeIpPayload = d.Get("oversize_ip_payload").(int)
	ret.Inst.RuntIpHdr = d.Get("runt_ip_hdr").(int)
	ret.Inst.RuntTcpudpHdr = d.Get("runt_tcpudp_hdr").(int)
	ret.Inst.SizDrop = d.Get("siz_drop").(int)
	ret.Inst.TcpBadChksum = d.Get("tcp_bad_chksum").(int)
	ret.Inst.TcpBadIpLen = d.Get("tcp_bad_ip_len").(int)
	ret.Inst.TcpBadUrgOffet = d.Get("tcp_bad_urg_offet").(int)
	ret.Inst.TcpFinSin = d.Get("tcp_fin_sin").(int)
	ret.Inst.TcpFragHdr = d.Get("tcp_frag_hdr").(int)
	ret.Inst.TcpNullFlags = d.Get("tcp_null_flags").(int)
	ret.Inst.TcpNullScan = d.Get("tcp_null_scan").(int)
	ret.Inst.TcpShortHdr = d.Get("tcp_short_hdr").(int)
	ret.Inst.TcpSynFrag = d.Get("tcp_syn_frag").(int)
	ret.Inst.TcpXmasFlags = d.Get("tcp_xmas_flags").(int)
	ret.Inst.TcpXmasScan = d.Get("tcp_xmas_scan").(int)
	ret.Inst.TunMismatch = d.Get("tun_mismatch").(int)
	ret.Inst.UdpBadChksum = d.Get("udp_bad_chksum").(int)
	ret.Inst.UdpBadIpLen = d.Get("udp_bad_ip_len").(int)
	ret.Inst.UdpKbFrags = d.Get("udp_kb_frags").(int)
	ret.Inst.UdpPortLb = d.Get("udp_port_lb").(int)
	ret.Inst.UdpShortHdr = d.Get("udp_short_hdr").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
