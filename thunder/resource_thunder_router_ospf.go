package thunder

//Thunder resource RouterOspf

import (
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceRouterOspf() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspfCreate,
		Update: resourceRouterOspfUpdate,
		Read:   resourceRouterOspfRead,
		Delete: resourceRouterOspfDelete,
		Schema: map[string]*schema.Schema{
			"process_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"auto_cost_reference_bandwidth": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bfd_all_interfaces": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rfc1583_compatible": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"default_metric": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"distance": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"distance_value": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"distance_ospf": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"distance_external": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"distance_inter_area": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"distance_intra_area": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"distribute_internal_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"di_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"di_area_ipv4": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"di_area_num": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"di_cost": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"distribute_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"direction": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ospf_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"option": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ha_standby_extra_cost": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"extra_cost": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"group": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"host_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"area_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"area_ipv4": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"area_num": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"cost": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"log_adjacency_changes_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"max_concurrent_dd": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"maximum_area": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"neighbor_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cost": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"poll_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"priority": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"network_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_ipv4": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"network_ipv4_mask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"network_ipv4_cidr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"network_area": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_area_ipv4": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"network_area_num": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"instance_value": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"ospf_1": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"abr_type": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"option": {
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
			"router_id": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"overflow": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"database": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"count": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"limit": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"db_external": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"recovery_time": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"passive_interface": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"loopback_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"loopback": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"loopback_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"trunk_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trunk": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"trunk_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ve_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ve": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ve_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"tunnel_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tunnel": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"tunnel_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"lif_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"lif": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"lif_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"eth_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ethernet": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"eth_address": {
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
			"summary_address_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"summary_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"not_advertise": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"tag": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"timers": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"spf": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"exp": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"min_delay": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"max_delay": {
													Type:        schema.TypeInt,
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
			"default_information": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"originate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"always": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"metric": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"metric_type": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"route_map": {
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
			"area_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"area_ipv4": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"area_num": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"auth_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"message_digest": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"filter_lists": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"filter_list": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"acl_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"acl_direction": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"plist_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"plist_direction": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"nssa_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"nssa": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"no_redistribution": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"no_summary": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"translator_role": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"default_information_originate": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"metric": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"metric_type": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"default_cost": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"range_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"area_range_prefix": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"option": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"shortcut": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"stub_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"stub": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"no_summary": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"virtual_link_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"virtual_link_ip_addr": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"bfd": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"hello_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dead_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"retransmit_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"transmit_delay": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"virtual_link_authentication": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"virtual_link_auth_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"authentication_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"message_digest_key": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"md5": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
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
			"redistribute": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"redist_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"metric": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"metric_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"route_map": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"tag": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ospf_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ospf": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"process_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"metric_ospf": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"metric_type_ospf": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"route_map_ospf": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"tag_ospf": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ip_nat": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"metric_ip_nat": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"metric_type_ip_nat": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"route_map_ip_nat": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"tag_ip_nat": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ip_nat_floating_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_nat_prefix": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ip_nat_floating_ip_forward": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"vip_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type_vip": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"metric_vip": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"metric_type_vip": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"route_map_vip": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"tag_vip": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"vip_floating_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vip_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"vip_floating_ip_forward": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
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
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sequence": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceRouterOspfCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterOspf (Inside resourceRouterOspfCreate) ")
		name1 := d.Get("process_id").(int)
		data := dataToRouterOspf(d)
		logger.Println("[INFO] received formatted data from method data to RouterOspf --")
		d.SetId(strconv.Itoa(name1))
		go_thunder.PostRouterOspf(client.Token, data, client.Host)

		return resourceRouterOspfRead(d, meta)

	}
	return nil
}

func resourceRouterOspfRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterOspf (Inside resourceRouterOspfRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterOspf(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceRouterOspfUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying RouterOspf   (Inside resourceRouterOspfUpdate) ")
		data := dataToRouterOspf(d)
		logger.Println("[INFO] received formatted data from method data to RouterOspf ")
		go_thunder.PutRouterOspf(client.Token, name1, data, client.Host)

		return resourceRouterOspfRead(d, meta)

	}
	return nil
}

func resourceRouterOspfDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceRouterOspfDelete) " + name1)
		err := go_thunder.DeleteRouterOspf(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToRouterOspf(d *schema.ResourceData) go_thunder.RouterOspf {
	var vc go_thunder.RouterOspf
	var c go_thunder.RouterOspfInstance
	c.ProcessID = d.Get("process_id").(int)
	c.AutoCostReferenceBandwidth = d.Get("auto_cost_reference_bandwidth").(int)
	c.BfdAllInterfaces = d.Get("bfd_all_interfaces").(int)
	c.Rfc1583Compatible = d.Get("rfc1583_compatible").(int)
	c.DefaultMetric = d.Get("default_metric").(int)

	var obj1 go_thunder.RouterOspfDistance
	prefix1 := "distance.0."
	obj1.DistanceValue = d.Get(prefix1 + "distance_value").(int)

	var obj1_1 go_thunder.RouterOspfDistanceOspf
	prefix1_1 := prefix1 + "distance_ospf.0."
	obj1_1.DistanceExternal = d.Get(prefix1_1 + "distance_external").(int)
	obj1_1.DistanceInterArea = d.Get(prefix1_1 + "distance_inter_area").(int)
	obj1_1.DistanceIntraArea = d.Get(prefix1_1 + "distance_intra_area").(int)

	obj1.DistanceExternal = obj1_1

	c.DistanceValue = obj1

	DistributeInternalListCount := d.Get("distribute_internal_list.#").(int)
	c.DiType = make([]go_thunder.RouterOspfDistributeInternalList, 0, DistributeInternalListCount)

	for i := 0; i < DistributeInternalListCount; i++ {
		var obj2 go_thunder.RouterOspfDistributeInternalList
		prefix2 := fmt.Sprintf("distribute_internal_list.%d.", i)
		obj2.DiType = d.Get(prefix2 + "di_type").(string)
		obj2.DiAreaIpv4 = d.Get(prefix2 + "di_area_ipv4").(string)
		obj2.DiAreaNum = d.Get(prefix2 + "di_area_num").(int)
		obj2.DiCost = d.Get(prefix2 + "di_cost").(int)
		c.DiType = append(c.DiType, obj2)
	}

	DistributeListsCount := d.Get("distribute_lists.#").(int)
	c.Direction = make([]go_thunder.RouterOspfDistributeLists, 0, DistributeListsCount)

	for i := 0; i < DistributeListsCount; i++ {
		var obj3 go_thunder.RouterOspfDistributeLists
		prefix3 := fmt.Sprintf("distribute_lists.%d.", i)
		obj3.Value = d.Get(prefix3 + "value").(string)
		obj3.Direction = d.Get(prefix3 + "direction").(string)
		obj3.Protocol = d.Get(prefix3 + "protocol").(string)
		obj3.OspfID = d.Get(prefix3 + "ospf_id").(int)
		obj3.Option = d.Get(prefix3 + "option").(string)
		c.Direction = append(c.Direction, obj3)
	}

	HaStandbyExtraCostCount := d.Get("ha_standby_extra_cost.#").(int)
	c.ExtraCost = make([]go_thunder.RouterOspfHaStandbyExtraCost, 0, HaStandbyExtraCostCount)

	for i := 0; i < HaStandbyExtraCostCount; i++ {
		var obj4 go_thunder.RouterOspfHaStandbyExtraCost
		prefix4 := fmt.Sprintf("ha_standby_extra_cost.%d.", i)
		obj4.ExtraCost = d.Get(prefix4 + "extra_cost").(int)
		obj4.Group = d.Get(prefix4 + "group").(int)
		c.ExtraCost = append(c.ExtraCost, obj4)
	}

	HostListCount := d.Get("host_list.#").(int)
	c.HostAddress = make([]go_thunder.RouterOspfHostList, 0, HostListCount)

	for i := 0; i < HostListCount; i++ {
		var obj5 go_thunder.RouterOspfHostList
		prefix5 := fmt.Sprintf("host_list.%d.", i)
		obj5.HostAddress = d.Get(prefix5 + "host_address").(string)

		var obj5_1 go_thunder.RouterOspfAreaCfg
		prefix5_1 := prefix5 + "area_cfg.0."
		obj5_1.AreaIpv4 = d.Get(prefix5_1 + "area_ipv4").(string)
		obj5_1.AreaNum = d.Get(prefix5_1 + "area_num").(int)
		obj5_1.Cost = d.Get(prefix5_1 + "cost").(int)

		obj5.AreaIpv4 = obj5_1

		c.HostAddress = append(c.HostAddress, obj5)
	}

	var obj6 go_thunder.RouterOspfLogAdjacencyChangesCfg
	prefix6 := "log_adjacency_changes_cfg.0."
	obj6.State = d.Get(prefix6 + "state").(string)

	c.State = obj6

	c.MaxConcurrentDd = d.Get("max_concurrent_dd").(int)
	c.MaximumArea = d.Get("maximum_area").(int)

	NeighborListCount := d.Get("neighbor_list.#").(int)
	c.Address = make([]go_thunder.RouterOspfNeighborList, 0, NeighborListCount)

	for i := 0; i < NeighborListCount; i++ {
		var obj7 go_thunder.RouterOspfNeighborList
		prefix7 := fmt.Sprintf("neighbor_list.%d.", i)
		obj7.Address = d.Get(prefix7 + "address").(string)
		obj7.Cost = d.Get(prefix7 + "cost").(int)
		obj7.PollInterval = d.Get(prefix7 + "poll_interval").(int)
		obj7.Priority = d.Get(prefix7 + "priority").(int)
		c.Address = append(c.Address, obj7)
	}

	NetworkListCount := d.Get("network_list.#").(int)
	c.NetworkIpv4 = make([]go_thunder.RouterOspfNetworkList, 0, NetworkListCount)

	for i := 0; i < NetworkListCount; i++ {
		var obj8 go_thunder.RouterOspfNetworkList
		prefix8 := fmt.Sprintf("network_list.%d.", i)
		obj8.NetworkIpv4 = d.Get(prefix8 + "network_ipv4").(string)
		obj8.NetworkIpv4Mask = d.Get(prefix8 + "network_ipv4_mask").(string)
		obj8.NetworkIpv4Cidr = d.Get(prefix8 + "network_ipv4_cidr").(string)

		var obj8_1 go_thunder.RouterOspfNetworkArea
		prefix8_1 := prefix8 + "network_area.0."
		obj8_1.NetworkAreaIpv4 = d.Get(prefix8_1 + "network_area_ipv4").(string)
		obj8_1.NetworkAreaNum = d.Get(prefix8_1 + "network_area_num").(int)
		obj8_1.InstanceValue = d.Get(prefix8_1 + "instance_value").(int)

		obj8.NetworkAreaIpv4 = obj8_1

		c.NetworkIpv4 = append(c.NetworkIpv4, obj8)
	}

	var obj9 go_thunder.RouterOspfOspf1
	prefix9 := "ospf_1.0."

	var obj9_1 go_thunder.RouterOspfAbrType
	prefix9_1 := prefix9 + "abr_type.0."
	obj9_1.Option = d.Get(prefix9_1 + "option").(string)

	obj9.Option = obj9_1

	c.AbrType = obj9

	var obj10 go_thunder.RouterOspfRouterID
	prefix10 := "router_id.0."
	obj10.Value = d.Get(prefix10 + "value").(string)

	c.Value = obj10

	var obj11 go_thunder.RouterOspfOverflow
	prefix11 := "overflow.0."

	var obj11_1 go_thunder.RouterOspfDatabase
	prefix11_1 := prefix11 + "database.0."
	obj11_1.Count = d.Get(prefix11_1 + "count").(int)
	obj11_1.Limit = d.Get(prefix11_1 + "limit").(string)
	obj11_1.DbExternal = d.Get(prefix11_1 + "db_external").(int)
	obj11_1.RecoveryTime = d.Get(prefix11_1 + "recovery_time").(int)

	obj11.Count = obj11_1

	c.Database = obj11

	var obj12 go_thunder.RouterOspfPassiveInterface
	prefix12 := "passive_interface.0."

	LoopbackCfgCount := d.Get(prefix12 + "loopback_cfg.#").(int)
	obj12.Loopback = make([]go_thunder.RouterOspfLoopbackCfg, 0, LoopbackCfgCount)

	for i := 0; i < LoopbackCfgCount; i++ {
		var obj12_1 go_thunder.RouterOspfLoopbackCfg
		prefix12_1 := prefix12 + fmt.Sprintf("loopback_cfg.%d.", i)
		obj12_1.Loopback = d.Get(prefix12_1 + "loopback").(int)
		obj12_1.LoopbackAddress = d.Get(prefix12_1 + "loopback_address").(string)
		obj12.Loopback = append(obj12.Loopback, obj12_1)
	}

	TrunkCfgCount := d.Get(prefix12 + "trunk_cfg.#").(int)
	obj12.Trunk = make([]go_thunder.RouterOspfTrunkCfg, 0, TrunkCfgCount)

	for i := 0; i < TrunkCfgCount; i++ {
		var obj12_2 go_thunder.RouterOspfTrunkCfg
		prefix12_2 := prefix12 + fmt.Sprintf("trunk_cfg.%d.", i)
		obj12_2.Trunk = d.Get(prefix12_2 + "trunk").(int)
		obj12_2.TrunkAddress = d.Get(prefix12_2 + "trunk_address").(string)
		obj12.Trunk = append(obj12.Trunk, obj12_2)
	}

	VeCfgCount := d.Get(prefix12 + "ve_cfg.#").(int)
	obj12.Ve = make([]go_thunder.RouterOspfVeCfg, 0, VeCfgCount)

	for i := 0; i < VeCfgCount; i++ {
		var obj12_3 go_thunder.RouterOspfVeCfg
		prefix12_3 := prefix12 + fmt.Sprintf("ve_cfg.%d.", i)
		obj12_3.Ve = d.Get(prefix12_3 + "ve").(int)
		obj12_3.VeAddress = d.Get(prefix12_3 + "ve_address").(string)
		obj12.Ve = append(obj12.Ve, obj12_3)
	}

	TunnelCfgCount := d.Get(prefix12 + "tunnel_cfg.#").(int)
	obj12.Tunnel = make([]go_thunder.RouterOspfTunnelCfg, 0, TunnelCfgCount)

	for i := 0; i < TunnelCfgCount; i++ {
		var obj12_4 go_thunder.RouterOspfTunnelCfg
		prefix12_4 := prefix12 + fmt.Sprintf("tunnel_cfg.%d.", i)
		obj12_4.Tunnel = d.Get(prefix12_4 + "tunnel").(int)
		obj12_4.TunnelAddress = d.Get(prefix12_4 + "tunnel_address").(string)
		obj12.Tunnel = append(obj12.Tunnel, obj12_4)
	}

	LifCfgCount := d.Get(prefix12 + "lif_cfg.#").(int)
	obj12.Lif = make([]go_thunder.RouterOspfLifCfg, 0, LifCfgCount)

	for i := 0; i < LifCfgCount; i++ {
		var obj12_5 go_thunder.RouterOspfLifCfg
		prefix12_5 := prefix12 + fmt.Sprintf("lif_cfg.%d.", i)
		obj12_5.Lif = d.Get(prefix12_5 + "lif").(int)
		obj12_5.LifAddress = d.Get(prefix12_5 + "lif_address").(string)
		obj12.Lif = append(obj12.Lif, obj12_5)
	}

	EthCfgCount := d.Get(prefix12 + "eth_cfg.#").(int)
	obj12.Ethernet = make([]go_thunder.RouterOspfEthCfg, 0, EthCfgCount)

	for i := 0; i < EthCfgCount; i++ {
		var obj12_6 go_thunder.RouterOspfEthCfg
		prefix12_6 := prefix12 + fmt.Sprintf("eth_cfg.%d.", i)
		obj12_6.Ethernet = d.Get(prefix12_6 + "ethernet").(int)
		obj12_6.EthAddress = d.Get(prefix12_6 + "eth_address").(string)
		obj12.Ethernet = append(obj12.Ethernet, obj12_6)
	}

	c.LoopbackCfg = obj12

	SummaryAddressListCount := d.Get("summary_address_list.#").(int)
	c.SummaryAddress = make([]go_thunder.RouterOspfSummaryAddressList, 0, SummaryAddressListCount)

	for i := 0; i < SummaryAddressListCount; i++ {
		var obj13 go_thunder.RouterOspfSummaryAddressList
		prefix13 := fmt.Sprintf("summary_address_list.%d.", i)
		obj13.SummaryAddress = d.Get(prefix13 + "summary_address").(string)
		obj13.NotAdvertise = d.Get(prefix13 + "not_advertise").(int)
		obj13.Tag = d.Get(prefix13 + "tag").(int)
		c.SummaryAddress = append(c.SummaryAddress, obj13)
	}

	var obj14 go_thunder.RouterOspfTimers
	prefix14 := "timers.0."

	var obj14_1 go_thunder.RouterOspfSpf
	prefix14_1 := prefix14 + "spf.0."

	var obj14_1_1 go_thunder.RouterOspfExp
	prefix14_1_1 := prefix14_1 + "exp.0."
	obj14_1_1.MinDelay = d.Get(prefix14_1_1 + "min_delay").(int)
	obj14_1_1.MaxDelay = d.Get(prefix14_1_1 + "max_delay").(int)

	obj14_1.MinDelay = obj14_1_1

	obj14.Exp = obj14_1

	c.Spf = obj14

	c.UserTag = d.Get("user_tag").(string)

	var obj15 go_thunder.RouterOspfDefaultInformationOspf
	prefix15 := "default_information.0."
	obj15.Originate = d.Get(prefix15 + "originate").(int)
	obj15.Always = d.Get(prefix15 + "always").(int)
	obj15.Metric = d.Get(prefix15 + "metric").(int)
	obj15.MetricType = d.Get(prefix15 + "metric_type").(int)
	obj15.RouteMap = d.Get(prefix15 + "route_map").(string)

	c.Originate = obj15

	AreaListCount := d.Get("area_list.#").(int)
	c.AreaIpv4 = make([]go_thunder.RouterOspfAreaList, 0, AreaListCount)

	for i := 0; i < AreaListCount; i++ {
		var obj16 go_thunder.RouterOspfAreaList
		prefix16 := fmt.Sprintf("area_list.%d.", i)
		obj16.AreaIpv4 = d.Get(prefix16 + "area_ipv4").(string)
		obj16.AreaNum = d.Get(prefix16 + "area_num").(int)

		var obj16_1 go_thunder.RouterOspfAuthCfg
		prefix16_1 := prefix16 + "auth_cfg.0."
		obj16_1.Authentication = d.Get(prefix16_1 + "authentication").(int)
		obj16_1.MessageDigest = d.Get(prefix16_1 + "message_digest").(int)

		obj16.Authentication = obj16_1

		FilterListsCount := d.Get(prefix16 + "filter_lists.#").(int)
		obj16.FilterList = make([]go_thunder.RouterOspfFilterLists, 0, FilterListsCount)

		for i := 0; i < FilterListsCount; i++ {
			var obj16_2 go_thunder.RouterOspfFilterLists
			prefix16_2 := prefix16 + fmt.Sprintf("filter_lists.%d.", i)
			obj16_2.FilterList = d.Get(prefix16_2 + "filter_list").(int)
			obj16_2.AclName = d.Get(prefix16_2 + "acl_name").(string)
			obj16_2.AclDirection = d.Get(prefix16_2 + "acl_direction").(string)
			obj16_2.PlistName = d.Get(prefix16_2 + "plist_name").(string)
			obj16_2.PlistDirection = d.Get(prefix16_2 + "plist_direction").(string)
			obj16.FilterList = append(obj16.FilterList, obj16_2)
		}

		var obj16_3 go_thunder.RouterOspfNssaCfg
		prefix16_3 := prefix16 + "nssa_cfg.0."
		obj16_3.Nssa = d.Get(prefix16_3 + "nssa").(int)
		obj16_3.NoRedistribution = d.Get(prefix16_3 + "no_redistribution").(int)
		obj16_3.NoSummary = d.Get(prefix16_3 + "no_summary").(int)
		obj16_3.TranslatorRole = d.Get(prefix16_3 + "translator_role").(string)
		obj16_3.DefaultInformationOriginate = d.Get(prefix16_3 + "default_information_originate").(int)
		obj16_3.Metric = d.Get(prefix16_3 + "metric").(int)
		obj16_3.MetricType = d.Get(prefix16_3 + "metric_type").(int)

		obj16.Nssa = obj16_3

		obj16.DefaultCost = d.Get(prefix16 + "default_cost").(int)

		RangeListCount := d.Get(prefix16 + "range_list.#").(int)
		obj16.AreaRangePrefix = make([]go_thunder.RouterOspfRangeList, 0, RangeListCount)

		for i := 0; i < RangeListCount; i++ {
			var obj16_4 go_thunder.RouterOspfRangeList
			prefix16_4 := prefix16 + fmt.Sprintf("range_list.%d.", i)
			obj16_4.AreaRangePrefix = d.Get(prefix16_4 + "area_range_prefix").(string)
			obj16_4.Option = d.Get(prefix16_4 + "option").(string)
			obj16.AreaRangePrefix = append(obj16.AreaRangePrefix, obj16_4)
		}

		obj16.Shortcut = d.Get(prefix16 + "shortcut").(string)

		var obj16_5 go_thunder.RouterOspfStubCfg
		prefix16_5 := prefix16 + "stub_cfg.0."
		obj16_5.Stub = d.Get(prefix16_5 + "stub").(int)
		obj16_5.NoSummary = d.Get(prefix16_5 + "no_summary").(int)

		obj16.Stub = obj16_5

		VirtualLinkListCount := d.Get(prefix16 + "virtual_link_list.#").(int)
		obj16.VirtualLinkIPAddr = make([]go_thunder.RouterOspfVirtualLinkList, 0, VirtualLinkListCount)

		for i := 0; i < VirtualLinkListCount; i++ {
			var obj16_6 go_thunder.RouterOspfVirtualLinkList
			prefix16_6 := prefix16 + fmt.Sprintf("virtual_link_list.%d.", i)
			obj16_6.VirtualLinkIPAddr = d.Get(prefix16_6 + "virtual_link_ip_addr").(string)
			obj16_6.Bfd = d.Get(prefix16_6 + "bfd").(int)
			obj16_6.HelloInterval = d.Get(prefix16_6 + "hello_interval").(int)
			obj16_6.DeadInterval = d.Get(prefix16_6 + "dead_interval").(int)
			obj16_6.RetransmitInterval = d.Get(prefix16_6 + "retransmit_interval").(int)
			obj16_6.TransmitDelay = d.Get(prefix16_6 + "transmit_delay").(int)
			obj16_6.VirtualLinkAuthentication = d.Get(prefix16_6 + "virtual_link_authentication").(int)
			obj16_6.VirtualLinkAuthType = d.Get(prefix16_6 + "virtual_link_auth_type").(string)
			obj16_6.AuthenticationKey = d.Get(prefix16_6 + "authentication_key").(string)
			obj16_6.MessageDigestKey = d.Get(prefix16_6 + "message_digest_key").(int)
			obj16_6.Md5 = d.Get(prefix16_6 + "md5").(string)
			obj16.VirtualLinkIPAddr = append(obj16.VirtualLinkIPAddr, obj16_6)
		}

		c.AreaIpv4 = append(c.AreaIpv4, obj16)
	}

	var obj17 go_thunder.RouterOspfRedistributeOspf
	prefix17 := "redistribute.0."

	RedistListCount := d.Get(prefix17 + "redist_list.#").(int)
	obj17.Type = make([]go_thunder.RouterOspfRedistList, 0, RedistListCount)

	for i := 0; i < RedistListCount; i++ {
		var obj17_1 go_thunder.RouterOspfRedistList
		prefix17_1 := prefix17 + fmt.Sprintf("redist_list.%d.", i)
		obj17_1.Type = d.Get(prefix17_1 + "type").(string)
		obj17_1.Metric = d.Get(prefix17_1 + "metric").(int)
		obj17_1.MetricType = d.Get(prefix17_1 + "metric_type").(string)
		obj17_1.RouteMap = d.Get(prefix17_1 + "route_map").(string)
		obj17_1.Tag = d.Get(prefix17_1 + "tag").(int)
		obj17.Type = append(obj17.Type, obj17_1)
	}

	OspfListCount := d.Get(prefix17 + "ospf_list.#").(int)
	obj17.Ospf = make([]go_thunder.RouterOspfOspfList, 0, OspfListCount)

	for i := 0; i < OspfListCount; i++ {
		var obj17_2 go_thunder.RouterOspfOspfList
		prefix17_2 := prefix17 + fmt.Sprintf("ospf_list.%d.", i)
		obj17_2.Ospf = d.Get(prefix17_2 + "ospf").(int)
		obj17_2.ProcessID = d.Get(prefix17_2 + "process_id").(int)
		obj17_2.MetricOspf = d.Get(prefix17_2 + "metric_ospf").(int)
		obj17_2.MetricTypeOspf = d.Get(prefix17_2 + "metric_type_ospf").(string)
		obj17_2.RouteMapOspf = d.Get(prefix17_2 + "route_map_ospf").(string)
		obj17_2.TagOspf = d.Get(prefix17_2 + "tag_ospf").(int)
		obj17.Ospf = append(obj17.Ospf, obj17_2)
	}

	obj17.IPNat = d.Get(prefix17 + "ip_nat").(int)
	obj17.MetricIPNat = d.Get(prefix17 + "metric_ip_nat").(int)
	obj17.MetricTypeIPNat = d.Get(prefix17 + "metric_type_ip_nat").(string)
	obj17.RouteMapIPNat = d.Get(prefix17 + "route_map_ip_nat").(string)
	obj17.TagIPNat = d.Get(prefix17 + "tag_ip_nat").(int)

	IpNatFloatingListCount := d.Get(prefix17 + "ip_nat_floating_list.#").(int)
	obj17.IPNatPrefix = make([]go_thunder.RouterOspfIPNatFloatingList, 0, IpNatFloatingListCount)

	for i := 0; i < IpNatFloatingListCount; i++ {
		var obj17_3 go_thunder.RouterOspfIPNatFloatingList
		prefix17_3 := prefix17 + fmt.Sprintf("ip_nat_floating_list.%d.", i)
		obj17_3.IPNatPrefix = d.Get(prefix17_3 + "ip_nat_prefix").(string)
		obj17_3.IPNatFloatingIPForward = d.Get(prefix17_3 + "ip_nat_floating_ip_forward").(string)
		obj17.IPNatPrefix = append(obj17.IPNatPrefix, obj17_3)
	}

	VipListCount := d.Get(prefix17 + "vip_list.#").(int)
	obj17.TypeVip = make([]go_thunder.RouterOspfVipList, 0, VipListCount)

	for i := 0; i < VipListCount; i++ {
		var obj17_4 go_thunder.RouterOspfVipList
		prefix17_4 := prefix17 + fmt.Sprintf("vip_list.%d.", i)
		obj17_4.TypeVip = d.Get(prefix17_4 + "type_vip").(string)
		obj17_4.MetricVip = d.Get(prefix17_4 + "metric_vip").(int)
		obj17_4.MetricTypeVip = d.Get(prefix17_4 + "metric_type_vip").(string)
		obj17_4.RouteMapVip = d.Get(prefix17_4 + "route_map_vip").(string)
		obj17_4.TagVip = d.Get(prefix17_4 + "tag_vip").(int)
		obj17.TypeVip = append(obj17.TypeVip, obj17_4)
	}

	VipFloatingListCount := d.Get(prefix17 + "vip_floating_list.#").(int)
	obj17.VipAddress = make([]go_thunder.RouterOspfVipFloatingList, 0, VipFloatingListCount)

	for i := 0; i < VipFloatingListCount; i++ {
		var obj17_5 go_thunder.RouterOspfVipFloatingList
		prefix17_5 := prefix17 + fmt.Sprintf("vip_floating_list.%d.", i)
		obj17_5.VipAddress = d.Get(prefix17_5 + "vip_address").(string)
		obj17_5.VipFloatingIPForward = d.Get(prefix17_5 + "vip_floating_ip_forward").(string)
		obj17.VipAddress = append(obj17.VipAddress, obj17_5)
	}

	c.RedistList = obj17

	vc.ProcessID = c
	return vc
}
