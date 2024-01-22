package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstInterfaceIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_interface_ipv6`: Configure protections on local interface IPs\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstInterfaceIpv6Create,
		UpdateContext: resourceDdosDstInterfaceIpv6Update,
		ReadContext:   resourceDdosDstInterfaceIpv6Read,
		DeleteContext: resourceDdosDstInterfaceIpv6Delete,

		Schema: map[string]*schema.Schema{
			"addr": {
				Type: schema.TypeString, Required: true, Description: "IPv6 address of interface",
			},
			"ip_proto_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "IP protocol number",
						},
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
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
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
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
			"log_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging of limit exceed drops",
			},
			"port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'http-probe': http port for interface health check;",
						},
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosDstInterfaceIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstInterfaceIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstInterfaceIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstInterfaceIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstInterfaceIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstInterfaceIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstInterfaceIpv6IpProtoList(d []interface{}) []edpt.DdosDstInterfaceIpv6IpProtoList {

	count1 := len(d)
	ret := make([]edpt.DdosDstInterfaceIpv6IpProtoList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstInterfaceIpv6IpProtoList
		oi.PortNum = in["port_num"].(int)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstInterfaceIpv6L4TypeList(d []interface{}) []edpt.DdosDstInterfaceIpv6L4TypeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstInterfaceIpv6L4TypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstInterfaceIpv6L4TypeList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.TunnelDecap = getObjectDdosDstInterfaceIpv6L4TypeListTunnelDecap(in["tunnel_decap"].([]interface{}))
		oi.TunnelRateLimit = getObjectDdosDstInterfaceIpv6L4TypeListTunnelRateLimit(in["tunnel_rate_limit"].([]interface{}))
		oi.DropFragPkt = in["drop_frag_pkt"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstInterfaceIpv6L4TypeListTunnelDecap(d []interface{}) edpt.DdosDstInterfaceIpv6L4TypeListTunnelDecap {

	count1 := len(d)
	var ret edpt.DdosDstInterfaceIpv6L4TypeListTunnelDecap
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpDecap = in["ip_decap"].(int)
		ret.GreDecap = in["gre_decap"].(int)
		ret.KeyCfg = getSliceDdosDstInterfaceIpv6L4TypeListTunnelDecapKeyCfg(in["key_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstInterfaceIpv6L4TypeListTunnelDecapKeyCfg(d []interface{}) []edpt.DdosDstInterfaceIpv6L4TypeListTunnelDecapKeyCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstInterfaceIpv6L4TypeListTunnelDecapKeyCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstInterfaceIpv6L4TypeListTunnelDecapKeyCfg
		oi.Key = in["key"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstInterfaceIpv6L4TypeListTunnelRateLimit(d []interface{}) edpt.DdosDstInterfaceIpv6L4TypeListTunnelRateLimit {

	count1 := len(d)
	var ret edpt.DdosDstInterfaceIpv6L4TypeListTunnelRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpRateLimit = in["ip_rate_limit"].(int)
		ret.GreRateLimit = in["gre_rate_limit"].(int)
	}
	return ret
}

func getSliceDdosDstInterfaceIpv6PortList(d []interface{}) []edpt.DdosDstInterfaceIpv6PortList {

	count1 := len(d)
	ret := make([]edpt.DdosDstInterfaceIpv6PortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstInterfaceIpv6PortList
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstInterfaceIpv6(d *schema.ResourceData) edpt.DdosDstInterfaceIpv6 {
	var ret edpt.DdosDstInterfaceIpv6
	ret.Inst.Addr = d.Get("addr").(string)
	ret.Inst.IpProtoList = getSliceDdosDstInterfaceIpv6IpProtoList(d.Get("ip_proto_list").([]interface{}))
	ret.Inst.L4TypeList = getSliceDdosDstInterfaceIpv6L4TypeList(d.Get("l4_type_list").([]interface{}))
	ret.Inst.LogEnable = d.Get("log_enable").(int)
	ret.Inst.PortList = getSliceDdosDstInterfaceIpv6PortList(d.Get("port_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
