package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpNeighborPeerGroupNeighbor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_neighbor_peer_group_neighbor`: Specify a peer-group neighbor router\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpNeighborPeerGroupNeighborCreate,
		UpdateContext: resourceRouterBgpNeighborPeerGroupNeighborUpdate,
		ReadContext:   resourceRouterBgpNeighborPeerGroupNeighborRead,
		DeleteContext: resourceRouterBgpNeighborPeerGroupNeighborDelete,

		Schema: map[string]*schema.Schema{
			"activate": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable the Address Family for this Neighbor",
			},
			"advertisement_interval": {
				Type: schema.TypeInt, Optional: true, Description: "Minimum interval between sending BGP routing updates (time in seconds)",
			},
			"allowas_in": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept as-path with my AS present in it",
			},
			"allowas_in_count": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Number of occurrences of AS number",
			},
			"as_origination_interval": {
				Type: schema.TypeInt, Optional: true, Description: "Minimum interval between sending AS-origination routing updates (time in seconds)",
			},
			"bfd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
			},
			"collide_established": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include Neighbor in Established State for Collision Detection",
			},
			"connect": {
				Type: schema.TypeInt, Optional: true, Description: "BGP connect timer",
			},
			"default_originate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Originate default route to this neighbor",
			},
			"description": {
				Type: schema.TypeString, Optional: true, Description: "Neighbor specific description (Up to 80 characters describing this neighbor)",
			},
			"dont_capability_negotiate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not perform capability negotiation",
			},
			"dynamic": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertise dynamic capability to this neighbor",
			},
			"ebgp_multihop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow EBGP neighbors not on directly connected networks",
			},
			"ebgp_multihop_hop_count": {
				Type: schema.TypeInt, Optional: true, Description: "maximum hop count",
			},
			"enforce_multihop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enforce EBGP neighbors to perform multihop",
			},
			"ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
			},
			"extended_nexthop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertise extended-nexthop capability to this neighbor",
			},
			"inbound": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow inbound soft reconfiguration for this neighbor",
			},
			"lif": {
				Type: schema.TypeString, Optional: true, Description: "Logical interface (Lif interface name)",
			},
			"loopback": {
				Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
			},
			"maximum_prefix": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))",
			},
			"maximum_prefix_thres": {
				Type: schema.TypeInt, Optional: true, Description: "threshold-value, 1 to 100 percent",
			},
			"multihop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable multihop",
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
			"override_capability": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Override capability negotiation result",
			},
			"pass_value": {
				Type: schema.TypeString, Optional: true, Description: "Key String",
			},
			"passive": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't send open messages to this neighbor",
			},
			"peer_group": {
				Type: schema.TypeString, Required: true, Description: "Neighbor tag",
			},
			"peer_group_key": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure peer-group",
			},
			"peer_group_remote_as": {
				Type: schema.TypeString, Optional: true, Description: "Specify AS number of BGP neighbor",
			},
			"remove_private_as": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove private AS number from outbound updates",
			},
			"route_map": {
				Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria to originate default (route-map name)",
			},
			"route_refresh": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Advertise route-refresh capability to this neighbor",
			},
			"shutdown": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Administratively shut down this neighbor",
			},
			"strict_capability_match": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Strict capability negotiation match",
			},
			"timers_holdtime": {
				Type: schema.TypeInt, Optional: true, Default: 90, Description: "Holdtime",
			},
			"timers_keepalive": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Keepalive interval",
			},
			"trunk": {
				Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
			},
			"tunnel": {
				Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
			},
			"update_source_ip": {
				Type: schema.TypeString, Optional: true, Description: "IP address",
			},
			"update_source_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ve": {
				Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
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
func resourceRouterBgpNeighborPeerGroupNeighborCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborPeerGroupNeighborCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborPeerGroupNeighbor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpNeighborPeerGroupNeighborRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpNeighborPeerGroupNeighborUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborPeerGroupNeighborUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborPeerGroupNeighbor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpNeighborPeerGroupNeighborRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpNeighborPeerGroupNeighborDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborPeerGroupNeighborDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborPeerGroupNeighbor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpNeighborPeerGroupNeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborPeerGroupNeighborRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborPeerGroupNeighbor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterBgpNeighborPeerGroupNeighborNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpNeighborPeerGroupNeighborNeighborRouteMapLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborPeerGroupNeighborNeighborRouteMapLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborPeerGroupNeighborNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterBgpNeighborPeerGroupNeighbor(d *schema.ResourceData) edpt.RouterBgpNeighborPeerGroupNeighbor {
	var ret edpt.RouterBgpNeighborPeerGroupNeighbor
	ret.Inst.Activate = d.Get("activate").(int)
	ret.Inst.AdvertisementInterval = d.Get("advertisement_interval").(int)
	ret.Inst.AllowasIn = d.Get("allowas_in").(int)
	ret.Inst.AllowasInCount = d.Get("allowas_in_count").(int)
	ret.Inst.AsOriginationInterval = d.Get("as_origination_interval").(int)
	ret.Inst.Bfd = d.Get("bfd").(int)
	ret.Inst.CollideEstablished = d.Get("collide_established").(int)
	ret.Inst.Connect = d.Get("connect").(int)
	ret.Inst.DefaultOriginate = d.Get("default_originate").(int)
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.DontCapabilityNegotiate = d.Get("dont_capability_negotiate").(int)
	ret.Inst.Dynamic = d.Get("dynamic").(int)
	ret.Inst.EbgpMultihop = d.Get("ebgp_multihop").(int)
	ret.Inst.EbgpMultihopHopCount = d.Get("ebgp_multihop_hop_count").(int)
	ret.Inst.EnforceMultihop = d.Get("enforce_multihop").(int)
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.ExtendedNexthop = d.Get("extended_nexthop").(int)
	ret.Inst.Inbound = d.Get("inbound").(int)
	ret.Inst.Lif = d.Get("lif").(string)
	ret.Inst.Loopback = d.Get("loopback").(int)
	ret.Inst.MaximumPrefix = d.Get("maximum_prefix").(int)
	ret.Inst.MaximumPrefixThres = d.Get("maximum_prefix_thres").(int)
	ret.Inst.Multihop = d.Get("multihop").(int)
	ret.Inst.NeighborRouteMapLists = getSliceRouterBgpNeighborPeerGroupNeighborNeighborRouteMapLists(d.Get("neighbor_route_map_lists").([]interface{}))
	ret.Inst.OverrideCapability = d.Get("override_capability").(int)
	//omit pass_encrypted
	ret.Inst.PassValue = d.Get("pass_value").(string)
	ret.Inst.Passive = d.Get("passive").(int)
	ret.Inst.PeerGroup = d.Get("peer_group").(string)
	ret.Inst.PeerGroupKey = d.Get("peer_group_key").(int)
	ret.Inst.PeerGroupRemoteAs = d.Get("peer_group_remote_as").(string)
	ret.Inst.RemovePrivateAs = d.Get("remove_private_as").(int)
	ret.Inst.RouteMap = d.Get("route_map").(string)
	ret.Inst.RouteRefresh = d.Get("route_refresh").(int)
	ret.Inst.Shutdown = d.Get("shutdown").(int)
	ret.Inst.StrictCapabilityMatch = d.Get("strict_capability_match").(int)
	ret.Inst.TimersHoldtime = d.Get("timers_holdtime").(int)
	ret.Inst.TimersKeepalive = d.Get("timers_keepalive").(int)
	ret.Inst.Trunk = d.Get("trunk").(int)
	ret.Inst.Tunnel = d.Get("tunnel").(int)
	ret.Inst.UpdateSourceIp = d.Get("update_source_ip").(string)
	ret.Inst.UpdateSourceIpv6 = d.Get("update_source_ipv6").(string)
	//omit uuid
	ret.Inst.Ve = d.Get("ve").(int)
	ret.Inst.Weight = d.Get("weight").(int)
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
