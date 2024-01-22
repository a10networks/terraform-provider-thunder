package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpAddressFamilyIpv6NetworkMonitorDefault() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_address_family_ipv6_network_monitor_default`: help default\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpAddressFamilyIpv6NetworkMonitorDefaultCreate,
		UpdateContext: resourceRouterBgpAddressFamilyIpv6NetworkMonitorDefaultUpdate,
		ReadContext:   resourceRouterBgpAddressFamilyIpv6NetworkMonitorDefaultRead,
		DeleteContext: resourceRouterBgpAddressFamilyIpv6NetworkMonitorDefaultDelete,

		Schema: map[string]*schema.Schema{
			"network_monitor_default": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "default route monitoring",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"as_number": {
				Type: schema.TypeString, Required: true, Description: "AsNumber",
			},
		},
	}
}
func resourceRouterBgpAddressFamilyIpv6NetworkMonitorDefaultCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NetworkMonitorDefaultCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NetworkMonitorDefault(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6NetworkMonitorDefaultRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NetworkMonitorDefaultUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NetworkMonitorDefaultUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NetworkMonitorDefault(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6NetworkMonitorDefaultRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpAddressFamilyIpv6NetworkMonitorDefaultDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NetworkMonitorDefaultDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NetworkMonitorDefault(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NetworkMonitorDefaultRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NetworkMonitorDefaultRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NetworkMonitorDefault(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRouterBgpAddressFamilyIpv6NetworkMonitorDefault(d *schema.ResourceData) edpt.RouterBgpAddressFamilyIpv6NetworkMonitorDefault {
	var ret edpt.RouterBgpAddressFamilyIpv6NetworkMonitorDefault
	ret.Inst.NetworkMonitorDefault = d.Get("network_monitor_default").(int)
	//omit uuid
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
