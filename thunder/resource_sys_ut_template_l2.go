package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtTemplateL2() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_template_l2`: L2 packet paramters\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtTemplateL2Create,
		UpdateContext: resourceSysUtTemplateL2Update,
		ReadContext:   resourceSysUtTemplateL2Read,
		DeleteContext: resourceSysUtTemplateL2Delete,

		Schema: map[string]*schema.Schema{
			"ethertype": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "L2 frame type",
			},
			"mac_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_dst": {
							Type: schema.TypeString, Required: true, Description: "'dest': dest; 'src': src;",
						},
						"address_type": {
							Type: schema.TypeString, Optional: true, Description: "'broadcast': broadcast; 'multicast': multicast;",
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
						"value": {
							Type: schema.TypeString, Optional: true, Description: "Mac Address",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"protocol": {
				Type: schema.TypeString, Optional: true, Default: "ipv4", Description: "'arp': arp; 'ipv4': ipv4; 'ipv6': ipv6;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeInt, Optional: true, Description: "ethertype number",
			},
			"vlan": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Vlan ID on the packet. 0 is untagged",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSysUtTemplateL2Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL2Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL2(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateL2Read(ctx, d, meta)
	}
	return diags
}

func resourceSysUtTemplateL2Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL2Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL2(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateL2Read(ctx, d, meta)
	}
	return diags
}
func resourceSysUtTemplateL2Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL2Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL2(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtTemplateL2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL2Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL2(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSysUtTemplateL2MacList(d []interface{}) []edpt.SysUtTemplateL2MacList {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateL2MacList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateL2MacList
		oi.SrcDst = in["src_dst"].(string)
		oi.AddressType = in["address_type"].(string)
		oi.VirtualServer = in["virtual_server"].(string)
		oi.NatPool = in["nat_pool"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Value = in["value"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSysUtTemplateL2(d *schema.ResourceData) edpt.SysUtTemplateL2 {
	var ret edpt.SysUtTemplateL2
	ret.Inst.Ethertype = d.Get("ethertype").(int)
	ret.Inst.MacList = getSliceSysUtTemplateL2MacList(d.Get("mac_list").([]interface{}))
	ret.Inst.Protocol = d.Get("protocol").(string)
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	ret.Inst.Vlan = d.Get("vlan").(int)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
