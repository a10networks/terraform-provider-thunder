package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunkIpv6RouterRipng() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_trunk_ipv6_router_ripng`: ripng\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTrunkIpv6RouterRipngCreate,
		UpdateContext: resourceInterfaceTrunkIpv6RouterRipngUpdate,
		ReadContext:   resourceInterfaceTrunkIpv6RouterRipngRead,
		DeleteContext: resourceInterfaceTrunkIpv6RouterRipngDelete,

		Schema: map[string]*schema.Schema{
			"rip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "RIP Routing for IPv6",
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
func resourceInterfaceTrunkIpv6RouterRipngCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6RouterRipngCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6RouterRipng(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpv6RouterRipngRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTrunkIpv6RouterRipngUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6RouterRipngUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6RouterRipng(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpv6RouterRipngRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTrunkIpv6RouterRipngDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6RouterRipngDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6RouterRipng(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTrunkIpv6RouterRipngRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6RouterRipngRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6RouterRipng(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceTrunkIpv6RouterRipng(d *schema.ResourceData) edpt.InterfaceTrunkIpv6RouterRipng {
	var ret edpt.InterfaceTrunkIpv6RouterRipng
	ret.Inst.Rip = d.Get("rip").(int)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
