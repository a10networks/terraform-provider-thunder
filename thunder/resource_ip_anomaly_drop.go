package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAnomalyDrop() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_anomaly_drop`: Set IP anomaly drop policy\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAnomalyDropCreate,
		UpdateContext: resourceIpAnomalyDropUpdate,
		ReadContext:   resourceIpAnomalyDropRead,
		DeleteContext: resourceIpAnomalyDropDelete,

		Schema: map[string]*schema.Schema{
			"bad_content": {
				Type: schema.TypeInt, Optional: true, Description: "bad content threshold (threshold value)",
			},
			"drop_all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "drop all IP anomaly packets",
			},
			"frag": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "drop all fragmented packets",
			},
			"ip_option": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "drop packets with IP options",
			},
			"ipv6_ext_header": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_eh_frag": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Filter fragmentation extension header",
						},
						"ipv6_eh_auth": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Filter authentication extension header",
						},
						"ipv6_eh_esp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Filter ESP extension header",
						},
						"ipv6_eh_mobility": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Filter mobility extension header",
						},
						"ipv6_eh_nonext": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Filter no-next-header extension header",
						},
						"ipv6_eh_malformed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Filter malformed extension headers (check for order and occurrences)",
						},
						"ipv6_eh_hbh": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Filter hop by hop extension header",
						},
						"ipv6_eh_dest": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Filter destination extension header",
						},
						"ipv6_eh_routing": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Filter routing extension header",
						},
						"hbh_option_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hbh_otype_from": {
										Type: schema.TypeInt, Optional: true, Description: "Filter hop by hop option type (Option type value)",
									},
									"hbh_otype_to": {
										Type: schema.TypeInt, Optional: true, Description: "Option type range end",
									},
								},
							},
						},
						"dst_option_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_otype_from": {
										Type: schema.TypeInt, Optional: true, Description: "Filter destination header option type (Option type value)",
									},
									"dst_otype_to": {
										Type: schema.TypeInt, Optional: true, Description: "Option type range end",
									},
								},
							},
						},
						"routing_option_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"routing_otype_from": {
										Type: schema.TypeInt, Optional: true, Description: "Filter routing header option type (Option type value)",
									},
									"routing_otype_to": {
										Type: schema.TypeInt, Optional: true, Description: "Option type range end",
									},
								},
							},
						},
						"unknown_ext_header_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"eh_type_from": {
										Type: schema.TypeInt, Optional: true, Description: "Filter unknown extension header (eh) type (Extension header type value)",
									},
									"eh_type_to": {
										Type: schema.TypeInt, Optional: true, Description: "Extension header type range end",
									},
								},
							},
						},
					},
				},
			},
			"land_attack": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "drop IP packets with the same source and destination addresses",
			},
			"out_of_sequence": {
				Type: schema.TypeInt, Optional: true, Description: "out of sequence packet threshold (threshold value)",
			},
			"packet_deformity": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packet_deformity_layer_3": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "drop packets with layer 3 anomaly",
						},
						"packet_deformity_layer_4": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "drop packets with layer 4 anomaly",
						},
					},
				},
			},
			"ping_of_death": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "drop oversize ICMP packets",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'land': Land Attack Drop; 'emp_frg': Empty Fragment Drop; 'emp_mic_frg': Micro Fragment Drop; 'opt': IPv4 Options Drop; 'frg': IPv4 Fragment Drop; 'bad_ip_hdrlen': Bad IP Header Len Drop; 'bad_ip_flg': Bad IP Flags Drop; 'bad_ip_ttl': Bad IP TTL Drop; 'no_ip_payload': No IP Payload drop; 'over_ip_payload': Oversize IP Payload Drop; 'bad_ip_payload_len': Bad IP Payload Len Drop; 'bad_ip_frg_offset': Bad IP Fragment Offset Drop; 'csum': Bad IP Checksum Drop; 'pod': ICMP Ping of Death Drop; 'bad_tcp_urg_offset': TCP Bad Urgent Offset Drop; 'tcp_sht_hdr': TCP Short Header Drop; 'tcp_bad_iplen': TCP Bad IP Length Drop; 'tcp_null_frg': TCP Null Flags Drop; 'tcp_null_scan': TCP Null Scan Drop; 'tcp_syn_fin': TCP Syn and Fin Drop; 'tcp_xmas': TCP XMAS Flags Drop; 'tcp_xmas_scan': TCP XMAS Scan Drop; 'tcp_syn_frg': TCP Syn Fragment Drop; 'tcp_frg_hdr': TCP Fragmented Header Drop; 'tcp_bad_csum': TCP Bad Checksum Drop; 'udp_srt_hdr': UDP Short Header Drop; 'udp_bad_len': UDP Bad Length Drop; 'udp_kerb_frg': UDP Kerberos Fragment Drop; 'udp_port_lb': UDP Port Loopback Drop; 'udp_bad_csum': UDP Bad Checksum Drop; 'runt_ip_hdr': Runt IP Header Drop; 'runt_tcp_udp_hdr': Runt TCP/UDP Header Drop; 'ipip_tnl_msmtch': IP-over-IP Tunnel Mismatch Drop; 'tcp_opt_err': TCP Option Error Drop; 'ipip_tnl_err': IP-over-IP Tunnel Error Drop; 'vxlan_err': VXLAN Tunnel Error Drop; 'nvgre_err': GRE Tunnel Error Drop; 'gre_pptp_err': GRE PPTP Error Drop; 'ipv6_eh_hbh': IPv6 Hop by Hop Header Drop; 'ipv6_eh_dest': IPv6 Destination Header Drop; 'ipv6_eh_routing': IPv6 Routing Header Drop; 'ipv6_eh_frag': IPv6 Fragmentation Header Drop; 'ipv6_eh_ah': IPv6 Authentication Header Drop; 'ipv6_eh_esp': IPv6 ESP Header Drop; 'ipv6_eh_mobility': IPv6 Mobility Header Drop; 'ipv6_eh_none': IPv6 No Next Header Drop; 'ipv6_eh_other': IPv6 Unknown Extension Header Drop; 'ipv6_eh_malformed': IPv6 Malformed Extension Header Drop;",
						},
					},
				},
			},
			"security_attack": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"security_attack_layer_3": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "drop packets with layer 3 anomaly",
						},
						"security_attack_layer_4": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "drop packets with layer 4 anomaly",
						},
					},
				},
			},
			"tcp_no_flag": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "drop TCP packets with no flag",
			},
			"tcp_syn_fin": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "drop TCP packets with both syn and fin flags set",
			},
			"tcp_syn_frag": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "drop fragmented TCP packets with syn flag set",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zero_window": {
				Type: schema.TypeInt, Optional: true, Description: "zero window size threshold (threshold value)",
			},
		},
	}
}
func resourceIpAnomalyDropCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAnomalyDropCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAnomalyDrop(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAnomalyDropRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAnomalyDropUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAnomalyDropUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAnomalyDrop(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAnomalyDropRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAnomalyDropDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAnomalyDropDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAnomalyDrop(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAnomalyDropRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAnomalyDropRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAnomalyDrop(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectIpAnomalyDropIpv6ExtHeader(d []interface{}) edpt.IpAnomalyDropIpv6ExtHeader {

	count1 := len(d)
	var ret edpt.IpAnomalyDropIpv6ExtHeader
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv6EhFrag = in["ipv6_eh_frag"].(int)
		ret.Ipv6EhAuth = in["ipv6_eh_auth"].(int)
		ret.Ipv6EhEsp = in["ipv6_eh_esp"].(int)
		ret.Ipv6EhMobility = in["ipv6_eh_mobility"].(int)
		ret.Ipv6EhNonext = in["ipv6_eh_nonext"].(int)
		ret.Ipv6EhMalformed = in["ipv6_eh_malformed"].(int)
		ret.Ipv6EhHbh = in["ipv6_eh_hbh"].(int)
		ret.Ipv6EhDest = in["ipv6_eh_dest"].(int)
		ret.Ipv6EhRouting = in["ipv6_eh_routing"].(int)
		ret.HbhOptionList = getSliceIpAnomalyDropIpv6ExtHeaderHbhOptionList(in["hbh_option_list"].([]interface{}))
		ret.DstOptionList = getSliceIpAnomalyDropIpv6ExtHeaderDstOptionList(in["dst_option_list"].([]interface{}))
		ret.RoutingOptionList = getSliceIpAnomalyDropIpv6ExtHeaderRoutingOptionList(in["routing_option_list"].([]interface{}))
		ret.UnknownExtHeaderList = getSliceIpAnomalyDropIpv6ExtHeaderUnknownExtHeaderList(in["unknown_ext_header_list"].([]interface{}))
	}
	return ret
}

func getSliceIpAnomalyDropIpv6ExtHeaderHbhOptionList(d []interface{}) []edpt.IpAnomalyDropIpv6ExtHeaderHbhOptionList {

	count1 := len(d)
	ret := make([]edpt.IpAnomalyDropIpv6ExtHeaderHbhOptionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAnomalyDropIpv6ExtHeaderHbhOptionList
		oi.HbhOtypeFrom = in["hbh_otype_from"].(int)
		oi.HbhOtypeTo = in["hbh_otype_to"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpAnomalyDropIpv6ExtHeaderDstOptionList(d []interface{}) []edpt.IpAnomalyDropIpv6ExtHeaderDstOptionList {

	count1 := len(d)
	ret := make([]edpt.IpAnomalyDropIpv6ExtHeaderDstOptionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAnomalyDropIpv6ExtHeaderDstOptionList
		oi.DstOtypeFrom = in["dst_otype_from"].(int)
		oi.DstOtypeTo = in["dst_otype_to"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpAnomalyDropIpv6ExtHeaderRoutingOptionList(d []interface{}) []edpt.IpAnomalyDropIpv6ExtHeaderRoutingOptionList {

	count1 := len(d)
	ret := make([]edpt.IpAnomalyDropIpv6ExtHeaderRoutingOptionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAnomalyDropIpv6ExtHeaderRoutingOptionList
		oi.RoutingOtypeFrom = in["routing_otype_from"].(int)
		oi.RoutingOtypeTo = in["routing_otype_to"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpAnomalyDropIpv6ExtHeaderUnknownExtHeaderList(d []interface{}) []edpt.IpAnomalyDropIpv6ExtHeaderUnknownExtHeaderList {

	count1 := len(d)
	ret := make([]edpt.IpAnomalyDropIpv6ExtHeaderUnknownExtHeaderList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAnomalyDropIpv6ExtHeaderUnknownExtHeaderList
		oi.EhTypeFrom = in["eh_type_from"].(int)
		oi.EhTypeTo = in["eh_type_to"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectIpAnomalyDropPacketDeformity(d []interface{}) edpt.IpAnomalyDropPacketDeformity {

	count1 := len(d)
	var ret edpt.IpAnomalyDropPacketDeformity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PacketDeformityLayer3 = in["packet_deformity_layer_3"].(int)
		ret.PacketDeformityLayer4 = in["packet_deformity_layer_4"].(int)
	}
	return ret
}

func getSliceIpAnomalyDropSamplingEnable(d []interface{}) []edpt.IpAnomalyDropSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.IpAnomalyDropSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAnomalyDropSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectIpAnomalyDropSecurityAttack(d []interface{}) edpt.IpAnomalyDropSecurityAttack {

	count1 := len(d)
	var ret edpt.IpAnomalyDropSecurityAttack
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SecurityAttackLayer3 = in["security_attack_layer_3"].(int)
		ret.SecurityAttackLayer4 = in["security_attack_layer_4"].(int)
	}
	return ret
}

func dataToEndpointIpAnomalyDrop(d *schema.ResourceData) edpt.IpAnomalyDrop {
	var ret edpt.IpAnomalyDrop
	ret.Inst.BadContent = d.Get("bad_content").(int)
	ret.Inst.DropAll = d.Get("drop_all").(int)
	ret.Inst.Frag = d.Get("frag").(int)
	ret.Inst.IpOption = d.Get("ip_option").(int)
	ret.Inst.Ipv6ExtHeader = getObjectIpAnomalyDropIpv6ExtHeader(d.Get("ipv6_ext_header").([]interface{}))
	ret.Inst.LandAttack = d.Get("land_attack").(int)
	ret.Inst.OutOfSequence = d.Get("out_of_sequence").(int)
	ret.Inst.PacketDeformity = getObjectIpAnomalyDropPacketDeformity(d.Get("packet_deformity").([]interface{}))
	ret.Inst.PingOfDeath = d.Get("ping_of_death").(int)
	ret.Inst.SamplingEnable = getSliceIpAnomalyDropSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SecurityAttack = getObjectIpAnomalyDropSecurityAttack(d.Get("security_attack").([]interface{}))
	ret.Inst.TcpNoFlag = d.Get("tcp_no_flag").(int)
	ret.Inst.TcpSynFin = d.Get("tcp_syn_fin").(int)
	ret.Inst.TcpSynFrag = d.Get("tcp_syn_frag").(int)
	//omit uuid
	ret.Inst.ZeroWindow = d.Get("zero_window").(int)
	return ret
}
