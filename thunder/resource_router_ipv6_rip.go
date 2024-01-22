package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterIpv6Rip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_ipv6_rip`: Routing Information Protocol (RIPng)\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterIpv6RipCreate,
		UpdateContext: resourceRouterIpv6RipUpdate,
		ReadContext:   resourceRouterIpv6RipRead,
		DeleteContext: resourceRouterIpv6RipDelete,

		Schema: map[string]*schema.Schema{
			"aggregate_address_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aggregate_address": {
							Type: schema.TypeString, Optional: true, Description: "Set aggregate RIP route announcement (Aggregate network)",
						},
					},
				},
			},
			"cisco_metric_behavior": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enables updating metric consistent with Cisco; 'disable': Disables updating metric consistent with Cisco;  (Enable/Disable updating metric consistent with Cisco)",
			},
			"default_information": {
				Type: schema.TypeString, Optional: true, Description: "'originate': originate;  (Distribute default route)",
			},
			"default_metric": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set a metric of redistribute routes (Default metric)",
			},
			"distribute_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acl": {
										Type: schema.TypeString, Optional: true, Description: "Access-list name",
									},
									"acl_direction": {
										Type: schema.TypeString, Optional: true, Description: "'in': Filter incoming routing updates; 'out': Filter outgoing routing updates;",
									},
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
									},
									"loopback": {
										Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
									},
									"tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"prefix": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"prefix_list": {
													Type: schema.TypeString, Optional: true, Description: "Filter prefixes in routing updates (Name of a prefix list)",
												},
												"prefix_list_direction": {
													Type: schema.TypeString, Optional: true, Description: "'in': Filter incoming routing updates; 'out': Filter outgoing routing updates;",
												},
												"ethernet": {
													Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
												},
												"loopback": {
													Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
												},
												"trunk": {
													Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
												},
												"tunnel": {
													Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
												},
												"ve": {
													Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
												},
											},
										},
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
			"offset_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acl": {
										Type: schema.TypeString, Optional: true, Description: "Access-list name",
									},
									"offset_list_direction": {
										Type: schema.TypeString, Optional: true, Description: "'in': Filter incoming updates; 'out': Filter outgoing updates;",
									},
									"metric": {
										Type: schema.TypeInt, Optional: true, Description: "Metric value",
									},
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
									},
									"loopback": {
										Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
									},
									"tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"passive_interface_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethernet": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
						},
						"loopback": {
							Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
						},
						"trunk": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
						},
						"tunnel": {
							Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
						},
						"ve": {
							Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
						},
					},
				},
			},
			"recv_buffer_size": {
				Type: schema.TypeInt, Optional: true, Description: "Set the RIPNG UDP receive buffer size (the RIPNG UDP receive buffer size value)",
			},
			"redistribute": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"redist_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type: schema.TypeString, Optional: true, Description: "'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'ip-nat-list': IP NAT list; 'ip-nat': IP NAT; 'isis': ISO IS-IS; 'lw4o6': LW4O6 Prefix; 'nat-map': NAT MAP Prefix; 'nat64': NAT64 Prefix; 'static-nat': Static NAT; 'ospf': Open Shortest Path First (OSPF); 'static': Static routes;",
									},
									"metric": {
										Type: schema.TypeInt, Optional: true, Description: "Metric for redistributed routes (metric value)",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"vip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vip_type": {
										Type: schema.TypeString, Optional: true, Description: "'only-flagged': Selected Virtual IP (VIP); 'only-not-flagged': Only not flagged;",
									},
									"vip_metric": {
										Type: schema.TypeInt, Optional: true, Description: "Metric for redistributed routes (metric value)",
									},
									"vip_route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ripng_neighbor": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ripng_neighbor_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"neighbor_link_local_addr": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 link-local address",
									},
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
									},
									"loopback": {
										Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
									},
									"tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
									},
								},
							},
						},
					},
				},
			},
			"route_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"route": {
							Type: schema.TypeString, Optional: true, Description: "Static route advertisement (debugging purpose) (IP prefix)",
						},
					},
				},
			},
			"route_map": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"map_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"map": {
										Type: schema.TypeString, Optional: true, Description: "Route map name",
									},
									"route_map_direction": {
										Type: schema.TypeString, Optional: true, Description: "'in': Route map set for input filtering; 'out': Route map set for output filtering;",
									},
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
									},
									"loopback": {
										Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
									},
									"tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"timers": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timers_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"basic": {
										Type: schema.TypeInt, Optional: true, Description: "Basic routing protocol update timers (Routing table update timer value in second. Default is 30)",
									},
									"val_2": {
										Type: schema.TypeInt, Optional: true, Description: "Routing information timeout timer. Default is 180",
									},
									"val_3": {
										Type: schema.TypeInt, Optional: true, Description: "Garbage collection timer. Default is 120",
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
		},
	}
}
func resourceRouterIpv6RipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6Rip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6RipRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterIpv6RipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6Rip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6RipRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterIpv6RipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6Rip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterIpv6RipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6Rip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterIpv6RipAggregateAddressCfg(d []interface{}) []edpt.RouterIpv6RipAggregateAddressCfg {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipAggregateAddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipAggregateAddressCfg
		oi.AggregateAddress = in["aggregate_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIpv6RipDistributeList1257(d []interface{}) edpt.RouterIpv6RipDistributeList1257 {

	count1 := len(d)
	var ret edpt.RouterIpv6RipDistributeList1257
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AclCfg = getSliceRouterIpv6RipDistributeListAclCfg1258(in["acl_cfg"].([]interface{}))
		//omit uuid
		ret.Prefix = getObjectRouterIpv6RipDistributeListPrefix1259(in["prefix"].([]interface{}))
	}
	return ret
}

func getSliceRouterIpv6RipDistributeListAclCfg1258(d []interface{}) []edpt.RouterIpv6RipDistributeListAclCfg1258 {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipDistributeListAclCfg1258, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipDistributeListAclCfg1258
		oi.Acl = in["acl"].(string)
		oi.AclDirection = in["acl_direction"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Loopback = in["loopback"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Tunnel = in["tunnel"].(int)
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIpv6RipDistributeListPrefix1259(d []interface{}) edpt.RouterIpv6RipDistributeListPrefix1259 {

	count1 := len(d)
	var ret edpt.RouterIpv6RipDistributeListPrefix1259
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrefixCfg = getSliceRouterIpv6RipDistributeListPrefixPrefixCfg1260(in["prefix_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceRouterIpv6RipDistributeListPrefixPrefixCfg1260(d []interface{}) []edpt.RouterIpv6RipDistributeListPrefixPrefixCfg1260 {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipDistributeListPrefixPrefixCfg1260, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipDistributeListPrefixPrefixCfg1260
		oi.PrefixList = in["prefix_list"].(string)
		oi.PrefixListDirection = in["prefix_list_direction"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Loopback = in["loopback"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Tunnel = in["tunnel"].(int)
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIpv6RipOffsetList1261(d []interface{}) edpt.RouterIpv6RipOffsetList1261 {

	count1 := len(d)
	var ret edpt.RouterIpv6RipOffsetList1261
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AclCfg = getSliceRouterIpv6RipOffsetListAclCfg1262(in["acl_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceRouterIpv6RipOffsetListAclCfg1262(d []interface{}) []edpt.RouterIpv6RipOffsetListAclCfg1262 {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipOffsetListAclCfg1262, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipOffsetListAclCfg1262
		oi.Acl = in["acl"].(string)
		oi.OffsetListDirection = in["offset_list_direction"].(string)
		oi.Metric = in["metric"].(int)
		oi.Ethernet = in["ethernet"].(int)
		oi.Loopback = in["loopback"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Tunnel = in["tunnel"].(int)
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6RipPassiveInterfaceList(d []interface{}) []edpt.RouterIpv6RipPassiveInterfaceList {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipPassiveInterfaceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipPassiveInterfaceList
		oi.Ethernet = in["ethernet"].(int)
		oi.Loopback = in["loopback"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Tunnel = in["tunnel"].(int)
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIpv6RipRedistribute1263(d []interface{}) edpt.RouterIpv6RipRedistribute1263 {

	count1 := len(d)
	var ret edpt.RouterIpv6RipRedistribute1263
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RedistList = getSliceRouterIpv6RipRedistributeRedistList1264(in["redist_list"].([]interface{}))
		ret.VipList = getSliceRouterIpv6RipRedistributeVipList1265(in["vip_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceRouterIpv6RipRedistributeRedistList1264(d []interface{}) []edpt.RouterIpv6RipRedistributeRedistList1264 {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipRedistributeRedistList1264, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipRedistributeRedistList1264
		oi.Type = in["type"].(string)
		oi.Metric = in["metric"].(int)
		oi.RouteMap = in["route_map"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6RipRedistributeVipList1265(d []interface{}) []edpt.RouterIpv6RipRedistributeVipList1265 {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipRedistributeVipList1265, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipRedistributeVipList1265
		oi.VipType = in["vip_type"].(string)
		oi.VipMetric = in["vip_metric"].(int)
		oi.VipRouteMap = in["vip_route_map"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIpv6RipRipngNeighbor(d []interface{}) edpt.RouterIpv6RipRipngNeighbor {

	count1 := len(d)
	var ret edpt.RouterIpv6RipRipngNeighbor
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RipngNeighborCfg = getSliceRouterIpv6RipRipngNeighborRipngNeighborCfg(in["ripng_neighbor_cfg"].([]interface{}))
	}
	return ret
}

func getSliceRouterIpv6RipRipngNeighborRipngNeighborCfg(d []interface{}) []edpt.RouterIpv6RipRipngNeighborRipngNeighborCfg {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipRipngNeighborRipngNeighborCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipRipngNeighborRipngNeighborCfg
		oi.NeighborLinkLocalAddr = in["neighbor_link_local_addr"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Loopback = in["loopback"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Tunnel = in["tunnel"].(int)
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6RipRouteCfg(d []interface{}) []edpt.RouterIpv6RipRouteCfg {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipRouteCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipRouteCfg
		oi.Route = in["route"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIpv6RipRouteMap1266(d []interface{}) edpt.RouterIpv6RipRouteMap1266 {

	count1 := len(d)
	var ret edpt.RouterIpv6RipRouteMap1266
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MapCfg = getSliceRouterIpv6RipRouteMapMapCfg1267(in["map_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceRouterIpv6RipRouteMapMapCfg1267(d []interface{}) []edpt.RouterIpv6RipRouteMapMapCfg1267 {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipRouteMapMapCfg1267, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipRouteMapMapCfg1267
		oi.Map = in["map"].(string)
		oi.RouteMapDirection = in["route_map_direction"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Loopback = in["loopback"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Tunnel = in["tunnel"].(int)
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIpv6RipTimers(d []interface{}) edpt.RouterIpv6RipTimers {

	count1 := len(d)
	var ret edpt.RouterIpv6RipTimers
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TimersCfg = getObjectRouterIpv6RipTimersTimersCfg(in["timers_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRouterIpv6RipTimersTimersCfg(d []interface{}) edpt.RouterIpv6RipTimersTimersCfg {

	count1 := len(d)
	var ret edpt.RouterIpv6RipTimersTimersCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Basic = in["basic"].(int)
		ret.Val2 = in["val_2"].(int)
		ret.Val3 = in["val_3"].(int)
	}
	return ret
}

func dataToEndpointRouterIpv6Rip(d *schema.ResourceData) edpt.RouterIpv6Rip {
	var ret edpt.RouterIpv6Rip
	ret.Inst.AggregateAddressCfg = getSliceRouterIpv6RipAggregateAddressCfg(d.Get("aggregate_address_cfg").([]interface{}))
	ret.Inst.CiscoMetricBehavior = d.Get("cisco_metric_behavior").(string)
	ret.Inst.DefaultInformation = d.Get("default_information").(string)
	ret.Inst.DefaultMetric = d.Get("default_metric").(int)
	ret.Inst.DistributeList = getObjectRouterIpv6RipDistributeList1257(d.Get("distribute_list").([]interface{}))
	ret.Inst.OffsetList = getObjectRouterIpv6RipOffsetList1261(d.Get("offset_list").([]interface{}))
	ret.Inst.PassiveInterfaceList = getSliceRouterIpv6RipPassiveInterfaceList(d.Get("passive_interface_list").([]interface{}))
	ret.Inst.RecvBufferSize = d.Get("recv_buffer_size").(int)
	ret.Inst.Redistribute = getObjectRouterIpv6RipRedistribute1263(d.Get("redistribute").([]interface{}))
	ret.Inst.RipngNeighbor = getObjectRouterIpv6RipRipngNeighbor(d.Get("ripng_neighbor").([]interface{}))
	ret.Inst.RouteCfg = getSliceRouterIpv6RipRouteCfg(d.Get("route_cfg").([]interface{}))
	ret.Inst.RouteMap = getObjectRouterIpv6RipRouteMap1266(d.Get("route_map").([]interface{}))
	ret.Inst.Timers = getObjectRouterIpv6RipTimers(d.Get("timers").([]interface{}))
	//omit uuid
	return ret
}
