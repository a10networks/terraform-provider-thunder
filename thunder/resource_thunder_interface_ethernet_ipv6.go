package thunder

//Thunder resource InterfaceEthernetIPv6

import (
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceInterfaceEthernetIPv6() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceEthernetIPv6Create,
		Update: resourceInterfaceEthernetIPv6Update,
		Read:   resourceInterfaceEthernetIPv6Read,
		Delete: resourceInterfaceEthernetIPv6Delete,
		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stateful_firewall": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_list": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"outside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"inside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"class_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"acl_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"access_list_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inbound": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"v6_acl_name": {
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
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"tag": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"area_id_addr": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"area_id_num": {
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
					},
				},
			},
			"outside": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
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
			"router_adver": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"managed_config_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"hop_limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"other_config_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"rate_limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"prefix_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"valid_lifetime": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"not_autonomous": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"prefix": {
										Type:        schema.TypeString,
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
								},
							},
						},
						"use_floating_ip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"retransmit_timer": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"min_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"adver_vrid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"floating_ip_default_vrid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"adver_vrid_default": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"max_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"use_floating_ip_default_vrid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"adver_mtu_disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"reachable_time": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"floating_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"adver_mtu": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"default_lifetime": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"address_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ipv6_addr": {
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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv6_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ospf": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bfd": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"neighbor_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"neighbor_priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"neig_inst": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"neighbor_poll_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"neighbor_cost": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"neighbor": {
										Type:        schema.TypeString,
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
						"dead_interval_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dead_interval": {
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
									"instance_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"retransmit_interval": {
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
									"instance_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"hello_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
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
						"priority_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_id": {
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
						"network_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_instance_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
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
								},
							},
						},
					},
				},
			},
			"ttl_ignore": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceInterfaceEthernetIPv6Create(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceEthernetIPv6 (Inside resourceInterfaceEthernetIPv6Create) ")
		name := d.Get("ifnum").(int)
		data := dataToInterfaceEthernetIPv6(d)
		logger.Println("[INFO] received V from method data to InterfaceEthernetIPv6 --")
		d.SetId(strconv.Itoa(name))
		go_thunder.PostInterfaceEthernetIPv6(client.Token, name, data, client.Host)

		return resourceInterfaceEthernetIPv6Read(d, meta)

	}
	return nil
}

func resourceInterfaceEthernetIPv6Read(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading InterfaceEthernetIPv6 (Inside resourceInterfaceEthernetIPv6Read)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetInterfaceEthernetIPv6(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceInterfaceEthernetIPv6Update(d *schema.ResourceData, meta interface{}) error {

	return resourceInterfaceEthernetIPv6Read(d, meta)
}

func resourceInterfaceEthernetIPv6Delete(d *schema.ResourceData, meta interface{}) error {

	return resourceInterfaceEthernetIPv6Read(d, meta)
}
func dataToInterfaceEthernetIPv6(d *schema.ResourceData) go_thunder.EthernetIPv6 {
	var vc go_thunder.EthernetIPv6
	var c go_thunder.EthernetIPv6Instance
	AddressListCount := d.Get("address_list.#").(int)
	c.AddressType = make([]go_thunder.EthernetIPv6AddressList, 0, AddressListCount)

	for i := 0; i < AddressListCount; i++ {
		var obj1 go_thunder.EthernetIPv6AddressList
		prefix := fmt.Sprintf("address_list.%d.", i)
		obj1.AddressType = d.Get(prefix + "address_type").(string)
		obj1.Ipv6Addr = d.Get(prefix + "ipv6_addr").(string)
		c.AddressType = append(c.AddressType, obj1)
	}

	c.Inside = d.Get("inside").(int)
	c.Ipv6Enable = d.Get("ipv6_enable").(int)

	var obj2 go_thunder.EthernetIPv6Rip
	prefix := "rip.0."

	var obj2_1 go_thunder.EthernetIPv6SplitHorizonCfg
	prefix1 := prefix + "split_horizon_cfg.0."
	obj2_1.State = d.Get(prefix1 + "state").(string)
	obj2.State = obj2_1

	c.State = obj2

	c.Outside = d.Get("outside").(int)

	var obj3 go_thunder.EthernetIPv6StatefulFirewall
	prefix = "stateful_firewall.0."
	obj3.ClassList = d.Get(prefix + "class_list").(string)
	obj3.ACLName = d.Get(prefix + "acl_name").(string)
	obj3.Inside = d.Get(prefix + "inside").(int)
	obj3.Outside = d.Get(prefix + "outside").(int)
	obj3.AccessList = d.Get(prefix + "access_list").(int)
	c.ClassList = obj3

	c.TTLIgnore = d.Get("ttl_ignore").(int)

	var obj4 go_thunder.EthernetIPv6Router
	prefix = "router.0."

	var obj4_1 go_thunder.EthernetIPv6Ripng
	prefix1 = prefix + "ripng.0."
	obj4_1.Rip = d.Get(prefix1 + "rip").(int)
	obj4.Rip = obj4_1

	var obj4_2 go_thunder.EthernetIPv6Ospf1
	prefix1 = prefix + "ospf.0."

	AreaListCount := d.Get(prefix1 + "area_list.#").(int)
	obj4_2.AreaIDAddr = make([]go_thunder.EthernetIPv6AreaList, 0, AreaListCount)

	for i := 0; i < AreaListCount; i++ {
		var obj4_2_1 go_thunder.EthernetIPv6AreaList
		prefix2 := prefix1 + fmt.Sprintf("area_list.%d.", i)
		obj4_2_1.AreaIDAddr = d.Get(prefix2 + "area_id_addr").(string)
		obj4_2_1.Tag = d.Get(prefix2 + "tag").(string)
		obj4_2_1.InstanceID = d.Get(prefix2 + "instance_id").(int)
		obj4_2_1.AreaIDNum = d.Get(prefix2 + "area_id_num").(int)
		obj4_2.AreaIDAddr = append(obj4_2.AreaIDAddr, obj4_2_1)
	}

	obj4.AreaIDAddr = obj4_2

	var obj4_3 go_thunder.EthernetIPv6Isis
	prefix1 = "isis.0."
	obj4_3.Tag = d.Get(prefix1 + "tag").(string)
	obj4.Tag = obj4_3

	c.Rip = obj4

	var obj5 go_thunder.EthernetIPv6AccessListCfg
	prefix = "access_list_cfg.0."
	obj5.Inbound = d.Get(prefix + "inbound").(int)
	obj5.V6ACLName = d.Get(prefix + "v6_acl_name").(string)
	c.Inbound = obj5

	var obj6 go_thunder.EthernetIPv6Ospf
	prefix = "ospf.0."
	obj6.Bfd = d.Get(prefix + "bfd").(int)

	CostCfgCount := d.Get(prefix + "cost_cfg.#").(int)
	obj6.Cost = make([]go_thunder.EthernetIPv6CostCfg, 0, CostCfgCount)

	for i := 0; i < CostCfgCount; i++ {
		var obj6_1 go_thunder.EthernetIPv6CostCfg
		prefix1 := prefix + fmt.Sprintf("cost_cfg.%d.", i)
		obj6_1.Cost = d.Get(prefix1 + "cost").(int)
		obj6_1.InstanceID = d.Get(prefix1 + "instance_id").(int)
		obj6.Cost = append(obj6.Cost, obj6_1)
	}

	PriorityCfgCount := d.Get("priority_cfg.#").(int)
	obj6.Priority = make([]go_thunder.EthernetIPv6PriorityCfg, 0, PriorityCfgCount)

	for i := 0; i < PriorityCfgCount; i++ {
		var obj6_2 go_thunder.EthernetIPv6PriorityCfg
		prefix = fmt.Sprintf("priority_cfg.%d.", i)
		obj6_2.Priority = d.Get(prefix + "priority").(int)
		obj6_2.InstanceID = d.Get(prefix + "instance_id").(int)
		obj6.Priority = append(obj6.Priority, obj6_2)
	}

	HelloIntervalCfgCount := d.Get(prefix + "hello_interval_cfg.#").(int)
	obj6.HelloInterval = make([]go_thunder.EthernetIPv6HelloIntervalCfg, 0, HelloIntervalCfgCount)

	for i := 0; i < HelloIntervalCfgCount; i++ {
		var obj6_3 go_thunder.EthernetIPv6HelloIntervalCfg
		prefix1 := prefix + fmt.Sprintf("hello_interval_cfg.%d.", i)
		obj6_3.HelloInterval = d.Get(prefix1 + "hello_interval").(int)
		obj6_3.InstanceID = d.Get(prefix1 + "instance_id").(int)
		obj6.HelloInterval = append(obj6.HelloInterval, obj6_3)
	}

	MtuIgnoreCfgCount := d.Get(prefix + "mtu_ignore_cfg.#").(int)
	obj6.MtuIgnore = make([]go_thunder.EthernetIPv6MtuIgnoreCfg, 0, MtuIgnoreCfgCount)

	for i := 0; i < MtuIgnoreCfgCount; i++ {
		var obj6_4 go_thunder.EthernetIPv6MtuIgnoreCfg
		prefix1 := prefix + fmt.Sprintf("mtu_ignore_cfg.%d.", i)
		obj6_4.MtuIgnore = d.Get(prefix1 + "mtu_ignore").(int)
		obj6_4.InstanceID = d.Get(prefix1 + "instance_id").(int)
		obj6.MtuIgnore = append(obj6.MtuIgnore, obj6_4)
	}

	RetransmitIntervalCfgCount := d.Get(prefix + "retransmit_interval_cfg.#").(int)
	obj6.RetransmitInterval = make([]go_thunder.EthernetIPv6RetransmitIntervalCfg, 0, RetransmitIntervalCfgCount)

	for i := 0; i < RetransmitIntervalCfgCount; i++ {
		var obj6_5 go_thunder.EthernetIPv6RetransmitIntervalCfg
		prefix1 = prefix + fmt.Sprintf("retransmit_interval_cfg.%d.", i)
		obj6_5.RetransmitInterval = d.Get(prefix1 + "retransmit_interval").(int)
		obj6_5.InstanceID = d.Get(prefix1 + "instance_id").(int)
		obj6.RetransmitInterval = append(obj6.RetransmitInterval, obj6_5)
	}

	obj6.Disable = d.Get(prefix + "disable").(int)

	TransmitDelayCfgCount := d.Get(prefix + "transmit_delay_cfg.#").(int)
	obj6.TransmitDelay = make([]go_thunder.EthernetIPv6TransmitDelayCfg, 0, TransmitDelayCfgCount)

	for i := 0; i < TransmitDelayCfgCount; i++ {
		var obj6_6 go_thunder.EthernetIPv6TransmitDelayCfg
		prefix1 = prefix + fmt.Sprintf("transmit_delay_cfg.%d.", i)
		obj6_6.TransmitDelay = d.Get(prefix1 + "transmit_delay").(int)
		obj6_6.InstanceID = d.Get(prefix1 + "instance_id").(int)
		obj6.TransmitDelay = append(obj6.TransmitDelay, obj6_6)
	}

	NeighborCfgCount := d.Get(prefix + "neighbor_cfg.#").(int)
	obj6.NeighborPriority = make([]go_thunder.EthernetIPv6NeighborCfg, 0, NeighborCfgCount)

	for i := 0; i < NeighborCfgCount; i++ {
		var obj6_7 go_thunder.EthernetIPv6NeighborCfg
		prefix1 = prefix + fmt.Sprintf("neighbor_cfg.%d.", i)
		obj6_7.NeighborPriority = d.Get(prefix1 + "neighbor_priority").(int)
		obj6_7.NeigInst = d.Get(prefix1 + "neig_inst").(int)
		obj6_7.NeighborPollInterval = d.Get(prefix1 + "neighbor_poll_interval").(int)
		obj6_7.NeighborCost = d.Get(prefix1 + "neighbor_cost").(int)
		obj6_7.Neighbor = d.Get(prefix1 + "neighbor").(string)
		obj6.NeighborPriority = append(obj6.NeighborPriority, obj6_7)
	}

	NetworkListCount := d.Get(prefix + "network_list.#").(int)
	obj6.BroadcastType = make([]go_thunder.EthernetIPv6NetworkList, 0, NetworkListCount)

	for i := 0; i < NetworkListCount; i++ {
		var obj6_8 go_thunder.EthernetIPv6NetworkList
		prefix1 = prefix + fmt.Sprintf("network_list.%d.", i)
		obj6_8.BroadcastType = d.Get(prefix1 + "broadcast_type").(string)
		obj6_8.P2MpNbma = d.Get(prefix1 + "p2mp_nbma").(int)
		obj6_8.NetworkInstanceID = d.Get(prefix1 + "network_instance_id").(int)
		obj6.BroadcastType = append(obj6.BroadcastType, obj6_8)
	}

	DeadIntervalCfgCount := d.Get(prefix + "dead_interval_cfg.#").(int)
	obj6.DeadInterval = make([]go_thunder.EthernetIPv6DeadIntervalCfg, 0, DeadIntervalCfgCount)

	for i := 0; i < DeadIntervalCfgCount; i++ {
		var obj6_9 go_thunder.EthernetIPv6DeadIntervalCfg
		prefix1 = prefix + fmt.Sprintf("dead_interval_cfg.%d.", i)
		obj6_9.DeadInterval = d.Get(prefix1 + "dead_interval").(int)
		obj6_9.InstanceID = d.Get(prefix1 + "instance_id").(int)
		obj6.DeadInterval = append(obj6.DeadInterval, obj6_9)
	}

	c.Bfd = obj6

	var obj7 go_thunder.EthernetIPv6RouterAdver
	prefix = "router_adver.0."
	obj7.MaxInterval = d.Get(prefix + "max_interval").(int)
	obj7.DefaultLifetime = d.Get(prefix + "default_lifetime").(int)
	obj7.ReachableTime = d.Get(prefix + "reachable_time").(int)
	obj7.OtherConfigAction = d.Get(prefix + "other_config_action").(string)
	obj7.FloatingIPDefaultVrid = d.Get(prefix + "floating_ip_default_vrid").(string)
	obj7.ManagedConfigAction = d.Get(prefix + "managed_config_action").(string)
	obj7.MinInterval = d.Get(prefix + "min_interval").(int)
	obj7.RateLimit = d.Get(prefix + "rate_limit").(int)
	obj7.AdverMtuDisable = d.Get(prefix + "adver_mtu_disable").(int)

	PrefixListCount := d.Get("prefix_list.#").(int)
	obj7.NotAutonomous = make([]go_thunder.EthernetIPv6PrefixList, 0, PrefixListCount)

	for i := 0; i < PrefixListCount; i++ {
		var obj7_1 go_thunder.EthernetIPv6PrefixList
		prefix1 = prefix + fmt.Sprintf("prefix_list.%d.", i)
		obj7_1.NotAutonomous = d.Get(prefix1 + "not_autonomous").(int)
		obj7_1.ValidLifetime = d.Get(prefix1 + "valid_lifetime").(int)
		obj7_1.NotOnLink = d.Get(prefix1 + "not_on_link").(int)
		obj7_1.Prefix = d.Get(prefix1 + "prefix").(string)
		obj7_1.PreferredLifetime = d.Get(prefix1 + "preferred_lifetime").(int)
		obj7.NotAutonomous = append(obj7.NotAutonomous, obj7_1)
	}

	obj7.FloatingIP = d.Get(prefix + "floating_ip").(string)
	obj7.AdverVrid = d.Get(prefix + "adver_vrid").(int)
	obj7.UseFloatingIPDefaultVrid = d.Get(prefix + "use_floating_ip_default_vrid").(int)
	obj7.Action = d.Get(prefix + "action").(string)
	obj7.AdverVridDefault = d.Get(prefix + "adver_vrid_default").(int)
	obj7.AdverMtu = d.Get(prefix + "adver_mtu").(int)
	obj7.RetransmitTimer = d.Get(prefix + "retransmit_timer").(int)
	obj7.HopLimit = d.Get(prefix + "hop_limit").(int)
	obj7.UseFloatingIP = d.Get(prefix + "use_floating_ip").(int)
	c.MaxInterval = obj7

	vc.UUID = c
	return vc
}
