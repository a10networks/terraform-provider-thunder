package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpAddressFamilyIpv6NeighborIpv6Neighbor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_address_family_ipv6_neighbor_ipv6_neighbor`: Specify a peer-group neighbor router\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborCreate,
		UpdateContext: resourceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborUpdate,
		ReadContext:   resourceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborRead,
		DeleteContext: resourceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborDelete,

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
			"default_originate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Originate default route to this neighbor",
			},
			"distribute_lists": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"distribute_list": {
							Type: schema.TypeString, Optional: true, Description: "Filter updates to/from this neighbor (IP standard/extended/named access list)",
						},
						"distribute_list_direction": {
							Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
						},
					},
				},
			},
			"graceful_restart": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable graceful-restart helper for this neighbor",
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
			"neighbor_filter_lists": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_list": {
							Type: schema.TypeString, Optional: true, Description: "Establish BGP filters (AS path access-list name)",
						},
						"filter_list_direction": {
							Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
						},
					},
				},
			},
			"neighbor_ipv6": {
				Type: schema.TypeString, Required: true, Description: "Neighbor IPv6 address",
			},
			"neighbor_prefix_lists": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nbr_prefix_list": {
							Type: schema.TypeString, Optional: true, Description: "Filter updates to/from this neighbor (Name of a prefix list)",
						},
						"nbr_prefix_list_direction": {
							Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
						},
					},
				},
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
			"peer_group_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure peer-group (peer-group name)",
			},
			"prefix_list_direction": {
				Type: schema.TypeString, Optional: true, Description: "'both': both; 'receive': receive; 'send': send;",
			},
			"remove_private_as": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove private AS number from outbound updates",
			},
			"restart_min": {
				Type: schema.TypeInt, Optional: true, Description: "restart value, 1 to 1440 minutes",
			},
			"route_map": {
				Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria to originate default (route-map name)",
			},
			"send_community_val": {
				Type: schema.TypeString, Optional: true, Default: "both", Description: "'all': Send Standard, Extended, and Large Community attributes; 'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes; 'large': Send Large Community attributes;",
			},
			"unsuppress_map": {
				Type: schema.TypeString, Optional: true, Description: "Route-map to selectively unsuppress suppressed routes (Name of route map)",
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
func resourceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NeighborIpv6Neighbor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NeighborIpv6Neighbor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NeighborIpv6Neighbor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NeighborIpv6Neighbor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborDistributeLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborDistributeLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborDistributeLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborDistributeLists
		oi.DistributeList = in["distribute_list"].(string)
		oi.DistributeListDirection = in["distribute_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborFilterLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborFilterLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborFilterLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborFilterLists
		oi.FilterList = in["filter_list"].(string)
		oi.FilterListDirection = in["filter_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborPrefixLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborPrefixLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborPrefixLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborPrefixLists
		oi.NbrPrefixList = in["nbr_prefix_list"].(string)
		oi.NbrPrefixListDirection = in["nbr_prefix_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborRouteMapLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborRouteMapLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterBgpAddressFamilyIpv6NeighborIpv6Neighbor(d *schema.ResourceData) edpt.RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor {
	var ret edpt.RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor
	ret.Inst.Activate = d.Get("activate").(int)
	ret.Inst.AllowasIn = d.Get("allowas_in").(int)
	ret.Inst.AllowasInCount = d.Get("allowas_in_count").(int)
	ret.Inst.DefaultOriginate = d.Get("default_originate").(int)
	ret.Inst.DistributeLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborDistributeLists(d.Get("distribute_lists").([]interface{}))
	ret.Inst.GracefulRestart = d.Get("graceful_restart").(int)
	ret.Inst.Inbound = d.Get("inbound").(int)
	ret.Inst.MaximumPrefix = d.Get("maximum_prefix").(int)
	ret.Inst.MaximumPrefixThres = d.Get("maximum_prefix_thres").(int)
	ret.Inst.NeighborFilterLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborFilterLists(d.Get("neighbor_filter_lists").([]interface{}))
	ret.Inst.NeighborIpv6 = d.Get("neighbor_ipv6").(string)
	ret.Inst.NeighborPrefixLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborPrefixLists(d.Get("neighbor_prefix_lists").([]interface{}))
	ret.Inst.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborRouteMapLists(d.Get("neighbor_route_map_lists").([]interface{}))
	ret.Inst.NextHopSelf = d.Get("next_hop_self").(int)
	ret.Inst.PeerGroupName = d.Get("peer_group_name").(string)
	ret.Inst.PrefixListDirection = d.Get("prefix_list_direction").(string)
	ret.Inst.RemovePrivateAs = d.Get("remove_private_as").(int)
	ret.Inst.RestartMin = d.Get("restart_min").(int)
	ret.Inst.RouteMap = d.Get("route_map").(string)
	ret.Inst.SendCommunityVal = d.Get("send_community_val").(string)
	ret.Inst.UnsuppressMap = d.Get("unsuppress_map").(string)
	//omit uuid
	ret.Inst.Weight = d.Get("weight").(int)
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
