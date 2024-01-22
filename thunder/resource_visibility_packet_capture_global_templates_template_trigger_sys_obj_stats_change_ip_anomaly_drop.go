package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDrop() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_ip_anomaly_drop`: Configure triggers for ip.anomaly-drop object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropDelete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"land": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Land Attack Drop",
						},
						"emp_frg": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Empty Fragment Drop",
						},
						"emp_mic_frg": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Micro Fragment Drop",
						},
						"opt": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv4 Options Drop",
						},
						"frg": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv4 Fragment Drop",
						},
						"bad_ip_hdrlen": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Header Len Drop",
						},
						"bad_ip_flg": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Flags Drop",
						},
						"bad_ip_ttl": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP TTL Drop",
						},
						"no_ip_payload": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No IP Payload drop",
						},
						"over_ip_payload": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Oversize IP Payload Drop",
						},
						"bad_ip_payload_len": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Payload Len Drop",
						},
						"bad_ip_frg_offset": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Fragment Offset Drop",
						},
						"csum": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Checksum Drop",
						},
						"pod": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP Ping of Death Drop",
						},
						"bad_tcp_urg_offset": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Bad Urgent Offset Drop",
						},
						"tcp_sht_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Short Header Drop",
						},
						"tcp_bad_iplen": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Bad IP Length Drop",
						},
						"tcp_null_frg": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Null Flags Drop",
						},
						"tcp_null_scan": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Null Scan Drop",
						},
						"tcp_syn_fin": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Syn and Fin Drop",
						},
						"tcp_xmas": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP XMAS Flags Drop",
						},
						"tcp_xmas_scan": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP XMAS Scan Drop",
						},
						"tcp_syn_frg": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Syn Fragment Drop",
						},
						"tcp_frg_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Fragmented Header Drop",
						},
						"tcp_bad_csum": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Bad Checksum Drop",
						},
						"udp_srt_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Short Header Drop",
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
						"udp_bad_csum": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Bad Checksum Drop",
						},
						"runt_ip_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Runt IP Header Drop",
						},
						"runt_tcp_udp_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Runt TCP/UDP Header Drop",
						},
						"ipip_tnl_msmtch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IP-over-IP Tunnel Mismatch Drop",
						},
						"tcp_opt_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Option Error Drop",
						},
						"ipip_tnl_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IP-over-IP Tunnel Error Drop",
						},
						"vxlan_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for VXLAN Tunnel Error Drop",
						},
						"nvgre_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GRE Tunnel Error Drop",
						},
						"gre_pptp_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GRE PPTP Error Drop",
						},
						"ipv6_eh_hbh": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Hop by Hop Header Drop",
						},
						"ipv6_eh_dest": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Destination Header Drop",
						},
						"ipv6_eh_routing": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Routing Header Drop",
						},
						"ipv6_eh_frag": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Fragmentation Header Drop",
						},
						"ipv6_eh_ah": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Authentication Header Drop",
						},
						"ipv6_eh_esp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 ESP Header Drop",
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
						"ipv6_eh_malformed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Malformed Extension Header Drop",
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
						"land": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Land Attack Drop",
						},
						"emp_frg": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Empty Fragment Drop",
						},
						"emp_mic_frg": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Micro Fragment Drop",
						},
						"opt": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv4 Options Drop",
						},
						"frg": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv4 Fragment Drop",
						},
						"bad_ip_hdrlen": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Header Len Drop",
						},
						"bad_ip_flg": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Flags Drop",
						},
						"bad_ip_ttl": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP TTL Drop",
						},
						"no_ip_payload": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No IP Payload drop",
						},
						"over_ip_payload": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Oversize IP Payload Drop",
						},
						"bad_ip_payload_len": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Payload Len Drop",
						},
						"bad_ip_frg_offset": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Fragment Offset Drop",
						},
						"csum": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Checksum Drop",
						},
						"pod": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP Ping of Death Drop",
						},
						"bad_tcp_urg_offset": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Bad Urgent Offset Drop",
						},
						"tcp_sht_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Short Header Drop",
						},
						"tcp_bad_iplen": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Bad IP Length Drop",
						},
						"tcp_null_frg": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Null Flags Drop",
						},
						"tcp_null_scan": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Null Scan Drop",
						},
						"tcp_syn_fin": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Syn and Fin Drop",
						},
						"tcp_xmas": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP XMAS Flags Drop",
						},
						"tcp_xmas_scan": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP XMAS Scan Drop",
						},
						"tcp_syn_frg": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Syn Fragment Drop",
						},
						"tcp_frg_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Fragmented Header Drop",
						},
						"tcp_bad_csum": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Bad Checksum Drop",
						},
						"udp_srt_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Short Header Drop",
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
						"udp_bad_csum": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Bad Checksum Drop",
						},
						"runt_ip_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Runt IP Header Drop",
						},
						"runt_tcp_udp_hdr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Runt TCP/UDP Header Drop",
						},
						"ipip_tnl_msmtch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IP-over-IP Tunnel Mismatch Drop",
						},
						"tcp_opt_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Option Error Drop",
						},
						"ipip_tnl_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IP-over-IP Tunnel Error Drop",
						},
						"vxlan_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for VXLAN Tunnel Error Drop",
						},
						"nvgre_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GRE Tunnel Error Drop",
						},
						"gre_pptp_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GRE PPTP Error Drop",
						},
						"ipv6_eh_hbh": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Hop by Hop Header Drop",
						},
						"ipv6_eh_dest": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Destination Header Drop",
						},
						"ipv6_eh_routing": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Routing Header Drop",
						},
						"ipv6_eh_frag": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Fragmentation Header Drop",
						},
						"ipv6_eh_ah": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Authentication Header Drop",
						},
						"ipv6_eh_esp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 ESP Header Drop",
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
						"ipv6_eh_malformed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Malformed Extension Header Drop",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDrop(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDrop(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDrop(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDrop(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc2017(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc2017 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc2017
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Land = in["land"].(int)
		ret.Emp_frg = in["emp_frg"].(int)
		ret.Emp_mic_frg = in["emp_mic_frg"].(int)
		ret.Opt = in["opt"].(int)
		ret.Frg = in["frg"].(int)
		ret.Bad_ip_hdrlen = in["bad_ip_hdrlen"].(int)
		ret.Bad_ip_flg = in["bad_ip_flg"].(int)
		ret.Bad_ip_ttl = in["bad_ip_ttl"].(int)
		ret.No_ip_payload = in["no_ip_payload"].(int)
		ret.Over_ip_payload = in["over_ip_payload"].(int)
		ret.Bad_ip_payload_len = in["bad_ip_payload_len"].(int)
		ret.Bad_ip_frg_offset = in["bad_ip_frg_offset"].(int)
		ret.Csum = in["csum"].(int)
		ret.Pod = in["pod"].(int)
		ret.Bad_tcp_urg_offset = in["bad_tcp_urg_offset"].(int)
		ret.Tcp_sht_hdr = in["tcp_sht_hdr"].(int)
		ret.Tcp_bad_iplen = in["tcp_bad_iplen"].(int)
		ret.Tcp_null_frg = in["tcp_null_frg"].(int)
		ret.Tcp_null_scan = in["tcp_null_scan"].(int)
		ret.Tcp_syn_fin = in["tcp_syn_fin"].(int)
		ret.Tcp_xmas = in["tcp_xmas"].(int)
		ret.Tcp_xmas_scan = in["tcp_xmas_scan"].(int)
		ret.Tcp_syn_frg = in["tcp_syn_frg"].(int)
		ret.Tcp_frg_hdr = in["tcp_frg_hdr"].(int)
		ret.Tcp_bad_csum = in["tcp_bad_csum"].(int)
		ret.Udp_srt_hdr = in["udp_srt_hdr"].(int)
		ret.Udp_bad_len = in["udp_bad_len"].(int)
		ret.Udp_kerb_frg = in["udp_kerb_frg"].(int)
		ret.Udp_port_lb = in["udp_port_lb"].(int)
		ret.Udp_bad_csum = in["udp_bad_csum"].(int)
		ret.Runt_ip_hdr = in["runt_ip_hdr"].(int)
		ret.Runt_tcp_udp_hdr = in["runt_tcp_udp_hdr"].(int)
		ret.Ipip_tnl_msmtch = in["ipip_tnl_msmtch"].(int)
		ret.Tcp_opt_err = in["tcp_opt_err"].(int)
		ret.Ipip_tnl_err = in["ipip_tnl_err"].(int)
		ret.Vxlan_err = in["vxlan_err"].(int)
		ret.Nvgre_err = in["nvgre_err"].(int)
		ret.Gre_pptp_err = in["gre_pptp_err"].(int)
		ret.Ipv6_eh_hbh = in["ipv6_eh_hbh"].(int)
		ret.Ipv6_eh_dest = in["ipv6_eh_dest"].(int)
		ret.Ipv6_eh_routing = in["ipv6_eh_routing"].(int)
		ret.Ipv6_eh_frag = in["ipv6_eh_frag"].(int)
		ret.Ipv6_eh_ah = in["ipv6_eh_ah"].(int)
		ret.Ipv6_eh_esp = in["ipv6_eh_esp"].(int)
		ret.Ipv6_eh_mobility = in["ipv6_eh_mobility"].(int)
		ret.Ipv6_eh_none = in["ipv6_eh_none"].(int)
		ret.Ipv6_eh_other = in["ipv6_eh_other"].(int)
		ret.Ipv6_eh_malformed = in["ipv6_eh_malformed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsRate2018(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsRate2018 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsRate2018
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Land = in["land"].(int)
		ret.Emp_frg = in["emp_frg"].(int)
		ret.Emp_mic_frg = in["emp_mic_frg"].(int)
		ret.Opt = in["opt"].(int)
		ret.Frg = in["frg"].(int)
		ret.Bad_ip_hdrlen = in["bad_ip_hdrlen"].(int)
		ret.Bad_ip_flg = in["bad_ip_flg"].(int)
		ret.Bad_ip_ttl = in["bad_ip_ttl"].(int)
		ret.No_ip_payload = in["no_ip_payload"].(int)
		ret.Over_ip_payload = in["over_ip_payload"].(int)
		ret.Bad_ip_payload_len = in["bad_ip_payload_len"].(int)
		ret.Bad_ip_frg_offset = in["bad_ip_frg_offset"].(int)
		ret.Csum = in["csum"].(int)
		ret.Pod = in["pod"].(int)
		ret.Bad_tcp_urg_offset = in["bad_tcp_urg_offset"].(int)
		ret.Tcp_sht_hdr = in["tcp_sht_hdr"].(int)
		ret.Tcp_bad_iplen = in["tcp_bad_iplen"].(int)
		ret.Tcp_null_frg = in["tcp_null_frg"].(int)
		ret.Tcp_null_scan = in["tcp_null_scan"].(int)
		ret.Tcp_syn_fin = in["tcp_syn_fin"].(int)
		ret.Tcp_xmas = in["tcp_xmas"].(int)
		ret.Tcp_xmas_scan = in["tcp_xmas_scan"].(int)
		ret.Tcp_syn_frg = in["tcp_syn_frg"].(int)
		ret.Tcp_frg_hdr = in["tcp_frg_hdr"].(int)
		ret.Tcp_bad_csum = in["tcp_bad_csum"].(int)
		ret.Udp_srt_hdr = in["udp_srt_hdr"].(int)
		ret.Udp_bad_len = in["udp_bad_len"].(int)
		ret.Udp_kerb_frg = in["udp_kerb_frg"].(int)
		ret.Udp_port_lb = in["udp_port_lb"].(int)
		ret.Udp_bad_csum = in["udp_bad_csum"].(int)
		ret.Runt_ip_hdr = in["runt_ip_hdr"].(int)
		ret.Runt_tcp_udp_hdr = in["runt_tcp_udp_hdr"].(int)
		ret.Ipip_tnl_msmtch = in["ipip_tnl_msmtch"].(int)
		ret.Tcp_opt_err = in["tcp_opt_err"].(int)
		ret.Ipip_tnl_err = in["ipip_tnl_err"].(int)
		ret.Vxlan_err = in["vxlan_err"].(int)
		ret.Nvgre_err = in["nvgre_err"].(int)
		ret.Gre_pptp_err = in["gre_pptp_err"].(int)
		ret.Ipv6_eh_hbh = in["ipv6_eh_hbh"].(int)
		ret.Ipv6_eh_dest = in["ipv6_eh_dest"].(int)
		ret.Ipv6_eh_routing = in["ipv6_eh_routing"].(int)
		ret.Ipv6_eh_frag = in["ipv6_eh_frag"].(int)
		ret.Ipv6_eh_ah = in["ipv6_eh_ah"].(int)
		ret.Ipv6_eh_esp = in["ipv6_eh_esp"].(int)
		ret.Ipv6_eh_mobility = in["ipv6_eh_mobility"].(int)
		ret.Ipv6_eh_none = in["ipv6_eh_none"].(int)
		ret.Ipv6_eh_other = in["ipv6_eh_other"].(int)
		ret.Ipv6_eh_malformed = in["ipv6_eh_malformed"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDrop(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDrop {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDrop
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc2017(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsRate2018(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
