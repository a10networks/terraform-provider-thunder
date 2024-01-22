package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpNetworkIpCidr() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_network_ip_cidr`: Specify a ip network to announce via BGP\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpNetworkIpCidrCreate,
		UpdateContext: resourceRouterBgpNetworkIpCidrUpdate,
		ReadContext:   resourceRouterBgpNetworkIpCidrRead,
		DeleteContext: resourceRouterBgpNetworkIpCidrDelete,

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
			"network_ipv4_cidr": {
				Type: schema.TypeString, Required: true, Description: "Specify network mask",
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
func resourceRouterBgpNetworkIpCidrCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNetworkIpCidrCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNetworkIpCidr(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpNetworkIpCidrRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpNetworkIpCidrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNetworkIpCidrUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNetworkIpCidr(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpNetworkIpCidrRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpNetworkIpCidrDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNetworkIpCidrDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNetworkIpCidr(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpNetworkIpCidrRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNetworkIpCidrRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNetworkIpCidr(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRouterBgpNetworkIpCidr(d *schema.ResourceData) edpt.RouterBgpNetworkIpCidr {
	var ret edpt.RouterBgpNetworkIpCidr
	ret.Inst.Backdoor = d.Get("backdoor").(int)
	ret.Inst.CommValue = d.Get("comm_value").(string)
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.LcommValue = d.Get("lcomm_value").(string)
	ret.Inst.NetworkIpv4Cidr = d.Get("network_ipv4_cidr").(string)
	ret.Inst.RouteMap = d.Get("route_map").(string)
	//omit uuid
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
