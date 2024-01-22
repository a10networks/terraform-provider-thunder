package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstInterfaceIpv6L4Type() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_interface_ipv6_l4_type`: Configure L4 on local interface IPs\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstInterfaceIpv6L4TypeCreate,
		UpdateContext: resourceDdosDstInterfaceIpv6L4TypeUpdate,
		ReadContext:   resourceDdosDstInterfaceIpv6L4TypeRead,
		DeleteContext: resourceDdosDstInterfaceIpv6L4TypeDelete,

		Schema: map[string]*schema.Schema{
			"deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
			},
			"drop_frag_pkt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop fragmented packets",
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"addr": {
				Type: schema.TypeString, Required: true, Description: "Addr",
			},
		},
	}
}
func resourceDdosDstInterfaceIpv6L4TypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpv6L4TypeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIpv6L4Type(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstInterfaceIpv6L4TypeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstInterfaceIpv6L4TypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpv6L4TypeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIpv6L4Type(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstInterfaceIpv6L4TypeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstInterfaceIpv6L4TypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpv6L4TypeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIpv6L4Type(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstInterfaceIpv6L4TypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstInterfaceIpv6L4TypeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstInterfaceIpv6L4Type(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstInterfaceIpv6L4TypeTunnelDecap(d []interface{}) edpt.DdosDstInterfaceIpv6L4TypeTunnelDecap {

	count1 := len(d)
	var ret edpt.DdosDstInterfaceIpv6L4TypeTunnelDecap
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpDecap = in["ip_decap"].(int)
		ret.GreDecap = in["gre_decap"].(int)
		ret.KeyCfg = getSliceDdosDstInterfaceIpv6L4TypeTunnelDecapKeyCfg(in["key_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstInterfaceIpv6L4TypeTunnelDecapKeyCfg(d []interface{}) []edpt.DdosDstInterfaceIpv6L4TypeTunnelDecapKeyCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstInterfaceIpv6L4TypeTunnelDecapKeyCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstInterfaceIpv6L4TypeTunnelDecapKeyCfg
		oi.Key = in["key"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstInterfaceIpv6L4TypeTunnelRateLimit(d []interface{}) edpt.DdosDstInterfaceIpv6L4TypeTunnelRateLimit {

	count1 := len(d)
	var ret edpt.DdosDstInterfaceIpv6L4TypeTunnelRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpRateLimit = in["ip_rate_limit"].(int)
		ret.GreRateLimit = in["gre_rate_limit"].(int)
	}
	return ret
}

func dataToEndpointDdosDstInterfaceIpv6L4Type(d *schema.ResourceData) edpt.DdosDstInterfaceIpv6L4Type {
	var ret edpt.DdosDstInterfaceIpv6L4Type
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.DropFragPkt = d.Get("drop_frag_pkt").(int)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.TunnelDecap = getObjectDdosDstInterfaceIpv6L4TypeTunnelDecap(d.Get("tunnel_decap").([]interface{}))
	ret.Inst.TunnelRateLimit = getObjectDdosDstInterfaceIpv6L4TypeTunnelRateLimit(d.Get("tunnel_rate_limit").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Addr = d.Get("addr").(string)
	return ret
}
