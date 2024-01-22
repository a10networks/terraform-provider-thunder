package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpAddressFamilyIpv6NetworkIpv6Network() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_address_family_ipv6_network_ipv6_network`: Specify a ip address mask network to announce via BGP\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkCreate,
		UpdateContext: resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkUpdate,
		ReadContext:   resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkRead,
		DeleteContext: resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkDelete,

		Schema: map[string]*schema.Schema{
			"backdoor": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify a BGP backdoor route",
			},
			"comm_value": {
				Type: schema.TypeString, Optional: true, Description: "community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export",
			},
			"description": {
				Type: schema.TypeString, Optional: true, Description: "Network specific description (Up to 80 characters describing this network)",
			},
			"lcomm_value": {
				Type: schema.TypeString, Optional: true, Description: "Large community value in the format XX:YY:ZZ",
			},
			"network_ipv6": {
				Type: schema.TypeString, Required: true, Description: "Specify a network to announce via BGP",
			},
			"route_map": {
				Type: schema.TypeString, Optional: true, Description: "Route-map to modify the attributes (Name of the route map)",
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
func resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NetworkIpv6Network(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NetworkIpv6Network(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NetworkIpv6Network(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NetworkIpv6Network(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRouterBgpAddressFamilyIpv6NetworkIpv6Network(d *schema.ResourceData) edpt.RouterBgpAddressFamilyIpv6NetworkIpv6Network {
	var ret edpt.RouterBgpAddressFamilyIpv6NetworkIpv6Network
	ret.Inst.Backdoor = d.Get("backdoor").(int)
	ret.Inst.CommValue = d.Get("comm_value").(string)
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.LcommValue = d.Get("lcomm_value").(string)
	ret.Inst.NetworkIpv6 = d.Get("network_ipv6").(string)
	ret.Inst.RouteMap = d.Get("route_map").(string)
	//omit uuid
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
