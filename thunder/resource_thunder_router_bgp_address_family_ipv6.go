package thunder

//Thunder resource RouterBgpAddressFamilyIpv6

import (
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceRouterBgpAddressFamilyIpv6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpAddressFamilyIpv6Create,
		Update: resourceRouterBgpAddressFamilyIpv6Update,
		Read:   resourceRouterBgpAddressFamilyIpv6Read,
		Delete: resourceRouterBgpAddressFamilyIpv6Delete,
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
			"as_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"process_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceRouterBgpAddressFamilyIpv6Create(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpAddressFamilyIpv6 (Inside resourceRouterBgpAddressFamilyIpv6Create) ")
		name1 := d.Get("as_number").(int)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpAddressFamilyIpv6(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6 --")
		d.SetId(strconv.Itoa(name1))
		go_thunder.PostRouterBgpAddressFamilyIpv6(client.Token, name, data, client.Host)

		return resourceRouterBgpAddressFamilyIpv6Read(d, meta)

	}
	return nil
}

func resourceRouterBgpAddressFamilyIpv6Read(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterBgpAddressFamilyIpv6 (Inside resourceRouterBgpAddressFamilyIpv6Read)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpAddressFamilyIpv6(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceRouterBgpAddressFamilyIpv6Update(d *schema.ResourceData, meta interface{}) error {

	return resourceRouterBgpAddressFamilyIpv6Read(d, meta)
}

func resourceRouterBgpAddressFamilyIpv6Delete(d *schema.ResourceData, meta interface{}) error {

	return resourceRouterBgpAddressFamilyIpv6Read(d, meta)
}
func dataToRouterBgpAddressFamilyIpv6(d *schema.ResourceData) go_thunder.RouterBgpAddressFamilyIpv6 {
	var vc go_thunder.RouterBgpAddressFamilyIpv6
	var c go_thunder.RouterBgpAddressFamilyIpv6Instance

	var obj1 go_thunder.RouterBgpAddressFamilyIpv6Bgp
	prefix1 := "bgp.0."
	obj1.Dampening = d.Get(prefix1 + "dampening").(int)
	obj1.DampeningHalf = d.Get(prefix1 + "dampening_half").(int)
	obj1.DampeningStartReuse = d.Get(prefix1 + "dampening_start_reuse").(int)
	obj1.DampeningStartSupress = d.Get(prefix1 + "dampening_start_supress").(int)
	obj1.DampeningMaxSupress = d.Get(prefix1 + "dampening_max_supress").(int)
	obj1.DampeningUnreachability = d.Get(prefix1 + "dampening_unreachability").(int)
	obj1.RouteMap = d.Get(prefix1 + "route_map").(string)

	c.Dampening = obj1

	var obj2 go_thunder.RouterBgpAddressFamilyIpv6Ipv6Distance
	prefix2 := "distance.0."
	obj2.DistanceExt = d.Get(prefix2 + "distance_ext").(int)
	obj2.DistanceInt = d.Get(prefix2 + "distance_int").(int)
	obj2.DistanceLocal = d.Get(prefix2 + "distance_local").(int)

	c.DistanceExt = obj2

	c.MaximumPathsValue = d.Get("maximum_paths_value").(int)
	c.Originate = d.Get("originate").(int)

	AggregateAddressListCount := d.Get("aggregate_address_list.#").(int)
	c.AggregateAddress = make([]go_thunder.RouterBgpAddressFamilyIpv6AggregateAddressList, 0, AggregateAddressListCount)

	for i := 0; i < AggregateAddressListCount; i++ {
		var obj3 go_thunder.RouterBgpAddressFamilyIpv6AggregateAddressList
		prefix3 := fmt.Sprintf("aggregate_address_list.%d.", i)
		obj3.AggregateAddress = d.Get(prefix3 + "aggregate_address").(string)
		obj3.AsSet = d.Get(prefix3 + "as_set").(int)
		obj3.SummaryOnly = d.Get(prefix3 + "summary_only").(int)
		c.AggregateAddress = append(c.AggregateAddress, obj3)
	}

	c.AutoSummary = d.Get("auto_summary").(int)
	c.Synchronization = d.Get("synchronization").(int)

	var obj4 go_thunder.RouterBgpAddressFamilyIpv6Network
	prefix4 := "network.0."

	var obj4_1 go_thunder.RouterBgpAddressFamilyIpv6Synchronization
	prefix4_1 := prefix4 + "synchronization.0."
	obj4_1.NetworkSynchronization = d.Get(prefix4_1 + "network_synchronization").(int)

	obj4.NetworkSynchronization = obj4_1

	Ipv6NetworkListCount := d.Get(prefix4 + "ipv6_network_list.#").(int)
	obj4.NetworkIpv6 = make([]go_thunder.RouterBgpAddressFamilyIpv6NetworkList, 0, Ipv6NetworkListCount)

	for i := 0; i < Ipv6NetworkListCount; i++ {
		var obj4_2 go_thunder.RouterBgpAddressFamilyIpv6NetworkList
		prefix4_2 := prefix4 + fmt.Sprintf("ipv6_network_list.%d.", i)
		obj4_2.NetworkIpv6 = d.Get(prefix4_2 + "network_ipv6").(string)
		obj4_2.RouteMap = d.Get(prefix4_2 + "route_map").(string)
		obj4_2.Backdoor = d.Get(prefix4_2 + "backdoor").(int)
		obj4_2.Description = d.Get(prefix4_2 + "description").(string)
		obj4_2.CommValue = d.Get(prefix4_2 + "comm_value").(string)
		obj4.NetworkIpv6 = append(obj4.NetworkIpv6, obj4_2)
	}

	c.NetworkSynchronization = obj4

	var obj5 go_thunder.RouterBgpAddressFamilyIpv6Ipv6Neighbor
	prefix5 := "neighbor.0."

	var obj5_44 go_thunder.RouterBgpAddressFamilyIpv6Synchronization
	prefix5_44 := prefix5 + "synchronization.0."
	obj5_44.NetworkSynchronization = d.Get(prefix5_44 + "network_synchronization").(int)

	obj5.NetworkSynchronization = obj5_44

	PeerGroupNeighborListCount := d.Get(prefix5 + "peer_group_neighbor_list.#").(int)
	obj5.PeerGroup = make([]go_thunder.RouterBgpAddressFamilyIpv6PeerGroupNeighborList, 0, PeerGroupNeighborListCount)

	for i := 0; i < PeerGroupNeighborListCount; i++ {
		var obj5_1 go_thunder.RouterBgpAddressFamilyIpv6PeerGroupNeighborList
		prefix5_1 := prefix5 + fmt.Sprintf("peer_group_neighbor_list.%d.", i)
		obj5_1.PeerGroup = d.Get(prefix5_1 + "peer_group").(string)
		obj5_1.Activate = d.Get(prefix5_1 + "activate").(int)
		obj5_1.AllowasIn = d.Get(prefix5_1 + "allowas_in").(int)
		obj5_1.AllowasInCount = d.Get(prefix5_1 + "allowas_in_count").(int)
		obj5_1.PrefixListDirection = d.Get(prefix5_1 + "prefix_list_direction").(string)
		obj5_1.DefaultOriginate = d.Get(prefix5_1 + "default_originate").(int)
		obj5_1.RouteMap = d.Get(prefix5_1 + "route_map").(string)

		DistributeListsCount := d.Get(prefix5_1 + "distribute_lists.#").(int)
		obj5_1.DistributeList = make([]go_thunder.RouterBgpAddressFamilyIpv6DistributeLists, 0, DistributeListsCount)

		for i := 0; i < DistributeListsCount; i++ {
			var obj5_1_1 go_thunder.RouterBgpAddressFamilyIpv6DistributeLists
			prefix5_1_1 := prefix5_1 + fmt.Sprintf("distribute_lists.%d.", i)
			obj5_1_1.DistributeList = d.Get(prefix5_1_1 + "distribute_list").(string)
			obj5_1_1.DistributeListDirection = d.Get(prefix5_1_1 + "distribute_list_direction").(string)
			obj5_1.DistributeList = append(obj5_1.DistributeList, obj5_1_1)
		}

		NeighborFilterListsCount := d.Get(prefix5_1 + "neighbor_filter_lists.#").(int)
		obj5_1.FilterList = make([]go_thunder.RouterBgpAddressFamilyIpv6NeighborFilterLists, 0, NeighborFilterListsCount)

		for i := 0; i < NeighborFilterListsCount; i++ {
			var obj5_1_2 go_thunder.RouterBgpAddressFamilyIpv6NeighborFilterLists
			prefix5_1_2 := prefix5_1 + fmt.Sprintf("neighbor_filter_lists.%d.", i)
			obj5_1_2.FilterList = d.Get(prefix5_1_2 + "filter_list").(string)
			obj5_1_2.FilterListDirection = d.Get(prefix5_1_2 + "filter_list_direction").(string)
			obj5_1.FilterList = append(obj5_1.FilterList, obj5_1_2)
		}

		obj5_1.MaximumPrefix = d.Get(prefix5_1 + "maximum_prefix").(int)
		obj5_1.MaximumPrefixThres = d.Get(prefix5_1 + "maximum_prefix_thres").(int)
		obj5_1.NextHopSelf = d.Get(prefix5_1 + "next_hop_self").(int)

		NeighborPrefixListsCount := d.Get(prefix5_1 + "neighbor_prefix_lists.#").(int)
		obj5_1.NbrPrefixList = make([]go_thunder.RouterBgpAddressFamilyIpv6NeighborPrefixLists, 0, NeighborPrefixListsCount)

		for i := 0; i < NeighborPrefixListsCount; i++ {
			var obj5_1_3 go_thunder.RouterBgpAddressFamilyIpv6NeighborPrefixLists
			prefix5_1_3 := prefix5_1 + fmt.Sprintf("neighbor_prefix_lists.%d.", i)
			obj5_1_3.NbrPrefixList = d.Get(prefix5_1_3 + "nbr_prefix_list").(string)
			obj5_1_3.NbrPrefixListDirection = d.Get(prefix5_1_3 + "nbr_prefix_list_direction").(string)
			obj5_1.NbrPrefixList = append(obj5_1.NbrPrefixList, obj5_1_3)
		}

		obj5_1.RemovePrivateAs = d.Get(prefix5_1 + "remove_private_as").(int)

		NeighborRouteMapListsCount := d.Get("neighbor_route_map_lists.#").(int)
		obj5_1.NbrRouteMap = make([]go_thunder.RouterBgpAddressFamilyIpv6NeighborRouteMapLists, 0, NeighborRouteMapListsCount)

		for i := 0; i < NeighborRouteMapListsCount; i++ {
			var obj5_1_4 go_thunder.RouterBgpAddressFamilyIpv6NeighborRouteMapLists
			prefix5_1_4 := prefix5_1 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj5_1_4.NbrRouteMap = d.Get(prefix5_1_4 + "nbr_route_map").(string)
			obj5_1_4.NbrRmapDirection = d.Get(prefix5_1_4 + "nbr_rmap_direction").(string)
			obj5_1.NbrRouteMap = append(obj5_1.NbrRouteMap, obj5_1_4)
		}

		obj5_1.SendCommunityVal = d.Get(prefix5_1 + "send_community_val").(string)
		obj5_1.Inbound = d.Get(prefix5_1 + "inbound").(int)
		obj5_1.UnsuppressMap = d.Get(prefix5_1 + "unsuppress_map").(string)
		obj5_1.Weight = d.Get(prefix5_1 + "weight").(int)
		obj5.PeerGroup = append(obj5.PeerGroup, obj5_1)
	}

	Ipv4NeighborListCount := d.Get(prefix5 + "ipv4_neighbor_list.#").(int)
	obj5.NeighborIpv4 = make([]go_thunder.RouterBgpAddressFamilyIpv6Ipv4NeighborList, 0, Ipv4NeighborListCount)

	for i := 0; i < Ipv4NeighborListCount; i++ {
		var obj5_2 go_thunder.RouterBgpAddressFamilyIpv6Ipv4NeighborList
		prefix5_2 := prefix5 + fmt.Sprintf("ipv4_neighbor_list.%d.", i)
		obj5_2.NeighborIpv4 = d.Get(prefix5_2 + "neighbor_ipv4").(string)
		obj5_2.PeerGroupName = d.Get(prefix5_2 + "peer_group_name").(string)
		obj5_2.Activate = d.Get(prefix5_2 + "activate").(int)
		obj5_2.AllowasIn = d.Get(prefix5_2 + "allowas_in").(int)
		obj5_2.AllowasInCount = d.Get(prefix5_2 + "allowas_in_count").(int)
		obj5_2.PrefixListDirection = d.Get(prefix5_2 + "prefix_list_direction").(string)
		obj5_2.GracefulRestart = d.Get(prefix5_2 + "graceful_restart").(int)
		obj5_2.DefaultOriginate = d.Get(prefix5_2 + "default_originate").(int)
		obj5_2.RouteMap = d.Get(prefix5_2 + "route_map").(string)

		DistributeListsCount := d.Get(prefix5_2 + "distribute_lists.#").(int)
		obj5_2.DistributeList = make([]go_thunder.RouterBgpAddressFamilyIpv6DistributeLists, 0, DistributeListsCount)

		for i := 0; i < DistributeListsCount; i++ {
			var obj5_2_1 go_thunder.RouterBgpAddressFamilyIpv6DistributeLists
			prefix5_2_1 := prefix5_2 + fmt.Sprintf("distribute_lists.%d.", i)
			obj5_2_1.DistributeList = d.Get(prefix5_2_1 + "distribute_list").(string)
			obj5_2_1.DistributeListDirection = d.Get(prefix5_2_1 + "distribute_list_direction").(string)
			obj5_2.DistributeList = append(obj5_2.DistributeList, obj5_2_1)
		}

		NeighborFilterListsCount := d.Get(prefix5_2 + "neighbor_filter_lists.#").(int)
		obj5_2.FilterList = make([]go_thunder.RouterBgpAddressFamilyIpv6NeighborFilterLists, 0, NeighborFilterListsCount)

		for i := 0; i < NeighborFilterListsCount; i++ {
			var obj5_2_2 go_thunder.RouterBgpAddressFamilyIpv6NeighborFilterLists
			prefix5_2_2 := prefix5_2 + fmt.Sprintf("neighbor_filter_lists.%d.", i)
			obj5_2_2.FilterList = d.Get(prefix5_2_2 + "filter_list").(string)
			obj5_2_2.FilterListDirection = d.Get(prefix5_2_2 + "filter_list_direction").(string)
			obj5_2.FilterList = append(obj5_2.FilterList, obj5_2_2)
		}

		obj5_2.MaximumPrefix = d.Get(prefix5_2 + "maximum_prefix").(int)
		obj5_2.MaximumPrefixThres = d.Get(prefix5_2 + "maximum_prefix_thres").(int)
		obj5_2.NextHopSelf = d.Get(prefix5_2 + "next_hop_self").(int)

		NeighborPrefixListsCount := d.Get(prefix5_2 + "neighbor_prefix_lists.#").(int)
		obj5_2.NbrPrefixList = make([]go_thunder.RouterBgpAddressFamilyIpv6NeighborPrefixLists, 0, NeighborPrefixListsCount)

		for i := 0; i < NeighborPrefixListsCount; i++ {
			var obj5_2_3 go_thunder.RouterBgpAddressFamilyIpv6NeighborPrefixLists
			prefix5_2_3 := prefix5_2 + fmt.Sprintf("neighbor_prefix_lists.%d.", i)
			obj5_2_3.NbrPrefixList = d.Get(prefix5_2_3 + "nbr_prefix_list").(string)
			obj5_2_3.NbrPrefixListDirection = d.Get(prefix5_2_3 + "nbr_prefix_list_direction").(string)
			obj5_2.NbrPrefixList = append(obj5_2.NbrPrefixList, obj5_2_3)
		}

		obj5_2.RemovePrivateAs = d.Get(prefix5_2 + "remove_private_as").(int)

		NeighborRouteMapListsCount := d.Get(prefix5_2 + "neighbor_route_map_lists.#").(int)
		obj5_2.NbrRouteMap = make([]go_thunder.RouterBgpAddressFamilyIpv6NeighborRouteMapLists, 0, NeighborRouteMapListsCount)

		for i := 0; i < NeighborRouteMapListsCount; i++ {
			var obj5_2_4 go_thunder.RouterBgpAddressFamilyIpv6NeighborRouteMapLists
			prefix5_2_4 := prefix5_2 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj5_2_4.NbrRouteMap = d.Get(prefix5_2_4 + "nbr_route_map").(string)
			obj5_2_4.NbrRmapDirection = d.Get(prefix5_2_4 + "nbr_rmap_direction").(string)
			obj5_2.NbrRouteMap = append(obj5_2.NbrRouteMap, obj5_2_4)
		}

		obj5_2.SendCommunityVal = d.Get(prefix5_2 + "send_community_val").(string)
		obj5_2.Inbound = d.Get(prefix5_2 + "inbound").(int)
		obj5_2.UnsuppressMap = d.Get(prefix5_2 + "unsuppress_map").(string)
		obj5_2.Weight = d.Get(prefix5_2 + "weight").(int)
		obj5.NeighborIpv4 = append(obj5.NeighborIpv4, obj5_2)
	}

	Ipv6NeighborListCount := d.Get(prefix5 + "ipv6_neighbor_list.#").(int)
	obj5.NeighborIpv6 = make([]go_thunder.RouterBgpAddressFamilyIpv6Ipv6NeighborList, 0, Ipv6NeighborListCount)

	for i := 0; i < Ipv6NeighborListCount; i++ {
		var obj5_3 go_thunder.RouterBgpAddressFamilyIpv6Ipv6NeighborList
		prefix5_3 := prefix5 + fmt.Sprintf("ipv6_neighbor_list.%d.", i)
		obj5_3.NeighborIpv6 = d.Get(prefix5_3 + "neighbor_ipv6").(string)
		obj5_3.PeerGroupName = d.Get(prefix5_3 + "peer_group_name").(string)
		obj5_3.Activate = d.Get(prefix5_3 + "activate").(int)
		obj5_3.AllowasIn = d.Get(prefix5_3 + "allowas_in").(int)
		obj5_3.AllowasInCount = d.Get(prefix5_3 + "allowas_in_count").(int)
		obj5_3.PrefixListDirection = d.Get(prefix5_3 + "prefix_list_direction").(string)
		obj5_3.GracefulRestart = d.Get(prefix5_3 + "graceful_restart").(int)
		obj5_3.DefaultOriginate = d.Get(prefix5_3 + "default_originate").(int)
		obj5_3.RouteMap = d.Get(prefix5_3 + "route_map").(string)

		DistributeListsCount := d.Get(prefix5_3 + "distribute_lists.#").(int)
		obj5_3.DistributeList = make([]go_thunder.RouterBgpAddressFamilyIpv6DistributeLists, 0, DistributeListsCount)

		for i := 0; i < DistributeListsCount; i++ {
			var obj5_3_1 go_thunder.RouterBgpAddressFamilyIpv6DistributeLists
			prefix5_3_1 := prefix5_3 + fmt.Sprintf("distribute_lists.%d.", i)
			obj5_3_1.DistributeList = d.Get(prefix5_3_1 + "distribute_list").(string)
			obj5_3_1.DistributeListDirection = d.Get(prefix5_3_1 + "distribute_list_direction").(string)
			obj5_3.DistributeList = append(obj5_3.DistributeList, obj5_3_1)
		}

		NeighborFilterListsCount := d.Get(prefix5_3 + "neighbor_filter_lists.#").(int)
		obj5_3.FilterList = make([]go_thunder.RouterBgpAddressFamilyIpv6NeighborFilterLists, 0, NeighborFilterListsCount)

		for i := 0; i < NeighborFilterListsCount; i++ {
			var obj5_3_2 go_thunder.RouterBgpAddressFamilyIpv6NeighborFilterLists
			prefix5_3_2 := fmt.Sprintf("neighbor_filter_lists.%d.", i)
			obj5_3_2.FilterList = d.Get(prefix5_3_2 + "filter_list").(string)
			obj5_3_2.FilterListDirection = d.Get(prefix5_3_2 + "filter_list_direction").(string)
			obj5_3.FilterList = append(obj5_3.FilterList, obj5_3_2)
		}

		obj5_3.MaximumPrefix = d.Get(prefix5_3 + "maximum_prefix").(int)
		obj5_3.MaximumPrefixThres = d.Get(prefix5_3 + "maximum_prefix_thres").(int)
		obj5_3.NextHopSelf = d.Get(prefix5_3 + "next_hop_self").(int)

		NeighborPrefixListsCount := d.Get(prefix5_3 + "neighbor_prefix_lists.#").(int)
		obj5_3.NbrPrefixList = make([]go_thunder.RouterBgpAddressFamilyIpv6NeighborPrefixLists, 0, NeighborPrefixListsCount)

		for i := 0; i < NeighborPrefixListsCount; i++ {
			var obj5_3_3 go_thunder.RouterBgpAddressFamilyIpv6NeighborPrefixLists
			prefix5_3_3 := prefix5_3 + fmt.Sprintf("neighbor_prefix_lists.%d.", i)
			obj5_3_3.NbrPrefixList = d.Get(prefix5_3_3 + "nbr_prefix_list").(string)
			obj5_3_3.NbrPrefixListDirection = d.Get(prefix5_3_3 + "nbr_prefix_list_direction").(string)
			obj5_3.NbrPrefixList = append(obj5_3.NbrPrefixList, obj5_3_3)
		}

		obj5_3.RemovePrivateAs = d.Get(prefix5_3 + "remove_private_as").(int)

		NeighborRouteMapListsCount := d.Get(prefix5_3 + "neighbor_route_map_lists.#").(int)
		obj5_3.NbrRouteMap = make([]go_thunder.RouterBgpAddressFamilyIpv6NeighborRouteMapLists, 0, NeighborRouteMapListsCount)

		for i := 0; i < NeighborRouteMapListsCount; i++ {
			var obj5_3_4 go_thunder.RouterBgpAddressFamilyIpv6NeighborRouteMapLists
			prefix5_3_4 := prefix5_3 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj5_3_4.NbrRouteMap = d.Get(prefix5_3_4 + "nbr_route_map").(string)
			obj5_3_4.NbrRmapDirection = d.Get(prefix5_3_4 + "nbr_rmap_direction").(string)
			obj5_3.NbrRouteMap = append(obj5_3.NbrRouteMap, obj5_3_4)
		}

		obj5_3.SendCommunityVal = d.Get(prefix5_3 + "send_community_val").(string)
		obj5_3.Inbound = d.Get(prefix5_3 + "inbound").(int)
		obj5_3.UnsuppressMap = d.Get(prefix5_3 + "unsuppress_map").(string)
		obj5_3.Weight = d.Get(prefix5_3 + "weight").(int)
		obj5.NeighborIpv6 = append(obj5.NeighborIpv6, obj5_3)
	}

	EthernetNeighborIpv6ListCount := d.Get(prefix5 + "ethernet_neighbor_ipv6_list.#").(int)
	obj5.Ethernet = make([]go_thunder.RouterBgpAddressFamilyIpv6EthernetNeighborIpv6List, 0, EthernetNeighborIpv6ListCount)

	for i := 0; i < EthernetNeighborIpv6ListCount; i++ {
		var obj5_4 go_thunder.RouterBgpAddressFamilyIpv6EthernetNeighborIpv6List
		prefix5_4 := prefix5 + fmt.Sprintf("ethernet_neighbor_ipv6_list.%d.", i)
		obj5_4.Ethernet = d.Get(prefix5_4 + "ethernet").(int)
		obj5_4.PeerGroupName = d.Get(prefix5_4 + "peer_group_name").(string)
		obj5.Ethernet = append(obj5.Ethernet, obj5_4)
	}

	VeNeighborIpv6ListCount := d.Get(prefix5 + "ve_neighbor_ipv6_list.#").(int)
	obj5.Ve = make([]go_thunder.RouterBgpAddressFamilyIpv6VeNeighborIpv6List, 0, VeNeighborIpv6ListCount)

	for i := 0; i < VeNeighborIpv6ListCount; i++ {
		var obj5_5 go_thunder.RouterBgpAddressFamilyIpv6VeNeighborIpv6List
		prefix5_5 := prefix5 + fmt.Sprintf("ve_neighbor_ipv6_list.%d.", i)
		obj5_5.Ve = d.Get(prefix5_5 + "ve").(int)
		obj5_5.PeerGroupName = d.Get(prefix5_5 + "peer_group_name").(string)
		obj5.Ve = append(obj5.Ve, obj5_5)
	}

	TrunkNeighborIpv6ListCount := d.Get(prefix5 + "trunk_neighbor_ipv6_list.#").(int)
	obj5.Trunk = make([]go_thunder.RouterBgpAddressFamilyIpv6TrunkNeighborIpv6List, 0, TrunkNeighborIpv6ListCount)

	for i := 0; i < TrunkNeighborIpv6ListCount; i++ {
		var obj5_6 go_thunder.RouterBgpAddressFamilyIpv6TrunkNeighborIpv6List
		prefix5_6 := prefix5 + fmt.Sprintf("trunk_neighbor_ipv6_list.%d.", i)
		obj5_6.Trunk = d.Get(prefix5_6 + "trunk").(int)
		obj5_6.PeerGroupName = d.Get(prefix5_6 + "peer_group_name").(string)
		obj5.Trunk = append(obj5.Trunk, obj5_6)
	}

	c.PeerGroupNeighborList = obj5

	var obj6 go_thunder.RouterBgpAddressFamilyIpv6RedistributeIpv6
	prefix6 := "redistribute.0."

	var obj6_1 go_thunder.RouterBgpAddressFamilyIpv6ConnectedCfg
	prefix6_1 := prefix6 + "connected_cfg.0."
	obj6_1.Connected = d.Get(prefix6_1 + "connected").(int)
	obj6_1.RouteMap = d.Get(prefix6_1 + "route_map").(string)

	obj6.Connected = obj6_1

	var obj6_2 go_thunder.RouterBgpAddressFamilyIpv6FloatingIPCfg
	prefix6_2 := prefix6 + "floating_ip_cfg.0."
	obj6_2.FloatingIP = d.Get(prefix6_2 + "floating_ip").(int)
	obj6_2.RouteMap = d.Get(prefix6_2 + "route_map").(string)

	obj6.FloatingIP = obj6_2

	var obj6_3 go_thunder.RouterBgpAddressFamilyIpv6Nat64Cfg
	prefix6_3 := prefix6 + "nat64_cfg.0."
	obj6_3.Nat64 = d.Get(prefix6_3 + "nat64").(int)
	obj6_3.RouteMap = d.Get(prefix6_3 + "route_map").(string)

	obj6.Nat64 = obj6_3

	var obj6_4 go_thunder.RouterBgpAddressFamilyIpv6NatMapCfg
	prefix6_4 := prefix6 + "nat_map_cfg.0."
	obj6_4.NatMap = d.Get(prefix6_4 + "nat_map").(int)
	obj6_4.RouteMap = d.Get(prefix6_4 + "route_map").(string)

	obj6.NatMap = obj6_4

	var obj6_5 go_thunder.RouterBgpAddressFamilyIpv6Lw4O6Cfg
	prefix6_5 := prefix6 + "lw4o6_cfg.0."
	obj6_5.Lw4O6 = d.Get(prefix6_5 + "lw4o6").(int)
	obj6_5.RouteMap = d.Get(prefix6_5 + "route_map").(string)

	obj6.Lw4O6 = obj6_5

	var obj6_6 go_thunder.RouterBgpAddressFamilyIpv6StaticNatCfg
	prefix6_6 := prefix6 + "static_nat_cfg.0."
	obj6_6.StaticNat = d.Get(prefix6_6 + "static_nat").(int)
	obj6_6.RouteMap = d.Get(prefix6_6 + "route_map").(string)

	obj6.StaticNat = obj6_6

	var obj6_7 go_thunder.RouterBgpAddressFamilyIpv6IPNatCfg
	prefix6_7 := prefix6 + "ip_nat_cfg.0."
	obj6_7.IPNat = d.Get(prefix6_7 + "ip_nat").(int)
	obj6_7.RouteMap = d.Get(prefix6_7 + "route_map").(string)

	obj6.IPNat = obj6_7

	var obj6_8 go_thunder.RouterBgpAddressFamilyIpv6IPNatListCfg
	prefix6_8 := prefix6 + "ip_nat_list_cfg.0."
	obj6_8.IPNatList = d.Get(prefix6_8 + "ip_nat_list").(int)
	obj6_8.RouteMap = d.Get(prefix6_8 + "route_map").(string)

	obj6.IPNatList = obj6_8

	var obj6_9 go_thunder.RouterBgpAddressFamilyIpv6IsisCfg
	prefix6_9 := prefix6 + "isis_cfg.0."
	obj6_9.Isis = d.Get(prefix6_9 + "isis").(int)
	obj6_9.RouteMap = d.Get(prefix6_9 + "route_map").(string)

	obj6.Isis = obj6_9

	var obj6_10 go_thunder.RouterBgpAddressFamilyIpv6OspfCfg
	prefix6_10 := prefix6 + "ospf_cfg.0."
	obj6_10.Ospf = d.Get(prefix6_10 + "ospf").(int)
	obj6_10.RouteMap = d.Get(prefix6_10 + "route_map").(string)

	obj6.Ospf = obj6_10

	var obj6_11 go_thunder.RouterBgpAddressFamilyIpv6RipCfg
	prefix6_11 := prefix6 + "rip_cfg.0."
	obj6_11.Rip = d.Get(prefix6_11 + "rip").(int)
	obj6_11.RouteMap = d.Get(prefix6_11 + "route_map").(string)

	obj6.Rip = obj6_11

	var obj6_12 go_thunder.RouterBgpAddressFamilyIpv6StaticCfg
	prefix6_12 := prefix6 + "static_cfg.0."
	obj6_12.Static = d.Get(prefix6_12 + "static").(int)
	obj6_12.RouteMap = d.Get(prefix6_12 + "route_map").(string)

	obj6.Static = obj6_12

	var obj6_13 go_thunder.RouterBgpAddressFamilyIpv6Vip
	prefix6_13 := prefix6 + "vip.0."

	var obj6_13_1 go_thunder.RouterBgpAddressFamilyIpv6OnlyFlaggedCfg
	prefix6_13_1 := prefix6_13 + "only_flagged_cfg.0."
	obj6_13_1.OnlyFlagged = d.Get(prefix6_13_1 + "only_flagged").(int)
	obj6_13_1.RouteMap = d.Get(prefix6_13_1 + "route_map").(string)

	obj6_13.OnlyFlagged = obj6_13_1

	var obj6_13_2 go_thunder.RouterBgpAddressFamilyIpv6OnlyNotFlaggedCfg
	prefix6_13_2 := prefix6_13 + "only_not_flagged_cfg.0."
	obj6_13_2.OnlyNotFlagged = d.Get(prefix6_13_2 + "only_not_flagged").(int)
	obj6_13_2.RouteMap = d.Get(prefix6_13_2 + "route_map").(string)

	obj6_13.OnlyNotFlagged = obj6_13_2

	obj6.OnlyFlaggedCfg = obj6_13

	c.ConnectedCfg = obj6

	vc.Bgp = c
	return vc
}
