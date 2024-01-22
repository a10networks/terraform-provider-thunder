package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpAddressFamilyIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_address_family_ipv6`: ipv6 Address family\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpAddressFamilyIpv6Create,
		UpdateContext: resourceRouterBgpAddressFamilyIpv6Update,
		ReadContext:   resourceRouterBgpAddressFamilyIpv6Read,
		DeleteContext: resourceRouterBgpAddressFamilyIpv6Delete,

		Schema: map[string]*schema.Schema{
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
			"originate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Distribute an IPv6 default route",
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
			"synchronization": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Perform IGP synchronization",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"as_number": {
				Type: schema.TypeString, Required: true, Description: "AsNumber",
			},
		},
	}
}
func resourceRouterBgpAddressFamilyIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpAddressFamilyIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterBgpAddressFamilyIpv6AggregateAddressList(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6AggregateAddressList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6AggregateAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6AggregateAddressList
		oi.AggregateAddress = in["aggregate_address"].(string)
		oi.AsSet = in["as_set"].(int)
		oi.SummaryOnly = in["summary_only"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Bgp(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Bgp {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Bgp
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

func getObjectRouterBgpAddressFamilyIpv6Distance(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Distance {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Distance
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DistanceExt = in["distance_ext"].(int)
		ret.DistanceInt = in["distance_int"].(int)
		ret.DistanceLocal = in["distance_local"].(int)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Neighbor1153(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Neighbor1153 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Neighbor1153
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PeerGroupNeighborList = getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList(in["peer_group_neighbor_list"].([]interface{}))
		ret.Ipv4NeighborList = getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborList(in["ipv4_neighbor_list"].([]interface{}))
		ret.Ipv6NeighborList = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborList(in["ipv6_neighbor_list"].([]interface{}))
		ret.EthernetNeighborIpv6List = getSliceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List(in["ethernet_neighbor_ipv6_list"].([]interface{}))
		ret.VeNeighborIpv6List = getSliceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List(in["ve_neighbor_ipv6_list"].([]interface{}))
		ret.TrunkNeighborIpv6List = getSliceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List(in["trunk_neighbor_ipv6_list"].([]interface{}))
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList
		oi.PeerGroup = in["peer_group"].(string)
		oi.Activate = in["activate"].(int)
		oi.AllowasIn = in["allowas_in"].(int)
		oi.AllowasInCount = in["allowas_in_count"].(int)
		oi.MaximumPrefix = in["maximum_prefix"].(int)
		oi.MaximumPrefixThres = in["maximum_prefix_thres"].(int)
		oi.NextHopSelf = in["next_hop_self"].(int)
		oi.RemovePrivateAs = in["remove_private_as"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists(in["neighbor_route_map_lists"].([]interface{}))
		oi.Inbound = in["inbound"].(int)
		oi.Weight = in["weight"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborList(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList
		oi.NeighborIpv4 = in["neighbor_ipv4"].(string)
		oi.PeerGroupName = in["peer_group_name"].(string)
		oi.Activate = in["activate"].(int)
		oi.AllowasIn = in["allowas_in"].(int)
		oi.AllowasInCount = in["allowas_in_count"].(int)
		oi.PrefixListDirection = in["prefix_list_direction"].(string)
		oi.GracefulRestart = in["graceful_restart"].(int)
		oi.DefaultOriginate = in["default_originate"].(int)
		oi.RouteMap = in["route_map"].(string)
		oi.DistributeLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists(in["distribute_lists"].([]interface{}))
		oi.NeighborFilterLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists(in["neighbor_filter_lists"].([]interface{}))
		oi.MaximumPrefix = in["maximum_prefix"].(int)
		oi.MaximumPrefixThres = in["maximum_prefix_thres"].(int)
		oi.RestartMin = in["restart_min"].(int)
		oi.NextHopSelf = in["next_hop_self"].(int)
		oi.NeighborPrefixLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists(in["neighbor_prefix_lists"].([]interface{}))
		oi.RemovePrivateAs = in["remove_private_as"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		oi.Inbound = in["inbound"].(int)
		oi.UnsuppressMap = in["unsuppress_map"].(string)
		oi.Weight = in["weight"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists
		oi.DistributeList = in["distribute_list"].(string)
		oi.DistributeListDirection = in["distribute_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists
		oi.FilterList = in["filter_list"].(string)
		oi.FilterListDirection = in["filter_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists
		oi.NbrPrefixList = in["nbr_prefix_list"].(string)
		oi.NbrPrefixListDirection = in["nbr_prefix_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborList(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList
		oi.NeighborIpv6 = in["neighbor_ipv6"].(string)
		oi.PeerGroupName = in["peer_group_name"].(string)
		oi.Activate = in["activate"].(int)
		oi.AllowasIn = in["allowas_in"].(int)
		oi.AllowasInCount = in["allowas_in_count"].(int)
		oi.PrefixListDirection = in["prefix_list_direction"].(string)
		oi.GracefulRestart = in["graceful_restart"].(int)
		oi.DefaultOriginate = in["default_originate"].(int)
		oi.RouteMap = in["route_map"].(string)
		oi.DistributeLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists(in["distribute_lists"].([]interface{}))
		oi.NeighborFilterLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists(in["neighbor_filter_lists"].([]interface{}))
		oi.MaximumPrefix = in["maximum_prefix"].(int)
		oi.MaximumPrefixThres = in["maximum_prefix_thres"].(int)
		oi.RestartMin = in["restart_min"].(int)
		oi.NextHopSelf = in["next_hop_self"].(int)
		oi.NeighborPrefixLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists(in["neighbor_prefix_lists"].([]interface{}))
		oi.RemovePrivateAs = in["remove_private_as"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		oi.Inbound = in["inbound"].(int)
		oi.UnsuppressMap = in["unsuppress_map"].(string)
		oi.Weight = in["weight"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists
		oi.DistributeList = in["distribute_list"].(string)
		oi.DistributeListDirection = in["distribute_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists
		oi.FilterList = in["filter_list"].(string)
		oi.FilterListDirection = in["filter_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists
		oi.NbrPrefixList = in["nbr_prefix_list"].(string)
		oi.NbrPrefixListDirection = in["nbr_prefix_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List
		oi.Ethernet = in["ethernet"].(int)
		oi.PeerGroupName = in["peer_group_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List
		oi.Ve = in["ve"].(int)
		oi.PeerGroupName = in["peer_group_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List
		oi.Trunk = in["trunk"].(int)
		oi.PeerGroupName = in["peer_group_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Network1154(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Network1154 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Network1154
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Synchronization = getObjectRouterBgpAddressFamilyIpv6NetworkSynchronization1155(in["synchronization"].([]interface{}))
		ret.Monitor = getObjectRouterBgpAddressFamilyIpv6NetworkMonitor1156(in["monitor"].([]interface{}))
		ret.Ipv6NetworkList = getSliceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkList(in["ipv6_network_list"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6NetworkSynchronization1155(d []interface{}) edpt.RouterBgpAddressFamilyIpv6NetworkSynchronization1155 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6NetworkSynchronization1155
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkSynchronization = in["network_synchronization"].(int)
		//omit uuid
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6NetworkMonitor1156(d []interface{}) edpt.RouterBgpAddressFamilyIpv6NetworkMonitor1156 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6NetworkMonitor1156
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Default = getObjectRouterBgpAddressFamilyIpv6NetworkMonitorDefault1157(in["default"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6NetworkMonitorDefault1157(d []interface{}) edpt.RouterBgpAddressFamilyIpv6NetworkMonitorDefault1157 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6NetworkMonitorDefault1157
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkMonitorDefault = in["network_monitor_default"].(int)
		//omit uuid
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkList(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList {

	count1 := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList
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

func getObjectRouterBgpAddressFamilyIpv6Redistribute1158(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Redistribute1158 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6Redistribute1158
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConnectedCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeConnectedCfg1159(in["connected_cfg"].([]interface{}))
		ret.FloatingIpCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1160(in["floating_ip_cfg"].([]interface{}))
		ret.Nat64Cfg = getObjectRouterBgpAddressFamilyIpv6RedistributeNat64Cfg1161(in["nat64_cfg"].([]interface{}))
		ret.NatMapCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeNatMapCfg1162(in["nat_map_cfg"].([]interface{}))
		ret.Lw4o6Cfg = getObjectRouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1163(in["lw4o6_cfg"].([]interface{}))
		ret.StaticNatCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1164(in["static_nat_cfg"].([]interface{}))
		ret.IpNatCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatCfg1165(in["ip_nat_cfg"].([]interface{}))
		ret.IpNatListCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1166(in["ip_nat_list_cfg"].([]interface{}))
		ret.IsisCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeIsisCfg1167(in["isis_cfg"].([]interface{}))
		ret.OspfCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeOspfCfg1168(in["ospf_cfg"].([]interface{}))
		ret.RipCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeRipCfg1169(in["rip_cfg"].([]interface{}))
		ret.StaticCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeStaticCfg1170(in["static_cfg"].([]interface{}))
		ret.Vip = getObjectRouterBgpAddressFamilyIpv6RedistributeVip1171(in["vip"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeConnectedCfg1159(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeConnectedCfg1159 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeConnectedCfg1159
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Connected = in["connected"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1160(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1160 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1160
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FloatingIp = in["floating_ip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeNat64Cfg1161(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeNat64Cfg1161 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeNat64Cfg1161
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Nat64 = in["nat64"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeNatMapCfg1162(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeNatMapCfg1162 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeNatMapCfg1162
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatMap = in["nat_map"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1163(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1163 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1163
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Lw4o6 = in["lw4o6"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1164(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1164 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1164
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticNat = in["static_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatCfg1165(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatCfg1165 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatCfg1165
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpNat = in["ip_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1166(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1166 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1166
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpNatList = in["ip_nat_list"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeIsisCfg1167(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeIsisCfg1167 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeIsisCfg1167
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isis = in["isis"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeOspfCfg1168(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeOspfCfg1168 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeOspfCfg1168
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ospf = in["ospf"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeRipCfg1169(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeRipCfg1169 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeRipCfg1169
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeStaticCfg1170(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeStaticCfg1170 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeStaticCfg1170
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Static = in["static"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeVip1171(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeVip1171 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeVip1171
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyFlaggedCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1172(in["only_flagged_cfg"].([]interface{}))
		ret.OnlyNotFlaggedCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1173(in["only_not_flagged_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1172(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1172 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1172
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyFlagged = in["only_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1173(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1173 {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1173
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyNotFlagged = in["only_not_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func dataToEndpointRouterBgpAddressFamilyIpv6(d *schema.ResourceData) edpt.RouterBgpAddressFamilyIpv6 {
	var ret edpt.RouterBgpAddressFamilyIpv6
	ret.Inst.AggregateAddressList = getSliceRouterBgpAddressFamilyIpv6AggregateAddressList(d.Get("aggregate_address_list").([]interface{}))
	ret.Inst.AutoSummary = d.Get("auto_summary").(int)
	ret.Inst.Bgp = getObjectRouterBgpAddressFamilyIpv6Bgp(d.Get("bgp").([]interface{}))
	ret.Inst.Distance = getObjectRouterBgpAddressFamilyIpv6Distance(d.Get("distance").([]interface{}))
	ret.Inst.MaximumPathsValue = d.Get("maximum_paths_value").(int)
	ret.Inst.Neighbor = getObjectRouterBgpAddressFamilyIpv6Neighbor1153(d.Get("neighbor").([]interface{}))
	ret.Inst.Network = getObjectRouterBgpAddressFamilyIpv6Network1154(d.Get("network").([]interface{}))
	ret.Inst.Originate = d.Get("originate").(int)
	ret.Inst.Redistribute = getObjectRouterBgpAddressFamilyIpv6Redistribute1158(d.Get("redistribute").([]interface{}))
	ret.Inst.Synchronization = d.Get("synchronization").(int)
	//omit uuid
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
