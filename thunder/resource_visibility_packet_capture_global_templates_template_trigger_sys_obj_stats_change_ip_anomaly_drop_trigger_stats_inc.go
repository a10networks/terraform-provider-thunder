package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_ip_anomaly_drop_trigger_stats_inc`: Configure stats to trigger packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"bad_ip_flg": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Flags Drop",
			},
			"bad_ip_frg_offset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Fragment Offset Drop",
			},
			"bad_ip_hdrlen": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Header Len Drop",
			},
			"bad_ip_payload_len": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Payload Len Drop",
			},
			"bad_ip_ttl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP TTL Drop",
			},
			"bad_tcp_urg_offset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Bad Urgent Offset Drop",
			},
			"csum": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Checksum Drop",
			},
			"emp_frg": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Empty Fragment Drop",
			},
			"emp_mic_frg": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Micro Fragment Drop",
			},
			"frg": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv4 Fragment Drop",
			},
			"gre_pptp_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GRE PPTP Error Drop",
			},
			"ipip_tnl_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IP-over-IP Tunnel Error Drop",
			},
			"ipip_tnl_msmtch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IP-over-IP Tunnel Mismatch Drop",
			},
			"ipv6_eh_ah": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Authentication Header Drop",
			},
			"ipv6_eh_dest": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Destination Header Drop",
			},
			"ipv6_eh_esp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 ESP Header Drop",
			},
			"ipv6_eh_frag": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Fragmentation Header Drop",
			},
			"ipv6_eh_hbh": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Hop by Hop Header Drop",
			},
			"ipv6_eh_malformed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Malformed Extension Header Drop",
			},
			"ipv6_eh_mobility": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Mobility Header Drop",
			},
			"ipv6_eh_none": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 No Next Header Drop",
			},
			"ipv6_eh_other": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Unknown Extension Header Drop",
			},
			"ipv6_eh_routing": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Routing Header Drop",
			},
			"land": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Land Attack Drop",
			},
			"no_ip_payload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No IP Payload drop",
			},
			"nvgre_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GRE Tunnel Error Drop",
			},
			"opt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv4 Options Drop",
			},
			"over_ip_payload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Oversize IP Payload Drop",
			},
			"pod": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP Ping of Death Drop",
			},
			"runt_ip_hdr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Runt IP Header Drop",
			},
			"runt_tcp_udp_hdr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Runt TCP/UDP Header Drop",
			},
			"tcp_bad_csum": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Bad Checksum Drop",
			},
			"tcp_bad_iplen": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Bad IP Length Drop",
			},
			"tcp_frg_hdr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Fragmented Header Drop",
			},
			"tcp_null_frg": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Null Flags Drop",
			},
			"tcp_null_scan": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Null Scan Drop",
			},
			"tcp_opt_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Option Error Drop",
			},
			"tcp_sht_hdr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Short Header Drop",
			},
			"tcp_syn_fin": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Syn and Fin Drop",
			},
			"tcp_syn_frg": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Syn Fragment Drop",
			},
			"tcp_xmas": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP XMAS Flags Drop",
			},
			"tcp_xmas_scan": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP XMAS Scan Drop",
			},
			"udp_bad_csum": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Bad Checksum Drop",
			},
			"udp_bad_len": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Bad Length Drop",
			},
			"udp_kerb_frg": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Kerberos Fragment Drop",
			},
			"udp_port_lb": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Port Loopback Drop",
			},
			"udp_srt_hdr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Short Header Drop",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vxlan_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for VXLAN Tunnel Error Drop",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc
	ret.Inst.Bad_ip_flg = d.Get("bad_ip_flg").(int)
	ret.Inst.Bad_ip_frg_offset = d.Get("bad_ip_frg_offset").(int)
	ret.Inst.Bad_ip_hdrlen = d.Get("bad_ip_hdrlen").(int)
	ret.Inst.Bad_ip_payload_len = d.Get("bad_ip_payload_len").(int)
	ret.Inst.Bad_ip_ttl = d.Get("bad_ip_ttl").(int)
	ret.Inst.Bad_tcp_urg_offset = d.Get("bad_tcp_urg_offset").(int)
	ret.Inst.Csum = d.Get("csum").(int)
	ret.Inst.Emp_frg = d.Get("emp_frg").(int)
	ret.Inst.Emp_mic_frg = d.Get("emp_mic_frg").(int)
	ret.Inst.Frg = d.Get("frg").(int)
	ret.Inst.Gre_pptp_err = d.Get("gre_pptp_err").(int)
	ret.Inst.Ipip_tnl_err = d.Get("ipip_tnl_err").(int)
	ret.Inst.Ipip_tnl_msmtch = d.Get("ipip_tnl_msmtch").(int)
	ret.Inst.Ipv6_eh_ah = d.Get("ipv6_eh_ah").(int)
	ret.Inst.Ipv6_eh_dest = d.Get("ipv6_eh_dest").(int)
	ret.Inst.Ipv6_eh_esp = d.Get("ipv6_eh_esp").(int)
	ret.Inst.Ipv6_eh_frag = d.Get("ipv6_eh_frag").(int)
	ret.Inst.Ipv6_eh_hbh = d.Get("ipv6_eh_hbh").(int)
	ret.Inst.Ipv6_eh_malformed = d.Get("ipv6_eh_malformed").(int)
	ret.Inst.Ipv6_eh_mobility = d.Get("ipv6_eh_mobility").(int)
	ret.Inst.Ipv6_eh_none = d.Get("ipv6_eh_none").(int)
	ret.Inst.Ipv6_eh_other = d.Get("ipv6_eh_other").(int)
	ret.Inst.Ipv6_eh_routing = d.Get("ipv6_eh_routing").(int)
	ret.Inst.Land = d.Get("land").(int)
	ret.Inst.No_ip_payload = d.Get("no_ip_payload").(int)
	ret.Inst.Nvgre_err = d.Get("nvgre_err").(int)
	ret.Inst.Opt = d.Get("opt").(int)
	ret.Inst.Over_ip_payload = d.Get("over_ip_payload").(int)
	ret.Inst.Pod = d.Get("pod").(int)
	ret.Inst.Runt_ip_hdr = d.Get("runt_ip_hdr").(int)
	ret.Inst.Runt_tcp_udp_hdr = d.Get("runt_tcp_udp_hdr").(int)
	ret.Inst.Tcp_bad_csum = d.Get("tcp_bad_csum").(int)
	ret.Inst.Tcp_bad_iplen = d.Get("tcp_bad_iplen").(int)
	ret.Inst.Tcp_frg_hdr = d.Get("tcp_frg_hdr").(int)
	ret.Inst.Tcp_null_frg = d.Get("tcp_null_frg").(int)
	ret.Inst.Tcp_null_scan = d.Get("tcp_null_scan").(int)
	ret.Inst.Tcp_opt_err = d.Get("tcp_opt_err").(int)
	ret.Inst.Tcp_sht_hdr = d.Get("tcp_sht_hdr").(int)
	ret.Inst.Tcp_syn_fin = d.Get("tcp_syn_fin").(int)
	ret.Inst.Tcp_syn_frg = d.Get("tcp_syn_frg").(int)
	ret.Inst.Tcp_xmas = d.Get("tcp_xmas").(int)
	ret.Inst.Tcp_xmas_scan = d.Get("tcp_xmas_scan").(int)
	ret.Inst.Udp_bad_csum = d.Get("udp_bad_csum").(int)
	ret.Inst.Udp_bad_len = d.Get("udp_bad_len").(int)
	ret.Inst.Udp_kerb_frg = d.Get("udp_kerb_frg").(int)
	ret.Inst.Udp_port_lb = d.Get("udp_port_lb").(int)
	ret.Inst.Udp_srt_hdr = d.Get("udp_srt_hdr").(int)
	//omit uuid
	ret.Inst.Vxlan_err = d.Get("vxlan_err").(int)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
