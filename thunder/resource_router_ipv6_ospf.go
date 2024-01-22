package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterIpv6Ospf() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_ipv6_ospf`: Open Shortest Path First (OSPFv3)\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterIpv6OspfCreate,
		UpdateContext: resourceRouterIpv6OspfUpdate,
		ReadContext:   resourceRouterIpv6OspfRead,
		DeleteContext: resourceRouterIpv6OspfDelete,

		Schema: map[string]*schema.Schema{
			"abr_type_option": {
				Type: schema.TypeString, Optional: true, Default: "cisco", Description: "'cisco': Alternative ABR, Cisco implementation (RFC3509); 'ibm': Alternative ABR, IBM implementation (RFC3509); 'standard': Standard behavior (RFC2328);",
			},
			"area_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"area_ipv4": {
							Type: schema.TypeString, Required: true, Description: "OSPFv3 area ID in IP address format",
						},
						"area_num": {
							Type: schema.TypeInt, Required: true, Description: "OSPFv3 area ID as a decimal value",
						},
						"default_cost": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set the summary-default cost of a NSSA or stub area (Stub's advertised default summary cost)",
						},
						"range_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeString, Optional: true, Description: "Area range for IPv6 prefix",
									},
									"option": {
										Type: schema.TypeString, Optional: true, Description: "'advertise': Advertise this range (default); 'not-advertise': DoNotAdvertise this range;",
									},
								},
							},
						},
						"stub": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure OSPFv3 area as stub",
						},
						"no_summary": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not inject inter-area routes into area",
						},
						"virtual_link_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeString, Optional: true, Description: "ID (IP addr) associated with virtual link neighbor",
									},
									"dead_interval": {
										Type: schema.TypeInt, Optional: true, Description: "Dead router detection time (Seconds)",
									},
									"bfd": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
									},
									"hello_interval": {
										Type: schema.TypeInt, Optional: true, Description: "Hello packet interval (Seconds)",
									},
									"retransmit_interval": {
										Type: schema.TypeInt, Optional: true, Description: "LSA retransmit interval (Seconds)",
									},
									"transmit_delay": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "LSA transmission delay (Seconds)",
									},
									"instance_id": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPFv3 instance ID",
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
			"auto_cost_reference_bandwidth": {
				Type: schema.TypeInt, Optional: true, Default: 10000, Description: "Use reference bandwidth method to assign OSPF cost (The reference bandwidth in terms of Mbits per second)",
			},
			"bfd_all_interfaces": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable BFD on all interfaces",
			},
			"default_information": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"originate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Distribute a default route",
						},
						"always": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always advertise default route",
						},
						"metric": {
							Type: schema.TypeInt, Optional: true, Description: "OSPF default metric (OSPF metric)",
						},
						"metric_type": {
							Type: schema.TypeInt, Optional: true, Default: 2, Description: "OSPF metric type for default routes",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"default_metric": {
				Type: schema.TypeInt, Optional: true, Default: 20, Description: "Set metric of redistributed routes (Default metric)",
			},
			"distribute_internal_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Optional: true, Description: "'lw4o6': LW4O6 Prefix; 'nat64': NAT64 Prefix; 'static-nat': Static NAT; 'floating-ip': Floating IP; 'ip-nat': IP NAT; 'ip-nat-list': IP NAT list; 'vip': Only not flagged Virtual IP (VIP); 'vip-only-flagged': Selected Virtual IP (VIP);",
						},
						"area_ipv4": {
							Type: schema.TypeString, Optional: true, Default: "255.255.255.255", Description: "OSPF area ID in IP address format",
						},
						"area_num": {
							Type: schema.TypeInt, Optional: true, Description: "OSPF area ID as a decimal value",
						},
						"cost": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Cost",
						},
					},
				},
			},
			"distribute_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeString, Optional: true, Description: "Prefix-list name",
									},
									"direction": {
										Type: schema.TypeString, Optional: true, Description: "'in': Filter incoming routing updates;",
									},
								},
							},
						},
					},
				},
			},
			"ha_standby_extra_cost": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"extra_cost": {
							Type: schema.TypeInt, Optional: true, Description: "The extra cost value",
						},
						"group": {
							Type: schema.TypeInt, Optional: true, Description: "Group (Group ID)",
						},
					},
				},
			},
			"log_adjacency_changes": {
				Type: schema.TypeString, Optional: true, Description: "'detail': Log changes in adjacency state; 'disable': Disable logging;",
			},
			"max_concurrent_dd": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Maximum number allowed to process DD concurrently (Number of DD process)",
			},
			"passive_interface": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"loopback_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"loopback": {
										Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
									},
								},
							},
						},
						"trunk_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
									},
								},
							},
						},
						"ve_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
									},
								},
							},
						},
						"tunnel_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
									},
								},
							},
						},
						"eth_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
									},
								},
							},
						},
					},
				},
			},
			"process_id": {
				Type: schema.TypeString, Required: true, Description: "OSPFv3 process tag",
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
										Type: schema.TypeString, Optional: true, Description: "'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'ip-nat-list': IP NAT list; 'nat-map': NAT MAP Prefix; 'static-nat': Static NAT; 'nat64': NAT64 Prefix; 'lw4o6': LW4O6 Prefix; 'isis': ISO IS-IS; 'rip': Routing Information Protocol (RIP); 'static': Static routes;",
									},
									"metric": {
										Type: schema.TypeInt, Optional: true, Description: "OSPFV3 default metric (OSPFV3 metric)",
									},
									"metric_type": {
										Type: schema.TypeString, Optional: true, Default: "2", Description: "'1': Set OSPFV3 External Type 1 metrics; '2': Set OSPFV3 External Type 2 metrics;",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"ospf_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ospf": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Open Shortest Path First (OSPF)",
									},
									"process_id": {
										Type: schema.TypeString, Optional: true, Description: "OSPFV3 process tag",
									},
									"metric_ospf": {
										Type: schema.TypeInt, Optional: true, Description: "OSPFV3 default metric (OSPFV3 metric)",
									},
									"metric_type_ospf": {
										Type: schema.TypeString, Optional: true, Default: "2", Description: "'1': Set OSPFV3 External Type 1 metrics; '2': Set OSPFV3 External Type 2 metrics;",
									},
									"route_map_ospf": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"ip_nat": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP-NAT",
						},
						"metric_ip_nat": {
							Type: schema.TypeInt, Optional: true, Description: "OSPFV3 default metric (OSPFV3 metric)",
						},
						"metric_type_ip_nat": {
							Type: schema.TypeString, Optional: true, Default: "2", Description: "'1': Set OSPFV3 External Type 1 metrics; '2': Set OSPFV3 External Type 2 metrics;",
						},
						"route_map_ip_nat": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
						"ip_nat_floating_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_nat_prefix": {
										Type: schema.TypeString, Optional: true, Description: "Address",
									},
									"ip_nat_floating_ip_forward": {
										Type: schema.TypeString, Optional: true, Description: "Floating-IP as forward address",
									},
								},
							},
						},
						"vip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type_vip": {
										Type: schema.TypeString, Optional: true, Description: "'only-flagged': Selected Virtual IP (VIP); 'only-not-flagged': Only not flagged;",
									},
									"metric_vip": {
										Type: schema.TypeInt, Optional: true, Description: "OSPFV3 default metric (OSPFV3 metric)",
									},
									"metric_type_vip": {
										Type: schema.TypeString, Optional: true, Default: "2", Description: "'1': Set OSPFV3 External Type 1 metrics; '2': Set OSPFV3 External Type 2 metrics;",
									},
									"route_map_vip": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"vip_floating_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vip_address": {
										Type: schema.TypeString, Optional: true, Description: "Address",
									},
									"vip_floating_ip_forward": {
										Type: schema.TypeString, Optional: true, Description: "Floating-IP as forward address",
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
			"router_id": {
				Type: schema.TypeString, Optional: true, Description: "router-id for the OSPF process (OSPFv3 router-id in IPv4 address format)",
			},
			"timers": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"spf": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"exp": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"min_delay": {
													Type: schema.TypeInt, Optional: true, Default: 500, Description: "Minimum delay between receiving a change to SPF calculation in milliseconds",
												},
												"max_delay": {
													Type: schema.TypeInt, Optional: true, Default: 50000, Description: "Maximum delay between receiving a change to SPF calculation in milliseconds",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceRouterIpv6OspfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6OspfCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6Ospf(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6OspfRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterIpv6OspfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6OspfUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6Ospf(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6OspfRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterIpv6OspfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6OspfDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6Ospf(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterIpv6OspfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6OspfRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6Ospf(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterIpv6OspfAreaList(d []interface{}) []edpt.RouterIpv6OspfAreaList {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfAreaList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfAreaList
		oi.AreaIpv4 = in["area_ipv4"].(string)
		oi.AreaNum = in["area_num"].(int)
		oi.DefaultCost = in["default_cost"].(int)
		oi.RangeList = getSliceRouterIpv6OspfAreaListRangeList(in["range_list"].([]interface{}))
		oi.Stub = in["stub"].(int)
		oi.NoSummary = in["no_summary"].(int)
		oi.VirtualLinkList = getSliceRouterIpv6OspfAreaListVirtualLinkList(in["virtual_link_list"].([]interface{}))
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6OspfAreaListRangeList(d []interface{}) []edpt.RouterIpv6OspfAreaListRangeList {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfAreaListRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfAreaListRangeList
		oi.Value = in["value"].(string)
		oi.Option = in["option"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6OspfAreaListVirtualLinkList(d []interface{}) []edpt.RouterIpv6OspfAreaListVirtualLinkList {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfAreaListVirtualLinkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfAreaListVirtualLinkList
		oi.Value = in["value"].(string)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.Bfd = in["bfd"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIpv6OspfDefaultInformation1248(d []interface{}) edpt.RouterIpv6OspfDefaultInformation1248 {

	count1 := len(d)
	var ret edpt.RouterIpv6OspfDefaultInformation1248
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Originate = in["originate"].(int)
		ret.Always = in["always"].(int)
		ret.Metric = in["metric"].(int)
		ret.MetricType = in["metric_type"].(int)
		ret.RouteMap = in["route_map"].(string)
		//omit uuid
	}
	return ret
}

func getSliceRouterIpv6OspfDistributeInternalList(d []interface{}) []edpt.RouterIpv6OspfDistributeInternalList {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfDistributeInternalList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfDistributeInternalList
		oi.Type = in["type"].(string)
		oi.AreaIpv4 = in["area_ipv4"].(string)
		oi.AreaNum = in["area_num"].(int)
		oi.Cost = in["cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIpv6OspfDistributeList(d []interface{}) edpt.RouterIpv6OspfDistributeList {

	count1 := len(d)
	var ret edpt.RouterIpv6OspfDistributeList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrefixList = getObjectRouterIpv6OspfDistributeListPrefixList(in["prefix_list"].([]interface{}))
	}
	return ret
}

func getObjectRouterIpv6OspfDistributeListPrefixList(d []interface{}) edpt.RouterIpv6OspfDistributeListPrefixList {

	count1 := len(d)
	var ret edpt.RouterIpv6OspfDistributeListPrefixList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(string)
		ret.Direction = in["direction"].(string)
	}
	return ret
}

func getSliceRouterIpv6OspfHaStandbyExtraCost(d []interface{}) []edpt.RouterIpv6OspfHaStandbyExtraCost {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfHaStandbyExtraCost, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfHaStandbyExtraCost
		oi.ExtraCost = in["extra_cost"].(int)
		oi.Group = in["group"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIpv6OspfPassiveInterface(d []interface{}) edpt.RouterIpv6OspfPassiveInterface {

	count1 := len(d)
	var ret edpt.RouterIpv6OspfPassiveInterface
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LoopbackCfg = getSliceRouterIpv6OspfPassiveInterfaceLoopbackCfg(in["loopback_cfg"].([]interface{}))
		ret.TrunkCfg = getSliceRouterIpv6OspfPassiveInterfaceTrunkCfg(in["trunk_cfg"].([]interface{}))
		ret.VeCfg = getSliceRouterIpv6OspfPassiveInterfaceVeCfg(in["ve_cfg"].([]interface{}))
		ret.TunnelCfg = getSliceRouterIpv6OspfPassiveInterfaceTunnelCfg(in["tunnel_cfg"].([]interface{}))
		ret.EthCfg = getSliceRouterIpv6OspfPassiveInterfaceEthCfg(in["eth_cfg"].([]interface{}))
	}
	return ret
}

func getSliceRouterIpv6OspfPassiveInterfaceLoopbackCfg(d []interface{}) []edpt.RouterIpv6OspfPassiveInterfaceLoopbackCfg {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfPassiveInterfaceLoopbackCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfPassiveInterfaceLoopbackCfg
		oi.Loopback = in["loopback"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6OspfPassiveInterfaceTrunkCfg(d []interface{}) []edpt.RouterIpv6OspfPassiveInterfaceTrunkCfg {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfPassiveInterfaceTrunkCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfPassiveInterfaceTrunkCfg
		oi.Trunk = in["trunk"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6OspfPassiveInterfaceVeCfg(d []interface{}) []edpt.RouterIpv6OspfPassiveInterfaceVeCfg {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfPassiveInterfaceVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfPassiveInterfaceVeCfg
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6OspfPassiveInterfaceTunnelCfg(d []interface{}) []edpt.RouterIpv6OspfPassiveInterfaceTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfPassiveInterfaceTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfPassiveInterfaceTunnelCfg
		oi.Tunnel = in["tunnel"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6OspfPassiveInterfaceEthCfg(d []interface{}) []edpt.RouterIpv6OspfPassiveInterfaceEthCfg {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfPassiveInterfaceEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfPassiveInterfaceEthCfg
		oi.Ethernet = in["ethernet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIpv6OspfRedistribute1249(d []interface{}) edpt.RouterIpv6OspfRedistribute1249 {

	count1 := len(d)
	var ret edpt.RouterIpv6OspfRedistribute1249
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RedistList = getSliceRouterIpv6OspfRedistributeRedistList1250(in["redist_list"].([]interface{}))
		ret.OspfList = getSliceRouterIpv6OspfRedistributeOspfList1251(in["ospf_list"].([]interface{}))
		ret.IpNat = in["ip_nat"].(int)
		ret.MetricIpNat = in["metric_ip_nat"].(int)
		ret.MetricTypeIpNat = in["metric_type_ip_nat"].(string)
		ret.RouteMapIpNat = in["route_map_ip_nat"].(string)
		ret.IpNatFloatingList = getSliceRouterIpv6OspfRedistributeIpNatFloatingList1252(in["ip_nat_floating_list"].([]interface{}))
		ret.VipList = getSliceRouterIpv6OspfRedistributeVipList1253(in["vip_list"].([]interface{}))
		ret.VipFloatingList = getSliceRouterIpv6OspfRedistributeVipFloatingList1254(in["vip_floating_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceRouterIpv6OspfRedistributeRedistList1250(d []interface{}) []edpt.RouterIpv6OspfRedistributeRedistList1250 {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfRedistributeRedistList1250, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfRedistributeRedistList1250
		oi.Type = in["type"].(string)
		oi.Metric = in["metric"].(int)
		oi.MetricType = in["metric_type"].(string)
		oi.RouteMap = in["route_map"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6OspfRedistributeOspfList1251(d []interface{}) []edpt.RouterIpv6OspfRedistributeOspfList1251 {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfRedistributeOspfList1251, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfRedistributeOspfList1251
		oi.Ospf = in["ospf"].(int)
		oi.ProcessId = in["process_id"].(string)
		oi.MetricOspf = in["metric_ospf"].(int)
		oi.MetricTypeOspf = in["metric_type_ospf"].(string)
		oi.RouteMapOspf = in["route_map_ospf"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6OspfRedistributeIpNatFloatingList1252(d []interface{}) []edpt.RouterIpv6OspfRedistributeIpNatFloatingList1252 {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfRedistributeIpNatFloatingList1252, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfRedistributeIpNatFloatingList1252
		oi.IpNatPrefix = in["ip_nat_prefix"].(string)
		oi.IpNatFloatingIpForward = in["ip_nat_floating_ip_forward"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6OspfRedistributeVipList1253(d []interface{}) []edpt.RouterIpv6OspfRedistributeVipList1253 {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfRedistributeVipList1253, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfRedistributeVipList1253
		oi.TypeVip = in["type_vip"].(string)
		oi.MetricVip = in["metric_vip"].(int)
		oi.MetricTypeVip = in["metric_type_vip"].(string)
		oi.RouteMapVip = in["route_map_vip"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6OspfRedistributeVipFloatingList1254(d []interface{}) []edpt.RouterIpv6OspfRedistributeVipFloatingList1254 {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfRedistributeVipFloatingList1254, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfRedistributeVipFloatingList1254
		oi.VipAddress = in["vip_address"].(string)
		oi.VipFloatingIpForward = in["vip_floating_ip_forward"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIpv6OspfTimers(d []interface{}) edpt.RouterIpv6OspfTimers {

	count1 := len(d)
	var ret edpt.RouterIpv6OspfTimers
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Spf = getObjectRouterIpv6OspfTimersSpf(in["spf"].([]interface{}))
	}
	return ret
}

func getObjectRouterIpv6OspfTimersSpf(d []interface{}) edpt.RouterIpv6OspfTimersSpf {

	count1 := len(d)
	var ret edpt.RouterIpv6OspfTimersSpf
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Exp = getObjectRouterIpv6OspfTimersSpfExp(in["exp"].([]interface{}))
	}
	return ret
}

func getObjectRouterIpv6OspfTimersSpfExp(d []interface{}) edpt.RouterIpv6OspfTimersSpfExp {

	count1 := len(d)
	var ret edpt.RouterIpv6OspfTimersSpfExp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MinDelay = in["min_delay"].(int)
		ret.MaxDelay = in["max_delay"].(int)
	}
	return ret
}

func dataToEndpointRouterIpv6Ospf(d *schema.ResourceData) edpt.RouterIpv6Ospf {
	var ret edpt.RouterIpv6Ospf
	ret.Inst.AbrTypeOption = d.Get("abr_type_option").(string)
	ret.Inst.AreaList = getSliceRouterIpv6OspfAreaList(d.Get("area_list").([]interface{}))
	ret.Inst.AutoCostReferenceBandwidth = d.Get("auto_cost_reference_bandwidth").(int)
	ret.Inst.BfdAllInterfaces = d.Get("bfd_all_interfaces").(int)
	ret.Inst.DefaultInformation = getObjectRouterIpv6OspfDefaultInformation1248(d.Get("default_information").([]interface{}))
	ret.Inst.DefaultMetric = d.Get("default_metric").(int)
	ret.Inst.DistributeInternalList = getSliceRouterIpv6OspfDistributeInternalList(d.Get("distribute_internal_list").([]interface{}))
	ret.Inst.DistributeList = getObjectRouterIpv6OspfDistributeList(d.Get("distribute_list").([]interface{}))
	ret.Inst.HaStandbyExtraCost = getSliceRouterIpv6OspfHaStandbyExtraCost(d.Get("ha_standby_extra_cost").([]interface{}))
	ret.Inst.LogAdjacencyChanges = d.Get("log_adjacency_changes").(string)
	ret.Inst.MaxConcurrentDd = d.Get("max_concurrent_dd").(int)
	ret.Inst.PassiveInterface = getObjectRouterIpv6OspfPassiveInterface(d.Get("passive_interface").([]interface{}))
	ret.Inst.ProcessId = d.Get("process_id").(string)
	ret.Inst.Redistribute = getObjectRouterIpv6OspfRedistribute1249(d.Get("redistribute").([]interface{}))
	ret.Inst.RouterId = d.Get("router_id").(string)
	ret.Inst.Timers = getObjectRouterIpv6OspfTimers(d.Get("timers").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
