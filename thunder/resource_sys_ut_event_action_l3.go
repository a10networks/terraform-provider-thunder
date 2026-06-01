package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtEventActionL3() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_event_action_l3`: L3 packet paramters\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtEventActionL3Create,
		UpdateContext: resourceSysUtEventActionL3Update,
		ReadContext:   resourceSysUtEventActionL3Read,
		DeleteContext: resourceSysUtEventActionL3Delete,

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
						"ipv4_address": {
							Type: schema.TypeString, Optional: true, Description: "IP address",
						},
						"ipv6_address": {
							Type: schema.TypeString, Optional: true, Description: "Ipv6 address",
						},
						"nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "Nat pool",
						},
						"virtual_server": {
							Type: schema.TypeString, Optional: true, Description: "vip",
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
			"event_number": {
				Type: schema.TypeString, Required: true, Description: "EventNumber",
			},
			"direction": {
				Type: schema.TypeString, Required: true, Description: "Direction",
			},
		},
	}
}
func resourceSysUtEventActionL3Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionL3Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionL3(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtEventActionL3Read(ctx, d, meta)
	}
	return diags
}

func resourceSysUtEventActionL3Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionL3Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionL3(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtEventActionL3Read(ctx, d, meta)
	}
	return diags
}
func resourceSysUtEventActionL3Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionL3Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionL3(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtEventActionL3Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionL3Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionL3(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSysUtEventActionL3IpList(d []interface{}) []edpt.SysUtEventActionL3IpList {

	count1 := len(d)
	ret := make([]edpt.SysUtEventActionL3IpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtEventActionL3IpList
		oi.SrcDst = in["src_dst"].(string)
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.NatPool = in["nat_pool"].(string)
		oi.VirtualServer = in["virtual_server"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSysUtEventActionL3(d *schema.ResourceData) edpt.SysUtEventActionL3 {
	var ret edpt.SysUtEventActionL3
	ret.Inst.Checksum = d.Get("checksum").(string)
	ret.Inst.IpList = getSliceSysUtEventActionL3IpList(d.Get("ip_list").([]interface{}))
	ret.Inst.Protocol = d.Get("protocol").(int)
	ret.Inst.Ttl = d.Get("ttl").(int)
	ret.Inst.Type = d.Get("type").(string)
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	ret.Inst.EventNumber = d.Get("event_number").(string)
	ret.Inst.Direction = d.Get("direction").(string)
	return ret
}
