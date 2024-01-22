package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpAddressFamilyIpv6Flowspec() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_address_family_ipv6_flowspec`: ipv6-flowspec\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpAddressFamilyIpv6FlowspecCreate,
		UpdateContext: resourceRouterBgpAddressFamilyIpv6FlowspecUpdate,
		ReadContext:   resourceRouterBgpAddressFamilyIpv6FlowspecRead,
		DeleteContext: resourceRouterBgpAddressFamilyIpv6FlowspecDelete,

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
func resourceRouterBgpAddressFamilyIpv6FlowspecCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6FlowspecCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6Flowspec(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6FlowspecRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6FlowspecUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6FlowspecUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6Flowspec(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6FlowspecRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpAddressFamilyIpv6FlowspecDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6FlowspecDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6Flowspec(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6FlowspecRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6FlowspecRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6Flowspec(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRouterBgpAddressFamilyIpv6FlowspecNeighbor1152(d []interface{}) edpt.RouterBgpAddressFamilyIpv6FlowspecNeighbor1152 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6FlowspecNeighbor1152
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4NeighborList = getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList(in["ipv4_neighbor_list"].([]interface{}))
		ret.Ipv6NeighborList = getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList(in["ipv6_neighbor_list"].([]interface{}))
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList
		oi.NeighborIpv4 = in["neighbor_ipv4"].(string)
		oi.Activate = in["activate"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList
		oi.NeighborIpv6 = in["neighbor_ipv6"].(string)
		oi.Activate = in["activate"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterBgpAddressFamilyIpv6Flowspec(d *schema.ResourceData) edpt.RouterBgpAddressFamilyIpv6Flowspec {
	var ret edpt.RouterBgpAddressFamilyIpv6Flowspec
	ret.Inst.Neighbor = getObjectRouterBgpAddressFamilyIpv6FlowspecNeighbor1152(d.Get("neighbor").([]interface{}))
	//omit uuid
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
