package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtEventActionL3Ip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_event_action_l3_ip`: IP address\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtEventActionL3IpCreate,
		UpdateContext: resourceSysUtEventActionL3IpUpdate,
		ReadContext:   resourceSysUtEventActionL3IpRead,
		DeleteContext: resourceSysUtEventActionL3IpDelete,

		Schema: map[string]*schema.Schema{
			"ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
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
			"src_dst": {
				Type: schema.TypeString, Required: true, Description: "'dest': dest; 'src': src;",
			},
			"trunk": {
				Type: schema.TypeInt, Optional: true, Description: "Trunk number",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ve": {
				Type: schema.TypeInt, Optional: true, Description: "Virtual Ethernet interface",
			},
			"virtual_server": {
				Type: schema.TypeString, Optional: true, Description: "vip",
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
func resourceSysUtEventActionL3IpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionL3IpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionL3Ip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtEventActionL3IpRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtEventActionL3IpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionL3IpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionL3Ip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtEventActionL3IpRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtEventActionL3IpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionL3IpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionL3Ip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtEventActionL3IpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionL3IpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionL3Ip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSysUtEventActionL3Ip(d *schema.ResourceData) edpt.SysUtEventActionL3Ip {
	var ret edpt.SysUtEventActionL3Ip
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.Ipv4Address = d.Get("ipv4_address").(string)
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	ret.Inst.NatPool = d.Get("nat_pool").(string)
	ret.Inst.SrcDst = d.Get("src_dst").(string)
	ret.Inst.Trunk = d.Get("trunk").(int)
	//omit uuid
	ret.Inst.Ve = d.Get("ve").(int)
	ret.Inst.VirtualServer = d.Get("virtual_server").(string)
	ret.Inst.EventNumber = d.Get("event_number").(string)
	ret.Inst.Direction = d.Get("direction").(string)
	return ret
}
