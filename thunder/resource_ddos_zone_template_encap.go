package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateEncap() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_encap`: Encapsulation template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateEncapCreate,
		UpdateContext: resourceDdosZoneTemplateEncapUpdate,
		ReadContext:   resourceDdosZoneTemplateEncapRead,
		DeleteContext: resourceDdosZoneTemplateEncapDelete,

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
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Tunnel encapsulation using IP in IP",
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
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Tunnel encapsulation using GRE",
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
func resourceDdosZoneTemplateEncapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateEncapCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateEncap(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateEncapRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateEncapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateEncapUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateEncap(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateEncapRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateEncapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateEncapDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateEncap(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateEncapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateEncapRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateEncap(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosZoneTemplateEncapTunnelEncap(d []interface{}) edpt.DdosZoneTemplateEncapTunnelEncap {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateEncapTunnelEncap
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpCfg = getObjectDdosZoneTemplateEncapTunnelEncapIpCfg(in["ip_cfg"].([]interface{}))
		ret.GreCfg = getObjectDdosZoneTemplateEncapTunnelEncapGreCfg(in["gre_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateEncapTunnelEncapIpCfg(d []interface{}) edpt.DdosZoneTemplateEncapTunnelEncapIpCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateEncapTunnelEncapIpCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpEncap = in["ip_encap"].(int)
		ret.Always = getObjectDdosZoneTemplateEncapTunnelEncapIpCfgAlways(in["always"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateEncapTunnelEncapIpCfgAlways(d []interface{}) edpt.DdosZoneTemplateEncapTunnelEncapIpCfgAlways {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateEncapTunnelEncapIpCfgAlways
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateEncapTunnelEncapGreCfg(d []interface{}) edpt.DdosZoneTemplateEncapTunnelEncapGreCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateEncapTunnelEncapGreCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GreEncap = in["gre_encap"].(int)
		ret.GreAlways = getObjectDdosZoneTemplateEncapTunnelEncapGreCfgGreAlways(in["gre_always"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateEncapTunnelEncapGreCfgGreAlways(d []interface{}) edpt.DdosZoneTemplateEncapTunnelEncapGreCfgGreAlways {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateEncapTunnelEncapGreCfgGreAlways
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GreIpv4 = in["gre_ipv4"].(string)
		ret.KeyIpv4 = in["key_ipv4"].(string)
		ret.GreIpv6 = in["gre_ipv6"].(string)
		ret.KeyIpv6 = in["key_ipv6"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneTemplateEncap(d *schema.ResourceData) edpt.DdosZoneTemplateEncap {
	var ret edpt.DdosZoneTemplateEncap
	ret.Inst.EncapTmplName = d.Get("encap_tmpl_name").(string)
	ret.Inst.PreserveSourceIp = d.Get("preserve_source_ip").(int)
	ret.Inst.TunnelEncap = getObjectDdosZoneTemplateEncapTunnelEncap(d.Get("tunnel_encap").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
