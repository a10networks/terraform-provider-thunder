package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryPort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_port`: DDOS Port & Protocol configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntryPortCreate,
		UpdateContext: resourceDdosDstEntryPortUpdate,
		ReadContext:   resourceDdosDstEntryPortRead,
		DeleteContext: resourceDdosDstEntryPortDelete,

		Schema: map[string]*schema.Schema{
			"capture_config": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"capture_config_name": {
							Type: schema.TypeString, Optional: true, Description: "Capture-config name",
						},
						"capture_config_mode": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Apply capture-config to dropped packets; 'forward': Apply capture-config to forwarded packets; 'all': Apply capture-config to both dropped and forwarded packets;",
						},
					},
				},
			},
			"deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
			},
			"detection_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos detection",
			},
			"dns_cache": {
				Type: schema.TypeString, Optional: true, Description: "DNS Cache Instance",
			},
			"enable_top_k": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos top-k entries",
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"glid_exceed_action": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stateless_encap_action_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"stateless_encap_action": {
										Type: schema.TypeString, Optional: true, Description: "'stateless-tunnel-encap': Encapsulate all packets; 'stateless-tunnel-encap-scrubbed': Encapsulate all packets and allow packets to go through other DDoS checks before sent (conn-limit exceeded packet can not be scrubbed, it will default to stateless-tunnel-encap);",
									},
									"encap_template": {
										Type: schema.TypeString, Optional: true, Description: "Apply legacy encap template for encap action",
									},
								},
							},
						},
					},
				},
			},
			"ip_filtering_policy": {
				Type: schema.TypeString, Optional: true, Description: "Configure IP Filter",
			},
			"ip_filtering_policy_oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"pattern_recognition": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"algorithm": {
							Type: schema.TypeString, Optional: true, Description: "'heuristic': heuristic algorithm;",
						},
						"mode": {
							Type: schema.TypeString, Optional: true, Description: "'capture-never-expire': War-time capture without rate exceeding and never expires; 'manual': Manual mode;",
						},
						"sensitivity": {
							Type: schema.TypeString, Optional: true, Description: "'high': High Sensitivity; 'medium': Medium Sensitivity; 'low': Low Sensitivity;",
						},
						"filter_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Extracted filter threshold",
						},
						"filter_inactive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Extracted filter inactive threshold",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"pattern_recognition_pu_details": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"port_ind": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'ip-proto-type': IP Protocol Type; 'ddet_ind_pkt_rate_current': Pkt Rate Current; 'ddet_ind_pkt_rate_min': Pkt Rate Min; 'ddet_ind_pkt_rate_max': Pkt Rate Max; 'ddet_ind_pkt_drop_rate_current': Pkt Drop Rate Current; 'ddet_ind_pkt_drop_rate_min': Pkt Drop Rate Min; 'ddet_ind_pkt_drop_rate_max': Pkt Drop Rate Max; 'ddet_ind_syn_rate_current': TCP SYN Rate Current; 'ddet_ind_syn_rate_min': TCP SYN Rate Min; 'ddet_ind_syn_rate_max': TCP SYN Rate Max; 'ddet_ind_fin_rate_current': TCP FIN Rate Current; 'ddet_ind_fin_rate_min': TCP FIN Rate Min; 'ddet_ind_fin_rate_max': TCP FIN Rate Max; 'ddet_ind_rst_rate_current': TCP RST Rate Current; 'ddet_ind_rst_rate_min': TCP RST Rate Min; 'ddet_ind_rst_rate_max': TCP RST Rate Max; 'ddet_ind_small_window_ack_rate_current': TCP Small Window ACK Rate Current; 'ddet_ind_small_window_ack_rate_min': TCP Small Window ACK Rate Min; 'ddet_ind_small_window_ack_rate_max': TCP Small Window ACK Rate Max; 'ddet_ind_empty_ack_rate_current': TCP Empty ACK Rate Current; 'ddet_ind_empty_ack_rate_min': TCP Empty ACK Rate Min; 'ddet_ind_empty_ack_rate_max': TCP Empty ACK Rate Max; 'ddet_ind_small_payload_rate_current': TCP Small Payload Rate Current; 'ddet_ind_small_payload_rate_min': TCP Small Payload Rate Min; 'ddet_ind_small_payload_rate_max': TCP Small Payload Rate Max; 'ddet_ind_pkt_drop_ratio_current': Pkt Drop / Pkt Rcvd Current; 'ddet_ind_pkt_drop_ratio_min': Pkt Drop / Pkt Rcvd Min; 'ddet_ind_pkt_drop_ratio_max': Pkt Drop / Pkt Rcvd Max; 'ddet_ind_inb_per_outb_current': Bytes-to / Bytes-from Current; 'ddet_ind_inb_per_outb_min': Bytes-to / Bytes-from Min; 'ddet_ind_inb_per_outb_max': Bytes-to / Bytes-from Max; 'ddet_ind_syn_per_fin_rate_current': TCP SYN Rate / FIN Rate Current; 'ddet_ind_syn_per_fin_rate_min': TCP SYN Rate / FIN Rate Min; 'ddet_ind_syn_per_fin_rate_max': TCP SYN Rate / FIN Rate Max; 'ddet_ind_conn_miss_rate_current': TCP Session Miss Rate Current; 'ddet_ind_conn_miss_rate_min': TCP Session Miss Rate Min; 'ddet_ind_conn_miss_rate_max': TCP Session Miss Rate Max; 'ddet_ind_concurrent_conns_current': TCP/UDP Concurrent Sessions Current; 'ddet_ind_concurrent_conns_min': TCP/UDP Concurrent Sessions Min; 'ddet_ind_concurrent_conns_max': TCP/UDP Concurrent Sessions Max; 'ddet_ind_data_cpu_util_current': Data CPU Utilization Current; 'ddet_ind_data_cpu_util_min': Data CPU Utilization Min; 'ddet_ind_data_cpu_util_max': Data CPU Utilization Max; 'ddet_ind_outside_intf_util_current': Outside Interface Utilization Current; 'ddet_ind_outside_intf_util_min': Outside Interface Utilization Min; 'ddet_ind_outside_intf_util_max': Outside Interface Utilization Max; 'ddet_ind_frag_rate_current': Frag Pkt Rate Current; 'ddet_ind_frag_rate_min': Frag Pkt Rate Min; 'ddet_ind_frag_rate_max': Frag Pkt Rate Max; 'ddet_ind_bit_rate_current': Bit Rate Current; 'ddet_ind_bit_rate_min': Bit Rate Min; 'ddet_ind_bit_rate_max': Bit Rate Max;",
									},
								},
							},
						},
					},
				},
			},
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "Port Number",
			},
			"progression_tracking": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port; 'sip-udp': SIP-UDP Port; 'sip-tcp': SIP-TCP Port;",
			},
			"set_counter_base_val": {
				Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
			},
			"sflow": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"polling": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sflow_packets": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow packet-level counter polling",
									},
									"sflow_tcp": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sflow_tcp_basic": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow basic TCP counter polling",
												},
												"sflow_tcp_stateful": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow stateful TCP counter polling",
												},
											},
										},
									},
									"sflow_http": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sFlow HTTP counter polling",
									},
								},
							},
						},
					},
				},
			},
			"signature_extraction": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"algorithm": {
							Type: schema.TypeString, Optional: true, Description: "'heuristic': heuristic algorithm;",
						},
						"manual_mode": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable manual mode",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns": {
							Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
						},
						"http": {
							Type: schema.TypeString, Optional: true, Description: "DDOS http template",
						},
						"ssl_l4": {
							Type: schema.TypeString, Optional: true, Description: "DDOS SSL-L4 template",
						},
						"sip": {
							Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
						},
						"tcp": {
							Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
						},
						"udp": {
							Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
						},
					},
				},
			},
			"topk_num_records": {
				Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
			},
			"topk_sources": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}
func resourceDdosDstEntryPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryPortRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntryPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryPortRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntryPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntryPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstEntryPortCaptureConfig(d []interface{}) edpt.DdosDstEntryPortCaptureConfig {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortCaptureConfig
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CaptureConfigName = in["capture_config_name"].(string)
		ret.CaptureConfigMode = in["capture_config_mode"].(string)
	}
	return ret
}

func getObjectDdosDstEntryPortGlidExceedAction(d []interface{}) edpt.DdosDstEntryPortGlidExceedAction {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortGlidExceedAction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapActionCfg = getObjectDdosDstEntryPortGlidExceedActionStatelessEncapActionCfg(in["stateless_encap_action_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortGlidExceedActionStatelessEncapActionCfg(d []interface{}) edpt.DdosDstEntryPortGlidExceedActionStatelessEncapActionCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortGlidExceedActionStatelessEncapActionCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapAction = in["stateless_encap_action"].(string)
		ret.EncapTemplate = in["encap_template"].(string)
	}
	return ret
}

func getObjectDdosDstEntryPortIpFilteringPolicyOper170(d []interface{}) edpt.DdosDstEntryPortIpFilteringPolicyOper170 {

	var ret edpt.DdosDstEntryPortIpFilteringPolicyOper170
	return ret
}

func getObjectDdosDstEntryPortPatternRecognition171(d []interface{}) edpt.DdosDstEntryPortPatternRecognition171 {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortPatternRecognition171
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Algorithm = in["algorithm"].(string)
		ret.Mode = in["mode"].(string)
		ret.Sensitivity = in["sensitivity"].(string)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterInactiveThreshold = in["filter_inactive_threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDstEntryPortPatternRecognitionPuDetails172(d []interface{}) edpt.DdosDstEntryPortPatternRecognitionPuDetails172 {

	var ret edpt.DdosDstEntryPortPatternRecognitionPuDetails172
	return ret
}

func getObjectDdosDstEntryPortPortInd173(d []interface{}) edpt.DdosDstEntryPortPortInd173 {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortPortInd173
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceDdosDstEntryPortPortIndSamplingEnable174(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryPortPortIndSamplingEnable174(d []interface{}) []edpt.DdosDstEntryPortPortIndSamplingEnable174 {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortPortIndSamplingEnable174, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortPortIndSamplingEnable174
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryPortProgressionTracking175(d []interface{}) edpt.DdosDstEntryPortProgressionTracking175 {

	var ret edpt.DdosDstEntryPortProgressionTracking175
	return ret
}

func getObjectDdosDstEntryPortSflow(d []interface{}) edpt.DdosDstEntryPortSflow {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortSflow
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Polling = getObjectDdosDstEntryPortSflowPolling(in["polling"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortSflowPolling(d []interface{}) edpt.DdosDstEntryPortSflowPolling {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortSflowPolling
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowPackets = in["sflow_packets"].(int)
		ret.SflowTcp = getObjectDdosDstEntryPortSflowPollingSflowTcp(in["sflow_tcp"].([]interface{}))
		ret.SflowHttp = in["sflow_http"].(int)
	}
	return ret
}

func getObjectDdosDstEntryPortSflowPollingSflowTcp(d []interface{}) edpt.DdosDstEntryPortSflowPollingSflowTcp {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortSflowPollingSflowTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowTcpBasic = in["sflow_tcp_basic"].(int)
		ret.SflowTcpStateful = in["sflow_tcp_stateful"].(int)
	}
	return ret
}

func getObjectDdosDstEntryPortSignatureExtraction176(d []interface{}) edpt.DdosDstEntryPortSignatureExtraction176 {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortSignatureExtraction176
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Algorithm = in["algorithm"].(string)
		ret.ManualMode = in["manual_mode"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDstEntryPortTemplate(d []interface{}) edpt.DdosDstEntryPortTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
	}
	return ret
}

func getObjectDdosDstEntryPortTopkSources177(d []interface{}) edpt.DdosDstEntryPortTopkSources177 {

	var ret edpt.DdosDstEntryPortTopkSources177
	return ret
}

func dataToEndpointDdosDstEntryPort(d *schema.ResourceData) edpt.DdosDstEntryPort {
	var ret edpt.DdosDstEntryPort
	ret.Inst.CaptureConfig = getObjectDdosDstEntryPortCaptureConfig(d.Get("capture_config").([]interface{}))
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.DetectionEnable = d.Get("detection_enable").(int)
	ret.Inst.DnsCache = d.Get("dns_cache").(string)
	ret.Inst.EnableTopK = d.Get("enable_top_k").(int)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.GlidExceedAction = getObjectDdosDstEntryPortGlidExceedAction(d.Get("glid_exceed_action").([]interface{}))
	ret.Inst.IpFilteringPolicy = d.Get("ip_filtering_policy").(string)
	ret.Inst.IpFilteringPolicyOper = getObjectDdosDstEntryPortIpFilteringPolicyOper170(d.Get("ip_filtering_policy_oper").([]interface{}))
	ret.Inst.PatternRecognition = getObjectDdosDstEntryPortPatternRecognition171(d.Get("pattern_recognition").([]interface{}))
	ret.Inst.PatternRecognitionPuDetails = getObjectDdosDstEntryPortPatternRecognitionPuDetails172(d.Get("pattern_recognition_pu_details").([]interface{}))
	ret.Inst.PortInd = getObjectDdosDstEntryPortPortInd173(d.Get("port_ind").([]interface{}))
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.ProgressionTracking = getObjectDdosDstEntryPortProgressionTracking175(d.Get("progression_tracking").([]interface{}))
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SetCounterBaseVal = d.Get("set_counter_base_val").(int)
	ret.Inst.Sflow = getObjectDdosDstEntryPortSflow(d.Get("sflow").([]interface{}))
	ret.Inst.SignatureExtraction = getObjectDdosDstEntryPortSignatureExtraction176(d.Get("signature_extraction").([]interface{}))
	ret.Inst.Template = getObjectDdosDstEntryPortTemplate(d.Get("template").([]interface{}))
	ret.Inst.TopkNumRecords = d.Get("topk_num_records").(int)
	ret.Inst.TopkSources = getObjectDdosDstEntryPortTopkSources177(d.Get("topk_sources").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
