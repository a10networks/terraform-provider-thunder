package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateEncap() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_encap`: Encapsulation template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateEncapCreate,
		UpdateContext: resourceDdosTemplateEncapUpdate,
		ReadContext:   resourceDdosTemplateEncapRead,
		DeleteContext: resourceDdosTemplateEncapDelete,

		Schema: map[string]*schema.Schema{
			"encap_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "DDOS Tunnel Template Name",
			},
			"preserve_source_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use original source ip for encapsulation",
			},
			"tunnel_encap": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_encap": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Tunnel encap for IP packets",
									},
									"always": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv4_addr": {
													Type: schema.TypeString, Optional: true, Description: "IPv4 address (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
												},
												"ipv6_addr": {
													Type: schema.TypeString, Optional: true, Description: "IPv6 address (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
												},
											},
										},
									},
								},
							},
						},
						"gre_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gre_encap": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Tunnel encap for GRE packets",
									},
									"gre_always": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"gre_ipv4": {
													Type: schema.TypeString, Optional: true, Description: "IPv4 address (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
												},
												"key_ipv4": {
													Type: schema.TypeString, Optional: true, Description: "Encapsulate with key (Hexadecimal 0x0-0xFFFFFFFF,decimal 0-4294967295)",
												},
												"gre_ipv6": {
													Type: schema.TypeString, Optional: true, Description: "IPv6 address (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
												},
												"key_ipv6": {
													Type: schema.TypeString, Optional: true, Description: "Encapsulate with key (Hexadecimal 0x0-0xFFFFFFFF,decimal 0-4294967295)",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosTemplateEncapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateEncapCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateEncap(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateEncapRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateEncapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateEncapUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateEncap(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateEncapRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateEncapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateEncapDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateEncap(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateEncapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateEncapRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateEncap(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosTemplateEncapTunnelEncap(d []interface{}) edpt.DdosTemplateEncapTunnelEncap {

	count1 := len(d)
	var ret edpt.DdosTemplateEncapTunnelEncap
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpCfg = getObjectDdosTemplateEncapTunnelEncapIpCfg(in["ip_cfg"].([]interface{}))
		ret.GreCfg = getObjectDdosTemplateEncapTunnelEncapGreCfg(in["gre_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateEncapTunnelEncapIpCfg(d []interface{}) edpt.DdosTemplateEncapTunnelEncapIpCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateEncapTunnelEncapIpCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpEncap = in["ip_encap"].(int)
		ret.Always = getObjectDdosTemplateEncapTunnelEncapIpCfgAlways(in["always"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateEncapTunnelEncapIpCfgAlways(d []interface{}) edpt.DdosTemplateEncapTunnelEncapIpCfgAlways {

	count1 := len(d)
	var ret edpt.DdosTemplateEncapTunnelEncapIpCfgAlways
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
	}
	return ret
}

func getObjectDdosTemplateEncapTunnelEncapGreCfg(d []interface{}) edpt.DdosTemplateEncapTunnelEncapGreCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateEncapTunnelEncapGreCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GreEncap = in["gre_encap"].(int)
		ret.GreAlways = getObjectDdosTemplateEncapTunnelEncapGreCfgGreAlways(in["gre_always"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateEncapTunnelEncapGreCfgGreAlways(d []interface{}) edpt.DdosTemplateEncapTunnelEncapGreCfgGreAlways {

	count1 := len(d)
	var ret edpt.DdosTemplateEncapTunnelEncapGreCfgGreAlways
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GreIpv4 = in["gre_ipv4"].(string)
		ret.KeyIpv4 = in["key_ipv4"].(string)
		ret.GreIpv6 = in["gre_ipv6"].(string)
		ret.KeyIpv6 = in["key_ipv6"].(string)
	}
	return ret
}

func dataToEndpointDdosTemplateEncap(d *schema.ResourceData) edpt.DdosTemplateEncap {
	var ret edpt.DdosTemplateEncap
	ret.Inst.EncapTmplName = d.Get("encap_tmpl_name").(string)
	ret.Inst.PreserveSourceIp = d.Get("preserve_source_ip").(int)
	ret.Inst.TunnelEncap = getObjectDdosTemplateEncapTunnelEncap(d.Get("tunnel_encap").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
