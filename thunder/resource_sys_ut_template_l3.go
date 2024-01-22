package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtTemplateL3() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_template_l3`: L3 packet paramters\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtTemplateL3Create,
		UpdateContext: resourceSysUtTemplateL3Update,
		ReadContext:   resourceSysUtTemplateL3Read,
		DeleteContext: resourceSysUtTemplateL3Delete,

		Schema: map[string]*schema.Schema{
			"checksum": {
				Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
			},
			"ip_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_dst": {
							Type: schema.TypeString, Required: true, Description: "'dest': dest; 'src': src;",
						},
						"ipv4_start_address": {
							Type: schema.TypeString, Optional: true, Description: "IP address",
						},
						"ipv4_end_address": {
							Type: schema.TypeString, Optional: true, Description: "IP end address",
						},
						"ipv6_start_address": {
							Type: schema.TypeString, Optional: true, Description: "Ipv6 address",
						},
						"ipv6_end_address": {
							Type: schema.TypeString, Optional: true, Description: "Ipv6 end address",
						},
						"virtual_server": {
							Type: schema.TypeString, Optional: true, Description: "vip",
						},
						"nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "Nat pool",
						},
						"ethernet": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
						},
						"ve": {
							Type: schema.TypeInt, Optional: true, Description: "Virtual Ethernet interface",
						},
						"trunk": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk number",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"protocol": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "L4 Protocol",
			},
			"ttl": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"type": {
				Type: schema.TypeString, Optional: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeInt, Optional: true, Description: "protocol number",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSysUtTemplateL3Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL3Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL3(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateL3Read(ctx, d, meta)
	}
	return diags
}

func resourceSysUtTemplateL3Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL3Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL3(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateL3Read(ctx, d, meta)
	}
	return diags
}
func resourceSysUtTemplateL3Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL3Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL3(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtTemplateL3Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL3Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL3(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSysUtTemplateL3IpList(d []interface{}) []edpt.SysUtTemplateL3IpList {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateL3IpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateL3IpList
		oi.SrcDst = in["src_dst"].(string)
		oi.Ipv4StartAddress = in["ipv4_start_address"].(string)
		oi.Ipv4EndAddress = in["ipv4_end_address"].(string)
		oi.Ipv6StartAddress = in["ipv6_start_address"].(string)
		oi.Ipv6EndAddress = in["ipv6_end_address"].(string)
		oi.VirtualServer = in["virtual_server"].(string)
		oi.NatPool = in["nat_pool"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSysUtTemplateL3(d *schema.ResourceData) edpt.SysUtTemplateL3 {
	var ret edpt.SysUtTemplateL3
	ret.Inst.Checksum = d.Get("checksum").(string)
	ret.Inst.IpList = getSliceSysUtTemplateL3IpList(d.Get("ip_list").([]interface{}))
	ret.Inst.Protocol = d.Get("protocol").(int)
	ret.Inst.Ttl = d.Get("ttl").(int)
	ret.Inst.Type = d.Get("type").(string)
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
