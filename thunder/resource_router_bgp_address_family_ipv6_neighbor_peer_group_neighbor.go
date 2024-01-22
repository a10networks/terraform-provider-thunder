package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_address_family_ipv6_neighbor_peer_group_neighbor`: Specify a peer-group neighbor router\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborCreate,
		UpdateContext: resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborUpdate,
		ReadContext:   resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborRead,
		DeleteContext: resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborDelete,

		Schema: map[string]*schema.Schema{
			"activate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the Address Family for this Neighbor",
			},
			"allowas_in": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept as-path with my AS present in it",
			},
			"allowas_in_count": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Number of occurrences of AS number",
			},
			"inbound": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow inbound soft reconfiguration for this neighbor",
			},
			"maximum_prefix": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))",
			},
			"maximum_prefix_thres": {
				Type: schema.TypeInt, Optional: true, Description: "threshold-value, 1 to 100 percent",
			},
			"neighbor_route_map_lists": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nbr_route_map": {
							Type: schema.TypeString, Optional: true, Description: "Apply route map to neighbor (Name of route map)",
						},
						"nbr_rmap_direction": {
							Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
						},
					},
				},
			},
			"next_hop_self": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable the next hop calculation for this neighbor",
			},
			"peer_group": {
				Type: schema.TypeString, Required: true, Description: "Neighbor tag",
			},
			"remove_private_as": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove private AS number from outbound updates",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"weight": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set default weight for routes from this neighbor",
			},
			"as_number": {
				Type: schema.TypeString, Required: true, Description: "AsNumber",
			},
		},
	}
}
func resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborNeighborRouteMapLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborNeighborRouteMapLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor(d *schema.ResourceData) edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor {
	var ret edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor
	ret.Inst.Activate = d.Get("activate").(int)
	ret.Inst.AllowasIn = d.Get("allowas_in").(int)
	ret.Inst.AllowasInCount = d.Get("allowas_in_count").(int)
	ret.Inst.Inbound = d.Get("inbound").(int)
	ret.Inst.MaximumPrefix = d.Get("maximum_prefix").(int)
	ret.Inst.MaximumPrefixThres = d.Get("maximum_prefix_thres").(int)
	ret.Inst.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborNeighborRouteMapLists(d.Get("neighbor_route_map_lists").([]interface{}))
	ret.Inst.NextHopSelf = d.Get("next_hop_self").(int)
	ret.Inst.PeerGroup = d.Get("peer_group").(string)
	ret.Inst.RemovePrivateAs = d.Get("remove_private_as").(int)
	//omit uuid
	ret.Inst.Weight = d.Get("weight").(int)
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
