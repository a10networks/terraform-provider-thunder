package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLoopbackIpv6RouterRipng() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_loopback_ipv6_router_ripng`: ripng\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLoopbackIpv6RouterRipngCreate,
		UpdateContext: resourceInterfaceLoopbackIpv6RouterRipngUpdate,
		ReadContext:   resourceInterfaceLoopbackIpv6RouterRipngRead,
		DeleteContext: resourceInterfaceLoopbackIpv6RouterRipngDelete,

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
func resourceInterfaceLoopbackIpv6RouterRipngCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpv6RouterRipngCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpv6RouterRipng(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLoopbackIpv6RouterRipngRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLoopbackIpv6RouterRipngUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpv6RouterRipngUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpv6RouterRipng(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLoopbackIpv6RouterRipngRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLoopbackIpv6RouterRipngDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpv6RouterRipngDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpv6RouterRipng(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLoopbackIpv6RouterRipngRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpv6RouterRipngRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpv6RouterRipng(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceLoopbackIpv6RouterRipng(d *schema.ResourceData) edpt.InterfaceLoopbackIpv6RouterRipng {
	var ret edpt.InterfaceLoopbackIpv6RouterRipng
	ret.Inst.Rip = d.Get("rip").(int)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
