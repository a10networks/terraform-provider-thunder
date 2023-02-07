package thunder

//Thunder resource InterfaceVeIPv6

import (
	"context"
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceVeIPv6() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfaceVeIPv6Create,
		UpdateContext: resourceInterfaceVeIPv6Update,
		ReadContext:   resourceInterfaceVeIPv6Read,
		DeleteContext: resourceInterfaceVeIPv6Delete,
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
			"inbound": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
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
			"v6_acl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
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
			"ttl_ignore": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceInterfaceVeIPv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceVeIPv6 (Inside resourceInterfaceVeIPv6Create) ")
		name := d.Get("ifnum").(int)
		data := dataToInterfaceVeIPv6(d)
		logger.Println("[INFO] received V from method data to InterfaceVeIPv6 --")
		d.SetId(strconv.Itoa(name))
		err := go_thunder.PostInterfaceVeIPv6(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceInterfaceVeIPv6Read(ctx, d, meta)

	}
	return diags
}

func resourceInterfaceVeIPv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading InterfaceVeIPv6 (Inside resourceInterfaceVeIPv6Read)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetInterfaceVeIPv6(client.Token, name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceInterfaceVeIPv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceInterfaceVeIPv6Read(ctx, d, meta)
}

func resourceInterfaceVeIPv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceInterfaceVeIPv6Read(ctx, d, meta)
}
func dataToInterfaceVeIPv6(d *schema.ResourceData) go_thunder.VeIPv6 {
	var vc go_thunder.VeIPv6
	var c go_thunder.VeIPv6Instance
	c.Inbound = d.Get("inbound").(int)

	AddressListCount := d.Get("address_list.#").(int)
	c.AddressType = make([]go_thunder.VeIPv6AddressList, 0, AddressListCount)

	for i := 0; i < AddressListCount; i++ {
		var obj1 go_thunder.VeIPv6AddressList
		prefix := fmt.Sprintf("address_list.%d.", i)
		obj1.AddressType = d.Get(prefix + "address_type").(string)
		obj1.Ipv6Addr = d.Get(prefix + "ipv6_addr").(string)
		c.AddressType = append(c.AddressType, obj1)
	}

	c.Inside = d.Get("inside").(int)
	c.Ipv6Enable = d.Get("ipv6_enable").(int)

	var obj2 go_thunder.VeIPv6Rip
	prefix := "rip.0."

	var obj2_1 go_thunder.VeIPv6SplitHorizonCfg
	prefix1 := prefix + "split_horizon_cfg.0."
	obj2_1.State = d.Get(prefix1 + "state").(string)
	obj2.State = obj2_1

	c.State = obj2

	c.Outside = d.Get("outside").(int)

	var obj3 go_thunder.VeIPv6StatefulFirewall
	prefix = "stateful_firewall.0."
	obj3.ClassList = d.Get(prefix + "class_list").(string)
	obj3.ACLName = d.Get(prefix + "acl_name").(string)
	obj3.Inside = d.Get(prefix + "inside").(int)
	obj3.Outside = d.Get(prefix + "outside").(int)
	obj3.AccessList = d.Get(prefix + "access_list").(int)
	c.ClassList = obj3

	c.V6ACLName = d.Get("v6_acl_name").(string)
	c.TTLIgnore = d.Get("ttl_ignore").(int)

	var obj4 go_thunder.VeIPv6Router
	prefix = "router.0."

	var obj4_1 go_thunder.VeIPv6Ripng
	prefix1 = prefix + "ripng.0."
	obj4_1.Rip = d.Get(prefix1 + "rip").(int)
	obj4.Rip = obj4_1

	var obj4_2 go_thunder.VeIPv6Ospf1
	prefix1 = prefix + "ospf.0."

	AreaListCount := d.Get(prefix1 + "area_list.#").(int)
	obj4_2.AreaIDAddr = make([]go_thunder.VeIPv6AreaList, 0, AreaListCount)

	for i := 0; i < AreaListCount; i++ {
		var obj4_2_1 go_thunder.VeIPv6AreaList
		prefix2 := prefix1 + fmt.Sprintf("area_list.%d.", i)
		obj4_2_1.AreaIDAddr = d.Get(prefix2 + "area_id_addr").(string)
		obj4_2_1.Tag = d.Get(prefix2 + "tag").(string)
		obj4_2_1.InstanceID = d.Get(prefix2 + "instance_id").(int)
		obj4_2_1.AreaIDNum = d.Get(prefix2 + "area_id_num").(int)
		obj4_2.AreaIDAddr = append(obj4_2.AreaIDAddr, obj4_2_1)
	}

	obj4.AreaIDAddr = obj4_2

	var obj4_3 go_thunder.VeIPv6Isis
	prefix1 = prefix + "isis.0."
	obj4_3.Tag = d.Get(prefix1 + "tag").(string)
	obj4.Tag = obj4_3

	c.Router = obj4

	var obj5 go_thunder.VeIPv6Ospf
	prefix = "ospf.0."
	obj5.Bfd = d.Get(prefix + "bfd").(int)

	CostCfgCount := d.Get(prefix + "cost_cfg.#").(int)
	obj5.Cost = make([]go_thunder.VeIPv6CostCfg, 0, CostCfgCount)

	for i := 0; i < CostCfgCount; i++ {
		var obj5_1 go_thunder.VeIPv6CostCfg
		prefix1 = prefix + fmt.Sprintf("cost_cfg.%d.", i)
		obj5_1.Cost = d.Get(prefix1 + "cost").(int)
		obj5_1.InstanceID = d.Get(prefix1 + "instance_id").(int)
		obj5.Cost = append(obj5.Cost, obj5_1)
	}

	PriorityCfgCount := d.Get(prefix + "priority_cfg.#").(int)
	obj5.Priority = make([]go_thunder.VeIPv6PriorityCfg, 0, PriorityCfgCount)

	for i := 0; i < PriorityCfgCount; i++ {
		var obj5_2 go_thunder.VeIPv6PriorityCfg
		prefix1 = prefix + fmt.Sprintf("priority_cfg.%d.", i)
		obj5_2.Priority = d.Get(prefix1 + "priority").(int)
		obj5_2.InstanceID = d.Get(prefix1 + "instance_id").(int)
		obj5.Priority = append(obj5.Priority, obj5_2)
	}

	HelloIntervalCfgCount := d.Get(prefix + "hello_interval_cfg.#").(int)
	obj5.HelloInterval = make([]go_thunder.VeIPv6HelloIntervalCfg, 0, HelloIntervalCfgCount)

	for i := 0; i < HelloIntervalCfgCount; i++ {
		var obj5_3 go_thunder.VeIPv6HelloIntervalCfg
		prefix1 = prefix + fmt.Sprintf("hello_interval_cfg.%d.", i)
		obj5_3.HelloInterval = d.Get(prefix1 + "hello_interval").(int)
		obj5_3.InstanceID = d.Get(prefix1 + "instance_id").(int)
		obj5.HelloInterval = append(obj5.HelloInterval, obj5_3)
	}

	MtuIgnoreCfgCount := d.Get(prefix + "mtu_ignore_cfg.#").(int)
	obj5.MtuIgnore = make([]go_thunder.VeIPv6MtuIgnoreCfg, 0, MtuIgnoreCfgCount)

	for i := 0; i < MtuIgnoreCfgCount; i++ {
		var obj5_4 go_thunder.VeIPv6MtuIgnoreCfg
		prefix1 = prefix + fmt.Sprintf("mtu_ignore_cfg.%d.", i)
		obj5_4.MtuIgnore = d.Get(prefix1 + "mtu_ignore").(int)
		obj5_4.InstanceID = d.Get(prefix1 + "instance_id").(int)
		obj5.MtuIgnore = append(obj5.MtuIgnore, obj5_4)
	}

	RetransmitIntervalCfgCount := d.Get(prefix + "retransmit_interval_cfg.#").(int)
	obj5.RetransmitInterval = make([]go_thunder.VeIPv6RetransmitIntervalCfg, 0, RetransmitIntervalCfgCount)

	for i := 0; i < RetransmitIntervalCfgCount; i++ {
		var obj5_5 go_thunder.VeIPv6RetransmitIntervalCfg
		prefix1 = prefix + fmt.Sprintf("retransmit_interval_cfg.%d.", i)
		obj5_5.RetransmitInterval = d.Get(prefix1 + "retransmit_interval").(int)
		obj5_5.InstanceID = d.Get(prefix1 + "instance_id").(int)
		obj5.RetransmitInterval = append(obj5.RetransmitInterval, obj5_5)
	}

	obj5.Disable = d.Get(prefix + "disable").(int)

	TransmitDelayCfgCount := d.Get(prefix + "transmit_delay_cfg.#").(int)
	obj5.TransmitDelay = make([]go_thunder.VeIPv6TransmitDelayCfg, 0, TransmitDelayCfgCount)

	for i := 0; i < TransmitDelayCfgCount; i++ {
		var obj5_6 go_thunder.VeIPv6TransmitDelayCfg
		prefix1 = prefix + fmt.Sprintf("transmit_delay_cfg.%d.", i)
		obj5_6.TransmitDelay = d.Get(prefix1 + "transmit_delay").(int)
		obj5_6.InstanceID = d.Get(prefix1 + "instance_id").(int)
		obj5.TransmitDelay = append(obj5.TransmitDelay, obj5_6)
	}

	NeighborCfgCount := d.Get(prefix + "neighbor_cfg.#").(int)
	obj5.NeighborPriority = make([]go_thunder.VeIPv6NeighborCfg, 0, NeighborCfgCount)

	for i := 0; i < NeighborCfgCount; i++ {
		var obj5_7 go_thunder.VeIPv6NeighborCfg
		prefix1 = prefix + fmt.Sprintf("neighbor_cfg.%d.", i)
		obj5_7.NeighborPriority = d.Get(prefix1 + "neighbor_priority").(int)
		obj5_7.NeigInst = d.Get(prefix1 + "neig_inst").(int)
		obj5_7.NeighborPollInterval = d.Get(prefix1 + "neighbor_poll_interval").(int)
		obj5_7.NeighborCost = d.Get(prefix1 + "neighbor_cost").(int)
		obj5_7.Neighbor = d.Get(prefix1 + "neighbor").(string)
		obj5.NeighborPriority = append(obj5.NeighborPriority, obj5_7)
	}

	NetworkListCount := d.Get(prefix + "network_list.#").(int)
	obj5.BroadcastType = make([]go_thunder.VeIPv6NetworkList, 0, NetworkListCount)

	for i := 0; i < NetworkListCount; i++ {
		var obj5_8 go_thunder.VeIPv6NetworkList
		prefix1 = prefix + fmt.Sprintf("network_list.%d.", i)
		obj5_8.BroadcastType = d.Get(prefix1 + "broadcast_type").(string)
		obj5_8.P2MpNbma = d.Get(prefix1 + "p2mp_nbma").(int)
		obj5_8.NetworkInstanceID = d.Get(prefix1 + "network_instance_id").(int)
		obj5.BroadcastType = append(obj5.BroadcastType, obj5_8)
	}

	DeadIntervalCfgCount := d.Get(prefix + "dead_interval_cfg.#").(int)
	obj5.DeadInterval = make([]go_thunder.VeIPv6DeadIntervalCfg, 0, DeadIntervalCfgCount)

	for i := 0; i < DeadIntervalCfgCount; i++ {
		var obj5_9 go_thunder.VeIPv6DeadIntervalCfg
		prefix1 = prefix + fmt.Sprintf("dead_interval_cfg.%d.", i)
		obj5_9.DeadInterval = d.Get(prefix1 + "dead_interval").(int)
		obj5_9.InstanceID = d.Get(prefix1 + "instance_id").(int)
		obj5.DeadInterval = append(obj5.DeadInterval, obj5_9)
	}

	c.Bfd = obj5

	var obj6 go_thunder.VeIPv6RouterAdver
	prefix = "router_adver.0."
	obj6.MaxInterval = d.Get(prefix + "max_interval").(int)
	obj6.DefaultLifetime = d.Get(prefix + "default_lifetime").(int)
	obj6.ReachableTime = d.Get(prefix + "reachable_time").(int)
	obj6.OtherConfigAction = d.Get(prefix + "other_config_action").(string)
	obj6.FloatingIPDefaultVrid = d.Get(prefix + "floating_ip_default_vrid").(string)
	obj6.ManagedConfigAction = d.Get(prefix + "managed_config_action").(string)
	obj6.MinInterval = d.Get(prefix + "min_interval").(int)
	obj6.RateLimit = d.Get(prefix + "rate_limit").(int)
	obj6.AdverMtuDisable = d.Get(prefix + "adver_mtu_disable").(int)

	PrefixListCount := d.Get(prefix + "prefix_list.#").(int)
	obj6.NotAutonomous = make([]go_thunder.VeIPv6PrefixList, 0, PrefixListCount)

	for i := 0; i < PrefixListCount; i++ {
		var obj6_1 go_thunder.VeIPv6PrefixList
		prefix1 = prefix + fmt.Sprintf("prefix_list.%d.", i)
		obj6_1.NotAutonomous = d.Get(prefix1 + "not_autonomous").(int)
		obj6_1.ValidLifetime = d.Get(prefix1 + "valid_lifetime").(int)
		obj6_1.NotOnLink = d.Get(prefix1 + "not_on_link").(int)
		obj6_1.Prefix = d.Get(prefix1 + "prefix").(string)
		obj6_1.PreferredLifetime = d.Get(prefix1 + "preferred_lifetime").(int)
		obj6.NotAutonomous = append(obj6.NotAutonomous, obj6_1)
	}

	obj6.FloatingIP = d.Get(prefix + "floating_ip").(string)
	obj6.AdverVrid = d.Get(prefix + "adver_vrid").(int)
	obj6.UseFloatingIPDefaultVrid = d.Get(prefix + "use_floating_ip_default_vrid").(int)
	obj6.Action = d.Get(prefix + "action").(string)
	obj6.AdverVridDefault = d.Get(prefix + "adver_vrid_default").(int)
	obj6.AdverMtu = d.Get(prefix + "adver_mtu").(int)
	obj6.RetransmitTimer = d.Get(prefix + "retransmit_timer").(int)
	obj6.HopLimit = d.Get(prefix + "hop_limit").(int)
	obj6.UseFloatingIP = d.Get(prefix + "use_floating_ip").(int)
	c.MaxInterval = obj6

	vc.UUID = c
	return vc
}
