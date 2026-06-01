package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp`: Border Gateway Protocol (BGP)\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpCreate,
		UpdateContext: resourceRouterBgpUpdate,
		ReadContext:   resourceRouterBgpRead,
		DeleteContext: resourceRouterBgpDelete,

		Schema: map[string]*schema.Schema{
			"address_family": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bgp": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dampening": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable route-flap dampening",
												},
												"dampening_half": {
													Type: schema.TypeInt, Optional: true, Description: "Reachability Half-life time for the penalty(minutes)",
												},
												"dampening_start_reuse": {
													Type: schema.TypeInt, Optional: true, Description: "Value to start reusing a route",
												},
												"dampening_start_supress": {
													Type: schema.TypeInt, Optional: true, Description: "Value to start suppressing a route",
												},
												"dampening_max_supress": {
													Type: schema.TypeInt, Optional: true, Description: "Maximum duration to suppress a stable route(minutes)",
												},
												"dampening_unreachability": {
													Type: schema.TypeInt, Optional: true, Description: "Un-reachability Half-life time for the penalty(minutes)",
												},
												"route_map": {
													Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria for dampening (Route-map name)",
												},
											},
										},
									},
									"distance": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"distance_ext": {
													Type: schema.TypeInt, Optional: true, Description: "Distance for routes external to the AS",
												},
												"distance_int": {
													Type: schema.TypeInt, Optional: true, Description: "Distance for routes internal to the AS",
												},
												"distance_local": {
													Type: schema.TypeInt, Optional: true, Description: "Distance for local routes",
												},
											},
										},
									},
									"maximum_paths_value": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Supported BGP multipath numbers",
									},
									"originate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Distribute an IPv6 default route",
									},
									"prefer_global": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Prefer Global IPv6 Nexthop address",
									},
									"aggregate_address_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"aggregate_address": {
													Type: schema.TypeString, Optional: true, Description: "Configure BGP aggregate entries (Aggregate IPv6 prefix)",
												},
												"as_set": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Generate AS set path information",
												},
												"summary_only": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Filter more specific routes from updates",
												},
											},
										},
									},
									"auto_summary": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic network number summarization",
									},
									"synchronization": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Perform IGP synchronization",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"network": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"synchronization": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"network_synchronization": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Perform IGP synchronization",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
														},
													},
												},
												"monitor": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"default": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"network_monitor_default": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "default route monitoring",
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
												"ipv6_network_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"network_ipv6": {
																Type: schema.TypeString, Required: true, Description: "Specify a network to announce via BGP",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route-map to modify the attributes (Name of the route map)",
															},
															"backdoor": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify a BGP backdoor route",
															},
															"description": {
																Type: schema.TypeString, Optional: true, Description: "Network specific description (Up to 80 characters describing this network)",
															},
															"comm_value": {
																Type: schema.TypeString, Optional: true, Description: "community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export",
															},
															"lcomm_value": {
																Type: schema.TypeString, Optional: true, Description: "Large community value in the format XX:YY:ZZ",
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
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"peer_group_neighbor_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"peer_group": {
																Type: schema.TypeString, Required: true, Description: "Neighbor tag",
															},
															"activate": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the Address Family for this Neighbor",
															},
															"allowas_in": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept as-path with my AS present in it",
															},
															"allowas_in_count": {
																Type: schema.TypeInt, Optional: true, Default: 3, Description: "Number of occurrences of AS number",
															},
															"maximum_prefix": {
																Type: schema.TypeInt, Optional: true, Description: "Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))",
															},
															"maximum_prefix_thres": {
																Type: schema.TypeInt, Optional: true, Description: "threshold-value, 1 to 100 percent",
															},
															"next_hop_self": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable the next hop calculation for this neighbor",
															},
															"remove_private_as": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove private AS number from outbound updates",
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
															"inbound": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow inbound soft reconfiguration for this neighbor",
															},
															"weight": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set default weight for routes from this neighbor",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
														},
													},
												},
												"ipv4_neighbor_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"neighbor_ipv4": {
																Type: schema.TypeString, Required: true, Description: "Neighbor address",
															},
															"peer_group_name": {
																Type: schema.TypeString, Optional: true, Description: "Configure peer-group (peer-group name)",
															},
															"activate": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the Address Family for this Neighbor",
															},
															"allowas_in": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept as-path with my AS present in it",
															},
															"allowas_in_count": {
																Type: schema.TypeInt, Optional: true, Default: 3, Description: "Number of occurrences of AS number",
															},
															"prefix_list_direction": {
																Type: schema.TypeString, Optional: true, Description: "'both': both; 'receive': receive; 'send': send;",
															},
															"graceful_restart": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable graceful-restart helper for this neighbor",
															},
															"default_originate": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Originate default route to this neighbor",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria to originate default (route-map name)",
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
															"maximum_prefix": {
																Type: schema.TypeInt, Optional: true, Description: "Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))",
															},
															"maximum_prefix_thres": {
																Type: schema.TypeInt, Optional: true, Description: "threshold-value, 1 to 100 percent",
															},
															"restart_min": {
																Type: schema.TypeInt, Optional: true, Description: "restart value, 1 to 1440 minutes",
															},
															"next_hop_self": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable the next hop calculation for this neighbor",
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
															"remove_private_as": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove private AS number from outbound updates",
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
															"inbound": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow inbound soft reconfiguration for this neighbor",
															},
															"unsuppress_map": {
																Type: schema.TypeString, Optional: true, Description: "Route-map to selectively unsuppress suppressed routes (Name of route map)",
															},
															"weight": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set default weight for routes from this neighbor",
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
															"peer_group_name": {
																Type: schema.TypeString, Optional: true, Description: "Configure peer-group (peer-group name)",
															},
															"activate": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the Address Family for this Neighbor",
															},
															"allowas_in": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept as-path with my AS present in it",
															},
															"allowas_in_count": {
																Type: schema.TypeInt, Optional: true, Default: 3, Description: "Number of occurrences of AS number",
															},
															"prefix_list_direction": {
																Type: schema.TypeString, Optional: true, Description: "'both': both; 'receive': receive; 'send': send;",
															},
															"graceful_restart": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable graceful-restart helper for this neighbor",
															},
															"default_originate": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Originate default route to this neighbor",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria to originate default (route-map name)",
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
															"maximum_prefix": {
																Type: schema.TypeInt, Optional: true, Description: "Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))",
															},
															"maximum_prefix_thres": {
																Type: schema.TypeInt, Optional: true, Description: "threshold-value, 1 to 100 percent",
															},
															"restart_min": {
																Type: schema.TypeInt, Optional: true, Description: "restart value, 1 to 1440 minutes",
															},
															"next_hop_self": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable the next hop calculation for this neighbor",
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
															"remove_private_as": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove private AS number from outbound updates",
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
															"inbound": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow inbound soft reconfiguration for this neighbor",
															},
															"unsuppress_map": {
																Type: schema.TypeString, Optional: true, Description: "Route-map to selectively unsuppress suppressed routes (Name of route map)",
															},
															"weight": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set default weight for routes from this neighbor",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
														},
													},
												},
												"ethernet_neighbor_ipv6_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ethernet": {
																Type: schema.TypeInt, Required: true, Description: "Ethernet interface number",
															},
															"peer_group_name": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
														},
													},
												},
												"ve_neighbor_ipv6_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ve": {
																Type: schema.TypeInt, Required: true, Description: "Virtual ethernet interface number",
															},
															"peer_group_name": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
														},
													},
												},
												"trunk_neighbor_ipv6_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"trunk": {
																Type: schema.TypeInt, Required: true, Description: "Trunk interface number",
															},
															"peer_group_name": {
																Type: schema.TypeString, Optional: true, Description: "",
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
									"redistribute": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"connected_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"connected": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Connected",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
															},
														},
													},
												},
												"floating_ip_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"floating_ip": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Floating IP",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
															},
														},
													},
												},
												"nat64_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"nat64": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "NAT64 Prefix",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
															},
														},
													},
												},
												"nat_map_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"nat_map": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "NAT MAP Prefix",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
															},
														},
													},
												},
												"lw4o6_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"lw4o6": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "LW4O6 Prefix",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
															},
														},
													},
												},
												"static_nat_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static_nat": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Static NAT Prefix",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
															},
														},
													},
												},
												"ip_nat_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_nat": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP NAT",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
															},
														},
													},
												},
												"ip_nat_list_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_nat_list": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP NAT list",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
															},
														},
													},
												},
												"isis_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"isis": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "ISO IS-IS",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
															},
														},
													},
												},
												"ospf_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ospf": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Open Shortest Path First (OSPF)",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
															},
														},
													},
												},
												"rip_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"rip": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Routing Information Protocol (RIP)",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
															},
														},
													},
												},
												"static_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Static routes",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
															},
														},
													},
												},
												"public_ip_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"public_ip": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Public IPv6/IPv4 Prefixes",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
															},
														},
													},
												},
												"vip": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"only_flagged_cfg": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"only_flagged": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Selected Virtual IP (VIP)",
																		},
																		"route_map": {
																			Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
																		},
																	},
																},
															},
															"only_not_flagged_cfg": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"only_not_flagged": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only not flagged",
																		},
																		"route_map": {
																			Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
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
										},
									},
								},
							},
						},
						"ipv4_flowspec": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
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
								},
							},
						},
						"ipv6_flowspec": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
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
								},
							},
						},
					},
				},
			},
			"aggregate_address_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aggregate_address": {
							Type: schema.TypeString, Optional: true, Description: "Configure BGP aggregate entries (Aggregate prefix)",
						},
						"as_set": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Generate AS set path information",
						},
						"summary_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Filter more specific routes from updates",
						},
					},
				},
			},
			"as_number": {
				Type: schema.TypeString, Required: true, Description: "AS number",
			},
			"auto_summary": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic network number summarization",
			},
			"bgp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"always_compare_med": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow comparing MED from different neighbors",
						},
						"bestpath_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ignore": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore as-path length in selecting a route",
									},
									"compare_routerid": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Compare router-id for identical EBGP paths",
									},
									"remove_recv_med": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "To remove rcvd MED attribute",
									},
									"remove_send_med": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "To remove send MED attribute",
									},
									"missing_as_worst": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Treat missing MED as the least preferred one",
									},
								},
							},
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
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria for dampening (Route-map name)",
									},
								},
							},
						},
						"local_preference_value": {
							Type: schema.TypeInt, Optional: true, Default: 100, Description: "Configure default local preference value",
						},
						"deterministic_med": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Pick the best-MED path among paths advertised from the neighboring AS",
						},
						"enforce_first_as": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enforce the first AS for EBGP routes",
						},
						"fast_external_failover": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Immediately reset session if a link to a directly connected external peer goes down",
						},
						"log_neighbor_changes": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log neighbor up/down and reset reason",
						},
						"nexthop_trigger_count": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "BGP nexthop-tracking status (count)",
						},
						"router_id": {
							Type: schema.TypeString, Optional: true, Description: "Override current router identifier (peers will reset) (Manually configured router identifier)",
						},
						"scan_time": {
							Type: schema.TypeInt, Optional: true, Default: 60, Description: "Configure background scan interval (Scan interval (sec) [Default:60 Disable:0])",
						},
						"graceful_restart": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure BGP BGP Graceful Restart",
						},
						"bgp_restart_time": {
							Type: schema.TypeInt, Optional: true, Default: 90, Description: "BGP Peer Graceful Restart time in seconds (default 90)",
						},
						"bgp_stalepath_time": {
							Type: schema.TypeInt, Optional: true, Default: 360, Description: "BGP Graceful Restart Stalepath retention time in seconds (default 360)",
						},
					},
				},
			},
			"distance_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin_distance": {
							Type: schema.TypeInt, Optional: true, Description: "Define an administrative distance",
						},
						"src_prefix": {
							Type: schema.TypeString, Optional: true, Description: "IP source prefix",
						},
						"acl_str": {
							Type: schema.TypeString, Optional: true, Description: "Access list name",
						},
						"ext_routes_dist": {
							Type: schema.TypeInt, Optional: true, Description: "Distance for routes external to the AS",
						},
						"int_routes_dist": {
							Type: schema.TypeInt, Optional: true, Description: "Distance for routes internal to the AS",
						},
						"local_routes_dist": {
							Type: schema.TypeInt, Optional: true, Description: "Distance for local routes",
						},
					},
				},
			},
			"maximum_paths_value": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Supported BGP multipath numbers",
			},
			"neighbor": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"peer_group_neighbor_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"peer_group": {
										Type: schema.TypeString, Required: true, Description: "Neighbor tag",
									},
									"peer_group_key": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure peer-group",
									},
									"peer_group_remote_as": {
										Type: schema.TypeString, Optional: true, Description: "Specify AS number of BGP neighbor",
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
									"dynamic": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertise dynamic capability to this neighbor",
									},
									"route_refresh": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Advertise route-refresh capability to this neighbor",
									},
									"extended_nexthop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertise extended-nexthop capability to this neighbor",
									},
									"collide_established": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include Neighbor in Established State for Collision Detection",
									},
									"default_originate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Originate default route to this neighbor",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria to originate default (route-map name)",
									},
									"description": {
										Type: schema.TypeString, Optional: true, Description: "Neighbor specific description (Up to 80 characters describing this neighbor)",
									},
									"dont_capability_negotiate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not perform capability negotiation",
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
									"bfd": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
									},
									"multihop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable multihop",
									},
									"maximum_prefix": {
										Type: schema.TypeInt, Optional: true, Description: "Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))",
									},
									"maximum_prefix_thres": {
										Type: schema.TypeInt, Optional: true, Description: "threshold-value, 1 to 100 percent",
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
									"remove_private_as": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove private AS number from outbound updates",
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
									"inbound": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow inbound soft reconfiguration for this neighbor",
									},
									"shutdown": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Administratively shut down this neighbor",
									},
									"strict_capability_match": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Strict capability negotiation match",
									},
									"timers_keepalive": {
										Type: schema.TypeInt, Optional: true, Default: 30, Description: "Keepalive interval",
									},
									"timers_holdtime": {
										Type: schema.TypeInt, Optional: true, Default: 90, Description: "Holdtime",
									},
									"connect": {
										Type: schema.TypeInt, Optional: true, Description: "BGP connect timer",
									},
									"update_source_ip": {
										Type: schema.TypeString, Optional: true, Description: "IP address",
									},
									"update_source_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 address",
									},
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
									},
									"loopback": {
										Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
									},
									"lif": {
										Type: schema.TypeString, Optional: true, Description: "Logical interface (Lif interface name)",
									},
									"tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
									},
									"weight": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set default weight for routes from this neighbor",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ipv4_neighbor_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"neighbor_ipv4": {
										Type: schema.TypeString, Required: true, Description: "Neighbor address",
									},
									"nbr_remote_as": {
										Type: schema.TypeString, Optional: true, Description: "Specify AS number of BGP neighbor",
									},
									"peer_group_name": {
										Type: schema.TypeString, Optional: true, Description: "Configure peer-group (peer-group name)",
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
									"dynamic": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertise dynamic capability to this neighbor",
									},
									"prefix_list_direction": {
										Type: schema.TypeString, Optional: true, Description: "'both': both; 'receive': receive; 'send': send;",
									},
									"route_refresh": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Advertise route-refresh capability to this neighbor",
									},
									"graceful_restart": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable graceful-restart helper for this neighbor",
									},
									"collide_established": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include Neighbor in Established State for Collision Detection",
									},
									"default_originate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Originate default route to this neighbor",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria to originate default (route-map name)",
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
									"acos_application_only": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send BGP update to ACOS application",
									},
									"telemetry": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send BGP update to telemetry db",
									},
									"dont_capability_negotiate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not perform capability negotiation",
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
									"bfd": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
									},
									"multihop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable multihop",
									},
									"key_id": {
										Type: schema.TypeInt, Optional: true, Description: "Key ID",
									},
									"key_type": {
										Type: schema.TypeString, Optional: true, Description: "'md5': md5; 'meticulous-md5': meticulous-md5; 'meticulous-sha1': meticulous-sha1; 'sha1': sha1; 'simple': simple;  (Keyed MD5/Meticulous Keyed MD5/Meticulous Keyed SHA1/Keyed SHA1/Simple Password)",
									},
									"bfd_value": {
										Type: schema.TypeString, Optional: true, Description: "Key String",
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
									"maximum_prefix": {
										Type: schema.TypeInt, Optional: true, Description: "Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))",
									},
									"maximum_prefix_thres": {
										Type: schema.TypeInt, Optional: true, Description: "threshold-value, 1 to 100 percent",
									},
									"restart_min": {
										Type: schema.TypeInt, Optional: true, Description: "restart value, 1 to 1440 minutes",
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
									"remove_private_as": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove private AS number from outbound updates",
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
									"inbound": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow inbound soft reconfiguration for this neighbor",
									},
									"shutdown": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Administratively shut down this neighbor",
									},
									"strict_capability_match": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Strict capability negotiation match",
									},
									"timers_keepalive": {
										Type: schema.TypeInt, Optional: true, Default: 30, Description: "Keepalive interval",
									},
									"timers_holdtime": {
										Type: schema.TypeInt, Optional: true, Default: 90, Description: "Holdtime",
									},
									"connect": {
										Type: schema.TypeInt, Optional: true, Description: "BGP connect timer",
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
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
									},
									"loopback": {
										Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
									},
									"lif": {
										Type: schema.TypeString, Optional: true, Description: "Logical interface (Lif interface name)",
									},
									"tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
									},
									"weight": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set default weight for routes from this neighbor",
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
									"nbr_remote_as": {
										Type: schema.TypeString, Optional: true, Description: "Specify AS number of BGP neighbor",
									},
									"peer_group_name": {
										Type: schema.TypeString, Optional: true, Description: "Configure peer-group (peer-group name)",
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
									"dynamic": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertise dynamic capability to this neighbor",
									},
									"prefix_list_direction": {
										Type: schema.TypeString, Optional: true, Description: "'both': both; 'receive': receive; 'send': send;",
									},
									"route_refresh": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Advertise route-refresh capability to this neighbor",
									},
									"graceful_restart": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable graceful-restart helper for this neighbor",
									},
									"extended_nexthop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertise extended-nexthop capability to this neighbor",
									},
									"collide_established": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include Neighbor in Established State for Collision Detection",
									},
									"default_originate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Originate default route to this neighbor",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria to originate default (route-map name)",
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
									"acos_application_only": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send BGP update to ACOS application",
									},
									"telemetry": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send BGP update to telemetry db",
									},
									"dont_capability_negotiate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not perform capability negotiation",
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
									"bfd": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
									},
									"multihop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable multihop",
									},
									"key_id": {
										Type: schema.TypeInt, Optional: true, Description: "Key ID",
									},
									"key_type": {
										Type: schema.TypeString, Optional: true, Description: "'md5': md5; 'meticulous-md5': meticulous-md5; 'meticulous-sha1': meticulous-sha1; 'sha1': sha1; 'simple': simple;  (Keyed MD5/Meticulous Keyed MD5/Meticulous Keyed SHA1/Keyed SHA1/Simple Password)",
									},
									"bfd_value": {
										Type: schema.TypeString, Optional: true, Description: "Key String",
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
									"maximum_prefix": {
										Type: schema.TypeInt, Optional: true, Description: "Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))",
									},
									"maximum_prefix_thres": {
										Type: schema.TypeInt, Optional: true, Description: "threshold-value, 1 to 100 percent",
									},
									"restart_min": {
										Type: schema.TypeInt, Optional: true, Description: "restart value, 1 to 1440 minutes",
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
									"remove_private_as": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove private AS number from outbound updates",
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
									"inbound": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow inbound soft reconfiguration for this neighbor",
									},
									"shutdown": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Administratively shut down this neighbor",
									},
									"strict_capability_match": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Strict capability negotiation match",
									},
									"timers_keepalive": {
										Type: schema.TypeInt, Optional: true, Default: 30, Description: "Keepalive interval",
									},
									"timers_holdtime": {
										Type: schema.TypeInt, Optional: true, Default: 90, Description: "Holdtime",
									},
									"connect": {
										Type: schema.TypeInt, Optional: true, Description: "BGP connect timer",
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
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
									},
									"loopback": {
										Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
									},
									"lif": {
										Type: schema.TypeString, Optional: true, Description: "Logical interface (Lif interface name)",
									},
									"tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
									},
									"weight": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set default weight for routes from this neighbor",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ethernet_neighbor_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ethernet": {
										Type: schema.TypeInt, Required: true, Description: "Ethernet interface number",
									},
									"unnumbered": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
									},
									"peer_group_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ve_neighbor_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ve": {
										Type: schema.TypeInt, Required: true, Description: "Virtual ethernet interface number",
									},
									"unnumbered": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
									},
									"peer_group_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"trunk_neighbor_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trunk": {
										Type: schema.TypeInt, Required: true, Description: "Trunk interface number",
									},
									"unnumbered": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
									},
									"peer_group_name": {
										Type: schema.TypeString, Optional: true, Description: "",
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
			"network": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"synchronization": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_synchronization": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Perform IGP synchronization",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"monitor": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"network_monitor_default": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "default route monitoring",
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
						"ip_cidr_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_ipv4_cidr": {
										Type: schema.TypeString, Required: true, Description: "Specify network mask",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route-map to modify the attributes (Name of the route map)",
									},
									"backdoor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify a BGP backdoor route",
									},
									"description": {
										Type: schema.TypeString, Optional: true, Description: "Network specific description (Up to 80 characters describing this network)",
									},
									"comm_value": {
										Type: schema.TypeString, Optional: true, Description: "community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export",
									},
									"lcomm_value": {
										Type: schema.TypeString, Optional: true, Description: "Large community value in the format XX:YY:ZZ",
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
			"originate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Distribute a default route",
			},
			"redistribute": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connected_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connected": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Connected",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"floating_ip_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"floating_ip": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Floating IP",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"lw4o6_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"lw4o6": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "LW4O6 Prefix",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"static_nat_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"static_nat": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Static NAT Prefix",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"ip_nat_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_nat": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP NAT",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"ip_nat_list_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_nat_list": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP NAT list",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"isis_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"isis": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "ISO IS-IS",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"ospf_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ospf": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Open Shortest Path First (OSPF)",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"rip_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rip": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Routing Information Protocol (RIP)",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"static_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"static": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Static routes",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"nat_map_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"nat_map": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "NAT MAP Prefix",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"public_ip_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"public_ip": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Public IPv6/IPv4 Prefixes",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"vip": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"only_flagged_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"only_flagged": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Selected Virtual IP (VIP)",
												},
												"route_map": {
													Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
												},
											},
										},
									},
									"only_not_flagged_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"only_not_flagged": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only not flagged",
												},
												"route_map": {
													Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
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
				},
			},
			"synchronization": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Perform IGP synchronization",
			},
			"timers": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bgp_keepalive": {
							Type: schema.TypeInt, Optional: true, Default: 30, Description: "Keepalive interval",
						},
						"bgp_holdtime": {
							Type: schema.TypeInt, Optional: true, Default: 90, Description: "Holdtime",
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
func resourceRouterBgpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRouterBgpAddressFamily1258(d []interface{}) edpt.RouterBgpAddressFamily1258 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamily1258
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv6 = getObjectRouterBgpAddressFamilyIpv61259(in["ipv6"].([]interface{}))
		ret.Ipv4Flowspec = getObjectRouterBgpAddressFamilyIpv4Flowspec1301(in["ipv4_flowspec"].([]interface{}))
		ret.Ipv6Flowspec = getObjectRouterBgpAddressFamilyIpv6Flowspec1307(in["ipv6_flowspec"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv61259(d []interface{}) edpt.RouterBgpAddressFamilyIpv61259 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv61259
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bgp = getObjectRouterBgpAddressFamilyIpv6Bgp1260(in["bgp"].([]interface{}))
		ret.Distance = getObjectRouterBgpAddressFamilyIpv6Distance1261(in["distance"].([]interface{}))
		ret.MaximumPathsValue = in["maximum_paths_value"].(int)
		ret.Originate = in["originate"].(int)
		ret.PreferGlobal = in["prefer_global"].(int)
		ret.AggregateAddressList = getSliceRouterBgpAddressFamilyIpv6AggregateAddressList1262(in["aggregate_address_list"].([]interface{}))
		ret.AutoSummary = in["auto_summary"].(int)
		ret.Synchronization = in["synchronization"].(int)
		//omit uuid
		ret.Network = getObjectRouterBgpAddressFamilyIpv6Network1263(in["network"].([]interface{}))
		ret.Neighbor = getObjectRouterBgpAddressFamilyIpv6Neighbor1268(in["neighbor"].([]interface{}))
		ret.Redistribute = getObjectRouterBgpAddressFamilyIpv6Redistribute1284(in["redistribute"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Bgp1260(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Bgp1260 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Bgp1260
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dampening = in["dampening"].(int)
		ret.DampeningHalf = in["dampening_half"].(int)
		ret.DampeningStartReuse = in["dampening_start_reuse"].(int)
		ret.DampeningStartSupress = in["dampening_start_supress"].(int)
		ret.DampeningMaxSupress = in["dampening_max_supress"].(int)
		ret.DampeningUnreachability = in["dampening_unreachability"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Distance1261(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Distance1261 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Distance1261
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DistanceExt = in["distance_ext"].(int)
		ret.DistanceInt = in["distance_int"].(int)
		ret.DistanceLocal = in["distance_local"].(int)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6AggregateAddressList1262(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6AggregateAddressList1262 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6AggregateAddressList1262, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6AggregateAddressList1262
		oi.AggregateAddress = in["aggregate_address"].(string)
		oi.AsSet = in["as_set"].(int)
		oi.SummaryOnly = in["summary_only"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Network1263(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Network1263 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Network1263
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Synchronization = getObjectRouterBgpAddressFamilyIpv6NetworkSynchronization1264(in["synchronization"].([]interface{}))
		ret.Monitor = getObjectRouterBgpAddressFamilyIpv6NetworkMonitor1265(in["monitor"].([]interface{}))
		ret.Ipv6NetworkList = getSliceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkList1267(in["ipv6_network_list"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6NetworkSynchronization1264(d []interface{}) edpt.RouterBgpAddressFamilyIpv6NetworkSynchronization1264 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6NetworkSynchronization1264
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkSynchronization = in["network_synchronization"].(int)
		//omit uuid
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6NetworkMonitor1265(d []interface{}) edpt.RouterBgpAddressFamilyIpv6NetworkMonitor1265 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6NetworkMonitor1265
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Default = getObjectRouterBgpAddressFamilyIpv6NetworkMonitorDefault1266(in["default"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6NetworkMonitorDefault1266(d []interface{}) edpt.RouterBgpAddressFamilyIpv6NetworkMonitorDefault1266 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6NetworkMonitorDefault1266
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkMonitorDefault = in["network_monitor_default"].(int)
		//omit uuid
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkList1267(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList1267 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList1267, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList1267
		oi.NetworkIpv6 = in["network_ipv6"].(string)
		oi.RouteMap = in["route_map"].(string)
		oi.Backdoor = in["backdoor"].(int)
		oi.Description = in["description"].(string)
		oi.CommValue = in["comm_value"].(string)
		oi.LcommValue = in["lcomm_value"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Neighbor1268(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Neighbor1268 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Neighbor1268
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PeerGroupNeighborList = getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList1269(in["peer_group_neighbor_list"].([]interface{}))
		ret.Ipv4NeighborList = getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborList1271(in["ipv4_neighbor_list"].([]interface{}))
		ret.Ipv6NeighborList = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborList1276(in["ipv6_neighbor_list"].([]interface{}))
		ret.EthernetNeighborIpv6List = getSliceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List1281(in["ethernet_neighbor_ipv6_list"].([]interface{}))
		ret.VeNeighborIpv6List = getSliceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List1282(in["ve_neighbor_ipv6_list"].([]interface{}))
		ret.TrunkNeighborIpv6List = getSliceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List1283(in["trunk_neighbor_ipv6_list"].([]interface{}))
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList1269(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList1269 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList1269, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList1269
		oi.PeerGroup = in["peer_group"].(string)
		oi.Activate = in["activate"].(int)
		oi.AllowasIn = in["allowas_in"].(int)
		oi.AllowasInCount = in["allowas_in_count"].(int)
		oi.MaximumPrefix = in["maximum_prefix"].(int)
		oi.MaximumPrefixThres = in["maximum_prefix_thres"].(int)
		oi.NextHopSelf = in["next_hop_self"].(int)
		oi.RemovePrivateAs = in["remove_private_as"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists1270(in["neighbor_route_map_lists"].([]interface{}))
		oi.Inbound = in["inbound"].(int)
		oi.Weight = in["weight"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists1270(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists1270 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists1270, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists1270
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborList1271(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList1271 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList1271, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList1271
		oi.NeighborIpv4 = in["neighbor_ipv4"].(string)
		oi.PeerGroupName = in["peer_group_name"].(string)
		oi.Activate = in["activate"].(int)
		oi.AllowasIn = in["allowas_in"].(int)
		oi.AllowasInCount = in["allowas_in_count"].(int)
		oi.PrefixListDirection = in["prefix_list_direction"].(string)
		oi.GracefulRestart = in["graceful_restart"].(int)
		oi.DefaultOriginate = in["default_originate"].(int)
		oi.RouteMap = in["route_map"].(string)
		oi.DistributeLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists1272(in["distribute_lists"].([]interface{}))
		oi.NeighborFilterLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists1273(in["neighbor_filter_lists"].([]interface{}))
		oi.MaximumPrefix = in["maximum_prefix"].(int)
		oi.MaximumPrefixThres = in["maximum_prefix_thres"].(int)
		oi.RestartMin = in["restart_min"].(int)
		oi.NextHopSelf = in["next_hop_self"].(int)
		oi.NeighborPrefixLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists1274(in["neighbor_prefix_lists"].([]interface{}))
		oi.RemovePrivateAs = in["remove_private_as"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists1275(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		oi.Inbound = in["inbound"].(int)
		oi.UnsuppressMap = in["unsuppress_map"].(string)
		oi.Weight = in["weight"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists1272(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists1272 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists1272, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists1272
		oi.DistributeList = in["distribute_list"].(string)
		oi.DistributeListDirection = in["distribute_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists1273(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists1273 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists1273, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists1273
		oi.FilterList = in["filter_list"].(string)
		oi.FilterListDirection = in["filter_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists1274(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists1274 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists1274, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists1274
		oi.NbrPrefixList = in["nbr_prefix_list"].(string)
		oi.NbrPrefixListDirection = in["nbr_prefix_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists1275(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists1275 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists1275, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists1275
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborList1276(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList1276 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList1276, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList1276
		oi.NeighborIpv6 = in["neighbor_ipv6"].(string)
		oi.PeerGroupName = in["peer_group_name"].(string)
		oi.Activate = in["activate"].(int)
		oi.AllowasIn = in["allowas_in"].(int)
		oi.AllowasInCount = in["allowas_in_count"].(int)
		oi.PrefixListDirection = in["prefix_list_direction"].(string)
		oi.GracefulRestart = in["graceful_restart"].(int)
		oi.DefaultOriginate = in["default_originate"].(int)
		oi.RouteMap = in["route_map"].(string)
		oi.DistributeLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists1277(in["distribute_lists"].([]interface{}))
		oi.NeighborFilterLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists1278(in["neighbor_filter_lists"].([]interface{}))
		oi.MaximumPrefix = in["maximum_prefix"].(int)
		oi.MaximumPrefixThres = in["maximum_prefix_thres"].(int)
		oi.RestartMin = in["restart_min"].(int)
		oi.NextHopSelf = in["next_hop_self"].(int)
		oi.NeighborPrefixLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists1279(in["neighbor_prefix_lists"].([]interface{}))
		oi.RemovePrivateAs = in["remove_private_as"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists1280(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		oi.Inbound = in["inbound"].(int)
		oi.UnsuppressMap = in["unsuppress_map"].(string)
		oi.Weight = in["weight"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists1277(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists1277 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists1277, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists1277
		oi.DistributeList = in["distribute_list"].(string)
		oi.DistributeListDirection = in["distribute_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists1278(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists1278 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists1278, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists1278
		oi.FilterList = in["filter_list"].(string)
		oi.FilterListDirection = in["filter_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists1279(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists1279 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists1279, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists1279
		oi.NbrPrefixList = in["nbr_prefix_list"].(string)
		oi.NbrPrefixListDirection = in["nbr_prefix_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists1280(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists1280 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists1280, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists1280
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List1281(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List1281 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List1281, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List1281
		oi.Ethernet = in["ethernet"].(int)
		oi.PeerGroupName = in["peer_group_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List1282(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List1282 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List1282, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List1282
		oi.Ve = in["ve"].(int)
		oi.PeerGroupName = in["peer_group_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List1283(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List1283 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List1283, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List1283
		oi.Trunk = in["trunk"].(int)
		oi.PeerGroupName = in["peer_group_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Redistribute1284(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Redistribute1284 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Redistribute1284
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConnectedCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeConnectedCfg1285(in["connected_cfg"].([]interface{}))
		ret.FloatingIpCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1286(in["floating_ip_cfg"].([]interface{}))
		ret.Nat64Cfg = getObjectRouterBgpAddressFamilyIpv6RedistributeNat64Cfg1287(in["nat64_cfg"].([]interface{}))
		ret.NatMapCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeNatMapCfg1288(in["nat_map_cfg"].([]interface{}))
		ret.Lw4o6Cfg = getObjectRouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1289(in["lw4o6_cfg"].([]interface{}))
		ret.StaticNatCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1290(in["static_nat_cfg"].([]interface{}))
		ret.IpNatCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatCfg1291(in["ip_nat_cfg"].([]interface{}))
		ret.IpNatListCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1292(in["ip_nat_list_cfg"].([]interface{}))
		ret.IsisCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeIsisCfg1293(in["isis_cfg"].([]interface{}))
		ret.OspfCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeOspfCfg1294(in["ospf_cfg"].([]interface{}))
		ret.RipCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeRipCfg1295(in["rip_cfg"].([]interface{}))
		ret.StaticCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeStaticCfg1296(in["static_cfg"].([]interface{}))
		ret.PublicIpCfg = getObjectRouterBgpAddressFamilyIpv6RedistributePublicIpCfg1297(in["public_ip_cfg"].([]interface{}))
		ret.Vip = getObjectRouterBgpAddressFamilyIpv6RedistributeVip1298(in["vip"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeConnectedCfg1285(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeConnectedCfg1285 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeConnectedCfg1285
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Connected = in["connected"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1286(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1286 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1286
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FloatingIp = in["floating_ip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeNat64Cfg1287(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeNat64Cfg1287 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeNat64Cfg1287
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Nat64 = in["nat64"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeNatMapCfg1288(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeNatMapCfg1288 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeNatMapCfg1288
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatMap = in["nat_map"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1289(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1289 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1289
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Lw4o6 = in["lw4o6"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1290(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1290 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1290
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticNat = in["static_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatCfg1291(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatCfg1291 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatCfg1291
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpNat = in["ip_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1292(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1292 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1292
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpNatList = in["ip_nat_list"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeIsisCfg1293(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeIsisCfg1293 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeIsisCfg1293
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isis = in["isis"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeOspfCfg1294(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeOspfCfg1294 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeOspfCfg1294
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ospf = in["ospf"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeRipCfg1295(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeRipCfg1295 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeRipCfg1295
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeStaticCfg1296(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeStaticCfg1296 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeStaticCfg1296
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Static = in["static"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributePublicIpCfg1297(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributePublicIpCfg1297 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributePublicIpCfg1297
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PublicIp = in["public_ip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeVip1298(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeVip1298 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeVip1298
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyFlaggedCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1299(in["only_flagged_cfg"].([]interface{}))
		ret.OnlyNotFlaggedCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1300(in["only_not_flagged_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1299(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1299 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1299
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyFlagged = in["only_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1300(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1300 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1300
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyNotFlagged = in["only_not_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv4Flowspec1301(d []interface{}) edpt.RouterBgpAddressFamilyIpv4Flowspec1301 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv4Flowspec1301
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.Neighbor = getObjectRouterBgpAddressFamilyIpv4FlowspecNeighbor1302(in["neighbor"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv4FlowspecNeighbor1302(d []interface{}) edpt.RouterBgpAddressFamilyIpv4FlowspecNeighbor1302 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv4FlowspecNeighbor1302
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4NeighborList = getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList1303(in["ipv4_neighbor_list"].([]interface{}))
		ret.Ipv6NeighborList = getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList1305(in["ipv6_neighbor_list"].([]interface{}))
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList1303(d []interface{}) []edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList1303 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList1303, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList1303
		oi.NeighborIpv4 = in["neighbor_ipv4"].(string)
		oi.Activate = in["activate"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1304(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1304(d []interface{}) []edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1304 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1304, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1304
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList1305(d []interface{}) []edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList1305 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList1305, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList1305
		oi.NeighborIpv6 = in["neighbor_ipv6"].(string)
		oi.Activate = in["activate"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1306(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1306(d []interface{}) []edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1306 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1306, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1306
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Flowspec1307(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Flowspec1307 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Flowspec1307
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.Neighbor = getObjectRouterBgpAddressFamilyIpv6FlowspecNeighbor1308(in["neighbor"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6FlowspecNeighbor1308(d []interface{}) edpt.RouterBgpAddressFamilyIpv6FlowspecNeighbor1308 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6FlowspecNeighbor1308
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4NeighborList = getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList1309(in["ipv4_neighbor_list"].([]interface{}))
		ret.Ipv6NeighborList = getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList1311(in["ipv6_neighbor_list"].([]interface{}))
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList1309(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList1309 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList1309, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList1309
		oi.NeighborIpv4 = in["neighbor_ipv4"].(string)
		oi.Activate = in["activate"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1310(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1310(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1310 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1310, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1310
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList1311(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList1311 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList1311, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList1311
		oi.NeighborIpv6 = in["neighbor_ipv6"].(string)
		oi.Activate = in["activate"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1312(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1312(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1312 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1312, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1312
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAggregateAddressList(d []interface{}) []edpt.RouterBgpAggregateAddressList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAggregateAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAggregateAddressList
		oi.AggregateAddress = in["aggregate_address"].(string)
		oi.AsSet = in["as_set"].(int)
		oi.SummaryOnly = in["summary_only"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpBgp(d []interface{}) edpt.RouterBgpBgp {

	count1 := len(d)
	var ret edpt.RouterBgpBgp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AlwaysCompareMed = in["always_compare_med"].(int)
		ret.BestpathCfg = getObjectRouterBgpBgpBestpathCfg(in["bestpath_cfg"].([]interface{}))
		ret.DampeningCfg = getObjectRouterBgpBgpDampeningCfg(in["dampening_cfg"].([]interface{}))
		ret.LocalPreferenceValue = in["local_preference_value"].(int)
		ret.DeterministicMed = in["deterministic_med"].(int)
		ret.EnforceFirstAs = in["enforce_first_as"].(int)
		ret.FastExternalFailover = in["fast_external_failover"].(int)
		ret.LogNeighborChanges = in["log_neighbor_changes"].(int)
		ret.NexthopTriggerCount = in["nexthop_trigger_count"].(int)
		ret.RouterId = in["router_id"].(string)
		ret.ScanTime = in["scan_time"].(int)
		ret.GracefulRestart = in["graceful_restart"].(int)
		ret.BgpRestartTime = in["bgp_restart_time"].(int)
		ret.BgpStalepathTime = in["bgp_stalepath_time"].(int)
	}
	return ret
}

func getObjectRouterBgpBgpBestpathCfg(d []interface{}) edpt.RouterBgpBgpBestpathCfg {

	count1 := len(d)
	var ret edpt.RouterBgpBgpBestpathCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ignore = in["ignore"].(int)
		ret.CompareRouterid = in["compare_routerid"].(int)
		ret.RemoveRecvMed = in["remove_recv_med"].(int)
		ret.RemoveSendMed = in["remove_send_med"].(int)
		ret.MissingAsWorst = in["missing_as_worst"].(int)
	}
	return ret
}

func getObjectRouterBgpBgpDampeningCfg(d []interface{}) edpt.RouterBgpBgpDampeningCfg {

	count1 := len(d)
	var ret edpt.RouterBgpBgpDampeningCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dampening = in["dampening"].(int)
		ret.DampeningHalfTime = in["dampening_half_time"].(int)
		ret.DampeningReuse = in["dampening_reuse"].(int)
		ret.DampeningSupress = in["dampening_supress"].(int)
		ret.DampeningMaxSupress = in["dampening_max_supress"].(int)
		ret.DampeningPenalty = in["dampening_penalty"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getSliceRouterBgpDistanceList(d []interface{}) []edpt.RouterBgpDistanceList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpDistanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpDistanceList
		oi.AdminDistance = in["admin_distance"].(int)
		oi.SrcPrefix = in["src_prefix"].(string)
		oi.AclStr = in["acl_str"].(string)
		oi.ExtRoutesDist = in["ext_routes_dist"].(int)
		oi.IntRoutesDist = in["int_routes_dist"].(int)
		oi.LocalRoutesDist = in["local_routes_dist"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpNeighbor1313(d []interface{}) edpt.RouterBgpNeighbor1313 {

	count1 := len(d)
	var ret edpt.RouterBgpNeighbor1313
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PeerGroupNeighborList = getSliceRouterBgpNeighborPeerGroupNeighborList(in["peer_group_neighbor_list"].([]interface{}))
		ret.Ipv4NeighborList = getSliceRouterBgpNeighborIpv4NeighborList(in["ipv4_neighbor_list"].([]interface{}))
		ret.Ipv6NeighborList = getSliceRouterBgpNeighborIpv6NeighborList(in["ipv6_neighbor_list"].([]interface{}))
		ret.EthernetNeighborList = getSliceRouterBgpNeighborEthernetNeighborList(in["ethernet_neighbor_list"].([]interface{}))
		ret.VeNeighborList = getSliceRouterBgpNeighborVeNeighborList(in["ve_neighbor_list"].([]interface{}))
		ret.TrunkNeighborList = getSliceRouterBgpNeighborTrunkNeighborList(in["trunk_neighbor_list"].([]interface{}))
	}
	return ret
}

func getSliceRouterBgpNeighborPeerGroupNeighborList(d []interface{}) []edpt.RouterBgpNeighborPeerGroupNeighborList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborPeerGroupNeighborList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborPeerGroupNeighborList
		oi.PeerGroup = in["peer_group"].(string)
		oi.PeerGroupKey = in["peer_group_key"].(int)
		oi.PeerGroupRemoteAs = in["peer_group_remote_as"].(string)
		oi.Activate = in["activate"].(int)
		oi.AdvertisementInterval = in["advertisement_interval"].(int)
		oi.AllowasIn = in["allowas_in"].(int)
		oi.AllowasInCount = in["allowas_in_count"].(int)
		oi.AsOriginationInterval = in["as_origination_interval"].(int)
		oi.Dynamic = in["dynamic"].(int)
		oi.RouteRefresh = in["route_refresh"].(int)
		oi.ExtendedNexthop = in["extended_nexthop"].(int)
		oi.CollideEstablished = in["collide_established"].(int)
		oi.DefaultOriginate = in["default_originate"].(int)
		oi.RouteMap = in["route_map"].(string)
		oi.Description = in["description"].(string)
		oi.DontCapabilityNegotiate = in["dont_capability_negotiate"].(int)
		oi.EbgpMultihop = in["ebgp_multihop"].(int)
		oi.EbgpMultihopHopCount = in["ebgp_multihop_hop_count"].(int)
		oi.EnforceMultihop = in["enforce_multihop"].(int)
		oi.Bfd = in["bfd"].(int)
		oi.Multihop = in["multihop"].(int)
		oi.MaximumPrefix = in["maximum_prefix"].(int)
		oi.MaximumPrefixThres = in["maximum_prefix_thres"].(int)
		oi.OverrideCapability = in["override_capability"].(int)
		oi.PassValue = in["pass_value"].(string)
		//omit pass_encrypted
		oi.Passive = in["passive"].(int)
		oi.RemovePrivateAs = in["remove_private_as"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpNeighborPeerGroupNeighborListNeighborRouteMapLists(in["neighbor_route_map_lists"].([]interface{}))
		oi.Inbound = in["inbound"].(int)
		oi.Shutdown = in["shutdown"].(int)
		oi.StrictCapabilityMatch = in["strict_capability_match"].(int)
		oi.TimersKeepalive = in["timers_keepalive"].(int)
		oi.TimersHoldtime = in["timers_holdtime"].(int)
		oi.Connect = in["connect"].(int)
		oi.UpdateSourceIp = in["update_source_ip"].(string)
		oi.UpdateSourceIpv6 = in["update_source_ipv6"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Loopback = in["loopback"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Lif = in["lif"].(string)
		oi.Tunnel = in["tunnel"].(int)
		oi.Weight = in["weight"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborPeerGroupNeighborListNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpNeighborPeerGroupNeighborListNeighborRouteMapLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborPeerGroupNeighborListNeighborRouteMapLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborPeerGroupNeighborListNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv4NeighborList(d []interface{}) []edpt.RouterBgpNeighborIpv4NeighborList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv4NeighborList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv4NeighborList
		oi.NeighborIpv4 = in["neighbor_ipv4"].(string)
		oi.NbrRemoteAs = in["nbr_remote_as"].(string)
		oi.PeerGroupName = in["peer_group_name"].(string)
		oi.Activate = in["activate"].(int)
		oi.AdvertisementInterval = in["advertisement_interval"].(int)
		oi.AllowasIn = in["allowas_in"].(int)
		oi.AllowasInCount = in["allowas_in_count"].(int)
		oi.AsOriginationInterval = in["as_origination_interval"].(int)
		oi.Dynamic = in["dynamic"].(int)
		oi.PrefixListDirection = in["prefix_list_direction"].(string)
		oi.RouteRefresh = in["route_refresh"].(int)
		oi.GracefulRestart = in["graceful_restart"].(int)
		oi.CollideEstablished = in["collide_established"].(int)
		oi.DefaultOriginate = in["default_originate"].(int)
		oi.RouteMap = in["route_map"].(string)
		oi.Description = in["description"].(string)
		oi.DisallowInfiniteHoldtime = in["disallow_infinite_holdtime"].(int)
		oi.DistributeLists = getSliceRouterBgpNeighborIpv4NeighborListDistributeLists(in["distribute_lists"].([]interface{}))
		oi.AcosApplicationOnly = in["acos_application_only"].(int)
		oi.Telemetry = in["telemetry"].(int)
		oi.DontCapabilityNegotiate = in["dont_capability_negotiate"].(int)
		oi.EbgpMultihop = in["ebgp_multihop"].(int)
		oi.EbgpMultihopHopCount = in["ebgp_multihop_hop_count"].(int)
		oi.EnforceMultihop = in["enforce_multihop"].(int)
		oi.Bfd = in["bfd"].(int)
		oi.Multihop = in["multihop"].(int)
		oi.KeyId = in["key_id"].(int)
		oi.KeyType = in["key_type"].(string)
		oi.BfdValue = in["bfd_value"].(string)
		//omit bfd_encrypted
		oi.NeighborFilterLists = getSliceRouterBgpNeighborIpv4NeighborListNeighborFilterLists(in["neighbor_filter_lists"].([]interface{}))
		oi.MaximumPrefix = in["maximum_prefix"].(int)
		oi.MaximumPrefixThres = in["maximum_prefix_thres"].(int)
		oi.RestartMin = in["restart_min"].(int)
		oi.NextHopSelf = in["next_hop_self"].(int)
		oi.OverrideCapability = in["override_capability"].(int)
		oi.PassValue = in["pass_value"].(string)
		//omit pass_encrypted
		oi.Passive = in["passive"].(int)
		oi.NeighborPrefixLists = getSliceRouterBgpNeighborIpv4NeighborListNeighborPrefixLists(in["neighbor_prefix_lists"].([]interface{}))
		oi.RemovePrivateAs = in["remove_private_as"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpNeighborIpv4NeighborListNeighborRouteMapLists(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		oi.Inbound = in["inbound"].(int)
		oi.Shutdown = in["shutdown"].(int)
		oi.StrictCapabilityMatch = in["strict_capability_match"].(int)
		oi.TimersKeepalive = in["timers_keepalive"].(int)
		oi.TimersHoldtime = in["timers_holdtime"].(int)
		oi.Connect = in["connect"].(int)
		oi.UnsuppressMap = in["unsuppress_map"].(string)
		oi.UpdateSourceIp = in["update_source_ip"].(string)
		oi.UpdateSourceIpv6 = in["update_source_ipv6"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Loopback = in["loopback"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Lif = in["lif"].(string)
		oi.Tunnel = in["tunnel"].(int)
		oi.Weight = in["weight"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv4NeighborListDistributeLists(d []interface{}) []edpt.RouterBgpNeighborIpv4NeighborListDistributeLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv4NeighborListDistributeLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv4NeighborListDistributeLists
		oi.DistributeList = in["distribute_list"].(string)
		oi.DistributeListDirection = in["distribute_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv4NeighborListNeighborFilterLists(d []interface{}) []edpt.RouterBgpNeighborIpv4NeighborListNeighborFilterLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv4NeighborListNeighborFilterLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv4NeighborListNeighborFilterLists
		oi.FilterList = in["filter_list"].(string)
		oi.FilterListDirection = in["filter_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv4NeighborListNeighborPrefixLists(d []interface{}) []edpt.RouterBgpNeighborIpv4NeighborListNeighborPrefixLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv4NeighborListNeighborPrefixLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv4NeighborListNeighborPrefixLists
		oi.NbrPrefixList = in["nbr_prefix_list"].(string)
		oi.NbrPrefixListDirection = in["nbr_prefix_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv4NeighborListNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpNeighborIpv4NeighborListNeighborRouteMapLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv4NeighborListNeighborRouteMapLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv4NeighborListNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv6NeighborList(d []interface{}) []edpt.RouterBgpNeighborIpv6NeighborList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv6NeighborList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv6NeighborList
		oi.NeighborIpv6 = in["neighbor_ipv6"].(string)
		oi.NbrRemoteAs = in["nbr_remote_as"].(string)
		oi.PeerGroupName = in["peer_group_name"].(string)
		oi.Activate = in["activate"].(int)
		oi.AdvertisementInterval = in["advertisement_interval"].(int)
		oi.AllowasIn = in["allowas_in"].(int)
		oi.AllowasInCount = in["allowas_in_count"].(int)
		oi.AsOriginationInterval = in["as_origination_interval"].(int)
		oi.Dynamic = in["dynamic"].(int)
		oi.PrefixListDirection = in["prefix_list_direction"].(string)
		oi.RouteRefresh = in["route_refresh"].(int)
		oi.GracefulRestart = in["graceful_restart"].(int)
		oi.ExtendedNexthop = in["extended_nexthop"].(int)
		oi.CollideEstablished = in["collide_established"].(int)
		oi.DefaultOriginate = in["default_originate"].(int)
		oi.RouteMap = in["route_map"].(string)
		oi.Description = in["description"].(string)
		oi.DisallowInfiniteHoldtime = in["disallow_infinite_holdtime"].(int)
		oi.DistributeLists = getSliceRouterBgpNeighborIpv6NeighborListDistributeLists(in["distribute_lists"].([]interface{}))
		oi.AcosApplicationOnly = in["acos_application_only"].(int)
		oi.Telemetry = in["telemetry"].(int)
		oi.DontCapabilityNegotiate = in["dont_capability_negotiate"].(int)
		oi.EbgpMultihop = in["ebgp_multihop"].(int)
		oi.EbgpMultihopHopCount = in["ebgp_multihop_hop_count"].(int)
		oi.EnforceMultihop = in["enforce_multihop"].(int)
		oi.Bfd = in["bfd"].(int)
		oi.Multihop = in["multihop"].(int)
		oi.KeyId = in["key_id"].(int)
		oi.KeyType = in["key_type"].(string)
		oi.BfdValue = in["bfd_value"].(string)
		//omit bfd_encrypted
		oi.NeighborFilterLists = getSliceRouterBgpNeighborIpv6NeighborListNeighborFilterLists(in["neighbor_filter_lists"].([]interface{}))
		oi.MaximumPrefix = in["maximum_prefix"].(int)
		oi.MaximumPrefixThres = in["maximum_prefix_thres"].(int)
		oi.RestartMin = in["restart_min"].(int)
		oi.NextHopSelf = in["next_hop_self"].(int)
		oi.OverrideCapability = in["override_capability"].(int)
		oi.PassValue = in["pass_value"].(string)
		//omit pass_encrypted
		oi.Passive = in["passive"].(int)
		oi.NeighborPrefixLists = getSliceRouterBgpNeighborIpv6NeighborListNeighborPrefixLists(in["neighbor_prefix_lists"].([]interface{}))
		oi.RemovePrivateAs = in["remove_private_as"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpNeighborIpv6NeighborListNeighborRouteMapLists(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		oi.Inbound = in["inbound"].(int)
		oi.Shutdown = in["shutdown"].(int)
		oi.StrictCapabilityMatch = in["strict_capability_match"].(int)
		oi.TimersKeepalive = in["timers_keepalive"].(int)
		oi.TimersHoldtime = in["timers_holdtime"].(int)
		oi.Connect = in["connect"].(int)
		oi.UnsuppressMap = in["unsuppress_map"].(string)
		oi.UpdateSourceIp = in["update_source_ip"].(string)
		oi.UpdateSourceIpv6 = in["update_source_ipv6"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Loopback = in["loopback"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Lif = in["lif"].(string)
		oi.Tunnel = in["tunnel"].(int)
		oi.Weight = in["weight"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv6NeighborListDistributeLists(d []interface{}) []edpt.RouterBgpNeighborIpv6NeighborListDistributeLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv6NeighborListDistributeLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv6NeighborListDistributeLists
		oi.DistributeList = in["distribute_list"].(string)
		oi.DistributeListDirection = in["distribute_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv6NeighborListNeighborFilterLists(d []interface{}) []edpt.RouterBgpNeighborIpv6NeighborListNeighborFilterLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv6NeighborListNeighborFilterLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv6NeighborListNeighborFilterLists
		oi.FilterList = in["filter_list"].(string)
		oi.FilterListDirection = in["filter_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv6NeighborListNeighborPrefixLists(d []interface{}) []edpt.RouterBgpNeighborIpv6NeighborListNeighborPrefixLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv6NeighborListNeighborPrefixLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv6NeighborListNeighborPrefixLists
		oi.NbrPrefixList = in["nbr_prefix_list"].(string)
		oi.NbrPrefixListDirection = in["nbr_prefix_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv6NeighborListNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpNeighborIpv6NeighborListNeighborRouteMapLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv6NeighborListNeighborRouteMapLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv6NeighborListNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborEthernetNeighborList(d []interface{}) []edpt.RouterBgpNeighborEthernetNeighborList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborEthernetNeighborList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborEthernetNeighborList
		oi.Ethernet = in["ethernet"].(int)
		oi.Unnumbered = in["unnumbered"].(int)
		oi.PeerGroupName = in["peer_group_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborVeNeighborList(d []interface{}) []edpt.RouterBgpNeighborVeNeighborList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborVeNeighborList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborVeNeighborList
		oi.Ve = in["ve"].(int)
		oi.Unnumbered = in["unnumbered"].(int)
		oi.PeerGroupName = in["peer_group_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborTrunkNeighborList(d []interface{}) []edpt.RouterBgpNeighborTrunkNeighborList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNeighborTrunkNeighborList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborTrunkNeighborList
		oi.Trunk = in["trunk"].(int)
		oi.Unnumbered = in["unnumbered"].(int)
		oi.PeerGroupName = in["peer_group_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpNetwork1314(d []interface{}) edpt.RouterBgpNetwork1314 {

	count1 := len(d)
	var ret edpt.RouterBgpNetwork1314
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Synchronization = getObjectRouterBgpNetworkSynchronization1315(in["synchronization"].([]interface{}))
		ret.Monitor = getObjectRouterBgpNetworkMonitor1316(in["monitor"].([]interface{}))
		ret.IpCidrList = getSliceRouterBgpNetworkIpCidrList(in["ip_cidr_list"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpNetworkSynchronization1315(d []interface{}) edpt.RouterBgpNetworkSynchronization1315 {

	count1 := len(d)
	var ret edpt.RouterBgpNetworkSynchronization1315
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkSynchronization = in["network_synchronization"].(int)
		//omit uuid
	}
	return ret
}

func getObjectRouterBgpNetworkMonitor1316(d []interface{}) edpt.RouterBgpNetworkMonitor1316 {

	count1 := len(d)
	var ret edpt.RouterBgpNetworkMonitor1316
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Default = getObjectRouterBgpNetworkMonitorDefault1317(in["default"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpNetworkMonitorDefault1317(d []interface{}) edpt.RouterBgpNetworkMonitorDefault1317 {

	count1 := len(d)
	var ret edpt.RouterBgpNetworkMonitorDefault1317
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkMonitorDefault = in["network_monitor_default"].(int)
		//omit uuid
	}
	return ret
}

func getSliceRouterBgpNetworkIpCidrList(d []interface{}) []edpt.RouterBgpNetworkIpCidrList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpNetworkIpCidrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNetworkIpCidrList
		oi.NetworkIpv4Cidr = in["network_ipv4_cidr"].(string)
		oi.RouteMap = in["route_map"].(string)
		oi.Backdoor = in["backdoor"].(int)
		oi.Description = in["description"].(string)
		oi.CommValue = in["comm_value"].(string)
		oi.LcommValue = in["lcomm_value"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpRedistribute1318(d []interface{}) edpt.RouterBgpRedistribute1318 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistribute1318
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConnectedCfg = getObjectRouterBgpRedistributeConnectedCfg1319(in["connected_cfg"].([]interface{}))
		ret.FloatingIpCfg = getObjectRouterBgpRedistributeFloatingIpCfg1320(in["floating_ip_cfg"].([]interface{}))
		ret.Lw4o6Cfg = getObjectRouterBgpRedistributeLw4o6Cfg1321(in["lw4o6_cfg"].([]interface{}))
		ret.StaticNatCfg = getObjectRouterBgpRedistributeStaticNatCfg1322(in["static_nat_cfg"].([]interface{}))
		ret.IpNatCfg = getObjectRouterBgpRedistributeIpNatCfg1323(in["ip_nat_cfg"].([]interface{}))
		ret.IpNatListCfg = getObjectRouterBgpRedistributeIpNatListCfg1324(in["ip_nat_list_cfg"].([]interface{}))
		ret.IsisCfg = getObjectRouterBgpRedistributeIsisCfg1325(in["isis_cfg"].([]interface{}))
		ret.OspfCfg = getObjectRouterBgpRedistributeOspfCfg1326(in["ospf_cfg"].([]interface{}))
		ret.RipCfg = getObjectRouterBgpRedistributeRipCfg1327(in["rip_cfg"].([]interface{}))
		ret.StaticCfg = getObjectRouterBgpRedistributeStaticCfg1328(in["static_cfg"].([]interface{}))
		ret.NatMapCfg = getObjectRouterBgpRedistributeNatMapCfg1329(in["nat_map_cfg"].([]interface{}))
		ret.PublicIpCfg = getObjectRouterBgpRedistributePublicIpCfg1330(in["public_ip_cfg"].([]interface{}))
		ret.Vip = getObjectRouterBgpRedistributeVip1331(in["vip"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectRouterBgpRedistributeConnectedCfg1319(d []interface{}) edpt.RouterBgpRedistributeConnectedCfg1319 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeConnectedCfg1319
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Connected = in["connected"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeFloatingIpCfg1320(d []interface{}) edpt.RouterBgpRedistributeFloatingIpCfg1320 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeFloatingIpCfg1320
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FloatingIp = in["floating_ip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeLw4o6Cfg1321(d []interface{}) edpt.RouterBgpRedistributeLw4o6Cfg1321 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeLw4o6Cfg1321
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Lw4o6 = in["lw4o6"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeStaticNatCfg1322(d []interface{}) edpt.RouterBgpRedistributeStaticNatCfg1322 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeStaticNatCfg1322
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticNat = in["static_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeIpNatCfg1323(d []interface{}) edpt.RouterBgpRedistributeIpNatCfg1323 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeIpNatCfg1323
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpNat = in["ip_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeIpNatListCfg1324(d []interface{}) edpt.RouterBgpRedistributeIpNatListCfg1324 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeIpNatListCfg1324
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpNatList = in["ip_nat_list"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeIsisCfg1325(d []interface{}) edpt.RouterBgpRedistributeIsisCfg1325 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeIsisCfg1325
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isis = in["isis"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeOspfCfg1326(d []interface{}) edpt.RouterBgpRedistributeOspfCfg1326 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeOspfCfg1326
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ospf = in["ospf"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeRipCfg1327(d []interface{}) edpt.RouterBgpRedistributeRipCfg1327 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeRipCfg1327
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeStaticCfg1328(d []interface{}) edpt.RouterBgpRedistributeStaticCfg1328 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeStaticCfg1328
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Static = in["static"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeNatMapCfg1329(d []interface{}) edpt.RouterBgpRedistributeNatMapCfg1329 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeNatMapCfg1329
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatMap = in["nat_map"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributePublicIpCfg1330(d []interface{}) edpt.RouterBgpRedistributePublicIpCfg1330 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributePublicIpCfg1330
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PublicIp = in["public_ip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeVip1331(d []interface{}) edpt.RouterBgpRedistributeVip1331 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeVip1331
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyFlaggedCfg = getObjectRouterBgpRedistributeVipOnlyFlaggedCfg1332(in["only_flagged_cfg"].([]interface{}))
		ret.OnlyNotFlaggedCfg = getObjectRouterBgpRedistributeVipOnlyNotFlaggedCfg1333(in["only_not_flagged_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpRedistributeVipOnlyFlaggedCfg1332(d []interface{}) edpt.RouterBgpRedistributeVipOnlyFlaggedCfg1332 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeVipOnlyFlaggedCfg1332
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyFlagged = in["only_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeVipOnlyNotFlaggedCfg1333(d []interface{}) edpt.RouterBgpRedistributeVipOnlyNotFlaggedCfg1333 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeVipOnlyNotFlaggedCfg1333
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyNotFlagged = in["only_not_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpTimers(d []interface{}) edpt.RouterBgpTimers {

	count1 := len(d)
	var ret edpt.RouterBgpTimers
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BgpKeepalive = in["bgp_keepalive"].(int)
		ret.BgpHoldtime = in["bgp_holdtime"].(int)
	}
	return ret
}

func dataToEndpointRouterBgp(d *schema.ResourceData) edpt.RouterBgp {
	var ret edpt.RouterBgp
	ret.Inst.AddressFamily = getObjectRouterBgpAddressFamily1258(d.Get("address_family").([]interface{}))
	ret.Inst.AggregateAddressList = getSliceRouterBgpAggregateAddressList(d.Get("aggregate_address_list").([]interface{}))
	ret.Inst.AsNumber = d.Get("as_number").(string)
	ret.Inst.AutoSummary = d.Get("auto_summary").(int)
	ret.Inst.Bgp = getObjectRouterBgpBgp(d.Get("bgp").([]interface{}))
	ret.Inst.DistanceList = getSliceRouterBgpDistanceList(d.Get("distance_list").([]interface{}))
	ret.Inst.MaximumPathsValue = d.Get("maximum_paths_value").(int)
	ret.Inst.Neighbor = getObjectRouterBgpNeighbor1313(d.Get("neighbor").([]interface{}))
	ret.Inst.Network = getObjectRouterBgpNetwork1314(d.Get("network").([]interface{}))
	ret.Inst.Originate = d.Get("originate").(int)
	ret.Inst.Redistribute = getObjectRouterBgpRedistribute1318(d.Get("redistribute").([]interface{}))
	ret.Inst.Synchronization = d.Get("synchronization").(int)
	ret.Inst.Timers = getObjectRouterBgpTimers(d.Get("timers").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
