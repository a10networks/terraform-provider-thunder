package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstDefaultL4Type() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_default_l4_type`: DDOS L4 type\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstDefaultL4TypeCreate,
		UpdateContext: resourceDdosDstDefaultL4TypeUpdate,
		ReadContext:   resourceDdosDstDefaultL4TypeRead,
		DeleteContext: resourceDdosDstDefaultL4TypeDelete,

		Schema: map[string]*schema.Schema{
			"deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
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
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"max_rexmit_syn_per_flow": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum number of re-transmit SYN per flow. Exceed action set to Drop",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
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
			"default_address_type": {
				Type: schema.TypeString, Required: true, Description: "DefaultAddressType",
			},
		},
	}
}
func resourceDdosDstDefaultL4TypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDefaultL4TypeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDefaultL4Type(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstDefaultL4TypeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstDefaultL4TypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDefaultL4TypeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDefaultL4Type(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstDefaultL4TypeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstDefaultL4TypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDefaultL4TypeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDefaultL4Type(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstDefaultL4TypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDefaultL4TypeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDefaultL4Type(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstDefaultL4TypeTunnelDecap(d []interface{}) edpt.DdosDstDefaultL4TypeTunnelDecap {

	count1 := len(d)
	var ret edpt.DdosDstDefaultL4TypeTunnelDecap
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpDecap = in["ip_decap"].(int)
		ret.GreDecap = in["gre_decap"].(int)
		ret.KeyCfg = getSliceDdosDstDefaultL4TypeTunnelDecapKeyCfg(in["key_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstDefaultL4TypeTunnelDecapKeyCfg(d []interface{}) []edpt.DdosDstDefaultL4TypeTunnelDecapKeyCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstDefaultL4TypeTunnelDecapKeyCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstDefaultL4TypeTunnelDecapKeyCfg
		oi.Key = in["key"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstDefaultL4TypeTunnelRateLimit(d []interface{}) edpt.DdosDstDefaultL4TypeTunnelRateLimit {

	count1 := len(d)
	var ret edpt.DdosDstDefaultL4TypeTunnelRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpRateLimit = in["ip_rate_limit"].(int)
		ret.GreRateLimit = in["gre_rate_limit"].(int)
	}
	return ret
}

func dataToEndpointDdosDstDefaultL4Type(d *schema.ResourceData) edpt.DdosDstDefaultL4Type {
	var ret edpt.DdosDstDefaultL4Type
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.DisableSynAuth = d.Get("disable_syn_auth").(int)
	ret.Inst.DropFragPkt = d.Get("drop_frag_pkt").(int)
	ret.Inst.DropOnNoPortMatch = d.Get("drop_on_no_port_match").(string)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.MaxRexmitSynPerFlow = d.Get("max_rexmit_syn_per_flow").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Stateful = d.Get("stateful").(int)
	ret.Inst.SynAuth = d.Get("syn_auth").(string)
	ret.Inst.SynCookie = d.Get("syn_cookie").(int)
	ret.Inst.TcpResetClient = d.Get("tcp_reset_client").(int)
	ret.Inst.TcpResetServer = d.Get("tcp_reset_server").(int)
	ret.Inst.TunnelDecap = getObjectDdosDstDefaultL4TypeTunnelDecap(d.Get("tunnel_decap").([]interface{}))
	ret.Inst.TunnelRateLimit = getObjectDdosDstDefaultL4TypeTunnelRateLimit(d.Get("tunnel_rate_limit").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.DefaultAddressType = d.Get("default_address_type").(string)
	return ret
}
