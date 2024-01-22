package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpNeighborIpv4Neighbor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_neighbor_ipv4_neighbor`: Specify a ipv4 neighbor router\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpNeighborIpv4NeighborCreate,
		UpdateContext: resourceRouterBgpNeighborIpv4NeighborUpdate,
		ReadContext:   resourceRouterBgpNeighborIpv4NeighborRead,
		DeleteContext: resourceRouterBgpNeighborIpv4NeighborDelete,

		Schema: map[string]*schema.Schema{
			"acos_application_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send BGP update to ACOS application",
			},
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
			"bfd_value": {
				Type: schema.TypeString, Optional: true, Description: "Key String",
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
			"disallow_infinite_holdtime": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "BGP per neighbor disallow-infinite-holdtime",
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
			"graceful_restart": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable graceful-restart helper for this neighbor",
			},
			"inbound": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow inbound soft reconfiguration for this neighbor",
			},
			"key_id": {
				Type: schema.TypeInt, Optional: true, Description: "Key ID",
			},
			"key_type": {
				Type: schema.TypeString, Optional: true, Description: "'md5': md5; 'meticulous-md5': meticulous-md5; 'meticulous-sha1': meticulous-sha1; 'sha1': sha1; 'simple': simple;  (Keyed MD5/Meticulous Keyed MD5/Meticulous Keyed SHA1/Keyed SHA1/Simple Password)",
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
			"nbr_remote_as": {
				Type: schema.TypeString, Optional: true, Description: "Specify AS number of BGP neighbor",
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
			"neighbor_ipv4": {
				Type: schema.TypeString, Required: true, Description: "Neighbor address",
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
			"override_capability": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Override capability negotiation result",
			},
			"pass_value": {
				Type: schema.TypeString, Optional: true, Description: "Key String",
			},
			"passive": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't send open messages to this neighbor",
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
			"route_refresh": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Advertise route-refresh capability to this neighbor",
			},
			"send_community_val": {
				Type: schema.TypeString, Optional: true, Default: "both", Description: "'all': Send Standard, Extended, and Large Community attributes; 'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes; 'large': Send Large Community attributes;",
			},
			"shutdown": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Administratively shut down this neighbor",
			},
			"strict_capability_match": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Strict capability negotiation match",
			},
			"telemetry": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send BGP update to telemetry db",
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
			"unsuppress_map": {
				Type: schema.TypeString, Optional: true, Description: "Route-map to selectively unsuppress suppressed routes (Name of route map)",
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
func resourceRouterBgpNeighborIpv4NeighborCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborIpv4NeighborCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborIpv4Neighbor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpNeighborIpv4NeighborRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpNeighborIpv4NeighborUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborIpv4NeighborUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborIpv4Neighbor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpNeighborIpv4NeighborRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpNeighborIpv4NeighborDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborIpv4NeighborDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborIpv4Neighbor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpNeighborIpv4NeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpNeighborIpv4NeighborRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpNeighborIpv4Neighbor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterBgpNeighborIpv4NeighborDistributeLists(d []interface{}) []edpt.RouterBgpNeighborIpv4NeighborDistributeLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv4NeighborDistributeLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv4NeighborDistributeLists
		oi.DistributeList = in["distribute_list"].(string)
		oi.DistributeListDirection = in["distribute_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv4NeighborNeighborFilterLists(d []interface{}) []edpt.RouterBgpNeighborIpv4NeighborNeighborFilterLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv4NeighborNeighborFilterLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv4NeighborNeighborFilterLists
		oi.FilterList = in["filter_list"].(string)
		oi.FilterListDirection = in["filter_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv4NeighborNeighborPrefixLists(d []interface{}) []edpt.RouterBgpNeighborIpv4NeighborNeighborPrefixLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv4NeighborNeighborPrefixLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv4NeighborNeighborPrefixLists
		oi.NbrPrefixList = in["nbr_prefix_list"].(string)
		oi.NbrPrefixListDirection = in["nbr_prefix_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv4NeighborNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpNeighborIpv4NeighborNeighborRouteMapLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv4NeighborNeighborRouteMapLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv4NeighborNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterBgpNeighborIpv4Neighbor(d *schema.ResourceData) edpt.RouterBgpNeighborIpv4Neighbor {
	var ret edpt.RouterBgpNeighborIpv4Neighbor
	ret.Inst.AcosApplicationOnly = d.Get("acos_application_only").(int)
	ret.Inst.Activate = d.Get("activate").(int)
	ret.Inst.AdvertisementInterval = d.Get("advertisement_interval").(int)
	ret.Inst.AllowasIn = d.Get("allowas_in").(int)
	ret.Inst.AllowasInCount = d.Get("allowas_in_count").(int)
	ret.Inst.AsOriginationInterval = d.Get("as_origination_interval").(int)
	ret.Inst.Bfd = d.Get("bfd").(int)
	//omit bfd_encrypted
	ret.Inst.BfdValue = d.Get("bfd_value").(string)
	ret.Inst.CollideEstablished = d.Get("collide_established").(int)
	ret.Inst.Connect = d.Get("connect").(int)
	ret.Inst.DefaultOriginate = d.Get("default_originate").(int)
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.DisallowInfiniteHoldtime = d.Get("disallow_infinite_holdtime").(int)
	ret.Inst.DistributeLists = getSliceRouterBgpNeighborIpv4NeighborDistributeLists(d.Get("distribute_lists").([]interface{}))
	ret.Inst.DontCapabilityNegotiate = d.Get("dont_capability_negotiate").(int)
	ret.Inst.Dynamic = d.Get("dynamic").(int)
	ret.Inst.EbgpMultihop = d.Get("ebgp_multihop").(int)
	ret.Inst.EbgpMultihopHopCount = d.Get("ebgp_multihop_hop_count").(int)
	ret.Inst.EnforceMultihop = d.Get("enforce_multihop").(int)
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.GracefulRestart = d.Get("graceful_restart").(int)
	ret.Inst.Inbound = d.Get("inbound").(int)
	ret.Inst.KeyId = d.Get("key_id").(int)
	ret.Inst.KeyType = d.Get("key_type").(string)
	ret.Inst.Lif = d.Get("lif").(string)
	ret.Inst.Loopback = d.Get("loopback").(int)
	ret.Inst.MaximumPrefix = d.Get("maximum_prefix").(int)
	ret.Inst.MaximumPrefixThres = d.Get("maximum_prefix_thres").(int)
	ret.Inst.Multihop = d.Get("multihop").(int)
	ret.Inst.NbrRemoteAs = d.Get("nbr_remote_as").(string)
	ret.Inst.NeighborFilterLists = getSliceRouterBgpNeighborIpv4NeighborNeighborFilterLists(d.Get("neighbor_filter_lists").([]interface{}))
	ret.Inst.NeighborIpv4 = d.Get("neighbor_ipv4").(string)
	ret.Inst.NeighborPrefixLists = getSliceRouterBgpNeighborIpv4NeighborNeighborPrefixLists(d.Get("neighbor_prefix_lists").([]interface{}))
	ret.Inst.NeighborRouteMapLists = getSliceRouterBgpNeighborIpv4NeighborNeighborRouteMapLists(d.Get("neighbor_route_map_lists").([]interface{}))
	ret.Inst.NextHopSelf = d.Get("next_hop_self").(int)
	ret.Inst.OverrideCapability = d.Get("override_capability").(int)
	//omit pass_encrypted
	ret.Inst.PassValue = d.Get("pass_value").(string)
	ret.Inst.Passive = d.Get("passive").(int)
	ret.Inst.PeerGroupName = d.Get("peer_group_name").(string)
	ret.Inst.PrefixListDirection = d.Get("prefix_list_direction").(string)
	ret.Inst.RemovePrivateAs = d.Get("remove_private_as").(int)
	ret.Inst.RestartMin = d.Get("restart_min").(int)
	ret.Inst.RouteMap = d.Get("route_map").(string)
	ret.Inst.RouteRefresh = d.Get("route_refresh").(int)
	ret.Inst.SendCommunityVal = d.Get("send_community_val").(string)
	ret.Inst.Shutdown = d.Get("shutdown").(int)
	ret.Inst.StrictCapabilityMatch = d.Get("strict_capability_match").(int)
	ret.Inst.Telemetry = d.Get("telemetry").(int)
	ret.Inst.TimersHoldtime = d.Get("timers_holdtime").(int)
	ret.Inst.TimersKeepalive = d.Get("timers_keepalive").(int)
	ret.Inst.Trunk = d.Get("trunk").(int)
	ret.Inst.Tunnel = d.Get("tunnel").(int)
	ret.Inst.UnsuppressMap = d.Get("unsuppress_map").(string)
	ret.Inst.UpdateSourceIp = d.Get("update_source_ip").(string)
	ret.Inst.UpdateSourceIpv6 = d.Get("update_source_ipv6").(string)
	//omit uuid
	ret.Inst.Ve = d.Get("ve").(int)
	ret.Inst.Weight = d.Get("weight").(int)
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
