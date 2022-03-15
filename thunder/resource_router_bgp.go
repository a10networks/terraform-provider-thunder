package thunder

import (
	"context"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBgp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp`: Border Gateway Protocol (BGP)\n\n",
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
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"dampening_half": {
													Type: schema.TypeInt, Optional: true, Description: "Reachability Half-life time for the penalty(minutes)",
													ValidateFunc: validation.IntBetween(1, 45),
												},
												"dampening_start_reuse": {
													Type: schema.TypeInt, Optional: true, Description: "Value to start reusing a route",
													ValidateFunc: validation.IntBetween(1, 20000),
												},
												"dampening_start_supress": {
													Type: schema.TypeInt, Optional: true, Description: "Value to start suppressing a route",
													ValidateFunc: validation.IntBetween(1, 20000),
												},
												"dampening_max_supress": {
													Type: schema.TypeInt, Optional: true, Description: "Maximum duration to suppress a stable route(minutes)",
													ValidateFunc: validation.IntBetween(1, 255),
												},
												"dampening_unreachability": {
													Type: schema.TypeInt, Optional: true, Description: "Un-reachability Half-life time for the penalty(minutes)",
													ValidateFunc: validation.IntBetween(1, 45),
												},
												"route_map": {
													Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria for dampening (Route-map name)",
													ValidateFunc: validation.StringLenBetween(1, 128),
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
													ValidateFunc: validation.IntBetween(1, 255),
												},
												"distance_int": {
													Type: schema.TypeInt, Optional: true, Description: "Distance for routes internal to the AS",
													ValidateFunc: validation.IntBetween(1, 255),
												},
												"distance_local": {
													Type: schema.TypeInt, Optional: true, Description: "Distance for local routes",
													ValidateFunc: validation.IntBetween(1, 255),
												},
											},
										},
									},
									"maximum_paths_value": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Supported BGP multipath numbers",
										ValidateFunc: validation.IntBetween(1, 64),
									},
									"originate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Distribute an IPv6 default route",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"aggregate_address_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"aggregate_address": {
													Type: schema.TypeString, Optional: true, Description: "Configure BGP aggregate entries (Aggregate IPv6 prefix)",
													ValidateFunc: axapi.IsIpv6AddressPlen,
												},
												"as_set": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Generate AS set path information",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"summary_only": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Filter more specific routes from updates",
													ValidateFunc: validation.IntBetween(0, 1),
												},
											},
										},
									},
									"auto_summary": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic network number summarization",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"synchronization": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Perform IGP synchronization",
										ValidateFunc: validation.IntBetween(0, 1),
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
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
																ValidateFunc: axapi.IsIpv6AddressPlen,
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route-map to modify the attributes (Name of the route map)",
																ValidateFunc: validation.StringLenBetween(1, 128),
															},
															"backdoor": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify a BGP backdoor route",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"description": {
																Type: schema.TypeString, Optional: true, Description: "Network specific description (Up to 80 characters describing this network)",
																ValidateFunc: validation.StringLenBetween(1, 80),
															},
															"comm_value": {
																Type: schema.TypeString, Optional: true, Description: "community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export",
																ValidateFunc: validation.StringLenBetween(1, 128),
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
																ValidateFunc: validation.StringLenBetween(1, 128),
															},
															"activate": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the Address Family for this Neighbor",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"allowas_in": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept as-path with my AS present in it",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"allowas_in_count": {
																Type: schema.TypeInt, Optional: true, Default: 3, Description: "Number of occurrences of AS number",
																ValidateFunc: validation.IntBetween(1, 10),
															},
															"prefix_list_direction": {
																Type: schema.TypeString, Optional: true, Description: "'both': both; 'receive': receive; 'send': send;",
																ValidateFunc: validation.StringInSlice([]string{"both", "receive", "send"}, false),
															},
															"default_originate": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Originate default route to this neighbor",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria to originate default (route-map name)",
																ValidateFunc: validation.StringLenBetween(1, 128),
															},
															"distribute_lists": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"distribute_list": {
																			Type: schema.TypeString, Optional: true, Description: "Filter updates to/from this neighbor (IP standard/extended/named access list)",
																			ValidateFunc: validation.StringLenBetween(1, 128),
																		},
																		"distribute_list_direction": {
																			Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
																			ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
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
																			ValidateFunc: validation.StringLenBetween(1, 128),
																		},
																		"filter_list_direction": {
																			Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
																			ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
																		},
																	},
																},
															},
															"maximum_prefix": {
																Type: schema.TypeInt, Optional: true, Description: "Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))",
															},
															"maximum_prefix_thres": {
																Type: schema.TypeInt, Optional: true, Description: "threshold-value, 1 to 100 percent",
																ValidateFunc: validation.IntBetween(1, 100),
															},
															"next_hop_self": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable the next hop calculation for this neighbor",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"neighbor_prefix_lists": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"nbr_prefix_list": {
																			Type: schema.TypeString, Optional: true, Description: "Filter updates to/from this neighbor (Name of a prefix list)",
																			ValidateFunc: validation.StringLenBetween(1, 128),
																		},
																		"nbr_prefix_list_direction": {
																			Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
																			ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
																		},
																	},
																},
															},
															"remove_private_as": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove private AS number from outbound updates",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"neighbor_route_map_lists": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"nbr_route_map": {
																			Type: schema.TypeString, Optional: true, Description: "Apply route map to neighbor (Name of route map)",
																			ValidateFunc: validation.StringLenBetween(1, 128),
																		},
																		"nbr_rmap_direction": {
																			Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
																			ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
																		},
																	},
																},
															},
															"send_community_val": {
																Type: schema.TypeString, Optional: true, Default: "both", Description: "'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;",
																ValidateFunc: validation.StringInSlice([]string{"both", "none", "standard", "extended"}, false),
															},
															"inbound": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow inbound soft reconfiguration for this neighbor",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"unsuppress_map": {
																Type: schema.TypeString, Optional: true, Description: "Route-map to selectively unsuppress suppressed routes (Name of route map)",
																ValidateFunc: validation.StringLenBetween(1, 128),
															},
															"weight": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set default weight for routes from this neighbor",
																ValidateFunc: validation.IntBetween(0, 65535),
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
																ValidateFunc: validation.IsIPv4Address,
															},
															"peer_group_name": {
																Type: schema.TypeString, Optional: true, Description: "Configure peer-group (peer-group name)",
																ValidateFunc: validation.StringLenBetween(1, 128),
															},
															"activate": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the Address Family for this Neighbor",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"allowas_in": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept as-path with my AS present in it",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"allowas_in_count": {
																Type: schema.TypeInt, Optional: true, Default: 3, Description: "Number of occurrences of AS number",
																ValidateFunc: validation.IntBetween(1, 10),
															},
															"prefix_list_direction": {
																Type: schema.TypeString, Optional: true, Description: "'both': both; 'receive': receive; 'send': send;",
																ValidateFunc: validation.StringInSlice([]string{"both", "receive", "send"}, false),
															},
															"graceful_restart": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable graceful-restart helper for this neighbor",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"default_originate": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Originate default route to this neighbor",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria to originate default (route-map name)",
																ValidateFunc: validation.StringLenBetween(1, 128),
															},
															"distribute_lists": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"distribute_list": {
																			Type: schema.TypeString, Optional: true, Description: "Filter updates to/from this neighbor (IP standard/extended/named access list)",
																			ValidateFunc: validation.StringLenBetween(1, 128),
																		},
																		"distribute_list_direction": {
																			Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
																			ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
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
																			ValidateFunc: validation.StringLenBetween(1, 128),
																		},
																		"filter_list_direction": {
																			Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
																			ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
																		},
																	},
																},
															},
															"maximum_prefix": {
																Type: schema.TypeInt, Optional: true, Description: "Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))",
															},
															"maximum_prefix_thres": {
																Type: schema.TypeInt, Optional: true, Description: "threshold-value, 1 to 100 percent",
																ValidateFunc: validation.IntBetween(1, 100),
															},
															"restart_min": {
																Type: schema.TypeInt, Optional: true, Description: "restart value, 1 to 1440 minutes",
																ValidateFunc: validation.IntBetween(1, 1440),
															},
															"next_hop_self": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable the next hop calculation for this neighbor",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"neighbor_prefix_lists": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"nbr_prefix_list": {
																			Type: schema.TypeString, Optional: true, Description: "Filter updates to/from this neighbor (Name of a prefix list)",
																			ValidateFunc: validation.StringLenBetween(1, 128),
																		},
																		"nbr_prefix_list_direction": {
																			Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
																			ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
																		},
																	},
																},
															},
															"remove_private_as": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove private AS number from outbound updates",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"neighbor_route_map_lists": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"nbr_route_map": {
																			Type: schema.TypeString, Optional: true, Description: "Apply route map to neighbor (Name of route map)",
																			ValidateFunc: validation.StringLenBetween(1, 128),
																		},
																		"nbr_rmap_direction": {
																			Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
																			ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
																		},
																	},
																},
															},
															"send_community_val": {
																Type: schema.TypeString, Optional: true, Default: "both", Description: "'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;",
																ValidateFunc: validation.StringInSlice([]string{"both", "none", "standard", "extended"}, false),
															},
															"inbound": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow inbound soft reconfiguration for this neighbor",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"unsuppress_map": {
																Type: schema.TypeString, Optional: true, Description: "Route-map to selectively unsuppress suppressed routes (Name of route map)",
																ValidateFunc: validation.StringLenBetween(1, 128),
															},
															"weight": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set default weight for routes from this neighbor",
																ValidateFunc: validation.IntBetween(0, 65535),
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
																ValidateFunc: validation.IsIPv6Address,
															},
															"peer_group_name": {
																Type: schema.TypeString, Optional: true, Description: "Configure peer-group (peer-group name)",
																ValidateFunc: validation.StringLenBetween(1, 128),
															},
															"activate": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the Address Family for this Neighbor",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"allowas_in": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept as-path with my AS present in it",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"allowas_in_count": {
																Type: schema.TypeInt, Optional: true, Default: 3, Description: "Number of occurrences of AS number",
																ValidateFunc: validation.IntBetween(1, 10),
															},
															"prefix_list_direction": {
																Type: schema.TypeString, Optional: true, Description: "'both': both; 'receive': receive; 'send': send;",
																ValidateFunc: validation.StringInSlice([]string{"both", "receive", "send"}, false),
															},
															"graceful_restart": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable graceful-restart helper for this neighbor",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"default_originate": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Originate default route to this neighbor",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria to originate default (route-map name)",
																ValidateFunc: validation.StringLenBetween(1, 128),
															},
															"distribute_lists": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"distribute_list": {
																			Type: schema.TypeString, Optional: true, Description: "Filter updates to/from this neighbor (IP standard/extended/named access list)",
																			ValidateFunc: validation.StringLenBetween(1, 128),
																		},
																		"distribute_list_direction": {
																			Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
																			ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
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
																			ValidateFunc: validation.StringLenBetween(1, 128),
																		},
																		"filter_list_direction": {
																			Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
																			ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
																		},
																	},
																},
															},
															"maximum_prefix": {
																Type: schema.TypeInt, Optional: true, Description: "Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))",
															},
															"maximum_prefix_thres": {
																Type: schema.TypeInt, Optional: true, Description: "threshold-value, 1 to 100 percent",
																ValidateFunc: validation.IntBetween(1, 100),
															},
															"restart_min": {
																Type: schema.TypeInt, Optional: true, Description: "restart value, 1 to 1440 minutes",
																ValidateFunc: validation.IntBetween(1, 1440),
															},
															"next_hop_self": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable the next hop calculation for this neighbor",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"neighbor_prefix_lists": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"nbr_prefix_list": {
																			Type: schema.TypeString, Optional: true, Description: "Filter updates to/from this neighbor (Name of a prefix list)",
																			ValidateFunc: validation.StringLenBetween(1, 128),
																		},
																		"nbr_prefix_list_direction": {
																			Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
																			ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
																		},
																	},
																},
															},
															"remove_private_as": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove private AS number from outbound updates",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"neighbor_route_map_lists": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"nbr_route_map": {
																			Type: schema.TypeString, Optional: true, Description: "Apply route map to neighbor (Name of route map)",
																			ValidateFunc: validation.StringLenBetween(1, 128),
																		},
																		"nbr_rmap_direction": {
																			Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
																			ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
																		},
																	},
																},
															},
															"send_community_val": {
																Type: schema.TypeString, Optional: true, Default: "both", Description: "'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;",
																ValidateFunc: validation.StringInSlice([]string{"both", "none", "standard", "extended"}, false),
															},
															"inbound": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow inbound soft reconfiguration for this neighbor",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"unsuppress_map": {
																Type: schema.TypeString, Optional: true, Description: "Route-map to selectively unsuppress suppressed routes (Name of route map)",
																ValidateFunc: validation.StringLenBetween(1, 128),
															},
															"weight": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set default weight for routes from this neighbor",
																ValidateFunc: validation.IntBetween(0, 65535),
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
																ValidateFunc: validation.IntAtLeast(1),
															},
															"peer_group_name": {
																Type: schema.TypeString, Optional: true, Description: "",
																ValidateFunc: validation.StringLenBetween(1, 128),
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
																ValidateFunc: validation.IntAtLeast(1),
															},
															"peer_group_name": {
																Type: schema.TypeString, Optional: true, Description: "",
																ValidateFunc: validation.StringLenBetween(1, 128),
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
																ValidateFunc: validation.IntAtLeast(1),
															},
															"peer_group_name": {
																Type: schema.TypeString, Optional: true, Description: "",
																ValidateFunc: validation.StringLenBetween(1, 128),
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
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
																ValidateFunc: validation.StringLenBetween(1, 128),
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
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
																ValidateFunc: validation.StringLenBetween(1, 128),
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
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
																ValidateFunc: validation.StringLenBetween(1, 128),
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
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
																ValidateFunc: validation.StringLenBetween(1, 128),
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
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
																ValidateFunc: validation.StringLenBetween(1, 128),
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
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
																ValidateFunc: validation.StringLenBetween(1, 128),
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
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
																ValidateFunc: validation.StringLenBetween(1, 128),
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
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
																ValidateFunc: validation.StringLenBetween(1, 128),
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
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
																ValidateFunc: validation.StringLenBetween(1, 128),
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
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
																ValidateFunc: validation.StringLenBetween(1, 128),
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
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
																ValidateFunc: validation.StringLenBetween(1, 128),
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
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
																ValidateFunc: validation.StringLenBetween(1, 128),
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
																			ValidateFunc: validation.IntBetween(0, 1),
																		},
																		"route_map": {
																			Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
																			ValidateFunc: validation.StringLenBetween(1, 128),
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
																			ValidateFunc: validation.IntBetween(0, 1),
																		},
																		"route_map": {
																			Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
																			ValidateFunc: validation.StringLenBetween(1, 128),
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
					},
				},
			},
			"aggregate_address_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aggregate_address": {
							Type: schema.TypeString, Optional: true, Description: "Configure BGP aggregate entries (Aggregate prefix)",
							ValidateFunc: validation.IsCIDR,
						},
						"as_set": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Generate AS set path information",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"summary_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Filter more specific routes from updates",
							ValidateFunc: validation.IntBetween(0, 1),
						},
					},
				},
			},
			"as_number": {
				Type: schema.TypeInt, Required: true, ForceNew: true, Description: "AS number",
				ValidateFunc: validation.IntBetween(1, 4294967295),
			},
			"auto_summary": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic network number summarization",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"bgp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"always_compare_med": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow comparing MED from different neighbors",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"bestpath_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ignore": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore as-path length in selecting a route",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"compare_routerid": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Compare router-id for identical EBGP paths",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"remove_recv_med": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "To remove rcvd MED attribute",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"remove_send_med": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "To remove send MED attribute",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"missing_as_worst": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Treat missing MED as the least preferred one",
										ValidateFunc: validation.IntBetween(0, 1),
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
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"dampening_half_time": {
										Type: schema.TypeInt, Optional: true, Description: "Reachability Half-life time for the penalty(minutes)",
										ValidateFunc: validation.IntBetween(1, 45),
									},
									"dampening_reuse": {
										Type: schema.TypeInt, Optional: true, Description: "Value to start reusing a route",
										ValidateFunc: validation.IntBetween(1, 20000),
									},
									"dampening_supress": {
										Type: schema.TypeInt, Optional: true, Description: "Value to start suppressing a route",
										ValidateFunc: validation.IntBetween(1, 20000),
									},
									"dampening_max_supress": {
										Type: schema.TypeInt, Optional: true, Description: "Maximum duration to suppress a stable route(minutes)",
										ValidateFunc: validation.IntBetween(1, 255),
									},
									"dampening_penalty": {
										Type: schema.TypeInt, Optional: true, Description: "Un-reachability Half-life time for the penalty(minutes)",
										ValidateFunc: validation.IntBetween(1, 45),
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria for dampening (Route-map name)",
										ValidateFunc: validation.StringLenBetween(1, 128),
									},
								},
							},
						},
						"local_preference_value": {
							Type: schema.TypeInt, Optional: true, Default: 100, Description: "Configure default local preference value",
							ValidateFunc: validation.IntBetween(0, 4294967295),
						},
						"deterministic_med": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Pick the best-MED path among paths advertised from the neighboring AS",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"enforce_first_as": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enforce the first AS for EBGP routes",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"fast_external_failover": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Immediately reset session if a link to a directly connected external peer goes down",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"log_neighbor_changes": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log neighbor up/down and reset reason",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"nexthop_trigger_count": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "BGP nexthop-tracking status (count)",
							ValidateFunc: validation.IntBetween(0, 127),
						},
						"router_id": {
							Type: schema.TypeString, Optional: true, Description: "Override current router identifier (peers will reset) (Manually configured router identifier)",
							ValidateFunc: validation.IsIPv4Address,
						},
						"override_validation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "override router-id validation",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"scan_time": {
							Type: schema.TypeInt, Optional: true, Default: 60, Description: "Configure background scan interval (Scan interval (sec) [Default:60 Disable:0])",
							ValidateFunc: validation.IntBetween(0, 60),
						},
						"graceful_restart": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure BGP BGP Graceful Restart",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"bgp_restart_time": {
							Type: schema.TypeInt, Optional: true, Default: 90, Description: "BGP Peer Graceful Restart time in seconds (default 90)",
							ValidateFunc: validation.IntBetween(1, 3600),
						},
						"bgp_stalepath_time": {
							Type: schema.TypeInt, Optional: true, Default: 360, Description: "BGP Graceful Restart Stalepath retention time in seconds (default 360)",
							ValidateFunc: validation.IntBetween(1, 3600),
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
							ValidateFunc: validation.IntBetween(1, 255),
						},
						"src_prefix": {
							Type: schema.TypeString, Optional: true, Description: "IP source prefix",
							ValidateFunc: validation.IsCIDR,
						},
						"acl_str": {
							Type: schema.TypeString, Optional: true, Description: "Access list name",
							ValidateFunc: validation.StringLenBetween(1, 128),
						},
						"ext_routes_dist": {
							Type: schema.TypeInt, Optional: true, Description: "Distance for routes external to the AS",
							ValidateFunc: validation.IntBetween(1, 255),
						},
						"int_routes_dist": {
							Type: schema.TypeInt, Optional: true, Description: "Distance for routes internal to the AS",
							ValidateFunc: validation.IntBetween(1, 255),
						},
						"local_routes_dist": {
							Type: schema.TypeInt, Optional: true, Description: "Distance for local routes",
							ValidateFunc: validation.IntBetween(1, 255),
						},
					},
				},
			},
			"maximum_paths_value": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Supported BGP multipath numbers",
				ValidateFunc: validation.IntBetween(1, 64),
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
										ValidateFunc: validation.StringLenBetween(1, 128),
									},
									"peer_group_key": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure peer-group",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"peer_group_remote_as": {
										Type: schema.TypeInt, Optional: true, Description: "Specify AS number of BGP neighbor",
										ValidateFunc: validation.IntBetween(1, 4294967295),
									},
									"activate": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable the Address Family for this Neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"advertisement_interval": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum interval between sending BGP routing updates (time in seconds)",
										ValidateFunc: validation.IntBetween(1, 600),
									},
									"allowas_in": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept as-path with my AS present in it",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"allowas_in_count": {
										Type: schema.TypeInt, Optional: true, Default: 3, Description: "Number of occurrences of AS number",
										ValidateFunc: validation.IntBetween(1, 10),
									},
									"as_origination_interval": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum interval between sending AS-origination routing updates (time in seconds)",
										ValidateFunc: validation.IntBetween(1, 600),
									},
									"dynamic": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertise dynamic capability to this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"prefix_list_direction": {
										Type: schema.TypeString, Optional: true, Description: "'both': both; 'receive': receive; 'send': send;",
										ValidateFunc: validation.StringInSlice([]string{"both", "receive", "send"}, false),
									},
									"route_refresh": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Advertise route-refresh capability to this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"extended_nexthop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertise extended-nexthop capability to this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"collide_established": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include Neighbor in Established State for Collision Detection",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"default_originate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Originate default route to this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria to originate default (route-map name)",
										ValidateFunc: validation.StringLenBetween(1, 128),
									},
									"description": {
										Type: schema.TypeString, Optional: true, Description: "Neighbor specific description (Up to 80 characters describing this neighbor)",
										ValidateFunc: validation.StringLenBetween(1, 80),
									},
									"disallow_infinite_holdtime": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "BGP per neighbor disallow-infinite-holdtime",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"distribute_lists": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"distribute_list": {
													Type: schema.TypeString, Optional: true, Description: "Filter updates to/from this neighbor (IP standard/extended/named access list)",
													ValidateFunc: validation.StringLenBetween(1, 128),
												},
												"distribute_list_direction": {
													Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
													ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
												},
											},
										},
									},
									"dont_capability_negotiate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not perform capability negotiation",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ebgp_multihop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow EBGP neighbors not on directly connected networks",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ebgp_multihop_hop_count": {
										Type: schema.TypeInt, Optional: true, Description: "maximum hop count",
										ValidateFunc: validation.IntBetween(1, 255),
									},
									"enforce_multihop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enforce EBGP neighbors to perform multihop",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"bfd": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"multihop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable multihop",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"neighbor_filter_lists": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"filter_list": {
													Type: schema.TypeString, Optional: true, Description: "Establish BGP filters (AS path access-list name)",
													ValidateFunc: validation.StringLenBetween(1, 128),
												},
												"filter_list_direction": {
													Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
													ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
												},
											},
										},
									},
									"maximum_prefix": {
										Type: schema.TypeInt, Optional: true, Description: "Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))",
									},
									"maximum_prefix_thres": {
										Type: schema.TypeInt, Optional: true, Description: "threshold-value, 1 to 100 percent",
										ValidateFunc: validation.IntBetween(1, 100),
									},
									"next_hop_self": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable the next hop calculation for this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"override_capability": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Override capability negotiation result",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"pass_value": {
										Type: schema.TypeString, Optional: true, Description: "Key String",
									},
									//omit encrypted field: 'pass_encrypted'
									"passive": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't send open messages to this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"neighbor_prefix_lists": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nbr_prefix_list": {
													Type: schema.TypeString, Optional: true, Description: "Filter updates to/from this neighbor (Name of a prefix list)",
													ValidateFunc: validation.StringLenBetween(1, 128),
												},
												"nbr_prefix_list_direction": {
													Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
													ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
												},
											},
										},
									},
									"remove_private_as": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove private AS number from outbound updates",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"neighbor_route_map_lists": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nbr_route_map": {
													Type: schema.TypeString, Optional: true, Description: "Apply route map to neighbor (Name of route map)",
													ValidateFunc: validation.StringLenBetween(1, 128),
												},
												"nbr_rmap_direction": {
													Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
													ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
												},
											},
										},
									},
									"send_community_val": {
										Type: schema.TypeString, Optional: true, Default: "both", Description: "'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;",
										ValidateFunc: validation.StringInSlice([]string{"both", "none", "standard", "extended"}, false),
									},
									"inbound": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow inbound soft reconfiguration for this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"shutdown": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Administratively shut down this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"strict_capability_match": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Strict capability negotiation match",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"timers_keepalive": {
										Type: schema.TypeInt, Optional: true, Default: 30, Description: "Keepalive interval",
										ValidateFunc: validation.IntBetween(0, 65535),
									},
									"timers_holdtime": {
										Type: schema.TypeInt, Optional: true, Default: 90, Description: "Holdtime",
										ValidateFunc: validation.IntBetween(0, 65535),
									},
									"connect": {
										Type: schema.TypeInt, Optional: true, Description: "BGP connect timer",
										ValidateFunc: validation.IntBetween(1, 65535),
									},
									"unsuppress_map": {
										Type: schema.TypeString, Optional: true, Description: "Route-map to selectively unsuppress suppressed routes (Name of route map)",
										ValidateFunc: validation.StringLenBetween(1, 128),
									},
									"update_source_ip": {
										Type: schema.TypeString, Optional: true, Description: "IP address",
										ValidateFunc:  validation.IsIPv4Address,
									},
									"update_source_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 address",
										ValidateFunc:  validation.IsIPv6Address,
									},
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
										ValidateFunc:  validation.IntAtLeast(1),
									},
									"loopback": {
										Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
										ValidateFunc:  validation.IntAtLeast(1),
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
										ValidateFunc:  validation.IntAtLeast(1),
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
										ValidateFunc:  validation.IntAtLeast(1),
									},
									"lif": {
										Type: schema.TypeString, Optional: true, Description: "Logical interface (Lif interface name)",
										ValidateFunc:  validation.StringLenBetween(1, 15),
									},
									"tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
										ValidateFunc:  validation.IntAtLeast(1),
									},
									"weight": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set default weight for routes from this neighbor",
										ValidateFunc: validation.IntBetween(0, 65535),
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
										ValidateFunc: validation.IsIPv4Address,
									},
									"nbr_remote_as": {
										Type: schema.TypeInt, Optional: true, Description: "Specify AS number of BGP neighbor",
										ValidateFunc: validation.IntBetween(1, 4294967295),
									},
									"peer_group_name": {
										Type: schema.TypeString, Optional: true, Description: "Configure peer-group (peer-group name)",
										ValidateFunc:  validation.StringLenBetween(1, 128),
									},
									"activate": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable the Address Family for this Neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"advertisement_interval": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum interval between sending BGP routing updates (time in seconds)",
										ValidateFunc:  validation.IntBetween(1, 600),
									},
									"allowas_in": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept as-path with my AS present in it",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"allowas_in_count": {
										Type: schema.TypeInt, Optional: true, Default: 3, Description: "Number of occurrences of AS number",
										ValidateFunc: validation.IntBetween(1, 10),
									},
									"as_origination_interval": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum interval between sending AS-origination routing updates (time in seconds)",
										ValidateFunc:  validation.IntBetween(1, 600),
									},
									"dynamic": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertise dynamic capability to this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"prefix_list_direction": {
										Type: schema.TypeString, Optional: true, Description: "'both': both; 'receive': receive; 'send': send;",
										ValidateFunc: validation.StringInSlice([]string{"both", "receive", "send"}, false),
									},
									"route_refresh": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Advertise route-refresh capability to this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"graceful_restart": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable graceful-restart helper for this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"collide_established": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include Neighbor in Established State for Collision Detection",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"default_originate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Originate default route to this neighbor",
										ValidateFunc:  validation.IntBetween(0, 1),
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria to originate default (route-map name)",
										ValidateFunc: validation.StringLenBetween(1, 128),
									},
									"description": {
										Type: schema.TypeString, Optional: true, Description: "Neighbor specific description (Up to 80 characters describing this neighbor)",
										ValidateFunc: validation.StringLenBetween(1, 80),
									},
									"disallow_infinite_holdtime": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "BGP per neighbor disallow-infinite-holdtime",
										ValidateFunc:  validation.IntBetween(0, 1),
									},
									"distribute_lists": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"distribute_list": {
													Type: schema.TypeString, Optional: true, Description: "Filter updates to/from this neighbor (IP standard/extended/named access list)",
													ValidateFunc: validation.StringLenBetween(1, 128),
												},
												"distribute_list_direction": {
													Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
													ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
												},
											},
										},
									},
									"acos_application_only": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send BGP update to ACOS application",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"dont_capability_negotiate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not perform capability negotiation",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ebgp_multihop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow EBGP neighbors not on directly connected networks",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ebgp_multihop_hop_count": {
										Type: schema.TypeInt, Optional: true, Description: "maximum hop count",
										ValidateFunc: validation.IntBetween(1, 255),
									},
									"enforce_multihop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enforce EBGP neighbors to perform multihop",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"bfd": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"multihop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable multihop",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"key_id": {
										Type: schema.TypeInt, Optional: true, Description: "Key ID",
										ValidateFunc: validation.IntBetween(0, 255),
									},
									"key_type": {
										Type: schema.TypeString, Optional: true, Description: "'md5': md5; 'meticulous-md5': meticulous-md5; 'meticulous-sha1': meticulous-sha1; 'sha1': sha1; 'simple': simple;  (Keyed MD5/Meticulous Keyed MD5/Meticulous Keyed SHA1/Keyed SHA1/Simple Password)",
										ValidateFunc: validation.StringInSlice([]string{"md5", "meticulous-md5", "meticulous-sha1", "sha1", "simple"}, false),
									},
									"bfd_value": {
										Type: schema.TypeString, Optional: true, Description: "Key String",
									},
									//omit encrypted field: 'bfd_encrypted'
									"neighbor_filter_lists": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"filter_list": {
													Type: schema.TypeString, Optional: true, Description: "Establish BGP filters (AS path access-list name)",
													ValidateFunc: validation.StringLenBetween(1, 128),
												},
												"filter_list_direction": {
													Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
													ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
												},
											},
										},
									},
									"maximum_prefix": {
										Type: schema.TypeInt, Optional: true, Description: "Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))",
									},
									"maximum_prefix_thres": {
										Type: schema.TypeInt, Optional: true, Description: "threshold-value, 1 to 100 percent",
										ValidateFunc: validation.IntBetween(1, 100),
									},
									"restart_min": {
										Type: schema.TypeInt, Optional: true, Description: "restart value, 1 to 1440 minutes",
										ValidateFunc: validation.IntBetween(1, 1440),
									},
									"next_hop_self": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable the next hop calculation for this neighbor",
										ValidateFunc:  validation.IntBetween(0, 1),
									},
									"override_capability": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Override capability negotiation result",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"pass_value": {
										Type: schema.TypeString, Optional: true, Description: "Key String",
									},
									//omit encrypted field: 'pass_encrypted'
									"passive": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't send open messages to this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"neighbor_prefix_lists": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nbr_prefix_list": {
													Type: schema.TypeString, Optional: true, Description: "Filter updates to/from this neighbor (Name of a prefix list)",
													ValidateFunc: validation.StringLenBetween(1, 128),
												},
												"nbr_prefix_list_direction": {
													Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
													ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
												},
											},
										},
									},
									"remove_private_as": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove private AS number from outbound updates",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"neighbor_route_map_lists": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nbr_route_map": {
													Type: schema.TypeString, Optional: true, Description: "Apply route map to neighbor (Name of route map)",
													ValidateFunc: validation.StringLenBetween(1, 128),
												},
												"nbr_rmap_direction": {
													Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
													ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
												},
											},
										},
									},
									"send_community_val": {
										Type: schema.TypeString, Optional: true, Default: "both", Description: "'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;",
										ValidateFunc: validation.StringInSlice([]string{"both", "none", "standard", "extended"}, false),
									},
									"inbound": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow inbound soft reconfiguration for this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"shutdown": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Administratively shut down this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"strict_capability_match": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Strict capability negotiation match",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"timers_keepalive": {
										Type: schema.TypeInt, Optional: true, Default: 30, Description: "Keepalive interval",
										ValidateFunc: validation.IntBetween(0, 65535),
									},
									"timers_holdtime": {
										Type: schema.TypeInt, Optional: true, Default: 90, Description: "Holdtime",
										ValidateFunc: validation.IntBetween(0, 65535),
									},
									"connect": {
										Type: schema.TypeInt, Optional: true, Description: "BGP connect timer",
										ValidateFunc: validation.IntBetween(1, 65535),
									},
									"unsuppress_map": {
										Type: schema.TypeString, Optional: true, Description: "Route-map to selectively unsuppress suppressed routes (Name of route map)",
										ValidateFunc:  validation.StringLenBetween(1, 128),
									},
									"update_source_ip": {
										Type: schema.TypeString, Optional: true, Description: "IP address",
										ValidateFunc:  validation.IsIPv4Address,
									},
									"update_source_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 address",
										ValidateFunc:  validation.IsIPv6Address,
									},
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
										ValidateFunc:  validation.IntAtLeast(1),
									},
									"loopback": {
										Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
										ValidateFunc:  validation.IntAtLeast(1),
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
										ValidateFunc:  validation.IntAtLeast(1),
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
										ValidateFunc:  validation.IntAtLeast(1),
									},
									"lif": {
										Type: schema.TypeString, Optional: true, Description: "Logical interface (Lif interface name)",
										ValidateFunc:  validation.StringLenBetween(1, 15),
									},
									"tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
										ValidateFunc:  validation.IntAtLeast(1),
									},
									"weight": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set default weight for routes from this neighbor",
										ValidateFunc: validation.IntBetween(0, 65535),
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
										ValidateFunc: validation.IsIPv6Address,
									},
									"nbr_remote_as": {
										Type: schema.TypeInt, Optional: true, Description: "Specify AS number of BGP neighbor",
										ValidateFunc: validation.IntBetween(1, 4294967295),
									},
									"peer_group_name": {
										Type: schema.TypeString, Optional: true, Description: "Configure peer-group (peer-group name)",
										ValidateFunc:  validation.StringLenBetween(1, 128),
									},
									"activate": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable the Address Family for this Neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"advertisement_interval": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum interval between sending BGP routing updates (time in seconds)",
										ValidateFunc:  validation.IntBetween(1, 600),
									},
									"allowas_in": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept as-path with my AS present in it",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"allowas_in_count": {
										Type: schema.TypeInt, Optional: true, Default: 3, Description: "Number of occurrences of AS number",
										ValidateFunc: validation.IntBetween(1, 10),
									},
									"as_origination_interval": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum interval between sending AS-origination routing updates (time in seconds)",
										ValidateFunc:  validation.IntBetween(1, 600),
									},
									"dynamic": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertise dynamic capability to this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"prefix_list_direction": {
										Type: schema.TypeString, Optional: true, Description: "'both': both; 'receive': receive; 'send': send;",
										ValidateFunc: validation.StringInSlice([]string{"both", "receive", "send"}, false),
									},
									"route_refresh": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Advertise route-refresh capability to this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"graceful_restart": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable graceful-restart helper for this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"extended_nexthop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertise extended-nexthop capability to this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"collide_established": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include Neighbor in Established State for Collision Detection",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"default_originate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Originate default route to this neighbor",
										ValidateFunc:  validation.IntBetween(0, 1),
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route-map to specify criteria to originate default (route-map name)",
										ValidateFunc: validation.StringLenBetween(1, 128),
									},
									"description": {
										Type: schema.TypeString, Optional: true, Description: "Neighbor specific description (Up to 80 characters describing this neighbor)",
										ValidateFunc: validation.StringLenBetween(1, 80),
									},
									"disallow_infinite_holdtime": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "BGP per neighbor disallow-infinite-holdtime",
										ValidateFunc:  validation.IntBetween(0, 1),
									},
									"distribute_lists": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"distribute_list": {
													Type: schema.TypeString, Optional: true, Description: "Filter updates to/from this neighbor (IP standard/extended/named access list)",
													ValidateFunc: validation.StringLenBetween(1, 128),
												},
												"distribute_list_direction": {
													Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
													ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
												},
											},
										},
									},
									"acos_application_only": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send BGP update to ACOS application",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"dont_capability_negotiate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not perform capability negotiation",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ebgp_multihop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow EBGP neighbors not on directly connected networks",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ebgp_multihop_hop_count": {
										Type: schema.TypeInt, Optional: true, Description: "maximum hop count",
										ValidateFunc: validation.IntBetween(1, 255),
									},
									"enforce_multihop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enforce EBGP neighbors to perform multihop",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"bfd": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"multihop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable multihop",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"key_id": {
										Type: schema.TypeInt, Optional: true, Description: "Key ID",
										ValidateFunc: validation.IntBetween(0, 255),
									},
									"key_type": {
										Type: schema.TypeString, Optional: true, Description: "'md5': md5; 'meticulous-md5': meticulous-md5; 'meticulous-sha1': meticulous-sha1; 'sha1': sha1; 'simple': simple;  (Keyed MD5/Meticulous Keyed MD5/Meticulous Keyed SHA1/Keyed SHA1/Simple Password)",
										ValidateFunc: validation.StringInSlice([]string{"md5", "meticulous-md5", "meticulous-sha1", "sha1", "simple"}, false),
									},
									"bfd_value": {
										Type: schema.TypeString, Optional: true, Description: "Key String",
									},
									//omit encrypted field: 'bfd_encrypted'
									"neighbor_filter_lists": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"filter_list": {
													Type: schema.TypeString, Optional: true, Description: "Establish BGP filters (AS path access-list name)",
													ValidateFunc: validation.StringLenBetween(1, 128),
												},
												"filter_list_direction": {
													Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
													ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
												},
											},
										},
									},
									"maximum_prefix": {
										Type: schema.TypeInt, Optional: true, Description: "Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))",
									},
									"maximum_prefix_thres": {
										Type: schema.TypeInt, Optional: true, Description: "threshold-value, 1 to 100 percent",
										ValidateFunc: validation.IntBetween(1, 100),
									},
									"restart_min": {
										Type: schema.TypeInt, Optional: true, Description: "restart value, 1 to 1440 minutes",
										ValidateFunc: validation.IntBetween(1, 1440),
									},
									"next_hop_self": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable the next hop calculation for this neighbor",
										ValidateFunc:  validation.IntBetween(0, 1),
									},
									"override_capability": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Override capability negotiation result",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"pass_value": {
										Type: schema.TypeString, Optional: true, Description: "Key String",
									},
									//omit encrypted field: 'pass_encrypted'
									"passive": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't send open messages to this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"neighbor_prefix_lists": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nbr_prefix_list": {
													Type: schema.TypeString, Optional: true, Description: "Filter updates to/from this neighbor (Name of a prefix list)",
													ValidateFunc: validation.StringLenBetween(1, 128),
												},
												"nbr_prefix_list_direction": {
													Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
													ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
												},
											},
										},
									},
									"remove_private_as": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove private AS number from outbound updates",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"neighbor_route_map_lists": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nbr_route_map": {
													Type: schema.TypeString, Optional: true, Description: "Apply route map to neighbor (Name of route map)",
													ValidateFunc: validation.StringLenBetween(1, 128),
												},
												"nbr_rmap_direction": {
													Type: schema.TypeString, Optional: true, Description: "'in': in; 'out': out;",
													ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
												},
											},
										},
									},
									"send_community_val": {
										Type: schema.TypeString, Optional: true, Default: "both", Description: "'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;",
										ValidateFunc: validation.StringInSlice([]string{"both", "none", "standard", "extended"}, false),
									},
									"inbound": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow inbound soft reconfiguration for this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"shutdown": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Administratively shut down this neighbor",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"strict_capability_match": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Strict capability negotiation match",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"timers_keepalive": {
										Type: schema.TypeInt, Optional: true, Default: 30, Description: "Keepalive interval",
										ValidateFunc: validation.IntBetween(0, 65535),
									},
									"timers_holdtime": {
										Type: schema.TypeInt, Optional: true, Default: 90, Description: "Holdtime",
										ValidateFunc: validation.IntBetween(0, 65535),
									},
									"connect": {
										Type: schema.TypeInt, Optional: true, Description: "BGP connect timer",
										ValidateFunc: validation.IntBetween(1, 65535),
									},
									"unsuppress_map": {
										Type: schema.TypeString, Optional: true, Description: "Route-map to selectively unsuppress suppressed routes (Name of route map)",
										ValidateFunc:  validation.StringLenBetween(1, 128),
									},
									"update_source_ip": {
										Type: schema.TypeString, Optional: true, Description: "IP address",
										ValidateFunc:  validation.IsIPv4Address,
									},
									"update_source_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 address",
										ValidateFunc:  validation.IsIPv6Address,
									},
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
										ValidateFunc:  validation.IntAtLeast(1),
									},
									"loopback": {
										Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
										ValidateFunc:  validation.IntAtLeast(1),
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
										ValidateFunc:  validation.IntAtLeast(1),
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
										ValidateFunc:  validation.IntAtLeast(1),
									},
									"lif": {
										Type: schema.TypeString, Optional: true, Description: "Logical interface (Lif interface name)",
										ValidateFunc:  validation.StringLenBetween(1, 15),
									},
									"tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
										ValidateFunc:  validation.IntAtLeast(1),
									},
									"weight": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set default weight for routes from this neighbor",
										ValidateFunc: validation.IntBetween(0, 65535),
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
										ValidateFunc: validation.IntAtLeast(1),
									},
									"unnumbered": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"peer_group_name": {
										Type: schema.TypeString, Optional: true, Description: "",
										ValidateFunc: validation.StringLenBetween(1, 128),
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
										ValidateFunc: validation.IntAtLeast(1),
									},
									"unnumbered": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"peer_group_name": {
										Type: schema.TypeString, Optional: true, Description: "",
										ValidateFunc: validation.StringLenBetween(1, 128),
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
										ValidateFunc: validation.IntAtLeast(1),
									},
									"unnumbered": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"peer_group_name": {
										Type: schema.TypeString, Optional: true, Description: "",
										ValidateFunc: validation.StringLenBetween(1, 128),
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
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
										ValidateFunc: validation.IsCIDR,
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route-map to modify the attributes (Name of the route map)",
										ValidateFunc: validation.StringLenBetween(1, 128),
									},
									"backdoor": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify a BGP backdoor route",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"description": {
										Type: schema.TypeString, Optional: true, Description: "Network specific description (Up to 80 characters describing this network)",
										ValidateFunc: validation.StringLenBetween(1, 80),
									},
									"comm_value": {
										Type: schema.TypeString, Optional: true, Description: "community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export",
										ValidateFunc: validation.StringLenBetween(1, 128),
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
				ValidateFunc: validation.IntBetween(0, 1),
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
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
										ValidateFunc: validation.StringLenBetween(1, 128),
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
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
										ValidateFunc: validation.StringLenBetween(1, 128),
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
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
										ValidateFunc: validation.StringLenBetween(1, 128),
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
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
										ValidateFunc: validation.StringLenBetween(1, 128),
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
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
										ValidateFunc: validation.StringLenBetween(1, 128),
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
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
										ValidateFunc: validation.StringLenBetween(1, 128),
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
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
										ValidateFunc: validation.StringLenBetween(1, 128),
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
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
										ValidateFunc: validation.StringLenBetween(1, 128),
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
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
										ValidateFunc: validation.StringLenBetween(1, 128),
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
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
										ValidateFunc: validation.StringLenBetween(1, 128),
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
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
										ValidateFunc: validation.StringLenBetween(1, 128),
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
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"route_map": {
													Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
													ValidateFunc: validation.StringLenBetween(1, 128),
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
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"route_map": {
													Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
													ValidateFunc: validation.StringLenBetween(1, 128),
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
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"timers": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bgp_keepalive": {
							Type: schema.TypeInt, Optional: true, Default: 30, Description: "Keepalive interval",
							ValidateFunc: validation.IntBetween(0, 65535),
						},
						"bgp_holdtime": {
							Type: schema.TypeInt, Optional: true, Default: 90, Description: "Holdtime",
							ValidateFunc: validation.IntBetween(0, 65535),
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
				ValidateFunc: validation.StringLenBetween(1, 127),
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

func getObjectRouterBgpAddressFamily(d []interface{}) edpt.RouterBgpAddressFamily {
	var ret edpt.RouterBgpAddressFamily
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Ipv6 = getObjectRouterBgpAddressFamilyIpv6(in["ipv6"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6(d []interface{}) edpt.RouterBgpAddressFamilyIpv6 {
	var ret edpt.RouterBgpAddressFamilyIpv6
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Bgp = getObjectRouterBgpAddressFamilyIpv6Bgp(in["bgp"].([]interface{}))
		ret.Distance = getObjectRouterBgpAddressFamilyIpv6Distance(in["distance"].([]interface{}))
		ret.MaximumPathsValue = in["maximum_paths_value"].(int)
		ret.Originate = in["originate"].(int)
		ret.AggregateAddressList = getSliceRouterBgpAddressFamilyIpv6AggregateAddressList(in["aggregate_address_list"].([]interface{}))
		ret.AutoSummary = in["auto_summary"].(int)
		ret.Synchronization = in["synchronization"].(int)
		//omit uuid
		ret.Network = getObjectRouterBgpAddressFamilyIpv6Network(in["network"].([]interface{}))
		ret.Neighbor = getObjectRouterBgpAddressFamilyIpv6Neighbor(in["neighbor"].([]interface{}))
		ret.Redistribute = getObjectRouterBgpAddressFamilyIpv6Redistribute(in["redistribute"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Bgp(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Bgp {
	var ret edpt.RouterBgpAddressFamilyIpv6Bgp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
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
	var ret edpt.RouterBgpAddressFamilyIpv6Distance
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.DistanceExt = in["distance_ext"].(int)
		ret.DistanceInt = in["distance_int"].(int)
		ret.DistanceLocal = in["distance_local"].(int)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6AggregateAddressList(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6AggregateAddressList {
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6AggregateAddressList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6AggregateAddressList
		oi.AggregateAddress = in["aggregate_address"].(string)
		oi.AsSet = in["as_set"].(int)
		oi.SummaryOnly = in["summary_only"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Network(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Network {
	var ret edpt.RouterBgpAddressFamilyIpv6Network
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Synchronization = getObjectRouterBgpAddressFamilyIpv6NetworkSynchronization(in["synchronization"].([]interface{}))
		ret.Ipv6NetworkList = getSliceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkList(in["ipv6_network_list"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6NetworkSynchronization(d []interface{}) edpt.RouterBgpAddressFamilyIpv6NetworkSynchronization {
	var ret edpt.RouterBgpAddressFamilyIpv6NetworkSynchronization
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.NetworkSynchronization = in["network_synchronization"].(int)
		//omit uuid
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkList(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList {
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList
		oi.NetworkIpv6 = in["network_ipv6"].(string)
		oi.RouteMap = in["route_map"].(string)
		oi.Backdoor = in["backdoor"].(int)
		oi.Description = in["description"].(string)
		oi.CommValue = in["comm_value"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Neighbor(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Neighbor {
	var ret edpt.RouterBgpAddressFamilyIpv6Neighbor
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
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
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList
		oi.PeerGroup = in["peer_group"].(string)
		oi.Activate = in["activate"].(int)
		oi.AllowasIn = in["allowas_in"].(int)
		oi.AllowasInCount = in["allowas_in_count"].(int)
		oi.PrefixListDirection = in["prefix_list_direction"].(string)
		oi.DefaultOriginate = in["default_originate"].(int)
		oi.RouteMap = in["route_map"].(string)
		oi.DistributeLists = getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListDistributeLists(in["distribute_lists"].([]interface{}))
		oi.NeighborFilterLists = getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborFilterLists(in["neighbor_filter_lists"].([]interface{}))
		oi.MaximumPrefix = in["maximum_prefix"].(int)
		oi.MaximumPrefixThres = in["maximum_prefix_thres"].(int)
		oi.NextHopSelf = in["next_hop_self"].(int)
		oi.NeighborPrefixLists = getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborPrefixLists(in["neighbor_prefix_lists"].([]interface{}))
		oi.RemovePrivateAs = in["remove_private_as"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists(in["neighbor_route_map_lists"].([]interface{}))
		oi.SendCommunityVal = in["send_community_val"].(string)
		oi.Inbound = in["inbound"].(int)
		oi.UnsuppressMap = in["unsuppress_map"].(string)
		oi.Weight = in["weight"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListDistributeLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListDistributeLists {
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListDistributeLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListDistributeLists
		oi.DistributeList = in["distribute_list"].(string)
		oi.DistributeListDirection = in["distribute_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborFilterLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborFilterLists {
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborFilterLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborFilterLists
		oi.FilterList = in["filter_list"].(string)
		oi.FilterListDirection = in["filter_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborPrefixLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborPrefixLists {
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborPrefixLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborPrefixLists
		oi.NbrPrefixList = in["nbr_prefix_list"].(string)
		oi.NbrPrefixListDirection = in["nbr_prefix_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists {
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborList(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList {
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
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
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists
		oi.DistributeList = in["distribute_list"].(string)
		oi.DistributeListDirection = in["distribute_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists {
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists
		oi.FilterList = in["filter_list"].(string)
		oi.FilterListDirection = in["filter_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists {
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists
		oi.NbrPrefixList = in["nbr_prefix_list"].(string)
		oi.NbrPrefixListDirection = in["nbr_prefix_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists {
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborList(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList {
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
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
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists
		oi.DistributeList = in["distribute_list"].(string)
		oi.DistributeListDirection = in["distribute_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists {
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists
		oi.FilterList = in["filter_list"].(string)
		oi.FilterListDirection = in["filter_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists {
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists
		oi.NbrPrefixList = in["nbr_prefix_list"].(string)
		oi.NbrPrefixListDirection = in["nbr_prefix_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists {
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List(d []interface{}) []edpt.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List {
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
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
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
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
	count := len(d)
	ret := make([]edpt.RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List
		oi.Trunk = in["trunk"].(int)
		oi.PeerGroupName = in["peer_group_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6Redistribute(d []interface{}) edpt.RouterBgpAddressFamilyIpv6Redistribute {
	var ret edpt.RouterBgpAddressFamilyIpv6Redistribute
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ConnectedCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeConnectedCfg(in["connected_cfg"].([]interface{}))
		ret.FloatingIpCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg(in["floating_ip_cfg"].([]interface{}))
		ret.Nat64Cfg = getObjectRouterBgpAddressFamilyIpv6RedistributeNat64Cfg(in["nat64_cfg"].([]interface{}))
		ret.NatMapCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeNatMapCfg(in["nat_map_cfg"].([]interface{}))
		ret.Lw4o6Cfg = getObjectRouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg(in["lw4o6_cfg"].([]interface{}))
		ret.StaticNatCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeStaticNatCfg(in["static_nat_cfg"].([]interface{}))
		ret.IpNatCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatCfg(in["ip_nat_cfg"].([]interface{}))
		ret.IpNatListCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatListCfg(in["ip_nat_list_cfg"].([]interface{}))
		ret.IsisCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeIsisCfg(in["isis_cfg"].([]interface{}))
		ret.OspfCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeOspfCfg(in["ospf_cfg"].([]interface{}))
		ret.RipCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeRipCfg(in["rip_cfg"].([]interface{}))
		ret.StaticCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeStaticCfg(in["static_cfg"].([]interface{}))
		ret.Vip = getObjectRouterBgpAddressFamilyIpv6RedistributeVip(in["vip"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeConnectedCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeConnectedCfg {
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeConnectedCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Connected = in["connected"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg {
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.FloatingIp = in["floating_ip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeNat64Cfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeNat64Cfg {
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeNat64Cfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Nat64 = in["nat64"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeNatMapCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeNatMapCfg {
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeNatMapCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.NatMap = in["nat_map"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg {
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Lw4o6 = in["lw4o6"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeStaticNatCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg {
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.StaticNat = in["static_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatCfg {
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.IpNat = in["ip_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatListCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg {
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.IpNatList = in["ip_nat_list"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeIsisCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeIsisCfg {
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeIsisCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Isis = in["isis"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeOspfCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeOspfCfg {
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeOspfCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Ospf = in["ospf"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeRipCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeRipCfg {
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeRipCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Rip = in["rip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeStaticCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeStaticCfg {
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeStaticCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Static = in["static"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeVip(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeVip {
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeVip
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.OnlyFlaggedCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg(in["only_flagged_cfg"].([]interface{}))
		ret.OnlyNotFlaggedCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg(in["only_not_flagged_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg {
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.OnlyFlagged = in["only_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg {
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.OnlyNotFlagged = in["only_not_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getSliceRouterBgpAggregateAddressList(d []interface{}) []edpt.RouterBgpAggregateAddressList {
	count := len(d)
	ret := make([]edpt.RouterBgpAggregateAddressList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
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
	var ret edpt.RouterBgpBgp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
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
		ret.OverrideValidation = in["override_validation"].(int)
		ret.ScanTime = in["scan_time"].(int)
		ret.GracefulRestart = in["graceful_restart"].(int)
		ret.BgpRestartTime = in["bgp_restart_time"].(int)
		ret.BgpStalepathTime = in["bgp_stalepath_time"].(int)
	}
	return ret
}

func getObjectRouterBgpBgpBestpathCfg(d []interface{}) edpt.RouterBgpBgpBestpathCfg {
	var ret edpt.RouterBgpBgpBestpathCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Ignore = in["ignore"].(int)
		ret.CompareRouterid = in["compare_routerid"].(int)
		ret.RemoveRecvMed = in["remove_recv_med"].(int)
		ret.RemoveSendMed = in["remove_send_med"].(int)
		ret.MissingAsWorst = in["missing_as_worst"].(int)
	}
	return ret
}

func getObjectRouterBgpBgpDampeningCfg(d []interface{}) edpt.RouterBgpBgpDampeningCfg {
	var ret edpt.RouterBgpBgpDampeningCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
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
	count := len(d)
	ret := make([]edpt.RouterBgpDistanceList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
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

func getObjectRouterBgpNeighbor(d []interface{}) edpt.RouterBgpNeighbor {
	var ret edpt.RouterBgpNeighbor
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
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
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborPeerGroupNeighborList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborPeerGroupNeighborList
		oi.PeerGroup = in["peer_group"].(string)
		oi.PeerGroupKey = in["peer_group_key"].(int)
		oi.PeerGroupRemoteAs = in["peer_group_remote_as"].(int)
		oi.Activate = in["activate"].(int)
		oi.AdvertisementInterval = in["advertisement_interval"].(int)
		oi.AllowasIn = in["allowas_in"].(int)
		oi.AllowasInCount = in["allowas_in_count"].(int)
		oi.AsOriginationInterval = in["as_origination_interval"].(int)
		oi.Dynamic = in["dynamic"].(int)
		oi.PrefixListDirection = in["prefix_list_direction"].(string)
		oi.RouteRefresh = in["route_refresh"].(int)
		oi.ExtendedNexthop = in["extended_nexthop"].(int)
		oi.CollideEstablished = in["collide_established"].(int)
		oi.DefaultOriginate = in["default_originate"].(int)
		oi.RouteMap = in["route_map"].(string)
		oi.Description = in["description"].(string)
		oi.DisallowInfiniteHoldtime = in["disallow_infinite_holdtime"].(int)
		oi.DistributeLists = getSliceRouterBgpNeighborPeerGroupNeighborListDistributeLists(in["distribute_lists"].([]interface{}))
		oi.DontCapabilityNegotiate = in["dont_capability_negotiate"].(int)
		oi.EbgpMultihop = in["ebgp_multihop"].(int)
		oi.EbgpMultihopHopCount = in["ebgp_multihop_hop_count"].(int)
		oi.EnforceMultihop = in["enforce_multihop"].(int)
		oi.Bfd = in["bfd"].(int)
		oi.Multihop = in["multihop"].(int)
		oi.NeighborFilterLists = getSliceRouterBgpNeighborPeerGroupNeighborListNeighborFilterLists(in["neighbor_filter_lists"].([]interface{}))
		oi.MaximumPrefix = in["maximum_prefix"].(int)
		oi.MaximumPrefixThres = in["maximum_prefix_thres"].(int)
		oi.NextHopSelf = in["next_hop_self"].(int)
		oi.OverrideCapability = in["override_capability"].(int)
		oi.PassValue = in["pass_value"].(string)
		//omit pass_encrypted
		oi.Passive = in["passive"].(int)
		oi.NeighborPrefixLists = getSliceRouterBgpNeighborPeerGroupNeighborListNeighborPrefixLists(in["neighbor_prefix_lists"].([]interface{}))
		oi.RemovePrivateAs = in["remove_private_as"].(int)
		oi.NeighborRouteMapLists = getSliceRouterBgpNeighborPeerGroupNeighborListNeighborRouteMapLists(in["neighbor_route_map_lists"].([]interface{}))
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

func getSliceRouterBgpNeighborPeerGroupNeighborListDistributeLists(d []interface{}) []edpt.RouterBgpNeighborPeerGroupNeighborListDistributeLists {
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborPeerGroupNeighborListDistributeLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborPeerGroupNeighborListDistributeLists
		oi.DistributeList = in["distribute_list"].(string)
		oi.DistributeListDirection = in["distribute_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborPeerGroupNeighborListNeighborFilterLists(d []interface{}) []edpt.RouterBgpNeighborPeerGroupNeighborListNeighborFilterLists {
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborPeerGroupNeighborListNeighborFilterLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborPeerGroupNeighborListNeighborFilterLists
		oi.FilterList = in["filter_list"].(string)
		oi.FilterListDirection = in["filter_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborPeerGroupNeighborListNeighborPrefixLists(d []interface{}) []edpt.RouterBgpNeighborPeerGroupNeighborListNeighborPrefixLists {
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborPeerGroupNeighborListNeighborPrefixLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborPeerGroupNeighborListNeighborPrefixLists
		oi.NbrPrefixList = in["nbr_prefix_list"].(string)
		oi.NbrPrefixListDirection = in["nbr_prefix_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborPeerGroupNeighborListNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpNeighborPeerGroupNeighborListNeighborRouteMapLists {
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborPeerGroupNeighborListNeighborRouteMapLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborPeerGroupNeighborListNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv4NeighborList(d []interface{}) []edpt.RouterBgpNeighborIpv4NeighborList {
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv4NeighborList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv4NeighborList
		oi.NeighborIpv4 = in["neighbor_ipv4"].(string)
		oi.NbrRemoteAs = in["nbr_remote_as"].(int)
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
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv4NeighborListDistributeLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv4NeighborListDistributeLists
		oi.DistributeList = in["distribute_list"].(string)
		oi.DistributeListDirection = in["distribute_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv4NeighborListNeighborFilterLists(d []interface{}) []edpt.RouterBgpNeighborIpv4NeighborListNeighborFilterLists {
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv4NeighborListNeighborFilterLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv4NeighborListNeighborFilterLists
		oi.FilterList = in["filter_list"].(string)
		oi.FilterListDirection = in["filter_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv4NeighborListNeighborPrefixLists(d []interface{}) []edpt.RouterBgpNeighborIpv4NeighborListNeighborPrefixLists {
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv4NeighborListNeighborPrefixLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv4NeighborListNeighborPrefixLists
		oi.NbrPrefixList = in["nbr_prefix_list"].(string)
		oi.NbrPrefixListDirection = in["nbr_prefix_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv4NeighborListNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpNeighborIpv4NeighborListNeighborRouteMapLists {
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv4NeighborListNeighborRouteMapLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv4NeighborListNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv6NeighborList(d []interface{}) []edpt.RouterBgpNeighborIpv6NeighborList {
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv6NeighborList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv6NeighborList
		oi.NeighborIpv6 = in["neighbor_ipv6"].(string)
		oi.NbrRemoteAs = in["nbr_remote_as"].(int)
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
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv6NeighborListDistributeLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv6NeighborListDistributeLists
		oi.DistributeList = in["distribute_list"].(string)
		oi.DistributeListDirection = in["distribute_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv6NeighborListNeighborFilterLists(d []interface{}) []edpt.RouterBgpNeighborIpv6NeighborListNeighborFilterLists {
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv6NeighborListNeighborFilterLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv6NeighborListNeighborFilterLists
		oi.FilterList = in["filter_list"].(string)
		oi.FilterListDirection = in["filter_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv6NeighborListNeighborPrefixLists(d []interface{}) []edpt.RouterBgpNeighborIpv6NeighborListNeighborPrefixLists {
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv6NeighborListNeighborPrefixLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv6NeighborListNeighborPrefixLists
		oi.NbrPrefixList = in["nbr_prefix_list"].(string)
		oi.NbrPrefixListDirection = in["nbr_prefix_list_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborIpv6NeighborListNeighborRouteMapLists(d []interface{}) []edpt.RouterBgpNeighborIpv6NeighborListNeighborRouteMapLists {
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborIpv6NeighborListNeighborRouteMapLists, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNeighborIpv6NeighborListNeighborRouteMapLists
		oi.NbrRouteMap = in["nbr_route_map"].(string)
		oi.NbrRmapDirection = in["nbr_rmap_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterBgpNeighborEthernetNeighborList(d []interface{}) []edpt.RouterBgpNeighborEthernetNeighborList {
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborEthernetNeighborList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
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
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborVeNeighborList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
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
	count := len(d)
	ret := make([]edpt.RouterBgpNeighborTrunkNeighborList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
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

func getObjectRouterBgpNetwork(d []interface{}) edpt.RouterBgpNetwork {
	var ret edpt.RouterBgpNetwork
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Synchronization = getObjectRouterBgpNetworkSynchronization(in["synchronization"].([]interface{}))
		ret.IpCidrList = getSliceRouterBgpNetworkIpCidrList(in["ip_cidr_list"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpNetworkSynchronization(d []interface{}) edpt.RouterBgpNetworkSynchronization {
	var ret edpt.RouterBgpNetworkSynchronization
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.NetworkSynchronization = in["network_synchronization"].(int)
		//omit uuid
	}
	return ret
}

func getSliceRouterBgpNetworkIpCidrList(d []interface{}) []edpt.RouterBgpNetworkIpCidrList {
	count := len(d)
	ret := make([]edpt.RouterBgpNetworkIpCidrList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.RouterBgpNetworkIpCidrList
		oi.NetworkIpv4Cidr = in["network_ipv4_cidr"].(string)
		oi.RouteMap = in["route_map"].(string)
		oi.Backdoor = in["backdoor"].(int)
		oi.Description = in["description"].(string)
		oi.CommValue = in["comm_value"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterBgpRedistribute(d []interface{}) edpt.RouterBgpRedistribute {
	var ret edpt.RouterBgpRedistribute
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ConnectedCfg = getObjectRouterBgpRedistributeConnectedCfg(in["connected_cfg"].([]interface{}))
		ret.FloatingIpCfg = getObjectRouterBgpRedistributeFloatingIpCfg(in["floating_ip_cfg"].([]interface{}))
		ret.Lw4o6Cfg = getObjectRouterBgpRedistributeLw4o6Cfg(in["lw4o6_cfg"].([]interface{}))
		ret.StaticNatCfg = getObjectRouterBgpRedistributeStaticNatCfg(in["static_nat_cfg"].([]interface{}))
		ret.IpNatCfg = getObjectRouterBgpRedistributeIpNatCfg(in["ip_nat_cfg"].([]interface{}))
		ret.IpNatListCfg = getObjectRouterBgpRedistributeIpNatListCfg(in["ip_nat_list_cfg"].([]interface{}))
		ret.IsisCfg = getObjectRouterBgpRedistributeIsisCfg(in["isis_cfg"].([]interface{}))
		ret.OspfCfg = getObjectRouterBgpRedistributeOspfCfg(in["ospf_cfg"].([]interface{}))
		ret.RipCfg = getObjectRouterBgpRedistributeRipCfg(in["rip_cfg"].([]interface{}))
		ret.StaticCfg = getObjectRouterBgpRedistributeStaticCfg(in["static_cfg"].([]interface{}))
		ret.NatMapCfg = getObjectRouterBgpRedistributeNatMapCfg(in["nat_map_cfg"].([]interface{}))
		ret.Vip = getObjectRouterBgpRedistributeVip(in["vip"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectRouterBgpRedistributeConnectedCfg(d []interface{}) edpt.RouterBgpRedistributeConnectedCfg {
	var ret edpt.RouterBgpRedistributeConnectedCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Connected = in["connected"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeFloatingIpCfg(d []interface{}) edpt.RouterBgpRedistributeFloatingIpCfg {
	var ret edpt.RouterBgpRedistributeFloatingIpCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.FloatingIp = in["floating_ip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeLw4o6Cfg(d []interface{}) edpt.RouterBgpRedistributeLw4o6Cfg {
	var ret edpt.RouterBgpRedistributeLw4o6Cfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Lw4o6 = in["lw4o6"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeStaticNatCfg(d []interface{}) edpt.RouterBgpRedistributeStaticNatCfg {
	var ret edpt.RouterBgpRedistributeStaticNatCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.StaticNat = in["static_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeIpNatCfg(d []interface{}) edpt.RouterBgpRedistributeIpNatCfg {
	var ret edpt.RouterBgpRedistributeIpNatCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.IpNat = in["ip_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeIpNatListCfg(d []interface{}) edpt.RouterBgpRedistributeIpNatListCfg {
	var ret edpt.RouterBgpRedistributeIpNatListCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.IpNatList = in["ip_nat_list"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeIsisCfg(d []interface{}) edpt.RouterBgpRedistributeIsisCfg {
	var ret edpt.RouterBgpRedistributeIsisCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Isis = in["isis"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeOspfCfg(d []interface{}) edpt.RouterBgpRedistributeOspfCfg {
	var ret edpt.RouterBgpRedistributeOspfCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Ospf = in["ospf"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeRipCfg(d []interface{}) edpt.RouterBgpRedistributeRipCfg {
	var ret edpt.RouterBgpRedistributeRipCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Rip = in["rip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeStaticCfg(d []interface{}) edpt.RouterBgpRedistributeStaticCfg {
	var ret edpt.RouterBgpRedistributeStaticCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Static = in["static"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeNatMapCfg(d []interface{}) edpt.RouterBgpRedistributeNatMapCfg {
	var ret edpt.RouterBgpRedistributeNatMapCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.NatMap = in["nat_map"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeVip(d []interface{}) edpt.RouterBgpRedistributeVip {
	var ret edpt.RouterBgpRedistributeVip
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.OnlyFlaggedCfg = getObjectRouterBgpRedistributeVipOnlyFlaggedCfg(in["only_flagged_cfg"].([]interface{}))
		ret.OnlyNotFlaggedCfg = getObjectRouterBgpRedistributeVipOnlyNotFlaggedCfg(in["only_not_flagged_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpRedistributeVipOnlyFlaggedCfg(d []interface{}) edpt.RouterBgpRedistributeVipOnlyFlaggedCfg {
	var ret edpt.RouterBgpRedistributeVipOnlyFlaggedCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.OnlyFlagged = in["only_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeVipOnlyNotFlaggedCfg(d []interface{}) edpt.RouterBgpRedistributeVipOnlyNotFlaggedCfg {
	var ret edpt.RouterBgpRedistributeVipOnlyNotFlaggedCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.OnlyNotFlagged = in["only_not_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpTimers(d []interface{}) edpt.RouterBgpTimers {
	var ret edpt.RouterBgpTimers
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.BgpKeepalive = in["bgp_keepalive"].(int)
		ret.BgpHoldtime = in["bgp_holdtime"].(int)
	}
	return ret
}

func dataToEndpointRouterBgp(d *schema.ResourceData) edpt.RouterBgp {
	var ret edpt.RouterBgp
	ret.Inst.AddressFamily = getObjectRouterBgpAddressFamily(d.Get("address_family").([]interface{}))
	ret.Inst.AggregateAddressList = getSliceRouterBgpAggregateAddressList(d.Get("aggregate_address_list").([]interface{}))
	ret.Inst.AsNumber = d.Get("as_number").(int)
	ret.Inst.AutoSummary = d.Get("auto_summary").(int)
	ret.Inst.Bgp = getObjectRouterBgpBgp(d.Get("bgp").([]interface{}))
	ret.Inst.DistanceList = getSliceRouterBgpDistanceList(d.Get("distance_list").([]interface{}))
	ret.Inst.MaximumPathsValue = d.Get("maximum_paths_value").(int)
	ret.Inst.Neighbor = getObjectRouterBgpNeighbor(d.Get("neighbor").([]interface{}))
	ret.Inst.Network = getObjectRouterBgpNetwork(d.Get("network").([]interface{}))
	ret.Inst.Originate = d.Get("originate").(int)
	ret.Inst.Redistribute = getObjectRouterBgpRedistribute(d.Get("redistribute").([]interface{}))
	ret.Inst.Synchronization = d.Get("synchronization").(int)
	ret.Inst.Timers = getObjectRouterBgpTimers(d.Get("timers").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
