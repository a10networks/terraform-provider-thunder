package thunder

//Thunder resource RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborCreate,
		UpdateContext: resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborUpdate,
		ReadContext:   resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborRead,
		DeleteContext: resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborDelete,
		Schema: map[string]*schema.Schema{
			"peer_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"activate": {
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
			"prefix_list_direction": {
				Type:        schema.TypeString,
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
			"unsuppress_map": {
				Type:        schema.TypeString,
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
		},
	}
}

func resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor (Inside resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborCreate) ")
		name1 := d.Get("as_number").(int)
		name2 := d.Get("peer_group").(string)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor --")
		d.SetId(strconv.Itoa(name1) + "," + name2)
		err := go_thunder.PostRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor (Inside resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor(client.Token, name1, name2, client.Host)
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

func resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor   (Inside resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborUpdate) ")
		data := dataToRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor ")
		err := go_thunder.PutRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor(client.Token, name1, name2, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborRead(ctx, d, meta)

}

func dataToRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor(d *schema.ResourceData) go_thunder.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor {
	var vc go_thunder.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor
	var c go_thunder.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborInstance
	c.PeerGroup = d.Get("peer_group").(string)
	c.Activate = d.Get("activate").(int)
	c.AllowasIn = d.Get("allowas_in").(int)
	c.AllowasInCount = d.Get("allowas_in_count").(int)
	c.PrefixListDirection = d.Get("prefix_list_direction").(string)
	c.DefaultOriginate = d.Get("default_originate").(int)
	c.RouteMap = d.Get("route_map").(string)

	DistributeListsCount := d.Get("distribute_lists.#").(int)
	c.DistributeList = make([]go_thunder.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborDistributeLists, 0, DistributeListsCount)

	for i := 0; i < DistributeListsCount; i++ {
		var obj1 go_thunder.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborDistributeLists
		prefix1 := fmt.Sprintf("distribute_lists.%d.", i)
		obj1.DistributeList = d.Get(prefix1 + "distribute_list").(string)
		obj1.DistributeListDirection = d.Get(prefix1 + "distribute_list_direction").(string)
		c.DistributeList = append(c.DistributeList, obj1)
	}

	NeighborFilterListsCount := d.Get("neighbor_filter_lists.#").(int)
	c.FilterList = make([]go_thunder.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborFilterLists, 0, NeighborFilterListsCount)

	for i := 0; i < NeighborFilterListsCount; i++ {
		var obj2 go_thunder.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborFilterLists
		prefix2 := fmt.Sprintf("neighbor_filter_lists.%d.", i)
		obj2.FilterList = d.Get(prefix2 + "filter_list").(string)
		obj2.FilterListDirection = d.Get(prefix2 + "filter_list_direction").(string)
		c.FilterList = append(c.FilterList, obj2)
	}

	c.MaximumPrefix = d.Get("maximum_prefix").(int)
	c.MaximumPrefixThres = d.Get("maximum_prefix_thres").(int)
	c.NextHopSelf = d.Get("next_hop_self").(int)

	NeighborPrefixListsCount := d.Get("neighbor_prefix_lists.#").(int)
	c.NbrPrefixList = make([]go_thunder.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborPrefixLists, 0, NeighborPrefixListsCount)

	for i := 0; i < NeighborPrefixListsCount; i++ {
		var obj3 go_thunder.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborPrefixLists
		prefix3 := fmt.Sprintf("neighbor_prefix_lists.%d.", i)
		obj3.NbrPrefixList = d.Get(prefix3 + "nbr_prefix_list").(string)
		obj3.NbrPrefixListDirection = d.Get(prefix3 + "nbr_prefix_list_direction").(string)
		c.NbrPrefixList = append(c.NbrPrefixList, obj3)
	}

	c.RemovePrivateAs = d.Get("remove_private_as").(int)

	NeighborRouteMapListsCount := d.Get("neighbor_route_map_lists.#").(int)
	c.NbrRouteMap = make([]go_thunder.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborRouteMapLists, 0, NeighborRouteMapListsCount)

	for i := 0; i < NeighborRouteMapListsCount; i++ {
		var obj4 go_thunder.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborRouteMapLists
		prefix4 := fmt.Sprintf("neighbor_route_map_lists.%d.", i)
		obj4.NbrRouteMap = d.Get(prefix4 + "nbr_route_map").(string)
		obj4.NbrRmapDirection = d.Get(prefix4 + "nbr_rmap_direction").(string)
		c.NbrRouteMap = append(c.NbrRouteMap, obj4)
	}

	c.SendCommunityVal = d.Get("send_community_val").(string)
	c.Inbound = d.Get("inbound").(int)
	c.UnsuppressMap = d.Get("unsuppress_map").(string)
	c.Weight = d.Get("weight").(int)

	vc.PeerGroup = c
	return vc
}
