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

func getObjectRouterBgpAddressFamily1174(d []interface{}) edpt.RouterBgpAddressFamily1174 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamily1174
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv6 = getObjectRouterBgpAddressFamilyIpv61175(in["ipv6"].([]interface{}))
		ret.Ipv4Flowspec = getObjectRouterBgpAddressFamilyIpv4Flowspec1216(in["ipv4_flowspec"].([]interface{}))
		ret.Ipv6Flowspec = getObjectRouterBgpAddressFamilyIpv6Flowspec1222(in["ipv6_flowspec"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv61175(d []interface{}) edpt.RouterBgpAddressFamilyIpv61175 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv61175
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bgp = getObjectRouterBgpAddressFamilyIpv6Bgp1176(in["bgp"].([]interface{}))
		ret.Distance = getObjectRouterBgpAddressFamilyIpv6Distance1177(in["distance"].([]interface{}))
		ret.MaximumPathsValue = in["maximum_paths_value"].(int)
		ret.Originate = in["originate"].(int)
		ret.AggregateAddressList = getSliceRouterBgpAddressFamilyIpv6AggregateAddressList1178(in["aggregate_address_list"].([]interface{}))
		ret.AutoSummary = in["auto_summary"].(int)
		ret.Synchronization = in["synchronization"].(int)
		//omit uuid
		ret.Network = getObjectRouterBgpAddressFamilyIpv6Network1179(in["network"].([]interface{}))
		ret.Neighbor = getObjectRouterBgpAddressFamilyIpv6Neighbor1184(in["neighbor"].([]interface{}))
		ret.Redistribute = getObjectRouterBgpAddressFamilyIpv6Redistribute1200(in["redistribute"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Bgp1176(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Bgp1176 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Bgp1176
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

func getObjectRouterBgpAddressFamilyIpv6Distance1177(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Distance1177 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Distance1177
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DistanceExt = in["distance_ext"].(int)
		ret.DistanceInt = in["distance_int"].(int)
		ret.DistanceLocal = in["distance_local"].(int)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6AggregateAddressList1178(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6AggregateAddressList1178 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6AggregateAddressList1178, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6AggregateAddressList1178
		oi.AggregateAddress = in["aggregate_address"].(string)
		oi.AsSet = in["as_set"].(int)
		oi.SummaryOnly = in["summary_only"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Network1179(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Network1179 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Network1179
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Synchronization = getObjectRouterBgpAddressFamilyIpv6NetworkSynchronization1180(in["synchronization"].([]interface{}))
		ret.Monitor = getObjectRouterBgpAddressFamilyIpv6NetworkMonitor1181(in["monitor"].([]interface{}))
		ret.Ipv6NetworkList = getSliceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkList1183(in["ipv6_network_list"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6NetworkSynchronization1180(d []interface{}) edpt.RouterBgpAddressFamilyIpv6NetworkSynchronization1180 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6NetworkSynchronization1180
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkSynchronization = in["network_synchronization"].(int)
		//omit uuid
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6NetworkMonitor1181(d []interface{}) edpt.RouterBgpAddressFamilyIpv6NetworkMonitor1181 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6NetworkMonitor1181
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Default = getObjectRouterBgpAddressFamilyIpv6NetworkMonitorDefault1182(in["default"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6NetworkMonitorDefault1182(d []interface{}) edpt.RouterBgpAddressFamilyIpv6NetworkMonitorDefault1182 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6NetworkMonitorDefault1182
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkMonitorDefault = in["network_monitor_default"].(int)
		//omit uuid
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkList1183(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList1183 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList1183, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList1183
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

func getObjectRouterBgpAddressFamilyIpv6Neighbor1184(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Neighbor1184 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Neighbor1184
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PeerGroupNeighborList = getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList1185(in["peer_group_neighbor_list"].([]interface{}))
		ret.Ipv4NeighborList = getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborList1187(in["ipv4_neighbor_list"].([]interface{}))
		ret.Ipv6NeighborList = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborList1192(in["ipv6_neighbor_list"].([]interface{}))
		ret.EthernetNeighborIpv6List = getSliceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List1197(in["ethernet_neighbor_ipv6_list"].([]interface{}))
		ret.VeNeighborIpv6List = getSliceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List1198(in["ve_neighbor_ipv6_list"].([]interface{}))
		ret.TrunkNeighborIpv6List = getSliceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List1199(in["trunk_neighbor_ipv6_list"].([]interface{}))
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList1185(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList1185 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList1185, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList1185
		oi.PeerGroup = in["peer_group"].(string)
		oi.Activate = in["activate"].(int)
		oi.AllowasIn = in["allowas_in"].(int)
		oi.AllowasInCount = in["allowas_in_count"].(int)
		oi.MaximumPrefix = in["maximum_prefix"].(int)
		oi.MaximumPrefixThres = in["maximum_prefix_thres"].(int)
		oi.NextHopSelf = in["next_hop_self"].(int)
		oi.RemovePrivateAs = in["remove_private_as"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists1186(in["neighbor_route_map_lists"].([]interface{}))
		oi.Inbound = in["inbound"].(int)
		oi.Weight = in["weight"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists1186(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists1186 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists1186, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists1186
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborList1187(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList1187 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList1187, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList1187
		oi.NeighborIpv4 = in["neighbor_ipv4"].(string)
		oi.PeerGroupName = in["peer_group_name"].(string)
		oi.Activate = in["activate"].(int)
		oi.AllowasIn = in["allowas_in"].(int)
		oi.AllowasInCount = in["allowas_in_count"].(int)
		oi.PrefixListDirection = in["prefix_list_direction"].(string)
		oi.GracefulRestart = in["graceful_restart"].(int)
		oi.DefaultOriginate = in["default_originate"].(int)
		oi.RouteMap = in["route_map"].(string)
		oi.DistributeLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists1188(in["distribute_lists"].([]interface{}))
		oi.NeighborFilterLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists1189(in["neighbor_filter_lists"].([]interface{}))
		oi.MaximumPrefix = in["maximum_prefix"].(int)
		oi.MaximumPrefixThres = in["maximum_prefix_thres"].(int)
		oi.RestartMin = in["restart_min"].(int)
		oi.NextHopSelf = in["next_hop_self"].(int)
		oi.NeighborPrefixLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists1190(in["neighbor_prefix_lists"].([]interface{}))
		oi.RemovePrivateAs = in["remove_private_as"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists1191(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		oi.Inbound = in["inbound"].(int)
		oi.UnsuppressMap = in["unsuppress_map"].(string)
		oi.Weight = in["weight"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists1188(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists1188 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists1188, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists1188
		oi.DistributeList = in["distribute_list"].(string)
		oi.DistributeListDirection = in["distribute_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists1189(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists1189 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists1189, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists1189
		oi.FilterList = in["filter_list"].(string)
		oi.FilterListDirection = in["filter_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists1190(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists1190 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists1190, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists1190
		oi.NbrPrefixList = in["nbr_prefix_list"].(string)
		oi.NbrPrefixListDirection = in["nbr_prefix_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists1191(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists1191 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists1191, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists1191
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborList1192(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList1192 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList1192, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList1192
		oi.NeighborIpv6 = in["neighbor_ipv6"].(string)
		oi.PeerGroupName = in["peer_group_name"].(string)
		oi.Activate = in["activate"].(int)
		oi.AllowasIn = in["allowas_in"].(int)
		oi.AllowasInCount = in["allowas_in_count"].(int)
		oi.PrefixListDirection = in["prefix_list_direction"].(string)
		oi.GracefulRestart = in["graceful_restart"].(int)
		oi.DefaultOriginate = in["default_originate"].(int)
		oi.RouteMap = in["route_map"].(string)
		oi.DistributeLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists1193(in["distribute_lists"].([]interface{}))
		oi.NeighborFilterLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists1194(in["neighbor_filter_lists"].([]interface{}))
		oi.MaximumPrefix = in["maximum_prefix"].(int)
		oi.MaximumPrefixThres = in["maximum_prefix_thres"].(int)
		oi.RestartMin = in["restart_min"].(int)
		oi.NextHopSelf = in["next_hop_self"].(int)
		oi.NeighborPrefixLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists1195(in["neighbor_prefix_lists"].([]interface{}))
		oi.RemovePrivateAs = in["remove_private_as"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists1196(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		oi.Inbound = in["inbound"].(int)
		oi.UnsuppressMap = in["unsuppress_map"].(string)
		oi.Weight = in["weight"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists1193(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists1193 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists1193, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists1193
		oi.DistributeList = in["distribute_list"].(string)
		oi.DistributeListDirection = in["distribute_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists1194(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists1194 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists1194, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists1194
		oi.FilterList = in["filter_list"].(string)
		oi.FilterListDirection = in["filter_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists1195(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists1195 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists1195, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists1195
		oi.NbrPrefixList = in["nbr_prefix_list"].(string)
		oi.NbrPrefixListDirection = in["nbr_prefix_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists1196(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists1196 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists1196, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists1196
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List1197(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List1197 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List1197, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List1197
		oi.Ethernet = in["ethernet"].(int)
		oi.PeerGroupName = in["peer_group_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List1198(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List1198 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List1198, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List1198
		oi.Ve = in["ve"].(int)
		oi.PeerGroupName = in["peer_group_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List1199(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List1199 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List1199, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List1199
		oi.Trunk = in["trunk"].(int)
		oi.PeerGroupName = in["peer_group_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Redistribute1200(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Redistribute1200 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Redistribute1200
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConnectedCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeConnectedCfg1201(in["connected_cfg"].([]interface{}))
		ret.FloatingIpCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1202(in["floating_ip_cfg"].([]interface{}))
		ret.Nat64Cfg = getObjectRouterBgpAddressFamilyIpv6RedistributeNat64Cfg1203(in["nat64_cfg"].([]interface{}))
		ret.NatMapCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeNatMapCfg1204(in["nat_map_cfg"].([]interface{}))
		ret.Lw4o6Cfg = getObjectRouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1205(in["lw4o6_cfg"].([]interface{}))
		ret.StaticNatCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1206(in["static_nat_cfg"].([]interface{}))
		ret.IpNatCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatCfg1207(in["ip_nat_cfg"].([]interface{}))
		ret.IpNatListCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1208(in["ip_nat_list_cfg"].([]interface{}))
		ret.IsisCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeIsisCfg1209(in["isis_cfg"].([]interface{}))
		ret.OspfCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeOspfCfg1210(in["ospf_cfg"].([]interface{}))
		ret.RipCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeRipCfg1211(in["rip_cfg"].([]interface{}))
		ret.StaticCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeStaticCfg1212(in["static_cfg"].([]interface{}))
		ret.Vip = getObjectRouterBgpAddressFamilyIpv6RedistributeVip1213(in["vip"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeConnectedCfg1201(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeConnectedCfg1201 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeConnectedCfg1201
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Connected = in["connected"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1202(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1202 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1202
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FloatingIp = in["floating_ip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeNat64Cfg1203(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeNat64Cfg1203 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeNat64Cfg1203
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Nat64 = in["nat64"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeNatMapCfg1204(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeNatMapCfg1204 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeNatMapCfg1204
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatMap = in["nat_map"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1205(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1205 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1205
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Lw4o6 = in["lw4o6"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1206(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1206 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1206
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticNat = in["static_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatCfg1207(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatCfg1207 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatCfg1207
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpNat = in["ip_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1208(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1208 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1208
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpNatList = in["ip_nat_list"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeIsisCfg1209(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeIsisCfg1209 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeIsisCfg1209
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isis = in["isis"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeOspfCfg1210(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeOspfCfg1210 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeOspfCfg1210
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ospf = in["ospf"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeRipCfg1211(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeRipCfg1211 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeRipCfg1211
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeStaticCfg1212(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeStaticCfg1212 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeStaticCfg1212
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Static = in["static"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeVip1213(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeVip1213 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeVip1213
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyFlaggedCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1214(in["only_flagged_cfg"].([]interface{}))
		ret.OnlyNotFlaggedCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1215(in["only_not_flagged_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1214(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1214 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1214
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyFlagged = in["only_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1215(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1215 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1215
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyNotFlagged = in["only_not_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv4Flowspec1216(d []interface{}) edpt.RouterBgpAddressFamilyIpv4Flowspec1216 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv4Flowspec1216
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.Neighbor = getObjectRouterBgpAddressFamilyIpv4FlowspecNeighbor1217(in["neighbor"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv4FlowspecNeighbor1217(d []interface{}) edpt.RouterBgpAddressFamilyIpv4FlowspecNeighbor1217 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv4FlowspecNeighbor1217
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4NeighborList = getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList1218(in["ipv4_neighbor_list"].([]interface{}))
		ret.Ipv6NeighborList = getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList1220(in["ipv6_neighbor_list"].([]interface{}))
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList1218(d []interface{}) []edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList1218 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList1218, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList1218
		oi.NeighborIpv4 = in["neighbor_ipv4"].(string)
		oi.Activate = in["activate"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1219(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1219(d []interface{}) []edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1219 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1219, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1219
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList1220(d []interface{}) []edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList1220 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList1220, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList1220
		oi.NeighborIpv6 = in["neighbor_ipv6"].(string)
		oi.Activate = in["activate"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1221(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1221(d []interface{}) []edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1221 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1221, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1221
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Flowspec1222(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Flowspec1222 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Flowspec1222
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.Neighbor = getObjectRouterBgpAddressFamilyIpv6FlowspecNeighbor1223(in["neighbor"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6FlowspecNeighbor1223(d []interface{}) edpt.RouterBgpAddressFamilyIpv6FlowspecNeighbor1223 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6FlowspecNeighbor1223
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4NeighborList = getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList1224(in["ipv4_neighbor_list"].([]interface{}))
		ret.Ipv6NeighborList = getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList1226(in["ipv6_neighbor_list"].([]interface{}))
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList1224(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList1224 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList1224, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList1224
		oi.NeighborIpv4 = in["neighbor_ipv4"].(string)
		oi.Activate = in["activate"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1225(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1225(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1225 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1225, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1225
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList1226(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList1226 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList1226, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList1226
		oi.NeighborIpv6 = in["neighbor_ipv6"].(string)
		oi.Activate = in["activate"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1227(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1227(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1227 {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1227, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1227
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

func getObjectRouterBgpNeighbor1228(d []interface{}) edpt.RouterBgpNeighbor1228 {

	count1 := len(d)
	var ret edpt.RouterBgpNeighbor1228
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

func getObjectRouterBgpNetwork1229(d []interface{}) edpt.RouterBgpNetwork1229 {

	count1 := len(d)
	var ret edpt.RouterBgpNetwork1229
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Synchronization = getObjectRouterBgpNetworkSynchronization1230(in["synchronization"].([]interface{}))
		ret.Monitor = getObjectRouterBgpNetworkMonitor1231(in["monitor"].([]interface{}))
		ret.IpCidrList = getSliceRouterBgpNetworkIpCidrList(in["ip_cidr_list"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpNetworkSynchronization1230(d []interface{}) edpt.RouterBgpNetworkSynchronization1230 {

	count1 := len(d)
	var ret edpt.RouterBgpNetworkSynchronization1230
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkSynchronization = in["network_synchronization"].(int)
		//omit uuid
	}
	return ret
}

func getObjectRouterBgpNetworkMonitor1231(d []interface{}) edpt.RouterBgpNetworkMonitor1231 {

	count1 := len(d)
	var ret edpt.RouterBgpNetworkMonitor1231
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Default = getObjectRouterBgpNetworkMonitorDefault1232(in["default"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpNetworkMonitorDefault1232(d []interface{}) edpt.RouterBgpNetworkMonitorDefault1232 {

	count1 := len(d)
	var ret edpt.RouterBgpNetworkMonitorDefault1232
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

func getObjectRouterBgpRedistribute1233(d []interface{}) edpt.RouterBgpRedistribute1233 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistribute1233
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConnectedCfg = getObjectRouterBgpRedistributeConnectedCfg1234(in["connected_cfg"].([]interface{}))
		ret.FloatingIpCfg = getObjectRouterBgpRedistributeFloatingIpCfg1235(in["floating_ip_cfg"].([]interface{}))
		ret.Lw4o6Cfg = getObjectRouterBgpRedistributeLw4o6Cfg1236(in["lw4o6_cfg"].([]interface{}))
		ret.StaticNatCfg = getObjectRouterBgpRedistributeStaticNatCfg1237(in["static_nat_cfg"].([]interface{}))
		ret.IpNatCfg = getObjectRouterBgpRedistributeIpNatCfg1238(in["ip_nat_cfg"].([]interface{}))
		ret.IpNatListCfg = getObjectRouterBgpRedistributeIpNatListCfg1239(in["ip_nat_list_cfg"].([]interface{}))
		ret.IsisCfg = getObjectRouterBgpRedistributeIsisCfg1240(in["isis_cfg"].([]interface{}))
		ret.OspfCfg = getObjectRouterBgpRedistributeOspfCfg1241(in["ospf_cfg"].([]interface{}))
		ret.RipCfg = getObjectRouterBgpRedistributeRipCfg1242(in["rip_cfg"].([]interface{}))
		ret.StaticCfg = getObjectRouterBgpRedistributeStaticCfg1243(in["static_cfg"].([]interface{}))
		ret.NatMapCfg = getObjectRouterBgpRedistributeNatMapCfg1244(in["nat_map_cfg"].([]interface{}))
		ret.Vip = getObjectRouterBgpRedistributeVip1245(in["vip"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectRouterBgpRedistributeConnectedCfg1234(d []interface{}) edpt.RouterBgpRedistributeConnectedCfg1234 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeConnectedCfg1234
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Connected = in["connected"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeFloatingIpCfg1235(d []interface{}) edpt.RouterBgpRedistributeFloatingIpCfg1235 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeFloatingIpCfg1235
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FloatingIp = in["floating_ip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeLw4o6Cfg1236(d []interface{}) edpt.RouterBgpRedistributeLw4o6Cfg1236 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeLw4o6Cfg1236
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Lw4o6 = in["lw4o6"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeStaticNatCfg1237(d []interface{}) edpt.RouterBgpRedistributeStaticNatCfg1237 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeStaticNatCfg1237
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticNat = in["static_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeIpNatCfg1238(d []interface{}) edpt.RouterBgpRedistributeIpNatCfg1238 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeIpNatCfg1238
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpNat = in["ip_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeIpNatListCfg1239(d []interface{}) edpt.RouterBgpRedistributeIpNatListCfg1239 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeIpNatListCfg1239
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpNatList = in["ip_nat_list"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeIsisCfg1240(d []interface{}) edpt.RouterBgpRedistributeIsisCfg1240 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeIsisCfg1240
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isis = in["isis"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeOspfCfg1241(d []interface{}) edpt.RouterBgpRedistributeOspfCfg1241 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeOspfCfg1241
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ospf = in["ospf"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeRipCfg1242(d []interface{}) edpt.RouterBgpRedistributeRipCfg1242 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeRipCfg1242
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeStaticCfg1243(d []interface{}) edpt.RouterBgpRedistributeStaticCfg1243 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeStaticCfg1243
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Static = in["static"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeNatMapCfg1244(d []interface{}) edpt.RouterBgpRedistributeNatMapCfg1244 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeNatMapCfg1244
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatMap = in["nat_map"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeVip1245(d []interface{}) edpt.RouterBgpRedistributeVip1245 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeVip1245
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyFlaggedCfg = getObjectRouterBgpRedistributeVipOnlyFlaggedCfg1246(in["only_flagged_cfg"].([]interface{}))
		ret.OnlyNotFlaggedCfg = getObjectRouterBgpRedistributeVipOnlyNotFlaggedCfg1247(in["only_not_flagged_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpRedistributeVipOnlyFlaggedCfg1246(d []interface{}) edpt.RouterBgpRedistributeVipOnlyFlaggedCfg1246 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeVipOnlyFlaggedCfg1246
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyFlagged = in["only_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeVipOnlyNotFlaggedCfg1247(d []interface{}) edpt.RouterBgpRedistributeVipOnlyNotFlaggedCfg1247 {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeVipOnlyNotFlaggedCfg1247
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
	ret.Inst.AddressFamily = getObjectRouterBgpAddressFamily1174(d.Get("address_family").([]interface{}))
	ret.Inst.AggregateAddressList = getSliceRouterBgpAggregateAddressList(d.Get("aggregate_address_list").([]interface{}))
	ret.Inst.AsNumber = d.Get("as_number").(string)
	ret.Inst.AutoSummary = d.Get("auto_summary").(int)
	ret.Inst.Bgp = getObjectRouterBgpBgp(d.Get("bgp").([]interface{}))
	ret.Inst.DistanceList = getSliceRouterBgpDistanceList(d.Get("distance_list").([]interface{}))
	ret.Inst.MaximumPathsValue = d.Get("maximum_paths_value").(int)
	ret.Inst.Neighbor = getObjectRouterBgpNeighbor1228(d.Get("neighbor").([]interface{}))
	ret.Inst.Network = getObjectRouterBgpNetwork1229(d.Get("network").([]interface{}))
	ret.Inst.Originate = d.Get("originate").(int)
	ret.Inst.Redistribute = getObjectRouterBgpRedistribute1233(d.Get("redistribute").([]interface{}))
	ret.Inst.Synchronization = d.Get("synchronization").(int)
	ret.Inst.Timers = getObjectRouterBgpTimers(d.Get("timers").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
