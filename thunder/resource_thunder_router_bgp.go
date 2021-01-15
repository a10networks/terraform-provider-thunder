package thunder

//Thunder resource RouterBgp

import (
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceRouterBgp() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpCreate,
		Update: resourceRouterBgpUpdate,
		Read:   resourceRouterBgpRead,
		Delete: resourceRouterBgpDelete,
		Schema: map[string]*schema.Schema{
			"as_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"aggregate_address_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aggregate_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"as_set": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"summary_only": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"bgp": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"always_compare_med": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bestpath_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ignore": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"compare_routerid": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"remove_recv_med": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"remove_send_med": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"missing_as_worst": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"dampening_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dampening": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dampening_half_time": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dampening_reuse": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dampening_supress": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dampening_max_supress": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dampening_penalty": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"route_map": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"local_preference_value": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"deterministic_med": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"enforce_first_as": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"fast_external_failover": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"log_neighbor_changes": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"nexthop_trigger_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"router_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"override_validation": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"scan_time": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"graceful_restart": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bgp_restart_time": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bgp_stalepath_time": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"distance_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin_distance": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"src_prefix": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"acl_str": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ext_routes_dist": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"int_routes_dist": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"local_routes_dist": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"maximum_paths_value": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"originate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"timers": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bgp_keepalive": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bgp_holdtime": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"synchronization": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"auto_summary": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"network": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"synchronization": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_synchronization": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ip_cidr_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_ipv4_cidr": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"route_map": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"backdoor": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"description": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"comm_value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"neighbor": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"distance": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"distance_ext": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"distance_int": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"distance_local": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},

						"temp_val": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},

						"peer_group_neighbor_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
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
								},
							},
						},
						"ipv4_neighbor_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"neighbor_ipv4": {
										Type:        schema.TypeString,
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
									"graceful_restart": {
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
								},
							},
						},
						"ipv6_neighbor_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"neighbor_ipv6": {
										Type:        schema.TypeString,
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
									"graceful_restart": {
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
								},
							},
						},
						"ethernet_neighbor_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ethernet": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"unnumbered": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"peer_group_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ve_neighbor_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ve": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"unnumbered": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"peer_group_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"trunk_neighbor_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trunk": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"unnumbered": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"peer_group_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"redistribute": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connected_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connected": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"route_map": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"floating_ip_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"floating_ip": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"route_map": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"lw4o6_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"lw4o6": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"route_map": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"static_nat_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"static_nat": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"route_map": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ip_nat_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_nat": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"route_map": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ip_nat_list_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_nat_list": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"route_map": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"isis_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"isis": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"route_map": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ospf_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ospf": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"route_map": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"rip_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rip": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"route_map": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"static_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"static": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"route_map": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"nat_map_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"nat_map": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"route_map": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"vip": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"only_flagged_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"only_flagged": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"route_map": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"only_not_flagged_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"only_not_flagged": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"route_map": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
								},
							},
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"address_family": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bgp": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dampening": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"dampening_half": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"dampening_start_reuse": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"dampening_start_supress": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"dampening_max_supress": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"dampening_unreachability": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"route_map": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"distance": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"distance_ext": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"distance_int": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"distance_local": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"maximum_paths_value": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"originate": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"aggregate_address_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"aggregate_address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"as_set": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"summary_only": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"auto_summary": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"synchronization": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"network": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"synchronization": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"network_synchronization": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"uuid": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"ipv6_network_list": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"network_ipv6": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"route_map": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"backdoor": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"description": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"comm_value": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"uuid": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
											},
										},
									},
									"neighbor": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"distance": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"distance_ext": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"distance_int": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"distance_local": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"temp_val": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},

												"peer_group_neighbor_list": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
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
														},
													},
												},
												"ipv4_neighbor_list": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"neighbor_ipv4": {
																Type:        schema.TypeString,
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
															"graceful_restart": {
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
														},
													},
												},
												"ipv6_neighbor_list": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"neighbor_ipv6": {
																Type:        schema.TypeString,
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
															"graceful_restart": {
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
														},
													},
												},
												"ethernet_neighbor_ipv6_list": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ethernet": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"peer_group_name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"uuid": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"ve_neighbor_ipv6_list": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ve": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"peer_group_name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"uuid": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"trunk_neighbor_ipv6_list": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"trunk": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"peer_group_name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"uuid": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
											},
										},
									},
									"redistribute": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"connected_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"connected": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"route_map": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"floating_ip_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"floating_ip": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"route_map": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"nat64_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"nat64": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"route_map": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"nat_map_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"nat_map": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"route_map": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"lw4o6_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"lw4o6": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"route_map": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"static_nat_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static_nat": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"route_map": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"ip_nat_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_nat": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"route_map": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"ip_nat_list_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_nat_list": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"route_map": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"isis_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"isis": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"route_map": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"ospf_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ospf": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"route_map": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"rip_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"rip": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"route_map": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"static_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"route_map": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"vip": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"only_flagged_cfg": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"only_flagged": {
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "",
																		},
																		"route_map": {
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "",
																		},
																	},
																},
															},
															"only_not_flagged_cfg": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"only_not_flagged": {
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "",
																		},
																		"route_map": {
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "",
																		},
																	},
																},
															},
														},
													},
												},
												"uuid": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
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

			"process_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceRouterBgpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgp (Inside resourceRouterBgpCreate) ")
		name1 := d.Get("as_number").(int)
		data := dataToRouterBgp(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgp --")
		d.SetId(strconv.Itoa(name1))
		go_thunder.PostRouterBgp(client.Token, data, client.Host)

		return resourceRouterBgpRead(d, meta)

	}
	return nil
}

func resourceRouterBgpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterBgp (Inside resourceRouterBgpRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgp(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceRouterBgpUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying RouterBgp   (Inside resourceRouterBgpUpdate) ")
		data := dataToRouterBgp(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgp ")
		go_thunder.PutRouterBgp(client.Token, name1, data, client.Host)

		return resourceRouterBgpRead(d, meta)

	}
	return nil
}

func resourceRouterBgpDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceRouterBgpDelete) " + name1)
		err := go_thunder.DeleteRouterBgp(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToRouterBgp(d *schema.ResourceData) go_thunder.RouterBgp {
	var vc go_thunder.RouterBgp
	var c go_thunder.RouterBgpInstance
	c.AsNumber = d.Get("as_number").(int)

	AggregateAddressListCount := d.Get("aggregate_address_list.#").(int)
	c.AggregateAddress = make([]go_thunder.RouterBgpAggregateAddressList, 0, AggregateAddressListCount)

	for i := 0; i < AggregateAddressListCount; i++ {
		var obj1 go_thunder.RouterBgpAggregateAddressList
		prefix1 := fmt.Sprintf("aggregate_address_list.%d.", i)
		obj1.AggregateAddress = d.Get(prefix1 + "aggregate_address").(string)
		obj1.AsSet = d.Get(prefix1 + "as_set").(int)
		obj1.SummaryOnly = d.Get(prefix1 + "summary_only").(int)
		c.AggregateAddress = append(c.AggregateAddress, obj1)
	}

	var obj2 go_thunder.RouterBgpBgp
	prefix2 := "bgp.0."
	obj2.AlwaysCompareMed = d.Get(prefix2 + "always_compare_med").(int)

	var obj2_1 go_thunder.RouterBgpBestpathCfg
	prefix2_1 := prefix2 + "bestpath_cfg.0."
	obj2_1.Ignore = d.Get(prefix2_1 + "ignore").(int)
	obj2_1.CompareRouterid = d.Get(prefix2_1 + "compare_routerid").(int)
	obj2_1.RemoveRecvMed = d.Get(prefix2_1 + "remove_recv_med").(int)
	obj2_1.RemoveSendMed = d.Get(prefix2_1 + "remove_send_med").(int)
	obj2_1.MissingAsWorst = d.Get(prefix2_1 + "missing_as_worst").(int)

	obj2.Ignore = obj2_1

	var obj2_2 go_thunder.RouterBgpDampeningCfg
	prefix2_2 := prefix2 + "dampening_cfg.0."
	obj2_2.Dampening = d.Get(prefix2_2 + "dampening").(int)
	obj2_2.DampeningHalfTime = d.Get(prefix2_2 + "dampening_half_time").(int)
	obj2_2.DampeningReuse = d.Get(prefix2_2 + "dampening_reuse").(int)
	obj2_2.DampeningSupress = d.Get(prefix2_2 + "dampening_supress").(int)
	obj2_2.DampeningMaxSupress = d.Get(prefix2_2 + "dampening_max_supress").(int)
	obj2_2.DampeningPenalty = d.Get(prefix2_2 + "dampening_penalty").(int)
	obj2_2.RouteMap = d.Get(prefix2_2 + "route_map").(string)

	obj2.Dampening = obj2_2

	obj2.LocalPreferenceValue = d.Get(prefix2 + "local_preference_value").(int)
	obj2.DeterministicMed = d.Get(prefix2 + "deterministic_med").(int)
	obj2.EnforceFirstAs = d.Get(prefix2 + "enforce_first_as").(int)
	obj2.FastExternalFailover = d.Get(prefix2 + "fast_external_failover").(int)
	obj2.LogNeighborChanges = d.Get(prefix2 + "log_neighbor_changes").(int)
	obj2.NexthopTriggerCount = d.Get(prefix2 + "nexthop_trigger_count").(int)
	obj2.RouterID = d.Get(prefix2 + "router_id").(string)
	obj2.OverrideValidation = d.Get(prefix2 + "override_validation").(int)
	obj2.ScanTime = d.Get(prefix2 + "scan_time").(int)
	obj2.GracefulRestart = d.Get(prefix2 + "graceful_restart").(int)
	obj2.BgpRestartTime = d.Get(prefix2 + "bgp_restart_time").(int)
	obj2.BgpStalepathTime = d.Get(prefix2 + "bgp_stalepath_time").(int)

	c.AlwaysCompareMed = obj2

	DistanceListCount := d.Get("distance_list.#").(int)
	c.AdminDistance = make([]go_thunder.RouterBgpDistanceList, 0, DistanceListCount)

	for i := 0; i < DistanceListCount; i++ {
		var obj3 go_thunder.RouterBgpDistanceList
		prefix3 := fmt.Sprintf("distance_list.%d.", i)
		obj3.AdminDistance = d.Get(prefix3 + "admin_distance").(int)
		obj3.SrcPrefix = d.Get(prefix3 + "src_prefix").(string)
		obj3.AclStr = d.Get(prefix3 + "acl_str").(string)
		obj3.ExtRoutesDist = d.Get(prefix3 + "ext_routes_dist").(int)
		obj3.IntRoutesDist = d.Get(prefix3 + "int_routes_dist").(int)
		obj3.LocalRoutesDist = d.Get(prefix3 + "local_routes_dist").(int)
		c.AdminDistance = append(c.AdminDistance, obj3)
	}

	c.MaximumPathsValue = d.Get("maximum_paths_value").(int)
	c.Originate = d.Get("originate").(int)

	var obj4 go_thunder.RouterBgpTimers
	prefix4 := "timers.0."
	obj4.BgpKeepalive = d.Get(prefix4 + "bgp_keepalive").(int)
	obj4.BgpHoldtime = d.Get(prefix4 + "bgp_holdtime").(int)

	c.BgpKeepalive = obj4

	c.Synchronization = d.Get("synchronization").(int)
	c.AutoSummary = d.Get("auto_summary").(int)
	c.UserTag = d.Get("user_tag").(string)

	var obj5 go_thunder.RouterBgpNetwork
	prefix5 := "network.0."

	var obj5_1 go_thunder.RouterBgpSynchronization
	prefix5_1 := prefix5 + "synchronization.0."
	obj5_1.NetworkSynchronization = d.Get(prefix5_1 + "network_synchronization").(int)

	obj5.NetworkSynchronization = obj5_1

	IpCidrListCount := d.Get(prefix5 + "ip_cidr_list.#").(int)
	obj5.NetworkIpv4Cidr = make([]go_thunder.RouterBgpIPCidrList, 0, IpCidrListCount)

	for i := 0; i < IpCidrListCount; i++ {
		var obj5_2 go_thunder.RouterBgpIPCidrList
		prefix5_2 := prefix2 + fmt.Sprintf("ip_cidr_list.%d.", i)
		obj5_2.NetworkIpv4Cidr = d.Get(prefix5_2 + "network_ipv4_cidr").(string)
		obj5_2.RouteMap = d.Get(prefix5_2 + "route_map").(string)
		obj5_2.Backdoor = d.Get(prefix5_2 + "backdoor").(int)
		obj5_2.Description = d.Get(prefix5_2 + "description").(string)
		obj5_2.CommValue = d.Get(prefix5_2 + "comm_value").(string)
		obj5.NetworkIpv4Cidr = append(obj5.NetworkIpv4Cidr, obj5_2)
	}

	c.NetworkSynchronization = obj5

	var obj6 go_thunder.RouterBgpNeighbor
	prefix6 := "neighbor.0."

	//obj6.TempVal = d.Get(prefix6 + "temp_val").(string)

	var obj6_44 go_thunder.RouterBgpDistance
	prefix6_44 := prefix6 + "distance.0."
	obj6_44.DistanceExt = d.Get(prefix6_44 + "distance_ext").(int)
	obj6_44.DistanceInt = d.Get(prefix6_44 + "distance_int").(int)
	obj6_44.DistanceLocal = d.Get(prefix6_44 + "distance_local").(int)

	obj6.DistanceExt = obj6_44

	PeerGroupNeighborListCount := d.Get(prefix6 + "peer_group_neighbor_list.#").(int)
	obj6.PeerGroup = make([]go_thunder.RouterBgpPeerGroupNeighborList, 0, PeerGroupNeighborListCount)

	for i := 0; i < PeerGroupNeighborListCount; i++ {
		var obj6_1 go_thunder.RouterBgpPeerGroupNeighborList
		prefix6_1 := prefix6 + fmt.Sprintf("peer_group_neighbor_list.%d.", i)
		obj6_1.PeerGroup = d.Get(prefix6_1 + "peer_group").(string)
		obj6_1.Activate = d.Get(prefix6_1 + "activate").(int)
		obj6_1.AllowasIn = d.Get(prefix6_1 + "allowas_in").(int)
		obj6_1.AllowasInCount = d.Get(prefix6_1 + "allowas_in_count").(int)
		obj6_1.PrefixListDirection = d.Get(prefix6_1 + "prefix_list_direction").(string)
		obj6_1.DefaultOriginate = d.Get(prefix6_1 + "default_originate").(int)
		obj6_1.RouteMap = d.Get(prefix6_1 + "route_map").(string)

		DistributeListsCount := d.Get(prefix6_1 + "distribute_lists.#").(int)
		obj6_1.DistributeList = make([]go_thunder.RouterBgpDistributeLists, 0, DistributeListsCount)

		for i := 0; i < DistributeListsCount; i++ {
			var obj6_1_1 go_thunder.RouterBgpDistributeLists
			prefix6_1_1 := prefix6_1 + fmt.Sprintf("distribute_lists.%d.", i)
			obj6_1_1.DistributeList = d.Get(prefix6_1_1 + "distribute_list").(string)
			obj6_1_1.DistributeListDirection = d.Get(prefix6_1_1 + "distribute_list_direction").(string)
			obj6_1.DistributeList = append(obj6_1.DistributeList, obj6_1_1)
		}

		NeighborFilterListsCount := d.Get(prefix6_1 + "neighbor_filter_lists.#").(int)
		obj6_1.FilterList = make([]go_thunder.RouterBgpNeighborFilterLists, 0, NeighborFilterListsCount)

		for i := 0; i < NeighborFilterListsCount; i++ {
			var obj6_1_2 go_thunder.RouterBgpNeighborFilterLists
			prefix6_1_2 := prefix6_1 + fmt.Sprintf("neighbor_filter_lists.%d.", i)
			obj6_1_2.FilterList = d.Get(prefix6_1_2 + "filter_list").(string)
			obj6_1_2.FilterListDirection = d.Get(prefix6_1_2 + "filter_list_direction").(string)
			obj6_1.FilterList = append(obj6_1.FilterList, obj6_1_2)
		}

		obj6_1.MaximumPrefix = d.Get(prefix6_1 + "maximum_prefix").(int)
		obj6_1.MaximumPrefixThres = d.Get(prefix6_1 + "maximum_prefix_thres").(int)
		obj6_1.NextHopSelf = d.Get(prefix6_1 + "next_hop_self").(int)

		NeighborPrefixListsCount := d.Get(prefix6_1 + "neighbor_prefix_lists.#").(int)
		obj6_1.NbrPrefixList = make([]go_thunder.RouterBgpNeighborPrefixLists, 0, NeighborPrefixListsCount)

		for i := 0; i < NeighborPrefixListsCount; i++ {
			var obj6_1_3 go_thunder.RouterBgpNeighborPrefixLists
			prefix6_1_3 := prefix6_1 + fmt.Sprintf("neighbor_prefix_lists.%d.", i)
			obj6_1_3.NbrPrefixList = d.Get(prefix6_1_3 + "nbr_prefix_list").(string)
			obj6_1_3.NbrPrefixListDirection = d.Get(prefix6_1_3 + "nbr_prefix_list_direction").(string)
			obj6_1.NbrPrefixList = append(obj6_1.NbrPrefixList, obj6_1_3)
		}

		obj6_1.RemovePrivateAs = d.Get(prefix6_1 + "remove_private_as").(int)

		NeighborRouteMapListsCount := d.Get(prefix6_1 + "neighbor_route_map_lists.#").(int)
		obj6_1.NbrRouteMap = make([]go_thunder.RouterBgpNeighborRouteMapLists, 0, NeighborRouteMapListsCount)

		for i := 0; i < NeighborRouteMapListsCount; i++ {
			var obj6_1_4 go_thunder.RouterBgpNeighborRouteMapLists
			prefix6_1_4 := prefix6_1 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj6_1_4.NbrRouteMap = d.Get(prefix6_1_4 + "nbr_route_map").(string)
			obj6_1_4.NbrRmapDirection = d.Get(prefix6_1_4 + "nbr_rmap_direction").(string)
			obj6_1.NbrRouteMap = append(obj6_1.NbrRouteMap, obj6_1_4)
		}

		obj6_1.SendCommunityVal = d.Get(prefix6_1 + "send_community_val").(string)
		obj6_1.Inbound = d.Get(prefix6_1 + "inbound").(int)
		obj6_1.UnsuppressMap = d.Get(prefix6_1 + "unsuppress_map").(string)
		obj6_1.Weight = d.Get(prefix6_1 + "weight").(int)
		obj6.PeerGroup = append(obj6.PeerGroup, obj6_1)
	}

	Ipv4NeighborListCount := d.Get(prefix6 + "ipv4_neighbor_list.#").(int)
	obj6.NeighborIpv4 = make([]go_thunder.RouterBgpIpv4NeighborList, 0, Ipv4NeighborListCount)

	for i := 0; i < Ipv4NeighborListCount; i++ {
		var obj6_2 go_thunder.RouterBgpIpv4NeighborList
		prefix6_2 := prefix2 + fmt.Sprintf("ipv4_neighbor_list.%d.", i)
		obj6_2.NeighborIpv4 = d.Get(prefix6_2 + "neighbor_ipv4").(string)
		obj6_2.PeerGroupName = d.Get(prefix6_2 + "peer_group_name").(string)
		obj6_2.Activate = d.Get(prefix6_2 + "activate").(int)
		obj6_2.AllowasIn = d.Get(prefix6_2 + "allowas_in").(int)
		obj6_2.AllowasInCount = d.Get(prefix6_2 + "allowas_in_count").(int)
		obj6_2.PrefixListDirection = d.Get(prefix6_2 + "prefix_list_direction").(string)
		obj6_2.GracefulRestart = d.Get(prefix6_2 + "graceful_restart").(int)
		obj6_2.DefaultOriginate = d.Get(prefix6_2 + "default_originate").(int)
		obj6_2.RouteMap = d.Get(prefix6_2 + "route_map").(string)

		DistributeListsCount := d.Get(prefix6_2 + "distribute_lists.#").(int)
		obj6_2.DistributeList = make([]go_thunder.RouterBgpDistributeLists, 0, DistributeListsCount)

		for i := 0; i < DistributeListsCount; i++ {
			var obj6_2_1 go_thunder.RouterBgpDistributeLists
			prefix6_2_1 := prefix6_2 + fmt.Sprintf("distribute_lists.%d.", i)
			obj6_2_1.DistributeList = d.Get(prefix6_2_1 + "distribute_list").(string)
			obj6_2_1.DistributeListDirection = d.Get(prefix6_2_1 + "distribute_list_direction").(string)
			obj6_2.DistributeList = append(obj6_2.DistributeList, obj6_2_1)
		}

		NeighborFilterListsCount := d.Get(prefix6_2 + "neighbor_filter_lists.#").(int)
		obj6_2.FilterList = make([]go_thunder.RouterBgpNeighborFilterLists, 0, NeighborFilterListsCount)

		for i := 0; i < NeighborFilterListsCount; i++ {
			var obj6_2_2 go_thunder.RouterBgpNeighborFilterLists
			prefix6_2_2 := prefix6_2 + fmt.Sprintf("neighbor_filter_lists.%d.", i)
			obj6_2_2.FilterList = d.Get(prefix6_2_2 + "filter_list").(string)
			obj6_2_2.FilterListDirection = d.Get(prefix6_2_2 + "filter_list_direction").(string)
			obj6_2.FilterList = append(obj6_2.FilterList, obj6_2_2)
		}

		obj6_2.MaximumPrefix = d.Get(prefix6_2 + "maximum_prefix").(int)
		obj6_2.MaximumPrefixThres = d.Get(prefix6_2 + "maximum_prefix_thres").(int)
		obj6_2.NextHopSelf = d.Get(prefix6_2 + "next_hop_self").(int)

		NeighborPrefixListsCount := d.Get(prefix6_2 + "neighbor_prefix_lists.#").(int)
		obj6_2.NbrPrefixList = make([]go_thunder.RouterBgpNeighborPrefixLists, 0, NeighborPrefixListsCount)

		for i := 0; i < NeighborPrefixListsCount; i++ {
			var obj6_2_3 go_thunder.RouterBgpNeighborPrefixLists
			prefix6_2_3 := prefix6_2 + fmt.Sprintf("neighbor_prefix_lists.%d.", i)
			obj6_2_3.NbrPrefixList = d.Get(prefix6_2_3 + "nbr_prefix_list").(string)
			obj6_2_3.NbrPrefixListDirection = d.Get(prefix6_2_3 + "nbr_prefix_list_direction").(string)
			obj6_2.NbrPrefixList = append(obj6_2.NbrPrefixList, obj6_2_3)
		}

		obj6_2.RemovePrivateAs = d.Get(prefix6_2 + "remove_private_as").(int)

		NeighborRouteMapListsCount := d.Get(prefix6_2 + "neighbor_route_map_lists.#").(int)
		obj6_2.NbrRouteMap = make([]go_thunder.RouterBgpNeighborRouteMapLists, 0, NeighborRouteMapListsCount)

		for i := 0; i < NeighborRouteMapListsCount; i++ {
			var obj6_2_4 go_thunder.RouterBgpNeighborRouteMapLists
			prefix6_2_4 := prefix6_2 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj6_2_4.NbrRouteMap = d.Get(prefix6_2_4 + "nbr_route_map").(string)
			obj6_2_4.NbrRmapDirection = d.Get(prefix6_2_4 + "nbr_rmap_direction").(string)
			obj6_2.NbrRouteMap = append(obj6_2.NbrRouteMap, obj6_2_4)
		}

		obj6_2.SendCommunityVal = d.Get(prefix6_2 + "send_community_val").(string)
		obj6_2.Inbound = d.Get(prefix6_2 + "inbound").(int)
		obj6_2.UnsuppressMap = d.Get(prefix6_2 + "unsuppress_map").(string)
		obj6_2.Weight = d.Get(prefix6_2 + "weight").(int)
		obj6.NeighborIpv4 = append(obj6.NeighborIpv4, obj6_2)
	}

	Ipv6NeighborListCount := d.Get(prefix6 + "ipv6_neighbor_list.#").(int)
	obj6.NeighborIpv6 = make([]go_thunder.RouterBgpIpv6NeighborList, 0, Ipv6NeighborListCount)

	for i := 0; i < Ipv6NeighborListCount; i++ {
		var obj6_3 go_thunder.RouterBgpIpv6NeighborList
		prefix6_3 := prefix6 + fmt.Sprintf("ipv6_neighbor_list.%d.", i)
		obj6_3.NeighborIpv6 = d.Get(prefix6_3 + "neighbor_ipv6").(string)
		obj6_3.PeerGroupName = d.Get(prefix6_3 + "peer_group_name").(string)
		obj6_3.Activate = d.Get(prefix6_3 + "activate").(int)
		obj6_3.AllowasIn = d.Get(prefix6_3 + "allowas_in").(int)
		obj6_3.AllowasInCount = d.Get(prefix6_3 + "allowas_in_count").(int)
		obj6_3.PrefixListDirection = d.Get(prefix6_3 + "prefix_list_direction").(string)
		obj6_3.GracefulRestart = d.Get(prefix6_3 + "graceful_restart").(int)
		obj6_3.DefaultOriginate = d.Get(prefix6_3 + "default_originate").(int)
		obj6_3.RouteMap = d.Get(prefix6_3 + "route_map").(string)

		DistributeListsCount := d.Get(prefix6_3 + "distribute_lists.#").(int)
		obj6_3.DistributeList = make([]go_thunder.RouterBgpDistributeLists, 0, DistributeListsCount)

		for i := 0; i < DistributeListsCount; i++ {
			var obj6_3_1 go_thunder.RouterBgpDistributeLists
			prefix6_3_1 := prefix6_3 + fmt.Sprintf("distribute_lists.%d.", i)
			obj6_3_1.DistributeList = d.Get(prefix6_3_1 + "distribute_list").(string)
			obj6_3_1.DistributeListDirection = d.Get(prefix6_3_1 + "distribute_list_direction").(string)
			obj6_3.DistributeList = append(obj6_3.DistributeList, obj6_3_1)
		}

		NeighborFilterListsCount := d.Get(prefix6_3 + "neighbor_filter_lists.#").(int)
		obj6_3.FilterList = make([]go_thunder.RouterBgpNeighborFilterLists, 0, NeighborFilterListsCount)

		for i := 0; i < NeighborFilterListsCount; i++ {
			var obj6_3_2 go_thunder.RouterBgpNeighborFilterLists
			prefix6_3_2 := prefix6_3 + fmt.Sprintf("neighbor_filter_lists.%d.", i)
			obj6_3_2.FilterList = d.Get(prefix6_3_2 + "filter_list").(string)
			obj6_3_2.FilterListDirection = d.Get(prefix6_3_2 + "filter_list_direction").(string)
			obj6_3.FilterList = append(obj6_3.FilterList, obj6_3_2)
		}

		obj6_3.MaximumPrefix = d.Get(prefix6_3 + "maximum_prefix").(int)
		obj6_3.MaximumPrefixThres = d.Get(prefix6_3 + "maximum_prefix_thres").(int)
		obj6_3.NextHopSelf = d.Get(prefix6_3 + "next_hop_self").(int)

		NeighborPrefixListsCount := d.Get(prefix6_3 + "neighbor_prefix_lists.#").(int)
		obj6_3.NbrPrefixList = make([]go_thunder.RouterBgpNeighborPrefixLists, 0, NeighborPrefixListsCount)

		for i := 0; i < NeighborPrefixListsCount; i++ {
			var obj6_3_3 go_thunder.RouterBgpNeighborPrefixLists
			prefix6_3_3 := prefix6_3 + fmt.Sprintf("neighbor_prefix_lists.%d.", i)
			obj6_3_3.NbrPrefixList = d.Get(prefix6_3_3 + "nbr_prefix_list").(string)
			obj6_3_3.NbrPrefixListDirection = d.Get(prefix6_3_3 + "nbr_prefix_list_direction").(string)
			obj6_3.NbrPrefixList = append(obj6_3.NbrPrefixList, obj6_3_3)
		}

		obj6_3.RemovePrivateAs = d.Get(prefix6_3 + "remove_private_as").(int)

		NeighborRouteMapListsCount := d.Get(prefix6_3 + "neighbor_route_map_lists.#").(int)
		obj6_3.NbrRouteMap = make([]go_thunder.RouterBgpNeighborRouteMapLists, 0, NeighborRouteMapListsCount)

		for i := 0; i < NeighborRouteMapListsCount; i++ {
			var obj6_3_4 go_thunder.RouterBgpNeighborRouteMapLists
			prefix6_3_4 := prefix6_3 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj6_3_4.NbrRouteMap = d.Get(prefix6_3_4 + "nbr_route_map").(string)
			obj6_3_4.NbrRmapDirection = d.Get(prefix6_3_4 + "nbr_rmap_direction").(string)
			obj6_3.NbrRouteMap = append(obj6_3.NbrRouteMap, obj6_3_4)
		}

		obj6_3.SendCommunityVal = d.Get(prefix6_3 + "send_community_val").(string)
		obj6_3.Inbound = d.Get(prefix6_3 + "inbound").(int)
		obj6_3.UnsuppressMap = d.Get(prefix6_3 + "unsuppress_map").(string)
		obj6_3.Weight = d.Get(prefix6_3 + "weight").(int)
		obj6.NeighborIpv6 = append(obj6.NeighborIpv6, obj6_3)
	}

	EthernetNeighborListCount := d.Get(prefix6 + "ethernet_neighbor_list.#").(int)
	obj6.Ethernet = make([]go_thunder.RouterBgpEthernetNeighborList, 0, EthernetNeighborListCount)

	for i := 0; i < EthernetNeighborListCount; i++ {
		var obj6_4 go_thunder.RouterBgpEthernetNeighborList
		prefix6_4 := prefix6 + fmt.Sprintf("ethernet_neighbor_list.%d.", i)
		obj6_4.Ethernet = d.Get(prefix6_4 + "ethernet").(int)
		obj6_4.Unnumbered = d.Get(prefix6_4 + "unnumbered").(int)
		obj6_4.PeerGroupName = d.Get(prefix6_4 + "peer_group_name").(string)
		obj6.Ethernet = append(obj6.Ethernet, obj6_4)
	}

	VeNeighborListCount := d.Get(prefix6 + "ve_neighbor_list.#").(int)
	obj6.Ve = make([]go_thunder.RouterBgpVeNeighborList, 0, VeNeighborListCount)

	for i := 0; i < VeNeighborListCount; i++ {
		var obj6_5 go_thunder.RouterBgpVeNeighborList
		prefix6_5 := prefix6 + fmt.Sprintf("ve_neighbor_list.%d.", i)
		obj6_5.Ve = d.Get(prefix6_5 + "ve").(int)
		obj6_5.Unnumbered = d.Get(prefix6_5 + "unnumbered").(int)
		obj6_5.PeerGroupName = d.Get(prefix6_5 + "peer_group_name").(string)
		obj6.Ve = append(obj6.Ve, obj6_5)
	}

	TrunkNeighborListCount := d.Get(prefix6 + "trunk_neighbor_list.#").(int)
	obj6.Trunk = make([]go_thunder.RouterBgpTrunkNeighborList, 0, TrunkNeighborListCount)

	for i := 0; i < TrunkNeighborListCount; i++ {
		var obj6_6 go_thunder.RouterBgpTrunkNeighborList
		prefix6_6 := prefix6 + fmt.Sprintf("trunk_neighbor_list.%d.", i)
		obj6_6.Trunk = d.Get(prefix6_6 + "trunk").(int)
		obj6_6.Unnumbered = d.Get(prefix6_6 + "unnumbered").(int)
		obj6_6.PeerGroupName = d.Get(prefix6_6 + "peer_group_name").(string)
		obj6.Trunk = append(obj6.Trunk, obj6_6)
	}

	c.Ethernet = obj6

	var obj7 go_thunder.RouterBgpRedistributeBgp
	prefix7 := "redistribute.0."

	var obj7_1 go_thunder.RouterBgpConnectedCfg
	prefix7_1 := prefix7 + "connected_cfg.0."
	obj7_1.Connected = d.Get(prefix7_1 + "connected").(int)
	obj7_1.RouteMap = d.Get(prefix7_1 + "route_map").(string)

	obj7.Connected = obj7_1

	var obj7_2 go_thunder.RouterBgpFloatingIPCfg
	prefix7_2 := prefix7 + "floating_ip_cfg.0."
	obj7_2.FloatingIP = d.Get(prefix7_2 + "floating_ip").(int)
	obj7_2.RouteMap = d.Get(prefix7_2 + "route_map").(string)

	obj7.FloatingIP = obj7_2

	var obj7_3 go_thunder.RouterBgpLw4O6Cfg
	prefix7_3 := prefix7 + "lw4o6_cfg.0."
	obj7_3.Lw4O6 = d.Get(prefix7_3 + "lw4o6").(int)
	obj7_3.RouteMap = d.Get(prefix7_3 + "route_map").(string)

	obj7.Lw4O6 = obj7_3

	var obj7_4 go_thunder.RouterBgpStaticNatCfg
	prefix7_4 := prefix7 + "static_nat_cfg.0."
	obj7_4.StaticNat = d.Get(prefix7_4 + "static_nat").(int)
	obj7_4.RouteMap = d.Get(prefix7_4 + "route_map").(string)

	obj7.StaticNat = obj7_4

	var obj7_5 go_thunder.RouterBgpIPNatCfg
	prefix7_5 := prefix7 + "ip_nat_cfg.0."
	obj7_5.IPNat = d.Get(prefix7_5 + "ip_nat").(int)
	obj7_5.RouteMap = d.Get(prefix7_5 + "route_map").(string)

	obj7.IPNat = obj7_5

	var obj7_6 go_thunder.RouterBgpIPNatListCfg
	prefix7_6 := prefix7 + "ip_nat_list_cfg.0."
	obj7_6.IPNatList = d.Get(prefix7_6 + "ip_nat_list").(int)
	obj7_6.RouteMap = d.Get(prefix7_6 + "route_map").(string)

	obj7.IPNatList = obj7_6

	var obj7_7 go_thunder.RouterBgpIsisCfg
	prefix7_7 := prefix7 + "isis_cfg.0."
	obj7_7.Isis = d.Get(prefix7_7 + "isis").(int)
	obj7_7.RouteMap = d.Get(prefix7_7 + "route_map").(string)

	obj7.Isis = obj7_7

	var obj7_8 go_thunder.RouterBgpOspfCfg
	prefix7_8 := prefix7 + "ospf_cfg.0."
	obj7_8.Ospf = d.Get(prefix7_8 + "ospf").(int)
	obj7_8.RouteMap = d.Get(prefix7_8 + "route_map").(string)

	obj7.Ospf = obj7_8

	var obj7_9 go_thunder.RouterBgpRipCfg
	prefix7_9 := prefix7 + "rip_cfg.0."
	obj7_9.Rip = d.Get(prefix7_9 + "rip").(int)
	obj7_9.RouteMap = d.Get(prefix7_9 + "route_map").(string)

	obj7.Rip = obj7_9

	var obj7_10 go_thunder.RouterBgpStaticCfg
	prefix7_10 := prefix7 + "static_cfg.0."
	obj7_10.Static = d.Get(prefix7_10 + "static").(int)
	obj7_10.RouteMap = d.Get(prefix7_10 + "route_map").(string)

	obj7.Static = obj7_10

	var obj7_11 go_thunder.RouterBgpNatMapCfg
	prefix7_11 := prefix7 + "nat_map_cfg.0."
	obj7_11.NatMap = d.Get(prefix7_11 + "nat_map").(int)
	obj7_11.RouteMap = d.Get(prefix7_11 + "route_map").(string)

	obj7.NatMap = obj7_11

	var obj7_12 go_thunder.RouterBgpVip
	prefix7_12 := prefix7 + "vip.0."

	var obj7_12_1 go_thunder.RouterBgpOnlyFlaggedCfg
	prefix7_12_1 := prefix7_12 + "only_flagged_cfg.0."
	obj7_12_1.OnlyFlagged = d.Get(prefix7_12_1 + "only_flagged").(int)
	obj7_12_1.RouteMap = d.Get(prefix7_12_1 + "route_map").(string)

	obj7_12.OnlyFlagged = obj7_12_1

	var obj7_12_2 go_thunder.RouterBgpOnlyNotFlaggedCfg
	prefix7_12_2 := prefix7_12 + "only_not_flagged_cfg.0."
	obj7_12_2.OnlyNotFlagged = d.Get(prefix7_12_2 + "only_not_flagged").(int)
	obj7_12_2.RouteMap = d.Get(prefix7_12_2 + "route_map").(string)

	obj7_12.OnlyNotFlagged = obj7_12_2

	obj7.OnlyFlaggedCfg = obj7_12

	c.ConnectedCfg = obj7

	var obj8 go_thunder.RouterBgpAddressFamily
	prefix8 := "address_family.0."

	var obj8_1 go_thunder.RouterBgpIpv6
	prefix8_1 := prefix8 + "ipv6.0."

	var obj8_1_1 go_thunder.RouterBgpBgp1
	prefix8_1_1 := prefix8_1 + "bgp.0."
	obj8_1_1.Dampening = d.Get(prefix8_1_1 + "dampening").(int)
	obj8_1_1.DampeningHalf = d.Get(prefix8_1_1 + "dampening_half").(int)
	obj8_1_1.DampeningStartReuse = d.Get(prefix8_1_1 + "dampening_start_reuse").(int)
	obj8_1_1.DampeningStartSupress = d.Get(prefix8_1_1 + "dampening_start_supress").(int)
	obj8_1_1.DampeningMaxSupress = d.Get(prefix8_1_1 + "dampening_max_supress").(int)
	obj8_1_1.DampeningUnreachability = d.Get(prefix8_1_1 + "dampening_unreachability").(int)
	obj8_1_1.RouteMap = d.Get(prefix8_1_1 + "route_map").(string)

	obj8_1.Dampening = obj8_1_1

	var obj8_1_2 go_thunder.RouterBgpDistance
	prefix8_1_2 := prefix8_1 + "distance.0."
	obj8_1_2.DistanceExt = d.Get(prefix8_1_2 + "distance_ext").(int)
	obj8_1_2.DistanceInt = d.Get(prefix8_1_2 + "distance_int").(int)
	obj8_1_2.DistanceLocal = d.Get(prefix8_1_2 + "distance_local").(int)

	obj8_1.DistanceExt = obj8_1_2

	obj8_1.MaximumPathsValue = d.Get(prefix8_1 + "maximum_paths_value").(int)
	obj8_1.Originate = d.Get(prefix8_1 + "originate").(int)

	AggregateAddressListCount = d.Get(prefix8_1 + "aggregate_address_list.#").(int)
	obj8_1.AggregateAddress = make([]go_thunder.RouterBgpAggregateAddressList, 0, AggregateAddressListCount)

	for i := 0; i < AggregateAddressListCount; i++ {
		var obj8_1_3 go_thunder.RouterBgpAggregateAddressList
		prefix8_1_3 := prefix8_1 + fmt.Sprintf("aggregate_address_list.%d.", i)
		obj8_1_3.AggregateAddress = d.Get(prefix8_1_3 + "aggregate_address").(string)
		obj8_1_3.AsSet = d.Get(prefix8_1_3 + "as_set").(int)
		obj8_1_3.SummaryOnly = d.Get(prefix8_1_3 + "summary_only").(int)
		obj8_1.AggregateAddress = append(obj8_1.AggregateAddress, obj8_1_3)
	}

	obj8_1.AutoSummary = d.Get(prefix8_1 + "auto_summary").(int)
	obj8_1.Synchronization = d.Get(prefix8_1 + "synchronization").(int)

	var obj8_1_4 go_thunder.RouterBgpNetworkIpv6
	prefix8_1_4 := prefix8_1 + "network.0."

	var obj8_1_4_1 go_thunder.RouterBgpSynchronization
	prefix8_1_4_1 := prefix8_1_4 + "synchronization.0."
	obj8_1_4_1.NetworkSynchronization = d.Get(prefix8_1_4_1 + "network_synchronization").(int)

	obj8_1_4.NetworkSynchronization = obj8_1_4_1

	Ipv6NetworkListCount := d.Get(prefix8_1_4 + "ipv6_network_list.#").(int)
	obj8_1_4.NetworkIpv6 = make([]go_thunder.RouterBgpIpv6NetworkList, 0, Ipv6NetworkListCount)

	for i := 0; i < Ipv6NetworkListCount; i++ {
		var obj8_1_4_2 go_thunder.RouterBgpIpv6NetworkList
		prefix8_1_4_2 := prefix8_1_4 + fmt.Sprintf("ipv6_network_list.%d.", i)
		obj8_1_4_2.NetworkIpv6 = d.Get(prefix8_1_4_2 + "network_ipv6").(string)
		obj8_1_4_2.RouteMap = d.Get(prefix8_1_4_2 + "route_map").(string)
		obj8_1_4_2.Backdoor = d.Get(prefix8_1_4_2 + "backdoor").(int)
		obj8_1_4_2.Description = d.Get(prefix8_1_4_2 + "description").(string)
		obj8_1_4_2.CommValue = d.Get(prefix8_1_4_2 + "comm_value").(string)
		obj8_1_4.NetworkIpv6 = append(obj8_1_4.NetworkIpv6, obj8_1_4_2)
	}

	obj8_1.NetworkSynchronization = obj8_1_4

	var obj8_1_5 go_thunder.RouterBgpNeighbor6
	prefix8_1_5 := prefix8_1 + "neighbor.0."

	//obj8_1_5.TempVal = d.Get(prefix8_1_5 + "temp_val").(string)
	var obj8_1_5_44 go_thunder.RouterBgpDistance
	prefix8_1_5_44 := prefix8_1_5 + "distance.0."
	obj8_1_5_44.DistanceExt = d.Get(prefix8_1_5_44 + "distance_ext").(int)
	obj8_1_5_44.DistanceInt = d.Get(prefix8_1_5_44 + "distance_int").(int)
	obj8_1_5_44.DistanceLocal = d.Get(prefix8_1_5_44 + "distance_local").(int)

	obj8_1_5.DistanceExt = obj8_1_5_44

	PeerGroupNeighborListCount = d.Get(prefix8_1_5 + "peer_group_neighbor_list.#").(int)
	obj8_1_5.PeerGroup = make([]go_thunder.RouterBgpPeerGroupNeighborList, 0, PeerGroupNeighborListCount)

	for i := 0; i < PeerGroupNeighborListCount; i++ {
		var obj8_1_5_1 go_thunder.RouterBgpPeerGroupNeighborList
		prefix8_1_5_1 := prefix8_1_5 + fmt.Sprintf("peer_group_neighbor_list.%d.", i)
		obj8_1_5_1.PeerGroup = d.Get(prefix8_1_5_1 + "peer_group").(string)
		obj8_1_5_1.Activate = d.Get(prefix8_1_5_1 + "activate").(int)
		obj8_1_5_1.AllowasIn = d.Get(prefix8_1_5_1 + "allowas_in").(int)
		obj8_1_5_1.AllowasInCount = d.Get(prefix8_1_5_1 + "allowas_in_count").(int)
		obj8_1_5_1.PrefixListDirection = d.Get(prefix8_1_5_1 + "prefix_list_direction").(string)
		obj8_1_5_1.DefaultOriginate = d.Get(prefix8_1_5_1 + "default_originate").(int)
		obj8_1_5_1.RouteMap = d.Get(prefix8_1_5_1 + "route_map").(string)

		DistributeListsCount := d.Get(prefix8_1_5_1 + "distribute_lists.#").(int)
		obj8_1_5_1.DistributeList = make([]go_thunder.RouterBgpDistributeLists, 0, DistributeListsCount)

		for i := 0; i < DistributeListsCount; i++ {
			var obj8_1_5_1_1 go_thunder.RouterBgpDistributeLists
			prefix8_1_5_1_1 := prefix8_1_5_1 + fmt.Sprintf("distribute_lists.%d.", i)
			obj8_1_5_1_1.DistributeList = d.Get(prefix8_1_5_1_1 + "distribute_list").(string)
			obj8_1_5_1_1.DistributeListDirection = d.Get(prefix8_1_5_1_1 + "distribute_list_direction").(string)
			obj8_1_5_1.DistributeList = append(obj8_1_5_1.DistributeList, obj8_1_5_1_1)
		}

		NeighborFilterListsCount := d.Get(prefix8_1_5_1 + "neighbor_filter_lists.#").(int)
		obj8_1_5_1.FilterList = make([]go_thunder.RouterBgpNeighborFilterLists, 0, NeighborFilterListsCount)

		for i := 0; i < NeighborFilterListsCount; i++ {
			var obj8_1_5_1_2 go_thunder.RouterBgpNeighborFilterLists
			prefix8_1_5_1_2 := prefix8_1_5_1 + fmt.Sprintf("neighbor_filter_lists.%d.", i)
			obj8_1_5_1_2.FilterList = d.Get(prefix8_1_5_1_2 + "filter_list").(string)
			obj8_1_5_1_2.FilterListDirection = d.Get(prefix8_1_5_1_2 + "filter_list_direction").(string)
			obj8_1_5_1.FilterList = append(obj8_1_5_1.FilterList, obj8_1_5_1_2)
		}

		obj8_1_5_1.MaximumPrefix = d.Get(prefix8_1_5_1 + "maximum_prefix").(int)
		obj8_1_5_1.MaximumPrefixThres = d.Get(prefix8_1_5_1 + "maximum_prefix_thres").(int)
		obj8_1_5_1.NextHopSelf = d.Get(prefix8_1_5_1 + "next_hop_self").(int)

		NeighborPrefixListsCount := d.Get(prefix8_1_5_1 + "neighbor_prefix_lists.#").(int)
		obj8_1_5_1.NbrPrefixList = make([]go_thunder.RouterBgpNeighborPrefixLists, 0, NeighborPrefixListsCount)

		for i := 0; i < NeighborPrefixListsCount; i++ {
			var obj8_1_5_1_3 go_thunder.RouterBgpNeighborPrefixLists
			prefix8_1_5_1_3 := prefix8_1_5_1 + fmt.Sprintf("neighbor_prefix_lists.%d.", i)
			obj8_1_5_1_3.NbrPrefixList = d.Get(prefix8_1_5_1_3 + "nbr_prefix_list").(string)
			obj8_1_5_1_3.NbrPrefixListDirection = d.Get(prefix8_1_5_1_3 + "nbr_prefix_list_direction").(string)
			obj8_1_5_1.NbrPrefixList = append(obj8_1_5_1.NbrPrefixList, obj8_1_5_1_3)
		}

		obj8_1_5_1.RemovePrivateAs = d.Get(prefix8_1_5_1 + "remove_private_as").(int)

		NeighborRouteMapListsCount := d.Get(prefix8_1_5_1 + "neighbor_route_map_lists.#").(int)
		obj8_1_5_1.NbrRouteMap = make([]go_thunder.RouterBgpNeighborRouteMapLists, 0, NeighborRouteMapListsCount)

		for i := 0; i < NeighborRouteMapListsCount; i++ {
			var obj8_1_5_1_4 go_thunder.RouterBgpNeighborRouteMapLists
			prefix8_1_5_1_4 := prefix8_1_5_1 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj8_1_5_1_4.NbrRouteMap = d.Get(prefix8_1_5_1_4 + "nbr_route_map").(string)
			obj8_1_5_1_4.NbrRmapDirection = d.Get(prefix8_1_5_1_4 + "nbr_rmap_direction").(string)
			obj8_1_5_1.NbrRouteMap = append(obj8_1_5_1.NbrRouteMap, obj8_1_5_1_4)
		}

		obj8_1_5_1.SendCommunityVal = d.Get(prefix8_1_5_1 + "send_community_val").(string)
		obj8_1_5_1.Inbound = d.Get(prefix8_1_5_1 + "inbound").(int)
		obj8_1_5_1.UnsuppressMap = d.Get(prefix8_1_5_1 + "unsuppress_map").(string)
		obj8_1_5_1.Weight = d.Get(prefix8_1_5_1 + "weight").(int)
		obj8_1_5.PeerGroup = append(obj8_1_5.PeerGroup, obj8_1_5_1)
	}

	Ipv4NeighborListCount = d.Get(prefix8_1_5 + "ipv4_neighbor_list.#").(int)
	obj8_1_5.NeighborIpv4 = make([]go_thunder.RouterBgpIpv4NeighborList, 0, Ipv4NeighborListCount)

	for i := 0; i < Ipv4NeighborListCount; i++ {
		var obj8_1_5_2 go_thunder.RouterBgpIpv4NeighborList
		prefix8_1_5_2 := prefix8_1_5 + fmt.Sprintf("ipv4_neighbor_list.%d.", i)
		obj8_1_5_2.NeighborIpv4 = d.Get(prefix8_1_5_2 + "neighbor_ipv4").(string)
		obj8_1_5_2.PeerGroupName = d.Get(prefix8_1_5_2 + "peer_group_name").(string)
		obj8_1_5_2.Activate = d.Get(prefix8_1_5_2 + "activate").(int)
		obj8_1_5_2.AllowasIn = d.Get(prefix8_1_5_2 + "allowas_in").(int)
		obj8_1_5_2.AllowasInCount = d.Get(prefix8_1_5_2 + "allowas_in_count").(int)
		obj8_1_5_2.PrefixListDirection = d.Get(prefix8_1_5_2 + "prefix_list_direction").(string)
		obj8_1_5_2.GracefulRestart = d.Get(prefix8_1_5_2 + "graceful_restart").(int)
		obj8_1_5_2.DefaultOriginate = d.Get(prefix8_1_5_2 + "default_originate").(int)
		obj8_1_5_2.RouteMap = d.Get(prefix8_1_5_2 + "route_map").(string)

		DistributeListsCount := d.Get(prefix8_1_5_2 + "distribute_lists.#").(int)
		obj8_1_5_2.DistributeList = make([]go_thunder.RouterBgpDistributeLists, 0, DistributeListsCount)

		for i := 0; i < DistributeListsCount; i++ {
			var obj8_1_5_2_1 go_thunder.RouterBgpDistributeLists
			prefix8_1_5_2_1 := prefix8_1_5_2 + fmt.Sprintf("distribute_lists.%d.", i)
			obj8_1_5_2_1.DistributeList = d.Get(prefix8_1_5_2_1 + "distribute_list").(string)
			obj8_1_5_2_1.DistributeListDirection = d.Get(prefix8_1_5_2_1 + "distribute_list_direction").(string)
			obj8_1_5_2.DistributeList = append(obj8_1_5_2.DistributeList, obj8_1_5_2_1)
		}

		NeighborFilterListsCount := d.Get(prefix8_1_5_2 + "neighbor_filter_lists.#").(int)
		obj8_1_5_2.FilterList = make([]go_thunder.RouterBgpNeighborFilterLists, 0, NeighborFilterListsCount)

		for i := 0; i < NeighborFilterListsCount; i++ {
			var obj8_1_5_2_2 go_thunder.RouterBgpNeighborFilterLists
			prefix8_1_5_2_2 := prefix8_1_5_2 + fmt.Sprintf("neighbor_filter_lists.%d.", i)
			obj8_1_5_2_2.FilterList = d.Get(prefix8_1_5_2_2 + "filter_list").(string)
			obj8_1_5_2_2.FilterListDirection = d.Get(prefix8_1_5_2_2 + "filter_list_direction").(string)
			obj8_1_5_2.FilterList = append(obj8_1_5_2.FilterList, obj8_1_5_2_2)
		}

		obj8_1_5_2.MaximumPrefix = d.Get(prefix8_1_5_2 + "maximum_prefix").(int)
		obj8_1_5_2.MaximumPrefixThres = d.Get(prefix8_1_5_2 + "maximum_prefix_thres").(int)
		obj8_1_5_2.NextHopSelf = d.Get(prefix8_1_5_2 + "next_hop_self").(int)

		NeighborPrefixListsCount := d.Get(prefix8_1_5_2 + "neighbor_prefix_lists.#").(int)
		obj8_1_5_2.NbrPrefixList = make([]go_thunder.RouterBgpNeighborPrefixLists, 0, NeighborPrefixListsCount)

		for i := 0; i < NeighborPrefixListsCount; i++ {
			var obj8_1_5_2_3 go_thunder.RouterBgpNeighborPrefixLists
			prefix8_1_5_2_3 := prefix8_1_5_2 + fmt.Sprintf("neighbor_prefix_lists.%d.", i)
			obj8_1_5_2_3.NbrPrefixList = d.Get(prefix8_1_5_2_3 + "nbr_prefix_list").(string)
			obj8_1_5_2_3.NbrPrefixListDirection = d.Get(prefix8_1_5_2_3 + "nbr_prefix_list_direction").(string)
			obj8_1_5_2.NbrPrefixList = append(obj8_1_5_2.NbrPrefixList, obj8_1_5_2_3)
		}

		obj8_1_5_2.RemovePrivateAs = d.Get(prefix8_1_5_2 + "remove_private_as").(int)

		NeighborRouteMapListsCount := d.Get(prefix8_1_5_2 + "neighbor_route_map_lists.#").(int)
		obj8_1_5_2.NbrRouteMap = make([]go_thunder.RouterBgpNeighborRouteMapLists, 0, NeighborRouteMapListsCount)

		for i := 0; i < NeighborRouteMapListsCount; i++ {
			var obj8_1_5_2_4 go_thunder.RouterBgpNeighborRouteMapLists
			prefix8_1_5_2_4 := prefix8_1_5_2 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj8_1_5_2_4.NbrRouteMap = d.Get(prefix8_1_5_2_4 + "nbr_route_map").(string)
			obj8_1_5_2_4.NbrRmapDirection = d.Get(prefix8_1_5_2_4 + "nbr_rmap_direction").(string)
			obj8_1_5_2.NbrRouteMap = append(obj8_1_5_2.NbrRouteMap, obj8_1_5_2_4)
		}

		obj8_1_5_2.SendCommunityVal = d.Get(prefix8_1_5_2 + "send_community_val").(string)
		obj8_1_5_2.Inbound = d.Get(prefix8_1_5_2 + "inbound").(int)
		obj8_1_5_2.UnsuppressMap = d.Get(prefix8_1_5_2 + "unsuppress_map").(string)
		obj8_1_5_2.Weight = d.Get(prefix8_1_5_2 + "weight").(int)
		obj8_1_5.NeighborIpv4 = append(obj8_1_5.NeighborIpv4, obj8_1_5_2)
	}

	Ipv6NeighborListCount = d.Get(prefix8_1_5 + "ipv6_neighbor_list.#").(int)
	obj8_1_5.NeighborIpv6 = make([]go_thunder.RouterBgpIpv6NeighborList, 0, Ipv6NeighborListCount)

	for i := 0; i < Ipv6NeighborListCount; i++ {
		var obj8_1_5_3 go_thunder.RouterBgpIpv6NeighborList
		prefix8_1_5_3 := prefix8_1_5 + fmt.Sprintf("ipv6_neighbor_list.%d.", i)
		obj8_1_5_3.NeighborIpv6 = d.Get(prefix8_1_5_3 + "neighbor_ipv6").(string)
		obj8_1_5_3.PeerGroupName = d.Get(prefix8_1_5_3 + "peer_group_name").(string)
		obj8_1_5_3.Activate = d.Get(prefix8_1_5_3 + "activate").(int)
		obj8_1_5_3.AllowasIn = d.Get(prefix8_1_5_3 + "allowas_in").(int)
		obj8_1_5_3.AllowasInCount = d.Get(prefix8_1_5_3 + "allowas_in_count").(int)
		obj8_1_5_3.PrefixListDirection = d.Get(prefix8_1_5_3 + "prefix_list_direction").(string)
		obj8_1_5_3.GracefulRestart = d.Get(prefix8_1_5_3 + "graceful_restart").(int)
		obj8_1_5_3.DefaultOriginate = d.Get(prefix8_1_5_3 + "default_originate").(int)
		obj8_1_5_3.RouteMap = d.Get(prefix8_1_5_3 + "route_map").(string)

		DistributeListsCount := d.Get(prefix8_1_5_3 + "distribute_lists.#").(int)
		obj8_1_5_3.DistributeList = make([]go_thunder.RouterBgpDistributeLists, 0, DistributeListsCount)

		for i := 0; i < DistributeListsCount; i++ {
			var obj8_1_5_3_1 go_thunder.RouterBgpDistributeLists
			prefix8_1_5_3_1 := prefix8_1_5_3 + fmt.Sprintf("distribute_lists.%d.", i)
			obj8_1_5_3_1.DistributeList = d.Get(prefix8_1_5_3_1 + "distribute_list").(string)
			obj8_1_5_3_1.DistributeListDirection = d.Get(prefix8_1_5_3_1 + "distribute_list_direction").(string)
			obj8_1_5_3.DistributeList = append(obj8_1_5_3.DistributeList, obj8_1_5_3_1)
		}

		NeighborFilterListsCount := d.Get(prefix8_1_5_3 + "neighbor_filter_lists.#").(int)
		obj8_1_5_3.FilterList = make([]go_thunder.RouterBgpNeighborFilterLists, 0, NeighborFilterListsCount)

		for i := 0; i < NeighborFilterListsCount; i++ {
			var obj8_1_5_3_2 go_thunder.RouterBgpNeighborFilterLists
			prefix8_1_5_3_2 := prefix8_1_5_3 + fmt.Sprintf("neighbor_filter_lists.%d.", i)
			obj8_1_5_3_2.FilterList = d.Get(prefix8_1_5_3_2 + "filter_list").(string)
			obj8_1_5_3_2.FilterListDirection = d.Get(prefix8_1_5_3_2 + "filter_list_direction").(string)
			obj8_1_5_3.FilterList = append(obj8_1_5_3.FilterList, obj8_1_5_3_2)
		}

		obj8_1_5_3.MaximumPrefix = d.Get(prefix8_1_5_3 + "maximum_prefix").(int)
		obj8_1_5_3.MaximumPrefixThres = d.Get(prefix8_1_5_3 + "maximum_prefix_thres").(int)
		obj8_1_5_3.NextHopSelf = d.Get(prefix8_1_5_3 + "next_hop_self").(int)

		NeighborPrefixListsCount := d.Get(prefix8_1_5_3 + "neighbor_prefix_lists.#").(int)
		obj8_1_5_3.NbrPrefixList = make([]go_thunder.RouterBgpNeighborPrefixLists, 0, NeighborPrefixListsCount)

		for i := 0; i < NeighborPrefixListsCount; i++ {
			var obj8_1_5_3_3 go_thunder.RouterBgpNeighborPrefixLists
			prefix8_1_5_3_3 := prefix8_1_5_3 + fmt.Sprintf("neighbor_prefix_lists.%d.", i)
			obj8_1_5_3_3.NbrPrefixList = d.Get(prefix8_1_5_3_3 + "nbr_prefix_list").(string)
			obj8_1_5_3_3.NbrPrefixListDirection = d.Get(prefix8_1_5_3_3 + "nbr_prefix_list_direction").(string)
			obj8_1_5_3.NbrPrefixList = append(obj8_1_5_3.NbrPrefixList, obj8_1_5_3_3)
		}

		obj8_1_5_3.RemovePrivateAs = d.Get(prefix8_1_5_3 + "remove_private_as").(int)

		NeighborRouteMapListsCount := d.Get(prefix8_1_5_3 + "neighbor_route_map_lists.#").(int)
		obj8_1_5_3.NbrRouteMap = make([]go_thunder.RouterBgpNeighborRouteMapLists, 0, NeighborRouteMapListsCount)

		for i := 0; i < NeighborRouteMapListsCount; i++ {
			var obj8_1_5_3_4 go_thunder.RouterBgpNeighborRouteMapLists
			prefix8_1_5_3_4 := prefix8_1_5_3 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj8_1_5_3_4.NbrRouteMap = d.Get(prefix8_1_5_3_4 + "nbr_route_map").(string)
			obj8_1_5_3_4.NbrRmapDirection = d.Get(prefix8_1_5_3_4 + "nbr_rmap_direction").(string)
			obj8_1_5_3.NbrRouteMap = append(obj8_1_5_3.NbrRouteMap, obj8_1_5_3_4)
		}

		obj8_1_5_3.SendCommunityVal = d.Get(prefix8_1_5_3 + "send_community_val").(string)
		obj8_1_5_3.Inbound = d.Get(prefix8_1_5_3 + "inbound").(int)
		obj8_1_5_3.UnsuppressMap = d.Get(prefix8_1_5_3 + "unsuppress_map").(string)
		obj8_1_5_3.Weight = d.Get(prefix8_1_5_3 + "weight").(int)
		obj8_1_5.NeighborIpv6 = append(obj8_1_5.NeighborIpv6, obj8_1_5_3)
	}

	EthernetNeighborIpv6ListCount := d.Get(prefix8_1_5 + "ethernet_neighbor_ipv6_list.#").(int)
	obj8_1_5.Ethernet = make([]go_thunder.RouterBgpEthernetNeighborIpv6List, 0, EthernetNeighborIpv6ListCount)

	for i := 0; i < EthernetNeighborIpv6ListCount; i++ {
		var obj8_1_5_4 go_thunder.RouterBgpEthernetNeighborIpv6List
		prefix8_1_5_4 := prefix8_1_5 + fmt.Sprintf("ethernet_neighbor_ipv6_list.%d.", i)
		obj8_1_5_4.Ethernet = d.Get(prefix8_1_5_4 + "ethernet").(int)
		obj8_1_5_4.PeerGroupName = d.Get(prefix8_1_5_4 + "peer_group_name").(string)
		obj8_1_5.Ethernet = append(obj8_1_5.Ethernet, obj8_1_5_4)
	}

	VeNeighborIpv6ListCount := d.Get(prefix8_1_5 + "ve_neighbor_ipv6_list.#").(int)
	obj8_1_5.Ve = make([]go_thunder.RouterBgpVeNeighborIpv6List, 0, VeNeighborIpv6ListCount)

	for i := 0; i < VeNeighborIpv6ListCount; i++ {
		var obj8_1_5_5 go_thunder.RouterBgpVeNeighborIpv6List
		prefix8_1_5_5 := prefix8_1_5 + fmt.Sprintf("ve_neighbor_ipv6_list.%d.", i)
		obj8_1_5_5.Ve = d.Get(prefix8_1_5_5 + "ve").(int)
		obj8_1_5_5.PeerGroupName = d.Get(prefix8_1_5_5 + "peer_group_name").(string)
		obj8_1_5.Ve = append(obj8_1_5.Ve, obj8_1_5_5)
	}

	TrunkNeighborIpv6ListCount := d.Get(prefix8_1_5 + "trunk_neighbor_ipv6_list.#").(int)
	obj8_1_5.Trunk = make([]go_thunder.RouterBgpTrunkNeighborIpv6List, 0, TrunkNeighborIpv6ListCount)

	for i := 0; i < TrunkNeighborIpv6ListCount; i++ {
		var obj8_1_5_6 go_thunder.RouterBgpTrunkNeighborIpv6List
		prefix8_1_5_6 := prefix8_1_5 + fmt.Sprintf("trunk_neighbor_ipv6_list.%d.", i)
		obj8_1_5_6.Trunk = d.Get(prefix8_1_5_6 + "trunk").(int)
		obj8_1_5_6.PeerGroupName = d.Get(prefix8_1_5_6 + "peer_group_name").(string)
		obj8_1_5.Trunk = append(obj8_1_5.Trunk, obj8_1_5_6)
	}

	obj8_1.Ethernet = obj8_1_5

	var obj8_1_6 go_thunder.RouterBgpRedistributeIpv6
	prefix8_1_6 := prefix8_1 + "redistribute.0."

	var obj8_1_6_1 go_thunder.RouterBgpConnectedCfg
	prefix8_1_6_1 := prefix8_1_6 + "connected_cfg.0."
	obj8_1_6_1.Connected = d.Get(prefix8_1_6_1 + "connected").(int)
	obj8_1_6_1.RouteMap = d.Get(prefix8_1_6_1 + "route_map").(string)

	obj8_1_6.Connected = obj8_1_6_1

	var obj8_1_6_2 go_thunder.RouterBgpFloatingIPCfg
	prefix8_1_6_2 := prefix8_1_6 + "floating_ip_cfg.0."
	obj8_1_6_2.FloatingIP = d.Get(prefix8_1_6_2 + "floating_ip").(int)
	obj8_1_6_2.RouteMap = d.Get(prefix8_1_6_2 + "route_map").(string)

	obj8_1_6.FloatingIP = obj8_1_6_2

	var obj8_1_6_3 go_thunder.RouterBgpNat64Cfg
	prefix8_1_6_3 := prefix8_1_6 + "nat64_cfg.0."
	obj8_1_6_3.Nat64 = d.Get(prefix8_1_6_3 + "nat64").(int)
	obj8_1_6_3.RouteMap = d.Get(prefix8_1_6_3 + "route_map").(string)

	obj8_1_6.Nat64 = obj8_1_6_3

	var obj8_1_6_4 go_thunder.RouterBgpNatMapCfg
	prefix8_1_6_4 := prefix8_1_6 + "nat_map_cfg.0."
	obj8_1_6_4.NatMap = d.Get(prefix8_1_6_4 + "nat_map").(int)
	obj8_1_6_4.RouteMap = d.Get(prefix8_1_6_4 + "route_map").(string)

	obj8_1_6.NatMap = obj8_1_6_4

	var obj8_1_6_5 go_thunder.RouterBgpLw4O6Cfg
	prefix8_1_6_5 := prefix8_1_6 + "lw4o6_cfg.0."
	obj8_1_6_5.Lw4O6 = d.Get(prefix8_1_6_5 + "lw4o6").(int)
	obj8_1_6_5.RouteMap = d.Get(prefix8_1_6_5 + "route_map").(string)

	obj8_1_6.Lw4O6 = obj8_1_6_5

	var obj8_1_6_6 go_thunder.RouterBgpStaticNatCfg
	prefix8_1_6_6 := prefix8_1_6 + "static_nat_cfg.0."
	obj8_1_6_6.StaticNat = d.Get(prefix8_1_6_6 + "static_nat").(int)
	obj8_1_6_6.RouteMap = d.Get(prefix8_1_6_6 + "route_map").(string)

	obj8_1_6.StaticNat = obj8_1_6_6

	var obj8_1_6_7 go_thunder.RouterBgpIPNatCfg
	prefix8_1_6_7 := prefix8_1_6 + "ip_nat_cfg.0."
	obj8_1_6_7.IPNat = d.Get(prefix8_1_6_7 + "ip_nat").(int)
	obj8_1_6_7.RouteMap = d.Get(prefix8_1_6_7 + "route_map").(string)

	obj8_1_6.IPNat = obj8_1_6_7

	var obj8_1_6_8 go_thunder.RouterBgpIPNatListCfg
	prefix8_1_6_8 := prefix8_1_6 + "ip_nat_list_cfg.0."
	obj8_1_6_8.IPNatList = d.Get(prefix8_1_6_8 + "ip_nat_list").(int)
	obj8_1_6_8.RouteMap = d.Get(prefix8_1_6_8 + "route_map").(string)

	obj8_1_6.IPNatList = obj8_1_6_8

	var obj8_1_6_9 go_thunder.RouterBgpIsisCfg
	prefix8_1_6_9 := prefix8_1_6 + "isis_cfg.0."
	obj8_1_6_9.Isis = d.Get(prefix8_1_6_9 + "isis").(int)
	obj8_1_6_9.RouteMap = d.Get(prefix8_1_6_9 + "route_map").(string)

	obj8_1_6.Isis = obj8_1_6_9

	var obj8_1_6_10 go_thunder.RouterBgpOspfCfg
	prefix8_1_6_10 := prefix8_1_6 + "ospf_cfg.0."
	obj8_1_6_10.Ospf = d.Get(prefix8_1_6_10 + "ospf").(int)
	obj8_1_6_10.RouteMap = d.Get(prefix8_1_6_10 + "route_map").(string)

	obj8_1_6.Ospf = obj8_1_6_10

	var obj8_1_6_11 go_thunder.RouterBgpRipCfg
	prefix8_1_6_11 := prefix8_1_6 + "rip_cfg.0."
	obj8_1_6_11.Rip = d.Get(prefix8_1_6_11 + "rip").(int)
	obj8_1_6_11.RouteMap = d.Get(prefix8_1_6_11 + "route_map").(string)

	obj8_1_6.Rip = obj8_1_6_11

	var obj8_1_6_12 go_thunder.RouterBgpStaticCfg
	prefix8_1_6_12 := prefix8_1_6 + "static_cfg.0."
	obj8_1_6_12.Static = d.Get(prefix8_1_6_12 + "static").(int)
	obj8_1_6_12.RouteMap = d.Get(prefix8_1_6_12 + "route_map").(string)

	obj8_1_6.Static = obj8_1_6_12

	var obj8_1_6_13 go_thunder.RouterBgpVip
	prefix8_1_6_13 := prefix8_1_6 + "vip.0."

	var obj8_1_6_13_1 go_thunder.RouterBgpOnlyFlaggedCfg
	prefix8_1_6_13_1 := prefix8_1_6_13 + "only_flagged_cfg.0."
	obj8_1_6_13_1.OnlyFlagged = d.Get(prefix8_1_6_13_1 + "only_flagged").(int)
	obj8_1_6_13_1.RouteMap = d.Get(prefix8_1_6_13_1 + "route_map").(string)

	obj8_1_6_13.OnlyFlagged = obj8_1_6_13_1

	var obj8_1_6_13_2 go_thunder.RouterBgpOnlyNotFlaggedCfg
	prefix8_1_6_13_2 := prefix8_1_6_13 + "only_not_flagged_cfg.0."
	obj8_1_6_13_2.OnlyNotFlagged = d.Get(prefix8_1_6_13_2 + "only_not_flagged").(int)
	obj8_1_6_13_2.RouteMap = d.Get(prefix8_1_6_13_2 + "route_map").(string)

	obj8_1_6_13.OnlyNotFlagged = obj8_1_6_13_2

	obj8_1_6.OnlyFlaggedCfg = obj8_1_6_13

	obj8_1.Connected = obj8_1_6

	obj8.Bgp = obj8_1

	c.Ipv6 = obj8

	vc.AsNumber = c
	return vc
}
