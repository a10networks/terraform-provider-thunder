package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_address_family_ipv6_neighbor_trunk_neighbor_ipv6`: Specify a trunk unnumbered neighbor\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Create,
		UpdateContext: resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Update,
		ReadContext:   resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Read,
		DeleteContext: resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Delete,

		Schema: map[string]*schema.Schema{
			"peer_group_name": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"trunk": {
				Type: schema.TypeInt, Required: true, Description: "Trunk interface number",
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
func resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6(d *schema.ResourceData) edpt.RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6 {
	var ret edpt.RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6
	ret.Inst.PeerGroupName = d.Get("peer_group_name").(string)
	ret.Inst.Trunk = d.Get("trunk").(int)
	//omit uuid
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
