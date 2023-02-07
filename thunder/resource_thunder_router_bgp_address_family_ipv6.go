package thunder

//Thunder resource RouterBgpAddressFamilyIpv6

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceRouterBgpAddressFamilyIpv6() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterBgpAddressFamilyIpv6Create,
		UpdateContext: resourceRouterBgpAddressFamilyIpv6Update,
		ReadContext:   resourceRouterBgpAddressFamilyIpv6Read,
		DeleteContext: resourceRouterBgpAddressFamilyIpv6Delete,
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
			"as_number": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceRouterBgpAddressFamilyIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpAddressFamilyIpv6 (Inside resourceRouterBgpAddressFamilyIpv6Create) ")
		name1 := d.Get("as_number").(string)
		data := dataToRouterBgpAddressFamilyIpv6(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6 --")
		d.SetId(name1)
		err := go_thunder.PostRouterBgpAddressFamilyIpv6(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpAddressFamilyIpv6Read(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterBgpAddressFamilyIpv6 (Inside resourceRouterBgpAddressFamilyIpv6Read)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpAddressFamilyIpv6(client.Token, name1, client.Host)
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

func resourceRouterBgpAddressFamilyIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRouterBgpAddressFamilyIpv6Read(ctx, d, meta)
}

func resourceRouterBgpAddressFamilyIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRouterBgpAddressFamilyIpv6Read(ctx, d, meta)
}
func dataToRouterBgpAddressFamilyIpv6(d *schema.ResourceData) go_thunder.RouterBgpAddressFamilyIpv6 {
	var vc go_thunder.RouterBgpAddressFamilyIpv6
	var c go_thunder.RouterBgpAddressFamilyIpv6Instance

	var obj1 go_thunder.RouterBgpAddressFamilyIpv6InstanceBgp
	prefix1 := "bgp.0."
	obj1.RouterBgpAddressFamilyIpv6InstanceBgpDampening = d.Get(prefix1 + "dampening").(int)
	obj1.RouterBgpAddressFamilyIpv6InstanceBgpDampeningHalf = d.Get(prefix1 + "dampening_half").(int)
	obj1.RouterBgpAddressFamilyIpv6InstanceBgpDampeningStartReuse = d.Get(prefix1 + "dampening_start_reuse").(int)
	obj1.RouterBgpAddressFamilyIpv6InstanceBgpDampeningStartSupress = d.Get(prefix1 + "dampening_start_supress").(int)
	obj1.RouterBgpAddressFamilyIpv6InstanceBgpDampeningMaxSupress = d.Get(prefix1 + "dampening_max_supress").(int)
	obj1.RouterBgpAddressFamilyIpv6InstanceBgpDampeningUnreachability = d.Get(prefix1 + "dampening_unreachability").(int)
	obj1.RouterBgpAddressFamilyIpv6InstanceBgpRouteMap = d.Get(prefix1 + "route_map").(string)

	c.RouterBgpAddressFamilyIpv6InstanceBgpDampening = obj1

	var obj2 go_thunder.RouterBgpAddressFamilyIpv6InstanceDistance
	prefix2 := "distance.0."
	obj2.RouterBgpAddressFamilyIpv6InstanceDistanceDistanceExt = d.Get(prefix2 + "distance_ext").(int)
	obj2.RouterBgpAddressFamilyIpv6InstanceDistanceDistanceInt = d.Get(prefix2 + "distance_int").(int)
	obj2.RouterBgpAddressFamilyIpv6InstanceDistanceDistanceLocal = d.Get(prefix2 + "distance_local").(int)

	c.RouterBgpAddressFamilyIpv6InstanceDistanceDistanceExt = obj2

	c.RouterBgpAddressFamilyIpv6InstanceMaximumPathsValue = d.Get("maximum_paths_value").(int)
	c.RouterBgpAddressFamilyIpv6InstanceOriginate = d.Get("originate").(int)

	RouterBgpAddressFamilyIpv6InstanceAggregateAddressListCount := d.Get("aggregate_address_list.#").(int)
	c.RouterBgpAddressFamilyIpv6InstanceAggregateAddressListAggregateAddress = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceAggregateAddressList, 0, RouterBgpAddressFamilyIpv6InstanceAggregateAddressListCount)

	for i := 0; i < RouterBgpAddressFamilyIpv6InstanceAggregateAddressListCount; i++ {
		var obj3 go_thunder.RouterBgpAddressFamilyIpv6InstanceAggregateAddressList
		prefix3 := fmt.Sprintf("aggregate_address_list.%d.", i)
		obj3.RouterBgpAddressFamilyIpv6InstanceAggregateAddressListAggregateAddress = d.Get(prefix3 + "aggregate_address").(string)
		obj3.RouterBgpAddressFamilyIpv6InstanceAggregateAddressListAsSet = d.Get(prefix3 + "as_set").(int)
		obj3.RouterBgpAddressFamilyIpv6InstanceAggregateAddressListSummaryOnly = d.Get(prefix3 + "summary_only").(int)
		c.RouterBgpAddressFamilyIpv6InstanceAggregateAddressListAggregateAddress = append(c.RouterBgpAddressFamilyIpv6InstanceAggregateAddressListAggregateAddress, obj3)
	}

	c.RouterBgpAddressFamilyIpv6InstanceAutoSummary = d.Get("auto_summary").(int)
	c.RouterBgpAddressFamilyIpv6InstanceSynchronization = d.Get("synchronization").(int)

	var obj4 go_thunder.RouterBgpAddressFamilyIpv6InstanceNetwork
	prefix4 := "network.0."

	var obj4_1 go_thunder.RouterBgpAddressFamilyIpv6InstanceNetworkSynchronization
	prefix4_1 := prefix4 + "synchronization.0."
	obj4_1.RouterBgpAddressFamilyIpv6InstanceNetworkSynchronizationNetworkSynchronization = d.Get(prefix4_1 + "network_synchronization").(int)

	obj4.RouterBgpAddressFamilyIpv6InstanceNetworkSynchronizationNetworkSynchronization = obj4_1

	RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListCount := d.Get(prefix4 + "ipv6_network_list.#").(int)
	obj4.RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListNetworkIpv6 = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkList, 0, RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListCount)

	for i := 0; i < RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListCount; i++ {
		var obj4_2 go_thunder.RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkList
		prefix4_2 := prefix4 + fmt.Sprintf("ipv6_network_list.%d.", i)
		obj4_2.RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListNetworkIpv6 = d.Get(prefix4_2 + "network_ipv6").(string)
		obj4_2.RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListRouteMap = d.Get(prefix4_2 + "route_map").(string)
		obj4_2.RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListBackdoor = d.Get(prefix4_2 + "backdoor").(int)
		obj4_2.RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListDescription = d.Get(prefix4_2 + "description").(string)
		obj4_2.RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListCommValue = d.Get(prefix4_2 + "comm_value").(string)
		obj4.RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListNetworkIpv6 = append(obj4.RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListNetworkIpv6, obj4_2)
	}

	c.RouterBgpAddressFamilyIpv6InstanceNetworkSynchronization = obj4

	var obj5 go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighbor
	prefix5 := "neighbor.0."

	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListCount := d.Get(prefix5 + "peer_group_neighbor_list.#").(int)
	obj5.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListPeerGroup = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborList, 0, RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListCount)

	for i := 0; i < RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListCount; i++ {
		var obj5_1 go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborList
		prefix5_1 := prefix5 + fmt.Sprintf("peer_group_neighbor_list.%d.", i)
		obj5_1.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListPeerGroup = d.Get(prefix5_1 + "peer_group").(string)
		obj5_1.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListActivate = d.Get(prefix5_1 + "activate").(int)
		obj5_1.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListAllowasIn = d.Get(prefix5_1 + "allowas_in").(int)
		obj5_1.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListAllowasInCount = d.Get(prefix5_1 + "allowas_in_count").(int)
		obj5_1.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListMaximumPrefix = d.Get(prefix5_1 + "maximum_prefix").(int)
		obj5_1.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListMaximumPrefixThres = d.Get(prefix5_1 + "maximum_prefix_thres").(int)
		obj5_1.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNextHopSelf = d.Get(prefix5_1 + "next_hop_self").(int)
		obj5_1.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListRemovePrivateAs = d.Get(prefix5_1 + "remove_private_as").(int)

		RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNeighborRouteMapListsCount := d.Get(prefix5_1 + "neighbor_route_map_lists.#").(int)
		obj5_1.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNeighborRouteMapLists, 0, RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNeighborRouteMapListsCount)

		for i := 0; i < RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNeighborRouteMapListsCount; i++ {
			var obj5_1_1 go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNeighborRouteMapLists
			prefix5_1_1 := prefix5_1 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj5_1_1.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap = d.Get(prefix5_1_1 + "nbr_route_map").(string)
			obj5_1_1.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNeighborRouteMapListsNbrRmapDirection = d.Get(prefix5_1_1 + "nbr_rmap_direction").(string)
			obj5_1.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap = append(obj5_1.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap, obj5_1_1)
		}

		obj5_1.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListInbound = d.Get(prefix5_1 + "inbound").(int)
		obj5_1.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListWeight = d.Get(prefix5_1 + "weight").(int)
		obj5.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListPeerGroup = append(obj5.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListPeerGroup, obj5_1)
	}

	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListCount := d.Get(prefix5 + "ipv4_neighbor_list.#").(int)
	obj5.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborIpv4 = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborList, 0, RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListCount)

	for i := 0; i < RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListCount; i++ {
		var obj5_2 go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborList
		prefix5_2 := prefix5 + fmt.Sprintf("ipv4_neighbor_list.%d.", i)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborIpv4 = d.Get(prefix5_2 + "neighbor_ipv4").(string)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListPeerGroupName = d.Get(prefix5_2 + "peer_group_name").(string)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListActivate = d.Get(prefix5_2 + "activate").(int)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListAllowasIn = d.Get(prefix5_2 + "allowas_in").(int)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListAllowasInCount = d.Get(prefix5_2 + "allowas_in_count").(int)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListPrefixListDirection = d.Get(prefix5_2 + "prefix_list_direction").(string)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListGracefulRestart = d.Get(prefix5_2 + "graceful_restart").(int)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDefaultOriginate = d.Get(prefix5_2 + "default_originate").(int)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListRouteMap = d.Get(prefix5_2 + "route_map").(string)

		RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDistributeListsCount := d.Get(prefix5_2 + "distribute_lists.#").(int)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDistributeListsDistributeList = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDistributeLists, 0, RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDistributeListsCount)

		for i := 0; i < RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDistributeListsCount; i++ {
			var obj5_2_1 go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDistributeLists
			prefix5_2_1 := prefix5_2 + fmt.Sprintf("distribute_lists.%d.", i)
			obj5_2_1.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDistributeListsDistributeList = d.Get(prefix5_2_1 + "distribute_list").(string)
			obj5_2_1.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDistributeListsDistributeListDirection = d.Get(prefix5_2_1 + "distribute_list_direction").(string)
			obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDistributeListsDistributeList = append(obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDistributeListsDistributeList, obj5_2_1)
		}

		RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborFilterListsCount := d.Get(prefix5_2 + "neighbor_filter_lists.#").(int)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborFilterListsFilterList = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborFilterLists, 0, RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborFilterListsCount)

		for i := 0; i < RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborFilterListsCount; i++ {
			var obj5_2_2 go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborFilterLists
			prefix5_2_2 := prefix5_2 + fmt.Sprintf("neighbor_filter_lists.%d.", i)
			obj5_2_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborFilterListsFilterList = d.Get(prefix5_2_2 + "filter_list").(string)
			obj5_2_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborFilterListsFilterListDirection = d.Get(prefix5_2_2 + "filter_list_direction").(string)
			obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborFilterListsFilterList = append(obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborFilterListsFilterList, obj5_2_2)
		}

		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListMaximumPrefix = d.Get(prefix5_2 + "maximum_prefix").(int)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListMaximumPrefixThres = d.Get(prefix5_2 + "maximum_prefix_thres").(int)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListRestartMin = d.Get(prefix5_2 + "restart_min").(int)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNextHopSelf = d.Get(prefix5_2 + "next_hop_self").(int)

		RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborPrefixListsCount := d.Get(prefix5_2 + "neighbor_prefix_lists.#").(int)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborPrefixLists, 0, RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborPrefixListsCount)

		for i := 0; i < RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborPrefixListsCount; i++ {
			var obj5_2_3 go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborPrefixLists
			prefix5_2_3 := prefix5_2 + fmt.Sprintf("neighbor_prefix_lists.%d.", i)
			obj5_2_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList = d.Get(prefix5_2_3 + "nbr_prefix_list").(string)
			obj5_2_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborPrefixListsNbrPrefixListDirection = d.Get(prefix5_2_3 + "nbr_prefix_list_direction").(string)
			obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList = append(obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList, obj5_2_3)
		}

		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListRemovePrivateAs = d.Get(prefix5_2 + "remove_private_as").(int)

		RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborRouteMapListsCount := d.Get(prefix5_2 + "neighbor_route_map_lists.#").(int)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborRouteMapLists, 0, RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborRouteMapListsCount)

		for i := 0; i < RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborRouteMapListsCount; i++ {
			var obj5_2_4 go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborRouteMapLists
			prefix5_2_4 := prefix5_2 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj5_2_4.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap = d.Get(prefix5_2_4 + "nbr_route_map").(string)
			obj5_2_4.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborRouteMapListsNbrRmapDirection = d.Get(prefix5_2_4 + "nbr_rmap_direction").(string)
			obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap = append(obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap, obj5_2_4)
		}

		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListSendCommunityVal = d.Get(prefix5_2 + "send_community_val").(string)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListInbound = d.Get(prefix5_2 + "inbound").(int)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListUnsuppressMap = d.Get(prefix5_2 + "unsuppress_map").(string)
		obj5_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListWeight = d.Get(prefix5_2 + "weight").(int)
		obj5.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborIpv4 = append(obj5.RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborIpv4, obj5_2)
	}

	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListCount := d.Get(prefix5 + "ipv6_neighbor_list.#").(int)
	obj5.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborIpv6 = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborList, 0, RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListCount)

	for i := 0; i < RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListCount; i++ {
		var obj5_3 go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborList
		prefix5_3 := prefix5 + fmt.Sprintf("ipv6_neighbor_list.%d.", i)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborIpv6 = d.Get(prefix5_3 + "neighbor_ipv6").(string)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListPeerGroupName = d.Get(prefix5_3 + "peer_group_name").(string)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListActivate = d.Get(prefix5_3 + "activate").(int)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListAllowasIn = d.Get(prefix5_3 + "allowas_in").(int)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListAllowasInCount = d.Get(prefix5_3 + "allowas_in_count").(int)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListPrefixListDirection = d.Get(prefix5_3 + "prefix_list_direction").(string)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListGracefulRestart = d.Get(prefix5_3 + "graceful_restart").(int)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDefaultOriginate = d.Get(prefix5_3 + "default_originate").(int)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListRouteMap = d.Get(prefix5_3 + "route_map").(string)

		RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDistributeListsCount := d.Get(prefix5_3 + "distribute_lists.#").(int)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDistributeListsDistributeList = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDistributeLists, 0, RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDistributeListsCount)

		for i := 0; i < RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDistributeListsCount; i++ {
			var obj5_3_1 go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDistributeLists
			prefix5_3_1 := prefix5_3 + fmt.Sprintf("distribute_lists.%d.", i)
			obj5_3_1.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDistributeListsDistributeList = d.Get(prefix5_3_1 + "distribute_list").(string)
			obj5_3_1.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDistributeListsDistributeListDirection = d.Get(prefix5_3_1 + "distribute_list_direction").(string)
			obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDistributeListsDistributeList = append(obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDistributeListsDistributeList, obj5_3_1)
		}

		RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborFilterListsCount := d.Get(prefix5_3 + "neighbor_filter_lists.#").(int)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborFilterListsFilterList = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborFilterLists, 0, RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborFilterListsCount)

		for i := 0; i < RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborFilterListsCount; i++ {
			var obj5_3_2 go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborFilterLists
			prefix5_3_2 := prefix5_3 + fmt.Sprintf("neighbor_filter_lists.%d.", i)
			obj5_3_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborFilterListsFilterList = d.Get(prefix5_3_2 + "filter_list").(string)
			obj5_3_2.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborFilterListsFilterListDirection = d.Get(prefix5_3_2 + "filter_list_direction").(string)
			obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborFilterListsFilterList = append(obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborFilterListsFilterList, obj5_3_2)
		}

		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListMaximumPrefix = d.Get(prefix5_3 + "maximum_prefix").(int)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListMaximumPrefixThres = d.Get(prefix5_3 + "maximum_prefix_thres").(int)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListRestartMin = d.Get(prefix5_3 + "restart_min").(int)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNextHopSelf = d.Get(prefix5_3 + "next_hop_self").(int)

		RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborPrefixListsCount := d.Get(prefix5_3 + "neighbor_prefix_lists.#").(int)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborPrefixLists, 0, RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborPrefixListsCount)

		for i := 0; i < RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborPrefixListsCount; i++ {
			var obj5_3_3 go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborPrefixLists
			prefix5_3_3 := prefix5_3 + fmt.Sprintf("neighbor_prefix_lists.%d.", i)
			obj5_3_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList = d.Get(prefix5_3_3 + "nbr_prefix_list").(string)
			obj5_3_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborPrefixListsNbrPrefixListDirection = d.Get(prefix5_3_3 + "nbr_prefix_list_direction").(string)
			obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList = append(obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList, obj5_3_3)
		}

		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListRemovePrivateAs = d.Get(prefix5_3 + "remove_private_as").(int)

		RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborRouteMapListsCount := d.Get(prefix5_3 + "neighbor_route_map_lists.#").(int)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborRouteMapLists, 0, RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborRouteMapListsCount)

		for i := 0; i < RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborRouteMapListsCount; i++ {
			var obj5_3_4 go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborRouteMapLists
			prefix5_3_4 := prefix5_3 + fmt.Sprintf("neighbor_route_map_lists.%d.", i)
			obj5_3_4.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap = d.Get(prefix5_3_4 + "nbr_route_map").(string)
			obj5_3_4.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborRouteMapListsNbrRmapDirection = d.Get(prefix5_3_4 + "nbr_rmap_direction").(string)
			obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap = append(obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap, obj5_3_4)
		}

		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListSendCommunityVal = d.Get(prefix5_3 + "send_community_val").(string)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListInbound = d.Get(prefix5_3 + "inbound").(int)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListUnsuppressMap = d.Get(prefix5_3 + "unsuppress_map").(string)
		obj5_3.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListWeight = d.Get(prefix5_3 + "weight").(int)
		obj5.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborIpv6 = append(obj5.RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborIpv6, obj5_3)
	}

	RouterBgpAddressFamilyIpv6InstanceNeighborEthernetNeighborIpv6ListCount := d.Get(prefix5 + "ethernet_neighbor_ipv6_list.#").(int)
	obj5.RouterBgpAddressFamilyIpv6InstanceNeighborEthernetNeighborIpv6ListEthernet = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborEthernetNeighborIpv6List, 0, RouterBgpAddressFamilyIpv6InstanceNeighborEthernetNeighborIpv6ListCount)

	for i := 0; i < RouterBgpAddressFamilyIpv6InstanceNeighborEthernetNeighborIpv6ListCount; i++ {
		var obj5_4 go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborEthernetNeighborIpv6List
		prefix5_4 := prefix5 + fmt.Sprintf("ethernet_neighbor_ipv6_list.%d.", i)
		obj5_4.RouterBgpAddressFamilyIpv6InstanceNeighborEthernetNeighborIpv6ListEthernet = d.Get(prefix5_4 + "ethernet").(int)
		obj5_4.RouterBgpAddressFamilyIpv6InstanceNeighborEthernetNeighborIpv6ListPeerGroupName = d.Get(prefix5_4 + "peer_group_name").(string)
		obj5.RouterBgpAddressFamilyIpv6InstanceNeighborEthernetNeighborIpv6ListEthernet = append(obj5.RouterBgpAddressFamilyIpv6InstanceNeighborEthernetNeighborIpv6ListEthernet, obj5_4)
	}

	RouterBgpAddressFamilyIpv6InstanceNeighborVeNeighborIpv6ListCount := d.Get(prefix5 + "ve_neighbor_ipv6_list.#").(int)
	obj5.RouterBgpAddressFamilyIpv6InstanceNeighborVeNeighborIpv6ListVe = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborVeNeighborIpv6List, 0, RouterBgpAddressFamilyIpv6InstanceNeighborVeNeighborIpv6ListCount)

	for i := 0; i < RouterBgpAddressFamilyIpv6InstanceNeighborVeNeighborIpv6ListCount; i++ {
		var obj5_5 go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborVeNeighborIpv6List
		prefix5_5 := prefix5 + fmt.Sprintf("ve_neighbor_ipv6_list.%d.", i)
		obj5_5.RouterBgpAddressFamilyIpv6InstanceNeighborVeNeighborIpv6ListVe = d.Get(prefix5_5 + "ve").(int)
		obj5_5.RouterBgpAddressFamilyIpv6InstanceNeighborVeNeighborIpv6ListPeerGroupName = d.Get(prefix5_5 + "peer_group_name").(string)
		obj5.RouterBgpAddressFamilyIpv6InstanceNeighborVeNeighborIpv6ListVe = append(obj5.RouterBgpAddressFamilyIpv6InstanceNeighborVeNeighborIpv6ListVe, obj5_5)
	}

	RouterBgpAddressFamilyIpv6InstanceNeighborTrunkNeighborIpv6ListCount := d.Get(prefix5 + "trunk_neighbor_ipv6_list.#").(int)
	obj5.RouterBgpAddressFamilyIpv6InstanceNeighborTrunkNeighborIpv6ListTrunk = make([]go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborTrunkNeighborIpv6List, 0, RouterBgpAddressFamilyIpv6InstanceNeighborTrunkNeighborIpv6ListCount)

	for i := 0; i < RouterBgpAddressFamilyIpv6InstanceNeighborTrunkNeighborIpv6ListCount; i++ {
		var obj5_6 go_thunder.RouterBgpAddressFamilyIpv6InstanceNeighborTrunkNeighborIpv6List
		prefix5_6 := prefix5 + fmt.Sprintf("trunk_neighbor_ipv6_list.%d.", i)
		obj5_6.RouterBgpAddressFamilyIpv6InstanceNeighborTrunkNeighborIpv6ListTrunk = d.Get(prefix5_6 + "trunk").(int)
		obj5_6.RouterBgpAddressFamilyIpv6InstanceNeighborTrunkNeighborIpv6ListPeerGroupName = d.Get(prefix5_6 + "peer_group_name").(string)
		obj5.RouterBgpAddressFamilyIpv6InstanceNeighborTrunkNeighborIpv6ListTrunk = append(obj5.RouterBgpAddressFamilyIpv6InstanceNeighborTrunkNeighborIpv6ListTrunk, obj5_6)
	}

	c.RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborList = obj5

	var obj6 go_thunder.RouterBgpAddressFamilyIpv6InstanceRedistribute
	prefix6 := "redistribute.0."

	var obj6_1 go_thunder.RouterBgpAddressFamilyIpv6InstanceRedistributeConnectedCfg
	prefix6_1 := prefix6 + "connected_cfg.0."
	obj6_1.RouterBgpAddressFamilyIpv6InstanceRedistributeConnectedCfgConnected = d.Get(prefix6_1 + "connected").(int)
	obj6_1.RouterBgpAddressFamilyIpv6InstanceRedistributeConnectedCfgRouteMap = d.Get(prefix6_1 + "route_map").(string)

	obj6.RouterBgpAddressFamilyIpv6InstanceRedistributeConnectedCfgConnected = obj6_1

	var obj6_2 go_thunder.RouterBgpAddressFamilyIpv6InstanceRedistributeFloatingIPCfg
	prefix6_2 := prefix6 + "floating_ip_cfg.0."
	obj6_2.RouterBgpAddressFamilyIpv6InstanceRedistributeFloatingIPCfgFloatingIP = d.Get(prefix6_2 + "floating_ip").(int)
	obj6_2.RouterBgpAddressFamilyIpv6InstanceRedistributeFloatingIPCfgRouteMap = d.Get(prefix6_2 + "route_map").(string)

	obj6.RouterBgpAddressFamilyIpv6InstanceRedistributeFloatingIPCfgFloatingIP = obj6_2

	var obj6_3 go_thunder.RouterBgpAddressFamilyIpv6InstanceRedistributeNat64Cfg
	prefix6_3 := prefix6 + "nat64_cfg.0."
	obj6_3.RouterBgpAddressFamilyIpv6InstanceRedistributeNat64CfgNat64 = d.Get(prefix6_3 + "nat64").(int)
	obj6_3.RouterBgpAddressFamilyIpv6InstanceRedistributeNat64CfgRouteMap = d.Get(prefix6_3 + "route_map").(string)

	obj6.RouterBgpAddressFamilyIpv6InstanceRedistributeNat64CfgNat64 = obj6_3

	var obj6_4 go_thunder.RouterBgpAddressFamilyIpv6InstanceRedistributeNatMapCfg
	prefix6_4 := prefix6 + "nat_map_cfg.0."
	obj6_4.RouterBgpAddressFamilyIpv6InstanceRedistributeNatMapCfgNatMap = d.Get(prefix6_4 + "nat_map").(int)
	obj6_4.RouterBgpAddressFamilyIpv6InstanceRedistributeNatMapCfgRouteMap = d.Get(prefix6_4 + "route_map").(string)

	obj6.RouterBgpAddressFamilyIpv6InstanceRedistributeNatMapCfgNatMap = obj6_4

	var obj6_5 go_thunder.RouterBgpAddressFamilyIpv6InstanceRedistributeLw4O6Cfg
	prefix6_5 := prefix6 + "lw4o6_cfg.0."
	obj6_5.RouterBgpAddressFamilyIpv6InstanceRedistributeLw4O6CfgLw4O6 = d.Get(prefix6_5 + "lw4o6").(int)
	obj6_5.RouterBgpAddressFamilyIpv6InstanceRedistributeLw4O6CfgRouteMap = d.Get(prefix6_5 + "route_map").(string)

	obj6.RouterBgpAddressFamilyIpv6InstanceRedistributeLw4O6CfgLw4O6 = obj6_5

	var obj6_6 go_thunder.RouterBgpAddressFamilyIpv6InstanceRedistributeStaticNatCfg
	prefix6_6 := prefix6 + "static_nat_cfg.0."
	obj6_6.RouterBgpAddressFamilyIpv6InstanceRedistributeStaticNatCfgStaticNat = d.Get(prefix6_6 + "static_nat").(int)
	obj6_6.RouterBgpAddressFamilyIpv6InstanceRedistributeStaticNatCfgRouteMap = d.Get(prefix6_6 + "route_map").(string)

	obj6.RouterBgpAddressFamilyIpv6InstanceRedistributeStaticNatCfgStaticNat = obj6_6

	var obj6_7 go_thunder.RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatCfg
	prefix6_7 := prefix6 + "ip_nat_cfg.0."
	obj6_7.RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatCfgIPNat = d.Get(prefix6_7 + "ip_nat").(int)
	obj6_7.RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatCfgRouteMap = d.Get(prefix6_7 + "route_map").(string)

	obj6.RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatCfgIPNat = obj6_7

	var obj6_8 go_thunder.RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatListCfg
	prefix6_8 := prefix6 + "ip_nat_list_cfg.0."
	obj6_8.RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatListCfgIPNatList = d.Get(prefix6_8 + "ip_nat_list").(int)
	obj6_8.RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatListCfgRouteMap = d.Get(prefix6_8 + "route_map").(string)

	obj6.RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatListCfgIPNatList = obj6_8

	var obj6_9 go_thunder.RouterBgpAddressFamilyIpv6InstanceRedistributeIsisCfg
	prefix6_9 := prefix6 + "isis_cfg.0."
	obj6_9.RouterBgpAddressFamilyIpv6InstanceRedistributeIsisCfgIsis = d.Get(prefix6_9 + "isis").(int)
	obj6_9.RouterBgpAddressFamilyIpv6InstanceRedistributeIsisCfgRouteMap = d.Get(prefix6_9 + "route_map").(string)

	obj6.RouterBgpAddressFamilyIpv6InstanceRedistributeIsisCfgIsis = obj6_9

	var obj6_10 go_thunder.RouterBgpAddressFamilyIpv6InstanceRedistributeOspfCfg
	prefix6_10 := prefix6 + "ospf_cfg.0."
	obj6_10.RouterBgpAddressFamilyIpv6InstanceRedistributeOspfCfgOspf = d.Get(prefix6_10 + "ospf").(int)
	obj6_10.RouterBgpAddressFamilyIpv6InstanceRedistributeOspfCfgRouteMap = d.Get(prefix6_10 + "route_map").(string)

	obj6.RouterBgpAddressFamilyIpv6InstanceRedistributeOspfCfgOspf = obj6_10

	var obj6_11 go_thunder.RouterBgpAddressFamilyIpv6InstanceRedistributeRipCfg
	prefix6_11 := prefix6 + "rip_cfg.0."
	obj6_11.RouterBgpAddressFamilyIpv6InstanceRedistributeRipCfgRip = d.Get(prefix6_11 + "rip").(int)
	obj6_11.RouterBgpAddressFamilyIpv6InstanceRedistributeRipCfgRouteMap = d.Get(prefix6_11 + "route_map").(string)

	obj6.RouterBgpAddressFamilyIpv6InstanceRedistributeRipCfgRip = obj6_11

	var obj6_12 go_thunder.RouterBgpAddressFamilyIpv6InstanceRedistributeStaticCfg
	prefix6_12 := prefix6 + "static_cfg.0."
	obj6_12.RouterBgpAddressFamilyIpv6InstanceRedistributeStaticCfgStatic = d.Get(prefix6_12 + "static").(int)
	obj6_12.RouterBgpAddressFamilyIpv6InstanceRedistributeStaticCfgRouteMap = d.Get(prefix6_12 + "route_map").(string)

	obj6.RouterBgpAddressFamilyIpv6InstanceRedistributeStaticCfgStatic = obj6_12

	var obj6_13 go_thunder.RouterBgpAddressFamilyIpv6InstanceRedistributeVip
	prefix6_13 := prefix6 + "vip.0."

	var obj6_13_1 go_thunder.RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyFlaggedCfg
	prefix6_13_1 := prefix6_13 + "only_flagged_cfg.0."
	obj6_13_1.RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyFlaggedCfgOnlyFlagged = d.Get(prefix6_13_1 + "only_flagged").(int)
	obj6_13_1.RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyFlaggedCfgRouteMap = d.Get(prefix6_13_1 + "route_map").(string)

	obj6_13.RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyFlaggedCfgOnlyFlagged = obj6_13_1

	var obj6_13_2 go_thunder.RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyNotFlaggedCfg
	prefix6_13_2 := prefix6_13 + "only_not_flagged_cfg.0."
	obj6_13_2.RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyNotFlaggedCfgOnlyNotFlagged = d.Get(prefix6_13_2 + "only_not_flagged").(int)
	obj6_13_2.RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyNotFlaggedCfgRouteMap = d.Get(prefix6_13_2 + "route_map").(string)

	obj6_13.RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyNotFlaggedCfgOnlyNotFlagged = obj6_13_2

	obj6.RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyFlaggedCfg = obj6_13

	c.RouterBgpAddressFamilyIpv6InstanceRedistributeConnectedCfg = obj6

	vc.RouterBgpAddressFamilyIpv6InstanceBgp = c
	return vc
}
