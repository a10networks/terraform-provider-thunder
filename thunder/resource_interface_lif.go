package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLif() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_lif`: Logical interface\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLifCreate,
		UpdateContext: resourceInterfaceLifUpdate,
		ReadContext:   resourceInterfaceLifRead,
		DeleteContext: resourceInterfaceLifDelete,

		Schema: map[string]*schema.Schema{
			"access_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_id": {
							Type: schema.TypeInt, Optional: true, Description: "ACL id",
						},
						"acl_name": {
							Type: schema.TypeString, Optional: true, Description: "Apply an access list (Named Access List)",
						},
					},
				},
			},
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable; 'disable': Disable;",
			},
			"bfd": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key_id": {
										Type: schema.TypeInt, Optional: true, Description: "Key ID",
									},
									"method": {
										Type: schema.TypeString, Optional: true, Description: "'md5': Keyed MD5; 'meticulous-md5': Meticulous Keyed MD5; 'meticulous-sha1': Meticulous Keyed SHA1; 'sha1': Keyed SHA1; 'simple': Simple Password;",
									},
									"password": {
										Type: schema.TypeString, Optional: true, Description: "Key String",
									},
								},
							},
						},
						"echo": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable BFD Echo",
						},
						"demand": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Demand mode",
						},
						"interval_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interval": {
										Type: schema.TypeInt, Optional: true, Description: "Transmit interval between BFD packets (Milliseconds)",
									},
									"min_rx": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum receive interval capability (Milliseconds)",
									},
									"multiplier": {
										Type: schema.TypeInt, Optional: true, Description: "Multiplier value used to compute holddown (value used to multiply the interval)",
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
			"encapsulation": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dot1q": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag": {
										Type: schema.TypeInt, Optional: true, Description: "Tag ID",
									},
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet Interface (Ethernet port number)",
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk Interface (Trunk number)",
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
			"ifname": {
				Type: schema.TypeString, Required: true, Description: "Lif interface name",
			},
			"ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dhcp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use DHCP to configure IP address",
						},
						"address_list": {
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
						"allow_promiscuous_vip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow traffic to be associated with promiscuous VIP",
						},
						"cache_spoofing_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "This interface connects to spoofing cache",
						},
						"client": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Client facing interface for IPv4/v6 traffic",
						},
						"server": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Server facing interface for IPv4/v6 traffic",
						},
						"dmz": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "DMZ network facing interface for IPv4/v6 traffic",
						},
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as inside",
						},
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as outside",
						},
						"generate_membership_query": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Membership Query",
						},
						"query_interval": {
							Type: schema.TypeInt, Optional: true, Default: 125, Description: "1 - 255 (Default is 125)",
						},
						"max_resp_time": {
							Type: schema.TypeInt, Optional: true, Default: 100, Description: "Maximum Response Time (Max Response Time (Default is 100))",
						},
						"unnumbered": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set the interface as unnumbered",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"router": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
													Type: schema.TypeString, Optional: true, Description: "'1': RIP version 1; '2': RIP version 2; '1-compatible': RIPv1-compatible; '1-2': RIP version 1 & 2;",
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
													Type: schema.TypeString, Optional: true, Description: "'1': RIP version 1; '2': RIP version 2; '1-2': RIP version 1 & 2;",
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
						"address_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_addr": {
										Type: schema.TypeString, Optional: true, Description: "Set the IPv6 address of an interface",
									},
									"anycast": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure an IPv6 anycast address",
									},
									"link_local": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure an IPv6 link local address",
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
			"isis": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"send_only_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"send_only": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Authentication send-only",
												},
												"level": {
													Type: schema.TypeString, Optional: true, Description: "'level-1': Specify authentication send-only for level-1 PDUs; 'level-2': Specify authentication send-only for level-2 PDUs;",
												},
											},
										},
									},
									"mode_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mode": {
													Type: schema.TypeString, Optional: true, Description: "'md5': Keyed message digest;",
												},
												"level": {
													Type: schema.TypeString, Optional: true, Description: "'level-1': Specify authentication mode for level-1 PDUs; 'level-2': Specify authentication mode for level-2 PDUs;",
												},
											},
										},
									},
									"key_chain_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key_chain": {
													Type: schema.TypeString, Optional: true, Description: "Authentication key-chain (Name of key-chain)",
												},
												"level": {
													Type: schema.TypeString, Optional: true, Description: "'level-1': Specify authentication key-chain for level-1 PDUs; 'level-2': Specify authentication key-chain for level-2 PDUs;",
												},
											},
										},
									},
								},
							},
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
						"circuit_type": {
							Type: schema.TypeString, Optional: true, Default: "level-1-2", Description: "'level-1': Level-1 only adjacencies are formed; 'level-1-2': Level-1-2 adjacencies are formed; 'level-2-only': Level-2 only adjacencies are formed;",
						},
						"csnp_interval_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"csnp_interval": {
										Type: schema.TypeInt, Optional: true, Default: 10, Description: "Set CSNP interval in seconds (CSNP interval value)",
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Speficy interval for level-1 CSNPs; 'level-2': Specify interval for level-2 CSNPs;",
									},
								},
							},
						},
						"padding": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Add padding to IS-IS hello packets",
						},
						"hello_interval_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hello_interval": {
										Type: schema.TypeInt, Optional: true, Default: 10, Description: "Set Hello interval in seconds (Hello interval value)",
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Specify hello-interval for level-1 IIHs; 'level-2': Specify hello-interval for level-2 IIHs;",
									},
								},
							},
						},
						"hello_interval_minimal_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hello_interval_minimal": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set Hello holdtime 1 second, interval depends on multiplier",
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Specify hello-interval for level-1 IIHs; 'level-2': Specify hello-interval for level-2 IIHs;",
									},
								},
							},
						},
						"hello_multiplier_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hello_multiplier": {
										Type: schema.TypeInt, Optional: true, Default: 3, Description: "Set multiplier for Hello holding time (Hello multiplier value)",
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Specify hello multiplier for level-1 IIHs; 'level-2': Specify hello multiplier for level-2 IIHs;",
									},
								},
							},
						},
						"lsp_interval": {
							Type: schema.TypeInt, Optional: true, Default: 33, Description: "Set LSP transmission interval (LSP transmission interval (milliseconds))",
						},
						"mesh_group": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "Mesh group number",
									},
									"blocked": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Block LSPs on this interface",
									},
								},
							},
						},
						"metric_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"metric": {
										Type: schema.TypeInt, Optional: true, Default: 10, Description: "Configure the metric for interface (Default metric)",
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Apply metric to level-1 links; 'level-2': Apply metric to level-2 links;",
									},
								},
							},
						},
						"network": {
							Type: schema.TypeString, Optional: true, Description: "'broadcast': Specify IS-IS broadcast multi-access network; 'point-to-point': Specify IS-IS point-to-point network;",
						},
						"password_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"password": {
										Type: schema.TypeString, Optional: true, Description: "Configure the authentication password for interface",
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Specify password for level-1 PDUs; 'level-2': Specify password for level-2 PDUs;",
									},
								},
							},
						},
						"priority_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"priority": {
										Type: schema.TypeInt, Optional: true, Default: 64, Description: "Set priority for Designated Router election (Priority value)",
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Specify priority for level-1 routing; 'level-2': Specify priority for level-2 routing;",
									},
								},
							},
						},
						"retransmit_interval": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set per-LSP retransmission interval (Interval between retransmissions of the same LSP (seconds))",
						},
						"wide_metric_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"wide_metric": {
										Type: schema.TypeInt, Optional: true, Default: 10, Description: "Configure the wide metric for interface",
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Apply metric to level-1 links; 'level-2': Apply metric to level-2 links;",
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
			"mtu": {
				Type: schema.TypeInt, Optional: true, Description: "Interface mtu (Interface MTU, default 1 (min MTU is 1280 for IPv6))",
			},
			"name": {
				Type: schema.TypeString, Optional: true, Description: "Name for the interface",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'num_pkts': num_pkts; 'num_total_bytes': num_total_bytes; 'num_unicast_pkts': num_unicast_pkts; 'num_broadcast_pkts': num_broadcast_pkts; 'num_multicast_pkts': num_multicast_pkts; 'num_tx_pkts': num_tx_pkts; 'num_total_tx_bytes': num_total_tx_bytes; 'num_unicast_tx_pkts': num_unicast_tx_pkts; 'num_broadcast_tx_pkts': num_broadcast_tx_pkts; 'num_multicast_tx_pkts': num_multicast_tx_pkts; 'dropped_dis_rx_pkts': dropped_dis_rx_pkts; 'dropped_rx_pkts': dropped_rx_pkts; 'dropped_dis_tx_pkts': dropped_dis_tx_pkts; 'dropped_tx_pkts': dropped_tx_pkts; 'dropped_rx_pkts_gre_key': dropped_rx_pkts_gre_key;",
						},
					},
				},
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
func resourceInterfaceLifCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLif(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLifUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLif(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLifDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLif(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLifRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLif(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceLifAccessList(d []interface{}) edpt.InterfaceLifAccessList {

	count1 := len(d)
	var ret edpt.InterfaceLifAccessList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AclId = in["acl_id"].(int)
		ret.AclName = in["acl_name"].(string)
	}
	return ret
}

func getObjectInterfaceLifBfd674(d []interface{}) edpt.InterfaceLifBfd674 {

	count1 := len(d)
	var ret edpt.InterfaceLifBfd674
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceLifBfdAuthentication675(in["authentication"].([]interface{}))
		ret.Echo = in["echo"].(int)
		ret.Demand = in["demand"].(int)
		ret.IntervalCfg = getObjectInterfaceLifBfdIntervalCfg676(in["interval_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifBfdAuthentication675(d []interface{}) edpt.InterfaceLifBfdAuthentication675 {

	count1 := len(d)
	var ret edpt.InterfaceLifBfdAuthentication675
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyId = in["key_id"].(int)
		ret.Method = in["method"].(string)
		ret.Password = in["password"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceLifBfdIntervalCfg676(d []interface{}) edpt.InterfaceLifBfdIntervalCfg676 {

	count1 := len(d)
	var ret edpt.InterfaceLifBfdIntervalCfg676
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interval = in["interval"].(int)
		ret.MinRx = in["min_rx"].(int)
		ret.Multiplier = in["multiplier"].(int)
	}
	return ret
}

func getObjectInterfaceLifEncapsulation677(d []interface{}) edpt.InterfaceLifEncapsulation677 {

	count1 := len(d)
	var ret edpt.InterfaceLifEncapsulation677
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dot1q = getObjectInterfaceLifEncapsulationDot1q678(in["dot1q"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLifEncapsulationDot1q678(d []interface{}) edpt.InterfaceLifEncapsulationDot1q678 {

	count1 := len(d)
	var ret edpt.InterfaceLifEncapsulationDot1q678
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(int)
		ret.Ethernet = in["ethernet"].(int)
		ret.Trunk = in["trunk"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifIp679(d []interface{}) edpt.InterfaceLifIp679 {

	count1 := len(d)
	var ret edpt.InterfaceLifIp679
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dhcp = in["dhcp"].(int)
		ret.AddressList = getSliceInterfaceLifIpAddressList680(in["address_list"].([]interface{}))
		ret.AllowPromiscuousVip = in["allow_promiscuous_vip"].(int)
		ret.CacheSpoofingPort = in["cache_spoofing_port"].(int)
		ret.Client = in["client"].(int)
		ret.Server = in["server"].(int)
		ret.Dmz = in["dmz"].(int)
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
		ret.GenerateMembershipQuery = in["generate_membership_query"].(int)
		ret.QueryInterval = in["query_interval"].(int)
		ret.MaxRespTime = in["max_resp_time"].(int)
		ret.Unnumbered = in["unnumbered"].(int)
		//omit uuid
		ret.Router = getObjectInterfaceLifIpRouter681(in["router"].([]interface{}))
		ret.Rip = getObjectInterfaceLifIpRip683(in["rip"].([]interface{}))
		ret.Ospf = getObjectInterfaceLifIpOspf691(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceLifIpAddressList680(d []interface{}) []edpt.InterfaceLifIpAddressList680 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpAddressList680, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpAddressList680
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpRouter681(d []interface{}) edpt.InterfaceLifIpRouter681 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRouter681
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isis = getObjectInterfaceLifIpRouterIsis682(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLifIpRouterIsis682(d []interface{}) edpt.InterfaceLifIpRouterIsis682 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRouterIsis682
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifIpRip683(d []interface{}) edpt.InterfaceLifIpRip683 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRip683
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceLifIpRipAuthentication684(in["authentication"].([]interface{}))
		ret.SendPacket = in["send_packet"].(int)
		ret.ReceivePacket = in["receive_packet"].(int)
		ret.SendCfg = getObjectInterfaceLifIpRipSendCfg688(in["send_cfg"].([]interface{}))
		ret.ReceiveCfg = getObjectInterfaceLifIpRipReceiveCfg689(in["receive_cfg"].([]interface{}))
		ret.SplitHorizonCfg = getObjectInterfaceLifIpRipSplitHorizonCfg690(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifIpRipAuthentication684(d []interface{}) edpt.InterfaceLifIpRipAuthentication684 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipAuthentication684
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Str = getObjectInterfaceLifIpRipAuthenticationStr685(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceLifIpRipAuthenticationMode686(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceLifIpRipAuthenticationKeyChain687(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLifIpRipAuthenticationStr685(d []interface{}) edpt.InterfaceLifIpRipAuthenticationStr685 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipAuthenticationStr685
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipAuthenticationMode686(d []interface{}) edpt.InterfaceLifIpRipAuthenticationMode686 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipAuthenticationMode686
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipAuthenticationKeyChain687(d []interface{}) edpt.InterfaceLifIpRipAuthenticationKeyChain687 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipAuthenticationKeyChain687
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipSendCfg688(d []interface{}) edpt.InterfaceLifIpRipSendCfg688 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipSendCfg688
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipReceiveCfg689(d []interface{}) edpt.InterfaceLifIpRipReceiveCfg689 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipReceiveCfg689
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipSplitHorizonCfg690(d []interface{}) edpt.InterfaceLifIpRipSplitHorizonCfg690 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipSplitHorizonCfg690
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpOspf691(d []interface{}) edpt.InterfaceLifIpOspf691 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspf691
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OspfGlobal = getObjectInterfaceLifIpOspfOspfGlobal692(in["ospf_global"].([]interface{}))
		ret.OspfIpList = getSliceInterfaceLifIpOspfOspfIpList699(in["ospf_ip_list"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobal692(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobal692 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobal692
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthenticationCfg = getObjectInterfaceLifIpOspfOspfGlobalAuthenticationCfg693(in["authentication_cfg"].([]interface{}))
		ret.AuthenticationKey = in["authentication_key"].(string)
		ret.BfdCfg = getObjectInterfaceLifIpOspfOspfGlobalBfdCfg694(in["bfd_cfg"].([]interface{}))
		ret.Cost = in["cost"].(int)
		ret.DatabaseFilterCfg = getObjectInterfaceLifIpOspfOspfGlobalDatabaseFilterCfg695(in["database_filter_cfg"].([]interface{}))
		ret.DeadInterval = in["dead_interval"].(int)
		ret.Disable = in["disable"].(string)
		ret.HelloInterval = in["hello_interval"].(int)
		ret.MessageDigestCfg = getSliceInterfaceLifIpOspfOspfGlobalMessageDigestCfg696(in["message_digest_cfg"].([]interface{}))
		ret.Mtu = in["mtu"].(int)
		ret.MtuIgnore = in["mtu_ignore"].(int)
		ret.Network = getObjectInterfaceLifIpOspfOspfGlobalNetwork698(in["network"].([]interface{}))
		ret.Priority = in["priority"].(int)
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalAuthenticationCfg693(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalAuthenticationCfg693 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobalAuthenticationCfg693
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = in["authentication"].(int)
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalBfdCfg694(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalBfdCfg694 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobalBfdCfg694
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalDatabaseFilterCfg695(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg695 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg695
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DatabaseFilter = in["database_filter"].(string)
		ret.Out = in["out"].(int)
	}
	return ret
}

func getSliceInterfaceLifIpOspfOspfGlobalMessageDigestCfg696(d []interface{}) []edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfg696 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfg696, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfg696
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = getObjectInterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5697(in["md5"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5697(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5697 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5697
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Md5Value = in["md5_value"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalNetwork698(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalNetwork698 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobalNetwork698
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

func getSliceInterfaceLifIpOspfOspfIpList699(d []interface{}) []edpt.InterfaceLifIpOspfOspfIpList699 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpOspfOspfIpList699, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpOspfOspfIpList699
		oi.IpAddr = in["ip_addr"].(string)
		oi.Authentication = in["authentication"].(int)
		oi.Value = in["value"].(string)
		oi.AuthenticationKey = in["authentication_key"].(string)
		oi.Cost = in["cost"].(int)
		oi.DatabaseFilter = in["database_filter"].(string)
		oi.Out = in["out"].(int)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.MessageDigestCfg = getSliceInterfaceLifIpOspfOspfIpListMessageDigestCfg700(in["message_digest_cfg"].([]interface{}))
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.Priority = in["priority"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpOspfOspfIpListMessageDigestCfg700(d []interface{}) []edpt.InterfaceLifIpOspfOspfIpListMessageDigestCfg700 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpOspfOspfIpListMessageDigestCfg700, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpOspfOspfIpListMessageDigestCfg700
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5Value = in["md5_value"].(string)
		//omit encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpv6701(d []interface{}) edpt.InterfaceLifIpv6701 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpv6701
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AddressList = getSliceInterfaceLifIpv6AddressList702(in["address_list"].([]interface{}))
		ret.Ipv6Enable = in["ipv6_enable"].(int)
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
		//omit uuid
		ret.Router = getObjectInterfaceLifIpv6Router703(in["router"].([]interface{}))
		ret.Ospf = getObjectInterfaceLifIpv6Ospf708(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceLifIpv6AddressList702(d []interface{}) []edpt.InterfaceLifIpv6AddressList702 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6AddressList702, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6AddressList702
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.Anycast = in["anycast"].(int)
		oi.LinkLocal = in["link_local"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpv6Router703(d []interface{}) edpt.InterfaceLifIpv6Router703 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpv6Router703
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ripng = getObjectInterfaceLifIpv6RouterRipng704(in["ripng"].([]interface{}))
		ret.Ospf = getObjectInterfaceLifIpv6RouterOspf705(in["ospf"].([]interface{}))
		ret.Isis = getObjectInterfaceLifIpv6RouterIsis707(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLifIpv6RouterRipng704(d []interface{}) edpt.InterfaceLifIpv6RouterRipng704 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpv6RouterRipng704
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifIpv6RouterOspf705(d []interface{}) edpt.InterfaceLifIpv6RouterOspf705 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpv6RouterOspf705
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AreaList = getSliceInterfaceLifIpv6RouterOspfAreaList706(in["area_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceLifIpv6RouterOspfAreaList706(d []interface{}) []edpt.InterfaceLifIpv6RouterOspfAreaList706 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6RouterOspfAreaList706, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6RouterOspfAreaList706
		oi.AreaIdNum = in["area_id_num"].(int)
		oi.AreaIdAddr = in["area_id_addr"].(string)
		oi.Tag = in["tag"].(string)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpv6RouterIsis707(d []interface{}) edpt.InterfaceLifIpv6RouterIsis707 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpv6RouterIsis707
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifIpv6Ospf708(d []interface{}) edpt.InterfaceLifIpv6Ospf708 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpv6Ospf708
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkList = getSliceInterfaceLifIpv6OspfNetworkList709(in["network_list"].([]interface{}))
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
		ret.CostCfg = getSliceInterfaceLifIpv6OspfCostCfg710(in["cost_cfg"].([]interface{}))
		ret.DeadIntervalCfg = getSliceInterfaceLifIpv6OspfDeadIntervalCfg711(in["dead_interval_cfg"].([]interface{}))
		ret.HelloIntervalCfg = getSliceInterfaceLifIpv6OspfHelloIntervalCfg712(in["hello_interval_cfg"].([]interface{}))
		ret.MtuIgnoreCfg = getSliceInterfaceLifIpv6OspfMtuIgnoreCfg713(in["mtu_ignore_cfg"].([]interface{}))
		ret.NeighborCfg = getSliceInterfaceLifIpv6OspfNeighborCfg714(in["neighbor_cfg"].([]interface{}))
		ret.PriorityCfg = getSliceInterfaceLifIpv6OspfPriorityCfg715(in["priority_cfg"].([]interface{}))
		ret.RetransmitIntervalCfg = getSliceInterfaceLifIpv6OspfRetransmitIntervalCfg716(in["retransmit_interval_cfg"].([]interface{}))
		ret.TransmitDelayCfg = getSliceInterfaceLifIpv6OspfTransmitDelayCfg717(in["transmit_delay_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfNetworkList709(d []interface{}) []edpt.InterfaceLifIpv6OspfNetworkList709 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfNetworkList709, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfNetworkList709
		oi.BroadcastType = in["broadcast_type"].(string)
		oi.P2mpNbma = in["p2mp_nbma"].(int)
		oi.NetworkInstanceId = in["network_instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfCostCfg710(d []interface{}) []edpt.InterfaceLifIpv6OspfCostCfg710 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfCostCfg710, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfCostCfg710
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfDeadIntervalCfg711(d []interface{}) []edpt.InterfaceLifIpv6OspfDeadIntervalCfg711 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfDeadIntervalCfg711, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfDeadIntervalCfg711
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfHelloIntervalCfg712(d []interface{}) []edpt.InterfaceLifIpv6OspfHelloIntervalCfg712 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfHelloIntervalCfg712, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfHelloIntervalCfg712
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfMtuIgnoreCfg713(d []interface{}) []edpt.InterfaceLifIpv6OspfMtuIgnoreCfg713 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfMtuIgnoreCfg713, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfMtuIgnoreCfg713
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfNeighborCfg714(d []interface{}) []edpt.InterfaceLifIpv6OspfNeighborCfg714 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfNeighborCfg714, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfNeighborCfg714
		oi.Neighbor = in["neighbor"].(string)
		oi.NeigInst = in["neig_inst"].(int)
		oi.NeighborCost = in["neighbor_cost"].(int)
		oi.NeighborPollInterval = in["neighbor_poll_interval"].(int)
		oi.NeighborPriority = in["neighbor_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfPriorityCfg715(d []interface{}) []edpt.InterfaceLifIpv6OspfPriorityCfg715 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfPriorityCfg715, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfPriorityCfg715
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfRetransmitIntervalCfg716(d []interface{}) []edpt.InterfaceLifIpv6OspfRetransmitIntervalCfg716 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfRetransmitIntervalCfg716, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfRetransmitIntervalCfg716
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfTransmitDelayCfg717(d []interface{}) []edpt.InterfaceLifIpv6OspfTransmitDelayCfg717 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfTransmitDelayCfg717, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfTransmitDelayCfg717
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIsis718(d []interface{}) edpt.InterfaceLifIsis718 {

	count1 := len(d)
	var ret edpt.InterfaceLifIsis718
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceLifIsisAuthentication719(in["authentication"].([]interface{}))
		ret.BfdCfg = getObjectInterfaceLifIsisBfdCfg723(in["bfd_cfg"].([]interface{}))
		ret.CircuitType = in["circuit_type"].(string)
		ret.CsnpIntervalList = getSliceInterfaceLifIsisCsnpIntervalList724(in["csnp_interval_list"].([]interface{}))
		ret.Padding = in["padding"].(int)
		ret.HelloIntervalList = getSliceInterfaceLifIsisHelloIntervalList725(in["hello_interval_list"].([]interface{}))
		ret.HelloIntervalMinimalList = getSliceInterfaceLifIsisHelloIntervalMinimalList726(in["hello_interval_minimal_list"].([]interface{}))
		ret.HelloMultiplierList = getSliceInterfaceLifIsisHelloMultiplierList727(in["hello_multiplier_list"].([]interface{}))
		ret.LspInterval = in["lsp_interval"].(int)
		ret.MeshGroup = getObjectInterfaceLifIsisMeshGroup728(in["mesh_group"].([]interface{}))
		ret.MetricList = getSliceInterfaceLifIsisMetricList729(in["metric_list"].([]interface{}))
		ret.Network = in["network"].(string)
		ret.PasswordList = getSliceInterfaceLifIsisPasswordList730(in["password_list"].([]interface{}))
		ret.PriorityList = getSliceInterfaceLifIsisPriorityList731(in["priority_list"].([]interface{}))
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.WideMetricList = getSliceInterfaceLifIsisWideMetricList732(in["wide_metric_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifIsisAuthentication719(d []interface{}) edpt.InterfaceLifIsisAuthentication719 {

	count1 := len(d)
	var ret edpt.InterfaceLifIsisAuthentication719
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SendOnlyList = getSliceInterfaceLifIsisAuthenticationSendOnlyList720(in["send_only_list"].([]interface{}))
		ret.ModeList = getSliceInterfaceLifIsisAuthenticationModeList721(in["mode_list"].([]interface{}))
		ret.KeyChainList = getSliceInterfaceLifIsisAuthenticationKeyChainList722(in["key_chain_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceLifIsisAuthenticationSendOnlyList720(d []interface{}) []edpt.InterfaceLifIsisAuthenticationSendOnlyList720 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisAuthenticationSendOnlyList720, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisAuthenticationSendOnlyList720
		oi.SendOnly = in["send_only"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisAuthenticationModeList721(d []interface{}) []edpt.InterfaceLifIsisAuthenticationModeList721 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisAuthenticationModeList721, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisAuthenticationModeList721
		oi.Mode = in["mode"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisAuthenticationKeyChainList722(d []interface{}) []edpt.InterfaceLifIsisAuthenticationKeyChainList722 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisAuthenticationKeyChainList722, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisAuthenticationKeyChainList722
		oi.KeyChain = in["key_chain"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIsisBfdCfg723(d []interface{}) edpt.InterfaceLifIsisBfdCfg723 {

	count1 := len(d)
	var ret edpt.InterfaceLifIsisBfdCfg723
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getSliceInterfaceLifIsisCsnpIntervalList724(d []interface{}) []edpt.InterfaceLifIsisCsnpIntervalList724 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisCsnpIntervalList724, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisCsnpIntervalList724
		oi.CsnpInterval = in["csnp_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisHelloIntervalList725(d []interface{}) []edpt.InterfaceLifIsisHelloIntervalList725 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisHelloIntervalList725, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisHelloIntervalList725
		oi.HelloInterval = in["hello_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisHelloIntervalMinimalList726(d []interface{}) []edpt.InterfaceLifIsisHelloIntervalMinimalList726 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisHelloIntervalMinimalList726, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisHelloIntervalMinimalList726
		oi.HelloIntervalMinimal = in["hello_interval_minimal"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisHelloMultiplierList727(d []interface{}) []edpt.InterfaceLifIsisHelloMultiplierList727 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisHelloMultiplierList727, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisHelloMultiplierList727
		oi.HelloMultiplier = in["hello_multiplier"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIsisMeshGroup728(d []interface{}) edpt.InterfaceLifIsisMeshGroup728 {

	count1 := len(d)
	var ret edpt.InterfaceLifIsisMeshGroup728
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		ret.Blocked = in["blocked"].(int)
	}
	return ret
}

func getSliceInterfaceLifIsisMetricList729(d []interface{}) []edpt.InterfaceLifIsisMetricList729 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisMetricList729, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisMetricList729
		oi.Metric = in["metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisPasswordList730(d []interface{}) []edpt.InterfaceLifIsisPasswordList730 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisPasswordList730, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisPasswordList730
		oi.Password = in["password"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisPriorityList731(d []interface{}) []edpt.InterfaceLifIsisPriorityList731 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisPriorityList731, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisPriorityList731
		oi.Priority = in["priority"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisWideMetricList732(d []interface{}) []edpt.InterfaceLifIsisWideMetricList732 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisWideMetricList732, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisWideMetricList732
		oi.WideMetric = in["wide_metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifSamplingEnable(d []interface{}) []edpt.InterfaceLifSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceLif(d *schema.ResourceData) edpt.InterfaceLif {
	var ret edpt.InterfaceLif
	ret.Inst.AccessList = getObjectInterfaceLifAccessList(d.Get("access_list").([]interface{}))
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Bfd = getObjectInterfaceLifBfd674(d.Get("bfd").([]interface{}))
	ret.Inst.Encapsulation = getObjectInterfaceLifEncapsulation677(d.Get("encapsulation").([]interface{}))
	ret.Inst.Ifname = d.Get("ifname").(string)
	ret.Inst.Ip = getObjectInterfaceLifIp679(d.Get("ip").([]interface{}))
	ret.Inst.Ipv6 = getObjectInterfaceLifIpv6701(d.Get("ipv6").([]interface{}))
	ret.Inst.Isis = getObjectInterfaceLifIsis718(d.Get("isis").([]interface{}))
	ret.Inst.Mtu = d.Get("mtu").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.SamplingEnable = getSliceInterfaceLifSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
