package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstDynamicEntryOverflowPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_dynamic_entry_overflow_policy`: Configure IP/IPv6 Policy Used When Dynamic Dst Entry Count overflows\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstDynamicEntryOverflowPolicyCreate,
		UpdateContext: resourceDdosDstDynamicEntryOverflowPolicyUpdate,
		ReadContext:   resourceDdosDstDynamicEntryOverflowPolicyRead,
		DeleteContext: resourceDdosDstDynamicEntryOverflowPolicyDelete,

		Schema: map[string]*schema.Schema{
			"default_address_type": {
				Type: schema.TypeString, Required: true, Description: "'ip': ip; 'ipv6': ipv6;",
			},
			"drop_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable certain drops during packet processing",
			},
			"drop_disable_fwd_immediate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Immediately forward L4 drops",
			},
			"exceed_log_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging of limit exceed drop's",
						},
						"with_sflow_sample": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Turn on sflow sample with log",
						},
					},
				},
			},
			"exceed_log_dep_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exceed_log_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "(Deprecated)Enable logging of limit exceed drop's",
						},
						"log_with_sflow_dep": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Turn on sflow sample with log",
						},
					},
				},
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"inbound_forward_dscp": {
				Type: schema.TypeInt, Optional: true, Description: "To set dscp value for inbound packets (DSCP Value for the clear traffic marking)",
			},
			"ip_proto_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Protocol Number",
						},
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"template": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"other": {
										Type: schema.TypeString, Optional: true, Description: "DDOS other template",
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
			"l4_type_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
						},
						"max_rexmit_syn_per_flow": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum number of re-transmit SYN per flow. Exceed action set to Drop",
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
						"drop_on_no_port_match": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'disable': disable; 'enable': enable;",
						},
						"stateful": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable stateful tracking of sessions (Default is stateless)",
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
						"drop_frag_pkt": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop fragmented packets",
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
			"log_periodic": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable periodic log while event is continuing",
			},
			"outbound_forward_dscp": {
				Type: schema.TypeInt, Optional: true, Description: "To set dscp value for outbound",
			},
			"port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'dns-tcp': dns-tcp; 'dns-udp': dns-udp; 'http': http; 'tcp': tcp; 'udp': udp; 'ssl-l4': ssl-l4; 'sip-udp': sip-udp; 'sip-tcp': sip-tcp;",
						},
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
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
										Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"src_port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'udp': udp; 'tcp': tcp;",
						},
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"template": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_udp": {
										Type: schema.TypeString, Optional: true, Description: "DDOS udp src template",
									},
									"src_tcp": {
										Type: schema.TypeString, Optional: true, Description: "DDOS tcp src template",
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
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logging": {
							Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
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
		},
	}
}
func resourceDdosDstDynamicEntryOverflowPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDynamicEntryOverflowPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDynamicEntryOverflowPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstDynamicEntryOverflowPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstDynamicEntryOverflowPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDynamicEntryOverflowPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDynamicEntryOverflowPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstDynamicEntryOverflowPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstDynamicEntryOverflowPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDynamicEntryOverflowPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDynamicEntryOverflowPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstDynamicEntryOverflowPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDynamicEntryOverflowPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDynamicEntryOverflowPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstDynamicEntryOverflowPolicyExceedLogCfg(d []interface{}) edpt.DdosDstDynamicEntryOverflowPolicyExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosDstDynamicEntryOverflowPolicyExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
		ret.WithSflowSample = in["with_sflow_sample"].(int)
	}
	return ret
}

func getObjectDdosDstDynamicEntryOverflowPolicyExceedLogDepCfg(d []interface{}) edpt.DdosDstDynamicEntryOverflowPolicyExceedLogDepCfg {

	count1 := len(d)
	var ret edpt.DdosDstDynamicEntryOverflowPolicyExceedLogDepCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ExceedLogEnable = in["exceed_log_enable"].(int)
		ret.LogWithSflowDep = in["log_with_sflow_dep"].(int)
	}
	return ret
}

func getSliceDdosDstDynamicEntryOverflowPolicyIpProtoList(d []interface{}) []edpt.DdosDstDynamicEntryOverflowPolicyIpProtoList {

	count1 := len(d)
	ret := make([]edpt.DdosDstDynamicEntryOverflowPolicyIpProtoList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstDynamicEntryOverflowPolicyIpProtoList
		oi.PortNum = in["port_num"].(int)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstDynamicEntryOverflowPolicyIpProtoListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstDynamicEntryOverflowPolicyIpProtoListTemplate(d []interface{}) edpt.DdosDstDynamicEntryOverflowPolicyIpProtoListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstDynamicEntryOverflowPolicyIpProtoListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Other = in["other"].(string)
	}
	return ret
}

func getSliceDdosDstDynamicEntryOverflowPolicyL4TypeList(d []interface{}) []edpt.DdosDstDynamicEntryOverflowPolicyL4TypeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstDynamicEntryOverflowPolicyL4TypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstDynamicEntryOverflowPolicyL4TypeList
		oi.Protocol = in["protocol"].(string)
		oi.Glid = in["glid"].(string)
		oi.Deny = in["deny"].(int)
		oi.MaxRexmitSynPerFlow = in["max_rexmit_syn_per_flow"].(int)
		oi.SynAuth = in["syn_auth"].(string)
		oi.SynCookie = in["syn_cookie"].(int)
		oi.TcpResetClient = in["tcp_reset_client"].(int)
		oi.TcpResetServer = in["tcp_reset_server"].(int)
		oi.DropOnNoPortMatch = in["drop_on_no_port_match"].(string)
		oi.Stateful = in["stateful"].(int)
		oi.TunnelDecap = getObjectDdosDstDynamicEntryOverflowPolicyL4TypeListTunnelDecap(in["tunnel_decap"].([]interface{}))
		oi.TunnelRateLimit = getObjectDdosDstDynamicEntryOverflowPolicyL4TypeListTunnelRateLimit(in["tunnel_rate_limit"].([]interface{}))
		oi.DropFragPkt = in["drop_frag_pkt"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstDynamicEntryOverflowPolicyL4TypeListTunnelDecap(d []interface{}) edpt.DdosDstDynamicEntryOverflowPolicyL4TypeListTunnelDecap {

	count1 := len(d)
	var ret edpt.DdosDstDynamicEntryOverflowPolicyL4TypeListTunnelDecap
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpDecap = in["ip_decap"].(int)
		ret.GreDecap = in["gre_decap"].(int)
		ret.KeyCfg = getSliceDdosDstDynamicEntryOverflowPolicyL4TypeListTunnelDecapKeyCfg(in["key_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstDynamicEntryOverflowPolicyL4TypeListTunnelDecapKeyCfg(d []interface{}) []edpt.DdosDstDynamicEntryOverflowPolicyL4TypeListTunnelDecapKeyCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstDynamicEntryOverflowPolicyL4TypeListTunnelDecapKeyCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstDynamicEntryOverflowPolicyL4TypeListTunnelDecapKeyCfg
		oi.Key = in["key"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstDynamicEntryOverflowPolicyL4TypeListTunnelRateLimit(d []interface{}) edpt.DdosDstDynamicEntryOverflowPolicyL4TypeListTunnelRateLimit {

	count1 := len(d)
	var ret edpt.DdosDstDynamicEntryOverflowPolicyL4TypeListTunnelRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpRateLimit = in["ip_rate_limit"].(int)
		ret.GreRateLimit = in["gre_rate_limit"].(int)
	}
	return ret
}

func getSliceDdosDstDynamicEntryOverflowPolicyPortList(d []interface{}) []edpt.DdosDstDynamicEntryOverflowPolicyPortList {

	count1 := len(d)
	ret := make([]edpt.DdosDstDynamicEntryOverflowPolicyPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstDynamicEntryOverflowPolicyPortList
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstDynamicEntryOverflowPolicyPortListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstDynamicEntryOverflowPolicyPortListTemplate(d []interface{}) edpt.DdosDstDynamicEntryOverflowPolicyPortListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstDynamicEntryOverflowPolicyPortListTemplate
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

func getSliceDdosDstDynamicEntryOverflowPolicySrcPortList(d []interface{}) []edpt.DdosDstDynamicEntryOverflowPolicySrcPortList {

	count1 := len(d)
	ret := make([]edpt.DdosDstDynamicEntryOverflowPolicySrcPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstDynamicEntryOverflowPolicySrcPortList
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstDynamicEntryOverflowPolicySrcPortListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstDynamicEntryOverflowPolicySrcPortListTemplate(d []interface{}) edpt.DdosDstDynamicEntryOverflowPolicySrcPortListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstDynamicEntryOverflowPolicySrcPortListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcUdp = in["src_udp"].(string)
		ret.SrcTcp = in["src_tcp"].(string)
	}
	return ret
}

func getObjectDdosDstDynamicEntryOverflowPolicyTemplate(d []interface{}) edpt.DdosDstDynamicEntryOverflowPolicyTemplate {

	count1 := len(d)
	var ret edpt.DdosDstDynamicEntryOverflowPolicyTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func dataToEndpointDdosDstDynamicEntryOverflowPolicy(d *schema.ResourceData) edpt.DdosDstDynamicEntryOverflowPolicy {
	var ret edpt.DdosDstDynamicEntryOverflowPolicy
	ret.Inst.DefaultAddressType = d.Get("default_address_type").(string)
	ret.Inst.DropDisable = d.Get("drop_disable").(int)
	ret.Inst.DropDisableFwdImmediate = d.Get("drop_disable_fwd_immediate").(int)
	ret.Inst.ExceedLogCfg = getObjectDdosDstDynamicEntryOverflowPolicyExceedLogCfg(d.Get("exceed_log_cfg").([]interface{}))
	ret.Inst.ExceedLogDepCfg = getObjectDdosDstDynamicEntryOverflowPolicyExceedLogDepCfg(d.Get("exceed_log_dep_cfg").([]interface{}))
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.InboundForwardDscp = d.Get("inbound_forward_dscp").(int)
	ret.Inst.IpProtoList = getSliceDdosDstDynamicEntryOverflowPolicyIpProtoList(d.Get("ip_proto_list").([]interface{}))
	ret.Inst.L4TypeList = getSliceDdosDstDynamicEntryOverflowPolicyL4TypeList(d.Get("l4_type_list").([]interface{}))
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.OutboundForwardDscp = d.Get("outbound_forward_dscp").(int)
	ret.Inst.PortList = getSliceDdosDstDynamicEntryOverflowPolicyPortList(d.Get("port_list").([]interface{}))
	ret.Inst.SrcPortList = getSliceDdosDstDynamicEntryOverflowPolicySrcPortList(d.Get("src_port_list").([]interface{}))
	ret.Inst.Template = getObjectDdosDstDynamicEntryOverflowPolicyTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
