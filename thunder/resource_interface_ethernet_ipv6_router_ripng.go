package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernetIpv6RouterRipng() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ethernet_ipv6_router_ripng`: ripng\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceEthernetIpv6RouterRipngCreate,
		UpdateContext: resourceInterfaceEthernetIpv6RouterRipngUpdate,
		ReadContext:   resourceInterfaceEthernetIpv6RouterRipngRead,
		DeleteContext: resourceInterfaceEthernetIpv6RouterRipngDelete,

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
func resourceInterfaceEthernetIpv6RouterRipngCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpv6RouterRipngCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIpv6RouterRipng(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetIpv6RouterRipngRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceEthernetIpv6RouterRipngUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpv6RouterRipngUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIpv6RouterRipng(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetIpv6RouterRipngRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceEthernetIpv6RouterRipngDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpv6RouterRipngDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIpv6RouterRipng(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceEthernetIpv6RouterRipngRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpv6RouterRipngRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIpv6RouterRipng(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceEthernetIpv6RouterRipng(d *schema.ResourceData) edpt.InterfaceEthernetIpv6RouterRipng {
	var ret edpt.InterfaceEthernetIpv6RouterRipng
	ret.Inst.Rip = d.Get("rip").(int)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
