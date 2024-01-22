package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouteMap() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_route_map`: Configure route-map\n\n__PLACEHOLDER__",
		CreateContext: resourceRouteMapCreate,
		UpdateContext: resourceRouteMapUpdate,
		ReadContext:   resourceRouteMapRead,
		DeleteContext: resourceRouteMapDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Required: true, Description: "'permit': Route map permits set operations; 'deny': Route map denies set operations;",
			},
			"match": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"as_path": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "AS path access-list name",
									},
								},
							},
						},
						"community": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "One or more Community Lists (numbered or named)",
												},
												"exact_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do exact matching of communities",
												},
											},
										},
									},
								},
							},
						},
						"extcommunity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"extcommunity_l_name": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "One or more Extended Community Lists (numbered or named)",
												},
												"exact_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do exact matching of ecommunities",
												},
											},
										},
									},
								},
							},
						},
						"large_community": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"l_name_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "One or more Large Community Lists (numbered or named)",
												},
												"exact_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do exact matching of large-communities",
												},
											},
										},
									},
								},
							},
						},
						"group": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group_id": {
										Type: schema.TypeInt, Optional: true, Description: "HA or VRRP-A group id",
									},
									"ha_state": {
										Type: schema.TypeString, Optional: true, Description: "'active': HA or VRRP-A in Active State; 'standby': HA or VRRP-A in Standby State;",
									},
								},
							},
						},
						"scaleout": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_id": {
										Type: schema.TypeInt, Optional: true, Description: "Scaleout Cluster-id",
									},
									"operational_state": {
										Type: schema.TypeString, Optional: true, Description: "'up': Scaleout is up and running; 'down': Scaleout is down or disabled;",
									},
								},
							},
						},
						"interface": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
									},
									"loopback": {
										Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk Interface (Trunk interface number)",
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
									},
									"tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
									},
								},
							},
						},
						"local_preference": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"val": {
										Type: schema.TypeInt, Optional: true, Description: "Preference value",
									},
								},
							},
						},
						"origin": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"egp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "remote EGP",
									},
									"igp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "local IGP",
									},
									"incomplete": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "unknown heritage",
									},
								},
							},
						},
						"ip": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"acl1": {
													Type: schema.TypeInt, Optional: true, Description: "IP access-list number",
												},
												"acl2": {
													Type: schema.TypeInt, Optional: true, Description: "IP access-list number (expanded range)",
												},
												"name": {
													Type: schema.TypeString, Optional: true, Description: "IP access-list name",
												},
												"prefix_list": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type: schema.TypeString, Optional: true, Description: "IP prefix-list name",
															},
														},
													},
												},
											},
										},
									},
									"next_hop": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"acl1": {
													Type: schema.TypeInt, Optional: true, Description: "IP access-list number",
												},
												"acl2": {
													Type: schema.TypeInt, Optional: true, Description: "IP access-list number (expanded range)",
												},
												"name": {
													Type: schema.TypeString, Optional: true, Description: "IP access-list name",
												},
												"prefix_list_1": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type: schema.TypeString, Optional: true, Description: "IP prefix-list name",
															},
														},
													},
												},
											},
										},
									},
									"peer": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"acl1": {
													Type: schema.TypeInt, Optional: true, Description: "IP access-list number",
												},
												"acl2": {
													Type: schema.TypeInt, Optional: true, Description: "IP access-list number (expanded range)",
												},
												"name": {
													Type: schema.TypeString, Optional: true, Description: "IP access-list name",
												},
											},
										},
									},
									"rib": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"exact": {
													Type: schema.TypeString, Optional: true, Description: "Exact match a prefix (Prefix IPv4 address)",
												},
												"reachable": {
													Type: schema.TypeString, Optional: true, Description: "IP address is reachable",
												},
												"unreachable": {
													Type: schema.TypeString, Optional: true, Description: "IP address is unreachable",
												},
											},
										},
									},
								},
							},
						},
						"ipv6": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address_1": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "IPv6 access-list name",
												},
												"prefix_list_2": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type: schema.TypeString, Optional: true, Description: "IPv6 prefix-list name",
															},
														},
													},
												},
											},
										},
									},
									"next_hop_1": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"next_hop_acl_name": {
													Type: schema.TypeString, Optional: true, Description: "IPv6 access-list name",
												},
												"v6_addr": {
													Type: schema.TypeString, Optional: true, Description: "IPv6 address of next hop",
												},
												"prefix_list_name": {
													Type: schema.TypeString, Optional: true, Description: "IPv6 prefix-list name",
												},
											},
										},
									},
									"peer_1": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"acl1": {
													Type: schema.TypeInt, Optional: true, Description: "IPv6 access-list number",
												},
												"acl2": {
													Type: schema.TypeInt, Optional: true, Description: "IP access-list number (expanded range)",
												},
												"name": {
													Type: schema.TypeString, Optional: true, Description: "IP access-list name",
												},
											},
										},
									},
									"rib": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"exact": {
													Type: schema.TypeString, Optional: true, Description: "Exact match a prefix (Prefix IPv6 address)",
												},
												"reachable": {
													Type: schema.TypeString, Optional: true, Description: "IPv6 address is reachable",
												},
												"unreachable": {
													Type: schema.TypeString, Optional: true, Description: "IPv6 address is unreachable",
												},
											},
										},
									},
								},
							},
						},
						"metric": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "Metric value",
									},
								},
							},
						},
						"route_type": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"external": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type: schema.TypeString, Optional: true, Description: "'type-1': Match OSPF External Type 1 metrics; 'type-2': Match OSPF External Type 2 metrics;",
												},
											},
										},
									},
								},
							},
						},
						"tag": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "Tag value",
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
			"sequence": {
				Type: schema.TypeInt, Required: true, Description: "Sequence to insert to/delete from existing route-map entry",
			},
			"set": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"next_hop": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": {
													Type: schema.TypeString, Optional: true, Description: "IP address of next hop",
												},
											},
										},
									},
								},
							},
						},
						"ddos": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list_name": {
										Type: schema.TypeString, Optional: true, Description: "Class-List Name",
									},
									"class_list_cid": {
										Type: schema.TypeInt, Optional: true, Description: "Class-List Cid",
									},
									"zone": {
										Type: schema.TypeString, Optional: true, Description: "Zone Name",
									},
								},
							},
						},
						"ipv6": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"next_hop_1": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": {
													Type: schema.TypeString, Optional: true, Description: "global address of next hop",
												},
												"local": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"address": {
																Type: schema.TypeString, Optional: true, Description: "IPv6 address of next hop",
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
						"level": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Export into a level-1 area; 'level-1-2': Export into level-1 and level-2; 'level-2': Export into level-2 sub-domain;",
									},
								},
							},
						},
						"metric": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeString, Optional: true, Description: "Metric Value (from -4294967295 to 4294967295)",
									},
								},
							},
						},
						"metric_type": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeString, Optional: true, Description: "'external': IS-IS external metric type; 'internal': IS-IS internal metric type; 'type-1': OSPF external type 1 metric; 'type-2': OSPF external type 2 metric;",
									},
								},
							},
						},
						"tag": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "Tag value",
									},
								},
							},
						},
						"aggregator": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aggregator_as": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"asn": {
													Type: schema.TypeInt, Optional: true, Description: "AS number",
												},
												"ip": {
													Type: schema.TypeString, Optional: true, Description: "IP address of aggregator",
												},
											},
										},
									},
								},
							},
						},
						"as_path": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prepend": {
										Type: schema.TypeString, Optional: true, Description: "Prepend to the as-path (AS number)",
									},
									"num": {
										Type: schema.TypeString, Optional: true, Description: "AS number",
									},
									"num2": {
										Type: schema.TypeString, Optional: true, Description: "AS number",
									},
								},
							},
						},
						"atomic_aggregate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "BGP atomic aggregate attribute",
						},
						"comm_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"v_std": {
										Type: schema.TypeInt, Optional: true, Description: "Community-list number (standard)",
									},
									"delete": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete matching communities",
									},
									"v_exp": {
										Type: schema.TypeInt, Optional: true, Description: "Community-list number (expanded)",
									},
									"v_exp_delete": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete matching communities",
									},
									"name": {
										Type: schema.TypeString, Optional: true, Description: "Community-list name",
									},
									"name_delete": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete matching communities",
									},
								},
							},
						},
						"community": {
							Type: schema.TypeString, Optional: true, Description: "BGP community attribute",
						},
						"dampening_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dampening": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable route-flap dampening",
									},
									"dampening_half_time": {
										Type: schema.TypeInt, Optional: true, Description: "Reachability Half-life time for the penalty(minutes)",
									},
									"dampening_reuse": {
										Type: schema.TypeInt, Optional: true, Description: "Value to start reusing a route",
									},
									"dampening_supress": {
										Type: schema.TypeInt, Optional: true, Description: "Value to start suppressing a route",
									},
									"dampening_max_supress": {
										Type: schema.TypeInt, Optional: true, Description: "Maximum duration to suppress a stable route(minutes)",
									},
									"dampening_penalty": {
										Type: schema.TypeInt, Optional: true, Description: "Un-reachability Half-life time for the penalty(minutes)",
									},
								},
							},
						},
						"extcommunity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rt": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type: schema.TypeString, Optional: true, Description: "VPN extended community",
												},
											},
										},
									},
									"soo": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type: schema.TypeString, Optional: true, Description: "VPN extended community",
												},
											},
										},
									},
								},
							},
						},
						"local_preference": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"val": {
										Type: schema.TypeInt, Optional: true, Description: "Preference value",
									},
								},
							},
						},
						"originator_id": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"originator_ip": {
										Type: schema.TypeString, Optional: true, Description: "IP address of originator",
									},
								},
							},
						},
						"weight": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"weight_val": {
										Type: schema.TypeInt, Optional: true, Description: "Weight value",
									},
								},
							},
						},
						"origin": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"egp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "remote EGP",
									},
									"igp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "local IGP",
									},
									"incomplete": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "unknown heritage",
									},
								},
							},
						},
						"large_comm_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"l_v_std": {
										Type: schema.TypeInt, Optional: true, Description: "Large Community-list number (standard)",
									},
									"l_v_std_delete": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete matching large communities",
									},
									"l_v_exp": {
										Type: schema.TypeInt, Optional: true, Description: "Large Community-list number (expanded)",
									},
									"l_v_exp_delete": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete matching large communities",
									},
									"l_name": {
										Type: schema.TypeString, Optional: true, Description: "Large Community-list name",
									},
									"large_name_delete": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete matching large communities",
									},
								},
							},
						},
						"large_community": {
							Type: schema.TypeString, Optional: true, Description: "BGP large community attribute",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"tag": {
				Type: schema.TypeString, Required: true, Description: "Route map tag",
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
func resourceRouteMapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouteMapCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouteMap(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouteMapRead(ctx, d, meta)
	}
	return diags
}

func resourceRouteMapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouteMapUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouteMap(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouteMapRead(ctx, d, meta)
	}
	return diags
}
func resourceRouteMapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouteMapDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouteMap(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouteMapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouteMapRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouteMap(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRouteMapMatch1097(d []interface{}) edpt.RouteMapMatch1097 {

	count1 := len(d)
	var ret edpt.RouteMapMatch1097
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AsPath = getObjectRouteMapMatchAsPath1098(in["as_path"].([]interface{}))
		ret.Community = getObjectRouteMapMatchCommunity1099(in["community"].([]interface{}))
		ret.Extcommunity = getObjectRouteMapMatchExtcommunity1101(in["extcommunity"].([]interface{}))
		ret.LargeCommunity = getObjectRouteMapMatchLargeCommunity1103(in["large_community"].([]interface{}))
		ret.Group = getObjectRouteMapMatchGroup1105(in["group"].([]interface{}))
		ret.Scaleout = getObjectRouteMapMatchScaleout1106(in["scaleout"].([]interface{}))
		ret.Interface = getObjectRouteMapMatchInterface1107(in["interface"].([]interface{}))
		ret.LocalPreference = getObjectRouteMapMatchLocalPreference1108(in["local_preference"].([]interface{}))
		ret.Origin = getObjectRouteMapMatchOrigin1109(in["origin"].([]interface{}))
		ret.Ip = getObjectRouteMapMatchIp1110(in["ip"].([]interface{}))
		ret.Ipv6 = getObjectRouteMapMatchIpv61117(in["ipv6"].([]interface{}))
		ret.Metric = getObjectRouteMapMatchMetric1123(in["metric"].([]interface{}))
		ret.RouteType = getObjectRouteMapMatchRouteType1124(in["route_type"].([]interface{}))
		ret.Tag = getObjectRouteMapMatchTag1126(in["tag"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectRouteMapMatchAsPath1098(d []interface{}) edpt.RouteMapMatchAsPath1098 {

	count1 := len(d)
	var ret edpt.RouteMapMatchAsPath1098
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
	}
	return ret
}

func getObjectRouteMapMatchCommunity1099(d []interface{}) edpt.RouteMapMatchCommunity1099 {

	count1 := len(d)
	var ret edpt.RouteMapMatchCommunity1099
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NameCfg = getObjectRouteMapMatchCommunityNameCfg1100(in["name_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchCommunityNameCfg1100(d []interface{}) edpt.RouteMapMatchCommunityNameCfg1100 {

	count1 := len(d)
	var ret edpt.RouteMapMatchCommunityNameCfg1100
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.ExactMatch = in["exact_match"].(int)
	}
	return ret
}

func getObjectRouteMapMatchExtcommunity1101(d []interface{}) edpt.RouteMapMatchExtcommunity1101 {

	count1 := len(d)
	var ret edpt.RouteMapMatchExtcommunity1101
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ExtcommunityLName = getObjectRouteMapMatchExtcommunityExtcommunityLName1102(in["extcommunity_l_name"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchExtcommunityExtcommunityLName1102(d []interface{}) edpt.RouteMapMatchExtcommunityExtcommunityLName1102 {

	count1 := len(d)
	var ret edpt.RouteMapMatchExtcommunityExtcommunityLName1102
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.ExactMatch = in["exact_match"].(int)
	}
	return ret
}

func getObjectRouteMapMatchLargeCommunity1103(d []interface{}) edpt.RouteMapMatchLargeCommunity1103 {

	count1 := len(d)
	var ret edpt.RouteMapMatchLargeCommunity1103
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LNameCfg = getObjectRouteMapMatchLargeCommunityLNameCfg1104(in["l_name_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchLargeCommunityLNameCfg1104(d []interface{}) edpt.RouteMapMatchLargeCommunityLNameCfg1104 {

	count1 := len(d)
	var ret edpt.RouteMapMatchLargeCommunityLNameCfg1104
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.ExactMatch = in["exact_match"].(int)
	}
	return ret
}

func getObjectRouteMapMatchGroup1105(d []interface{}) edpt.RouteMapMatchGroup1105 {

	count1 := len(d)
	var ret edpt.RouteMapMatchGroup1105
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GroupId = in["group_id"].(int)
		ret.HaState = in["ha_state"].(string)
	}
	return ret
}

func getObjectRouteMapMatchScaleout1106(d []interface{}) edpt.RouteMapMatchScaleout1106 {

	count1 := len(d)
	var ret edpt.RouteMapMatchScaleout1106
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClusterId = in["cluster_id"].(int)
		ret.OperationalState = in["operational_state"].(string)
	}
	return ret
}

func getObjectRouteMapMatchInterface1107(d []interface{}) edpt.RouteMapMatchInterface1107 {

	count1 := len(d)
	var ret edpt.RouteMapMatchInterface1107
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ethernet = in["ethernet"].(int)
		ret.Loopback = in["loopback"].(int)
		ret.Trunk = in["trunk"].(int)
		ret.Ve = in["ve"].(int)
		ret.Tunnel = in["tunnel"].(int)
	}
	return ret
}

func getObjectRouteMapMatchLocalPreference1108(d []interface{}) edpt.RouteMapMatchLocalPreference1108 {

	count1 := len(d)
	var ret edpt.RouteMapMatchLocalPreference1108
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Val = in["val"].(int)
	}
	return ret
}

func getObjectRouteMapMatchOrigin1109(d []interface{}) edpt.RouteMapMatchOrigin1109 {

	count1 := len(d)
	var ret edpt.RouteMapMatchOrigin1109
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Egp = in["egp"].(int)
		ret.Igp = in["igp"].(int)
		ret.Incomplete = in["incomplete"].(int)
	}
	return ret
}

func getObjectRouteMapMatchIp1110(d []interface{}) edpt.RouteMapMatchIp1110 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIp1110
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Address = getObjectRouteMapMatchIpAddress1111(in["address"].([]interface{}))
		ret.NextHop = getObjectRouteMapMatchIpNextHop1113(in["next_hop"].([]interface{}))
		ret.Peer = getObjectRouteMapMatchIpPeer1115(in["peer"].([]interface{}))
		ret.Rib = getObjectRouteMapMatchIpRib1116(in["rib"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchIpAddress1111(d []interface{}) edpt.RouteMapMatchIpAddress1111 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpAddress1111
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Acl1 = in["acl1"].(int)
		ret.Acl2 = in["acl2"].(int)
		ret.Name = in["name"].(string)
		ret.PrefixList = getObjectRouteMapMatchIpAddressPrefixList1112(in["prefix_list"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchIpAddressPrefixList1112(d []interface{}) edpt.RouteMapMatchIpAddressPrefixList1112 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpAddressPrefixList1112
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
	}
	return ret
}

func getObjectRouteMapMatchIpNextHop1113(d []interface{}) edpt.RouteMapMatchIpNextHop1113 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpNextHop1113
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Acl1 = in["acl1"].(int)
		ret.Acl2 = in["acl2"].(int)
		ret.Name = in["name"].(string)
		ret.PrefixList1 = getObjectRouteMapMatchIpNextHopPrefixList11114(in["prefix_list_1"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchIpNextHopPrefixList11114(d []interface{}) edpt.RouteMapMatchIpNextHopPrefixList11114 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpNextHopPrefixList11114
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
	}
	return ret
}

func getObjectRouteMapMatchIpPeer1115(d []interface{}) edpt.RouteMapMatchIpPeer1115 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpPeer1115
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Acl1 = in["acl1"].(int)
		ret.Acl2 = in["acl2"].(int)
		ret.Name = in["name"].(string)
	}
	return ret
}

func getObjectRouteMapMatchIpRib1116(d []interface{}) edpt.RouteMapMatchIpRib1116 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpRib1116
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Exact = in["exact"].(string)
		ret.Reachable = in["reachable"].(string)
		ret.Unreachable = in["unreachable"].(string)
	}
	return ret
}

func getObjectRouteMapMatchIpv61117(d []interface{}) edpt.RouteMapMatchIpv61117 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpv61117
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Address1 = getObjectRouteMapMatchIpv6Address11118(in["address_1"].([]interface{}))
		ret.NextHop1 = getObjectRouteMapMatchIpv6NextHop11120(in["next_hop_1"].([]interface{}))
		ret.Peer1 = getObjectRouteMapMatchIpv6Peer11121(in["peer_1"].([]interface{}))
		ret.Rib = getObjectRouteMapMatchIpv6Rib1122(in["rib"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchIpv6Address11118(d []interface{}) edpt.RouteMapMatchIpv6Address11118 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpv6Address11118
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.PrefixList2 = getObjectRouteMapMatchIpv6Address1PrefixList21119(in["prefix_list_2"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchIpv6Address1PrefixList21119(d []interface{}) edpt.RouteMapMatchIpv6Address1PrefixList21119 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpv6Address1PrefixList21119
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
	}
	return ret
}

func getObjectRouteMapMatchIpv6NextHop11120(d []interface{}) edpt.RouteMapMatchIpv6NextHop11120 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpv6NextHop11120
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NextHopAclName = in["next_hop_acl_name"].(string)
		ret.V6Addr = in["v6_addr"].(string)
		ret.PrefixListName = in["prefix_list_name"].(string)
	}
	return ret
}

func getObjectRouteMapMatchIpv6Peer11121(d []interface{}) edpt.RouteMapMatchIpv6Peer11121 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpv6Peer11121
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Acl1 = in["acl1"].(int)
		ret.Acl2 = in["acl2"].(int)
		ret.Name = in["name"].(string)
	}
	return ret
}

func getObjectRouteMapMatchIpv6Rib1122(d []interface{}) edpt.RouteMapMatchIpv6Rib1122 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpv6Rib1122
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Exact = in["exact"].(string)
		ret.Reachable = in["reachable"].(string)
		ret.Unreachable = in["unreachable"].(string)
	}
	return ret
}

func getObjectRouteMapMatchMetric1123(d []interface{}) edpt.RouteMapMatchMetric1123 {

	count1 := len(d)
	var ret edpt.RouteMapMatchMetric1123
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
	}
	return ret
}

func getObjectRouteMapMatchRouteType1124(d []interface{}) edpt.RouteMapMatchRouteType1124 {

	count1 := len(d)
	var ret edpt.RouteMapMatchRouteType1124
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.External = getObjectRouteMapMatchRouteTypeExternal1125(in["external"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchRouteTypeExternal1125(d []interface{}) edpt.RouteMapMatchRouteTypeExternal1125 {

	count1 := len(d)
	var ret edpt.RouteMapMatchRouteTypeExternal1125
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectRouteMapMatchTag1126(d []interface{}) edpt.RouteMapMatchTag1126 {

	count1 := len(d)
	var ret edpt.RouteMapMatchTag1126
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
	}
	return ret
}

func getObjectRouteMapSet1127(d []interface{}) edpt.RouteMapSet1127 {

	count1 := len(d)
	var ret edpt.RouteMapSet1127
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ip = getObjectRouteMapSetIp1128(in["ip"].([]interface{}))
		ret.Ddos = getObjectRouteMapSetDdos1130(in["ddos"].([]interface{}))
		ret.Ipv6 = getObjectRouteMapSetIpv61131(in["ipv6"].([]interface{}))
		ret.Level = getObjectRouteMapSetLevel1134(in["level"].([]interface{}))
		ret.Metric = getObjectRouteMapSetMetric1135(in["metric"].([]interface{}))
		ret.MetricType = getObjectRouteMapSetMetricType1136(in["metric_type"].([]interface{}))
		ret.Tag = getObjectRouteMapSetTag1137(in["tag"].([]interface{}))
		ret.Aggregator = getObjectRouteMapSetAggregator1138(in["aggregator"].([]interface{}))
		ret.AsPath = getObjectRouteMapSetAsPath1140(in["as_path"].([]interface{}))
		ret.AtomicAggregate = in["atomic_aggregate"].(int)
		ret.CommList = getObjectRouteMapSetCommList1141(in["comm_list"].([]interface{}))
		ret.Community = in["community"].(string)
		ret.DampeningCfg = getObjectRouteMapSetDampeningCfg1142(in["dampening_cfg"].([]interface{}))
		ret.Extcommunity = getObjectRouteMapSetExtcommunity1143(in["extcommunity"].([]interface{}))
		ret.LocalPreference = getObjectRouteMapSetLocalPreference1146(in["local_preference"].([]interface{}))
		ret.OriginatorId = getObjectRouteMapSetOriginatorId1147(in["originator_id"].([]interface{}))
		ret.Weight = getObjectRouteMapSetWeight1148(in["weight"].([]interface{}))
		ret.Origin = getObjectRouteMapSetOrigin1149(in["origin"].([]interface{}))
		ret.LargeCommList = getObjectRouteMapSetLargeCommList1150(in["large_comm_list"].([]interface{}))
		ret.LargeCommunity = in["large_community"].(string)
		//omit uuid
	}
	return ret
}

func getObjectRouteMapSetIp1128(d []interface{}) edpt.RouteMapSetIp1128 {

	count1 := len(d)
	var ret edpt.RouteMapSetIp1128
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NextHop = getObjectRouteMapSetIpNextHop1129(in["next_hop"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapSetIpNextHop1129(d []interface{}) edpt.RouteMapSetIpNextHop1129 {

	count1 := len(d)
	var ret edpt.RouteMapSetIpNextHop1129
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Address = in["address"].(string)
	}
	return ret
}

func getObjectRouteMapSetDdos1130(d []interface{}) edpt.RouteMapSetDdos1130 {

	count1 := len(d)
	var ret edpt.RouteMapSetDdos1130
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListName = in["class_list_name"].(string)
		ret.ClassListCid = in["class_list_cid"].(int)
		ret.Zone = in["zone"].(string)
	}
	return ret
}

func getObjectRouteMapSetIpv61131(d []interface{}) edpt.RouteMapSetIpv61131 {

	count1 := len(d)
	var ret edpt.RouteMapSetIpv61131
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NextHop1 = getObjectRouteMapSetIpv6NextHop11132(in["next_hop_1"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapSetIpv6NextHop11132(d []interface{}) edpt.RouteMapSetIpv6NextHop11132 {

	count1 := len(d)
	var ret edpt.RouteMapSetIpv6NextHop11132
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Address = in["address"].(string)
		ret.Local = getObjectRouteMapSetIpv6NextHop1Local1133(in["local"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapSetIpv6NextHop1Local1133(d []interface{}) edpt.RouteMapSetIpv6NextHop1Local1133 {

	count1 := len(d)
	var ret edpt.RouteMapSetIpv6NextHop1Local1133
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Address = in["address"].(string)
	}
	return ret
}

func getObjectRouteMapSetLevel1134(d []interface{}) edpt.RouteMapSetLevel1134 {

	count1 := len(d)
	var ret edpt.RouteMapSetLevel1134
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectRouteMapSetMetric1135(d []interface{}) edpt.RouteMapSetMetric1135 {

	count1 := len(d)
	var ret edpt.RouteMapSetMetric1135
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectRouteMapSetMetricType1136(d []interface{}) edpt.RouteMapSetMetricType1136 {

	count1 := len(d)
	var ret edpt.RouteMapSetMetricType1136
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectRouteMapSetTag1137(d []interface{}) edpt.RouteMapSetTag1137 {

	count1 := len(d)
	var ret edpt.RouteMapSetTag1137
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
	}
	return ret
}

func getObjectRouteMapSetAggregator1138(d []interface{}) edpt.RouteMapSetAggregator1138 {

	count1 := len(d)
	var ret edpt.RouteMapSetAggregator1138
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AggregatorAs = getObjectRouteMapSetAggregatorAggregatorAs1139(in["aggregator_as"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapSetAggregatorAggregatorAs1139(d []interface{}) edpt.RouteMapSetAggregatorAggregatorAs1139 {

	count1 := len(d)
	var ret edpt.RouteMapSetAggregatorAggregatorAs1139
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Asn = in["asn"].(int)
		ret.Ip = in["ip"].(string)
	}
	return ret
}

func getObjectRouteMapSetAsPath1140(d []interface{}) edpt.RouteMapSetAsPath1140 {

	count1 := len(d)
	var ret edpt.RouteMapSetAsPath1140
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Prepend = in["prepend"].(string)
		ret.Num = in["num"].(string)
		ret.Num2 = in["num2"].(string)
	}
	return ret
}

func getObjectRouteMapSetCommList1141(d []interface{}) edpt.RouteMapSetCommList1141 {

	count1 := len(d)
	var ret edpt.RouteMapSetCommList1141
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VStd = in["v_std"].(int)
		ret.Delete = in["delete"].(int)
		ret.VExp = in["v_exp"].(int)
		ret.VExpDelete = in["v_exp_delete"].(int)
		ret.Name = in["name"].(string)
		ret.NameDelete = in["name_delete"].(int)
	}
	return ret
}

func getObjectRouteMapSetDampeningCfg1142(d []interface{}) edpt.RouteMapSetDampeningCfg1142 {

	count1 := len(d)
	var ret edpt.RouteMapSetDampeningCfg1142
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dampening = in["dampening"].(int)
		ret.DampeningHalfTime = in["dampening_half_time"].(int)
		ret.DampeningReuse = in["dampening_reuse"].(int)
		ret.DampeningSupress = in["dampening_supress"].(int)
		ret.DampeningMaxSupress = in["dampening_max_supress"].(int)
		ret.DampeningPenalty = in["dampening_penalty"].(int)
	}
	return ret
}

func getObjectRouteMapSetExtcommunity1143(d []interface{}) edpt.RouteMapSetExtcommunity1143 {

	count1 := len(d)
	var ret edpt.RouteMapSetExtcommunity1143
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rt = getObjectRouteMapSetExtcommunityRt1144(in["rt"].([]interface{}))
		ret.Soo = getObjectRouteMapSetExtcommunitySoo1145(in["soo"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapSetExtcommunityRt1144(d []interface{}) edpt.RouteMapSetExtcommunityRt1144 {

	count1 := len(d)
	var ret edpt.RouteMapSetExtcommunityRt1144
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectRouteMapSetExtcommunitySoo1145(d []interface{}) edpt.RouteMapSetExtcommunitySoo1145 {

	count1 := len(d)
	var ret edpt.RouteMapSetExtcommunitySoo1145
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectRouteMapSetLocalPreference1146(d []interface{}) edpt.RouteMapSetLocalPreference1146 {

	count1 := len(d)
	var ret edpt.RouteMapSetLocalPreference1146
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Val = in["val"].(int)
	}
	return ret
}

func getObjectRouteMapSetOriginatorId1147(d []interface{}) edpt.RouteMapSetOriginatorId1147 {

	count1 := len(d)
	var ret edpt.RouteMapSetOriginatorId1147
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OriginatorIp = in["originator_ip"].(string)
	}
	return ret
}

func getObjectRouteMapSetWeight1148(d []interface{}) edpt.RouteMapSetWeight1148 {

	count1 := len(d)
	var ret edpt.RouteMapSetWeight1148
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.WeightVal = in["weight_val"].(int)
	}
	return ret
}

func getObjectRouteMapSetOrigin1149(d []interface{}) edpt.RouteMapSetOrigin1149 {

	count1 := len(d)
	var ret edpt.RouteMapSetOrigin1149
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Egp = in["egp"].(int)
		ret.Igp = in["igp"].(int)
		ret.Incomplete = in["incomplete"].(int)
	}
	return ret
}

func getObjectRouteMapSetLargeCommList1150(d []interface{}) edpt.RouteMapSetLargeCommList1150 {

	count1 := len(d)
	var ret edpt.RouteMapSetLargeCommList1150
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LVStd = in["l_v_std"].(int)
		ret.LVStdDelete = in["l_v_std_delete"].(int)
		ret.LVExp = in["l_v_exp"].(int)
		ret.LVExpDelete = in["l_v_exp_delete"].(int)
		ret.LName = in["l_name"].(string)
		ret.LargeNameDelete = in["large_name_delete"].(int)
	}
	return ret
}

func dataToEndpointRouteMap(d *schema.ResourceData) edpt.RouteMap {
	var ret edpt.RouteMap
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Match = getObjectRouteMapMatch1097(d.Get("match").([]interface{}))
	ret.Inst.Sequence = d.Get("sequence").(int)
	ret.Inst.Set = getObjectRouteMapSet1127(d.Get("set").([]interface{}))
	ret.Inst.Tag = d.Get("tag").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
