package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterRip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_rip`: Routing Information Protocol (RIP)\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterRipCreate,
		UpdateContext: resourceRouterRipUpdate,
		ReadContext:   resourceRouterRipRead,
		DeleteContext: resourceRouterRipDelete,

		Schema: map[string]*schema.Schema{
			"cisco_metric_behavior": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enables updating metric consistent with Cisco; 'disable': Disables updating metric consistent with Cisco;  (Enable/Disable updating metric consistent with Cisco)",
			},
			"default_information": {
				Type: schema.TypeString, Optional: true, Description: "'originate': originate;  (Distribute default route)",
			},
			"default_metric": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set a metric of redistribute routes (Default metric)",
			},
			"distance_list_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"distance": {
							Type: schema.TypeInt, Optional: true, Default: 120, Description: "Administrative distance (Distance value)",
						},
						"distance_ipv4_mask": {
							Type: schema.TypeString, Optional: true, Description: "IP source prefix",
						},
						"distance_acl": {
							Type: schema.TypeString, Optional: true, Description: "Access list name",
						},
					},
				},
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
			"neighbor": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type: schema.TypeString, Optional: true, Description: "Neighbor address",
						},
					},
				},
			},
			"network_addresses": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_ipv4_mask": {
							Type: schema.TypeString, Optional: true, Description: "IP prefix network/length, e.g., 35.0.0.0/8",
						},
					},
				},
			},
			"network_interface_list_cfg": {
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
				Type: schema.TypeInt, Optional: true, Description: "Set the RIP UDP receive buffer size (the RIP UDP receive buffer size value)",
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
										Type: schema.TypeString, Optional: true, Description: "'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'ip-nat-list': IP NAT list; 'ip-nat': IP NAT; 'isis': ISO IS-IS; 'lw4o6': LW4O6 Prefix; 'nat-map': NAT MAP Prefix; 'static-nat': Static NAT; 'ospf': Open Shortest Path First (OSPF); 'static': Static routes;",
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
			"rip_maximum_prefix_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"maximum_prefix": {
							Type: schema.TypeInt, Optional: true, Description: "Set the maximum number of RIP routes",
						},
						"maximum_prefix_thres": {
							Type: schema.TypeInt, Optional: true, Default: 75, Description: "Percentage of maximum routes to generate a warning (Default 75%)",
						},
					},
				},
			},
			"route_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"route": {
							Type: schema.TypeString, Optional: true, Description: "Static route advertisement (debugging purpose) (IP prefix network/length)",
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
			"version": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Set routing protocol version (RIP version)",
			},
		},
	}
}
func resourceRouterRipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterRipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterRip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterRipRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterRipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterRipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterRip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterRipRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterRipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterRipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterRip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterRipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterRipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterRip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterRipDistanceListCfg(d []interface{}) []edpt.RouterRipDistanceListCfg {

	count1 := len(d)
	ret := make([]edpt.RouterRipDistanceListCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterRipDistanceListCfg
		oi.Distance = in["distance"].(int)
		oi.DistanceIpv4Mask = in["distance_ipv4_mask"].(string)
		oi.DistanceAcl = in["distance_acl"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterRipDistributeList1306(d []interface{}) edpt.RouterRipDistributeList1306 {

	count1 := len(d)
	var ret edpt.RouterRipDistributeList1306
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AclCfg = getSliceRouterRipDistributeListAclCfg1307(in["acl_cfg"].([]interface{}))
		//omit uuid
		ret.Prefix = getObjectRouterRipDistributeListPrefix1308(in["prefix"].([]interface{}))
	}
	return ret
}

func getSliceRouterRipDistributeListAclCfg1307(d []interface{}) []edpt.RouterRipDistributeListAclCfg1307 {

	count1 := len(d)
	ret := make([]edpt.RouterRipDistributeListAclCfg1307, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterRipDistributeListAclCfg1307
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

func getObjectRouterRipDistributeListPrefix1308(d []interface{}) edpt.RouterRipDistributeListPrefix1308 {

	count1 := len(d)
	var ret edpt.RouterRipDistributeListPrefix1308
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrefixCfg = getSliceRouterRipDistributeListPrefixPrefixCfg1309(in["prefix_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceRouterRipDistributeListPrefixPrefixCfg1309(d []interface{}) []edpt.RouterRipDistributeListPrefixPrefixCfg1309 {

	count1 := len(d)
	ret := make([]edpt.RouterRipDistributeListPrefixPrefixCfg1309, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterRipDistributeListPrefixPrefixCfg1309
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

func getSliceRouterRipNeighbor(d []interface{}) []edpt.RouterRipNeighbor {

	count1 := len(d)
	ret := make([]edpt.RouterRipNeighbor, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterRipNeighbor
		oi.Value = in["value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterRipNetworkAddresses(d []interface{}) []edpt.RouterRipNetworkAddresses {

	count1 := len(d)
	ret := make([]edpt.RouterRipNetworkAddresses, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterRipNetworkAddresses
		oi.NetworkIpv4Mask = in["network_ipv4_mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterRipNetworkInterfaceListCfg(d []interface{}) []edpt.RouterRipNetworkInterfaceListCfg {

	count1 := len(d)
	ret := make([]edpt.RouterRipNetworkInterfaceListCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterRipNetworkInterfaceListCfg
		oi.Ethernet = in["ethernet"].(int)
		oi.Loopback = in["loopback"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Tunnel = in["tunnel"].(int)
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterRipOffsetList1310(d []interface{}) edpt.RouterRipOffsetList1310 {

	count1 := len(d)
	var ret edpt.RouterRipOffsetList1310
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AclCfg = getSliceRouterRipOffsetListAclCfg1311(in["acl_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceRouterRipOffsetListAclCfg1311(d []interface{}) []edpt.RouterRipOffsetListAclCfg1311 {

	count1 := len(d)
	ret := make([]edpt.RouterRipOffsetListAclCfg1311, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterRipOffsetListAclCfg1311
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

func getSliceRouterRipPassiveInterfaceList(d []interface{}) []edpt.RouterRipPassiveInterfaceList {

	count1 := len(d)
	ret := make([]edpt.RouterRipPassiveInterfaceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterRipPassiveInterfaceList
		oi.Ethernet = in["ethernet"].(int)
		oi.Loopback = in["loopback"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Tunnel = in["tunnel"].(int)
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterRipRedistribute1312(d []interface{}) edpt.RouterRipRedistribute1312 {

	count1 := len(d)
	var ret edpt.RouterRipRedistribute1312
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RedistList = getSliceRouterRipRedistributeRedistList1313(in["redist_list"].([]interface{}))
		ret.VipList = getSliceRouterRipRedistributeVipList1314(in["vip_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceRouterRipRedistributeRedistList1313(d []interface{}) []edpt.RouterRipRedistributeRedistList1313 {

	count1 := len(d)
	ret := make([]edpt.RouterRipRedistributeRedistList1313, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterRipRedistributeRedistList1313
		oi.Type = in["type"].(string)
		oi.Metric = in["metric"].(int)
		oi.RouteMap = in["route_map"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterRipRedistributeVipList1314(d []interface{}) []edpt.RouterRipRedistributeVipList1314 {

	count1 := len(d)
	ret := make([]edpt.RouterRipRedistributeVipList1314, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterRipRedistributeVipList1314
		oi.VipType = in["vip_type"].(string)
		oi.VipMetric = in["vip_metric"].(int)
		oi.VipRouteMap = in["vip_route_map"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterRipRipMaximumPrefixCfg(d []interface{}) edpt.RouterRipRipMaximumPrefixCfg {

	count1 := len(d)
	var ret edpt.RouterRipRipMaximumPrefixCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MaximumPrefix = in["maximum_prefix"].(int)
		ret.MaximumPrefixThres = in["maximum_prefix_thres"].(int)
	}
	return ret
}

func getSliceRouterRipRouteCfg(d []interface{}) []edpt.RouterRipRouteCfg {

	count1 := len(d)
	ret := make([]edpt.RouterRipRouteCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterRipRouteCfg
		oi.Route = in["route"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterRipTimers(d []interface{}) edpt.RouterRipTimers {

	count1 := len(d)
	var ret edpt.RouterRipTimers
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TimersCfg = getObjectRouterRipTimersTimersCfg(in["timers_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRouterRipTimersTimersCfg(d []interface{}) edpt.RouterRipTimersTimersCfg {

	count1 := len(d)
	var ret edpt.RouterRipTimersTimersCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Basic = in["basic"].(int)
		ret.Val2 = in["val_2"].(int)
		ret.Val3 = in["val_3"].(int)
	}
	return ret
}

func dataToEndpointRouterRip(d *schema.ResourceData) edpt.RouterRip {
	var ret edpt.RouterRip
	ret.Inst.CiscoMetricBehavior = d.Get("cisco_metric_behavior").(string)
	ret.Inst.DefaultInformation = d.Get("default_information").(string)
	ret.Inst.DefaultMetric = d.Get("default_metric").(int)
	ret.Inst.DistanceListCfg = getSliceRouterRipDistanceListCfg(d.Get("distance_list_cfg").([]interface{}))
	ret.Inst.DistributeList = getObjectRouterRipDistributeList1306(d.Get("distribute_list").([]interface{}))
	ret.Inst.Neighbor = getSliceRouterRipNeighbor(d.Get("neighbor").([]interface{}))
	ret.Inst.NetworkAddresses = getSliceRouterRipNetworkAddresses(d.Get("network_addresses").([]interface{}))
	ret.Inst.NetworkInterfaceListCfg = getSliceRouterRipNetworkInterfaceListCfg(d.Get("network_interface_list_cfg").([]interface{}))
	ret.Inst.OffsetList = getObjectRouterRipOffsetList1310(d.Get("offset_list").([]interface{}))
	ret.Inst.PassiveInterfaceList = getSliceRouterRipPassiveInterfaceList(d.Get("passive_interface_list").([]interface{}))
	ret.Inst.RecvBufferSize = d.Get("recv_buffer_size").(int)
	ret.Inst.Redistribute = getObjectRouterRipRedistribute1312(d.Get("redistribute").([]interface{}))
	ret.Inst.RipMaximumPrefixCfg = getObjectRouterRipRipMaximumPrefixCfg(d.Get("rip_maximum_prefix_cfg").([]interface{}))
	ret.Inst.RouteCfg = getSliceRouterRipRouteCfg(d.Get("route_cfg").([]interface{}))
	ret.Inst.Timers = getObjectRouterRipTimers(d.Get("timers").([]interface{}))
	//omit uuid
	ret.Inst.Version = d.Get("version").(int)
	return ret
}
