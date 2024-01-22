package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryL4Type() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_l4_type`: DDOS L4 type\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntryL4TypeCreate,
		UpdateContext: resourceDdosDstEntryL4TypeUpdate,
		ReadContext:   resourceDdosDstEntryL4TypeRead,
		DeleteContext: resourceDdosDstEntryL4TypeDelete,

		Schema: map[string]*schema.Schema{
			"deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
			},
			"detection_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos detection",
			},
			"disable_syn_auth": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable TCP SYN Authentication",
			},
			"drop_frag_pkt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop fragmented packets",
			},
			"drop_on_no_port_match": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'disable': disable; 'enable': enable;",
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
			"max_rexmit_syn_per_flow": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum number of re-transmit SYN per flow",
			},
			"max_rexmit_syn_per_flow_exceed_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop the packet; 'black-list': Add the source IP into black list;",
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
				Type: schema.TypeString, Required: true, Description: "'tcp': L4-Type TCP; 'udp': L4-Type UDP; 'icmp': L4-Type ICMP; 'other': L4-Type OTHER;",
			},
			"set_counter_base_val": {
				Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
			},
			"stateful": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable stateful tracking of sessions (Default is stateless)",
			},
			"syn_auth": {
				Type: schema.TypeString, Optional: true, Default: "send-rst", Description: "'send-rst': Send RST to client upon client ACK; 'force-rst-by-ack': Force client RST via the use of ACK; 'force-rst-by-synack': Force client RST via the use of bad SYN|ACK; 'disable': Disable TCP SYN Authentication;",
			},
			"syn_cookie": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SYN Cookie",
			},
			"tcp_reset_client": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send reset to client when rate exceeds or session ages out",
			},
			"tcp_reset_server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send reset to server when rate exceeds or session ages out",
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template_icmp_v4": {
							Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
						},
						"template_icmp_v6": {
							Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
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
			"tunnel_decap": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_decap": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IP Tunnel decapsulation",
						},
						"gre_decap": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable GRE Tunnel decapsulation",
						},
						"key_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type: schema.TypeString, Optional: true, Description: "Only decapsulate GRE packet with this key (Hexadecimal 0x0-0xFFFFFFFF,decimal 0-4294967295)",
									},
								},
							},
						},
					},
				},
			},
			"tunnel_rate_limit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_rate_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable inner IP rate limiting on IPinIP traffic",
						},
						"gre_rate_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable inner IP rate limiting on GRE traffic",
						},
					},
				},
			},
			"undefined_port_hit_statistics": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"undefined_port_hit_statistics": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable port scanning statistics",
						},
						"reset_interval": {
							Type: schema.TypeInt, Optional: true, Default: 60, Description: "Configure port scanning counter reset interval (minutes), Default 60 mins",
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
func resourceDdosDstEntryL4TypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryL4TypeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryL4Type(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryL4TypeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntryL4TypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryL4TypeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryL4Type(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryL4TypeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntryL4TypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryL4TypeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryL4Type(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntryL4TypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryL4TypeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryL4Type(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstEntryL4TypeGlidExceedAction(d []interface{}) edpt.DdosDstEntryL4TypeGlidExceedAction {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeGlidExceedAction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapActionCfg = getObjectDdosDstEntryL4TypeGlidExceedActionStatelessEncapActionCfg(in["stateless_encap_action_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryL4TypeGlidExceedActionStatelessEncapActionCfg(d []interface{}) edpt.DdosDstEntryL4TypeGlidExceedActionStatelessEncapActionCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeGlidExceedActionStatelessEncapActionCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapAction = in["stateless_encap_action"].(string)
		ret.EncapTemplate = in["encap_template"].(string)
	}
	return ret
}

func getObjectDdosDstEntryL4TypeIpFilteringPolicyOper158(d []interface{}) edpt.DdosDstEntryL4TypeIpFilteringPolicyOper158 {

	var ret edpt.DdosDstEntryL4TypeIpFilteringPolicyOper158
	return ret
}

func getObjectDdosDstEntryL4TypePortInd159(d []interface{}) edpt.DdosDstEntryL4TypePortInd159 {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypePortInd159
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceDdosDstEntryL4TypePortIndSamplingEnable160(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryL4TypePortIndSamplingEnable160(d []interface{}) []edpt.DdosDstEntryL4TypePortIndSamplingEnable160 {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryL4TypePortIndSamplingEnable160, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryL4TypePortIndSamplingEnable160
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryL4TypeProgressionTracking161(d []interface{}) edpt.DdosDstEntryL4TypeProgressionTracking161 {

	var ret edpt.DdosDstEntryL4TypeProgressionTracking161
	return ret
}

func getObjectDdosDstEntryL4TypeTemplate(d []interface{}) edpt.DdosDstEntryL4TypeTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TemplateIcmpV4 = in["template_icmp_v4"].(string)
		ret.TemplateIcmpV6 = in["template_icmp_v6"].(string)
	}
	return ret
}

func getObjectDdosDstEntryL4TypeTopkSources162(d []interface{}) edpt.DdosDstEntryL4TypeTopkSources162 {

	var ret edpt.DdosDstEntryL4TypeTopkSources162
	return ret
}

func getObjectDdosDstEntryL4TypeTunnelDecap(d []interface{}) edpt.DdosDstEntryL4TypeTunnelDecap {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeTunnelDecap
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpDecap = in["ip_decap"].(int)
		ret.GreDecap = in["gre_decap"].(int)
		ret.KeyCfg = getSliceDdosDstEntryL4TypeTunnelDecapKeyCfg(in["key_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryL4TypeTunnelDecapKeyCfg(d []interface{}) []edpt.DdosDstEntryL4TypeTunnelDecapKeyCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryL4TypeTunnelDecapKeyCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryL4TypeTunnelDecapKeyCfg
		oi.Key = in["key"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryL4TypeTunnelRateLimit(d []interface{}) edpt.DdosDstEntryL4TypeTunnelRateLimit {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeTunnelRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpRateLimit = in["ip_rate_limit"].(int)
		ret.GreRateLimit = in["gre_rate_limit"].(int)
	}
	return ret
}

func getObjectDdosDstEntryL4TypeUndefinedPortHitStatistics(d []interface{}) edpt.DdosDstEntryL4TypeUndefinedPortHitStatistics {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeUndefinedPortHitStatistics
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UndefinedPortHitStatistics = in["undefined_port_hit_statistics"].(int)
		ret.ResetInterval = in["reset_interval"].(int)
	}
	return ret
}

func dataToEndpointDdosDstEntryL4Type(d *schema.ResourceData) edpt.DdosDstEntryL4Type {
	var ret edpt.DdosDstEntryL4Type
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.DetectionEnable = d.Get("detection_enable").(int)
	ret.Inst.DisableSynAuth = d.Get("disable_syn_auth").(int)
	ret.Inst.DropFragPkt = d.Get("drop_frag_pkt").(int)
	ret.Inst.DropOnNoPortMatch = d.Get("drop_on_no_port_match").(string)
	ret.Inst.EnableTopK = d.Get("enable_top_k").(int)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.GlidExceedAction = getObjectDdosDstEntryL4TypeGlidExceedAction(d.Get("glid_exceed_action").([]interface{}))
	ret.Inst.IpFilteringPolicy = d.Get("ip_filtering_policy").(string)
	ret.Inst.IpFilteringPolicyOper = getObjectDdosDstEntryL4TypeIpFilteringPolicyOper158(d.Get("ip_filtering_policy_oper").([]interface{}))
	ret.Inst.MaxRexmitSynPerFlow = d.Get("max_rexmit_syn_per_flow").(int)
	ret.Inst.MaxRexmitSynPerFlowExceedAction = d.Get("max_rexmit_syn_per_flow_exceed_action").(string)
	ret.Inst.PortInd = getObjectDdosDstEntryL4TypePortInd159(d.Get("port_ind").([]interface{}))
	ret.Inst.ProgressionTracking = getObjectDdosDstEntryL4TypeProgressionTracking161(d.Get("progression_tracking").([]interface{}))
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SetCounterBaseVal = d.Get("set_counter_base_val").(int)
	ret.Inst.Stateful = d.Get("stateful").(int)
	ret.Inst.SynAuth = d.Get("syn_auth").(string)
	ret.Inst.SynCookie = d.Get("syn_cookie").(int)
	ret.Inst.TcpResetClient = d.Get("tcp_reset_client").(int)
	ret.Inst.TcpResetServer = d.Get("tcp_reset_server").(int)
	ret.Inst.Template = getObjectDdosDstEntryL4TypeTemplate(d.Get("template").([]interface{}))
	ret.Inst.TopkNumRecords = d.Get("topk_num_records").(int)
	ret.Inst.TopkSources = getObjectDdosDstEntryL4TypeTopkSources162(d.Get("topk_sources").([]interface{}))
	ret.Inst.TunnelDecap = getObjectDdosDstEntryL4TypeTunnelDecap(d.Get("tunnel_decap").([]interface{}))
	ret.Inst.TunnelRateLimit = getObjectDdosDstEntryL4TypeTunnelRateLimit(d.Get("tunnel_rate_limit").([]interface{}))
	ret.Inst.UndefinedPortHitStatistics = getObjectDdosDstEntryL4TypeUndefinedPortHitStatistics(d.Get("undefined_port_hit_statistics").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
