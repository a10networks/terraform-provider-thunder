package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstInterfaceIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_interface_ip`: Configure protections on local interface IPs\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstInterfaceIpCreate,
		UpdateContext: resourceDdosDstInterfaceIpUpdate,
		ReadContext:   resourceDdosDstInterfaceIpRead,
		DeleteContext: resourceDdosDstInterfaceIpDelete,

		Schema: map[string]*schema.Schema{
			"addr": {
				Type: schema.TypeString, Required: true, Description: "IP address of interface",
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
			"log_periodic": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable periodic log while event is continuing",
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
func resourceDdosDstInterfaceIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstInterfaceIpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstInterfaceIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstInterfaceIpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstInterfaceIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstInterfaceIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstInterfaceIpIpProtoList(d []interface{}) []edpt.DdosDstInterfaceIpIpProtoList {

	count1 := len(d)
	ret := make([]edpt.DdosDstInterfaceIpIpProtoList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstInterfaceIpIpProtoList
		oi.PortNum = in["port_num"].(int)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstInterfaceIpL4TypeList(d []interface{}) []edpt.DdosDstInterfaceIpL4TypeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstInterfaceIpL4TypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstInterfaceIpL4TypeList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.TunnelDecap = getObjectDdosDstInterfaceIpL4TypeListTunnelDecap(in["tunnel_decap"].([]interface{}))
		oi.TunnelRateLimit = getObjectDdosDstInterfaceIpL4TypeListTunnelRateLimit(in["tunnel_rate_limit"].([]interface{}))
		oi.DropFragPkt = in["drop_frag_pkt"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstInterfaceIpL4TypeListTunnelDecap(d []interface{}) edpt.DdosDstInterfaceIpL4TypeListTunnelDecap {

	count1 := len(d)
	var ret edpt.DdosDstInterfaceIpL4TypeListTunnelDecap
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpDecap = in["ip_decap"].(int)
		ret.GreDecap = in["gre_decap"].(int)
		ret.KeyCfg = getSliceDdosDstInterfaceIpL4TypeListTunnelDecapKeyCfg(in["key_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstInterfaceIpL4TypeListTunnelDecapKeyCfg(d []interface{}) []edpt.DdosDstInterfaceIpL4TypeListTunnelDecapKeyCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstInterfaceIpL4TypeListTunnelDecapKeyCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstInterfaceIpL4TypeListTunnelDecapKeyCfg
		oi.Key = in["key"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstInterfaceIpL4TypeListTunnelRateLimit(d []interface{}) edpt.DdosDstInterfaceIpL4TypeListTunnelRateLimit {

	count1 := len(d)
	var ret edpt.DdosDstInterfaceIpL4TypeListTunnelRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpRateLimit = in["ip_rate_limit"].(int)
		ret.GreRateLimit = in["gre_rate_limit"].(int)
	}
	return ret
}

func getSliceDdosDstInterfaceIpPortList(d []interface{}) []edpt.DdosDstInterfaceIpPortList {

	count1 := len(d)
	ret := make([]edpt.DdosDstInterfaceIpPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstInterfaceIpPortList
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

func dataToEndpointDdosDstInterfaceIp(d *schema.ResourceData) edpt.DdosDstInterfaceIp {
	var ret edpt.DdosDstInterfaceIp
	ret.Inst.Addr = d.Get("addr").(string)
	ret.Inst.IpProtoList = getSliceDdosDstInterfaceIpIpProtoList(d.Get("ip_proto_list").([]interface{}))
	ret.Inst.L4TypeList = getSliceDdosDstInterfaceIpL4TypeList(d.Get("l4_type_list").([]interface{}))
	ret.Inst.LogEnable = d.Get("log_enable").(int)
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.PortList = getSliceDdosDstInterfaceIpPortList(d.Get("port_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
