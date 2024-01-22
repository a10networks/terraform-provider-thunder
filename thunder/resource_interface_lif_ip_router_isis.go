package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLifIpRouterIsis() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_lif_ip_router_isis`: IS-IS Routing for IP\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLifIpRouterIsisCreate,
		UpdateContext: resourceInterfaceLifIpRouterIsisUpdate,
		ReadContext:   resourceInterfaceLifIpRouterIsisRead,
		DeleteContext: resourceInterfaceLifIpRouterIsisDelete,

		Schema: map[string]*schema.Schema{
			"tag": {
				Type: schema.TypeString, Optional: true, Description: "ISO routing area tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifname": {
				Type: schema.TypeString, Required: true, Description: "Ifname",
			},
		},
	}
}
func resourceInterfaceLifIpRouterIsisCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpRouterIsisCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpRouterIsis(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIpRouterIsisRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLifIpRouterIsisUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpRouterIsisUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpRouterIsis(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIpRouterIsisRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLifIpRouterIsisDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpRouterIsisDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpRouterIsis(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLifIpRouterIsisRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpRouterIsisRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpRouterIsis(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceLifIpRouterIsis(d *schema.ResourceData) edpt.InterfaceLifIpRouterIsis {
	var ret edpt.InterfaceLifIpRouterIsis
	ret.Inst.Tag = d.Get("tag").(string)
	//omit uuid
	ret.Inst.Ifname = d.Get("ifname").(string)
	return ret
}
