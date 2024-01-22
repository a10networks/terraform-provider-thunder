package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryPortRange() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_port_range`: DDOS Port-Range & Protocol configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntryPortRangeCreate,
		UpdateContext: resourceDdosDstEntryPortRangeUpdate,
		ReadContext:   resourceDdosDstEntryPortRangeRead,
		DeleteContext: resourceDdosDstEntryPortRangeDelete,

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
			"port_range_end": {
				Type: schema.TypeInt, Required: true, Description: "Port-Range End Port Number",
			},
			"port_range_start": {
				Type: schema.TypeInt, Required: true, Description: "Port-Range Start Port Number",
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
func resourceDdosDstEntryPortRangeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortRangeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortRange(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryPortRangeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntryPortRangeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortRangeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortRange(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryPortRangeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntryPortRangeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortRangeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortRange(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntryPortRangeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortRangeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortRange(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstEntryPortRangeCaptureConfig(d []interface{}) edpt.DdosDstEntryPortRangeCaptureConfig {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeCaptureConfig
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CaptureConfigName = in["capture_config_name"].(string)
		ret.CaptureConfigMode = in["capture_config_mode"].(string)
	}
	return ret
}

func getObjectDdosDstEntryPortRangeGlidExceedAction(d []interface{}) edpt.DdosDstEntryPortRangeGlidExceedAction {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeGlidExceedAction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapActionCfg = getObjectDdosDstEntryPortRangeGlidExceedActionStatelessEncapActionCfg(in["stateless_encap_action_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortRangeGlidExceedActionStatelessEncapActionCfg(d []interface{}) edpt.DdosDstEntryPortRangeGlidExceedActionStatelessEncapActionCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeGlidExceedActionStatelessEncapActionCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapAction = in["stateless_encap_action"].(string)
		ret.EncapTemplate = in["encap_template"].(string)
	}
	return ret
}

func getObjectDdosDstEntryPortRangeIpFilteringPolicyOper163(d []interface{}) edpt.DdosDstEntryPortRangeIpFilteringPolicyOper163 {

	var ret edpt.DdosDstEntryPortRangeIpFilteringPolicyOper163
	return ret
}

func getObjectDdosDstEntryPortRangePatternRecognition164(d []interface{}) edpt.DdosDstEntryPortRangePatternRecognition164 {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangePatternRecognition164
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

func getObjectDdosDstEntryPortRangePatternRecognitionPuDetails165(d []interface{}) edpt.DdosDstEntryPortRangePatternRecognitionPuDetails165 {

	var ret edpt.DdosDstEntryPortRangePatternRecognitionPuDetails165
	return ret
}

func getObjectDdosDstEntryPortRangePortInd166(d []interface{}) edpt.DdosDstEntryPortRangePortInd166 {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangePortInd166
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceDdosDstEntryPortRangePortIndSamplingEnable167(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryPortRangePortIndSamplingEnable167(d []interface{}) []edpt.DdosDstEntryPortRangePortIndSamplingEnable167 {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangePortIndSamplingEnable167, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangePortIndSamplingEnable167
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryPortRangeProgressionTracking168(d []interface{}) edpt.DdosDstEntryPortRangeProgressionTracking168 {

	var ret edpt.DdosDstEntryPortRangeProgressionTracking168
	return ret
}

func getObjectDdosDstEntryPortRangeSflow(d []interface{}) edpt.DdosDstEntryPortRangeSflow {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeSflow
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Polling = getObjectDdosDstEntryPortRangeSflowPolling(in["polling"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryPortRangeSflowPolling(d []interface{}) edpt.DdosDstEntryPortRangeSflowPolling {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeSflowPolling
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowPackets = in["sflow_packets"].(int)
		ret.SflowTcp = getObjectDdosDstEntryPortRangeSflowPollingSflowTcp(in["sflow_tcp"].([]interface{}))
		ret.SflowHttp = in["sflow_http"].(int)
	}
	return ret
}

func getObjectDdosDstEntryPortRangeSflowPollingSflowTcp(d []interface{}) edpt.DdosDstEntryPortRangeSflowPollingSflowTcp {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeSflowPollingSflowTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowTcpBasic = in["sflow_tcp_basic"].(int)
		ret.SflowTcpStateful = in["sflow_tcp_stateful"].(int)
	}
	return ret
}

func getObjectDdosDstEntryPortRangeTemplate(d []interface{}) edpt.DdosDstEntryPortRangeTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeTemplate
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

func getObjectDdosDstEntryPortRangeTopkSources169(d []interface{}) edpt.DdosDstEntryPortRangeTopkSources169 {

	var ret edpt.DdosDstEntryPortRangeTopkSources169
	return ret
}

func dataToEndpointDdosDstEntryPortRange(d *schema.ResourceData) edpt.DdosDstEntryPortRange {
	var ret edpt.DdosDstEntryPortRange
	ret.Inst.CaptureConfig = getObjectDdosDstEntryPortRangeCaptureConfig(d.Get("capture_config").([]interface{}))
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.DetectionEnable = d.Get("detection_enable").(int)
	ret.Inst.EnableTopK = d.Get("enable_top_k").(int)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.GlidExceedAction = getObjectDdosDstEntryPortRangeGlidExceedAction(d.Get("glid_exceed_action").([]interface{}))
	ret.Inst.IpFilteringPolicy = d.Get("ip_filtering_policy").(string)
	ret.Inst.IpFilteringPolicyOper = getObjectDdosDstEntryPortRangeIpFilteringPolicyOper163(d.Get("ip_filtering_policy_oper").([]interface{}))
	ret.Inst.PatternRecognition = getObjectDdosDstEntryPortRangePatternRecognition164(d.Get("pattern_recognition").([]interface{}))
	ret.Inst.PatternRecognitionPuDetails = getObjectDdosDstEntryPortRangePatternRecognitionPuDetails165(d.Get("pattern_recognition_pu_details").([]interface{}))
	ret.Inst.PortInd = getObjectDdosDstEntryPortRangePortInd166(d.Get("port_ind").([]interface{}))
	ret.Inst.PortRangeEnd = d.Get("port_range_end").(int)
	ret.Inst.PortRangeStart = d.Get("port_range_start").(int)
	ret.Inst.ProgressionTracking = getObjectDdosDstEntryPortRangeProgressionTracking168(d.Get("progression_tracking").([]interface{}))
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SetCounterBaseVal = d.Get("set_counter_base_val").(int)
	ret.Inst.Sflow = getObjectDdosDstEntryPortRangeSflow(d.Get("sflow").([]interface{}))
	ret.Inst.Template = getObjectDdosDstEntryPortRangeTemplate(d.Get("template").([]interface{}))
	ret.Inst.TopkNumRecords = d.Get("topk_num_records").(int)
	ret.Inst.TopkSources = getObjectDdosDstEntryPortRangeTopkSources169(d.Get("topk_sources").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
