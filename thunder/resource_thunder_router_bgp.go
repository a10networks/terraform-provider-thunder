package thunder

//Thunder resource RouterBgp

import (
	"context"
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterBgpCreate,
		UpdateContext: resourceRouterBgpUpdate,
		ReadContext:   resourceRouterBgpRead,
		DeleteContext: resourceRouterBgpDelete,
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
									"inbound": {
										Type:        schema.TypeInt,
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
									"restart_min": {
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
									"restart_min": {
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
															"inbound": {
																Type:        schema.TypeInt,
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
															"restart_min": {
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
															"restart_min": {
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
		},
	}
}

func resourceRouterBgpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgp (Inside resourceRouterBgpCreate) ")
		name1 := d.Get("as_number").(int)
		data := dataToRouterBgp(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgp --")
		d.SetId(strconv.Itoa(name1))
		err := go_thunder.PostRouterBgp(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterBgp (Inside resourceRouterBgpRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgp(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return diags
}

func resourceRouterBgpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying RouterBgp   (Inside resourceRouterBgpUpdate) ")
		data := dataToRouterBgp(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgp ")
		err := go_thunder.PutRouterBgp(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceRouterBgpDelete) " + name1)
		err := go_thunder.DeleteRouterBgp(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToRouterBgp(d *schema.ResourceData) go_thunder.RouterBgp {
	var vc go_thunder.RouterBgp
	var c go_thunder.RouterBgpInstance
	c.RouterBgpInstanceAsNumber = d.Get("as_number").(int)

	RouterBgpInstanceAggregateAddressListCount := d.Get("aggregate_address_list.#").(int)
	c.RouterBgpInstanceAggregateAddressListAggregateAddress = make([]go_thunder.RouterBgpInstanceAggregateAddressList, 0, RouterBgpInstanceAggregateAddressListCount)

	for i := 0; i < RouterBgpInstanceAggregateAddressListCount; i++ {
		var obj1 go_thunder.RouterBgpInstanceAggregateAddressList
		prefix1 := fmt.Sprintf("aggregate_address_list.%d.", i)
		obj1.RouterBgpInstanceAggregateAddressListAggregateAddress = d.Get(prefix1 + "aggregate_address").(string)
		obj1.RouterBgpInstanceAggregateAddressListAsSet = d.Get(prefix1 + "as_set").(int)
		obj1.RouterBgpInstanceAggregateAddressListSummaryOnly = d.Get(prefix1 + "summary_only").(int)
		c.RouterBgpInstanceAggregateAddressListAggregateAddress = append(c.RouterBgpInstanceAggregateAddressListAggregateAddress, obj1)
	}

	var obj2 go_thunder.RouterBgpInstanceBgp
	prefix2 := "bgp.0."
	obj2.RouterBgpInstanceBgpAlwaysCompareMed = d.Get(prefix2 + "always_compare_med").(int)

	var obj2_1 go_thunder.RouterBgpInstanceBgpBestpathCfg
	prefix2_1 := prefix2 + "bestpath_cfg.0."
	obj2_1.RouterBgpInstanceBgpBestpathCfgIgnore = d.Get(prefix2_1 + "ignore").(int)
	obj2_1.RouterBgpInstanceBgpBestpathCfgCompareRouterid = d.Get(prefix2_1 + "compare_routerid").(int)
	obj2_1.RouterBgpInstanceBgpBestpathCfgRemoveRecvMed = d.Get(prefix2_1 + "remove_recv_med").(int)
	obj2_1.RouterBgpInstanceBgpBestpathCfgRemoveSendMed = d.Get(prefix2_1 + "remove_send_med").(int)
	obj2_1.RouterBgpInstanceBgpBestpathCfgMissingAsWorst = d.Get(prefix2_1 + "missing_as_worst").(int)

	obj2.RouterBgpInstanceBgpBestpathCfgIgnore = obj2_1

	var obj2_2 go_thunder.RouterBgpInstanceBgpDampeningCfg
	prefix2_2 := prefix2 + "dampening_cfg.0."
	obj2_2.RouterBgpInstanceBgpDampeningCfgDampening = d.Get(prefix2_2 + "dampening").(int)
	obj2_2.RouterBgpInstanceBgpDampeningCfgDampeningHalfTime = d.Get(prefix2_2 + "dampening_half_time").(int)
	obj2_2.RouterBgpInstanceBgpDampeningCfgDampeningReuse = d.Get(prefix2_2 + "dampening_reuse").(int)
	obj2_2.RouterBgpInstanceBgpDampeningCfgDampeningSupress = d.Get(prefix2_2 + "dampening_supress").(int)
	obj2_2.RouterBgpInstanceBgpDampeningCfgDampeningMaxSupress = d.Get(prefix2_2 + "dampening_max_supress").(int)
	obj2_2.RouterBgpInstanceBgpDampeningCfgDampeningPenalty = d.Get(prefix2_2 + "dampening_penalty").(int)
	obj2_2.RouterBgpInstanceBgpDampeningCfgRouteMap = d.Get(prefix2_2 + "route_map").(string)

	obj2.RouterBgpInstanceBgpDampeningCfgDampening = obj2_2

	obj2.RouterBgpInstanceBgpLocalPreferenceValue = d.Get(prefix2 + "local_preference_value").(int)
	obj2.RouterBgpInstanceBgpDeterministicMed = d.Get(prefix2 + "deterministic_med").(int)
	obj2.RouterBgpInstanceBgpEnforceFirstAs = d.Get(prefix2 + "enforce_first_as").(int)
	obj2.RouterBgpInstanceBgpFastExternalFailover = d.Get(prefix2 + "fast_external_failover").(int)
	obj2.RouterBgpInstanceBgpLogNeighborChanges = d.Get(prefix2 + "log_neighbor_changes").(int)
	obj2.RouterBgpInstanceBgpNexthopTriggerCount = d.Get(prefix2 + "nexthop_trigger_count").(int)
	obj2.RouterBgpInstanceBgpRouterID = d.Get(prefix2 + "router_id").(string)
	obj2.RouterBgpInstanceBgpOverrideValidation = d.Get(prefix2 + "override_validation").(int)
	obj2.RouterBgpInstanceBgpScanTime = d.Get(prefix2 + "scan_time").(int)
	obj2.RouterBgpInstanceBgpGracefulRestart = d.Get(prefix2 + "graceful_restart").(int)
	obj2.RouterBgpInstanceBgpBgpRestartTime = d.Get(prefix2 + "bgp_restart_time").(int)
	obj2.RouterBgpInstanceBgpBgpStalepathTime = d.Get(prefix2 + "bgp_stalepath_time").(int)

	c.RouterBgpInstanceBgpAlwaysCompareMed = obj2

	RouterBgpInstanceDistanceListCount := d.Get("distance_list.#").(int)
	c.RouterBgpInstanceDistanceListAdminDistance = make([]go_thunder.RouterBgpInstanceDistanceList, 0, RouterBgpInstanceDistanceListCount)

	for i := 0; i < RouterBgpInstanceDistanceListCount; i++ {
		var obj3 go_thunder.RouterBgpInstanceDistanceList
		prefix3 := fmt.Sprintf("distance_list.%d.", i)
		obj3.RouterBgpInstanceDistanceListAdminDistance = d.Get(prefix3 + "admin_distance").(int)
		obj3.RouterBgpInstanceDistanceListSrcPrefix = d.Get(prefix3 + "src_prefix").(string)
		obj3.RouterBgpInstanceDistanceListAclStr = d.Get(prefix3 + "acl_str").(string)
		obj3.RouterBgpInstanceDistanceListExtRoutesDist = d.Get(prefix3 + "ext_routes_dist").(int)
		obj3.RouterBgpInstanceDistanceListIntRoutesDist = d.Get(prefix3 + "int_routes_dist").(int)
		obj3.RouterBgpInstanceDistanceListLocalRoutesDist = d.Get(prefix3 + "local_routes_dist").(int)
		c.RouterBgpInstanceDistanceListAdminDistance = append(c.RouterBgpInstanceDistanceListAdminDistance, obj3)
	}

	c.RouterBgpInstanceMaximumPathsValue = d.Get("maximum_paths_value").(int)
	c.RouterBgpInstanceOriginate = d.Get("originate").(int)

	var obj4 go_thunder.RouterBgpInstanceTimers
	prefix4 := "timers.0."
	obj4.RouterBgpInstanceTimersBgpKeepalive = d.Get(prefix4 + "bgp_keepalive").(int)
	obj4.RouterBgpInstanceTimersBgpHoldtime = d.Get(prefix4 + "bgp_holdtime").(int)

	c.RouterBgpInstanceTimersBgpKeepalive = obj4

	c.RouterBgpInstanceSynchronization = d.Get("synchronization").(int)
	c.RouterBgpInstanceAutoSummary = d.Get("auto_summary").(int)
	c.RouterBgpInstanceUserTag = d.Get("user_tag").(string)

	var obj5 go_thunder.RouterBgpInstanceNetwork
	prefix5 := "network.0."

	var obj5_1 go_thunder.RouterBgpInstanceNetworkSynchronization
	prefix5_1 := prefix5 + "synchronization.0."
	obj5_1.RouterBgpInstanceNetworkSynchronizationNetworkSynchronization = d.Get(prefix5_1 + "network_synchronization").(int)

	obj5.RouterBgpInstanceNetworkSynchronizationNetworkSynchronization = obj5_1

	RouterBgpInstanceNetworkIPCidrListCount := d.Get(prefix5 + "ip_cidr_list.#").(int)
	obj5.RouterBgpInstanceNetworkIPCidrListNetworkIpv4Cidr = make([]go_thunder.RouterBgpInstanceNetworkIPCidrList, 0, RouterBgpInstanceNetworkIPCidrListCount)

	for i := 0; i < RouterBgpInstanceNetworkIPCidrListCount; i++ {
		var obj5_2 go_thunder.RouterBgpInstanceNetworkIPCidrList
		prefix5_2 := prefix5 + fmt.Sprintf("ip_cidr_list.%d.", i)
		obj5_2.RouterBgpInstanceNetworkIPCidrListNetworkIpv4Cidr = d.Get(prefix5_2 + "network_ipv4_cidr").(string)
		obj5_2.RouterBgpInstanceNetworkIPCidrListRouteMap = d.Get(prefix5_2 + "route_map").(string)
		obj5_2.RouterBgpInstanceNetworkIPCidrListBackdoor = d.Get(prefix5_2 + "backdoor").(int)
		obj5_2.RouterBgpInstanceNetworkIPCidrListDescription = d.Get(prefix5_2 + "description").(string)
		obj5_2.RouterBgpInstanceNetworkIPCidrListCommValue = d.Get(prefix5_2 + "comm_value").(string)
		obj5.RouterBgpInstanceNetworkIPCidrListNetworkIpv4Cidr = append(obj5.RouterBgpInstanceNetworkIPCidrListNetworkIpv4Cidr, obj5_2)
	}

	c.RouterBgpInstanceNetworkSynchronization = obj5

	var obj6 go_thunder.RouterBgpInstanceNeighbor
	prefix6 := "neighbor.0."

	RouterBgpInstanceNeighborPeerGroupNeighborListCount := d.Get(prefix6 + "peer_group_neighbor_list.#").(int)
	obj6.RouterBgpInstanceNeighborPeerGroupNeighborListPeerGroup = make([]go_thunder.RouterBgpInstanceNeighborPeerGroupNeighborList, 0, RouterBgpInstanceNeighborPeerGroupNeighborListCount)

	for i := 0; i < RouterBgpInstanceNeighborPeerGroupNeighborListCount; i++ {
		var obj6_1 go_thunder.RouterBgpInstanceNeighborPeerGroupNeighborList
		prefix6_1 := prefix6 + fmt.Sprintf("peer_group_neighbor_list.%d.", i)
		obj6_1.RouterBgpInstanceNeighborPeerGroupNeighborListPeerGroup = d.Get(prefix6_1 + "peer_group").(string)
		obj6_1.RouterBgpInstanceNeighborPeerGroupNeighborListActivate = d.Get(prefix6_1 + "activate").(int)
		obj6_1.RouterBgpInstanceNeighborPeerGroupNeighborListAllowasIn = d.Get(prefix6_1 + "allowas_in").(int)
		obj6_1.RouterBgpInstanceNeighborPeerGroupNeighborListAllowasInCount = d.Get(prefix6_1 + "allowas_in_count").(int)
		obj6_1.RouterBgpInstanceNeighborPeerGroupNeighborListMaximumPrefix = d.Get(prefix6_1 + "maximum_prefix").(int)
		obj6_1.RouterBgpInstanceNeighborPeerGroupNeighborListMaximumPrefixThres = d.Get(prefix6_1 + "maximum_prefix_thres").(int)
		obj6_1.RouterBgpInstanceNeighborPeerGroupNeighborListNextHopSelf = d.Get(prefix6_1 + "next_hop_self").(int)
		obj6_1.RouterBgpInstanceNeighborPeerGroupNeighborListRemovePrivateAs = d.Get(prefix6_1 + "remove_private_as").(int)

		RouterBgpInstanceNeighborPeerGroupNeighborListNeighborRouteMapListsCount := d.Get(prefix6_1 + "neighbor_route_map_lists.#").(int)
		obj6_1.RouterBgpInstanceNeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap = make([]go_thunder.RouterBgpInstanceNeighborPeerGroupNeighborListNeighborRouteMapLists, 0, RouterBgpInstanceNeighborPeerGroupNeighborListNeighborRouteMapListsCount)

		for i := 0; i < RouterBgpInstanceNeighborPeerGroupNeighborListNeighborRouteMapListsCount; i++ {
			var obj6_1_1 go_thunder.RouterBgpInstanceNeighborPeerGroupNeighborListNeighborRouteMapLists
			prefix6_1_1 := prefix6_1 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj6_1_1.RouterBgpInstanceNeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap = d.Get(prefix6_1_1 + "nbr_route_map").(string)
			obj6_1_1.RouterBgpInstanceNeighborPeerGroupNeighborListNeighborRouteMapListsNbrRmapDirection = d.Get(prefix6_1_1 + "nbr_rmap_direction").(string)
			obj6_1.RouterBgpInstanceNeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap = append(obj6_1.RouterBgpInstanceNeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap, obj6_1_1)
		}

		obj6_1.RouterBgpInstanceNeighborPeerGroupNeighborListInbound = d.Get(prefix6_1 + "inbound").(int)
		obj6_1.RouterBgpInstanceNeighborPeerGroupNeighborListWeight = d.Get(prefix6_1 + "weight").(int)
		obj6.RouterBgpInstanceNeighborPeerGroupNeighborListPeerGroup = append(obj6.RouterBgpInstanceNeighborPeerGroupNeighborListPeerGroup, obj6_1)
	}

	RouterBgpInstanceNeighborIpv4NeighborListCount := d.Get(prefix6 + "ipv4_neighbor_list.#").(int)
	obj6.RouterBgpInstanceNeighborIpv4NeighborListNeighborIpv4 = make([]go_thunder.RouterBgpInstanceNeighborIpv4NeighborList, 0, RouterBgpInstanceNeighborIpv4NeighborListCount)

	for i := 0; i < RouterBgpInstanceNeighborIpv4NeighborListCount; i++ {
		var obj6_2 go_thunder.RouterBgpInstanceNeighborIpv4NeighborList
		prefix6_2 := prefix6 + fmt.Sprintf("ipv4_neighbor_list.%d.", i)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListNeighborIpv4 = d.Get(prefix6_2 + "neighbor_ipv4").(string)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListPeerGroupName = d.Get(prefix6_2 + "peer_group_name").(string)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListActivate = d.Get(prefix6_2 + "activate").(int)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListAllowasIn = d.Get(prefix6_2 + "allowas_in").(int)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListAllowasInCount = d.Get(prefix6_2 + "allowas_in_count").(int)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListPrefixListDirection = d.Get(prefix6_2 + "prefix_list_direction").(string)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListGracefulRestart = d.Get(prefix6_2 + "graceful_restart").(int)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListDefaultOriginate = d.Get(prefix6_2 + "default_originate").(int)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListRouteMap = d.Get(prefix6_2 + "route_map").(string)

		RouterBgpInstanceNeighborIpv4NeighborListDistributeListsCount := d.Get(prefix6_2 + "distribute_lists.#").(int)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListDistributeListsDistributeList = make([]go_thunder.RouterBgpInstanceNeighborIpv4NeighborListDistributeLists, 0, RouterBgpInstanceNeighborIpv4NeighborListDistributeListsCount)

		for i := 0; i < RouterBgpInstanceNeighborIpv4NeighborListDistributeListsCount; i++ {
			var obj6_2_1 go_thunder.RouterBgpInstanceNeighborIpv4NeighborListDistributeLists
			prefix6_2_1 := prefix6_2 + fmt.Sprintf("distribute_lists.%d.", i)
			obj6_2_1.RouterBgpInstanceNeighborIpv4NeighborListDistributeListsDistributeList = d.Get(prefix6_2_1 + "distribute_list").(string)
			obj6_2_1.RouterBgpInstanceNeighborIpv4NeighborListDistributeListsDistributeListDirection = d.Get(prefix6_2_1 + "distribute_list_direction").(string)
			obj6_2.RouterBgpInstanceNeighborIpv4NeighborListDistributeListsDistributeList = append(obj6_2.RouterBgpInstanceNeighborIpv4NeighborListDistributeListsDistributeList, obj6_2_1)
		}

		RouterBgpInstanceNeighborIpv4NeighborListNeighborFilterListsCount := d.Get(prefix6_2 + "neighbor_filter_lists.#").(int)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListNeighborFilterListsFilterList = make([]go_thunder.RouterBgpInstanceNeighborIpv4NeighborListNeighborFilterLists, 0, RouterBgpInstanceNeighborIpv4NeighborListNeighborFilterListsCount)

		for i := 0; i < RouterBgpInstanceNeighborIpv4NeighborListNeighborFilterListsCount; i++ {
			var obj6_2_2 go_thunder.RouterBgpInstanceNeighborIpv4NeighborListNeighborFilterLists
			prefix6_2_2 := prefix6_2 + fmt.Sprintf("neighbor_filter_lists.%d.", i)
			obj6_2_2.RouterBgpInstanceNeighborIpv4NeighborListNeighborFilterListsFilterList = d.Get(prefix6_2_2 + "filter_list").(string)
			obj6_2_2.RouterBgpInstanceNeighborIpv4NeighborListNeighborFilterListsFilterListDirection = d.Get(prefix6_2_2 + "filter_list_direction").(string)
			obj6_2.RouterBgpInstanceNeighborIpv4NeighborListNeighborFilterListsFilterList = append(obj6_2.RouterBgpInstanceNeighborIpv4NeighborListNeighborFilterListsFilterList, obj6_2_2)
		}

		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListMaximumPrefix = d.Get(prefix6_2 + "maximum_prefix").(int)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListMaximumPrefixThres = d.Get(prefix6_2 + "maximum_prefix_thres").(int)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListRestartMin = d.Get(prefix6_2 + "restart_min").(int)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListNextHopSelf = d.Get(prefix6_2 + "next_hop_self").(int)

		RouterBgpInstanceNeighborIpv4NeighborListNeighborPrefixListsCount := d.Get(prefix6_2 + "neighbor_prefix_lists.#").(int)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList = make([]go_thunder.RouterBgpInstanceNeighborIpv4NeighborListNeighborPrefixLists, 0, RouterBgpInstanceNeighborIpv4NeighborListNeighborPrefixListsCount)

		for i := 0; i < RouterBgpInstanceNeighborIpv4NeighborListNeighborPrefixListsCount; i++ {
			var obj6_2_3 go_thunder.RouterBgpInstanceNeighborIpv4NeighborListNeighborPrefixLists
			prefix6_2_3 := prefix6_2 + fmt.Sprintf("neighbor_prefix_lists.%d.", i)
			obj6_2_3.RouterBgpInstanceNeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList = d.Get(prefix6_2_3 + "nbr_prefix_list").(string)
			obj6_2_3.RouterBgpInstanceNeighborIpv4NeighborListNeighborPrefixListsNbrPrefixListDirection = d.Get(prefix6_2_3 + "nbr_prefix_list_direction").(string)
			obj6_2.RouterBgpInstanceNeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList = append(obj6_2.RouterBgpInstanceNeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList, obj6_2_3)
		}

		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListRemovePrivateAs = d.Get(prefix6_2 + "remove_private_as").(int)

		RouterBgpInstanceNeighborIpv4NeighborListNeighborRouteMapListsCount := d.Get(prefix6_2 + "neighbor_route_map_lists.#").(int)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap = make([]go_thunder.RouterBgpInstanceNeighborIpv4NeighborListNeighborRouteMapLists, 0, RouterBgpInstanceNeighborIpv4NeighborListNeighborRouteMapListsCount)

		for i := 0; i < RouterBgpInstanceNeighborIpv4NeighborListNeighborRouteMapListsCount; i++ {
			var obj6_2_4 go_thunder.RouterBgpInstanceNeighborIpv4NeighborListNeighborRouteMapLists
			prefix6_2_4 := prefix6_2 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj6_2_4.RouterBgpInstanceNeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap = d.Get(prefix6_2_4 + "nbr_route_map").(string)
			obj6_2_4.RouterBgpInstanceNeighborIpv4NeighborListNeighborRouteMapListsNbrRmapDirection = d.Get(prefix6_2_4 + "nbr_rmap_direction").(string)
			obj6_2.RouterBgpInstanceNeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap = append(obj6_2.RouterBgpInstanceNeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap, obj6_2_4)
		}

		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListSendCommunityVal = d.Get(prefix6_2 + "send_community_val").(string)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListInbound = d.Get(prefix6_2 + "inbound").(int)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListUnsuppressMap = d.Get(prefix6_2 + "unsuppress_map").(string)
		obj6_2.RouterBgpInstanceNeighborIpv4NeighborListWeight = d.Get(prefix6_2 + "weight").(int)
		obj6.RouterBgpInstanceNeighborIpv4NeighborListNeighborIpv4 = append(obj6.RouterBgpInstanceNeighborIpv4NeighborListNeighborIpv4, obj6_2)
	}

	RouterBgpInstanceNeighborIpv6NeighborListCount := d.Get(prefix6 + "ipv6_neighbor_list.#").(int)
	obj6.RouterBgpInstanceNeighborIpv6NeighborListNeighborIpv6 = make([]go_thunder.RouterBgpInstanceNeighborIpv6NeighborList, 0, RouterBgpInstanceNeighborIpv6NeighborListCount)

	for i := 0; i < RouterBgpInstanceNeighborIpv6NeighborListCount; i++ {
		var obj6_3 go_thunder.RouterBgpInstanceNeighborIpv6NeighborList
		prefix6_3 := prefix6 + fmt.Sprintf("ipv6_neighbor_list.%d.", i)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListNeighborIpv6 = d.Get(prefix6_3 + "neighbor_ipv6").(string)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListPeerGroupName = d.Get(prefix6_3 + "peer_group_name").(string)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListActivate = d.Get(prefix6_3 + "activate").(int)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListAllowasIn = d.Get(prefix6_3 + "allowas_in").(int)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListAllowasInCount = d.Get(prefix6_3 + "allowas_in_count").(int)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListPrefixListDirection = d.Get(prefix6_3 + "prefix_list_direction").(string)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListGracefulRestart = d.Get(prefix6_3 + "graceful_restart").(int)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListDefaultOriginate = d.Get(prefix6_3 + "default_originate").(int)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListRouteMap = d.Get(prefix6_3 + "route_map").(string)

		RouterBgpInstanceNeighborIpv6NeighborListDistributeListsCount := d.Get(prefix6_3 + "distribute_lists.#").(int)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListDistributeListsDistributeList = make([]go_thunder.RouterBgpInstanceNeighborIpv6NeighborListDistributeLists, 0, RouterBgpInstanceNeighborIpv6NeighborListDistributeListsCount)

		for i := 0; i < RouterBgpInstanceNeighborIpv6NeighborListDistributeListsCount; i++ {
			var obj6_3_1 go_thunder.RouterBgpInstanceNeighborIpv6NeighborListDistributeLists
			prefix6_3_1 := prefix6_3 + fmt.Sprintf("distribute_lists.%d.", i)
			obj6_3_1.RouterBgpInstanceNeighborIpv6NeighborListDistributeListsDistributeList = d.Get(prefix6_3_1 + "distribute_list").(string)
			obj6_3_1.RouterBgpInstanceNeighborIpv6NeighborListDistributeListsDistributeListDirection = d.Get(prefix6_3_1 + "distribute_list_direction").(string)
			obj6_3.RouterBgpInstanceNeighborIpv6NeighborListDistributeListsDistributeList = append(obj6_3.RouterBgpInstanceNeighborIpv6NeighborListDistributeListsDistributeList, obj6_3_1)
		}

		RouterBgpInstanceNeighborIpv6NeighborListNeighborFilterListsCount := d.Get(prefix6_3 + "neighbor_filter_lists.#").(int)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListNeighborFilterListsFilterList = make([]go_thunder.RouterBgpInstanceNeighborIpv6NeighborListNeighborFilterLists, 0, RouterBgpInstanceNeighborIpv6NeighborListNeighborFilterListsCount)

		for i := 0; i < RouterBgpInstanceNeighborIpv6NeighborListNeighborFilterListsCount; i++ {
			var obj6_3_2 go_thunder.RouterBgpInstanceNeighborIpv6NeighborListNeighborFilterLists
			prefix6_3_2 := prefix6_3 + fmt.Sprintf("neighbor_filter_lists.%d.", i)
			obj6_3_2.RouterBgpInstanceNeighborIpv6NeighborListNeighborFilterListsFilterList = d.Get(prefix6_3_2 + "filter_list").(string)
			obj6_3_2.RouterBgpInstanceNeighborIpv6NeighborListNeighborFilterListsFilterListDirection = d.Get(prefix6_3_2 + "filter_list_direction").(string)
			obj6_3.RouterBgpInstanceNeighborIpv6NeighborListNeighborFilterListsFilterList = append(obj6_3.RouterBgpInstanceNeighborIpv6NeighborListNeighborFilterListsFilterList, obj6_3_2)
		}

		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListMaximumPrefix = d.Get(prefix6_3 + "maximum_prefix").(int)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListMaximumPrefixThres = d.Get(prefix6_3 + "maximum_prefix_thres").(int)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListRestartMin = d.Get(prefix6_3 + "restart_min").(int)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListNextHopSelf = d.Get(prefix6_3 + "next_hop_self").(int)

		RouterBgpInstanceNeighborIpv6NeighborListNeighborPrefixListsCount := d.Get(prefix6_3 + "neighbor_prefix_lists.#").(int)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList = make([]go_thunder.RouterBgpInstanceNeighborIpv6NeighborListNeighborPrefixLists, 0, RouterBgpInstanceNeighborIpv6NeighborListNeighborPrefixListsCount)

		for i := 0; i < RouterBgpInstanceNeighborIpv6NeighborListNeighborPrefixListsCount; i++ {
			var obj6_3_3 go_thunder.RouterBgpInstanceNeighborIpv6NeighborListNeighborPrefixLists
			prefix6_3_3 := prefix6_3 + fmt.Sprintf("neighbor_prefix_lists.%d.", i)
			obj6_3_3.RouterBgpInstanceNeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList = d.Get(prefix6_3_3 + "nbr_prefix_list").(string)
			obj6_3_3.RouterBgpInstanceNeighborIpv6NeighborListNeighborPrefixListsNbrPrefixListDirection = d.Get(prefix6_3_3 + "nbr_prefix_list_direction").(string)
			obj6_3.RouterBgpInstanceNeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList = append(obj6_3.RouterBgpInstanceNeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList, obj6_3_3)
		}

		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListRemovePrivateAs = d.Get(prefix6_3 + "remove_private_as").(int)

		RouterBgpInstanceNeighborIpv6NeighborListNeighborRouteMapListsCount := d.Get(prefix6_3 + "neighbor_route_map_lists.#").(int)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap = make([]go_thunder.RouterBgpInstanceNeighborIpv6NeighborListNeighborRouteMapLists, 0, RouterBgpInstanceNeighborIpv6NeighborListNeighborRouteMapListsCount)

		for i := 0; i < RouterBgpInstanceNeighborIpv6NeighborListNeighborRouteMapListsCount; i++ {
			var obj6_3_4 go_thunder.RouterBgpInstanceNeighborIpv6NeighborListNeighborRouteMapLists
			prefix6_3_4 := prefix6_3 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj6_3_4.RouterBgpInstanceNeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap = d.Get(prefix6_3_4 + "nbr_route_map").(string)
			obj6_3_4.RouterBgpInstanceNeighborIpv6NeighborListNeighborRouteMapListsNbrRmapDirection = d.Get(prefix6_3_4 + "nbr_rmap_direction").(string)
			obj6_3.RouterBgpInstanceNeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap = append(obj6_3.RouterBgpInstanceNeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap, obj6_3_4)
		}

		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListSendCommunityVal = d.Get(prefix6_3 + "send_community_val").(string)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListInbound = d.Get(prefix6_3 + "inbound").(int)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListUnsuppressMap = d.Get(prefix6_3 + "unsuppress_map").(string)
		obj6_3.RouterBgpInstanceNeighborIpv6NeighborListWeight = d.Get(prefix6_3 + "weight").(int)
		obj6.RouterBgpInstanceNeighborIpv6NeighborListNeighborIpv6 = append(obj6.RouterBgpInstanceNeighborIpv6NeighborListNeighborIpv6, obj6_3)
	}

	RouterBgpInstanceNeighborEthernetNeighborListCount := d.Get(prefix6 + "ethernet_neighbor_list.#").(int)
	obj6.RouterBgpInstanceNeighborEthernetNeighborListEthernet = make([]go_thunder.RouterBgpInstanceNeighborEthernetNeighborList, 0, RouterBgpInstanceNeighborEthernetNeighborListCount)

	for i := 0; i < RouterBgpInstanceNeighborEthernetNeighborListCount; i++ {
		var obj6_4 go_thunder.RouterBgpInstanceNeighborEthernetNeighborList
		prefix6_4 := prefix6 + fmt.Sprintf("ethernet_neighbor_list.%d.", i)
		obj6_4.RouterBgpInstanceNeighborEthernetNeighborListEthernet = d.Get(prefix6_4 + "ethernet").(int)
		obj6_4.RouterBgpInstanceNeighborEthernetNeighborListUnnumbered = d.Get(prefix6_4 + "unnumbered").(int)
		obj6_4.RouterBgpInstanceNeighborEthernetNeighborListPeerGroupName = d.Get(prefix6_4 + "peer_group_name").(string)
		obj6.RouterBgpInstanceNeighborEthernetNeighborListEthernet = append(obj6.RouterBgpInstanceNeighborEthernetNeighborListEthernet, obj6_4)
	}

	RouterBgpInstanceNeighborVeNeighborListCount := d.Get(prefix6 + "ve_neighbor_list.#").(int)
	obj6.RouterBgpInstanceNeighborVeNeighborListVe = make([]go_thunder.RouterBgpInstanceNeighborVeNeighborList, 0, RouterBgpInstanceNeighborVeNeighborListCount)

	for i := 0; i < RouterBgpInstanceNeighborVeNeighborListCount; i++ {
		var obj6_5 go_thunder.RouterBgpInstanceNeighborVeNeighborList
		prefix6_5 := prefix6 + fmt.Sprintf("ve_neighbor_list.%d.", i)
		obj6_5.RouterBgpInstanceNeighborVeNeighborListVe = d.Get(prefix6_5 + "ve").(int)
		obj6_5.RouterBgpInstanceNeighborVeNeighborListUnnumbered = d.Get(prefix6_5 + "unnumbered").(int)
		obj6_5.RouterBgpInstanceNeighborVeNeighborListPeerGroupName = d.Get(prefix6_5 + "peer_group_name").(string)
		obj6.RouterBgpInstanceNeighborVeNeighborListVe = append(obj6.RouterBgpInstanceNeighborVeNeighborListVe, obj6_5)
	}

	RouterBgpInstanceNeighborTrunkNeighborListCount := d.Get(prefix6 + "trunk_neighbor_list.#").(int)
	obj6.RouterBgpInstanceNeighborTrunkNeighborListTrunk = make([]go_thunder.RouterBgpInstanceNeighborTrunkNeighborList, 0, RouterBgpInstanceNeighborTrunkNeighborListCount)

	for i := 0; i < RouterBgpInstanceNeighborTrunkNeighborListCount; i++ {
		var obj6_6 go_thunder.RouterBgpInstanceNeighborTrunkNeighborList
		prefix6_6 := prefix6 + fmt.Sprintf("trunk_neighbor_list.%d.", i)
		obj6_6.RouterBgpInstanceNeighborTrunkNeighborListTrunk = d.Get(prefix6_6 + "trunk").(int)
		obj6_6.RouterBgpInstanceNeighborTrunkNeighborListUnnumbered = d.Get(prefix6_6 + "unnumbered").(int)
		obj6_6.RouterBgpInstanceNeighborTrunkNeighborListPeerGroupName = d.Get(prefix6_6 + "peer_group_name").(string)
		obj6.RouterBgpInstanceNeighborTrunkNeighborListTrunk = append(obj6.RouterBgpInstanceNeighborTrunkNeighborListTrunk, obj6_6)
	}

	c.RouterBgpInstanceNeighborPeerGroupNeighborList = obj6

	var obj7 go_thunder.RouterBgpInstanceRedistribute
	prefix7 := "redistribute.0."

	var obj7_1 go_thunder.RouterBgpInstanceRedistributeConnectedCfg
	prefix7_1 := prefix7 + "connected_cfg.0."
	obj7_1.RouterBgpInstanceRedistributeConnectedCfgConnected = d.Get(prefix7_1 + "connected").(int)
	obj7_1.RouterBgpInstanceRedistributeConnectedCfgRouteMap = d.Get(prefix7_1 + "route_map").(string)

	obj7.RouterBgpInstanceRedistributeConnectedCfgConnected = obj7_1

	var obj7_2 go_thunder.RouterBgpInstanceRedistributeFloatingIPCfg
	prefix7_2 := prefix7 + "floating_ip_cfg.0."
	obj7_2.RouterBgpInstanceRedistributeFloatingIPCfgFloatingIP = d.Get(prefix7_2 + "floating_ip").(int)
	obj7_2.RouterBgpInstanceRedistributeFloatingIPCfgRouteMap = d.Get(prefix7_2 + "route_map").(string)

	obj7.RouterBgpInstanceRedistributeFloatingIPCfgFloatingIP = obj7_2

	var obj7_3 go_thunder.RouterBgpInstanceRedistributeLw4O6Cfg
	prefix7_3 := prefix7 + "lw4o6_cfg.0."
	obj7_3.RouterBgpInstanceRedistributeLw4O6CfgLw4O6 = d.Get(prefix7_3 + "lw4o6").(int)
	obj7_3.RouterBgpInstanceRedistributeLw4O6CfgRouteMap = d.Get(prefix7_3 + "route_map").(string)

	obj7.RouterBgpInstanceRedistributeLw4O6CfgLw4O6 = obj7_3

	var obj7_4 go_thunder.RouterBgpInstanceRedistributeStaticNatCfg
	prefix7_4 := prefix7 + "static_nat_cfg.0."
	obj7_4.RouterBgpInstanceRedistributeStaticNatCfgStaticNat = d.Get(prefix7_4 + "static_nat").(int)
	obj7_4.RouterBgpInstanceRedistributeStaticNatCfgRouteMap = d.Get(prefix7_4 + "route_map").(string)

	obj7.RouterBgpInstanceRedistributeStaticNatCfgStaticNat = obj7_4

	var obj7_5 go_thunder.RouterBgpInstanceRedistributeIPNatCfg
	prefix7_5 := prefix7 + "ip_nat_cfg.0."
	obj7_5.RouterBgpInstanceRedistributeIPNatCfgIPNat = d.Get(prefix7_5 + "ip_nat").(int)
	obj7_5.RouterBgpInstanceRedistributeIPNatCfgRouteMap = d.Get(prefix7_5 + "route_map").(string)

	obj7.RouterBgpInstanceRedistributeIPNatCfgIPNat = obj7_5

	var obj7_6 go_thunder.RouterBgpInstanceRedistributeIPNatListCfg
	prefix7_6 := prefix7 + "ip_nat_list_cfg.0."
	obj7_6.RouterBgpInstanceRedistributeIPNatListCfgIPNatList = d.Get(prefix7_6 + "ip_nat_list").(int)
	obj7_6.RouterBgpInstanceRedistributeIPNatListCfgRouteMap = d.Get(prefix7_6 + "route_map").(string)

	obj7.RouterBgpInstanceRedistributeIPNatListCfgIPNatList = obj7_6

	var obj7_7 go_thunder.RouterBgpInstanceRedistributeIsisCfg
	prefix7_7 := prefix7 + "isis_cfg.0."
	obj7_7.RouterBgpInstanceRedistributeIsisCfgIsis = d.Get(prefix7_7 + "isis").(int)
	obj7_7.RouterBgpInstanceRedistributeIsisCfgRouteMap = d.Get(prefix7_7 + "route_map").(string)

	obj7.RouterBgpInstanceRedistributeIsisCfgIsis = obj7_7

	var obj7_8 go_thunder.RouterBgpInstanceRedistributeOspfCfg
	prefix7_8 := prefix7 + "ospf_cfg.0."
	obj7_8.RouterBgpInstanceRedistributeOspfCfgOspf = d.Get(prefix7_8 + "ospf").(int)
	obj7_8.RouterBgpInstanceRedistributeOspfCfgRouteMap = d.Get(prefix7_8 + "route_map").(string)

	obj7.RouterBgpInstanceRedistributeOspfCfgOspf = obj7_8

	var obj7_9 go_thunder.RouterBgpInstanceRedistributeRipCfg
	prefix7_9 := prefix7 + "rip_cfg.0."
	obj7_9.RouterBgpInstanceRedistributeRipCfgRip = d.Get(prefix7_9 + "rip").(int)
	obj7_9.RouterBgpInstanceRedistributeRipCfgRouteMap = d.Get(prefix7_9 + "route_map").(string)

	obj7.RouterBgpInstanceRedistributeRipCfgRip = obj7_9

	var obj7_10 go_thunder.RouterBgpInstanceRedistributeStaticCfg
	prefix7_10 := prefix7 + "static_cfg.0."
	obj7_10.RouterBgpInstanceRedistributeStaticCfgStatic = d.Get(prefix7_10 + "static").(int)
	obj7_10.RouterBgpInstanceRedistributeStaticCfgRouteMap = d.Get(prefix7_10 + "route_map").(string)

	obj7.RouterBgpInstanceRedistributeStaticCfgStatic = obj7_10

	var obj7_11 go_thunder.RouterBgpInstanceRedistributeNatMapCfg
	prefix7_11 := prefix7 + "nat_map_cfg.0."
	obj7_11.RouterBgpInstanceRedistributeNatMapCfgNatMap = d.Get(prefix7_11 + "nat_map").(int)
	obj7_11.RouterBgpInstanceRedistributeNatMapCfgRouteMap = d.Get(prefix7_11 + "route_map").(string)

	obj7.RouterBgpInstanceRedistributeNatMapCfgNatMap = obj7_11

	var obj7_12 go_thunder.RouterBgpInstanceRedistributeVip
	prefix7_12 := prefix7 + "vip.0."

	var obj7_12_1 go_thunder.RouterBgpInstanceRedistributeVipOnlyFlaggedCfg
	prefix7_12_1 := prefix7_12 + "only_flagged_cfg.0."
	obj7_12_1.RouterBgpInstanceRedistributeVipOnlyFlaggedCfgOnlyFlagged = d.Get(prefix7_12_1 + "only_flagged").(int)
	obj7_12_1.RouterBgpInstanceRedistributeVipOnlyFlaggedCfgRouteMap = d.Get(prefix7_12_1 + "route_map").(string)

	obj7_12.RouterBgpInstanceRedistributeVipOnlyFlaggedCfgOnlyFlagged = obj7_12_1

	var obj7_12_2 go_thunder.RouterBgpInstanceRedistributeVipOnlyNotFlaggedCfg
	prefix7_12_2 := prefix7_12 + "only_not_flagged_cfg.0."
	obj7_12_2.RouterBgpInstanceRedistributeVipOnlyNotFlaggedCfgOnlyNotFlagged = d.Get(prefix7_12_2 + "only_not_flagged").(int)
	obj7_12_2.RouterBgpInstanceRedistributeVipOnlyNotFlaggedCfgRouteMap = d.Get(prefix7_12_2 + "route_map").(string)

	obj7_12.RouterBgpInstanceRedistributeVipOnlyNotFlaggedCfgOnlyNotFlagged = obj7_12_2

	obj7.RouterBgpInstanceRedistributeVipOnlyFlaggedCfg = obj7_12

	c.RouterBgpInstanceRedistributeConnectedCfg = obj7

	var obj8 go_thunder.RouterBgpInstanceAddressFamily
	prefix8 := "address_family.0."

	var obj8_1 go_thunder.RouterBgpInstanceAddressFamilyIpv6
	prefix8_1 := prefix8 + "ipv6.0."

	var obj8_1_1 go_thunder.RouterBgpInstanceAddressFamilyIpv6Bgp
	prefix8_1_1 := prefix8_1 + "bgp.0."
	obj8_1_1.RouterBgpInstanceAddressFamilyIpv6BgpDampening = d.Get(prefix8_1_1 + "dampening").(int)
	obj8_1_1.RouterBgpInstanceAddressFamilyIpv6BgpDampeningHalf = d.Get(prefix8_1_1 + "dampening_half").(int)
	obj8_1_1.RouterBgpInstanceAddressFamilyIpv6BgpDampeningStartReuse = d.Get(prefix8_1_1 + "dampening_start_reuse").(int)
	obj8_1_1.RouterBgpInstanceAddressFamilyIpv6BgpDampeningStartSupress = d.Get(prefix8_1_1 + "dampening_start_supress").(int)
	obj8_1_1.RouterBgpInstanceAddressFamilyIpv6BgpDampeningMaxSupress = d.Get(prefix8_1_1 + "dampening_max_supress").(int)
	obj8_1_1.RouterBgpInstanceAddressFamilyIpv6BgpDampeningUnreachability = d.Get(prefix8_1_1 + "dampening_unreachability").(int)
	obj8_1_1.RouterBgpInstanceAddressFamilyIpv6BgpRouteMap = d.Get(prefix8_1_1 + "route_map").(string)

	obj8_1.RouterBgpInstanceAddressFamilyIpv6BgpDampening = obj8_1_1

	var obj8_1_2 go_thunder.RouterBgpInstanceAddressFamilyIpv6Distance
	prefix8_1_2 := prefix8_1 + "distance.0."
	obj8_1_2.RouterBgpInstanceAddressFamilyIpv6DistanceDistanceExt = d.Get(prefix8_1_2 + "distance_ext").(int)
	obj8_1_2.RouterBgpInstanceAddressFamilyIpv6DistanceDistanceInt = d.Get(prefix8_1_2 + "distance_int").(int)
	obj8_1_2.RouterBgpInstanceAddressFamilyIpv6DistanceDistanceLocal = d.Get(prefix8_1_2 + "distance_local").(int)

	obj8_1.RouterBgpInstanceAddressFamilyIpv6DistanceDistanceExt = obj8_1_2

	obj8_1.RouterBgpInstanceAddressFamilyIpv6MaximumPathsValue = d.Get(prefix8_1 + "maximum_paths_value").(int)
	obj8_1.RouterBgpInstanceAddressFamilyIpv6Originate = d.Get(prefix8_1 + "originate").(int)

	RouterBgpInstanceAddressFamilyIpv6AggregateAddressListCount := d.Get(prefix8_1 + "aggregate_address_list.#").(int)
	obj8_1.RouterBgpInstanceAddressFamilyIpv6AggregateAddressListAggregateAddress = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6AggregateAddressList, 0, RouterBgpInstanceAddressFamilyIpv6AggregateAddressListCount)

	for i := 0; i < RouterBgpInstanceAddressFamilyIpv6AggregateAddressListCount; i++ {
		var obj8_1_3 go_thunder.RouterBgpInstanceAddressFamilyIpv6AggregateAddressList
		prefix8_1_3 := prefix8_1 + fmt.Sprintf("aggregate_address_list.%d.", i)
		obj8_1_3.RouterBgpInstanceAddressFamilyIpv6AggregateAddressListAggregateAddress = d.Get(prefix8_1_3 + "aggregate_address").(string)
		obj8_1_3.RouterBgpInstanceAddressFamilyIpv6AggregateAddressListAsSet = d.Get(prefix8_1_3 + "as_set").(int)
		obj8_1_3.RouterBgpInstanceAddressFamilyIpv6AggregateAddressListSummaryOnly = d.Get(prefix8_1_3 + "summary_only").(int)
		obj8_1.RouterBgpInstanceAddressFamilyIpv6AggregateAddressListAggregateAddress = append(obj8_1.RouterBgpInstanceAddressFamilyIpv6AggregateAddressListAggregateAddress, obj8_1_3)
	}

	obj8_1.RouterBgpInstanceAddressFamilyIpv6AutoSummary = d.Get(prefix8_1 + "auto_summary").(int)
	obj8_1.RouterBgpInstanceAddressFamilyIpv6Synchronization = d.Get(prefix8_1 + "synchronization").(int)

	var obj8_1_4 go_thunder.RouterBgpInstanceAddressFamilyIpv6Network
	prefix8_1_4 := prefix8_1 + "network.0."

	var obj8_1_4_1 go_thunder.RouterBgpInstanceAddressFamilyIpv6NetworkSynchronization
	prefix8_1_4_1 := prefix8_1_4 + "synchronization.0."
	obj8_1_4_1.RouterBgpInstanceAddressFamilyIpv6NetworkSynchronizationNetworkSynchronization = d.Get(prefix8_1_4_1 + "network_synchronization").(int)

	obj8_1_4.RouterBgpInstanceAddressFamilyIpv6NetworkSynchronizationNetworkSynchronization = obj8_1_4_1

	RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListCount := d.Get(prefix8_1_4 + "ipv6_network_list.#").(int)
	obj8_1_4.RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListNetworkIpv6 = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkList, 0, RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListCount)

	for i := 0; i < RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListCount; i++ {
		var obj8_1_4_2 go_thunder.RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkList
		prefix8_1_4_2 := prefix8_1_4 + fmt.Sprintf("ipv6_network_list.%d.", i)
		obj8_1_4_2.RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListNetworkIpv6 = d.Get(prefix8_1_4_2 + "network_ipv6").(string)
		obj8_1_4_2.RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListRouteMap = d.Get(prefix8_1_4_2 + "route_map").(string)
		obj8_1_4_2.RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListBackdoor = d.Get(prefix8_1_4_2 + "backdoor").(int)
		obj8_1_4_2.RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListDescription = d.Get(prefix8_1_4_2 + "description").(string)
		obj8_1_4_2.RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListCommValue = d.Get(prefix8_1_4_2 + "comm_value").(string)
		obj8_1_4.RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListNetworkIpv6 = append(obj8_1_4.RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListNetworkIpv6, obj8_1_4_2)
	}

	obj8_1.RouterBgpInstanceAddressFamilyIpv6NetworkSynchronization = obj8_1_4

	var obj8_1_5 go_thunder.RouterBgpInstanceAddressFamilyIpv6Neighbor
	prefix8_1_5 := prefix8_1 + "neighbor.0."

	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListCount := d.Get(prefix8_1_5 + "peer_group_neighbor_list.#").(int)
	obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListPeerGroup = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborList, 0, RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListCount)

	for i := 0; i < RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListCount; i++ {
		var obj8_1_5_1 go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborList
		prefix8_1_5_1 := prefix8_1_5 + fmt.Sprintf("peer_group_neighbor_list.%d.", i)
		obj8_1_5_1.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListPeerGroup = d.Get(prefix8_1_5_1 + "peer_group").(string)
		obj8_1_5_1.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListActivate = d.Get(prefix8_1_5_1 + "activate").(int)
		obj8_1_5_1.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListAllowasIn = d.Get(prefix8_1_5_1 + "allowas_in").(int)
		obj8_1_5_1.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListAllowasInCount = d.Get(prefix8_1_5_1 + "allowas_in_count").(int)
		obj8_1_5_1.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListMaximumPrefix = d.Get(prefix8_1_5_1 + "maximum_prefix").(int)
		obj8_1_5_1.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListMaximumPrefixThres = d.Get(prefix8_1_5_1 + "maximum_prefix_thres").(int)
		obj8_1_5_1.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNextHopSelf = d.Get(prefix8_1_5_1 + "next_hop_self").(int)
		obj8_1_5_1.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListRemovePrivateAs = d.Get(prefix8_1_5_1 + "remove_private_as").(int)

		RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapListsCount := d.Get(prefix8_1_5_1 + "neighbor_route_map_lists.#").(int)
		obj8_1_5_1.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists, 0, RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapListsCount)

		for i := 0; i < RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapListsCount; i++ {
			var obj8_1_5_1_1 go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists
			prefix8_1_5_1_1 := prefix8_1_5_1 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj8_1_5_1_1.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap = d.Get(prefix8_1_5_1_1 + "nbr_route_map").(string)
			obj8_1_5_1_1.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapListsNbrRmapDirection = d.Get(prefix8_1_5_1_1 + "nbr_rmap_direction").(string)
			obj8_1_5_1.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap = append(obj8_1_5_1.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap, obj8_1_5_1_1)
		}

		obj8_1_5_1.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListInbound = d.Get(prefix8_1_5_1 + "inbound").(int)
		obj8_1_5_1.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListWeight = d.Get(prefix8_1_5_1 + "weight").(int)
		obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListPeerGroup = append(obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListPeerGroup, obj8_1_5_1)
	}

	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListCount := d.Get(prefix8_1_5 + "ipv4_neighbor_list.#").(int)
	obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborIpv4 = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborList, 0, RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListCount)

	for i := 0; i < RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListCount; i++ {
		var obj8_1_5_2 go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborList
		prefix8_1_5_2 := prefix8_1_5 + fmt.Sprintf("ipv4_neighbor_list.%d.", i)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborIpv4 = d.Get(prefix8_1_5_2 + "neighbor_ipv4").(string)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListPeerGroupName = d.Get(prefix8_1_5_2 + "peer_group_name").(string)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListActivate = d.Get(prefix8_1_5_2 + "activate").(int)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListAllowasIn = d.Get(prefix8_1_5_2 + "allowas_in").(int)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListAllowasInCount = d.Get(prefix8_1_5_2 + "allowas_in_count").(int)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListPrefixListDirection = d.Get(prefix8_1_5_2 + "prefix_list_direction").(string)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListGracefulRestart = d.Get(prefix8_1_5_2 + "graceful_restart").(int)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDefaultOriginate = d.Get(prefix8_1_5_2 + "default_originate").(int)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListRouteMap = d.Get(prefix8_1_5_2 + "route_map").(string)

		RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDistributeListsCount := d.Get(prefix8_1_5_2 + "distribute_lists.#").(int)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDistributeListsDistributeList = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists, 0, RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDistributeListsCount)

		for i := 0; i < RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDistributeListsCount; i++ {
			var obj8_1_5_2_1 go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists
			prefix8_1_5_2_1 := prefix8_1_5_2 + fmt.Sprintf("distribute_lists.%d.", i)
			obj8_1_5_2_1.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDistributeListsDistributeList = d.Get(prefix8_1_5_2_1 + "distribute_list").(string)
			obj8_1_5_2_1.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDistributeListsDistributeListDirection = d.Get(prefix8_1_5_2_1 + "distribute_list_direction").(string)
			obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDistributeListsDistributeList = append(obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDistributeListsDistributeList, obj8_1_5_2_1)
		}

		RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterListsCount := d.Get(prefix8_1_5_2 + "neighbor_filter_lists.#").(int)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterListsFilterList = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists, 0, RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterListsCount)

		for i := 0; i < RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterListsCount; i++ {
			var obj8_1_5_2_2 go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists
			prefix8_1_5_2_2 := prefix8_1_5_2 + fmt.Sprintf("neighbor_filter_lists.%d.", i)
			obj8_1_5_2_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterListsFilterList = d.Get(prefix8_1_5_2_2 + "filter_list").(string)
			obj8_1_5_2_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterListsFilterListDirection = d.Get(prefix8_1_5_2_2 + "filter_list_direction").(string)
			obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterListsFilterList = append(obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterListsFilterList, obj8_1_5_2_2)
		}

		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListMaximumPrefix = d.Get(prefix8_1_5_2 + "maximum_prefix").(int)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListMaximumPrefixThres = d.Get(prefix8_1_5_2 + "maximum_prefix_thres").(int)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListRestartMin = d.Get(prefix8_1_5_2 + "restart_min").(int)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNextHopSelf = d.Get(prefix8_1_5_2 + "next_hop_self").(int)

		RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixListsCount := d.Get(prefix8_1_5_2 + "neighbor_prefix_lists.#").(int)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists, 0, RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixListsCount)

		for i := 0; i < RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixListsCount; i++ {
			var obj8_1_5_2_3 go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists
			prefix8_1_5_2_3 := prefix8_1_5_2 + fmt.Sprintf("neighbor_prefix_lists.%d.", i)
			obj8_1_5_2_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList = d.Get(prefix8_1_5_2_3 + "nbr_prefix_list").(string)
			obj8_1_5_2_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixListsNbrPrefixListDirection = d.Get(prefix8_1_5_2_3 + "nbr_prefix_list_direction").(string)
			obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList = append(obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList, obj8_1_5_2_3)
		}

		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListRemovePrivateAs = d.Get(prefix8_1_5_2 + "remove_private_as").(int)

		RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapListsCount := d.Get(prefix8_1_5_2 + "neighbor_route_map_lists.#").(int)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists, 0, RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapListsCount)

		for i := 0; i < RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapListsCount; i++ {
			var obj8_1_5_2_4 go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists
			prefix8_1_5_2_4 := prefix8_1_5_2 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj8_1_5_2_4.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap = d.Get(prefix8_1_5_2_4 + "nbr_route_map").(string)
			obj8_1_5_2_4.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapListsNbrRmapDirection = d.Get(prefix8_1_5_2_4 + "nbr_rmap_direction").(string)
			obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap = append(obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap, obj8_1_5_2_4)
		}

		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListSendCommunityVal = d.Get(prefix8_1_5_2 + "send_community_val").(string)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListInbound = d.Get(prefix8_1_5_2 + "inbound").(int)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListUnsuppressMap = d.Get(prefix8_1_5_2 + "unsuppress_map").(string)
		obj8_1_5_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListWeight = d.Get(prefix8_1_5_2 + "weight").(int)
		obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborIpv4 = append(obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborIpv4, obj8_1_5_2)
	}

	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListCount := d.Get(prefix8_1_5 + "ipv6_neighbor_list.#").(int)
	obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborIpv6 = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborList, 0, RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListCount)

	for i := 0; i < RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListCount; i++ {
		var obj8_1_5_3 go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborList
		prefix8_1_5_3 := prefix8_1_5 + fmt.Sprintf("ipv6_neighbor_list.%d.", i)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborIpv6 = d.Get(prefix8_1_5_3 + "neighbor_ipv6").(string)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListPeerGroupName = d.Get(prefix8_1_5_3 + "peer_group_name").(string)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListActivate = d.Get(prefix8_1_5_3 + "activate").(int)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListAllowasIn = d.Get(prefix8_1_5_3 + "allowas_in").(int)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListAllowasInCount = d.Get(prefix8_1_5_3 + "allowas_in_count").(int)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListPrefixListDirection = d.Get(prefix8_1_5_3 + "prefix_list_direction").(string)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListGracefulRestart = d.Get(prefix8_1_5_3 + "graceful_restart").(int)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDefaultOriginate = d.Get(prefix8_1_5_3 + "default_originate").(int)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListRouteMap = d.Get(prefix8_1_5_3 + "route_map").(string)

		RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDistributeListsCount := d.Get(prefix8_1_5_3 + "distribute_lists.#").(int)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDistributeListsDistributeList = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists, 0, RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDistributeListsCount)

		for i := 0; i < RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDistributeListsCount; i++ {
			var obj8_1_5_3_1 go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists
			prefix8_1_5_3_1 := prefix8_1_5_3 + fmt.Sprintf("distribute_lists.%d.", i)
			obj8_1_5_3_1.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDistributeListsDistributeList = d.Get(prefix8_1_5_3_1 + "distribute_list").(string)
			obj8_1_5_3_1.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDistributeListsDistributeListDirection = d.Get(prefix8_1_5_3_1 + "distribute_list_direction").(string)
			obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDistributeListsDistributeList = append(obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDistributeListsDistributeList, obj8_1_5_3_1)
		}

		RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterListsCount := d.Get(prefix8_1_5_3 + "neighbor_filter_lists.#").(int)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterListsFilterList = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists, 0, RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterListsCount)

		for i := 0; i < RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterListsCount; i++ {
			var obj8_1_5_3_2 go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists
			prefix8_1_5_3_2 := prefix8_1_5_3 + fmt.Sprintf("neighbor_filter_lists.%d.", i)
			obj8_1_5_3_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterListsFilterList = d.Get(prefix8_1_5_3_2 + "filter_list").(string)
			obj8_1_5_3_2.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterListsFilterListDirection = d.Get(prefix8_1_5_3_2 + "filter_list_direction").(string)
			obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterListsFilterList = append(obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterListsFilterList, obj8_1_5_3_2)
		}

		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListMaximumPrefix = d.Get(prefix8_1_5_3 + "maximum_prefix").(int)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListMaximumPrefixThres = d.Get(prefix8_1_5_3 + "maximum_prefix_thres").(int)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListRestartMin = d.Get(prefix8_1_5_3 + "restart_min").(int)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNextHopSelf = d.Get(prefix8_1_5_3 + "next_hop_self").(int)

		RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixListsCount := d.Get(prefix8_1_5_3 + "neighbor_prefix_lists.#").(int)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists, 0, RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixListsCount)

		for i := 0; i < RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixListsCount; i++ {
			var obj8_1_5_3_3 go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists
			prefix8_1_5_3_3 := prefix8_1_5_3 + fmt.Sprintf("neighbor_prefix_lists.%d.", i)
			obj8_1_5_3_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList = d.Get(prefix8_1_5_3_3 + "nbr_prefix_list").(string)
			obj8_1_5_3_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixListsNbrPrefixListDirection = d.Get(prefix8_1_5_3_3 + "nbr_prefix_list_direction").(string)
			obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList = append(obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList, obj8_1_5_3_3)
		}

		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListRemovePrivateAs = d.Get(prefix8_1_5_3 + "remove_private_as").(int)

		RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapListsCount := d.Get(prefix8_1_5_3 + "neighbor_route_map_lists.#").(int)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists, 0, RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapListsCount)

		for i := 0; i < RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapListsCount; i++ {
			var obj8_1_5_3_4 go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists
			prefix8_1_5_3_4 := prefix8_1_5_3 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj8_1_5_3_4.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap = d.Get(prefix8_1_5_3_4 + "nbr_route_map").(string)
			obj8_1_5_3_4.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapListsNbrRmapDirection = d.Get(prefix8_1_5_3_4 + "nbr_rmap_direction").(string)
			obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap = append(obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap, obj8_1_5_3_4)
		}

		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListSendCommunityVal = d.Get(prefix8_1_5_3 + "send_community_val").(string)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListInbound = d.Get(prefix8_1_5_3 + "inbound").(int)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListUnsuppressMap = d.Get(prefix8_1_5_3 + "unsuppress_map").(string)
		obj8_1_5_3.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListWeight = d.Get(prefix8_1_5_3 + "weight").(int)
		obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborIpv6 = append(obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborIpv6, obj8_1_5_3)
	}

	RouterBgpInstanceAddressFamilyIpv6NeighborEthernetNeighborIpv6ListCount := d.Get(prefix8_1_5 + "ethernet_neighbor_ipv6_list.#").(int)
	obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborEthernetNeighborIpv6ListEthernet = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborEthernetNeighborIpv6List, 0, RouterBgpInstanceAddressFamilyIpv6NeighborEthernetNeighborIpv6ListCount)

	for i := 0; i < RouterBgpInstanceAddressFamilyIpv6NeighborEthernetNeighborIpv6ListCount; i++ {
		var obj8_1_5_4 go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborEthernetNeighborIpv6List
		prefix8_1_5_4 := prefix8_1_5 + fmt.Sprintf("ethernet_neighbor_ipv6_list.%d.", i)
		obj8_1_5_4.RouterBgpInstanceAddressFamilyIpv6NeighborEthernetNeighborIpv6ListEthernet = d.Get(prefix8_1_5_4 + "ethernet").(int)
		obj8_1_5_4.RouterBgpInstanceAddressFamilyIpv6NeighborEthernetNeighborIpv6ListPeerGroupName = d.Get(prefix8_1_5_4 + "peer_group_name").(string)
		obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborEthernetNeighborIpv6ListEthernet = append(obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborEthernetNeighborIpv6ListEthernet, obj8_1_5_4)
	}

	RouterBgpInstanceAddressFamilyIpv6NeighborVeNeighborIpv6ListCount := d.Get(prefix8_1_5 + "ve_neighbor_ipv6_list.#").(int)
	obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborVeNeighborIpv6ListVe = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborVeNeighborIpv6List, 0, RouterBgpInstanceAddressFamilyIpv6NeighborVeNeighborIpv6ListCount)

	for i := 0; i < RouterBgpInstanceAddressFamilyIpv6NeighborVeNeighborIpv6ListCount; i++ {
		var obj8_1_5_5 go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborVeNeighborIpv6List
		prefix8_1_5_5 := prefix8_1_5 + fmt.Sprintf("ve_neighbor_ipv6_list.%d.", i)
		obj8_1_5_5.RouterBgpInstanceAddressFamilyIpv6NeighborVeNeighborIpv6ListVe = d.Get(prefix8_1_5_5 + "ve").(int)
		obj8_1_5_5.RouterBgpInstanceAddressFamilyIpv6NeighborVeNeighborIpv6ListPeerGroupName = d.Get(prefix8_1_5_5 + "peer_group_name").(string)
		obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborVeNeighborIpv6ListVe = append(obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborVeNeighborIpv6ListVe, obj8_1_5_5)
	}

	RouterBgpInstanceAddressFamilyIpv6NeighborTrunkNeighborIpv6ListCount := d.Get(prefix8_1_5 + "trunk_neighbor_ipv6_list.#").(int)
	obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborTrunkNeighborIpv6ListTrunk = make([]go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborTrunkNeighborIpv6List, 0, RouterBgpInstanceAddressFamilyIpv6NeighborTrunkNeighborIpv6ListCount)

	for i := 0; i < RouterBgpInstanceAddressFamilyIpv6NeighborTrunkNeighborIpv6ListCount; i++ {
		var obj8_1_5_6 go_thunder.RouterBgpInstanceAddressFamilyIpv6NeighborTrunkNeighborIpv6List
		prefix8_1_5_6 := prefix8_1_5 + fmt.Sprintf("trunk_neighbor_ipv6_list.%d.", i)
		obj8_1_5_6.RouterBgpInstanceAddressFamilyIpv6NeighborTrunkNeighborIpv6ListTrunk = d.Get(prefix8_1_5_6 + "trunk").(int)
		obj8_1_5_6.RouterBgpInstanceAddressFamilyIpv6NeighborTrunkNeighborIpv6ListPeerGroupName = d.Get(prefix8_1_5_6 + "peer_group_name").(string)
		obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborTrunkNeighborIpv6ListTrunk = append(obj8_1_5.RouterBgpInstanceAddressFamilyIpv6NeighborTrunkNeighborIpv6ListTrunk, obj8_1_5_6)
	}

	obj8_1.RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborList = obj8_1_5

	var obj8_1_6 go_thunder.RouterBgpInstanceAddressFamilyIpv6Redistribute
	prefix8_1_6 := prefix8_1 + "redistribute.0."

	var obj8_1_6_1 go_thunder.RouterBgpInstanceAddressFamilyIpv6RedistributeConnectedCfg
	prefix8_1_6_1 := prefix8_1_6 + "connected_cfg.0."
	obj8_1_6_1.RouterBgpInstanceAddressFamilyIpv6RedistributeConnectedCfgConnected = d.Get(prefix8_1_6_1 + "connected").(int)
	obj8_1_6_1.RouterBgpInstanceAddressFamilyIpv6RedistributeConnectedCfgRouteMap = d.Get(prefix8_1_6_1 + "route_map").(string)

	obj8_1_6.RouterBgpInstanceAddressFamilyIpv6RedistributeConnectedCfgConnected = obj8_1_6_1

	var obj8_1_6_2 go_thunder.RouterBgpInstanceAddressFamilyIpv6RedistributeFloatingIPCfg
	prefix8_1_6_2 := prefix8_1_6 + "floating_ip_cfg.0."
	obj8_1_6_2.RouterBgpInstanceAddressFamilyIpv6RedistributeFloatingIPCfgFloatingIP = d.Get(prefix8_1_6_2 + "floating_ip").(int)
	obj8_1_6_2.RouterBgpInstanceAddressFamilyIpv6RedistributeFloatingIPCfgRouteMap = d.Get(prefix8_1_6_2 + "route_map").(string)

	obj8_1_6.RouterBgpInstanceAddressFamilyIpv6RedistributeFloatingIPCfgFloatingIP = obj8_1_6_2

	var obj8_1_6_3 go_thunder.RouterBgpInstanceAddressFamilyIpv6RedistributeNat64Cfg
	prefix8_1_6_3 := prefix8_1_6 + "nat64_cfg.0."
	obj8_1_6_3.RouterBgpInstanceAddressFamilyIpv6RedistributeNat64CfgNat64 = d.Get(prefix8_1_6_3 + "nat64").(int)
	obj8_1_6_3.RouterBgpInstanceAddressFamilyIpv6RedistributeNat64CfgRouteMap = d.Get(prefix8_1_6_3 + "route_map").(string)

	obj8_1_6.RouterBgpInstanceAddressFamilyIpv6RedistributeNat64CfgNat64 = obj8_1_6_3

	var obj8_1_6_4 go_thunder.RouterBgpInstanceAddressFamilyIpv6RedistributeNatMapCfg
	prefix8_1_6_4 := prefix8_1_6 + "nat_map_cfg.0."
	obj8_1_6_4.RouterBgpInstanceAddressFamilyIpv6RedistributeNatMapCfgNatMap = d.Get(prefix8_1_6_4 + "nat_map").(int)
	obj8_1_6_4.RouterBgpInstanceAddressFamilyIpv6RedistributeNatMapCfgRouteMap = d.Get(prefix8_1_6_4 + "route_map").(string)

	obj8_1_6.RouterBgpInstanceAddressFamilyIpv6RedistributeNatMapCfgNatMap = obj8_1_6_4

	var obj8_1_6_5 go_thunder.RouterBgpInstanceAddressFamilyIpv6RedistributeLw4O6Cfg
	prefix8_1_6_5 := prefix8_1_6 + "lw4o6_cfg.0."
	obj8_1_6_5.RouterBgpInstanceAddressFamilyIpv6RedistributeLw4O6CfgLw4O6 = d.Get(prefix8_1_6_5 + "lw4o6").(int)
	obj8_1_6_5.RouterBgpInstanceAddressFamilyIpv6RedistributeLw4O6CfgRouteMap = d.Get(prefix8_1_6_5 + "route_map").(string)

	obj8_1_6.RouterBgpInstanceAddressFamilyIpv6RedistributeLw4O6CfgLw4O6 = obj8_1_6_5

	var obj8_1_6_6 go_thunder.RouterBgpInstanceAddressFamilyIpv6RedistributeStaticNatCfg
	prefix8_1_6_6 := prefix8_1_6 + "static_nat_cfg.0."
	obj8_1_6_6.RouterBgpInstanceAddressFamilyIpv6RedistributeStaticNatCfgStaticNat = d.Get(prefix8_1_6_6 + "static_nat").(int)
	obj8_1_6_6.RouterBgpInstanceAddressFamilyIpv6RedistributeStaticNatCfgRouteMap = d.Get(prefix8_1_6_6 + "route_map").(string)

	obj8_1_6.RouterBgpInstanceAddressFamilyIpv6RedistributeStaticNatCfgStaticNat = obj8_1_6_6

	var obj8_1_6_7 go_thunder.RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatCfg
	prefix8_1_6_7 := prefix8_1_6 + "ip_nat_cfg.0."
	obj8_1_6_7.RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatCfgIPNat = d.Get(prefix8_1_6_7 + "ip_nat").(int)
	obj8_1_6_7.RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatCfgRouteMap = d.Get(prefix8_1_6_7 + "route_map").(string)

	obj8_1_6.RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatCfgIPNat = obj8_1_6_7

	var obj8_1_6_8 go_thunder.RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatListCfg
	prefix8_1_6_8 := prefix8_1_6 + "ip_nat_list_cfg.0."
	obj8_1_6_8.RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatListCfgIPNatList = d.Get(prefix8_1_6_8 + "ip_nat_list").(int)
	obj8_1_6_8.RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatListCfgRouteMap = d.Get(prefix8_1_6_8 + "route_map").(string)

	obj8_1_6.RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatListCfgIPNatList = obj8_1_6_8

	var obj8_1_6_9 go_thunder.RouterBgpInstanceAddressFamilyIpv6RedistributeIsisCfg
	prefix8_1_6_9 := prefix8_1_6 + "isis_cfg.0."
	obj8_1_6_9.RouterBgpInstanceAddressFamilyIpv6RedistributeIsisCfgIsis = d.Get(prefix8_1_6_9 + "isis").(int)
	obj8_1_6_9.RouterBgpInstanceAddressFamilyIpv6RedistributeIsisCfgRouteMap = d.Get(prefix8_1_6_9 + "route_map").(string)

	obj8_1_6.RouterBgpInstanceAddressFamilyIpv6RedistributeIsisCfgIsis = obj8_1_6_9

	var obj8_1_6_10 go_thunder.RouterBgpInstanceAddressFamilyIpv6RedistributeOspfCfg
	prefix8_1_6_10 := prefix8_1_6 + "ospf_cfg.0."
	obj8_1_6_10.RouterBgpInstanceAddressFamilyIpv6RedistributeOspfCfgOspf = d.Get(prefix8_1_6_10 + "ospf").(int)
	obj8_1_6_10.RouterBgpInstanceAddressFamilyIpv6RedistributeOspfCfgRouteMap = d.Get(prefix8_1_6_10 + "route_map").(string)

	obj8_1_6.RouterBgpInstanceAddressFamilyIpv6RedistributeOspfCfgOspf = obj8_1_6_10

	var obj8_1_6_11 go_thunder.RouterBgpInstanceAddressFamilyIpv6RedistributeRipCfg
	prefix8_1_6_11 := prefix8_1_6 + "rip_cfg.0."
	obj8_1_6_11.RouterBgpInstanceAddressFamilyIpv6RedistributeRipCfgRip = d.Get(prefix8_1_6_11 + "rip").(int)
	obj8_1_6_11.RouterBgpInstanceAddressFamilyIpv6RedistributeRipCfgRouteMap = d.Get(prefix8_1_6_11 + "route_map").(string)

	obj8_1_6.RouterBgpInstanceAddressFamilyIpv6RedistributeRipCfgRip = obj8_1_6_11

	var obj8_1_6_12 go_thunder.RouterBgpInstanceAddressFamilyIpv6RedistributeStaticCfg
	prefix8_1_6_12 := prefix8_1_6 + "static_cfg.0."
	obj8_1_6_12.RouterBgpInstanceAddressFamilyIpv6RedistributeStaticCfgStatic = d.Get(prefix8_1_6_12 + "static").(int)
	obj8_1_6_12.RouterBgpInstanceAddressFamilyIpv6RedistributeStaticCfgRouteMap = d.Get(prefix8_1_6_12 + "route_map").(string)

	obj8_1_6.RouterBgpInstanceAddressFamilyIpv6RedistributeStaticCfgStatic = obj8_1_6_12

	var obj8_1_6_13 go_thunder.RouterBgpInstanceAddressFamilyIpv6RedistributeVip
	prefix8_1_6_13 := prefix8_1_6 + "vip.0."

	var obj8_1_6_13_1 go_thunder.RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg
	prefix8_1_6_13_1 := prefix8_1_6_13 + "only_flagged_cfg.0."
	obj8_1_6_13_1.RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyFlaggedCfgOnlyFlagged = d.Get(prefix8_1_6_13_1 + "only_flagged").(int)
	obj8_1_6_13_1.RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyFlaggedCfgRouteMap = d.Get(prefix8_1_6_13_1 + "route_map").(string)

	obj8_1_6_13.RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyFlaggedCfgOnlyFlagged = obj8_1_6_13_1

	var obj8_1_6_13_2 go_thunder.RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg
	prefix8_1_6_13_2 := prefix8_1_6_13 + "only_not_flagged_cfg.0."
	obj8_1_6_13_2.RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfgOnlyNotFlagged = d.Get(prefix8_1_6_13_2 + "only_not_flagged").(int)
	obj8_1_6_13_2.RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfgRouteMap = d.Get(prefix8_1_6_13_2 + "route_map").(string)

	obj8_1_6_13.RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfgOnlyNotFlagged = obj8_1_6_13_2

	obj8_1_6.RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg = obj8_1_6_13

	obj8_1.RouterBgpInstanceAddressFamilyIpv6RedistributeConnectedCfg = obj8_1_6

	obj8.RouterBgpInstanceAddressFamilyIpv6Bgp = obj8_1

	c.RouterBgpInstanceAddressFamilyIpv6 = obj8

	vc.RouterBgpInstanceAsNumber = c
	return vc
}
