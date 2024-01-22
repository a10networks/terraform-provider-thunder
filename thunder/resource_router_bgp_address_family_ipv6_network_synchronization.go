package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpAddressFamilyIpv6NetworkSynchronization() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_address_family_ipv6_network_synchronization`: help Perform IGP synchronization\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationCreate,
		UpdateContext: resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationUpdate,
		ReadContext:   resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead,
		DeleteContext: resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationDelete,

		Schema: map[string]*schema.Schema{
			"network_synchronization": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Perform IGP synchronization",
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
func resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NetworkSynchronization(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NetworkSynchronization(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NetworkSynchronization(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NetworkSynchronization(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRouterBgpAddressFamilyIpv6NetworkSynchronization(d *schema.ResourceData) edpt.RouterBgpAddressFamilyIpv6NetworkSynchronization {
	var ret edpt.RouterBgpAddressFamilyIpv6NetworkSynchronization
	ret.Inst.NetworkSynchronization = d.Get("network_synchronization").(int)
	//omit uuid
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
