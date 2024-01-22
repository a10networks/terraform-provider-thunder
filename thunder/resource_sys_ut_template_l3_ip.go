package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtTemplateL3Ip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_template_l3_ip`: IP address\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtTemplateL3IpCreate,
		UpdateContext: resourceSysUtTemplateL3IpUpdate,
		ReadContext:   resourceSysUtTemplateL3IpRead,
		DeleteContext: resourceSysUtTemplateL3IpDelete,

		Schema: map[string]*schema.Schema{
			"ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
			},
			"ipv4_end_address": {
				Type: schema.TypeString, Optional: true, Description: "IP end address",
			},
			"ipv4_start_address": {
				Type: schema.TypeString, Optional: true, Description: "IP address",
			},
			"ipv6_end_address": {
				Type: schema.TypeString, Optional: true, Description: "Ipv6 end address",
			},
			"ipv6_start_address": {
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSysUtTemplateL3IpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL3IpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL3Ip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateL3IpRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtTemplateL3IpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL3IpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL3Ip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateL3IpRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtTemplateL3IpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL3IpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL3Ip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtTemplateL3IpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL3IpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL3Ip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSysUtTemplateL3Ip(d *schema.ResourceData) edpt.SysUtTemplateL3Ip {
	var ret edpt.SysUtTemplateL3Ip
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.Ipv4EndAddress = d.Get("ipv4_end_address").(string)
	ret.Inst.Ipv4StartAddress = d.Get("ipv4_start_address").(string)
	ret.Inst.Ipv6EndAddress = d.Get("ipv6_end_address").(string)
	ret.Inst.Ipv6StartAddress = d.Get("ipv6_start_address").(string)
	ret.Inst.NatPool = d.Get("nat_pool").(string)
	ret.Inst.SrcDst = d.Get("src_dst").(string)
	ret.Inst.Trunk = d.Get("trunk").(int)
	//omit uuid
	ret.Inst.Ve = d.Get("ve").(int)
	ret.Inst.VirtualServer = d.Get("virtual_server").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
