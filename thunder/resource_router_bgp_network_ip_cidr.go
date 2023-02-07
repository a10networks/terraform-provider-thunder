package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBgpNetworkIpCidr() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_network_ip_cidr`: Specify a ip network to announce via BGP\n\n",
		CreateContext: resourceRouterBgpNetworkIpCidrCreate,
		UpdateContext: resourceRouterBgpNetworkIpCidrUpdate,
		ReadContext:   resourceRouterBgpNetworkIpCidrRead,
		DeleteContext: resourceRouterBgpNetworkIpCidrDelete,
		Schema: map[string]*schema.Schema{
			"as_number": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "BGP AS number",
				ValidateFunc: validation.StringLenBetween(1, 11),
			},
			"backdoor": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify a BGP backdoor route",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"comm_value": {
				Type: schema.TypeString, Optional: true, Description: "community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export",
				ValidateFunc: validation.StringLenBetween(1, 128),
			},
			"description": {
				Type: schema.TypeString, Optional: true, Description: "Network specific description (Up to 80 characters describing this network)",
				ValidateFunc: validation.StringLenBetween(1, 80),
			},
			"network_ipv4_cidr": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Specify network mask",
				ValidateFunc: validation.IsCIDR,
			},
			"route_map": {
				Type: schema.TypeString, Optional: true, Description: "Route-map to modify the attributes (Name of the route map)",
				ValidateFunc: validation.StringLenBetween(1, 128),
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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

func resourceRouterBgpNetworkIpCidrRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNetworkIpCidrRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := edpt.RouterBgpNetworkIpCidr{}
		obj.Inst.AsNumber = d.Get("as_number").(string)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
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
		obj := edpt.RouterBgpNetworkIpCidr{}
		obj.Inst.AsNumber = d.Get("as_number").(string)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRouterBgpNetworkIpCidr(d *schema.ResourceData) edpt.RouterBgpNetworkIpCidr {
	var ret edpt.RouterBgpNetworkIpCidr
	ret.Inst.AsNumber = d.Get("as_number").(string)
	ret.Inst.Backdoor = d.Get("backdoor").(int)
	ret.Inst.CommValue = d.Get("comm_value").(string)
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.NetworkIpv4Cidr = d.Get("network_ipv4_cidr").(string)
	ret.Inst.RouteMap = d.Get("route_map").(string)
	//omit uuid
	return ret
}
