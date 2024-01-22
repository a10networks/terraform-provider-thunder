package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunkIpRouterIsis() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_trunk_ip_router_isis`: IS-IS Routing for IP\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTrunkIpRouterIsisCreate,
		UpdateContext: resourceInterfaceTrunkIpRouterIsisUpdate,
		ReadContext:   resourceInterfaceTrunkIpRouterIsisRead,
		DeleteContext: resourceInterfaceTrunkIpRouterIsisDelete,

		Schema: map[string]*schema.Schema{
			"tag": {
				Type: schema.TypeString, Optional: true, Description: "ISO routing area tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceTrunkIpRouterIsisCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpRouterIsisCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpRouterIsis(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpRouterIsisRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTrunkIpRouterIsisUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpRouterIsisUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpRouterIsis(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpRouterIsisRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTrunkIpRouterIsisDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpRouterIsisDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpRouterIsis(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTrunkIpRouterIsisRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpRouterIsisRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpRouterIsis(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceTrunkIpRouterIsis(d *schema.ResourceData) edpt.InterfaceTrunkIpRouterIsis {
	var ret edpt.InterfaceTrunkIpRouterIsis
	ret.Inst.Tag = d.Get("tag").(string)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
