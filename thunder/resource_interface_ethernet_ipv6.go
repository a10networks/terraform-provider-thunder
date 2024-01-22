package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernetIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ethernet_ipv6`: Global IPv6 configuration subcommands\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceEthernetIpv6Create,
		UpdateContext: resourceInterfaceEthernetIpv6Update,
		ReadContext:   resourceInterfaceEthernetIpv6Read,
		DeleteContext: resourceInterfaceEthernetIpv6Delete,

		Schema: map[string]*schema.Schema{
			"access_list_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"v6_acl_name": {
							Type: schema.TypeString, Optional: true, Description: "Apply ACL rules to incoming packets on this interface (Named Access List)",
						},
						"inbound": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "ACL applied on incoming packets to this interface",
						},
					},
				},
			},
			"address_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "Set the IPv6 address of an interface",
						},
						"address_type": {
							Type: schema.TypeString, Optional: true, Description: "'anycast': Configure an IPv6 anycast address; 'link-local': Configure an IPv6 link local address;",
						},
					},
				},
			},
			"inside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as inside",
			},
			"ipv6_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IPv6 processing",
			},
			"ospf": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"broadcast_type": {
										Type: schema.TypeString, Optional: true, Description: "'broadcast': Specify OSPF broadcast multi-access network; 'non-broadcast': Specify OSPF NBMA network; 'point-to-point': Specify OSPF point-to-point network; 'point-to-multipoint': Specify OSPF point-to-multipoint network;",
									},
									"p2mp_nbma": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify non-broadcast point-to-multipoint network",
									},
									"network_instance_id": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
									},
								},
							},
						},
						"bfd": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable BFD",
						},
						"cost_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cost": {
										Type: schema.TypeInt, Optional: true, Description: "Interface cost",
									},
									"instance_id": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
									},
								},
							},
						},
						"dead_interval_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dead_interval": {
										Type: schema.TypeInt, Optional: true, Default: 40, Description: "Interval after which a neighbor is declared dead (Seconds)",
									},
									"instance_id": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
									},
								},
							},
						},
						"hello_interval_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hello_interval": {
										Type: schema.TypeInt, Optional: true, Default: 10, Description: "Time between HELLO packets (Seconds)",
									},
									"instance_id": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
									},
								},
							},
						},
						"mtu_ignore_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mtu_ignore": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignores the MTU in DBD packets",
									},
									"instance_id": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
									},
								},
							},
						},
						"neighbor_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"neighbor": {
										Type: schema.TypeString, Optional: true, Default: "::", Description: "OSPFv3 neighbor (Neighbor IPv6 address)",
									},
									"neig_inst": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
									},
									"neighbor_cost": {
										Type: schema.TypeInt, Optional: true, Description: "OSPF cost for point-to-multipoint neighbor (metric)",
									},
									"neighbor_poll_interval": {
										Type: schema.TypeInt, Optional: true, Description: "OSPF dead-router polling interval (Seconds)",
									},
									"neighbor_priority": {
										Type: schema.TypeInt, Optional: true, Description: "OSPF priority of non-broadcast neighbor",
									},
								},
							},
						},
						"priority_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"priority": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Router priority",
									},
									"instance_id": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
									},
								},
							},
						},
						"retransmit_interval_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"retransmit_interval": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Time between retransmitting lost link state advertisements (Seconds)",
									},
									"instance_id": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
									},
								},
							},
						},
						"transmit_delay_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"transmit_delay": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Link state transmit delay (Seconds)",
									},
									"instance_id": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
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
			"outside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as outside",
			},
			"rip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"split_horizon_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"state": {
										Type: schema.TypeString, Optional: true, Default: "poisoned", Description: "'poisoned': Perform split horizon with poisoned reverse; 'disable': Disable split horizon; 'enable': Perform split horizon without poisoned reverse;",
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
			"router": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ripng": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rip": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "RIP Routing for IPv6",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ospf": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"area_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"area_id_num": {
													Type: schema.TypeInt, Optional: true, Description: "OSPF area ID as a decimal value",
												},
												"area_id_addr": {
													Type: schema.TypeString, Optional: true, Description: "OSPF area ID in IP address format",
												},
												"tag": {
													Type: schema.TypeString, Optional: true, Description: "Set the OSPFv3 process tag",
												},
												"instance_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set the interface instance ID",
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
						"isis": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag": {
										Type: schema.TypeString, Optional: true, Description: "ISO routing area tag",
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
			"router_adver": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Router Advertisements on this interface; 'disable': Disable Router Advertisements on this interface;",
						},
						"hop_limit": {
							Type: schema.TypeInt, Optional: true, Default: 255, Description: "Set Router Advertisement Hop Limit (default: 255)",
						},
						"max_interval": {
							Type: schema.TypeInt, Optional: true, Default: 600, Description: "Set Router Advertisement Max Interval (default: 600) (Max Router Advertisement Interval (seconds))",
						},
						"min_interval": {
							Type: schema.TypeInt, Optional: true, Default: 200, Description: "Set Router Advertisement Min Interval (default: 200) (Min Router Advertisement Interval (seconds))",
						},
						"default_lifetime": {
							Type: schema.TypeInt, Optional: true, Default: 1800, Description: "Set Router Advertisement Default Lifetime (default: 1800) (Default Lifetime (seconds))",
						},
						"rate_limit": {
							Type: schema.TypeInt, Optional: true, Default: 100000, Description: "Rate Limit the processing of incoming Router Solicitations (Max Number of Router Solicitations to process per second)",
						},
						"reachable_time": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set Router Advertisement Reachable ime (default: 0) (Reachable Time (milliseconds))",
						},
						"retransmit_timer": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set Router Advertisement Retransmit Timer (default: 0)",
						},
						"adver_mtu_disable": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Disable Router Advertisement MTU Option",
						},
						"adver_mtu": {
							Type: schema.TypeInt, Optional: true, Description: "Set Router Advertisement MTU Option",
						},
						"prefix_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type: schema.TypeString, Optional: true, Description: "Set Router Advertisement On-Link Prefix (IPv6 On-Link Prefix)",
									},
									"not_autonomous": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify that the Prefix is not usable for autoconfiguration (default:autonomous)",
									},
									"not_on_link": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify that the Prefix is not On-Link (default: on-link)",
									},
									"preferred_lifetime": {
										Type: schema.TypeInt, Optional: true, Default: 604800, Description: "Specify Prefix Preferred Lifetime (default:604800) (Prefix Advertised Preferred Lifetime (default: 604800))",
									},
									"valid_lifetime": {
										Type: schema.TypeInt, Optional: true, Default: 2592000, Description: "Specify Valid Lifetime (default:2592000) (Prefix Advertised Valid Lifetime (default: 2592000))",
									},
								},
							},
						},
						"managed_config_action": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Managed Address Configuration flag; 'disable': Disable the Managed Address Configuration flag (default);",
						},
						"other_config_action": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Other Stateful Configuration flag; 'disable': Disable the Other Stateful Configuration flag (default);",
						},
						"adver_vrid": {
							Type: schema.TypeInt, Optional: true, Description: "Specify ha VRRP-A vrid",
						},
						"use_floating_ip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use a floating IP as the source address for Router advertisements",
						},
						"floating_ip": {
							Type: schema.TypeString, Optional: true, Description: "Use a floating IP as the source address for Router advertisements",
						},
						"adver_vrid_default": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Default VRRP-A vrid",
						},
						"use_floating_ip_default_vrid": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use a floating IP as the source address for Router advertisements",
						},
						"floating_ip_default_vrid": {
							Type: schema.TypeString, Optional: true, Description: "Use a floating IP as the source address for Router advertisements",
						},
					},
				},
			},
			"stateful_firewall": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inside (private) interface for stateful firewall",
						},
						"class_list": {
							Type: schema.TypeString, Optional: true, Description: "Class List (Class List Name)",
						},
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Outside (public) interface for stateful firewall",
						},
						"access_list": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Access-list for traffic from the outside",
						},
						"acl_name": {
							Type: schema.TypeString, Optional: true, Description: "Access-list Name",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ttl_ignore": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore TTL decrement for a received packet before sending out",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceEthernetIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceEthernetIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceEthernetIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceEthernetIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceEthernetIpv6AccessListCfg(d []interface{}) edpt.InterfaceEthernetIpv6AccessListCfg {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6AccessListCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.V6AclName = in["v6_acl_name"].(string)
		ret.Inbound = in["inbound"].(int)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6AddressList(d []interface{}) []edpt.InterfaceEthernetIpv6AddressList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6AddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6AddressList
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.AddressType = in["address_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIpv6Ospf466(d []interface{}) edpt.InterfaceEthernetIpv6Ospf466 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6Ospf466
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkList = getSliceInterfaceEthernetIpv6OspfNetworkList467(in["network_list"].([]interface{}))
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
		ret.CostCfg = getSliceInterfaceEthernetIpv6OspfCostCfg468(in["cost_cfg"].([]interface{}))
		ret.DeadIntervalCfg = getSliceInterfaceEthernetIpv6OspfDeadIntervalCfg469(in["dead_interval_cfg"].([]interface{}))
		ret.HelloIntervalCfg = getSliceInterfaceEthernetIpv6OspfHelloIntervalCfg470(in["hello_interval_cfg"].([]interface{}))
		ret.MtuIgnoreCfg = getSliceInterfaceEthernetIpv6OspfMtuIgnoreCfg471(in["mtu_ignore_cfg"].([]interface{}))
		ret.NeighborCfg = getSliceInterfaceEthernetIpv6OspfNeighborCfg472(in["neighbor_cfg"].([]interface{}))
		ret.PriorityCfg = getSliceInterfaceEthernetIpv6OspfPriorityCfg473(in["priority_cfg"].([]interface{}))
		ret.RetransmitIntervalCfg = getSliceInterfaceEthernetIpv6OspfRetransmitIntervalCfg474(in["retransmit_interval_cfg"].([]interface{}))
		ret.TransmitDelayCfg = getSliceInterfaceEthernetIpv6OspfTransmitDelayCfg475(in["transmit_delay_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfNetworkList467(d []interface{}) []edpt.InterfaceEthernetIpv6OspfNetworkList467 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfNetworkList467, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfNetworkList467
		oi.BroadcastType = in["broadcast_type"].(string)
		oi.P2mpNbma = in["p2mp_nbma"].(int)
		oi.NetworkInstanceId = in["network_instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfCostCfg468(d []interface{}) []edpt.InterfaceEthernetIpv6OspfCostCfg468 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfCostCfg468, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfCostCfg468
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfDeadIntervalCfg469(d []interface{}) []edpt.InterfaceEthernetIpv6OspfDeadIntervalCfg469 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfDeadIntervalCfg469, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfDeadIntervalCfg469
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfHelloIntervalCfg470(d []interface{}) []edpt.InterfaceEthernetIpv6OspfHelloIntervalCfg470 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfHelloIntervalCfg470, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfHelloIntervalCfg470
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfMtuIgnoreCfg471(d []interface{}) []edpt.InterfaceEthernetIpv6OspfMtuIgnoreCfg471 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfMtuIgnoreCfg471, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfMtuIgnoreCfg471
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfNeighborCfg472(d []interface{}) []edpt.InterfaceEthernetIpv6OspfNeighborCfg472 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfNeighborCfg472, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfNeighborCfg472
		oi.Neighbor = in["neighbor"].(string)
		oi.NeigInst = in["neig_inst"].(int)
		oi.NeighborCost = in["neighbor_cost"].(int)
		oi.NeighborPollInterval = in["neighbor_poll_interval"].(int)
		oi.NeighborPriority = in["neighbor_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfPriorityCfg473(d []interface{}) []edpt.InterfaceEthernetIpv6OspfPriorityCfg473 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfPriorityCfg473, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfPriorityCfg473
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfRetransmitIntervalCfg474(d []interface{}) []edpt.InterfaceEthernetIpv6OspfRetransmitIntervalCfg474 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfRetransmitIntervalCfg474, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfRetransmitIntervalCfg474
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfTransmitDelayCfg475(d []interface{}) []edpt.InterfaceEthernetIpv6OspfTransmitDelayCfg475 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfTransmitDelayCfg475, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfTransmitDelayCfg475
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIpv6Rip476(d []interface{}) edpt.InterfaceEthernetIpv6Rip476 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6Rip476
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SplitHorizonCfg = getObjectInterfaceEthernetIpv6RipSplitHorizonCfg477(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetIpv6RipSplitHorizonCfg477(d []interface{}) edpt.InterfaceEthernetIpv6RipSplitHorizonCfg477 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6RipSplitHorizonCfg477
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpv6Router478(d []interface{}) edpt.InterfaceEthernetIpv6Router478 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6Router478
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ripng = getObjectInterfaceEthernetIpv6RouterRipng479(in["ripng"].([]interface{}))
		ret.Ospf = getObjectInterfaceEthernetIpv6RouterOspf480(in["ospf"].([]interface{}))
		ret.Isis = getObjectInterfaceEthernetIpv6RouterIsis482(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceEthernetIpv6RouterRipng479(d []interface{}) edpt.InterfaceEthernetIpv6RouterRipng479 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6RouterRipng479
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetIpv6RouterOspf480(d []interface{}) edpt.InterfaceEthernetIpv6RouterOspf480 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6RouterOspf480
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AreaList = getSliceInterfaceEthernetIpv6RouterOspfAreaList481(in["area_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceEthernetIpv6RouterOspfAreaList481(d []interface{}) []edpt.InterfaceEthernetIpv6RouterOspfAreaList481 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6RouterOspfAreaList481, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6RouterOspfAreaList481
		oi.AreaIdNum = in["area_id_num"].(int)
		oi.AreaIdAddr = in["area_id_addr"].(string)
		oi.Tag = in["tag"].(string)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIpv6RouterIsis482(d []interface{}) edpt.InterfaceEthernetIpv6RouterIsis482 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6RouterIsis482
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetIpv6RouterAdver(d []interface{}) edpt.InterfaceEthernetIpv6RouterAdver {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6RouterAdver
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		ret.HopLimit = in["hop_limit"].(int)
		ret.MaxInterval = in["max_interval"].(int)
		ret.MinInterval = in["min_interval"].(int)
		ret.DefaultLifetime = in["default_lifetime"].(int)
		ret.RateLimit = in["rate_limit"].(int)
		ret.ReachableTime = in["reachable_time"].(int)
		ret.RetransmitTimer = in["retransmit_timer"].(int)
		ret.AdverMtuDisable = in["adver_mtu_disable"].(int)
		ret.AdverMtu = in["adver_mtu"].(int)
		ret.PrefixList = getSliceInterfaceEthernetIpv6RouterAdverPrefixList(in["prefix_list"].([]interface{}))
		ret.ManagedConfigAction = in["managed_config_action"].(string)
		ret.OtherConfigAction = in["other_config_action"].(string)
		ret.AdverVrid = in["adver_vrid"].(int)
		ret.UseFloatingIp = in["use_floating_ip"].(int)
		ret.FloatingIp = in["floating_ip"].(string)
		ret.AdverVridDefault = in["adver_vrid_default"].(int)
		ret.UseFloatingIpDefaultVrid = in["use_floating_ip_default_vrid"].(int)
		ret.FloatingIpDefaultVrid = in["floating_ip_default_vrid"].(string)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6RouterAdverPrefixList(d []interface{}) []edpt.InterfaceEthernetIpv6RouterAdverPrefixList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6RouterAdverPrefixList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6RouterAdverPrefixList
		oi.Prefix = in["prefix"].(string)
		oi.NotAutonomous = in["not_autonomous"].(int)
		oi.NotOnLink = in["not_on_link"].(int)
		oi.PreferredLifetime = in["preferred_lifetime"].(int)
		oi.ValidLifetime = in["valid_lifetime"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIpv6StatefulFirewall483(d []interface{}) edpt.InterfaceEthernetIpv6StatefulFirewall483 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6StatefulFirewall483
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Inside = in["inside"].(int)
		ret.ClassList = in["class_list"].(string)
		ret.Outside = in["outside"].(int)
		ret.AccessList = in["access_list"].(int)
		ret.AclName = in["acl_name"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointInterfaceEthernetIpv6(d *schema.ResourceData) edpt.InterfaceEthernetIpv6 {
	var ret edpt.InterfaceEthernetIpv6
	ret.Inst.AccessListCfg = getObjectInterfaceEthernetIpv6AccessListCfg(d.Get("access_list_cfg").([]interface{}))
	ret.Inst.AddressList = getSliceInterfaceEthernetIpv6AddressList(d.Get("address_list").([]interface{}))
	ret.Inst.Inside = d.Get("inside").(int)
	ret.Inst.Ipv6Enable = d.Get("ipv6_enable").(int)
	ret.Inst.Ospf = getObjectInterfaceEthernetIpv6Ospf466(d.Get("ospf").([]interface{}))
	ret.Inst.Outside = d.Get("outside").(int)
	ret.Inst.Rip = getObjectInterfaceEthernetIpv6Rip476(d.Get("rip").([]interface{}))
	ret.Inst.Router = getObjectInterfaceEthernetIpv6Router478(d.Get("router").([]interface{}))
	ret.Inst.RouterAdver = getObjectInterfaceEthernetIpv6RouterAdver(d.Get("router_adver").([]interface{}))
	ret.Inst.StatefulFirewall = getObjectInterfaceEthernetIpv6StatefulFirewall483(d.Get("stateful_firewall").([]interface{}))
	ret.Inst.TtlIgnore = d.Get("ttl_ignore").(int)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
