package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtTemplateL2Mac() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_template_l2_mac`: Mac Address\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtTemplateL2MacCreate,
		UpdateContext: resourceSysUtTemplateL2MacUpdate,
		ReadContext:   resourceSysUtTemplateL2MacRead,
		DeleteContext: resourceSysUtTemplateL2MacDelete,

		Schema: map[string]*schema.Schema{
			"address_type": {
				Type: schema.TypeString, Optional: true, Description: "'broadcast': broadcast; 'multicast': multicast;",
			},
			"ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
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
			"value": {
				Type: schema.TypeString, Optional: true, Description: "Mac Address",
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
func resourceSysUtTemplateL2MacCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL2MacCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL2Mac(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateL2MacRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtTemplateL2MacUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL2MacUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL2Mac(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateL2MacRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtTemplateL2MacDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL2MacDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL2Mac(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtTemplateL2MacRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL2MacRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL2Mac(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSysUtTemplateL2Mac(d *schema.ResourceData) edpt.SysUtTemplateL2Mac {
	var ret edpt.SysUtTemplateL2Mac
	ret.Inst.AddressType = d.Get("address_type").(string)
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.NatPool = d.Get("nat_pool").(string)
	ret.Inst.SrcDst = d.Get("src_dst").(string)
	ret.Inst.Trunk = d.Get("trunk").(int)
	//omit uuid
	ret.Inst.Value = d.Get("value").(string)
	ret.Inst.Ve = d.Get("ve").(int)
	ret.Inst.VirtualServer = d.Get("virtual_server").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
