package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLifIpv6RouterRipng() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_lif_ipv6_router_ripng`: ripng\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLifIpv6RouterRipngCreate,
		UpdateContext: resourceInterfaceLifIpv6RouterRipngUpdate,
		ReadContext:   resourceInterfaceLifIpv6RouterRipngRead,
		DeleteContext: resourceInterfaceLifIpv6RouterRipngDelete,

		Schema: map[string]*schema.Schema{
			"rip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "RIP Routing for IPv6",
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
func resourceInterfaceLifIpv6RouterRipngCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpv6RouterRipngCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpv6RouterRipng(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIpv6RouterRipngRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLifIpv6RouterRipngUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpv6RouterRipngUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpv6RouterRipng(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIpv6RouterRipngRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLifIpv6RouterRipngDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpv6RouterRipngDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpv6RouterRipng(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLifIpv6RouterRipngRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpv6RouterRipngRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpv6RouterRipng(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceLifIpv6RouterRipng(d *schema.ResourceData) edpt.InterfaceLifIpv6RouterRipng {
	var ret edpt.InterfaceLifIpv6RouterRipng
	ret.Inst.Rip = d.Get("rip").(int)
	//omit uuid
	ret.Inst.Ifname = d.Get("ifname").(string)
	return ret
}
