package thunder

//Thunder resource RouterBgpNeighborIpv6Neighbor

import (
	"fmt"
	"strconv"
	"strings"
	"util"
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpNeighborIpv6Neighbor() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterBgpNeighborIpv6NeighborCreate,
		UpdateContext: resourceRouterBgpNeighborIpv6NeighborUpdate,
		ReadContext:   resourceRouterBgpNeighborIpv6NeighborRead,
		DeleteContext: resourceRouterBgpNeighborIpv6NeighborDelete,
		Schema: map[string]*schema.Schema{
			"neighbor_ipv6": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"nbr_remote_as": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"peer_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"activate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"advertisement_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"allowas_in": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"allowas_in_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"as_origination_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dynamic": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"prefix_list_direction": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"route_refresh": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"graceful_restart": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"extended_nexthop": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"collide_established": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"default_originate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"route_map": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"disallow_infinite_holdtime": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"distribute_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"distribute_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"distribute_list_direction": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"acos_application_only": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dont_capability_negotiate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ebgp_multihop": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ebgp_multihop_hop_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"enforce_multihop": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bfd": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"multihop": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"key_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"key_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"bfd_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"neighbor_filter_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"filter_list_direction": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"maximum_prefix": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"maximum_prefix_thres": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"next_hop_self": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"override_capability": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"pass_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"passive": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"neighbor_prefix_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nbr_prefix_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"nbr_prefix_list_direction": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"remove_private_as": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"neighbor_route_map_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nbr_route_map": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"nbr_rmap_direction": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"send_community_val": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"inbound": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shutdown": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"strict_capability_match": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"timers_keepalive": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"timers_holdtime": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"connect": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"unsuppress_map": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"update_source_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"update_source_ipv6": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ethernet": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"loopback": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ve": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"trunk": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"lif": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"tunnel": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"weight": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"as_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"process_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sequence": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceRouterBgpNeighborIpv6NeighborCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpNeighborIpv6Neighbor (Inside resourceRouterBgpNeighborIpv6NeighborCreate) ")
		name1 := d.Get("as_number").(int)
		name2 := d.Get("neighbor_ipv6").(string)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpNeighborIpv6Neighbor(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNeighborIpv6Neighbor --")
		d.SetId(strconv.Itoa(name1) + "," + name2)
		err := go_thunder.PostRouterBgpNeighborIpv6Neighbor(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpNeighborIpv6NeighborRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpNeighborIpv6NeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading RouterBgpNeighborIpv6Neighbor (Inside resourceRouterBgpNeighborIpv6NeighborRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpNeighborIpv6Neighbor(client.Token, name1, name2, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return nil
}

func resourceRouterBgpNeighborIpv6NeighborUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying RouterBgpNeighborIpv6Neighbor   (Inside resourceRouterBgpNeighborIpv6NeighborUpdate) ")
		data := dataToRouterBgpNeighborIpv6Neighbor(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNeighborIpv6Neighbor ")
		err := go_thunder.PutRouterBgpNeighborIpv6Neighbor(client.Token, name1, name2, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpNeighborIpv6NeighborRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpNeighborIpv6NeighborDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceRouterBgpNeighborIpv6NeighborDelete) " + name1)
		err := go_thunder.DeleteRouterBgpNeighborIpv6Neighbor(client.Token, name1, name2, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToRouterBgpNeighborIpv6Neighbor(d *schema.ResourceData) go_thunder.RouterBgpNeighborIpv6Neighbor {
	var vc go_thunder.RouterBgpNeighborIpv6Neighbor
	var c go_thunder.RouterBgpNeighborIpv6NeighborInstance
	c.NeighborIpv6 = d.Get("neighbor_ipv6").(string)
	c.NbrRemoteAs = d.Get("nbr_remote_as").(int)
	c.PeerGroupName = d.Get("peer_group_name").(string)
	c.Activate = d.Get("activate").(int)
	c.AdvertisementInterval = d.Get("advertisement_interval").(int)
	c.AllowasIn = d.Get("allowas_in").(int)
	c.AllowasInCount = d.Get("allowas_in_count").(int)
	c.AsOriginationInterval = d.Get("as_origination_interval").(int)
	c.Dynamic = d.Get("dynamic").(int)
	c.PrefixListDirection = d.Get("prefix_list_direction").(string)
	c.RouteRefresh = d.Get("route_refresh").(int)
	c.GracefulRestart = d.Get("graceful_restart").(int)
	c.ExtendedNexthop = d.Get("extended_nexthop").(int)
	c.CollideEstablished = d.Get("collide_established").(int)
	c.DefaultOriginate = d.Get("default_originate").(int)
	c.RouteMap = d.Get("route_map").(string)
	c.Description = d.Get("description").(string)
	c.DisallowInfiniteHoldtime = d.Get("disallow_infinite_holdtime").(int)

	DistributeListsCount := d.Get("distribute_lists.#").(int)
	c.DistributeList = make([]go_thunder.RouterBgpNeighborIpv6NeighborDistributeLists, 0, DistributeListsCount)

	for i := 0; i < DistributeListsCount; i++ {
		var obj1 go_thunder.RouterBgpNeighborIpv6NeighborDistributeLists
		prefix1 := fmt.Sprintf("distribute_lists.%d.", i)
		obj1.DistributeList = d.Get(prefix1 + "distribute_list").(string)
		obj1.DistributeListDirection = d.Get(prefix1 + "distribute_list_direction").(string)
		c.DistributeList = append(c.DistributeList, obj1)
	}

	c.AcosApplicationOnly = d.Get("acos_application_only").(int)
	c.DontCapabilityNegotiate = d.Get("dont_capability_negotiate").(int)
	c.EbgpMultihop = d.Get("ebgp_multihop").(int)
	c.EbgpMultihopHopCount = d.Get("ebgp_multihop_hop_count").(int)
	c.EnforceMultihop = d.Get("enforce_multihop").(int)
	c.Bfd = d.Get("bfd").(int)
	c.Multihop = d.Get("multihop").(int)
	c.KeyID = d.Get("key_id").(int)
	c.KeyType = d.Get("key_type").(string)
	c.BfdValue = d.Get("bfd_value").(string)

	NeighborFilterListsCount := d.Get("neighbor_filter_lists.#").(int)
	c.FilterList = make([]go_thunder.RouterBgpNeighborIpv6NeighborNeighborFilterLists, 0, NeighborFilterListsCount)

	for i := 0; i < NeighborFilterListsCount; i++ {
		var obj2 go_thunder.RouterBgpNeighborIpv6NeighborNeighborFilterLists
		prefix2 := fmt.Sprintf("neighbor_filter_lists.%d.", i)
		obj2.FilterList = d.Get(prefix2 + "filter_list").(string)
		obj2.FilterListDirection = d.Get(prefix2 + "filter_list_direction").(string)
		c.FilterList = append(c.FilterList, obj2)
	}

	c.MaximumPrefix = d.Get("maximum_prefix").(int)
	c.MaximumPrefixThres = d.Get("maximum_prefix_thres").(int)
	c.NextHopSelf = d.Get("next_hop_self").(int)
	c.OverrideCapability = d.Get("override_capability").(int)
	c.PassValue = d.Get("pass_value").(string)
	c.Passive = d.Get("passive").(int)

	NeighborPrefixListsCount := d.Get("neighbor_prefix_lists.#").(int)
	c.NbrPrefixList = make([]go_thunder.RouterBgpNeighborIpv6NeighborNeighborPrefixLists, 0, NeighborPrefixListsCount)

	for i := 0; i < NeighborPrefixListsCount; i++ {
		var obj3 go_thunder.RouterBgpNeighborIpv6NeighborNeighborPrefixLists
		prefix3 := fmt.Sprintf("neighbor_prefix_lists.%d.", i)
		obj3.NbrPrefixList = d.Get(prefix3 + "nbr_prefix_list").(string)
		obj3.NbrPrefixListDirection = d.Get(prefix3 + "nbr_prefix_list_direction").(string)
		c.NbrPrefixList = append(c.NbrPrefixList, obj3)
	}

	c.RemovePrivateAs = d.Get("remove_private_as").(int)

	NeighborRouteMapListsCount := d.Get("neighbor_route_map_lists.#").(int)
	c.NbrRouteMap = make([]go_thunder.RouterBgpNeighborIpv6NeighborNeighborRouteMapLists, 0, NeighborRouteMapListsCount)

	for i := 0; i < NeighborRouteMapListsCount; i++ {
		var obj4 go_thunder.RouterBgpNeighborIpv6NeighborNeighborRouteMapLists
		prefix4 := fmt.Sprintf("neighbor_route_map_lists.%d.", i)
		obj4.NbrRouteMap = d.Get(prefix4 + "nbr_route_map").(string)
		obj4.NbrRmapDirection = d.Get(prefix4 + "nbr_rmap_direction").(string)
		c.NbrRouteMap = append(c.NbrRouteMap, obj4)
	}

	c.SendCommunityVal = d.Get("send_community_val").(string)
	c.Inbound = d.Get("inbound").(int)
	c.Shutdown = d.Get("shutdown").(int)
	c.StrictCapabilityMatch = d.Get("strict_capability_match").(int)
	c.TimersKeepalive = d.Get("timers_keepalive").(int)
	c.TimersHoldtime = d.Get("timers_holdtime").(int)
	c.Connect = d.Get("connect").(int)
	c.UnsuppressMap = d.Get("unsuppress_map").(string)
	c.UpdateSourceIP = d.Get("update_source_ip").(string)
	c.UpdateSourceIpv6 = d.Get("update_source_ipv6").(string)
	c.Ethernet = d.Get("ethernet").(int)
	c.Loopback = d.Get("loopback").(int)
	c.Ve = d.Get("ve").(int)
	c.Trunk = d.Get("trunk").(int)
	c.Lif = d.Get("lif").(int)
	c.Tunnel = d.Get("tunnel").(int)
	c.Weight = d.Get("weight").(int)

	vc.NeighborIpv6 = c
	return vc
}
