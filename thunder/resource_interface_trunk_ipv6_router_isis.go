package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunkIpv6RouterIsis() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_trunk_ipv6_router_isis`: IS-IS Routing for IP\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTrunkIpv6RouterIsisCreate,
		UpdateContext: resourceInterfaceTrunkIpv6RouterIsisUpdate,
		ReadContext:   resourceInterfaceTrunkIpv6RouterIsisRead,
		DeleteContext: resourceInterfaceTrunkIpv6RouterIsisDelete,

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
func resourceInterfaceTrunkIpv6RouterIsisCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6RouterIsisCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6RouterIsis(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpv6RouterIsisRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTrunkIpv6RouterIsisUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6RouterIsisUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6RouterIsis(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpv6RouterIsisRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTrunkIpv6RouterIsisDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6RouterIsisDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6RouterIsis(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTrunkIpv6RouterIsisRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6RouterIsisRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6RouterIsis(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceTrunkIpv6RouterIsis(d *schema.ResourceData) edpt.InterfaceTrunkIpv6RouterIsis {
	var ret edpt.InterfaceTrunkIpv6RouterIsis
	ret.Inst.Tag = d.Get("tag").(string)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
