package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTunnel() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_tunnel`: Tunnel interface\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTunnelCreate,
		UpdateContext: resourceInterfaceTunnelUpdate,
		ReadContext:   resourceInterfaceTunnelRead,
		DeleteContext: resourceInterfaceTunnelDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable; 'disable': Disable;",
			},
			"ifnum": {
				Type: schema.TypeInt, Required: true, Description: "Tunnel interface number",
			},
			"ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dhcp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use DHCP to configure IP address",
									},
									"ip_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv4_address": {
													Type: schema.TypeString, Optional: true, Description: "IP address",
												},
												"ipv4_netmask": {
													Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
												},
											},
										},
									},
								},
							},
						},
						"generate_membership_query": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Membership Query",
						},
						"generate_membership_query_val": {
							Type: schema.TypeInt, Optional: true, Default: 125, Description: "1 - 255 (Default is 125)",
						},
						"max_resp_time": {
							Type: schema.TypeInt, Optional: true, Default: 100, Description: "Max Response Time (Default is 100)",
						},
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as inside",
						},
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as outside",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"rip": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"str": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"string": {
																Type: schema.TypeString, Optional: true, Description: "The RIP authentication string",
															},
														},
													},
												},
												"mode": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": {
																Type: schema.TypeString, Optional: true, Default: "text", Description: "'md5': Keyed message digest; 'text': Clear text authentication;",
															},
														},
													},
												},
												"key_chain": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key_chain": {
																Type: schema.TypeString, Optional: true, Description: "Authentication key-chain (Name of key-chain)",
															},
														},
													},
												},
											},
										},
									},
									"send_packet": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable sending packets through the specified interface",
									},
									"receive_packet": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable receiving packet through the specified interface",
									},
									"send_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"send": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertisement transmission",
												},
												"version": {
													Type: schema.TypeString, Optional: true, Description: "'2': RIP version 2;",
												},
											},
										},
									},
									"receive_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"receive": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertisement reception",
												},
												"version": {
													Type: schema.TypeString, Optional: true, Description: "'2': RIP version 2;",
												},
											},
										},
									},
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
						"ospf": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ospf_global": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"authentication_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"authentication": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable authentication",
															},
															"value": {
																Type: schema.TypeString, Optional: true, Description: "'message-digest': Use message-digest authentication; 'null': Use no authentication;",
															},
														},
													},
												},
												"authentication_key": {
													Type: schema.TypeString, Optional: true, Description: "Authentication password (key) (The OSPF password (key))",
												},
												"bfd_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bfd": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
															},
															"disable": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable BFD",
															},
														},
													},
												},
												"cost": {
													Type: schema.TypeInt, Optional: true, Description: "Interface cost",
												},
												"database_filter_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"database_filter": {
																Type: schema.TypeString, Optional: true, Description: "'all': Filter all LSA;",
															},
															"out": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Outgoing LSA",
															},
														},
													},
												},
												"dead_interval": {
													Type: schema.TypeInt, Optional: true, Default: 40, Description: "Interval after which a neighbor is declared dead (Seconds)",
												},
												"disable": {
													Type: schema.TypeString, Optional: true, Description: "'all': All functionality;",
												},
												"hello_interval": {
													Type: schema.TypeInt, Optional: true, Default: 10, Description: "Time between HELLO packets (Seconds)",
												},
												"message_digest_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"message_digest_key": {
																Type: schema.TypeInt, Optional: true, Description: "Message digest authentication password (key) (Key id)",
															},
															"md5": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"md5_value": {
																			Type: schema.TypeString, Optional: true, Description: "The OSPF password (1-16)",
																		},
																	},
																},
															},
														},
													},
												},
												"mtu": {
													Type: schema.TypeInt, Optional: true, Description: "OSPF interface MTU (MTU size)",
												},
												"mtu_ignore": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignores the MTU in DBD packets",
												},
												"network": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"broadcast": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF broadcast multi-access network",
															},
															"non_broadcast": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF NBMA network",
															},
															"point_to_point": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF point-to-point network",
															},
															"point_to_multipoint": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF point-to-multipoint network",
															},
															"p2mp_nbma": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify non-broadcast point-to-multipoint network",
															},
														},
													},
												},
												"priority": {
													Type: schema.TypeInt, Optional: true, Default: 1, Description: "Router priority",
												},
												"retransmit_interval": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Time between retransmitting lost link state advertisements (Seconds)",
												},
												"transmit_delay": {
													Type: schema.TypeInt, Optional: true, Default: 1, Description: "Link state transmit delay (Seconds)",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"ospf_ip_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_addr": {
													Type: schema.TypeString, Required: true, Description: "Address of interface",
												},
												"authentication": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable authentication",
												},
												"value": {
													Type: schema.TypeString, Optional: true, Description: "'message-digest': Use message-digest authentication; 'null': Use no authentication;",
												},
												"authentication_key": {
													Type: schema.TypeString, Optional: true, Description: "Authentication password (key) (The OSPF password (key))",
												},
												"cost": {
													Type: schema.TypeInt, Optional: true, Description: "Interface cost",
												},
												"database_filter": {
													Type: schema.TypeString, Optional: true, Description: "'all': Filter all LSA;",
												},
												"out": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Outgoing LSA",
												},
												"dead_interval": {
													Type: schema.TypeInt, Optional: true, Default: 40, Description: "Interval after which a neighbor is declared dead (Seconds)",
												},
												"hello_interval": {
													Type: schema.TypeInt, Optional: true, Default: 10, Description: "Time between HELLO packets (Seconds)",
												},
												"message_digest_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"message_digest_key": {
																Type: schema.TypeInt, Optional: true, Description: "Message digest authentication password (key) (Key id)",
															},
															"md5_value": {
																Type: schema.TypeString, Optional: true, Description: "The OSPF password (1-16)",
															},
														},
													},
												},
												"mtu_ignore": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignores the MTU in DBD packets",
												},
												"priority": {
													Type: schema.TypeInt, Optional: true, Default: 1, Description: "Router priority",
												},
												"retransmit_interval": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Time between retransmitting lost link state advertisements (Seconds)",
												},
												"transmit_delay": {
													Type: schema.TypeInt, Optional: true, Default: 1, Description: "Link state transmit delay (Seconds)",
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
					},
				},
			},
			"ipv6": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address_cfg": {
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
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as inside",
						},
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as outside",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
					},
				},
			},
			"load_interval": {
				Type: schema.TypeInt, Optional: true, Default: 300, Description: "Configure Load Interval (Seconds (5-300, Multiple of 5), default 300)",
			},
			"lw_4o6": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure LW-4over6 inside interface",
						},
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure LW-4over6 outside interface",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"map": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure MAP inside interface (connected to MAP domains)",
						},
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure MAP outside interface",
						},
						"map_t_inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure MAP inside interface (connected to MAP domains)",
						},
						"map_t_outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure MAP outside interface",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"mtu": {
				Type: schema.TypeInt, Optional: true, Description: "Interface mtu (Interface MTU, default 1 (min MTU is 1280 for IPv6))",
			},
			"name": {
				Type: schema.TypeString, Optional: true, Description: "Name for the interface",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'num-rx-pkts': received packets; 'num-total-rx-bytes': received bytes; 'num-tx-pkts': sent packets; 'num-total-tx-bytes': sent bytes; 'num-rx-err-pkts': received error packets; 'num-tx-err-pkts': sent error packets; 'rate_pkt_sent': Packet sent rate packets/sec; 'rate_byte_sent': Byte sent rate bits/sec; 'rate_pkt_rcvd': Packet received rate packets/sec; 'rate_byte_rcvd': Byte received rate bits/sec;",
						},
					},
				},
			},
			"speed": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Speed in Gbit/Sec (Default 10 Gbps)",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceInterfaceTunnelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnel(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTunnelRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTunnelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnel(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTunnelRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTunnelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnel(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnel(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceTunnelIp967(d []interface{}) edpt.InterfaceTunnelIp967 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIp967
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Address = getObjectInterfaceTunnelIpAddress968(in["address"].([]interface{}))
		ret.GenerateMembershipQuery = in["generate_membership_query"].(int)
		ret.GenerateMembershipQueryVal = in["generate_membership_query_val"].(int)
		ret.MaxRespTime = in["max_resp_time"].(int)
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
		//omit uuid
		ret.Rip = getObjectInterfaceTunnelIpRip970(in["rip"].([]interface{}))
		ret.Ospf = getObjectInterfaceTunnelIpOspf978(in["ospf"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTunnelIpAddress968(d []interface{}) edpt.InterfaceTunnelIpAddress968 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpAddress968
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dhcp = in["dhcp"].(int)
		ret.IpCfg = getSliceInterfaceTunnelIpAddressIpCfg969(in["ip_cfg"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceTunnelIpAddressIpCfg969(d []interface{}) []edpt.InterfaceTunnelIpAddressIpCfg969 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpAddressIpCfg969, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpAddressIpCfg969
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTunnelIpRip970(d []interface{}) edpt.InterfaceTunnelIpRip970 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRip970
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceTunnelIpRipAuthentication971(in["authentication"].([]interface{}))
		ret.SendPacket = in["send_packet"].(int)
		ret.ReceivePacket = in["receive_packet"].(int)
		ret.SendCfg = getObjectInterfaceTunnelIpRipSendCfg975(in["send_cfg"].([]interface{}))
		ret.ReceiveCfg = getObjectInterfaceTunnelIpRipReceiveCfg976(in["receive_cfg"].([]interface{}))
		ret.SplitHorizonCfg = getObjectInterfaceTunnelIpRipSplitHorizonCfg977(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTunnelIpRipAuthentication971(d []interface{}) edpt.InterfaceTunnelIpRipAuthentication971 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipAuthentication971
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Str = getObjectInterfaceTunnelIpRipAuthenticationStr972(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceTunnelIpRipAuthenticationMode973(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceTunnelIpRipAuthenticationKeyChain974(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTunnelIpRipAuthenticationStr972(d []interface{}) edpt.InterfaceTunnelIpRipAuthenticationStr972 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipAuthenticationStr972
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpRipAuthenticationMode973(d []interface{}) edpt.InterfaceTunnelIpRipAuthenticationMode973 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipAuthenticationMode973
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpRipAuthenticationKeyChain974(d []interface{}) edpt.InterfaceTunnelIpRipAuthenticationKeyChain974 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipAuthenticationKeyChain974
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpRipSendCfg975(d []interface{}) edpt.InterfaceTunnelIpRipSendCfg975 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipSendCfg975
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpRipReceiveCfg976(d []interface{}) edpt.InterfaceTunnelIpRipReceiveCfg976 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipReceiveCfg976
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpRipSplitHorizonCfg977(d []interface{}) edpt.InterfaceTunnelIpRipSplitHorizonCfg977 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipSplitHorizonCfg977
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpOspf978(d []interface{}) edpt.InterfaceTunnelIpOspf978 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpOspf978
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OspfGlobal = getObjectInterfaceTunnelIpOspfOspfGlobal979(in["ospf_global"].([]interface{}))
		ret.OspfIpList = getSliceInterfaceTunnelIpOspfOspfIpList986(in["ospf_ip_list"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTunnelIpOspfOspfGlobal979(d []interface{}) edpt.InterfaceTunnelIpOspfOspfGlobal979 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpOspfOspfGlobal979
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthenticationCfg = getObjectInterfaceTunnelIpOspfOspfGlobalAuthenticationCfg980(in["authentication_cfg"].([]interface{}))
		ret.AuthenticationKey = in["authentication_key"].(string)
		ret.BfdCfg = getObjectInterfaceTunnelIpOspfOspfGlobalBfdCfg981(in["bfd_cfg"].([]interface{}))
		ret.Cost = in["cost"].(int)
		ret.DatabaseFilterCfg = getObjectInterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg982(in["database_filter_cfg"].([]interface{}))
		ret.DeadInterval = in["dead_interval"].(int)
		ret.Disable = in["disable"].(string)
		ret.HelloInterval = in["hello_interval"].(int)
		ret.MessageDigestCfg = getSliceInterfaceTunnelIpOspfOspfGlobalMessageDigestCfg983(in["message_digest_cfg"].([]interface{}))
		ret.Mtu = in["mtu"].(int)
		ret.MtuIgnore = in["mtu_ignore"].(int)
		ret.Network = getObjectInterfaceTunnelIpOspfOspfGlobalNetwork985(in["network"].([]interface{}))
		ret.Priority = in["priority"].(int)
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTunnelIpOspfOspfGlobalAuthenticationCfg980(d []interface{}) edpt.InterfaceTunnelIpOspfOspfGlobalAuthenticationCfg980 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpOspfOspfGlobalAuthenticationCfg980
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = in["authentication"].(int)
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpOspfOspfGlobalBfdCfg981(d []interface{}) edpt.InterfaceTunnelIpOspfOspfGlobalBfdCfg981 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpOspfOspfGlobalBfdCfg981
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectInterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg982(d []interface{}) edpt.InterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg982 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg982
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DatabaseFilter = in["database_filter"].(string)
		ret.Out = in["out"].(int)
	}
	return ret
}

func getSliceInterfaceTunnelIpOspfOspfGlobalMessageDigestCfg983(d []interface{}) []edpt.InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg983 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg983, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg983
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = getObjectInterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5984(in["md5"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5984(d []interface{}) edpt.InterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5984 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5984
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Md5Value = in["md5_value"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceTunnelIpOspfOspfGlobalNetwork985(d []interface{}) edpt.InterfaceTunnelIpOspfOspfGlobalNetwork985 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpOspfOspfGlobalNetwork985
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Broadcast = in["broadcast"].(int)
		ret.NonBroadcast = in["non_broadcast"].(int)
		ret.PointToPoint = in["point_to_point"].(int)
		ret.PointToMultipoint = in["point_to_multipoint"].(int)
		ret.P2mpNbma = in["p2mp_nbma"].(int)
	}
	return ret
}

func getSliceInterfaceTunnelIpOspfOspfIpList986(d []interface{}) []edpt.InterfaceTunnelIpOspfOspfIpList986 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpOspfOspfIpList986, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpOspfOspfIpList986
		oi.IpAddr = in["ip_addr"].(string)
		oi.Authentication = in["authentication"].(int)
		oi.Value = in["value"].(string)
		oi.AuthenticationKey = in["authentication_key"].(string)
		oi.Cost = in["cost"].(int)
		oi.DatabaseFilter = in["database_filter"].(string)
		oi.Out = in["out"].(int)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.MessageDigestCfg = getSliceInterfaceTunnelIpOspfOspfIpListMessageDigestCfg987(in["message_digest_cfg"].([]interface{}))
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.Priority = in["priority"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpOspfOspfIpListMessageDigestCfg987(d []interface{}) []edpt.InterfaceTunnelIpOspfOspfIpListMessageDigestCfg987 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpOspfOspfIpListMessageDigestCfg987, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpOspfOspfIpListMessageDigestCfg987
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5Value = in["md5_value"].(string)
		//omit encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTunnelIpv6988(d []interface{}) edpt.InterfaceTunnelIpv6988 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpv6988
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AddressCfg = getSliceInterfaceTunnelIpv6AddressCfg989(in["address_cfg"].([]interface{}))
		ret.Ipv6Enable = in["ipv6_enable"].(int)
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
		//omit uuid
		ret.Router = getObjectInterfaceTunnelIpv6Router990(in["router"].([]interface{}))
		ret.Ospf = getObjectInterfaceTunnelIpv6Ospf994(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceTunnelIpv6AddressCfg989(d []interface{}) []edpt.InterfaceTunnelIpv6AddressCfg989 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6AddressCfg989, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6AddressCfg989
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.AddressType = in["address_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTunnelIpv6Router990(d []interface{}) edpt.InterfaceTunnelIpv6Router990 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpv6Router990
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ripng = getObjectInterfaceTunnelIpv6RouterRipng991(in["ripng"].([]interface{}))
		ret.Ospf = getObjectInterfaceTunnelIpv6RouterOspf992(in["ospf"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTunnelIpv6RouterRipng991(d []interface{}) edpt.InterfaceTunnelIpv6RouterRipng991 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpv6RouterRipng991
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTunnelIpv6RouterOspf992(d []interface{}) edpt.InterfaceTunnelIpv6RouterOspf992 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpv6RouterOspf992
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AreaList = getSliceInterfaceTunnelIpv6RouterOspfAreaList993(in["area_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceTunnelIpv6RouterOspfAreaList993(d []interface{}) []edpt.InterfaceTunnelIpv6RouterOspfAreaList993 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6RouterOspfAreaList993, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6RouterOspfAreaList993
		oi.AreaIdNum = in["area_id_num"].(int)
		oi.AreaIdAddr = in["area_id_addr"].(string)
		oi.Tag = in["tag"].(string)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTunnelIpv6Ospf994(d []interface{}) edpt.InterfaceTunnelIpv6Ospf994 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpv6Ospf994
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkList = getSliceInterfaceTunnelIpv6OspfNetworkList995(in["network_list"].([]interface{}))
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
		ret.CostCfg = getSliceInterfaceTunnelIpv6OspfCostCfg996(in["cost_cfg"].([]interface{}))
		ret.DeadIntervalCfg = getSliceInterfaceTunnelIpv6OspfDeadIntervalCfg997(in["dead_interval_cfg"].([]interface{}))
		ret.HelloIntervalCfg = getSliceInterfaceTunnelIpv6OspfHelloIntervalCfg998(in["hello_interval_cfg"].([]interface{}))
		ret.MtuIgnoreCfg = getSliceInterfaceTunnelIpv6OspfMtuIgnoreCfg999(in["mtu_ignore_cfg"].([]interface{}))
		ret.NeighborCfg = getSliceInterfaceTunnelIpv6OspfNeighborCfg1000(in["neighbor_cfg"].([]interface{}))
		ret.PriorityCfg = getSliceInterfaceTunnelIpv6OspfPriorityCfg1001(in["priority_cfg"].([]interface{}))
		ret.RetransmitIntervalCfg = getSliceInterfaceTunnelIpv6OspfRetransmitIntervalCfg1002(in["retransmit_interval_cfg"].([]interface{}))
		ret.TransmitDelayCfg = getSliceInterfaceTunnelIpv6OspfTransmitDelayCfg1003(in["transmit_delay_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfNetworkList995(d []interface{}) []edpt.InterfaceTunnelIpv6OspfNetworkList995 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfNetworkList995, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfNetworkList995
		oi.BroadcastType = in["broadcast_type"].(string)
		oi.P2mpNbma = in["p2mp_nbma"].(int)
		oi.NetworkInstanceId = in["network_instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfCostCfg996(d []interface{}) []edpt.InterfaceTunnelIpv6OspfCostCfg996 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfCostCfg996, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfCostCfg996
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfDeadIntervalCfg997(d []interface{}) []edpt.InterfaceTunnelIpv6OspfDeadIntervalCfg997 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfDeadIntervalCfg997, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfDeadIntervalCfg997
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfHelloIntervalCfg998(d []interface{}) []edpt.InterfaceTunnelIpv6OspfHelloIntervalCfg998 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfHelloIntervalCfg998, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfHelloIntervalCfg998
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfMtuIgnoreCfg999(d []interface{}) []edpt.InterfaceTunnelIpv6OspfMtuIgnoreCfg999 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfMtuIgnoreCfg999, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfMtuIgnoreCfg999
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfNeighborCfg1000(d []interface{}) []edpt.InterfaceTunnelIpv6OspfNeighborCfg1000 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfNeighborCfg1000, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfNeighborCfg1000
		oi.Neighbor = in["neighbor"].(string)
		oi.NeigInst = in["neig_inst"].(int)
		oi.NeighborCost = in["neighbor_cost"].(int)
		oi.NeighborPollInterval = in["neighbor_poll_interval"].(int)
		oi.NeighborPriority = in["neighbor_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfPriorityCfg1001(d []interface{}) []edpt.InterfaceTunnelIpv6OspfPriorityCfg1001 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfPriorityCfg1001, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfPriorityCfg1001
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfRetransmitIntervalCfg1002(d []interface{}) []edpt.InterfaceTunnelIpv6OspfRetransmitIntervalCfg1002 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfRetransmitIntervalCfg1002, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfRetransmitIntervalCfg1002
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfTransmitDelayCfg1003(d []interface{}) []edpt.InterfaceTunnelIpv6OspfTransmitDelayCfg1003 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfTransmitDelayCfg1003, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfTransmitDelayCfg1003
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTunnelLw4o61004(d []interface{}) edpt.InterfaceTunnelLw4o61004 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelLw4o61004
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Outside = in["outside"].(int)
		ret.Inside = in["inside"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTunnelMap1005(d []interface{}) edpt.InterfaceTunnelMap1005 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelMap1005
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
		ret.MapTInside = in["map_t_inside"].(int)
		ret.MapTOutside = in["map_t_outside"].(int)
		//omit uuid
	}
	return ret
}

func getSliceInterfaceTunnelSamplingEnable(d []interface{}) []edpt.InterfaceTunnelSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceTunnel(d *schema.ResourceData) edpt.InterfaceTunnel {
	var ret edpt.InterfaceTunnel
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Ifnum = d.Get("ifnum").(int)
	ret.Inst.Ip = getObjectInterfaceTunnelIp967(d.Get("ip").([]interface{}))
	ret.Inst.Ipv6 = getObjectInterfaceTunnelIpv6988(d.Get("ipv6").([]interface{}))
	ret.Inst.LoadInterval = d.Get("load_interval").(int)
	ret.Inst.Lw4o6 = getObjectInterfaceTunnelLw4o61004(d.Get("lw_4o6").([]interface{}))
	ret.Inst.Map = getObjectInterfaceTunnelMap1005(d.Get("map").([]interface{}))
	ret.Inst.Mtu = d.Get("mtu").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.SamplingEnable = getSliceInterfaceTunnelSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Speed = d.Get("speed").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
