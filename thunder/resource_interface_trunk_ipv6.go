package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunkIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_trunk_ipv6`: Global IPv6 configuration subcommands\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTrunkIpv6Create,
		UpdateContext: resourceInterfaceTrunkIpv6Update,
		ReadContext:   resourceInterfaceTrunkIpv6Read,
		DeleteContext: resourceInterfaceTrunkIpv6Delete,

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
			"ipv6_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IPv6 processing",
			},
			"nat": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as NAT inside",
						},
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as NAT outside",
						},
					},
				},
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
						"default_lifetime": {
							Type: schema.TypeInt, Optional: true, Default: 1800, Description: "Set Router Advertisement Default Lifetime (default: 1800) (Default Lifetime (seconds))",
						},
						"hop_limit": {
							Type: schema.TypeInt, Optional: true, Default: 255, Description: "Set Router Advertisement Hop Limit (default: 255) (Max Router Advertisement Interval (seconds))",
						},
						"max_interval": {
							Type: schema.TypeInt, Optional: true, Default: 600, Description: "Set Router Advertisement Max Interval (default: 600) (Min Router Advertisement Interval (seconds))",
						},
						"min_interval": {
							Type: schema.TypeInt, Optional: true, Default: 200, Description: "Set Router Advertisement Min Interval (default: 200) (Max Number of Router Solicitations to process per second)",
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
						"mtu": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"adver_mtu_disable": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Disable Router Advertisement MTU Option",
									},
									"adver_mtu": {
										Type: schema.TypeInt, Optional: true, Description: "Set Router Advertisement MTU Option",
									},
								},
							},
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
						"vrid": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
					},
				},
			},
			"ttl_ignore": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore TTL decrement for a received packet",
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
func resourceInterfaceTrunkIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTrunkIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTrunkIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTrunkIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceTrunkIpv6AccessListCfg(d []interface{}) edpt.InterfaceTrunkIpv6AccessListCfg {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6AccessListCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.V6AclName = in["v6_acl_name"].(string)
		ret.Inbound = in["inbound"].(int)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6AddressList(d []interface{}) []edpt.InterfaceTrunkIpv6AddressList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6AddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6AddressList
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.AddressType = in["address_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIpv6Nat(d []interface{}) edpt.InterfaceTrunkIpv6Nat {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6Nat
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
	}
	return ret
}

func getObjectInterfaceTrunkIpv6Ospf851(d []interface{}) edpt.InterfaceTrunkIpv6Ospf851 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6Ospf851
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkList = getSliceInterfaceTrunkIpv6OspfNetworkList852(in["network_list"].([]interface{}))
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
		ret.CostCfg = getSliceInterfaceTrunkIpv6OspfCostCfg853(in["cost_cfg"].([]interface{}))
		ret.DeadIntervalCfg = getSliceInterfaceTrunkIpv6OspfDeadIntervalCfg854(in["dead_interval_cfg"].([]interface{}))
		ret.HelloIntervalCfg = getSliceInterfaceTrunkIpv6OspfHelloIntervalCfg855(in["hello_interval_cfg"].([]interface{}))
		ret.MtuIgnoreCfg = getSliceInterfaceTrunkIpv6OspfMtuIgnoreCfg856(in["mtu_ignore_cfg"].([]interface{}))
		ret.NeighborCfg = getSliceInterfaceTrunkIpv6OspfNeighborCfg857(in["neighbor_cfg"].([]interface{}))
		ret.PriorityCfg = getSliceInterfaceTrunkIpv6OspfPriorityCfg858(in["priority_cfg"].([]interface{}))
		ret.RetransmitIntervalCfg = getSliceInterfaceTrunkIpv6OspfRetransmitIntervalCfg859(in["retransmit_interval_cfg"].([]interface{}))
		ret.TransmitDelayCfg = getSliceInterfaceTrunkIpv6OspfTransmitDelayCfg860(in["transmit_delay_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfNetworkList852(d []interface{}) []edpt.InterfaceTrunkIpv6OspfNetworkList852 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfNetworkList852, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfNetworkList852
		oi.BroadcastType = in["broadcast_type"].(string)
		oi.P2mpNbma = in["p2mp_nbma"].(int)
		oi.NetworkInstanceId = in["network_instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfCostCfg853(d []interface{}) []edpt.InterfaceTrunkIpv6OspfCostCfg853 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfCostCfg853, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfCostCfg853
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfDeadIntervalCfg854(d []interface{}) []edpt.InterfaceTrunkIpv6OspfDeadIntervalCfg854 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfDeadIntervalCfg854, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfDeadIntervalCfg854
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfHelloIntervalCfg855(d []interface{}) []edpt.InterfaceTrunkIpv6OspfHelloIntervalCfg855 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfHelloIntervalCfg855, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfHelloIntervalCfg855
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfMtuIgnoreCfg856(d []interface{}) []edpt.InterfaceTrunkIpv6OspfMtuIgnoreCfg856 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfMtuIgnoreCfg856, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfMtuIgnoreCfg856
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfNeighborCfg857(d []interface{}) []edpt.InterfaceTrunkIpv6OspfNeighborCfg857 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfNeighborCfg857, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfNeighborCfg857
		oi.Neighbor = in["neighbor"].(string)
		oi.NeigInst = in["neig_inst"].(int)
		oi.NeighborCost = in["neighbor_cost"].(int)
		oi.NeighborPollInterval = in["neighbor_poll_interval"].(int)
		oi.NeighborPriority = in["neighbor_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfPriorityCfg858(d []interface{}) []edpt.InterfaceTrunkIpv6OspfPriorityCfg858 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfPriorityCfg858, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfPriorityCfg858
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfRetransmitIntervalCfg859(d []interface{}) []edpt.InterfaceTrunkIpv6OspfRetransmitIntervalCfg859 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfRetransmitIntervalCfg859, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfRetransmitIntervalCfg859
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfTransmitDelayCfg860(d []interface{}) []edpt.InterfaceTrunkIpv6OspfTransmitDelayCfg860 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfTransmitDelayCfg860, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfTransmitDelayCfg860
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIpv6Rip861(d []interface{}) edpt.InterfaceTrunkIpv6Rip861 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6Rip861
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SplitHorizonCfg = getObjectInterfaceTrunkIpv6RipSplitHorizonCfg862(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkIpv6RipSplitHorizonCfg862(d []interface{}) edpt.InterfaceTrunkIpv6RipSplitHorizonCfg862 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6RipSplitHorizonCfg862
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpv6Router863(d []interface{}) edpt.InterfaceTrunkIpv6Router863 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6Router863
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ripng = getObjectInterfaceTrunkIpv6RouterRipng864(in["ripng"].([]interface{}))
		ret.Ospf = getObjectInterfaceTrunkIpv6RouterOspf865(in["ospf"].([]interface{}))
		ret.Isis = getObjectInterfaceTrunkIpv6RouterIsis867(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTrunkIpv6RouterRipng864(d []interface{}) edpt.InterfaceTrunkIpv6RouterRipng864 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6RouterRipng864
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkIpv6RouterOspf865(d []interface{}) edpt.InterfaceTrunkIpv6RouterOspf865 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6RouterOspf865
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AreaList = getSliceInterfaceTrunkIpv6RouterOspfAreaList866(in["area_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceTrunkIpv6RouterOspfAreaList866(d []interface{}) []edpt.InterfaceTrunkIpv6RouterOspfAreaList866 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6RouterOspfAreaList866, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6RouterOspfAreaList866
		oi.AreaIdNum = in["area_id_num"].(int)
		oi.AreaIdAddr = in["area_id_addr"].(string)
		oi.Tag = in["tag"].(string)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIpv6RouterIsis867(d []interface{}) edpt.InterfaceTrunkIpv6RouterIsis867 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6RouterIsis867
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkIpv6RouterAdver(d []interface{}) edpt.InterfaceTrunkIpv6RouterAdver {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6RouterAdver
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		ret.DefaultLifetime = in["default_lifetime"].(int)
		ret.HopLimit = in["hop_limit"].(int)
		ret.MaxInterval = in["max_interval"].(int)
		ret.MinInterval = in["min_interval"].(int)
		ret.RateLimit = in["rate_limit"].(int)
		ret.ReachableTime = in["reachable_time"].(int)
		ret.RetransmitTimer = in["retransmit_timer"].(int)
		ret.Mtu = getObjectInterfaceTrunkIpv6RouterAdverMtu(in["mtu"].([]interface{}))
		ret.PrefixList = getSliceInterfaceTrunkIpv6RouterAdverPrefixList(in["prefix_list"].([]interface{}))
		ret.ManagedConfigAction = in["managed_config_action"].(string)
		ret.OtherConfigAction = in["other_config_action"].(string)
		ret.Vrid = getObjectInterfaceTrunkIpv6RouterAdverVrid(in["vrid"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTrunkIpv6RouterAdverMtu(d []interface{}) edpt.InterfaceTrunkIpv6RouterAdverMtu {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6RouterAdverMtu
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AdverMtuDisable = in["adver_mtu_disable"].(int)
		ret.AdverMtu = in["adver_mtu"].(int)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6RouterAdverPrefixList(d []interface{}) []edpt.InterfaceTrunkIpv6RouterAdverPrefixList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6RouterAdverPrefixList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6RouterAdverPrefixList
		oi.Prefix = in["prefix"].(string)
		oi.NotAutonomous = in["not_autonomous"].(int)
		oi.NotOnLink = in["not_on_link"].(int)
		oi.PreferredLifetime = in["preferred_lifetime"].(int)
		oi.ValidLifetime = in["valid_lifetime"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIpv6RouterAdverVrid(d []interface{}) edpt.InterfaceTrunkIpv6RouterAdverVrid {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6RouterAdverVrid
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AdverVrid = in["adver_vrid"].(int)
		ret.UseFloatingIp = in["use_floating_ip"].(int)
		ret.FloatingIp = in["floating_ip"].(string)
		ret.AdverVridDefault = in["adver_vrid_default"].(int)
		ret.UseFloatingIpDefaultVrid = in["use_floating_ip_default_vrid"].(int)
		ret.FloatingIpDefaultVrid = in["floating_ip_default_vrid"].(string)
	}
	return ret
}

func dataToEndpointInterfaceTrunkIpv6(d *schema.ResourceData) edpt.InterfaceTrunkIpv6 {
	var ret edpt.InterfaceTrunkIpv6
	ret.Inst.AccessListCfg = getObjectInterfaceTrunkIpv6AccessListCfg(d.Get("access_list_cfg").([]interface{}))
	ret.Inst.AddressList = getSliceInterfaceTrunkIpv6AddressList(d.Get("address_list").([]interface{}))
	ret.Inst.Ipv6Enable = d.Get("ipv6_enable").(int)
	ret.Inst.Nat = getObjectInterfaceTrunkIpv6Nat(d.Get("nat").([]interface{}))
	ret.Inst.Ospf = getObjectInterfaceTrunkIpv6Ospf851(d.Get("ospf").([]interface{}))
	ret.Inst.Rip = getObjectInterfaceTrunkIpv6Rip861(d.Get("rip").([]interface{}))
	ret.Inst.Router = getObjectInterfaceTrunkIpv6Router863(d.Get("router").([]interface{}))
	ret.Inst.RouterAdver = getObjectInterfaceTrunkIpv6RouterAdver(d.Get("router_adver").([]interface{}))
	ret.Inst.TtlIgnore = d.Get("ttl_ignore").(int)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
