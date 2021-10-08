package thunder

//Thunder resource InterfaceEthernetIpv6

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceInterfaceEthernetIpv6() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfaceEthernetIpv6Create,
		UpdateContext: resourceInterfaceEthernetIpv6Update,
		ReadContext:   resourceInterfaceEthernetIpv6Read,
		DeleteContext: resourceInterfaceEthernetIpv6Delete,
		Schema: map[string]*schema.Schema{
			"address_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_addr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"address_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"inside": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"outside": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ipv6_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ttl_ignore": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"access_list_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"v6_acl_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"inbound": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"router_adver": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"hop_limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"max_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"min_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"default_lifetime": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"rate_limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"reachable_time": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"retransmit_timer": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"adver_mtu_disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"adver_mtu": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"prefix_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"not_autonomous": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"not_on_link": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"preferred_lifetime": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"valid_lifetime": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"managed_config_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"other_config_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"adver_vrid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"use_floating_ip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"floating_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"adver_vrid_default": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"use_floating_ip_default_vrid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"floating_ip_default_vrid": {
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
			"stateful_firewall": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"class_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"outside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"access_list": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"acl_name": {
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
			"router": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ripng": {
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
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ospf": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"area_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"area_id_num": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"area_id_addr": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"tag": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"instance_id": {
													Type:        schema.TypeInt,
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
						"isis": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag": {
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
			"rip": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"split_horizon_cfg": {
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
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ospf": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"broadcast_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"p2mp_nbma": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"network_instance_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"bfd": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"cost_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cost": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"instance_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"dead_interval_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dead_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"instance_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"hello_interval_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hello_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"instance_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"mtu_ignore_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mtu_ignore": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"instance_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"neighbor_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"neighbor": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"neig_inst": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"neighbor_cost": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"neighbor_poll_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"neighbor_priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"priority_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"instance_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"retransmit_interval_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"retransmit_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"instance_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"transmit_delay_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"transmit_delay": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"instance_id": {
										Type:        schema.TypeInt,
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
			"ifnum": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceInterfaceEthernetIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceEthernetIpv6 (Inside resourceInterfaceEthernetIpv6Create) ")
		name1 := d.Get("ifnum").(string)
		data := dataToInterfaceEthernetIpv6(d)
		logger.Println("[INFO] received formatted data from method data to InterfaceEthernetIpv6 --")
		d.SetId(name1)
		err := go_thunder.PostInterfaceEthernetIpv6(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceInterfaceEthernetIpv6Read(ctx, d, meta)

	}
	return diags
}

func resourceInterfaceEthernetIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading InterfaceEthernetIpv6 (Inside resourceInterfaceEthernetIpv6Read)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetInterfaceEthernetIpv6(client.Token, name1, client.Host)
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

func resourceInterfaceEthernetIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceInterfaceEthernetIpv6Read(ctx, d, meta)
}

func resourceInterfaceEthernetIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceInterfaceEthernetIpv6Read(ctx, d, meta)
}
func dataToInterfaceEthernetIpv6(d *schema.ResourceData) go_thunder.InterfaceEthernetIpv6 {
	var vc go_thunder.InterfaceEthernetIpv6
	var c go_thunder.InterfaceEthernetIpv6Instance

	InterfaceEthernetIpv6InstanceAddressListCount := d.Get("address_list.#").(int)
	c.InterfaceEthernetIpv6InstanceAddressListIpv6Addr = make([]go_thunder.InterfaceEthernetIpv6InstanceAddressList, 0, InterfaceEthernetIpv6InstanceAddressListCount)

	for i := 0; i < InterfaceEthernetIpv6InstanceAddressListCount; i++ {
		var obj1 go_thunder.InterfaceEthernetIpv6InstanceAddressList
		prefix1 := fmt.Sprintf("address_list.%d.", i)
		obj1.InterfaceEthernetIpv6InstanceAddressListIpv6Addr = d.Get(prefix1 + "ipv6_addr").(string)
		obj1.InterfaceEthernetIpv6InstanceAddressListAddressType = d.Get(prefix1 + "address_type").(string)
		c.InterfaceEthernetIpv6InstanceAddressListIpv6Addr = append(c.InterfaceEthernetIpv6InstanceAddressListIpv6Addr, obj1)
	}

	c.InterfaceEthernetIpv6InstanceInside = d.Get("inside").(int)
	c.InterfaceEthernetIpv6InstanceOutside = d.Get("outside").(int)
	c.InterfaceEthernetIpv6InstanceIpv6Enable = d.Get("ipv6_enable").(int)
	c.InterfaceEthernetIpv6InstanceTTLIgnore = d.Get("ttl_ignore").(int)

	var obj2 go_thunder.InterfaceEthernetIpv6InstanceAccessListCfg
	prefix2 := "access_list_cfg.0."
	obj2.InterfaceEthernetIpv6InstanceAccessListCfgV6AclName = d.Get(prefix2 + "v6_acl_name").(string)
	obj2.InterfaceEthernetIpv6InstanceAccessListCfgInbound = d.Get(prefix2 + "inbound").(int)

	c.InterfaceEthernetIpv6InstanceAccessListCfgV6AclName = obj2

	var obj3 go_thunder.InterfaceEthernetIpv6InstanceRouterAdver
	prefix3 := "router_adver.0."
	obj3.InterfaceEthernetIpv6InstanceRouterAdverAction = d.Get(prefix3 + "action").(string)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverHopLimit = d.Get(prefix3 + "hop_limit").(int)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverMaxInterval = d.Get(prefix3 + "max_interval").(int)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverMinInterval = d.Get(prefix3 + "min_interval").(int)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverDefaultLifetime = d.Get(prefix3 + "default_lifetime").(int)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverRateLimit = d.Get(prefix3 + "rate_limit").(int)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverReachableTime = d.Get(prefix3 + "reachable_time").(int)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverRetransmitTimer = d.Get(prefix3 + "retransmit_timer").(int)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverAdverMtuDisable = d.Get(prefix3 + "adver_mtu_disable").(int)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverAdverMtu = d.Get(prefix3 + "adver_mtu").(int)

	InterfaceEthernetIpv6InstanceRouterAdverPrefixListCount := d.Get(prefix3 + "prefix_list.#").(int)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverPrefixListPrefix = make([]go_thunder.InterfaceEthernetIpv6InstanceRouterAdverPrefixList, 0, InterfaceEthernetIpv6InstanceRouterAdverPrefixListCount)

	for i := 0; i < InterfaceEthernetIpv6InstanceRouterAdverPrefixListCount; i++ {
		var obj3_1 go_thunder.InterfaceEthernetIpv6InstanceRouterAdverPrefixList
		prefix3_1 := prefix3 + fmt.Sprintf("prefix_list.%d.", i)
		obj3_1.InterfaceEthernetIpv6InstanceRouterAdverPrefixListPrefix = d.Get(prefix3_1 + "prefix").(string)
		obj3_1.InterfaceEthernetIpv6InstanceRouterAdverPrefixListNotAutonomous = d.Get(prefix3_1 + "not_autonomous").(int)
		obj3_1.InterfaceEthernetIpv6InstanceRouterAdverPrefixListNotOnLink = d.Get(prefix3_1 + "not_on_link").(int)
		obj3_1.InterfaceEthernetIpv6InstanceRouterAdverPrefixListPreferredLifetime = d.Get(prefix3_1 + "preferred_lifetime").(int)
		obj3_1.InterfaceEthernetIpv6InstanceRouterAdverPrefixListValidLifetime = d.Get(prefix3_1 + "valid_lifetime").(int)
		obj3.InterfaceEthernetIpv6InstanceRouterAdverPrefixListPrefix = append(obj3.InterfaceEthernetIpv6InstanceRouterAdverPrefixListPrefix, obj3_1)
	}

	obj3.InterfaceEthernetIpv6InstanceRouterAdverManagedConfigAction = d.Get(prefix3 + "managed_config_action").(string)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverOtherConfigAction = d.Get(prefix3 + "other_config_action").(string)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverAdverVrid = d.Get(prefix3 + "adver_vrid").(int)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverUseFloatingIP = d.Get(prefix3 + "use_floating_ip").(int)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverFloatingIP = d.Get(prefix3 + "floating_ip").(string)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverAdverVridDefault = d.Get(prefix3 + "adver_vrid_default").(int)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverUseFloatingIPDefaultVrid = d.Get(prefix3 + "use_floating_ip_default_vrid").(int)
	obj3.InterfaceEthernetIpv6InstanceRouterAdverFloatingIPDefaultVrid = d.Get(prefix3 + "floating_ip_default_vrid").(string)

	c.InterfaceEthernetIpv6InstanceRouterAdverAction = obj3

	var obj4 go_thunder.InterfaceEthernetIpv6InstanceStatefulFirewall
	prefix4 := "stateful_firewall.0."
	obj4.InterfaceEthernetIpv6InstanceStatefulFirewallInside = d.Get(prefix4 + "inside").(int)
	obj4.InterfaceEthernetIpv6InstanceStatefulFirewallClassList = d.Get(prefix4 + "class_list").(string)
	obj4.InterfaceEthernetIpv6InstanceStatefulFirewallOutside = d.Get(prefix4 + "outside").(int)
	obj4.InterfaceEthernetIpv6InstanceStatefulFirewallAccessList = d.Get(prefix4 + "access_list").(int)
	obj4.InterfaceEthernetIpv6InstanceStatefulFirewallAclName = d.Get(prefix4 + "acl_name").(string)

	c.InterfaceEthernetIpv6InstanceStatefulFirewallInside = obj4

	var obj5 go_thunder.InterfaceEthernetIpv6InstanceRouter
	prefix5 := "router.0."

	var obj5_1 go_thunder.InterfaceEthernetIpv6InstanceRouterRipng
	prefix5_1 := prefix5 + "ripng.0."
	obj5_1.InterfaceEthernetIpv6InstanceRouterRipngRip = d.Get(prefix5_1 + "rip").(int)

	obj5.InterfaceEthernetIpv6InstanceRouterRipngRip = obj5_1

	var obj5_2 go_thunder.InterfaceEthernetIpv6InstanceRouterOspf
	prefix5_2 := prefix5 + "ospf.0."

	InterfaceEthernetIpv6InstanceRouterOspfAreaListCount := d.Get(prefix5_2 + "area_list.#").(int)
	obj5_2.InterfaceEthernetIpv6InstanceRouterOspfAreaListAreaIDNum = make([]go_thunder.InterfaceEthernetIpv6InstanceRouterOspfAreaList, 0, InterfaceEthernetIpv6InstanceRouterOspfAreaListCount)

	for i := 0; i < InterfaceEthernetIpv6InstanceRouterOspfAreaListCount; i++ {
		var obj5_2_1 go_thunder.InterfaceEthernetIpv6InstanceRouterOspfAreaList
		prefix5_2_1 := prefix5_2 + fmt.Sprintf("area_list.%d.", i)
		obj5_2_1.InterfaceEthernetIpv6InstanceRouterOspfAreaListAreaIDNum = d.Get(prefix5_2_1 + "area_id_num").(int)
		obj5_2_1.InterfaceEthernetIpv6InstanceRouterOspfAreaListAreaIDAddr = d.Get(prefix5_2_1 + "area_id_addr").(string)
		obj5_2_1.InterfaceEthernetIpv6InstanceRouterOspfAreaListTag = d.Get(prefix5_2_1 + "tag").(string)
		obj5_2_1.InterfaceEthernetIpv6InstanceRouterOspfAreaListInstanceID = d.Get(prefix5_2_1 + "instance_id").(int)
		obj5_2.InterfaceEthernetIpv6InstanceRouterOspfAreaListAreaIDNum = append(obj5_2.InterfaceEthernetIpv6InstanceRouterOspfAreaListAreaIDNum, obj5_2_1)
	}

	obj5.InterfaceEthernetIpv6InstanceRouterOspfAreaList = obj5_2

	var obj5_3 go_thunder.InterfaceEthernetIpv6InstanceRouterIsis
	prefix5_3 := prefix5 + "isis.0."
	obj5_3.InterfaceEthernetIpv6InstanceRouterIsisTag = d.Get(prefix5_3 + "tag").(string)

	obj5.InterfaceEthernetIpv6InstanceRouterIsisTag = obj5_3

	c.InterfaceEthernetIpv6InstanceRouterRipng = obj5

	var obj6 go_thunder.InterfaceEthernetIpv6InstanceRip
	prefix6 := "rip.0."

	var obj6_1 go_thunder.InterfaceEthernetIpv6InstanceRipSplitHorizonCfg
	prefix6_1 := prefix6 + "split_horizon_cfg.0."
	obj6_1.InterfaceEthernetIpv6InstanceRipSplitHorizonCfgState = d.Get(prefix6_1 + "state").(string)

	obj6.InterfaceEthernetIpv6InstanceRipSplitHorizonCfgState = obj6_1

	c.InterfaceEthernetIpv6InstanceRipSplitHorizonCfg = obj6

	var obj7 go_thunder.InterfaceEthernetIpv6InstanceOspf
	prefix7 := "ospf.0."

	InterfaceEthernetIpv6InstanceOspfNetworkListCount := d.Get(prefix7 + "network_list.#").(int)
	obj7.InterfaceEthernetIpv6InstanceOspfNetworkListBroadcastType = make([]go_thunder.InterfaceEthernetIpv6InstanceOspfNetworkList, 0, InterfaceEthernetIpv6InstanceOspfNetworkListCount)

	for i := 0; i < InterfaceEthernetIpv6InstanceOspfNetworkListCount; i++ {
		var obj7_1 go_thunder.InterfaceEthernetIpv6InstanceOspfNetworkList
		prefix7_1 := prefix7 + fmt.Sprintf("network_list.%d.", i)
		obj7_1.InterfaceEthernetIpv6InstanceOspfNetworkListBroadcastType = d.Get(prefix7_1 + "broadcast_type").(string)
		obj7_1.InterfaceEthernetIpv6InstanceOspfNetworkListP2MpNbma = d.Get(prefix7_1 + "p2mp_nbma").(int)
		obj7_1.InterfaceEthernetIpv6InstanceOspfNetworkListNetworkInstanceID = d.Get(prefix7_1 + "network_instance_id").(int)
		obj7.InterfaceEthernetIpv6InstanceOspfNetworkListBroadcastType = append(obj7.InterfaceEthernetIpv6InstanceOspfNetworkListBroadcastType, obj7_1)
	}

	obj7.InterfaceEthernetIpv6InstanceOspfBfd = d.Get(prefix7 + "bfd").(int)
	obj7.InterfaceEthernetIpv6InstanceOspfDisable = d.Get(prefix7 + "disable").(int)

	InterfaceEthernetIpv6InstanceOspfCostCfgCount := d.Get(prefix7 + "cost_cfg.#").(int)
	obj7.InterfaceEthernetIpv6InstanceOspfCostCfgCost = make([]go_thunder.InterfaceEthernetIpv6InstanceOspfCostCfg, 0, InterfaceEthernetIpv6InstanceOspfCostCfgCount)

	for i := 0; i < InterfaceEthernetIpv6InstanceOspfCostCfgCount; i++ {
		var obj7_2 go_thunder.InterfaceEthernetIpv6InstanceOspfCostCfg
		prefix7_2 := prefix7 + fmt.Sprintf("cost_cfg.%d.", i)
		obj7_2.InterfaceEthernetIpv6InstanceOspfCostCfgCost = d.Get(prefix7_2 + "cost").(int)
		obj7_2.InterfaceEthernetIpv6InstanceOspfCostCfgInstanceID = d.Get(prefix7_2 + "instance_id").(int)
		obj7.InterfaceEthernetIpv6InstanceOspfCostCfgCost = append(obj7.InterfaceEthernetIpv6InstanceOspfCostCfgCost, obj7_2)
	}

	InterfaceEthernetIpv6InstanceOspfDeadIntervalCfgCount := d.Get(prefix7 + "dead_interval_cfg.#").(int)
	obj7.InterfaceEthernetIpv6InstanceOspfDeadIntervalCfgDeadInterval = make([]go_thunder.InterfaceEthernetIpv6InstanceOspfDeadIntervalCfg, 0, InterfaceEthernetIpv6InstanceOspfDeadIntervalCfgCount)

	for i := 0; i < InterfaceEthernetIpv6InstanceOspfDeadIntervalCfgCount; i++ {
		var obj7_3 go_thunder.InterfaceEthernetIpv6InstanceOspfDeadIntervalCfg
		prefix7_3 := prefix7 + fmt.Sprintf("dead_interval_cfg.%d.", i)
		obj7_3.InterfaceEthernetIpv6InstanceOspfDeadIntervalCfgDeadInterval = d.Get(prefix7_3 + "dead_interval").(int)
		obj7_3.InterfaceEthernetIpv6InstanceOspfDeadIntervalCfgInstanceID = d.Get(prefix7_3 + "instance_id").(int)
		obj7.InterfaceEthernetIpv6InstanceOspfDeadIntervalCfgDeadInterval = append(obj7.InterfaceEthernetIpv6InstanceOspfDeadIntervalCfgDeadInterval, obj7_3)
	}

	InterfaceEthernetIpv6InstanceOspfHelloIntervalCfgCount := d.Get(prefix7 + "hello_interval_cfg.#").(int)
	obj7.InterfaceEthernetIpv6InstanceOspfHelloIntervalCfgHelloInterval = make([]go_thunder.InterfaceEthernetIpv6InstanceOspfHelloIntervalCfg, 0, InterfaceEthernetIpv6InstanceOspfHelloIntervalCfgCount)

	for i := 0; i < InterfaceEthernetIpv6InstanceOspfHelloIntervalCfgCount; i++ {
		var obj7_4 go_thunder.InterfaceEthernetIpv6InstanceOspfHelloIntervalCfg
		prefix7_4 := prefix7 + fmt.Sprintf("hello_interval_cfg.%d.", i)
		obj7_4.InterfaceEthernetIpv6InstanceOspfHelloIntervalCfgHelloInterval = d.Get(prefix7_4 + "hello_interval").(int)
		obj7_4.InterfaceEthernetIpv6InstanceOspfHelloIntervalCfgInstanceID = d.Get(prefix7_4 + "instance_id").(int)
		obj7.InterfaceEthernetIpv6InstanceOspfHelloIntervalCfgHelloInterval = append(obj7.InterfaceEthernetIpv6InstanceOspfHelloIntervalCfgHelloInterval, obj7_4)
	}

	InterfaceEthernetIpv6InstanceOspfMtuIgnoreCfgCount := d.Get(prefix7 + "mtu_ignore_cfg.#").(int)
	obj7.InterfaceEthernetIpv6InstanceOspfMtuIgnoreCfgMtuIgnore = make([]go_thunder.InterfaceEthernetIpv6InstanceOspfMtuIgnoreCfg, 0, InterfaceEthernetIpv6InstanceOspfMtuIgnoreCfgCount)

	for i := 0; i < InterfaceEthernetIpv6InstanceOspfMtuIgnoreCfgCount; i++ {
		var obj7_5 go_thunder.InterfaceEthernetIpv6InstanceOspfMtuIgnoreCfg
		prefix7_5 := prefix7 + fmt.Sprintf("mtu_ignore_cfg.%d.", i)
		obj7_5.InterfaceEthernetIpv6InstanceOspfMtuIgnoreCfgMtuIgnore = d.Get(prefix7_5 + "mtu_ignore").(int)
		obj7_5.InterfaceEthernetIpv6InstanceOspfMtuIgnoreCfgInstanceID = d.Get(prefix7_5 + "instance_id").(int)
		obj7.InterfaceEthernetIpv6InstanceOspfMtuIgnoreCfgMtuIgnore = append(obj7.InterfaceEthernetIpv6InstanceOspfMtuIgnoreCfgMtuIgnore, obj7_5)
	}

	InterfaceEthernetIpv6InstanceOspfNeighborCfgCount := d.Get(prefix7 + "neighbor_cfg.#").(int)
	obj7.InterfaceEthernetIpv6InstanceOspfNeighborCfgNeighbor = make([]go_thunder.InterfaceEthernetIpv6InstanceOspfNeighborCfg, 0, InterfaceEthernetIpv6InstanceOspfNeighborCfgCount)

	for i := 0; i < InterfaceEthernetIpv6InstanceOspfNeighborCfgCount; i++ {
		var obj7_6 go_thunder.InterfaceEthernetIpv6InstanceOspfNeighborCfg
		prefix7_6 := prefix7 + fmt.Sprintf("neighbor_cfg.%d.", i)
		obj7_6.InterfaceEthernetIpv6InstanceOspfNeighborCfgNeighbor = d.Get(prefix7_6 + "neighbor").(string)
		obj7_6.InterfaceEthernetIpv6InstanceOspfNeighborCfgNeigInst = d.Get(prefix7_6 + "neig_inst").(int)
		obj7_6.InterfaceEthernetIpv6InstanceOspfNeighborCfgNeighborCost = d.Get(prefix7_6 + "neighbor_cost").(int)
		obj7_6.InterfaceEthernetIpv6InstanceOspfNeighborCfgNeighborPollInterval = d.Get(prefix7_6 + "neighbor_poll_interval").(int)
		obj7_6.InterfaceEthernetIpv6InstanceOspfNeighborCfgNeighborPriority = d.Get(prefix7_6 + "neighbor_priority").(int)
		obj7.InterfaceEthernetIpv6InstanceOspfNeighborCfgNeighbor = append(obj7.InterfaceEthernetIpv6InstanceOspfNeighborCfgNeighbor, obj7_6)
	}

	InterfaceEthernetIpv6InstanceOspfPriorityCfgCount := d.Get(prefix7 + "priority_cfg.#").(int)
	obj7.InterfaceEthernetIpv6InstanceOspfPriorityCfgPriority = make([]go_thunder.InterfaceEthernetIpv6InstanceOspfPriorityCfg, 0, InterfaceEthernetIpv6InstanceOspfPriorityCfgCount)

	for i := 0; i < InterfaceEthernetIpv6InstanceOspfPriorityCfgCount; i++ {
		var obj7_7 go_thunder.InterfaceEthernetIpv6InstanceOspfPriorityCfg
		prefix7_7 := prefix7 + fmt.Sprintf("priority_cfg.%d.", i)
		obj7_7.InterfaceEthernetIpv6InstanceOspfPriorityCfgPriority = d.Get(prefix7_7 + "priority").(int)
		obj7_7.InterfaceEthernetIpv6InstanceOspfPriorityCfgInstanceID = d.Get(prefix7_7 + "instance_id").(int)
		obj7.InterfaceEthernetIpv6InstanceOspfPriorityCfgPriority = append(obj7.InterfaceEthernetIpv6InstanceOspfPriorityCfgPriority, obj7_7)
	}

	InterfaceEthernetIpv6InstanceOspfRetransmitIntervalCfgCount := d.Get(prefix7 + "retransmit_interval_cfg.#").(int)
	obj7.InterfaceEthernetIpv6InstanceOspfRetransmitIntervalCfgRetransmitInterval = make([]go_thunder.InterfaceEthernetIpv6InstanceOspfRetransmitIntervalCfg, 0, InterfaceEthernetIpv6InstanceOspfRetransmitIntervalCfgCount)

	for i := 0; i < InterfaceEthernetIpv6InstanceOspfRetransmitIntervalCfgCount; i++ {
		var obj7_8 go_thunder.InterfaceEthernetIpv6InstanceOspfRetransmitIntervalCfg
		prefix7_8 := prefix7 + fmt.Sprintf("retransmit_interval_cfg.%d.", i)
		obj7_8.InterfaceEthernetIpv6InstanceOspfRetransmitIntervalCfgRetransmitInterval = d.Get(prefix7_8 + "retransmit_interval").(int)
		obj7_8.InterfaceEthernetIpv6InstanceOspfRetransmitIntervalCfgInstanceID = d.Get(prefix7_8 + "instance_id").(int)
		obj7.InterfaceEthernetIpv6InstanceOspfRetransmitIntervalCfgRetransmitInterval = append(obj7.InterfaceEthernetIpv6InstanceOspfRetransmitIntervalCfgRetransmitInterval, obj7_8)
	}

	InterfaceEthernetIpv6InstanceOspfTransmitDelayCfgCount := d.Get(prefix7 + "transmit_delay_cfg.#").(int)
	obj7.InterfaceEthernetIpv6InstanceOspfTransmitDelayCfgTransmitDelay = make([]go_thunder.InterfaceEthernetIpv6InstanceOspfTransmitDelayCfg, 0, InterfaceEthernetIpv6InstanceOspfTransmitDelayCfgCount)

	for i := 0; i < InterfaceEthernetIpv6InstanceOspfTransmitDelayCfgCount; i++ {
		var obj7_9 go_thunder.InterfaceEthernetIpv6InstanceOspfTransmitDelayCfg
		prefix7_9 := prefix7 + fmt.Sprintf("transmit_delay_cfg.%d.", i)
		obj7_9.InterfaceEthernetIpv6InstanceOspfTransmitDelayCfgTransmitDelay = d.Get(prefix7_9 + "transmit_delay").(int)
		obj7_9.InterfaceEthernetIpv6InstanceOspfTransmitDelayCfgInstanceID = d.Get(prefix7_9 + "instance_id").(int)
		obj7.InterfaceEthernetIpv6InstanceOspfTransmitDelayCfgTransmitDelay = append(obj7.InterfaceEthernetIpv6InstanceOspfTransmitDelayCfgTransmitDelay, obj7_9)
	}

	c.InterfaceEthernetIpv6InstanceOspfNetworkList = obj7

	vc.InterfaceEthernetIpv6InstanceAddressList = c
	return vc
}
