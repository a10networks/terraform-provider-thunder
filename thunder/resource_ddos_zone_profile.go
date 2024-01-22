package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneProfile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_profile`: Profile for DDoS zone thresholds\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneProfileCreate,
		UpdateContext: resourceDdosZoneProfileUpdate,
		ReadContext:   resourceDdosZoneProfileRead,
		DeleteContext: resourceDdosZoneProfileDelete,

		Schema: map[string]*schema.Schema{
			"ip_proto": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proto_number_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol_num": {
										Type: schema.TypeInt, Required: true, Description: "Protocol Number",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"indicator_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"indicator_name": {
													Type: schema.TypeString, Required: true, Description: "'pkt-rate': pkt-rate; 'pkt-drop-rate': pkt-drop-rate; 'bit-rate': bit-rate; 'pkt-drop-ratio': pkt-drop-ratio; 'bytes-to-bytes-from-ratio': bytes-to-bytes-from-ratio; 'frag-rate': frag-rate; 'cpu-utilization': cpu-utilization; 'interface-utilization': interface-utilization;",
												},
												"src_threshold_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"src_threshold_num": {
																Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
															},
															"src_threshold_large_num": {
																Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
															},
															"src_threshold_str": {
																Type: schema.TypeString, Optional: true, Description: "Indicator per-src threshold (Non-zero floating point)",
															},
														},
													},
												},
												"zone_threshold_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"zone_threshold_num": {
																Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
															},
															"zone_threshold_large_num": {
																Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
															},
															"zone_threshold_str": {
																Type: schema.TypeString, Optional: true, Description: "Threshold for the entire zone (Non-zero floating point)",
															},
														},
													},
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
											},
										},
									},
								},
							},
						},
						"proto_name_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'icmp-v4': ip-proto icmp-v4; 'icmp-v6': ip-proto icmp-v6; 'gre': ip-proto gre; 'ipv4-encap': ip-proto IPv4 Encapsulation; 'ipv6-encap': ip-proto IPv6 Encapsulation;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"indicator_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"indicator_name": {
													Type: schema.TypeString, Required: true, Description: "'pkt-rate': pkt-rate; 'pkt-drop-rate': pkt-drop-rate; 'bit-rate': bit-rate; 'pkt-drop-ratio': pkt-drop-ratio; 'bytes-to-bytes-from-ratio': bytes-to-bytes-from-ratio; 'frag-rate': frag-rate; 'cpu-utilization': cpu-utilization; 'interface-utilization': interface-utilization;",
												},
												"src_threshold_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"src_threshold_num": {
																Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
															},
															"src_threshold_large_num": {
																Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
															},
															"src_threshold_str": {
																Type: schema.TypeString, Optional: true, Description: "Indicator per-src threshold (Non-zero floating point)",
															},
														},
													},
												},
												"zone_threshold_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"zone_threshold_num": {
																Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
															},
															"zone_threshold_large_num": {
																Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
															},
															"zone_threshold_str": {
																Type: schema.TypeString, Optional: true, Description: "Threshold for the entire zone (Non-zero floating point)",
															},
														},
													},
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"port_protocol": {
							Type: schema.TypeString, Required: true, Description: "'dns-tcp': dns-tcp; 'dns-udp': dns-udp; 'sip-tcp': sip-tcp; 'sip-udp': sip-udp; 'http': http; 'tcp': tcp; 'udp': udp; 'ssl-l4': ssl-l4; 'quic': quic;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"indicator_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"indicator_name": {
										Type: schema.TypeString, Required: true, Description: "'pkt-rate': pkt-rate; 'pkt-drop-rate': pkt-drop-rate; 'bit-rate': bit-rate; 'pkt-drop-ratio': pkt-drop-ratio; 'bytes-to-bytes-from-ratio': bytes-to-bytes-from-ratio; 'concurrent-conns': concurrent-conns; 'conn-miss-rate': conn-miss-rate; 'syn-rate': syn-rate; 'fin-rate': fin-rate; 'rst-rate': rst-rate; 'small-window-ack-rate': small-window-ack-rate; 'empty-ack-rate': empty-ack-rate; 'small-payload-rate': small-payload-rate; 'syn-fin-ratio': syn-fin-ratio; 'cpu-utilization': cpu-utilization; 'interface-utilization': interface-utilization;",
									},
									"src_threshold_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_threshold_num": {
													Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
												},
												"src_threshold_large_num": {
													Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
												},
												"src_threshold_str": {
													Type: schema.TypeString, Optional: true, Description: "Indicator per-src threshold (Non-zero floating point)",
												},
											},
										},
									},
									"zone_threshold_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"zone_threshold_num": {
													Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
												},
												"zone_threshold_large_num": {
													Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
												},
												"zone_threshold_str": {
													Type: schema.TypeString, Optional: true, Description: "Threshold for the entire zone (Non-zero floating point)",
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
								},
							},
						},
					},
				},
			},
			"port_range_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_range_start": {
							Type: schema.TypeInt, Required: true, Description: "Port-Range Start Port Number",
						},
						"port_range_end": {
							Type: schema.TypeInt, Required: true, Description: "Port-Range End Port Number",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port; 'sip-tcp': SIP-TCP Port; 'sip-udp': SIP-UDP Port; 'quic': QUIC Port;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"indicator_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"indicator_name": {
										Type: schema.TypeString, Required: true, Description: "'pkt-rate': pkt-rate; 'pkt-drop-rate': pkt-drop-rate; 'bit-rate': bit-rate; 'pkt-drop-ratio': pkt-drop-ratio; 'bytes-to-bytes-from-ratio': bytes-to-bytes-from-ratio; 'concurrent-conns': concurrent-conns; 'conn-miss-rate': conn-miss-rate; 'syn-rate': syn-rate; 'fin-rate': fin-rate; 'rst-rate': rst-rate; 'small-window-ack-rate': small-window-ack-rate; 'empty-ack-rate': empty-ack-rate; 'small-payload-rate': small-payload-rate; 'syn-fin-ratio': syn-fin-ratio; 'cpu-utilization': cpu-utilization; 'interface-utilization': interface-utilization;",
									},
									"src_threshold_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_threshold_num": {
													Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
												},
												"src_threshold_large_num": {
													Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
												},
												"src_threshold_str": {
													Type: schema.TypeString, Optional: true, Description: "Indicator per-src threshold (Non-zero floating point)",
												},
											},
										},
									},
									"zone_threshold_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"zone_threshold_num": {
													Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
												},
												"zone_threshold_large_num": {
													Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
												},
												"zone_threshold_str": {
													Type: schema.TypeString, Optional: true, Description: "Threshold for the entire zone (Non-zero floating point)",
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
								},
							},
						},
					},
				},
			},
			"profile_name": {
				Type: schema.TypeString, Required: true, Description: "Profile for DDoS zone thresholds",
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
func resourceDdosZoneProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfileRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneProfileRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneProfileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneProfile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosZoneProfileIpProto305(d []interface{}) edpt.DdosZoneProfileIpProto305 {

	count1 := len(d)
	var ret edpt.DdosZoneProfileIpProto305
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProtoNumberList = getSliceDdosZoneProfileIpProtoProtoNumberList(in["proto_number_list"].([]interface{}))
		ret.ProtoNameList = getSliceDdosZoneProfileIpProtoProtoNameList(in["proto_name_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosZoneProfileIpProtoProtoNumberList(d []interface{}) []edpt.DdosZoneProfileIpProtoProtoNumberList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneProfileIpProtoProtoNumberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneProfileIpProtoProtoNumberList
		oi.ProtocolNum = in["protocol_num"].(int)
		//omit uuid
		oi.IndicatorList = getSliceDdosZoneProfileIpProtoProtoNumberListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneProfileIpProtoProtoNumberListIndicatorList(d []interface{}) []edpt.DdosZoneProfileIpProtoProtoNumberListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneProfileIpProtoProtoNumberListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneProfileIpProtoProtoNumberListIndicatorList
		oi.IndicatorName = in["indicator_name"].(string)
		oi.SrcThresholdCfg = getObjectDdosZoneProfileIpProtoProtoNumberListIndicatorListSrcThresholdCfg(in["src_threshold_cfg"].([]interface{}))
		oi.ZoneThresholdCfg = getObjectDdosZoneProfileIpProtoProtoNumberListIndicatorListZoneThresholdCfg(in["zone_threshold_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneProfileIpProtoProtoNumberListIndicatorListSrcThresholdCfg(d []interface{}) edpt.DdosZoneProfileIpProtoProtoNumberListIndicatorListSrcThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfileIpProtoProtoNumberListIndicatorListSrcThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcThresholdNum = in["src_threshold_num"].(int)
		ret.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		ret.SrcThresholdStr = in["src_threshold_str"].(string)
	}
	return ret
}

func getObjectDdosZoneProfileIpProtoProtoNumberListIndicatorListZoneThresholdCfg(d []interface{}) edpt.DdosZoneProfileIpProtoProtoNumberListIndicatorListZoneThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfileIpProtoProtoNumberListIndicatorListZoneThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneThresholdNum = in["zone_threshold_num"].(int)
		ret.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		ret.ZoneThresholdStr = in["zone_threshold_str"].(string)
	}
	return ret
}

func getSliceDdosZoneProfileIpProtoProtoNameList(d []interface{}) []edpt.DdosZoneProfileIpProtoProtoNameList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneProfileIpProtoProtoNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneProfileIpProtoProtoNameList
		oi.Protocol = in["protocol"].(string)
		//omit uuid
		oi.IndicatorList = getSliceDdosZoneProfileIpProtoProtoNameListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneProfileIpProtoProtoNameListIndicatorList(d []interface{}) []edpt.DdosZoneProfileIpProtoProtoNameListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneProfileIpProtoProtoNameListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneProfileIpProtoProtoNameListIndicatorList
		oi.IndicatorName = in["indicator_name"].(string)
		oi.SrcThresholdCfg = getObjectDdosZoneProfileIpProtoProtoNameListIndicatorListSrcThresholdCfg(in["src_threshold_cfg"].([]interface{}))
		oi.ZoneThresholdCfg = getObjectDdosZoneProfileIpProtoProtoNameListIndicatorListZoneThresholdCfg(in["zone_threshold_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneProfileIpProtoProtoNameListIndicatorListSrcThresholdCfg(d []interface{}) edpt.DdosZoneProfileIpProtoProtoNameListIndicatorListSrcThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfileIpProtoProtoNameListIndicatorListSrcThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcThresholdNum = in["src_threshold_num"].(int)
		ret.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		ret.SrcThresholdStr = in["src_threshold_str"].(string)
	}
	return ret
}

func getObjectDdosZoneProfileIpProtoProtoNameListIndicatorListZoneThresholdCfg(d []interface{}) edpt.DdosZoneProfileIpProtoProtoNameListIndicatorListZoneThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfileIpProtoProtoNameListIndicatorListZoneThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneThresholdNum = in["zone_threshold_num"].(int)
		ret.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		ret.ZoneThresholdStr = in["zone_threshold_str"].(string)
	}
	return ret
}

func getSliceDdosZoneProfilePortList(d []interface{}) []edpt.DdosZoneProfilePortList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneProfilePortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneProfilePortList
		oi.PortNum = in["port_num"].(int)
		oi.PortProtocol = in["port_protocol"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IndicatorList = getSliceDdosZoneProfilePortListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneProfilePortListIndicatorList(d []interface{}) []edpt.DdosZoneProfilePortListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneProfilePortListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneProfilePortListIndicatorList
		oi.IndicatorName = in["indicator_name"].(string)
		oi.SrcThresholdCfg = getObjectDdosZoneProfilePortListIndicatorListSrcThresholdCfg(in["src_threshold_cfg"].([]interface{}))
		oi.ZoneThresholdCfg = getObjectDdosZoneProfilePortListIndicatorListZoneThresholdCfg(in["zone_threshold_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneProfilePortListIndicatorListSrcThresholdCfg(d []interface{}) edpt.DdosZoneProfilePortListIndicatorListSrcThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfilePortListIndicatorListSrcThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcThresholdNum = in["src_threshold_num"].(int)
		ret.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		ret.SrcThresholdStr = in["src_threshold_str"].(string)
	}
	return ret
}

func getObjectDdosZoneProfilePortListIndicatorListZoneThresholdCfg(d []interface{}) edpt.DdosZoneProfilePortListIndicatorListZoneThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfilePortListIndicatorListZoneThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneThresholdNum = in["zone_threshold_num"].(int)
		ret.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		ret.ZoneThresholdStr = in["zone_threshold_str"].(string)
	}
	return ret
}

func getSliceDdosZoneProfilePortRangeList(d []interface{}) []edpt.DdosZoneProfilePortRangeList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneProfilePortRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneProfilePortRangeList
		oi.PortRangeStart = in["port_range_start"].(int)
		oi.PortRangeEnd = in["port_range_end"].(int)
		oi.Protocol = in["protocol"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IndicatorList = getSliceDdosZoneProfilePortRangeListIndicatorList(in["indicator_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneProfilePortRangeListIndicatorList(d []interface{}) []edpt.DdosZoneProfilePortRangeListIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneProfilePortRangeListIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneProfilePortRangeListIndicatorList
		oi.IndicatorName = in["indicator_name"].(string)
		oi.SrcThresholdCfg = getObjectDdosZoneProfilePortRangeListIndicatorListSrcThresholdCfg(in["src_threshold_cfg"].([]interface{}))
		oi.ZoneThresholdCfg = getObjectDdosZoneProfilePortRangeListIndicatorListZoneThresholdCfg(in["zone_threshold_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneProfilePortRangeListIndicatorListSrcThresholdCfg(d []interface{}) edpt.DdosZoneProfilePortRangeListIndicatorListSrcThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfilePortRangeListIndicatorListSrcThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcThresholdNum = in["src_threshold_num"].(int)
		ret.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		ret.SrcThresholdStr = in["src_threshold_str"].(string)
	}
	return ret
}

func getObjectDdosZoneProfilePortRangeListIndicatorListZoneThresholdCfg(d []interface{}) edpt.DdosZoneProfilePortRangeListIndicatorListZoneThresholdCfg {

	count1 := len(d)
	var ret edpt.DdosZoneProfilePortRangeListIndicatorListZoneThresholdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneThresholdNum = in["zone_threshold_num"].(int)
		ret.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		ret.ZoneThresholdStr = in["zone_threshold_str"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneProfile(d *schema.ResourceData) edpt.DdosZoneProfile {
	var ret edpt.DdosZoneProfile
	ret.Inst.IpProto = getObjectDdosZoneProfileIpProto305(d.Get("ip_proto").([]interface{}))
	ret.Inst.PortList = getSliceDdosZoneProfilePortList(d.Get("port_list").([]interface{}))
	ret.Inst.PortRangeList = getSliceDdosZoneProfilePortRangeList(d.Get("port_range_list").([]interface{}))
	ret.Inst.ProfileName = d.Get("profile_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
