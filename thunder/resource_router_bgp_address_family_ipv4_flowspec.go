package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpAddressFamilyIpv4Flowspec() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_address_family_ipv4_flowspec`: ipv4-flowspec\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpAddressFamilyIpv4FlowspecCreate,
		UpdateContext: resourceRouterBgpAddressFamilyIpv4FlowspecUpdate,
		ReadContext:   resourceRouterBgpAddressFamilyIpv4FlowspecRead,
		DeleteContext: resourceRouterBgpAddressFamilyIpv4FlowspecDelete,

		Schema: map[string]*schema.Schema{
			"neighbor": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_neighbor_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"neighbor_ipv4": {
										Type: schema.TypeString, Required: true, Description: "Neighbor address",
									},
									"activate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the Address Family for this Neighbor",
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
									"send_community_val": {
										Type: schema.TypeString, Optional: true, Default: "both", Description: "'all': Send Standard, Extended, and Large Community attributes; 'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes; 'large': Send Large Community attributes;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ipv6_neighbor_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"neighbor_ipv6": {
										Type: schema.TypeString, Required: true, Description: "Neighbor IPv6 address",
									},
									"activate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the Address Family for this Neighbor",
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
									"send_community_val": {
										Type: schema.TypeString, Optional: true, Default: "both", Description: "'all': Send Standard, Extended, and Large Community attributes; 'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes; 'large': Send Large Community attributes;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
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
func resourceRouterBgpAddressFamilyIpv4FlowspecCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv4FlowspecCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv4Flowspec(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv4FlowspecRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv4FlowspecUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv4FlowspecUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv4Flowspec(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv4FlowspecRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpAddressFamilyIpv4FlowspecDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv4FlowspecDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv4Flowspec(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv4FlowspecRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv4FlowspecRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv4Flowspec(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRouterBgpAddressFamilyIpv4FlowspecNeighbor1151(d []interface{}) edpt.RouterBgpAddressFamilyIpv4FlowspecNeighbor1151 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv4FlowspecNeighbor1151
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4NeighborList = getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList(in["ipv4_neighbor_list"].([]interface{}))
		ret.Ipv6NeighborList = getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList(in["ipv6_neighbor_list"].([]interface{}))
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList(d []interface{}) []edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList
		oi.NeighborIpv4 = in["neighbor_ipv4"].(string)
		oi.Activate = in["activate"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList(d []interface{}) []edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList
		oi.NeighborIpv6 = in["neighbor_ipv6"].(string)
		oi.Activate = in["activate"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterBgpAddressFamilyIpv4Flowspec(d *schema.ResourceData) edpt.RouterBgpAddressFamilyIpv4Flowspec {
	var ret edpt.RouterBgpAddressFamilyIpv4Flowspec
	ret.Inst.Neighbor = getObjectRouterBgpAddressFamilyIpv4FlowspecNeighbor1151(d.Get("neighbor").([]interface{}))
	//omit uuid
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
