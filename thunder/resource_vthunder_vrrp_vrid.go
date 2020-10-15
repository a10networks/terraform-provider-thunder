package thunder

//Thunder resource vrrp vrid

import (
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceVrrpVrid() *schema.Resource {
	return &schema.Resource{
		Create: resourceVrrpVridCreate,
		Update: resourceVrrpVridUpdate,
		Read:   resourceVrrpVridRead,
		Delete: resourceVrrpVridDelete,

		Schema: map[string]*schema.Schema{
			"preempt_mode": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"floating_ip": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_address_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ethernet": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"trunk": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ipv6_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ve": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ip_address_part_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address_partition": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ipv6_address_part_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ethernet": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ipv6_address_partition": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"trunk": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ve": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ip_address_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": {
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
			"follow": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrid_lead": {
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
			"blade_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fail_over_policy_template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"tracking_options": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"route": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_destination_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"protocol": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"ip_destination": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"distance": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"priority_cost": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"gateway": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"mask": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"ipv6_destination_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"protocol": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"distance": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"gatewayv6": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"priority_cost": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"ipv6_destination": {
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
									"bgp": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bgp_ipv6_address_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bgp_ipv6_address": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"priority_cost": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"bgp_ipv4_address_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"priority_cost": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"bgp_ipv4_address": {
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
									"trunk_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"priority_cost": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"trunk": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"per_port_pri": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"interface": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ethernet": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"priority_cost": {
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
									"gateway": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv4_gateway_list": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_address": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"priority_cost": {
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
												"ipv6_gateway_list": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"priority_cost": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"uuid": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"ipv6_address": {
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
									"vlan_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vlan": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"priority_cost": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"timeout": {
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
						"priority": {
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
			"vrid_val": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceVrrpVridCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	logger.Println("[INFO] Creating vrrp vrid (Inside resourceVrrpVridCreate)")

	if client.Host != "" {
		name := d.Get("vrid_val").(int)
		vc := dataToVrrpVrid(d)
		d.SetId(strconv.Itoa(name))
		go_thunder.PostVrrpVrid(client.Token, vc, client.Host)

		return resourceVrrpVridRead(d, meta)
	}
	return nil
}

func resourceVrrpVridRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading vrrp common (Inside resourceVrrpVridRead)")

	client := meta.(Thunder)

	if client.Host != "" {

		name := d.Id()

		logger.Println("[INFO] Fetching vrrp vrid" + name)

		vc, err := go_thunder.GetVrrpVrid(client.Token, name, client.Host)

		if vc == nil {
			logger.Println("[INFO] No vrrp vrid found")
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceVrrpVridUpdate(d *schema.ResourceData, meta interface{}) error {

	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Modifying vrrp vrid   (Inside resourceVrrpVridUpdate " + name)
		v := dataToVrrpVrid(d)
		d.SetId(name)
		go_thunder.PutVrrpVrid(client.Token, name, v, client.Host)

		return resourceVrrpVridRead(d, meta)
	}
	return nil
}

func resourceVrrpVridDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(Thunder)
	logger := util.GetLoggerInstance()

	if client.Host != "" {

		name := d.Id()
		logger.Println("[INFO] Deleting vrrp vrid (Inside resourceVrrpVridDelete) " + name)

		err := go_thunder.DeleteVrrpVrid(client.Token, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete vrrp vrid  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

//Utility method to instantiate Vrrp Vrid Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToVrrpVrid(d *schema.ResourceData) go_thunder.VridInstance {
	var vc go_thunder.VridInstance

	var c go_thunder.Vrid

	var obj1 go_thunder.VridPairFollow
	prefix := "pair_follow.0."
	obj1.PairFollow = d.Get(prefix + "pair_follow").(int)
	obj1.VridLead = d.Get(prefix + "vrid_lead").(string)
	c.PairFollow = obj1

	var obj2 go_thunder.VridBladeParameters
	prefix = "blade_parameters.0."
	obj2.Priority = d.Get(prefix + "priority").(int)
	obj2.FailOverPolicyTemplate = d.Get(prefix + "fail_over_policy_template").(string)

	var obj2_1 go_thunder.VridTrackingOptions
	prefix1 := prefix + "tracking_options.0."

	VlanCfgCount := d.Get(prefix1 + "vlan_cfg.#").(int)
	obj2_1.Vlan = make([]go_thunder.VridVlanCfg, 0, VlanCfgCount)

	for i := 0; i < VlanCfgCount; i++ {
		var obj2_1_1 go_thunder.VridVlanCfg
		prefix2 := prefix1 + fmt.Sprintf("vlan_cfg.%d.", i)
		obj2_1_1.Vlan = d.Get(prefix2 + "vlan").(int)
		obj2_1_1.Timeout = d.Get(prefix2 + "timeout").(int)
		obj2_1_1.PriorityCost = d.Get(prefix2 + "priority_cost").(int)
		obj2_1.Vlan = append(obj2_1.Vlan, obj2_1_1)
	}

	var obj2_1_2 go_thunder.VridRoute
	prefix2 := prefix1 + "route.0."

	Ipv6DestinationCfgCount := d.Get(prefix2 + "ipv6_destination_cfg.#").(int)
	obj2_1_2.Distance = make([]go_thunder.VridIpv6DestinationCfg, 0, Ipv6DestinationCfgCount)

	for i := 0; i < Ipv6DestinationCfgCount; i++ {
		var obj2_1_2_1 go_thunder.VridIpv6DestinationCfg
		prefix3 := prefix2 + fmt.Sprintf("ipv6_destination_cfg.%d.", i)
		obj2_1_2_1.Distance = d.Get(prefix3 + "distance").(int)
		obj2_1_2_1.Protocol = d.Get(prefix3 + "protocol").(string)
		obj2_1_2_1.PriorityCost = d.Get(prefix3 + "priority_cost").(int)
		obj2_1_2_1.Ipv6Destination = d.Get(prefix3 + "ipv6_destination").(string)
		obj2_1_2_1.Gatewayv6 = d.Get(prefix3 + "gatewayv6").(string)
		obj2_1_2.Distance = append(obj2_1_2.Distance, obj2_1_2_1)
	}

	IpDestinationCfgCount := d.Get(prefix2 + "ip_destination_cfg.#").(int)
	obj2_1_2.Protocol = make([]go_thunder.VridIPDestinationCfg, 0, IpDestinationCfgCount)

	for i := 0; i < IpDestinationCfgCount; i++ {
		var obj2_1_2_2 go_thunder.VridIPDestinationCfg
		prefix3 := prefix2 + fmt.Sprintf("ip_destination_cfg.%d.", i)
		obj2_1_2_2.Distance = d.Get(prefix3 + "distance").(int)
		obj2_1_2_2.Protocol = d.Get(prefix3 + "protocol").(string)
		obj2_1_2_2.Mask = d.Get(prefix3 + "mask").(string)
		obj2_1_2_2.PriorityCost = d.Get(prefix3 + "priority_cost").(int)
		obj2_1_2_2.IPDestination = d.Get(prefix3 + "ip_destination").(string)
		obj2_1_2_2.Gateway = d.Get(prefix3 + "gateway").(string)
		obj2_1_2.Protocol = append(obj2_1_2.Protocol, obj2_1_2_2)
	}

	obj2_1.Distance = obj2_1_2

	var obj2_1_3 go_thunder.VridBgp
	prefix2 = prefix1 + "bgp.0."

	BgpIpv4AddressCfgCount := d.Get(prefix2 + "bgp_ipv4_address_cfg.#").(int)
	obj2_1_3.BgpIpv4Address = make([]go_thunder.VridBgpIpv4AddressCfg, 0, BgpIpv4AddressCfgCount)

	for i := 0; i < BgpIpv4AddressCfgCount; i++ {
		var obj2_1_3_1 go_thunder.VridBgpIpv4AddressCfg
		prefix3 := prefix2 + fmt.Sprintf("bgp_ipv4_address_cfg.%d.", i)
		obj2_1_3_1.BgpIpv4Address = d.Get(prefix3 + "bgp_ipv4_address").(string)
		obj2_1_3_1.PriorityCost = d.Get(prefix3 + "priority_cost").(int)
		obj2_1_3.BgpIpv4Address = append(obj2_1_3.BgpIpv4Address, obj2_1_3_1)
	}

	BgpIpv6AddressCfgCount := d.Get(prefix2 + "bgp_ipv6_address_cfg.#").(int)
	obj2_1_3.BgpIpv6Address = make([]go_thunder.VridBgpIpv6AddressCfg, 0, BgpIpv6AddressCfgCount)

	for i := 0; i < BgpIpv6AddressCfgCount; i++ {
		var obj2_1_3_2 go_thunder.VridBgpIpv6AddressCfg
		prefix3 := prefix2 + fmt.Sprintf("bgp_ipv6_address_cfg.%d.", i)
		obj2_1_3_2.BgpIpv6Address = d.Get(prefix3 + "bgp_ipv6_address").(string)
		obj2_1_3_2.PriorityCost = d.Get(prefix3 + "priority_cost").(int)
		obj2_1_3.BgpIpv6Address = append(obj2_1_3.BgpIpv6Address, obj2_1_3_2)
	}

	obj2_1.BgpIpv4Address = obj2_1_3

	InterfaceCount := d.Get(prefix2 + "interface.#").(int)
	obj2_1.Ethernet = make([]go_thunder.VridInterface, 0, InterfaceCount)

	for i := 0; i < InterfaceCount; i++ {
		var obj2_1_4 go_thunder.VridInterface
		prefix2 = prefix1 + fmt.Sprintf("interface.%d.", i)
		obj2_1_4.Ethernet = d.Get(prefix2 + "ethernet").(int)
		obj2_1_4.PriorityCost = d.Get(prefix2 + "priority_cost").(int)
		obj2_1.Ethernet = append(obj2_1.Ethernet, obj2_1_4)
	}

	var obj2_1_5 go_thunder.VridGateway
	prefix2 = prefix1 + "gateway.0."

	Ipv4GatewayListCount := d.Get(prefix2 + "ipv4_gateway_list.#").(int)
	obj2_1_5.IPAddress = make([]go_thunder.VridIpv4GatewayList, 0, Ipv4GatewayListCount)

	for i := 0; i < Ipv4GatewayListCount; i++ {
		var obj2_1_5_1 go_thunder.VridIpv4GatewayList
		prefix3 := prefix2 + fmt.Sprintf("ipv4_gateway_list.%d.", i)
		obj2_1_5_1.IPAddress = d.Get(prefix3 + "ip_address").(string)
		obj2_1_5_1.PriorityCost = d.Get(prefix3 + "priority_cost").(int)
		obj2_1_5.IPAddress = append(obj2_1_5.IPAddress, obj2_1_5_1)
	}

	Ipv6GatewayListCount := d.Get(prefix2 + "ipv6_gateway_list.#").(int)
	obj2_1_5.Ipv6Address = make([]go_thunder.VridIpv6GatewayList, 0, Ipv6GatewayListCount)

	for i := 0; i < Ipv6GatewayListCount; i++ {
		var obj2_1_5_2 go_thunder.VridIpv6GatewayList
		prefix3 := prefix2 + fmt.Sprintf("ipv6_gateway_list.%d.", i)
		obj2_1_5_2.Ipv6Address = d.Get(prefix3 + "ipv6_address").(string)
		obj2_1_5_2.PriorityCost = d.Get(prefix3 + "priority_cost").(int)
		obj2_1_5.Ipv6Address = append(obj2_1_5.Ipv6Address, obj2_1_5_2)
	}

	obj2_1.IPAddress = obj2_1_5

	TrunkCfgCount := d.Get(prefix1 + "trunk_cfg.#").(int)
	obj2_1.PriorityCost = make([]go_thunder.VridTrunkCfg, 0, TrunkCfgCount)

	for i := 0; i < TrunkCfgCount; i++ {
		var obj2_1_6 go_thunder.VridTrunkCfg
		prefix2 := prefix1 + fmt.Sprintf("trunk_cfg.%d.", i)
		obj2_1_6.PriorityCost = d.Get(prefix2 + "priority_cost").(int)
		obj2_1_6.Trunk = d.Get(prefix2 + "trunk").(int)
		obj2_1_6.PerPortPri = d.Get(prefix2 + "per_port_pri").(int)
		obj2_1.PriorityCost = append(obj2_1.PriorityCost, obj2_1_6)
	}

	obj2.Vlan = obj2_1

	c.Priority = obj2

	c.VridVal = d.Get("vrid_val").(int)
	c.UserTag = d.Get("user_tag").(string)

	var obj3 go_thunder.VridPreemptMode
	prefix = "preempt_mode.0."
	obj3.Threshold = d.Get(prefix + "threshold").(int)
	obj3.Disable = d.Get(prefix + "disable").(int)
	c.Threshold = obj3

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.VridSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj4 go_thunder.VridSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj4.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj4)
	}

	var obj5 go_thunder.VridFloatingIP
	prefix = "floating_ip.0."

	Ipv6AddressPartCfgCount := d.Get("ipv6_address_part_cfg.#").(int)
	obj5.Ethernet = make([]go_thunder.VridIpv6AddressPartCfg, 0, Ipv6AddressPartCfgCount)

	for i := 0; i < Ipv6AddressPartCfgCount; i++ {
		var obj5_1 go_thunder.VridIpv6AddressPartCfg
		prefix1 = prefix + fmt.Sprintf("ipv6_address_part_cfg.%d.", i)
		obj5_1.Ethernet = d.Get(prefix1 + "ethernet").(int)
		obj5_1.Ipv6AddressPartition = d.Get(prefix1 + "ipv6_address_partition").(string)
		obj5_1.Ve = d.Get(prefix1 + "ve").(int)
		obj5_1.Trunk = d.Get(prefix1 + "trunk").(int)
		obj5.Ethernet = append(obj5.Ethernet, obj5_1)
	}

	IpAddressCfgCount := d.Get(prefix + "ip_address_cfg.#").(int)
	obj5.IPAddress = make([]go_thunder.VridIPAddressCfg, 0, IpAddressCfgCount)

	for i := 0; i < IpAddressCfgCount; i++ {
		var obj5_2 go_thunder.VridIPAddressCfg
		prefix1 := prefix + fmt.Sprintf("ip_address_cfg.%d.", i)
		obj5_2.IPAddress = d.Get(prefix1 + "ip_address").(string)
		obj5.IPAddress = append(obj5.IPAddress, obj5_2)
	}

	IpAddressPartCfgCount := d.Get(prefix + "ip_address_part_cfg.#").(int)
	obj5.IPAddressPartition = make([]go_thunder.VridIPAddressPartCfg, 0, IpAddressPartCfgCount)

	for i := 0; i < IpAddressPartCfgCount; i++ {
		var obj5_3 go_thunder.VridIPAddressPartCfg
		prefix1 := prefix + fmt.Sprintf("ip_address_part_cfg.%d.", i)
		obj5_3.IPAddressPartition = d.Get(prefix1 + "ip_address_partition").(string)
		obj5.IPAddressPartition = append(obj5.IPAddressPartition, obj5_3)
	}

	Ipv6AddressCfgCount := d.Get(prefix + "ipv6_address_cfg.#").(int)
	obj5.Ipv6Address = make([]go_thunder.VridIpv6AddressCfg, 0, Ipv6AddressCfgCount)

	for i := 0; i < Ipv6AddressCfgCount; i++ {
		var obj5_4 go_thunder.VridIpv6AddressCfg
		prefix1 := prefix + fmt.Sprintf("ipv6_address_cfg.%d.", i)
		obj5_4.Ipv6Address = d.Get(prefix1 + "ipv6_address").(string)
		obj5_4.Ethernet = d.Get(prefix1 + "ethernet").(int)
		obj5_4.Ve = d.Get(prefix1 + "ve").(int)
		obj5_4.Trunk = d.Get(prefix1 + "trunk").(int)
		obj5.Ipv6Address = append(obj5.Ipv6Address, obj5_4)
	}

	c.Ethernet = obj5

	var obj6 go_thunder.VridFollow
	prefix = "follow.0."
	obj6.VridLead = d.Get(prefix + "vrid_lead").(string)
	c.VridLead = obj6
	vc.UUID = c

	return vc
}
